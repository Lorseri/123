/**
 * Copyright (c) Facebook, Inc. and its affiliates.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the root directory of this source tree.
 */

#include <faiss/gpu/utils/blockselect/BlockSelectImpl.cuh>
#include <faiss/gpu/utils/DeviceDefs.cuh>

namespace faiss { namespace gpu {

#if GPU_MAX_SELECTION_K >= 2048
#ifdef FAISS_USE_FLOAT16
BLOCK_SELECT_IMPL(half, false, 2048, 8);
#endif
#endif

} } // namespace
