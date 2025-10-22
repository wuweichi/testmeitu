package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type Calculator struct {
	history []string
}

func (c *Calculator) add(a, b float64) float64 {
	result := a + b
	c.history = append(c.history, fmt.Sprintf("%.2f + %.2f = %.2f", a, b, result))
	return result
}

func (c *Calculator) subtract(a, b float64) float64 {
	result := a - b
	c.history = append(c.history, fmt.Sprintf("%.2f - %.2f = %.2f", a, b, result))
	return result
}

func (c *Calculator) multiply(a, b float64) float64 {
	result := a * b
	c.history = append(c.history, fmt.Sprintf("%.2f * %.2f = %.2f", a, b, result))
	return result
}

func (c *Calculator) divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("division by zero")
	}
	result := a / b
	c.history = append(c.history, fmt.Sprintf("%.2f / %.2f = %.2f", a, b, result))
	return result, nil
}

func (c *Calculator) power(a, b float64) float64 {
	result := math.Pow(a, b)
	c.history = append(c.history, fmt.Sprintf("%.2f ^ %.2f = %.2f", a, b, result))
	return result
}

func (c *Calculator) squareRoot(a float64) (float64, error) {
	if a < 0 {
		return 0, fmt.Errorf("square root of negative number")
	}
	result := math.Sqrt(a)
	c.history = append(c.history, fmt.Sprintf("sqrt(%.2f) = %.2f", a, result))
	return result, nil
}

func (c *Calculator) factorial(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("factorial of negative number")
	}
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	c.history = append(c.history, fmt.Sprintf("%d! = %d", n, result))
	return result, nil
}

func (c *Calculator) showHistory() {
	if len(c.history) == 0 {
		fmt.Println("No history available.")
		return
	}
	fmt.Println("Calculation History:")
	for i, entry := range c.history {
		fmt.Printf("%d: %s\n", i+1, entry)
	}
}

func (c *Calculator) clearHistory() {
	c.history = []string{}
	fmt.Println("History cleared.")
}

func getInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

func parseFloat(input string) (float64, error) {
	return strconv.ParseFloat(input, 64)
}

func parseInt(input string) (int, error) {
	return strconv.Atoi(input)
}

func showMenu() {
	fmt.Println("\n=== Interactive Calculator ===")
	fmt.Println("1. Addition")
	fmt.Println("2. Subtraction")
	fmt.Println("3. Multiplication")
	fmt.Println("4. Division")
	fmt.Println("5. Power")
	fmt.Println("6. Square Root")
	fmt.Println("7. Factorial")
	fmt.Println("8. Show History")
	fmt.Println("9. Clear History")
	fmt.Println("0. Exit")
	fmt.Print("Choose an option: ")
}

func main() {
	calc := Calculator{}
	fmt.Println("Welcome to the Interactive Calculator!")
	fmt.Println("This program supports basic arithmetic operations and more.")
	
	for {
		showMenu()
		choice := getInput("")
		
		switch choice {
		case "1":
			fmt.Println("\n--- Addition ---")
			num1Str := getInput("Enter first number: ")
			num2Str := getInput("Enter second number: ")
			
			num1, err1 := parseFloat(num1Str)
			num2, err2 := parseFloat(num2Str)
			
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input. Please enter valid numbers.")
				continue
			}
			
			result := calc.add(num1, num2)
			fmt.Printf("Result: %.2f\n", result)
			
		case "2":
			fmt.Println("\n--- Subtraction ---")
			num1Str := getInput("Enter first number: ")
			num2Str := getInput("Enter second number: ")
			
			num1, err1 := parseFloat(num1Str)
			num2, err2 := parseFloat(num2Str)
			
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input. Please enter valid numbers.")
				continue
			}
			
			result := calc.subtract(num1, num2)
			fmt.Printf("Result: %.2f\n", result)
			
		case "3":
			fmt.Println("\n--- Multiplication ---")
			num1Str := getInput("Enter first number: ")
			num2Str := getInput("Enter second number: ")
			
			num1, err1 := parseFloat(num1Str)
			num2, err2 := parseFloat(num2Str)
			
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input. Please enter valid numbers.")
				continue
			}
			
			result := calc.multiply(num1, num2)
			fmt.Printf("Result: %.2f\n", result)
			
		case "4":
			fmt.Println("\n--- Division ---")
			num1Str := getInput("Enter dividend: ")
			num2Str := getInput("Enter divisor: ")
			
			num1, err1 := parseFloat(num1Str)
			num2, err2 := parseFloat(num2Str)
			
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input. Please enter valid numbers.")
				continue
			}
			
			result, err := calc.divide(num1, num2)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			
			fmt.Printf("Result: %.2f\n", result)
			
		case "5":
			fmt.Println("\n--- Power ---")
			baseStr := getInput("Enter base: ")
			exponentStr := getInput("Enter exponent: ")
			
			base, err1 := parseFloat(baseStr)
			exponent, err2 := parseFloat(exponentStr)
			
			if err1 != nil || err2 != nil {
				fmt.Println("Invalid input. Please enter valid numbers.")
				continue
			}
			
			result := calc.power(base, exponent)
			fmt.Printf("Result: %.2f\n", result)
			
		case "6":
			fmt.Println("\n--- Square Root ---")
			numStr := getInput("Enter number: ")
			
			num, err := parseFloat(numStr)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid number.")
				continue
			}
			
			result, err := calc.squareRoot(num)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			
			fmt.Printf("Result: %.2f\n", result)
			
		case "7":
			fmt.Println("\n--- Factorial ---")
			numStr := getInput("Enter a non-negative integer: ")
			
			num, err := parseInt(numStr)
			if err != nil {
				fmt.Println("Invalid input. Please enter a valid integer.")
				continue
			}
			
			result, err := calc.factorial(num)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			
			fmt.Printf("Result: %d\n", result)
			
		case "8":
			fmt.Println("\n--- History ---")
			calc.showHistory()
			
		case "9":
			fmt.Println("\n--- Clear History ---")
			calc.clearHistory()
			
		case "0":
			fmt.Println("\nThank you for using the Interactive Calculator!")
			fmt.Println("Goodbye!")
			time.Sleep(1 * time.Second)
			return
			
		default:
			fmt.Println("Invalid option. Please choose a valid menu item.")
		}
	}
}
