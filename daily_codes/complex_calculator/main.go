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
	for i := range numbers {
		numbers[i] = rand.Float64() * 100
	}
	return numbers
}

func calculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func calculateStandardDeviation(numbers []float64, average float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += math.Pow(num-average, 2)
	}
	variance := sum / float64(len(numbers))
	return math.Sqrt(variance)
}

func main() {
	numbers := generateRandomNumbers(1000)
	average := calculateAverage(numbers)
	standardDeviation := calculateStandardDeviation(numbers, average)

	fmt.Printf("Generated %d random numbers\n", len(numbers))
	fmt.Printf("Average: %.2f\n", average)
	fmt.Printf("Standard Deviation: %.2f\n", standardDeviation)
}
