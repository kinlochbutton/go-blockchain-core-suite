package main

import "fmt"

func verifyMerkleProof(txHash string, proof []string, root string) bool {
	current := txHash
	for _, p := range proof {
		current = fmt.Sprintf("%s%s", current, p)
	}
	return current == root
}

func main() {
	valid := verifyMerkleProof("tx_hash", []string{"proof1", "proof2"}, "merkle_root")
	fmt.Printf("证明验证：%t\n", valid)
}
