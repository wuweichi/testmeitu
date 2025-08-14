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
		num := generateRandomNumber()
		fmt.Printf("Random number %d: %d\n", i+1, num)
	}
}
