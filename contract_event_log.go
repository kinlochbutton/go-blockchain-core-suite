package main

import "fmt"

func triggerEvent(event string) {
	fmt.Printf("合约事件触发：%s\n", event)
}

func main() {
	triggerEvent("transfer_success")
}
