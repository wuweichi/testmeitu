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
	a := rand.Float64() * 100
	b := rand.Float64() * 100

	// Perform and print various mathematical operations
	fmt.Printf("Generated numbers: %.2f and %.2f\n", a, b)
	fmt.Printf("Addition: %.2f\n", a+b)
	fmt.Printf("Subtraction: %.2f\n", a-b)
	fmt.Printf("Multiplication: %.2f\n", a*b)
	fmt.Printf("Division: %.2f\n", a/b)
	fmt.Printf("Square root of first number: %.2f\n", math.Sqrt(a))
	fmt.Printf("Power (first to the second): %.2f\n", math.Pow(a, b))
	fmt.Printf("Sine of first number: %.2f\n", math.Sin(a))
	fmt.Printf("Cosine of first number: %.2f\n", math.Cos(a))
	fmt.Printf("Tangent of first number: %.2f\n", math.Tan(a))
	fmt.Printf("Logarithm (base 10) of first number: %.2f\n", math.Log10(a))
	fmt.Printf("Natural logarithm of first number: %.2f\n", math.Log(a))
	fmt.Printf("Absolute value of subtraction: %.2f\n", math.Abs(a-b))
	fmt.Printf("Ceiling of first number: %.2f\n", math.Ceil(a))
	fmt.Printf("Floor of first number: %.2f\n", math.Floor(a))
	fmt.Printf("Round of first number: %.2f\n", math.Round(a))
	fmt.Printf("Maximum of the two numbers: %.2f\n", math.Max(a, b))
	fmt.Printf("Minimum of the two numbers: %.2f\n", math.Min(a, b))
	fmt.Printf("Exponential of first number: %.2f\n", math.Exp(a))
	fmt.Printf("Modulo (remainder of division): %.2f\n", math.Mod(a, b))
	fmt.Printf("Hyperbolic sine of first number: %.2f\n", math.Sinh(a))
	fmt.Printf("Hyperbolic cosine of first number: %.2f\n", math.Cosh(a))
	fmt.Printf("Hyperbolic tangent of first number: %.2f\n", math.Tanh(a))
}
