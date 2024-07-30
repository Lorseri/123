// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package compaction

import (
	"context"
	"testing"

	"github.com/cockroachdb/errors"
	"github.com/samber/lo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/internal/datanode/allocator"
	"github.com/milvus-io/milvus/internal/datanode/io"
	"github.com/milvus-io/milvus/internal/datanode/metacache"
	"github.com/milvus-io/milvus/internal/mocks"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
	"github.com/milvus-io/milvus/pkg/util/timerecord"
)

func TestLevelZeroCompactionTaskSuite(t *testing.T) {
	paramtable.Init()
	suite.Run(t, new(LevelZeroCompactionTaskSuite))
}

type LevelZeroCompactionTaskSuite struct {
	suite.Suite

	mockBinlogIO *io.MockBinlogIO
	mockAlloc    *allocator.MockAllocator
	task         *LevelZeroCompactionTask

	dData *storage.DeleteData
	dBlob []byte
}

func (s *LevelZeroCompactionTaskSuite) SetupTest() {
	s.mockAlloc = allocator.NewMockAllocator(s.T())
	s.mockBinlogIO = io.NewMockBinlogIO(s.T())
	// plan of the task is unset
	s.task = NewLevelZeroCompactionTask(context.Background(), s.mockBinlogIO, s.mockAlloc, nil, nil)

	pk2ts := map[int64]uint64{
		1: 20000,
		2: 20001,
		3: 20002,
	}

	dData := storage.NewEmptyDeleteData()
	for pk, ts := range pk2ts {
		dData.Append(storage.NewInt64PrimaryKey(pk), ts)
	}

	blob, err := storage.NewDeleteCodec().Serialize(0, 0, 0, dData)
	s.Require().NoError(err)
	s.dBlob = blob.GetValue()

	_, _, serializedData, err := storage.NewDeleteCodec().DeserializeWithSerialized([]*storage.Blob{{Value: s.dBlob}})
	s.Require().NoError(err)
	s.dData = serializedData
}

func (s *LevelZeroCompactionTaskSuite) TestGetMaxBatchSize() {
	tests := []struct {
		baseMem        float64
		memLimit       float64
		batchSizeLimit string

		expected    int
		description string
	}{
		{10, 100, "-1", 10, "no limitation on maxBatchSize"},
		{10, 100, "0", 10, "no limitation on maxBatchSize v2"},
		{10, 100, "11", 10, "maxBatchSize == 11"},
		{10, 100, "1", 1, "maxBatchSize == 1"},
		{10, 12, "-1", 1, "no limitation on maxBatchSize"},
		{10, 12, "100", 1, "maxBatchSize == 100"},
	}

	maxSizeK := paramtable.Get().DataNodeCfg.L0CompactionMaxBatchSize.Key
	defer paramtable.Get().Reset(maxSizeK)
	for _, test := range tests {
		s.Run(test.description, func() {
			paramtable.Get().Save(maxSizeK, test.batchSizeLimit)
			defer paramtable.Get().Reset(maxSizeK)

			actual := getMaxBatchSize(test.baseMem, test.memLimit)
			s.Equal(test.expected, actual)
		})
	}
}

func (s *LevelZeroCompactionTaskSuite) TestProcessLoadDeltaFail() {
	plan := &datapb.CompactionPlan{
		PlanID: 19530,
		Type:   datapb.CompactionType_Level0DeleteCompaction,
		SegmentBinlogs: []*datapb.CompactionSegmentBinlogs{
			{
				SegmentID: 100, Level: datapb.SegmentLevel_L0, Deltalogs: []*datapb.FieldBinlog{
					{
						Binlogs: []*datapb.Binlog{
							{LogPath: "a/b/c1", LogSize: 100},
							{LogPath: "a/b/c2", LogSize: 100},
							{LogPath: "a/b/c3", LogSize: 100},
							{LogPath: "a/b/c4", LogSize: 100},
						},
					},
				},
			},
			{SegmentID: 200, Level: datapb.SegmentLevel_L1},
		},
		Schema: &schemapb.CollectionSchema{
			Fields: []*schemapb.FieldSchema{
				{
					IsPrimaryKey: true,
				},
			},
		},
	}

	s.task.plan = plan
	s.task.tr = timerecord.NewTimeRecorder("test")
	s.mockBinlogIO.EXPECT().Download(mock.Anything, mock.Anything).Return(nil, errors.New("mock download fail")).Once()

	targetSegments := lo.Filter(plan.SegmentBinlogs, func(s *datapb.CompactionSegmentBinlogs, _ int) bool {
		return s.Level == datapb.SegmentLevel_L1
	})
	deltaLogs := map[int64][]string{100: {"a/b/c1"}}

	segments, err := s.task.process(context.Background(), 1, targetSegments, lo.Values(deltaLogs)...)
	s.Error(err)
	s.Empty(segments)
}

func (s *LevelZeroCompactionTaskSuite) TestProcessUploadFail() {
	plan := &datapb.CompactionPlan{
		PlanID: 19530,
		Type:   datapb.CompactionType_Level0DeleteCompaction,
		SegmentBinlogs: []*datapb.CompactionSegmentBinlogs{
			{
				SegmentID: 100, Level: datapb.SegmentLevel_L0, Deltalogs: []*datapb.FieldBinlog{
					{
						Binlogs: []*datapb.Binlog{
							{LogPath: "a/b/c1", LogSize: 100},
							{LogPath: "a/b/c2", LogSize: 100},
							{LogPath: "a/b/c3", LogSize: 100},
							{LogPath: "a/b/c4", LogSize: 100},
						},
					},
				},
			},
			{SegmentID: 200, Level: datapb.SegmentLevel_L1, Field2StatslogPaths: []*datapb.FieldBinlog{
				{
					Binlogs: []*datapb.Binlog{
						{LogID: 9999, LogSize: 100},
					},
				},
			}},
		},
		Schema: &schemapb.CollectionSchema{
			Fields: []*schemapb.FieldSchema{
				{
					IsPrimaryKey: true,
				},
			},
		},
	}
	s.task.plan = plan

	data := &storage.Int64FieldData{
		Data: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	sw := &storage.StatsWriter{}
	err := sw.GenerateByData(common.RowIDField, schemapb.DataType_Int64, data)
	s.NoError(err)
	cm := mocks.NewChunkManager(s.T())
	cm.EXPECT().MultiRead(mock.Anything, mock.Anything).Return([][]byte{sw.GetBuffer()}, nil)
	s.task.cm = cm

	s.mockBinlogIO.EXPECT().Download(mock.Anything, mock.Anything).Return([][]byte{s.dBlob}, nil).Once()
	mockAlloc := allocator.NewMockAllocator(s.T())
	mockAlloc.EXPECT().AllocOne().Return(0, errors.New("mock alloc err"))
	s.task.allocator = mockAlloc

	targetSegments := lo.Filter(plan.SegmentBinlogs, func(s *datapb.CompactionSegmentBinlogs, _ int) bool {
		return s.Level == datapb.SegmentLevel_L1
	})
	deltaLogs := map[int64][]string{100: {"a/b/c1"}}

	segments, err := s.task.process(context.Background(), 2, targetSegments, lo.Values(deltaLogs)...)
	s.Error(err)
	s.Empty(segments)
}

func (s *LevelZeroCompactionTaskSuite) TestCompactLinear() {
	plan := &datapb.CompactionPlan{
		PlanID: 19530,
		Type:   datapb.CompactionType_Level0DeleteCompaction,
		SegmentBinlogs: []*datapb.CompactionSegmentBinlogs{
			{
				CollectionID: 1,
				SegmentID:    100, Level: datapb.SegmentLevel_L0, Deltalogs: []*datapb.FieldBinlog{
					{
						Binlogs: []*datapb.Binlog{
							{LogPath: "a/b/c1", LogSize: 100},
							{LogPath: "a/b/c2", LogSize: 100},
							{LogPath: "a/b/c3", LogSize: 100},
							{LogPath: "a/b/c4", LogSize: 100},
						},
					},
				},
			},
			{
				CollectionID: 1,
				SegmentID:    101, Level: datapb.SegmentLevel_L0, Deltalogs: []*datapb.FieldBinlog{
					{
						Binlogs: []*datapb.Binlog{
							{LogPath: "a/d/c1", LogSize: 100},
							{LogPath: "a/d/c2", LogSize: 100},
							{LogPath: "a/d/c3", LogSize: 100},
							{LogPath: "a/d/c4", LogSize: 100},
						},
					},
				},
			},
			{
				CollectionID: 1,
				SegmentID:    200, Level: datapb.SegmentLevel_L1, Field2StatslogPaths: []*datapb.FieldBinlog{
					{
						Binlogs: []*datapb.Binlog{
							{LogID: 9999, LogSize: 100},
						},
					},
				},
			},
			{
				CollectionID: 1,
				SegmentID:    201, Level: datapb.SegmentLevel_L1, Field2StatslogPaths: []*datapb.FieldBinlog{
					{
						Binlogs: []*datapb.Binlog{
							{LogID: 9999, LogSize: 100},
						},
					},
				},
			},
		},
		Schema: &schemapb.CollectionSchema{
			Fields: []*schemapb.FieldSchema{
				{
					IsPrimaryKey: true,
				},
			},
		},
	}

	s.task.plan = plan
	s.task.tr = timerecord.NewTimeRecorder("test")

	data := &storage.Int64FieldData{
		Data: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	sw := &storage.StatsWriter{}
	err := sw.GenerateByData(common.RowIDField, schemapb.DataType_Int64, data)
	s.NoError(err)
	cm := mocks.NewChunkManager(s.T())
	cm.EXPECT().MultiRead(mock.Anything, mock.Anything).Return([][]byte{sw.GetBuffer()}, nil)
	s.task.cm = cm

	s.mockBinlogIO.EXPECT().Download(mock.Anything, mock.Anything).Return([][]byte{s.dBlob}, nil).Times(1)
	s.mockBinlogIO.EXPECT().Upload(mock.Anything, mock.Anything).Return(nil).Once()
	s.mockAlloc.EXPECT().AllocOne().Return(19530, nil).Times(2)

	s.Require().Equal(plan.GetPlanID(), s.task.GetPlanID())
	s.Require().Equal(plan.GetChannel(), s.task.GetChannelName())
	s.Require().EqualValues(1, s.task.GetCollection())

	l0Segments := lo.Filter(s.task.plan.GetSegmentBinlogs(), func(s *datapb.CompactionSegmentBinlogs, _ int) bool {
		return s.Level == datapb.SegmentLevel_L0
	})

	targetSegments := lo.Filter(s.task.plan.GetSegmentBinlogs(), func(s *datapb.CompactionSegmentBinlogs, _ int) bool {
		return s.Level == datapb.SegmentLevel_L1
	})
	totalDeltalogs := make(map[int64][]string)

	for _, s := range l0Segments {
		paths := []string{}
		for _, d := range s.GetDeltalogs() {
			for _, l := range d.GetBinlogs() {
				paths = append(paths, l.GetLogPath())
			}
		}
		if len(paths) > 0 {
			totalDeltalogs[s.GetSegmentID()] = paths
		}
	}
	segments, err := s.task.process(context.Background(), 1, targetSegments, lo.Values(totalDeltalogs)...)
	s.NoError(err)
	s.NotEmpty(segments)
	s.Equal(2, len(segments))
	s.ElementsMatch([]int64{200, 201},
		lo.Map(segments, func(seg *datapb.CompactionSegment, _ int) int64 {
			return seg.GetSegmentID()
		}))
	for _, segment := range segments {
		s.NotNil(segment.GetDeltalogs())
	}

	log.Info("test segment results", zap.Any("result", segments))
}

func (s *LevelZeroCompactionTaskSuite) TestCompactBatch() {
	plan := &datapb.CompactionPlan{
		PlanID: 19530,
		Type:   datapb.CompactionType_Level0DeleteCompaction,
		SegmentBinlogs: []*datapb.CompactionSegmentBinlogs{
			{
				SegmentID: 100, Level: datapb.SegmentLevel_L0, Deltalogs: []*datapb.FieldBinlog{
					{
						Binlogs: []*datapb.Binlog{
							{LogPath: "a/b/c1", LogSize: 100},
							{LogPath: "a/b/c2", LogSize: 100},
							{LogPath: "a/b/c3", LogSize: 100},
							{LogPath: "a/b/c4", LogSize: 100},
						},
					},
				},
			},
			{
				SegmentID: 101, Level: datapb.SegmentLevel_L0, Deltalogs: []*datapb.FieldBinlog{
					{
						Binlogs: []*datapb.Binlog{
							{LogPath: "a/d/c1", LogSize: 100},
							{LogPath: "a/d/c2", LogSize: 100},
							{LogPath: "a/d/c3", LogSize: 100},
							{LogPath: "a/d/c4", LogSize: 100},
						},
					},
				},
			},
			{SegmentID: 200, Level: datapb.SegmentLevel_L1, Field2StatslogPaths: []*datapb.FieldBinlog{
				{
					Binlogs: []*datapb.Binlog{
						{LogID: 9999, LogSize: 100},
					},
				},
			}},
			{SegmentID: 201, Level: datapb.SegmentLevel_L1, Field2StatslogPaths: []*datapb.FieldBinlog{
				{
					Binlogs: []*datapb.Binlog{
						{LogID: 9999, LogSize: 100},
					},
				},
			}},
		},
		Schema: &schemapb.CollectionSchema{
			Fields: []*schemapb.FieldSchema{
				{
					IsPrimaryKey: true,
				},
			},
		},
	}

	s.task.plan = plan
	s.task.tr = timerecord.NewTimeRecorder("test")

	data := &storage.Int64FieldData{
		Data: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	sw := &storage.StatsWriter{}
	err := sw.GenerateByData(common.RowIDField, schemapb.DataType_Int64, data)
	s.NoError(err)
	cm := mocks.NewChunkManager(s.T())
	cm.EXPECT().MultiRead(mock.Anything, mock.Anything).Return([][]byte{sw.GetBuffer()}, nil)
	s.task.cm = cm

	s.mockBinlogIO.EXPECT().Download(mock.Anything, mock.Anything).Return([][]byte{s.dBlob}, nil).Once()
	s.mockAlloc.EXPECT().AllocOne().Return(19530, nil).Times(2)
	s.mockBinlogIO.EXPECT().Upload(mock.Anything, mock.Anything).Return(nil).Once()

	l0Segments := lo.Filter(s.task.plan.GetSegmentBinlogs(), func(s *datapb.CompactionSegmentBinlogs, _ int) bool {
		return s.Level == datapb.SegmentLevel_L0
	})

	targetSegments := lo.Filter(s.task.plan.GetSegmentBinlogs(), func(s *datapb.CompactionSegmentBinlogs, _ int) bool {
		return s.Level == datapb.SegmentLevel_L1
	})

	totalDeltalogs := make(map[int64][]string)
	for _, s := range l0Segments {
		paths := []string{}
		for _, d := range s.GetDeltalogs() {
			for _, l := range d.GetBinlogs() {
				paths = append(paths, l.GetLogPath())
			}
		}
		if len(paths) > 0 {
			totalDeltalogs[s.GetSegmentID()] = paths
		}
	}
	segments, err := s.task.process(context.TODO(), 2, targetSegments, lo.Values(totalDeltalogs)...)
	s.NoError(err)
	s.NotEmpty(segments)
	s.Equal(2, len(segments))
	s.ElementsMatch([]int64{200, 201},
		lo.Map(segments, func(seg *datapb.CompactionSegment, _ int) int64 {
			return seg.GetSegmentID()
		}))
	for _, segment := range segments {
		s.NotNil(segment.GetDeltalogs())
	}

	log.Info("test segment results", zap.Any("result", segments))
}

func (s *LevelZeroCompactionTaskSuite) TestSerializeUpload() {
	ctx := context.Background()
	plan := &datapb.CompactionPlan{
		SegmentBinlogs: []*datapb.CompactionSegmentBinlogs{
			{
				SegmentID: 100,
			},
			{
				SegmentID: 101,
			},
		},
	}

	s.Run("serializeUpload allocator Alloc failed", func() {
		s.SetupTest()
		s.mockAlloc.EXPECT().AllocOne().Return(0, errors.New("mock alloc wrong"))
		writer := NewSegmentDeltaWriter(100, 10, 1)

		for i := range s.dData.Pks {
			writer.WriteSerialized(s.dData.Serialized[i], s.dData.Pks[i], s.dData.Tss[i])
		}
		writers := map[int64]*SegmentDeltaWriter{100: writer}

		result, err := s.task.serializeUpload(ctx, writers)
		s.Error(err)
		s.Equal(0, len(result))
	})

	s.Run("serializeUpload Upload failed", func() {
		s.SetupTest()
		s.task.plan = plan
		s.mockBinlogIO.EXPECT().Upload(mock.Anything, mock.Anything).Return(errors.New("mock upload failed"))
		s.mockAlloc.EXPECT().AllocOne().Return(19530, nil)

		writer := NewSegmentDeltaWriter(100, 10, 1)
		for i := range s.dData.Pks {
			writer.WriteSerialized(s.dData.Serialized[i], s.dData.Pks[i], s.dData.Tss[i])
		}
		writers := map[int64]*SegmentDeltaWriter{100: writer}
		results, err := s.task.serializeUpload(ctx, writers)
		s.Error(err)
		s.Equal(0, len(results))
	})

	s.Run("upload success", func() {
		s.SetupTest()
		s.task.plan = plan
		s.mockBinlogIO.EXPECT().Upload(mock.Anything, mock.Anything).Return(nil)
		s.mockAlloc.EXPECT().AllocOne().Return(19530, nil)
		writer := NewSegmentDeltaWriter(100, 10, 1)
		for i := range s.dData.Pks {
			writer.WriteSerialized(s.dData.Serialized[i], s.dData.Pks[i], s.dData.Tss[i])
		}
		writers := map[int64]*SegmentDeltaWriter{100: writer}
		results, err := s.task.serializeUpload(ctx, writers)
		s.NoError(err)
		s.Equal(1, len(results))

		seg1 := results[0]
		s.EqualValues(100, seg1.GetSegmentID())
		s.Equal(1, len(seg1.GetDeltalogs()))
		s.Equal(1, len(seg1.GetDeltalogs()[0].GetBinlogs()))
	})
}

func (s *LevelZeroCompactionTaskSuite) TestSplitDelta() {
	plan := &datapb.CompactionPlan{
		SegmentBinlogs: []*datapb.CompactionSegmentBinlogs{
			{
				CollectionID: 1,
			},
		},
	}
	s.task.plan = plan

	bfs1 := metacache.NewBloomFilterSetWithBatchSize(100)
	bfs1.UpdatePKRange(&storage.Int64FieldData{Data: []int64{1, 3}})
	bfs2 := metacache.NewBloomFilterSetWithBatchSize(100)
	bfs2.UpdatePKRange(&storage.Int64FieldData{Data: []int64{3}})
	bfs3 := metacache.NewBloomFilterSetWithBatchSize(100)
	bfs3.UpdatePKRange(&storage.Int64FieldData{Data: []int64{3}})

	predicted := []int64{100, 101, 102}

	expectedSegPK := map[int64][]int64{
		100: {1, 3},
		101: {3},
		102: {3},
	}
	segmentBfs := map[int64]*metacache.BloomFilterSet{
		100: bfs1,
		101: bfs2,
		102: bfs3,
	}
	deltaWriters, err := s.task.splitDelta(context.TODO(), s.dData, segmentBfs)
	s.NoError(err)
	s.NotEmpty(deltaWriters)
	s.ElementsMatch(predicted, lo.Keys(deltaWriters))
	s.EqualValues(2, deltaWriters[100].GetRowNum())
	s.EqualValues(1, deltaWriters[101].GetRowNum())
	s.EqualValues(1, deltaWriters[102].GetRowNum())

	for segID, writer := range deltaWriters {
		gotBytes, _, err := writer.Finish()
		s.NoError(err)

		_, _, gotData, err := storage.NewDeleteCodec().Deserialize([]*storage.Blob{{Value: gotBytes}})
		s.NoError(err)
		s.ElementsMatch(expectedSegPK[segID], lo.Map(gotData.Pks, func(pk storage.PrimaryKey, _ int) int64 { return pk.(*storage.Int64PrimaryKey).Value }))
	}
}

func (s *LevelZeroCompactionTaskSuite) TestLoadDelta() {
	ctx := context.TODO()

	s.mockBinlogIO.EXPECT().Download(mock.Anything, mock.MatchedBy(
		func(paths []string) bool {
			return len(paths) > 0 && paths[0] == "correct"
		})).Return([][]byte{s.dBlob}, nil).Once()

	s.mockBinlogIO.EXPECT().Download(mock.Anything, mock.MatchedBy(
		func(paths []string) bool {
			return len(paths) > 0 && paths[0] == "error"
		})).Return(nil, errors.New("mock err")).Once()

	s.mockBinlogIO.EXPECT().Download(mock.Anything, mock.MatchedBy(
		func(paths []string) bool {
			return len(paths) > 0 && paths[0] == "invalid-blobs"
		})).Return([][]byte{{1}}, nil).Once()

	tests := []struct {
		description string
		paths       []string

		expectError bool
	}{
		{"no error", []string{"correct"}, false},
		{"download error", []string{"error"}, true},
		{"deserialize error", []string{"invalid-blobs"}, true},
	}

	for _, test := range tests {
		dData, err := s.task.loadDelta(ctx, test.paths)

		if test.expectError {
			s.Error(err)
		} else {
			s.NoError(err)
			s.NotEmpty(dData)
			s.NotNil(dData)
			s.ElementsMatch(s.dData.Pks, dData.Pks)
			s.Equal(s.dData.RowCount, dData.RowCount)
		}
	}
}

func (s *LevelZeroCompactionTaskSuite) TestLoadBF() {
	plan := &datapb.CompactionPlan{
		PlanID: 19530,
		Type:   datapb.CompactionType_Level0DeleteCompaction,
		SegmentBinlogs: []*datapb.CompactionSegmentBinlogs{
			{SegmentID: 201, Level: datapb.SegmentLevel_L1, Field2StatslogPaths: []*datapb.FieldBinlog{
				{
					Binlogs: []*datapb.Binlog{
						{LogID: 9999, LogSize: 100},
					},
				},
			}},
		},
		Schema: &schemapb.CollectionSchema{
			Fields: []*schemapb.FieldSchema{
				{
					IsPrimaryKey: true,
				},
			},
		},
	}

	s.task.plan = plan

	data := &storage.Int64FieldData{
		Data: []int64{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	sw := &storage.StatsWriter{}
	err := sw.GenerateByData(common.RowIDField, schemapb.DataType_Int64, data)
	s.NoError(err)
	cm := mocks.NewChunkManager(s.T())
	cm.EXPECT().MultiRead(mock.Anything, mock.Anything).Return([][]byte{sw.GetBuffer()}, nil)
	s.task.cm = cm

	bfs, err := s.task.loadBF(plan.SegmentBinlogs)
	s.NoError(err)

	s.Len(bfs, 1)
	for _, pk := range s.dData.Pks {
		lc := storage.NewLocationsCache(pk)
		s.True(bfs[201].PkExists(lc))
	}
}

func (s *LevelZeroCompactionTaskSuite) TestFailed() {
	s.Run("no primary key", func() {
		plan := &datapb.CompactionPlan{
			PlanID: 19530,
			Type:   datapb.CompactionType_Level0DeleteCompaction,
			SegmentBinlogs: []*datapb.CompactionSegmentBinlogs{
				{SegmentID: 201, Level: datapb.SegmentLevel_L1, Field2StatslogPaths: []*datapb.FieldBinlog{
					{
						Binlogs: []*datapb.Binlog{
							{LogID: 9999, LogSize: 100},
						},
					},
				}},
			},
			Schema: &schemapb.CollectionSchema{
				Fields: []*schemapb.FieldSchema{
					{
						IsPrimaryKey: false,
					},
				},
			},
		}

		s.task.plan = plan

		_, err := s.task.loadBF(plan.SegmentBinlogs)
		s.Error(err)
	})

	s.Run("no l1 segments", func() {
		plan := &datapb.CompactionPlan{
			PlanID: 19530,
			Type:   datapb.CompactionType_Level0DeleteCompaction,
			SegmentBinlogs: []*datapb.CompactionSegmentBinlogs{
				{SegmentID: 201, Level: datapb.SegmentLevel_L0},
			},
		}

		s.task.plan = plan

		_, err := s.task.Compact()
		s.Error(err)
	})
}
