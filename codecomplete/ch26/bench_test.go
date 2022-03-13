package ch26

import (
	"math/rand"
	"testing"
)

// go test -v -run="none" -bench="."

func BenchmarkSwitched(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Switched(100)
	}
}

func BenchmarkUnswitched(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Unswitched(100)
	}
}

var (
	int100 []int
	table  [][]int
)

func initTestData() {
	for i := 0; i < 100; i++ {
		int100 = append(int100, i*i)
	}

	rows, columns := 5, 100
	table = make([][]int, rows)
	for i := range table {
		table[i] = make([]int, columns)
	}
}

func BenchmarkNoSentinel(b *testing.B) {
	b.ReportAllocs()
	if int100 == nil {
		b.StopTimer()
		initTestData()
		b.StartTimer()
	}
	for i := 0; i < b.N; i++ {
		NoSentinel(int100, i)
	}
}

func BenchmarkWithSentinel(b *testing.B) {
	b.ReportAllocs()
	if int100 == nil {
		b.StopTimer()
		initTestData()
		b.StartTimer()
	}
	for i := 0; i < b.N; i++ {
		WithSentinel(int100, i)
	}
}

func BenchmarkBusiestOutside(b *testing.B) {
	b.ReportAllocs()
	if table == nil {
		b.StopTimer()
		initTestData()
		b.StartTimer()
	}
	for i := 0; i < b.N; i++ {
		BusiestOutside(table)
	}
}

func BenchmarkBusiestInside(b *testing.B) {
	b.ReportAllocs()
	if table == nil {
		b.StopTimer()
		initTestData()
		b.StartTimer()
	}
	for i := 0; i < b.N; i++ {
		BusiestInside(table)
	}
}

func BenchmarkCalCommission(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CalCommission(20)
	}
}

func BenchmarkCalCommissionTuned(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CalCommissionTuned(20)
	}
}

func BenchmarkHypotenuse_rand(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Hypotenuse(rand.NormFloat64(), rand.NormFloat64())
	}
}

func BenchmarkHypotenuseWithCache_rand(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		HypotenuseWithCache(rand.NormFloat64(), rand.NormFloat64())
	}
}

func BenchmarkHypotenuse_same(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Hypotenuse(3.0, 4.0)
	}
}

func BenchmarkHypotenuseWithCache_same(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		HypotenuseWithCache(3.0, 4.0)
	}
}

var sides = [][]float64{{3, 0, 4.0}, {8.0, 6.0}}

func BenchmarkHypotenuse_nocache(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		side := sides[i%2]
		Hypotenuse(side[0], side[1])
	}
}

func BenchmarkHypotenuseWithCache_nocache(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		side := sides[i%2]
		HypotenuseWithCache(side[0], side[1])
	}
}
