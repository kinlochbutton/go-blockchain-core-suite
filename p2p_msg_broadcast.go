package main

import "fmt"

func broadcastMsg(msg string, nodes []string) {
	for _, node := range nodes {
		fmt.Printf("向节点%s广播消息：%s\n", node, msg)
	}
}

func main() {
	broadcastMsg("new_block_100", []string{"node1", "node2", "node3"})
}
