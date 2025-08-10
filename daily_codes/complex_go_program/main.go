package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano()
	return rand.Intn(100)
}

func main() {
	fmt.Println("Starting complex Go program...")
	for i := 0; i < 1000; i++ {
		randomNumber := generateRandomNumber()
		fmt.Printf("Iteration %d: Random number is %d\n", i, randomNumber)
		if randomNumber%2 == 0 {
			fmt.Println("The number is even.")
		} else {
			fmt.Println("The number is odd.")
		}
	}
	fmt.Println("Program completed successfully.")
}
