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

#include "server/delivery/request/CreateCollectionReq.h"
#include "db/Utils.h"
#include "server/DBWrapper.h"
#include "server/ValidationUtil.h"
#include "utils/Log.h"
#include "utils/TimeRecorder.h"

#include <fiu-local.h>

namespace milvus {
namespace server {

CreateCollectionReq::CreateCollectionReq(const ContextPtr& context, const std::string& collection_name,
                                         FieldsType& fields, milvus::json& extra_params)
    : BaseReq(context, ReqType::kCreateCollection),
      collection_name_(collection_name),
      fields_(fields),
      extra_params_(extra_params) {
}

BaseReqPtr
CreateCollectionReq::Create(const ContextPtr& context, const std::string& collection_name, FieldsType& fields,
                            milvus::json& extra_params) {
    return std::shared_ptr<BaseReq>(new CreateCollectionReq(context, collection_name, fields, extra_params));
}

Status
CreateCollectionReq::OnExecute() {
    try {
        std::string hdr = "CreateCollectionReq(collection=" + collection_name_ + ")";
        TimeRecorderAuto rc(hdr);

        // step 1: check arguments
        auto status = ValidateCollectionName(collection_name_);
        fiu_do_on("CreateCollectionReq.OnExecute.invalid_collection_name",
                  status = Status(milvus::SERVER_UNEXPECTED_ERROR, ""));
        if (!status.ok()) {
            return status;
        }

        if (!extra_params_.contains(engine::PARAM_SEGMENT_ROW_COUNT)) {
            extra_params_[engine::PARAM_SEGMENT_ROW_COUNT] = engine::DEFAULT_SEGMENT_ROW_COUNT;
        } else {
            auto segment_row = extra_params_[engine::PARAM_SEGMENT_ROW_COUNT].get<int64_t>();
            STATUS_CHECK(ValidateSegmentRowCount(segment_row));
        }

        rc.RecordSection("check validation");

        // step 2: create snapshot collection context
        engine::snapshot::CreateCollectionContext create_collection_context;
        LOG_SERVER_DEBUG_ << "make collection_schema";
        auto collection_schema = std::make_shared<engine::snapshot::Collection>(collection_name_, extra_params_);
        if (collection_schema == nullptr) {
            LOG_SERVER_DEBUG_ << "collection_schema null";
        }

        LOG_SERVER_DEBUG_ << "create_collection_context";
        std::set<std::string> unique_field_names;
        create_collection_context.collection = collection_schema;
        for (auto& field_kv : fields_) {
            auto& field_name = field_kv.first;
            unique_field_names.insert(field_name);
            auto& field_schema = field_kv.second;

            auto& field_type = field_schema.field_type_;
            auto& field_params = field_schema.field_params_;
            auto& index_params = field_schema.index_params_;

            STATUS_CHECK(ValidateFieldName(field_name));

            std::string index_name;
            if (index_params.contains("name")) {
                index_name = index_params["name"];
            }

            LOG_SERVER_DEBUG_ << "checkout Default_UID_NAME";
            // validate id field
            if (field_name == engine::FIELD_UID) {
                if (field_type != engine::DataType::INT64) {
                    return Status(DB_ERROR, "Field '_id' data type must be int64");
                }
            }

            // validate vector field dimension
            if (field_type == engine::DataType::VECTOR_FLOAT || field_type == engine::DataType::VECTOR_BINARY) {
                if (!field_params.contains(engine::PARAM_DIMENSION)) {
                    return Status(SERVER_INVALID_VECTOR_DIMENSION, "Dimension not defined in field_params");
                } else {
                    auto dim = field_params[engine::PARAM_DIMENSION].get<int64_t>();
                    if (field_type == engine::DataType::VECTOR_FLOAT) {
                        STATUS_CHECK(ValidateDimension(dim, false));
                    } else {
                        STATUS_CHECK(ValidateDimension(dim, true));
                    }
                }
            }

            auto field = std::make_shared<engine::snapshot::Field>(field_name, 0, field_type, field_params);
            create_collection_context.fields_schema[field] = {};
        }

        // not allow duplicate field name
        if (unique_field_names.size() != fields_.size()) {
            return Status(DB_ERROR, "Duplicate field name");
        }

        // step 3: create collection
        LOG_SERVER_FATAL_ << "create collection";
        status = DBWrapper::DB()->CreateCollection(create_collection_context);
        fiu_do_on("CreateCollectionReq.OnExecute.invalid_db_execute",
                  status = Status(milvus::SERVER_UNEXPECTED_ERROR, ""));
        LOG_SERVER_FATAL_ << "create collection end";
        if (!status.ok()) {
            // collection could exist
            if (status.code() == DB_ALREADY_EXIST) {
                return Status(SERVER_INVALID_COLLECTION_NAME, status.message());
            }
            return status;
        }

        rc.ElapseFromBegin("done");
    } catch (std::exception& ex) {
        return Status(SERVER_UNEXPECTED_ERROR, ex.what());
    }

    return Status::OK();
}

}  // namespace server
}  // namespace milvus
