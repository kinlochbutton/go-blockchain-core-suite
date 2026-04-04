package main

import "fmt"

func udpTransfer(data string) {
	fmt.Printf("UDP高速传输：%s\n", data)
}

func main() {
	udpTransfer("large_block_data")
}
