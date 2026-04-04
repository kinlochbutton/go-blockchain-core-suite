package main

import "fmt"

func startAPIServer(port int) {
	fmt.Printf("DApp API服务启动，端口：%d\n", port)
}

func main() {
	startAPIServer(8080)
}
