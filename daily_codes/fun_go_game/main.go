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
	Mana     int
	Strength int
	Agility  int
	Intellect int
	Inventory []string
	Gold     int
}

type Monster struct {
	Name     string
	Health   int
	Damage   int
	Defense  int
	Loot     []string
	GoldDrop int
}

type Item struct {
	Name        string
	Description string
	Value       int
	Type        string
}

type Location struct {
	Name        string
	Description string
	Monsters    []Monster
	Items       []Item
	Connections []string
}

var player Player
var locations map[string]Location
var currentLocation string

func initGame() {
	player = Player{
		Name:     "Hero",
		Health:   100,
		Mana:     50,
		Strength: 10,
		Agility:  8,
		Intellect: 6,
		Inventory: []string{"Health Potion", "Mana Potion"},
		Gold:     50,
	}

	locations = make(map[string]Location)

	locations["forest"] = Location{
		Name:        "Enchanted Forest",
		Description: "A mystical forest filled with ancient trees and magical creatures.",
		Monsters: []Monster{
			{Name: "Goblin", Health: 30, Damage: 5, Defense: 2, Loot: []string{"Goblin Ear", "Rusty Dagger"}, GoldDrop: 10},
			{Name: "Wolf", Health: 25, Damage: 8, Defense: 1, Loot: []string{"Wolf Pelt", "Sharp Fang"}, GoldDrop: 15},
		},
		Items: []Item{
			{Name: "Health Potion", Description: "Restores 20 health", Value: 10, Type: "Consumable"},
			{Name: "Mana Potion", Description: "Restores 15 mana", Value: 12, Type: "Consumable"},
		},
		Connections: []string{"village", "cave"},
	}

	locations["village"] = Location{
		Name:        "Peaceful Village",
		Description: "A small village with friendly inhabitants and a market.",
		Monsters:    []Monster{},
		Items: []Item{
			{Name: "Iron Sword", Description: "A basic sword for combat", Value: 50, Type: "Weapon"},
			{Name: "Leather Armor", Description: "Light armor for protection", Value: 40, Type: "Armor"},
		},
		Connections: []string{"forest"},
	}

	locations["cave"] = Location{
		Name:        "Dark Cave",
		Description: "A dark and dangerous cave with hidden treasures.",
		Monsters: []Monster{
			{Name: "Troll", Health: 80, Damage: 15, Defense: 5, Loot: []string{"Troll Hide", "Massive Club"}, GoldDrop: 50},
			{Name: "Bat Swarm", Health: 20, Damage: 3, Defense: 0, Loot: []string{"Bat Wing"}, GoldDrop: 5},
		},
		Items: []Item{
			{Name: "Treasure Chest", Description: "Contains valuable items", Value: 100, Type: "Container"},
			{Name: "Magic Ring", Description: "Enhances intellect", Value: 75, Type: "Accessory"},
		},
		Connections: []string{"forest"},
	}

	currentLocation = "forest"
}

func displayStatus() {
	fmt.Printf("\n=== Player Status ===\n")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Health: %d\n", player.Health)
	fmt.Printf("Mana: %d\n", player.Mana)
	fmt.Printf("Strength: %d\n", player.Strength)
	fmt.Printf("Agility: %d\n", player.Agility)
	fmt.Printf("Intellect: %d\n", player.Intellect)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Inventory: %v\n", player.Inventory)
	fmt.Printf("Current Location: %s\n", locations[currentLocation].Name)
}

func displayLocation() {
	loc := locations[currentLocation]
	fmt.Printf("\n=== %s ===\n", loc.Name)
	fmt.Printf("Description: %s\n", loc.Description)
	if len(loc.Monsters) > 0 {
		fmt.Printf("Monsters here: ")
		for _, m := range loc.Monsters {
			fmt.Printf("%s ", m.Name)
		}
		fmt.Println()
	}
	if len(loc.Items) > 0 {
		fmt.Printf("Items here: ")
		for _, item := range loc.Items {
			fmt.Printf("%s ", item.Name)
		}
		fmt.Println()
	}
	fmt.Printf("Connections: %v\n", loc.Connections)
}

func moveTo(newLocation string) {
	if !contains(locations[currentLocation].Connections, newLocation) {
		fmt.Println("You cannot go there from here.")
		return
	}
	currentLocation = newLocation
	fmt.Printf("You have moved to %s.\n", locations[currentLocation].Name)
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func fightMonster() {
	loc := locations[currentLocation]
	if len(loc.Monsters) == 0 {
		fmt.Println("There are no monsters here to fight.")
		return
	}

	monster := loc.Monsters[rand.Intn(len(loc.Monsters))]
	fmt.Printf("You encounter a %s!\n", monster.Name)

	for player.Health > 0 && monster.Health > 0 {
		fmt.Printf("\nYour Health: %d, %s's Health: %d\n", player.Health, monster.Name, monster.Health)
		fmt.Println("Choose an action: (1) Attack (2) Use Item (3) Flee")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			playerDamage := player.Strength + rand.Intn(5)
			actualDamage := playerDamage - monster.Defense
			if actualDamage < 0 {
				actualDamage = 0
			}
			monster.Health -= actualDamage
			fmt.Printf("You attack the %s for %d damage!\n", monster.Name, actualDamage)

			if monster.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", monster.Name)
				player.Gold += monster.GoldDrop
				if len(monster.Loot) > 0 {
					loot := monster.Loot[rand.Intn(len(monster.Loot))]
					player.Inventory = append(player.Inventory, loot)
					fmt.Printf("You found: %s\n", loot)
				}
				break
			}

			monsterDamage := monster.Damage + rand.Intn(3)
			player.Health -= monsterDamage
			fmt.Printf("The %s attacks you for %d damage!\n", monster.Name, monsterDamage)

			if player.Health <= 0 {
				fmt.Println("You have been defeated! Game Over.")
				return
			}

		case "2":
			useItem()
		case "3":
			if rand.Float32() < 0.5 {
				fmt.Println("You successfully fled from the battle!")
				return
			} else {
				fmt.Println("You failed to flee!")
				monsterDamage := monster.Damage + rand.Intn(3)
				player.Health -= monsterDamage
				fmt.Printf("The %s attacks you for %d damage!\n", monster.Name, monsterDamage)
				if player.Health <= 0 {
					fmt.Println("You have been defeated! Game Over.")
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
		fmt.Println("You have no items to use.")
		return
	}

	fmt.Println("Your inventory:")
	for i, item := range player.Inventory {
		fmt.Printf("%d: %s\n", i+1, item)
	}
	fmt.Println("Choose an item to use (number) or 0 to cancel:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 0 || choice > len(player.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}

	if choice == 0 {
		return
	}

	itemName := player.Inventory[choice-1]
	if itemName == "Health Potion" {
		player.Health += 20
		fmt.Println("You used a Health Potion and restored 20 health.")
	} else if itemName == "Mana Potion" {
		player.Mana += 15
		fmt.Println("You used a Mana Potion and restored 15 mana.")
	} else {
		fmt.Printf("You cannot use %s right now.\n", itemName)
		return
	}

	player.Inventory = append(player.Inventory[:choice-1], player.Inventory[choice:]...)
}

func pickUpItem() {
	loc := locations[currentLocation]
	if len(loc.Items) == 0 {
		fmt.Println("There are no items here to pick up.")
		return
	}

	item := loc.Items[rand.Intn(len(loc.Items))]
	player.Inventory = append(player.Inventory, item.Name)
	fmt.Printf("You picked up: %s\n", item.Name)
	loc.Items = append(loc.Items[:0], loc.Items[1:]...)
	locations[currentLocation] = loc
}

func shop() {
	if currentLocation != "village" {
		fmt.Println("There is no shop here.")
		return
	}

	fmt.Println("\n=== Village Shop ===")
	fmt.Println("Items for sale:")
	fmt.Println("1. Health Potion - 10 gold")
	fmt.Println("2. Mana Potion - 12 gold")
	fmt.Println("3. Iron Sword - 50 gold")
	fmt.Println("4. Leather Armor - 40 gold")
	fmt.Println("Enter the number to buy or 0 to exit:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 0 || choice > 4 {
		fmt.Println("Invalid choice.")
		return
	}

	if choice == 0 {
		return
	}

	var itemName string
	var cost int
	switch choice {
	case 1:
		itemName = "Health Potion"
		cost = 10
	case 2:
		itemName = "Mana Potion"
		cost = 12
	case 3:
		itemName = "Iron Sword"
		cost = 50
	case 4:
		itemName = "Leather Armor"
		cost = 40
	}

	if player.Gold < cost {
		fmt.Println("Not enough gold!")
		return
	}

	player.Gold -= cost
	player.Inventory = append(player.Inventory, itemName)
	fmt.Printf("You bought %s for %d gold.\n", itemName, cost)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initGame()

	fmt.Println("Welcome to the Fun Go Game!")
	fmt.Println("You are an adventurer in a fantasy world.")
	fmt.Println("Type 'help' for a list of commands.")

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("\n> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("status - Show player status")
			fmt.Println("look - Describe current location")
			fmt.Println("go [location] - Move to a connected location")
			fmt.Println("fight - Fight a monster")
			fmt.Println("pickup - Pick up an item")
			fmt.Println("use - Use an item from inventory")
			fmt.Println("shop - Buy items in the village")
			fmt.Println("quit - Exit the game")
		case "status":
			displayStatus()
		case "look":
			displayLocation()
		case "go forest":
			moveTo("forest")
		case "go village":
			moveTo("village")
		case "go cave":
			moveTo("cave")
		case "fight":
			fightMonster()
		case "pickup":
			pickUpItem()
		case "use":
			useItem()
		case "shop":
			shop()
		case "quit":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}
}
