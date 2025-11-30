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

type Character struct {
	Name      string
	Health    int
	MaxHealth int
	Attack    int
	Defense   int
	Level     int
	Exp       int
	Gold      int
	Inventory []string
}

type Enemy struct {
	Name    string
	Health  int
	Attack  int
	Defense int
	Exp     int
	Gold    int
}

type Location struct {
	Name        string
	Description string
	Enemies     []Enemy
	Items       []string
	Connections []string
}

var player Character
var locations map[string]Location
var currentLocation string

func initGame() {
	rand.Seed(time.Now().UnixNano())
	
	player = Character{
		Name:      "Hero",
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
		Gold:      50,
		Inventory: []string{"Health Potion", "Health Potion"},
	}
	
	locations = make(map[string]Location)
	
	locations["forest"] = Location{
		Name:        "Mystic Forest",
		Description: "A dense forest with ancient trees and mysterious sounds.",
		Enemies: []Enemy{
			{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Exp: 10, Gold: 5},
			{Name: "Wolf", Health: 25, Attack: 7, Defense: 1, Exp: 8, Gold: 3},
		},
		Items:       []string{"Herbs", "Health Potion"},
		Connections: []string{"village", "cave"},
	}
	
	locations["village"] = Location{
		Name:        "Peaceful Village",
		Description: "A small village with friendly inhabitants and a market.",
		Enemies:     []Enemy{},
		Items:       []string{"Bread", "Water"},
		Connections: []string{"forest"},
	}
	
	locations["cave"] = Location{
		Name:        "Dark Cave",
		Description: "A dark and dangerous cave filled with treasures and monsters.",
		Enemies: []Enemy{
			{Name: "Troll", Health: 80, Attack: 15, Defense: 8, Exp: 30, Gold: 20},
			{Name: "Bat Swarm", Health: 20, Attack: 10, Defense: 0, Exp: 5, Gold: 2},
		},
		Items:       []string{"Gold Coin", "Magic Ring"},
		Connections: []string{"forest"},
	}
	
	currentLocation = "village"
}

func showStatus() {
	fmt.Printf("=== %s Status ===\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Exp: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Location: %s\n", locations[currentLocation].Name)
	fmt.Printf("Inventory: %v\n", player.Inventory)
	fmt.Println()
}

func showLocation() {
	loc := locations[currentLocation]
	fmt.Printf("=== %s ===\n", loc.Name)
	fmt.Println(loc.Description)
	
	if len(loc.Enemies) > 0 {
		fmt.Println("Enemies here:")
		for _, enemy := range loc.Enemies {
			fmt.Printf("- %s (Health: %d)\n", enemy.Name, enemy.Health)
		}
	}
	
	if len(loc.Items) > 0 {
		fmt.Println("Items here:")
		for _, item := range loc.Items {
			fmt.Printf("- %s\n", item)
		}
	}
	
	if len(loc.Connections) > 0 {
		fmt.Println("Connected locations:")
		for _, conn := range loc.Connections {
			fmt.Printf("- %s\n", locations[conn].Name)
		}
	}
	fmt.Println()
}

func moveTo(location string) {
	loc := locations[currentLocation]
	for _, conn := range loc.Connections {
		if conn == location {
			currentLocation = location
			fmt.Printf("You moved to %s.\n", locations[location].Name)
			return
		}
	}
	fmt.Println("You can't go there from here.")
}

func fightEnemy(enemyName string) {
	loc := locations[currentLocation]
	var enemy Enemy
	found := false
	
	for i, e := range loc.Enemies {
		if e.Name == enemyName {
			enemy = e
			loc.Enemies = append(loc.Enemies[:i], loc.Enemies[i+1:]...)
			locations[currentLocation] = loc
			found = true
			break
		}
	}
	
	if !found {
		fmt.Println("Enemy not found here.")
		return
	}
	
	fmt.Printf("You are fighting a %s!\n", enemy.Name)
	
	for player.Health > 0 && enemy.Health > 0 {
		// Player attack
		damage := player.Attack - enemy.Defense
		if damage < 1 {
			damage = 1
		}
		enemy.Health -= damage
		fmt.Printf("You hit the %s for %d damage.\n", enemy.Name, damage)
		
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			player.Exp += enemy.Exp
			player.Gold += enemy.Gold
			fmt.Printf("Gained %d exp and %d gold.\n", enemy.Exp, enemy.Gold)
			checkLevelUp()
			return
		}
		
		// Enemy attack
		damage = enemy.Attack - player.Defense
		if damage < 1 {
			damage = 1
		}
		player.Health -= damage
		fmt.Printf("The %s hits you for %d damage.\n", enemy.Name, damage)
		
		if player.Health <= 0 {
			fmt.Println("You have been defeated!")
			return
		}
	}
}

func checkLevelUp() {
	for player.Exp >= player.Level*100 {
		player.Exp -= player.Level * 100
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("Level up! You are now level %d.\n", player.Level)
	}
}

func useItem(itemName string) {
	for i, item := range player.Inventory {
		if item == itemName {
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			
			switch itemName {
			case "Health Potion":
				healAmount := 30
				player.Health += healAmount
				if player.Health > player.MaxHealth {
					player.Health = player.MaxHealth
				}
				fmt.Printf("Used Health Potion. Healed %d health.\n", healAmount)
			default:
				fmt.Printf("Used %s.\n", itemName)
			}
			return
		}
	}
	fmt.Println("Item not found in inventory.")
}

func pickUpItem(itemName string) {
	loc := locations[currentLocation]
	for i, item := range loc.Items {
		if item == itemName {
			loc.Items = append(loc.Items[:i], loc.Items[i+1:]...)
			locations[currentLocation] = loc
			player.Inventory = append(player.Inventory, item)
			fmt.Printf("You picked up %s.\n", item)
			return
		}
	}
	fmt.Println("Item not found here.")
}

func showHelp() {
	fmt.Println("=== Adventure Game Commands ===")
	fmt.Println("status - Show player status")
	fmt.Println("look - Show current location")
	fmt.Println("move [location] - Move to connected location")
	fmt.Println("fight [enemy] - Fight an enemy")
	fmt.Println("use [item] - Use an item from inventory")
	fmt.Println("pickup [item] - Pick up an item from location")
	fmt.Println("help - Show this help message")
	fmt.Println("quit - Quit the game")
	fmt.Println()
}

func main() {
	initGame()
	
	fmt.Println("Welcome to the Adventure Game!")
	fmt.Println("Type 'help' for a list of commands.")
	fmt.Println()
	
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		
		input := scanner.Text()
		parts := strings.Fields(input)
		
		if len(parts) == 0 {
			continue
		}
		
		command := strings.ToLower(parts[0])
		
		switch command {
		case "quit":
			fmt.Println("Thanks for playing!")
			return
		case "help":
			showHelp()
		case "status":
			showStatus()
		case "look":
			showLocation()
		case "move":
			if len(parts) < 2 {
				fmt.Println("Please specify a location to move to.")
			} else {
				moveTo(parts[1])
			}
		case "fight":
			if len(parts) < 2 {
				fmt.Println("Please specify an enemy to fight.")
			} else {
				fightEnemy(parts[1])
			}
		case "use":
			if len(parts) < 2 {
				fmt.Println("Please specify an item to use.")
			} else {
				useItem(parts[1])
			}
		case "pickup":
			if len(parts) < 2 {
				fmt.Println("Please specify an item to pick up.")
			} else {
				pickUpItem(parts[1])
			}
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
		
		if player.Health <= 0 {
			fmt.Println("Game Over!")
			return
		}
	}
}
