package json_test

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// JSON只能处理公共的字段（首字母大写）
// 结构（struct）可以小写
type message struct {
	Name string
	Body string
	Time int
}

func ExampleEncoding() {
	m := message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m)
	bb := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	//fmt.Println(string(b))
	//fmt.Println(string(bb))
	fmt.Println(reflect.DeepEqual(b, bb))

	// Output:
	// true
}

func ExampleDecoding() {
	var m message
	b := []byte(`{"Name":"Bob","Food":"Pickle"}`)
	err := json.Unmarshal(b, &m)
	fmt.Println(err == nil)
	fmt.Printf("%v", m)

	// Output:
	// true
	// {Bob  0}
}

// 对任意JSON串的解析
func ExampleDecoding_arbitraryData() {
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var a interface{}
	json.Unmarshal(b, &a)
	m := a.(map[string]interface{})
	for k, v := range m {
		switch vv := v.(type) {
		case string:
			fmt.Println(k, "is string:", vv)
		case float64:
			fmt.Println(k, "is float64:", vv)
		case []interface{}:
			fmt.Println(k, "is an array:")
			for i, u := range vv {
				fmt.Println(i, u)
			}
		default:
			fmt.Println(k, "is of a type I don't know how to handle")
		}
	}

	// random order
	// Output:
	// Name is string: Wednesday
	// Age is float64: 6
	// Parents is an array:
	// 0 Gomez
	// 1 Morticia
}
