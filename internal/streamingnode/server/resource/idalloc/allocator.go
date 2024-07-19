// Licensed to the LF AI & Data foundation under one
// or more contributor license agreements. See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership. The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License. You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package idalloc

import (
	"sync"

	"github.com/milvus-io/milvus/internal/allocator"
	"github.com/milvus-io/milvus/internal/types"
)

// batchAllocateSize is the size of batch allocate from remote allocator.
const batchAllocateSize = 1000

var _ Allocator = (*allocatorImpl)(nil)

// NewTSOAllocator creates a new allocator.
func NewTSOAllocator(rc types.RootCoordClient) Allocator {
	return &allocatorImpl{
		mu:              sync.Mutex{},
		remoteAllocator: newTSOAllocator(rc),
		localAllocator:  newLocalAllocator(),
	}
}

// NewIDAllocator creates a new allocator.
func NewIDAllocator(rc types.RootCoordClient) Allocator {
	return &allocatorImpl{
		mu:              sync.Mutex{},
		remoteAllocator: newIDAllocator(rc),
		localAllocator:  newLocalAllocator(),
	}
}

type remoteBatchAllocator interface {
	batchAllocate(count uint32) (int64, int, error)
}

type Allocator interface {
	allocator.Interface
	// Sync expire the local allocator messages,
	// syncs the local allocator and remote allocator.
	Sync()
}

type allocatorImpl struct {
	mu              sync.Mutex
	remoteAllocator remoteBatchAllocator
	localAllocator  *localAllocator
}

func (ta *allocatorImpl) Alloc(count uint32) (int64, int64, error) {
	panic("TODO: implement me")
}

// AllocOne allocates a timestamp.
func (ta *allocatorImpl) AllocOne() (int64, error) {
	ta.mu.Lock()
	defer ta.mu.Unlock()

	// allocate one from local allocator first.
	if id, err := ta.localAllocator.allocateOne(); err == nil {
		return id, nil
	}
	// allocate from remote.
	return ta.allocateRemote()
}

// Sync expire the local allocator messages,
// syncs the local allocator and remote allocator.
func (ta *allocatorImpl) Sync() {
	ta.mu.Lock()
	defer ta.mu.Unlock()

	ta.localAllocator.exhausted()
}

// allocateRemote allocates timestamp from remote root coordinator.
func (ta *allocatorImpl) allocateRemote() (int64, error) {
	// Update local allocator from remote.
	start, count, err := ta.remoteAllocator.batchAllocate(batchAllocateSize)
	if err != nil {
		return 0, err
	}
	ta.localAllocator.update(start, count)

	// Get from local again.
	return ta.localAllocator.allocateOne()
}
