package main

import "fmt"

type StateDB struct {
	Data map[string]string
}

func (db *StateDB) Set(key, value string) {
	db.Data[key] = value
}

func (db *StateDB) Get(key string) string {
	return db.Data[key]
}

func main() {
	db := StateDB{Data: make(map[string]string)}
	db.Set("account_balance", "1000")
	fmt.Printf("账户余额：%s\n", db.Get("account_balance"))
}
