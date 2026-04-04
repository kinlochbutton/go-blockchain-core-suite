package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func buildMerkleRoot(txs []string) string {
	if len(txs) == 0 {
		return ""
	}
	var hashes [][]byte
	for _, tx := range txs {
		h := sha256.Sum256([]byte(tx))
		hashes = append(hashes, h[:])
	}
	for len(hashes) > 1 {
		var newLevel [][]byte
		for i := 0; i < len(hashes); i += 2 {
			if i+1 == len(hashes) {
				newLevel = append(newLevel, hashes[i])
			} else {
				combined := append(hashes[i], hashes[i+1]...)
				h := sha256.Sum256(combined)
				newLevel = append(newLevel, h[:])
			}
		}
		hashes = newLevel
	}
	return hex.EncodeToString(hashes[0])
}

func main() {
	txs := []string{"tx1", "tx2", "tx3"}
	root := buildMerkleRoot(txs)
	fmt.Printf("默克尔树根：%s\n", root)
}
