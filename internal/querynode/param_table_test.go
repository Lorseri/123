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
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParamTable_PulsarAddress(t *testing.T) {
	address := Params.PulsarAddress
	split := strings.Split(address, ":")
	assert.Equal(t, "pulsar", split[0])
	assert.Equal(t, "6650", split[len(split)-1])
}

func TestParamTable_QueryNode(t *testing.T) {
	t.Run("Test id", func(t *testing.T) {
		id := Params.QueryNodeID
		assert.Contains(t, Params.QueryNodeIDList(), id)
	})

	t.Run("Test time tick channel", func(t *testing.T) {
		ch := Params.QueryTimeTickChannelName
		assert.Equal(t, ch, "queryTimeTick")
	})

	t.Run("Test time tick ReceiveBufSize", func(t *testing.T) {
		size := Params.QueryTimeTickReceiveBufSize
		assert.Equal(t, size, int64(64))
	})
}

func TestParamTable_minio(t *testing.T) {
	t.Run("Test endPoint", func(t *testing.T) {
		endPoint := Params.MinioEndPoint
		equal := endPoint == "localhost:9000" || endPoint == "minio:9000"
		assert.Equal(t, equal, true)
	})

	t.Run("Test accessKeyID", func(t *testing.T) {
		accessKeyID := Params.MinioAccessKeyID
		assert.Equal(t, accessKeyID, "minioadmin")
	})

	t.Run("Test secretAccessKey", func(t *testing.T) {
		secretAccessKey := Params.MinioSecretAccessKey
		assert.Equal(t, secretAccessKey, "minioadmin")
	})

	t.Run("Test useSSL", func(t *testing.T) {
		useSSL := Params.MinioUseSSLStr
		assert.Equal(t, useSSL, false)
	})
}

func TestParamTable_insertChannelRange(t *testing.T) {
	channelRange := Params.InsertChannelRange
	assert.Equal(t, 2, len(channelRange))
}

func TestParamTable_statsServiceTimeInterval(t *testing.T) {
	interval := Params.StatsPublishInterval
	assert.Equal(t, 1000, interval)
}

func TestParamTable_statsMsgStreamReceiveBufSize(t *testing.T) {
	bufSize := Params.StatsReceiveBufSize
	assert.Equal(t, int64(64), bufSize)
}

func TestParamTable_insertMsgStreamReceiveBufSize(t *testing.T) {
	bufSize := Params.InsertReceiveBufSize
	assert.Equal(t, int64(1024), bufSize)
}

func TestParamTable_ddMsgStreamReceiveBufSize(t *testing.T) {
	bufSize := Params.DDReceiveBufSize
	assert.Equal(t, bufSize, int64(64))
}

func TestParamTable_searchMsgStreamReceiveBufSize(t *testing.T) {
	bufSize := Params.SearchReceiveBufSize
	assert.Equal(t, int64(512), bufSize)
}

func TestParamTable_searchResultMsgStreamReceiveBufSize(t *testing.T) {
	bufSize := Params.SearchResultReceiveBufSize
	assert.Equal(t, int64(64), bufSize)
}

func TestParamTable_searchPulsarBufSize(t *testing.T) {
	bufSize := Params.SearchPulsarBufSize
	assert.Equal(t, int64(512), bufSize)
}

func TestParamTable_insertPulsarBufSize(t *testing.T) {
	bufSize := Params.InsertPulsarBufSize
	assert.Equal(t, int64(1024), bufSize)
}

func TestParamTable_ddPulsarBufSize(t *testing.T) {
	bufSize := Params.DDPulsarBufSize
	assert.Equal(t, bufSize, int64(64))
}

func TestParamTable_flowGraphMaxQueueLength(t *testing.T) {
	length := Params.FlowGraphMaxQueueLength
	assert.Equal(t, int32(1024), length)
}

func TestParamTable_flowGraphMaxParallelism(t *testing.T) {
	maxParallelism := Params.FlowGraphMaxParallelism
	assert.Equal(t, int32(1024), maxParallelism)
}

func TestParamTable_insertChannelNames(t *testing.T) {
	names := Params.InsertChannelNames
	channelRange := Params.InsertChannelRange
	num := channelRange[1] - channelRange[0]
	num = num / Params.QueryNodeNum
	assert.Equal(t, num, len(names))
	start := num * Params.SliceIndex
	contains := strings.Contains(names[0], fmt.Sprintf("insert-%d", channelRange[start]))
	assert.Equal(t, contains, true)
}

func TestParamTable_searchChannelNames(t *testing.T) {
	names := Params.SearchChannelNames
	assert.Equal(t, len(names), 1)
	contains := strings.Contains(names[0], "search-0")
	assert.Equal(t, contains, true)
}

func TestParamTable_searchResultChannelNames(t *testing.T) {
	names := Params.SearchResultChannelNames
	assert.NotNil(t, names)
}

func TestParamTable_msgChannelSubName(t *testing.T) {
	name := Params.MsgChannelSubName
	expectName := fmt.Sprintf("queryNode-%d", Params.QueryNodeID)
	assert.Equal(t, expectName, name)
}

func TestParamTable_statsChannelName(t *testing.T) {
	name := Params.StatsChannelName
	contains := strings.Contains(name, "query-node-stats")
	assert.Equal(t, contains, true)
}

func TestParamTable_metaRootPath(t *testing.T) {
	path := Params.MetaRootPath
	fmt.Println(path)
}

func TestParamTable_ddChannelName(t *testing.T) {
	names := Params.DDChannelNames
	contains := strings.Contains(names[0], "data-definition")
	assert.Equal(t, contains, true)
}
