package main

import "fmt"

func linkIPFSHash(hash string) {
	fmt.Printf("IPFS哈希上链：%s\n", hash)
}

func main() {
	linkIPFSHash("QmXYZ123456789")
}
