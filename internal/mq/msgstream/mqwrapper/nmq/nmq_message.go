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

package nmq

import (
	"github.com/milvus-io/milvus/pkg/mq/msgstream/mqwrapper"
)

// Check nmqMessage implements ConsumerMessage
var _ mqwrapper.Message = (*nmqMessage)(nil)

// Message is the message content of a consumer message
type Message struct {
	MsgID      MessageID
	Topic      string
	Payload    []byte
	Properties map[string]string
}

// nmqMessage wraps the message for natsmq
type nmqMessage struct {
	msg Message
}

// Topic returns the topic name of natsmq message
func (nm *nmqMessage) Topic() string {
	return nm.msg.Topic
}

// Properties returns the properties of natsmq message
func (nm *nmqMessage) Properties() map[string]string {
	return nm.msg.Properties
}

// Payload returns the payload of natsmq message
func (nm *nmqMessage) Payload() []byte {
	return nm.msg.Payload
}

// ID returns the id of natsmq message
func (nm *nmqMessage) ID() mqwrapper.MessageID {
	return &nmqID{messageID: nm.msg.MsgID}
}
