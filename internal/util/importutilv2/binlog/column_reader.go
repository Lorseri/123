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

package binlog

import (
	"github.com/milvus-io/milvus-proto/go-api/v2/schemapb"
	"github.com/milvus-io/milvus/internal/storage"
	"github.com/milvus-io/milvus/pkg/util/typeutil"
	"github.com/samber/lo"
)

type columnReader struct {
	reader      *storage.BinlogReader
	fieldSchema *schemapb.FieldSchema
}

func newColumnReader(cm storage.ChunkManager, fieldSchema *schemapb.FieldSchema, path string) (*columnReader, error) {
	reader, err := newBinlogReader(cm, path)
	if err != nil {
		return nil, err
	}
	return &columnReader{
		reader:      reader,
		fieldSchema: fieldSchema,
	}, nil
}

func (r *columnReader) Next(_ int64) (storage.FieldData, error) {
	fieldData, err := storage.NewFieldData(r.fieldSchema.GetDataType(), r.fieldSchema)
	if err != nil {
		return nil, err
	}
	result, err := readData(r.reader, storage.InsertEventType)
	if err != nil {
		return nil, err
	}
	if typeutil.IsVectorType(r.fieldSchema.GetDataType()) {
		dim, err := typeutil.GetDim(r.fieldSchema)
		if err != nil {
			return nil, err
		}
		chunks := lo.Chunk(result, int(dim))
		result = make([]any, 0, len(chunks))
		for _, chunk := range chunks {
			result = append(result, chunk)
		}
	}
	for _, v := range result {
		err = fieldData.AppendRow(v)
		if err != nil {
			return nil, err
		}
	}
	return fieldData, nil
}
