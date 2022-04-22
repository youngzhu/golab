package iswitch

import "fmt"

func main() {
	var t interface{}
	t = funcOfSomeType()

	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t)
	case bool:
		fmt.Printf("boolean %t\n", t)
	case int:
		fmt.Printf("integer %d\n", t)
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t)
	case *int:
		fmt.Printf("pointer to integer %d\n", *t)
	}
}

func funcOfSomeType() interface{} {
	return nil
}
