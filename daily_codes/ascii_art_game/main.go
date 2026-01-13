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
	Score      int
	Level      int
	Lives      int
	TimeLeft   int
	IsRunning  bool
	PlayerX    int
	PlayerY    int
	Enemies    []Enemy
	Projectiles []Projectile
	PowerUps   []PowerUp
	Obstacles  []Obstacle
	Messages   []string
}

type Enemy struct {
	X, Y       int
	Symbol     string
	Speed      int
	Health     int
	Direction  int
	IsActive   bool
}

type Projectile struct {
	X, Y       int
	Symbol     string
	Speed      int
	Direction  int
	IsActive   bool
}

type PowerUp struct {
	X, Y       int
	Symbol     string
	Type       string
	IsActive   bool
}

type Obstacle struct {
	X, Y       int
	Symbol     string
	IsActive   bool
}

const (
	ScreenWidth  = 80
	ScreenHeight = 24
	PlayerSymbol = "@"
	EnemySymbol  = "*"
	ProjectileSymbol = "|"
	PowerUpSymbol = "+"
	ObstacleSymbol = "#"
	InitialLives = 3
	InitialTime = 60
)

var gameState GameState

func initGame() {
	rand.Seed(time.Now().UnixNano())
	gameState = GameState{
		Score:      0,
		Level:      1,
		Lives:      InitialLives,
		TimeLeft:   InitialTime,
		IsRunning:  true,
		PlayerX:    ScreenWidth / 2,
		PlayerY:    ScreenHeight - 2,
		Enemies:    []Enemy{},
		Projectiles: []Projectile{},
		PowerUps:   []PowerUp{},
		Obstacles:  []Obstacle{},
		Messages:   []string{"Welcome to ASCII Art Game!", "Use WASD to move, Space to shoot."},
	}
	spawnEnemies(5)
	spawnPowerUps(2)
	spawnObstacles(10)
}

func spawnEnemies(count int) {
	for i := 0; i < count; i++ {
		x := rand.Intn(ScreenWidth-2) + 1
		y := rand.Intn(ScreenHeight/2) + 1
		enemy := Enemy{
			X:         x,
			Y:         y,
			Symbol:    EnemySymbol,
			Speed:     1,
			Health:    1,
			Direction: rand.Intn(4),
			IsActive:  true,
		}
		gameState.Enemies = append(gameState.Enemies, enemy)
	}
}

func spawnPowerUps(count int) {
	for i := 0; i < count; i++ {
		x := rand.Intn(ScreenWidth-2) + 1
		y := rand.Intn(ScreenHeight-2) + 1
		powerUp := PowerUp{
			X:        x,
			Y:        y,
			Symbol:   PowerUpSymbol,
			Type:     []string{"health", "time", "score"}[rand.Intn(3)],
			IsActive: true,
		}
		gameState.PowerUps = append(gameState.PowerUps, powerUp)
	}
}

func spawnObstacles(count int) {
	for i := 0; i < count; i++ {
		x := rand.Intn(ScreenWidth-2) + 1
		y := rand.Intn(ScreenHeight-2) + 1
		obstacle := Obstacle{
			X:        x,
			Y:        y,
			Symbol:   ObstacleSymbol,
			IsActive: true,
		}
		gameState.Obstacles = append(gameState.Obstacles, obstacle)
	}
}

func updateGame() {
	if !gameState.IsRunning {
		return
	}
	gameState.TimeLeft--
	if gameState.TimeLeft <= 0 {
		gameState.Messages = append(gameState.Messages, "Time's up! Game Over.")
		gameState.IsRunning = false
		return
	}
	updateEnemies()
	updateProjectiles()
	checkCollisions()
	if len(gameState.Enemies) == 0 {
		gameState.Level++
		gameState.Messages = append(gameState.Messages, "Level up! Level "+strconv.Itoa(gameState.Level))
		spawnEnemies(5 + gameState.Level)
		spawnPowerUps(2)
		spawnObstacles(10)
	}
}

func updateEnemies() {
	for i := range gameState.Enemies {
		if !gameState.Enemies[i].IsActive {
			continue
		}
		switch gameState.Enemies[i].Direction {
		case 0:
			gameState.Enemies[i].Y -= gameState.Enemies[i].Speed
		case 1:
			gameState.Enemies[i].X += gameState.Enemies[i].Speed
		case 2:
			gameState.Enemies[i].Y += gameState.Enemies[i].Speed
		case 3:
			gameState.Enemies[i].X -= gameState.Enemies[i].Speed
		}
		if gameState.Enemies[i].X < 1 || gameState.Enemies[i].X >= ScreenWidth-1 ||
			gameState.Enemies[i].Y < 1 || gameState.Enemies[i].Y >= ScreenHeight-1 {
			gameState.Enemies[i].Direction = rand.Intn(4)
		}
	}
}

func updateProjectiles() {
	for i := range gameState.Projectiles {
		if !gameState.Projectiles[i].IsActive {
			continue
		}
		switch gameState.Projectiles[i].Direction {
		case 0:
			gameState.Projectiles[i].Y -= gameState.Projectiles[i].Speed
		case 1:
			gameState.Projectiles[i].X += gameState.Projectiles[i].Speed
		case 2:
			gameState.Projectiles[i].Y += gameState.Projectiles[i].Speed
		case 3:
			gameState.Projectiles[i].X -= gameState.Projectiles[i].Speed
		}
		if gameState.Projectiles[i].X < 0 || gameState.Projectiles[i].X >= ScreenWidth ||
			gameState.Projectiles[i].Y < 0 || gameState.Projectiles[i].Y >= ScreenHeight {
			gameState.Projectiles[i].IsActive = false
		}
	}
	activeProjectiles := []Projectile{}
	for _, p := range gameState.Projectiles {
		if p.IsActive {
			activeProjectiles = append(activeProjectiles, p)
		}
	}
	gameState.Projectiles = activeProjectiles
}

func checkCollisions() {
	for i, enemy := range gameState.Enemies {
		if !enemy.IsActive {
			continue
		}
		if enemy.X == gameState.PlayerX && enemy.Y == gameState.PlayerY {
			gameState.Lives--
			gameState.Messages = append(gameState.Messages, "Hit by enemy! Lives: "+strconv.Itoa(gameState.Lives))
			gameState.Enemies[i].IsActive = false
			if gameState.Lives <= 0 {
				gameState.Messages = append(gameState.Messages, "No lives left! Game Over.")
				gameState.IsRunning = false
			}
		}
		for j, projectile := range gameState.Projectiles {
			if !projectile.IsActive {
				continue
			}
			if enemy.X == projectile.X && enemy.Y == projectile.Y {
				gameState.Enemies[i].Health--
				gameState.Projectiles[j].IsActive = false
				if gameState.Enemies[i].Health <= 0 {
					gameState.Enemies[i].IsActive = false
					gameState.Score += 10
					gameState.Messages = append(gameState.Messages, "Enemy destroyed! Score: "+strconv.Itoa(gameState.Score))
				}
			}
		}
	}
	activeEnemies := []Enemy{}
	for _, e := range gameState.Enemies {
		if e.IsActive {
			activeEnemies = append(activeEnemies, e)
		}
	}
	gameState.Enemies = activeEnemies
	for i, powerUp := range gameState.PowerUps {
		if !powerUp.IsActive {
			continue
		}
		if powerUp.X == gameState.PlayerX && powerUp.Y == gameState.PlayerY {
			gameState.PowerUps[i].IsActive = false
			switch powerUp.Type {
			case "health":
				gameState.Lives++
				gameState.Messages = append(gameState.Messages, "Health power-up! Lives: "+strconv.Itoa(gameState.Lives))
			case "time":
				gameState.TimeLeft += 10
				gameState.Messages = append(gameState.Messages, "Time power-up! Time left: "+strconv.Itoa(gameState.TimeLeft))
			case "score":
				gameState.Score += 50
				gameState.Messages = append(gameState.Messages, "Score power-up! Score: "+strconv.Itoa(gameState.Score))
			}
		}
	}
	activePowerUps := []PowerUp{}
	for _, p := range gameState.PowerUps {
		if p.IsActive {
			activePowerUps = append(activePowerUps, p)
		}
	}
	gameState.PowerUps = activePowerUps
	for _, obstacle := range gameState.Obstacles {
		if obstacle.X == gameState.PlayerX && obstacle.Y == gameState.PlayerY {
			gameState.Messages = append(gameState.Messages, "Hit obstacle!")
		}
	}
}

func shootProjectile() {
	if !gameState.IsRunning {
		return
	}
	projectile := Projectile{
		X:         gameState.PlayerX,
		Y:         gameState.PlayerY - 1,
		Symbol:    ProjectileSymbol,
		Speed:     2,
		Direction: 0,
		IsActive:  true,
	}
	gameState.Projectiles = append(gameState.Projectiles, projectile)
}

func movePlayer(dx, dy int) {
	if !gameState.IsRunning {
		return
	}
	newX := gameState.PlayerX + dx
	newY := gameState.PlayerY + dy
	if newX >= 0 && newX < ScreenWidth && newY >= 0 && newY < ScreenHeight {
		gameState.PlayerX = newX
		gameState.PlayerY = newY
	}
}

func renderScreen() {
	fmt.Print("\033[2J\033[H")
	fmt.Println("ASCII Art Game - Score:", gameState.Score, "Level:", gameState.Level, "Lives:", gameState.Lives, "Time:", gameState.TimeLeft)
	fmt.Println(strings.Repeat("=", ScreenWidth))
	screen := make([][]string, ScreenHeight)
	for i := range screen {
		screen[i] = make([]string, ScreenWidth)
		for j := range screen[i] {
			screen[i][j] = " "
		}
	}
	screen[gameState.PlayerY][gameState.PlayerX] = PlayerSymbol
	for _, enemy := range gameState.Enemies {
		if enemy.IsActive {
			screen[enemy.Y][enemy.X] = enemy.Symbol
		}
	}
	for _, projectile := range gameState.Projectiles {
		if projectile.IsActive {
			screen[projectile.Y][projectile.X] = projectile.Symbol
		}
	}
	for _, powerUp := range gameState.PowerUps {
		if powerUp.IsActive {
			screen[powerUp.Y][powerUp.X] = powerUp.Symbol
		}
	}
	for _, obstacle := range gameState.Obstacles {
		if obstacle.IsActive {
			screen[obstacle.Y][obstacle.X] = obstacle.Symbol
		}
	}
	for i := 0; i < ScreenHeight; i++ {
		for j := 0; j < ScreenWidth; j++ {
			fmt.Print(screen[i][j])
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("=", ScreenWidth))
	for _, msg := range gameState.Messages {
		fmt.Println(msg)
	}
	gameState.Messages = []string{}
}

func handleInput(input string) {
	switch input {
	case "w":
		movePlayer(0, -1)
	case "a":
		movePlayer(-1, 0)
	case "s":
		movePlayer(0, 1)
	case "d":
		movePlayer(1, 0)
	case " ":
		shootProjectile()
	case "q":
		gameState.IsRunning = false
		gameState.Messages = append(gameState.Messages, "Game quit.")
	}
}

func main() {
	initGame()
	reader := bufio.NewReader(os.Stdin)
	inputChan := make(chan string)
	go func() {
		for {
			char, _, err := reader.ReadRune()
			if err != nil {
				close(inputChan)
				return
			}
			inputChan <- string(char)
		}
	}()
	gameLoop:
	for {
		select {
		case input := <-inputChan:
			handleInput(strings.ToLower(input))
		default:
			updateGame()
			renderScreen()
			if !gameState.IsRunning {
				break gameLoop
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	fmt.Println("Final Score:", gameState.Score, "Level:", gameState.Level)
	fmt.Println("Thanks for playing!")
}
