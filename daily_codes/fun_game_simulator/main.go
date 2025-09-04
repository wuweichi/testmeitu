package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"os"
	"bufio"
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

// Game struct to manage the game state
type Game struct {
	Player      Player
	Enemies     []Enemy
	CurrentRoom int
	Rooms       []Room
}

// Room struct to define a room in the game
type Room struct {
	Description string
	Items       []string
	Enemies     []Enemy
	Exits       map[string]int // direction to room index
}

// Function to initialize the game
func (g *Game) InitializeGame() {
	g.Player = Player{
		Name:     "Hero",
		Health:   100,
		Score:    0,
		Inventory: []string{"sword", "potion"},
	}
	
	// Create some enemies
	enemy1 := Enemy{Name: "Goblin", Health: 30, Damage: 10}
	enemy2 := Enemy{Name: "Dragon", Health: 100, Damage: 25}
	enemy3 := Enemy{Name: "Skeleton", Health: 20, Damage: 5}
	
	// Define rooms
	room1 := Room{
		Description: "You are in a dark forest. There is a path to the north.",
		Items:       []string{"key"},
		Enemies:     []Enemy{enemy1},
		Exits:       map[string]int{"north": 1},
	}
	room2 := Room{
		Description: "You are in a clearing. A castle is visible to the east.",
		Items:       []string{"shield"},
		Enemies:     []Enemy{enemy2},
		Exits:       map[string]int{"east": 2, "south": 0},
	}
	room3 := Room{
		Description: "You are inside the castle throne room. The king's treasure is here!",
		Items:       []string{"treasure"},
		Enemies:     []Enemy{enemy3},
		Exits:       map[string]int{"west": 1},
	}
	
	g.Rooms = []Room{room1, room2, room3}
	g.CurrentRoom = 0
	g.Enemies = []Enemy{enemy1, enemy2, enemy3}
}

// Function to handle player movement
func (g *Game) Move(direction string) {
	room := g.Rooms[g.CurrentRoom]
	if nextRoom, ok := room.Exits[direction]; ok {
		g.CurrentRoom = nextRoom
		fmt.Printf("You move %s.\n", direction)
		g.DescribeRoom()
	} else {
		fmt.Println("You can't go that way.")
	}
}

// Function to describe the current room
func (g *Game) DescribeRoom() {
	room := g.Rooms[g.CurrentRoom]
	fmt.Println(room.Description)
	if len(room.Items) > 0 {
		fmt.Printf("You see: %v\n", room.Items)
	}
	if len(room.Enemies) > 0 {
		fmt.Printf("Enemies here: %v\n", room.Enemies)
	}
	fmt.Printf("Exits: %v\n", getKeys(room.Exits))
}

// Helper function to get keys from a map
func getKeys(m map[string]int) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Function to handle combat
func (g *Game) Combat(enemyIndex int) {
	enemy := g.Rooms[g.CurrentRoom].Enemies[enemyIndex]
	fmt.Printf("You are fighting a %s!\n", enemy.Name)
	
	for g.Player.Health > 0 && enemy.Health > 0 {
		// Player's turn
		fmt.Println("Your turn. Choose action: attack (a) or use item (i)")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1] // Remove newline
		
		switch input {
		case "a":
			damage := rand.Intn(20) + 10 // Random damage between 10 and 30
			enemy.Health -= damage
			fmt.Printf("You attack the %s for %d damage. %s health: %d\n", enemy.Name, damage, enemy.Name, enemy.Health)
		case "i":
			if len(g.Player.Inventory) > 0 {
				fmt.Printf("Your inventory: %v\n", g.Player.Inventory)
				fmt.Println("Choose item to use (enter index):")
				itemInput, _ := reader.ReadString('\n')
				itemInput = itemInput[:len(itemInput)-1]
				index, err := strconv.Atoi(itemInput)
				if err == nil && index >= 0 && index < len(g.Player.Inventory) {
					item := g.Player.Inventory[index]
					if item == "potion" {
						g.Player.Health += 30
						fmt.Println("You used a potion and gained 30 health.")
						g.Player.Inventory = append(g.Player.Inventory[:index], g.Player.Inventory[index+1:]...)
					} else {
						fmt.Println("That item cannot be used in combat.")
					}
				} else {
					fmt.Println("Invalid item index.")
				}
			} else {
				fmt.Println("No items in inventory.")
			}
		default:
			fmt.Println("Invalid action.")
		}
		
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			g.Player.Score += 10
			g.Rooms[g.CurrentRoom].Enemies = append(g.Rooms[g.CurrentRoom].Enemies[:enemyIndex], g.Rooms[g.CurrentRoom].Enemies[enemyIndex+1:]...)
			break
		}
		
		// Enemy's turn
		if enemy.Health > 0 {
			enemyDamage := rand.Intn(enemy.Damage) + 1
			g.Player.Health -= enemyDamage
			fmt.Printf("%s attacks you for %d damage. Your health: %d\n", enemy.Name, enemyDamage, g.Player.Health)
		}
		
		if g.Player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			os.Exit(0)
		}
	}
}

// Function to pick up items
func (g *Game) PickUpItem(itemIndex int) {
	room := g.Rooms[g.CurrentRoom]
	if itemIndex >= 0 && itemIndex < len(room.Items) {
		item := room.Items[itemIndex]
		g.Player.Inventory = append(g.Player.Inventory, item)
		g.Rooms[g.CurrentRoom].Items = append(room.Items[:itemIndex], room.Items[itemIndex+1:]...)
		fmt.Printf("You picked up: %s\n", item)
	} else {
		fmt.Println("Invalid item index.")
	}
}

// Function to display player status
func (g *Game) ShowStatus() {
	fmt.Printf("Name: %s, Health: %d, Score: %d, Inventory: %v\n", g.Player.Name, g.Player.Health, g.Player.Score, g.Player.Inventory)
}

// Main function to run the game
func main() {
	rand.Seed(time.Now().UnixNano())
	
	game := Game{}
	game.InitializeGame()
	
	fmt.Println("Welcome to the Fun Game Simulator!")
	fmt.Println("You are an adventurer exploring a mysterious world.")
	game.DescribeRoom()
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println("\nChoose an action: move (m), combat (c), pickup (p), status (s), quit (q)")
		input, _ := reader.ReadString('\n')
		input = input[:len(input)-1]
		
		switch input {
		case "m":
			fmt.Println("Enter direction (e.g., north, south, east, west):")
			direction, _ := reader.ReadString('\n')
			direction = direction[:len(direction)-1]
			game.Move(direction)
		case "c":
			room := game.Rooms[game.CurrentRoom]
			if len(room.Enemies) > 0 {
				fmt.Printf("Choose enemy to fight (enter index 0 to %d):\n", len(room.Enemies)-1)
				enemyInput, _ := reader.ReadString('\n')
				enemyInput = enemyInput[:len(enemyInput)-1]
				index, err := strconv.Atoi(enemyInput)
				if err == nil && index >= 0 && index < len(room.Enemies) {
					game.Combat(index)
				} else {
					fmt.Println("Invalid enemy index.")
				}
			} else {
				fmt.Println("No enemies in this room.")
			}
		case "p":
			room := game.Rooms[game.CurrentRoom]
			if len(room.Items) > 0 {
				fmt.Printf("Choose item to pickup (enter index 0 to %d):\n", len(room.Items)-1)
				itemInput, _ := reader.ReadString('\n')
				itemInput = itemInput[:len(itemInput)-1]
				index, err := strconv.Atoi(itemInput)
				if err == nil && index >= 0 && index < len(room.Items) {
					game.PickUpItem(index)
				} else {
					fmt.Println("Invalid item index.")
				}
			} else {
				fmt.Println("No items in this room.")
			}
		case "s":
			game.ShowStatus()
		case "q":
			fmt.Println("Thanks for playing!")
			os.Exit(0)
		default:
			fmt.Println("Invalid action. Please try again.")
		}
	}
}
