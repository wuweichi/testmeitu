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

type Player struct {
	Name     string
	Health   int
	Strength int
	Agility  int
	Level    int
	Exp      int
	Gold     int
	Inventory []Item
}

type Item struct {
	Name        string
	Description string
	Value       int
	Rarity      string
}

type Enemy struct {
	Name     string
	Health   int
	Strength int
	Agility  int
	Reward   Reward
}

type Reward struct {
	Exp  int
	Gold int
	Item Item
}

type GameState struct {
	Player      Player
	Enemies     []Enemy
	Locations   []Location
	CurrentTime time.Time
	DayCount    int
}

type Location struct {
	Name        string
	Description string
	Enemies     []Enemy
	Items       []Item
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	game := initializeGame()
	fmt.Println("Welcome to the Complex Game Simulator!")
	fmt.Println("Type 'help' for a list of commands.")
	gameLoop(game)
}

func initializeGame() *GameState {
	player := Player{
		Name:     "Hero",
		Health:   100,
		Strength: 10,
		Agility:  10,
		Level:    1,
		Exp:      0,
		Gold:     50,
		Inventory: []Item{
			{Name: "Health Potion", Description: "Restores 50 health", Value: 25, Rarity: "Common"},
		},
	}

	enemies := generateEnemies()
	locations := generateLocations()

	return &GameState{
		Player:      player,
		Enemies:     enemies,
		Locations:   locations,
		CurrentTime: time.Now(),
		DayCount:    1,
	}
}

func generateEnemies() []Enemy {
	return []Enemy{
		{
			Name:     "Goblin",
			Health:   30,
			Strength: 5,
			Agility:  8,
			Reward:   Reward{Exp: 10, Gold: 5, Item: Item{Name: "Goblin Ear", Description: "A trophy from a goblin", Value: 2, Rarity: "Common"}},
		},
		{
			Name:     "Orc",
			Health:   60,
			Strength: 12,
			Agility:  4,
			Reward:   Reward{Exp: 25, Gold: 15, Item: Item{Name: "Orc Tooth", Description: "A sharp orc tooth", Value: 10, Rarity: "Uncommon"}},
		},
		{
			Name:     "Dragon",
			Health:   200,
			Strength: 25,
			Agility:  15,
			Reward:   Reward{Exp: 100, Gold: 100, Item: Item{Name: "Dragon Scale", Description: "A rare dragon scale", Value: 100, Rarity: "Rare"}},
		},
	}
}

func generateLocations() []Location {
	return []Location{
		{
			Name:        "Forest",
			Description: "A dense forest with various creatures.",
			Enemies:     []Enemy{generateEnemies()[0], generateEnemies()[1]},
			Items: []Item{
				{Name: "Herbs", Description: "Medicinal herbs for healing", Value: 5, Rarity: "Common"},
			},
		},
		{
			Name:        "Cave",
			Description: "A dark cave with dangerous enemies.",
			Enemies:     []Enemy{generateEnemies()[1], generateEnemies()[2]},
			Items: []Item{
				{Name: "Treasure Chest", Description: "A chest containing valuable items", Value: 50, Rarity: "Rare"},
			},
		},
		{
			Name:        "Village",
			Description: "A peaceful village with shops and NPCs.",
			Enemies:     []Enemy{},
			Items: []Item{
				{Name: "Bread", Description: "Freshly baked bread", Value: 2, Rarity: "Common"},
			},
		},
	}
}

func gameLoop(game *GameState) {
	for {
		fmt.Printf("\nDay %d - Current Time: %s\n", game.DayCount, game.CurrentTime.Format("15:04"))
		fmt.Print("Enter command: ")
		var input string
		fmt.Scanln(&input)
		processCommand(input, game)
		if game.Player.Health <= 0 {
			fmt.Println("Game Over! You have been defeated.")
			break
		}
		game.CurrentTime = game.CurrentTime.Add(30 * time.Minute)
		if game.CurrentTime.Hour() >= 24 {
			game.CurrentTime = time.Date(game.CurrentTime.Year(), game.CurrentTime.Month(), game.CurrentTime.Day()+1, 0, 0, 0, 0, time.UTC)
			game.DayCount++
		}
	}
}

func processCommand(input string, game *GameState) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}
	command := strings.ToLower(parts[0])
	switch command {
	case "help":
		printHelp()
	case "status":
		printStatus(game)
	case "inventory":
		printInventory(game)
	case "explore":
		exploreLocation(game)
	case "fight":
		if len(parts) < 2 {
			fmt.Println("Usage: fight <enemy_name>")
			return
		}
		fightEnemy(parts[1], game)
	case "use":
		if len(parts) < 2 {
			fmt.Println("Usage: use <item_name>")
			return
		}
		useItem(parts[1], game)
	case "shop":
		shopMenu(game)
	case "save":
		saveGame(game)
	case "load":
		loadGame(game)
	case "quit":
		fmt.Println("Thanks for playing!")
		os.Exit(0)
	default:
		fmt.Println("Unknown command. Type 'help' for a list of commands.")
	}
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  help      - Show this help message")
	fmt.Println("  status    - Show player status")
	fmt.Println("  inventory - Show player inventory")
	fmt.Println("  explore   - Explore current location")
	fmt.Println("  fight     - Fight an enemy (e.g., fight Goblin)")
	fmt.Println("  use       - Use an item from inventory (e.g., use Health Potion)")
	fmt.Println("  shop      - Access the shop")
	fmt.Println("  save      - Save the game")
	fmt.Println("  load      - Load a saved game")
	fmt.Println("  quit      - Exit the game")
}

func printStatus(game *GameState) {
	player := &game.Player
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Health: %d\n", player.Health)
	fmt.Printf("Strength: %d\n", player.Strength)
	fmt.Printf("Agility: %d\n", player.Agility)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Exp: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
}

func printInventory(game *GameState) {
	player := &game.Player
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("Inventory:")
	for i, item := range player.Inventory {
		fmt.Printf("  %d. %s - %s (Value: %d, Rarity: %s)\n", i+1, item.Name, item.Description, item.Value, item.Rarity)
	}
}

func exploreLocation(game *GameState) {
	currentLocation := game.Locations[rand.Intn(len(game.Locations))]
	fmt.Printf("You are exploring the %s. %s\n", currentLocation.Name, currentLocation.Description)
	if len(currentLocation.Enemies) > 0 {
		fmt.Println("Enemies here:")
		for _, enemy := range currentLocation.Enemies {
			fmt.Printf("  - %s\n", enemy.Name)
		}
	}
	if len(currentLocation.Items) > 0 {
		fmt.Println("Items found:")
		for _, item := range currentLocation.Items {
			fmt.Printf("  - %s\n", item.Name)
			game.Player.Inventory = append(game.Player.Inventory, item)
		}
	}
}

func fightEnemy(enemyName string, game *GameState) {
	var enemy *Enemy
	for i := range game.Enemies {
		if strings.EqualFold(game.Enemies[i].Name, enemyName) {
			enemy = &game.Enemies[i]
			break
		}
	}
	if enemy == nil {
		fmt.Printf("Enemy '%s' not found.\n", enemyName)
		return
	}
	fmt.Printf("You are fighting a %s!\n", enemy.Name)
	for enemy.Health > 0 && game.Player.Health > 0 {
		playerAttack := rand.Intn(game.Player.Strength) + 1
		enemy.Health -= playerAttack
		fmt.Printf("You hit the %s for %d damage. %s's health: %d\n", enemy.Name, playerAttack, enemy.Name, enemy.Health)
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			game.Player.Exp += enemy.Reward.Exp
			game.Player.Gold += enemy.Reward.Gold
			if enemy.Reward.Item.Name != "" {
				game.Player.Inventory = append(game.Player.Inventory, enemy.Reward.Item)
				fmt.Printf("You found: %s\n", enemy.Reward.Item.Name)
			}
			checkLevelUp(game)
			return
		}
		enemyAttack := rand.Intn(enemy.Strength) + 1
		game.Player.Health -= enemyAttack
		fmt.Printf("The %s hits you for %d damage. Your health: %d\n", enemy.Name, enemyAttack, game.Player.Health)
	}
	if game.Player.Health <= 0 {
		fmt.Println("You have been defeated!")
	}
}

func useItem(itemName string, game *GameState) {
	for i, item := range game.Player.Inventory {
		if strings.EqualFold(item.Name, itemName) {
			switch item.Name {
			case "Health Potion":
				game.Player.Health += 50
				if game.Player.Health > 100 {
					game.Player.Health = 100
				}
				fmt.Println("You used a Health Potion and restored 50 health.")
			case "Herbs":
				game.Player.Health += 10
				if game.Player.Health > 100 {
					game.Player.Health = 100
				}
				fmt.Println("You used Herbs and restored 10 health.")
			default:
				fmt.Printf("You cannot use %s.\n", item.Name)
				return
			}
			game.Player.Inventory = append(game.Player.Inventory[:i], game.Player.Inventory[i+1:]...)
			return
		}
	}
	fmt.Printf("Item '%s' not found in inventory.\n", itemName)
}

func shopMenu(game *GameState) {
	fmt.Println("Welcome to the shop!")
	fmt.Println("Available items:")
	items := []Item{
		{Name: "Health Potion", Description: "Restores 50 health", Value: 25, Rarity: "Common"},
		{Name: "Strength Potion", Description: "Increases strength by 5", Value: 50, Rarity: "Uncommon"},
		{Name: "Agility Potion", Description: "Increases agility by 5", Value: 50, Rarity: "Uncommon"},
	}
	for i, item := range items {
		fmt.Printf("  %d. %s - %s (Cost: %d gold)\n", i+1, item.Name, item.Description, item.Value)
	}
	fmt.Print("Enter item number to buy (or 'back' to leave): ")
	var input string
	fmt.Scanln(&input)
	if input == "back" {
		return
	}
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(items) {
		fmt.Println("Invalid selection.")
		return
	}
	item := items[index-1]
	if game.Player.Gold < item.Value {
		fmt.Println("Not enough gold!")
		return
	}
	game.Player.Gold -= item.Value
	game.Player.Inventory = append(game.Player.Inventory, item)
	fmt.Printf("You bought a %s for %d gold.\n", item.Name, item.Value)
}

func checkLevelUp(game *GameState) {
	player := &game.Player
	requiredExp := player.Level * 100
	if player.Exp >= requiredExp {
		player.Level++
		player.Exp -= requiredExp
		player.Health += 20
		player.Strength += 2
		player.Agility += 2
		fmt.Printf("Level up! You are now level %d. Health, Strength, and Agility increased.\n", player.Level)
	}
}

func saveGame(game *GameState) {
	data, err := json.Marshal(game)
	if err != nil {
		fmt.Println("Error saving game:", err)
		return
	}
	err = os.WriteFile("savegame.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing save file:", err)
		return
	}
	fmt.Println("Game saved successfully.")
}

func loadGame(game *GameState) {
	data, err := os.ReadFile("savegame.json")
	if err != nil {
		fmt.Println("Error loading game:", err)
		return
	}
	err = json.Unmarshal(data, game)
	if err != nil {
		fmt.Println("Error parsing save file:", err)
		return
	}
	fmt.Println("Game loaded successfully.")
}
