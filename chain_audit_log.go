package main

import "fmt"

func writeAuditLog(op string) {
	fmt.Printf("审计日志：%s\n", op)
}

func main() {
	writeAuditLog("deploy_contract")
}
