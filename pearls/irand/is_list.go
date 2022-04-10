package irand

type node struct {
	val  int
	next *node
}

func newNode(val int, next *node) *node {
	return &node{val: val, next: next}
}

type IntSetList struct {
	n              int
	head, sentinel *node
}

// k 个值
// n 最大值（不包含）
func newIntSetList(k, n int) *IntSetList {
	sentinel := newNode(n, nil)
	return &IntSetList{head: sentinel, sentinel: sentinel}
}

func (s *IntSetList) Insert(t int) bool {
	nn := s.n
	s.head = s.insert(s.head, t)
	return nn < s.n
}

func (s *IntSetList) insert(p *node, t int) *node {
	if p.val < t {
		p.next = s.insert(p.next, t)
	} else if p.val > t {
		p = newNode(t, p)
		s.n++
	}
	return p
}

func (s *IntSetList) Ints() []int {
	ints := make([]int, s.n)
	j := 0
	for p := s.head; p != s.sentinel; p = p.next {
		ints[j] = p.val
		j++
	}
	return ints
}
