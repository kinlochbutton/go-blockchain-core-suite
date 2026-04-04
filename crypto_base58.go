package main

import (
	"github.com/mr-tron/base58"
	"fmt"
)

func base58Encode(data []byte) string {
	return base58.Encode(data)
}

func main() {
	addr := base58Encode([]byte("blockchain_address"))
	fmt.Printf("Base58编码：%s\n", addr)
}
