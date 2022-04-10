package irand

import "testing"

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
