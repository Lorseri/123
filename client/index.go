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

package client

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"google.golang.org/grpc"

	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"github.com/milvus-io/milvus/client/v2/entity"
	"github.com/milvus-io/milvus/client/v2/index"
	"github.com/milvus-io/milvus/pkg/util/merr"
)

type CreateIndexTask struct {
	client         *Client
	collectionName string
	fieldName      string
	indexName      string
	interval       time.Duration
}

func (t *CreateIndexTask) Await(ctx context.Context) error {
	ticker := time.NewTicker(t.interval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			finished := false
			err := t.client.callService(func(milvusService milvuspb.MilvusServiceClient) error {
				resp, err := milvusService.DescribeIndex(ctx, &milvuspb.DescribeIndexRequest{
					CollectionName: t.collectionName,
					FieldName:      t.fieldName,
					IndexName:      t.indexName,
				})
				err = merr.CheckRPCCall(resp, err)
				if err != nil {
					return err
				}

				for _, info := range resp.GetIndexDescriptions() {
					if (t.indexName == "" && info.GetFieldName() == t.fieldName) || t.indexName == info.GetIndexName() {
						switch info.GetState() {
						case commonpb.IndexState_Finished:
							finished = true
							return nil
						case commonpb.IndexState_Failed:
							return fmt.Errorf("create index failed, reason: %s", info.GetIndexStateFailReason())
						}
					}
				}
				return nil
			})
			if err != nil {
				return err
			}
			if finished {
				return nil
			}
			ticker.Reset(t.interval)
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func (c *Client) CreateIndex(ctx context.Context, option CreateIndexOption, callOptions ...grpc.CallOption) (*CreateIndexTask, error) {
	req := option.Request()
	var task *CreateIndexTask

	err := c.callService(func(milvusService milvuspb.MilvusServiceClient) error {
		resp, err := milvusService.CreateIndex(ctx, req, callOptions...)
		if err = merr.CheckRPCCall(resp, err); err != nil {
			return err
		}

		task = &CreateIndexTask{
			client:         c,
			collectionName: req.GetCollectionName(),
			fieldName:      req.GetFieldName(),
			indexName:      req.GetIndexName(),
			interval:       time.Millisecond * 100,
		}

		return nil
	})

	return task, err
}

func (c *Client) ListIndexes(ctx context.Context, opt ListIndexOption, callOptions ...grpc.CallOption) ([]string, error) {
	req := opt.Request()

	var indexes []string

	err := c.callService(func(milvusService milvuspb.MilvusServiceClient) error {
		resp, err := milvusService.DescribeIndex(ctx, req, callOptions...)
		if err = merr.CheckRPCCall(resp, err); err != nil {
			return err
		}
		for _, idxDef := range resp.GetIndexDescriptions() {
			if opt.Matches(idxDef) {
				indexes = append(indexes, idxDef.GetIndexName())
			}
		}
		return nil
	})
	return indexes, err
}

func (c *Client) DescribeIndex(ctx context.Context, opt DescribeIndexOption, callOptions ...grpc.CallOption) (index.Index, error) {
	req := opt.Request()
	var idx index.Index

	err := c.callService(func(milvusService milvuspb.MilvusServiceClient) error {
		resp, err := milvusService.DescribeIndex(ctx, req, callOptions...)
		if err = merr.CheckRPCCall(resp, err); err != nil {
			return err
		}

		if len(resp.GetIndexDescriptions()) == 0 {
			return merr.WrapErrIndexNotFound(req.GetIndexName())
		}
		for _, idxDef := range resp.GetIndexDescriptions() {
			if idxDef.GetIndexName() == req.GetIndexName() {
				params := entity.KvPairsMap(idxDef.GetParams())
				params["total_rows"] = strconv.FormatInt(idxDef.GetTotalRows(), 10)
				params["indexed_rows"] = strconv.FormatInt(idxDef.GetIndexedRows(), 10)
				params["pending_index_rows"] = strconv.FormatInt(idxDef.GetPendingIndexRows(), 10)
				params["state"] = idxDef.GetState().String()

				idx = index.NewGenericIndex(idxDef.GetIndexName(), params)
			}
		}
		return nil
	})

	return idx, err
}

func (c *Client) DropIndex(ctx context.Context, opt DropIndexOption, callOptions ...grpc.CallOption) error {
	req := opt.Request()

	return c.callService(func(milvusService milvuspb.MilvusServiceClient) error {
		resp, err := milvusService.DropIndex(ctx, req, callOptions...)
		return merr.CheckRPCCall(resp, err)
	})
}
