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

#include "server/delivery/request/GetCollectionInfoReq.h"
#include "db/Utils.h"
#include "server/DBWrapper.h"
#include "server/web_impl/Constants.h"
#include "utils/Log.h"
#include "utils/TimeRecorder.h"

#include <fiu-local.h>
#include <memory>
#include <string>
#include <unordered_map>
#include <utility>
#include <vector>

namespace milvus {
namespace server {

GetCollectionInfoReq::GetCollectionInfoReq(const std::shared_ptr<milvus::server::Context>& context,
                                           const std::string& collection_name, CollectionSchema& collection_schema)
    : BaseReq(context, BaseReq::kGetCollectionInfo),
      collection_name_(collection_name),
      collection_schema_(collection_schema) {
}

BaseReqPtr
GetCollectionInfoReq::Create(const std::shared_ptr<milvus::server::Context>& context,
                             const std::string& collection_name, CollectionSchema& collection_schema) {
    return std::shared_ptr<BaseReq>(new GetCollectionInfoReq(context, collection_name, collection_schema));
}

Status
GetCollectionInfoReq::OnExecute() {
    std::string hdr = "GetCollectionInfoReq(collection=" + collection_name_ + ")";
    TimeRecorderAuto rc(hdr);

    try {
        engine::snapshot::CollectionPtr collection;
        engine::snapshot::CollectionMappings collection_mappings;
        STATUS_CHECK(DBWrapper::DB()->GetCollectionInfo(collection_name_, collection, collection_mappings));

        collection_schema_.collection_name_ = collection_name_;
        collection_schema_.extra_params_ = collection->GetParams();
        for (auto& field_kv : collection_mappings) {
            auto field = field_kv.first;
            if (field->GetFtype() == (engine::snapshot::FTYPE_TYPE)engine::meta::hybrid::DataType::UID) {
                continue;
            }

            milvus::json field_index_param;
            auto field_elements = field_kv.second;
            for (const auto& element : field_elements) {
                if (element->GetFtype() == (engine::snapshot::FTYPE_TYPE)engine::FieldElementType::FET_INDEX) {
                    field_index_param = element->GetParams();
                    break;
                }
            }

            auto field_name = field->GetName();
            FieldSchema field_schema;
            field_schema.field_type_ = (engine::meta::hybrid::DataType)field->GetFtype();
            field_schema.field_params_ = field->GetParams();
            field_schema.index_params_ = field_index_param;

            collection_schema_.fields_.insert(std::make_pair(field_name, field_schema));
        }

        rc.ElapseFromBegin("done");
    } catch (std::exception& ex) {
        return Status(SERVER_UNEXPECTED_ERROR, ex.what());
    }

    return Status::OK();
}

}  // namespace server
}  // namespace milvus
