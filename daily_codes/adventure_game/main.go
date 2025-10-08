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
	Gold      int
	Inventory []string
	Level     int
	Exp       int
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

func main() {
	fmt.Println("Welcome to the Adventure Game!")
	fmt.Println("You find yourself in a mysterious land...")
	
	player := createPlayer()
	world := createWorld()
	currentLocation := "forest"
	
	gameLoop(player, world, currentLocation)
}

func createPlayer() Player {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	
	return Player{
		Name:      name,
		Health:    100,
		MaxHealth: 100,
		Strength:  10,
		Defense:   5,
		Gold:      50,
		Inventory: []string{"Health Potion", "Rusty Sword"},
		Level:     1,
		Exp:       0,
	}
}

func createWorld() map[string]Location {
	world := make(map[string]Location)
	
	world["forest"] = Location{
		Name:        "Mysterious Forest",
		Description: "A dense forest with tall trees and strange sounds.",
		Enemies: []Enemy{
			{Name: "Goblin", Health: 30, Strength: 8, Defense: 3, GoldDrop: 10, ExpDrop: 25},
			{Name: "Wolf", Health: 25, Strength: 12, Defense: 2, GoldDrop: 8, ExpDrop: 20},
		},
		Items:       []string{"Health Potion", "Iron Sword"},
		Connections: []string{"cave", "village"},
	}
	
	world["cave"] = Location{
		Name:        "Dark Cave",
		Description: "A dark and damp cave with glowing mushrooms.",
		Enemies: []Enemy{
			{Name: "Bat Swarm", Health: 40, Strength: 6, Defense: 1, GoldDrop: 5, ExpDrop: 15},
			{Name: "Cave Troll", Health: 80, Strength: 18, Defense: 8, GoldDrop: 30, ExpDrop: 50},
		},
		Items:       []string{"Gold Coin", "Magic Amulet"},
		Connections: []string{"forest"},
	}
	
	world["village"] = Location{
		Name:        "Peaceful Village",
		Description: "A small village with friendly inhabitants.",
		Enemies:     []Enemy{},
		Items:       []string{"Bread", "Water"},
		Connections: []string{"forest", "castle"},
	}
	
	world["castle"] = Location{
		Name:        "Ancient Castle",
		Description: "A massive castle with towering spires.",
		Enemies: []Enemy{
			{Name: "Knight", Health: 120, Strength: 22, Defense: 12, GoldDrop: 50, ExpDrop: 100},
			{Name: "Dragon", Health: 200, Strength: 30, Defense: 15, GoldDrop: 100, ExpDrop: 200},
		},
		Items:       []string{"Dragon Scale", "Royal Crown"},
		Connections: []string{"village"},
	}
	
	return world
}

func gameLoop(player Player, world map[string]Location, currentLocation string) {
	reader := bufio.NewReader(os.Stdin)
	
	for {
		location := world[currentLocation]
		
		fmt.Printf("\n=== %s ===\n", location.Name)
		fmt.Println(location.Description)
		
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1. Explore")
		fmt.Println("2. Check Inventory")
		fmt.Println("3. Check Stats")
		fmt.Println("4. Travel")
		fmt.Println("5. Quit")
		
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		
		switch choice {
		case "1":
			exploreLocation(player, location)
		case "2":
			checkInventory(player)
		case "3":
			checkStats(player)
		case "4":
			currentLocation = travel(player, world, currentLocation)
		case "5":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		
		// Check if player is dead
		if player.Health <= 0 {
			fmt.Println("\nYou have died! Game over.")
			return
		}
	}
}

func exploreLocation(player Player, location Location) {
	if len(location.Enemies) == 0 && len(location.Items) == 0 {
		fmt.Println("There's nothing interesting here.")
		return
	}
	
	// Random encounter
	if len(location.Enemies) > 0 {
		enemy := location.Enemies[rand.Intn(len(location.Enemies))]
		fmt.Printf("\nA wild %s appears!\n", enemy.Name)
		combat(player, enemy)
	}
	
	// Find items
	if len(location.Items) > 0 && rand.Float32() < 0.5 {
		item := location.Items[rand.Intn(len(location.Items))]
		fmt.Printf("\nYou found a %s!\n", item)
		player.Inventory = append(player.Inventory, item)
	}
}

func combat(player Player, enemy Enemy) {
	reader := bufio.NewReader(os.Stdin)
	
	for player.Health > 0 && enemy.Health > 0 {
		fmt.Printf("\nYour Health: %d/%d\n", player.Health, player.MaxHealth)
		fmt.Printf("%s's Health: %d\n", enemy.Name, enemy.Health)
		
		fmt.Println("\n1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Run")
		
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)
		
		switch choice {
		case "1":
			playerAttack := player.Strength + rand.Intn(6)
			enemyAttack := enemy.Strength + rand.Intn(6)
			
			playerDamage := playerAttack - enemy.Defense
			if playerDamage < 1 {
				playerDamage = 1
			}
			
			enemyDamage := enemyAttack - player.Defense
			if enemyDamage < 1 {
				enemyDamage = 1
			}
			
			enemy.Health -= playerDamage
			player.Health -= enemyDamage
			
			fmt.Printf("You deal %d damage to the %s!\n", playerDamage, enemy.Name)
			fmt.Printf("The %s deals %d damage to you!\n", enemy.Name, enemyDamage)
			
		case "2":
			useItem(player)
		case "3":
			if rand.Float32() < 0.7 {
				fmt.Println("You successfully ran away!")
				return
			} else {
				fmt.Println("You failed to run away!")
			}
		default:
			fmt.Println("Invalid choice. You hesitate and lose your turn!")
		}
	}
	
	if enemy.Health <= 0 {
		fmt.Printf("\nYou defeated the %s!\n", enemy.Name)
		player.Gold += enemy.GoldDrop
		player.Exp += enemy.ExpDrop
		fmt.Printf("You gained %d gold and %d experience!\n", enemy.GoldDrop, enemy.ExpDrop)
		
		// Check for level up
		if player.Exp >= player.Level*100 {
			player.Level++
			player.MaxHealth += 20
			player.Health = player.MaxHealth
			player.Strength += 3
			player.Defense += 2
			fmt.Printf("\nLevel Up! You are now level %d!\n", player.Level)
		}
	}
}

func useItem(player Player) {
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty!")
		return
	}
	
	fmt.Println("\nYour Inventory:")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of the item to use (or 0 to cancel): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	
	index, err := strconv.Atoi(choice)
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
		fmt.Printf("You used a Health Potion and recovered %d health!\n", healAmount)
		// Remove item from inventory
		player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
	default:
		fmt.Printf("You can't use the %s right now.\n", item)
	}
}

func checkInventory(player Player) {
	fmt.Println("\n=== Inventory ===")
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
	} else {
		for i, item := range player.Inventory {
			fmt.Printf("%d. %s\n", i+1, item)
		}
	}
	fmt.Printf("Gold: %d\n", player.Gold)
}

func checkStats(player Player) {
	fmt.Println("\n=== Player Stats ===")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Strength: %d\n", player.Strength)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Experience: %d/%d\n", player.Exp, player.Level*100)
}

func travel(player Player, world map[string]Location, currentLocation string) string {
	location := world[currentLocation]
	
	if len(location.Connections) == 0 {
		fmt.Println("There are no other locations to travel to from here.")
		return currentLocation
	}
	
	fmt.Println("\nAvailable locations:")
	for i, conn := range location.Connections {
		fmt.Printf("%d. %s\n", i+1, world[conn].Name)
	}
	
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of the location to travel to (or 0 to cancel): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	
	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(location.Connections) {
		fmt.Println("Invalid choice.")
		return currentLocation
	}
	
	newLocation := location.Connections[index-1]
	fmt.Printf("You travel to %s.\n", world[newLocation].Name)
	return newLocation
}

func init() {
	rand.Seed(time.Now().UnixNano())
}