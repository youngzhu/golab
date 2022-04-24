package iprint

import "fmt"

// Sprintf 调用的是 类型的 String 方法
// 所以，下面的方法是错误的，导致无限循环

type MyString string

func (m MyString) String() string {
	// 编译时就有提示
	//return fmt.Sprintf("MyString=%s", m)

	return fmt.Sprintf("MyString=%s", string(m)) // 正确
}
