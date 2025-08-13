package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Fun Go Program!")
	fmt.Println("Generating a random number between 1 and 100...")
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) + 1
	fmt.Printf("The random number is: %d\n", randomNumber)
	fmt.Println("Thanks for using the Fun Go Program!")
}
