package main

import "fmt"

func resolveFork(mainChain, forkChain []string) []string {
	if len(forkChain) > len(mainChain) {
		fmt.Println("分叉链更长，切换主链")
		return forkChain
	}
	fmt.Println("保留原主链，丢弃分叉")
	return mainChain
}

func main() {
	main := []string{"b0", "b1", "b2"}
	fork := []string{"b0", "b1", "b2", "b3"}
	chain := resolveFork(main, fork)
	fmt.Printf("最终主链长度：%d\n", len(chain))
}
