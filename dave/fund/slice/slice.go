package slice

// 数组，永远都是copy

// 切片
// 1. 引用。即使是从切片（a）分出的切片（b），也是指向同一个内存区域（slicing.go)
// 2. 在函数中是引用传递（passing.go）
// 3. 只要重新赋值（=）就使用了副本（copy），（assign.go)
