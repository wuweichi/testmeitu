package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate two random numbers
	a := rand.Intn(100)
	b := rand.Intn(100)

	// Perform and print various calculations
	fmt.Printf("Generated numbers: %d and %d\n", a, b)
	fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, a*b)
	if b != 0 {
		fmt.Printf("Division: %d / %d = %.2f\n", a, b, float64(a)/float64(b))
	} else {
		fmt.Println("Division by zero is undefined")
	}
	fmt.Printf("Square root of %d: %.2f\n", a, math.Sqrt(float64(a)))
	fmt.Printf("Square root of %d: %.2f\n", b, math.Sqrt(float64(b)))
	fmt.Printf("%d to the power of %d: %.2f\n", a, b, math.Pow(float64(a), float64(b)))

	// A loop to print a multiplication table for a
	fmt.Println("Multiplication table for the first generated number:")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", a, i, a*i)
	}

	// Generate a slice of random numbers and find the max and min
	numbers := make([]int, 10)
	for i := range numbers {
		numbers[i] = rand.Intn(1000)
	}
	max, min := numbers[0], numbers[0]
	for _, num := range numbers {
		if num > max {
			max = num
		}
		if num < min {
			min = num
		}
	}
	fmt.Printf("Generated numbers: %v\n", numbers)
	fmt.Printf("Max: %d, Min: %d\n", max, min)

	// Calculate factorial of a
	factorial := 1
	for i := 1; i <= a; i++ {
		factorial *= i
	}
	fmt.Printf("Factorial of %d: %d\n", a, factorial)

	// Check if numbers are prime
	isPrime := func(n int) bool {
		if n <= 1 {
			return false
		}
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	fmt.Printf("Is %d prime? %t\n", a, isPrime(a))
	fmt.Printf("Is %d prime? %t\n", b, isPrime(b))

	// Generate Fibonacci sequence up to n terms
	n := 10
	fib := make([]int, n)
	fib[0], fib[1] = 1, 1
	for i := 2; i < n; i++ {
		fib[i] = fib[i-1] + fib[i-2]
	}
	fmt.Printf("First %d terms of Fibonacci sequence: %v\n", n, fib)

	// More complex calculations and operations can be added here to extend the program
}