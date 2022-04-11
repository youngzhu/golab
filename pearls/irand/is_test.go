package irand

import (
	"fmt"
	"math/rand"
	"testing"
)

var n1, k1 = 10, 6

func floydTest(is IntSet, n, k int) {
	for j := n - k; j < n; j++ {
		t := rand.Int() % (j + 1)
		if ok := is.Insert(t); !ok {
			is.Insert(j)
		}
	}

	fmt.Println(is.Ints())
}

func TestIntSetArray(t *testing.T) {
	is := NewIntSetArray(k1, n1)
	floydTest(is, n1, k1)
}

func TestIntSetList(t *testing.T) {
	is := newIntSetList(k1, n1)
	floydTest(is, n1, k1)
}
