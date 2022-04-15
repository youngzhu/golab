package main

// 关闭的channel不会产生阻塞

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	//close(ch)

	//for i := 0; i < cap(ch)+2; i++ {
	//	v, ok := <-ch
	//	println(v, ok)
	//}

	for v := range ch {
		println(v) // called twice
	}
}
