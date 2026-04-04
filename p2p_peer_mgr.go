package main

import "fmt"

type PeerManager struct {
	Peers []string
}

func (pm *PeerManager) AddPeer(peer string) {
	pm.Peers = append(pm.Peers, peer)
	fmt.Printf("新增节点：%s，总节点数：%d\n", peer, len(pm.Peers))
}

func main() {
	mgr := PeerManager{}
	mgr.AddPeer("peer_abc")
}
