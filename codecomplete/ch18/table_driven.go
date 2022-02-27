package ch18

// 表驱动编程/表驱动开发
// 所谓表驱动开发，就是把复杂的逻辑转化成查表（数组/map等）的操作

/*
获得某个月的天数（为了方便讨论，忽略闰年）
*/

// DaysOfMonthGeneral 获得某个月的天数（常规方法）
func DaysOfMonthGeneral(month int) int {
	switch month {
	case 1:
		return 31
	case 2:
		return 28
	case 3:
		return 31
	case 4:
		return 30
	case 5:
		return 31
	case 6:
		return 30
	case 7:
		return 31
	case 8:
		return 31
	case 9:
		return 30
	case 10:
		return 31
	case 11:
		return 30
	case 12:
		return 31
	default:
		panic("invalid month")
	}
}

var daysPerMonth = []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}

// DaysOfMonth 获得某个月的天数（表驱动法）
func DaysOfMonth(month int) int {
	return daysPerMonth[month-1]
}

/*
评分系统
>=90.0	A
<90.0	B
<75.0	C
<65.0	D
<50.0	F
*/

// GetGradeGeneral 常见方法
func GetGradeGeneral(score float32) string {
	if score >= 90.0 {
		return "A"
	} else if score >= 75.0 {
		return "B"
	} else if score >= 65.0 {
		return "C"
	} else if score >= 50.0 {
		return "D"
	} else {
		return "F"
	}
}

var (
	rangeUpper = []float32{50.0, 65.0, 75.0, 90.0, 100.0}
	grades     = []string{"F", "D", "C", "B", "A"}
)

// GetGrade 表驱动法
func GetGrade(score float32) string {
	maxGradeLevel := len(grades) - 1

	gradeLevel := 0
	studentGrade := "A"
	// 数据量大时可考虑二分查找
	for studentGrade == "A" && gradeLevel < maxGradeLevel {
		if score < rangeUpper[gradeLevel] {
			studentGrade = grades[gradeLevel]
		}
		gradeLevel++
	}
	return studentGrade
}
