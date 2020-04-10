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

#include <fiu-control.h>
#include <fiu-local.h>
#include <gtest/gtest.h>

#include <boost/filesystem.hpp>
#include <random>
#include <thread>

#include "cache/CpuCacheMgr.h"
#include "config/Config.h"
#include "db/Constants.h"
#include "db/DB.h"
#include "db/DBFactory.h"
#include "db/DBImpl.h"
#include "db/IDGenerator.h"
#include "db/meta/MetaConsts.h"
#include "db/utils.h"
#include "utils/CommonUtil.h"

namespace {
static const char* TABLE_NAME = "test_hybrid";
static constexpr int64_t TABLE_DIM = 128;
static constexpr int64_t SECONDS_EACH_HOUR = 3600;
static constexpr int64_t FIELD_NUM = 4;
static constexpr int64_t NQ = 10;
static constexpr int64_t TOPK = 100;

void
BuildTableSchema(milvus::engine::meta::TableSchema& table_schema,
                 milvus::engine::meta::hybrid::FieldsSchema& fields_schema,
                 std::unordered_map<std::string, milvus::engine::meta::hybrid::DataType>& attr_type) {

    table_schema.dimension_ = TABLE_DIM;
    table_schema.table_id_ = TABLE_NAME;

    std::vector<milvus::engine::meta::hybrid::FieldSchema> fields;
    fields.resize(FIELD_NUM);
    for (uint64_t i = 0; i < FIELD_NUM; ++i) {
        fields[i].collection_id_ = TABLE_NAME;
        fields[i].field_name_ = "field_" + std::to_string(i + 1);
    }
    fields[0].field_type_ = (int)milvus::engine::meta::hybrid::DataType::INT32;
    fields[1].field_type_ = (int)milvus::engine::meta::hybrid::DataType::INT64;
    fields[2].field_type_ = (int)milvus::engine::meta::hybrid::DataType::FLOAT;
    fields[3].field_type_ = (int)milvus::engine::meta::hybrid::DataType::VECTOR;
    fields_schema.fields_schema_ = fields;

    attr_type.insert(std::make_pair("field_0", milvus::engine::meta::hybrid::DataType::INT32));
    attr_type.insert(std::make_pair("field_1", milvus::engine::meta::hybrid::DataType::INT64));
    attr_type.insert(std::make_pair("field_2", milvus::engine::meta::hybrid::DataType::FLOAT));
}

void
BuildVectors(uint64_t n, uint64_t batch_index, milvus::engine::VectorsData& vectors) {
    vectors.vector_count_ = n;
    vectors.float_data_.clear();
    vectors.float_data_.resize(n * TABLE_DIM);
    float* data = vectors.float_data_.data();
    for (uint64_t i = 0; i < n; i++) {
        for (int64_t j = 0; j < TABLE_DIM; j++) data[TABLE_DIM * i + j] = drand48();
        data[TABLE_DIM * i] += i / 2000.;

        vectors.id_array_.push_back(n * batch_index + i);
    }
}

void
BuildEntity(uint64_t n, uint64_t batch_index, milvus::engine::Entity& entity) {
    milvus::engine::VectorsData vectors;
    vectors.vector_count_ = n;
    vectors.float_data_.clear();
    vectors.float_data_.resize(n * TABLE_DIM);
    float* data = vectors.float_data_.data();
    for (uint64_t i = 0; i < n; i++) {
        for (int64_t j = 0; j < TABLE_DIM; j++) data[TABLE_DIM * i + j] = drand48();
        data[TABLE_DIM * i] += i / 2000.;

        vectors.id_array_.push_back(n * batch_index + i);
    }
    entity.vector_data_.insert(std::make_pair("field_3", vectors));
    std::vector<std::string> value_0, value_1, value_2;
    value_0.resize(n);
    value_1.resize(n);
    value_2.resize(n);
    for (uint64_t i = 0; i < n; ++i) {
        value_0[i] = std::to_string(i);
        value_1[i] = std::to_string(i + n);
        value_2[i] = std::to_string((i + 100) / (n + 1));
    }
    entity.entity_count_ = n;
    entity.attr_data_.insert(std::make_pair("field_0", value_0));
    entity.attr_data_.insert(std::make_pair("field_1", value_1));
    entity.attr_data_.insert(std::make_pair("field_2", value_2));
}

void
ConstructGeneralQuery(milvus::query::GeneralQueryPtr& general_query) {
    general_query->bin->relation = milvus::query::QueryRelation::AND;
    general_query->bin->left_query = std::make_shared<milvus::query::GeneralQuery>();
    general_query->bin->right_query = std::make_shared<milvus::query::GeneralQuery>();
    auto left = general_query->bin->left_query;
    auto right = general_query->bin->right_query;
    left->bin->relation = milvus::query::QueryRelation::AND;


    auto term_query = std::make_shared<milvus::query::TermQuery>();
    term_query->field_name = "field_0";
    term_query->field_value = {"10", "20", "30", "40", "50"};
    term_query->boost = 1;

    auto range_query = std::make_shared<milvus::query::RangeQuery>();
    range_query->field_name = "field_1";
    std::vector<milvus::query::CompareExpr> compare_expr;
    compare_expr.resize(2);
    compare_expr[0].compare_operator = milvus::query::CompareOperator::GTE;
    compare_expr[0].operand = "1000";
    compare_expr[1].compare_operator = milvus::query::CompareOperator::LTE;
    compare_expr[1].operand = "5000";
    range_query->compare_expr = compare_expr;
    range_query->boost = 2;

    auto vector_query = std::make_shared<milvus::query::VectorQuery>();
    vector_query->field_name = "field_3";
    vector_query->topk = 100;
    vector_query->boost = 3;
    milvus::query::VectorRecord record;
    record.float_data.resize(NQ * TABLE_DIM);
    float* data = record.float_data.data();
    for (uint64_t i = 0; i < NQ; i++) {
        for (int64_t j = 0; j < TABLE_DIM; j++) data[TABLE_DIM * i + j] = drand48();
        data[TABLE_DIM * i] += i / 2000.;
    }
    vector_query->query_vector = record;


    left->bin->left_query = std::make_shared<milvus::query::GeneralQuery>();
    left->bin->right_query = std::make_shared<milvus::query::GeneralQuery>();
    left->bin->left_query->leaf = std::make_shared<milvus::query::LeafQuery>();
    left->bin->right_query->leaf = std::make_shared<milvus::query::LeafQuery>();
    left->bin->left_query->leaf->term_query = term_query;
    left->bin->right_query->leaf->range_query = range_query;

    right->leaf = std::make_shared<milvus::query::LeafQuery>();
    right->leaf->vector_query = vector_query;
}

TEST_F(DBTest, HYBRID_DB_TEST) {
    milvus::engine::meta::TableSchema table_info;
    milvus::engine::meta::hybrid::FieldsSchema fields_info;
    std::unordered_map<std::string, milvus::engine::meta::hybrid::DataType> attr_type;
    BuildTableSchema(table_info, fields_info, attr_type);

    auto stat = db_->CreateHybridCollection(table_info, fields_info);
    ASSERT_TRUE(stat.ok());
    milvus::engine::meta::TableSchema table_info_get;
    milvus::engine::meta::hybrid::FieldsSchema fields_info_get;
    table_info_get.table_id_ = TABLE_NAME;
    stat = db_->DescribeHybridCollection(table_info_get, fields_info_get);
    ASSERT_TRUE(stat.ok());
    ASSERT_EQ(table_info_get.dimension_, TABLE_DIM);

    uint64_t qb = 1000;
    milvus::engine::Entity entity;
    BuildEntity(qb, 0, entity);

    stat = db_->InsertEntities(TABLE_NAME, "", entity, attr_type);
    ASSERT_TRUE(stat.ok());

    stat = db_->Flush();
    ASSERT_TRUE(stat.ok());

    milvus::json json_params = {{"nprobe", 10}};
    milvus::engine::TableIndex index;
    index.engine_type_ = (int)milvus::engine::EngineType::FAISS_IDMAP;
    index.extra_params_ = {{"nlist", 16384}};

    stat = db_->CreateIndex(TABLE_NAME, index);
    ASSERT_TRUE(stat.ok());
}

TEST_F(DBTest, HYBRID_SEARCH_TEST) {
    milvus::engine::meta::TableSchema table_info;
    milvus::engine::meta::hybrid::FieldsSchema fields_info;
    std::unordered_map<std::string, milvus::engine::meta::hybrid::DataType> attr_type;
    BuildTableSchema(table_info, fields_info, attr_type);

    auto stat = db_->CreateHybridCollection(table_info, fields_info);
    ASSERT_TRUE(stat.ok());
    milvus::engine::meta::TableSchema table_info_get;
    milvus::engine::meta::hybrid::FieldsSchema fields_info_get;
    table_info_get.table_id_ = TABLE_NAME;
    stat = db_->DescribeHybridCollection(table_info_get, fields_info_get);
    ASSERT_TRUE(stat.ok());
    ASSERT_EQ(table_info_get.dimension_, TABLE_DIM);

    uint64_t qb = 1000;
    milvus::engine::Entity entity;
    BuildEntity(qb, 0, entity);

    stat = db_->InsertEntities(TABLE_NAME, "", entity, attr_type);
    ASSERT_TRUE(stat.ok());

    stat = db_->Flush();
    ASSERT_TRUE(stat.ok());

    // Construct general query
    milvus::query::GeneralQueryPtr general_query = std::make_shared<milvus::query::GeneralQuery>();
    ConstructGeneralQuery(general_query);

    std::vector<std::string> tags;
    milvus::context::HybridSearchContextPtr hybrid_context = std::make_shared<milvus::context::HybridSearchContext>();
    milvus::engine::ResultIds result_ids;
    milvus::engine::ResultDistances result_distances;
    stat = db_->HybridQuery(dummy_context_,
                            TABLE_NAME,
                            tags,
                            hybrid_context,
                            general_query,
                            attr_type,
                            result_ids,
                            result_distances);

}

}