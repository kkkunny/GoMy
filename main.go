package main

import (
	"GoMy/crypto"
	"fmt"
)

func main() {
	content := "Hello World"
	result, err := crypto.EncodeSha3_256([]byte(content), []byte{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	result, err = crypto.EncodeSha256([]byte(content), []byte{})
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
