package mutex

import "sync"

type AtomicVariable struct {
	mu  sync.Mutex
	val uint64
}

func (av *AtomicVariable) Inc() {
	av.mu.Lock()
	av.val++
	av.mu.Unlock()
}
