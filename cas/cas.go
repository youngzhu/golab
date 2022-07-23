package cas

import "sync"

type Simulator struct {
	value int
	mu    sync.Mutex
}

func (s *Simulator) Get() int {
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.value
}

func (s *Simulator) CompareAndSwap(expectedVal, newVal int) int {
	s.mu.Lock()
	defer s.mu.Unlock()

	oldVal := s.value
	if oldVal == expectedVal {
		s.value = newVal
	}
	return oldVal
}

func (s *Simulator) CompareAndSet(expectedVal, newVal int) bool {
	return expectedVal == s.CompareAndSwap(expectedVal, newVal)
}

// Counter 基于CAS实现的非阻塞的计数器
type Counter struct {
	value *Simulator
}

func NewCounter() *Counter {
	return &Counter{&Simulator{}}
}

func (c *Counter) Get() int {
	return c.value.Get()
}

func (c *Counter) Increment() int {
	var v int
	for {
		v = c.value.Get()
		if c.value.CompareAndSet(v, v+1) {
			break
		}
	}
	return v + 1
}
