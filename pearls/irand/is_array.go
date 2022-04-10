package irand

// IntSetArray IntSet的线性（array）实现
type IntSetArray struct {
	n int
	x []int
}

// NewIntSetArray
// k 个值
// n 最大值（不包含）
func NewIntSetArray(k, n int) *IntSetArray {
	if k > n {
		panic("k>n")
	}
	x := make([]int, k+1) // 留一个哨兵位
	x[0] = n              // 哨兵
	return &IntSetArray{x: x}
}

func (s *IntSetArray) Insert(t int) bool {
	i := 0
	for ; s.x[i] < t; i++ {
	}
	if s.x[i] == t {
		return false
	}
	// 将i右边的值（包括i）向右移一位
	for j := s.n; j >= i; j-- {
		s.x[j+1] = s.x[j]
	}
	s.x[i] = t
	s.n++
	return true
}

func (s *IntSetArray) Ints() []int {
	return s.x[0:s.n]
}
