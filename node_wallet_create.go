package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func createWallet() (string, string) {
	priv := make([]byte, 32)
	rand.Read(priv)
	pub := make([]byte, 32)
	rand.Read(pub)
	return hex.EncodeToString(priv), hex.EncodeToString(pub)
}

func main() {
	priv, pub := createWallet()
	fmt.Printf("钱包私钥：%s\n公钥：%s\n", priv, pub)
}
