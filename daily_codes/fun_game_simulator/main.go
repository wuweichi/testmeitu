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

// Player represents a player in the game
type Player struct {
	Name     string
	Health   int
	Strength int
	Level    int
	Exp      int
}

// Enemy represents an enemy in the game
type Enemy struct {
	Name     string
	Health   int
	Strength int
}

// Item represents an item that can be used or collected
type Item struct {
	Name        string
	Description string
	Effect      func(*Player)
}

// GameState holds the current state of the game
type GameState struct {
	Player      *Player
	Enemies     []*Enemy
	Inventory   []*Item
	GameOver    bool
	CurrentRoom int
}

// Constants for game setup
const (
	StartingHealth = 100
	StartingStrength = 10
	ExpPerLevel    = 100
	MaxLevel       = 10
)

// Global random generator
var rng *rand.Rand

func init() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
}

// NewPlayer creates a new player with default values
func NewPlayer(name string) *Player {
	return &Player{
		Name:     name,
		Health:   StartingHealth,
		Strength: StartingStrength,
		Level:    1,
		Exp:      0,
	}
}

// NewEnemy creates a new enemy with random stats
func NewEnemy() *Enemy {
	enemyTypes := []string{"Goblin", "Orc", "Dragon", "Skeleton"}
	name := enemyTypes[rng.Intn(len(enemyTypes))]
	health := rng.Intn(50) + 30
	strength := rng.Intn(10) + 5
	return &Enemy{
		Name:     name,
		Health:   health,
		Strength: strength,
	}
}

// NewItem creates a new item with a random effect
func NewItem() *Item {
	itemNames := []string{"Health Potion", "Strength Elixir", "Magic Scroll"}
	name := itemNames[rng.Intn(len(itemNames))]
	var effect func(*Player)
	switch name {
	case "Health Potion":
		effect = func(p *Player) { p.Health += 20; fmt.Println("Health increased by 20!") }
	case "Strength Elixir":
		effect = func(p *Player) { p.Strength += 5; fmt.Println("Strength increased by 5!") }
	case "Magic Scroll":
		effect = func(p *Player) { p.Exp += 50; fmt.Println("Gained 50 experience!"); p.CheckLevelUp() }
	}
	return &Item{
		Name:        name,
		Description: fmt.Sprintf("A useful %s", name),
		Effect:      effect,
	}
}

// CheckLevelUp checks if the player has enough experience to level up
func (p *Player) CheckLevelUp() {
	for p.Exp >= p.Level*ExpPerLevel && p.Level < MaxLevel {
		p.Exp -= p.Level * ExpPerLevel
		p.Level++
		p.Health += 20
		p.Strength += 5
		fmt.Printf("Level up! You are now level %d. Health: %d, Strength: %d\n", p.Level, p.Health, p.Strength)
	}
}

// Attack simulates an attack between two entities
func Attack(attacker, defender interface{}) int {
	var damage int
	switch a := attacker.(type) {
	case *Player:
		damage = a.Strength + rng.Intn(10)
	case *Enemy:
		damage = a.Strength + rng.Intn(5)
	}
	switch d := defender.(type) {
	case *Player:
		d.Health -= damage
		if d.Health < 0 {
			d.Health = 0
		}
	case *Enemy:
		d.Health -= damage
		if d.Health < 0 {
			d.Health = 0
		}
	}
	return damage
}

// DisplayStatus shows the current status of the player
func (gs *GameState) DisplayStatus() {
	fmt.Printf("Player: %s, Health: %d, Strength: %d, Level: %d, Exp: %d/%d\n",
		gs.Player.Name, gs.Player.Health, gs.Player.Strength, gs.Player.Level, gs.Player.Exp, gs.Player.Level*ExpPerLevel)
	fmt.Printf("Inventory: %d items\n", len(gs.Inventory))
}

// HandleCombat handles a combat encounter with an enemy
func (gs *GameState) HandleCombat() {
	enemy := NewEnemy()
	fmt.Printf("A wild %s appears! Health: %d, Strength: %d\n", enemy.Name, enemy.Health, enemy.Strength)
	for gs.Player.Health > 0 && enemy.Health > 0 {
		fmt.Println("Choose action: (1) Attack, (2) Use Item, (3) Flee")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			damage := Attack(gs.Player, enemy)
			fmt.Printf("You attack the %s for %d damage. Enemy health: %d\n", enemy.Name, damage, enemy.Health)
			if enemy.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", enemy.Name)
				gs.Player.Exp += 30
				gs.Player.CheckLevelUp()
				if rng.Intn(100) < 30 { // 30% chance to drop an item
					item := NewItem()
					gs.Inventory = append(gs.Inventory, item)
					fmt.Printf("The enemy dropped a %s!\n", item.Name)
				}
				return
			}
			// Enemy attacks back
			damage = Attack(enemy, gs.Player)
			fmt.Printf("The %s attacks you for %d damage. Your health: %d\n", enemy.Name, damage, gs.Player.Health)
			if gs.Player.Health <= 0 {
				fmt.Println("You have been defeated! Game over.")
				gs.GameOver = true
				return
			}
		case "2":
			if len(gs.Inventory) == 0 {
				fmt.Println("No items in inventory!")
				continue
			}
			fmt.Println("Select an item to use:")
			for i, item := range gs.Inventory {
				fmt.Printf("%d: %s - %s\n", i+1, item.Name, item.Description)
			}
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			index, err := strconv.Atoi(input)
			if err != nil || index < 1 || index > len(gs.Inventory) {
				fmt.Println("Invalid selection!")
				continue
			}
			item := gs.Inventory[index-1]
			item.Effect(gs.Player)
			// Remove used item from inventory
			gs.Inventory = append(gs.Inventory[:index-1], gs.Inventory[index:]...)
		case "3":
			if rng.Intn(100) < 50 { // 50% chance to flee successfully
				fmt.Println("You successfully fled from combat!")
				return
			} else {
				fmt.Println("Failed to flee! The enemy attacks.")
				damage := Attack(enemy, gs.Player)
				fmt.Printf("The %s attacks you for %d damage. Your health: %d\n", enemy.Name, damage, gs.Player.Health)
				if gs.Player.Health <= 0 {
					fmt.Println("You have been defeated! Game over.")
					gs.GameOver = true
					return
				}
			}
		default:
			fmt.Println("Invalid action!")
		}
	}
}

// ExploreRoom simulates exploring a room in the game
func (gs *GameState) ExploreRoom() {
	fmt.Printf("You are in room %d.\n", gs.CurrentRoom)
	events := []string{"combat", "item", "nothing"}
	event := events[rng.Intn(len(events))]
	switch event {
	case "combat":
		gs.HandleCombat()
	case "item":
		item := NewItem()
		gs.Inventory = append(gs.Inventory, item)
		fmt.Printf("You found a %s!\n", item.Name)
	case "nothing":
		fmt.Println("You explore the room but find nothing of interest.")
	}
	gs.CurrentRoom++
}

// Main game loop
func (gs *GameState) RunGame() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the Fun Game Simulator!")
	fmt.Print("Enter your player name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	gs.Player = NewPlayer(name)
	gs.Enemies = make([]*Enemy, 0)
	gs.Inventory = make([]*Item, 0)
	gs.GameOver = false
	gs.CurrentRoom = 1

	for !gs.GameOver {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Explore next room")
		fmt.Println("2. Check status")
		fmt.Println("3. Use item")
		fmt.Println("4. Quit game")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			gs.ExploreRoom()
		case "2":
			gs.DisplayStatus()
		case "3":
			if len(gs.Inventory) == 0 {
				fmt.Println("No items in inventory!")
				continue
			}
			fmt.Println("Select an item to use:")
			for i, item := range gs.Inventory {
				fmt.Printf("%d: %s - %s\n", i+1, item.Name, item.Description)
			}
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			index, err := strconv.Atoi(input)
			if err != nil || index < 1 || index > len(gs.Inventory) {
				fmt.Println("Invalid selection!")
				continue
			}
			item := gs.Inventory[index-1]
			item.Effect(gs.Player)
			// Remove used item from inventory
			gs.Inventory = append(gs.Inventory[:index-1], gs.Inventory[index:]...)
		case "4":
			fmt.Println("Thanks for playing! Goodbye.")
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}

func main() {
	game := &GameState{}
	game.RunGame()
}
