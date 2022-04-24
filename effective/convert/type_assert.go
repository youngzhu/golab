package main

import "fmt"

var value interface{}

func main() {
	//value = "hello"
	value = 9.99
	str, ok := value.(string)
	if ok {
		fmt.Printf("string value is: %q", str)
	} else {
		// 转换失败了，str仍然有效，为对象的空值
		fmt.Printf("value is not a string: %q", str)
	}
}
