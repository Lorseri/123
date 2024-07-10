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

#pragma once

#include <string>

#include "common/Types.h"

namespace milvus::storage {

using DataType = milvus::DataType;
using Timestamp = milvus::Timestamp;

const int32_t MAGIC_NUM = 0xfffabc;

enum StorageType {
    None = 0,
    Memory = 1,
    LocalDisk = 2,
    Remote = 3,
};

enum class FileType : int32_t {
    InsertBinlog = 0,     // InsertBinlog FileType for insert data
    DeleteBinlog = 1,     // DeleteBinlog FileType for delete data
    DDLBinlog = 2,        // DDLBinlog FileType for DDL
    IndexFileBinlog = 3,  // IndexFileBinlog FileType for index

    KWInsertBinlog =
        100,  // InsertBinlog FileType for insert data prepared for knowhere
    KWIndexFileBinlog =
        101,  // IndexFileBinlog FileType for index generated by knowhere
};

enum class EventType : int8_t {
    DescriptorEvent = 0,
    InsertEvent = 1,
    DeleteEvent = 2,
    CreateCollectionEvent = 3,
    DropCollectionEvent = 4,
    CreatePartitionEvent = 5,
    DropPartitionEvent = 6,
    IndexFileEvent = 7,
    EventTypeEnd = 8,
};

// segment/field meta information corresponding to binlog file data
struct FieldDataMeta {
    int64_t collection_id;
    int64_t partition_id;
    int64_t segment_id;
    int64_t field_id;
};

enum CodecType {
    InvalidCodecType = 0,
    InsertDataType = 1,
    IndexDataType = 2,
};

// index meta information corresponding to index file data
struct IndexMeta {
    int64_t segment_id;
    int64_t field_id;
    int64_t build_id;
    int64_t index_version;
    std::string key;
};

struct StorageConfig {
    std::string address = "localhost:9000";
    std::string bucket_name = "a-bucket";
    std::string access_key_id = "minioadmin";
    std::string access_key_value = "minioadmin";
    std::string root_path = "files";
    std::string storage_type = "minio";
    std::string cloud_provider = "aws";
    std::string iam_endpoint = "";
    std::string log_level = "warn";
    std::string region = "";
    bool useSSL = false;
    std::string sslCACert = "";
    bool useIAM = false;
    bool useVirtualHost = false;
    int64_t requestTimeoutMs = 3000;
    bool useCollectionIdIndexPath = false;

    std::string
    ToString() const {
        std::stringstream ss;
        ss << "[address=" << address << ", bucket_name=" << bucket_name
           << ", root_path=" << root_path << ", storage_type=" << storage_type
           << ", cloud_provider=" << cloud_provider
           << ", iam_endpoint=" << iam_endpoint << ", log_level=" << log_level
           << ", region=" << region << ", useSSL=" << std::boolalpha << useSSL
           << ", sslCACert=" << sslCACert.size()  // only print cert length
           << ", useIAM=" << std::boolalpha << useIAM
           << ", useVirtualHost=" << std::boolalpha << useVirtualHost
           << ", requestTimeoutMs=" << requestTimeoutMs
           << ", useCollectionIdIndexPath=" << std::boolalpha << useCollectionIdIndexPath

        return ss.str();
    }
};

}  // namespace milvus::storage

template <>
struct fmt::formatter<milvus::storage::EventType> : formatter<string_view> {
    auto
    format(milvus::storage::EventType c, format_context& ctx) const {
        string_view name = "unknown";
        switch (c) {
            case milvus::storage::EventType::DescriptorEvent:
                name = "DescriptorEvent";
                break;
            case milvus::storage::EventType::InsertEvent:
                name = "InsertEvent";
                break;
            case milvus::storage::EventType::DeleteEvent:
                name = "DeleteEvent";
                break;
            case milvus::storage::EventType::CreateCollectionEvent:
                name = "CreateCollectionEvent";
                break;
            case milvus::storage::EventType::DropCollectionEvent:
                name = "DropCollectionEvent";
                break;
            case milvus::storage::EventType::CreatePartitionEvent:
                name = "CreatePartitionEvent";
                break;
            case milvus::storage::EventType::DropPartitionEvent:
                name = "DropPartitionEvent";
                break;
            case milvus::storage::EventType::IndexFileEvent:
                name = "IndexFileEvent";
                break;
            case milvus::storage::EventType::EventTypeEnd:
                name = "EventTypeEnd";
                break;
        }
        return formatter<string_view>::format(name, ctx);
    }
};

template <>
struct fmt::formatter<milvus::storage::StorageType> : formatter<string_view> {
    auto
    format(milvus::storage::StorageType c, format_context& ctx) const {
        switch (c) {
            case milvus::storage::StorageType::None:
                return formatter<string_view>::format("None", ctx);
            case milvus::storage::StorageType::Memory:
                return formatter<string_view>::format("Memory", ctx);
            case milvus::storage::StorageType::LocalDisk:
                return formatter<string_view>::format("LocalDisk", ctx);
            case milvus::storage::StorageType::Remote:
                return formatter<string_view>::format("Remote", ctx);
        }
        return formatter<string_view>::format("unknown", ctx);
    }
};
