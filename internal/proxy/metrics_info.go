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

package proxy

import (
	"context"
	"os"

	"github.com/milvus-io/milvus/internal/proto/commonpb"

	"github.com/milvus-io/milvus/internal/util/typeutil"

	"github.com/milvus-io/milvus/internal/util/metricsinfo"

	"github.com/milvus-io/milvus/internal/proto/milvuspb"
)

func getSystemInfoMetrics(
	ctx context.Context,
	request *milvuspb.GetMetricsRequest,
	node *Proxy,
) (*milvuspb.GetMetricsResponse, error) {

	var err error

	systemTopology := metricsinfo.SystemTopology{
		NodesInfo: make([]metricsinfo.SystemTopologyNode, 0),
	}

	identifierMap := make(map[string]int)

	proxyRoleName := metricsinfo.ConstructComponentName(typeutil.ProxyRole, Params.ProxyID)
	identifierMap[proxyRoleName] = getUniqueIntGeneratorIns().get()
	proxyTopologyNode := metricsinfo.SystemTopologyNode{
		Identifier: identifierMap[proxyRoleName],
		Connected:  make([]metricsinfo.ConnectionEdge, 0),
		Infos: &metricsinfo.ProxyInfos{
			BaseComponentInfos: metricsinfo.BaseComponentInfos{
				HasError:    false,
				ErrorReason: "",
				Name:        proxyRoleName,
				HardwareInfos: metricsinfo.HardwareMetrics{
					IP:           node.session.Address,
					CPUCoreCount: metricsinfo.GetCPUCoreCount(false),
					CPUCoreUsage: metricsinfo.GetCPUUsage(),
					Memory:       metricsinfo.GetMemoryCount(),
					MemoryUsage:  metricsinfo.GetUsedMemoryCount(),
					Disk:         metricsinfo.GetDiskCount(),
					DiskUsage:    metricsinfo.GetDiskUsage(),
				},
				SystemInfo: metricsinfo.DeployMetrics{
					SystemVersion: os.Getenv(metricsinfo.GitCommitEnvKey),
					DeployMode:    os.Getenv(metricsinfo.DeployModeEnvKey),
				},
				// TODO(dragondriver): CreatedTime & UpdatedTime, easy but time-costing
				Type: typeutil.ProxyRole,
			},
			SystemConfigurations: metricsinfo.ProxyConfiguration{
				DefaultPartitionName: Params.DefaultPartitionName,
				DefaultIndexName:     Params.DefaultIndexName,
			},
		},
	}

	queryCoordResp, queryCoordErr := node.queryCoord.GetMetrics(ctx, request)
	queryCoordRoleName := ""
	if queryCoordErr == nil && queryCoordResp != nil {
		queryCoordRoleName = queryCoordResp.ComponentName
		identifierMap[queryCoordRoleName] = getUniqueIntGeneratorIns().get()
	}

	dataCoordResp, dataCoordErr := node.dataCoord.GetMetrics(ctx, request)
	dataCoordRoleName := ""
	if dataCoordErr == nil && dataCoordResp != nil {
		dataCoordRoleName = dataCoordResp.ComponentName
		identifierMap[dataCoordRoleName] = getUniqueIntGeneratorIns().get()
	}

	indexCoordResp, indexCoordErr := node.indexCoord.GetMetrics(ctx, request)
	indexCoordRoleName := ""
	if indexCoordErr == nil && indexCoordResp != nil {
		indexCoordRoleName = indexCoordResp.ComponentName
		identifierMap[indexCoordRoleName] = getUniqueIntGeneratorIns().get()
	}

	rootCoordResp, rootCoordErr := node.rootCoord.GetMetrics(ctx, request)
	rootCoordRoleName := ""
	if rootCoordErr == nil && rootCoordResp != nil {
		rootCoordRoleName = rootCoordResp.ComponentName
		identifierMap[rootCoordRoleName] = getUniqueIntGeneratorIns().get()
	}

	if queryCoordErr == nil && queryCoordResp != nil {
		proxyTopologyNode.Connected = append(proxyTopologyNode.Connected, metricsinfo.ConnectionEdge{
			ConnectedIdentifier: identifierMap[queryCoordRoleName],
			Type:                metricsinfo.Forward,
			TargetType:          typeutil.QueryCoordRole,
		})

		queryCoordTopology := metricsinfo.QueryCoordTopology{}
		err = metricsinfo.UnmarshalTopology(queryCoordResp.Response, &queryCoordTopology)
		if err == nil {
			// query coord in system topology graph
			queryCoordTopologyNode := metricsinfo.SystemTopologyNode{
				Identifier: identifierMap[queryCoordRoleName],
				Connected:  make([]metricsinfo.ConnectionEdge, 0),
				Infos:      &queryCoordTopology.Cluster.Self,
			}

			// fill connection edge, a little trick here
			for _, edge := range queryCoordTopology.Connections.ConnectedComponents {
				switch edge.TargetType {
				case typeutil.RootCoordRole:
					if rootCoordErr == nil && rootCoordResp != nil {
						queryCoordTopologyNode.Connected = append(queryCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
							ConnectedIdentifier: identifierMap[rootCoordRoleName],
							Type:                metricsinfo.Forward,
							TargetType:          typeutil.RootCoordRole,
						})
					}
				case typeutil.DataCoordRole:
					if dataCoordErr == nil && dataCoordResp != nil {
						queryCoordTopologyNode.Connected = append(queryCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
							ConnectedIdentifier: identifierMap[dataCoordRoleName],
							Type:                metricsinfo.Forward,
							TargetType:          typeutil.DataCoordRole,
						})
					}
				case typeutil.IndexCoordRole:
					if indexCoordErr == nil && indexCoordResp != nil {
						queryCoordTopologyNode.Connected = append(queryCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
							ConnectedIdentifier: identifierMap[indexCoordRoleName],
							Type:                metricsinfo.Forward,
							TargetType:          typeutil.IndexCoordRole,
						})
					}
				case typeutil.QueryCoordRole:
					queryCoordTopologyNode.Connected = append(queryCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
						ConnectedIdentifier: identifierMap[queryCoordRoleName],
						Type:                metricsinfo.Forward,
						TargetType:          typeutil.QueryCoordRole,
					})
				}
			}

			// add query nodes to system topology graph
			for _, queryNode := range queryCoordTopology.Cluster.ConnectedNodes {
				identifier := getUniqueIntGeneratorIns().get()
				identifierMap[queryNode.Name] = identifier
				queryNodeTopologyNode := metricsinfo.SystemTopologyNode{
					Identifier: identifier,
					Connected:  nil,
					Infos:      &queryNode,
				}
				systemTopology.NodesInfo = append(systemTopology.NodesInfo, queryNodeTopologyNode)
				queryCoordTopologyNode.Connected = append(queryCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
					ConnectedIdentifier: identifier,
					Type:                metricsinfo.CoordConnectToNode,
					TargetType:          typeutil.QueryNodeRole,
				})
			}

			// add query coord to system topology graph
			systemTopology.NodesInfo = append(systemTopology.NodesInfo, queryCoordTopologyNode)
		}
	}

	if dataCoordErr == nil && dataCoordResp != nil {
		proxyTopologyNode.Connected = append(proxyTopologyNode.Connected, metricsinfo.ConnectionEdge{
			ConnectedIdentifier: identifierMap[dataCoordRoleName],
			Type:                metricsinfo.Forward,
			TargetType:          typeutil.DataCoordRole,
		})

		dataCoordTopology := metricsinfo.DataCoordTopology{}
		err = metricsinfo.UnmarshalTopology(dataCoordResp.Response, &dataCoordTopology)
		if err == nil {
			// data coord in system topology graph
			dataCoordTopologyNode := metricsinfo.SystemTopologyNode{
				Identifier: identifierMap[dataCoordRoleName],
				Connected:  make([]metricsinfo.ConnectionEdge, 0),
				Infos:      &dataCoordTopology.Cluster.Self,
			}

			// fill connection edge, a little trick here
			for _, edge := range dataCoordTopology.Connections.ConnectedComponents {
				switch edge.TargetType {
				case typeutil.RootCoordRole:
					if rootCoordErr == nil && rootCoordResp != nil {
						dataCoordTopologyNode.Connected = append(dataCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
							ConnectedIdentifier: identifierMap[rootCoordRoleName],
							Type:                metricsinfo.Forward,
							TargetType:          typeutil.RootCoordRole,
						})
					}
				case typeutil.DataCoordRole:
					dataCoordTopologyNode.Connected = append(dataCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
						ConnectedIdentifier: identifierMap[dataCoordRoleName],
						Type:                metricsinfo.Forward,
						TargetType:          typeutil.DataCoordRole,
					})
				case typeutil.IndexCoordRole:
					if indexCoordErr == nil && indexCoordResp != nil {
						dataCoordTopologyNode.Connected = append(dataCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
							ConnectedIdentifier: identifierMap[indexCoordRoleName],
							Type:                metricsinfo.Forward,
							TargetType:          typeutil.IndexCoordRole,
						})
					}
				case typeutil.QueryCoordRole:
					if queryCoordErr == nil && queryCoordResp != nil {
						dataCoordTopologyNode.Connected = append(dataCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
							ConnectedIdentifier: identifierMap[queryCoordRoleName],
							Type:                metricsinfo.Forward,
							TargetType:          typeutil.QueryCoordRole,
						})
					}
				}
			}

			// add data nodes to system topology graph
			for _, dataNode := range dataCoordTopology.Cluster.ConnectedNodes {
				identifier := getUniqueIntGeneratorIns().get()
				identifierMap[dataNode.Name] = identifier
				dataNodeTopologyNode := metricsinfo.SystemTopologyNode{
					Identifier: identifier,
					Connected:  nil,
					Infos:      &dataNode,
				}
				systemTopology.NodesInfo = append(systemTopology.NodesInfo, dataNodeTopologyNode)
				dataCoordTopologyNode.Connected = append(dataCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
					ConnectedIdentifier: identifier,
					Type:                metricsinfo.CoordConnectToNode,
					TargetType:          typeutil.DataNodeRole,
				})
			}

			// add data coord to system topology graph
			systemTopology.NodesInfo = append(systemTopology.NodesInfo, dataCoordTopologyNode)
		}
	}

	if indexCoordErr == nil && indexCoordResp != nil {
		proxyTopologyNode.Connected = append(proxyTopologyNode.Connected, metricsinfo.ConnectionEdge{
			ConnectedIdentifier: identifierMap[indexCoordRoleName],
			Type:                metricsinfo.Forward,
			TargetType:          typeutil.IndexCoordRole,
		})

		indexCoordTopology := metricsinfo.IndexCoordTopology{}
		err = metricsinfo.UnmarshalTopology(indexCoordResp.Response, &indexCoordTopology)
		if err == nil {
			// index coord in system topology graph
			indexCoordTopologyNode := metricsinfo.SystemTopologyNode{
				Identifier: identifierMap[indexCoordRoleName],
				Connected:  make([]metricsinfo.ConnectionEdge, 0),
				Infos:      &indexCoordTopology.Cluster.Self,
			}

			// fill connection edge, a little trick here
			for _, edge := range indexCoordTopology.Connections.ConnectedComponents {
				switch edge.TargetType {
				case typeutil.RootCoordRole:
					if rootCoordErr == nil && rootCoordResp != nil {
						indexCoordTopologyNode.Connected = append(indexCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
							ConnectedIdentifier: identifierMap[rootCoordRoleName],
							Type:                metricsinfo.Forward,
							TargetType:          typeutil.RootCoordRole,
						})
					}
				case typeutil.DataCoordRole:
					if dataCoordErr == nil && dataCoordResp != nil {
						indexCoordTopologyNode.Connected = append(indexCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
							ConnectedIdentifier: identifierMap[dataCoordRoleName],
							Type:                metricsinfo.Forward,
							TargetType:          typeutil.DataCoordRole,
						})
					}
				case typeutil.IndexCoordRole:
					indexCoordTopologyNode.Connected = append(indexCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
						ConnectedIdentifier: identifierMap[indexCoordRoleName],
						Type:                metricsinfo.Forward,
						TargetType:          typeutil.IndexCoordRole,
					})
				case typeutil.QueryCoordRole:
					if queryCoordErr == nil && queryCoordResp != nil {
						indexCoordTopologyNode.Connected = append(indexCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
							ConnectedIdentifier: identifierMap[queryCoordRoleName],
							Type:                metricsinfo.Forward,
							TargetType:          typeutil.QueryCoordRole,
						})
					}
				}
			}

			// add index nodes to system topology graph
			for _, indexNode := range indexCoordTopology.Cluster.ConnectedNodes {
				identifier := getUniqueIntGeneratorIns().get()
				identifierMap[indexNode.Name] = identifier
				indexNodeTopologyNode := metricsinfo.SystemTopologyNode{
					Identifier: identifier,
					Connected:  nil,
					Infos:      &indexNode,
				}
				systemTopology.NodesInfo = append(systemTopology.NodesInfo, indexNodeTopologyNode)
				indexCoordTopologyNode.Connected = append(indexCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
					ConnectedIdentifier: identifier,
					Type:                metricsinfo.CoordConnectToNode,
					TargetType:          typeutil.IndexNodeRole,
				})
			}

			// add index coord to system topology graph
			systemTopology.NodesInfo = append(systemTopology.NodesInfo, indexCoordTopologyNode)
		}
	}

	if rootCoordErr == nil && rootCoordResp != nil {
		proxyTopologyNode.Connected = append(proxyTopologyNode.Connected, metricsinfo.ConnectionEdge{
			ConnectedIdentifier: identifierMap[rootCoordRoleName],
			Type:                metricsinfo.Forward,
			TargetType:          typeutil.RootCoordRole,
		})

		rootCoordTopology := metricsinfo.RootCoordTopology{}
		err = metricsinfo.UnmarshalTopology(rootCoordResp.Response, &rootCoordTopology)
		if err == nil {
			// root coord in system topology graph
			rootCoordTopologyNode := metricsinfo.SystemTopologyNode{
				Identifier: identifierMap[rootCoordRoleName],
				Connected:  make([]metricsinfo.ConnectionEdge, 0),
				Infos:      &rootCoordTopology.Self,
			}

			// fill connection edge, a little trick here
			for _, edge := range rootCoordTopology.Connections.ConnectedComponents {
				switch edge.TargetType {
				case typeutil.RootCoordRole:
					rootCoordTopologyNode.Connected = append(rootCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
						ConnectedIdentifier: identifierMap[rootCoordRoleName],
						Type:                metricsinfo.Forward,
						TargetType:          typeutil.RootCoordRole,
					})
				case typeutil.DataCoordRole:
					if dataCoordErr == nil && dataCoordResp != nil {
						rootCoordTopologyNode.Connected = append(rootCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
							ConnectedIdentifier: identifierMap[dataCoordRoleName],
							Type:                metricsinfo.Forward,
							TargetType:          typeutil.DataCoordRole,
						})
					}
				case typeutil.IndexCoordRole:
					rootCoordTopologyNode.Connected = append(rootCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
						ConnectedIdentifier: identifierMap[indexCoordRoleName],
						Type:                metricsinfo.Forward,
						TargetType:          typeutil.IndexCoordRole,
					})
				case typeutil.QueryCoordRole:
					if queryCoordErr == nil && queryCoordResp != nil {
						rootCoordTopologyNode.Connected = append(rootCoordTopologyNode.Connected, metricsinfo.ConnectionEdge{
							ConnectedIdentifier: identifierMap[queryCoordRoleName],
							Type:                metricsinfo.Forward,
							TargetType:          typeutil.QueryCoordRole,
						})
					}
				}
			}

			// add root coord to system topology graph
			systemTopology.NodesInfo = append(systemTopology.NodesInfo, rootCoordTopologyNode)
		}
	}

	// add proxy to system topology graph
	systemTopology.NodesInfo = append(systemTopology.NodesInfo, proxyTopologyNode)

	resp, err := metricsinfo.MarshalTopology(systemTopology)
	if err != nil {
		return &milvuspb.GetMetricsResponse{
			Status: &commonpb.Status{
				ErrorCode: commonpb.ErrorCode_UnexpectedError,
				Reason:    err.Error(),
			},
			Response:      "",
			ComponentName: metricsinfo.ConstructComponentName(typeutil.ProxyRole, Params.ProxyID),
		}, nil
	}

	return &milvuspb.GetMetricsResponse{
		Status: &commonpb.Status{
			ErrorCode: commonpb.ErrorCode_Success,
			Reason:    "",
		},
		Response:      resp,
		ComponentName: metricsinfo.ConstructComponentName(typeutil.ProxyRole, Params.ProxyID),
	}, nil
}
