package main

import "fmt"

type P2PNode struct {
	ID     string
	Status string
}

func dialNode(nodeID string) P2PNode {
	fmt.Printf("节点%s连接成功\n", nodeID)
	return P2PNode{ID: nodeID, Status: "connected"}
}

func main() {
	node := dialNode("p2p_node_001")
	fmt.Printf("节点状态：%s\n", node.Status)
}
