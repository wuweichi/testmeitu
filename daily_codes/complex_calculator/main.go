package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Complex Calculator!")
	fmt.Println("This program demonstrates a variety of mathematical operations.")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate two random numbers
	a := rand.Float64() * 100
	b := rand.Float64() * 100

	// Perform and display basic operations
	fmt.Printf("a = %.2f, b = %.2f\n", a, b)
	fmt.Printf("a + b = %.2f\n", a+b)
	fmt.Printf("a - b = %.2f\n", a-b)
	fmt.Printf("a * b = %.2f\n", a*b)
	fmt.Printf("a / b = %.2f\n", a/b)

	// Perform and display advanced operations
	fmt.Printf("sqrt(a) = %.2f\n", math.Sqrt(a))
	fmt.Printf("pow(a, b) = %.2f\n", math.Pow(a, b))
	fmt.Printf("sin(a) = %.2f\n", math.Sin(a))
	fmt.Printf("cos(b) = %.2f\n", math.Cos(b))

	// Demonstrate loops and conditionals
	fmt.Println("\nGenerating 10 random numbers between 0 and 100:")
	for i := 0; i < 10; i++ {
		num := rand.Float64() * 100
		fmt.Printf("%.2f ", num)
		if num > 50 {
			fmt.Println("(greater than 50)")
		} else {
			fmt.Println("(50 or less)")
		}
	}

	// Demonstrate a slice and its operations
	numbers := []float64{a, b, a + b, a - b, a * b, a / b}
	fmt.Println("\nNumbers slice:", numbers)
	fmt.Println("Sum of numbers:", sumSlice(numbers))
	fmt.Println("Average of numbers:", averageSlice(numbers))

	// End of program
	fmt.Println("\nThank you for using the Complex Calculator!")
}

// sumSlice calculates the sum of all elements in a slice of float64
func sumSlice(s []float64) float64 {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	return sum
}

// averageSlice calculates the average of all elements in a slice of float64
func averageSlice(s []float64) float64 {
	if len(s) == 0 {
		return 0
	}
	return sumSlice(s) / float64(len(s))
}
