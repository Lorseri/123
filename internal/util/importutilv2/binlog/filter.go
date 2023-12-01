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
	"github.com/milvus-io/milvus/pkg/common"
	"github.com/milvus-io/milvus/pkg/util/typeutil"
)

type Filter func(row map[int64]interface{}) bool

func FilterWithDelete(r *reader) (Filter, error) {
	pkField, err := typeutil.GetPrimaryFieldSchema(r.schema)
	if err != nil {
		return nil, err
	}
	return func(row map[int64]interface{}) bool {
		rowPk := row[pkField.GetFieldID()]
		for _, pk := range r.delData.Pks {
			if pk.GetValue() == rowPk {
				return false
			}
		}
		return true
	}, nil
}

func FilterWithTimerange(r *reader) Filter {
	return func(row map[int64]interface{}) bool {
		ts := row[common.TimeStampField].(int64)
		return uint64(ts) >= r.tsBegin && uint64(ts) <= r.tsEnd
	}
}
