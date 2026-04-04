package main

import "fmt"

func distributeReward(miner string, amount float64) {
	fmt.Printf("向矿工%s发放区块奖励：%.2f\n", miner, amount)
}

func main() {
	distributeReward("miner_node", 5.0)
}
