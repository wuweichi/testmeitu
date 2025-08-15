package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Fun Go Program!")
	fmt.Println("This program generates random numbers and checks if they're prime.")

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		n := rand.Intn(100) + 1
		if isPrime(n) {
			fmt.Printf("%d is a prime number.\n", n)
		} else {
			fmt.Printf("%d is not a prime number.\n", n)
		}
	}
}

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
