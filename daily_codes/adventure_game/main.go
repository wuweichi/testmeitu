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

type Item struct {
	Name        string
	Description string
	Effect      string
	Value       int
}

type Location struct {
	Name        string
	Description string
	Enemies     []Enemy
	Items       []Item
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
		Inventory: []string{"Health Potion", "Health Potion"},
	}

	locations = make(map[string]Location)

	locations["forest"] = Location{
		Name:        "Forest",
		Description: "A dense forest with tall trees and mysterious sounds.",
		Enemies: []Enemy{
			{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Exp: 10, Gold: 5},
			{Name: "Wolf", Health: 25, Attack: 8, Defense: 1, Exp: 8, Gold: 3},
		},
		Items: []Item{
			{Name: "Health Potion", Description: "Restores 30 health", Effect: "heal", Value: 30},
			{Name: "Iron Sword", Description: "A basic sword", Effect: "weapon", Value: 5},
		},
		Connections: []string{"village", "cave"},
	}

	locations["village"] = Location{
		Name:        "Village",
		Description: "A peaceful village with friendly inhabitants.",
		Enemies:     []Enemy{},
		Items: []Item{
			{Name: "Health Potion", Description: "Restores 30 health", Effect: "heal", Value: 30},
			{Name: "Leather Armor", Description: "Basic armor", Effect: "armor", Value: 3},
		},
		Connections: []string{"forest", "shop"},
	}

	locations["cave"] = Location{
		Name:        "Cave",
		Description: "A dark cave with eerie echoes.",
		Enemies: []Enemy{
			{Name: "Bat", Health: 15, Attack: 3, Defense: 0, Exp: 5, Gold: 1},
			{Name: "Spider", Health: 20, Attack: 6, Defense: 1, Exp: 7, Gold: 2},
			{Name: "Troll", Health: 50, Attack: 12, Defense: 5, Exp: 25, Gold: 15},
		},
		Items: []Item{
			{Name: "Magic Ring", Description: "Increases attack", Effect: "attack", Value: 5},
			{Name: "Gold Coin", Description: "Worth 10 gold", Effect: "gold", Value: 10},
		},
		Connections: []string{"forest", "dungeon"},
	}

	locations["shop"] = Location{
		Name:        "Shop",
		Description: "A small shop selling various items.",
		Enemies:     []Enemy{},
		Items: []Item{
			{Name: "Health Potion", Description: "Restores 30 health", Effect: "heal", Value: 30},
			{Name: "Strength Potion", Description: "Increases attack", Effect: "attack", Value: 5},
			{Name: "Defense Potion", Description: "Increases defense", Effect: "defense", Value: 3},
		},
		Connections: []string{"village"},
	}

	locations["dungeon"] = Location{
		Name:        "Dungeon",
		Description: "A dangerous dungeon filled with powerful enemies.",
		Enemies: []Enemy{
			{Name: "Skeleton", Health: 40, Attack: 10, Defense: 3, Exp: 15, Gold: 8},
			{Name: "Zombie", Health: 45, Attack: 8, Defense: 4, Exp: 12, Gold: 6},
			{Name: "Dragon", Health: 100, Attack: 20, Defense: 10, Exp: 50, Gold: 50},
		},
		Items: []Item{
			{Name: "Dragon Scale", Description: "A rare item", Effect: "special", Value: 100},
			{Name: "Magic Staff", Description: "Powerful weapon", Effect: "weapon", Value: 15},
		},
		Connections: []string{"cave"},
	}

	currentLocation = "forest"
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
	fmt.Printf("Connections: %v\n", loc.Connections)
	if len(loc.Enemies) > 0 {
		fmt.Printf("Enemies: ")
		for i, enemy := range loc.Enemies {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", enemy.Name)
		}
		fmt.Println()
	}
	if len(loc.Items) > 0 {
		fmt.Printf("Items: ")
		for i, item := range loc.Items {
			if i > 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", item.Name)
		}
		fmt.Println()
	}
	fmt.Println()
}

func moveTo(newLocation string) {
	loc := locations[currentLocation]
	for _, conn := range loc.Connections {
		if conn == newLocation {
			currentLocation = newLocation
			fmt.Printf("Moved to %s.\n\n", newLocation)
			showLocation()
			return
		}
	}
	fmt.Println("Cannot move to that location from here.")
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
		fmt.Println("Enemy not found in this location.")
		return
	}

	fmt.Printf("Fighting %s!\n", enemy.Name)
	for player.Health > 0 && enemy.Health > 0 {
		playerDamage := player.Attack - enemy.Defense
		if playerDamage < 1 {
			playerDamage = 1
		}
		enemy.Health -= playerDamage
		fmt.Printf("You hit %s for %d damage. %s health: %d\n", enemy.Name, playerDamage, enemy.Name, enemy.Health)

		if enemy.Health <= 0 {
			break
		}

		enemyDamage := enemy.Attack - player.Defense
		if enemyDamage < 1 {
			enemyDamage = 1
		}
		player.Health -= enemyDamage
		fmt.Printf("%s hits you for %d damage. Your health: %d\n", enemy.Name, enemyDamage, player.Health)
	}

	if player.Health <= 0 {
		fmt.Println("You have been defeated!")
		os.Exit(0)
	} else {
		fmt.Printf("You defeated %s!\n", enemy.Name)
		player.Exp += enemy.Exp
		player.Gold += enemy.Gold
		fmt.Printf("Gained %d exp and %d gold.\n", enemy.Exp, enemy.Gold)
		checkLevelUp()
		removeEnemy(enemyName)
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
			fmt.Printf("Used %s.\n", itemName)
			if itemName == "Health Potion" {
				player.Health += 30
				if player.Health > player.MaxHealth {
					player.Health = player.MaxHealth
				}
				fmt.Printf("Restored 30 health. Current health: %d\n", player.Health)
			}
			return
		}
	}
	fmt.Println("Item not found in inventory.")
}

func pickUpItem(itemName string) {
	loc := locations[currentLocation]
	for i, item := range loc.Items {
		if item.Name == itemName {
			player.Inventory = append(player.Inventory, item.Name)
			loc.Items = append(loc.Items[:i], loc.Items[i+1:]...)
			locations[currentLocation] = loc
			fmt.Printf("Picked up %s.\n", item.Name)
			return
		}
	}
	fmt.Println("Item not found in this location.")
}

func buyItem(itemName string) {
	if currentLocation != "shop" {
		fmt.Println("You can only buy items in the shop.")
		return
	}
	loc := locations[currentLocation]
	for _, item := range loc.Items {
		if item.Name == itemName {
			if player.Gold >= 10 {
				player.Gold -= 10
				player.Inventory = append(player.Inventory, item.Name)
				fmt.Printf("Bought %s for 10 gold.\n", item.Name)
			} else {
				fmt.Println("Not enough gold.")
			}
			return
		}
	}
	fmt.Println("Item not available in shop.")
}

func main() {
	initGame()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Adventure Game!")
	fmt.Println("Type 'help' for a list of commands.")
	showLocation()

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Fields(input)

		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  status - Show player status")
			fmt.Println("  look - Show current location")
			fmt.Println("  move <location> - Move to another location")
			fmt.Println("  fight <enemy> - Fight an enemy")
			fmt.Println("  use <item> - Use an item from inventory")
			fmt.Println("  pickup <item> - Pick up an item from location")
			fmt.Println("  buy <item> - Buy an item from shop")
			fmt.Println("  quit - Exit the game")
		case "status":
			showStatus()
		case "look":
			showLocation()
		case "move":
			if len(args) < 2 {
				fmt.Println("Usage: move <location>")
			} else {
				moveTo(args[1])
			}
		case "fight":
			if len(args) < 2 {
				fmt.Println("Usage: fight <enemy>")
			} else {
				fightEnemy(args[1])
			}
		case "use":
			if len(args) < 2 {
				fmt.Println("Usage: use <item>")
			} else {
				useItem(args[1])
			}
		case "pickup":
			if len(args) < 2 {
				fmt.Println("Usage: pickup <item>")
			} else {
				pickUpItem(args[1])
			}
		case "buy":
			if len(args) < 2 {
				fmt.Println("Usage: buy <item>")
			} else {
				buyItem(args[1])
			}
		case "quit":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}