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

package querynode

import (
	"context"
	"io/ioutil"
	"net/url"
	"os"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.etcd.io/etcd/server/v3/embed"

	"github.com/milvus-io/milvus/internal/util/dependency"

	etcdkv "github.com/milvus-io/milvus/internal/kv/etcd"
	"github.com/milvus-io/milvus/internal/types"
	"github.com/milvus-io/milvus/internal/util/etcd"
)

var embedetcdServer *embed.Etcd

// mock of query coordinator client
type queryCoordMock struct {
	types.QueryCoord
}

func setup() {
	os.Setenv("QUERY_NODE_ID", "1")
	Params.Init()
	Params.EtcdCfg.MetaRootPath = "/etcd/test/root/querynode"
}

func TestInitHook(t *testing.T) {
	Params.QueryNodeCfg.SoPath = ""
	a, err := initHook()
	assert.Nil(t, a)
	assert.NotNil(t, err)

	Params.QueryNodeCfg.SoPath = "/a/b/hook.so"
	a, err = initHook()
	assert.Nil(t, a)
	assert.NotNil(t, err)
	Params.QueryNodeCfg.SoPath = ""

	a = &mockHook1{}
	assert.Equal(t, a.(*mockHook1).mockString, "")
	assert.Error(t, a.Init("test"))
	assert.NoError(t, a.Init("t"))
	assert.Equal(t, a.(*mockHook1).mockString, "t")

	var hoo interface{} = &mockWrongHook{}
	a, ok := hoo.(Hook)
	assert.False(t, ok)
}

func initTestMeta(t *testing.T, node *QueryNode, collectionID UniqueID, segmentID UniqueID, optional ...bool) {
	schema := genTestCollectionSchema()

	node.metaReplica.addCollection(defaultCollectionID, schema, genTestIndexMeta(defaultCollectionID))

	collection, err := node.metaReplica.getCollectionByID(collectionID)
	assert.NoError(t, err)
	assert.Equal(t, collection.ID(), collectionID)
	assert.Equal(t, node.metaReplica.getCollectionNum(), 1)

	err = node.metaReplica.addPartition(collection.ID(), defaultPartitionID)
	assert.NoError(t, err)

	err = node.metaReplica.addSegment(segmentID, defaultPartitionID, collectionID, "", defaultSegmentVersion, defaultSegmentStartPosition, segmentTypeSealed)
	assert.NoError(t, err)
}

func newQueryNodeMock() *QueryNode {

	var ctx context.Context

	if debugUT {
		ctx = context.Background()
	} else {
		var cancel context.CancelFunc
		d := time.Now().Add(ctxTimeInMillisecond * time.Millisecond)
		ctx, cancel = context.WithDeadline(context.Background(), d)
		go func() {
			<-ctx.Done()
			cancel()
		}()
	}
	etcdCli, err := etcd.GetEtcdClient(
		Params.EtcdCfg.UseEmbedEtcd,
		Params.EtcdCfg.EtcdUseSSL,
		Params.EtcdCfg.Endpoints,
		Params.EtcdCfg.EtcdTLSCert,
		Params.EtcdCfg.EtcdTLSKey,
		Params.EtcdCfg.EtcdTLSCACert,
		Params.EtcdCfg.EtcdTLSMinVersion)
	if err != nil {
		panic(err)
	}
	etcdKV := etcdkv.NewEtcdKV(etcdCli, Params.EtcdCfg.MetaRootPath)

	factory := newMessageStreamFactory()
	svr := NewQueryNode(ctx, factory)
	tsReplica := newTSafeReplica()

	replica := newCollectionReplica()
	svr.metaReplica = replica
	svr.dataSyncService = newDataSyncService(ctx, svr.metaReplica, tsReplica, factory)
	svr.vectorStorage, err = factory.NewPersistentStorageChunkManager(ctx)
	if err != nil {
		panic(err)
	}
	svr.loader = newSegmentLoader(svr.metaReplica, etcdKV, svr.vectorStorage, factory)
	svr.etcdKV = etcdKV
	svr.etcdCli = etcdCli

	return svr
}

func newMessageStreamFactory() dependency.Factory {
	return dependency.NewDefaultFactory(true)
}

func startEmbedEtcdServer() (*embed.Etcd, error) {
	dir, err := ioutil.TempDir(os.TempDir(), "milvus_ut")
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(dir)
	config := embed.NewConfig()

	config.Dir = os.TempDir()
	config.LogLevel = "warn"
	config.LogOutputs = []string{"default"}
	u, err := url.Parse("http://localhost:8989")
	if err != nil {
		return nil, err
	}
	config.LCUrls = []url.URL{*u}
	u, err = url.Parse("http://localhost:8990")
	if err != nil {
		return nil, err
	}
	config.LPUrls = []url.URL{*u}

	return embed.StartEtcd(config)
}

func TestMain(m *testing.M) {
	setup()
	var err error
	rateCol, err = newRateCollector()
	if err != nil {
		panic("init test failed, err = " + err.Error())
	}
	// init embed etcd
	embedetcdServer, err = startEmbedEtcdServer()
	if err != nil {
		os.Exit(1)
	}
	defer embedetcdServer.Close()
	exitCode := m.Run()
	os.Exit(exitCode)
}

// NOTE: start pulsar and etcd before test
func TestQueryNode_Start(t *testing.T) {
	localNode := newQueryNodeMock()
	assert.Nil(t, localNode.queryHook)
	localNode.Start()
	<-localNode.queryNodeLoopCtx.Done()
	localNode.Stop()
}

func TestQueryNode_register(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	node, err := genSimpleQueryNode(ctx)
	assert.NoError(t, err)

	etcdcli, err := etcd.GetEtcdClient(
		Params.EtcdCfg.UseEmbedEtcd,
		Params.EtcdCfg.EtcdUseSSL,
		Params.EtcdCfg.Endpoints,
		Params.EtcdCfg.EtcdTLSCert,
		Params.EtcdCfg.EtcdTLSKey,
		Params.EtcdCfg.EtcdTLSCACert,
		Params.EtcdCfg.EtcdTLSMinVersion)
	assert.NoError(t, err)
	defer etcdcli.Close()
	node.SetEtcdClient(etcdcli)
	err = node.initSession()
	assert.NoError(t, err)

	node.session.TriggerKill = false
	err = node.Register()
	assert.NoError(t, err)
}

func TestQueryNode_init(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	node, err := genSimpleQueryNode(ctx)
	assert.NoError(t, err)
	defer node.Stop()

	etcdcli, err := etcd.GetEtcdClient(
		Params.EtcdCfg.UseEmbedEtcd,
		Params.EtcdCfg.EtcdUseSSL,
		Params.EtcdCfg.Endpoints,
		Params.EtcdCfg.EtcdTLSCert,
		Params.EtcdCfg.EtcdTLSKey,
		Params.EtcdCfg.EtcdTLSCACert,
		Params.EtcdCfg.EtcdTLSMinVersion)
	assert.NoError(t, err)
	defer etcdcli.Close()
	node.SetEtcdClient(etcdcli)
	err = node.Init()
	assert.Nil(t, err)
}

func genSimpleQueryNodeToTestWatchChangeInfo(ctx context.Context) (*QueryNode, error) {
	node, err := genSimpleQueryNode(ctx)
	if err != nil {
		return nil, err
	}

	/*
		err = node.queryService.addQueryCollection(defaultCollectionID)
		if err != nil {
			return nil, err
		}

		qc, err := node.queryService.getQueryCollection(defaultCollectionID)
		if err != nil {
			return nil, err
		}*/
	//qc.globalSegmentManager.addGlobalSegmentInfo(genSimpleSegmentInfo())
	return node, nil
}

func TestQueryNode_adjustByChangeInfo(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	t.Run("test cleanup segments", func(t *testing.T) {
		defer wg.Done()
		_, err := genSimpleQueryNodeToTestWatchChangeInfo(ctx)
		assert.NoError(t, err)

	})

	wg.Add(1)
	t.Run("test cleanup segments no segment", func(t *testing.T) {
		defer wg.Done()
		node, err := genSimpleQueryNodeToTestWatchChangeInfo(ctx)
		assert.NoError(t, err)

		node.metaReplica.removeSegment(defaultSegmentID, segmentTypeSealed)
		segmentChangeInfos := genSimpleChangeInfo()
		segmentChangeInfos.Infos[0].OnlineSegments = nil
		segmentChangeInfos.Infos[0].OfflineNodeID = Params.QueryNodeCfg.GetNodeID()

		/*
			qc, err := node.queryService.getQueryCollection(defaultCollectionID)
			assert.NoError(t, err)
			qc.globalSegmentManager.removeGlobalSealedSegmentInfo(defaultSegmentID)
		*/

	})
	wg.Wait()
}
