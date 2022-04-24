package interfaces

import (
	"fmt"
	"sort"
)

type Two []int

// Methods required by sort.Interface
func (s Two) Len() int {
	return len(s)
}
func (s Two) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Two) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Copy returns a copy
func (s Two) Copy() Two {
	copy := make(Two, 0, len(s))
	return append(copy, s...)
}

// Method for printing - sorts the elements before printing
func (s Two) String() string {
	s = s.Copy() // make a copy; don't overwrite argument
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}
