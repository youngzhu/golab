package interfaces

import (
	"fmt"
)

var (
	ints = []int{3, 3, 1, 2, 6}

	one   = One(ints)
	two   = Two(ints)
	three = Three(ints)
	four  = Four(ints)
	five  = Five(ints)
)

func ExampleOne() {
	fmt.Println(one) // 自动调了String方法，所以排序了
	fmt.Println(one.String())
	fmt.Println([]int(one))

	// output:
	//[1 2 3 3 6]
	//[1 2 3 3 6]
	//[3 3 1 2 6]
}

func ExampleTwo() {
	fmt.Println(two) // 自动调了String方法，所以排序了
	fmt.Println(two.String())
	fmt.Println([]int(two))

	// output:
	//[1 2 3 3 6]
	//[1 2 3 3 6]
	//[3 3 1 2 6]
}

func ExampleThree() {
	fmt.Println(three) // 自动调了String方法，所以排序了
	fmt.Println(three.String())
	fmt.Println([]int(three))

	// output:
	//[1 2 3 3 6]
	//[1 2 3 3 6]
	//[3 3 1 2 6]
}

func ExampleFour() {
	// 没有String方法，两个输出是一样的
	fmt.Println(four)
	//fmt.Println(four.String())
	fmt.Println([]int(four))

	// output:
	//[3 3 1 2 6]
	//[3 3 1 2 6]
}

func ExampleFive() {
	// 没有Copy，caller也被改变了
	// 所以输出的都是有序的
	fmt.Println(five) // 自动调了String方法，所以排序了
	fmt.Println(five.String())
	fmt.Println([]int(five))

	// output:
	//[1 2 3 3 6]
	//[1 2 3 3 6]
	//[1 2 3 3 6]
}
