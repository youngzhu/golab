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
