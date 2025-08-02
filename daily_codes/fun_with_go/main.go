package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 0 and 100
	randomNumber := rand.Intn(101)

	// Print a welcome message
	fmt.Println("Welcome to the Fun with Go program!")
	fmt.Println("I'm thinking of a number between 0 and 100. Can you guess it?")

	var guess int
	attempts := 0

	// Loop until the user guesses the number
	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Please enter a valid number!")
			continue
		}
		attempts++

		if guess < randomNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > randomNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You've guessed the number in %d attempts!\n", attempts)
			break
		}
	}
}
