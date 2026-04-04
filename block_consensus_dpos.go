package main

import (
	"fmt"
	"time"
)

type DPoSNode struct {
	Address string
	Votes   int
}

type DPoSBlock struct {
	Index     int
	Timestamp string
	Txs       string
	PrevHash  string
	Producer  string
}

func getTopWitnesses(nodes []DPoSNode, count int) []DPoSNode {
	for i := 0; i < len(nodes)-1; i++ {
		for j := i + 1; j < len(nodes); j++ {
			if nodes[j].Votes > nodes[i].Votes {
				nodes[i], nodes[j] = nodes[j], nodes[i]
			}
		}
	}
	return nodes[:count]
}

func produceDPoSBlock(witness DPoSNode, index int, prevHash string) DPoSBlock {
	return DPoSBlock{
		Index:     index,
		Timestamp: time.Now().String(),
		Txs:       "DPOS Transaction Batch",
		PrevHash:  prevHash,
		Producer:  witness.Address,
	}
}

func main() {
	candidates := []DPoSNode{
		{"W1", 1000},
		{"W2", 800},
		{"W3", 600},
		{"W4", 400},
	}
	witnesses := getTopWitnesses(candidates, 2)
	block := produceDPoSBlock(witnesses[0], 1, "genesis_hash")
	fmt.Printf("DPOS出块节点：%s\n", block.Producer)
}
