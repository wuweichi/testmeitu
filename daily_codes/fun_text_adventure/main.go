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
	Exits       map[string]string
	Items       []string
	Enemies     []Enemy
}

var player Player
var locations map[string]Location
var currentLocation string

func initGame() {
	rand.Seed(time.Now().UnixNano())
	player = Player{Name: "Hero", Health: 100, Strength: 10, Gold: 0, Inventory: []string{}}
	locations = make(map[string]Location)
	locations["forest"] = Location{
		Name:        "Forest",
		Description: "You are in a dense forest. Trees tower above you, and the air is fresh.",
		Exits:       map[string]string{"north": "cave", "east": "river"},
		Items:       []string{"sword", "potion"},
		Enemies:     []Enemy{{Name: "Goblin", Health: 30, Strength: 5}},
	}
	locations["cave"] = Location{
		Name:        "Cave",
		Description: "A dark and musty cave. You hear dripping water echoing.",
		Exits:       map[string]string{"south": "forest", "west": "mountain"},
		Items:       []string{"torch", "gold"},
		Enemies:     []Enemy{{Name: "Bat", Health: 20, Strength: 3}},
	}
	locations["river"] = Location{
		Name:        "River",
		Description: "A flowing river with clear water. Fish swim by occasionally.",
		Exits:       map[string]string{"west": "forest", "north": "village"},
		Items:       []string{"fishing_rod", "fish"},
		Enemies:     []Enemy{{Name: "Piranha", Health: 15, Strength: 2}},
	}
	locations["mountain"] = Location{
		Name:        "Mountain",
		Description: "A steep mountain path. The view from here is breathtaking.",
		Exits:       map[string]string{"east": "cave", "up": "peak"},
		Items:       []string{"rope", "gem"},
		Enemies:     []Enemy{{Name: "Eagle", Health: 25, Strength: 4}},
	}
	locations["peak"] = Location{
		Name:        "Peak",
		Description: "The highest point of the mountain. You can see the entire world below.",
		Exits:       map[string]string{"down": "mountain"},
		Items:       []string{"treasure_chest"},
		Enemies:     []Enemy{},
	}
	locations["village"] = Location{
		Name:        "Village",
		Description: "A peaceful village with friendly inhabitants. Smoke rises from chimneys.",
		Exits:       map[string]string{"south": "river", "east": "shop"},
		Items:       []string{"bread", "map"},
		Enemies:     []Enemy{},
	}
	locations["shop"] = Location{
		Name:        "Shop",
		Description: "A small shop selling various goods. The shopkeeper smiles at you.",
		Exits:       map[string]string{"west": "village"},
		Items:       []string{"potion", "armor"},
		Enemies:     []Enemy{},
	}
	currentLocation = "forest"
}

func displayLocation() {
	loc := locations[currentLocation]
	fmt.Printf("You are at: %s\n", loc.Name)
	fmt.Println(loc.Description)
	if len(loc.Exits) > 0 {
		fmt.Print("Exits: ")
		for dir := range loc.Exits {
			fmt.Printf("%s ", dir)
		}
		fmt.Println()
	}
	if len(loc.Items) > 0 {
		fmt.Printf("Items here: %s\n", strings.Join(loc.Items, ", "))
	}
	if len(loc.Enemies) > 0 {
		fmt.Printf("Enemies here: ")
		for _, enemy := range loc.Enemies {
			fmt.Printf("%s (Health: %d) ", enemy.Name, enemy.Health)
		}
		fmt.Println()
	}
}

func move(direction string) {
	loc := locations[currentLocation]
	if newLoc, ok := loc.Exits[direction]; ok {
		currentLocation = newLoc
		fmt.Printf("You move %s to %s.\n", direction, newLoc)
	} else {
		fmt.Println("You can't go that way.")
	}
}

func takeItem(itemName string) {
	loc := locations[currentLocation]
	for i, item := range loc.Items {
		if item == itemName {
			player.Inventory = append(player.Inventory, item)
			loc.Items = append(loc.Items[:i], loc.Items[i+1:]...)
			locations[currentLocation] = loc
			fmt.Printf("You took the %s.\n", itemName)
			return
		}
	}
	fmt.Println("Item not found here.")
}

func useItem(itemName string) {
	for i, item := range player.Inventory {
		if item == itemName {
			switch itemName {
			case "potion":
				player.Health += 20
				fmt.Println("You used a potion and gained 20 health.")
			case "sword":
				player.Strength += 5
				fmt.Println("You equipped the sword, strength increased by 5.")
			default:
				fmt.Printf("You used the %s, but nothing happened.\n", itemName)
			}
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			return
		}
	}
	fmt.Println("Item not in inventory.")
}

func fight() {
	loc := locations[currentLocation]
	if len(loc.Enemies) == 0 {
		fmt.Println("No enemies to fight here.")
		return
	}
	enemy := &loc.Enemies[0]
	fmt.Printf("You are fighting a %s!\n", enemy.Name)
	for player.Health > 0 && enemy.Health > 0 {
		// Player attacks
		damage := rand.Intn(player.Strength) + 1
		enemy.Health -= damage
		fmt.Printf("You hit the %s for %d damage.\n", enemy.Name, damage)
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			loc.Enemies = loc.Enemies[1:]
			locations[currentLocation] = loc
			player.Gold += rand.Intn(10) + 5
			fmt.Printf("You gained %d gold.\n", player.Gold)
			return
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

func showStatus() {
	fmt.Printf("Name: %s, Health: %d, Strength: %d, Gold: %d\n", player.Name, player.Health, player.Strength, player.Gold)
	fmt.Printf("Inventory: %s\n", strings.Join(player.Inventory, ", "))
}

func handleCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		return
	}
	switch parts[0] {
	case "go":
		if len(parts) < 2 {
			fmt.Println("Go where?")
		} else {
			move(parts[1])
		}
	case "take":
		if len(parts) < 2 {
			fmt.Println("Take what?")
		} else {
			takeItem(parts[1])
		}
	case "use":
		if len(parts) < 2 {
			fmt.Println("Use what?")
		} else {
			useItem(parts[1])
		}
	case "fight":
		fight()
	case "status":
		showStatus()
	case "look":
		displayLocation()
	case "quit":
		fmt.Println("Thanks for playing!")
		os.Exit(0)
	default:
		fmt.Println("Unknown command. Try: go, take, use, fight, status, look, quit")
	}
}

func main() {
	initGame()
	fmt.Println("Welcome to the Text Adventure Game!")
	fmt.Println("You find yourself in a forest. Explore, fight enemies, and find treasure.")
	fmt.Println("Commands: go [direction], take [item], use [item], fight, status, look, quit")
	reader := bufio.NewReader(os.Stdin)
	for {
		displayLocation()
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		handleCommand(input)
	}
}
