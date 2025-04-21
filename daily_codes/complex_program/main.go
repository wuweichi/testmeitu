package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber() int {
	return rand.Intn(100)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		number := generateRandomNumber()
		fmt.Printf("Generated number: %d\n", number)
	}
}
