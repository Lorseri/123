package grpcdataserviceclient

import (
	"context"
	"fmt"
	"io"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/zilliztech/milvus-distributed/internal/logutil"

	"go.uber.org/zap"

	"google.golang.org/grpc"

	otgrpc "github.com/opentracing-contrib/go-grpc"
	msc "github.com/zilliztech/milvus-distributed/internal/distributed/masterservice/client"
	"github.com/zilliztech/milvus-distributed/internal/log"

	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"github.com/zilliztech/milvus-distributed/internal/dataservice"
	"github.com/zilliztech/milvus-distributed/internal/msgstream"
	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/datapb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	"github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"
	"github.com/zilliztech/milvus-distributed/internal/util/funcutil"
)

type Server struct {
	ctx    context.Context
	cancel context.CancelFunc

	grpcErrChan chan error
	wg          sync.WaitGroup

	impl         *dataservice.Server
	grpcServer   *grpc.Server
	masterClient *msc.GrpcClient

	closer io.Closer
}

func NewServer(ctx context.Context, factory msgstream.Factory) (*Server, error) {
	var err error
	ctx1, cancel := context.WithCancel(ctx)

	s := &Server{
		ctx:         ctx1,
		cancel:      cancel,
		grpcErrChan: make(chan error),
	}

	// TODO
	cfg := &config.Configuration{
		ServiceName: "data_service",
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
	}
	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(fmt.Sprintf("ERROR: cannot init Jaeger: %v\n", err))
	}
	opentracing.SetGlobalTracer(tracer)
	s.closer = closer

	s.impl, err = dataservice.CreateServer(s.ctx, factory)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Server) init() error {
	Params.Init()
	Params.LoadFromEnv()

	s.wg.Add(1)
	go s.startGrpcLoop(Params.Port)
	// wait for grpc server loop start
	if err := <-s.grpcErrChan; err != nil {
		return err
	}

	log.Debug("master address", zap.String("address", Params.MasterAddress))
	client, err := msc.NewClient(Params.MasterAddress, 10*time.Second)
	if err != nil {
		panic(err)
	}
	log.Debug("master client create complete")
	if err = client.Init(); err != nil {
		panic(err)
	}
	if err = client.Start(); err != nil {
		panic(err)
	}
	s.impl.UpdateStateCode(internalpb2.StateCode_INITIALIZING)

	ctx := context.Background()
	err = funcutil.WaitForComponentInitOrHealthy(ctx, client, "MasterService", 100, time.Millisecond*200)

	if err != nil {
		panic(err)
	}
	s.impl.SetMasterClient(client)

	dataservice.Params.Init()
	if err := s.impl.Init(); err != nil {
		log.Error("impl init error", zap.Error(err))
		return err
	}
	return nil
}

func (s *Server) startGrpcLoop(grpcPort int) {
	defer logutil.LogPanic()
	defer s.wg.Done()

	log.Debug("network port", zap.Int("port", grpcPort))
	lis, err := net.Listen("tcp", ":"+strconv.Itoa(grpcPort))
	if err != nil {
		log.Error("grpc server failed to listen error", zap.Error(err))
		s.grpcErrChan <- err
		return
	}

	ctx, cancel := context.WithCancel(s.ctx)
	defer cancel()

	tracer := opentracing.GlobalTracer()
	s.grpcServer = grpc.NewServer(grpc.UnaryInterceptor(
		otgrpc.OpenTracingServerInterceptor(tracer)),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(tracer)))
	datapb.RegisterDataServiceServer(s.grpcServer, s)

	go funcutil.CheckGrpcReady(ctx, s.grpcErrChan)
	if err := s.grpcServer.Serve(lis); err != nil {
		s.grpcErrChan <- err
	}
}

func (s *Server) start() error {
	return s.impl.Start()
}

func (s *Server) Stop() error {
	var err error
	if err = s.closer.Close(); err != nil {
		return err
	}
	s.cancel()

	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}

	err = s.impl.Stop()
	if err != nil {
		return err
	}

	s.wg.Wait()

	return nil
}

func (s *Server) Run() error {

	if err := s.init(); err != nil {
		return err
	}
	log.Debug("dataservice init done ...")

	if err := s.start(); err != nil {
		return err
	}
	return nil
}

func (s *Server) GetSegmentInfo(ctx context.Context, request *datapb.SegmentInfoRequest) (*datapb.SegmentInfoResponse, error) {
	return s.impl.GetSegmentInfo(ctx, request)
}

func (s *Server) RegisterNode(ctx context.Context, request *datapb.RegisterNodeRequest) (*datapb.RegisterNodeResponse, error) {
	return s.impl.RegisterNode(ctx, request)
}

func (s *Server) Flush(ctx context.Context, request *datapb.FlushRequest) (*commonpb.Status, error) {
	return s.impl.Flush(ctx, request)
}

func (s *Server) AssignSegmentID(ctx context.Context, request *datapb.AssignSegIDRequest) (*datapb.AssignSegIDResponse, error) {
	return s.impl.AssignSegmentID(ctx, request)
}

func (s *Server) ShowSegments(ctx context.Context, request *datapb.ShowSegmentRequest) (*datapb.ShowSegmentResponse, error) {
	return s.impl.ShowSegments(ctx, request)
}

func (s *Server) GetSegmentStates(ctx context.Context, request *datapb.SegmentStatesRequest) (*datapb.SegmentStatesResponse, error) {
	return s.impl.GetSegmentStates(ctx, request)
}

func (s *Server) GetInsertBinlogPaths(ctx context.Context, request *datapb.InsertBinlogPathRequest) (*datapb.InsertBinlogPathsResponse, error) {
	return s.impl.GetInsertBinlogPaths(ctx, request)
}

func (s *Server) GetInsertChannels(ctx context.Context, request *datapb.InsertChannelRequest) (*internalpb2.StringList, error) {
	return s.impl.GetInsertChannels(ctx, request)
}

func (s *Server) GetCollectionStatistics(ctx context.Context, request *datapb.CollectionStatsRequest) (*datapb.CollectionStatsResponse, error) {
	return s.impl.GetCollectionStatistics(ctx, request)
}

func (s *Server) GetPartitionStatistics(ctx context.Context, request *datapb.PartitionStatsRequest) (*datapb.PartitionStatsResponse, error) {
	return s.impl.GetPartitionStatistics(ctx, request)
}

func (s *Server) GetComponentStates(ctx context.Context, empty *commonpb.Empty) (*internalpb2.ComponentStates, error) {
	return s.impl.GetComponentStates(ctx)
}

func (s *Server) GetTimeTickChannel(ctx context.Context, empty *commonpb.Empty) (*milvuspb.StringResponse, error) {
	return s.impl.GetTimeTickChannel(ctx)
}

func (s *Server) GetStatisticsChannel(ctx context.Context, empty *commonpb.Empty) (*milvuspb.StringResponse, error) {
	return s.impl.GetStatisticsChannel(ctx)
}

func (s *Server) GetSegmentInfoChannel(ctx context.Context, empty *commonpb.Empty) (*milvuspb.StringResponse, error) {
	return s.impl.GetSegmentInfoChannel(ctx)
}

func (s *Server) GetCount(ctx context.Context, request *datapb.CollectionCountRequest) (*datapb.CollectionCountResponse, error) {
	return s.impl.GetCount(ctx, request)
}
