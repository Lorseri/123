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
#include <condition_variable>
#include <list>
#include <map>
#include <memory>
#include <mutex>
#include <set>
#include <string>
#include <thread>
#include <vector>

#include "DB.h"
#include "db/IndexFailedChecker.h"
#include "db/OngoingFileChecker.h"
#include "db/Types.h"
#include "db/insert/MemManager.h"
#include "utils/ThreadPool.h"
#include "wal/WalManager.h"

namespace milvus {
namespace engine {

namespace meta {
class Meta;
}

class DBImpl : public DB {
 public:
    explicit DBImpl(const DBOptions& options);
    ~DBImpl();

    Status
    Start() override;
    Status
    Stop() override;
    Status
    DropAll() override;

    Status
    CreateTable(meta::TableSchema& table_schema) override;

    Status
    DropTable(const std::string& table_id, const meta::DatesT& dates) override;

    Status
    DescribeTable(meta::TableSchema& table_schema) override;

    Status
    HasTable(const std::string& table_id, bool& has_or_not) override;

    Status
    AllTables(std::vector<meta::TableSchema>& table_schema_array) override;

    Status
    GetTableInfo(const std::string& table_id, TableInfo& table_info) override;

    Status
    PreloadTable(const std::string& table_id) override;

    Status
    UpdateTableFlag(const std::string& table_id, int64_t flag);

    Status
    GetTableRowCount(const std::string& table_id, uint64_t& row_count) override;

    Status
    CreatePartition(const std::string& table_id, const std::string& partition_name,
                    const std::string& partition_tag) override;

    Status
    DropPartition(const std::string& partition_name) override;

    Status
    DropPartitionByTag(const std::string& table_id, const std::string& partition_tag) override;

    Status
    ShowPartitions(const std::string& table_id, std::vector<meta::TableSchema>& partition_schema_array) override;

    Status
    InsertVectors(const std::string& table_id, const std::string& partition_tag, VectorsData& vectors) override;

    Status
    DeleteVector(const std::string& table_id, IDNumber vector_id) override;

    Status
    DeleteVectors(const std::string& table_id, IDNumbers vector_ids) override;

    Status
    Flush(const std::string& table_id) override;

    Status
    Flush() override;

    Status
    Compact(const std::string& table_id) override;

    //    Status
    //    Merge(const std::set<std::string>& table_ids) override;

    Status
    CreateIndex(const std::string& table_id, const TableIndex& index) override;

    Status
    DescribeIndex(const std::string& table_id, TableIndex& index) override;

    Status
    DropIndex(const std::string& table_id) override;

    Status
    QueryByID(const std::shared_ptr<server::Context>& context, const std::string& table_id,
              const std::vector<std::string>& partition_tags, uint64_t k, uint64_t nprobe, IDNumber vector_id,
              ResultIds& result_ids, ResultDistances& result_distances) override;

    Status
    Query(const std::shared_ptr<server::Context>& context, const std::string& table_id,
          const std::vector<std::string>& partition_tags, uint64_t k, uint64_t nprobe, const VectorsData& vectors,
          ResultIds& result_ids, ResultDistances& result_distances) override;

    Status
    Query(const std::shared_ptr<server::Context>& context, const std::string& table_id,
          const std::vector<std::string>& partition_tags, uint64_t k, uint64_t nprobe, const VectorsData& vectors,
          const meta::DatesT& dates, ResultIds& result_ids, ResultDistances& result_distances) override;

    Status
    QueryByFileID(const std::shared_ptr<server::Context>& context, const std::string& table_id,
                  const std::vector<std::string>& file_ids, uint64_t k, uint64_t nprobe, const VectorsData& vectors,
                  const meta::DatesT& dates, ResultIds& result_ids, ResultDistances& result_distances) override;

    Status
    Size(uint64_t& result) override;

 private:
    Status
    QueryAsync(const std::shared_ptr<server::Context>& context, const std::string& table_id,
               const meta::TableFilesSchema& files, uint64_t k, uint64_t nprobe, const VectorsData& vectors,
               ResultIds& result_ids, ResultDistances& result_distances);

    void
    BackgroundTimerTask();
    void
    WaitMergeFileFinish();
    void
    WaitBuildIndexFinish();

    void
    StartMetricTask();

    void
    StartCompactionTask();

    Status
    MergeFiles(const std::string& table_id, const meta::DateT& date, const meta::TableFilesSchema& files);
    Status
    BackgroundMergeFiles(const std::string& table_id);
    void
    BackgroundCompaction(std::set<std::string> table_ids);

    void
    StartBuildIndexTask(bool force = false);
    void
    BackgroundBuildIndex();

    Status
    CompactFile(const std::string& table_id, const milvus::engine::meta::TableFileSchema& file);

    /*
    Status
    SyncMemData(std::set<std::string>& sync_table_ids);
    */

    Status
    GetFilesToBuildIndex(const std::string& table_id, const std::vector<int>& file_types,
                         meta::TableFilesSchema& files);

    Status
    GetFilesToSearch(const std::string& table_id, const std::vector<size_t>& file_ids, const meta::DatesT& dates,
                     meta::TableFilesSchema& files);

    Status
    GetPartitionByTag(const std::string& table_id, const std::string& partition_tags,
                      std::string& partition_name_array);

    Status
    GetPartitionsByTags(const std::string& table_id, const std::vector<std::string>& partition_tags,
                        std::set<std::string>& partition_name_array);

    Status
    DropTableRecursively(const std::string& table_id, const meta::DatesT& dates);

    Status
    UpdateTableIndexRecursively(const std::string& table_id, const TableIndex& index);

    Status
    BuildTableIndexRecursively(const std::string& table_id, const TableIndex& index);

    Status
    DropTableIndexRecursively(const std::string& table_id);

    Status
    GetTableRowCountRecursively(const std::string& table_id, uint64_t& row_count);

    Status
    ExecWalRecord(const wal::MXLogRecord& record);

    void
    BackgroundWalTask();

 private:
    const DBOptions options_;

    std::atomic<bool> initialized_;

    std::thread bg_timer_thread_;

    meta::MetaPtr meta_ptr_;
    MemManagerPtr mem_mgr_;
    std::mutex mem_serialize_mutex_;

    bool wal_enable_ = false;
    std::shared_ptr<wal::WalManager> wal_mgr_;
    std::thread bg_wal_thread_;

    struct SimpleWaitNotify {
        bool notified_ = false;
        std::mutex mutex_;
        std::condition_variable cv_;

        void
        Wait() {
            std::unique_lock<std::mutex> lck(mutex_);
            if (!notified_) {
                cv_.wait(lck);
            }
            notified_ = false;
        }

        void
        Wait_Until(const std::chrono::system_clock::time_point& tm_pint) {
            std::unique_lock<std::mutex> lck(mutex_);
            if (!notified_) {
                cv_.wait_until(lck, tm_pint);
            }
            notified_ = false;
        }

        void
        Notify() {
            std::unique_lock<std::mutex> lck(mutex_);
            notified_ = true;
            lck.unlock();
            cv_.notify_one();
        }
    };

    SimpleWaitNotify wal_task_swn_;
    SimpleWaitNotify flush_task_swn_;

    ThreadPool compact_thread_pool_;
    std::mutex compact_result_mutex_;
    std::list<std::future<void>> compact_thread_results_;
    std::set<std::string> compact_table_ids_;
    //    std::mutex compact_table_ids_mutex_;

    ThreadPool index_thread_pool_;
    std::mutex index_result_mutex_;
    std::list<std::future<void>> index_thread_results_;

    std::mutex build_index_mutex_;

    IndexFailedChecker index_failed_checker_;
    OngoingFileChecker ongoing_files_checker_;
};  // DBImpl

}  // namespace engine
}  // namespace milvus
