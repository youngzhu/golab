package irand

import "math/rand"

// Intsn 生成count个[0,n)的随机数
func Intsn(n, count int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = rand.Intn(n)
	}
	return s
}

func Ints(count int) []int {
	s := make([]int, count)
	for i := range s {
		s[i] = rand.Int()
	}
	return s
}

func RangeInt(a, b int) int {
	if b <= a {
		panic("b<=a")
	}

	return a + rand.Intn(b-a)
}

func UniqueInts(n, k int) []int {
	if k > n {
		panic("k>n")
	}

	sn := make([]int, n)

	for i := range sn {
		sn[i] = i + 1
	}

	for i := 0; i < k; i++ {
		r := RangeInt(i, n)
		sn[i], sn[r] = sn[r], sn[i]
	}

	return sn[0:k]
}
