package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to the Complex Go Program!")
	fmt.Println("Generating a random number between 1 and 100...")
	randomNumber := generateRandomNumber(1, 100)
	fmt.Printf("The random number is: %d\n", randomNumber)
	fmt.Println("This is a simple example, but the program could be extended to over 1000 lines with more complex functionality.")
}
