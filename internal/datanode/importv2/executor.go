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

package importv2

import (
	"context"
	"fmt"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/internal/datanode/metacache"
	"github.com/milvus-io/milvus/internal/datanode/syncmgr"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/conc"
	"github.com/milvus-io/milvus/pkg/util/merr"
	"github.com/milvus-io/milvus/pkg/util/typeutil"
	"github.com/samber/lo"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
)

type Executor interface {
	Start()
	Close()
}

type executor struct {
	metaCache metacache.MetaCache
	manager   TaskManager
	handler   Handler
	cm        storage.ChunkManager
}

func (e *executor) estimateReadRows(schema *schemapb.CollectionSchema) (int64, error) {
	const BufferSize = 16 * 1024 * 1024 // TODO: dyh, make it configurable
	sizePerRow, err := typeutil.EstimateSizePerRecord(schema)
	if err != nil {
		return 0, err
	}
	return int64(BufferSize / sizePerRow), nil
}

func (e *executor) handleErr(task Task, err error, msg string) {
	log.Warn(msg, zap.Int64("taskID", task.GetTaskID()),
		zap.Int64("requestID", task.GetRequestID()),
		zap.Int64("collectionID", task.GetCollectionID()),
		zap.String("state", task.GetState().String()),
		zap.String("type", task.GetType().String()),
		zap.Error(err))
	e.manager.Update(task.GetTaskID(), UpdateState(datapb.ImportState_Failed), UpdateReason(err.Error()))
}

func (e *executor) runPreImportTask(task Task) {
	e.manager.Update(task.GetTaskID(), UpdateState(datapb.ImportState_InProgress))
	files := lo.Map(task.(*PreImportTask).GetFileStats(), func(fileStat *datapb.ImportFileStats, _ int) *datapb.ImportFile {
		return fileStat.GetImportFile()
	})

	wg, _ := errgroup.WithContext(context.TODO()) // TODO: dyh, set timeout
	for i, file := range files {
		i := i
		file := file
		wg.Go(func() error {
			reader := NewReader(e.cm, task.GetSchema(), file)
			stat, err := reader.ReadStats()
			if err != nil {
				e.handleErr(task, err, "read stats failed")
				return err
			}
			e.manager.Update(task.GetTaskID(), UpdateFileStat(i, stat))
			return nil
		})
	}
	err := wg.Wait()
	if err != nil {
		e.handleErr(task, err, "preimport failed")
		return
	}
	e.manager.Update(task.GetTaskID(), UpdateState(datapb.ImportState_Completed))
}

func (e *executor) runImportTask(task Task) {
	e.manager.Update(task.GetTaskID(), UpdateState(datapb.ImportState_InProgress))
	count, err := e.estimateReadRows(task.GetSchema())
	if err != nil {
		e.handleErr(task, err, fmt.Sprintf("estimate rows size failed"))
		return
	}
	for _, fileInfo := range task.(*ImportTask).req.GetFilesInfo() {
		for {
			rows, err := e.doImportOnFile(count, task, fileInfo)
			if err != nil {
				e.handleErr(task, err, fmt.Sprintf("do import on file %s failed"))
				return
			}
			if rows == 0 {
				break
			}
		}
	}
	e.manager.Update(task.GetTaskID(), UpdateState(datapb.ImportState_Completed))
}

func (e *executor) doImportOnFile(count int64, task Task, fileInfo *datapb.ImportFileRequestInfo) (int, error) {
	reader := NewReader(e.cm, task.GetSchema(), fileInfo.GetImportFile())
	insertData, err := reader.Next(count)
	if err != nil {
		e.handleErr(task, err, fmt.Sprintf(""))
		return 0, err
	}
	readRows := insertData.GetRowNum()
	if readRows == 0 {
		return 0, nil
	}
	hashedData := e.handler.Hash(insertData)
	futures := make([]*conc.Future[error], 0)
	syncTasks := make([]*syncmgr.SyncTask, 0)
	for vchannel, datas := range hashedData {
		for partitionID, data := range datas {
			segmentID := PickSegment(task, fileInfo, vchannel, partitionID)
			AddSegment(e.metaCache, vchannel, segmentID, partitionID, task.GetCollectionID())
			syncTask := NewSyncTask(segmentID, partitionID, task.GetCollectionID(), vchannel, data)
			future := e.handler.SyncData(context.TODO(), syncTask)
			futures = append(futures, future)
			syncTasks = append(syncTasks, syncTask)
		}
	}
	err = conc.AwaitAll(futures...) // TODO: dyh, return futures and syncTasks to increase concurrence
	if err != nil {
		return 0, err
	}
	for _, syncTask := range syncTasks {
		segmentID := syncTask.SegmentID()
		insertBinlogs, statsBinlog, _ := syncTask.Binlogs()
		segment, ok := e.metaCache.GetSegmentByID(segmentID)
		if !ok {
			return 0, merr.WrapErrSegmentNotFound(segmentID, "import failed")
		}
		segmentInfo := &datapb.ImportSegmentInfo{
			SegmentID:    segmentID,
			ImportedRows: segment.FlushedRows(),
			Binlogs:      lo.Values(insertBinlogs),
			Statslogs:    lo.Values(statsBinlog),
		}
		e.manager.Update(task.GetTaskID(), UpdateSegmentInfo(segmentInfo))
	}
	return readRows, nil
}
