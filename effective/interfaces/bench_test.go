package interfaces

import "testing"

// Two 效果最好
// 虽然Three省去了sort.Interface

var (
	benchInts = []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}

	benchOne   = One(benchInts)
	benchTwo   = Two(benchInts)
	benchThree = Three(benchInts)
)

func BenchmarkOne_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchOne.String()
	}
}

func BenchmarkTwo_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchTwo.String()
	}
}

func BenchmarkThree_String(b *testing.B) {
	for i := 0; i < b.N; i++ {
		benchThree.String()
	}
}
