package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Game constants
const (
	MaxPlayers = 10
	MaxRounds = 100
	BoardSize = 20
)

// Player represents a game participant
type Player struct {
	ID          int
	Name        string
	Score       int
	Position    int
	IsActive    bool
	SpecialPowers []string
}

// GameState holds the current state of the game
type GameState struct {
	Round      int
	Players    []Player
	Board      []string
	GameLog    []string
	IsGameOver bool
}

// GameEvent represents an event that occurs during gameplay
type GameEvent struct {
	Type        string
	Description string
	PlayerID    int
	Round       int
}

// GameManager handles the game logic
type GameManager struct {
	State      GameState
	Events     []GameEvent
	RandSource *rand.Rand
}

// InitializeGame sets up a new game
func (gm *GameManager) InitializeGame(playerNames []string) {
	gm.RandSource = rand.New(rand.NewSource(time.Now().UnixNano()))
	gm.State.Round = 1
	gm.State.Board = make([]string, BoardSize)
	for i := range gm.State.Board {
		gm.State.Board[i] = "_"
	}
	gm.State.Players = make([]Player, len(playerNames))
	for i, name := range playerNames {
		gm.State.Players[i] = Player{
			ID:          i + 1,
			Name:        name,
			Score:       0,
			Position:    gm.RandSource.Intn(BoardSize),
			IsActive:    true,
			SpecialPowers: []string{"Double Points", "Teleport", "Shield"},
		}
	}
	gm.logEvent("Game Initialized", "Game started with "+strconv.Itoa(len(playerNames))+" players")
}

// PlayRound simulates one round of the game
func (gm *GameManager) PlayRound() {
	if gm.State.IsGameOver || gm.State.Round > MaxRounds {
		gm.State.IsGameOver = true
		return
	}

	gm.logEvent("Round Start", "Round "+strconv.Itoa(gm.State.Round)+" begins")

	// Process each active player
	for i := range gm.State.Players {
		if !gm.State.Players[i].IsActive {
			continue
		}

		// Move player
		oldPos := gm.State.Players[i].Position
		move := gm.RandSource.Intn(6) + 1 // Dice roll 1-6
		gm.State.Players[i].Position = (oldPos + move) % BoardSize

		// Check for special events
		if gm.RandSource.Float32() < 0.2 {
			gm.triggerSpecialEvent(&gm.State.Players[i])
		}

		// Update board
		gm.updateBoard()

		// Score points
		points := gm.calculatePoints(&gm.State.Players[i])
		gm.State.Players[i].Score += points

		gm.logEvent("Player Move", fmt.Sprintf("%s moved from %d to %d, scored %d points",
			gm.State.Players[i].Name, oldPos, gm.State.Players[i].Position, points))
	}

	// Check for eliminations
	gm.checkEliminations()

	gm.State.Round++
	if gm.State.Round > MaxRounds || len(gm.getActivePlayers()) <= 1 {
		gm.State.IsGameOver = true
		gm.logEvent("Game Over", "Game ended after "+strconv.Itoa(gm.State.Round-1)+" rounds")
	}
}

// calculatePoints determines points earned by a player
func (gm *GameManager) calculatePoints(player *Player) int {
	basePoints := gm.RandSource.Intn(10) + 1
	// Bonus for landing on special positions
	if player.Position%5 == 0 {
		basePoints *= 2
	}
	// Check for special power activation
	if len(player.SpecialPowers) > 0 && gm.RandSource.Float32() < 0.3 {
		power := player.SpecialPowers[gm.RandSource.Intn(len(player.SpecialPowers))]
		gm.logEvent("Power Activated", fmt.Sprintf("%s used %s", player.Name, power))
		if power == "Double Points" {
			basePoints *= 2
		} else if power == "Teleport" {
			player.Position = gm.RandSource.Intn(BoardSize)
		} else if power == "Shield" {
			// Shield prevents elimination in next round
			gm.logEvent("Shield Activated", player.Name+" is shielded for one round")
		}
	}
	return basePoints
}

// triggerSpecialEvent triggers a random special event
func (gm *GameManager) triggerSpecialEvent(player *Player) {
	events := []string{
		"Found a treasure chest!",
		"Encountered a wild monster!",
		"Stumbled upon a secret passage!",
		"Weather changed dramatically!",
		"Mysterious stranger appears!",
	}
	event := events[gm.RandSource.Intn(len(events))]
	gm.logEvent("Special Event", fmt.Sprintf("%s: %s", player.Name, event))
	// Random effect
	effect := gm.RandSource.Intn(3)
	switch effect {
	case 0:
		player.Score += 20
	case 1:
		player.Position = (player.Position + 10) % BoardSize
	case 2:
		player.SpecialPowers = append(player.SpecialPowers, "Extra Life")
	}
}

// checkEliminations removes players with low scores
func (gm *GameManager) checkEliminations() {
	activePlayers := gm.getActivePlayers()
	if len(activePlayers) <= 1 {
		return
	}
	// Find lowest score
	lowestScore := activePlayers[0].Score
	for _, p := range activePlayers {
		if p.Score < lowestScore {
			lowestScore = p.Score
		}
	}
	// Eliminate players with lowest score (if more than one)
	for i := range gm.State.Players {
		if gm.State.Players[i].IsActive && gm.State.Players[i].Score == lowestScore {
			gm.State.Players[i].IsActive = false
			gm.logEvent("Elimination", fmt.Sprintf("%s was eliminated with score %d",
				gm.State.Players[i].Name, lowestScore))
		}
	}
}

// updateBoard updates the visual representation of the board
func (gm *GameManager) updateBoard() {
	// Reset board
	for i := range gm.State.Board {
		gm.State.Board[i] = "_"
	}
	// Place players on board
	for _, player := range gm.State.Players {
		if player.IsActive {
			gm.State.Board[player.Position] = strconv.Itoa(player.ID)
		}
	}
}

// getActivePlayers returns a slice of active players
func (gm *GameManager) getActivePlayers() []Player {
	var active []Player
	for _, p := range gm.State.Players {
		if p.IsActive {
			active = append(active, p)
		}
	}
	return active
}

// logEvent adds an event to the game log
func (gm *GameManager) logEvent(eventType, description string) {
	event := GameEvent{
		Type:        eventType,
		Description: description,
		Round:       gm.State.Round,
	}
	gm.Events = append(gm.Events, event)
	gm.State.GameLog = append(gm.State.GameLog, fmt.Sprintf("Round %d: [%s] %s",
		gm.State.Round, eventType, description))
}

// DisplayGameState prints the current game state
func (gm *GameManager) DisplayGameState() {
	fmt.Println("\n=== Game State ===")
	fmt.Printf("Round: %d\n", gm.State.Round)
	fmt.Println("Board:", strings.Join(gm.State.Board, " "))
	fmt.Println("Players:")
	for _, p := range gm.State.Players {
		status := "Active"
		if !p.IsActive {
			status = "Eliminated"
		}
		fmt.Printf("  %d. %s - Score: %d, Position: %d, Status: %s\n",
			p.ID, p.Name, p.Score, p.Position, status)
	}
	fmt.Println("==================")
}

// DisplayGameLog prints the game event log
func (gm *GameManager) DisplayGameLog() {
	fmt.Println("\n=== Game Log ===")
	for _, entry := range gm.State.GameLog {
		fmt.Println(entry)
	}
	fmt.Println("================")
}

// SaveGameState saves the current game state to a JSON file
func (gm *GameManager) SaveGameState(filename string) error {
	data, err := json.MarshalIndent(gm.State, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// LoadGameState loads a game state from a JSON file
func (gm *GameManager) LoadGameState(filename string) error {
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &gm.State)
}

// SimulateFullGame runs a complete game simulation
func (gm *GameManager) SimulateFullGame() {
	for !gm.State.IsGameOver {
		gm.PlayRound()
		if gm.State.Round%10 == 0 {
			gm.DisplayGameState()
		}
	}
	gm.DisplayGameState()
	gm.DisplayGameLog()
	gm.announceWinner()
}

// announceWinner declares the game winner
func (gm *GameManager) announceWinner() {
	activePlayers := gm.getActivePlayers()
	if len(activePlayers) == 0 {
		fmt.Println("\nNo winner - all players eliminated!")
		return
	}
	winner := activePlayers[0]
	for _, p := range activePlayers {
		if p.Score > winner.Score {
			winner = p
		}
	}
	fmt.Printf("\n🎉 Winner: %s with %d points! 🎉\n", winner.Name, winner.Score)
}

// Helper function to generate random player names
func generatePlayerNames(count int) []string {
	names := []string{
		"Alice", "Bob", "Charlie", "Diana", "Eve",
		"Frank", "Grace", "Henry", "Ivy", "Jack",
		"Kara", "Leo", "Mona", "Nina", "Oscar",
		"Paul", "Quinn", "Rita", "Sam", "Tina",
	}
	if count > len(names) {
		count = len(names)
	}
	return names[:count]
}

// Main function - entry point of the program
func main() {
	fmt.Println("=== Advanced Go Game Simulator ===")
	fmt.Println("A complex game simulation with multiple players, rounds, and events")

	// Initialize game manager
	gm := &GameManager{}

	// Get number of players
	playerCount := 5
	if len(os.Args) > 1 {
		if count, err := strconv.Atoi(os.Args[1]); err == nil && count > 0 && count <= MaxPlayers {
			playerCount = count
		}
	}

	// Generate player names
	playerNames := generatePlayerNames(playerCount)

	// Initialize game
	gm.InitializeGame(playerNames)
	gm.DisplayGameState()

	// Ask user for simulation mode
	fmt.Print("\nChoose mode: (1) Auto-simulate full game, (2) Step through rounds: ")
	var choice string
	fmt.Scanln(&choice)

	if choice == "1" {
		// Auto-simulate
		gm.SimulateFullGame()
	} else {
		// Step through rounds
		fmt.Println("\nStepping through rounds. Press Enter to continue, 'q' to quit.")
		for !gm.State.IsGameOver {
			gm.PlayRound()
			gm.DisplayGameState()
			fmt.Print("Press Enter to continue...")
			var input string
			fmt.Scanln(&input)
			if strings.ToLower(input) == "q" {
				break
			}
		}
		if gm.State.IsGameOver {
			gm.announceWinner()
		}
	}

	// Save game state
	if err := gm.SaveGameState("game_state.json"); err != nil {
		fmt.Printf("Error saving game: %v\n", err)
	} else {
		fmt.Println("Game state saved to game_state.json")
	}

	// Demonstrate loading (optional)
	fmt.Print("\nLoad saved game? (y/n): ")
	var loadChoice string
	fmt.Scanln(&loadChoice)
	if strings.ToLower(loadChoice) == "y" {
		loadedGM := &GameManager{}
		if err := loadedGM.LoadGameState("game_state.json"); err != nil {
			fmt.Printf("Error loading game: %v\n", err)
		} else {
			fmt.Println("Game loaded successfully!")
			loadedGM.DisplayGameState()
		}
	}

	fmt.Println("\n=== Game Simulator Ended ===")
}
