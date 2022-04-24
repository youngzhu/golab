package interfaces

import (
	"fmt"
	"sort"
)

type One []int

// Methods required by sort.Interface
func (s One) Len() int {
	return len(s)
}
func (s One) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s One) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Copy returns a copy
func (s One) Copy() One {
	copy := make(One, 0, len(s))
	return append(copy, s...)
}

// Method for printing - sorts the elements before printing
func (s One) String() string {
	s = s.Copy() // make a copy; don't overwrite argument
	sort.Sort(s)
	str := "["
	for i, elem := range s { // O(N^2)
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}
