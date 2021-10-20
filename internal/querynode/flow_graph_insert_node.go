// Copyright (C) 2019-2020 Zilliz. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software distributed under the License
// is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
// or implied. See the License for the specific language governing permissions and limitations under the License.

package querynode

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"strconv"
	"sync"

	"github.com/opentracing/opentracing-go"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/log"
	"github.com/milvus-io/milvus/internal/msgstream"
	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/proto/schemapb"
	"github.com/milvus-io/milvus/internal/util/flowgraph"
	"github.com/milvus-io/milvus/internal/util/trace"
)

type insertNode struct {
	baseNode
	replica ReplicaInterface
}

type insertData struct {
	insertIDs        map[UniqueID][]int64
	insertTimestamps map[UniqueID][]Timestamp
	insertRecords    map[UniqueID][]*commonpb.Blob
	insertOffset     map[UniqueID]int64
	insertPKs        map[UniqueID][]int64
}

type deleteData struct {
	deleteIDs        map[UniqueID][]int64
	deleteTimestamps map[UniqueID][]Timestamp
	deleteOffset     map[UniqueID]int64
}

func (iNode *insertNode) Name() string {
	return "iNode"
}

func (iNode *insertNode) Operate(in []flowgraph.Msg) []flowgraph.Msg {
	//log.Debug("Do insertNode operation")

	if len(in) != 1 {
		log.Error("Invalid operate message input in insertNode", zap.Int("input length", len(in)))
		// TODO: add error handling
	}

	iMsg, ok := in[0].(*insertMsg)
	if !ok {
		log.Warn("type assertion failed for insertMsg")
		// TODO: add error handling
	}

	iData := insertData{
		insertIDs:        make(map[UniqueID][]int64),
		insertTimestamps: make(map[UniqueID][]Timestamp),
		insertRecords:    make(map[UniqueID][]*commonpb.Blob),
		insertOffset:     make(map[UniqueID]int64),
		insertPKs:        make(map[UniqueID][]int64),
	}

	if iMsg == nil {
		return []Msg{}
	}

	var spans []opentracing.Span
	for _, msg := range iMsg.insertMessages {
		sp, ctx := trace.StartSpanFromContext(msg.TraceCtx())
		spans = append(spans, sp)
		msg.SetTraceCtx(ctx)
	}

	// 1. hash insertMessages to insertData
	for _, task := range iMsg.insertMessages {
		// check if partition exists, if not, create partition
		if hasPartition := iNode.replica.hasPartition(task.PartitionID); !hasPartition {
			err := iNode.replica.addPartition(task.CollectionID, task.PartitionID)
			if err != nil {
				log.Warn(err.Error())
				continue
			}
		}

		// check if segment exists, if not, create this segment
		if !iNode.replica.hasSegment(task.SegmentID) {
			err := iNode.replica.addSegment(task.SegmentID, task.PartitionID, task.CollectionID, task.ShardName, segmentTypeGrowing, true)
			if err != nil {
				log.Warn(err.Error())
				continue
			}
		}

		iData.insertIDs[task.SegmentID] = append(iData.insertIDs[task.SegmentID], task.RowIDs...)
		iData.insertTimestamps[task.SegmentID] = append(iData.insertTimestamps[task.SegmentID], task.Timestamps...)
		iData.insertRecords[task.SegmentID] = append(iData.insertRecords[task.SegmentID], task.RowData...)
		iData.insertPKs[task.SegmentID] = iNode.getPrimaryKeys(task)
	}

	// 2. do preInsert
	for segmentID := range iData.insertRecords {
		var targetSegment, err = iNode.replica.getSegmentByID(segmentID)
		if err != nil {
			log.Warn(err.Error())
		}

		var numOfRecords = len(iData.insertRecords[segmentID])
		if targetSegment != nil {
			offset, err := targetSegment.segmentPreInsert(numOfRecords)
			if err != nil {
				log.Warn(err.Error())
			}
			iData.insertOffset[segmentID] = offset
			log.Debug("insertNode operator", zap.Int("insert size", numOfRecords), zap.Int64("insert offset", offset), zap.Int64("segment id", segmentID))
			targetSegment.updateBloomFilter(iData.insertPKs[segmentID])
		}
	}

	// 3. do insert
	wg := sync.WaitGroup{}
	for segmentID := range iData.insertRecords {
		wg.Add(1)
		go iNode.insert(&iData, segmentID, &wg)
	}
	wg.Wait()

	delData := &deleteData{
		deleteIDs:        make(map[UniqueID][]int64),
		deleteTimestamps: make(map[UniqueID][]Timestamp),
		deleteOffset:     make(map[UniqueID]int64),
	}
	// 1. filter segment by bloom filter
	for _, delMsg := range iMsg.deleteMessages {
		var partitionIDs []UniqueID
		var err error
		if delMsg.PartitionID != -1 {
			partitionIDs = []UniqueID{delMsg.PartitionID}
		} else {
			partitionIDs, err = iNode.replica.getPartitionIDs(delMsg.CollectionID)
			if err != nil {
				log.Warn(err.Error())
				continue
			}
		}
		resultSegmentIDs := make([]UniqueID, 0)
		for _, partitionID := range partitionIDs {
			segmentIDs, err := iNode.replica.getSegmentIDs(partitionID)
			if err != nil {
				log.Warn(err.Error())
				continue
			}
			resultSegmentIDs = append(resultSegmentIDs, segmentIDs...)
		}
		for _, segmentID := range resultSegmentIDs {
			segment, err := iNode.replica.getSegmentByID(segmentID)
			if err != nil {
				log.Warn(err.Error())
				continue
			}
			pks, err := filterSegmentsByPKs(delMsg.PrimaryKeys, segment)
			if err != nil {
				log.Warn(err.Error())
				continue
			}
			if len(pks) > 0 {
				offset := segment.segmentPreDelete(len(pks))
				if err != nil {
					log.Warn(err.Error())
					continue
				}
				delData.deleteIDs[segmentID] = append(delData.deleteIDs[segmentID], pks...)
				// TODO(yukun) get offset of pks
				delData.deleteTimestamps[segmentID] = append(delData.deleteTimestamps[segmentID], delMsg.Timestamps[:len(pks)]...)
				delData.deleteOffset[segmentID] = offset
			}
		}
	}

	// 2. do delete
	for segmentID := range delData.deleteIDs {
		wg.Add(1)
		go iNode.delete(delData, segmentID, &wg)
	}
	wg.Wait()

	var res Msg = &serviceTimeMsg{
		timeRange: iMsg.timeRange,
	}
	for _, sp := range spans {
		sp.Finish()
	}

	return []Msg{res}
}

func filterSegmentsByPKs(pks []int64, segment *Segment) ([]int64, error) {
	if pks == nil {
		return nil, fmt.Errorf("pks is nil when getSegmentsByPKs")
	}
	if segment == nil {
		return nil, fmt.Errorf("segments is nil when getSegmentsByPKs")
	}
	buf := make([]byte, 8)
	res := make([]int64, 0)
	for _, pk := range pks {
		binary.BigEndian.PutUint64(buf, uint64(pk))
		exist := segment.pkFilter.Test(buf)
		if exist {
			res = append(res, pk)
		}
	}
	log.Debug("In filterSegmentsByPKs", zap.Any("pk", res), zap.Any("segment", segment.segmentID))
	return res, nil
}

func (iNode *insertNode) insert(iData *insertData, segmentID UniqueID, wg *sync.WaitGroup) {
	log.Debug("QueryNode::iNode::insert", zap.Any("SegmentID", segmentID))
	var targetSegment, err = iNode.replica.getSegmentByID(segmentID)
	if err != nil {
		log.Warn("cannot find segment:", zap.Int64("segmentID", segmentID))
		// TODO: add error handling
		wg.Done()
		return
	}

	if targetSegment.segmentType != segmentTypeGrowing {
		wg.Done()
		return
	}

	ids := iData.insertIDs[segmentID]
	timestamps := iData.insertTimestamps[segmentID]
	records := iData.insertRecords[segmentID]
	offsets := iData.insertOffset[segmentID]

	err = targetSegment.segmentInsert(offsets, &ids, &timestamps, &records)
	if err != nil {
		log.Debug("QueryNode: targetSegmentInsert failed", zap.Error(err))
		// TODO: add error handling
		wg.Done()
		return
	}

	log.Debug("Do insert done", zap.Int("len", len(iData.insertIDs[segmentID])),
		zap.Int64("segmentID", segmentID), zap.Any("IDS", iData.insertPKs))
	wg.Done()
}

func (iNode *insertNode) delete(deleteData *deleteData, segmentID UniqueID, wg *sync.WaitGroup) {
	defer wg.Done()
	log.Debug("QueryNode::iNode::delete", zap.Any("SegmentID", segmentID))
	var targetSegment, err = iNode.replica.getSegmentByID(segmentID)
	if err != nil {
		log.Warn("Cannot find segment:", zap.Int64("segmentID", segmentID))
		return
	}

	if targetSegment.segmentType != segmentTypeGrowing {
		return
	}

	ids := deleteData.deleteIDs[segmentID]
	timestamps := deleteData.deleteTimestamps[segmentID]
	offset := deleteData.deleteOffset[segmentID]

	err = targetSegment.segmentDelete(offset, &ids, &timestamps)
	if err != nil {
		log.Warn("QueryNode: targetSegmentDelete failed", zap.Error(err))
		return
	}

	log.Debug("Do delete done", zap.Int("len", len(deleteData.deleteIDs[segmentID])), zap.Int64("segmentID", segmentID))
}

func (iNode *insertNode) getPrimaryKeys(msg *msgstream.InsertMsg) []int64 {
	if len(msg.RowIDs) != len(msg.Timestamps) || len(msg.RowIDs) != len(msg.RowData) {
		log.Warn("misaligned messages detected")
		return nil
	}
	collectionID := msg.GetCollectionID()

	collection, err := iNode.replica.getCollectionByID(collectionID)
	if err != nil {
		log.Warn("collection cannot be found")
		return nil
	}

	offset := 0
	for _, field := range collection.schema.Fields {
		if field.IsPrimaryKey {
			break
		}
		switch field.DataType {
		case schemapb.DataType_Bool:
			offset++
		case schemapb.DataType_Int8:
			offset++
		case schemapb.DataType_Int16:
			offset += 2
		case schemapb.DataType_Int32:
			offset += 4
		case schemapb.DataType_Int64:
			offset += 8
		case schemapb.DataType_Float:
			offset += 4
		case schemapb.DataType_Double:
			offset += 8
		case schemapb.DataType_FloatVector:
			for _, t := range field.TypeParams {
				if t.Key == "dim" {
					dim, err := strconv.Atoi(t.Value)
					if err != nil {
						log.Error("strconv wrong on get dim", zap.Error(err))
						break
					}
					offset += dim * 4
					break
				}
			}
		case schemapb.DataType_BinaryVector:
			var dim int
			for _, t := range field.TypeParams {
				if t.Key == "dim" {
					dim, err = strconv.Atoi(t.Value)
					if err != nil {
						log.Error("strconv wrong on get dim", zap.Error(err))
						return nil
					}
					offset += dim / 8
					break
				}
			}
		}
	}

	blobReaders := make([]io.Reader, len(msg.RowData))
	for i, blob := range msg.RowData {
		blobReaders[i] = bytes.NewReader(blob.GetValue()[offset : offset+8])
	}
	pks := make([]int64, len(blobReaders))

	for i, reader := range blobReaders {
		err = binary.Read(reader, binary.LittleEndian, &pks[i])
		if err != nil {
			log.Warn("binary read blob value failed", zap.Error(err))
		}
	}

	return pks
}
func newInsertNode(replica ReplicaInterface) *insertNode {
	maxQueueLength := Params.FlowGraphMaxQueueLength
	maxParallelism := Params.FlowGraphMaxParallelism

	baseNode := baseNode{}
	baseNode.SetMaxQueueLength(maxQueueLength)
	baseNode.SetMaxParallelism(maxParallelism)

	return &insertNode{
		baseNode: baseNode,
		replica:  replica,
	}
}
