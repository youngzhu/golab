package insert

func Sort1(s []int) {
	n := len(s)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}

func Sort2(s []int) {
	n := len(s)
	for i := 1; i < n; i++ {
		t := s[i]
		j := i
		for ; j > 0 && s[j-1] > t; j-- {
			// 这里少了赋值
			// 无序数组差别比较大，快了近4倍
			s[j] = s[j-1]
		}
		s[j] = t
	}
}
