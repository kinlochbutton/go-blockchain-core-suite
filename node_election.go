package main

import "fmt"

func electNode(nodes []string) string {
	return nodes[0]
}

func main() {
	node := electNode([]string{"node1", "node2"})
	fmt.Printf("当选节点：%s\n", node)
}
