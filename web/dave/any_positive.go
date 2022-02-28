package dave

func doSomething() {
	var num1, num2, num3, num4 int

	if num1 > 0 || num2 > 0 || num3 > 0 || num4 > 0 {
		// 任意一个数值大于0
	}
}

// 重构
func anyPositive(values ...int) bool {
	for _, v := range values {
		if v > 0 {
			return true
		}
	}
	return false
}

func doSomething2() {
	var num1, num2, num3, num4 int

	if anyPositive(num1, num2, num3, num4) {
		// 任意一个数值大于0
	}
}

// 但这存在一个问题，可能被这样调用：anyPositive()
// 即，不给任何参数

// 以下方法要求至少一个参数

// anyPositive indicates if any value is greater than zero.
func anyPositiveX(first int, rest ...int) bool {
	if first > 0 {
		return true
	}
	for _, v := range rest {
		if v > 0 {
			return true
		}
	}
	return false
}
