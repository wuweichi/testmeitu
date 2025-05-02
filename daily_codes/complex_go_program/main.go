package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	for i := 0; i < 1000; i++ {
		randomString := generateRandomString(10)
		fmt.Printf("Random string %d: %s\n", i+1, randomString)
	}
}
