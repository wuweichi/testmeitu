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
	fmt.Println("Starting the complex program...")
	for i := 0; i < 1000; i++ {
		randomNum := generateRandomNumber(1, 100)
		fmt.Printf("Iteration %d: Random number is %d\n", i+1, randomNum)
	}
	fmt.Println("Program completed successfully.")
}
