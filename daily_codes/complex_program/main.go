package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber() int {
	rand.Seed(time.Now().UnixNano()
	return rand.Intn(100)
}

func main() {
	fmt.Println("Starting complex program...")
	for i := 0; i < 1000; i++ {
		num := generateRandomNumber()
		fmt.Printf("Generated random number %d: %d\n", i+1, num)
	}
	fmt.Println("Program completed successfully.")
}
