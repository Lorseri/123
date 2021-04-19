
package id

import (
	"github.com/zilliztech/milvus-distributed/internal/kv"
	"github.com/zilliztech/milvus-distributed/internal/master/tso"
	"github.com/zilliztech/milvus-distributed/internal/util/typeutil"
)


type UniqueID = typeutil.UniqueID

// GlobalTSOAllocator is the global single point TSO allocator.
type GlobalIdAllocator struct {
	allocator tso.Allocator
}

var allocator *GlobalIdAllocator

func InitGlobalIdAllocator(key string, base kv.KVBase){
	allocator = NewGlobalIdAllocator(key, base)
}

func NewGlobalIdAllocator(key string, base kv.KVBase) * GlobalIdAllocator{
	return &GlobalIdAllocator{
		allocator: tso.NewGlobalTSOAllocator( key, base),
	}
}

// Initialize will initialize the created global TSO allocator.
func (gia *GlobalIdAllocator) Initialize() error {
	return gia.allocator.Initialize()
}

// GenerateTSO is used to generate a given number of TSOs.
// Make sure you have initialized the TSO allocator before calling.
func (gia *GlobalIdAllocator) Alloc(count uint32) (UniqueID, UniqueID, error) {
	timestamp, err := gia.allocator.GenerateTSO(count)
	if err != nil {
		return 0, 0, err
	}
	idStart := UniqueID(timestamp)
	idEnd := idStart + int64(count)
	return idStart, idEnd, nil
}

func (gia *GlobalIdAllocator) AllocOne() (UniqueID, error) {
	timestamp, err := gia.allocator.GenerateTSO(1)
	if err != nil {
		return 0, err
	}
	idStart := UniqueID(timestamp)
	return idStart, nil
}

func AllocOne() (UniqueID, error) {
	return allocator.AllocOne()
}

func Alloc(count uint32) (UniqueID, UniqueID, error) {
	return allocator.Alloc(count)
}
