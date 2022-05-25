package json_test

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Message struct {
	Name string
	Body string
	Time int
}

func ExampleEncoding() {
	m := Message{"Alice", "Hello", 1294706395881547000}
	b, _ := json.Marshal(m)
	bb := []byte(`{"Name":"Alice","Body":"Hello","Time":1294706395881547000}`)
	fmt.Println(reflect.DeepEqual(b, bb))
	fmt.Println(string(b))
	fmt.Println(string(bb))

	// Output:
	// true
}
