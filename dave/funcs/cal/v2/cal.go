package main

import (
	"fmt"
	"math"
)

// Sqrt有点奇怪

type Calculator struct {
	acc float64
}

type opFunc func(op1, op2 float64) float64

func Add(a, b float64) float64 {
	return a + b
}

func Sub(a, b float64) float64 {
	return a - b
}

func Mul(a, b float64) float64 {
	return a * b
}

func Sqrt(n, _ float64) float64 {
	return math.Sqrt(n)
}

func (c *Calculator) Do(op opFunc, v float64) float64 {
	c.acc = op(c.acc, v)
	return c.acc
}

func main() {
	var c Calculator
	fmt.Println(c.Do(Add, 16)) // 0+16=16
	// 0 是被忽略的，而且显得奇怪
	fmt.Println(c.Do(Sqrt, 0)) // Sqrt(16)=4
}
