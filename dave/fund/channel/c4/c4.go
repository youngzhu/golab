package main

import (
	"fmt"
	"time"
)

// waits for a and b to close.

// 关闭的chan可以一直获取，所以这个实现是不对的
//func waitMany(a, b chan bool) {
//	var aclosed, bclosed bool
//	for !aclosed || !bclosed {
//		select {
//		case <-a:
//			aclosed = true
//		case <-b:
//			bclosed = true
//		}
//	}
//}

func waitMany(a, b chan bool) {
	for a != nil || b != nil {
		select {
		case <-a:
			a = nil
		case <-b:
			b = nil
		}
	}
}
func main() {
	a := make(chan bool)
	b := make(chan bool)
	t0 := time.Now()
	go func() {
		close(a)
		close(b)
	}()
	waitMany(a, b)
	fmt.Printf("wait %v for waitMany", time.Since(t0))
}
