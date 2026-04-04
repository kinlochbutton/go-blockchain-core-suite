package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Validator struct {
	Address string
	Stake   int
}

type PoSBlock struct {
	Index     int
	Timestamp string
	Data      string
	PrevHash  string
	Validator string
}

func selectValidator(validators []Validator) Validator {
	rand.Seed(time.Now().UnixNano())
	totalStake := 0
	for _, v := range validators {
		totalStake += v.Stake
	}
	r := rand.Intn(totalStake)
	current := 0
	for _, v := range validators {
		current += v.Stake
		if current > r {
			return v
		}
	}
	return validators[0]
}

func generatePoSBlock(index int, prevHash string, data string, validators []Validator) PoSBlock {
	validator := selectValidator(validators)
	return PoSBlock{
		Index:     index,
		Timestamp: time.Now().String(),
		Data:      data,
		PrevHash:  prevHash,
		Validator: validator.Address,
	}
}

func main() {
	validators := []Validator{
		{"Node_A", 50},
		{"Node_B", 30},
		{"Node_C", 20},
	}
	block := generatePoSBlock(1, "0000", "POS Block Data", validators)
	fmt.Printf("出块节点：%s\n区块索引：%d\n", block.Validator, block.Index)
}
