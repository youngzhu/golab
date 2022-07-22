package main

import (
	"log"
	"sync"
	"time"
)

// Wait() 放在 for 里和 if 里有啥区别？
// 要用for，每次Broadcast()都会唤醒一次

var sharedRsc = make(map[string]interface{})

func main() {
	c := sync.NewCond(&sync.Mutex{})

	go func() {
		for {
			time.Sleep(time.Second)
			c.L.Lock()
			c.Broadcast()
			c.L.Unlock()
		}
	}()

	c.L.Lock()
	//for len(sharedRsc) == 0 {
	if len(sharedRsc) == 0 {
		log.Println("wait...")
		c.Wait()
	}
	c.L.Unlock()

	log.Println("done")
}
