package doc

// 文档中的重复

type account struct {
	returnedCheckCount int
	overdraftDays      int
	averageBalance     float64
}

/******* BEFORE *******/
// 计算这个账户的费用
// 每张退票算20
// 如果账户透支超过3天，每天的费用为 10
// 如果平均账户余额超过 2000，减免费用50%
func fee(a account) float64 {
	f := 0.0
	if a.returnedCheckCount > 0 {
		f += 20 * float64(a.returnedCheckCount)
	}
	if a.overdraftDays > 3 {
		f += 10 * float64(a.overdraftDays)
	}
	if a.averageBalance > 2_000 {
		f /= 2
	}

	return f
}

// 问题：一旦费用计算规则发生变化，备注和代码都要改

/******* AFTER *******/
// 去掉注释，算法也不复杂，没必要用文字再描述一遍
func calculateAccountFee(a account) float64 {
	f := 20.0 * float64(a.returnedCheckCount)
	if a.overdraftDays > 3 {
		f += 10 * float64(a.overdraftDays)
	}
	if a.averageBalance > 2_000 {
		f /= 2
	}
	return f
}
