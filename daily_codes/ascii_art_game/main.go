package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
    "unicode/utf8"
)

const (
    screenWidth  = 80
    screenHeight = 24
    maxEnemies   = 50
    maxStars     = 100
    maxParticles = 200
    playerSpeed  = 2
    bulletSpeed  = 3
    enemySpeed   = 1
    gameFPS      = 30
)

type Vector2 struct {
    X, Y int
}

type Player struct {
    Position Vector2
    Health   int
    Score    int
    Symbol   rune
}

type Bullet struct {
    Position Vector2
    Active   bool
    Symbol   rune
}

type Enemy struct {
    Position Vector2
    Active   bool
    Health   int
    Symbol   rune
    Type     int
}

type Star struct {
    Position Vector2
    Brightness int
}

type Particle struct {
    Position Vector2
    Velocity Vector2
    Life     int
    Symbol   rune
}

type Game struct {
    Player     Player
    Bullets    [100]Bullet
    Enemies    [maxEnemies]Enemy
    Stars      [maxStars]Star
    Particles  [maxParticles]Particle
    GameOver   bool
    Level      int
    FrameCount int
    Input      string
}

func (g *Game) Init() {
    rand.Seed(time.Now().UnixNano())
    g.Player = Player{
        Position: Vector2{screenWidth / 2, screenHeight - 2},
        Health:   100,
        Score:    0,
        Symbol:   'A',
    }
    for i := range g.Bullets {
        g.Bullets[i].Active = false
        g.Bullets[i].Symbol = '|'
    }
    for i := range g.Enemies {
        g.Enemies[i].Active = false
    }
    for i := range g.Stars {
        g.Stars[i].Position = Vector2{rand.Intn(screenWidth), rand.Intn(screenHeight)}
        g.Stars[i].Brightness = rand.Intn(2) + 1
    }
    g.GameOver = false
    g.Level = 1
    g.FrameCount = 0
    g.SpawnEnemies()
}

func (g *Game) SpawnEnemies() {
    count := 5 + g.Level*2
    if count > maxEnemies {
        count = maxEnemies
    }
    spawned := 0
    for i := range g.Enemies {
        if !g.Enemies[i].Active && spawned < count {
            g.Enemies[i].Active = true
            g.Enemies[i].Position = Vector2{rand.Intn(screenWidth - 2) + 1, rand.Intn(5) + 1}
            g.Enemies[i].Health = 10 + g.Level*5
            g.Enemies[i].Type = rand.Intn(3)
            switch g.Enemies[i].Type {
            case 0:
                g.Enemies[i].Symbol = 'V'
            case 1:
                g.Enemies[i].Symbol = 'W'
            case 2:
                g.Enemies[i].Symbol = 'M'
            }
            spawned++
        }
    }
}

func (g *Game) Update() {
    if g.GameOver {
        return
    }
    g.FrameCount++
    g.UpdateStars()
    g.UpdateParticles()
    g.HandleInput()
    g.UpdateBullets()
    g.UpdateEnemies()
    g.CheckCollisions()
    if g.FrameCount%300 == 0 {
        g.Level++
        g.SpawnEnemies()
    }
}

func (g *Game) UpdateStars() {
    for i := range g.Stars {
        g.Stars[i].Position.Y++
        if g.Stars[i].Position.Y >= screenHeight {
            g.Stars[i].Position = Vector2{rand.Intn(screenWidth), 0}
            g.Stars[i].Brightness = rand.Intn(2) + 1
        }
    }
}

func (g *Game) UpdateParticles() {
    for i := range g.Particles {
        if g.Particles[i].Life > 0 {
            g.Particles[i].Position.X += g.Particles[i].Velocity.X
            g.Particles[i].Position.Y += g.Particles[i].Velocity.Y
            g.Particles[i].Life--
            if g.Particles[i].Position.X < 0 || g.Particles[i].Position.X >= screenWidth ||
                g.Particles[i].Position.Y < 0 || g.Particles[i].Position.Y >= screenHeight {
                g.Particles[i].Life = 0
            }
        }
    }
}

func (g *Game) HandleInput() {
    switch g.Input {
    case "a":
        if g.Player.Position.X > 1 {
            g.Player.Position.X -= playerSpeed
        }
    case "d":
        if g.Player.Position.X < screenWidth-2 {
            g.Player.Position.X += playerSpeed
        }
    case "w":
        if g.Player.Position.Y > 1 {
            g.Player.Position.Y -= playerSpeed
        }
    case "s":
        if g.Player.Position.Y < screenHeight-2 {
            g.Player.Position.Y += playerSpeed
        }
    case " ":
        g.ShootBullet()
    case "q":
        g.GameOver = true
    }
    g.Input = ""
}

func (g *Game) ShootBullet() {
    for i := range g.Bullets {
        if !g.Bullets[i].Active {
            g.Bullets[i].Active = true
            g.Bullets[i].Position = Vector2{g.Player.Position.X, g.Player.Position.Y - 1}
            break
        }
    }
}

func (g *Game) UpdateBullets() {
    for i := range g.Bullets {
        if g.Bullets[i].Active {
            g.Bullets[i].Position.Y -= bulletSpeed
            if g.Bullets[i].Position.Y < 0 {
                g.Bullets[i].Active = false
            }
        }
    }
}

func (g *Game) UpdateEnemies() {
    for i := range g.Enemies {
        if g.Enemies[i].Active {
            g.Enemies[i].Position.Y += enemySpeed
            if g.Enemies[i].Position.Y >= screenHeight-1 {
                g.Enemies[i].Active = false
                g.Player.Health -= 10
                if g.Player.Health <= 0 {
                    g.GameOver = true
                }
            }
            if g.FrameCount%20 == 0 {
                g.Enemies[i].Position.X += rand.Intn(3) - 1
                if g.Enemies[i].Position.X < 1 {
                    g.Enemies[i].Position.X = 1
                }
                if g.Enemies[i].Position.X >= screenWidth-1 {
                    g.Enemies[i].Position.X = screenWidth - 2
                }
            }
        }
    }
}

func (g *Game) CheckCollisions() {
    for bi := range g.Bullets {
        if !g.Bullets[bi].Active {
            continue
        }
        for ei := range g.Enemies {
            if !g.Enemies[ei].Active {
                continue
            }
            if g.Bullets[bi].Position.X == g.Enemies[ei].Position.X &&
                g.Bullets[bi].Position.Y == g.Enemies[ei].Position.Y {
                g.Bullets[bi].Active = false
                g.Enemies[ei].Health -= 25
                g.CreateExplosion(g.Enemies[ei].Position.X, g.Enemies[ei].Position.Y)
                if g.Enemies[ei].Health <= 0 {
                    g.Enemies[ei].Active = false
                    g.Player.Score += 100 * g.Level
                }
                break
            }
        }
    }
    for ei := range g.Enemies {
        if !g.Enemies[ei].Active {
            continue
        }
        if g.Enemies[ei].Position.X == g.Player.Position.X &&
            g.Enemies[ei].Position.Y == g.Player.Position.Y {
            g.Enemies[ei].Active = false
            g.Player.Health -= 30
            g.CreateExplosion(g.Player.Position.X, g.Player.Position.Y)
            if g.Player.Health <= 0 {
                g.GameOver = true
            }
        }
    }
}

func (g *Game) CreateExplosion(x, y int) {
    symbols := []rune{'*', '+', '.', 'o', 'O'}
    for i := 0; i < 10; i++ {
        for pi := range g.Particles {
            if g.Particles[pi].Life == 0 {
                g.Particles[pi].Position = Vector2{x, y}
                g.Particles[pi].Velocity = Vector2{rand.Intn(5) - 2, rand.Intn(5) - 2}
                g.Particles[pi].Life = rand.Intn(20) + 10
                g.Particles[pi].Symbol = symbols[rand.Intn(len(symbols))]
                break
            }
        }
    }
}

func (g *Game) Render() {
    screen := make([][]rune, screenHeight)
    for i := range screen {
        screen[i] = make([]rune, screenWidth)
        for j := range screen[i] {
            screen[i][j] = ' '
        }
    }
    for _, star := range g.Stars {
        if star.Position.Y >= 0 && star.Position.Y < screenHeight &&
            star.Position.X >= 0 && star.Position.X < screenWidth {
            brightnessChar := '.'
            if star.Brightness == 2 {
                brightnessChar = '*'
            }
            screen[star.Position.Y][star.Position.X] = brightnessChar
        }
    }
    for _, particle := range g.Particles {
        if particle.Life > 0 && particle.Position.Y >= 0 && particle.Position.Y < screenHeight &&
            particle.Position.X >= 0 && particle.Position.X < screenWidth {
            screen[particle.Position.Y][particle.Position.X] = particle.Symbol
        }
    }
    for _, bullet := range g.Bullets {
        if bullet.Active && bullet.Position.Y >= 0 && bullet.Position.Y < screenHeight &&
            bullet.Position.X >= 0 && bullet.Position.X < screenWidth {
            screen[bullet.Position.Y][bullet.Position.X] = bullet.Symbol
        }
    }
    for _, enemy := range g.Enemies {
        if enemy.Active && enemy.Position.Y >= 0 && enemy.Position.Y < screenHeight &&
            enemy.Position.X >= 0 && enemy.Position.X < screenWidth {
            screen[enemy.Position.Y][enemy.Position.X] = enemy.Symbol
        }
    }
    if g.Player.Position.Y >= 0 && g.Player.Position.Y < screenHeight &&
        g.Player.Position.X >= 0 && g.Player.Position.X < screenWidth {
        screen[g.Player.Position.Y][g.Player.Position.X] = g.Player.Symbol
    }
    fmt.Print("\033[H\033[2J")
    for y := 0; y < screenHeight; y++ {
        line := ""
        for x := 0; x < screenWidth; x++ {
            line += string(screen[y][x])
        }
        fmt.Println(line)
    }
    fmt.Printf("Health: %d | Score: %d | Level: %d\n", g.Player.Health, g.Player.Score, g.Level)
    fmt.Println("Controls: A-left, D-right, W-up, S-down, Space-shoot, Q-quit")
    if g.GameOver {
        fmt.Println("GAME OVER! Press Enter to exit.")
    }
}

func main() {
    game := Game{}
    game.Init()
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("\033[?25l")
    defer fmt.Print("\033[?25h")
    for !game.GameOver {
        game.Render()
        game.Update()
        time.Sleep(time.Second / gameFPS)
        if game.GameOver {
            break
        }
        fmt.Print("Input: ")
        input, _ := reader.ReadString('\n')
        game.Input = strings.TrimSpace(input)
    }
    game.Render()
    fmt.Println("Final Score:", game.Player.Score)
    fmt.Println("Press Enter to exit...")
    reader.ReadString('\n')
}
