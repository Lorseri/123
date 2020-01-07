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

#pragma once

#include <atomic>
#include <thread>
#include <condition_variable>
//#include <src/sdk/include/MilvusApi.h>
#include "WalDefinations.h"
#include "WalFileHandler.h"
#include "WalMetaHandler.h"
#include "WalBuffer.h"

namespace milvus {
namespace engine {
namespace wal {

class WalManager {
 public:
    WalManager* GetInstance();

    void Init();
    void Start();
    void Stop();
    //todo: return error code
    bool
    Insert(const std::string &table_id,
           size_t n,
           const float *vectors,
           milvus::engine::IDNumbers &vector_ids);
    void DeleteById(const std::string& table_id, const milvus::engine::IDNumbers& vector_ids);
    void Flush(const std::string& table_id = "");
    void Apply(const uint64_t& apply_lsn);
    void Dispatch(std::string &table_id,
                  MXLogType& mxl_type,
                  size_t &n,
                  size_t &dim,
                  float *vectors,
                  milvus::engine::IDNumbers &vector_ids,
                  const uint64_t& last_applied_lsn,
                  uint64_t &lsn);

    void Recovery();

    uint64_t GetCurrentLsn();

 private:
    WalManager();
    ~WalManager();
    WalManager operator = (WalManager&);

    bool is_running_;
    MXLogConfiguration mxlog_config_;
    uint64_t last_applied_lsn_;
    MXLogBufferPtr p_buffer_;
    MXLogMetaHandlerPtr p_meta_handler_;

    std::thread reader_;

};
} // wal
} // engine
} // milvus

