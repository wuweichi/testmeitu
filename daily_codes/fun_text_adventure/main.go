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
	Name     string
	Health   int
	MaxHealth int
	Attack   int
	Defense  int
	Gold     int
	Level    int
	Exp      int
	Inventory []string
}

type Enemy struct {
	Name     string
	Health   int
	Attack   int
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
		Name:     "",
		Health:   100,
		MaxHealth: 100,
		Attack:   10,
		Defense:  5,
		Gold:     50,
		Level:    1,
		Exp:      0,
		Inventory: []string{"Health Potion", "Health Potion"},
	}

	locations = map[string]Location{
		"forest": {
			Name:        "Mysterious Forest",
			Description: "You are in a dense forest. The trees are tall and the air is fresh. You can hear birds chirping in the distance.",
			Enemies: []Enemy{
				{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, GoldDrop: 10, ExpDrop: 20},
				{Name: "Wolf", Health: 25, Attack: 7, Defense: 1, GoldDrop: 5, ExpDrop: 15},
			},
			Items:       []string{"Health Potion", "Iron Sword"},
			Connections: []string{"village", "cave"},
		},
		"village": {
			Name:        "Peaceful Village",
			Description: "You are in a small village. People are going about their daily lives. There's a shop and an inn here.",
			Enemies:     []Enemy{},
			Items:       []string{"Bread", "Water"},
			Connections: []string{"forest"},
		},
		"cave": {
			Name:        "Dark Cave",
			Description: "You are in a dark and damp cave. It's hard to see, and you can hear strange noises echoing.",
			Enemies: []Enemy{
				{Name: "Bat", Health: 15, Attack: 3, Defense: 0, GoldDrop: 2, ExpDrop: 10},
				{Name: "Spider", Health: 20, Attack: 4, Defense: 1, GoldDrop: 3, ExpDrop: 12},
				{Name: "Troll", Health: 50, Attack: 12, Defense: 5, GoldDrop: 30, ExpDrop: 50},
			},
			Items:       []string{"Torch", "Gold Coin"},
			Connections: []string{"forest"},
		},
	}

	currentLocation = "forest"
}

func printStatus() {
	fmt.Printf("\n=== Status ===\n")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Exp: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Inventory: %v\n", player.Inventory)
	fmt.Printf("Location: %s\n", locations[currentLocation].Name)
}

func printLocation() {
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
	fmt.Println("You can go to:")
	for _, conn := range loc.Connections {
		fmt.Printf("  - %s\n", locations[conn].Name)
	}
}

func moveTo(location string) {
	if contains(locations[currentLocation].Connections, location) {
		currentLocation = location
		fmt.Printf("You moved to %s.\n", locations[location].Name)
		printLocation()
	} else {
		fmt.Println("You can't go there from here.")
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
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
		fmt.Println("Choose an action: (1) Attack (2) Use Item (3) Run")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		action := scanner.Text()

		switch action {
		case "1":
			playerDamage := player.Attack - enemy.Defense
			if playerDamage < 1 {
				playerDamage = 1
			}
			enemy.Health -= playerDamage
			fmt.Printf("You hit the %s for %d damage.\n", enemy.Name, playerDamage)
			if enemy.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", enemy.Name)
				player.Gold += enemy.GoldDrop
				player.Exp += enemy.ExpDrop
				fmt.Printf("You gained %d gold and %d experience.\n", enemy.GoldDrop, enemy.ExpDrop)
				checkLevelUp()
				removeEnemy(enemyName)
				return
			}
		case "2":
			useItem()
			continue
		case "3":
			fmt.Println("You ran away!")
			return
		default:
			fmt.Println("Invalid action.")
			continue
		}

		enemyDamage := enemy.Attack - player.Defense
		if enemyDamage < 1 {
			enemyDamage = 1
		}
		player.Health -= enemyDamage
		fmt.Printf("The %s hit you for %d damage.\n", enemy.Name, enemyDamage)
		if player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
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
	fmt.Println("Choose an item to use (or 0 to cancel):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice, err := strconv.Atoi(scanner.Text())
	if err != nil || choice < 0 || choice > len(player.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}
	if choice == 0 {
		return
	}
	item := player.Inventory[choice-1]
	if item == "Health Potion" {
		healAmount := 30
		player.Health += healAmount
		if player.Health > player.MaxHealth {
			player.Health = player.MaxHealth
		}
		fmt.Printf("You used a Health Potion and restored %d health.\n", healAmount)
		player.Inventory = append(player.Inventory[:choice-1], player.Inventory[choice:]...)
	} else {
		fmt.Println("You can't use that item in battle.")
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
		fmt.Printf("\n=== Level Up! ===\n")
		fmt.Printf("You are now level %d!\n", player.Level)
		fmt.Printf("Max Health increased to %d.\n", player.MaxHealth)
		fmt.Printf("Attack increased to %d.\n", player.Attack)
		fmt.Printf("Defense increased to %d.\n", player.Defense)
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

func shop() {
	if currentLocation != "village" {
		fmt.Println("There's no shop here.")
		return
	}
	fmt.Println("\n=== Shop ===")
	fmt.Println("Items for sale:")
	fmt.Println("1. Health Potion - 10 gold")
	fmt.Println("2. Iron Sword - 50 gold")
	fmt.Println("3. Leather Armor - 40 gold")
	fmt.Println("Enter the number to buy (or 0 to leave):")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice, err := strconv.Atoi(scanner.Text())
	if err != nil || choice < 0 || choice > 3 {
		fmt.Println("Invalid choice.")
		return
	}
	if choice == 0 {
		return
	}
	switch choice {
	case 1:
		if player.Gold >= 10 {
			player.Gold -= 10
			player.Inventory = append(player.Inventory, "Health Potion")
			fmt.Println("You bought a Health Potion.")
		} else {
			fmt.Println("Not enough gold.")
		}
	case 2:
		if player.Gold >= 50 {
			player.Gold -= 50
			player.Attack += 5
			fmt.Println("You bought an Iron Sword. Attack increased by 5.")
		} else {
			fmt.Println("Not enough gold.")
		}
	case 3:
		if player.Gold >= 40 {
			player.Gold -= 40
			player.Defense += 3
			fmt.Println("You bought Leather Armor. Defense increased by 3.")
		} else {
			fmt.Println("Not enough gold.")
		}
	}
}

func rest() {
	if currentLocation != "village" {
		fmt.Println("You can only rest in the village.")
		return
	}
	player.Health = player.MaxHealth
	fmt.Println("You rested and restored your health to full.")
}

func main() {
	fmt.Println("Welcome to the Text Adventure Game!")
	fmt.Print("Enter your name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	player.Name = scanner.Text()
	if player.Name == "" {
		player.Name = "Adventurer"
	}

	initGame()
	fmt.Printf("Hello, %s! Let's begin your adventure.\n", player.Name)
	printLocation()

	for {
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("Commands: status, look, move [location], fight [enemy], pickup [item], shop, rest, quit")
		scanner.Scan()
		input := scanner.Text()
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}

		switch parts[0] {
		case "status":
			printStatus()
		case "look":
			printLocation()
		case "move":
			if len(parts) < 2 {
				fmt.Println("Please specify a location.")
			} else {
				moveTo(parts[1])
			}
		case "fight":
			if len(parts) < 2 {
				fmt.Println("Please specify an enemy.")
			} else {
				fightEnemy(parts[1])
			}
		case "pickup":
			if len(parts) < 2 {
				fmt.Println("Please specify an item.")
			} else {
				pickUpItem(parts[1])
			}
		case "shop":
			shop()
		case "rest":
			rest()
		case "quit":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Unknown command.")
		}
	}
}
