package irand

import (
	"math/rand"
	"testing"
)

var n, k = 100000, 10000

// 执行的结果，接不接没什么差别
func BenchmarkIntSetArray_noVar(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		OrderedUniqueInts(n, k)
	}
}

var Result []int

func BenchmarkIntSetArray_withVar(b *testing.B) {
	var r []int
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r = OrderedUniqueInts(n, k)
	}
	Result = r
}

var (
	k1000 = 1000
	k2000 = 2000
	k4000 = 4000
)

func floydBench(is IntSet, n, k int) {
	for j := n - k; j < n; j++ {
		t := rand.Int() % (j + 1)
		if ok := is.Insert(t); !ok {
			is.Insert(j)
		}
	}
}

var is IntSet

func BenchmarkIntSetArray_k1000(b *testing.B) {
	kk := k1000
	for i := 0; i < b.N; i++ {
		is = NewIntSetArray(kk, n)
		floydBench(is, n, kk)
	}
}

func BenchmarkIntSetArray_k2000(b *testing.B) {
	kk := k2000
	for i := 0; i < b.N; i++ {
		is = NewIntSetArray(kk, n)
		floydBench(is, n, kk)
	}
}

func BenchmarkIntSetArray_k4000(b *testing.B) {
	kk := k4000
	for i := 0; i < b.N; i++ {
		is = NewIntSetArray(kk, n)
		floydBench(is, n, kk)
	}
}

func BenchmarkIntSetList_k1000(b *testing.B) {
	kk := k1000
	for i := 0; i < b.N; i++ {
		is = newIntSetList(kk, n)
		floydBench(is, n, kk)
	}
}

func BenchmarkIntSetList_k2000(b *testing.B) {
	kk := k2000
	for i := 0; i < b.N; i++ {
		is = newIntSetList(kk, n)
		floydBench(is, n, kk)
	}
}

func BenchmarkIntSetList_k4000(b *testing.B) {
	kk := k4000
	for i := 0; i < b.N; i++ {
		is = newIntSetList(kk, n)
		floydBench(is, n, kk)
	}
}

// 递归的基准测试
// 这样分开测试，要逼近真实很麻烦，还不如改代码跑两遍
func Benchmark_insert_normal_1000(b *testing.B) {
	kk := k1000
	for i := 0; i < b.N; i++ {
		list := newIntSetList(kk, n)
		list.insertNormal(kk)
	}
}

func Benchmark_insert_recursion_1000(b *testing.B) {
	kk := k1000
	for i := 0; i < b.N; i++ {
		list := newIntSetList(kk, n)
		list.insertRecursion(kk)
	}
}

func Benchmark_insert_normal_2000(b *testing.B) {
	kk := k2000
	for i := 0; i < b.N; i++ {
		list := newIntSetList(kk, n)
		list.insertNormal(kk)
	}
}

func Benchmark_insert_recursion_2000(b *testing.B) {
	kk := k2000
	for i := 0; i < b.N; i++ {
		list := newIntSetList(kk, n)
		list.insertRecursion(kk)
	}
}

func Benchmark_insert_normal_4000(b *testing.B) {
	kk := k4000
	for i := 0; i < b.N; i++ {
		list := newIntSetList(kk, n)
		list.insertNormal(kk)
	}
}

func Benchmark_insert_recursion_4000(b *testing.B) {
	kk := k4000
	for i := 0; i < b.N; i++ {
		list := newIntSetList(kk, n)
		list.insertRecursion(kk)
	}
}
