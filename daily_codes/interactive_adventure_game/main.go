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

type Player struct {
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

var player Player
var locations map[string]Location
var currentLocation string

func initGame() {
	player = Player{
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
			{Name: "Wolf", Health: 25, Attack: 7, Defense: 1, Exp: 8, Gold: 3},
		},
		Items:       []string{"Health Potion", "Wooden Shield"},
		Connections: []string{"village", "cave"},
	}
	locations["village"] = Location{
		Name:        "Village",
		Description: "A peaceful village with friendly villagers.",
		Enemies:     []Enemy{},
		Items:       []string{"Bread", "Health Potion"},
		Connections: []string{"forest"},
	}
	locations["cave"] = Location{
		Name:        "Cave",
		Description: "A dark cave with eerie echoes.",
		Enemies: []Enemy{
			{Name: "Bat", Health: 15, Attack: 3, Defense: 0, Exp: 5, Gold: 1},
			{Name: "Spider", Health: 20, Attack: 4, Defense: 1, Exp: 7, Gold: 2},
		},
		Items:       []string{"Torch", "Gold Coin"},
		Connections: []string{"forest", "dungeon"},
	}
	locations["dungeon"] = Location{
		Name:        "Dungeon",
		Description: "A dangerous dungeon filled with traps and monsters.",
		Enemies: []Enemy{
			{Name: "Skeleton", Health: 40, Attack: 8, Defense: 3, Exp: 15, Gold: 10},
			{Name: "Orc", Health: 50, Attack: 12, Defense: 4, Exp: 20, Gold: 15},
		},
		Items:       []string{"Steel Sword", "Magic Amulet"},
		Connections: []string{"cave"},
	}

	currentLocation = "village"
	rand.Seed(time.Now().UnixNano())
}

func showStatus() {
	fmt.Printf("\n=== Player Status ===\n")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Exp: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Inventory: %v\n", player.Inventory)
}

func showLocation() {
	loc := locations[currentLocation]
	fmt.Printf("\n=== %s ===\n", loc.Name)
	fmt.Println(loc.Description)
	if len(loc.Connections) > 0 {
		fmt.Printf("Connections: %v\n", loc.Connections)
	}
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
		fmt.Printf("Items: %v\n", loc.Items)
	}
}

func move(destination string) {
	loc := locations[currentLocation]
	for _, conn := range loc.Connections {
		if conn == destination {
			currentLocation = destination
			fmt.Printf("You moved to %s.\n", destination)
			showLocation()
			return
		}
	}
	fmt.Printf("Cannot move to %s from here.\n", destination)
}

func fight() {
	loc := locations[currentLocation]
	if len(loc.Enemies) == 0 {
		fmt.Println("No enemies here.")
		return
	}

	enemy := loc.Enemies[rand.Intn(len(loc.Enemies))]
	fmt.Printf("A wild %s appears!\n", enemy.Name)

	for player.Health > 0 && enemy.Health > 0 {
		fmt.Printf("\nYour Health: %d, %s's Health: %d\n", player.Health, enemy.Name, enemy.Health)
		fmt.Print("Choose action: (1) Attack (2) Use Item (3) Flee: ")
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
			fmt.Printf("You attack the %s for %d damage.\n", enemy.Name, damage)

			if enemy.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", enemy.Name)
				player.Exp += enemy.Exp
				player.Gold += enemy.Gold
				fmt.Printf("Gained %d exp and %d gold.\n", enemy.Exp, enemy.Gold)
				checkLevelUp()
				return
			}

			enemyDamage := enemy.Attack - player.Defense
			if enemyDamage < 1 {
				enemyDamage = 1
			}
			player.Health -= enemyDamage
			fmt.Printf("%s attacks you for %d damage.\n", enemy.Name, enemyDamage)

			if player.Health <= 0 {
				fmt.Println("You have been defeated!")
				return
			}

		case "2":
			useItem()
		case "3":
			if rand.Intn(2) == 0 {
				fmt.Println("You successfully fled!")
				return
			} else {
				fmt.Println("You failed to flee!")
				enemyDamage := enemy.Attack - player.Defense
				if enemyDamage < 1 {
					enemyDamage = 1
				}
				player.Health -= enemyDamage
				fmt.Printf("%s attacks you for %d damage.\n", enemy.Name, enemyDamage)
				if player.Health <= 0 {
					fmt.Println("You have been defeated!")
					return
				}
			}
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func useItem() {
	if len(player.Inventory) == 0 {
		fmt.Println("No items in inventory.")
		return
	}

	fmt.Println("Inventory:")
	for i, item := range player.Inventory {
		fmt.Printf("%d: %s\n", i+1, item)
	}
	fmt.Print("Choose item to use (number): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(player.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}

	item := player.Inventory[index-1]
	switch item {
	case "Health Potion":
		healAmount := 30
		player.Health += healAmount
		if player.Health > player.MaxHealth {
			player.Health = player.MaxHealth
		}
		fmt.Printf("Used Health Potion. Healed %d health. Current health: %d\n", healAmount, player.Health)
		player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
	default:
		fmt.Printf("Cannot use %s in battle.\n", item)
	}
}

func checkLevelUp() {
	requiredExp := player.Level * 100
	if player.Exp >= requiredExp {
		player.Level++
		player.Exp -= requiredExp
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("Level up! You are now level %d.\n", player.Level)
		showStatus()
	}
}

func explore() {
	loc := locations[currentLocation]
	if len(loc.Items) == 0 {
		fmt.Println("Nothing to find here.")
		return
	}

	item := loc.Items[rand.Intn(len(loc.Items))]
	fmt.Printf("You found a %s!\n", item)
	player.Inventory = append(player.Inventory, item)
	loc.Items = removeItem(loc.Items, item)
	locations[currentLocation] = loc
}

func removeItem(items []string, item string) []string {
	for i, it := range items {
		if it == item {
			return append(items[:i], items[i+1:]...)
		}
	}
	return items
}

func shop() {
	if currentLocation != "village" {
		fmt.Println("No shop here.")
		return
	}

	fmt.Println("\n=== Shop ===")
	fmt.Println("1. Health Potion - 10 gold")
	fmt.Println("2. Steel Sword - 50 gold")
	fmt.Println("3. Magic Amulet - 100 gold")
	fmt.Println("4. Exit")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		if player.Gold >= 10 {
			player.Gold -= 10
			player.Inventory = append(player.Inventory, "Health Potion")
			fmt.Println("Bought Health Potion.")
		} else {
			fmt.Println("Not enough gold.")
		}
	case "2":
		if player.Gold >= 50 {
			player.Gold -= 50
			player.Attack += 5
			fmt.Println("Bought Steel Sword. Attack increased by 5.")
		} else {
			fmt.Println("Not enough gold.")
		}
	case "3":
		if player.Gold >= 100 {
			player.Gold -= 100
			player.MaxHealth += 50
			player.Health = player.MaxHealth
			fmt.Println("Bought Magic Amulet. Max health increased by 50.")
		} else {
			fmt.Println("Not enough gold.")
		}
	case "4":
		fmt.Println("Exited shop.")
	default:
		fmt.Println("Invalid choice.")
	}
}

func main() {
	initGame()
	fmt.Println("Welcome to the Interactive Adventure Game!")
	fmt.Println("Type 'help' for a list of commands.")

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "help":
			fmt.Println("Commands: status, look, move <location>, fight, explore, shop, use, quit")
		case "status":
			showStatus()
		case "look":
			showLocation()
		case "fight":
			fight()
			if player.Health <= 0 {
				fmt.Println("Game Over!")
				return
			}
		case "explore":
			explore()
		case "shop":
			shop()
		case "use":
			useItem()
		case "quit":
			fmt.Println("Thanks for playing!")
			return
		default:
			if strings.HasPrefix(input, "move ") {
				destination := strings.TrimPrefix(input, "move ")
				move(destination)
			} else {
				fmt.Println("Unknown command. Type 'help' for a list of commands.")
			}
		}
	}
}
