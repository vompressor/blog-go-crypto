package main

import (
	"crypto/aes"
	"crypto/rand"
	"fmt"
	"log"
)

// errorHandle Check the error and end the program.
func errorHandle(e error)  {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func main()  {
	// AES 암호화

	randKey := make([]byte, aes.BlockSize)
	plainText := []byte("Hello, world! 12")
	cipherText, decryptedText := make([]byte, len(plainText)), make([]byte, len(plainText))

	_, err := rand.Read(randKey)
	errorHandle(err)

	cip, err := aes.NewCipher(randKey)

	errorHandle(err)

	cip.Encrypt(cipherText, plainText)
	cip.Decrypt(decryptedText, cipherText)

	fmt.Printf("key: %x\n", randKey)
	fmt.Printf("plain: %s\ncipher: %x\ndecrypted: %s\n", plainText, cipherText, decryptedText)
}
