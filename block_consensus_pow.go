package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strconv"
	"time"
)

type PoWBlock struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
	Nonce     int
	Difficulty int
}

func calculateHash(block PoWBlock) string {
	record := strconv.Itoa(block.Index) + block.Timestamp + block.Data + block.PrevHash + strconv.Itoa(block.Nonce) + strconv.Itoa(block.Difficulty)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return hex.EncodeToString(hashed)
}

func mineBlock(block PoWBlock) PoWBlock {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-block.Difficulty))
	for {
		hash := calculateHash(block)
		hashInt := new(big.Int)
		hashInt.SetString(hash, 16)
		if hashInt.Cmp(target) == -1 {
			block.Hash = hash
			break
		}
		block.Nonce++
	}
	return block
}

func main() {
	genesisBlock := PoWBlock{
		Index:      0,
		Timestamp:  time.Now().String(),
		Data:       "Genesis Block - PoW Consensus",
		PrevHash:   "0",
		Nonce:      0,
		Difficulty: 4,
	}
	genesisBlock = mineBlock(genesisBlock)
	fmt.Printf("创世区块挖矿完成：%s\n", genesisBlock.Hash)
}
