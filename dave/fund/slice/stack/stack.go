package main

import "fmt"

func f(s []string, level int) {
	if level > 5 {
		return
	}
	s = append(s, fmt.Sprint(level))
	f(s, level+1)
	fmt.Println("level:", level, "slice:", s)
}

func main() {
	f(nil, 0)
}
