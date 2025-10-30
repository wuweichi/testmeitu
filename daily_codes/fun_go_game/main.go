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
	Players      []Player
	CurrentRound int
	MaxRounds    int
}

func (g *Game) AddPlayer(name string) {
	g.Players = append(g.Players, Player{Name: name, Score: 0})
}

func (g *Game) StartGame() {
	fmt.Println("Welcome to the Fun Go Game!")
	g.SetupPlayers()
	g.MaxRounds = 5
	for g.CurrentRound < g.MaxRounds {
		g.PlayRound()
		g.CurrentRound++
	}
	g.EndGame()
}

func (g *Game) SetupPlayers() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter number of players: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	numPlayers, err := strconv.Atoi(input)
	if err != nil || numPlayers < 1 {
		numPlayers = 2
	}
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Enter name for player %d: ", i+1)
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		if name == "" {
			name = fmt.Sprintf("Player%d", i+1)
		}
		g.AddPlayer(name)
	}
}

func (g *Game) PlayRound() {
	fmt.Printf("\n--- Round %d ---\n", g.CurrentRound+1)
	for i := range g.Players {
		g.PlayTurn(&g.Players[i])
	}
	g.DisplayScores()
}

func (g *Game) PlayTurn(player *Player) {
	fmt.Printf("\n%s's turn!\n", player.Name)
	fmt.Println("Guess a number between 1 and 10:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. You lose this turn.")
		return
	}
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(10) + 1
	if guess == target {
		points := 10
		player.Score += points
		fmt.Printf("Correct! You earned %d points.\n", points)
	} else {
		fmt.Printf("Wrong! The number was %d.\n", target)
	}
}

func (g *Game) DisplayScores() {
	fmt.Println("\nCurrent Scores:")
	for _, player := range g.Players {
		fmt.Printf("%s: %d\n", player.Name, player.Score)
	}
}

func (g *Game) EndGame() {
	fmt.Println("\n--- Game Over ---")
	g.DisplayScores()
	winner := g.DetermineWinner()
	if winner != nil {
		fmt.Printf("Congratulations %s, you win!\n", winner.Name)
	} else {
		fmt.Println("It's a tie!")
	}
}

func (g *Game) DetermineWinner() *Player {
	if len(g.Players) == 0 {
		return nil
	}
	winner := &g.Players[0]
	for i := 1; i < len(g.Players); i++ {
		if g.Players[i].Score > winner.Score {
			winner = &g.Players[i]
		}
	}
	// Check for ties
	for _, player := range g.Players {
		if player.Score == winner.Score && player.Name != winner.Name {
			return nil
		}
	}
	return winner
}

func main() {
	game := Game{}
	game.StartGame()
}
