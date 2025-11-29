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
	player = Character{
		Name:      "Hero",
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
		Gold:      50,
		Inventory: []string{"Health Potion"},
	}

	locations = make(map[string]Location)

	locations["forest"] = Location{
		Name:        "Enchanted Forest",
		Description: "A mystical forest filled with ancient trees and magical creatures.",
		Enemies: []Enemy{
			{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Exp: 10, Gold: 5},
			{Name: "Wolf", Health: 25, Attack: 8, Defense: 1, Exp: 8, Gold: 3},
		},
		Items:       []string{"Health Potion", "Magic Herb"},
		Connections: []string{"village", "cave"},
	}

	locations["village"] = Location{
		Name:        "Peaceful Village",
		Description: "A small village with friendly inhabitants and basic amenities.",
		Enemies:     []Enemy{},
		Items:       []string{"Bread", "Water"},
		Connections: []string{"forest", "castle"},
	}

	locations["cave"] = Location{
		Name:        "Dark Cave",
		Description: "A dark and dangerous cave filled with treasures and monsters.",
		Enemies: []Enemy{
			{Name: "Bat", Health: 15, Attack: 3, Defense: 0, Exp: 5, Gold: 1},
			{Name: "Spider", Health: 20, Attack: 6, Defense: 1, Exp: 7, Gold: 2},
			{Name: "Troll", Health: 50, Attack: 12, Defense: 5, Exp: 20, Gold: 15},
		},
		Items:       []string{"Gold Coin", "Ancient Artifact"},
		Connections: []string{"forest"},
	}

	locations["castle"] = Location{
		Name:        "Royal Castle",
		Description: "A magnificent castle with royal guards and noble inhabitants.",
		Enemies:     []Enemy{},
		Items:       []string{"Royal Sword", "Healing Scroll"},
		Connections: []string{"village"},
	}

	currentLocation = "forest"
}

func showStatus() {
	fmt.Printf("\n=== Status ===\n")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Exp: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Location: %s\n", locations[currentLocation].Name)
	fmt.Printf("Inventory: %v\n", player.Inventory)
}

func moveTo(newLocation string) {
	current := locations[currentLocation]
	found := false
	for _, conn := range current.Connections {
		if conn == newLocation {
			found = true
			break
		}
	}
	if !found {
		fmt.Println("You cannot go there from here!")
		return
	}
	currentLocation = newLocation
	fmt.Printf("You have moved to %s.\n", locations[currentLocation].Name)
	fmt.Println(locations[currentLocation].Description)
}

func exploreLocation() {
	loc := locations[currentLocation]
	if len(loc.Enemies) > 0 {
		fmt.Println("You encounter some enemies!")
		for i, enemy := range loc.Enemies {
			fmt.Printf("%d. %s (Health: %d)\n", i+1, enemy.Name, enemy.Health)
		}
	} else {
		fmt.Println("This area seems peaceful.")
	}

	if len(loc.Items) > 0 {
		fmt.Println("You find some items:")
		for i, item := range loc.Items {
			fmt.Printf("%d. %s\n", i+1, item)
		}
	}
}

func fightEnemy(enemyIndex int) {
	loc := locations[currentLocation]
	if enemyIndex < 0 || enemyIndex >= len(loc.Enemies) {
		fmt.Println("Invalid enemy selection!")
		return
	}

	enemy := loc.Enemies[enemyIndex]
	fmt.Printf("You engage in combat with %s!\n", enemy.Name)

	for player.Health > 0 && enemy.Health > 0 {
		playerDamage := player.Attack - enemy.Defense
		if playerDamage < 1 {
			playerDamage = 1
		}
		enemy.Health -= playerDamage
		fmt.Printf("You attack %s for %d damage. %s's health: %d\n", enemy.Name, playerDamage, enemy.Name, enemy.Health)

		if enemy.Health <= 0 {
			break
		}

		enemyDamage := enemy.Attack - player.Defense
		if enemyDamage < 1 {
			enemyDamage = 1
		}
		player.Health -= enemyDamage
		fmt.Printf("%s attacks you for %d damage. Your health: %d\n", enemy.Name, enemyDamage, player.Health)

		if player.Health <= 0 {
			break
		}
	}

	if player.Health <= 0 {
		fmt.Println("You have been defeated! Game Over.")
		os.Exit(0)
	} else {
		fmt.Printf("You defeated %s!\n", enemy.Name)
		player.Exp += enemy.Exp
		player.Gold += enemy.Gold
		fmt.Printf("Gained %d experience and %d gold.\n", enemy.Exp, enemy.Gold)

		loc.Enemies = append(loc.Enemies[:enemyIndex], loc.Enemies[enemyIndex+1:]...)
		locations[currentLocation] = loc

		checkLevelUp()
	}
}

func checkLevelUp() {
	requiredExp := player.Level * 100
	if player.Exp >= requiredExp {
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		player.Exp -= requiredExp
		fmt.Printf("Congratulations! You reached level %d!\n", player.Level)
		fmt.Printf("Your health, attack, and defense have increased!\n")
	}
}

func useItem(itemIndex int) {
	if itemIndex < 0 || itemIndex >= len(player.Inventory) {
		fmt.Println("Invalid item selection!")
		return
	}

	item := player.Inventory[itemIndex]
	switch item {
	case "Health Potion":
		healAmount := 30
		player.Health += healAmount
		if player.Health > player.MaxHealth {
			player.Health = player.MaxHealth
		}
		fmt.Printf("You used a Health Potion and recovered %d health. Current health: %d\n", healAmount, player.Health)
		player.Inventory = append(player.Inventory[:itemIndex], player.Inventory[itemIndex+1:]...)
	case "Magic Herb":
		player.Attack += 2
		fmt.Printf("You used a Magic Herb. Your attack increased by 2. Current attack: %d\n", player.Attack)
		player.Inventory = append(player.Inventory[:itemIndex], player.Inventory[itemIndex+1:]...)
	default:
		fmt.Printf("You cannot use %s right now.\n", item)
	}
}

func pickUpItem(itemIndex int) {
	loc := locations[currentLocation]
	if itemIndex < 0 || itemIndex >= len(loc.Items) {
		fmt.Println("Invalid item selection!")
		return
	}

	item := loc.Items[itemIndex]
	player.Inventory = append(player.Inventory, item)
	loc.Items = append(loc.Items[:itemIndex], loc.Items[itemIndex+1:]...)
	locations[currentLocation] = loc
	fmt.Printf("You picked up %s.\n", item)
}

func showHelp() {
	fmt.Println("\n=== Available Commands ===")
	fmt.Println("status - Show player status")
	fmt.Println("explore - Explore current location")
	fmt.Println("move <location> - Move to another location")
	fmt.Println("fight <enemy_number> - Fight an enemy")
	fmt.Println("use <item_number> - Use an item from inventory")
	fmt.Println("pickup <item_number> - Pick up an item from location")
	fmt.Println("help - Show this help message")
	fmt.Println("quit - Exit the game")
}

func main() {
	fmt.Println("Welcome to the Interactive Adventure Game!")
	fmt.Println("Type 'help' for a list of commands.")

	initGame()
	fmt.Println(locations[currentLocation].Description)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("\n> ")
		if !scanner.Scan() {
			break
		}

		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}

		parts := strings.Fields(input)
		command := strings.ToLower(parts[0])

		switch command {
		case "quit":
			fmt.Println("Thanks for playing!")
			return
		case "help":
			showHelp()
		case "status":
			showStatus()
		case "explore":
			exploreLocation()
		case "move":
			if len(parts) < 2 {
				fmt.Println("Please specify a location to move to.")
			} else {
				moveTo(parts[1])
			}
		case "fight":
			if len(parts) < 2 {
				fmt.Println("Please specify an enemy number to fight.")
			} else {
				if index, err := strconv.Atoi(parts[1]); err == nil {
					fightEnemy(index - 1)
				} else {
					fmt.Println("Invalid enemy number!")
				}
			}
		case "use":
			if len(parts) < 2 {
				fmt.Println("Please specify an item number to use.")
			} else {
				if index, err := strconv.Atoi(parts[1]); err == nil {
					useItem(index - 1)
				} else {
					fmt.Println("Invalid item number!")
				}
			}
		case "pickup":
			if len(parts) < 2 {
				fmt.Println("Please specify an item number to pick up.")
			} else {
				if index, err := strconv.Atoi(parts[1]); err == nil {
					pickUpItem(index - 1)
				} else {
					fmt.Println("Invalid item number!")
				}
			}
		default:
			fmt.Println("Unknown command. Type 'help' for available commands.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
