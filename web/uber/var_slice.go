package uber

// 用 var 申明的切片可以立即使用，无需调用 make 创建

func bad() {
	nums := []int{}
	// or
	// nums := make([]int)

	nums = append(nums, 1)
}

func good() {
	var nums []int

	nums = append(nums, 1)
}