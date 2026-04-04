package main

import "fmt"

func ringSign(msg string, ring []string) string {
	sig := "ring_signature_" + msg
	fmt.Printf("环签名生成，环成员：%v\n", ring)
	return sig
}

func main() {
	ringSign("anon_tx", []string{"user1", "user2", "user3"})
}
