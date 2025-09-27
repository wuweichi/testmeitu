package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
	"github.com/eiannone/keyboard"
)

type Point struct {
	X, Y int
}

type Tetromino struct {
	Shape    [][]int
	Rotation int
	Color    int
}

type Game struct {
	Board          [][]int
	CurrentPiece   Tetromino
	NextPiece      Tetromino
	PiecePosition  Point
	Score          int
	Level          int
	LinesCleared   int
	GameOver       bool
	Paused         bool
}

const (
	BoardWidth  = 10
	BoardHeight = 20
	BlockSize   = 2
)

var colors = []string{
	"\033[0m",   // Reset
	"\033[31m", // Red
	"\033[32m", // Green
	"\033[33m", // Yellow
	"\033[34m", // Blue
	"\033[35m", // Magenta
	"\033[36m", // Cyan
	"\033[37m", // White
}

var tetrominoes = []Tetromino{
	// I
	{
		Shape: [][]int{
			{0, 0, 0, 0},
			{1, 1, 1, 1},
			{0, 0, 0, 0},
			{0, 0, 0, 0},
		},
		Color: 1,
	},
	// J
	{
		Shape: [][]int{
			{1, 0, 0},
			{1, 1, 1},
			{0, 0, 0},
		},
		Color: 2,
	},
	// L
	{
		Shape: [][]int{
			{0, 0, 1},
			{1, 1, 1},
			{0, 0, 0},
		},
		Color: 3,
	},
	// O
	{
		Shape: [][]int{
			{1, 1},
			{1, 1},
		},
		Color: 4,
	},
	// S
	{
		Shape: [][]int{
			{0, 1, 1},
			{1, 1, 0},
			{0, 0, 0},
		},
		Color: 5,
	},
	// T
	{
		Shape: [][]int{
			{0, 1, 0},
			{1, 1, 1},
			{0, 0, 0},
		},
		Color: 6,
	},
	// Z
	{
		Shape: [][]int{
			{1, 1, 0},
			{0, 1, 1},
			{0, 0, 0},
		},
		Color: 7,
	},
}

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func (g *Game) initBoard() {
	g.Board = make([][]int, BoardHeight)
	for i := range g.Board {
		g.Board[i] = make([]int, BoardWidth)
	}
}

func (g *Game) newPiece() {
	g.CurrentPiece = g.NextPiece
	g.NextPiece = tetrominoes[rand.Intn(len(tetrominoes))]
	g.PiecePosition = Point{X: BoardWidth/2 - len(g.CurrentPiece.Shape[0])/2, Y: 0}
	if g.checkCollision() {
		g.GameOver = true
	}
}

func (g *Game) rotatePiece() {
	oldRotation := g.CurrentPiece.Rotation
	g.CurrentPiece.Rotation = (g.CurrentPiece.Rotation + 1) % 4
	// Rotate the shape
	rows := len(g.CurrentPiece.Shape)
	cols := len(g.CurrentPiece.Shape[0])
	rotated := make([][]int, cols)
	for i := range rotated {
		rotated[i] = make([]int, rows)
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotated[j][rows-1-i] = g.CurrentPiece.Shape[i][j]
		}
	}
	g.CurrentPiece.Shape = rotated
	if g.checkCollision() {
		g.CurrentPiece.Rotation = oldRotation
		// Revert rotation
		g.CurrentPiece.Shape = make([][]int, rows)
		for i := range g.CurrentPiece.Shape {
			g.CurrentPiece.Shape[i] = make([]int, cols)
			for j := range g.CurrentPiece.Shape[i] {
				g.CurrentPiece.Shape[i][j] = rotated[cols-1-j][i]
			}
		}
	}
}

func (g *Game) movePiece(dx, dy int) {
	g.PiecePosition.X += dx
	g.PiecePosition.Y += dy
	if g.checkCollision() {
		g.PiecePosition.X -= dx
		g.PiecePosition.Y -= dy
		if dy > 0 {
			g.lockPiece()
			g.clearLines()
			g.newPiece()
		}
		return
	}
}

func (g *Game) checkCollision() bool {
	for y, row := range g.CurrentPiece.Shape {
		for x, cell := range row {
			if cell != 0 {
				boardX := g.PiecePosition.X + x
				boardY := g.PiecePosition.Y + y
				if boardX < 0 || boardX >= BoardWidth || boardY >= BoardHeight || (boardY >= 0 && g.Board[boardY][boardX] != 0) {
					return true
				}
			}
		}
	}
	return false
}

func (g *Game) lockPiece() {
	for y, row := range g.CurrentPiece.Shape {
		for x, cell := range row {
			if cell != 0 {
				boardY := g.PiecePosition.Y + y
				if boardY >= 0 {
					boardX := g.PiecePosition.X + x
					g.Board[boardY][boardX] = g.CurrentPiece.Color
				}
			}
		}
	}
}

func (g *Game) clearLines() {
	linesCleared := 0
	for y := BoardHeight - 1; y >= 0; y-- {
		full := true
		for x := 0; x < BoardWidth; x++ {
			if g.Board[y][x] == 0 {
				full = false
				break
			}
		}
		if full {
			linesCleared++
			// Remove the line
			for ny := y; ny > 0; ny-- {
				copy(g.Board[ny], g.Board[ny-1])
			}
			// Clear the top line
			for x := 0; x < BoardWidth; x++ {
				g.Board[0][x] = 0
			}
			y++ // Recheck the same line index after shifting
		}
	}
	if linesCleared > 0 {
		g.LinesCleared += linesCleared
		g.Score += linesCleared * 100 * g.Level
		g.Level = g.LinesCleared/10 + 1
	}
}

func (g *Game) draw() {
	clearScreen()
	fmt.Println("Advanced Tetris Game")
	fmt.Printf("Score: %d | Level: %d | Lines: %d\n", g.Score, g.Level, g.LinesCleared)
	fmt.Println("Controls: Arrow Keys to move, Up to rotate, P to pause, Q to quit")
	fmt.Println()

	// Draw the board with current piece
	for y := 0; y < BoardHeight; y++ {
		fmt.Print("| ")
		for x := 0; x < BoardWidth; x++ {
			cell := g.Board[y][x]
			// Check if current piece occupies this cell
			if y >= g.PiecePosition.Y && y < g.PiecePosition.Y+len(g.CurrentPiece.Shape) &&
				x >= g.PiecePosition.X && x < g.PiecePosition.X+len(g.CurrentPiece.Shape[0]) {
				pieceY := y - g.PiecePosition.Y
				pieceX := x - g.PiecePosition.X
				if pieceY >= 0 && pieceY < len(g.CurrentPiece.Shape) &&
					pieceX >= 0 && pieceX < len(g.CurrentPiece.Shape[0]) &&
					g.CurrentPiece.Shape[pieceY][pieceX] != 0 {
					cell = g.CurrentPiece.Color
				}
			}
			if cell == 0 {
				fmt.Print("  ")
			} else {
				fmt.Printf("%s██\033[0m", colors[cell])
			}
		}
		fmt.Println(" |")
	}
	fmt.Print("+")
	for x := 0; x < BoardWidth*2; x++ {
		fmt.Print("-")
	}
	fmt.Println("+")

	// Draw next piece preview
	fmt.Println("\nNext Piece:")
	for y := 0; y < len(g.NextPiece.Shape); y++ {
		for x := 0; x < len(g.NextPiece.Shape[0]); x++ {
			if g.NextPiece.Shape[y][x] != 0 {
				fmt.Printf("%s██\033[0m", colors[g.NextPiece.Color])
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

	if g.Paused {
		fmt.Println("\nGame Paused - Press P to resume")
	}
	if g.GameOver {
		fmt.Println("\nGame Over!")
		fmt.Printf("Final Score: %d\n", g.Score)
	}
}

func (g *Game) run() {
	err := keyboard.Open()
	if err != nil {
		fmt.Println("Error opening keyboard:", err)
		return
	}
	defer keyboard.Close()

	g.initBoard()
	g.NextPiece = tetrominoes[rand.Intn(len(tetrominoes))]
	g.newPiece()

	ticker := time.NewTicker(time.Second / time.Duration(g.Level))
	defer ticker.Stop()

	for !g.GameOver {
		select {
		case <-ticker.C:
			if !g.Paused {
				g.movePiece(0, 1)
			}
		default:
			if g.Paused {
				g.draw()
				time.Sleep(100 * time.Millisecond)
				continue
			}
			char, key, err := keyboard.GetKey()
			if err != nil {
				continue
			}
			switch {
			case key == keyboard.KeyArrowLeft:
				g.movePiece(-1, 0)
			case key == keyboard.KeyArrowRight:
				g.movePiece(1, 0)
			case key == keyboard.KeyArrowDown:
				g.movePiece(0, 1)
			case key == keyboard.KeyArrowUp:
				g.rotatePiece()
			case char == 'p' || char == 'P':
				g.Paused = !g.Paused
			case char == 'q' || char == 'Q':
				return
			}
		}
		g.draw()
	}

	// Game over loop
	for {
		g.draw()
		char, key, err := keyboard.GetKey()
		if err != nil {
			continue
		}
		if char == 'q' || char == 'Q' || key == keyboard.KeyEsc {
			break
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := &Game{}
	game.run()
}
