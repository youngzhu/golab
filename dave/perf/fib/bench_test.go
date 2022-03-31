package fib

import "testing"

// go test -bench="." ./fib/
// "" 不能省（Windows），不确定是否跟环境有关

// go test -bench="." -cpu="1,2,4" ./fib/

// go test -bench="." -benchtime=10s ./fib/

// go test -bench="." -count=10 ./fib/|tee old.txt
// tee 将结果写入文件

// go test -bench="." -benchtime=500x ./fib/
// 指定次数 500

func BenchmarkFib20(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(20)
	}
}

//func BenchmarkFib28(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		Fib(28)
//	}
//}

func BenchmarkExpensive(b *testing.B) {
	boringAndExpensiveSetup()
	b.ResetTimer() // 初始化之后重置
	for i := 0; i < b.N; i++ {

	}
}

func boringAndExpensiveSetup() {

}

func BenchmarkComplicated(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		complicatedSetup() // 每次循环都需要的设置
		b.StartTimer()

		// 测试主体
	}
}

func complicatedSetup() {

}

func BenchmarkRead(b *testing.B) {
	// 显示每次操作的 内存分配
	// 同 go test -bench=. -benchmem
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {

	}
}

func BenchmarkWrong(b *testing.B) {
	Fib(b.N)
}

func BenchmarkWrong2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(i)
	}
}
