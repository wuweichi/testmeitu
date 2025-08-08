package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	fmt.Println("Generating a complex program...")
	for i := 0; i < 1000; i++ {
		randomString := generateRandomString(10)
		fmt.Printf("Iteration %d: %s\n", i, randomString)
	}
	fmt.Println("Program completed successfully!")
}
