// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

#include "segment/SSSegmentReader.h"

#include <memory>
#include <utility>

#include "Vectors.h"
#include "codecs/snapshot/SSCodec.h"
#include "db/Types.h"
#include "db/snapshot/ResourceHelper.h"
#include "knowhere/index/vector_index/VecIndex.h"
#include "storage/disk/DiskIOReader.h"
#include "storage/disk/DiskIOWriter.h"
#include "storage/disk/DiskOperation.h"
#include "utils/Log.h"

namespace milvus {
namespace segment {

SSSegmentReader::SSSegmentReader(const std::string& dir_root, const engine::SegmentVisitorPtr& segment_visitor)
    : dir_root_(dir_root), segment_visitor_(segment_visitor) {
    auto& segment_ptr = segment_visitor_->GetSegment();
    std::string directory =
        engine::snapshot::GetResPath<engine::snapshot::Segment>(dir_root_, segment_visitor->GetSegment());

    storage::IOReaderPtr reader_ptr = std::make_shared<storage::DiskIOReader>();
    storage::IOWriterPtr writer_ptr = std::make_shared<storage::DiskIOWriter>();
    storage::OperationPtr operation_ptr = std::make_shared<storage::DiskOperation>(directory);
    fs_ptr_ = std::make_shared<storage::FSHandler>(reader_ptr, writer_ptr, operation_ptr);

    segment_ptr_ = std::make_shared<engine::Segment>();
}

Status
SSSegmentReader::Load() {
    STATUS_CHECK(LoadFields());

    STATUS_CHECK(LoadBloomFilter());

    STATUS_CHECK(LoadDeletedDocs());

    STATUS_CHECK(LoadVectorIndice());

    return Status::OK();
}

Status
SSSegmentReader::LoadField(const std::string& field_name, std::vector<uint8_t>& raw) {
    try {
        auto field_visitor = segment_visitor_->GetFieldVisitor(field_name);
        auto raw_visitor = field_visitor->GetElementVisitor(engine::FieldElementType::FET_RAW);
        std::string file_path =
            engine::snapshot::GetResPath<engine::snapshot::SegmentFile>(dir_root_, raw_visitor->GetFile());

        auto& ss_codec = codec::SSCodec::instance();
        ss_codec.GetBlockFormat()->read(fs_ptr_, file_path, raw);
    } catch (std::exception& e) {
        std::string err_msg = "Failed to load raw vectors: " + std::string(e.what());
        LOG_ENGINE_ERROR_ << err_msg;
        return Status(DB_ERROR, err_msg);
    }
    return Status::OK();
}

Status
SSSegmentReader::LoadField(const std::string& field_name, off_t offset, size_t num_bytes, std::vector<uint8_t>& raw) {
    try {
        auto field_visitor = segment_visitor_->GetFieldVisitor(field_name);
        auto raw_visitor = field_visitor->GetElementVisitor(engine::FieldElementType::FET_RAW);
        std::string file_path =
            engine::snapshot::GetResPath<engine::snapshot::SegmentFile>(dir_root_, raw_visitor->GetFile());

        auto& ss_codec = codec::SSCodec::instance();
        ss_codec.GetBlockFormat()->read(fs_ptr_, file_path, offset, num_bytes, raw);
    } catch (std::exception& e) {
        std::string err_msg = "Failed to load raw vectors: " + std::string(e.what());
        LOG_ENGINE_ERROR_ << err_msg;
        return Status(DB_ERROR, err_msg);
    }

    return Status::OK();
}

Status
SSSegmentReader::LoadFields() {
    engine::FIXEDX_FIELD_MAP& field_map = segment_ptr_->GetFixedFields();
    auto& field_visitors_map = segment_visitor_->GetFieldVisitors();
    for (auto& iter : field_visitors_map) {
        const engine::snapshot::FieldPtr& field = iter.second->GetField();
        std::string name = field->GetName();
        engine::FIXED_FIELD_DATA raw_data;
        segment_ptr_->GetFixedFieldData(name, raw_data);

        auto element_visitor = iter.second->GetElementVisitor(engine::FieldElementType::FET_RAW);
        std::string file_path =
            engine::snapshot::GetResPath<engine::snapshot::SegmentFile>(dir_root_, element_visitor->GetFile());
        STATUS_CHECK(LoadField(file_path, raw_data));

        field_map.insert(std::make_pair(name, raw_data));
    }

    return Status::OK();
}

Status
SSSegmentReader::LoadUids(std::vector<int64_t>& uids) {
    std::vector<uint8_t> raw;
    auto status = LoadField(engine::DEFAULT_UID_NAME, raw);
    if (!status.ok()) {
        LOG_ENGINE_ERROR_ << status.message();
        return status;
    }

    if (raw.size() % sizeof(int64_t) != 0) {
        std::string err_msg = "Failed to load uids: illegal file size";
        LOG_ENGINE_ERROR_ << err_msg;
        return Status(DB_ERROR, err_msg);
    }

    uids.clear();
    uids.resize(raw.size() / sizeof(int64_t));
    memcpy(uids.data(), raw.data(), raw.size());

    return Status::OK();
}

Status
SSSegmentReader::GetSegment(engine::SegmentPtr& segment_ptr) {
    segment_ptr = segment_ptr_;
    return Status::OK();
}

Status
SSSegmentReader::LoadVectorIndex(const std::string& field_name, segment::VectorIndexPtr& vector_index_ptr) {
    try {
        auto field_visitor = segment_visitor_->GetFieldVisitor(field_name);
        auto raw_visitor = field_visitor->GetElementVisitor(engine::FieldElementType::FET_RAW);
        std::string file_path =
            engine::snapshot::GetResPath<engine::snapshot::SegmentFile>(dir_root_, raw_visitor->GetFile());
        //        auto& ss_codec = codec::SSCodec::instance();
        //        ss_codec.GetVectorIndexFormat()->read(fs_ptr_, location, external_data, vector_index_ptr);
    } catch (std::exception& e) {
        std::string err_msg = "Failed to load vector index: " + std::string(e.what());
        LOG_ENGINE_ERROR_ << err_msg;
        return Status(DB_ERROR, err_msg);
    }

    return Status::OK();
}

Status
SSSegmentReader::LoadVectorIndice() {
    auto& field_visitors_map = segment_visitor_->GetFieldVisitors();
    for (auto& iter : field_visitors_map) {
        const engine::snapshot::FieldPtr& field = iter.second->GetField();
        std::string name = field->GetName();

        auto element_visitor = iter.second->GetElementVisitor(engine::FieldElementType::FET_INDEX);
        if (element_visitor == nullptr) {
            continue;
        }

        if (field->GetFtype() == engine::FIELD_TYPE::VECTOR || field->GetFtype() == engine::FIELD_TYPE::VECTOR_FLOAT ||
            field->GetFtype() == engine::FIELD_TYPE::VECTOR_BINARY) {
            std::string file_path =
                engine::snapshot::GetResPath<engine::snapshot::SegmentFile>(dir_root_, element_visitor->GetFile());

            segment::VectorIndexPtr vector_index_ptr;
            STATUS_CHECK(LoadVectorIndex(name, vector_index_ptr));

            segment_ptr_->SetVectorIndex(name, vector_index_ptr->GetVectorIndex());
        }
    }

    return Status::OK();
}

Status
SSSegmentReader::LoadBloomFilter(segment::IdBloomFilterPtr& id_bloom_filter_ptr) {
    try {
        auto uid_field_visitor = segment_visitor_->GetFieldVisitor(engine::DEFAULT_UID_NAME);
        auto visitor = uid_field_visitor->GetElementVisitor(engine::FieldElementType::FET_BLOOM_FILTER);
        std::string file_path =
            engine::snapshot::GetResPath<engine::snapshot::SegmentFile>(dir_root_, visitor->GetFile());

        auto& ss_codec = codec::SSCodec::instance();
        ss_codec.GetIdBloomFilterFormat()->read(fs_ptr_, file_path, id_bloom_filter_ptr);
    } catch (std::exception& e) {
        std::string err_msg = "Failed to load bloom filter: " + std::string(e.what());
        LOG_ENGINE_ERROR_ << err_msg;
        return Status(DB_ERROR, err_msg);
    }
    return Status::OK();
}

Status
SSSegmentReader::LoadBloomFilter() {
    segment::IdBloomFilterPtr id_bloom_filter_ptr;
    auto status = LoadBloomFilter(id_bloom_filter_ptr);
    if (!status.ok()) {
        return status;
    }

    segment_ptr_->SetBloomFilter(id_bloom_filter_ptr);
    return Status::OK();
}

Status
SSSegmentReader::LoadDeletedDocs(segment::DeletedDocsPtr& deleted_docs_ptr) {
    try {
        auto uid_field_visitor = segment_visitor_->GetFieldVisitor(engine::DEFAULT_UID_NAME);
        auto visitor = uid_field_visitor->GetElementVisitor(engine::FieldElementType::FET_BLOOM_FILTER);
        std::string file_path =
            engine::snapshot::GetResPath<engine::snapshot::SegmentFile>(dir_root_, visitor->GetFile());

        auto& ss_codec = codec::SSCodec::instance();
        ss_codec.GetDeletedDocsFormat()->read(fs_ptr_, file_path, deleted_docs_ptr);
    } catch (std::exception& e) {
        std::string err_msg = "Failed to load deleted docs: " + std::string(e.what());
        LOG_ENGINE_ERROR_ << err_msg;
        return Status(DB_ERROR, err_msg);
    }
    return Status::OK();
}

Status
SSSegmentReader::LoadDeletedDocs() {
    segment::DeletedDocsPtr deleted_docs_ptr;
    auto status = LoadDeletedDocs(deleted_docs_ptr);
    if (!status.ok()) {
        return status;
    }

    segment_ptr_->SetDeletedDocs(deleted_docs_ptr);
    return Status::OK();
}

Status
SSSegmentReader::ReadDeletedDocsSize(size_t& size) {
    try {
        auto& ss_codec = codec::SSCodec::instance();
        ss_codec.GetDeletedDocsFormat()->readSize(fs_ptr_, size);
    } catch (std::exception& e) {
        std::string err_msg = "Failed to read deleted docs size: " + std::string(e.what());
        LOG_ENGINE_ERROR_ << err_msg;
        return Status(DB_ERROR, err_msg);
    }
    return Status::OK();
}
}  // namespace segment
}  // namespace milvus
