package main

import "fmt"

func double(s []int) {
	s = append(s, s...)
}

func main() {
	a := []int{1, 2, 3}
	double(a)
	fmt.Println(a)
}
