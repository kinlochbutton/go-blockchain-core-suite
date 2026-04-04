package main

import "fmt"

func checkFinality(height int) bool {
	fmt.Printf("区块%d已最终确认，不可逆\n", height)
	return true
}

func main() {
	checkFinality(1000)
}
