package main

import (
	"fmt"
	"math"
)

//

type Calculator struct {
	acc float64
}

type opFunc func(n float64) float64

func Add(n float64) func(float64) float64 {
	return func(acc float64) float64 {
		return acc + n
	}
}
func Sqrt() func(float64) float64 {
	return func(n float64) float64 {
		return math.Sqrt(n)
	}
}

func (c *Calculator) Do(op opFunc) float64 {
	c.acc = op(c.acc)
	return c.acc
}

func main() {
	var c Calculator
	fmt.Println(c.Do(Add(10))) // 0+10=10
	fmt.Println(c.Do(Add(20))) // 10+20=30

	var c2 Calculator
	fmt.Println(c2.Do(Add(2)))
	//fmt.Println(c2.Do(Sqrt()))
	// math.Sqrt 刚好满足 opFunc
	fmt.Println(c2.Do(math.Sqrt))

}
