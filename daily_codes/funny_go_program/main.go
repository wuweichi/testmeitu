package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Funny Number Generator!")
	fmt.Println("Generating 1000 lines of random numbers...")
	
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		fmt.Printf("Line %d: %d\n", i+1, rand.Intn(1000000))
	}
	fmt.Println("Done! Thanks for using the Funny Number Generator!")
}
