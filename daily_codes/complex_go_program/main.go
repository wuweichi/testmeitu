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
	fmt.Println("Starting the complex Go program...")
	for i := 0; i < 1000; i++ {
		randomNum := generateRandomNumber()
		fmt.Printf("Iteration %d: Random number is %d\n", i, randomNum)
		time.Sleep(10 * time.Millisecond)
	}
	fmt.Println("Program completed successfully.")
}
