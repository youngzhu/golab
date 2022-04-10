package irand

type IntSet interface {
	Insert(i int) bool // true表示插入成功
	Ints() []int
}
