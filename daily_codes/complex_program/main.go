package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func main() {
	for i := 0; i < 1000; i++ {
		randomNumber := generateRandomNumber()
		fmt.Printf("Random number %d: %d\n", i+1, randomNumber)
	}
}
