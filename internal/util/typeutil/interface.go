package typeutil

import (
	"context"

	"github.com/zilliztech/milvus-distributed/internal/proto/commonpb"
	"github.com/zilliztech/milvus-distributed/internal/proto/internalpb2"
	"github.com/zilliztech/milvus-distributed/internal/proto/milvuspb"
	"github.com/zilliztech/milvus-distributed/internal/proto/querypb"
)

type TimeTickHandler interface {
	GetTimeTickChannel(ctx context.Context) (*milvuspb.StringResponse, error)
}
type Service interface {
	Init() error
	Start() error
	Stop() error
}

type Component interface {
	GetComponentStates(ctx context.Context) (*internalpb2.ComponentStates, error)
	GetStatisticsChannel(ctx context.Context) (*milvuspb.StringResponse, error)
}

type QueryServiceInterface interface {
	Service
	Component

	RegisterNode(ctx context.Context, req *querypb.RegisterNodeRequest) (*querypb.RegisterNodeResponse, error)
	ShowCollections(ctx context.Context, req *querypb.ShowCollectionRequest) (*querypb.ShowCollectionResponse, error)
	LoadCollection(ctx context.Context, req *querypb.LoadCollectionRequest) (*commonpb.Status, error)
	ReleaseCollection(ctx context.Context, req *querypb.ReleaseCollectionRequest) (*commonpb.Status, error)
	ShowPartitions(ctx context.Context, req *querypb.ShowPartitionRequest) (*querypb.ShowPartitionResponse, error)
	LoadPartitions(ctx context.Context, req *querypb.LoadPartitionRequest) (*commonpb.Status, error)
	ReleasePartitions(ctx context.Context, req *querypb.ReleasePartitionRequest) (*commonpb.Status, error)
	CreateQueryChannel(ctx context.Context) (*querypb.CreateQueryChannelResponse, error)
	GetPartitionStates(ctx context.Context, req *querypb.PartitionStatesRequest) (*querypb.PartitionStatesResponse, error)
	GetSegmentInfo(ctx context.Context, req *querypb.SegmentInfoRequest) (*querypb.SegmentInfoResponse, error)
}
