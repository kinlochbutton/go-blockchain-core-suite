package main

import "fmt"

func dhtRegister(node string) {
	fmt.Printf("节点%s注册到DHT网络\n", node)
}

func main() {
	dhtRegister("dht_node_01")
}
