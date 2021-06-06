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

package queryservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/milvus-io/milvus/internal/log"
	"go.uber.org/zap"
	"sync"

	"github.com/milvus-io/milvus/internal/proto/commonpb"
	"github.com/milvus-io/milvus/internal/proto/internalpb"
	"github.com/milvus-io/milvus/internal/proto/querypb"
)

type queryNodeCluster struct {
	sync.RWMutex
	clusterMeta *meta
	nodes       map[int64]*queryNode
}

func newQueryNodeCluster(clusterMeta *meta) *queryNodeCluster {
	nodes := make(map[int64]*queryNode)
	return &queryNodeCluster{
		clusterMeta: clusterMeta,
		nodes:       nodes,
	}
}

func (c *queryNodeCluster) GetComponentInfos(ctx context.Context) []*internalpb.ComponentInfo {
	c.RLock()
	defer c.RUnlock()
	subComponentInfos := make([]*internalpb.ComponentInfo, 0)
	for nodeID, node := range c.nodes {
		componentStates, err := node.client.GetComponentStates(ctx)
		if err != nil {
			subComponentInfos = append(subComponentInfos, &internalpb.ComponentInfo{
				NodeID:    nodeID,
				StateCode: internalpb.StateCode_Abnormal,
			})
			continue
		}
		subComponentInfos = append(subComponentInfos, componentStates.State)
	}

	return subComponentInfos
}

func (c *queryNodeCluster) LoadSegments(ctx context.Context, nodeID int64, in *querypb.LoadSegmentsRequest) (*commonpb.Status, error) {
	c.Lock()
	defer c.Unlock()
	if node, ok := c.nodes[nodeID]; ok {
		//TODO::etcd
		for _, info := range in.Infos {
			segmentID := info.SegmentID
			if info, ok := c.clusterMeta.segmentInfos[segmentID]; ok {
				info.SegmentState = querypb.SegmentState_sealing
			}
			segmentInfo := &querypb.SegmentInfo{
				SegmentID:    segmentID,
				CollectionID: in.CollectionID,
				PartitionID:  in.PartitionID,
				NodeID:       nodeID,
				SegmentState: querypb.SegmentState_sealing,
			}
			c.clusterMeta.segmentInfos[segmentID] = segmentInfo
		}
		status, err := node.client.LoadSegments(ctx, in)
		if err == nil && status.ErrorCode == commonpb.ErrorCode_Success {
			if !c.clusterMeta.hasCollection(in.CollectionID) {
				c.clusterMeta.addCollection(in.CollectionID, in.Schema)
			}
			c.clusterMeta.addPartition(in.CollectionID, in.PartitionID)

			if !node.hasCollection(in.CollectionID) {
				node.addCollection(in.CollectionID, in.Schema)
			}
			node.addPartition(in.CollectionID, in.PartitionID)
			return status, err
		}
		for _, info := range in.Infos {
			segmentID := info.SegmentID
			c.clusterMeta.deleteSegmentInfoByID(segmentID)
		}
		return status, err
	}
	return nil, errors.New("Can't find query node by nodeID ")
}

func (c *queryNodeCluster) ReleaseSegments(ctx context.Context, nodeID int64, in *querypb.ReleaseSegmentsRequest)(*commonpb.Status, error) {
	c.Lock()
	defer c.Unlock()

	if node, ok := c.nodes[nodeID]; ok {
		status, err := node.client.ReleaseSegments(ctx, in)
		if err == nil && status.ErrorCode == commonpb.ErrorCode_Success {
			for _, segmentID := range in.SegmentIDs {
				c.clusterMeta.deleteSegmentInfoByID(segmentID)
			}
		}
		return status, err
	}

	return nil, errors.New("Can't find query node by nodeID ")
}

func (c *queryNodeCluster) WatchDmChannels(ctx context.Context, nodeID int64, in *querypb.WatchDmChannelsRequest) (*commonpb.Status, error) {
	c.Lock()
	defer c.Unlock()
	if node, ok := c.nodes[nodeID]; ok {
		status, err := node.client.WatchDmChannels(ctx, in)
		if err == nil && status.ErrorCode == commonpb.ErrorCode_Success {
			collectionID := in.CollectionID
			if !c.clusterMeta.hasCollection(collectionID) {
				c.clusterMeta.addCollection(collectionID, in.Schema)
			}
			c.clusterMeta.addDmChannel(collectionID, nodeID, in.ChannelIDs)
			if !node.hasCollection(collectionID) {
				node.addCollection(collectionID, in.Schema)
			}
			node.addDmChannel(collectionID, in.ChannelIDs)
		}
		return status, err
	}
	return nil, errors.New("Can't find query node by nodeID ")
}

func (c *queryNodeCluster) AddQueryChannel(ctx context.Context, nodeID int64, in *querypb.AddQueryChannelRequest) (*commonpb.Status, error) {
	c.Lock()
	defer c.Unlock()
	if node, ok := c.nodes[nodeID]; ok {
		status, err := node.client.AddQueryChannel(ctx, in)
		if err == nil && status.ErrorCode == commonpb.ErrorCode_Success {
			collectionID := in.CollectionID
			if queryChannelInfo, ok := c.clusterMeta.queryChannelInfos[collectionID]; ok {
				node.addQueryChannel(collectionID, queryChannelInfo)
				return status, err
			}
			log.Error("queryChannel for collection not assigned", zap.Int64("collectionID", collectionID))
		}
		return status, err
	}

	return nil, errors.New("can't find query node by nodeID")
}
func (c *queryNodeCluster) removeQueryChannel(ctx context.Context, nodeID int64, in *querypb.RemoveQueryChannelRequest) (*commonpb.Status, error) {
	c.Lock()
	defer c.Unlock()

	if node, ok := c.nodes[nodeID]; ok {
		status, err := node.client.RemoveQueryChannel(ctx, in)
		if err == nil && status.ErrorCode == commonpb.ErrorCode_Success {
			collectionID := in.CollectionID
			if _, ok = node.watchedQueryChannels[collectionID]; ok {
				node.removeQueryChannel(collectionID)
				return status, err
			}
			log.Error("queryChannel for collection not watched", zap.Int64("collectionID", collectionID))
		}
		return status, err
	}

	return nil, errors.New("can't find query node by nodeID")
}

func (c *queryNodeCluster) releaseCollection(ctx context.Context, nodeID int64, in *querypb.ReleaseCollectionRequest) (*commonpb.Status, error) {
	c.Lock()
	defer c.Unlock()

	if node, ok := c.nodes[nodeID]; ok {
		status, err := node.client.ReleaseCollection(ctx, in)
		if err == nil && status.ErrorCode == commonpb.ErrorCode_Success {
			node.releaseCollection(in.CollectionID)
			c.clusterMeta.releaseCollection(in.CollectionID)
		}
		return status, err
	}

	return nil, errors.New("can't find query node by nodeID")
}

func (c *queryNodeCluster) releasePartitions(ctx context.Context, nodeID int64, in *querypb.ReleasePartitionsRequest) (*commonpb.Status, error) {
	c.Lock()
	defer c.Unlock()

	if node, ok := c.nodes[nodeID]; ok {
		status, err := node.client.ReleasePartitions(ctx, in)
		if err == nil && status.ErrorCode == commonpb.ErrorCode_Success {
			for _, partitionID := range in.PartitionIDs {
				node.releasePartition(in.CollectionID, partitionID)
				c.clusterMeta.releasePartition(in.CollectionID, partitionID)
			}
		}
		return status, err
	}

	return nil, errors.New("can't find query node by nodeID")
}

func (c *queryNodeCluster) getNumDmChannels(nodeID int64) (int, error) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.nodes[nodeID]; !ok {
		return 0, errors.New("Can't find query node by nodeID ")
	}

	numChannel := 0
	for _, info := range c.clusterMeta.collectionInfos {
		for _, channelInfo := range info.ChannelInfos {
			if channelInfo.NodeIDLoaded == nodeID {
				numChannel++
			}
		}
	}
	return numChannel, nil
}

func (c *queryNodeCluster) getNumSegments(nodeID int64) (int, error) {
	c.Lock()
	defer c.Unlock()

	if _, ok := c.nodes[nodeID]; !ok {
		return 0, errors.New("Can't find query node by nodeID ")
	}

	numSegment := 0
	for _, info := range c.clusterMeta.segmentInfos {
		if info.NodeID == nodeID {
			numSegment++
		}
	}
	return numSegment, nil
}

func (c *queryNodeCluster) RegisterNode(ip string, port int64, id UniqueID) error {
	node, err := newQueryNode(ip, port, id)
	if err != nil {
		return err
	}
	c.Lock()
	defer c.Unlock()
	if _, ok := c.nodes[id]; !ok {
		c.nodes[id] = node
		return nil
	}
	return errors.New(fmt.Sprintf("node %d alredy exists in cluster", id))
}
