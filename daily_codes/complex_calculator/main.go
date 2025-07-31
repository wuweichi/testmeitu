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

func calculateAverage(numbers []float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}

func calculateStandardDeviation(numbers []float64, average float64) float64 {
	sum := 0.0
	for _, number := range numbers {
		sum += math.Pow(number-average, 2)
	}
	return math.Sqrt(sum / float64(len(numbers)))
}

func main() {
	numbers := generateRandomNumbers(1000)
	average := calculateAverage(numbers)
	standardDeviation := calculateStandardDeviation(numbers, average)

	fmt.Printf("Generated %d random numbers\n", len(numbers))
	fmt.Printf("Average: %.2f\n", average)
	fmt.Printf("Standard Deviation: %.2f\n", standardDeviation)
}
