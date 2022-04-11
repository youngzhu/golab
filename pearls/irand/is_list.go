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
	// 经测试，常规算法更高效
	//return s.insertRecursion(t)
	return s.insertNormal(t)
}

func (s *IntSetList) insertNormal(t int) bool {
	if s.head.val == t {
		return false
	}
	if s.head.val > t {
		s.head = newNode(t, s.head)
		s.n++
		return true
	}
	p := s.head
	for ; p.next.val < t; p = p.next {
	}
	if p.next.val == t {
		return false
	}
	p.next = newNode(t, p.next)
	s.n++
	return true
}

// 递归
func (s *IntSetList) insertRecursion(t int) bool {
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
