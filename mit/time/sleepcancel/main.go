package main

import (
	"sync"
	"time"
)

var (
	done bool
	mu   sync.Mutex
)

func main() {
	time.Sleep(1 * time.Second)
	println("started")
	go periodic()
	// 为了观察 periodic()
	//time.Sleep(5 * time.Second) // 没有这个停顿，输出就很奇怪
	mu.Lock()
	done = true
	mu.Unlock()
	println("cancelled")
	time.Sleep(3 * time.Second) // 观察是否还有输出
}

func periodic() {
	for {
		println("tick")
		time.Sleep(1 * time.Second)
		mu.Lock()
		if done {
			mu.Unlock()
			return
		}
		mu.Unlock()
	}
}
