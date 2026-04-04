package main

import "fmt"

type Account struct {
	Addr    string
	Balance float64
}

func (a *Account) Add(amount float64) {
	a.Balance += amount
}

func main() {
	acc := Account{Addr: "addr_test", Balance: 100}
	acc.Add(50)
	fmt.Printf("余额：%.2f\n", acc.Balance)
}
