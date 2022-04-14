package main

import "time"

// default 防止永久性阻塞

func process(in chan string) {
	for {
		time.Sleep(4 * time.Second)
		in <- "input"
	}
}

func main() {
	in := make(chan string)

	go process(in)

	for {
		time.Sleep(2 * time.Second)
		select {
		case s := <-in:
			println("output:", s)
		default:
			println("wait...")
		}
	}
}
