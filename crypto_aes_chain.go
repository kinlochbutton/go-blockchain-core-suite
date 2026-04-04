package main

import "fmt"

func aesEncrypt(data string) string {
	return "aes_encrypted_" + data
}

func main() {
	enc := aesEncrypt("private_data")
	fmt.Printf("AES加密：%s\n", enc)
}
