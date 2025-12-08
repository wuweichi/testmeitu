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

const (
    boardWidth  = 10
    boardHeight = 20
    cellEmpty   = 0
    cellFilled  = 1
    cellCurrent = 2
)

type Point struct {
    X, Y int
}

type Tetromino struct {
    shape [][]int
    color string
}

var tetrominoes = []Tetromino{
    {shape: [][]int{{1, 1, 1, 1}}, color: "\033[36m"},
    {shape: [][]int{{1, 1}, {1, 1}}, color: "\033[33m"},
    {shape: [][]int{{0, 1, 0}, {1, 1, 1}}, color: "\033[35m"},
    {shape: [][]int{{1, 1, 0}, {0, 1, 1}}, color: "\033[32m"},
    {shape: [][]int{{0, 1, 1}, {1, 1, 0}}, color: "\033[31m"},
    {shape: [][]int{{1, 0, 0}, {1, 1, 1}}, color: "\033[34m"},
    {shape: [][]int{{0, 0, 1}, {1, 1, 1}}, color: "\033[37m"},
}

type Game struct {
    board      [boardHeight][boardWidth]int
    current    Tetromino
    position   Point
    score      int
    level      int
    lines      int
    gameOver   bool
    paused     bool
    startTime  time.Time
    lastDrop   time.Time
    dropSpeed  time.Duration
}

func (g *Game) init() {
    rand.Seed(time.Now().UnixNano())
    g.newPiece()
    g.score = 0
    g.level = 1
    g.lines = 0
    g.gameOver = false
    g.paused = false
    g.startTime = time.Now()
    g.lastDrop = time.Now()
    g.dropSpeed = 1000 * time.Millisecond
}

func (g *Game) newPiece() {
    g.current = tetrominoes[rand.Intn(len(tetrominoes))]
    g.position = Point{X: boardWidth/2 - len(g.current.shape[0])/2, Y: 0}
    if g.checkCollision() {
        g.gameOver = true
    }
}

func (g *Game) checkCollision() bool {
    for y, row := range g.current.shape {
        for x, cell := range row {
            if cell == 1 {
                boardX := g.position.X + x
                boardY := g.position.Y + y
                if boardX < 0 || boardX >= boardWidth || boardY >= boardHeight || (boardY >= 0 && g.board[boardY][boardX] == cellFilled) {
                    return true
                }
            }
        }
    }
    return false
}

func (g *Game) mergePiece() {
    for y, row := range g.current.shape {
        for x, cell := range row {
            if cell == 1 {
                boardX := g.position.X + x
                boardY := g.position.Y + y
                if boardY >= 0 {
                    g.board[boardY][boardX] = cellFilled
                }
            }
        }
    }
}

func (g *Game) clearLines() {
    linesCleared := 0
    for y := boardHeight - 1; y >= 0; y-- {
        full := true
        for x := 0; x < boardWidth; x++ {
            if g.board[y][x] == cellEmpty {
                full = false
                break
            }
        }
        if full {
            linesCleared++
            for yy := y; yy > 0; yy-- {
                for x := 0; x < boardWidth; x++ {
                    g.board[yy][x] = g.board[yy-1][x]
                }
            }
            for x := 0; x < boardWidth; x++ {
                g.board[0][x] = cellEmpty
            }
            y++
        }
    }
    if linesCleared > 0 {
        g.lines += linesCleared
        g.score += linesCleared * 100 * g.level
        g.level = g.lines/10 + 1
        g.dropSpeed = time.Duration(1000/g.level) * time.Millisecond
        if g.dropSpeed < 100*time.Millisecond {
            g.dropSpeed = 100 * time.Millisecond
        }
    }
}

func (g *Game) move(dx, dy int) bool {
    g.position.X += dx
    g.position.Y += dy
    if g.checkCollision() {
        g.position.X -= dx
        g.position.Y -= dy
        return false
    }
    return true
}

func (g *Game) rotate() {
    oldShape := g.current.shape
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
    g.current.shape = newShape
    if g.checkCollision() {
        g.current.shape = oldShape
    }
}

func (g *Game) drop() {
    for g.move(0, 1) {
    }
    g.mergePiece()
    g.clearLines()
    g.newPiece()
}

func (g *Game) update() {
    if g.gameOver || g.paused {
        return
    }
    if time.Since(g.lastDrop) > g.dropSpeed {
        if !g.move(0, 1) {
            g.mergePiece()
            g.clearLines()
            g.newPiece()
        }
        g.lastDrop = time.Now()
    }
}

func (g *Game) draw() {
    clearScreen()
    fmt.Println("Advanced Tetris Game")
    fmt.Printf("Score: %d | Level: %d | Lines: %d\n", g.score, g.level, g.lines)
    fmt.Printf("Time: %s\n", time.Since(g.startTime).Round(time.Second))
    fmt.Println("Controls: Arrow Keys (Move), Up (Rotate), Space (Drop), P (Pause), Q (Quit)")
    fmt.Println()
    for y := 0; y < boardHeight; y++ {
        fmt.Print("| ")
        for x := 0; x < boardWidth; x++ {
            cell := g.board[y][x]
            inCurrent := false
            if !g.gameOver {
                for cy, row := range g.current.shape {
                    for cx, ccell := range row {
                        if ccell == 1 && g.position.Y+cy == y && g.position.X+cx == x {
                            inCurrent = true
                            break
                        }
                    }
                    if inCurrent {
                        break
                    }
                }
            }
            if inCurrent {
                fmt.Printf("%s█\033[0m ", g.current.color)
            } else if cell == cellFilled {
                fmt.Print("█ ")
            } else {
                fmt.Print(". ")
            }
        }
        fmt.Println("|")
    }
    fmt.Println("+----------------------+")
    if g.gameOver {
        fmt.Println("GAME OVER!")
        fmt.Println("Press Q to quit.")
    } else if g.paused {
        fmt.Println("PAUSED")
    }
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

func main() {
    if err := keyboard.Open(); err != nil {
        fmt.Println("Error opening keyboard:", err)
        return
    }
    defer keyboard.Close()
    game := &Game{}
    game.init()
    ticker := time.NewTicker(50 * time.Millisecond)
    defer ticker.Stop()
    keyEvents := make(chan keyboard.Key)
    go func() {
        for {
            _, key, err := keyboard.GetKey()
            if err != nil {
                close(keyEvents)
                return
            }
            keyEvents <- key
        }
    }()
    for {
        select {
        case <-ticker.C:
            game.update()
            game.draw()
        case key := <-keyEvents:
            if game.gameOver {
                if key == keyboard.KeyEsc || key == keyboard.KeyCtrlC || key == keyboard.Key('q') || key == keyboard.Key('Q') {
                    fmt.Println("Thanks for playing!")
                    return
                }
                continue
            }
            switch key {
            case keyboard.KeyArrowLeft:
                game.move(-1, 0)
            case keyboard.KeyArrowRight:
                game.move(1, 0)
            case keyboard.KeyArrowDown:
                game.move(0, 1)
            case keyboard.KeyArrowUp:
                game.rotate()
            case keyboard.KeySpace:
                game.drop()
            case keyboard.Key('p'), keyboard.Key('P'):
                game.paused = !game.paused
            case keyboard.KeyEsc, keyboard.KeyCtrlC, keyboard.Key('q'), keyboard.Key('Q'):
                fmt.Println("Quitting game...")
                return
            }
        }
    }
}
