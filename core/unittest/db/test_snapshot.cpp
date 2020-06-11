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

#include <random>
#include <string>
#include <set>

#include "db/utils.h"
#include "db/snapshot/ReferenceProxy.h"
#include "db/snapshot/ScopedResource.h"
#include "db/snapshot/WrappedTypes.h"
#include "db/snapshot/ResourceHolders.h"
#include "db/snapshot/OperationExecutor.h"
#include "db/snapshot/Store.h"
#include "db/snapshot/Context.h"
#include "db/snapshot/CompoundOperations.h"
#include "db/snapshot/Snapshots.h"

int RandomInt(int start, int end) {
    std::random_device dev;
    std::mt19937 rng(dev());
    std::uniform_int_distribution<std::mt19937::result_type> dist(start, end);
    return dist(rng);
}

TEST_F(SnapshotTest, ReferenceProxyTest) {
    std::string status("raw");
    const std::string CALLED = "CALLED";
    auto callback = [&]() {
        status = CALLED;
    };

    auto proxy = milvus::engine::snapshot::ReferenceProxy();
    ASSERT_EQ(proxy.RefCnt(), 0);

    int refcnt = 3;
    for (auto i = 0; i < refcnt; ++i) {
        proxy.Ref();
    }
    ASSERT_EQ(proxy.RefCnt(), refcnt);

    proxy.RegisterOnNoRefCB(callback);

    for (auto i = 0; i < refcnt; ++i) {
        proxy.UnRef();
    }
    ASSERT_EQ(proxy.RefCnt(), 0);
    ASSERT_EQ(status, CALLED);
}

TEST_F(SnapshotTest, ScopedResourceTest) {
    auto inner = std::make_shared<milvus::engine::snapshot::Collection>("c1");
    ASSERT_EQ(inner->RefCnt(), 0);

    {
        auto not_scoped = milvus::engine::snapshot::CollectionScopedT(inner, false);
        ASSERT_EQ(not_scoped->RefCnt(), 0);
        not_scoped->Ref();
        ASSERT_EQ(not_scoped->RefCnt(), 1);
        ASSERT_EQ(inner->RefCnt(), 1);

        auto not_scoped_2 = not_scoped;
        ASSERT_EQ(not_scoped_2->RefCnt(), 1);
        ASSERT_EQ(not_scoped->RefCnt(), 1);
        ASSERT_EQ(inner->RefCnt(), 1);
    }
    ASSERT_EQ(inner->RefCnt(), 1);

    inner->UnRef();
    ASSERT_EQ(inner->RefCnt(), 0);

    {
        // Test scoped construct
        auto scoped = milvus::engine::snapshot::CollectionScopedT(inner);
        ASSERT_EQ(scoped->RefCnt(), 1);
        ASSERT_EQ(inner->RefCnt(), 1);

        {
            // Test bool operator
            decltype(scoped) other_scoped;
            ASSERT_EQ(other_scoped, false);
            // Test operator=
            other_scoped = scoped;
            ASSERT_EQ(other_scoped->RefCnt(), 2);
            ASSERT_EQ(scoped->RefCnt(), 2);
            ASSERT_EQ(inner->RefCnt(), 2);
        }
        ASSERT_EQ(scoped->RefCnt(), 1);
        ASSERT_EQ(inner->RefCnt(), 1);

        {
            // Test copy
            auto other_scoped(scoped);
            ASSERT_EQ(other_scoped->RefCnt(), 2);
            ASSERT_EQ(scoped->RefCnt(), 2);
            ASSERT_EQ(inner->RefCnt(), 2);
        }
        ASSERT_EQ(scoped->RefCnt(), 1);
        ASSERT_EQ(inner->RefCnt(), 1);
    }
    ASSERT_EQ(inner->RefCnt(), 0);
}

TEST_F(SnapshotTest, ResourceHoldersTest) {
    milvus::engine::snapshot::ID_TYPE collection_id = 1;
    auto collection = milvus::engine::snapshot::CollectionsHolder::GetInstance().GetResource(collection_id, false);
    auto prev_cnt = collection->RefCnt();
    {
        auto collection_2 = milvus::engine::snapshot::CollectionsHolder::GetInstance().GetResource(
                collection_id, false);
        ASSERT_EQ(collection->GetID(), collection_id);
        ASSERT_EQ(collection->RefCnt(), prev_cnt);
    }

    {
        auto collection = milvus::engine::snapshot::CollectionsHolder::GetInstance().GetResource(collection_id, true);
        ASSERT_EQ(collection->GetID(), collection_id);
        ASSERT_EQ(collection->RefCnt(), 1+prev_cnt);
    }

    if (prev_cnt == 0) {
        auto collection = milvus::engine::snapshot::CollectionsHolder::GetInstance().GetResource(collection_id, false);
        ASSERT_TRUE(!collection);
    }
}

milvus::engine::snapshot::ScopedSnapshotT
CreateCollection(const std::string& collection_name, milvus::engine::snapshot::LSN_TYPE lsn) {
    milvus::engine::snapshot::CreateCollectionContext context;
    context.lsn = lsn;
    auto collection_schema = std::make_shared<milvus::engine::snapshot::Collection>(collection_name);
    context.collection = collection_schema;
    auto vector_field = std::make_shared<milvus::engine::snapshot::Field>("vector", 0);
    auto vector_field_element = std::make_shared<milvus::engine::snapshot::FieldElement>(0, 0, "ivfsq8",
            milvus::engine::snapshot::FieldElementType::IVFSQ8);
    auto int_field = std::make_shared<milvus::engine::snapshot::Field>("int", 0);
    context.fields_schema[vector_field] = {vector_field_element};
    context.fields_schema[int_field] = {};

    auto op = std::make_shared<milvus::engine::snapshot::CreateCollectionOperation>(context);
    op->Push();
    milvus::engine::snapshot::ScopedSnapshotT ss;
    auto status = op->GetSnapshot(ss);
    std::cout << status.ToString() << std::endl;
    return ss;
}

TEST_F(SnapshotTest, CreateCollectionOperationTest) {
    milvus::engine::snapshot::ScopedSnapshotT expect_null;
    auto status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(expect_null, 100000);
    ASSERT_TRUE(!expect_null);

    std::string collection_name = "test_c1";
    milvus::engine::snapshot::LSN_TYPE lsn = 1;
    auto ss = CreateCollection(collection_name, lsn);
    ASSERT_TRUE(ss);

    milvus::engine::snapshot::ScopedSnapshotT latest_ss;
    status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(latest_ss, "xxxx");
    ASSERT_TRUE(!status.ok());

    status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(latest_ss, collection_name);
    ASSERT_TRUE(latest_ss);
    ASSERT_TRUE(latest_ss->GetName() == collection_name);

    milvus::engine::snapshot::IDS_TYPE ids;
    status = milvus::engine::snapshot::Snapshots::GetInstance().GetCollectionIds(ids);
    ASSERT_EQ(ids.size(), 6);
    ASSERT_EQ(ids[5], latest_ss->GetCollectionId());

    milvus::engine::snapshot::OperationContext sd_op_ctx;
    sd_op_ctx.collection = latest_ss->GetCollection();
    sd_op_ctx.lsn = latest_ss->GetMaxLsn() + 1;
    ASSERT_TRUE(sd_op_ctx.collection->IsActive());
    auto sd_op = std::make_shared<milvus::engine::snapshot::DropCollectionOperation>(sd_op_ctx, latest_ss);
    status = sd_op->Push();
    ASSERT_TRUE(status.ok());
    ASSERT_TRUE(sd_op->GetStatus().ok());
    ASSERT_TRUE(!sd_op_ctx.collection->IsActive());
    ASSERT_TRUE(!latest_ss->GetCollection()->IsActive());

    milvus::engine::snapshot::Snapshots::GetInstance().Reset();
}

TEST_F(SnapshotTest, DropCollectionTest) {
    std::string collection_name = "test_c1";
    milvus::engine::snapshot::LSN_TYPE lsn = 1;
    auto ss = CreateCollection(collection_name, lsn);
    ASSERT_TRUE(ss);
    milvus::engine::snapshot::ScopedSnapshotT lss;
    auto status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(lss, collection_name);
    std::cout << status.ToString() << std::endl;
    ASSERT_TRUE(status.ok());
    ASSERT_TRUE(lss);
    ASSERT_EQ(ss->GetID(), lss->GetID());
    auto prev_ss_id = ss->GetID();
    auto prev_c_id = ss->GetCollection()->GetID();
    lsn = ss->GetMaxLsn() + 1;
    status = milvus::engine::snapshot::Snapshots::GetInstance().DropCollection(collection_name, lsn);
    ASSERT_TRUE(status.ok());
    status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(lss, collection_name);
    ASSERT_TRUE(!status.ok());

    auto ss_2 = CreateCollection(collection_name, ++lsn);
    status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(lss, collection_name);
    ASSERT_TRUE(status.ok());
    ASSERT_EQ(ss_2->GetID(), lss->GetID());
    ASSERT_TRUE(prev_ss_id != ss_2->GetID());
    ASSERT_TRUE(prev_c_id != ss_2->GetCollection()->GetID());
    status = milvus::engine::snapshot::Snapshots::GetInstance().DropCollection(collection_name, ++lsn);
    ASSERT_TRUE(status.ok());
    status = milvus::engine::snapshot::Snapshots::GetInstance().DropCollection(collection_name, ++lsn);
    ASSERT_TRUE(!status.ok());
}

TEST_F(SnapshotTest, ConCurrentCollectionOperation) {
    std::string collection_name("c1");
    milvus::engine::snapshot::LSN_TYPE lsn = 1;

    milvus::engine::snapshot::ID_TYPE stale_ss_id;
    auto worker1 = [&]() {
        milvus::Status status;
        auto ss = CreateCollection(collection_name, ++lsn);
        ASSERT_TRUE(ss);
        ASSERT_EQ(ss->GetName(), collection_name);
        stale_ss_id = ss->GetID();
        decltype(ss) a_ss;
        status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(a_ss, collection_name);
        ASSERT_TRUE(status.ok());
        std::this_thread::sleep_for(std::chrono::milliseconds(80));
        ASSERT_TRUE(!ss->GetCollection()->IsActive());
        status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(a_ss, collection_name);
        ASSERT_TRUE(!status.ok());

        auto c_c = milvus::engine::snapshot::CollectionCommitsHolder::GetInstance().GetResource(stale_ss_id, false);
        ASSERT_TRUE(c_c);
        ASSERT_EQ(c_c->GetID(), stale_ss_id);
    };
    auto worker2 = [&] {
        std::this_thread::sleep_for(std::chrono::milliseconds(50));
        auto status = milvus::engine::snapshot::Snapshots::GetInstance().DropCollection(collection_name, ++lsn);
        ASSERT_TRUE(status.ok());
        milvus::engine::snapshot::ScopedSnapshotT a_ss;
        status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(a_ss, collection_name);
        ASSERT_TRUE(!status.ok());
    };
    auto worker3 = [&] {
        std::this_thread::sleep_for(std::chrono::milliseconds(20));
        auto ss = CreateCollection(collection_name, ++lsn);
        ASSERT_TRUE(!ss);
        std::this_thread::sleep_for(std::chrono::milliseconds(80));
        ss = CreateCollection(collection_name, ++lsn);
        ASSERT_TRUE(ss);
        ASSERT_EQ(ss->GetName(), collection_name);
    };
    std::thread t1 = std::thread(worker1);
    std::thread t2 = std::thread(worker2);
    std::thread t3 = std::thread(worker3);
    t1.join();
    t2.join();
    t3.join();

    auto c_c = milvus::engine::snapshot::CollectionCommitsHolder::GetInstance().GetResource(stale_ss_id, false);
    ASSERT_TRUE(!c_c);
}

milvus::engine::snapshot::ScopedSnapshotT
CreatePartition(const std::string& collection_name, const milvus::engine::snapshot::PartitionContext p_context,
        const milvus::engine::snapshot::LSN_TYPE& lsn) {
    milvus::engine::snapshot::ScopedSnapshotT curr_ss;
    milvus::engine::snapshot::ScopedSnapshotT ss;
    auto status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(ss, collection_name);
    if (!status.ok()) {
        return curr_ss;
    }

    milvus::engine::snapshot::OperationContext context;
    context.lsn = lsn;
    auto op = std::make_shared<milvus::engine::snapshot::CreatePartitionOperation>(context, ss);

    milvus::engine::snapshot::PartitionPtr partition;
    status = op->CommitNewPartition(p_context, partition);
    if (!status.ok()) {
        return curr_ss;
    }

    status = op->Push();
    if (!status.ok()) {
        return curr_ss;
    }

    status = op->GetSnapshot(curr_ss);
    return curr_ss;
}

TEST_F(SnapshotTest, PartitionTest) {
    std::string collection_name("c1");
    milvus::engine::snapshot::LSN_TYPE lsn = 1;
    auto ss = CreateCollection(collection_name, ++lsn);
    ASSERT_TRUE(ss);
    ASSERT_EQ(ss->GetName(), collection_name);
    ASSERT_EQ(ss->NumberOfPartitions(), 1);

    milvus::engine::snapshot::OperationContext context;
    context.lsn = ++lsn;
    auto op = std::make_shared<milvus::engine::snapshot::CreatePartitionOperation>(context, ss);

    std::string partition_name("p1");
    milvus::engine::snapshot::PartitionContext p_ctx;
    p_ctx.name = partition_name;
    milvus::engine::snapshot::PartitionPtr partition;
    auto status = op->CommitNewPartition(p_ctx, partition);
    ASSERT_TRUE(status.ok());
    ASSERT_TRUE(partition);
    ASSERT_EQ(partition->GetName(), partition_name);
    ASSERT_TRUE(!partition->IsActive());
    ASSERT_TRUE(partition->HasAssigned());

    status = op->Push();
    ASSERT_TRUE(status.ok());
    decltype(ss) curr_ss;
    status = op->GetSnapshot(curr_ss);
    ASSERT_TRUE(status.ok());
    ASSERT_TRUE(curr_ss);
    ASSERT_EQ(curr_ss->GetName(), ss->GetName());
    ASSERT_TRUE(curr_ss->GetID() > ss->GetID());
    ASSERT_EQ(curr_ss->NumberOfPartitions(), 2);

    p_ctx.lsn = ++lsn;
    auto drop_op = std::make_shared<milvus::engine::snapshot::DropPartitionOperation>(p_ctx, curr_ss);
    status = drop_op->Push();
    ASSERT_TRUE(status.ok());

    decltype(ss) latest_ss;
    status = drop_op->GetSnapshot(latest_ss);
    ASSERT_TRUE(status.ok());
    ASSERT_TRUE(latest_ss);
    ASSERT_EQ(latest_ss->GetName(), ss->GetName());
    ASSERT_TRUE(latest_ss->GetID() > curr_ss->GetID());
    ASSERT_EQ(latest_ss->NumberOfPartitions(), 1);

    p_ctx.lsn = ++lsn;
    drop_op = std::make_shared<milvus::engine::snapshot::DropPartitionOperation>(p_ctx, latest_ss);
    status = drop_op->Push();
    ASSERT_TRUE(!status.ok());
    std::cout << status.ToString() << std::endl;

    milvus::engine::snapshot::PartitionContext pp_ctx;
    pp_ctx.name = "p2";
    curr_ss = CreatePartition(collection_name, pp_ctx, lsn-1);
    ASSERT_FALSE(curr_ss);

    std::stringstream p_name_stream;

    auto num = RandomInt(20, 30);
    for (auto i = 0; i < num; ++i) {
        p_name_stream.str("");
        p_name_stream << "partition_" << i;
        pp_ctx.name = p_name_stream.str();
        curr_ss = CreatePartition(collection_name, pp_ctx, ++lsn);
        ASSERT_TRUE(curr_ss);
        ASSERT_EQ(curr_ss->NumberOfPartitions(), 2 + i);
    }

    auto total_partition_num = curr_ss->NumberOfPartitions();

    milvus::engine::snapshot::ID_TYPE partition_id;
    for (auto i = 0; i < num; ++i) {
        p_name_stream.str("");
        p_name_stream << "partition_" << i;

        status = curr_ss->GetPartitionId(p_name_stream.str(), partition_id);
        ASSERT_TRUE(status.ok());
        status = milvus::engine::snapshot::Snapshots::GetInstance().DropPartition(
                curr_ss->GetCollectionId(), partition_id, ++lsn);
        ASSERT_TRUE(status.ok());
        status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(
                curr_ss, curr_ss->GetCollectionId());
        ASSERT_TRUE(status.ok());
        ASSERT_EQ(curr_ss->NumberOfPartitions(), total_partition_num - i -1);
    }

}

TEST_F(SnapshotTest, PartitionTest2) {
    std::string collection_name("c1");
    milvus::engine::snapshot::LSN_TYPE lsn = 1;
    milvus::Status status;

    auto ss = CreateCollection(collection_name, ++lsn);
    ASSERT_TRUE(ss);
    ASSERT_EQ(lsn, ss->GetMaxLsn());

    milvus::engine::snapshot::OperationContext context;
    context.lsn = lsn;
    auto cp_op = std::make_shared<milvus::engine::snapshot::CreatePartitionOperation>(context, ss);
    std::string partition_name("p1");
    milvus::engine::snapshot::PartitionContext p_ctx;
    p_ctx.name = partition_name;
    milvus::engine::snapshot::PartitionPtr partition;
    status = cp_op->CommitNewPartition(p_ctx, partition);
    ASSERT_TRUE(status.ok());
    ASSERT_TRUE(partition);
    ASSERT_EQ(partition->GetName(), partition_name);
    ASSERT_TRUE(!partition->IsActive());
    ASSERT_TRUE(partition->HasAssigned());

    status = cp_op->Push();
    ASSERT_TRUE(!status.ok());
}

TEST_F(SnapshotTest, OperationTest) {
    milvus::Status status;
    std::string to_string;
    milvus::engine::snapshot::LSN_TYPE lsn;
    milvus::engine::snapshot::SegmentFileContext sf_context;
    sf_context.field_name = "f_1_1";
    sf_context.field_element_name = "fe_1_1";
    sf_context.segment_id = 1;
    sf_context.partition_id = 1;

    milvus::engine::snapshot::ScopedSnapshotT ss;
    status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(ss, 1);
    std::cout << status.ToString() << std::endl;
    ASSERT_TRUE(status.ok());
    auto ss_id = ss->GetID();
    lsn = ss->GetMaxLsn() + 1;
    ASSERT_TRUE(status.ok());

    // Check snapshot
    {
        auto collection_commit = milvus::engine::snapshot::CollectionCommitsHolder::GetInstance()
            .GetResource(ss_id, false);
        /* snapshot::SegmentCommitsHolder::GetInstance().GetResource(prev_segment_commit->GetID()); */
        ASSERT_TRUE(collection_commit);
        to_string = collection_commit->ToString();
        ASSERT_EQ(to_string, "");
    }

    milvus::engine::snapshot::OperationContext merge_ctx;
    std::set<milvus::engine::snapshot::ID_TYPE> stale_segment_commit_ids;

    decltype(sf_context.segment_id) new_seg_id;
    decltype(ss) new_ss;
    // Check build operation correctness
    {
        milvus::engine::snapshot::OperationContext context;
        context.lsn = ++lsn;
        auto build_op = std::make_shared<milvus::engine::snapshot::BuildOperation>(context, ss);
        milvus::engine::snapshot::SegmentFilePtr seg_file;
        status = build_op->CommitNewSegmentFile(sf_context, seg_file);
        ASSERT_TRUE(status.ok());
        ASSERT_TRUE(seg_file);
        auto prev_segment_commit = ss->GetSegmentCommit(seg_file->GetSegmentId());
        auto prev_segment_commit_mappings = prev_segment_commit->GetMappings();
        ASSERT_NE(prev_segment_commit->ToString(), "");

        build_op->Push();
        status = build_op->GetSnapshot(ss);
        ASSERT_TRUE(ss->GetID() > ss_id);

        auto segment_commit = ss->GetSegmentCommit(seg_file->GetSegmentId());
        auto segment_commit_mappings = segment_commit->GetMappings();
        milvus::engine::snapshot::MappingT expected_mappings = prev_segment_commit_mappings;
        expected_mappings.insert(seg_file->GetID());
        ASSERT_EQ(expected_mappings, segment_commit_mappings);

        auto seg = ss->GetResource<milvus::engine::snapshot::Segment>(seg_file->GetSegmentId());
        ASSERT_TRUE(seg);
        merge_ctx.stale_segments.push_back(seg);
        stale_segment_commit_ids.insert(segment_commit->GetID());
    }

    // Check stale snapshot has been deleted from store
    {
        auto collection_commit = milvus::engine::snapshot::CollectionCommitsHolder::GetInstance()
            .GetResource(ss_id, false);
        ASSERT_TRUE(!collection_commit);
    }

    ss_id = ss->GetID();
    milvus::engine::snapshot::ID_TYPE partition_id;
    {
        milvus::engine::snapshot::OperationContext context;
        context.lsn = ++lsn;
        context.prev_partition = ss->GetResource<milvus::engine::snapshot::Partition>(1);
        auto op = std::make_shared<milvus::engine::snapshot::NewSegmentOperation>(context, ss);
        milvus::engine::snapshot::SegmentPtr new_seg;
        status = op->CommitNewSegment(new_seg);
        ASSERT_TRUE(status.ok());
        ASSERT_NE(new_seg->ToString(), "");
        milvus::engine::snapshot::SegmentFilePtr seg_file;
        status = op->CommitNewSegmentFile(sf_context, seg_file);
        ASSERT_TRUE(status.ok());
        status = op->Push();
        ASSERT_TRUE(status.ok());

        status = op->GetSnapshot(ss);
        ASSERT_TRUE(ss->GetID() > ss_id);
        ASSERT_TRUE(status.ok());

        auto segment_commit = ss->GetSegmentCommit(seg_file->GetSegmentId());
        auto segment_commit_mappings = segment_commit->GetMappings();
        milvus::engine::snapshot::MappingT expected_segment_mappings;
        expected_segment_mappings.insert(seg_file->GetID());
        ASSERT_EQ(expected_segment_mappings, segment_commit_mappings);
        merge_ctx.stale_segments.push_back(new_seg);
        partition_id = segment_commit->GetPartitionId();
        stale_segment_commit_ids.insert(segment_commit->GetID());
        auto partition = ss->GetResource<milvus::engine::snapshot::Partition>(partition_id);
        merge_ctx.prev_partition = partition;
        new_seg_id = seg_file->GetSegmentId();
        new_ss = ss;
    }

    milvus::engine::snapshot::SegmentPtr merge_seg;
    ss_id = ss->GetID();
    {
        auto prev_partition_commit = ss->GetPartitionCommitByPartitionId(partition_id);
        auto expect_null = ss->GetPartitionCommitByPartitionId(11111111);
        ASSERT_TRUE(!expect_null);
        ASSERT_NE(prev_partition_commit->ToString(), "");
        merge_ctx.lsn = ++lsn;
        auto op = std::make_shared<milvus::engine::snapshot::MergeOperation>(merge_ctx, ss);
        milvus::engine::snapshot::SegmentPtr new_seg;
        status = op->CommitNewSegment(new_seg);
        sf_context.segment_id = new_seg->GetID();
        milvus::engine::snapshot::SegmentFilePtr seg_file;
        status = op->CommitNewSegmentFile(sf_context, seg_file);
        ASSERT_TRUE(status.ok());
        status = op->Push();
        ASSERT_TRUE(status.ok());
        std::cout << op->ToString() << std::endl;
        status = op->GetSnapshot(ss);
        ASSERT_TRUE(ss->GetID() > ss_id);
        ASSERT_TRUE(status.ok());

        auto segment_commit = ss->GetSegmentCommit(new_seg->GetID());
        auto new_partition_commit = ss->GetPartitionCommitByPartitionId(partition_id);
        auto new_mappings = new_partition_commit->GetMappings();
        auto prev_mappings = prev_partition_commit->GetMappings();
        auto expected_mappings = prev_mappings;
        for (auto id : stale_segment_commit_ids) {
            expected_mappings.erase(id);
        }
        expected_mappings.insert(segment_commit->GetID());
        ASSERT_EQ(expected_mappings, new_mappings);

        milvus::engine::snapshot::CollectionCommitsHolder::GetInstance().Dump();
        merge_seg = new_seg;
    }

    // 1. New seg1, seg2
    // 2. Build seg1 start
    // 3. Merge seg1, seg2 to seg3
    // 4. Commit new seg file of build operation -> Stale Segment Found Here!
    {
        milvus::engine::snapshot::OperationContext context;
        context.lsn = ++lsn;
        auto build_op = std::make_shared<milvus::engine::snapshot::BuildOperation>(context, new_ss);
        milvus::engine::snapshot::SegmentFilePtr seg_file;
        auto new_sf_context = sf_context;
        new_sf_context.segment_id = new_seg_id;
        status = build_op->CommitNewSegmentFile(new_sf_context, seg_file);
        ASSERT_TRUE(!status.ok());
    }

    // 1. Build start
    // 2. Commit new seg file of build operation
    // 3. Drop collection
    // 4. Commit build operation -> Stale Segment Found Here!
    {
        milvus::engine::snapshot::OperationContext context;
        context.lsn = ++lsn;
        auto build_op = std::make_shared<milvus::engine::snapshot::BuildOperation>(context, ss);
        milvus::engine::snapshot::SegmentFilePtr seg_file;
        auto new_sf_context = sf_context;
        new_sf_context.segment_id = merge_seg->GetID();
        status = build_op->CommitNewSegmentFile(new_sf_context, seg_file);
        ASSERT_TRUE(status.ok());
        std::cout << build_op->ToString() << std::endl;

        auto status = milvus::engine::snapshot::Snapshots::GetInstance().DropCollection(ss->GetName(),
                ++lsn);
        ASSERT_TRUE(status.ok());
        status = build_op->Push();
        ASSERT_TRUE(!status.ok());
        ASSERT_TRUE(!(build_op->GetStatus()).ok());
        std::cout << build_op->ToString() << std::endl;
    }
    milvus::engine::snapshot::Snapshots::GetInstance().Reset();
}

struct WaitableObj {
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
    Notify() {
        std::unique_lock<std::mutex> lck(mutex_);
        notified_ = true;
        lck.unlock();
        cv_.notify_one();
    }
};


#if 1
TEST_F(SnapshotTest, CompoundTest1) {
    milvus::Status status;
    milvus::engine::snapshot::LSN_TYPE lsn = 0;
    auto next_lsn = [&]() -> decltype(lsn) {
        return ++lsn;
    };
    std::string collection_name("c1");
    auto ss = CreateCollection(collection_name, next_lsn());
    ASSERT_TRUE(ss);
    ASSERT_EQ(lsn, ss->GetMaxLsn());

    using ID_TYPE = milvus::engine::snapshot::ID_TYPE;
    using SegmentFileContext = milvus::engine::snapshot::SegmentFileContext;
    using OperationContext =  milvus::engine::snapshot::OperationContext;
    using NewSegmentOperation = milvus::engine::snapshot::NewSegmentOperation;
    using Queue = milvus::server::BlockingQueue<ID_TYPE>;
    using Partition = milvus::engine::snapshot::Partition;
    using SegmentPtr = milvus::engine::snapshot::SegmentPtr;
    using SegmentFilePtr = milvus::engine::snapshot::SegmentFilePtr;
    Queue merge_queue;

    std::set<ID_TYPE> all_segments;
    std::set<ID_TYPE> segment_in_building;
    std::set<ID_TYPE> merge_segs;
    std::map<ID_TYPE, std::set<ID_TYPE>> merged_segs;

    std::mutex all_mtx;
    std::mutex building_mtx;
    std::mutex merging_mtx;

    WaitableObj w_l;

    SegmentFileContext sf_context;
    sf_context.field_name = "vector";
    sf_context.field_element_name = "ivfsq8";
    sf_context.segment_id = 1;
    sf_context.partition_id = 1;

    auto do_merge = [&] (std::set<ID_TYPE>& seg_ids) {
        if (seg_ids.size() == 0) {
            return;
        }
        decltype(ss) latest_ss;
        auto status = milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(latest_ss, collection_name);
        ASSERT_TRUE(status.ok());

        OperationContext context;
        for (auto& id : seg_ids) {
            auto seg = latest_ss->GetResource<milvus::engine::snapshot::Segment>(id);
            if (!seg) {
                std::cout << "Error seg=" << id << std::endl;
                ASSERT_TRUE(seg);
            }
            context.stale_segments.push_back(seg);
            if (!context.prev_partition) {
                context.prev_partition = latest_ss->GetResource<milvus::engine::snapshot::Partition>(
                        seg->GetPartitionId());
            }
        }

        context.lsn = next_lsn();
        auto op = std::make_shared<milvus::engine::snapshot::MergeOperation>(context, latest_ss);
        milvus::engine::snapshot::SegmentPtr new_seg;
        status = op->CommitNewSegment(new_seg);
        ASSERT_TRUE(status.ok());
        sf_context.segment_id = new_seg->GetID();
        milvus::engine::snapshot::SegmentFilePtr seg_file;
        status = op->CommitNewSegmentFile(sf_context, seg_file);
        ASSERT_TRUE(status.ok());
        status = op->Push();
        ASSERT_TRUE(status.ok());
        ID_TYPE ss_id = latest_ss->GetID();
        status = op->GetSnapshot(latest_ss);
        ASSERT_TRUE(status.ok());
        ASSERT_TRUE(latest_ss->GetID() > ss_id);
        latest_ss->DumpResource<milvus::engine::snapshot::Segment>("do_merge");
        merged_segs[new_seg->GetID()] = seg_ids;
    };

    // TODO: If any Compound Operation find larger Snapshot. This Operation should be rollback to latest
    auto normal_worker = [&] {
        auto to_build_segments = RandomInt(10, 11);
        decltype(ss) latest_ss;

        for (auto i=0; i<to_build_segments; ++i) {
            milvus::engine::snapshot::Snapshots::GetInstance().GetSnapshot(latest_ss, collection_name);
            OperationContext context;
            context.lsn = next_lsn();
            context.prev_partition = latest_ss->GetResource<Partition>(8);
            auto op = std::make_shared<NewSegmentOperation>(context, latest_ss);
            SegmentPtr new_seg;
            status = op->CommitNewSegment(new_seg);
            ASSERT_TRUE(status.ok());
            SegmentFilePtr seg_file;
            sf_context.segment_id = new_seg->GetID();
            op->CommitNewSegmentFile(sf_context, seg_file);
            op->Push();
            status = op->GetSnapshot(latest_ss);
            ASSERT_TRUE(status.ok());
            latest_ss->DumpResource<milvus::engine::snapshot::Segment>("normal_worker");

            {
                std::unique_lock<std::mutex> lock(all_mtx);
                all_segments.insert(new_seg->GetID());
            }
            merge_queue.Put(new_seg->GetID());
        }

        merge_queue.Put(0);
    };

    auto merge_worker = [&] {
        while (true) {
            auto seg_id = merge_queue.Take();
            if (seg_id == 0) {
                std::cout << "Exiting Merge Worker" << std::endl;
                break;
            }
            merge_segs.insert(seg_id);
            if ((merge_segs.size() >= 2) && (RandomInt(0, 10) >= 5)) {
                std::cout << "Merging (";
                for (auto seg : merge_segs) {
                    std::cout << seg << ",";
                }
                std::cout << ")" << std::endl;
                do_merge(merge_segs);
                merge_segs.clear();

            } else {
                continue;
            }
        }
        w_l.Notify();
    };

    std::thread t1 = std::thread(normal_worker);
    std::thread t2 = std::thread(merge_worker);
    t1.join();
    t2.join();

    for (auto sid : all_segments) {
        std::cout << "no seg " << sid << std::endl;
    }

    for (auto& kv : merged_segs) {
        std::cout << "merged: (";
        for (auto i : kv.second) {
            std::cout << i << ",";
        }
        std::cout << ") -> " << kv.first << std::endl;
    }

    w_l.Wait();
}
#endif
