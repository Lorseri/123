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

package grpcqueryservice

import (
	"sync"

	"github.com/milvus-io/milvus/internal/util/paramtable"
)

var Params ParamTable
var once sync.Once

type ParamTable struct {
	paramtable.BaseTable
	Port int

	IndexServiceAddress string
	MasterAddress       string
	DataCoordAddress    string
}

func (pt *ParamTable) Init() {
	once.Do(func() {
		pt.BaseTable.Init()
		pt.initPort()
		pt.initMasterAddress()
		pt.initIndexServiceAddress()
		pt.initDataCoordAddress()

	})
}

func (pt *ParamTable) initMasterAddress() {
	ret, err := pt.Load("_MasterAddress")
	if err != nil {
		panic(err)
	}
	pt.MasterAddress = ret
}

func (pt *ParamTable) initIndexServiceAddress() {
	ret, err := pt.Load("IndexServiceAddress")
	if err != nil {
		panic(err)
	}
	pt.IndexServiceAddress = ret
}

func (pt *ParamTable) initDataCoordAddress() {
	ret, err := pt.Load("_DataCoordAddress")
	if err != nil {
		panic(err)
	}
	pt.DataCoordAddress = ret
}

func (pt *ParamTable) initPort() {
	pt.Port = pt.ParseInt("queryService.port")
}
