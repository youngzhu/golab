package irand

import (
	"math/rand"
)

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

// RangeInt 返回[a,b)之间的随机数
func RangeInt(a, b int) int {
	if b <= a {
		panic("b<=a")
	}

	return a + rand.Intn(b-a)
}

// UniqueInts 返回k个不重复的[1,n]之间的随机数
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

// OrderedUniqueInts 返回[0,n)范围内 k 个有序的随机数
// Floyd算法
func OrderedUniqueInts(n, k int) []int {
	set := NewIntSetArray(k, n)

	for j := n - k; j < n; j++ {
		t := rand.Int() % (j + 1)
		if ok := set.Insert(t); !ok {
			set.Insert(j)
		}
	}

	return set.Ints()
}

func Shuffle(s []int, ratio float64) {
	n := int(float64(len(s)) * ratio)
	if n < 3 {
		panic("ratio too small")
	}
	rand.Shuffle(n, func(i, j int) {
		//log.Println("n:", n, "i:", i, "j:", j)
		s[i], s[j] = s[j], s[i]
	})
}
