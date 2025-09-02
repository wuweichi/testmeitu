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

// Player struct to hold player information
type Player struct {
	Name     string
	Health   int
	Strength int
	Level    int
	Exp      int
}

// Enemy struct to hold enemy information
type Enemy struct {
	Name     string
	Health   int
	Strength int
	Level    int
}

// Item struct to hold item information
type Item struct {
	Name        string
	Description string
	Effect      func(*Player)
}

// GameState struct to manage game state
type GameState struct {
	Player      *Player
	Enemies     []*Enemy
	Inventory   []*Item
	GameRunning bool
}

// InitializePlayer creates a new player with default values
func InitializePlayer(name string) *Player {
	return &Player{
		Name:     name,
		Health:   100,
		Strength: 10,
		Level:    1,
		Exp:      0,
	}
}

// InitializeEnemies creates a list of enemies
func InitializeEnemies() []*Enemy {
	return []*Enemy{
		{Name: "Goblin", Health: 30, Strength: 5, Level: 1},
		{Name: "Orc", Health: 50, Strength: 8, Level: 2},
		{Name: "Dragon", Health: 100, Strength: 15, Level: 3},
	}
}

// InitializeItems creates a list of items
func InitializeItems() []*Item {
	return []*Item{
		{
			Name:        "Health Potion",
			Description: "Restores 20 health points.",
			Effect: func(p *Player) {
				p.Health += 20
				fmt.Println("You used a Health Potion and restored 20 health!")
			},
		},
		{
			Name:        "Strength Elixir",
			Description: "Increases strength by 5 points.",
			Effect: func(p *Player) {
				p.Strength += 5
				fmt.Println("You used a Strength Elixir and gained 5 strength!")
			},
		},
	}
}

// Attack simulates an attack between two entities
func Attack(attacker, defender interface{}) {
	switch a := attacker.(type) {
	case *Player:
		switch d := defender.(type) {
		case *Enemy:
			damage := a.Strength + rand.Intn(5) // Random damage variation
			d.Health -= damage
			fmt.Printf("%s attacks %s for %d damage!\n", a.Name, d.Name, damage)
		}
	case *Enemy:
		switch d := defender.(type) {
		case *Player:
			damage := a.Strength + rand.Intn(5)
			d.Health -= damage
			fmt.Printf("%s attacks %s for %d damage!\n", a.Name, d.Name, damage)
		}
	}
}

// CheckLevelUp checks if the player should level up
func CheckLevelUp(p *Player) {
	if p.Exp >= p.Level*100 {
		p.Level++
		p.Health += 20
		p.Strength += 5
		fmt.Printf("Level up! You are now level %d. Health and strength increased.\n", p.Level)
		p.Exp = 0
	}
}

// DisplayStatus shows the player's current status
func DisplayStatus(p *Player) {
	fmt.Printf("Name: %s, Health: %d, Strength: %d, Level: %d, Exp: %d\n", p.Name, p.Health, p.Strength, p.Level, p.Exp)
}

// DisplayInventory shows the player's inventory
func DisplayInventory(inventory []*Item) {
	if len(inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("Inventory:")
	for i, item := range inventory {
		fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.Description)
	}
}

// UseItem allows the player to use an item from inventory
func UseItem(inventory []*Item, p *Player, index int) []*Item {
	if index < 0 || index >= len(inventory) {
		fmt.Println("Invalid item selection.")
		return inventory
	}
	item := inventory[index]
	item.Effect(p)
	// Remove the used item from inventory
	return append(inventory[:index], inventory[index+1:]...)
}

// Battle simulates a battle between player and an enemy
func Battle(p *Player, e *Enemy) {
	fmt.Printf("A wild %s appears!\n", e.Name)
	for p.Health > 0 && e.Health > 0 {
		fmt.Println("\nChoose an action: 1. Attack, 2. Use Item, 3. Run")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			Attack(p, e)
			if e.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", e.Name)
				p.Exp += e.Level * 20
				CheckLevelUp(p)
				return
			}
			Attack(e, p)
			if p.Health <= 0 {
				fmt.Println("You have been defeated! Game over.")
				return
			}
		case "2":
			// For simplicity, assume no inventory in battle; can extend
			fmt.Println("No items available in battle in this version.")
		case "3":
			fmt.Println("You ran away safely.")
			return
		default:
			fmt.Println("Invalid choice. Please choose 1, 2, or 3.")
		}
	}
}

// Explore allows the player to explore the game world
func Explore(gs *GameState) {
	fmt.Println("You are exploring...")
	// Random event: encounter enemy or find item
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Println("You found nothing.")
	case 1:
		enemyIndex := rand.Intn(len(gs.Enemies))
		Battle(gs.Player, gs.Enemies[enemyIndex])
	case 2:
		if len(gs.Inventory) < 5 { // Limit inventory size
			itemIndex := rand.Intn(len(InitializeItems()))
			newItem := InitializeItems()[itemIndex]
			gs.Inventory = append(gs.Inventory, newItem)
			fmt.Printf("You found a %s!\n", newItem.Name)
		} else {
			fmt.Println("Your inventory is full. Cannot pick up more items.")
		}
	}
}

// Main game loop
func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Fun CLI Game!")
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	player := InitializePlayer(name)
	enemies := InitializeEnemies()
	inventory := []*Item{}
	gameState := &GameState{
		Player:      player,
		Enemies:     enemies,
		Inventory:   inventory,
		GameRunning: true,
	}

	for gameState.GameRunning {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Explore")
		fmt.Println("2. Check Status")
		fmt.Println("3. Check Inventory")
		fmt.Println("4. Use Item")
		fmt.Println("5. Quit")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			Explore(gameState)
			if gameState.Player.Health <= 0 {
				gameState.GameRunning = false
			}
		case "2":
			DisplayStatus(gameState.Player)
		case "3":
			DisplayInventory(gameState.Inventory)
		case "4":
			if len(gameState.Inventory) == 0 {
				fmt.Println("No items to use.")
				break
			}
			DisplayInventory(gameState.Inventory)
			fmt.Print("Select an item to use (number): ")
			itemInput, _ := reader.ReadString('\n')
			itemInput = strings.TrimSpace(itemInput)
			index, err := strconv.Atoi(itemInput)
			if err != nil || index < 1 || index > len(gameState.Inventory) {
				fmt.Println("Invalid selection.")
				break
			}
			gameState.Inventory = UseItem(gameState.Inventory, gameState.Player, index-1)
		case "5":
			fmt.Println("Thanks for playing! Goodbye.")
			gameState.GameRunning = false
		default:
			fmt.Println("Invalid option. Please choose 1-5.")
		}
	}
}
