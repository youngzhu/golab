package main

import "fmt"

// 如果要新增一个“除法”呢？ 增加一个常量，再修改Do的逻辑？
// 再新增一个呢？

const (
	OpAdd = 1 << iota
	OpSub
	OpMul
)

type Calculator struct {
	acc float64
}

func (c *Calculator) Do(op int, v float64) float64 {
	switch op {
	case OpAdd:
		c.acc += v
	case OpSub:
		c.acc -= v
	case OpMul:
		c.acc *= v
	default:
		panic("unhandled operation")
	}
	return c.acc
}

func main() {
	var c Calculator
	fmt.Println(c.Do(OpAdd, 100)) // 0+100=100
	fmt.Println(c.Do(OpSub, 80))  // 100-80=20
	fmt.Println(c.Do(OpMul, 2))   // 20*2=40
}
