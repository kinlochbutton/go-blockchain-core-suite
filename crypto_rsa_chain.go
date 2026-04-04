package main

import (
	"crypto/rand"
	"crypto/rsa"
	"fmt"
)

func generateRSAKey() (*rsa.PrivateKey, error) {
	return rsa.GenerateKey(rand.Reader, 2048)
}

func main() {
	priv, _ := generateRSAKey()
	fmt.Printf("RSA私钥长度：%d bits\n", priv.N.BitLen())
}
