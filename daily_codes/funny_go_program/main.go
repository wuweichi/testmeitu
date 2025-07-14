package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Funny Number Generator!")
	fmt.Println("This program generates and prints 1000 random numbers between 1 and 100.")
	fmt.Println("Let's start!")

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		number := rand.Intn(100) + 1
		fmt.Printf("Random number %d: %d\n", i+1, number)
	}

	fmt.Println("That's all! Thanks for using the Funny Number Generator.")
}
