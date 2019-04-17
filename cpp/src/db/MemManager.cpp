/* #include <faiss/IndexFlat.h> */
/* #include <faiss/MetaIndexes.h> */
#include <faiss/AutoTune.h>
#include <faiss/index_io.h>
#include <iostream>
#include <sstream>
#include <thread>

#include "MemManager.h"
#include "Meta.h"


namespace zilliz {
namespace vecwise {
namespace engine {

MemVectors::MemVectors(const std::string& group_id,
        size_t dimension, const std::string& file_location) :
    group_id_(group_id),
    _file_location(file_location),
    _pIdGenerator(new SimpleIDGenerator()),
    _dimension(dimension),
    pIndex_(faiss::index_factory(_dimension, "IDMap,Flat")) {
}

void MemVectors::add(size_t n_, const float* vectors_, IDNumbers& vector_ids_) {
    _pIdGenerator->getNextIDNumbers(n_, vector_ids_);
    pIndex_->add_with_ids(n_, vectors_, &vector_ids_[0]);
    for(auto i=0 ; i<n_; i++) {
        vector_ids_.push_back(i);
    }
}

size_t MemVectors::total() const {
    return pIndex_->ntotal;
}

size_t MemVectors::approximate_size() const {
    return total() * _dimension;
}

Status MemVectors::serialize(std::string& group_id) {
    /* std::stringstream ss; */
    /* ss << "/tmp/test/" << _pIdGenerator->getNextIDNumber(); */
    /* faiss::write_index(pIndex_, ss.str().c_str()); */
    /* std::cout << pIndex_->ntotal << std::endl; */
    /* std::cout << _file_location << std::endl; */
    faiss::write_index(pIndex_, _file_location.c_str());
    group_id = group_id_;
    return Status::OK();
}

MemVectors::~MemVectors() {
    if (_pIdGenerator != nullptr) {
        delete _pIdGenerator;
        _pIdGenerator = nullptr;
    }
    if (pIndex_ != nullptr) {
        delete pIndex_;
        pIndex_ = nullptr;
    }
}

/*
 * MemManager
 */

VectorsPtr MemManager::get_mem_by_group(const std::string& group_id) {
    auto memIt = _memMap.find(group_id);
    if (memIt != _memMap.end()) {
        return memIt->second;
    }

    meta::GroupFileSchema group_file;
    auto status = _pMeta->add_group_file(group_id, group_file);
    if (!status.ok()) {
        return nullptr;
    }

    _memMap[group_id] = std::shared_ptr<MemVectors>(new MemVectors(group_file.group_id,
                group_file.dimension,
                group_file.location));
    return _memMap[group_id];
}

Status MemManager::add_vectors(const std::string& group_id_,
        size_t n_,
        const float* vectors_,
        IDNumbers& vector_ids_) {
    std::unique_lock<std::mutex> lock(_mutex);
    return add_vectors_no_lock(group_id_, n_, vectors_, vector_ids_);
}

Status MemManager::add_vectors_no_lock(const std::string& group_id,
        size_t n,
        const float* vectors,
        IDNumbers& vector_ids) {
    std::shared_ptr<MemVectors> mem = get_mem_by_group(group_id);
    if (mem == nullptr) {
        return Status::NotFound("Group " + group_id + " not found!");
    }
    mem->add(n, vectors, vector_ids);

    return Status::OK();
}

Status MemManager::mark_memory_as_immutable() {
    std::unique_lock<std::mutex> lock(_mutex);
    for (auto& kv: _memMap) {
        _immMems.push_back(kv.second);
    }

    _memMap.clear();
    return Status::OK();
}

/* bool MemManager::need_serialize(double interval) { */
/*     if (_immMems.size() > 0) { */
/*         return false; */
/*     } */

/*     auto diff = std::difftime(std::time(nullptr), _last_compact_time); */
/*     if (diff >= interval) { */
/*         return true; */
/*     } */

/*     return false; */
/* } */

Status MemManager::serialize(std::vector<std::string>& group_ids) {
    mark_memory_as_immutable();
    std::string group_id;
    group_ids.clear();
    for (auto& mem : _immMems) {
        mem->serialize(group_id);
        group_ids.push_back(group_id);
    }
    _immMems.clear();
    return Status::OK();
}


} // namespace engine
} // namespace vecwise
} // namespace zilliz
