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

	fmt.Printf("Generated numbers: %.2f and %.2f\n", a, b)

	// Perform and display basic operations
	fmt.Printf("Addition: %.2f + %.2f = %.2f\n", a, b, a+b)
	fmt.Printf("Subtraction: %.2f - %.2f = %.2f\n", a, b, a-b)
	fmt.Printf("Multiplication: %.2f * %.2f = %.2f\n", a, b, a*b)
	if b != 0 {
		fmt.Printf("Division: %.2f / %.2f = %.2f\n", a, b, a/b)
	} else {
		fmt.Println("Division by zero is undefined.")
	}

	// Perform and display advanced operations
	fmt.Printf("Square root of %.2f: %.2f\n", a, math.Sqrt(a))
	fmt.Printf("Square root of %.2f: %.2f\n", b, math.Sqrt(b))
	fmt.Printf("%.2f raised to the power of %.2f: %.2f\n", a, b, math.Pow(a, b))
	fmt.Printf("Natural logarithm of %.2f: %.2f\n", a, math.Log(a))
	fmt.Printf("Natural logarithm of %.2f: %.2f\n", b, math.Log(b))

	// Trigonometric functions
	fmt.Printf("Sine of %.2f radians: %.2f\n", a, math.Sin(a))
	fmt.Printf("Cosine of %.2f radians: %.2f\n", a, math.Cos(a))
	fmt.Printf("Tangent of %.2f radians: %.2f\n", a, math.Tan(a))

	// Generate and display a random integer
	randomInt := rand.Intn(100)
	fmt.Printf("Random integer between 0 and 100: %d\n", randomInt)

	// Calculate and display the factorial of the random integer
	factorial := 1
	for i := 1; i <= randomInt; i++ {
		factorial *= i
	}
	fmt.Printf("Factorial of %d: %d\n", randomInt, factorial)

	// Fibonacci sequence up to the random integer
	fmt.Println("Fibonacci sequence up to", randomInt, ":")
	fib1, fib2 := 0, 1
	for fib1 <= randomInt {
		fmt.Print(fib1, " ")
		fib1, fib2 = fib2, fib1+fib2
	}
	fmt.Println()

	// End of program
	fmt.Println("Thank you for using the Complex Calculator!")
}
