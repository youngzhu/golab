package code

import (
	"github.com/youngzhu/golab/dry/code/account"
	v0 "github.com/youngzhu/golab/dry/code/v0"
	v1 "github.com/youngzhu/golab/dry/code/v1"
	v2 "github.com/youngzhu/golab/dry/code/v2"
	"testing"
)

var a = account.Account{
	Debits:  1_000_000.00,
	Credits: 1_000_000_000.00,
	Fees:    -1500,
	Balance: 9999999.99,
}

var want = `Debits: 1000000.00
Credits: 1000000000.00
Fees: -   1500.00
-------
Balance: 9999999.99
`

func testPrintBalance(t *testing.T, f func(account2 account.Account) string) {
	got := f(a)
	if got != want {
		t.Errorf("got:\n%v", got)
		t.Errorf("want:\n%v", want)
	}
}

func TestPrintBalance_v0(t *testing.T) {
	testPrintBalance(t, v0.PrintBalance)
}

func TestPrintBalance_v1(t *testing.T) {
	testPrintBalance(t, v1.PrintBalance)
}

func TestPrintBalance_v2(t *testing.T) {
	testPrintBalance(t, v2.PrintBalance)
}
