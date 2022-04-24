package interfaces

// One 没有 String 方法

type Four []int

// Methods required by sort.Interface
func (s Four) Len() int {
	return len(s)
}
func (s Four) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Four) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Copy returns a copy
func (s Four) Copy() Four {
	copy := make(Four, 0, len(s))
	return append(copy, s...)
}
