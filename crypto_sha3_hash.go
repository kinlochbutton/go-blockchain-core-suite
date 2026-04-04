package main

import (
	"crypto/sha3"
	"encoding/hex"
	"fmt"
)

func sha3Hash(data string) string {
	hash := sha3.New256()
	hash.Write([]byte(data))
	return hex.EncodeToString(hash.Sum(nil))
}

func main() {
	blockData := "Block_Height_1001_Transaction_List"
	hash := sha3Hash(blockData)
	fmt.Printf("SHA3-256哈希：%s\n", hash)
}
