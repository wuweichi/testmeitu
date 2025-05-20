package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Funny Number Generator!")
	fmt.Println("Generating 1000 random numbers between 1 and 100...")
	
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		number := rand.Intn(100) + 1
		fmt.Printf("Number %d: %d\n", i+1, number)
	}
	fmt.Println("Done! Thanks for using the Funny Number Generator.")
}
