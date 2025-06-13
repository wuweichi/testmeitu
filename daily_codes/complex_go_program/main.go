package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min) + min
}

func main() {
	for i := 0; i < 1000; i++ {
		randomNum := generateRandomNumber(1, 100)
		fmt.Printf("Random number %d: %d\n", i+1, randomNum)
	}
}
