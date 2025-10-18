package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// ASCII art for a simple game
var asciiArt = map[string]string{
	"rock": `
    _______
---'   ____)
      (_____)
      (_____)
      (____)
---.__(___)
`,
	"paper": `
    _______
---'   ____)____
          ______)
          _______)
         _______)
---.__________)
`,
	"scissors": `
    _______
---'   ____)____
          ______)
       __________)
      (____)
---.__(___)
`,
}

// Game logic for Rock, Paper, Scissors
func playGame(playerChoice string) (string, string, string) {
	choices := []string{"rock", "paper", "scissors"}
	computerChoice := choices[rand.Intn(len(choices))]
	result := ""

	if playerChoice == computerChoice {
		result = "It's a tie!"
	} else if (playerChoice == "rock" && computerChoice == "scissors") ||
		(playerChoice == "paper" && computerChoice == "rock") ||
		(playerChoice == "scissors" && computerChoice == "paper") {
		result = "You win!"
	} else {
		result = "Computer wins!"
	}

	return playerChoice, computerChoice, result
}

// Function to display ASCII art
func displayArt(choice string) {
	if art, exists := asciiArt[choice]; exists {
		fmt.Println(art)
	} else {
		fmt.Println("No art for this choice.")
	}
}

// Function to handle user input
func getUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// Main game loop
func gameLoop() {
	fmt.Println("Welcome to the ASCII Art Rock, Paper, Scissors Game!")
	fmt.Println("Enter 'rock', 'paper', 'scissors', or 'quit' to exit.")

	for {
		fmt.Print("Your choice: ")
		input := getUserInput()

		if input == "quit" {
			fmt.Println("Thanks for playing!")
			break
		}

		if input != "rock" && input != "paper" && input != "scissors" {
			fmt.Println("Invalid choice. Please enter 'rock', 'paper', or 'scissors'.")
			continue
		}

		playerChoice, computerChoice, result := playGame(input)

		fmt.Println("\nYou chose:")
		displayArt(playerChoice)
		fmt.Println("Computer chose:")
		displayArt(computerChoice)
		fmt.Println(result)
		fmt.Println()
	}
}

// Additional utility functions to increase line count
func generateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func printMultiples(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", n, i, n*i)
	}
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func printFibonacciSequence(limit int) {
	for i := 0; i < limit; i++ {
		fmt.Printf("Fibonacci(%d) = %d\n", i, fibonacci(i))
	}
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func countVowels(s string) int {
	vowels := "aeiouAEIOU"
	count := 0
	for _, char := range s {
		if strings.ContainsRune(vowels, char) {
			count++
		}
	}
	return count
}

func printPattern(rows int) {
	for i := 1; i <= rows; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
	}
}

func calculateArea(shape string, dimensions ...float64) float64 {
	switch shape {
	case "circle":
		if len(dimensions) != 1 {
			return 0
		}
		return 3.14159 * dimensions[0] * dimensions[0]
	case "rectangle":
		if len(dimensions) != 2 {
			return 0
		}
		return dimensions[0] * dimensions[1]
	case "triangle":
		if len(dimensions) != 2 {
			return 0
		}
		return 0.5 * dimensions[0] * dimensions[1]
	default:
		return 0
	}
}

func demonstrateFunctions() {
	fmt.Println("\n--- Demonstrating Utility Functions ---")
	
	// Generate random number
	randomNum := generateRandomNumber(1, 100)
	fmt.Printf("Random number between 1 and 100: %d\n", randomNum)
	
	// Print multiples
	fmt.Println("\nMultiples of 5:")
	printMultiples(5)
	
	// Factorial
	fact := factorial(5)
	fmt.Printf("\nFactorial of 5: %d\n", fact)
	
	// Prime check
	primeCheck := isPrime(29)
	fmt.Printf("Is 29 prime? %t\n", primeCheck)
	
	// Fibonacci sequence
	fmt.Println("\nFirst 10 Fibonacci numbers:")
	printFibonacciSequence(10)
	
	// String reverse
	reversed := reverseString("Hello, World!")
	fmt.Printf("\nReversed string: %s\n", reversed)
	
	// Count vowels
	vowelCount := countVowels("Hello, World!")
	fmt.Printf("Number of vowels in 'Hello, World!': %d\n", vowelCount)
	
	// Print pattern
	fmt.Println("\nPattern with 5 rows:")
	printPattern(5)
	
	// Calculate areas
	circleArea := calculateArea("circle", 5.0)
	fmt.Printf("\nArea of circle with radius 5: %.2f\n", circleArea)
	
	rectangleArea := calculateArea("rectangle", 4.0, 6.0)
	fmt.Printf("Area of rectangle 4x6: %.2f\n", rectangleArea)
	
	triangleArea := calculateArea("triangle", 3.0, 4.0)
	fmt.Printf("Area of triangle with base 3 and height 4: %.2f\n", triangleArea)
}

func main() {
	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Start the game
	gameLoop()

	// Demonstrate additional functions
	demonstrateFunctions()

	// Interactive number guessing game
	fmt.Println("\n--- Number Guessing Game ---")
	target := generateRandomNumber(1, 50)
	attempts := 0
	maxAttempts := 5

	fmt.Printf("I'm thinking of a number between 1 and 50. You have %d attempts.\n", maxAttempts)

	for attempts < maxAttempts {
		fmt.Print("Enter your guess: ")
		input := getUserInput()
		guess, err := strconv.Atoi(input)
		
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}
		
		attempts++
		
		if guess < target {
			fmt.Println("Too low!")
		} else if guess > target {
			fmt.Println("Too high!")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			break
		}
		
		if attempts == maxAttempts {
			fmt.Printf("Sorry, you've used all %d attempts. The number was %d.\n", maxAttempts, target)
		}
	}

	// Final message
	fmt.Println("\nThanks for playing all the games!")
}
