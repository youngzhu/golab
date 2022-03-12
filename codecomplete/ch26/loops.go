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
