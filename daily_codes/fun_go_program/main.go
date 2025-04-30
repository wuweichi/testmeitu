package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Fun Go Program!")
	fmt.Println("Let's play a guessing game!")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100) + 1

	fmt.Println("I've picked a number between 1 and 100. Can you guess it?")

	var guess int
	attempts := 0

	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Please enter a valid number!")
			continue
		}
		attempts++

		if guess < secretNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > secretNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You've guessed the number in %d attempts!\n", attempts)
			break
		}
	}

	fmt.Println("Thanks for playing! Goodbye!")
}
