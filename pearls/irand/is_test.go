package irand

import (
	"fmt"
	"math/rand"
	"testing"
)

var n1, k1 = 10, 6

func floyd(is IntSet) {
	for j := n1 - k1; j < n1; j++ {
		t := rand.Int() % (j + 1)
		if ok := is.Insert(t); !ok {
			is.Insert(j)
		}
	}

	fmt.Println(is.Ints())
}

func TestIntSetArray(t *testing.T) {
	is := NewIntSetArray(k1, n1)
	floyd(is)
}

func TestIntSetList(t *testing.T) {
	is := newIntSetList(k1, n1)
	floyd(is)
}
