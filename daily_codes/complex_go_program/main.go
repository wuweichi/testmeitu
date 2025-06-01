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
	fmt.Println("Welcome to the Complex Go Program!")
	fmt.Println("Generating a series of random numbers...")
	for i := 0; i < 100; i++ {
		randomNum := generateRandomNumber(1, 1000)
		fmt.Printf("Random number %d: %d\n", i+1, randomNum)
	}
	fmt.Println("Program completed successfully.")
}
