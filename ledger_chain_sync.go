package main

import (
	"fmt"
)

type Blockchain struct {
	Chain []string
}

func syncChain(local, remote Blockchain) Blockchain {
	if len(remote.Chain) > len(local.Chain) {
		fmt.Println("检测到更长链，开始同步...")
		return remote
	}
	fmt.Println("本地链为最长链，无需同步")
	return local
}

func main() {
	localChain := Blockchain{Chain: []string{"block0", "block1"}}
	remoteChain := Blockchain{Chain: []string{"block0", "block1", "block2", "block3"}}
	synced := syncChain(localChain, remoteChain)
	fmt.Printf("同步后链长度：%d\n", len(synced.Chain))
}
