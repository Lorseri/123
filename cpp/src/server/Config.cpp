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

#include "server/Config.h"

#include <sys/stat.h>
#include <sys/types.h>
#include <unistd.h>
#include <stdlib.h>
#include <iostream>
#include <algorithm>
#include <vector>
#include <string>

#include "config/ConfigMgr.h"
#include "utils/CommonUtil.h"
#include "utils/ValidationUtil.h"

namespace zilliz {
namespace milvus {
namespace server {

constexpr uint64_t GB = 1UL << 30;

Config &
Config::GetInstance() {
    static Config config_inst;
    return config_inst;
}

Status
Config::LoadConfigFile(const std::string &filename) {
    if (filename.empty()) {
        std::cerr << "ERROR: need specify config file" << std::endl;
        exit(1);
    }
    struct stat dirStat;
    int statOK = stat(filename.c_str(), &dirStat);
    if (statOK != 0) {
        std::cerr << "ERROR: Config file not exist: " << filename << std::endl;
        exit(1);
    }

    try {
        ConfigMgr *mgr = const_cast<ConfigMgr *>(ConfigMgr::GetInstance());
        ErrorCode err = mgr->LoadConfigFile(filename);
        if (err != 0) {
            std::cerr << "Server failed to load config file: " << filename << std::endl;
            exit(1);
        }
    }
    catch (YAML::Exception &e) {
        std::cerr << "Server failed to load config file: " << filename << std::endl;
        exit(1);
    }

    return Status::OK();
}

void
Config::PrintConfigSection(const std::string &config_node_name) {
    std::cout << std::endl;
    std::cout << config_node_name << ":" << std::endl;
    if (config_map_.find(config_node_name) != config_map_.end()) {
        for (auto item : config_map_[config_node_name]) {
            std::cout << item.first << ": " << item.second << std::endl;
        }
    }
}

void
Config::PrintAll() {
    PrintConfigSection(CONFIG_SERVER);
    PrintConfigSection(CONFIG_DB);
    PrintConfigSection(CONFIG_CACHE);
    PrintConfigSection(CONFIG_METRIC);
    PrintConfigSection(CONFIG_ENGINE);
    PrintConfigSection(CONFIG_RESOURCE);
}

////////////////////////////////////////////////////////////////////////////////
Status
Config::CheckServerConfigAddress(const std::string &value) {
    if (!ValidationUtil::ValidateIpAddress(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid server config address: " + value);
    }
    return Status::OK();
}

Status
Config::CheckServerConfigPort(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsNumber(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid server config port: " + value);
    } else {
        int32_t port = std::stoi(value);
        if (!(port > 1024 && port < 65535)) {
            return Status(SERVER_INVALID_ARGUMENT, "Server config port out of range (1024, 65535): " + value);
        }
    }
    return Status::OK();
}

Status
Config::CheckServerConfigMode(const std::string &value) {
    if (value != "single" && value != "cluster" && value != "read_only") {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid server config mode [single, cluster, read_only]: " + value);
    }
    return Status::OK();
}

Status
Config::CheckServerConfigTimeZone(const std::string &value) {
    if (value.length() <= 3) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid server config time_zone: " + value);
    } else {
        if (value.substr(0, 3) != "UTC") {
            return Status(SERVER_INVALID_ARGUMENT, "Invalid server config time_zone: " + value);
        } else {
            try {
                stoi(value.substr(3));
            } catch (...) {
                return Status(SERVER_INVALID_ARGUMENT, "Invalid server config time_zone: " + value);
            }
        }
    }
    return Status::OK();
}

Status
Config::CheckDBConfigPath(const std::string &value) {
    if (value.empty()) {
        return Status(SERVER_INVALID_ARGUMENT, "DB config path empty");
    }
    return Status::OK();
}

Status
Config::CheckDBConfigSlavePath(const std::string &value) {
    return Status::OK();
}

Status
Config::CheckDBConfigBackendUrl(const std::string &value) {
    if (!ValidationUtil::ValidateDbURI(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid DB config backend_url: " + value);
    }
    return Status::OK();
}

Status
Config::CheckDBConfigArchiveDiskThreshold(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsNumber(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid DB config archive_disk_threshold: " + value);
    }
    return Status::OK();
}

Status
Config::CheckDBConfigArchiveDaysThreshold(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsNumber(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid DB config archive_days_threshold: " + value);
    }
    return Status::OK();
}

Status
Config::CheckDBConfigBufferSize(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsNumber(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid DB config buffer_size: " + value);
    } else {
        int64_t buffer_size = std::stoi(value) * GB;
        uint64_t total_mem = 0, free_mem = 0;
        CommonUtil::GetSystemMemInfo(total_mem, free_mem);
        if (buffer_size >= total_mem) {
            return Status(SERVER_INVALID_ARGUMENT, "DB config buffer_size exceed system memory: " + value);
        }
    }
    return Status::OK();
}

Status
Config::CheckDBConfigBuildIndexGPU(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsNumber(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid DB config build_index_gpu: " + value);
    } else {
        int32_t gpu_index = std::stoi(value);
        if (!ValidationUtil::ValidateGpuIndex(gpu_index).ok()) {
            return Status(SERVER_INVALID_ARGUMENT, "Invalid DB config build_index_gpu: " + value);
        }
    }
    return Status::OK();
}

Status
Config::CheckMetricConfigAutoBootup(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsBool(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid metric config auto_bootup: " + value);
    }
    return Status::OK();
}

Status
Config::CheckMetricConfigCollector(const std::string &value) {
    if (value != "prometheus") {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid metric config collector: " + value);
    }
    return Status::OK();
}

Status
Config::CheckMetricConfigPrometheusPort(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsNumber(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid metric config prometheus_port: " + value);
    }
    return Status::OK();
}

Status
Config::CheckCacheConfigCpuMemCapacity(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsNumber(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid cache config cpu_mem_capacity: " + value);
    } else {
        uint64_t cpu_cache_capacity = std::stoi(value) * GB;
        uint64_t total_mem = 0, free_mem = 0;
        CommonUtil::GetSystemMemInfo(total_mem, free_mem);
        if (cpu_cache_capacity >= total_mem) {
            return Status(SERVER_INVALID_ARGUMENT, "Cache config cpu_mem_capacity exceed system memory: " + value);
        } else if (cpu_cache_capacity > (double) total_mem * 0.9) {
            std::cerr << "Warning: cpu_mem_capacity value is too big" << std::endl;
        }

        int32_t buffer_size;
        Status s = GetDBConfigBufferSize(buffer_size);
        if (!s.ok()) return s;
        int64_t insert_buffer_size = buffer_size * GB;
        if (insert_buffer_size + cpu_cache_capacity >= total_mem) {
            return Status(SERVER_INVALID_ARGUMENT, "Sum of cpu_mem_capacity and buffer_size exceed system memory");
        }
    }
    return Status::OK();
}

Status
Config::CheckCacheConfigCpuMemThreshold(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsFloat(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid cache config cpu_mem_threshold: " + value);
    } else {
        float cpu_mem_threshold = std::stof(value);
        if (cpu_mem_threshold <= 0.0 || cpu_mem_threshold >= 1.0) {
            return Status(SERVER_INVALID_ARGUMENT, "Invalid cache config cpu_mem_threshold: " + value);
        }
    }
    return Status::OK();
}

Status
Config::CheckCacheConfigGpuMemCapacity(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsNumber(value).ok()) {
        std::cerr << "ERROR: gpu_cache_capacity " << value << " is not a number" << std::endl;
    } else {
        uint64_t gpu_cache_capacity = std::stoi(value) * GB;
        int gpu_index;
        Status s = GetDBConfigBuildIndexGPU(gpu_index);
        if (!s.ok()) return s;
        size_t gpu_memory;
        if (!ValidationUtil::GetGpuMemory(gpu_index, gpu_memory).ok()) {
            return Status(SERVER_UNEXPECTED_ERROR,
                          "Fail to get GPU memory for GPU device: " + std::to_string(gpu_index));
        } else if (gpu_cache_capacity >= gpu_memory) {
            return Status(SERVER_INVALID_ARGUMENT,
                          "Cache config gpu_mem_capacity exceed GPU memory: " + std::to_string(gpu_memory));
        } else if (gpu_cache_capacity > (double) gpu_memory * 0.9) {
            std::cerr << "Warning: gpu_mem_capacity value is too big" << std::endl;
        }
    }
    return Status::OK();
}

Status
Config::CheckCacheConfigGpuMemThreshold(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsFloat(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid cache config gpu_mem_threshold: " + value);
    } else {
        float gpu_mem_threshold = std::stof(value);
        if (gpu_mem_threshold <= 0.0 || gpu_mem_threshold >= 1.0) {
            return Status(SERVER_INVALID_ARGUMENT, "Invalid cache config gpu_mem_threshold: " + value);
        }
    }
    return Status::OK();
}

Status
Config::CheckCacheConfigCacheInsertData(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsBool(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid cache config cache_insert_data: " + value);
    }
    return Status::OK();
}

Status
Config::CheckEngineConfigBlasThreshold(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsNumber(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid engine config blas threshold: " + value);
    }
    return Status::OK();
}

Status
Config::CheckEngineConfigOmpThreadNum(const std::string &value) {
    if (!ValidationUtil::ValidateStringIsNumber(value).ok()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid engine config omp_thread_num: " + value);
    } else {
        int32_t omp_thread = std::stoi(value);
        uint32_t sys_thread_cnt = 8;
        if (omp_thread > CommonUtil::GetSystemAvailableThreads(sys_thread_cnt)) {
            return Status(SERVER_INVALID_ARGUMENT, "Invalid engine config omp_thread_num: " + value);
        }
    }
    return Status::OK();
}

Status
Config::CheckResourceConfigMode(const std::string &value) {
    if (value != "simple") {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid resource config mode: " + value);
    }
    return Status::OK();
}

Status
Config::CheckResourceConfigPool(const std::vector<std::string> &value) {
    if (value.empty()) {
        return Status(SERVER_INVALID_ARGUMENT, "Invalid resource config pool");
    }
    return Status::OK();
}

////////////////////////////////////////////////////////////////////////////////
ConfigNode &
Config::GetConfigNode(const std::string &name) {
    ConfigMgr *mgr = ConfigMgr::GetInstance();
    ConfigNode &root_node = mgr->GetRootNode();
    return root_node.GetChild(name);
}

Status
Config::GetConfigValueInMem(const std::string &parent_key,
                            const std::string &child_key,
                            std::string &value) {
    std::lock_guard<std::mutex> lock(mutex_);
    if (config_map_.find(parent_key) != config_map_.end() &&
        config_map_[parent_key].find(child_key) != config_map_[parent_key].end()) {
        value = config_map_[parent_key][child_key];
        return Status::OK();
    } else {
        return Status(SERVER_UNEXPECTED_ERROR, "key not exist");
    }
}

void
Config::SetConfigValueInMem(const std::string &parent_key,
                            const std::string &child_key,
                            const std::string &value) {
    std::lock_guard<std::mutex> lock(mutex_);
    config_map_[parent_key][child_key] = value;
}

////////////////////////////////////////////////////////////////////////////////
/* server config */
std::string
Config::GetServerConfigStrAddress() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_ADDRESS, value).ok()) {
        value = GetConfigNode(CONFIG_SERVER).GetValue(CONFIG_SERVER_ADDRESS,
                                                      CONFIG_SERVER_ADDRESS_DEFAULT);
        SetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_ADDRESS, value);
    }
    return value;
}

std::string
Config::GetServerConfigStrPort() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_PORT, value).ok()) {
        value = GetConfigNode(CONFIG_SERVER).GetValue(CONFIG_SERVER_PORT,
                                                      CONFIG_SERVER_PORT_DEFAULT);
        SetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_PORT, value);
    }
    return value;
}

std::string
Config::GetServerConfigStrMode() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_MODE, value).ok()) {
        value = GetConfigNode(CONFIG_SERVER).GetValue(CONFIG_SERVER_MODE,
                                                      CONFIG_SERVER_MODE_DEFAULT);
        SetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_MODE, value);
    }
    return value;
}

std::string
Config::GetServerConfigStrTimeZone() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_TIME_ZONE, value).ok()) {
        value = GetConfigNode(CONFIG_SERVER).GetValue(CONFIG_SERVER_TIME_ZONE,
                                                      CONFIG_SERVER_TIME_ZONE_DEFAULT);
        SetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_TIME_ZONE, value);
    }
    return value;
}

////////////////////////////////////////////////////////////////////////////////
/* db config */
std::string
Config::GetDBConfigStrPath() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_DB, CONFIG_DB_PATH, value).ok()) {
        value = GetConfigNode(CONFIG_DB).GetValue(CONFIG_DB_PATH,
                                                  CONFIG_DB_PATH_DEFAULT);
        SetConfigValueInMem(CONFIG_DB, CONFIG_DB_PATH, value);
    }
    return value;
}

std::string
Config::GetDBConfigStrSlavePath() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_DB, CONFIG_DB_SLAVE_PATH, value).ok()) {
        value = GetConfigNode(CONFIG_DB).GetValue(CONFIG_DB_SLAVE_PATH,
                                                  CONFIG_DB_SLAVE_PATH_DEFAULT);
        SetConfigValueInMem(CONFIG_DB, CONFIG_DB_SLAVE_PATH, value);
    }
    return value;
}

std::string
Config::GetDBConfigStrBackendUrl() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_DB, CONFIG_DB_BACKEND_URL, value).ok()) {
        value = GetConfigNode(CONFIG_DB).GetValue(CONFIG_DB_BACKEND_URL,
                                                  CONFIG_DB_BACKEND_URL_DEFAULT);
        SetConfigValueInMem(CONFIG_DB, CONFIG_DB_BACKEND_URL, value);
    }
    return value;
}

std::string
Config::GetDBConfigStrArchiveDiskThreshold() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_DB, CONFIG_DB_ARCHIVE_DISK_THRESHOLD, value).ok()) {
        value = GetConfigNode(CONFIG_DB).GetValue(CONFIG_DB_ARCHIVE_DISK_THRESHOLD,
                                                  CONFIG_DB_ARCHIVE_DISK_THRESHOLD_DEFAULT);
        SetConfigValueInMem(CONFIG_DB, CONFIG_DB_ARCHIVE_DISK_THRESHOLD, value);
    }
    return value;
}

std::string
Config::GetDBConfigStrArchiveDaysThreshold() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_DB, CONFIG_DB_ARCHIVE_DAYS_THRESHOLD, value).ok()) {
        value = GetConfigNode(CONFIG_DB).GetValue(CONFIG_DB_ARCHIVE_DAYS_THRESHOLD,
                                                  CONFIG_DB_ARCHIVE_DAYS_THRESHOLD_DEFAULT);
        SetConfigValueInMem(CONFIG_DB, CONFIG_DB_ARCHIVE_DAYS_THRESHOLD, value);
    }
    return value;
}

std::string
Config::GetDBConfigStrBufferSize() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_DB, CONFIG_DB_BUFFER_SIZE, value).ok()) {
        value = GetConfigNode(CONFIG_DB).GetValue(CONFIG_DB_BUFFER_SIZE,
                                                  CONFIG_DB_BUFFER_SIZE_DEFAULT);
        SetConfigValueInMem(CONFIG_DB, CONFIG_DB_BUFFER_SIZE, value);
    }
    return value;
}

std::string
Config::GetDBConfigStrBuildIndexGPU() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_DB, CONFIG_DB_BUILD_INDEX_GPU, value).ok()) {
        value = GetConfigNode(CONFIG_DB).GetValue(CONFIG_DB_BUILD_INDEX_GPU,
                                                  CONFIG_DB_BUILD_INDEX_GPU_DEFAULT);
        SetConfigValueInMem(CONFIG_DB, CONFIG_DB_BUILD_INDEX_GPU, value);
    }
    return value;
}

////////////////////////////////////////////////////////////////////////////////
/* metric config */
std::string
Config::GetMetricConfigStrAutoBootup() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_METRIC, CONFIG_METRIC_AUTO_BOOTUP, value).ok()) {
        value = GetConfigNode(CONFIG_METRIC).GetValue(CONFIG_METRIC_AUTO_BOOTUP,
                                                      CONFIG_METRIC_AUTO_BOOTUP_DEFAULT);
        SetConfigValueInMem(CONFIG_METRIC, CONFIG_METRIC_AUTO_BOOTUP, value);
    }
    return value;
}

std::string
Config::GetMetricConfigStrCollector() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_METRIC, CONFIG_METRIC_COLLECTOR, value).ok()) {
        value = GetConfigNode(CONFIG_METRIC).GetValue(CONFIG_METRIC_COLLECTOR,
                                                      CONFIG_METRIC_COLLECTOR_DEFAULT);
        SetConfigValueInMem(CONFIG_METRIC, CONFIG_METRIC_COLLECTOR, value);
    }
    return value;
}

std::string
Config::GetMetricConfigStrPrometheusPort() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_METRIC, CONFIG_METRIC_PROMETHEUS_PORT, value).ok()) {
        value = GetConfigNode(CONFIG_METRIC).GetValue(CONFIG_METRIC_PROMETHEUS_PORT,
                                                      CONFIG_METRIC_PROMETHEUS_PORT_DEFAULT);
        SetConfigValueInMem(CONFIG_METRIC, CONFIG_METRIC_PROMETHEUS_PORT, value);
    }
    return value;
}

////////////////////////////////////////////////////////////////////////////////
/* cache config */
std::string
Config::GetCacheConfigStrCpuMemCapacity() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_CACHE, CONFIG_CACHE_CPU_MEM_CAPACITY, value).ok()) {
        value = GetConfigNode(CONFIG_CACHE).GetValue(CONFIG_CACHE_CPU_MEM_CAPACITY,
                                                     CONFIG_CACHE_CPU_MEM_CAPACITY_DEFAULT);
        SetConfigValueInMem(CONFIG_CACHE, CONFIG_CACHE_CPU_MEM_CAPACITY, value);
    }
    return value;
}

std::string
Config::GetCacheConfigStrCpuMemThreshold() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_CACHE, CONFIG_CACHE_CPU_MEM_THRESHOLD, value).ok()) {
        value = GetConfigNode(CONFIG_CACHE).GetValue(CONFIG_CACHE_CPU_MEM_THRESHOLD,
                                                     CONFIG_CACHE_CPU_MEM_THRESHOLD_DEFAULT);
        SetConfigValueInMem(CONFIG_CACHE, CONFIG_CACHE_CPU_MEM_THRESHOLD, value);
    }
    return value;
}

std::string
Config::GetCacheConfigStrGpuMemCapacity() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_CACHE, CONFIG_CACHE_GPU_MEM_CAPACITY, value).ok()) {
        value = GetConfigNode(CONFIG_CACHE).GetValue(CONFIG_CACHE_GPU_MEM_CAPACITY,
                                                     CONFIG_CACHE_GPU_MEM_CAPACITY_DEFAULT);
        SetConfigValueInMem(CONFIG_CACHE, CONFIG_CACHE_GPU_MEM_CAPACITY, value);
    }
    return value;
}

std::string
Config::GetCacheConfigStrGpuMemThreshold() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_CACHE, CONFIG_CACHE_GPU_MEM_THRESHOLD, value).ok()) {
        value = GetConfigNode(CONFIG_CACHE).GetValue(CONFIG_CACHE_GPU_MEM_THRESHOLD,
                                                     CONFIG_CACHE_GPU_MEM_THRESHOLD_DEFAULT);
        SetConfigValueInMem(CONFIG_CACHE, CONFIG_CACHE_GPU_MEM_THRESHOLD, value);
    }
    return value;
}

std::string
Config::GetCacheConfigStrCacheInsertData() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_CACHE, CONFIG_CACHE_CACHE_INSERT_DATA, value).ok()) {
        value = GetConfigNode(CONFIG_CACHE).GetValue(CONFIG_CACHE_CACHE_INSERT_DATA,
                                                     CONFIG_CACHE_CACHE_INSERT_DATA_DEFAULT);
        SetConfigValueInMem(CONFIG_CACHE, CONFIG_CACHE_CACHE_INSERT_DATA, value);
    }
    return value;
}

////////////////////////////////////////////////////////////////////////////////
/* engine config */
std::string
Config::GetEngineConfigStrBlasThreshold() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_ENGINE, CONFIG_ENGINE_BLAS_THRESHOLD, value).ok()) {
        value = GetConfigNode(CONFIG_ENGINE).GetValue(CONFIG_ENGINE_BLAS_THRESHOLD,
                                                      CONFIG_ENGINE_BLAS_THRESHOLD_DEFAULT);
        SetConfigValueInMem(CONFIG_ENGINE, CONFIG_ENGINE_BLAS_THRESHOLD, value);
    }
    return value;
}

std::string
Config::GetEngineConfigStrOmpThreadNum() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_ENGINE, CONFIG_ENGINE_OMP_THREAD_NUM, value).ok()) {
        value = GetConfigNode(CONFIG_ENGINE).GetValue(CONFIG_ENGINE_OMP_THREAD_NUM,
                                                      CONFIG_ENGINE_OMP_THREAD_NUM_DEFAULT);
        SetConfigValueInMem(CONFIG_ENGINE, CONFIG_ENGINE_OMP_THREAD_NUM, value);
    }
    return value;
}

////////////////////////////////////////////////////////////////////////////////
/* resource config */
std::string
Config::GetResourceConfigStrMode() {
    std::string value;
    if (!GetConfigValueInMem(CONFIG_RESOURCE, CONFIG_RESOURCE_MODE, value).ok()) {
        value = GetConfigNode(CONFIG_RESOURCE).GetValue(CONFIG_RESOURCE_MODE,
                                                        CONFIG_RESOURCE_MODE_DEFAULT);
        SetConfigValueInMem(CONFIG_RESOURCE, CONFIG_RESOURCE_MODE, value);
    }
    return value;
}

////////////////////////////////////////////////////////////////////////////////
Status
Config::GetServerConfigAddress(std::string &value) {
    value = GetServerConfigStrAddress();
    return CheckServerConfigAddress(value);
}

Status
Config::GetServerConfigPort(std::string &value) {
    value = GetServerConfigStrPort();
    return CheckServerConfigPort(value);
}

Status
Config::GetServerConfigMode(std::string &value) {
    value = GetServerConfigStrMode();
    return CheckServerConfigMode(value);
}

Status
Config::GetServerConfigTimeZone(std::string &value) {
    value = GetServerConfigStrTimeZone();
    return CheckServerConfigTimeZone(value);
}

Status
Config::GetDBConfigPath(std::string &value) {
    value = GetDBConfigStrPath();
    return CheckDBConfigPath(value);
}

Status
Config::GetDBConfigSlavePath(std::string &value) {
    value = GetDBConfigStrSlavePath();
    return Status::OK();
}

Status
Config::GetDBConfigBackendUrl(std::string &value) {
    value = GetDBConfigStrBackendUrl();
    return CheckDBConfigBackendUrl(value);
}

Status
Config::GetDBConfigArchiveDiskThreshold(int32_t &value) {
    std::string str = GetDBConfigStrArchiveDiskThreshold();
    Status s = CheckDBConfigArchiveDiskThreshold(str);
    if (!s.ok()) return s;
    value = std::stoi(str);
    return Status::OK();
}

Status
Config::GetDBConfigArchiveDaysThreshold(int32_t &value) {
    std::string str = GetDBConfigStrArchiveDaysThreshold();
    Status s = CheckDBConfigArchiveDaysThreshold(str);
    if (!s.ok()) return s;
    value = std::stoi(str);
    return Status::OK();
}

Status
Config::GetDBConfigBufferSize(int32_t &value) {
    std::string str = GetDBConfigStrBufferSize();
    Status s = CheckDBConfigBufferSize(str);
    if (!s.ok()) return s;
    value = std::stoi(str);
    return Status::OK();
}

Status
Config::GetDBConfigBuildIndexGPU(int32_t &value) {
    std::string str = GetDBConfigStrBuildIndexGPU();
    Status s = CheckDBConfigBuildIndexGPU(str);
    if (!s.ok()) return s;
    value = std::stoi(str);
    return Status::OK();
}

Status
Config::GetMetricConfigAutoBootup(bool &value) {
    std::string str = GetMetricConfigStrAutoBootup();
    Status s = CheckMetricConfigPrometheusPort(str);
    if (!s.ok()) return s;
    std::transform(str.begin(), str.end(), str.begin(), ::tolower);
    value = (str == "true" || str == "on" || str == "yes" || str == "1");
    return Status::OK();
}

Status
Config::GetMetricConfigCollector(std::string &value) {
    value = GetMetricConfigStrCollector();
    return Status::OK();
}

Status
Config::GetMetricConfigPrometheusPort(std::string &value) {
    value = GetMetricConfigStrPrometheusPort();
    return CheckMetricConfigPrometheusPort(value);
}

Status
Config::GetCacheConfigCpuMemCapacity(int32_t &value) {
    std::string str = GetCacheConfigStrCpuMemCapacity();
    Status s = CheckCacheConfigCpuMemCapacity(str);
    if (!s.ok()) return s;
    value = std::stoi(str);
    return Status::OK();
}

Status
Config::GetCacheConfigCpuMemThreshold(float &value) {
    std::string str = GetCacheConfigStrCpuMemThreshold();
    Status s = CheckCacheConfigCpuMemThreshold(str);
    if (!s.ok()) return s;
    value = std::stof(str);
    return Status::OK();
}

Status
Config::GetCacheConfigGpuMemCapacity(int32_t &value) {
    std::string str = GetCacheConfigStrGpuMemCapacity();
    Status s = CheckCacheConfigGpuMemCapacity(str);
    if (!s.ok()) return s;
    value = std::stoi(str);
    return Status::OK();
}

Status
Config::GetCacheConfigGpuMemThreshold(float &value) {
    std::string str = GetCacheConfigStrGpuMemThreshold();
    Status s = CheckCacheConfigGpuMemThreshold(str);
    if (!s.ok()) return s;
    value = std::stof(str);
    return Status::OK();
}

Status
Config::GetCacheConfigCacheInsertData(bool &value) {
    std::string str = GetCacheConfigStrCacheInsertData();
    Status s = CheckCacheConfigCacheInsertData(str);
    if (!s.ok()) return s;
    std::transform(str.begin(), str.end(), str.begin(), ::tolower);
    value = (str == "true" || str == "on" || str == "yes" || str == "1");
    return Status::OK();
}

Status
Config::GetEngineConfigBlasThreshold(int32_t &value) {
    std::string str = GetEngineConfigStrBlasThreshold();
    Status s = CheckEngineConfigBlasThreshold(str);
    if (!s.ok()) return s;
    value = std::stoi(str);
    return Status::OK();
}

Status
Config::GetEngineConfigOmpThreadNum(int32_t &value) {
    std::string str = GetEngineConfigStrOmpThreadNum();
    Status s = CheckEngineConfigOmpThreadNum(str);
    if (!s.ok()) return s;
    value = std::stoi(str);
    return Status::OK();
}

Status
Config::GetResourceConfigMode(std::string &value) {
    value = GetResourceConfigStrMode();
    return CheckResourceConfigMode(value);
}

Status
Config::GetResourceConfigPool(std::vector<std::string> &value) {
    ConfigNode resource_config = GetConfigNode(CONFIG_RESOURCE);
    value = resource_config.GetSequence(CONFIG_RESOURCE_POOL);
    return CheckResourceConfigPool(value);
}

///////////////////////////////////////////////////////////////////////////////
/* server config */
Status
Config::SetServerConfigAddress(const std::string &value) {
    Status s = CheckServerConfigAddress(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_ADDRESS, value);
    return Status::OK();
}

Status
Config::SetServerConfigPort(const std::string &value) {
    Status s = CheckServerConfigPort(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_PORT, value);
    return Status::OK();
}

Status
Config::SetServerConfigMode(const std::string &value) {
    Status s = CheckServerConfigMode(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_MODE, value);
    return Status::OK();
}

Status
Config::SetServerConfigTimeZone(const std::string &value) {
    Status s = CheckServerConfigTimeZone(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_SERVER, CONFIG_SERVER_TIME_ZONE, value);
    return Status::OK();
}

/* db config */
Status
Config::SetDBConfigPath(const std::string &value) {
    Status s = CheckDBConfigPath(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_DB_PATH, value);
    return Status::OK();
}

Status
Config::SetDBConfigSlavePath(const std::string &value) {
    Status s = CheckDBConfigSlavePath(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_DB_SLAVE_PATH, value);
    return Status::OK();
}

Status
Config::SetDBConfigBackendUrl(const std::string &value) {
    Status s = CheckDBConfigBackendUrl(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_DB_BACKEND_URL, value);
    return Status::OK();
}

Status
Config::SetDBConfigArchiveDiskThreshold(const std::string &value) {
    Status s = CheckDBConfigArchiveDiskThreshold(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_DB_ARCHIVE_DISK_THRESHOLD, value);
    return Status::OK();
}

Status
Config::SetDBConfigArchiveDaysThreshold(const std::string &value) {
    Status s = CheckDBConfigArchiveDaysThreshold(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_DB_ARCHIVE_DAYS_THRESHOLD, value);
    return Status::OK();
}

Status
Config::SetDBConfigBufferSize(const std::string &value) {
    Status s = CheckDBConfigBufferSize(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_DB_BUFFER_SIZE, value);
    return Status::OK();
}

Status
Config::SetDBConfigBuildIndexGPU(const std::string &value) {
    Status s = CheckDBConfigBuildIndexGPU(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_DB_BUILD_INDEX_GPU, value);
    return Status::OK();
}

/* metric config */
Status
Config::SetMetricConfigAutoBootup(const std::string &value) {
    Status s = CheckMetricConfigAutoBootup(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_METRIC_AUTO_BOOTUP, value);
    return Status::OK();
}

Status
Config::SetMetricConfigCollector(const std::string &value) {
    Status s = CheckMetricConfigCollector(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_METRIC_COLLECTOR, value);
    return Status::OK();
}

Status
Config::SetMetricConfigPrometheusPort(const std::string &value) {
    Status s = CheckMetricConfigPrometheusPort(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_METRIC_PROMETHEUS_PORT, value);
    return Status::OK();
}

/* cache config */
Status
Config::SetCacheConfigCpuMemCapacity(const std::string &value) {
    Status s = CheckCacheConfigCpuMemCapacity(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_CACHE_CPU_MEM_CAPACITY, value);
    return Status::OK();
}

Status
Config::SetCacheConfigCpuMemThreshold(const std::string &value) {
    Status s = CheckCacheConfigCpuMemThreshold(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_CACHE_CPU_MEM_THRESHOLD, value);
    return Status::OK();
}

Status
Config::SetCacheConfigGpuMemCapacity(const std::string &value) {
    Status s = CheckCacheConfigGpuMemCapacity(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_CACHE_GPU_MEM_CAPACITY, value);
    return Status::OK();
}

Status
Config::SetCacheConfigGpuMemThreshold(const std::string &value) {
    Status s = CheckCacheConfigGpuMemThreshold(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_CACHE_GPU_MEM_THRESHOLD, value);
    return Status::OK();
}

Status
Config::SetCacheConfigCacheInsertData(const std::string &value) {
    Status s = CheckCacheConfigCacheInsertData(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_CACHE_CACHE_INSERT_DATA, value);
    return Status::OK();
}

/* engine config */
Status
Config::SetEngineConfigBlasThreshold(const std::string &value) {
    Status s = CheckEngineConfigBlasThreshold(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_ENGINE_BLAS_THRESHOLD, value);
    return Status::OK();
}

Status
Config::SetEngineConfigOmpThreadNum(const std::string &value) {
    Status s = CheckEngineConfigOmpThreadNum(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_ENGINE_OMP_THREAD_NUM, value);
    return Status::OK();
}

/* resource config */
Status
Config::SetResourceConfigMode(const std::string &value) {
    Status s = CheckResourceConfigMode(value);
    if (!s.ok()) return s;
    SetConfigValueInMem(CONFIG_DB, CONFIG_RESOURCE_MODE, value);
    return Status::OK();
}

} // namespace server
} // namespace milvus
} // namespace zilliz
