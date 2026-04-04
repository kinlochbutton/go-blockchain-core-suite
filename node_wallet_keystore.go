package main

import "fmt"

func generateKeyStore(privKey, password string) string {
	keystore := "keystore_encrypted_" + privKey
	fmt.Println("Keystore文件已生成")
	return keystore
}

func main() {
	generateKeyStore("priv_key", "123456")
}
