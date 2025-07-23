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

	// Perform and print various mathematical operations
	fmt.Printf("Generated numbers: %d and %d\n", a, b)
	fmt.Printf("Addition: %d + %d = %d\n", a, b, a+b)
	fmt.Printf("Subtraction: %d - %d = %d\n", a, b, a-b)
	fmt.Printf("Multiplication: %d * %d = %d\n", a, b, a*b)
	if b != 0 {
		fmt.Printf("Division: %d / %d = %.2f\n", a, b, float64(a)/float64(b))
	} else {
		fmt.Println("Division by zero is undefined")
	}
	fmt.Printf("Modulus: %d %% %d = %d\n", a, b, a%b)
	fmt.Printf("Power: %d^%d = %.2f\n", a, b, math.Pow(float64(a), float64(b)))
	fmt.Printf("Square root of %d: %.2f\n", a, math.Sqrt(float64(a)))
	fmt.Printf("Square root of %d: %.2f\n", b, math.Sqrt(float64(b)))

	// Generate a random angle in degrees
	angle := rand.Float64() * 360
	radians := angle * math.Pi / 180
	fmt.Printf("\nRandom angle: %.2f degrees\n", angle)
	fmt.Printf("Sine: %.2f\n", math.Sin(radians))
	fmt.Printf("Cosine: %.2f\n", math.Cos(radians))
	fmt.Printf("Tangent: %.2f\n", math.Tan(radians))

	// More operations can be added here to extend the program
}