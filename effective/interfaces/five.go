package interfaces

import (
	"fmt"
	"sort"
)

// One 没有Copy方法

type Five []int

// Methods required by sort.Interface
func (s Five) Len() int {
	return len(s)
}
func (s Five) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Five) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Method for printing - sorts the elements before printing
func (s Five) String() string {
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
