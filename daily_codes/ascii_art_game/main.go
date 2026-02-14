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

// ASCII art data structures and functions
const (
    screenWidth  = 80
    screenHeight = 24
)

type Pixel struct {
    Char  rune
    Color string
}

type Screen [screenHeight][screenWidth]Pixel

func NewScreen() Screen {
    var s Screen
    for y := 0; y < screenHeight; y++ {
        for x := 0; x < screenWidth; x++ {
            s[y][x] = Pixel{' ', ""}
        }
    }
    return s
}

func (s *Screen) Clear() {
    for y := 0; y < screenHeight; y++ {
        for x := 0; x < screenWidth; x++ {
            s[y][x] = Pixel{' ', ""}
        }
    }
}

func (s *Screen) DrawPixel(x, y int, char rune, color string) {
    if x >= 0 && x < screenWidth && y >= 0 && y < screenHeight {
        s[y][x] = Pixel{char, color}
    }
}

func (s *Screen) DrawText(x, y int, text string, color string) {
    for i, ch := range text {
        s.DrawPixel(x+i, y, ch, color)
    }
}

func (s *Screen) DrawRect(x1, y1, x2, y2 int, char rune, color string) {
    for y := y1; y <= y2; y++ {
        for x := x1; x <= x2; x++ {
            s.DrawPixel(x, y, char, color)
        }
    }
}

func (s *Screen) DrawLine(x1, y1, x2, y2 int, char rune, color string) {
    dx := abs(x2 - x1)
    dy := abs(y2 - y1)
    sx := 1
    if x1 > x2 {
        sx = -1
    }
    sy := 1
    if y1 > y2 {
        sy = -1
    }
    err := dx - dy
    for {
        s.DrawPixel(x1, y1, char, color)
        if x1 == x2 && y1 == y2 {
            break
        }
        e2 := 2 * err
        if e2 > -dy {
            err -= dy
            x1 += sx
        }
        if e2 < dx {
            err += dx
            y1 += sy
        }
    }
}

func (s *Screen) Render() {
    fmt.Print("\033[2J\033[H") // Clear screen and move cursor to top-left
    for y := 0; y < screenHeight; y++ {
        for x := 0; x < screenWidth; x++ {
            pixel := s[y][x]
            if pixel.Color != "" {
                fmt.Printf("\033[%sm%c\033[0m", pixel.Color, pixel.Char)
            } else {
                fmt.Printf("%c", pixel.Char)
            }
        }
        fmt.Println()
    }
}

// Game entities
type Player struct {
    X, Y int
    Char rune
}

func NewPlayer(x, y int) Player {
    return Player{x, y, '@'}
}

func (p *Player) Move(dx, dy int) {
    p.X += dx
    p.Y += dy
    if p.X < 0 {
        p.X = 0
    }
    if p.X >= screenWidth {
        p.X = screenWidth - 1
    }
    if p.Y < 0 {
        p.Y = 0
    }
    if p.Y >= screenHeight {
        p.Y = screenHeight - 1
    }
}

type Enemy struct {
    X, Y int
    Char rune
    Alive bool
}

func NewEnemy(x, y int) Enemy {
    return Enemy{x, y, 'E', true}
}

func (e *Enemy) Update(playerX, playerY int) {
    if !e.Alive {
        return
    }
    dx := 0
    dy := 0
    if e.X < playerX {
        dx = 1
    } else if e.X > playerX {
        dx = -1
    }
    if e.Y < playerY {
        dy = 1
    } else if e.Y > playerY {
        dy = -1
    }
    e.X += dx
    e.Y += dy
    if e.X < 0 {
        e.X = 0
    }
    if e.X >= screenWidth {
        e.X = screenWidth - 1
    }
    if e.Y < 0 {
        e.Y = 0
    }
    if e.Y >= screenHeight {
        e.Y = screenHeight - 1
    }
}

type Projectile struct {
    X, Y int
    DX, DY int
    Char rune
    Active bool
}

func NewProjectile(x, y, dx, dy int) Projectile {
    return Projectile{x, y, dx, dy, '*', true}
}

func (p *Projectile) Update() {
    if !p.Active {
        return
    }
    p.X += p.DX
    p.Y += p.DY
    if p.X < 0 || p.X >= screenWidth || p.Y < 0 || p.Y >= screenHeight {
        p.Active = false
    }
}

// Game state
type Game struct {
    Screen   Screen
    Player   Player
    Enemies  []Enemy
    Projectiles []Projectile
    Score    int
    Running  bool
}

func NewGame() *Game {
    rand.Seed(time.Now().UnixNano())
    game := &Game{
        Screen:   NewScreen(),
        Player:   NewPlayer(screenWidth/2, screenHeight/2),
        Enemies:  make([]Enemy, 0),
        Projectiles: make([]Projectile, 0),
        Score:    0,
        Running:  true,
    }
    for i := 0; i < 5; i++ {
        game.Enemies = append(game.Enemies, NewEnemy(rand.Intn(screenWidth), rand.Intn(screenHeight)))
    }
    return game
}

func (g *Game) HandleInput(input string) {
    switch input {
    case "w":
        g.Player.Move(0, -1)
    case "s":
        g.Player.Move(0, 1)
    case "a":
        g.Player.Move(-1, 0)
    case "d":
        g.Player.Move(1, 0)
    case " ":
        g.Projectiles = append(g.Projectiles, NewProjectile(g.Player.X, g.Player.Y, 1, 0))
        g.Projectiles = append(g.Projectiles, NewProjectile(g.Player.X, g.Player.Y, -1, 0))
        g.Projectiles = append(g.Projectiles, NewProjectile(g.Player.X, g.Player.Y, 0, 1))
        g.Projectiles = append(g.Projectiles, NewProjectile(g.Player.X, g.Player.Y, 0, -1))
    case "q":
        g.Running = false
    }
}

func (g *Game) Update() {
    // Update enemies
    for i := range g.Enemies {
        g.Enemies[i].Update(g.Player.X, g.Player.Y)
    }
    // Update projectiles
    for i := range g.Projectiles {
        g.Projectiles[i].Update()
    }
    // Check collisions
    for i := range g.Projectiles {
        if !g.Projectiles[i].Active {
            continue
        }
        for j := range g.Enemies {
            if !g.Enemies[j].Alive {
                continue
            }
            if g.Projectiles[i].X == g.Enemies[j].X && g.Projectiles[i].Y == g.Enemies[j].Y {
                g.Enemies[j].Alive = false
                g.Projectiles[i].Active = false
                g.Score += 10
            }
        }
    }
    // Check player-enemy collisions
    for i := range g.Enemies {
        if !g.Enemies[i].Alive {
            continue
        }
        if g.Player.X == g.Enemies[i].X && g.Player.Y == g.Enemies[i].Y {
            g.Running = false
        }
    }
    // Respawn enemies if all dead
    allDead := true
    for i := range g.Enemies {
        if g.Enemies[i].Alive {
            allDead = false
            break
        }
    }
    if allDead {
        g.Enemies = make([]Enemy, 0)
        for i := 0; i < 5; i++ {
            g.Enemies = append(g.Enemies, NewEnemy(rand.Intn(screenWidth), rand.Intn(screenHeight)))
        }
    }
}

func (g *Game) Render() {
    g.Screen.Clear()
    // Draw border
    g.Screen.DrawRect(0, 0, screenWidth-1, screenHeight-1, '#', "37")
    // Draw player
    g.Screen.DrawPixel(g.Player.X, g.Player.Y, g.Player.Char, "32")
    // Draw enemies
    for _, enemy := range g.Enemies {
        if enemy.Alive {
            g.Screen.DrawPixel(enemy.X, enemy.Y, enemy.Char, "31")
        }
    }
    // Draw projectiles
    for _, proj := range g.Projectiles {
        if proj.Active {
            g.Screen.DrawPixel(proj.X, proj.Y, proj.Char, "33")
        }
    }
    // Draw score
    scoreText := fmt.Sprintf("Score: %d", g.Score)
    g.Screen.DrawText(2, 1, scoreText, "36")
    // Draw instructions
    instructions := "WASD: Move, Space: Shoot, Q: Quit"
    g.Screen.DrawText(2, screenHeight-2, instructions, "35")
    g.Screen.Render()
}

// Utility functions
func abs(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

// Main game loop
func main() {
    // Set terminal to raw mode for immediate input (simplified for cross-platform)
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("\033[?25l") // Hide cursor
    defer fmt.Print("\033[?25h") // Show cursor on exit

    game := NewGame()

    // Game introduction
    introScreen := NewScreen()
    introScreen.DrawText(20, 5, "ASCII Art Game", "34")
    introScreen.DrawText(15, 10, "Control a character (@) with WASD keys", "37")
    introScreen.DrawText(15, 12, "Shoot enemies (E) with SPACE bar", "37")
    introScreen.DrawText(15, 14, "Avoid collisions with enemies", "37")
    introScreen.DrawText(15, 16, "Press any key to start...", "32")
    introScreen.Render()
    reader.ReadRune() // Wait for key press

    // Main loop
    for game.Running {
        game.Render()
        // Non-blocking input (simplified)
        fmt.Print("\033[?25l") // Ensure cursor hidden
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if len(input) > 0 {
            game.HandleInput(string(input[0]))
        }
        game.Update()
        time.Sleep(50 * time.Millisecond) // Frame rate control
    }

    // Game over screen
    gameOverScreen := NewScreen()
    gameOverScreen.DrawText(25, 10, "Game Over!", "31")
    scoreText := fmt.Sprintf("Final Score: %d", game.Score)
    gameOverScreen.DrawText(25, 12, scoreText, "33")
    gameOverScreen.DrawText(20, 15, "Press any key to exit...", "37")
    gameOverScreen.Render()
    reader.ReadRune() // Wait for key press
}
