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

#include "server/delivery/request/DeleteEntityByIDReq.h"
#include "src/server/delivery/ReqScheduler.h"
#include "server/MessageWrapper.h"
#include "server/MetaWrapper.h"

#include <memory>
#include <string>
#include <unordered_map>
#include <vector>

#include "server/ValidationUtil.h"
#include "utils/Log.h"
#include "utils/TimeRecorder.h"

namespace milvus {
namespace server {

DeleteEntityByIDReq::DeleteEntityByIDReq(const ContextPtr& context, const ::milvus::grpc::DeleteByIDParam *request)
    : BaseReq(context, ReqType::kDeleteEntityByID), request_(request) {
}

BaseReqPtr
DeleteEntityByIDReq::Create(const ContextPtr& context, const ::milvus::grpc::DeleteByIDParam *request) {
    return std::shared_ptr<BaseReq>(new DeleteEntityByIDReq(context, request));
}

Status
DeleteEntityByIDReq::OnExecute() {
  auto &msg_client = MessageWrapper::GetInstance().MessageClient();
  auto segment_id = [](const std::string &collection_name,
                       uint64_t channel_id,
                       uint64_t timestamp) {
    return MetaWrapper::GetInstance().AskSegmentId(collection_name, channel_id, timestamp);
  };
  Status status = msg_client->SendMutMessage(*request_, timestamp_, segment_id);
  return status;
}

Status DeleteEntityByIDReq::OnPostExecute() {
  ReqScheduler::GetInstance().UpdateLatestDeliveredReqTime(timestamp_);
  return Status::OK();
}

}  // namespace server
}  // namespace milvus
