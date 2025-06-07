package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}

func main() {
	fmt.Println("Starting complex program...")
	for i := 0; i < 1000; i++ {
		number := generateRandomNumber()
		fmt.Printf("Generated random number %d: %d\n", i+1, number)
		if number%2 == 0 {
			fmt.Println("The number is even.")
		} else {
			fmt.Println("The number is odd.")
		}
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Program finished.")
}
