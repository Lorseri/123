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

package utils

import (
	"fmt"

	"github.com/cockroachdb/errors"
	"go.uber.org/multierr"
	"go.uber.org/zap"

	"github.com/milvus-io/milvus/internal/proto/datapb"
	"github.com/milvus-io/milvus/internal/proto/querypb"
	"github.com/milvus-io/milvus/internal/querycoordv2/meta"
	"github.com/milvus-io/milvus/internal/querycoordv2/session"
	"github.com/milvus-io/milvus/pkg/log"
	"github.com/milvus-io/milvus/pkg/util/merr"
)

func GetReadableShardLeaders(metaMgr *meta.Meta,
	distMgr *meta.DistributionManager,
	targetMgr *meta.TargetManager,
	nodeMgr *session.NodeManager,
	collectionID int64,
) (map[string]*querypb.ShardLeadersList, error) {
	ret := make(map[string]*querypb.ShardLeadersList)
	percentage := metaMgr.CollectionManager.CalculateLoadPercentage(collectionID)
	if percentage < 0 {
		err := merr.WrapErrCollectionNotLoaded(collectionID)
		log.Warn("failed to GetShardLeaders", zap.Error(err))
		return nil, err
	}
	collection := metaMgr.CollectionManager.GetCollection(collectionID)
	if collection.GetStatus() == querypb.LoadStatus_Loaded {
		// when collection is loaded, regard collection as readable, set percentage == 100
		percentage = 100
	}
	if percentage < 100 {
		err := merr.WrapErrCollectionNotFullyLoaded(collectionID)
		msg := fmt.Sprintf("collection %v is not fully loaded", collectionID)
		log.Warn(msg)
		return nil, err
	}

	channels := targetMgr.GetDmChannelsByCollection(collectionID, meta.CurrentTarget)
	if len(channels) == 0 {
		msg := "failed to get channels"
		err := merr.WrapErrCollectionNotLoaded(collectionID)
		log.Warn(msg, zap.Error(err))
		return nil, err
	}

	currentTargets := targetMgr.GetSealedSegmentsByCollection(collectionID, meta.CurrentTarget)
	for _, channel := range channels {
		log := log.With(zap.String("channel", channel.GetChannelName()))

		leaders := distMgr.LeaderViewManager.GetLeadersByShard(channel.GetChannelName())

		readableLeaders := make(map[int64]*meta.LeaderView)

		var channelErr error
		if len(leaders) == 0 {
			channelErr = merr.WrapErrChannelLack("channel not subscribed")
		}

		for _, leader := range leaders {
			if err := CheckLeaderAvailable(nodeMgr, leader, currentTargets); err != nil {
				multierr.AppendInto(&channelErr, err)
				continue
			}

			readableLeaders[leader.ID] = leader
		}

		if len(readableLeaders) == 0 {
			msg := fmt.Sprintf("channel %s is not available in any replica", channel.GetChannelName())
			log.Warn(msg, zap.Error(channelErr))
			err := errors.Wrap(merr.WrapErrChannelNotAvailable(channel.GetChannelName()), channelErr.Error())
			return nil, err
		}

		readableLeaders = filterDupLeaders(metaMgr.ReplicaManager, readableLeaders)
		ids := make([]int64, 0, len(leaders))
		addrs := make([]string, 0, len(leaders))
		for _, leader := range readableLeaders {
			info := nodeMgr.Get(leader.ID)
			ids = append(ids, info.ID())
			addrs = append(addrs, info.Addr())
		}

		ret[channel.GetChannelName()] = &querypb.ShardLeadersList{
			ChannelName: channel.GetChannelName(),
			NodeIds:     ids,
			NodeAddrs:   addrs,
		}
	}

	return ret, nil
}

func filterDupLeaders(replicaManager *meta.ReplicaManager, leaders map[int64]*meta.LeaderView) map[int64]*meta.LeaderView {
	type leaderID struct {
		ReplicaID int64
		Shard     string
	}

	newLeaders := make(map[leaderID]*meta.LeaderView)
	for _, view := range leaders {
		replica := replicaManager.GetByCollectionAndNode(view.CollectionID, view.ID)
		if replica == nil {
			continue
		}

		id := leaderID{replica.GetID(), view.Channel}
		if old, ok := newLeaders[id]; ok && old.Version > view.Version {
			continue
		}

		newLeaders[id] = view
	}

	result := make(map[int64]*meta.LeaderView)
	for _, v := range newLeaders {
		result[v.ID] = v
	}
	return result
}

// In a replica, a shard is available, if and only if:
// 1. The leader is online
// 2. All QueryNodes in the distribution are online
// 3. The last heartbeat response time is within HeartbeatAvailableInterval for all QueryNodes(include leader) in the distribution
// 4. All segments of the shard in target should be in the distribution
func CheckLeaderAvailable(nodeMgr *session.NodeManager, leader *meta.LeaderView, currentTargets map[int64]*datapb.SegmentInfo) error {
	log := log.With(zap.Int64("leaderID", leader.ID))
	info := nodeMgr.Get(leader.ID)

	// Check whether leader is online
	err := CheckNodeAvailable(leader.ID, info)
	if err != nil {
		log.Info("leader is not available", zap.Error(err))
		return fmt.Errorf("leader not available: %w", err)
	}

	for id, version := range leader.Segments {
		info := nodeMgr.Get(version.GetNodeID())
		err = CheckNodeAvailable(version.GetNodeID(), info)
		if err != nil {
			log.Info("leader is not available due to QueryNode unavailable",
				zap.Int64("segmentID", id),
				zap.Error(err))
			return err
		}
	}

	// Check whether segments are fully loaded
	for segmentID, info := range currentTargets {
		if info.GetInsertChannel() != leader.Channel {
			continue
		}

		_, exist := leader.Segments[segmentID]
		if !exist {
			log.Info("leader is not available due to lack of segment", zap.Int64("segmentID", segmentID))
			return merr.WrapErrSegmentLack(segmentID)
		}
	}
	return nil
}

func CheckNodeAvailable(nodeID int64, info *session.NodeInfo) error {
	if info == nil {
		return merr.WrapErrNodeOffline(nodeID)
	}
	return nil
}
