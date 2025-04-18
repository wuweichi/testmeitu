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

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	return a / b, nil
}

func power(a, b float64) float64 {
	return math.Pow(a, b)
}

func sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("cannot take square root of negative number")
	}
	return math.Sqrt(a), nil
}

func main() {
	if len(os.Args) < 4 {
		fmt.Println("Usage: calculator <operation> <operand1> <operand2>")
		fmt.Println("Operations: add, subtract, multiply, divide, power, sqrt")
		os.Exit(1)
	}

	operation := os.Args[1]
	operand1, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Invalid operand1")
		os.Exit(1)
	}
	var operand2 float64
	if operation != "sqrt" {
		operand2, err = strconv.ParseFloat(os.Args[3], 64)
		if err != nil {
			fmt.Println("Invalid operand2")
			os.Exit(1)
		}
	}

	var result float64
	switch operation {
	case "add":
		result = add(operand1, operand2)
	case "subtract":
		result = subtract(operand1, operand2)
	case "multiply":
		result = multiply(operand1, operand2)
	case "divide":
		result, err = divide(operand1, operand2)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	case "power":
		result = power(operand1, operand2)
	case "sqrt":
		result, err = sqrt(operand1)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid operation")
		os.Exit(1)
	}

	fmt.Printf("Result: %v\n", result)
}
