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

	// Perform arithmetic operations
	sum := a + b
	difference := a - b
	product := a * b
	quotient := float64(a) / float64(b)

	// Perform some trigonometric operations
	angle := float64(a) * math.Pi / 180
	sin := math.Sin(angle)
	cos := math.Cos(angle)
	tan := math.Tan(angle)

	// Output the results
	fmt.Printf("Random numbers: %d and %d\n", a, b)
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Difference: %d\n", difference)
	fmt.Printf("Product: %d\n", product)
	fmt.Printf("Quotient: %.2f\n", quotient)
	fmt.Printf("Sin of %d degrees: %.2f\n", a, sin)
	fmt.Printf("Cos of %d degrees: %.2f\n", a, cos)
	fmt.Printf("Tan of %d degrees: %.2f\n", a, tan)

	// Generate and print a random string
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	randomString := make([]rune, 10)
	for i := range randomString {
		randomString[i] = letters[rand.Intn(len(letters))]
	}
	fmt.Printf("Random string: %s\n", string(randomString))

	// A simple loop to print numbers
	fmt.Println("Printing numbers from 1 to 10:")
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	// A simple if-else example
	if a > b {
		fmt.Printf("%d is greater than %d\n", a, b)
	} else if a < b {
		fmt.Printf("%d is less than %d\n", a, b)
	} else {
		fmt.Printf("%d is equal to %d\n", a, b)
	}

	// A simple switch example
	switch {
	case a > 50:
		fmt.Println("First number is greater than 50")
	case b > 50:
		fmt.Println("Second number is greater than 50")
	default:
		fmt.Println("Both numbers are 50 or less")
	}

	// A simple function call
	fmt.Printf("Factorial of 5: %d\n", factorial(5))


	// A simple struct example
	type Person struct {
		Name string
		Age  int
	}
	person := Person{Name: "John Doe", Age: 30}
	fmt.Printf("Person: %+v\n", person)

	// A simple slice example
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("Numbers slice: %v\n", numbers)

	// A simple map example
	ages := map[string]int{
		"John": 30,
		"Jane": 25,
	}
	fmt.Printf("Ages map: %v\n", ages)


	// A simple goroutine example
	go func() {
		fmt.Println("This is printed from a goroutine")
	}()
	time.Sleep(time.Second) // Wait for the goroutine to finish

	// A simple channel example
	messages := make(chan string)
	go func() { messages <- "ping" }()
	msg := <-messages
	fmt.Println(msg)
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
