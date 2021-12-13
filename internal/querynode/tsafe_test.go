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

package querynode

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/milvus-io/milvus/internal/util/typeutil"
)

func TestTSafe_TSafeWatcher(t *testing.T) {
	watcher := newTSafeWatcher()
	defer watcher.close()
	assert.NotNil(t, watcher)

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		watcher.notify()
		wg.Done()
	}()
	wg.Wait()

	// wait notify, expect non-block here
	<-watcher.watcherChan()
}

func TestTSafe_TSafe(t *testing.T) {
	safe := newTSafe("TestTSafe-channel")
	assert.NotNil(t, safe)

	timestamp := safe.get()
	assert.Equal(t, typeutil.ZeroTimestamp, timestamp)

	watcher := newTSafeWatcher()
	defer watcher.close()
	assert.NotNil(t, watcher)

	safe.registerTSafeWatcher(watcher)
	assert.Len(t, safe.watcherList, 1)

	targetTimestamp := Timestamp(1000)
	safe.set(targetTimestamp)

	timestamp = safe.get()
	assert.Equal(t, targetTimestamp, timestamp)
}
