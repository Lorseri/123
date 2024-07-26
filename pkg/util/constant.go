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

package util

import (
	"strings"

	"github.com/milvus-io/milvus-proto/go-api/v2/commonpb"
	"github.com/milvus-io/milvus/pkg/common"
)

// Meta Prefix consts
const (
	MetaStoreTypeEtcd = "etcd"
	MetaStoreTypeTiKV = "tikv"

	SegmentMetaPrefix    = "queryCoord-segmentMeta"
	ChangeInfoMetaPrefix = "queryCoord-sealedSegmentChangeInfo"

	// FlushedSegmentPrefix TODO @cai.zhang: remove this
	FlushedSegmentPrefix = "flushed-segment"
	// HandoffSegmentPrefix TODO @cai.zhang: remove this
	HandoffSegmentPrefix = "querycoord-handoff"
	// SegmentReferPrefix TODO @cai.zhang: remove this
	SegmentReferPrefix = "segmentRefer"

	SegmentIndexPrefix = "segment-index"
	FieldIndexPrefix   = "field-index"

	HeaderAuthorize = "authorization"
	// HeaderSourceID identify requests from Milvus members and client requests
	HeaderSourceID = "sourceId"
	// MemberCredID id for Milvus members (data/index/query node/coord component)
	MemberCredID        = "@@milvus-member@@"
	CredentialSeperator = ":"
	UserRoot            = "root"
	PasswordHolder      = "___"
	DefaultTenant       = ""
	RoleAdmin           = "admin"
	RolePublic          = "public"
	DefaultDBName       = "default"
	DefaultDBID         = int64(1)
	NonDBID             = int64(0)
	InvalidDBID         = int64(-1)

	PrivilegeWord = "Privilege"
	AnyWord       = "*"

	IdentifierKey = "identifier"

	HeaderUserAgent = "user-agent"
	HeaderDBName    = "dbName"

	RoleConfigPrivileges = "privileges"
	RoleConfigObjectType = "object_type"
	RoleConfigObjectName = "object_name"
	RoleConfigDBName     = "db_name"
	RoleConfigPrivilege  = "privilege"

	MaxEtcdTxnNum = 128
)

const (
	// ParamsKeyToParse is the key of the param to build index.
	ParamsKeyToParse = common.IndexParamsKey
)

var (
	DefaultRoles = []string{RoleAdmin, RolePublic}
	BuiltinRoles = []string{}

	ObjectPrivileges = map[string][]string{
		commonpb.ObjectType_Collection.String(): {
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeLoad.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeRelease.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeCompaction.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeInsert.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDelete.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeUpsert.String()),

			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeGetStatistics.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeCreateIndex.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeIndexDetail.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDropIndex.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeSearch.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeFlush.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeQuery.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeLoadBalance.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeImport.String()),

			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeGetLoadingProgress.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeGetLoadState.String()),

			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeCreatePartition.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDropPartition.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeShowPartitions.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeHasPartition.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeGetFlushState.String()),
		},
		commonpb.ObjectType_Global.String(): {
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeAll.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeCreateCollection.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDropCollection.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDescribeCollection.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeShowCollections.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeRenameCollection.String()),

			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeCreateOwnership.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDropOwnership.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeSelectOwnership.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeManageOwnership.String()),

			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeCreateResourceGroup.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeUpdateResourceGroups.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDropResourceGroup.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDescribeResourceGroup.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeListResourceGroups.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeTransferReplica.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeTransferNode.String()),

			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeFlushAll.String()),

			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeCreateDatabase.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDropDatabase.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeListDatabases.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeAlterDatabase.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDescribeDatabase.String()),

			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeCreateAlias.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDropAlias.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeDescribeAlias.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeListAliases.String()),
		},
		commonpb.ObjectType_User.String(): {
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeUpdateUser.String()),
			MetaStore2API(commonpb.ObjectPrivilege_PrivilegeSelectUser.String()),
		},
	}

	RelatedPrivileges = map[string][]string{
		commonpb.ObjectPrivilege_PrivilegeLoad.String(): {
			commonpb.ObjectPrivilege_PrivilegeGetLoadState.String(),
			commonpb.ObjectPrivilege_PrivilegeGetLoadingProgress.String(),
		},
		commonpb.ObjectPrivilege_PrivilegeFlush.String(): {
			commonpb.ObjectPrivilege_PrivilegeGetFlushState.String(),
		},
	}
)

// StringSet convert array to map for conveniently check if the array contains an element
func StringSet(strings []string) map[string]struct{} {
	stringsMap := make(map[string]struct{})
	for _, str := range strings {
		stringsMap[str] = struct{}{}
	}
	return stringsMap
}

func StringList(stringMap map[string]struct{}) []string {
	strs := make([]string, 0, len(stringMap))
	for k := range stringMap {
		strs = append(strs, k)
	}
	return strs
}

// MetaStore2API convert meta-store's privilege name to api's
// example: PrivilegeAll -> All
func MetaStore2API(name string) string {
	return name[strings.Index(name, PrivilegeWord)+len(PrivilegeWord):]
}

func PrivilegeNameForAPI(name string) string {
	_, ok := commonpb.ObjectPrivilege_value[name]
	if !ok {
		return ""
	}
	return MetaStore2API(name)
}

func PrivilegeNameForMetastore(name string) string {
	dbPrivilege := PrivilegeWord + name
	_, ok := commonpb.ObjectPrivilege_value[dbPrivilege]
	if !ok {
		return ""
	}
	return dbPrivilege
}

func IsAnyWord(word string) bool {
	return word == AnyWord
}

func IsBuiltinRole(roleName string) bool {
	for _, builtinRole := range BuiltinRoles {
		if builtinRole == roleName {
			return true
		}
	}
	return false
}

func IsSameDatabase(dbName1, dbName2 string) bool {
	if dbName1 == dbName2 {
		return true
	}
	if dbName1 == "" {
		return dbName2 == DefaultDBName
	}
	if dbName2 == "" {
		return dbName1 == DefaultDBName
	}
	return false
}
