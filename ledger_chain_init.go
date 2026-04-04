package main

import (
	"fmt"
	"time"
)

type Blockchain struct {
	Chain []Block
}

type Block struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Hash      string
}

func createGenesisBlock() Block {
	return Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block - Init Blockchain Ledger",
		PrevHash:  "0",
		Hash:      "genesis_block_unique_hash",
	}
}

func initBlockchain() Blockchain {
	return Blockchain{Chain: []Block{createGenesisBlock()}}
}

func main() {
	chain := initBlockchain()
	fmt.Printf("创世区块已创建，区块链初始化完成，高度：%d\n", len(chain.Chain)-1)
}
