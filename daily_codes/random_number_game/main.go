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

// GameState holds the state of the game
type GameState struct {
	TargetNumber int
	Attempts     int
	MaxAttempts  int
	Score        int
	PlayerName   string
}

// NewGameState initializes a new game state
func NewGameState(playerName string, maxAttempts int) *GameState {
	rand.Seed(time.Now().UnixNano())
	return &GameState{
		TargetNumber: rand.Intn(100) + 1,
		Attempts:     0,
		MaxAttempts:  maxAttempts,
		Score:        0,
		PlayerName:   playerName,
	}
}

// PlayRound processes a single guess
func (g *GameState) PlayRound(guess int) (bool, string) {
	g.Attempts++
	if guess == g.TargetNumber {
		g.Score += (g.MaxAttempts - g.Attempts + 1) * 10
		return true, fmt.Sprintf("Congratulations, %s! You guessed the number %d in %d attempts. Your score is %d.", g.PlayerName, g.TargetNumber, g.Attempts, g.Score)
	} else if guess < g.TargetNumber {
		return false, "Too low! Try again."
	} else {
		return false, "Too high! Try again."
	}
}

// IsGameOver checks if the game is over
func (g *GameState) IsGameOver() bool {
	return g.Attempts >= g.MaxAttempts
}

// GetHint provides a hint based on attempts
func (g *GameState) GetHint() string {
	if g.Attempts == 0 {
		return "No hints yet. Make a guess!"
	}
	// Simple hint logic: after 3 attempts, give a range hint
	if g.Attempts >= 3 {
		lower := g.TargetNumber - 10
		if lower < 1 {
			lower = 1
		}
		upper := g.TargetNumber + 10
		if upper > 100 {
			upper = 100
		}
		return fmt.Sprintf("Hint: The number is between %d and %d.", lower, upper)
	}
	return "Keep guessing! Hints available after more attempts."
}

// SaveScore saves the score to a file
func SaveScore(playerName string, score int) error {
	file, err := os.OpenFile("scores.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(fmt.Sprintf("%s: %d\n", playerName, score))
	return err
}

// LoadScores loads scores from file
func LoadScores() ([]string, error) {
	file, err := os.Open("scores.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return []string{"No scores yet."}, nil
		}
		return nil, err
	}
	defer file.Close()
	var scores []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scores = append(scores, scanner.Text())
	}
	return scores, scanner.Err()
}

// DisplayMenu shows the main menu
func DisplayMenu() {
	fmt.Println("\n=== Random Number Game ===")
	fmt.Println("1. Start New Game")
	fmt.Println("2. View High Scores")
	fmt.Println("3. Exit")
	fmt.Print("Choose an option: ")
}

// GetPlayerInput gets input from the player
func GetPlayerInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	input, _ := reader.ReadString('\n')
	return strings.TrimSpace(input)
}

// ValidateInput validates numeric input
func ValidateInput(input string, min, max int) (int, error) {
	num, err := strconv.Atoi(input)
	if err != nil {
		return 0, fmt.Errorf("please enter a valid number")
	}
	if num < min || num > max {
		return 0, fmt.Errorf("number must be between %d and %d", min, max)
	}
	return num, nil
}

// Main game function
func main() {
	for {
		DisplayMenu()
		choice := GetPlayerInput("")
		switch choice {
		case "1":
			playerName := GetPlayerInput("Enter your name: ")
			if playerName == "" {
				playerName = "Player"
			}
			maxAttempts := 10
			game := NewGameState(playerName, maxAttempts)
			fmt.Printf("\nWelcome, %s! I'm thinking of a number between 1 and 100. You have %d attempts.\n", playerName, maxAttempts)
			for !game.IsGameOver() {
				input := GetPlayerInput("Enter your guess: ")
				guess, err := ValidateInput(input, 1, 100)
				if err != nil {
					fmt.Println("Error:", err)
					continue
				}
				won, message := game.PlayRound(guess)
				fmt.Println(message)
				if won {
					break
				}
				if game.Attempts%2 == 0 {
					hint := game.GetHint()
					fmt.Println(hint)
				}
				fmt.Printf("Attempts left: %d\n", maxAttempts-game.Attempts)
			}
			if game.IsGameOver() && game.Score == 0 {
				fmt.Printf("Game over! The number was %d. Better luck next time!\n", game.TargetNumber)
			}
			err := SaveScore(playerName, game.Score)
			if err != nil {
				fmt.Println("Error saving score:", err)
			} else {
				fmt.Println("Score saved successfully.")
			}
		case "2":
			scores, err := LoadScores()
			if err != nil {
				fmt.Println("Error loading scores:", err)
			} else {
				fmt.Println("\n=== High Scores ===")
				for _, score := range scores {
					fmt.Println(score)
				}
			}
		case "3":
			fmt.Println("Thanks for playing! Goodbye.")
			return
		default:
			fmt.Println("Invalid option. Please choose 1, 2, or 3.")
		}
	}
}
