package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int)

	for i := 0; i < 4; i++ {
		go doWork(c)
	}

	for {
		v := <-c
		fmt.Printf("consumer:%v\n", v)
	}
}

func doWork(c chan int) {
	for {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		p := rand.Int()
		c <- p
		fmt.Printf("producer:%v\n", p)
	}
}
