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

#pragma once

#include <string>
#include <unordered_map>
#include <utility>
#include <vector>

#include "db/snapshot/ResourceTypes.h"

namespace milvus::engine {

enum ResourceContextOp {
    oAdd = 1,
    oUpdate,
    oDelete
};

struct SqlContext {
    std::string sql_;
    ResourceContextOp op_;
    snapshot::ID_TYPE id_;
};

struct DBQueryContext {
    std::string table_;
    bool all_required_ = true;
    std::vector<std::string> query_fields_;
    std::unordered_map<std::string, std::string> filter_attrs_;
};

struct DBApplyContext {
    std::string table_;
    ResourceContextOp op_;
    snapshot::ID_TYPE id_;
    std::unordered_map<std::string, std::string> attrs_;
    std::unordered_map<std::string, std::string> filter_attrs_;
    std::string sql_;
};

}