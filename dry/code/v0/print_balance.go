package v0

import (
	"fmt"
	"github.com/youngzhu/golab/dry/code/account"
	"strings"
)

// v0
// 问题：对正负数的处理有明显的重复

func PrintBalance(account account.Account) string {
	var sb strings.Builder
	fmt.Fprintf(&sb, "Debits: %10.2f\n", account.Debits)
	fmt.Fprintf(&sb, "Credits: %10.2f\n", account.Credits)
	if account.Fees < 0 {
		fmt.Fprintf(&sb, "Fees: -%10.2f\n", -account.Fees)
	} else {
		fmt.Fprintf(&sb, "Fees: %10.2f\n", account.Fees)
	}
	fmt.Fprintln(&sb, "-------")
	if account.Balance < 0 {
		fmt.Fprintf(&sb, "Balance: %10.2f-\n", -account.Balance)
	} else {
		fmt.Fprintf(&sb, "Balance: %10.2f\n", account.Balance)
	}
	return sb.String()
}
