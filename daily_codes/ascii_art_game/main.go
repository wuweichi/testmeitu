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

type GameState struct {
	Score       int
	Level       int
	Lives       int
	TimeLeft    int
	PlayerPos   int
	Enemies     []Enemy
	Projectiles []Projectile
	PowerUps    []PowerUp
	IsPaused    bool
	IsGameOver  bool
}

type Enemy struct {
	ID        int
	X         int
	Y         int
	Speed     int
	Direction int
	Health    int
	Symbol    string
}

type Projectile struct {
	X      int
	Y      int
	Speed  int
	Damage int
	Symbol string
}

type PowerUp struct {
	X      int
	Y      int
	Type   string
	Symbol string
}

const (
	ScreenWidth  = 80
	ScreenHeight = 24
	PlayerSymbol = "@"
	EmptySymbol  = " "
	WallSymbol   = "#"
	HeartSymbol  = "♥"
	StarSymbol   = "★"
)

var (
	gameState GameState
	inputChan = make(chan string, 1)
)

func initGame() {
	gameState = GameState{
		Score:     0,
		Level:     1,
		Lives:     3,
		TimeLeft:  60,
		PlayerPos: ScreenWidth / 2,
		Enemies:   []Enemy{},
		IsPaused:  false,
		IsGameOver: false,
	}
	spawnEnemies()
	spawnPowerUps()
}

func spawnEnemies() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5+gameState.Level; i++ {
		enemy := Enemy{
			ID:        i,
			X:         rand.Intn(ScreenWidth-2) + 1,
			Y:         rand.Intn(ScreenHeight/2) + 1,
			Speed:     1 + rand.Intn(2),
			Direction: rand.Intn(2)*2 - 1,
			Health:    2 + gameState.Level,
			Symbol:    "☠",
		}
		gameState.Enemies = append(gameState.Enemies, enemy)
	}
}

func spawnPowerUps() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 3; i++ {
		powerUp := PowerUp{
			X:      rand.Intn(ScreenWidth-2) + 1,
			Y:      rand.Intn(ScreenHeight-2) + 1,
			Type:   []string{"health", "time", "score"}[rand.Intn(3)],
			Symbol: []string{HeartSymbol, StarSymbol, "$"}[rand.Intn(3)],
		}
		gameState.PowerUps = append(gameState.PowerUps, powerUp)
	}
}

func updateGame() {
	if gameState.IsPaused || gameState.IsGameOver {
		return
	}
	gameState.TimeLeft--
	if gameState.TimeLeft <= 0 {
		gameState.IsGameOver = true
		return
	}
	updateEnemies()
	updateProjectiles()
	checkCollisions()
	if len(gameState.Enemies) == 0 {
		gameState.Level++
		gameState.TimeLeft += 30
		spawnEnemies()
		spawnPowerUps()
	}
}

func updateEnemies() {
	for i := range gameState.Enemies {
		gameState.Enemies[i].X += gameState.Enemies[i].Direction * gameState.Enemies[i].Speed
		if gameState.Enemies[i].X <= 0 || gameState.Enemies[i].X >= ScreenWidth-1 {
			gameState.Enemies[i].Direction *= -1
		}
		if rand.Intn(100) < 5 {
			fireProjectile(gameState.Enemies[i].X, gameState.Enemies[i].Y+1, 1, 1, "↓")
		}
	}
}

func updateProjectiles() {
	for i := 0; i < len(gameState.Projectiles); i++ {
		gameState.Projectiles[i].Y += gameState.Projectiles[i].Speed
		if gameState.Projectiles[i].Y < 0 || gameState.Projectiles[i].Y >= ScreenHeight {
			gameState.Projectiles = append(gameState.Projectiles[:i], gameState.Projectiles[i+1:]...)
			i--
		}
	}
}

func checkCollisions() {
	for i := 0; i < len(gameState.Projectiles); i++ {
		proj := &gameState.Projectiles[i]
		if proj.Y == ScreenHeight-2 && proj.X == gameState.PlayerPos {
			gameState.Lives--
			gameState.Projectiles = append(gameState.Projectiles[:i], gameState.Projectiles[i+1:]...)
			i--
			if gameState.Lives <= 0 {
				gameState.IsGameOver = true
			}
			continue
		}
		for j := 0; j < len(gameState.Enemies); j++ {
			enemy := &gameState.Enemies[j]
			if proj.X == enemy.X && proj.Y == enemy.Y {
				enemy.Health -= proj.Damage
				gameState.Projectiles = append(gameState.Projectiles[:i], gameState.Projectiles[i+1:]...)
				i--
				if enemy.Health <= 0 {
					gameState.Score += 10 * gameState.Level
					gameState.Enemies = append(gameState.Enemies[:j], gameState.Enemies[j+1:]...)
					j--
				}
				break
			}
		}
	}
	for i := 0; i < len(gameState.PowerUps); i++ {
		pUp := &gameState.PowerUps[i]
		if pUp.X == gameState.PlayerPos && pUp.Y == ScreenHeight-2 {
			switch pUp.Type {
			case "health":
				gameState.Lives++
			case "time":
				gameState.TimeLeft += 10
			case "score":
				gameState.Score += 50
			}
			gameState.PowerUps = append(gameState.PowerUps[:i], gameState.PowerUps[i+1:]...)
			i--
		}
	}
}

func fireProjectile(x, y, speed, damage int, symbol string) {
	proj := Projectile{
		X:      x,
		Y:      y,
		Speed:  speed,
		Damage: damage,
		Symbol: symbol,
	}
	gameState.Projectiles = append(gameState.Projectiles, proj)
}

func handleInput(input string) {
	switch input {
	case "a":
		if gameState.PlayerPos > 1 {
			gameState.PlayerPos--
		}
	case "d":
		if gameState.PlayerPos < ScreenWidth-2 {
			gameState.PlayerPos++
		}
	case " ":
		fireProjectile(gameState.PlayerPos, ScreenHeight-3, -1, 1, "↑")
	case "p":
		gameState.IsPaused = !gameState.IsPaused
	case "r":
		if gameState.IsGameOver {
			initGame()
		}
	case "q":
		os.Exit(0)
	}
}

func drawScreen() {
	fmt.Print("\033[2J\033[H") // Clear screen and move cursor to top-left
	var screen [ScreenHeight][ScreenWidth]string
	for y := 0; y < ScreenHeight; y++ {
		for x := 0; x < ScreenWidth; x++ {
			if y == 0 || y == ScreenHeight-1 || x == 0 || x == ScreenWidth-1 {
				screen[y][x] = WallSymbol
			} else {
				screen[y][x] = EmptySymbol
			}
		}
	}
	screen[ScreenHeight-2][gameState.PlayerPos] = PlayerSymbol
	for _, enemy := range gameState.Enemies {
		if enemy.Y >= 0 && enemy.Y < ScreenHeight && enemy.X >= 0 && enemy.X < ScreenWidth {
			screen[enemy.Y][enemy.X] = enemy.Symbol
		}
	}
	for _, proj := range gameState.Projectiles {
		if proj.Y >= 0 && proj.Y < ScreenHeight && proj.X >= 0 && proj.X < ScreenWidth {
			screen[proj.Y][proj.X] = proj.Symbol
		}
	}
	for _, pUp := range gameState.PowerUps {
		if pUp.Y >= 0 && pUp.Y < ScreenHeight && pUp.X >= 0 && pUp.X < ScreenWidth {
			screen[pUp.Y][pUp.X] = pUp.Symbol
		}
	}
	for y := 0; y < ScreenHeight; y++ {
		for x := 0; x < ScreenWidth; x++ {
			fmt.Print(screen[y][x])
		}
		fmt.Println()
	}
	drawHUD()
}

func drawHUD() {
	hud := fmt.Sprintf("Score: %d | Level: %d | Lives: %d | Time: %d",
		gameState.Score, gameState.Level, gameState.Lives, gameState.TimeLeft)
	fmt.Println(strings.Repeat("-", ScreenWidth))
	fmt.Println(hud)
	if gameState.IsPaused {
		fmt.Println("PAUSED - Press 'p' to resume")
	}
	if gameState.IsGameOver {
		fmt.Println("GAME OVER! Press 'r' to restart or 'q' to quit")
	}
	fmt.Println("Controls: A/D to move, SPACE to shoot, P to pause, Q to quit")
}

func inputListener() {
	reader := bufio.NewReader(os.Stdin)
	for {
		char, _, err := reader.ReadRune()
		if err != nil {
			close(inputChan)
			return
		}
		inputChan <- strings.ToLower(string(char))
	}
}

func main() {
	initGame()
	go inputListener()
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()
	for {
		select {
		case input := <-inputChan:
			handleInput(input)
		case <-ticker.C:
			updateGame()
			drawScreen()
		}
	}
}
