package ch26

import (
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
