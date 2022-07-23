package main

import (
	"sync"
	"time"
)

// 问题: A协程产生另一个协程B，A结束了，B还存在吗？
// 经验证：B不受A的影响

type Counter struct {
	count int
	mu    sync.Mutex
}

func (c *Counter) Increment() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.count++
}

func (c *Counter) Get() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.count
}

func main() {
	i := &Counter{}

	go count(i)

	for {
		i.Increment()
		println("main:", i.Get())
		time.Sleep(time.Second)
	}

}

func count(i *Counter) {

	for i.Get() < 10 {
		i.Increment()
		println("father:", i.Get())
		time.Sleep(time.Second)
	}

	go func() {
		for {
			i.Increment()
			println("son:", i.Get())
			time.Sleep(time.Second)
		}
	}()

	println("father eeeeee")
}
