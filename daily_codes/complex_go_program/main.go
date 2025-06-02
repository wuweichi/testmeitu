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
	fmt.Println("Welcome to the Complex Go Program!")
	fmt.Println("Generating a random number between 0 and 99...")
	randomNumber := generateRandomNumber()
	fmt.Printf("Your random number is: %d\n", randomNumber)
	
	for i := 0; i < 1000; i++ {
		fmt.Printf("Loop iteration %d: Still running...\n", i)
	}
	
	fmt.Println("Program completed successfully!")
}
