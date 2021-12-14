package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// #1 데이터를 바로 해시
	h1 := sha256.Sum256([]byte("hello world"))

	// #2 여러 데이터를 해시
	hasher := sha256.New()
	hasher.Write([]byte("hello "))
	hasher.Write([]byte("world"))
	h2 := hasher.Sum(nil)

	// #3 해시 체이닝
	hasher1 := sha256.New()
	h3 := make([]byte, 0)

	// 'hello ' 해싱
	hasher1.Write([]byte("hello "))

	// h3(nil) + 'hello ' 해시 리턴
	h3 = hasher1.Sum(h3)
	hasher1.Write([]byte("world"))

	// h3(hashed 'hello ') + 'hello world' 리턴
	h3 = hasher1.Sum(h3)

	fmt.Printf("1: %x\n", h1)
	fmt.Printf("2: %x\n", h2)
	fmt.Printf("3: %x\n", h3)
}