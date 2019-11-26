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

#include "server/grpc_impl/request/GrpcBaseRequest.h"

namespace milvus {
namespace server {
namespace grpc {

class DeleteByDateRequest : public GrpcBaseRequest {
 public:
    static BaseRequestPtr
    Create(const std::shared_ptr<Context>& context, const ::milvus::grpc::DeleteByDateParam* delete_by_range_param);

 protected:
    explicit DeleteByDateRequest(const std::shared_ptr<Context>& context,
                                 const ::milvus::grpc::DeleteByDateParam* delete_by_range_param);

    Status
    OnExecute() override;

 private:
    const ::milvus::grpc::DeleteByDateParam* delete_by_range_param_;
};

}  // namespace grpc
}  // namespace server
}  // namespace milvus
