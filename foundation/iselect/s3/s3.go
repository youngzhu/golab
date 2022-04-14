package main

import "time"

func process1(ch chan string) {
	ch <- "from p1"
}
func process2(ch chan string) {
	ch <- "from p2"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go process1(c1)
	go process2(c2)

	// 保证两个协程都执行过了
	time.Sleep(2 * time.Second)

	select {
	case s1 := <-c1:
		println(s1)
	case s2 := <-c2:
		println(s2)
	}
}
