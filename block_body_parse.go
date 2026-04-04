package main

import "fmt"

func parseBlockBody(txs []string) {
	fmt.Printf("区块体交易数：%d\n", len(txs))
}

func main() {
	parseBlockBody([]string{"tx1", "tx2", "tx3"})
}
