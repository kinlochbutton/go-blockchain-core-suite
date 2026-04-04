package main

import "fmt"

func buildBlockIndex(height int) string {
	index := fmt.Sprintf("index_block_%d", height)
	fmt.Printf("区块索引创建：%s\n", index)
	return index
}

func main() {
	buildBlockIndex(100)
}
