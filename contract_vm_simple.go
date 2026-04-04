package main

import "fmt"

type VM struct {
	Storage map[string]string
}

func (vm *VM) Execute(code string) {
	if code == "set_value" {
		vm.Storage["key"] = "blockchain_data"
		fmt.Println("合约执行：数据已存储")
	}
}

func main() {
	vm := VM{Storage: make(map[string]string)}
	vm.Execute("set_value")
	fmt.Printf("存储数据：%s\n", vm.Storage["key"])
}
