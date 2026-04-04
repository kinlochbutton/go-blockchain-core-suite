package main

import "fmt"

func createTimeLockTx(timestamp int64) string {
	return fmt.Sprintf("time_lock_tx_%d", timestamp)
}

func main() {
	tx := createTimeLockTx(1744000000)
	fmt.Printf("时间锁交易：%s\n", tx)
}
