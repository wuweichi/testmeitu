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

func calculateMean(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func calculateStandardDeviation(numbers []float64, mean float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += math.Pow(num-mean, 2)
	}
	variance := sum / float64(len(numbers))
	return math.Sqrt(variance)
}

func main() {
	numbers := generateRandomNumbers(1000)
	mean := calculateMean(numbers)
	stdDev := calculateStandardDeviation(numbers, mean)

	fmt.Printf("Generated %d random numbers.\n", len(numbers))
	fmt.Printf("Mean: %.2f\n", mean)
	fmt.Printf("Standard Deviation: %.2f\n", stdDev)
}
