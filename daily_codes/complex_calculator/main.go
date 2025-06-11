package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to the Complex Calculator!")
	fmt.Println("This program demonstrates various mathematical operations.")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate two random numbers
	a := rand.Float64() * 100
	b := rand.Float64() * 100

	// Perform and display basic operations
	fmt.Printf("Generated numbers: %.2f and %.2f\n", a, b)
	fmt.Printf("Addition: %.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("Subtraction: %.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("Multiplication: %.2f * %.2f = %.2f\n", a, b, a*b)
	fmt.Printf("Division: %.2f / %.2f = %.2f\n", a, b, a/b)

	// Perform and display advanced operations
	fmt.Printf("Square root of %.2f: %.2f\n", a, math.Sqrt(a))
	fmt.Printf("%.2f raised to the power of %.2f: %.2f\n", a, b, math.Pow(a, b))
	fmt.Printf("Natural logarithm of %.2f: %.2f\n", a, math.Log(a))
	fmt.Printf("Sine of %.2f radians: %.2f\n", a, math.Sin(a))

	// Demonstrate loop and conditional operations
	fmt.Println("\nNow demonstrating a loop that prints numbers divisible by 3 up to 100:")
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println("\n")

	// Demonstrate array operations
	fmt.Println("Demonstrating array operations with 10 random numbers:")
	var numbers [10]float64
	for i := 0; i < 10; i++ {
		numbers[i] = rand.Float64() * 100
		fmt.Printf("%.2f ", numbers[i])
	}
	fmt.Println("\n")

	// Find and display the maximum number in the array
	max := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
	}
	fmt.Printf("The maximum number in the array is: %.2f\n", max)

	// Demonstrate function usage
	fmt.Println("\nCalculating the factorial of 10:")
	factorial := 1
	for i := 1; i <= 10; i++ {
		factorial *= i
	}
	fmt.Printf("10! = %d\n", factorial)

	// End of program
	fmt.Println("\nThank you for using the Complex Calculator!")
}
