package main

import (
    "fmt"
    "math/rand"
    "os"
    "os/exec"
    "runtime"
    "time"
    "github.com/eiannone/keyboard"
    "github.com/fatih/color"
)

const (
    boardWidth  = 10
    boardHeight = 20
    cellWidth   = 2
    cellHeight  = 1
)

type Point struct {
    X, Y int
}

type Tetromino struct {
    shape [][]int
    color color.Attribute
    pos   Point
}

type Game struct {
    board         [][]int
    currentPiece  *Tetromino
    nextPiece     *Tetromino
    score         int
    level         int
    linesCleared  int
    gameOver      bool
    paused        bool
    startTime     time.Time
    lastDropTime  time.Time
    dropInterval  time.Duration
    inputChan     chan rune
    quitChan      chan bool
}

var tetrominoes = []Tetromino{
    {shape: [][]int{{1, 1, 1, 1}}, color: color.FgCyan},    // I
    {shape: [][]int{{1, 1}, {1, 1}}, color: color.FgYellow}, // O
    {shape: [][]int{{0, 1, 0}, {1, 1, 1}}, color: color.FgMagenta}, // T
    {shape: [][]int{{0, 1, 1}, {1, 1, 0}}, color: color.FgGreen}, // S
    {shape: [][]int{{1, 1, 0}, {0, 1, 1}}, color: color.FgRed}, // Z
    {shape: [][]int{{1, 0, 0}, {1, 1, 1}}, color: color.FgBlue}, // J
    {shape: [][]int{{0, 0, 1}, {1, 1, 1}}, color: color.FgWhite}, // L
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
    g.board = make([][]int, boardHeight)
    for i := range g.board {
        g.board[i] = make([]int, boardWidth)
    }
}

func (g *Game) newRandomPiece() *Tetromino {
    idx := rand.Intn(len(tetrominoes))
    piece := tetrominoes[idx]
    piece.pos = Point{X: boardWidth/2 - len(piece.shape[0])/2, Y: 0}
    return &piece
}

func (g *Game) spawnNewPiece() {
    g.currentPiece = g.nextPiece
    g.nextPiece = g.newRandomPiece()
    if g.checkCollision(g.currentPiece.pos.X, g.currentPiece.pos.Y, g.currentPiece.shape) {
        g.gameOver = true
    }
}

func (g *Game) checkCollision(x, y int, shape [][]int) bool {
    for dy, row := range shape {
        for dx, cell := range row {
            if cell != 0 {
                newX, newY := x+dx, y+dy
                if newX < 0 || newX >= boardWidth || newY >= boardHeight || (newY >= 0 && g.board[newY][newX] != 0) {
                    return true
                }
            }
        }
    }
    return false
}

func (g *Game) mergePiece() {
    for dy, row := range g.currentPiece.shape {
        for dx, cell := range row {
            if cell != 0 {
                y := g.currentPiece.pos.Y + dy
                x := g.currentPiece.pos.X + dx
                if y >= 0 && y < boardHeight && x >= 0 && x < boardWidth {
                    g.board[y][x] = int(g.currentPiece.color)
                }
            }
        }
    }
}

func (g *Game) movePiece(dx, dy int) bool {
    newX := g.currentPiece.pos.X + dx
    newY := g.currentPiece.pos.Y + dy
    if !g.checkCollision(newX, newY, g.currentPiece.shape) {
        g.currentPiece.pos.X = newX
        g.currentPiece.pos.Y = newY
        return true
    }
    return false
}

func (g *Game) rotatePiece() {
    oldShape := g.currentPiece.shape
    rows := len(oldShape)
    cols := len(oldShape[0])
    newShape := make([][]int, cols)
    for i := range newShape {
        newShape[i] = make([]int, rows)
    }
    for y := 0; y < rows; y++ {
        for x := 0; x < cols; x++ {
            newShape[x][rows-1-y] = oldShape[y][x]
        }
    }
    if !g.checkCollision(g.currentPiece.pos.X, g.currentPiece.pos.Y, newShape) {
        g.currentPiece.shape = newShape
    }
}

func (g *Game) clearLines() {
    linesToClear := []int{}
    for y := 0; y < boardHeight; y++ {
        full := true
        for x := 0; x < boardWidth; x++ {
            if g.board[y][x] == 0 {
                full = false
                break
            }
        }
        if full {
            linesToClear = append(linesToClear, y)
        }
    }
    if len(linesToClear) > 0 {
        for _, line := range linesToClear {
            for y := line; y > 0; y-- {
                copy(g.board[y], g.board[y-1])
            }
            for x := 0; x < boardWidth; x++ {
                g.board[0][x] = 0
            }
        }
        g.linesCleared += len(linesToClear)
        g.score += len(linesToClear) * 100 * (g.level + 1)
        g.level = g.linesCleared / 10
        g.dropInterval = time.Duration(1000-(g.level*50)) * time.Millisecond
        if g.dropInterval < 100*time.Millisecond {
            g.dropInterval = 100 * time.Millisecond
        }
    }
}

func (g *Game) draw() {
    clearScreen()
    c := color.New(color.FgHiWhite)
    c.Println("=== ADVANCED TETRIS GAME ===")
    fmt.Printf("Score: %d  Level: %d  Lines: %d\n", g.score, g.level, g.linesCleared)
    fmt.Printf("Time: %v\n", time.Since(g.startTime).Round(time.Second))
    fmt.Println("Controls: Arrow Keys (Move), Up (Rotate), P (Pause), Q (Quit)")
    fmt.Println()

    // Draw board with current piece
    for y := 0; y < boardHeight; y++ {
        fmt.Print("| ")
        for x := 0; x < boardWidth; x++ {
            cellValue := g.board[y][x]
            // Check if current piece occupies this cell
            if g.currentPiece != nil && !g.gameOver {
                pieceX := x - g.currentPiece.pos.X
                pieceY := y - g.currentPiece.pos.Y
                if pieceY >= 0 && pieceY < len(g.currentPiece.shape) &&
                    pieceX >= 0 && pieceX < len(g.currentPiece.shape[0]) &&
                    g.currentPiece.shape[pieceY][pieceX] != 0 {
                    cellValue = int(g.currentPiece.color)
                }
            }
            if cellValue == 0 {
                fmt.Print("  ")
            } else {
                c := color.New(color.Attribute(cellValue))
                c.Print("██")
            }
        }
        fmt.Println(" |")
    }
    fmt.Println("  " + repeat("──", boardWidth) + " ")

    // Draw next piece preview
    fmt.Println("\nNext Piece:")
    if g.nextPiece != nil {
        for y := 0; y < len(g.nextPiece.shape); y++ {
            fmt.Print("  ")
            for x := 0; x < len(g.nextPiece.shape[0]); x++ {
                if g.nextPiece.shape[y][x] != 0 {
                    c := color.New(g.nextPiece.color)
                    c.Print("██")
                } else {
                    fmt.Print("  ")
                }
            }
            fmt.Println()
        }
    }
    fmt.Println()
    if g.paused {
        color.New(color.FgYellow, color.Bold).Println("*** PAUSED ***")
    }
    if g.gameOver {
        color.New(color.FgRed, color.Bold).Println("*** GAME OVER ***")
        fmt.Println("Press 'Q' to quit.")
    }
}

func repeat(s string, n int) string {
    result := ""
    for i := 0; i < n; i++ {
        result += s
    }
    return result
}

func (g *Game) handleInput() {
    if err := keyboard.Open(); err != nil {
        panic(err)
    }
    defer keyboard.Close()

    for {
        select {
        case <-g.quitChan:
            return
        default:
            char, key, err := keyboard.GetKey()
            if err != nil {
                continue
            }
            if key == keyboard.KeyEsc || char == 'q' || char == 'Q' {
                g.quitChan <- true
                return
            }
            if !g.gameOver {
                if char == 'p' || char == 'P' {
                    g.paused = !g.paused
                }
                if !g.paused {
                    switch key {
                    case keyboard.KeyArrowLeft:
                        g.movePiece(-1, 0)
                    case keyboard.KeyArrowRight:
                        g.movePiece(1, 0)
                    case keyboard.KeyArrowDown:
                        g.movePiece(0, 1)
                    case keyboard.KeyArrowUp:
                        g.rotatePiece()
                    case keyboard.KeySpace:
                        // Hard drop
                        for g.movePiece(0, 1) {
                        }
                        g.mergePiece()
                        g.clearLines()
                        g.spawnNewPiece()
                    }
                }
            }
            g.inputChan <- char
        }
    }
}

func (g *Game) run() {
    g.startTime = time.Now()
    g.lastDropTime = time.Now()
    g.dropInterval = 1000 * time.Millisecond
    g.inputChan = make(chan rune, 10)
    g.quitChan = make(chan bool, 1)

    go g.handleInput()

    for {
        select {
        case <-g.quitChan:
            return
        default:
            if !g.gameOver && !g.paused {
                currentTime := time.Now()
                if currentTime.Sub(g.lastDropTime) > g.dropInterval {
                    if !g.movePiece(0, 1) {
                        g.mergePiece()
                        g.clearLines()
                        g.spawnNewPiece()
                    }
                    g.lastDropTime = currentTime
                }
            }
            g.draw()
            time.Sleep(50 * time.Millisecond)
        }
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    game := &Game{}
    game.initBoard()
    game.nextPiece = game.newRandomPiece()
    game.spawnNewPiece()
    game.run()
    clearScreen()
    color.New(color.FgHiGreen).Println("Thanks for playing Advanced Tetris!")
    fmt.Printf("Final Score: %d, Lines Cleared: %d, Level: %d\n", game.score, game.linesCleared, game.level)
    fmt.Printf("Total Time: %v\n", time.Since(game.startTime).Round(time.Second))
}
