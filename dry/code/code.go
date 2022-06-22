package code

import "reflect"

// 不是所有的代码重复都是"知识"的重复
// validateAge和validateQuantity看似一样的代码，
// 但表达了两个完全不同的概念：一个是年龄，一个是（订单）数量

func validateAge(value interface{}) {
	validateType(value, reflect.Int)
	validateMin(value, 0)
}

func validateQuantity(value interface{}) {
	validateType(value, reflect.Int)
	validateMin(value, 0)
}

func validateType(value interface{}, kind reflect.Kind) {

}

func validateMin(value interface{}, i int) {

}
