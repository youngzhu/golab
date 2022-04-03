package mutex

import "testing"

func BenchmarkAtomicVariable_Inc(b *testing.B) {
	var av AtomicVariable

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			av.Inc()
		}
	})
}

func BenchmarkAtomicCounter_Inc(b *testing.B) {
	var ac AtomicCounter

	b.ReportAllocs()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			ac.Inc()
		}
	})
}
