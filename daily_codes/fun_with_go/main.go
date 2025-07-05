package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Fun with Go!")
	fmt.Println("Generating a random number between 1 and 100...")
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(100) + 1
	fmt.Printf("Your random number is: %d\n", num)
	fmt.Println("Let's see if it's a prime number...")
	isPrime := true
	if num <= 1 {
		isPrime = false
	} else {
		for i := 2; i*i <= num; i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}
	}
	if isPrime {
		fmt.Printf("%d is a prime number!\n", num)
	} else {
		fmt.Printf("%d is not a prime number.\n", num)
	}
	fmt.Println("Thanks for playing!")
}
