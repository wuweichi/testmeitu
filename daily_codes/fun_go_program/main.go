package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Fun Go Program!")
	fmt.Println("This program generates random numbers and checks if they are prime.")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		number := rand.Intn(100) + 1
		fmt.Printf("Number: %d, Is Prime: %v\n", number, isPrime(number))
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}
	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
