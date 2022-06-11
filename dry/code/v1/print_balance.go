package v1

import (
	"fmt"
	"github.com/youngzhu/golab/dry/code/account"
	"math"
	"strings"
)

// v0
// 问题：对正负数的处理有明显的重复
// v1
// 改进：增加一个格式化输出金额的函数
// 问题：如果需要改变输出格式怎么办？例如在标签和金额之间多加一个空格

func formatAmount(amount float64) string {
	result := fmt.Sprintf("%10.2f", math.Abs(amount))
	if amount < 0 {
		return "-" + result
	} else {
		return result
	}
}

func PrintBalance(account account.Account) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Debits: %s\n", formatAmount(account.Debits))
	fmt.Fprintf(&sb, "Credits: %s\n", formatAmount(account.Credits))
	fmt.Fprintf(&sb, "Fees: %s\n", formatAmount(account.Fees))
	fmt.Fprintln(&sb, "-------")
	fmt.Fprintf(&sb, "Balance: %s\n", formatAmount(account.Balance))
	return sb.String()
}
