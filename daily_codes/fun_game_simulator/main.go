package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Player represents a game player
type Player struct {
	Name     string
	Health   int
	Strength int
	Agility  int
	Luck     int
	Level    int
	XP       int
	Gold     int
	Inventory []string
}

// Enemy represents an enemy in the game
type Enemy struct {
	Name     string
	Health   int
	Strength int
	RewardXP int
	RewardGold int
}

// Quest represents a game quest
type Quest struct {
	ID          int
	Name        string
	Description string
	Completed   bool
	RewardXP    int
	RewardGold  int
}

// GameState holds the overall game state
type GameState struct {
	Player      Player
	Enemies     []Enemy
	Quests      []Quest
	CurrentArea string
	GameTime    time.Time
	SaveFile    string
}

// Global game state
var gameState GameState

// Constants for game configuration
const (
	MaxHealth   = 100
	MaxStrength = 20
	MaxAgility  = 20
	MaxLuck     = 20
	XPPerLevel  = 100
	SaveFileName = "game_save.json"
)

// Initialize the game with default values
func initGame() {
	gameState = GameState{
		Player: Player{
			Name:     "Hero",
			Health:   MaxHealth,
			Strength: 10,
			Agility:  10,
			Luck:     10,
			Level:    1,
			XP:       0,
			Gold:     50,
			Inventory: []string{"Sword", "Health Potion"},
		},
		Enemies: []Enemy{
			{Name: "Goblin", Health: 30, Strength: 5, RewardXP: 20, RewardGold: 10},
			{Name: "Orc", Health: 50, Strength: 8, RewardXP: 40, RewardGold: 25},
			{Name: "Dragon", Health: 100, Strength: 15, RewardXP: 100, RewardGold: 100},
		},
		Quests: []Quest{
			{ID: 1, Name: "Slay the Goblin", Description: "Defeat the goblin in the forest", Completed: false, RewardXP: 50, RewardGold: 30},
			{ID: 2, Name: "Collect Herbs", Description: "Find 3 medicinal herbs in the meadow", Completed: false, RewardXP: 30, RewardGold: 20},
			{ID: 3, Name: "Defeat the Dragon", Description: "Slay the mighty dragon", Completed: false, RewardXP: 200, RewardGold: 150},
		},
		CurrentArea: "Town",
		GameTime:   time.Now(),
		SaveFile:   SaveFileName,
	}
}

// Display player status
func showStatus() {
	p := gameState.Player
	fmt.Println("\n=== Player Status ===")
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Health: %d/%d\n", p.Health, MaxHealth)
	fmt.Printf("Strength: %d\n", p.Strength)
	fmt.Printf("Agility: %d\n", p.Agility)
	fmt.Printf("Luck: %d\n", p.Luck)
	fmt.Printf("Level: %d\n", p.Level)
	fmt.Printf("XP: %d/%d\n", p.XP, p.Level*XPPerLevel)
	fmt.Printf("Gold: %d\n", p.Gold)
	fmt.Printf("Inventory: %v\n", p.Inventory)
	fmt.Printf("Current Area: %s\n", gameState.CurrentArea)
}

// Display available quests
func showQuests() {
	fmt.Println("\n=== Available Quests ===")
	for _, q := range gameState.Quests {
		status := "Incomplete"
		if q.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s - %s [%s]\n", q.ID, q.Name, q.Description, status)
	}
}

// Simulate a battle with an enemy
func battle(enemyIndex int) {
	if enemyIndex < 0 || enemyIndex >= len(gameState.Enemies) {
		fmt.Println("Invalid enemy selection.")
		return
	}
	
	enemy := gameState.Enemies[enemyIndex]
	player := &gameState.Player
	
	fmt.Printf("\n=== Battle with %s ===\n", enemy.Name)
	fmt.Printf("Your Health: %d, Enemy Health: %d\n", player.Health, enemy.Health)
	
	for player.Health > 0 && enemy.Health > 0 {
		// Player attack
		playerDamage := player.Strength + rand.Intn(6)
		enemy.Health -= playerDamage
		fmt.Printf("You attack %s for %d damage. Enemy health: %d\n", enemy.Name, playerDamage, enemy.Health)
		
		if enemy.Health <= 0 {
			fmt.Printf("You defeated %s!\n", enemy.Name)
			player.XP += enemy.RewardXP
			player.Gold += enemy.RewardGold
			fmt.Printf("Gained %d XP and %d gold.\n", enemy.RewardXP, enemy.RewardGold)
			checkLevelUp()
			// Check if any quests are completed
			checkQuestCompletion(enemy.Name)
			break
		}
		
		// Enemy attack
		enemyDamage := enemy.Strength + rand.Intn(4)
		player.Health -= enemyDamage
		fmt.Printf("%s attacks you for %d damage. Your health: %d\n", enemy.Name, enemyDamage, player.Health)
		
		if player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			os.Exit(0)
		}
		
		time.Sleep(1 * time.Second)
	}
}

// Check if player levels up
func checkLevelUp() {
	player := &gameState.Player
	requiredXP := player.Level * XPPerLevel
	if player.XP >= requiredXP {
		player.Level++
		player.Health = MaxHealth
		player.Strength += 2
		player.Agility += 1
		player.Luck += 1
		fmt.Printf("\nCongratulations! You reached level %d!\n", player.Level)
		fmt.Println("Health restored, attributes increased.")
	}
}

// Check quest completion based on enemy name
func checkQuestCompletion(enemyName string) {
	for i := range gameState.Quests {
		q := &gameState.Quests[i]
		if !q.Completed && strings.Contains(q.Description, enemyName) {
			q.Completed = true
			gameState.Player.XP += q.RewardXP
			gameState.Player.Gold += q.RewardGold
			fmt.Printf("Quest '%s' completed! Gained %d XP and %d gold.\n", q.Name, q.RewardXP, q.RewardGold)
		}
	}
}

// Save game state to file
func saveGame() {
	data, err := json.MarshalIndent(gameState, "", "  ")
	if err != nil {
		fmt.Println("Error saving game:", err)
		return
	}
	err = os.WriteFile(gameState.SaveFile, data, 0644)
	if err != nil {
		fmt.Println("Error writing save file:", err)
		return
	}
	fmt.Println("Game saved successfully.")
}

// Load game state from file
func loadGame() {
	data, err := os.ReadFile(gameState.SaveFile)
	if err != nil {
		fmt.Println("No save file found, starting new game.")
		return
	}
	err = json.Unmarshal(data, &gameState)
	if err != nil {
		fmt.Println("Error loading save file:", err)
		return
	}
	fmt.Println("Game loaded successfully.")
}

// Display game help
func showHelp() {
	fmt.Println("\n=== Game Commands ===")
	fmt.Println("status    - Show player status")
	fmt.Println("quests    - Show available quests")
	fmt.Println("battle    - Start a battle with an enemy")
	fmt.Println("areas     - Show available areas to explore")
	fmt.Println("explore   - Explore the current area")
	fmt.Println("shop      - Visit the shop to buy items")
	fmt.Println("save      - Save the current game")
	fmt.Println("load      - Load a saved game")
	fmt.Println("help      - Show this help message")
	fmt.Println("exit      - Exit the game")
}

// Show available areas
func showAreas() {
	fmt.Println("\n=== Available Areas ===")
	fmt.Println("1. Town - A safe hub with shops and quests")
	fmt.Println("2. Forest - Home to goblins and herbs")
	fmt.Println("3. Cave - Dangerous area with orcs")
	fmt.Println("4. Mountain - Lair of the dragon")
}

// Change current area
func changeArea(areaName string) {
	gameState.CurrentArea = areaName
	fmt.Printf("You have traveled to %s.\n", areaName)
}

// Explore the current area
func exploreArea() {
	area := gameState.CurrentArea
	fmt.Printf("\nExploring %s...\n", area)
	
	rand.Seed(time.Now().UnixNano())
	event := rand.Intn(5)
	
	switch event {
	case 0:
		fmt.Println("You found a hidden treasure chest!")
		goldFound := 10 + rand.Intn(20)
		gameState.Player.Gold += goldFound
		fmt.Printf("Gained %d gold.\n", goldFound)
	case 1:
		fmt.Println("You discovered a medicinal herb.")
		gameState.Player.Inventory = append(gameState.Player.Inventory, "Medicinal Herb")
	case 2:
		fmt.Println("A wild animal attacks!")
		// Simulate a minor battle
		gameState.Player.Health -= 5
		fmt.Println("Lost 5 health in the struggle.")
	case 3:
		fmt.Println("You meet a friendly traveler who gives you advice.")
		fmt.Println("Traveler: 'Beware of the dragon in the mountains!'")
	case 4:
		fmt.Println("Nothing of interest found.")
	}
}

// Simulate a shop where player can buy items
func visitShop() {
	if gameState.CurrentArea != "Town" {
		fmt.Println("Shop is only available in Town.")
		return
	}
	
	fmt.Println("\n=== Welcome to the Shop ===")
	fmt.Println("1. Health Potion - 20 gold (Restores 30 health)")
	fmt.Println("2. Strength Elixir - 50 gold (+2 Strength)")
	fmt.Println("3. Lucky Charm - 30 gold (+1 Luck)")
	fmt.Println("4. Exit shop")
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose an item to buy (1-4): ")
	choiceStr, _ := reader.ReadString('\n')
	choiceStr = strings.TrimSpace(choiceStr)
	choice, err := strconv.Atoi(choiceStr)
	if err != nil || choice < 1 || choice > 4 {
		fmt.Println("Invalid choice.")
		return
	}
	
	switch choice {
	case 1:
		if gameState.Player.Gold >= 20 {
			gameState.Player.Gold -= 20
			gameState.Player.Inventory = append(gameState.Player.Inventory, "Health Potion")
			fmt.Println("Purchased Health Potion.")
		} else {
			fmt.Println("Not enough gold.")
		}
	case 2:
		if gameState.Player.Gold >= 50 {
			gameState.Player.Gold -= 50
			gameState.Player.Strength += 2
			fmt.Println("Purchased Strength Elixir. Strength increased by 2.")
		} else {
			fmt.Println("Not enough gold.")
		}
	case 3:
		if gameState.Player.Gold >= 30 {
			gameState.Player.Gold -= 30
			gameState.Player.Luck += 1
			fmt.Println("Purchased Lucky Charm. Luck increased by 1.")
		} else {
			fmt.Println("Not enough gold.")
		}
	case 4:
		fmt.Println("Exiting shop.")
	}
}

// Main game loop
func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("=== Welcome to Fun Game Simulator ===")
	fmt.Println("Type 'help' for a list of commands.")
	
	// Initialize or load game
	initGame()
	loadGame()
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Print("\nEnter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		switch input {
		case "status":
			showStatus()
		case "quests":
			showQuests()
		case "battle":
			fmt.Println("\nSelect an enemy to battle:")
			for i, e := range gameState.Enemies {
				fmt.Printf("%d. %s (Health: %d, Strength: %d)\n", i+1, e.Name, e.Health, e.Strength)
			}
			fmt.Print("Enter enemy number: ")
			enemyStr, _ := reader.ReadString('\n')
			enemyStr = strings.TrimSpace(enemyStr)
			enemyNum, err := strconv.Atoi(enemyStr)
			if err != nil || enemyNum < 1 || enemyNum > len(gameState.Enemies) {
				fmt.Println("Invalid selection.")
				break
			}
			battle(enemyNum - 1)
		case "areas":
			showAreas()
		case "explore":
			exploreArea()
		case "shop":
			visitShop()
		case "save":
			saveGame()
		case "load":
			loadGame()
		case "help":
			showHelp()
		case "exit":
			fmt.Println("Thanks for playing! Goodbye.")
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}
}
