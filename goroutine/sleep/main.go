package main

import "time"

func main() {
	time.Sleep(1 * time.Second)
	println("started")
	go periodic()
	// 为了观察 periodic()
	time.Sleep(2 * time.Second)
	println("done")
}

func periodic() {
	for {
		println("tick")
		time.Sleep(2 * time.Second)
	}
}
