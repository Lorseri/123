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

#include <map>
#include <memory>
#include <string>
#include <unordered_map>
#include <vector>

#include <faiss/utils/ConcurrentBitset.h>

#include "knowhere/index/IndexType.h"
#include "query/GeneralQuery.h"
#include "utils/Json.h"
#include "utils/Status.h"

namespace milvus {

namespace scheduler {
class SearchJob;
using SearchJobPtr = std::shared_ptr<SearchJob>;
}  // namespace scheduler

namespace engine {

// TODO(linxj): replace with VecIndex::IndexType
enum class EngineType {
    INVALID = 0,
    FAISS_IDMAP = 1,
    FAISS_IVFFLAT,
    FAISS_IVFSQ8,
    NSG_MIX,
    FAISS_IVFSQ8H,
    FAISS_PQ,
#ifdef MILVUS_SUPPORT_SPTAG
    SPTAG_KDT,
    SPTAG_BKT,
#endif
    FAISS_BIN_IDMAP,
    FAISS_BIN_IVFFLAT,
    HNSW,
    ANNOY,
    FAISS_IVFSQ8NR,
    HNSW_SQ8NR,
    MAX_VALUE = HNSW_SQ8NR,
};

static std::map<std::string, EngineType> s_map_engine_type = {
    {knowhere::IndexEnum::INDEX_FAISS_IDMAP, EngineType::FAISS_IDMAP},
    {knowhere::IndexEnum::INDEX_FAISS_IVFFLAT, EngineType::FAISS_IVFFLAT},
    {knowhere::IndexEnum::INDEX_FAISS_IVFPQ, EngineType::FAISS_PQ},
    {knowhere::IndexEnum::INDEX_FAISS_IVFSQ8, EngineType::FAISS_IVFSQ8},
    {knowhere::IndexEnum::INDEX_FAISS_IVFSQ8NR, EngineType::FAISS_IVFSQ8NR},
    {knowhere::IndexEnum::INDEX_FAISS_IVFSQ8H, EngineType::FAISS_IVFSQ8H},
    {knowhere::IndexEnum::INDEX_FAISS_BIN_IDMAP, EngineType::FAISS_BIN_IDMAP},
    {knowhere::IndexEnum::INDEX_FAISS_BIN_IVFFLAT, EngineType::FAISS_BIN_IVFFLAT},
    {knowhere::IndexEnum::INDEX_NSG, EngineType::NSG_MIX},
#ifdef MILVUS_SUPPORT_SPTAG
    {knowhere::IndexEnum::INDEX_SPTAG_KDT_RNT, EngineType::SPTAG_KDT},
    {knowhere::IndexEnum::INDEX_SPTAG_BKT_RNT, EngineType::SPTAG_BKT},
#endif
    {knowhere::IndexEnum::INDEX_HNSW, EngineType::HNSW},
    {knowhere::IndexEnum::INDEX_HNSW_SQ8NR, EngineType::HNSW_SQ8NR},
    {knowhere::IndexEnum::INDEX_ANNOY, EngineType::ANNOY},
};

enum class MetricType {
    L2 = 1,              // Euclidean Distance
    IP = 2,              // Cosine Similarity
    HAMMING = 3,         // Hamming Distance
    JACCARD = 4,         // Jaccard Distance
    TANIMOTO = 5,        // Tanimoto Distance
    SUBSTRUCTURE = 6,    // Substructure Distance
    SUPERSTRUCTURE = 7,  // Superstructure Distance
    MAX_VALUE = SUPERSTRUCTURE
};

enum class DataType {
    INT8 = 1,
    INT16 = 2,
    INT32 = 3,
    INT64 = 4,

    STRING = 20,

    BOOL = 30,

    FLOAT = 40,
    DOUBLE = 41,

    VECTOR = 100,
    UNKNOWN = 9999,
};

class ExecutionEngine {
 public:
    virtual Status
    AddWithIds(int64_t n, const float* xdata, const int64_t* xids) = 0;

    virtual Status
    AddWithIds(int64_t n, const uint8_t* xdata, const int64_t* xids) = 0;

    virtual size_t
    Count() const = 0;

    virtual size_t
    Dimension() const = 0;

    virtual size_t
    Size() const = 0;

    virtual Status
    Serialize() = 0;

    virtual Status
    Load(bool to_cache = true) = 0;

    virtual Status
    LoadAttr(bool to_cache = true) = 0;

    virtual Status
    CopyToGpu(uint64_t device_id, bool hybrid) = 0;

    virtual Status
    CopyToIndexFileToGpu(uint64_t device_id) = 0;

    virtual Status
    CopyToCpu() = 0;

    //    virtual std::shared_ptr<ExecutionEngine>
    //    Clone() = 0;

    //    virtual Status
    //    Merge(const std::string& location) = 0;

#if 0
    virtual Status
    GetVectorByID(const int64_t id, float* vector, bool hybrid) = 0;

    virtual Status
    GetVectorByID(const int64_t id, uint8_t* vector, bool hybrid) = 0;
#endif

    virtual Status
    ExecBinaryQuery(query::GeneralQueryPtr general_query, faiss::ConcurrentBitsetPtr& bitset,
                    std::unordered_map<std::string, DataType>& attr_type, std::string& vector_placeholder) = 0;

    virtual Status
    HybridSearch(scheduler::SearchJobPtr job, std::unordered_map<std::string, DataType>& attr_type,
                 std::vector<float>& distances, std::vector<int64_t>& search_ids, bool hybrid) = 0;

    virtual Status
    Search(std::vector<int64_t>& ids, std::vector<float>& distances, scheduler::SearchJobPtr job, bool hybrid) = 0;

    virtual std::shared_ptr<ExecutionEngine>
    BuildIndex(const std::string& location, EngineType engine_type) = 0;

    virtual Status
    Cache() = 0;

    virtual Status
    AttrCache() = 0;

    virtual Status
    Init() = 0;

    virtual EngineType
    IndexEngineType() const = 0;

    virtual MetricType
    IndexMetricType() const = 0;

    virtual std::string
    GetLocation() const = 0;

    virtual std::string
    GetAttrLocation() const = 0;
};

using ExecutionEnginePtr = std::shared_ptr<ExecutionEngine>;

}  // namespace engine
}  // namespace milvus
