package main

import "fmt"

func main() {
	var a [5]int
	b := a
	b[2] = 2
	fmt.Println(a, b) // [0 0 0 0 0] [0 0 2 0 0]

}
