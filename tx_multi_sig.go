package main

import "fmt"

func multiSignTx(signers []string) bool {
	fmt.Printf("多签完成，签名人数：%d\n", len(signers))
	return true
}

func main() {
	multiSignTx([]string{"signer1", "signer2", "signer3"})
}
