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
	Strength int
	Gold     int
	Inventory []string
}

type Enemy struct {
	Name     string
	Health   int
	Strength int
}

type Location struct {
	Name        string
	Description string
	Connections map[string]string
	Enemies     []Enemy
	Items       []string
}

var player Player
var currentLocation string
var locations map[string]Location

func initGame() {
	rand.Seed(time.Now().UnixNano())
	player = Player{Name: "Hero", Health: 100, Strength: 10, Gold: 0, Inventory: []string{}}
	locations = make(map[string]Location)
	locations["forest"] = Location{
		Name:        "Forest",
		Description: "A dense forest with tall trees. You hear birds chirping.",
		Connections: map[string]string{"north": "mountain", "east": "river"},
		Enemies:     []Enemy{{Name: "Goblin", Health: 30, Strength: 5}},
		Items:       []string{"sword", "potion"},
	}
	locations["mountain"] = Location{
		Name:        "Mountain",
		Description: "A rocky mountain path. The air is thin here.",
		Connections: map[string]string{"south": "forest", "west": "cave"},
		Enemies:     []Enemy{{Name: "Troll", Health: 50, Strength: 8}},
		Items:       []string{"shield"},
	}
	locations["river"] = Location{
		Name:        "River",
		Description: "A flowing river with clear water. Fish are swimming.",
		Connections: map[string]string{"west": "forest", "south": "village"},
		Enemies:     []Enemy{{Name: "Pirate", Health: 40, Strength: 7}},
		Items:       []string{"fishing_rod"},
	}
	locations["village"] = Location{
		Name:        "Village",
		Description: "A peaceful village with friendly villagers.",
		Connections: map[string]string{"north": "river", "east": "shop"},
		Enemies:     []Enemy{},
		Items:       []string{"bread"},
	}
	locations["shop"] = Location{
		Name:        "Shop",
		Description: "A small shop selling various items.",
		Connections: map[string]string{"west": "village"},
		Enemies:     []Enemy{},
		Items:       []string{},
	}
	locations["cave"] = Location{
		Name:        "Cave",
		Description: "A dark cave. It smells damp and mysterious.",
		Connections: map[string]string{"east": "mountain"},
		Enemies:     []Enemy{{Name: "Dragon", Health: 100, Strength: 15}},
		Items:       []string{"treasure"},
	}
	currentLocation = "forest"
}

func describeLocation() {
	loc := locations[currentLocation]
	fmt.Printf("You are in the %s. %s\n", loc.Name, loc.Description)
	if len(loc.Connections) > 0 {
		fmt.Print("You can go to: ")
		first := true
		for dir := range loc.Connections {
			if !first {
				fmt.Print(", ")
			}
			fmt.Print(dir)
			first = false
		}
		fmt.Println()
	}
	if len(loc.Enemies) > 0 {
		fmt.Print("Enemies here: ")
		for i, enemy := range loc.Enemies {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s (Health: %d)", enemy.Name, enemy.Health)
		}
		fmt.Println()
	}
	if len(loc.Items) > 0 {
		fmt.Print("Items here: ")
		for i, item := range loc.Items {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(item)
		}
		fmt.Println()
	}
}

func move(direction string) {
	loc := locations[currentLocation]
	if newLoc, ok := loc.Connections[direction]; ok {
		currentLocation = newLoc
		fmt.Printf("You move %s to the %s.\n", direction, locations[currentLocation].Name)
		describeLocation()
	} else {
		fmt.Println("You can't go that way.")
	}
}

func fight() {
	loc := locations[currentLocation]
	if len(loc.Enemies) == 0 {
		fmt.Println("There are no enemies to fight here.")
		return
	}
	enemy := &loc.Enemies[0] // Simplified: fight first enemy
	fmt.Printf("You attack the %s!\n", enemy.Name)
	for player.Health > 0 && enemy.Health > 0 {
		// Player attacks
		damage := rand.Intn(player.Strength) + 1
		enemy.Health -= damage
		fmt.Printf("You hit the %s for %d damage. %s's health: %d\n", enemy.Name, damage, enemy.Name, enemy.Health)
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			goldEarned := rand.Intn(20) + 10
			player.Gold += goldEarned
			fmt.Printf("You found %d gold. Total gold: %d\n", goldEarned, player.Gold)
			loc.Enemies = loc.Enemies[1:] // Remove defeated enemy
			locations[currentLocation] = loc
			break
		}
		// Enemy attacks
		damage = rand.Intn(enemy.Strength) + 1
		player.Health -= damage
		fmt.Printf("The %s hits you for %d damage. Your health: %d\n", enemy.Name, damage, player.Health)
		if player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			os.Exit(0)
		}
	}
}

func takeItem(itemName string) {
	loc := locations[currentLocation]
	for i, item := range loc.Items {
		if item == itemName {
			player.Inventory = append(player.Inventory, item)
			loc.Items = append(loc.Items[:i], loc.Items[i+1:]...)
			locations[currentLocation] = loc
			fmt.Printf("You took the %s.\n", item)
			return
		}
	}
	fmt.Println("Item not found here.")
}

func useItem(itemName string) {
	for i, item := range player.Inventory {
		if item == itemName {
			switch item {
			case "potion":
				healAmount := 20
				player.Health += healAmount
				fmt.Printf("You used a potion and healed %d health. Your health: %d\n", healAmount, player.Health)
				player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			case "sword":
				player.Strength += 5
				fmt.Printf("You equipped the sword. Your strength is now %d.\n", player.Strength)
				player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			case "shield":
				player.Health += 10 // Simple bonus
				fmt.Printf("You equipped the shield. Your health is now %d.\n", player.Health)
				player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			default:
				fmt.Printf("You can't use the %s.\n", item)
			}
			return
		}
	}
	fmt.Println("Item not in inventory.")
}

func showInventory() {
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Print("Inventory: ")
	for i, item := range player.Inventory {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(item)
	}
	fmt.Printf(" | Gold: %d | Health: %d | Strength: %d\n", player.Gold, player.Health, player.Strength)
}

func buyItem() {
	if currentLocation != "shop" {
		fmt.Println("You can only buy items in the shop.")
		return
	}
	fmt.Println("Shop items available: sword (50 gold), potion (20 gold), shield (30 gold)")
	fmt.Print("Enter item to buy: ")
	reader := bufio.NewReader(os.Stdin)
	item, _ := reader.ReadString('\n')
	item = strings.TrimSpace(item)
	switch item {
	case "sword":
		if player.Gold >= 50 {
			player.Gold -= 50
			player.Inventory = append(player.Inventory, "sword")
			fmt.Println("You bought a sword.")
		} else {
			fmt.Println("Not enough gold.")
		}
	case "potion":
		if player.Gold >= 20 {
			player.Gold -= 20
			player.Inventory = append(player.Inventory, "potion")
			fmt.Println("You bought a potion.")
		} else {
			fmt.Println("Not enough gold.")
		}
	case "shield":
		if player.Gold >= 30 {
			player.Gold -= 30
			player.Inventory = append(player.Inventory, "shield")
			fmt.Println("You bought a shield.")
		} else {
			fmt.Println("Not enough gold.")
		}
	default:
		fmt.Println("Item not available.")
	}
}

func help() {
	fmt.Println("Available commands:")
	fmt.Println("  go <direction> - Move in a direction (e.g., go north)")
	fmt.Println("  look - Describe the current location")
	fmt.Println("  fight - Fight enemies in the location")
	fmt.Println("  take <item> - Take an item from the location")
	fmt.Println("  use <item> - Use an item from inventory")
	fmt.Println("  inventory - Show your inventory and stats")
	fmt.Println("  buy - Buy items in the shop")
	fmt.Println("  help - Show this help message")
	fmt.Println("  quit - Quit the game")
}

func main() {
	initGame()
	fmt.Println("Welcome to the Text Adventure Game! Type 'help' for commands.")
	describeLocation()
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}
		command := parts[0]
		args := parts[1:]
		switch command {
		case "go":
			if len(args) == 1 {
				move(args[0])
			} else {
				fmt.Println("Usage: go <direction>")
			}
		case "look":
			describeLocation()
		case "fight":
			fight()
		case "take":
			if len(args) == 1 {
				takeItem(args[0])
			} else {
				fmt.Println("Usage: take <item>")
			}
		case "use":
			if len(args) == 1 {
				useItem(args[0])
			} else {
				fmt.Println("Usage: use <item>")
			}
		case "inventory":
			showInventory()
		case "buy":
			buyItem()
		case "help":
			help()
		case "quit":
			fmt.Println("Thanks for playing!")
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}