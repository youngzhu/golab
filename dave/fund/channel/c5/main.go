package main

func main() {
	var c = make(chan int, 100)
	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				c <- i * j
			}
			close(c)
		}()
	}

	for i := range c {
		println(i)
	}
}
