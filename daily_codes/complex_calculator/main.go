package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func multiply(a, b float64) float64 {
	return a * b
}

func divide(a, b float64) float64 {
	if b == 0 {
		fmt.Println("Error: Division by zero")
		os.Exit(1)
	}
	return a / b
}

func power(a, b float64) float64 {
	return math.Pow(a, b)
}

func sqrt(a float64) float64 {
	if a < 0 {
		fmt.Println("Error: Square root of negative number")
		os.Exit(1)
	}
	return math.Sqrt(a)
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: calculator <operation> <operand1> <operand2>")
		fmt.Println("Operations: add, subtract, multiply, divide, power, sqrt")
		os.Exit(1)
	}

	operation := os.Args[1]
	a, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Error parsing operand1:", err)
		os.Exit(1)
	}

	var b float64
	if operation != "sqrt" {
		b, err = strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Println("Error parsing operand2:", err)
			os.Exit(1)
		}
	}

	var result float64
	switch operation {
	case "add":
		result = add(a, b)
	case "subtract":
		result = subtract(a, b)
	case "multiply":
		result = multiply(a, b)
	case "divide":
		result = divide(a, b)
	case "power":
		result = power(a, b)
	case "sqrt":
		result = sqrt(a)
	default:
		fmt.Println("Unknown operation:", operation)
		os.Exit(1)
	}

	fmt.Println("Result:", result)
}
