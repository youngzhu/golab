package insert

import (
	"github.com/youngzhu/golab/pearls/isort"
	"testing"
)

func BenchmarkSort1_rand(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := isort.Get1KRand() // 保证每次都是未排序且是一样的
		b.StartTimer()

		Sort1(s)
	}
}
func BenchmarkSort1_sorted(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := isort.Get1kSorted()
		b.StartTimer()

		Sort1(s)
	}
}
func BenchmarkSort1_almostSorted(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := isort.Get1kAlmostSorted()
		b.StartTimer()

		Sort1(s)
	}
}

func BenchmarkSort2_rand(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := isort.Get1KRand() // 保证每次都是未排序且是一样的
		b.StartTimer()

		Sort2(s)
	}
}
func BenchmarkSort2_sorted(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := isort.Get1kSorted()
		b.StartTimer()

		Sort2(s)
	}
}
func BenchmarkSort2_almostSorted(b *testing.B) {
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		s := isort.Get1kAlmostSorted()
		b.StartTimer()

		Sort2(s)
	}
}
