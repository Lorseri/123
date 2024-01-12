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
	"github.com/golang/protobuf/proto"
	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/util/importutilv2"
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/funcutil"
	"github.com/milvus-io/milvus/pkg/util/metric"
	"github.com/milvus-io/milvus/tests/integration"
	"github.com/stretchr/testify/suite"
	"go.uber.org/zap"
	"math/rand"
	"os"
	"testing"
	"time"
)

type BulkInsertSuite struct {
	integration.MiniClusterSuite

	fileType importutilv2.FileType
}

func (s *BulkInsertSuite) testNormal() {
	const (
		rowCount = 100
	)

	c := s.Cluster
	ctx, cancel := context.WithTimeout(c.GetContext(), 60*time.Second)
	defer cancel()

	collectionName := "TestBulkInsert" + funcutil.GenRandomStr()

	schema := integration.ConstructSchema(collectionName, dim, true,
		&schemapb.FieldSchema{FieldID: 100, Name: "id", DataType: schemapb.DataType_Int64, IsPrimaryKey: true, AutoID: true},
		&schemapb.FieldSchema{FieldID: 101, Name: "image_path", DataType: schemapb.DataType_VarChar, TypeParams: []*commonpb.KeyValuePair{{Key: common.MaxLengthKey, Value: "65535"}}},
		&schemapb.FieldSchema{FieldID: 102, Name: "embeddings", DataType: schemapb.DataType_FloatVector, TypeParams: []*commonpb.KeyValuePair{{Key: common.DimKey, Value: "128"}}},
	)
	marshaledSchema, err := proto.Marshal(schema)
	s.NoError(err)

	createCollectionStatus, err := c.Proxy.CreateCollection(ctx, &milvuspb.CreateCollectionRequest{
		CollectionName: collectionName,
		Schema:         marshaledSchema,
		ShardsNum:      common.DefaultShardsNum,
	})
	s.NoError(err)
	s.Equal(commonpb.ErrorCode_Success, createCollectionStatus.GetErrorCode())

	var files []*internalpb.ImportFile
	if s.fileType == importutilv2.Numpy {
		path1 := c.ChunkManager.RootPath() + "/" + "embeddings.npy"
		err = GenerateNumpyFile(path1, rowCount, schemapb.DataType_FloatVector)
		s.NoError(err)
		defer os.Remove(path1)
		path2 := c.ChunkManager.RootPath() + "/" + "image_path.npy"
		err = GenerateNumpyFile(path2, rowCount, schemapb.DataType_VarChar)
		s.NoError(err)
		defer os.Remove(path2)

		files = []*internalpb.ImportFile{
			{
				Paths: []string{
					c.ChunkManager.RootPath() + "/" + "embeddings.npy",
					c.ChunkManager.RootPath() + "/" + "image_path.npy",
				},
			},
		}
	} else if s.fileType == importutilv2.JSON {
		rowBasedFile := c.ChunkManager.RootPath() + "/" + "test.json"
		GenerateJsonFile(s.T(), rowBasedFile, schema, rowCount)
		defer os.Remove(rowBasedFile)
		files = []*internalpb.ImportFile{
			{
				Paths: []string{
					rowBasedFile,
				},
			},
		}
	} else if s.fileType == importutilv2.Parquet {
		filePath := fmt.Sprintf("/tmp/test_%d.parquet", rand.Int())
		err = GenerateParquetFile(filePath, schema, rowCount)
		s.NoError(err)
		defer os.Remove(filePath)
		files = []*internalpb.ImportFile{
			{
				Paths: []string{
					filePath,
				},
			},
		}
	}

	importResp, err := c.Proxy.ImportV2(ctx, &internalpb.ImportRequest{
		CollectionName: collectionName,
		Files:          files,
	})
	s.NoError(err)
	log.Info("Import result", zap.Any("importResp", importResp))

	request := importResp.GetRequestID()
	err = WaitForImportDone(ctx, c, request)
	s.NoError(err)

	segments, err := c.MetaWatcher.ShowSegments()
	s.NoError(err)
	s.NotEmpty(segments)

	// create index
	createIndexStatus, err := c.Proxy.CreateIndex(ctx, &milvuspb.CreateIndexRequest{
		CollectionName: collectionName,
		FieldName:      "embeddings",
		IndexName:      "_default",
		ExtraParams:    integration.ConstructIndexParam(dim, integration.IndexHNSW, metric.L2),
	})
	s.NoError(err)
	s.Equal(commonpb.ErrorCode_Success, createIndexStatus.GetErrorCode())

	s.WaitForIndexBuilt(ctx, collectionName, "embeddings")

	// load
	loadStatus, err := c.Proxy.LoadCollection(ctx, &milvuspb.LoadCollectionRequest{
		CollectionName: collectionName,
	})
	s.NoError(err)
	s.Equal(commonpb.ErrorCode_Success, loadStatus.GetErrorCode())
	s.WaitForLoad(ctx, collectionName)

	// search
	expr := ""
	nq := 10
	topk := 10
	roundDecimal := -1

	params := integration.GetSearchParams(integration.IndexHNSW, metric.L2)
	searchReq := integration.ConstructSearchRequest("", collectionName, expr,
		"embeddings", schemapb.DataType_FloatVector, nil, metric.L2, params, nq, dim, topk, roundDecimal)

	searchResult, err := c.Proxy.Search(ctx, searchReq)
	s.NoError(err)
	s.Equal(commonpb.ErrorCode_Success, searchResult.GetStatus().GetErrorCode())
}

func (s *BulkInsertSuite) TestNumpy() {
	s.fileType = importutilv2.Numpy
	s.testNormal()
}

func (s *BulkInsertSuite) TestJSON() {
	s.fileType = importutilv2.JSON
	s.testNormal()
}

func (s *BulkInsertSuite) TestParquet() {
	s.fileType = importutilv2.Parquet
	s.testNormal()
}

func (s *BulkInsertSuite) TestZeroRowCount() {
	const (
		rowCount = 0
	)

	c := s.Cluster
	ctx, cancel := context.WithTimeout(c.GetContext(), 60*time.Second)
	defer cancel()

	collectionName := "TestBulkInsert_" + funcutil.GenRandomStr()

	schema := integration.ConstructSchema(collectionName, dim, true,
		&schemapb.FieldSchema{FieldID: 100, Name: "id", DataType: schemapb.DataType_Int64, IsPrimaryKey: true, AutoID: true},
		&schemapb.FieldSchema{FieldID: 101, Name: "image_path", DataType: schemapb.DataType_VarChar, TypeParams: []*commonpb.KeyValuePair{{Key: common.MaxLengthKey, Value: "65535"}}},
		&schemapb.FieldSchema{FieldID: 102, Name: "embeddings", DataType: schemapb.DataType_FloatVector, TypeParams: []*commonpb.KeyValuePair{{Key: common.DimKey, Value: "128"}}},
	)
	marshaledSchema, err := proto.Marshal(schema)
	s.NoError(err)

	createCollectionStatus, err := c.Proxy.CreateCollection(ctx, &milvuspb.CreateCollectionRequest{
		CollectionName: collectionName,
		Schema:         marshaledSchema,
		ShardsNum:      common.DefaultShardsNum,
	})
	s.NoError(err)
	s.Equal(commonpb.ErrorCode_Success, createCollectionStatus.GetErrorCode())

	var files []*internalpb.ImportFile
	filePath := fmt.Sprintf("/tmp/test_%d.parquet", rand.Int())
	err = GenerateParquetFile(filePath, schema, rowCount)
	s.NoError(err)
	defer os.Remove(filePath)
	files = []*internalpb.ImportFile{
		{
			Paths: []string{
				filePath,
			},
		},
	}

	importResp, err := c.Proxy.ImportV2(ctx, &internalpb.ImportRequest{
		CollectionName: collectionName,
		Files:          files,
	})
	s.NoError(err)
	log.Info("Import result", zap.Any("importResp", importResp))

	request := importResp.GetRequestID()
	err = WaitForImportDone(ctx, c, request)
	s.NoError(err)

	segments, err := c.MetaWatcher.ShowSegments()
	s.NoError(err)
	s.Empty(segments)
}

func TestBulkInsert(t *testing.T) {
	suite.Run(t, new(BulkInsertSuite))
}
