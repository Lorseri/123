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

#include "db/wal/WalFile.h"
#include "db/Constants.h"
#include "db/Types.h"

namespace milvus {
namespace engine {

WalFile::WalFile() {
}

WalFile::~WalFile() {
    CloseFile();
}

Status
WalFile::OpenFile(const std::string& path, OpenMode mode) {
    CloseFile();

    try {
        std::string str_mode = (mode == OpenMode::READ) ? "rb" : "awb";
        file_ = fopen(path.c_str(), str_mode.c_str());
        if (file_ == nullptr) {
            std::string msg = "Failed to create wal file: " + path;
            return Status(DB_ERROR, msg);
        }
        file_path_ = path;
        mode_ = mode;
    } catch (std::exception& ex) {
        std::string msg = "Failed to create wal file, reason: " + std::string(ex.what());
        return Status(DB_ERROR, msg);
    }

    return Status::OK();
}

Status
WalFile::CloseFile() {
    if (file_ != nullptr) {
        fclose(file_);
        file_ = nullptr;
        file_size_ = 0;
        file_path_ = "";
    }

    return Status::OK();
}

bool
WalFile::ExceedMaxSize(int64_t append_size) {
    return (file_size_ + append_size) > MAX_WAL_FILE_SIZE;
}

idx_t
WalFile::ReadLastOpId() {
    if (file_ == nullptr) {
        return 0;
    }

    // current position
    auto cur_poz = ftell(file_);

    // get total lenth
    fseek(file_, 0, SEEK_END);
    auto end_poz = ftell(file_);

    // read last id
    idx_t last_id = 0;
    int64_t offset = end_poz - sizeof(last_id);
    fseek(file_, offset, SEEK_SET);

    fread(&last_id, 1, sizeof(last_id), file_);

    // back to current postiion
    fseek(file_, cur_poz, SEEK_SET);
    return last_id;
}

}  // namespace engine
}  // namespace milvus
