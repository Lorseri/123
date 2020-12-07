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

#include "knowhere/index/vector_index/IndexBinaryIVF.h"

#include <faiss/IndexBinaryFlat.h>
#include <faiss/IndexBinaryIVF.h>

#include <chrono>
#include <string>

#include "knowhere/common/Exception.h"
#include "knowhere/common/Log.h"
#include "knowhere/index/vector_index/adapter/VectorAdapter.h"

namespace milvus {
namespace knowhere {

using stdclock = std::chrono::high_resolution_clock;

BinarySet
BinaryIVF::Serialize(const Config& config) {
    if (!index_ || !index_->is_trained) {
        KNOWHERE_THROW_MSG("index not initialize or trained");
    }

    std::lock_guard<std::mutex> lk(mutex_);
    return SerializeImpl(index_type_);
}

void
BinaryIVF::Load(const BinarySet& index_binary) {
    std::lock_guard<std::mutex> lk(mutex_);
    LoadImpl(index_binary, index_type_);

    auto ivf_index = dynamic_cast<faiss::IndexBinaryIVF*>(index_.get());
    if (STATISTICS_LEVEL) {
        stats = std::make_shared<milvus::knowhere::IVFStatistics>(index_type_);
        ivf_index->nprobe_statistics.resize(ivf_index->nlist);
        ivf_index->nprobe_statistics.assign(ivf_index->nlist, 0);
        stats->Clear();
    }
}

DatasetPtr
BinaryIVF::Query(const DatasetPtr& dataset_ptr, const Config& config, const faiss::ConcurrentBitsetPtr& bitset) {
    if (!index_ || !index_->is_trained) {
        KNOWHERE_THROW_MSG("index not initialize or trained");
    }

    GET_TENSOR_DATA(dataset_ptr)

    try {
        auto k = config[meta::TOPK].get<int64_t>();
        auto elems = rows * k;

        size_t p_id_size = sizeof(int64_t) * elems;
        size_t p_dist_size = sizeof(float) * elems;
        auto p_id = static_cast<int64_t*>(malloc(p_id_size));
        auto p_dist = static_cast<float*>(malloc(p_dist_size));

        QueryImpl(rows, reinterpret_cast<const uint8_t*>(p_data), k, p_dist, p_id, config, bitset);

        auto ret_ds = std::make_shared<Dataset>();

        ret_ds->Set(meta::IDS, p_id);
        ret_ds->Set(meta::DISTANCE, p_dist);

        return ret_ds;
    } catch (faiss::FaissException& e) {
        KNOWHERE_THROW_MSG(e.what());
    } catch (std::exception& e) {
        KNOWHERE_THROW_MSG(e.what());
    }
}

int64_t
BinaryIVF::Count() {
    if (!index_) {
        KNOWHERE_THROW_MSG("index not initialize");
    }
    return index_->ntotal;
}

int64_t
BinaryIVF::Dim() {
    if (!index_) {
        KNOWHERE_THROW_MSG("index not initialize");
    }
    return index_->d;
}

void
BinaryIVF::UpdateIndexSize() {
    if (!index_) {
        KNOWHERE_THROW_MSG("index not initialize");
    }
    auto bin_ivf_index = dynamic_cast<faiss::IndexBinaryIVF*>(index_.get());
    auto nb = bin_ivf_index->invlists->compute_ntotal();
    auto nlist = bin_ivf_index->nlist;
    auto code_size = bin_ivf_index->code_size;

    // binary ivf codes, ids and quantizer
    index_size_ = nb * code_size + nb * sizeof(int64_t) + nlist * code_size;
}

void
BinaryIVF::ClearStatistics() {
    if (stats != nullptr) {
        auto ivf_stats = std::dynamic_pointer_cast<IVFStatistics>(stats);
        ivf_stats->Clear();
        auto ivf_index = dynamic_cast<faiss::IndexBinaryIVF*>(index_.get());
        ivf_index->clear_nprobe_statistics();
        ivf_index->index_ivf_stats.reset();
    }
}

void
BinaryIVF::Train(const DatasetPtr& dataset_ptr, const Config& config) {
    GET_TENSOR(dataset_ptr)

    int64_t nlist = config[IndexParams::nlist];
    faiss::MetricType metric_type = GetMetricType(config[Metric::TYPE].get<std::string>());
    faiss::IndexBinary* coarse_quantizer = new faiss::IndexBinaryFlat(dim, metric_type);
    auto index = std::make_shared<faiss::IndexBinaryIVF>(coarse_quantizer, dim, nlist, metric_type);
    index->own_fields = true;
    index->train(rows, static_cast<const uint8_t*>(p_data));
    index->add_with_ids(rows, static_cast<const uint8_t*>(p_data), p_ids);
    index_ = index;
}

std::shared_ptr<faiss::IVFSearchParameters>
BinaryIVF::GenParams(const Config& config) {
    auto params = std::make_shared<faiss::IVFSearchParameters>();
    params->nprobe = config[IndexParams::nprobe];
    // params->max_codes = config["max_code"];
    return params;
}

void
BinaryIVF::QueryImpl(int64_t n, const uint8_t* data, int64_t k, float* distances, int64_t* labels, const Config& config,
                     const faiss::ConcurrentBitsetPtr& bitset) {
    auto params = GenParams(config);
    auto ivf_index = dynamic_cast<faiss::IndexBinaryIVF*>(index_.get());
    ivf_index->nprobe = params->nprobe;

    stdclock::time_point before = stdclock::now();
    auto i_distances = reinterpret_cast<int32_t*>(distances);

    index_->search(n, data, k, i_distances, labels, bitset);

    stdclock::time_point after = stdclock::now();
    double search_cost = (std::chrono::duration<double, std::micro>(after - before)).count();
    if (STATISTICS_LEVEL) {
        auto ivf_stats = std::dynamic_pointer_cast<IVFStatistics>(stats);
        if (STATISTICS_LEVEL >= 1) {
            ivf_stats->nq_cnt += n;
            ivf_stats->batch_cnt += 1;
            ivf_stats->nprobe_access_count = ivf_index->index_ivf_stats.nlist;

            if (n > 2048) {
                ivf_stats->nq_stat[12]++;
            }
            else {
                ivf_stats->nq_stat[len_of_pow2(upper_bound_of_pow2((uint64_t)n))]++;
            }

            LOG_KNOWHERE_DEBUG_ << "IVF_NM search cost: " << search_cost
                                << ", quantization cost: " << ivf_index->index_ivf_stats.quantization_time
                                << ", data search cost: " << ivf_index->index_ivf_stats.search_time;
            ivf_stats->total_query_time +=
                ivf_index->index_ivf_stats.quantization_time + ivf_index->index_ivf_stats.search_time;
            ivf_index->index_ivf_stats.quantization_time = 0;
            ivf_index->index_ivf_stats.search_time = 0;
        }
        if (STATISTICS_LEVEL >= 2) {
            double fps = bitset ? (double)bitset->count_1() / bitset->count() : 0.0;
            if (fps > 1.0 || fps < 0.0) {
                LOG_KNOWHERE_ERROR_ << "in IndexIVF::Query, the percentage of 1 in bitset is " << fps
                                    << ", which is exceed 100% or negative!";
            }
            else {
                ivf_stats->filter_stat[(int)(fps * 100) / 5] += 1;
            }
        }
        if (STATISTICS_LEVEL >= 3) {
            ivf_stats->UpdateStatistics(ivf_index->nprobe_statistics);
        }
    }

    // if hamming, it need transform int32 to float
    if (ivf_index->metric_type == faiss::METRIC_Hamming) {
        int64_t num = n * k;
        for (int64_t i = 0; i < num; i++) {
            distances[i] = static_cast<float>(i_distances[i]);
        }
    }
}

}  // namespace knowhere
}  // namespace milvus
