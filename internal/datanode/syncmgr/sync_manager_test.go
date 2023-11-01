package syncmgr

import (
	"context"
	"math/rand"
	"sync/atomic"
	"testing"
	"time"

	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/msgpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/internal/allocator"
	"github.com/milvus-io/milvus/internal/datanode/broker"
	"github.com/milvus-io/milvus/internal/datanode/metacache"
	"github.com/milvus-io/milvus/internal/mocks"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
	"github.com/milvus-io/milvus/pkg/util/tsoutil"
)

type SyncManagerSuite struct {
	suite.Suite

	collectionID int64
	partitionID  int64
	segmentID    int64
	channelName  string

	metacache    *metacache.MockMetaCache
	allocator    *allocator.MockGIDAllocator
	schema       *schemapb.CollectionSchema
	chunkManager *mocks.ChunkManager
	broker       *broker.MockBroker
}

func (s *SyncManagerSuite) SetupSuite() {
	paramtable.Get().Init(paramtable.NewBaseTable())

	s.collectionID = 100
	s.partitionID = 101
	s.segmentID = 1001
	s.channelName = "by-dev-rootcoord-dml_0_100v0"

	s.schema = &schemapb.CollectionSchema{
		Name: "sync_task_test_col",
		Fields: []*schemapb.FieldSchema{
			{FieldID: common.RowIDField, DataType: schemapb.DataType_Int64},
			{FieldID: common.TimeStampField, DataType: schemapb.DataType_Int64},
			{
				FieldID:      100,
				Name:         "pk",
				DataType:     schemapb.DataType_Int64,
				IsPrimaryKey: true,
			},
			{
				FieldID:  101,
				Name:     "vector",
				DataType: schemapb.DataType_FloatVector,
				TypeParams: []*commonpb.KeyValuePair{
					{Key: common.DimKey, Value: "128"},
				},
			},
		},
	}
}

func (s *SyncManagerSuite) SetupTest() {
	s.allocator = allocator.NewMockGIDAllocator()
	s.allocator.AllocF = func(count uint32) (int64, int64, error) {
		return time.Now().Unix(), int64(count), nil
	}
	s.allocator.AllocOneF = func() (allocator.UniqueID, error) {
		return time.Now().Unix(), nil
	}

	s.chunkManager = mocks.NewChunkManager(s.T())
	s.chunkManager.EXPECT().RootPath().Return("files").Maybe()
	s.chunkManager.EXPECT().MultiWrite(mock.Anything, mock.Anything).Return(nil).Maybe()

	s.broker = broker.NewMockBroker(s.T())
	s.metacache = metacache.NewMockMetaCache(s.T())
}

func (s *SyncManagerSuite) getEmptyInsertBuffer() *storage.InsertData {
	buf, err := storage.NewInsertData(s.schema)
	s.Require().NoError(err)

	return buf
}

func (s *SyncManagerSuite) getInsertBuffer() *storage.InsertData {
	buf := s.getEmptyInsertBuffer()

	// generate data
	for i := 0; i < 10; i++ {
		data := make(map[storage.FieldID]any)
		data[common.RowIDField] = int64(i + 1)
		data[common.TimeStampField] = int64(i + 1)
		data[100] = int64(i + 1)
		vector := lo.RepeatBy(128, func(_ int) float32 {
			return rand.Float32()
		})
		data[101] = vector
		err := buf.Append(data)
		s.Require().NoError(err)
	}
	return buf
}

func (s *SyncManagerSuite) getDeleteBuffer() *storage.DeleteData {
	buf := &storage.DeleteData{}
	for i := 0; i < 10; i++ {
		pk := storage.NewInt64PrimaryKey(int64(i + 1))
		ts := tsoutil.ComposeTSByTime(time.Now(), 0)
		buf.Append(pk, ts)
	}
	return buf
}

func (s *SyncManagerSuite) getDeleteBufferZeroTs() *storage.DeleteData {
	buf := &storage.DeleteData{}
	for i := 0; i < 10; i++ {
		pk := storage.NewInt64PrimaryKey(int64(i + 1))
		buf.Append(pk, 0)
	}
	return buf
}

func (s *SyncManagerSuite) getSuiteSyncTask() *SyncTask {
	task := NewSyncTask().WithCollectionID(s.collectionID).
		WithPartitionID(s.partitionID).
		WithSegmentID(s.segmentID).
		WithChannelName(s.channelName).
		WithSchema(s.schema).
		WithChunkManager(s.chunkManager).
		WithAllocator(s.allocator).
		WithMetaCache(s.metacache)

	return task
}

func (s *SyncManagerSuite) TestSubmit() {
	sig := make(chan struct{})
	s.broker.EXPECT().SaveBinlogPaths(mock.Anything, mock.Anything).Run(func(_ context.Context, _ *datapb.SaveBinlogPathsRequest) {
		close(sig)
	}).Return(nil)
	bfs := metacache.NewBloomFilterSet()
	seg := metacache.NewSegmentInfo(&datapb.SegmentInfo{}, bfs)
	metacache.UpdateNumOfRows(1000)(seg)
	s.metacache.EXPECT().GetSegmentsBy(mock.Anything).Return([]*metacache.SegmentInfo{seg})

	manager, err := NewSyncManager(10, s.chunkManager, s.allocator)
	s.NoError(err)
	task := s.getSuiteSyncTask()
	task.WithMetaWriter(BrokerMetaWriter(s.broker))
	task.WithTimeRange(50, 100)
	task.WithCheckpoint(&msgpb.MsgPosition{
		ChannelName: s.channelName,
		MsgID:       []byte{1, 2, 3, 4},
		Timestamp:   100,
	})

	err = manager.SyncData(context.Background(), task)
	s.NoError(err)

	<-sig
}

func (s *SyncManagerSuite) TestCompacted() {
	sig := make(chan struct{})
	var segmentID atomic.Int64
	s.broker.EXPECT().SaveBinlogPaths(mock.Anything, mock.Anything).Run(func(_ context.Context, req *datapb.SaveBinlogPathsRequest) {
		close(sig)
		segmentID.Store(req.GetSegmentID())
	}).Return(nil)
	bfs := metacache.NewBloomFilterSet()
	seg := metacache.NewSegmentInfo(&datapb.SegmentInfo{}, bfs)
	metacache.UpdateNumOfRows(1000)(seg)
	metacache.CompactTo(1001)(seg)
	s.metacache.EXPECT().GetSegmentsBy(mock.Anything).Return([]*metacache.SegmentInfo{seg})

	manager, err := NewSyncManager(10, s.chunkManager, s.allocator)
	s.NoError(err)
	task := s.getSuiteSyncTask()
	task.WithMetaWriter(BrokerMetaWriter(s.broker))
	task.WithTimeRange(50, 100)
	task.WithCheckpoint(&msgpb.MsgPosition{
		ChannelName: s.channelName,
		MsgID:       []byte{1, 2, 3, 4},
		Timestamp:   100,
	})

	err = manager.SyncData(context.Background(), task)
	s.NoError(err)

	<-sig
	s.EqualValues(1001, segmentID.Load())
}

func (s *SyncManagerSuite) TestBlock() {
	sig := make(chan struct{})
	s.broker.EXPECT().SaveBinlogPaths(mock.Anything, mock.Anything).Run(func(_ context.Context, _ *datapb.SaveBinlogPathsRequest) {
		close(sig)
	}).Return(nil)
	bfs := metacache.NewBloomFilterSet()
	seg := metacache.NewSegmentInfo(&datapb.SegmentInfo{}, bfs)
	metacache.UpdateNumOfRows(1000)(seg)
	s.metacache.EXPECT().GetSegmentsBy(mock.Anything).Return([]*metacache.SegmentInfo{seg})

	manager, err := NewSyncManager(10, s.chunkManager, s.allocator)
	s.NoError(err)

	// block
	manager.Block(s.segmentID)

	go func() {
		task := s.getSuiteSyncTask()
		task.WithMetaWriter(BrokerMetaWriter(s.broker))
		task.WithTimeRange(50, 100)
		task.WithCheckpoint(&msgpb.MsgPosition{
			ChannelName: s.channelName,
			MsgID:       []byte{1, 2, 3, 4},
			Timestamp:   100,
		})
		manager.SyncData(context.Background(), task)
	}()

	select {
	case <-sig:
		s.FailNow("sync task done during block")
	default:
	}

	manager.Unblock(s.segmentID)
	<-sig
}

func TestSyncManager(t *testing.T) {
	suite.Run(t, new(SyncManagerSuite))
}
