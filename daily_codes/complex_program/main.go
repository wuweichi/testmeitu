package main

import (
	"fmt"
	"math/rand"
	"time"
)

func generateRandomNumbers(count int) []int {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]int, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Intn(100)
	}
	return numbers
}

func bubbleSort(numbers []int) []int {
	n := len(numbers)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
	return numbers
}

func main() {
	numbers := generateRandomNumbers(1000)
	sortedNumbers := bubbleSort(numbers)
	fmt.Println("Sorted Numbers:", sortedNumbers)
}
