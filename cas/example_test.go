package cas_test

import (
	"fmt"
	"github.com/youngzhu/golab/cas"
	"time"
)

func ExampleCounter_Increment() {
	counter := cas.NewCounter()

	go c1(counter)

	go c2(counter)

	for counter.Get() < 100 {
		//fmt.Println(counter.Get())
	}

	fmt.Println("done")

	// Output:
	// ...
}

func c1(counter *cas.Counter) {
	for i := 0; i < 50; i++ {
		val := counter.Increment()
		fmt.Println("c1:", val)
		time.Sleep(200 * time.Millisecond)
	}
}

func c2(counter *cas.Counter) {
	for i := 0; i < 50; i++ {
		val := counter.Increment()
		fmt.Println("c2:", val)
		time.Sleep(100 * time.Millisecond)
	}
}
