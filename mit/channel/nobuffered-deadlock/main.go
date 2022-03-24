package main

func main() {
	c := make(chan bool)
	//go func() {
	//	for {
	//	}
	//}()
	<-c
	c <- true
}
