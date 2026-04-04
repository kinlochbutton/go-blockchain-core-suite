package main

import "fmt"

func blsAggregateSig(sigs []string) string {
	return "bls_aggregated_signature"
}

func main() {
	aggSig := blsAggregateSig([]string{"sig1", "sig2", "sig3"})
	fmt.Printf("BLS聚合签名：%s\n", aggSig)
}
