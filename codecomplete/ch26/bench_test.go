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

var int100 []int

func initTestData() {
	for i := 0; i < 100; i++ {
		int100 = append(int100, i*i)
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
