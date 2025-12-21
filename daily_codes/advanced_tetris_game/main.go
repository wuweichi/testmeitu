package main

import (
    "fmt"
    "math/rand"
    "os"
    "time"
    "github.com/nsf/termbox-go"
)

const (
    boardWidth  = 10
    boardHeight = 20
    cellWidth   = 2
    cellHeight  = 1
    startX      = 5
    startY      = 0
    fps         = 60
    dropSpeed   = 1000 // milliseconds
)

type Point struct {
    X, Y int
}

type Tetromino struct {
    shape [][]int
    color termbox.Attribute
    x, y  int
}

var tetrominoes = []Tetromino{
    {shape: [][]int{{1, 1, 1, 1}}, color: termbox.ColorCyan},
    {shape: [][]int{{1, 1}, {1, 1}}, color: termbox.ColorYellow},
    {shape: [][]int{{0, 1, 0}, {1, 1, 1}}, color: termbox.ColorMagenta},
    {shape: [][]int{{1, 1, 0}, {0, 1, 1}}, color: termbox.ColorGreen},
    {shape: [][]int{{0, 1, 1}, {1, 1, 0}}, color: termbox.ColorRed},
    {shape: [][]int{{1, 0, 0}, {1, 1, 1}}, color: termbox.ColorBlue},
    {shape: [][]int{{0, 0, 1}, {1, 1, 1}}, color: termbox.ColorWhite},
}

type Game struct {
    board      [boardHeight][boardWidth]int
    current    Tetromino
    next       Tetromino
    score      int
    level      int
    lines      int
    gameOver   bool
    paused     bool
    lastDrop   time.Time
    startTime  time.Time
    elapsed    time.Duration
}

func (g *Game) init() {
    rand.Seed(time.Now().UnixNano())
    g.current = g.randomTetromino()
    g.next = g.randomTetromino()
    g.score = 0
    g.level = 1
    g.lines = 0
    g.gameOver = false
    g.paused = false
    g.lastDrop = time.Now()
    g.startTime = time.Now()
    g.elapsed = 0
    for y := 0; y < boardHeight; y++ {
        for x := 0; x < boardWidth; x++ {
            g.board[y][x] = 0
        }
    }
}

func (g *Game) randomTetromino() Tetromino {
    idx := rand.Intn(len(tetrominoes))
    t := tetrominoes[idx]
    t.x = startX
    t.y = startY
    return t
}

func (g *Game) drawBoard() {
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    for y := 0; y < boardHeight; y++ {
        for x := 0; x < boardWidth; x++ {
            if g.board[y][x] != 0 {
                termbox.SetCell(x*cellWidth, y*cellHeight, '█', termbox.ColorCyan, termbox.ColorDefault)
                termbox.SetCell(x*cellWidth+1, y*cellHeight, '█', termbox.ColorCyan, termbox.ColorDefault)
            } else {
                termbox.SetCell(x*cellWidth, y*cellHeight, '.', termbox.ColorDarkGray, termbox.ColorDefault)
                termbox.SetCell(x*cellWidth+1, y*cellHeight, '.', termbox.ColorDarkGray, termbox.ColorDefault)
            }
        }
    }
}

func (g *Game) drawTetromino(t Tetromino) {
    for y, row := range t.shape {
        for x, cell := range row {
            if cell != 0 {
                termbox.SetCell((t.x+x)*cellWidth, (t.y+y)*cellHeight, '█', t.color, termbox.ColorDefault)
                termbox.SetCell((t.x+x)*cellWidth+1, (t.y+y)*cellHeight, '█', t.color, termbox.ColorDefault)
            }
        }
    }
}

func (g *Game) drawUI() {
    infoX := boardWidth*cellWidth + 2
    fmt.Fprintf(os.Stdout, "\033[%d;%dHScore: %d", 1, infoX, g.score)
    fmt.Fprintf(os.Stdout, "\033[%d;%dHLevel: %d", 2, infoX, g.level)
    fmt.Fprintf(os.Stdout, "\033[%d;%dHLines: %d", 3, infoX, g.lines)
    fmt.Fprintf(os.Stdout, "\033[%d;%dHTime: %v", 4, infoX, g.elapsed)
    fmt.Fprintf(os.Stdout, "\033[%d;%dHNext:", 6, infoX)
    for y, row := range g.next.shape {
        for x, cell := range row {
            if cell != 0 {
                termbox.SetCell(infoX+x*cellWidth, 7+y, '█', g.next.color, termbox.ColorDefault)
            }
        }
    }
    if g.gameOver {
        msg := "GAME OVER! Press 'r' to restart."
        x := (boardWidth*cellWidth - len(msg)) / 2
        for i, ch := range msg {
            termbox.SetCell(x+i, boardHeight/2, ch, termbox.ColorRed, termbox.ColorDefault)
        }
    }
    if g.paused {
        msg := "PAUSED - Press 'p' to resume"
        x := (boardWidth*cellWidth - len(msg)) / 2
        for i, ch := range msg {
            termbox.SetCell(x+i, boardHeight/2-2, ch, termbox.ColorYellow, termbox.ColorDefault)
        }
    }
}

func (g *Game) collision(t Tetromino) bool {
    for y, row := range t.shape {
        for x, cell := range row {
            if cell != 0 {
                boardX := t.x + x
                boardY := t.y + y
                if boardX < 0 || boardX >= boardWidth || boardY >= boardHeight {
                    return true
                }
                if boardY >= 0 && g.board[boardY][boardX] != 0 {
                    return true
                }
            }
        }
    }
    return false
}

func (g *Game) mergeTetromino() {
    for y, row := range g.current.shape {
        for x, cell := range row {
            if cell != 0 {
                boardX := g.current.x + x
                boardY := g.current.y + y
                if boardY >= 0 {
                    g.board[boardY][boardX] = 1
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
            if g.board[y][x] == 0 {
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
                g.board[0][x] = 0
            }
            y++
        }
    }
    if linesCleared > 0 {
        g.lines += linesCleared
        g.score += linesCleared * 100 * g.level
        g.level = g.lines/10 + 1
    }
}

func (g *Game) rotateTetromino() {
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
    if g.collision(g.current) {
        g.current.shape = oldShape
    }
}

func (g *Game) moveTetromino(dx, dy int) {
    g.current.x += dx
    g.current.y += dy
    if g.collision(g.current) {
        g.current.x -= dx
        g.current.y -= dy
        if dy > 0 {
            g.mergeTetromino()
            g.clearLines()
            g.current = g.next
            g.next = g.randomTetromino()
            g.current.x = startX
            g.current.y = startY
            if g.collision(g.current) {
                g.gameOver = true
            }
        }
    }
}

func (g *Game) update() {
    if g.gameOver || g.paused {
        return
    }
    now := time.Now()
    g.elapsed = now.Sub(g.startTime)
    if now.Sub(g.lastDrop) > time.Duration(dropSpeed/g.level)*time.Millisecond {
        g.moveTetromino(0, 1)
        g.lastDrop = now
    }
}

func (g *Game) handleInput(ev termbox.Event) {
    switch ev.Key {
    case termbox.KeyArrowLeft:
        g.moveTetromino(-1, 0)
    case termbox.KeyArrowRight:
        g.moveTetromino(1, 0)
    case termbox.KeyArrowDown:
        g.moveTetromino(0, 1)
    case termbox.KeyArrowUp:
        g.rotateTetromino()
    case termbox.KeySpace:
        for !g.collision(Tetromino{shape: g.current.shape, x: g.current.x, y: g.current.y + 1}) {
            g.current.y++
        }
        g.moveTetromino(0, 1)
    case termbox.KeyEsc:
        g.gameOver = true
    case termbox.KeyCtrlC:
        termbox.Close()
        os.Exit(0)
    }
    switch ev.Ch {
    case 'p', 'P':
        g.paused = !g.paused
    case 'r', 'R':
        if g.gameOver {
            g.init()
        }
    }
}

func main() {
    err := termbox.Init()
    if err != nil {
        panic(err)
    }
    defer termbox.Close()
    termbox.SetOutputMode(termbox.OutputNormal)
    termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
    game := &Game{}
    game.init()
    eventQueue := make(chan termbox.Event)
    go func() {
        for {
            eventQueue <- termbox.PollEvent()
        }
    }()
    ticker := time.NewTicker(time.Second / fps)
    defer ticker.Stop()
    for {
        select {
        case ev := <-eventQueue:
            game.handleInput(ev)
        case <-ticker.C:
            game.update()
            game.drawBoard()
            game.drawTetromino(game.current)
            game.drawUI()
            termbox.Flush()
        }
        if game.gameOver && !game.paused {
            break
        }
    }
    termbox.Close()
    fmt.Println("Thanks for playing! Final score:", game.score)
}
