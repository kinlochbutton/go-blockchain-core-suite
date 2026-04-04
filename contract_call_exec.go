package main

import "fmt"

type Contract struct {
	Addr  string
	State map[string]int
}

func (c *Contract) Call(method string, param int) {
	if method == "add" {
		c.State["count"] += param
		fmt.Printf("合约调用成功，当前计数：%d\n", c.State["count"])
	}
}

func main() {
	contract := Contract{Addr: "contract_addr", State: map[string]int{"count": 0}}
	contract.Call("add", 5)
}
