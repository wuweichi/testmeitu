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
	fmt.Printf("Logarithm of %d: %.2f\n", a, math.Log(float64(a)))
	fmt.Printf("Logarithm of %d: %.2f\n", b, math.Log(float64(b)))

	// Generate and print a random number between 0 and 1
	fmt.Printf("Random number between 0 and 1: %.2f\n", rand.Float64())

	// Calculate and print the factorial of a
	factorial := 1
	for i := 1; i <= a; i++ {
		factorial *= i
	}
	fmt.Printf("Factorial of %d: %d\n", a, factorial)

	// Check if a is prime
	isPrime := true
	if a < 2 {
		isPrime = false
	} else {
		for i := 2; i <= int(math.Sqrt(float64(a))); i++ {
			if a%i == 0 {
				isPrime = false
				break
			}
		}
	}
	fmt.Printf("Is %d a prime number? %t\n", a, isPrime)

	// Generate a random string
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomString := make([]rune, 10)
	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Printf("Random string: %s\n", string(randomString))

	// Infinite loop to keep the program running (simulate a long program)
	for {
		fmt.Println("Program is running... Press CTRL+C to exit.")
		time.Sleep(10 * time.Second)
	}
}
