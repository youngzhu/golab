package solid

import "fmt"

// Open/Closed Principle

type A struct {
	year int
}

func (a A) Greet() {
	fmt.Println("Hello Golang", a.year)
}

type B struct {
	A
}

func (b B) Greet() {
	fmt.Println("Welcome to Golang", b.year)
}

// 嵌套，对扩展开放
func foo() {
	var a A
	a.year = 2016
	var b B
	b.year = 2022
	a.Greet() // Hello Golang 2016
	b.Greet() // Welcome to Golang 2022
}

//

type Cat struct {
	Name string
}

func (c Cat) Legs() int {
	return 4
}

func (c Cat) PrintLegs() {
	fmt.Printf("I have %d legs.", c.Legs())
}

type OctoCat struct {
	Cat
}

func (o OctoCat) Legs() int {
	return 5
}

func bar() {
	var octo OctoCat
	fmt.Println(octo.Legs()) // 5
	// 对修改封闭
	// PrintLegs的接收器是Cat
	// Go不支持重载，OctoCat不能代替普通的Cat，所以仍然是Cat.Legs，结果是4
	octo.PrintLegs() // I have 4 legs
}

// 接收器：Go语言中的语法糖
// 以下方法跟上面带接收器的方法比较，接收器就是函数的第一个参数

func PrintLegs(c Cat) {
	fmt.Printf("I have %d legs.", c.Legs())
}
