package ch26

const (
	SumTypeEven = iota
	SumTypeOdd
)

func Switched(n int) int {
	sum := 0

	sumType := SumTypeEven
	for i := 0; i < n; i++ {
		// 这个判断条件在循环中不会改变
		if sumType == SumTypeEven {
			if i%2 == 0 {
				sum += i
			}
		} else {
			if i%2 == 1 {
				sum += i
			}
		}
	}

	return sum
}

func Unswitched(n int) int {
	sum := 0

	sumType := SumTypeEven

	if sumType == SumTypeEven {
		for i := 0; i < n; i++ {
			if i%2 == 0 {
				sum += i
			}
		}
	} else {
		for i := 0; i < n; i++ {
			if i%2 == 1 {
				sum += i
			}
		}
	}

	return sum
}

// 哨兵

func NoSentinel(items []int, target int) bool {
	n := len(items)
	found := false
	i := 0
	for !found && i < n {
		if items[i] == target {
			found = true
		} else {
			i++
		}
	}
	return found
}

func WithSentinel(items []int, target int) bool {
	n := len(items)
	last := items[n-1]
	if last == target {
		return true
	}
	items[n-1] = target // 哨兵，省去边界值的判断
	i := 0
	for items[i] != target {
		i++
	}
	return i < n-2
}
