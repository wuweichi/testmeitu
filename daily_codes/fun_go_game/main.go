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

type Player struct {
	Name  string
	Score int
}

type Game struct {
	Players       []Player
	CurrentPlayer int
	Round         int
	MaxRounds     int
}

func (g *Game) AddPlayer(name string) {
	g.Players = append(g.Players, Player{Name: name, Score: 0})
}

func (g *Game) NextPlayer() {
	g.CurrentPlayer = (g.CurrentPlayer + 1) % len(g.Players)
}

func (g *Game) PlayRound() {
	player := &g.Players[g.CurrentPlayer]
	fmt.Printf("\nRound %d - %s's turn!\n", g.Round, player.Name)
	
	// Mini-game 1: Number Guessing
	fmt.Println("Mini-game 1: Guess the number between 1 and 10")
	target := rand.Intn(10) + 1
	for i := 0; i < 3; i++ {
		fmt.Print("Enter your guess: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue
		}
		if guess == target {
			fmt.Println("Correct! You earned 10 points.")
			player.Score += 10
			break
		} else if guess < target {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}
	}
	
	// Mini-game 2: Rock Paper Scissors
	fmt.Println("\nMini-game 2: Rock, Paper, Scissors")
	choices := []string{"rock", "paper", "scissors"}
	computerChoice := choices[rand.Intn(3)]
	fmt.Print("Enter your choice (rock/paper/scissors): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	playerChoice := strings.TrimSpace(strings.ToLower(input))
	
	validChoice := false
	for _, c := range choices {
		if playerChoice == c {
			validChoice = true
			break
		}
	}
	
	if !validChoice {
		fmt.Println("Invalid choice! No points awarded.")
	} else {
		fmt.Printf("Computer chose: %s\n", computerChoice)
		if playerChoice == computerChoice {
			fmt.Println("It's a tie! You earned 5 points.")
			player.Score += 5
		} else if (playerChoice == "rock" && computerChoice == "scissors") ||
			(playerChoice == "paper" && computerChoice == "rock") ||
			(playerChoice == "scissors" && computerChoice == "paper") {
			fmt.Println("You win! You earned 15 points.")
			player.Score += 15
		} else {
			fmt.Println("You lose! No points awarded.")
		}
	}
	
	// Mini-game 3: Math Quiz
	fmt.Println("\nMini-game 3: Math Quiz")
	a := rand.Intn(20) + 1
	b := rand.Intn(20) + 1
	operation := rand.Intn(3)
	var answer int
	var question string
	
	switch operation {
	case 0:
		question = fmt.Sprintf("What is %d + %d? ", a, b)
		answer = a + b
	case 1:
		question = fmt.Sprintf("What is %d - %d? ", a, b)
		answer = a - b
	case 2:
		question = fmt.Sprintf("What is %d * %d? ", a, b)
		answer = a * b
	}
	
	fmt.Print(question)
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	userAnswer, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input! No points awarded.")
	} else if userAnswer == answer {
		fmt.Println("Correct! You earned 20 points.")
		player.Score += 20
	} else {
		fmt.Printf("Wrong! The correct answer was %d. No points awarded.\n", answer)
	}
	
	g.NextPlayer()
	g.Round++
}

func (g *Game) DisplayScores() {
	fmt.Println("\nCurrent Scores:")
	for _, player := range g.Players {
		fmt.Printf("%s: %d points\n", player.Name, player.Score)
	}
}

func (g *Game) GetWinner() Player {
	winner := g.Players[0]
	for _, player := range g.Players {
		if player.Score > winner.Score {
			winner = player
		}
	}
	return winner
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
	
	// Create a new game
	game := Game{
		Players:       []Player{},
		CurrentPlayer: 0,
		Round:         1,
		MaxRounds:     3,
	}
	
	// Get player names
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Fun Go Game!")
	fmt.Print("Enter number of players (1-4): ")
	input, _ := reader.ReadString('\n')
	numPlayers, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || numPlayers < 1 || numPlayers > 4 {
		fmt.Println("Invalid number of players! Using default of 2.")
		numPlayers = 2
	}
	
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Enter name for player %d: ", i+1)
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			name = fmt.Sprintf("Player %d", i+1)
		}
		game.AddPlayer(name)
	}
	
	// Play the game
	fmt.Printf("\nStarting game with %d players for %d rounds!\n", len(game.Players), game.MaxRounds)
	
	for game.Round <= game.MaxRounds {
		game.PlayRound()
		game.DisplayScores()
		
		if game.Round <= game.MaxRounds {
			fmt.Print("\nPress Enter to continue to next round...")
			reader.ReadString('\n')
		}
	}
	
	// Determine winner
	winner := game.GetWinner()
	fmt.Printf("\nGame Over! The winner is %s with %d points!\n", winner.Name, winner.Score)
	
	// Ask if players want to play again
	fmt.Print("\nWould you like to play again? (y/n): ")
	input, _ = reader.ReadString('\n')
	if strings.TrimSpace(strings.ToLower(input)) == "y" {
		fmt.Println("\nRestarting game...")
		main()
	} else {
		fmt.Println("Thanks for playing!")
	}
}
