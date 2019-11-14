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

#include "wrapper/ConfAdapter.h"
#include "knowhere/index/vector_index/helpers/IndexParameter.h"
#include "utils/Log.h"

#include <cmath>
#include <memory>

// TODO(lxj): add conf checker

namespace milvus {
namespace engine {

#if CUDA_VERSION > 9000
#define GPU_MAX_NRPOBE 2048
#else
#define GPU_MAX_NRPOBE 1024
#endif

void
ConfAdapter::MatchBase(knowhere::Config conf) {
    if (conf->metric_type == knowhere::DEFAULT_TYPE)
        conf->metric_type = knowhere::METRICTYPE::L2;
    if (conf->gpu_id == knowhere::INVALID_VALUE)
        conf->gpu_id = 0;
}

knowhere::Config
ConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::Cfg>();
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    conf->gpu_id = conf->gpu_id;
    MatchBase(conf);
    return conf;
}

knowhere::Config
ConfAdapter::MatchSearch(const TempMetaConf& metaconf, const IndexType& type) {
    auto conf = std::make_shared<knowhere::Cfg>();
    conf->k = metaconf.k;
    return conf;
}

knowhere::Config
IVFConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::IVFCfg>();
    conf->nlist = MatchNlist(metaconf.size, metaconf.nlist);
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    conf->gpu_id = conf->gpu_id;
    MatchBase(conf);
    return conf;
}

static constexpr float TYPICAL_COUNT = 1000000.0;

int64_t
IVFConfAdapter::MatchNlist(const int64_t& size, const int64_t& nlist) {
    if (size <= TYPICAL_COUNT / 16384 + 1) {
        // handle less row count, avoid nlist set to 0
        return 1;
    } else if (int(size / TYPICAL_COUNT) * nlist <= 0) {
        // calculate a proper nlist if nlist not specified or size less than TYPICAL_COUNT
        return int(size / TYPICAL_COUNT * 16384);
    }
    return nlist;
}

knowhere::Config
IVFConfAdapter::MatchSearch(const TempMetaConf& metaconf, const IndexType& type) {
    auto conf = std::make_shared<knowhere::IVFCfg>();
    conf->k = metaconf.k;

    if (metaconf.nprobe <= 0)
        conf->nprobe = 16;  // hardcode here
    else
        conf->nprobe = metaconf.nprobe;

    switch (type) {
        case IndexType::FAISS_IVFFLAT_GPU:
        case IndexType::FAISS_IVFSQ8_GPU:
        case IndexType::FAISS_IVFPQ_GPU:
            if (conf->nprobe > GPU_MAX_NRPOBE) {
                WRAPPER_LOG_WARNING << "When search with GPU, nprobe shoud be no more than " << GPU_MAX_NRPOBE
                                    << ", but you passed " << conf->nprobe << ". Search with " << GPU_MAX_NRPOBE
                                    << " instead";
                conf->nprobe = GPU_MAX_NRPOBE;
            }
    }
    return conf;
}

knowhere::Config
IVFSQConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::IVFSQCfg>();
    conf->nlist = MatchNlist(metaconf.size, metaconf.nlist);
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    conf->gpu_id = conf->gpu_id;
    conf->nbits = 8;
    MatchBase(conf);
    return conf;
}

knowhere::Config
IVFPQConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::IVFPQCfg>();
    conf->nlist = MatchNlist(metaconf.size, metaconf.nlist);
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    conf->gpu_id = conf->gpu_id;
    conf->nbits = 8;

    if (!(conf->d % 4))
        conf->m = conf->d / 4;  // compression radio = 16
    else if (!(conf->d % 2))
        conf->m = conf->d / 2;  // compression radio = 8
    else if (!(conf->d % 3))
        conf->m = conf->d / 3;  // compression radio = 12
    else
        conf->m = conf->d;  // same as SQ8, compression radio = 4

    MatchBase(conf);
    return conf;
}

knowhere::Config
IVFPQConfAdapter::MatchSearch(const TempMetaConf& metaconf, const IndexType& type) {
    auto conf = std::make_shared<knowhere::IVFPQCfg>();
    conf->k = metaconf.k;

    if (metaconf.nprobe <= 0)
        conf->nprobe = 16;  // hardcode here
    else
        conf->nprobe = metaconf.nprobe;

    return conf;
}

int64_t
IVFPQConfAdapter::MatchNlist(const int64_t& size, const int64_t& nlist) {
    if (size <= TYPICAL_COUNT / 16384 + 1) {
        // handle less row count, avoid nlist set to 0
        return 1;
    } else if (int(size / TYPICAL_COUNT) * nlist <= 0) {
        // calculate a proper nlist if nlist not specified or size less than TYPICAL_COUNT
        return int(size / TYPICAL_COUNT * 16384);
    }
    return nlist;
}

knowhere::Config
NSGConfAdapter::Match(const TempMetaConf& metaconf) {
    auto conf = std::make_shared<knowhere::NSGCfg>();
    conf->nlist = MatchNlist(metaconf.size, metaconf.nlist);
    conf->d = metaconf.dim;
    conf->metric_type = metaconf.metric_type;
    conf->gpu_id = conf->gpu_id;

    double factor = metaconf.size / TYPICAL_COUNT;
    auto scale_factor = round(metaconf.dim / 128.0);
    scale_factor = scale_factor >= 4 ? 4 : scale_factor;
    conf->nprobe = conf->nlist > 10000 ? conf->nlist * 0.02 : conf->nlist * 0.1;
    conf->knng = (100 + 100 * scale_factor) * factor;
    conf->search_length = (40 + 5 * scale_factor) * factor;
    conf->out_degree = (50 + 5 * scale_factor) * factor;
    conf->candidate_pool_size = (200 + 100 * scale_factor) * factor;
    MatchBase(conf);

    //    WRAPPER_LOG_DEBUG << "nlist: " << conf->nlist
    //    << ", gpu_id: " << conf->gpu_id << ", d: " << conf->d
    //    << ", nprobe: " << conf->nprobe << ", knng: " << conf->knng;
    return conf;
}

knowhere::Config
NSGConfAdapter::MatchSearch(const TempMetaConf& metaconf, const IndexType& type) {
    auto conf = std::make_shared<knowhere::NSGCfg>();
    conf->k = metaconf.k;
    conf->search_length = metaconf.search_length;
    if (metaconf.search_length == TEMPMETA_DEFAULT_VALUE) {
        conf->search_length = 30;  // TODO(linxj): hardcode here.
    }
    return conf;
}

}  // namespace engine
}  // namespace milvus
