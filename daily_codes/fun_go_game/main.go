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
	Players    []Player
	Round      int
	MaxRounds  int
	Difficulty string
}

func (g *Game) Start() {
	fmt.Println("Welcome to the Fun Go Game!")
	g.setupPlayers()
	g.setupGame()
	g.playRounds()
	g.endGame()
}

func (g *Game) setupPlayers() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of players: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numPlayers, err := strconv.Atoi(input)
	if err != nil || numPlayers < 1 {
		numPlayers = 1
	}

	g.Players = make([]Player, numPlayers)
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Enter name for player %d: ", i+1)
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			name = fmt.Sprintf("Player%d", i+1)
		}
		g.Players[i] = Player{Name: name, Score: 0}
	}
}

func (g *Game) setupGame() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of rounds: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	maxRounds, err := strconv.Atoi(input)
	if err != nil || maxRounds < 1 {
		maxRounds = 5
	}
	g.MaxRounds = maxRounds

	fmt.Print("Choose difficulty (easy/medium/hard): ")
	diff, _ := reader.ReadString('\n')
	diff = strings.TrimSpace(diff)
	g.Difficulty = diff
}

func (g *Game) playRounds() {
	for g.Round = 1; g.Round <= g.MaxRounds; g.Round++ {
		fmt.Printf("\n--- Round %d ---\n", g.Round)
		g.playRound()
	}
}

func (g *Game) playRound() {
	for i := range g.Players {
		g.playTurn(&g.Players[i])
	}
}

func (g *Game) playTurn(player *Player) {
	fmt.Printf("\n%s's turn!\n", player.Name)
	
	// Different mini-games based on difficulty
	switch g.Difficulty {
	case "easy":
		g.easyGame(player)
	case "medium":
		g.mediumGame(player)
	case "hard":
		g.hardGame(player)
	default:
		g.easyGame(player)
	}
}

func (g *Game) easyGame(player *Player) {
	fmt.Println("Guess the number between 1-10!")
	target := rand.Intn(10) + 1
	
	reader := bufio.NewReader(os.Stdin)
	for attempts := 3; attempts > 0; attempts-- {
		fmt.Printf("Attempts left: %d. Enter your guess: ", attempts)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		
		if err != nil {
			fmt.Println("Invalid input! Please enter a number.")
			continue
		}
		
		if guess == target {
			fmt.Println("Correct! You earned 10 points!")
			player.Score += 10
			return
		} else if guess < target {
			fmt.Println("Too low!")
		} else {
			fmt.Println("Too high!")
		}
	}
	fmt.Printf("Out of attempts! The number was %d.\n", target)
}

func (g *Game) mediumGame(player *Player) {
	fmt.Println("Math challenge! Solve the equation.")
	
	// Generate random math problem
	a := rand.Intn(20) + 1
	b := rand.Intn(20) + 1
	ops := []string{"+", "-", "*"}
	op := ops[rand.Intn(len(ops))]
	
	var answer int
	var problem string
	
	switch op {
	case "+":
		answer = a + b
		problem = fmt.Sprintf("%d + %d", a, b)
	case "-":
		answer = a - b
		problem = fmt.Sprintf("%d - %d", a, b)
	case "*":
		answer = a * b
		problem = fmt.Sprintf("%d * %d", a, b)
	}
	
	fmt.Printf("What is %s? ", problem)
	
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	userAnswer, err := strconv.Atoi(input)
	
	if err != nil {
		fmt.Println("Invalid input! No points earned.")
		return
	}
	
	if userAnswer == answer {
		fmt.Println("Correct! You earned 15 points!")
		player.Score += 15
	} else {
		fmt.Printf("Wrong! The answer was %d.\n", answer)
	}
}

func (g *Game) hardGame(player *Player) {
	fmt.Println("Memory game! Remember the sequence.")
	
	sequenceLength := 4 + g.Round
	sequence := make([]int, sequenceLength)
	
	// Generate random sequence
	for i := 0; i < sequenceLength; i++ {
		sequence[i] = rand.Intn(9) + 1
	}
	
	// Show sequence briefly
	fmt.Printf("Remember this sequence: %v\n", sequence)
	time.Sleep(3 * time.Second)
	
	// Clear screen (simplified)
	for i := 0; i < 50; i++ {
		fmt.Println()
	}
	
	// Get user input
	fmt.Print("Enter the sequence (space separated): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	parts := strings.Split(input, " ")
	if len(parts) != sequenceLength {
		fmt.Printf("Wrong length! Expected %d numbers.\n", sequenceLength)
		return
	}
	
	// Check sequence
	correct := true
	for i, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil || num != sequence[i] {
			correct = false
			break
		}
	}
	
	if correct {
		points := 20 + g.Round*5
		fmt.Printf("Perfect! You earned %d points!\n", points)
		player.Score += points
	} else {
		fmt.Printf("Incorrect! The sequence was %v\n", sequence)
	}
}

func (g *Game) endGame() {
	fmt.Println("\n=== Game Over ===")
	fmt.Println("Final Scores:")
	
	// Sort players by score
	for i := 0; i < len(g.Players)-1; i++ {
		for j := i + 1; j < len(g.Players); j++ {
			if g.Players[j].Score > g.Players[i].Score {
				g.Players[i], g.Players[j] = g.Players[j], g.Players[i]
			}
		}
	}
	
	for i, player := range g.Players {
		fmt.Printf("%d. %s: %d points\n", i+1, player.Name, player.Score)
	}
	
	if len(g.Players) > 1 {
		fmt.Printf("\nCongratulations %s, you are the winner!\n", g.Players[0].Name)
	}
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
	
	game := Game{}
	game.Start()
	
	// Ask if user wants to play again
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\nPlay again? (y/n): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))
		
		if input == "y" || input == "yes" {
			game = Game{}
			game.Start()
		} else {
			fmt.Println("Thanks for playing!")
			break
		}
	}
}
