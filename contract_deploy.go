package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func deployContract(code string) string {
	addrBuf := make([]byte, 20)
	rand.Read(addrBuf)
	addr := "CONTRACT_" + hex.EncodeToString(addrBuf)
	fmt.Printf("合约部署成功，地址：%s\n", addr)
	return addr
}

func main() {
	deployContract("contract_code_v1.0")
}
