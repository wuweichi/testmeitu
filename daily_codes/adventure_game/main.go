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
	Strength  int
	Defense   int
	Level     int
	Experience int
	Gold      int
	Inventory []string
}

type Enemy struct {
	Name     string
	Health   int
	Strength int
	Defense  int
	GoldDrop int
	ExpDrop  int
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
		Strength:  10,
		Defense:   5,
		Level:     1,
		Experience: 0,
		Gold:      50,
		Inventory: []string{"Health Potion", "Health Potion"},
	}

	locations = make(map[string]Location)

	locations["forest"] = Location{
		Name:        "Forest",
		Description: "A dense forest with tall trees and mysterious sounds.",
		Enemies: []Enemy{
			{Name: "Goblin", Health: 30, Strength: 8, Defense: 2, GoldDrop: 10, ExpDrop: 20},
			{Name: "Wolf", Health: 25, Strength: 12, Defense: 1, GoldDrop: 5, ExpDrop: 15},
		},
		Items:       []string{"Health Potion", "Iron Sword"},
		Connections: []string{"village", "cave"},
	}

	locations["village"] = Location{
		Name:        "Village",
		Description: "A peaceful village with friendly inhabitants.",
		Enemies:     []Enemy{},
		Items:       []string{"Health Potion", "Leather Armor"},
		Connections: []string{"forest", "castle"},
	}

	locations["cave"] = Location{
		Name:        "Cave",
		Description: "A dark and damp cave with glowing mushrooms.",
		Enemies: []Enemy{
			{Name: "Bat", Health: 15, Strength: 5, Defense: 0, GoldDrop: 2, ExpDrop: 10},
			{Name: "Spider", Health: 20, Strength: 10, Defense: 1, GoldDrop: 8, ExpDrop: 25},
		},
		Items:       []string{"Gold Coin", "Magic Ring"},
		Connections: []string{"forest", "dungeon"},
	}

	locations["castle"] = Location{
		Name:        "Castle",
		Description: "A majestic castle with high walls and a grand entrance.",
		Enemies: []Enemy{
			{Name: "Guard", Health: 50, Strength: 15, Defense: 8, GoldDrop: 20, ExpDrop: 50},
			{Name: "Knight", Health: 80, Strength: 25, Defense: 12, GoldDrop: 50, ExpDrop: 100},
		},
		Items:       []string{"Steel Sword", "Health Potion"},
		Connections: []string{"village", "throne_room"},
	}

	locations["dungeon"] = Location{
		Name:        "Dungeon",
		Description: "A creepy dungeon with chains and torture devices.",
		Enemies: []Enemy{
			{Name: "Skeleton", Health: 40, Strength: 18, Defense: 5, GoldDrop: 15, ExpDrop: 40},
			{Name: "Zombie", Health: 60, Strength: 22, Defense: 3, GoldDrop: 25, ExpDrop: 60},
		},
		Items:       []string{"Health Potion", "Magic Amulet"},
		Connections: []string{"cave", "boss_room"},
	}

	locations["throne_room"] = Location{
		Name:        "Throne Room",
		Description: "The throne room of the castle, where the king resides.",
		Enemies:     []Enemy{},
		Items:       []string{"Crown", "Royal Scepter"},
		Connections: []string{"castle"},
	}

	locations["boss_room"] = Location{
		Name:        "Boss Room",
		Description: "A large chamber with a fearsome dragon.",
		Enemies: []Enemy{
			{Name: "Dragon", Health: 200, Strength: 40, Defense: 20, GoldDrop: 200, ExpDrop: 500},
		},
		Items:       []string{"Dragon Scale", "Treasure Chest"},
		Connections: []string{"dungeon"},
	}

	currentLocation = "forest"
}

func displayStatus() {
	fmt.Printf("\n=== Player Status ===\n")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Experience: %d/%d\n", player.Experience, player.Level*100)
	fmt.Printf("Strength: %d\n", player.Strength)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Inventory: %v\n", player.Inventory)
}

func displayLocation() {
	loc := locations[currentLocation]
	fmt.Printf("\n=== %s ===\n", loc.Name)
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
	fmt.Println("Connections:")
	for _, conn := range loc.Connections {
		fmt.Printf("  - %s\n", conn)
	}
}

func moveTo(newLocation string) {
	loc := locations[currentLocation]
	for _, conn := range loc.Connections {
		if conn == newLocation {
			currentLocation = newLocation
			fmt.Printf("You moved to %s.\n", newLocation)
			return
		}
	}
	fmt.Println("You can't go there from here.")
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
		fmt.Println("Choose action: (1) Attack (2) Use Item (3) Flee")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			playerDamage := player.Strength - enemy.Defense
			if playerDamage < 1 {
				playerDamage = 1
			}
			enemy.Health -= playerDamage
			fmt.Printf("You hit the %s for %d damage.\n", enemy.Name, playerDamage)
			if enemy.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", enemy.Name)
				player.Gold += enemy.GoldDrop
				player.Experience += enemy.ExpDrop
				fmt.Printf("You gained %d gold and %d experience.\n", enemy.GoldDrop, enemy.ExpDrop)
				checkLevelUp()
				removeEnemy(enemyName)
				return
			}
		case "2":
			useItem()
			continue
		case "3":
			fmt.Println("You fled from the battle.")
			return
		default:
			fmt.Println("Invalid choice.")
			continue
		}

		enemyDamage := enemy.Strength - player.Defense
		if enemyDamage < 1 {
			enemyDamage = 1
		}
		player.Health -= enemyDamage
		fmt.Printf("The %s hits you for %d damage.\n", enemy.Name, enemyDamage)
		if player.Health <= 0 {
			fmt.Println("You have been defeated! Game Over.")
			os.Exit(0)
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
	if item == "Health Potion" {
		player.Health += 30
		if player.Health > player.MaxHealth {
			player.Health = player.MaxHealth
		}
		fmt.Println("You used a Health Potion and restored 30 health.")
	} else {
		fmt.Printf("You used %s, but it has no effect in battle.\n", item)
		return
	}
	player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
}

func checkLevelUp() {
	for player.Experience >= player.Level*100 {
		player.Experience -= player.Level * 100
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Strength += 5
		player.Defense += 3
		fmt.Printf("Level up! You are now level %d.\n", player.Level)
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

func main() {
	initGame()
	fmt.Println("Welcome to the Adventure Game!")
	fmt.Println("Type 'help' for a list of commands.")

	reader := bufio.NewReader(os.Stdin)
	for {
		displayLocation()
		displayStatus()
		fmt.Print("\nWhat do you want to do? ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  help - Show this help message")
			fmt.Println("  status - Show player status")
			fmt.Println("  move <location> - Move to a connected location")
			fmt.Println("  fight <enemy> - Fight an enemy")
			fmt.Println("  pickup <item> - Pick up an item")
			fmt.Println("  use - Use an item from inventory")
			fmt.Println("  quit - Quit the game")
		case "status":
			displayStatus()
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
		case "pickup":
			if len(args) < 2 {
				fmt.Println("Usage: pickup <item>")
			} else {
				pickUpItem(args[1])
			}
		case "use":
			useItem()
		case "quit":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}