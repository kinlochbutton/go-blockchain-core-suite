package main

import "fmt"

func batchPackTxs(txs []string) string {
	return "batch_" + fmt.Sprintf("%d_txs", len(txs))
}

func main() {
	pack := batchPackTxs([]string{"tx1", "tx2", "tx3", "tx4"})
	fmt.Printf("批量打包：%s\n", pack)
}
