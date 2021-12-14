package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	// hasher 생성
	hasher := sha256.New()

	// 'hello ' 입력
	hasher.Write([]byte("hello "))
	// 입력된 값 해싱
	h := hasher.Sum(nil)

	// 'world' 입력
	hasher.Write([]byte("world"))
	// 입력된 값 해싱, 위에 입력한 'hello '도 같이 해싱됨
	w1 := hasher.Sum(nil)

	// hasher에 입력된 값 제거
	hasher.Reset()
	// 'world' 입력
	hasher.Write([]byte("world"))
	// 입력한값 해싱
	w2 := hasher.Sum(nil)

	// hasher에 입력된 값 제거
	hasher.Reset()
	// 'world' 입력
	hasher.Write([]byte("hello world"))
	// 입력한값 해싱
	w3 := hasher.Sum(nil)

	fmt.Printf("'hello ' hashed :\n%x\n", h)
	fmt.Printf("'world' hashed :\n%x\n", w1)
	fmt.Printf("hasher reseted, 'world' hashed :\n%x\n", w2)
	fmt.Printf("hasher reseted, 'hello world' hashed :\n%x\n", w3)
}
