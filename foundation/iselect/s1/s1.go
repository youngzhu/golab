package main

import "time"

func server1(ch chan string) {
	time.Sleep(4 * time.Second)
	ch <- "from server1"
}
func server2(ch chan string) {
	time.Sleep(2 * time.Second)
	ch <- "from server2"
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go server1(c1)
	go server2(c2)
	select {
	case s1 := <-c1:
		println(s1)
	case s2 := <-c2:
		println(s2)
	}

	//time.Sleep(5 * time.Second)
}
