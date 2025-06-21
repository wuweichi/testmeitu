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

	// Perform calculations
	add := a + b
	sub := a - b
	mul := a * b
	div := a / b
	pow := math.Pow(a, b)
	sqrtA := math.Sqrt(a)
	sqrtB := math.Sqrt(b)

	// Output results
	fmt.Printf("a = %.2f, b = %.2f\n", a, b)
	fmt.Printf("Addition: %.2f + %.2f = %.2f\n", a, b, add)
	fmt.Printf("Subtraction: %.2f - %.2f = %.2f\n", a, b, sub)
	fmt.Printf("Multiplication: %.2f * %.2f = %.2f\n", a, b, mul)
	fmt.Printf("Division: %.2f / %.2f = %.2f\n", a, b, div)
	fmt.Printf("Power: %.2f ^ %.2f = %.2f\n", a, b, pow)
	fmt.Printf("Square root of a: sqrt(%.2f) = %.2f\n", a, sqrtA)
	fmt.Printf("Square root of b: sqrt(%.2f) = %.2f\n", b, sqrtB)

	// Additional complex operations to extend the code length
	for i := 0; i < 100; i++ {
		sinA := math.Sin(a)
		cosA := math.Cos(a)
		tanA := math.Tan(a)
		logA := math.Log(a)
		log10A := math.Log10(a)

		// More calculations
		sinB := math.Sin(b)
		cosB := math.Cos(b)
		tanB := math.Tan(b)
		logB := math.Log(b)
		log10B := math.Log10(b)

		// Even more calculations
		hypot := math.Hypot(a, b)
		ceilA := math.Ceil(a)
		floorA := math.Floor(a)
		ceilB := math.Ceil(b)
		floorB := math.Floor(b)

		// Output more results
		fmt.Printf("Iteration %d: sin(%.2f) = %.2f, cos(%.2f) = %.2f, tan(%.2f) = %.2f\n", i, a, sinA, a, cosA, a, tanA)
		fmt.Printf("Iteration %d: log(%.2f) = %.2f, log10(%.2f) = %.2f\n", i, a, logA, a, log10A)
		fmt.Printf("Iteration %d: sin(%.2f) = %.2f, cos(%.2f) = %.2f, tan(%.2f) = %.2f\n", i, b, sinB, b, cosB, b, tanB)
		fmt.Printf("Iteration %d: log(%.2f) = %.2f, log10(%.2f) = %.2f\n", i, b, logB, b, log10B)
		fmt.Printf("Iteration %d: hypot(%.2f, %.2f) = %.2f\n", i, a, b, hypot)
		fmt.Printf("Iteration %d: ceil(%.2f) = %.2f, floor(%.2f) = %.2f\n", i, a, ceilA, a, floorA)
		fmt.Printf("Iteration %d: ceil(%.2f) = %.2f, floor(%.2f) = %.2f\n", i, b, ceilB, b, floorB)

		// Change a and b slightly for the next iteration
		a += 0.1
		b += 0.1
	}
}
