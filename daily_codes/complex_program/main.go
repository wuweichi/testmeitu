package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}

func main() {
	fmt.Println("Welcome to the Complex Number Generator!")
	for i := 0; i < 1000; i++ {
		randomNum := generateRandomNumber(1, 1000)
		fmt.Printf("Generated random number %d: %d\n", i+1, randomNum)
	}
	fmt.Println("Finished generating numbers!")
}
