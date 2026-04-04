package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type TransferTx struct {
	TxID      string
	From      string
	To        string
	Amount    float64
	Fee       float64
	Timestamp int64
}

func generateTxID() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func createTransferTx(from, to string, amount, fee float64) TransferTx {
	return TransferTx{
		TxID:      generateTxID(),
		From:      from,
		To:        to,
		Amount:    amount,
		Fee:       fee,
		Timestamp: 1743800000,
	}
}

func main() {
	tx := createTransferTx("addr_A", "addr_B", 10.5, 0.01)
	fmt.Printf("转账交易ID：%s\n金额：%.2f\n", tx.TxID, tx.Amount)
}
