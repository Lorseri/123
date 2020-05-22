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

using namespace milvus::engine;

TEST_F(SnapshotTest, ReferenceProxyTest) {
    std::string status("raw");
    const std::string CALLED = "CALLED";
    auto callback = [&]() {
        status = CALLED;
    };

    auto proxy = snapshot::ReferenceProxy();
    ASSERT_EQ(proxy.RefCnt(), 0);

    int refcnt = 3;
    for (auto i=0; i<refcnt; ++i) {
        proxy.Ref();
    }
    ASSERT_EQ(proxy.RefCnt(), refcnt);

    proxy.RegisterOnNoRefCB(callback);

    for (auto i=0; i<refcnt; ++i) {
        proxy.UnRef();
    }
    ASSERT_EQ(proxy.RefCnt(), 0);
    ASSERT_EQ(status, CALLED);
}

TEST_F(SnapshotTest, ScopedResourceTest) {
    auto inner = std::make_shared<snapshot::Collection>("c1");
    ASSERT_EQ(inner->RefCnt(), 0);

    {
        auto not_scoped = snapshot::CollectionScopedT(inner, false);
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
        auto scoped = snapshot::CollectionScopedT(inner);
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
    snapshot::ID_TYPE collection_id = 1;
    {
        auto collection = snapshot::CollectionsHolder::GetInstance().GetResource(collection_id, false);
        ASSERT_EQ(collection->GetID(), collection_id);
        ASSERT_EQ(collection->RefCnt(), 0);
    }

    {
        auto collection = snapshot::CollectionsHolder::GetInstance().GetResource(collection_id, true);
        ASSERT_EQ(collection->GetID(), collection_id);
        ASSERT_EQ(collection->RefCnt(), 1);
    }

    {
        auto collection = snapshot::CollectionsHolder::GetInstance().GetResource(collection_id, false);
        ASSERT_TRUE(!collection);
    }
}

TEST_F(SnapshotTest, OperationTest) {
    {
        snapshot::SegmentFileContext sf_context;
        sf_context.field_name = "f_1_1";
        sf_context.field_element_name = "fe_1_1";
        sf_context.segment_id = 1;
        sf_context.partition_id = 1;

        auto ss = snapshot::Snapshots::GetInstance().GetSnapshot(1);
        auto ss_id = ss->GetID();

        // Check snapshot
        {
            auto collection_commit = snapshot::CollectionCommitsHolder::GetInstance().GetResource(ss_id, false);
            ASSERT_TRUE(collection_commit);
        }

        // Check build operation correctness
        {
            snapshot::OperationContext context;
            auto build_op = std::make_shared<snapshot::BuildOperation>(context, ss);
            auto seg_file = build_op->CommitNewSegmentFile(sf_context);
            build_op->Push();
            ss = build_op->GetSnapshot();
            ASSERT_TRUE(ss->GetID() > ss_id);
        }

        // Check stale snapshot has been deleted from store
        {
            auto collection_commit = snapshot::CollectionCommitsHolder::GetInstance().GetResource(ss_id, false);
            ASSERT_TRUE(!collection_commit);
        }

    }
}
