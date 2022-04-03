package str

import (
	"bytes"
	"testing"
)

func BenchmarkBytesEqualInline(b *testing.B) {
	x := bytes.Repeat([]byte{'a'}, 1<<10)
	y := bytes.Repeat([]byte{'a'}, 1<<10)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if string(x) != string(y) {
			b.Fatalf("x!=y")
		}
	}
}
func BenchmarkBytesEqualExplicit(b *testing.B) {
	x := bytes.Repeat([]byte{'a'}, 1<<10)
	y := bytes.Repeat([]byte{'a'}, 1<<10)
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// 内联这么大威力啊，效率上差一个数量级
		p := string(x)
		q := string(y)
		if p != q {
			b.Fatalf("x!=y")
		}
	}
}
