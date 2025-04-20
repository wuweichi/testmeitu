package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		randomString := generateRandomString(10)
		fmt.Printf("Random string %d: %s\n", i+1, randomString)
	}
}
