package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

// ComplexNumber represents a complex number with real and imaginary parts.
type ComplexNumber struct {
	Real float64
	Imaginary float64
}

// Add adds two complex numbers.
func (c ComplexNumber) Add(other ComplexNumber) ComplexNumber {
	return ComplexNumber{
		Real: c.Real + other.Real,
		Imaginary: c.Imaginary + other.Imaginary,
	}
}

// Subtract subtracts two complex numbers.
func (c ComplexNumber) Subtract(other ComplexNumber) ComplexNumber {
	return ComplexNumber{
		Real: c.Real - other.Real,
		Imaginary: c.Imaginary - other.Imaginary,
	}
}

// Multiply multiplies two complex numbers.
func (c ComplexNumber) Multiply(other ComplexNumber) ComplexNumber {
	return ComplexNumber{
		Real: c.Real*other.Real - c.Imaginary*other.Imaginary,
		Imaginary: c.Real*other.Imaginary + c.Imaginary*other.Real,
	}
}

// Divide divides two complex numbers.
func (c ComplexNumber) Divide(other ComplexNumber) ComplexNumber {
	denominator := other.Real*other.Real + other.Imaginary*other.Imaginary
	return ComplexNumber{
		Real: (c.Real*other.Real + c.Imaginary*other.Imaginary) / denominator,
		Imaginary: (c.Imaginary*other.Real - c.Real*other.Imaginary) / denominator,
	}
}

// Magnitude returns the magnitude of the complex number.
func (c ComplexNumber) Magnitude() float64 {
	return math.Sqrt(c.Real*c.Real + c.Imaginary*c.Imaginary)
}

// Phase returns the phase of the complex number in radians.
func (c ComplexNumber) Phase() float64 {
	return math.Atan2(c.Imaginary, c.Real)
}

// GenerateRandomComplexNumber generates a random complex number.
func GenerateRandomComplexNumber() ComplexNumber {
	rand.Seed(time.Now().UnixNano())
	return ComplexNumber{
		Real: rand.Float64() * 100,
		Imaginary: rand.Float64() * 100,
	}
}

func main() {
	// Generate two random complex numbers
	a := GenerateRandomComplexNumber()
	b := GenerateRandomComplexNumber()

	// Perform operations
	sum := a.Add(b)
	difference := a.Subtract(b)
	product := a.Multiply(b)
	quotient := a.Divide(b)

	// Print results
	fmt.Printf("Complex Number A: %v\n", a)
	fmt.Printf("Complex Number B: %v\n", b)
	fmt.Printf("Sum: %v\n", sum)
	fmt.Printf("Difference: %v\n", difference)
	fmt.Printf("Product: %v\n", product)
	fmt.Printf("Quotient: %v\n", quotient)
	fmt.Printf("Magnitude of A: %v\n", a.Magnitude())
	fmt.Printf("Phase of A: %v radians\n", a.Phase())
}
