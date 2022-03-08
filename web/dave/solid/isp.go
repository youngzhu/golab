package solid

import (
	"io"
	"os"
)

// Interface Segregation Principle: 接口隔离原则

type Document struct {
}

func Save(f *os.File, doc *Document) error {
	return nil
}

// 上述方法的问题：
// 1. 排除了将数据写入网络的选项
// 2. f 包含了很多和 Save 无关的方法

// 更合适的定义

func SaveISP(w io.Writer, doc *Document) error {
	return nil
}
