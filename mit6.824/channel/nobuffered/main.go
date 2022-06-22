package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan bool)
	go func() {
		time.Sleep(2 * time.Second)
		<-c
	}()
	start := time.Now()
	c <- true // 阻塞，直到goroutine中chan有输出
	fmt.Printf("send took %v\n", time.Since(start))
}
