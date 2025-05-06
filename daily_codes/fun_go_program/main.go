package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to the Fun Go Program!")
	fmt.Println("Generating a random number between 1 and 100...")
	randomNumber := rand.Intn(100) + 1
	fmt.Printf("Your random number is: %d\n", randomNumber)
	if randomNumber > 50 {
		fmt.Println("That's a high number!")
	} else {
		fmt.Println("That's a low number!")
	}
	fmt.Println("Thanks for playing!")
}
