package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumber(min, max int) int {
	rand.Seed(time.Now().UnixNano()
	return rand.Intn(max-min+1) + min
}

func main() {
	fmt.Println("Welcome to the Complex Program!")
	fmt.Println("Generating a list of random numbers...")
	
	var numbers []int
	for i := 0; i < 1000; i++ {
		numbers = append(numbers, generateRandomNumber(1, 1000))
	}
	
	fmt.Println("Numbers generated:", numbers)
	
	fmt.Println("Calculating the sum of all numbers...")
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	
	fmt.Println("Total sum:", sum)
	
	fmt.Println("Finding the largest number...")
	largest := numbers[0]
	for _, num := range numbers {
		if num > largest {
			largest = num
		}
	}
	
	fmt.Println("Largest number:", largest)
	
	fmt.Println("Program finished.")
}
