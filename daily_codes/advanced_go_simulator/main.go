package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

// GameState represents the state of the simulation
type GameState struct {
	Players      []Player
	GameObjects  []GameObject
	Time         time.Time
	Score        int
	Level        int
	IsRunning    bool
	Mutex        sync.RWMutex
}

// Player represents a player in the simulation
type Player struct {
	ID           int
	Name         string
	Health       int
	Position     Position
	Inventory    []Item
	Skills       map[string]int
	LastActive   time.Time
}

// GameObject represents an object in the game world
type GameObject struct {
	ID         int
	Type       string
	Position   Position
	Value      int
	IsActive   bool
}

// Position represents a 2D position
type Position struct {
	X int
	Y int
}

// Item represents an item that can be in inventory
type Item struct {
	ID    int
	Name  string
	Type  string
	Value int
}

// Simulation constants
const (
	MaxPlayers     = 100
	MaxObjects     = 500
	WorldWidth     = 1000
	WorldHeight    = 1000
	TickDuration   = 100 * time.Millisecond
	SaveInterval   = 30 * time.Second
	LogFile        = "simulation.log"
)

var (
	gameState GameState
	logger    *os.File
)

func initGame() {
	gameState = GameState{
		Players:     make([]Player, 0, MaxPlayers),
		GameObjects: make([]GameObject, 0, MaxObjects),
		Time:        time.Now(),
		Score:       0,
		Level:       1,
		IsRunning:   true,
	}

	// Initialize players
	for i := 0; i < 50; i++ {
		player := Player{
			ID:         i,
			Name:       fmt.Sprintf("Player%d", i),
			Health:     100,
			Position:   Position{X: rand.Intn(WorldWidth), Y: rand.Intn(WorldHeight)},
			Inventory:  make([]Item, 0),
			Skills:     make(map[string]int),
			LastActive: time.Now(),
		}
		player.Skills["combat"] = rand.Intn(100)
		player.Skills["magic"] = rand.Intn(100)
		player.Skills["stealth"] = rand.Intn(100)
		gameState.Players = append(gameState.Players, player)
	}

	// Initialize game objects
	objectTypes := []string{"treasure", "enemy", "npc", "resource", "obstacle"}
	for i := 0; i < 300; i++ {
		objType := objectTypes[rand.Intn(len(objectTypes))]
		gameObject := GameObject{
			ID:       i,
			Type:     objType,
			Position: Position{X: rand.Intn(WorldWidth), Y: rand.Intn(WorldHeight)},
			Value:    rand.Intn(1000),
			IsActive: true,
		}
		gameState.GameObjects = append(gameState.GameObjects, gameObject)
	}

	// Initialize logger
	var err error
	logger, err = os.OpenFile(LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Failed to open log file: %v\n", err)
	}
}

func gameLoop() {
	ticker := time.NewTicker(TickDuration)
	saveTicker := time.NewTicker(SaveInterval)
	defer ticker.Stop()
	defer saveTicker.Stop()

	for gameState.IsRunning {
		select {
		case <-ticker.C:
			updateGameState()
		case <-saveTicker.C:
			saveGameState()
		}
	}
}

func updateGameState() {
	gameState.Mutex.Lock()
	defer gameState.Mutex.Unlock()

	gameState.Time = time.Now()
	gameState.Score += len(gameState.Players)

	// Update players
	for i := range gameState.Players {
		// Move players randomly
		gameState.Players[i].Position.X += rand.Intn(3) - 1
		gameState.Players[i].Position.Y += rand.Intn(3) - 1

		// Keep within bounds
		if gameState.Players[i].Position.X < 0 {
			gameState.Players[i].Position.X = 0
		}
		if gameState.Players[i].Position.X >= WorldWidth {
			gameState.Players[i].Position.X = WorldWidth - 1
		}
		if gameState.Players[i].Position.Y < 0 {
			gameState.Players[i].Position.Y = 0
		}
		if gameState.Players[i].Position.Y >= WorldHeight {
			gameState.Players[i].Position.Y = WorldHeight - 1
		}

		// Update last active time
		gameState.Players[i].LastActive = gameState.Time

		// Random health changes
		if rand.Intn(100) < 5 {
			gameState.Players[i].Health += rand.Intn(20) - 10
			if gameState.Players[i].Health > 100 {
				gameState.Players[i].Health = 100
			}
			if gameState.Players[i].Health < 0 {
				gameState.Players[i].Health = 0
			}
		}

		// Skill improvements
		for skill := range gameState.Players[i].Skills {
			if rand.Intn(100) < 2 {
				gameState.Players[i].Skills[skill] += 1
				if gameState.Players[i].Skills[skill] > 100 {
					gameState.Players[i].Skills[skill] = 100
				}
			}
		}
	}

	// Update game objects
	for i := range gameState.GameObjects {
		// Randomly deactivate/reactivate objects
		if rand.Intn(1000) < 2 {
			gameState.GameObjects[i].IsActive = !gameState.GameObjects[i].IsActive
		}

		// Move some objects
		if gameState.GameObjects[i].Type == "enemy" && rand.Intn(100) < 10 {
			gameState.GameObjects[i].Position.X += rand.Intn(5) - 2
			gameState.GameObjects[i].Position.Y += rand.Intn(5) - 2
		}
	}

	// Check for level up
	if gameState.Score > gameState.Level*1000 {
		gameState.Level++
		logEvent(fmt.Sprintf("Level up to %d!", gameState.Level))
	}

	// Random events
	if rand.Intn(1000) < 2 {
		spawnNewObject()
	}
	if rand.Intn(1000) < 1 {
		spawnNewPlayer()
	}
}

func spawnNewObject() {
	objectTypes := []string{"treasure", "enemy", "npc", "resource", "obstacle"}
	objType := objectTypes[rand.Intn(len(objectTypes))]
	newObj := GameObject{
		ID:       len(gameState.GameObjects),
		Type:     objType,
		Position: Position{X: rand.Intn(WorldWidth), Y: rand.Intn(WorldHeight)},
		Value:    rand.Intn(1000),
		IsActive: true,
	}
	gameState.GameObjects = append(gameState.GameObjects, newObj)
	logEvent(fmt.Sprintf("Spawned new %s at (%d, %d)", objType, newObj.Position.X, newObj.Position.Y))
}

func spawnNewPlayer() {
	if len(gameState.Players) >= MaxPlayers {
		return
	}

	newPlayer := Player{
		ID:         len(gameState.Players),
		Name:       fmt.Sprintf("Player%d", len(gameState.Players)),
		Health:     100,
		Position:   Position{X: rand.Intn(WorldWidth), Y: rand.Intn(WorldHeight)},
		Inventory:  make([]Item, 0),
		Skills:     make(map[string]int),
		LastActive: time.Now(),
	}
	newPlayer.Skills["combat"] = rand.Intn(100)
	newPlayer.Skills["magic"] = rand.Intn(100)
	newPlayer.Skills["stealth"] = rand.Intn(100)
	gameState.Players = append(gameState.Players, newPlayer)
	logEvent(fmt.Sprintf("New player %s joined the game", newPlayer.Name))
}

func saveGameState() {
	gameState.Mutex.RLock()
	defer gameState.Mutex.RUnlock()

	data, err := json.MarshalIndent(gameState, "", "  ")
	if err != nil {
		logEvent(fmt.Sprintf("Failed to marshal game state: %v", err))
		return
	}

	filename := fmt.Sprintf("save_%s.json", time.Now().Format("20060102_150405"))
	err = os.WriteFile(filename, data, 0644)
	if err != nil {
		logEvent(fmt.Sprintf("Failed to save game state: %v", err))
	} else {
		logEvent(fmt.Sprintf("Game state saved to %s", filename))
	}
}

func logEvent(message string) {
	if logger != nil {
		logEntry := fmt.Sprintf("[%s] %s\n", time.Now().Format("2006-01-02 15:04:05"), message)
		logger.WriteString(logEntry)
	}
	fmt.Println(message)
}

func displayStatus() {
	gameState.Mutex.RLock()
	defer gameState.Mutex.RUnlock()

	fmt.Println("\n=== Game Status ===")
	fmt.Printf("Time: %s\n", gameState.Time.Format("2006-01-02 15:04:05"))
	fmt.Printf("Level: %d\n", gameState.Level)
	fmt.Printf("Score: %d\n", gameState.Score)
	fmt.Printf("Active Players: %d\n", len(gameState.Players))
	fmt.Printf("Active Objects: %d\n", countActiveObjects())

	// Show top 5 players by skill
	fmt.Println("\nTop 5 Players:")
	players := make([]Player, len(gameState.Players))
	copy(players, gameState.Players)

	// Sort by total skill points
	for i := 0; i < len(players)-1 && i < 4; i++ {
		for j := i + 1; j < len(players); j++ {
			if totalSkills(players[j]) > totalSkills(players[i]) {
				players[i], players[j] = players[j], players[i]
			}
		}
	}

	for i := 0; i < 5 && i < len(players); i++ {
		fmt.Printf("%d. %s (Health: %d, Skills: %d)\n", 
			i+1, players[i].Name, players[i].Health, totalSkills(players[i]))
	}
}

func totalSkills(player Player) int {
	total := 0
	for _, skill := range player.Skills {
		total += skill
	}
	return total
}

func countActiveObjects() int {
	count := 0
	for _, obj := range gameState.GameObjects {
		if obj.IsActive {
			count++
		}
	}
	return count
}

func handleCommands() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("\nGame Commands:")
	fmt.Println("  status - Display game status")
	fmt.Println("  save - Manually save game state")
	fmt.Println("  players - List all players")
	fmt.Println("  objects - List all game objects")
	fmt.Println("  quit - Exit the game")
	fmt.Println("\nEnter commands (one per line):")

	for scanner.Scan() {
		command := strings.TrimSpace(scanner.Text())
		switch command {
		case "status":
			displayStatus()
		case "save":
			saveGameState()
		case "players":
			listPlayers()
		case "objects":
			listObjects()
		case "quit":
			gameState.IsRunning = false
			fmt.Println("Shutting down game...")
			return
		case "":
			// Do nothing for empty input
		default:
			fmt.Printf("Unknown command: %s\n", command)
		}
	}
}

func listPlayers() {
	gameState.Mutex.RLock()
	defer gameState.Mutex.RUnlock()

	fmt.Println("\n=== All Players ===")
	for _, player := range gameState.Players {
		fmt.Printf("ID: %d, Name: %s, Health: %d, Position: (%d, %d), Last Active: %s\n",
			player.ID, player.Name, player.Health, player.Position.X, player.Position.Y,
			player.LastActive.Format("15:04:05"))
	}
}

func listObjects() {
	gameState.Mutex.RLock()
	defer gameState.Mutex.RUnlock()

	fmt.Println("\n=== All Game Objects ===")
	for _, obj := range gameState.GameObjects {
		status := "inactive"
		if obj.IsActive {
			status = "active"
		}
		fmt.Printf("ID: %d, Type: %s, Position: (%d, %d), Value: %d, Status: %s\n",
			obj.ID, obj.Type, obj.Position.X, obj.Position.Y, obj.Value, status)
	}
}

func cleanup() {
	if logger != nil {
		logger.Close()
	}
	fmt.Println("Game cleanup completed.")
}

func main() {
	fmt.Println("Starting Advanced Go Simulator...")
	fmt.Println("This is a complex simulation with multiple systems running concurrently.")
	fmt.Println("The game will run until you enter 'quit' command.")

	// Seed random number generator
	rand.Seed(time.Now().UnixNano())

	// Initialize game
	initGame()
	logEvent("Game initialized")

	// Start game loop in background
	go gameLoop()

	// Handle commands in main goroutine
	handleCommands()

	// Cleanup
	cleanup()
	fmt.Println("Game ended. Goodbye!")
}
