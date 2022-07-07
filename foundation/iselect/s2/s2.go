package main

import (
	"log"
	"time"
)

// default 防止永久性阻塞

func process(in chan string) {
	for {
		time.Sleep(4 * time.Second)
		in <- "input"
	}
}

func process2(in chan string) {
	for {
		time.Sleep(5 * time.Second)
		in <- "input2"
	}
}

func main() {
	in := make(chan string)
	in2 := make(chan string)

	go process(in)
	go process2(in2)

	for {
		time.Sleep(2 * time.Second)
		select {
		case s := <-in:
			log.Println("output:", s)
		case s2 := <-in2:
			log.Println("output:", s2)
		default:
			log.Println("wait...")
		}
	}
}
