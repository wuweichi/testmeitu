package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func generateRandomNumbers(count int) []float64 {
	rand.Seed(time.Now().UnixNano())
	numbers := make([]float64, count)
	for i := 0; i < count; i++ {
		numbers[i] = rand.Float64() * 100
	}
	return numbers
}

func calculateSum(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func calculateAverage(numbers []float64) float64 {
	return calculateSum(numbers) / float64(len(numbers))
}

func calculateStandardDeviation(numbers []float64) float64 {
	average := calculateAverage(numbers)
	variance := 0.0
	for _, num := range numbers {
		variance += math.Pow(num-average, 2)
	}
	variance = variance / float64(len(numbers))
	return math.Sqrt(variance)
}

func main() {
	numbers := generateRandomNumbers(1000)
	fmt.Printf("Generated %d random numbers\n", len(numbers))
	fmt.Printf("Sum: %.2f\n", calculateSum(numbers))
	fmt.Printf("Average: %.2f\n", calculateAverage(numbers))
	fmt.Printf("Standard Deviation: %.2f\n", calculateStandardDeviation(numbers))
}
