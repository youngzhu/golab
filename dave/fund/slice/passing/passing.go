package main

import "fmt"

func negate(s []int) {
	for i := range s {
		s[i] = -s[i]
	}
}

func main() {
	var a = []int{1, 2, 3, 4, 5}
	negate(a)
	fmt.Println(a) // [-1 -2 -3 -4 -5]

}
