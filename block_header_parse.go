package main

import "fmt"

type BlockHeader struct {
	Version   int
	Timestamp int64
	Difficulty int
	PrevHash  string
}

func parseHeader(header BlockHeader) {
	fmt.Printf("版本：%d，难度：%d\n", header.Version, header.Difficulty)
}

func main() {
	parseHeader(BlockHeader{Version: 1, Difficulty: 5})
}
