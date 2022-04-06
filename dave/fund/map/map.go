package main

func f(m map[int]int) {
	m = make(map[int]int)
}

func main() {
	var m map[int]int
	f(m)
	// 跟slice一样，赋值后就是副本
	// 如果是引用传递，这里就应该是 false
	println(m == nil) // true
}
