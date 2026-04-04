package main

import "fmt"

func routeCrossChainTx(txID string, fromChain, toChain string) {
	fmt.Printf("跨链交易%s：%s -> %s 路由完成\n", txID, fromChain, toChain)
}

func main() {
	routeCrossChainTx("cross_tx_001", "ETH", "BSC")
}
