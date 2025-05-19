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

// Abs returns the absolute value of a complex number.
func (c ComplexNumber) Abs() float64 {
	return math.Sqrt(c.Real*c.Real + c.Imaginary*c.Imaginary)
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
	fmt.Printf("a = %v + %vi\n", a.Real, a.Imaginary)
	fmt.Printf("b = %v + %vi\n", b.Real, b.Imaginary)
	fmt.Printf("a + b = %v + %vi\n", sum.Real, sum.Imaginary)
	fmt.Printf("a - b = %v + %vi\n", difference.Real, difference.Imaginary)
	fmt.Printf("a * b = %v + %vi\n", product.Real, product.Imaginary)
	fmt.Printf("a / b = %v + %vi\n", quotient.Real, quotient.Imaginary)
	fmt.Printf("|a| = %v\n", a.Abs())
	fmt.Printf("|b| = %v\n", b.Abs())
}
