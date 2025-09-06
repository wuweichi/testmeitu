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
	Symbol string
	Score  int
}

type Game struct {
	Board      [3][3]string
	Players    [2]Player
	Current    int
	GameOver   bool
	Winner     *Player
}

func NewGame() *Game {
	g := &Game{
		Players: [2]Player{
			{Symbol: "X", Score: 0},
			{Symbol: "O", Score: 0},
		},
		Current:  0,
		GameOver: false,
		Winner:   nil,
	}
	g.ResetBoard()
	return g
}

func (g *Game) ResetBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			g.Board[i][j] = " "
		}
	}
	g.GameOver = false
	g.Winner = nil
}

func (g *Game) PrintBoard() {
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d ", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("%s", g.Board[i][j])
			if j < 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 2 {
			fmt.Println("  -----")
		}
	}
}

func (g *Game) IsValidMove(row, col int) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 {
		return false
	}
	return g.Board[row][col] == " "
}

func (g *Game) MakeMove(row, col int) bool {
	if !g.IsValidMove(row, col) {
		return false
	}
	g.Board[row][col] = g.Players[g.Current].Symbol
	g.CheckGameOver()
	if !g.GameOver {
		g.Current = 1 - g.Current
	}
	return true
}

func (g *Game) CheckGameOver() {
	// Check rows
	for i := 0; i < 3; i++ {
		if g.Board[i][0] != " " && g.Board[i][0] == g.Board[i][1] && g.Board[i][1] == g.Board[i][2] {
			g.GameOver = true
			g.Winner = &g.Players[g.Current]
			g.Winner.Score++
			return
		}
	}
	// Check columns
	for j := 0; j < 3; j++ {
		if g.Board[0][j] != " " && g.Board[0][j] == g.Board[1][j] && g.Board[1][j] == g.Board[2][j] {
			g.GameOver = true
			g.Winner = &g.Players[g.Current]
			g.Winner.Score++
			return
		}
	}
	// Check diagonals
	if g.Board[0][0] != " " && g.Board[0][0] == g.Board[1][1] && g.Board[1][1] == g.Board[2][2] {
		g.GameOver = true
		g.Winner = &g.Players[g.Current]
		g.Winner.Score++
		return
	}
	if g.Board[0][2] != " " && g.Board[0][2] == g.Board[1][1] && g.Board[1][1] == g.Board[2][0] {
		g.GameOver = true
		g.Winner = &g.Players[g.Current]
		g.Winner.Score++
		return
	}
	// Check for draw
	draw := true
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.Board[i][j] == " " {
				draw = false
				break
			}
		}
		if !draw {
			break
		}
	}
	if draw {
		g.GameOver = true
	}
}

func (g *Game) GetAIMove() (int, int) {
	// Simple AI: choose a random valid move
	var moves [][2]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if g.IsValidMove(i, j) {
				moves = append(moves, [2]int{i, j})
			}
		}
	}
	if len(moves) == 0 {
		return -1, -1
	}
	rand.Seed(time.Now().UnixNano())
	move := moves[rand.Intn(len(moves))]
	return move[0], move[1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	game := NewGame()
	fmt.Println("Welcome to Tic-Tac-Toe!")
	fmt.Println("Player 1: X, Player 2: O")
	fmt.Println("Enter moves as row and column (e.g., 0 1). Type 'quit' to exit.")

	for {
		game.PrintBoard()
		currentPlayer := game.Players[game.Current]
		fmt.Printf("Player %s's turn. Enter move: ", currentPlayer.Symbol)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "quit" {
			fmt.Println("Thanks for playing!")
			break
		}

		coords := strings.Fields(input)
		if len(coords) != 2 {
			fmt.Println("Invalid input. Please enter two numbers.")
			continue
		}

		row, err1 := strconv.Atoi(coords[0])
		col, err2 := strconv.Atoi(coords[1])
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid numbers. Please enter integers.")
			continue
		}

		if !game.MakeMove(row, col) {
			fmt.Println("Invalid move. Try again.")
			continue
		}

		if game.GameOver {
			game.PrintBoard()
			if game.Winner != nil {
				fmt.Printf("Player %s wins!\n", game.Winner.Symbol)
			} else {
				fmt.Println("It's a draw!")
			}
			fmt.Printf("Scores - X: %d, O: %d\n", game.Players[0].Score, game.Players[1].Score)
			fmt.Print("Play again? (y/n): ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			if strings.ToLower(input) == "y" {
				game.ResetBoard()
				continue
			} else {
				fmt.Println("Goodbye!")
				break
			}
		}

		// AI move for player O
		if game.Current == 1 {
			row, col := game.GetAIMove()
			if row != -1 && col != -1 {
				game.MakeMove(row, col)
				fmt.Printf("AI (O) plays at %d %d\n", row, col)
			}
			if game.GameOver {
				game.PrintBoard()
				if game.Winner != nil {
					fmt.Printf("Player %s wins!\n", game.Winner.Symbol)
				} else {
					fmt.Println("It's a draw!")
				}
				fmt.Printf("Scores - X: %d, O: %d\n", game.Players[0].Score, game.Players[1].Score)
				fmt.Print("Play again? (y/n): ")
				input, _ := reader.ReadString('\n')
				input = strings.TrimSpace(input)
				if strings.ToLower(input) == "y" {
					game.ResetBoard()
					continue
				} else {
					fmt.Println("Goodbye!")
					break
				}
			}
		}
	}
}
