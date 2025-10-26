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
var currentLocation string
var locations map[string]Location
var gameRunning bool

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
		Inventory: []string{"Health Potion", "Health Potion"},
	}

	locations = make(map[string]Location)

	locations["forest"] = Location{
		Name:        "Dark Forest",
		Description: "A dense, dark forest with tall trees and mysterious sounds.",
		Enemies: []Enemy{
			{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Exp: 10, Gold: 5},
			{Name: "Wolf", Health: 25, Attack: 7, Defense: 1, Exp: 8, Gold: 3},
		},
		Items:       []string{"Health Potion", "Iron Sword"},
		Connections: []string{"village", "cave"},
	}

	locations["village"] = Location{
		Name:        "Peaceful Village",
		Description: "A small, peaceful village with friendly inhabitants.",
		Enemies:     []Enemy{},
		Items:       []string{"Bread", "Health Potion"},
		Connections: []string{"forest", "castle"},
	}

	locations["cave"] = Location{
		Name:        "Dark Cave",
		Description: "A dark and damp cave with glowing mushrooms.",
		Enemies: []Enemy{
			{Name: "Bat", Health: 15, Attack: 3, Defense: 0, Exp: 5, Gold: 2},
			{Name: "Spider", Health: 20, Attack: 4, Defense: 1, Exp: 7, Gold: 4},
		},
		Items:       []string{"Gold Coin", "Magic Ring"},
		Connections: []string{"forest"},
	}

	locations["castle"] = Location{
		Name:        "Ancient Castle",
		Description: "An ancient castle with towering walls and a mysterious aura.",
		Enemies: []Enemy{
			{Name: "Skeleton", Health: 40, Attack: 8, Defense: 3, Exp: 15, Gold: 10},
			{Name: "Dragon", Health: 100, Attack: 20, Defense: 10, Exp: 50, Gold: 100},
		},
		Items:       []string{"Dragon Scale", "Ancient Artifact"},
		Connections: []string{"village"},
	}

	currentLocation = "village"
	gameRunning = true
}

func showStatus() {
	fmt.Printf("=== %s Status ===\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Exp: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
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
			fmt.Printf("  - %s (Health: %d)\n", enemy.Name, enemy.Health)
		}
	}

	if len(loc.Items) > 0 {
		fmt.Println("Items here:")
		for _, item := range loc.Items {
			fmt.Printf("  - %s\n", item)
		}
	}

	fmt.Println("Connected locations:")
	for _, conn := range loc.Connections {
		fmt.Printf("  - %s\n", conn)
	}
	fmt.Println()
}

func moveToLocation(target string) {
	loc := locations[currentLocation]
	for _, conn := range loc.Connections {
		if conn == target {
			currentLocation = target
			fmt.Printf("You moved to %s.\n", locations[target].Name)
			return
		}
	}
	fmt.Println("You can't move to that location from here.")
}

func fightEnemy(enemyName string) {
	loc := locations[currentLocation]
	var enemy Enemy
	found := false

	for _, e := range loc.Enemies {
		if e.Name == enemyName {
			enemy = e
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
		fmt.Printf("\nYour Health: %d, %s's Health: %d\n", player.Health, enemy.Name, enemy.Health)
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Run Away")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			damage := player.Attack - enemy.Defense
			if damage < 1 {
				damage = 1
			}
			enemy.Health -= damage
			fmt.Printf("You hit the %s for %d damage!\n", enemy.Name, damage)

			if enemy.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", enemy.Name)
				player.Exp += enemy.Exp
				player.Gold += enemy.Gold
				fmt.Printf("Gained %d exp and %d gold.\n", enemy.Exp, enemy.Gold)
				checkLevelUp()
				removeEnemy(enemyName)
				return
			}

			enemyDamage := enemy.Attack - player.Defense
			if enemyDamage < 1 {
				enemyDamage = 1
			}
			player.Health -= enemyDamage
			fmt.Printf("The %s hits you for %d damage!\n", enemy.Name, enemyDamage)

			if player.Health <= 0 {
				fmt.Println("You have been defeated!")
				gameRunning = false
				return
			}

		case "2":
			useItem()
		case "3":
			fmt.Println("You ran away from the fight!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func removeEnemy(enemyName string) {
	loc := locations[currentLocation]
	for i, enemy := range loc.Enemies {
		if enemy.Name == enemyName {
			loc.Enemies = append(loc.Enemies[:i], loc.Enemies[i+1:]...)
			locations[currentLocation] = loc
			break
		}
	}
}

func useItem() {
	if len(player.Inventory) == 0 {
		fmt.Println("You have no items.")
		return
	}

	fmt.Println("Your inventory:")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(player.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}

	item := player.Inventory[index-1]
	if item == "Health Potion" {
		healAmount := 30
		player.Health += healAmount
		if player.Health > player.MaxHealth {
			player.Health = player.MaxHealth
		}
		fmt.Printf("You used a Health Potion and recovered %d health.\n", healAmount)
	} else {
		fmt.Printf("You used %s, but it has no effect in battle.\n", item)
		return
	}

	player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
}

func checkLevelUp() {
	for player.Exp >= player.Level*100 {
		player.Exp -= player.Level * 100
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("Level up! You are now level %d!\n", player.Level)
	}
}

func pickUpItem(itemName string) {
	loc := locations[currentLocation]
	for i, item := range loc.Items {
		if item == itemName {
			player.Inventory = append(player.Inventory, item)
			loc.Items = append(loc.Items[:i], loc.Items[i+1:]...)
			locations[currentLocation] = loc
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
	fmt.Println("move <location> - Move to another location")
	fmt.Println("fight <enemy> - Fight an enemy")
	fmt.Println("pickup <item> - Pick up an item")
	fmt.Println("help - Show this help message")
	fmt.Println("quit - Quit the game")
	fmt.Println()
}

func main() {
	fmt.Println("Welcome to the Adventure Game!")
	fmt.Println("Type 'help' for a list of commands.")
	fmt.Println()

	initGame()
	rand.Seed(time.Now().UnixNano())

	reader := bufio.NewReader(os.Stdin)

	for gameRunning {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Fields(input)

		if len(parts) == 0 {
			continue
		}

		command := parts[0]

		switch command {
		case "status":
			showStatus()
		case "look":
			showLocation()
		case "move":
			if len(parts) < 2 {
				fmt.Println("Please specify a location to move to.")
			} else {
				moveToLocation(parts[1])
			}
		case "fight":
			if len(parts) < 2 {
				fmt.Println("Please specify an enemy to fight.")
			} else {
				fightEnemy(parts[1])
			}
		case "pickup":
			if len(parts) < 2 {
				fmt.Println("Please specify an item to pick up.")
			} else {
				pickUpItem(parts[1])
			}
		case "help":
			showHelp()
		case "quit":
			fmt.Println("Thanks for playing!")
			gameRunning = false
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}
