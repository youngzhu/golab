package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	const n = 1000
	//finish := make(chan bool)
	finish := make(chan struct{}) // 只关心是否关闭，不管里面的值
	var done sync.WaitGroup
	for i := 0; i < n; i++ {
		done.Add(1)

		go func() {
			select {
			case <-time.After(1 * time.Hour):
			case <-finish:
			}
			done.Done()
		}()
	}

	t0 := time.Now()
	close(finish) // close chan makes it ready to receive from
	done.Wait()   // wait for all goroutine to stop
	fmt.Printf("waited %v for %d goroutines to stop", time.Since(t0), n)
}
