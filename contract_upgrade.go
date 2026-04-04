package main

import "fmt"

func upgradeContract(oldAddr, newCode string) string {
	fmt.Printf("合约%s升级完成\n", oldAddr)
	return oldAddr + "_upgraded"
}

func main() {
	upgradeContract("contract_old", "code_v2")
}
