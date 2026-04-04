package main

import "fmt"

func issueToken(name string, total float64) {
	fmt.Printf("代币%s发行成功，总量：%.2f\n", name, total)
}

func main() {
	issueToken("GOCOIN", 21000000)
}
