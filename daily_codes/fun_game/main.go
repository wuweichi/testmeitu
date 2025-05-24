package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Fun Game!")
	fmt.Println("Guess the number between 1 and 100")

	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1

	var guess int
	attempts := 0

	for {
		fmt.Print("Enter your guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number.")
			continue
		}
		attempts++

		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			break
		}
	}
}
