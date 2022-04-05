package irand

import "testing"

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
