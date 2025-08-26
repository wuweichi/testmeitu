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

// Game interface defines the common methods for all games
type Game interface {
	Play()
	GetName() string
}

// NumberGuessingGame struct for a simple number guessing game
type NumberGuessingGame struct {
	Name string
}

func (g NumberGuessingGame) GetName() string {
	return g.Name
}

func (g NumberGuessingGame) Play() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 100. Can you guess it?")
	
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	attempts := 0
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("Enter your guess: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}
		attempts++
		if guess < target {
			fmt.Println("Too low! Try again.")
		} else if guess > target {
			fmt.Println("Too high! Try again.")
		} else {
			fmt.Printf("Congratulations! You guessed the number in %d attempts.\n", attempts)
			break
		}
	}
}

// RockPaperScissorsGame struct for a rock-paper-scissors game
type RockPaperScissorsGame struct {
	Name string
}

func (g RockPaperScissorsGame) GetName() string {
	return g.Name
}

func (g RockPaperScissorsGame) Play() {
	fmt.Println("Welcome to Rock-Paper-Scissors!")
	fmt.Println("Enter your choice (rock, paper, scissors):")
	
	choices := []string{"rock", "paper", "scissors"}
	rand.Seed(time.Now().UnixNano())
	computerChoice := choices[rand.Intn(len(choices))]
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("Your choice: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		playerChoice := strings.TrimSpace(strings.ToLower(input))
		if playerChoice != "rock" && playerChoice != "paper" && playerChoice != "scissors" {
			fmt.Println("Invalid choice. Please enter rock, paper, or scissors.")
			continue
		}
		fmt.Printf("Computer chose: %s\n", computerChoice)
		if playerChoice == computerChoice {
			fmt.Println("It's a tie!")
		} else if (playerChoice == "rock" && computerChoice == "scissors") ||
			(playerChoice == "paper" && computerChoice == "rock") ||
			(playerChoice == "scissors" && computerChoice == "paper") {
			fmt.Println("You win!")
		} else {
			fmt.Println("Computer wins!")
		}
		break
	}
}

// SimpleCalculator struct for a basic calculator
type SimpleCalculator struct {
	Name string
}

func (g SimpleCalculator) GetName() string {
	return g.Name
}

func (g SimpleCalculator) Play() {
	fmt.Println("Welcome to the Simple Calculator!")
	fmt.Println("Enter two numbers and an operator (+, -, *, /):")
	reader := bufio.NewReader(os.Stdin)
	
	fmt.Print("Enter first number: ")
	num1Str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	num1Str = strings.TrimSpace(num1Str)
	num1, err := strconv.ParseFloat(num1Str, 64)
	if err != nil {
		fmt.Println("Invalid number.")
		return
	}
	
	fmt.Print("Enter operator: ")
	op, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	op = strings.TrimSpace(op)
	
	fmt.Print("Enter second number: ")
	num2Str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input.")
		return
	}
	num2Str = strings.TrimSpace(num2Str)
	num2, err := strconv.ParseFloat(num2Str, 64)
	if err != nil {
		fmt.Println("Invalid number.")
		return
	}
	
	var result float64
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero.")
			return
		}
		result = num1 / num2
	default:
		fmt.Println("Invalid operator.")
		return
	}
	fmt.Printf("Result: %.2f\n", result)
}

// QuizGame struct for a simple quiz
type QuizGame struct {
	Name string
}

func (g QuizGame) GetName() string {
	return g.Name
}

func (g QuizGame) Play() {
	fmt.Println("Welcome to the Quiz Game!")
	questions := []struct {
		question string
		answer   string
	}{
		{"What is the capital of France?", "paris"},
		{"What is 2 + 2?", "4"},
		{"What is the largest planet in our solar system?", "jupiter"},
	}
	
	score := 0
	reader := bufio.NewReader(os.Stdin)
	for i, q := range questions {
		fmt.Printf("Question %d: %s\n", i+1, q.question)
		fmt.Print("Your answer: ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}
		userAnswer := strings.TrimSpace(strings.ToLower(input))
		if userAnswer == q.answer {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Wrong! The correct answer is %s.\n", q.answer)
		}
	}
	fmt.Printf("Your score: %d out of %d\n", score, len(questions))
}

// DiceRollingGame struct for rolling dice
type DiceRollingGame struct {
	Name string
}

func (g DiceRollingGame) GetName() string {
	return g.Name
}

func (g DiceRollingGame) Play() {
	fmt.Println("Welcome to Dice Rolling!")
	rand.Seed(time.Now().UnixNano())
	roll := rand.Intn(6) + 1
	fmt.Printf("You rolled a %d!\n", roll)
}

// Function to display menu and handle user input
func showMenu(games []Game) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nChoose a game to play:")
		for i, game := range games {
			fmt.Printf("%d. %s\n", i+1, game.GetName())
		}
		fmt.Println("0. Exit")
		fmt.Print("Enter your choice: ")
		
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input, please try again.")
			continue
		}
		choiceStr := strings.TrimSpace(input)
		choice, err := strconv.Atoi(choiceStr)
		if err != nil {
			fmt.Println("Please enter a valid number.")
			continue
		}
		
		if choice == 0 {
			fmt.Println("Goodbye!")
			return
		}
		if choice < 1 || choice > len(games) {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}
		selectedGame := games[choice-1]
		fmt.Printf("\nStarting %s...\n", selectedGame.GetName())
		selectedGame.Play()
	}
}

func main() {
	// Initialize games
	games := []Game{
		NumberGuessingGame{Name: "Number Guessing"},
		RockPaperScissorsGame{Name: "Rock Paper Scissors"},
		SimpleCalculator{Name: "Simple Calculator"},
		QuizGame{Name: "Quiz Game"},
		DiceRollingGame{Name: "Dice Rolling"},
	}
	
	fmt.Println("Welcome to the Interactive Game Suite!")
	showMenu(games)
}
