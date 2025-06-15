package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		num := generateRandomNumber(1, 100)
		fmt.Printf("Random number %d: %d\n", i+1, num)
	}
}
