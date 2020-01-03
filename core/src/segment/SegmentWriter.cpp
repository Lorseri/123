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

#include "SegmentWriter.h"

#include <memory>

namespace milvus {
namespace segment {

SegmentWriter::SegmentWriter(const std::string& directory) : directory_(directory) {
}

Status
SegmentWriter::AddVectors(const std::string& field_name, const std::vector<uint8_t>& data,
                          const std::vector<doc_id_t>& uids) {
    auto found = vectors_ptr_->vectors.find(field_name);
    if (found == vectors_ptr_->vectors.end()) {
        vectors_ptr_->vectors[field_name] = std::make_shared<Vector>(data, uids);
    } else {
        vectors_ptr_->vectors[field_name]->AddData(data);
        vectors_ptr_->vectors[field_name]->AddUids(uids);
    }
    return Status::OK();
}

Status
SegmentWriter::Serialize() {
    // TODO
    return Status::OK();
}

Status
SegmentWriter::Cache() {
    // TODO
    return Status::OK();
}

}  // namespace segment
}  // namespace milvus
