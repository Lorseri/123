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

package crossclusterrouting

import (
	"context"
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/milvus-io/milvus-proto/go-api/v2/milvuspb"
	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/proto/indexpb"
	"github.com/milvus-io/milvus/internal/proto/proxypb"
	"github.com/milvus-io/milvus/internal/proto/querypb"
	"github.com/milvus-io/milvus/internal/util/dependency"
	"github.com/milvus-io/milvus/internal/util/etcd"
	"github.com/milvus-io/milvus/internal/util/interceptor"
	"github.com/milvus-io/milvus/internal/util/paramtable"
	"github.com/milvus-io/milvus/internal/util/typeutil"

	grpcdatacoord "github.com/milvus-io/milvus/internal/distributed/datacoord"
	grpcdatacoordclient "github.com/milvus-io/milvus/internal/distributed/datacoord/client"
	grpcdatanode "github.com/milvus-io/milvus/internal/distributed/datanode"
	grpcdatanodeclient "github.com/milvus-io/milvus/internal/distributed/datanode/client"
	grpcindexcoord "github.com/milvus-io/milvus/internal/distributed/indexcoord"
	grpcindexcoordclient "github.com/milvus-io/milvus/internal/distributed/indexcoord/client"
	grpcindexnode "github.com/milvus-io/milvus/internal/distributed/indexnode"
	grpcindexnodeclient "github.com/milvus-io/milvus/internal/distributed/indexnode/client"
	grpcproxy "github.com/milvus-io/milvus/internal/distributed/proxy"
	grpcproxyclient "github.com/milvus-io/milvus/internal/distributed/proxy/client"
	grpcquerycoord "github.com/milvus-io/milvus/internal/distributed/querycoord"
	grpcquerycoordclient "github.com/milvus-io/milvus/internal/distributed/querycoord/client"
	grpcquerynode "github.com/milvus-io/milvus/internal/distributed/querynode"
	grpcquerynodeclient "github.com/milvus-io/milvus/internal/distributed/querynode/client"
	grpcrootcoord "github.com/milvus-io/milvus/internal/distributed/rootcoord"
	grpcrootcoordclient "github.com/milvus-io/milvus/internal/distributed/rootcoord/client"
)

type CrossClusterRoutingSuite struct {
	suite.Suite

	ctx    context.Context
	cancel context.CancelFunc

	factory dependency.Factory
	client  *clientv3.Client

	// clients
	rootCoordClient  *grpcrootcoordclient.Client
	proxyClient      *grpcproxyclient.Client
	dataCoordClient  *grpcdatacoordclient.Client
	indexCoordClient *grpcindexcoordclient.Client
	queryCoordClient *grpcquerycoordclient.Client
	dataNodeClient   *grpcdatanodeclient.Client
	queryNodeClient  *grpcquerynodeclient.Client
	indexNodeClient  *grpcindexnodeclient.Client

	// servers
	rootCoord  *grpcrootcoord.Server
	proxy      *grpcproxy.Server
	dataCoord  *grpcdatacoord.Server
	indexCoord *grpcindexcoord.Server
	queryCoord *grpcquerycoord.Server
	dataNode   *grpcdatanode.Server
	queryNode  *grpcquerynode.Server
	indexNode  *grpcindexnode.Server
}

func (s *CrossClusterRoutingSuite) SetupSuite() {
	s.ctx, s.cancel = context.WithTimeout(context.Background(), time.Second*180)
	rand.Seed(time.Now().UnixNano())

	s.factory = dependency.NewDefaultFactory(true)
}

func (s *CrossClusterRoutingSuite) TearDownSuite() {
}

func (s *CrossClusterRoutingSuite) SetupTest() {
	s.T().Logf("Setup test...")
	var err error

	// setup etcd client
	etcdConfig := interceptor.Params.EtcdCfg
	s.client, err = etcd.GetEtcdClient(
		etcdConfig.UseEmbedEtcd,
		etcdConfig.EtcdUseSSL,
		etcdConfig.Endpoints,
		etcdConfig.EtcdTLSCert,
		etcdConfig.EtcdTLSKey,
		etcdConfig.EtcdTLSCACert,
		etcdConfig.EtcdTLSMinVersion)
	s.NoError(err)

	// setup servers
	s.rootCoord, err = grpcrootcoord.NewServer(s.ctx, s.factory)
	s.NoError(err)
	err = s.rootCoord.Run()
	s.NoError(err)
	s.T().Logf("rootCoord server successfully started")

	s.dataCoord = grpcdatacoord.NewServer(s.ctx, s.factory)
	s.NotNil(s.dataCoord)
	err = s.dataCoord.Run()
	s.NoError(err)
	s.T().Logf("dataCoord server successfully started")

	s.indexCoord, err = grpcindexcoord.NewServer(s.ctx, s.factory)
	s.NoError(err)
	err = s.indexCoord.Run()
	s.NoError(err)
	s.T().Logf("indexCoord server successfully started")

	s.queryCoord, err = grpcquerycoord.NewServer(s.ctx, s.factory)
	s.NoError(err)
	err = s.queryCoord.Run()
	s.NoError(err)
	s.T().Logf("queryCoord server successfully started")

	s.proxy, err = grpcproxy.NewServer(s.ctx, s.factory)
	s.NoError(err)
	err = s.proxy.Run()
	s.NoError(err)
	s.T().Logf("proxy server successfully started")

	s.dataNode, err = grpcdatanode.NewServer(s.ctx, s.factory)
	s.NoError(err)
	err = s.dataNode.Run()
	s.NoError(err)
	s.T().Logf("dataNode server successfully started")

	s.queryNode, err = grpcquerynode.NewServer(s.ctx, s.factory)
	s.NoError(err)
	err = s.queryNode.Run()
	s.NoError(err)
	s.T().Logf("queryNode server successfully started")

	s.indexNode, err = grpcindexnode.NewServer(s.ctx, s.factory)
	s.NoError(err)
	err = s.indexNode.Run()
	s.NoError(err)
	s.T().Logf("indexNode server successfully started")

	metaRoot := interceptor.Params.EtcdCfg.MetaRootPath

	// setup clients
	s.rootCoordClient, err = grpcrootcoordclient.NewClient(s.ctx, metaRoot, s.client)
	s.NoError(err)
	s.dataCoordClient, err = grpcdatacoordclient.NewClient(s.ctx, metaRoot, s.client)
	s.NoError(err)
	s.indexCoordClient, err = grpcindexcoordclient.NewClient(s.ctx, metaRoot, s.client)
	s.NoError(err)
	s.queryCoordClient, err = grpcquerycoordclient.NewClient(s.ctx, metaRoot, s.client)
	s.NoError(err)

	var proxyGrpcServerParam paramtable.GrpcServerConfig
	proxyGrpcServerParam.InitOnce(typeutil.ProxyRole)
	s.proxyClient, err = grpcproxyclient.NewClient(s.ctx, proxyGrpcServerParam.GetInternalAddress())
	s.NoError(err)
	var dataNodeGrpcServerParam paramtable.GrpcServerConfig
	dataNodeGrpcServerParam.Init(typeutil.DataNodeRole)
	s.dataNodeClient, err = grpcdatanodeclient.NewClient(s.ctx, dataNodeGrpcServerParam.GetAddress())
	s.NoError(err)
	var queryNodeServerParam paramtable.GrpcServerConfig
	queryNodeServerParam.InitOnce(typeutil.QueryNodeRole)
	s.queryNodeClient, err = grpcquerynodeclient.NewClient(s.ctx, queryNodeServerParam.GetAddress())
	s.NoError(err)
	var indexNodeGrpcServerParam paramtable.GrpcServerConfig
	indexNodeGrpcServerParam.Init(typeutil.IndexNodeRole)
	s.indexNodeClient, err = grpcindexnodeclient.NewClient(s.ctx, indexNodeGrpcServerParam.GetAddress(), false)
	s.NoError(err)
}

func (s *CrossClusterRoutingSuite) TearDownTest() {
	err := s.rootCoord.Stop()
	s.NoError(err)
	err = s.proxy.Stop()
	s.NoError(err)
	err = s.dataCoord.Stop()
	s.NoError(err)
	err = s.indexCoord.Stop()
	s.NoError(err)
	err = s.queryCoord.Stop()
	s.NoError(err)
	err = s.dataNode.Stop()
	s.NoError(err)
	err = s.queryNode.Stop()
	s.NoError(err)
	err = s.indexNode.Stop()
	s.NoError(err)
	s.cancel()
}

func (s *CrossClusterRoutingSuite) TestCrossClusterRoutingSuite() {
	const (
		waitFor  = time.Second * 10
		duration = time.Millisecond * 10
	)

	go func() {
		for {
			select {
			case <-s.ctx.Done():
				return
			default:
				interceptor.Params.CommonCfg.SetClusterPrefix(fmt.Sprintf("%d", rand.Int()))
			}
		}
	}()

	// test rootCoord
	s.Eventually(func() bool {
		resp, err := s.rootCoordClient.ShowCollections(s.ctx, &milvuspb.ShowCollectionsRequest{})
		s.Suite.T().Logf("resp: %s, err: %s", resp, err)
		if err != nil {
			return strings.Contains(err.Error(), interceptor.ErrServiceUnavailable.Error())
		}
		return false
	}, waitFor, duration)

	// test dataCoord
	s.Eventually(func() bool {
		resp, err := s.dataCoordClient.GetRecoveryInfoV2(s.ctx, &datapb.GetRecoveryInfoRequestV2{})
		s.Suite.T().Logf("resp: %s, err: %s", resp, err)
		if err != nil {
			return strings.Contains(err.Error(), interceptor.ErrServiceUnavailable.Error())
		}
		return false
	}, waitFor, duration)

	// test indexCoord
	s.Eventually(func() bool {
		resp, err := s.indexCoordClient.CreateIndex(s.ctx, &indexpb.CreateIndexRequest{})
		s.Suite.T().Logf("resp: %s, err: %s", resp, err)
		if err != nil {
			return strings.Contains(err.Error(), interceptor.ErrServiceUnavailable.Error())
		}
		return false
	}, waitFor, duration)

	// test queryCoord
	s.Eventually(func() bool {
		resp, err := s.queryCoordClient.LoadCollection(s.ctx, &querypb.LoadCollectionRequest{})
		s.Suite.T().Logf("resp: %s, err: %s", resp, err)
		if err != nil {
			return strings.Contains(err.Error(), interceptor.ErrServiceUnavailable.Error())
		}
		return false
	}, waitFor, duration)

	// test proxy
	s.Eventually(func() bool {
		resp, err := s.proxyClient.InvalidateCollectionMetaCache(s.ctx, &proxypb.InvalidateCollMetaCacheRequest{})
		s.Suite.T().Logf("resp: %s, err: %s", resp, err)
		if err != nil {
			return strings.Contains(err.Error(), interceptor.ErrServiceUnavailable.Error())
		}
		return false
	}, waitFor, duration)

	// test dataNode
	s.Eventually(func() bool {
		resp, err := s.dataNodeClient.FlushSegments(s.ctx, &datapb.FlushSegmentsRequest{})
		s.Suite.T().Logf("resp: %s, err: %s", resp, err)
		if err != nil {
			return strings.Contains(err.Error(), interceptor.ErrServiceUnavailable.Error())
		}
		return false
	}, waitFor, duration)

	// test queryNode
	s.Eventually(func() bool {
		resp, err := s.queryNodeClient.Search(s.ctx, &querypb.SearchRequest{})
		s.Suite.T().Logf("resp: %s, err: %s", resp, err)
		if err != nil {
			return strings.Contains(err.Error(), interceptor.ErrServiceUnavailable.Error())
		}
		return false
	}, waitFor, duration)

	// test indexNode
	s.Eventually(func() bool {
		resp, err := s.indexNodeClient.CreateJob(s.ctx, &indexpb.CreateJobRequest{})
		s.Suite.T().Logf("resp: %s, err: %s", resp, err)
		if err != nil {
			return strings.Contains(err.Error(), interceptor.ErrServiceUnavailable.Error())
		}
		return false
	}, waitFor, duration)
}

func TestCrossClusterRoutingSuite(t *testing.T) {
	suite.Run(t, new(CrossClusterRoutingSuite))
}
