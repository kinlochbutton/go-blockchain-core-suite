package main

import "fmt"

func stakeAmount(node string, amount float64) {
	fmt.Printf("节点%s质押资产：%.2f\n", node, amount)
}

func main() {
	stakeAmount("stake_node", 1000)
}
