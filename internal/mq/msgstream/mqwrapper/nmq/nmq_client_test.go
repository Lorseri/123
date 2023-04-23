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

package nmq

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"path"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	nmqserver "github.com/milvus-io/milvus/internal/mq/mqimpl/natsmq/server"
	"github.com/milvus-io/milvus/pkg/mq/msgstream/mqwrapper"
)

var nats_server_address string

func TestMain(m *testing.M) {
	tmpDir := path.Join(os.TempDir(), "nmq_client_test")
	defer os.RemoveAll(tmpDir)
	nmqserver.InitNatsMQ(tmpDir)
	nats_server_address = nmqserver.Nmq.ClientURL()
	defer nmqserver.CloseNatsMQ()

	exitCode := m.Run()
	os.Exit(exitCode)
}

func createNmqClient() (*nmqClient, error) {
	return NewClient(nats_server_address)
}

func Test_NewNmqClient(t *testing.T) {
	client, err := createNmqClient()
	defer client.Close()
	assert.Nil(t, err)
	assert.NotNil(t, client)
}

func TestNmqClient_CreateProducer(t *testing.T) {
	client, err := createNmqClient()
	defer client.Close()
	assert.Nil(t, err)
	assert.NotNil(t, client)

	topic := "TestNmqClient_CreateProducer"
	proOpts := mqwrapper.ProducerOptions{Topic: topic}
	producer, err := client.CreateProducer(proOpts)

	defer producer.Close()
	assert.Nil(t, err)
	assert.NotNil(t, producer)

	nmqProducer := producer.(*nmqProducer)
	assert.Equal(t, nmqProducer.Topic(), topic)

	msg := &mqwrapper.ProducerMessage{
		Payload:    []byte{},
		Properties: nil,
	}
	_, err = nmqProducer.Send(context.TODO(), msg)
	assert.Nil(t, err)

	invalidOpts := mqwrapper.ProducerOptions{Topic: ""}
	producer, e := client.CreateProducer(invalidOpts)
	assert.Nil(t, producer)
	assert.Error(t, e)
}

func TestNmqClient_GetLatestMsg(t *testing.T) {
	client, err := createNmqClient()
	assert.Nil(t, err)
	defer client.Close()

	topic := fmt.Sprintf("t2GetLatestMsg-%d", rand.Int())
	proOpts := mqwrapper.ProducerOptions{Topic: topic}
	producer, err := client.CreateProducer(proOpts)
	assert.Nil(t, err)
	defer producer.Close()

	for i := 0; i < 10; i++ {
		msg := &mqwrapper.ProducerMessage{
			Payload:    []byte{byte(i)},
			Properties: nil,
		}
		_, err = producer.Send(context.TODO(), msg)
		assert.Nil(t, err)
	}

	subName := "subName"
	consumerOpts := mqwrapper.ConsumerOptions{
		Topic:                       topic,
		SubscriptionName:            subName,
		SubscriptionInitialPosition: mqwrapper.SubscriptionPositionEarliest,
		BufSize:                     1024,
	}

	consumer, err := client.Subscribe(consumerOpts)
	assert.Nil(t, err)

	expectLastMsg, err := consumer.GetLatestMsgID()
	assert.Nil(t, err)

	var actualLastMsg mqwrapper.Message
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	for i := 0; i < 10; i++ {
		select {
		case <-ctx.Done():
			fmt.Println(i)
			assert.FailNow(t, "consumer failed to yield message in 100 milliseconds")
		case msg := <-consumer.Chan():
			consumer.Ack(msg)
			actualLastMsg = msg
		}
	}
	require.NotNil(t, actualLastMsg)
	ret, err := expectLastMsg.LessOrEqualThan(actualLastMsg.ID().Serialize())
	assert.Nil(t, err)
	assert.True(t, ret)
}

func TestNmqClient_Subscribe(t *testing.T) {
	client, err := createNmqClient()
	defer client.Close()
	assert.Nil(t, err)
	assert.NotNil(t, client)

	topic := "TestNmqClient_Subscribe"
	proOpts := mqwrapper.ProducerOptions{Topic: topic}
	producer, err := client.CreateProducer(proOpts)
	defer producer.Close()
	assert.Nil(t, err)
	assert.NotNil(t, producer)

	subName := "subName"
	consumerOpts := mqwrapper.ConsumerOptions{
		Topic:                       "",
		SubscriptionName:            subName,
		SubscriptionInitialPosition: mqwrapper.SubscriptionPositionEarliest,
		BufSize:                     1024,
	}

	consumer, err := client.Subscribe(consumerOpts)
	assert.NotNil(t, err)
	assert.Nil(t, consumer)

	consumerOpts.Topic = topic
	consumer, err = client.Subscribe(consumerOpts)
	defer consumer.Close()
	assert.Nil(t, err)
	assert.NotNil(t, consumer)
	assert.Equal(t, consumer.Subscription(), subName)

	msg := &mqwrapper.ProducerMessage{
		Payload:    []byte{1},
		Properties: nil,
	}
	_, err = producer.Send(context.TODO(), msg)
	assert.Nil(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Millisecond)
	defer cancel()
	select {
	case <-ctx.Done():
		assert.FailNow(t, "consumer failed to yield message in 100 milliseconds")
	case msg := <-consumer.Chan():
		consumer.Ack(msg)
		nmqmsg := msg.(*nmqMessage)
		msgPayload := nmqmsg.Payload()
		assert.NotEmpty(t, msgPayload)
		msgTopic := nmqmsg.Topic()
		assert.Equal(t, msgTopic, topic)
		msgProp := nmqmsg.Properties()
		assert.Empty(t, msgProp)
		msgID := nmqmsg.ID()
		rID := msgID.(*nmqID)
		assert.Equal(t, rID.messageID, MessageIDType(1))
	}
}

func TestNmqClient_EarliestMessageID(t *testing.T) {
	client, _ := createNmqClient()
	defer client.Close()

	mid := client.EarliestMessageID()
	assert.NotNil(t, mid)
	nmqmsg := mid.(*nmqID)
	assert.Equal(t, nmqmsg.messageID, MessageIDType(1))
}

func TestNmqClient_StringToMsgID(t *testing.T) {
	client, _ := createNmqClient()
	defer client.Close()

	str := "5"
	res, err := client.StringToMsgID(str)
	assert.Nil(t, err)
	assert.NotNil(t, res)

	str = "X"
	res, err = client.StringToMsgID(str)
	assert.Nil(t, res)
	assert.NotNil(t, err)
}

func TestNmqClient_BytesToMsgID(t *testing.T) {
	client, _ := createNmqClient()
	defer client.Close()

	mid := client.EarliestMessageID()
	res, err := client.BytesToMsgID(mid.Serialize())
	assert.Nil(t, err)
	assert.NotNil(t, res)
}
