package main

import "fmt"

type MemPool struct {
	PendingTxs []string
}

func (mp *MemPool) AddTx(txID string) {
	mp.PendingTxs = append(mp.PendingTxs, txID)
	fmt.Printf("交易%s加入内存池\n", txID)
}

func (mp *MemPool) CleanTxs() {
	mp.PendingTxs = []string{}
	fmt.Println("内存池已清空")
}

func main() {
	pool := MemPool{}
	pool.AddTx("tx_123456")
	pool.CleanTxs()
}
