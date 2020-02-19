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

#include <mutex>
#include <string>
#include <vector>

#include "codecs/VectorsFormat.h"
#include "segment/Vectors.h"

namespace milvus {
namespace codec {

class DefaultVectorsFormat : public VectorsFormat {
 public:
    DefaultVectorsFormat() = default;

    void
    read(const store::DirectoryPtr& directory_ptr, segment::VectorsPtr& vectors_read) override;

    void
    write(const store::DirectoryPtr& directory_ptr, const segment::VectorsPtr& vectors) override;

    void
    readUids(const store::DirectoryPtr& directory_ptr, std::vector<segment::doc_id_t>& uids) override;

    // No copy and move
    DefaultVectorsFormat(const DefaultVectorsFormat&) = delete;
    DefaultVectorsFormat(DefaultVectorsFormat&&) = delete;

    DefaultVectorsFormat&
    operator=(const DefaultVectorsFormat&) = delete;
    DefaultVectorsFormat&
    operator=(DefaultVectorsFormat&&) = delete;

 private:
    std::mutex mutex_;

    const std::string raw_vector_extension_ = ".rv";
    const std::string user_id_extension_ = ".uid";
};

}  // namespace codec
}  // namespace milvus
