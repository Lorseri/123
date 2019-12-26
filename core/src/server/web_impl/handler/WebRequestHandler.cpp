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

#include "server/web_impl/handler/WebRequestHandler.h"

#include <string>
#include <vector>
#include <cmath>
#include <boost/algorithm/string.hpp>

#include "server/delivery/request/BaseRequest.h"
#include "server/web_impl/Types.h"
#include "server/web_impl/dto/PartitionDto.hpp"

#include "server/Config.h"

#include "metrics/SystemInfo.h"

namespace milvus {
namespace server {
namespace web {

StatusCode
WebErrorMap(ErrorCode code) {
    static const std::map<ErrorCode, StatusCode> code_map = {
        {SERVER_UNEXPECTED_ERROR, StatusCode::UNEXPECTED_ERROR},
        {SERVER_UNSUPPORTED_ERROR, StatusCode::UNEXPECTED_ERROR},
        {SERVER_NULL_POINTER, StatusCode::UNEXPECTED_ERROR},
        {SERVER_INVALID_ARGUMENT, StatusCode::ILLEGAL_ARGUMENT},
        {SERVER_FILE_NOT_FOUND, StatusCode::FILE_NOT_FOUND},
        {SERVER_NOT_IMPLEMENT, StatusCode::UNEXPECTED_ERROR},
        {SERVER_CANNOT_CREATE_FOLDER, StatusCode::CANNOT_CREATE_FOLDER},
        {SERVER_CANNOT_CREATE_FILE, StatusCode::CANNOT_CREATE_FILE},
        {SERVER_CANNOT_DELETE_FOLDER, StatusCode::CANNOT_DELETE_FOLDER},
        {SERVER_CANNOT_DELETE_FILE, StatusCode::CANNOT_DELETE_FILE},
        {SERVER_TABLE_NOT_EXIST, StatusCode::TABLE_NOT_EXISTS},
        {SERVER_INVALID_TABLE_NAME, StatusCode::ILLEGAL_TABLE_NAME},
        {SERVER_INVALID_TABLE_DIMENSION, StatusCode::ILLEGAL_DIMENSION},
        {SERVER_INVALID_TIME_RANGE, StatusCode::ILLEGAL_RANGE},
        {SERVER_INVALID_VECTOR_DIMENSION, StatusCode::ILLEGAL_DIMENSION},

        {SERVER_INVALID_INDEX_TYPE, StatusCode::ILLEGAL_INDEX_TYPE},
        {SERVER_INVALID_ROWRECORD, StatusCode::ILLEGAL_ROWRECORD},
        {SERVER_INVALID_ROWRECORD_ARRAY, StatusCode::ILLEGAL_ROWRECORD},
        {SERVER_INVALID_TOPK, StatusCode::ILLEGAL_TOPK},
        {SERVER_INVALID_NPROBE, StatusCode::ILLEGAL_ARGUMENT},
        {SERVER_INVALID_INDEX_NLIST, StatusCode::ILLEGAL_NLIST},
        {SERVER_INVALID_INDEX_METRIC_TYPE, StatusCode::ILLEGAL_METRIC_TYPE},
        {SERVER_INVALID_INDEX_FILE_SIZE, StatusCode::ILLEGAL_ARGUMENT},
        {SERVER_ILLEGAL_VECTOR_ID, StatusCode::ILLEGAL_VECTOR_ID},
        {SERVER_ILLEGAL_SEARCH_RESULT, StatusCode::ILLEGAL_SEARCH_RESULT},
        {SERVER_CACHE_FULL, StatusCode::CACHE_FAILED},
        {SERVER_BUILD_INDEX_ERROR, StatusCode::BUILD_INDEX_ERROR},
        {SERVER_OUT_OF_MEMORY, StatusCode::OUT_OF_MEMORY},

        {DB_NOT_FOUND, StatusCode::TABLE_NOT_EXISTS},
        {DB_META_TRANSACTION_FAILED, StatusCode::META_FAILED},
    };

    if (code_map.find(code) != code_map.end()) {
        return code_map.at(code);
    } else {
        return StatusCode::UNEXPECTED_ERROR;
    }
}

///////////////////////// WebRequestHandler methods ///////////////////////////////////////

Status
WebRequestHandler::getTaleInfo(const std::shared_ptr<Context>& context,
                               const std::string& table_name,
                               std::map<std::string, std::string>& table_info) {
    TableSchema schema;
    auto status = request_handler_.DescribeTable(context_ptr_, table_name, schema);
    if (!status.ok()) {
        return status;
    }

    int64_t count;
    status = request_handler_.CountTable(context_ptr_, table_name, count);
    if (!status.ok()) {
        return status;
    }

    IndexParam index_param;
    status = request_handler_.DescribeIndex(context_ptr_, table_name, index_param);
    if (!status.ok()) {
        return status;
    }

    table_info["table_name"] = schema.table_name_;
    table_info["dimension"] = std::to_string(schema.dimension_);
    table_info["metric_type"] = MetricMap.at(engine::MetricType(schema.metric_type_));
    table_info["index_file_size"] = schema.index_file_size_;

    table_info["index"] = IndexMap.at(engine::EngineType(index_param.index_type_));
    table_info["nlist"] = std::to_string(index_param.nlist_);

    table_info["count"] = std::to_string(count);
}

/////////////////////////////////////////// Router methods ////////////////////////////////////////////
StatusDto::ObjectWrapper
WebRequestHandler::GetDevices(DevicesDto::ObjectWrapper& devices_dto) {
    auto lamd = [](uint64_t x) -> uint64_t {
        return x / 1024 / 1024 / 1024;
    };
    auto system_info = SystemInfo::GetInstance();

    devices_dto->cpu = devices_dto->cpu->createShared();
    devices_dto->cpu->memory = lamd(system_info.GetPhysicalMemory());

    devices_dto->gpus = devices_dto->gpus->createShared();
    size_t count = system_info.num_device();
    std::vector<uint64_t> device_mems = system_info.GPUMemoryTotal();

    if (count != device_mems.size()) {
        ASSIGN_RETURN_STATUS_DTO(Status(UNEXPECTED_ERROR, "Can't obtain GPU info"));
    }

    for (size_t i = 0; i < count; i++) {
        auto device_dto = DeviceInfoDto::createShared();
        device_dto->memory = lamd(device_mems.at(i));
        devices_dto->gpus->put("GPU" + OString(std::to_string(i).c_str()), device_dto);
    }

    ASSIGN_RETURN_STATUS_DTO(Status::OK());
}

StatusDto::ObjectWrapper
WebRequestHandler::GetAdvancedConfig(AdvancedConfigDto::ObjectWrapper& advanced_config) {
    Config& config = Config::GetInstance();

//    advanced_config->cpu_cache_capacity =
    int64_t value;
    auto status = config.GetCacheConfigCpuCacheCapacity(value);
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status);
    }
    advanced_config->cpu_cache_capacity = value;

    bool ok;
    status = config.GetCacheConfigCacheInsertData(ok);
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status)
    }
    advanced_config->cache_insert_data = ok;

    status = config.GetEngineConfigUseBlasThreshold(value);
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status)
    }
    advanced_config->use_blas_threshold = value;

    status = config.GetEngineConfigGpuSearchThreshold(value);
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status)
    }
    advanced_config->gpu_search_threshold = value;

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::SetAdvancedConfig(const AdvancedConfigDto::ObjectWrapper& advanced_config) {
    Config& config = Config::GetInstance();

    auto
        status = config.SetCacheConfigCpuCacheCapacity(std::to_string(advanced_config->cpu_cache_capacity->getValue()));
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status)
    }

    status = config.SetCacheConfigCacheInsertData(std::to_string(advanced_config->cache_insert_data->getValue()));
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status)
    }

    status = config.SetEngineConfigUseBlasThreshold(std::to_string(advanced_config->use_blas_threshold->getValue()));
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status)
    }

    status =
        config.SetEngineConfigGpuSearchThreshold(std::to_string(advanced_config->gpu_search_threshold->getValue()));
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status)
    }

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::GetGpuConfig(GPUConfigDto::ObjectWrapper& gpu_config_dto) {
    Config& config = Config::GetInstance();

    bool enable;
    auto status = config.GetGpuResourceConfigEnable(enable);
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status);
    }

    gpu_config_dto->enable = enable;

    if (!enable) {
        ASSIGN_RETURN_STATUS_DTO(Status::OK());
    }

    int64_t capacity;
    status = config.GetGpuResourceConfigCacheCapacity(capacity);
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status);
    }
    gpu_config_dto->cache_capacity = capacity;

    std::vector<int64_t> values;
    status = config.GetGpuResourceConfigSearchResources(values);
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status);
    }

    gpu_config_dto->search_resources = gpu_config_dto->search_resources->createShared();
    for (auto& device_id : values) {
        gpu_config_dto->search_resources->pushBack("GPU" + OString(std::to_string(device_id).c_str()));
    }

    values.clear();
    status = config.GetGpuResourceConfigBuildIndexResources(values);
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status);
    }

    gpu_config_dto->build_index_resources = gpu_config_dto->build_index_resources->createShared();
    for (auto& device_id : values) {
        gpu_config_dto->build_index_resources->pushBack("GPU" + OString(std::to_string(device_id).c_str()));
    }

    ASSIGN_RETURN_STATUS_DTO(Status::OK());
}

StatusDto::ObjectWrapper
WebRequestHandler::SetGpuConfig(const GPUConfigDto::ObjectWrapper& gpu_config_dto) {
    auto str_lower_lamda = [](const std::string& str){
        std::string str_out;
        for (auto & item : str) {
            str_out += tolower(item);
        }
        return str_out;
    };

    Config& config = Config::GetInstance();

    auto status = config.SetGpuResourceConfigEnable(std::to_string(gpu_config_dto->enable->getValue()));
    if (!status.ok() || !gpu_config_dto->enable->getValue()) {
        ASSIGN_RETURN_STATUS_DTO(status);
    }

    status = config.SetGpuResourceConfigCacheCapacity(std::to_string(gpu_config_dto->cache_capacity->getValue()));
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status);
    }

    std::vector<std::string> search_resources;
    gpu_config_dto->search_resources->forEach([&search_resources, str_lower_lamda](const OString& res) {
        search_resources.emplace_back(str_lower_lamda(res->std_str()));
    });

    std::string search_resources_value;
    for (auto& res : search_resources) {
        search_resources_value += res + ",";
    }
    auto len = search_resources_value.size();
    search_resources_value.erase(len -1);
    status = config.SetGpuResourceConfigSearchResources(search_resources_value);
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status);
    }

    std::vector<std::string> build_resources;
    gpu_config_dto->build_index_resources->forEach([&build_resources, str_lower_lamda](const OString& res) {
        build_resources.emplace_back(str_lower_lamda(res->std_str()));
    });

    std::string build_resources_value;
    for (auto& res : build_resources) {
        build_resources_value += res + ",";
    }
    len = build_resources_value.size();
    build_resources_value.erase(len -1);
    status = config.SetGpuResourceConfigBuildIndexResources(build_resources_value);
    if (!status.ok()) {
        ASSIGN_RETURN_STATUS_DTO(status);
    }

    ASSIGN_RETURN_STATUS_DTO(Status::OK());
}

StatusDto::ObjectWrapper
WebRequestHandler::CreateTable(const TableRequestDto::ObjectWrapper& table_schema) {
    auto status =
        request_handler_.CreateTable(context_ptr_, table_schema->table_name->std_str(), table_schema->dimension,
                                     table_schema->index_file_size, table_schema->metric_type);

    ASSIGN_RETURN_STATUS_DTO(status)
}

/**
 * fields:
 *  - ALL: request all fields
 *  - COUNT: request number of vectors
 *  - *,*,*...:
 */
StatusDto::ObjectWrapper
WebRequestHandler::GetTable(const OString& table_name,
                            const OQueryParams& query_params,
                            TableFieldsDto::ObjectWrapper& fields_dto) {
    Status status = Status::OK();

    // TODO: query string field `fields` npt used here

    std::map<std::string, std::string> table_info;
    getTaleInfo(context_ptr_, table_name->std_str(), table_info);

    fields_dto->table_name = table_info["table_name"].c_str();
    fields_dto->dimension = atol(table_info["dimension"].c_str());
    fields_dto->index = table_info["index"].c_str();
    fields_dto->nlist = atol(table_info["nlist"].c_str());
    fields_dto->metric_type = table_info["metric_type"].c_str();
    fields_dto->num = atol(table_info["count"].c_str());
//    if (query_params.getSize() == 0) {
//        TableSchema schema;
//        status = request_handler_.DescribeTable(context_ptr_, table_name->std_str(), schema);
//        if (status.ok()) {
//            fields_dto->schema->put("table_name", schema.table_name_.c_str());
//            fields_dto->schema->put("dimension", OString(std::to_string(schema.dimension_).c_str()));
//            fields_dto->schema->put("index_file_size", OString(std::to_string(schema.index_file_size_).c_str()));
//            fields_dto->schema->put("metric_type", OString(std::to_string(schema.metric_type_).c_str()));
//        }
//    } else {
//        for (auto& param : query_params.getAll()) {
//            std::string key = param.first.std_str();
//            std::string value = param.second.std_str();
//
//            if ("fields" == key) {
//                if ("num" == value) {
//                    int64_t count;
//                    status = request_handler_.CountTable(context_ptr_, table_name->std_str(), count);
//                    if (status.ok()) {
//                        fields_dto->schema->put("num", OString(std::to_string(count).c_str()));
//                    }
//                }
//            }
//        }
//    }

    ASSIGN_RETURN_STATUS_DTO(status);
}

StatusDto::ObjectWrapper
WebRequestHandler::ShowTables(const OInt64& offset, const OInt64& page_size,
                              TableListFieldsDto::ObjectWrapper& response_dto) {
    std::vector<std::string> tables;
    Status status = Status::OK();

    response_dto->tables = response_dto->tables->createShared();

    if (offset < 0 || page_size < 0) {
        status = Status(SERVER_UNEXPECTED_ERROR, "Query param 'offset' or 'page_size' should bigger than 0");
    } else {
        status = request_handler_.ShowTables(context_ptr_, tables);
        if (status.ok() && offset < tables.size()) {
            int64_t size = (page_size->getValue() + offset->getValue() > tables.size()) ? tables.size() - offset
                                                                                        : page_size->getValue();
            for (int64_t i = offset->getValue(); i < size + offset->getValue(); i++) {
                std::map<std::string, std::string> table_info;

                status = getTaleInfo(context_ptr_, tables.at(i), table_info);
                if (!status.ok()) {
                    break;
                }

                auto table_fields_dto = TableFieldsDto::createShared();
                table_fields_dto->table_name = table_info["table_name"].c_str();
                table_fields_dto->dimension = atol(table_info["dimension"].c_str());
                table_fields_dto->index_file_size = atol(table_info["index_file_size"].c_str());
                table_fields_dto->index = table_info["index"].c_str();
                table_fields_dto->nlist = atol(table_info["nlist"].c_str());
                table_fields_dto->metric_type = table_info["metric_type"].c_str();
                table_fields_dto->num = atol(table_info["count"].c_str());

                response_dto->tables->pushBack(table_fields_dto);
            }

            response_dto->count = tables.size();
        }
    }

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::DropTable(const OString& table_name) {
    auto status = request_handler_.DropTable(context_ptr_, table_name->std_str());

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::CreateIndex(const OString& table_name, const IndexRequestDto::ObjectWrapper& index_param) {
    std::string index_type = index_param->index_type->std_str();
    auto status = request_handler_.CreateIndex(context_ptr_, table_name->std_str(),
                                               static_cast<int64_t>(IndexNameMap.at(index_type)),
                                               index_param->nlist->getValue());

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::GetIndex(const OString& table_name, IndexDto::ObjectWrapper& index_dto) {
    IndexParam param;
    auto status = request_handler_.DescribeIndex(context_ptr_, table_name->std_str(), param);

    if (status.ok()) {
        index_dto->index_type = param.index_type_;
        index_dto->nlist = param.nlist_;
    }

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::DropIndex(const OString& table_name) {
    auto status = request_handler_.DropIndex(context_ptr_, table_name->std_str());

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::CreatePartition(const OString& table_name, const PartitionRequestDto::ObjectWrapper& param) {
    auto status = request_handler_.CreatePartition(context_ptr_, table_name->std_str(),
                                                   param->partition_name->std_str(), param->tag->std_str());

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::ShowPartitions(const OInt64& offset, const OInt64& page_size, const OString& table_name,
                                  PartitionListDto::ObjectWrapper& partition_list_dto) {
    std::vector<PartitionParam> partitions;
    auto status = request_handler_.ShowPartitions(context_ptr_, table_name->std_str(), partitions);

    if (status.ok()) {
        partition_list_dto->partitions = partition_list_dto->partitions->createShared();

        if (offset->getValue() + 1 < partitions.size()) {
            int64_t size = (offset->getValue() + page_size->getValue() > partitions.size()) ? partitions.size()
                                                                                            : page_size->getValue();

            for (int64_t i = 0; i < size; i++) {
                auto partition_dto = PartitionFieldsDto::createShared();
                partition_dto->schema = partition_dto->schema->createShared();
                partition_dto->schema->put("partitin_name", partitions.at(i + offset).partition_name_.c_str());
                partition_dto->schema->put("tag", partitions.at(i + offset).tag_.c_str());

                partition_list_dto->partitions->pushBack(partition_dto);
            }
        }
    }

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::DropPartition(const OString& table_name, const OString& tag) {
    auto status = request_handler_.DropPartition(context_ptr_, table_name->std_str(), "", tag->std_str());

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::Insert(const InsertRequestDto::ObjectWrapper& param, VectorIdsDto::ObjectWrapper& ids_dto) {
    std::vector<int64_t> ids;
    if (param->ids->count() > 0) {
        for (int64_t i = 0; i < param->ids->count(); i++) {
            ids.emplace_back(param->ids->get(i)->getValue());
        }
    }

    std::vector<float> datas;
    for (int64_t j = 0; j < param->records->count(); j++) {
        for (int64_t k = 0; k < param->records->get(j)->record->count(); k++) {
            datas.emplace_back(param->records->get(j)->record->get(k)->getValue());
        }
    }

    auto status = request_handler_.Insert(context_ptr_, param->table_name->std_str(), param->records->count(), datas,
                                          param->tag->std_str(), ids);

    if (status.ok()) {
        ids_dto->ids = ids_dto->ids->createShared();
        for (auto& id : ids) {
            ids_dto->ids->pushBack(id);
        }
    }

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::Search(const OString& table_name, const OInt64& topk, const OInt64& nprobe,
                          const OQueryParams& query_params, const RecordsDto::ObjectWrapper& records,
                          ResultDto::ObjectWrapper& results_dto) {
    int64_t topk_t = topk->getValue();
    int64_t nprobe_t = nprobe->getValue();

    std::vector<std::string> tag_list;
    std::vector<std::string> file_id_list;

    for (auto& param : query_params.getAll()) {
        std::string key = param.first.std_str();
        std::string value = param.second.std_str();

        if ("tags" == key) {
            boost::split(tag_list, value, boost::is_any_of(","), boost::token_compress_on);
        } else if ("file_ids" == key) {
            boost::split(file_id_list, value, boost::is_any_of(","), boost::token_compress_on);
        }
    }

    std::vector<float> datas;
    for (int64_t j = 0; j < records->records->count(); j++) {
        for (int64_t k = 0; k < records->records->get(j)->record->count(); k++) {
            datas.emplace_back(records->records->get(j)->record->get(k)->getValue());
        }
    }

    std::vector<Range> range_list;

    TopKQueryResult result;
    auto context_ptr = GenContextPtr("Web Handler");
    auto status = request_handler_.Search(context_ptr, table_name->std_str(), records->records->count(), datas,
                                          range_list, topk_t, nprobe_t, tag_list, file_id_list, result);

    if (status.ok()) {
        results_dto->ids = results_dto->ids->createShared();
        for (auto& id : result.id_list_) {
            results_dto->ids->pushBack(id);
        }
        results_dto->dits = results_dto->dits->createShared();
        for (auto& dit : result.distance_list_) {
            results_dto->dits->pushBack(dit);
        }
        results_dto->num = result.row_num_;
    }

    ASSIGN_RETURN_STATUS_DTO(status)
}

StatusDto::ObjectWrapper
WebRequestHandler::Cmd(const OString& cmd, CommandDto::ObjectWrapper& cmd_dto) {
    std::string reply_str;
    auto status = request_handler_.Cmd(context_ptr_, cmd->std_str(), reply_str);

    if (status.ok()) {
        cmd_dto->reply = reply_str.c_str();
    }

    ASSIGN_RETURN_STATUS_DTO(status);
}

}  // namespace web
}  // namespace server
}  // namespace milvus
