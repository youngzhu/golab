package concat

import (
	"testing"
)

func BenchmarkPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Plus()
	}
}

func BenchmarkBuffer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Buffer()
	}
}
func BenchmarkSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sprintf()
	}
}
func BenchmarkBytes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Bytes()
	}
}
func BenchmarkBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Builder()
	}
}
