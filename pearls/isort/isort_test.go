package isort

import (
	"github.com/youngzhu/golab/pearls/isort/insert"
	"reflect"
	"sort"
	"testing"
)

// TODO 不是说都是副本吗？为什么排序后改变了原始数据呢？

//var array1 = [5]int{3, 2, 4, 1, 5}
var array1 = [...]int{3, 2, 4, 1, 5}
var array2 = [...]int{3, 2, 4, 1, 5}

var slice1 = []int{3, 2, 4, 1, 5}
var slice2 = []int{3, 2, 4, 1, 5}

func TestArray_1(t *testing.T) {
	t.Log("array1 before:", array1)
	sort.Ints(array1[0:])
	t.Log("array1 after:", array1)
}

func TestArray_2(t *testing.T) {
	t.Log("array1 before:", array1)
	sort.Ints(array1[0:])
	t.Log("array1 after:", array1)
}

func TestArray_3(t *testing.T) {
	t.Log("array2 before:", array2)
	a := array2[0:]
	sort.Ints(a)
	t.Log("array2 after:", array2)
}
func TestArray_4(t *testing.T) {
	t.Log("array2 before:", array2)
	a := array2[0:]
	sort.Ints(a)
	t.Log("array2 after:", array2)
}

func TestSlice_1(t *testing.T) {
	t.Log("slice1 before:", slice1)
	sort.Ints(slice1)
	t.Log("slice1 after:", slice1)
}
func TestSlice_2(t *testing.T) {
	t.Log("slice1 before:", slice1)
	sort.Ints(slice1)
	t.Log("slice1 after:", slice1)
}

func TestSlice_3(t *testing.T) {
	t.Log("slice2 before:", slice2)
	s := slice2
	sort.Ints(s)
	t.Log("slice2 after:", slice2)
}
func TestSlice_4(t *testing.T) {
	t.Log("slice2 before:", slice2)
	s := slice2
	sort.Ints(s)
	t.Log("slice2 after:", slice2)
}

var tests = []struct {
	before, after []int
}{
	{[]int{1, 4, 5, 3, 3, 1}, []int{1, 1, 3, 3, 4, 5}},
}

func TestInsertSort1(t *testing.T) {
	for _, tc := range tests {
		before := tc.before
		t.Log("before:", before)
		insert.Sort1(before)
		t.Log("after:", before)
		if !reflect.DeepEqual(before, tc.after) {
			t.Errorf("sort fail, want:%v, got:%v", tc.after, before)
		}
	}
}

func TestInsertSort2(t *testing.T) {
	for _, tc := range tests {
		before := tc.before
		t.Log("before:", before)
		insert.Sort2(before)
		t.Log("after:", before)
		if !reflect.DeepEqual(before, tc.after) {
			t.Errorf("sort fail, want:%v, got:%v", tc.after, before)
		}
	}
}
