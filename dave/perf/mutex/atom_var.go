package mutex

import (
	"sync"
	"sync/atomic"
)

type AtomicVariable struct {
	mu  sync.Mutex
	val uint64
}

func (av *AtomicVariable) Inc() {
	av.mu.Lock()
	av.val++
	av.mu.Unlock()
}

type AtomicCounter uint64

func (c *AtomicCounter) Get() uint64 {
	return atomic.LoadUint64((*uint64)(c))
}

func (c *AtomicCounter) Inc() uint64 {
	return atomic.AddUint64((*uint64)(c), 1)
}

func (c *AtomicCounter) Reset() uint64 {
	return atomic.SwapUint64((*uint64)(c), 0)
}
