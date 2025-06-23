package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Fun Go Program!")
	fmt.Println("This program generates random numbers and checks if they're prime.")
	fmt.Println("Let's get started!")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate and check 1000 random numbers
	for i := 0; i < 1000; i++ {
		n := rand.Intn(10000) + 1
		if isPrime(n) {
			fmt.Printf("%d is a prime number.\n", n)
		} else {
			fmt.Printf("%d is not a prime number.\n", n)
		}
	}
}

// isPrime checks if a number is prime
func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
