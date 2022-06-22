package main

// 效果同WaitGroup
func main() {
	done := make(chan bool)

	for i := 0; i < 5; i++ {
		go func(x int) {
			sendRPC(x)
			done <- true
		}(i)
	}

	for i := 0; i < 5; i++ {
		<-done
	}

	println("done")
}

func sendRPC(i int) {
	println(i)
}
