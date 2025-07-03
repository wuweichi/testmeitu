package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate and print 1000 random numbers
	for i := 0; i < 1000; i++ {
		fmt.Println(rand.Intn(100))
	}

	// A simple game: guess the number
	fmt.Println("Let's play a game! Guess the number between 1 and 100.")
	target := rand.Intn(100) + 1
	var guess int
	attempts := 0

	for {
		fmt.Print("Enter your guess: ")
		fmt.Scanf("%d", &guess)
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

	// Print a farewell message
	fmt.Println("Thanks for playing! Goodbye.")
}
