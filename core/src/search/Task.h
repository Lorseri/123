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

#include <iostream>
#include <memory>
#include <vector>
#include <string>

#include "db/Types.h"
#include "db/meta/MetaTypes.h"
#include "context/HybridSearchContext.h"
#include "server/context/Context.h"
#include "db/engine/ExecutionEngine.h"
#include "utils/Status.h"

namespace milvus {

namespace context {
struct HybridSearchContext;
using HybridSearchContextPtr = std::shared_ptr<HybridSearchContext>;
}

namespace search{

using TableFileSchemaPtr = engine::meta::TableFileSchemaPtr;

using Id2IndexMap = std::unordered_map<size_t, TableFileSchemaPtr>;

using ResultIds = engine::ResultIds;
using ResultDistances = engine::ResultDistances;

class Task {
 public:
    explicit Task(const std::shared_ptr<server::Context>& context,
                  TableFileSchemaPtr& file,
                  query::GeneralQueryPtr general_query,
                  std::unordered_map<std::string, engine::DataType>& attr_type,
                  context::HybridSearchContextPtr hybrid_search_context);

    void
    Load();

    void
    Execute();

 public:
    static void
    MergeTopkToResultSet(const ResultIds& src_ids, const ResultDistances& src_distances,
                         size_t src_k, size_t nq, size_t topk, bool ascending, ResultIds& tar_ids,
                         ResultDistances& tar_distances);

    const std::string&
    GetLocation() const;

    size_t
    GetIndexId() const;

 public:
    const std::shared_ptr<server::Context> context_;

    TableFileSchemaPtr file_;

    size_t index_id_ = 0;
    int index_type_ = 0;
    engine::ExecutionEnginePtr index_engine_ = nullptr;

    // distance -- value 0 means two vectors equal, ascending reduce, L2/HAMMING/JACCARD/TONIMOTO ...
    // similarity -- infinity value means two vectors equal, descending reduce, IP
    bool ascending_reduce = true;

    query::GeneralQueryPtr general_query_;
    std::unordered_map<std::string, engine::DataType> attr_type_;
    context::HybridSearchContextPtr hybrid_search_context_;

    ResultIds result_ids_;
    ResultDistances result_distances_;
};

using TaskPtr = std::shared_ptr<Task>;

} // namespace search
} // namespace milvus