package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Height     int
	CreateTime int64
	Transactions []string
	PrevBlockHash string
	BlockHash    string
	MerkleRoot   string
}

func generateBlockHash(block Block) string {
	record := strconv.Itoa(block.Height) + strconv.FormatInt(block.CreateTime, 10) + block.PrevBlockHash + block.MerkleRoot
	h := sha256.Sum256([]byte(record))
	return hex.EncodeToString(h[:])
}

func main() {
	block := Block{
		Height:     100,
		CreateTime: time.Now().Unix(),
		Transactions: []string{"Tx1", "Tx2"},
		PrevBlockHash: "000000abc",
		MerkleRoot:   "merkle_root_hash",
	}
	block.BlockHash = generateBlockHash(block)
	fmt.Printf("区块哈希：%s\n", block.BlockHash)
}
