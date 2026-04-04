package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

func generateECDSAKey() (*ecdsa.PrivateKey, error) {
	return ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
}

func signData(priv *ecdsa.PrivateKey, data []byte) (*big.Int, *big.Int, error) {
	hash := sha256.Sum256(data)
	return ecdsa.Sign(rand.Reader, priv, hash[:])
}

func verifySig(pub *ecdsa.PublicKey, data []byte, r, s *big.Int) bool {
	hash := sha256.Sum256(data)
	return ecdsa.Verify(pub, hash[:], r, s)
}

func main() {
	priv, _ := generateECDSAKey()
	data := []byte("Blockchain Transaction Data")
	r, s, _ := signData(priv, data)
	result := verifySig(&priv.PublicKey, data, r, s)
	fmt.Printf("签名验证结果：%t\n", result)
}
