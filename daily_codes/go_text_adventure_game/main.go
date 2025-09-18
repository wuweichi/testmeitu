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
	Inventory []string
}

type Enemy struct {
	Name     string
	Health   int
	Attack   int
	Defense  int
	LootGold int
	LootItem string
}

type Location struct {
	Name        string
	Description string
	Enemies     []Enemy
	Items       []string
	Exits       map[string]string
}

var player Player
var locations map[string]Location
var currentLocation string

func initGame() {
	rand.Seed(time.Now().UnixNano())
	player = Player{
		Name:      "Hero",
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Gold:      0,
		Inventory: []string{},
	}
	locations = map[string]Location{
		"forest": {
			Name:        "Forest",
			Description: "You are in a dense forest. Trees surround you, and you hear birds chirping.",
			Enemies: []Enemy{
				{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, LootGold: 10, LootItem: "Small Potion"},
			},
			Items: []string{"Herbs", "Stick"},
			Exits: map[string]string{"north": "cave", "east": "village"},
		},
		"cave": {
			Name:        "Cave",
			Description: "A dark and damp cave. You can see glowing mushrooms and hear dripping water.",
			Enemies: []Enemy{
				{Name: "Bat", Health: 20, Attack: 3, Defense: 1, LootGold: 5, LootItem: ""},
				{Name: "Spider", Health: 25, Attack: 4, Defense: 1, LootGold: 8, LootItem: "Spider Silk"},
			},
			Items: []string{"Torch", "Gold Nugget"},
			Exits: map[string]string{"south": "forest", "west": "dungeon"},
		},
		"village": {
			Name:        "Village",
			Description: "A peaceful village with wooden houses. Villagers are going about their day.",
			Enemies:     []Enemy{},
			Items:       []string{"Bread", "Water"},
			Exits:       map[string]string{"west": "forest", "north": "shop"},
		},
		"shop": {
			Name:        "Shop",
			Description: "A small shop selling various items. The shopkeeper smiles at you.",
			Enemies:     []Enemy{},
			Items:       []string{},
			Exits:       map[string]string{"south": "village"},
		},
		"dungeon": {
			Name:        "Dungeon",
			Description: "A creepy dungeon with chains on the walls. It smells of decay.",
			Enemies: []Enemy{
				{Name: "Skeleton", Health: 40, Attack: 6, Defense: 3, LootGold: 15, LootItem: "Bone"},
				{Name: "Zombie", Health: 35, Attack: 5, Defense: 2, LootGold: 12, LootItem: "Rotten Flesh"},
			},
			Items: []string{"Key", "Old Book"},
			Exits: map[string]string{"east": "cave"},
		},
	}
	currentLocation = "forest"
}

func main() {
	initGame()
	fmt.Println("Welcome to the Text Adventure Game!")
	fmt.Println("Type 'help' for a list of commands.")
	gameLoop()
}

func gameLoop() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("\nYou are at: %s\n", locations[currentLocation].Name)
		fmt.Println(locations[currentLocation].Description)
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if input == "quit" {
			fmt.Println("Thanks for playing!")
			break
		}
		handleInput(input)
	}
}

func handleInput(input string) {
	switch input {
	case "help":
		printHelp()
	case "look":
		lookAround()
	case "inventory":
		printInventory()
	case "stats":
		printStats()
	case "north", "south", "east", "west":
		move(input)
	case "fight":
		fightEnemy()
	case "take":
		takeItem()
	case "use":
		useItem()
	case "shop":
		if currentLocation == "shop" {
			shop()
		} else {
			fmt.Println("There is no shop here.")
		}
	default:
		fmt.Println("Unknown command. Type 'help' for options.")
	}
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  help       - Show this help message")
	fmt.Println("  look       - Look around the current location")
	fmt.Println("  inventory  - Show your inventory")
	fmt.Println("  stats      - Show your stats")
	fmt.Println("  north      - Move north")
	fmt.Println("  south      - Move south")
	fmt.Println("  east       - Move east")
	fmt.Println("  west       - Move west")
	fmt.Println("  fight      - Fight an enemy if present")
	fmt.Println("  take       - Take an item from the location")
	fmt.Println("  use        - Use an item from inventory")
	fmt.Println("  shop       - Access shop if in village shop")
	fmt.Println("  quit       - Quit the game")
}

func lookAround() {
	loc := locations[currentLocation]
	fmt.Printf("You see: %s\n", loc.Description)
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
	fmt.Println("Exits:")
	for direction, exit := range loc.Exits {
		fmt.Printf("  %s to %s\n", direction, exit)
	}
}

func printInventory() {
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("Inventory:")
	for _, item := range player.Inventory {
		fmt.Printf("  - %s\n", item)
	}
}

func printStats() {
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Gold: %d\n", player.Gold)
}

func move(direction string) {
	loc := locations[currentLocation]
	if exit, exists := loc.Exits[direction]; exists {
		currentLocation = exit
		fmt.Printf("You move %s to %s.\n", direction, exit)
	} else {
		fmt.Println("You cannot go that way.")
	}
}

func fightEnemy() {
	loc := locations[currentLocation]
	if len(loc.Enemies) == 0 {
		fmt.Println("There are no enemies to fight here.")
		return
	}
	enemy := &loc.Enemies[0] // Fight the first enemy
	fmt.Printf("You engage in combat with %s!\n", enemy.Name)
	for player.Health > 0 && enemy.Health > 0 {
		// Player attack
		damage := player.Attack - enemy.Defense
		if damage < 0 {
			damage = 0
		}
		enemy.Health -= damage
		fmt.Printf("You hit %s for %d damage. %s's health: %d\n", enemy.Name, damage, enemy.Name, enemy.Health)
		if enemy.Health <= 0 {
			fmt.Printf("You defeated %s!\n", enemy.Name)
			player.Gold += enemy.LootGold
			if enemy.LootItem != "" {
				player.Inventory = append(player.Inventory, enemy.LootItem)
				fmt.Printf("You found %s!\n", enemy.LootItem)
			}
			// Remove enemy from location
			loc.Enemies = loc.Enemies[1:]
			locations[currentLocation] = loc
			break
		}
		// Enemy attack
		damage = enemy.Attack - player.Defense
		if damage < 0 {
			damage = 0
		}
		player.Health -= damage
		fmt.Printf("%s hits you for %d damage. Your health: %d\n", enemy.Name, damage, player.Health)
		if player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			os.Exit(0)
		}
	}
}

func takeItem() {
	loc := locations[currentLocation]
	if len(loc.Items) == 0 {
		fmt.Println("There are no items to take here.")
		return
	}
	item := loc.Items[0]
	player.Inventory = append(player.Inventory, item)
	loc.Items = loc.Items[1:]
	locations[currentLocation] = loc
	fmt.Printf("You took %s.\n", item)
}

func useItem() {
	if len(player.Inventory) == 0 {
		fmt.Println("You have no items to use.")
		return
	}
	fmt.Println("Select an item to use:")
	for i, item := range player.Inventory {
		fmt.Printf("%d: %s\n", i+1, item)
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter item number: ")
	if !scanner.Scan() {
		return
	}
	input := strings.TrimSpace(scanner.Text())
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(player.Inventory) {
		fmt.Println("Invalid selection.")
		return
	}
	item := player.Inventory[index-1]
	switch item {
	case "Small Potion":
		healAmount := 20
		player.Health += healAmount
		if player.Health > player.MaxHealth {
			player.Health = player.MaxHealth
		}
		fmt.Printf("You used %s and healed %d health. Current health: %d\n", item, healAmount, player.Health)
		// Remove item from inventory
		player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
	default:
		fmt.Printf("You cannot use %s right now.\n", item)
	}
}

func shop() {
	fmt.Println("Welcome to the shop!")
	fmt.Println("Items for sale:")
	fmt.Println("1. Health Potion - 20 gold (Restores 50 health)")
	fmt.Println("2. Sword - 50 gold (+5 attack)")
	fmt.Println("3. Shield - 40 gold (+3 defense)")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter item number to buy (or 'leave' to exit): ")
	if !scanner.Scan() {
		return
	}
	input := strings.TrimSpace(scanner.Text())
	if input == "leave" {
		fmt.Println("You leave the shop.")
		return
	}
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > 3 {
		fmt.Println("Invalid selection.")
		return
	}
	switch index {
	case 1:
		if player.Gold >= 20 {
			player.Gold -= 20
			player.Inventory = append(player.Inventory, "Health Potion")
			fmt.Println("You bought a Health Potion.")
		} else {
			fmt.Println("Not enough gold!")
		}
	case 2:
		if player.Gold >= 50 {
			player.Gold -= 50
			player.Attack += 5
			fmt.Println("You bought a Sword. Attack increased by 5.")
		} else {
			fmt.Println("Not enough gold!")
		}
	case 3:
		if player.Gold >= 40 {
			player.Gold -= 40
			player.Defense += 3
			fmt.Println("You bought a Shield. Defense increased by 3.")
		} else {
			fmt.Println("Not enough gold!")
		}
	}
}
