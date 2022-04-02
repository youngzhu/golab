package main

import "fmt"

func main() {
	var a = []int{1, 2, 3, 4, 5}
	b := a[2:]
	b[0] = 0
	fmt.Println(a, b) //[1 2 0 4 5] [0 4 5]

}
