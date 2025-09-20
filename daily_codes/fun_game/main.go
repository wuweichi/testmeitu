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
	Score    int
	Inventory []string
}

// Enemy struct to hold enemy information
type Enemy struct {
	Name   string
	Health int
	Damage int
}

// Item struct to hold item information
type Item struct {
	Name        string
	Description string
	Effect      func(*Player)
}

// GameState struct to manage the game state
type GameState struct {
	Player      Player
	Enemies     []Enemy
	Items       []Item
	CurrentRoom string
	GameOver    bool
}

// Initialize the game state
func (gs *GameState) Init() {
	gs.Player = Player{
		Name:     "Hero",
		Health:   100,
		Score:    0,
		Inventory: []string{},
	}
	gs.Enemies = []Enemy{
		{Name: "Goblin", Health: 30, Damage: 10},
		{Name: "Orc", Health: 50, Damage: 15},
		{Name: "Dragon", Health: 100, Damage: 25},
	}
	gs.Items = []Item{
		{
			Name:        "Potion",
			Description: "Restores 20 health.",
			Effect: func(p *Player) {
				p.Health += 20
				fmt.Println("You used a Potion and restored 20 health!")
			},
		},
		{
			Name:        "Sword",
			Description: "Increases damage in combat.",
			Effect: func(p *Player) {
				// Placeholder effect, could be implemented in combat logic
				fmt.Println("You equipped the Sword. It feels powerful!")
			},
		},
	}
	gs.CurrentRoom = "start"
	gs.GameOver = false
}

// Display the current status of the player
func (gs *GameState) DisplayStatus() {
	fmt.Printf("Player: %s\n", gs.Player.Name)
	fmt.Printf("Health: %d\n", gs.Player.Health)
	fmt.Printf("Score: %d\n", gs.Player.Score)
	fmt.Printf("Inventory: %v\n", gs.Player.Inventory)
	fmt.Printf("Current Room: %s\n", gs.CurrentRoom)
}

// Handle player movement between rooms
func (gs *GameState) Move(direction string) {
	rooms := map[string][]string{
		"start":    {"north", "east"},
		"forest":   {"south", "west", "east"},
		"cave":     {"west", "north"},
		"treasure": {"south"},
	}
	if validDirections, exists := rooms[gs.CurrentRoom]; exists {
		valid := false
		for _, d := range validDirections {
			if d == direction {
				valid = true
				break
			}
		}
		if valid {
			switch direction {
			case "north":
				gs.CurrentRoom = "forest"
			case "south":
				gs.CurrentRoom = "start"
			case "east":
				gs.CurrentRoom = "cave"
			case "west":
				gs.CurrentRoom = "treasure"
			}
			fmt.Printf("You moved %s to the %s.\n", direction, gs.CurrentRoom)
		} else {
			fmt.Println("You can't go that way!")
		}
	} else {
		fmt.Println("Invalid room!")
	}
}

// Simulate combat with a random enemy
func (gs *GameState) Combat() {
	rand.Seed(time.Now().UnixNano())
	enemyIndex := rand.Intn(len(gs.Enemies))
	enemy := gs.Enemies[enemyIndex]
	fmt.Printf("A wild %s appears!\n", enemy.Name)
	for gs.Player.Health > 0 && enemy.Health > 0 {
		fmt.Printf("Your health: %d, %s's health: %d\n", gs.Player.Health, enemy.Name, enemy.Health)
		fmt.Print("Choose action: (1) Attack, (2) Run: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			damage := rand.Intn(20) + 10 // Player damage between 10-30
			enemy.Health -= damage
			fmt.Printf("You attacked the %s for %d damage!\n", enemy.Name, damage)
			if enemy.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", enemy.Name)
				gs.Player.Score += 10
				break
			}
			// Enemy attacks back
			gs.Player.Health -= enemy.Damage
			fmt.Printf("The %s attacked you for %d damage!\n", enemy.Name, enemy.Damage)
		case "2":
			fmt.Println("You ran away safely!")
			return
		default:
			fmt.Println("Invalid action!")
		}
	}
	if gs.Player.Health <= 0 {
		fmt.Println("You have been defeated! Game over.")
		gs.GameOver = true
	}
}

// Allow player to use an item from inventory
func (gs *GameState) UseItem(itemName string) {
	for _, item := range gs.Items {
		if item.Name == itemName {
			for i, invItem := range gs.Player.Inventory {
				if invItem == itemName {
					item.Effect(&gs.Player)
					// Remove item from inventory after use
					gs.Player.Inventory = append(gs.Player.Inventory[:i], gs.Player.Inventory[i+1:]...)
					return
				}
			}
			fmt.Println("Item not found in inventory!")
			return
		}
	}
	fmt.Println("Item does not exist!")
}

// Add an item to player's inventory
func (gs *GameState) AddItem(itemName string) {
	for _, item := range gs.Items {
		if item.Name == itemName {
			gs.Player.Inventory = append(gs.Player.Inventory, itemName)
			fmt.Printf("You obtained a %s!\n", itemName)
			return
		}
	}
	fmt.Println("Item not found!")
}

// Main game loop
func (gs *GameState) GameLoop() {
	reader := bufio.NewReader(os.Stdin)
	for !gs.GameOver {
		fmt.Println("\n--- Game Menu ---")
		fmt.Println("1. Display Status")
		fmt.Println("2. Move (north, south, east, west)")
		fmt.Println("3. Combat")
		fmt.Println("4. Use Item")
		fmt.Println("5. Add Item (for testing)")
		fmt.Println("6. Quit")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			gs.DisplayStatus()
		case "2":
			fmt.Print("Enter direction (north, south, east, west): ")
			direction, _ := reader.ReadString('\n')
			direction = strings.TrimSpace(direction)
			gs.Move(direction)
		case "3":
			gs.Combat()
		case "4":
			fmt.Print("Enter item name to use: ")
			itemName, _ := reader.ReadString('\n')
			itemName = strings.TrimSpace(itemName)
			gs.UseItem(itemName)
		case "5":
			fmt.Print("Enter item name to add: ")
			itemName, _ := reader.ReadString('\n')
			itemName = strings.TrimSpace(itemName)
			gs.AddItem(itemName)
		case "6":
			fmt.Println("Thanks for playing!")
			gs.GameOver = true
		default:
			fmt.Println("Invalid option!")
		}
	}
}

func main() {
	fmt.Println("Welcome to the Fun Game!")
	gameState := GameState{}
	gameState.Init()
	gameState.GameLoop()
}
