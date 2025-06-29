package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Board [3][3]string

func (b *Board) Initialize() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			b[i][j] = " "
		}
	}
}

func (b Board) Display() {
	fmt.Println("  0 1 2")
	for i := 0; i < 3; i++ {
		fmt.Printf("%d %s|%s|%s\n", i, b[i][0], b[i][1], b[i][2])
		if i < 2 {
			fmt.Println("  -----")
		}
	}
}

func (b Board) IsFull() bool {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if b[i][j] == " " {
				return false
			}
		}
	}
	return true
}

func (b Board) CheckWin() string {
	// Check rows
	for i := 0; i < 3; i++ {
		if b[i][0] != " " && b[i][0] == b[i][1] && b[i][1] == b[i][2] {
			return b[i][0]
		}
	}
	// Check columns
	for j := 0; j < 3; j++ {
		if b[0][j] != " " && b[0][j] == b[1][j] && b[1][j] == b[2][j] {
			return b[0][j]
		}
	}
	// Check diagonals
	if b[0][0] != " " && b[0][0] == b[1][1] && b[1][1] == b[2][2] {
		return b[0][0]
	}
	if b[0][2] != " " && b[0][2] == b[1][1] && b[1][1] == b[2][0] {
		return b[0][2]
	}
	return " "
}

func (b *Board) PlaceMark(row, col int, mark string) bool {
	if row < 0 || row >= 3 || col < 0 || col >= 3 || b[row][col] != " " {
		return false
	}
	b[row][col] = mark
	return true
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var board Board
	board.Initialize()
	currentPlayer := "X"

	for {
		board.Display()
		fmt.Printf("Player %s, enter row and column (0-2) separated by space: ", currentPlayer)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		var row, col int
		fmt.Sscanf(input, "%d %d", &row, &col)

		if !board.PlaceMark(row, col, currentPlayer) {
			fmt.Println("Invalid move. Try again.")
			continue
		}

		if winner := board.CheckWin(); winner != " " {
			board.Display()
			fmt.Printf("Player %s wins!\n", winner)
			break
		}

		if board.IsFull() {
			board.Display()
			fmt.Println("It's a draw!")
			break
		}

		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}
