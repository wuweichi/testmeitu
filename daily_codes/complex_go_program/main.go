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
	fmt.Println("Generating 1000 random numbers between 1 and 100:")
	for i := 0; i < 1000; i++ {
		randomNumber := generateRandomNumber(1, 100)
		fmt.Printf("%d ", randomNumber)
		if (i+1)%10 == 0 {
			fmt.Println()
		}
	}
	fmt.Println("\nDone generating random numbers.")
}
