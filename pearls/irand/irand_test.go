package irand

import (
	"fmt"
	"testing"
)

func TestIntsn(t *testing.T) {
	s := Intsn(100, 10)
	for _, i := range s {
		println(i)
	}
}

func TestInts(t *testing.T) {
	s := Ints(10)
	for _, i := range s {
		println(i)
	}
}

func TestRangeInt(t *testing.T) {
	println(RangeInt(10, 11))
	println(RangeInt(80, 100))
	println(RangeInt(100, 1000))
}

func TestUniqueInts(t *testing.T) {
	s := UniqueInts(10, 5)
	for i := range s {
		println(s[i])
	}
}

func TestOrderedUniqueInts(t *testing.T) {
	s := OrderedUniqueInts(10, 8)
	fmt.Println(s)
}

func TestShuffle(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	Shuffle(a, .2)
	t.Log(a)
}
