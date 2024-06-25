package syncmgr

import (
	"github.com/milvus-io/milvus-proto/go-api/v2/msgpb"
	"github.com/milvus-io/milvus/pkg/util/conc"
	"github.com/milvus-io/milvus/pkg/util/lock"
)

//go:generate mockery --name=Task --structname=MockTask --output=./  --filename=mock_task.go --with-expecter --inpackage
type Task interface {
	SegmentID() int64
	Checkpoint() *msgpb.MsgPosition
	StartPosition() *msgpb.MsgPosition
	ChannelName() string
	Run() error
	HandleError(error)
}

type keyLockDispatcher[K comparable] struct {
	keyLock    *lock.KeyLock[K]
	workerPool *conc.Pool[struct{}]
}

func newKeyLockDispatcher[K comparable](maxParallel int) *keyLockDispatcher[K] {
	dispatcher := &keyLockDispatcher[K]{
		workerPool: conc.NewPool[struct{}](maxParallel, conc.WithPreAlloc(false)),
		keyLock:    lock.NewKeyLock[K](),
	}
	return dispatcher
}

func (d *keyLockDispatcher[K]) Submit(key K, t Task, callbacks ...func(error) error) *conc.Future[struct{}] {
	d.keyLock.Lock(key)

	return d.workerPool.Submit(func() (struct{}, error) {
		defer d.keyLock.Unlock(key)
		err := t.Run()

		for _, callback := range callbacks {
			err = callback(err)
		}

		return struct{}{}, err
	})
}
