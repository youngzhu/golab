package main

import "fmt"

func main() {
	a := 200
	b := &a
	*b++
	fmt.Println(a, b) // 201 0xc000018088

	//
	var i int
	var j, k = &i, &i
	fmt.Println(j, k) // 0xc0000180a8 0xc0000180a8
	// 不会有两个变量指向同一个内存地址
	fmt.Println(&j, &k) // 0xc000006030 0xc000006038

}
