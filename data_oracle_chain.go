package main

import "fmt"

func oracleUploadData(data string) {
	fmt.Printf("预言机数据上链：%s\n", data)
}

func main() {
	oracleUploadData("price_ETH=3000")
}
