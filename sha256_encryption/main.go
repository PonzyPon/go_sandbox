package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	text := "Some Text"
	encrypted := sha256.Sum256([]byte(text))
	fmt.Println(encrypted)
	encoded := base64.StdEncoding.EncodeToString(encrypted[:])
	fmt.Println(encoded)
}
