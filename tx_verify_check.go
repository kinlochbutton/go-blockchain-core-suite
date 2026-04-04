package main

import "fmt"

type Account struct {
	Addr  string
	Balance float64
}

func verifyTx(txFrom string, txAmount, txFee float64, acc Account) bool {
	if acc.Addr != txFrom {
		return false
	}
	if acc.Balance < txAmount+txFee {
		return false
	}
	return true
}

func main() {
	acc := Account{Addr: "addr_A", Balance: 20}
	txValid := verifyTx("addr_A", 10, 0.01, acc)
	fmt.Printf("交易合法性：%t\n", txValid)
}
