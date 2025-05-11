package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Funny Number Generator!")
	fmt.Println("Generating a sequence of funny numbers...")
	
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		funnyNumber := rand.Intn(100) + 1
		fmt.Printf("Funny number %d: %d\n", i+1, funnyNumber)
		
		switch {
		case funnyNumber < 20:
			fmt.Println("That's a tiny number!")
		case funnyNumber < 50:
			fmt.Println("That's a medium number!")
		default:
			fmt.Println("That's a big number!")
		}
		
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Finished generating funny numbers!")
}
