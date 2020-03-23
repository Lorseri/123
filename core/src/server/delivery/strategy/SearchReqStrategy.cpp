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

#include "server/delivery/strategy/SearchReqStrategy.h"
#include "server/delivery/request/SearchCombineRequest.h"
#include "server/delivery/request/SearchRequest.h"
#include "utils/CommonUtil.h"
#include "utils/Error.h"
#include "utils/Log.h"
#include "utils/TimeRecorder.h"

#include <queue>

namespace milvus {
namespace server {

SearchReqStrategy::SearchReqStrategy() {
}

Status
SearchReqStrategy::ReScheduleQueue(const BaseRequestPtr& request, std::queue<BaseRequestPtr>& queue) {
    if (request->GetRequestType() != BaseRequest::kSearch) {
        SERVER_LOG_ERROR << "search strategy can only handle search request";
        return Status(SERVER_UNSUPPORTED_ERROR, "");
    }

    //    TimeRecorderAuto rc("SearchReqStrategy::ReScheduleQueue");
    SearchRequestPtr new_search_req = std::static_pointer_cast<SearchRequest>(request);

    BaseRequestPtr last_req = queue.back();
    if (last_req->GetRequestType() == BaseRequest::kSearch) {
        SearchRequestPtr last_search_req = std::static_pointer_cast<SearchRequest>(last_req);
        if (SearchCombineRequest::CanCombine(last_search_req, new_search_req)) {
            // pop last request
            queue.pop();

            // combine request
            SearchCombineRequestPtr combine_request = std::make_shared<SearchCombineRequest>();
            combine_request->Combine(last_search_req);
            combine_request->Combine(new_search_req);
            queue.push(combine_request);
            SERVER_LOG_DEBUG << "Combine search request strategy";
        } else {
            // directly put to queue
            queue.push(request);
        }
    } else if (last_req->GetRequestType() == BaseRequest::kSearchCombine) {
        SearchCombineRequestPtr combine_req = std::static_pointer_cast<SearchCombineRequest>(last_req);
        if (combine_req->CanCombine(new_search_req)) {
            // combine request
            combine_req->Combine(new_search_req);
            SERVER_LOG_DEBUG << "Combine search request";
        } else {
            // directly put to queue
            queue.push(request);
        }
    } else {
        SERVER_LOG_ERROR << "search strategy can only handle search request";
        return Status(SERVER_UNSUPPORTED_ERROR, "");
    }

    return Status::OK();
}

}  // namespace server
}  // namespace milvus
