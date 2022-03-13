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

func BusiestOutside(table [][]int) (sum int) {
	rowCount, columnCount := len(table), len(table[0])

	for column := 0; column < columnCount; column++ {
		for row := 0; row < rowCount; row++ {
			sum += table[row][column]
		}
	}

	return
}

func BusiestInside(table [][]int) (sum int) {
	rowCount, columnCount := len(table), len(table[0])

	for row := 0; row < rowCount; row++ {
		for column := 0; column < columnCount; column++ {
			sum += table[row][column]
		}
	}

	return
}

// CalCommission 计算佣金
// 直观。但是循环中使用了不必要的乘法
func CalCommission(saleCount int) (commission []float64) {
	commission = make([]float64, saleCount)

	revenue := 10000.00
	baseCommission := 3000.0
	discount := 0.85

	for i := 0; i < saleCount; i++ {
		count := float64(i + 1)
		commission[i] = count * revenue * baseCommission * discount
	}

	return
}

func CalCommissionTuned(saleCount int) (commission []float64) {
	commission = make([]float64, saleCount)
	
	revenue := 10000.00
	baseCommission := 3000.0
	discount := 0.85

	incrementalCommission := revenue * baseCommission * discount
	cumulativeCommission := incrementalCommission

	for i := 0; i < saleCount; i++ {
		commission[i] = cumulativeCommission
		cumulativeCommission += incrementalCommission
	}

	return
}
