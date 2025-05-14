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
	fmt.Println("Let's play a guessing game! Try to guess the number.")
	var guess int
	attempts := 0
	for {
		fmt.Print("Enter your guess: ")
		fmt.Scan(&guess)
		attempts++
		if guess < randomNumber {
			fmt.Println("Too low! Try again.")
		} else if guess > randomNumber {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			break
		}
	}
	fmt.Println("Thanks for playing the Fun Go Program!")
}
