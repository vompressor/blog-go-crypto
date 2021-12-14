package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
)

// errorHandle Check the error and end the program.
func errorHandle(e error)  {
	if e != nil {
		log.Fatal(e.Error())
	}
}

func PaddingPKCS7(src []byte, size int) []byte {
	if src == nil {
		return bytes.Repeat([]byte{byte(size)}, size)
	}

	padlen := 1
	for ((len(src) + padlen) % size) != 0 {
		padlen = padlen + 1
	}

	pad := bytes.Repeat([]byte{byte(padlen)}, padlen)
	return append(src, pad...)
}

func UnPaddingPKCS7(src []byte, size int) ([]byte, error) {
	if len(src) == 0 {
		return nil, errors.New("invalid padding")
	}

	padlen := int(src[len(src)-1])

	if padlen <= 0 || padlen > size {
		return nil, errors.New("invalid padding")
	}

	pad := src[len(src)-padlen:]

	for _, n := range pad {
		if int(n) != padlen {
			return nil, errors.New("invalid padding")
		}
	}

	return src[:len(src)-padlen], nil
}

func main()  {
	// AES CBC Mode 암호화

	randKey := make([]byte, aes.BlockSize)
	plainText := []byte(`동해 물과 백두산이 마르고 닳도록 
하느님이 보우하사 우리나라 만세.
무궁화 삼천리 화려 강산
대한 사람, 대한으로 길이 보전하세.`)

	_, err := rand.Read(randKey)
	errorHandle(err)

	cip, err := aes.NewCipher(randKey)

	errorHandle(err)

	iv := make([]byte, cip.BlockSize())
	_, err = rand.Read(iv)
	errorHandle(err)

	cbcEncrypter := cipher.NewCBCEncrypter(cip, iv)
	cbcDecrypter := cipher.NewCBCDecrypter(cip, iv)

	padded := PaddingPKCS7(plainText, cbcEncrypter.BlockSize())

	cipherText := make([]byte, len(padded))

	cbcEncrypter.CryptBlocks(cipherText, padded)
	fmt.Printf("CipherText:\n%x\n\n", cipherText)

	decrypted := make([]byte, len(cipherText))
	cbcDecrypter.CryptBlocks(decrypted, cipherText)
	unpadded, err := UnPaddingPKCS7(decrypted, cbcDecrypter.BlockSize())
	
	errorHandle(err)
	
	fmt.Printf("PlainText:\n%s\n\n", unpadded)

}
