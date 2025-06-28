package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100. Can you guess it?")

	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	var guess int
	attempts := 0

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanln(&guess)
		attempts++

		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			break
		}
	}
}
