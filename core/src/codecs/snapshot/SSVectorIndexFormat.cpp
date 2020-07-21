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

#include <boost/filesystem.hpp>
#include <memory>

#include "codecs/snapshot/SSCodec.h"
#include "codecs/snapshot/SSVectorIndexFormat.h"
#include "knowhere/common/BinarySet.h"
#include "knowhere/index/vector_index/VecIndex.h"
#include "knowhere/index/vector_index/VecIndexFactory.h"
#include "utils/Exception.h"
#include "utils/Log.h"
#include "utils/TimeRecorder.h"

namespace milvus {
namespace codec {

const char* VECTOR_INDEX_POSTFIX = ".idx";

std::string
SSVectorIndexFormat::FilePostfix() {
    std::string str = VECTOR_INDEX_POSTFIX;
    return str;
}

void
SSVectorIndexFormat::ReadRaw(const storage::FSHandlerPtr& fs_ptr, const std::string& file_path,
                             knowhere::BinaryPtr& data) {
    milvus::TimeRecorder recorder("SSVectorIndexFormat::ReadRaw");

    if (!fs_ptr->reader_ptr_->open(file_path.c_str())) {
        std::string err_msg = "Failed to open raw file: " + file_path + ", error: " + std::strerror(errno);
        LOG_ENGINE_ERROR_ << err_msg;
        throw Exception(SERVER_CANNOT_OPEN_FILE, err_msg);
    }

    size_t num_bytes;
    fs_ptr->reader_ptr_->read(&num_bytes, sizeof(size_t));

    data = std::make_shared<knowhere::Binary>();
    data->size = num_bytes;
    data->data = std::shared_ptr<uint8_t[]>(new uint8_t[num_bytes]);

    // Beginning of file is num_bytes
    fs_ptr->reader_ptr_->seekg(sizeof(size_t));
    fs_ptr->reader_ptr_->read(data->data.get(), num_bytes);
    fs_ptr->reader_ptr_->close();

    double span = recorder.RecordSection("End");
    double rate = num_bytes * 1000000.0 / span / 1024 / 1024;
    LOG_ENGINE_DEBUG_ << "SSVectorIndexFormat::ReadIndex(" << file_path << ") rate " << rate << "MB/s";
}

void
SSVectorIndexFormat::ReadIndex(const storage::FSHandlerPtr& fs_ptr, const std::string& file_path,
                               knowhere::BinarySet& data) {
    milvus::TimeRecorder recorder("SSVectorIndexFormat::ReadIndex");

    std::string full_file_path = file_path + VECTOR_INDEX_POSTFIX;
    if (!fs_ptr->reader_ptr_->open(full_file_path)) {
        std::string err_msg = "Failed to open vector index: " + full_file_path + ", error: " + std::strerror(errno);
        LOG_ENGINE_ERROR_ << err_msg;
        throw Exception(SERVER_CANNOT_OPEN_FILE, err_msg);
    }

    int64_t length = fs_ptr->reader_ptr_->length();
    if (length <= 0) {
        LOG_ENGINE_ERROR_ << "Invalid vector index length: " << full_file_path;
        return;
    }

    int64_t rp = 0;
    fs_ptr->reader_ptr_->seekg(0);

    int32_t current_type = 0;
    fs_ptr->reader_ptr_->read(&current_type, sizeof(current_type));
    rp += sizeof(current_type);
    fs_ptr->reader_ptr_->seekg(rp);

    LOG_ENGINE_DEBUG_ << "Start to ReadIndex(" << full_file_path << ") length: " << length << " bytes";
    while (rp < length) {
        size_t meta_length;
        fs_ptr->reader_ptr_->read(&meta_length, sizeof(meta_length));
        rp += sizeof(meta_length);
        fs_ptr->reader_ptr_->seekg(rp);

        auto meta = new char[meta_length];
        fs_ptr->reader_ptr_->read(meta, meta_length);
        rp += meta_length;
        fs_ptr->reader_ptr_->seekg(rp);

        size_t bin_length;
        fs_ptr->reader_ptr_->read(&bin_length, sizeof(bin_length));
        rp += sizeof(bin_length);
        fs_ptr->reader_ptr_->seekg(rp);

        auto bin = new uint8_t[bin_length];
        fs_ptr->reader_ptr_->read(bin, bin_length);
        rp += bin_length;
        fs_ptr->reader_ptr_->seekg(rp);

        std::shared_ptr<uint8_t[]> binptr(bin);
        data.Append(std::string(meta, meta_length), binptr, bin_length);
        delete[] meta;
    }
    fs_ptr->reader_ptr_->close();

    double span = recorder.RecordSection("End");
    double rate = length * 1000000.0 / span / 1024 / 1024;
    LOG_ENGINE_DEBUG_ << "SSVectorIndexFormat::ReadIndex(" << full_file_path << ") rate " << rate << "MB/s";
}

void
SSVectorIndexFormat::ReadCompress(const storage::FSHandlerPtr& fs_ptr, const std::string& file_path,
                                  knowhere::BinaryPtr& data) {
    auto& ss_codec = codec::SSCodec::instance();
    ss_codec.GetVectorCompressFormat()->Read(fs_ptr, file_path, data);
}

void
SSVectorIndexFormat::ConvertRaw(const std::vector<uint8_t>& raw, knowhere::BinaryPtr& data) {
    data = std::make_shared<knowhere::Binary>();
    data->size = raw.size();
    data->data = std::shared_ptr<uint8_t[]>(new uint8_t[data->size]);
}

void
SSVectorIndexFormat::ConstructIndex(const std::string& index_name, knowhere::BinarySet& index_data,
                                    knowhere::BinaryPtr& raw_data, knowhere::BinaryPtr& compress_data,
                                    knowhere::VecIndexPtr& index) {
    knowhere::VecIndexFactory& vec_index_factory = knowhere::VecIndexFactory::GetInstance();
    index = vec_index_factory.CreateVecIndex(index_name, knowhere::IndexMode::MODE_CPU);
    if (index != nullptr) {
        int64_t length = 0;
        for (auto& pair : index_data.binary_map_) {
            length += pair.second->size;
        }

        if (raw_data != nullptr) {
            LOG_ENGINE_DEBUG_ << "load index with " << RAW_DATA << " " << raw_data->size;
            index_data.Append(RAW_DATA, raw_data);
            length += raw_data->size;
        }

        if (compress_data != nullptr) {
            LOG_ENGINE_DEBUG_ << "load index with " << SQ8_DATA << " " << compress_data->size;
            index_data.Append(SQ8_DATA, compress_data);
            length += compress_data->size;
        }

        index->Load(index_data);
        index->SetIndexSize(length);
    } else {
        std::string err_msg = "Fail to create vector index";
        LOG_ENGINE_ERROR_ << err_msg;
        throw Exception(SERVER_UNEXPECTED_ERROR, err_msg);
    }
}

void
SSVectorIndexFormat::WriteIndex(const storage::FSHandlerPtr& fs_ptr, const std::string& file_path,
                                const knowhere::VecIndexPtr& index) {
    milvus::TimeRecorder recorder("SVectorIndexFormat::WriteIndex");

    std::string full_file_path = file_path + VECTOR_INDEX_POSTFIX;
    auto binaryset = index->Serialize(knowhere::Config());
    int32_t index_type = knowhere::StrToOldIndexType(index->index_type());

    if (!fs_ptr->writer_ptr_->open(full_file_path)) {
        LOG_ENGINE_ERROR_ << "Fail to open vector index: " << full_file_path;
        return;
    }

    fs_ptr->writer_ptr_->write(&index_type, sizeof(index_type));

    for (auto& iter : binaryset.binary_map_) {
        auto meta = iter.first.c_str();
        size_t meta_length = iter.first.length();
        fs_ptr->writer_ptr_->write(&meta_length, sizeof(meta_length));
        fs_ptr->writer_ptr_->write((void*)meta, meta_length);

        auto binary = iter.second;
        int64_t binary_length = binary->size;
        fs_ptr->writer_ptr_->write(&binary_length, sizeof(binary_length));
        fs_ptr->writer_ptr_->write((void*)binary->data.get(), binary_length);
    }
    fs_ptr->writer_ptr_->close();

    double span = recorder.RecordSection("End");
    double rate = fs_ptr->writer_ptr_->length() * 1000000.0 / span / 1024 / 1024;
    LOG_ENGINE_DEBUG_ << "SSVectorIndexFormat::WriteIndex(" << full_file_path << ") rate " << rate << "MB/s";
}

void
SSVectorIndexFormat::WriteCompress(const storage::FSHandlerPtr& fs_ptr, const std::string& file_path,
                                   const knowhere::VecIndexPtr& index) {
    milvus::TimeRecorder recorder("SSVectorIndexFormat::WriteCompress");

    auto binaryset = index->Serialize(knowhere::Config());

    auto sq8_data = binaryset.Erase(SQ8_DATA);
    if (sq8_data != nullptr) {
        auto& ss_codec = codec::SSCodec::instance();
        ss_codec.GetVectorCompressFormat()->Write(fs_ptr, file_path, sq8_data);
    }
}

}  // namespace codec
}  // namespace milvus
