package isort

import (
	"github.com/youngzhu/golab/pearls/irand"
	"math/rand"
)

const n1k = 1000

func Get1KRand() []int {
	return irand.Intsn(n1k, n1k)
}

func Get1kSorted() (s []int) {
	s = make([]int, n1k)
	for i := range s {
		s[i] = i
	}
	return
}

func Get1kAlmostSorted() []int {
	s := Get1kSorted()

	rand.Shuffle(n1k*.05, func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})

	return s
}
