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

package datanode

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"

	"github.com/milvus-io/milvus-proto/go-api/v2/msgpb"
	"github.com/milvus-io/milvus/internal/datanode/broker"
	"github.com/milvus-io/milvus/pkg/util/paramtable"
)

type ChannelCPUpdaterSuite struct {
	suite.Suite

	updater *channelCheckpointUpdater
}

func (s *ChannelCPUpdaterSuite) SetupTest() {
	s.updater = newChannelCheckpointUpdater(&DataNode{})
}

func (s *ChannelCPUpdaterSuite) TestUpdate() {
	paramtable.Get().Save(paramtable.Get().DataNodeCfg.ChannelCheckpointUpdaterTick.Key, "0.01")
	defer paramtable.Get().Save(paramtable.Get().DataNodeCfg.ChannelCheckpointUpdaterTick.Key, "5")

	b := broker.NewMockBroker(s.T())
	b.EXPECT().UpdateChannelCheckpoint(mock.Anything, mock.Anything).Return(nil)
	s.updater.dn.broker = b

	go s.updater.start()
	defer s.updater.close()

	tasksNum := 100000
	counter := 0
	for i := 0; i < tasksNum; i++ {
		s.updater.addTask(&msgpb.MsgPosition{
			ChannelName: fmt.Sprintf("ch-%d", i),
		}, func() {
			counter++
			fmt.Println(counter)
		})
	}
	s.Eventually(func() bool {
		return counter == tasksNum
	}, time.Second*10, time.Millisecond*100)
}

func TestChannelCPUpdater(t *testing.T) {
	suite.Run(t, new(ChannelCPUpdaterSuite))
}
