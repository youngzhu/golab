package main

import (
	"sync"
)

func main() {
	counter := 0
	var mu sync.Mutex
	for i := 0; i < 10000000; i++ {
		go func() {
			mu.Lock()
			//defer mu.Unlock()
			counter = counter + 1
			mu.Unlock()
		}()
	}
	//time.Sleep(1 * time.Second)
	// 这里应该不需要锁啊。不知道示例中为什么有
	// 用大数测试，有没有锁都是正确的
	// 但没有sleep，即使有锁，打印的值也是错误的。defer或另一种方式都是错
	mu.Lock()
	println(counter)
	mu.Unlock()
}
