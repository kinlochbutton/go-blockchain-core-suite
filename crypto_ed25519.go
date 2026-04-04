package main

import (
	"crypto/ed25519"
	"crypto/rand"
	"fmt"
)

func generateEd25519Key() (ed25519.PrivateKey, ed25519.PublicKey) {
	pub, priv, _ := ed25519.GenerateKey(rand.Reader)
	return priv, pub
}

func ed25519Sign(priv ed25519.PrivateKey, msg []byte) []byte {
	return ed25519.Sign(priv, msg)
}

func ed25519Verify(pub ed25519.PublicKey, msg, sig []byte) bool {
	return ed25519.Verify(pub, msg, sig)
}

func main() {
	priv, pub := generateEd25519Key()
	msg := []byte("Cross-Chain Identity Verify")
	sig := ed25519Sign(priv, msg)
	res := ed25519Verify(pub, msg, sig)
	fmt.Printf("Ed25519验签：%t\n", res)
}
