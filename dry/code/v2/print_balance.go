package v2

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

func printLine(label, value string) string {
	return fmt.Sprintf("%-9s%s\n", label, value)
}

func reportLine(label string, amount float64) string {
	return printLine(label+":", formatAmount(amount))
}

func formatAmount(amount float64) string {
	result := fmt.Sprintf("%10.2f", math.Abs(amount))
	if amount < 0 {
		return "-" + result
	} else {
		return result
	}
}

func PrintBalance(a account.Account) string {
	var sb strings.Builder
	sb.WriteString(reportLine("Debits", a.Debits))
	sb.WriteString(reportLine("Credits", a.Credits))
	sb.WriteString(reportLine("Fees", a.Fees))
	sb.WriteString(printLine("", "-------"))
	sb.WriteString(reportLine("Balance", a.Balance))
	return sb.String()
}
