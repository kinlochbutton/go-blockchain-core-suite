package main

import "fmt"

func checkWhiteList(node string) bool {
	fmt.Printf("节点%s在白名单内\n", node)
	return true
}

func main() {
	checkWhiteList("trust_node")
}
