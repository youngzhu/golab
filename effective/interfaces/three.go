package interfaces

import (
	"fmt"
	"sort"
)

type Three []int

// Copy returns a copy
func (s Three) Copy() Three {
	copy := make(Three, 0, len(s))
	return append(copy, s...)
}

// Method for printing - sorts the elements before printing
func (s Three) String() string {
	s = s.Copy() // make a copy; don't overwrite argument
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}
