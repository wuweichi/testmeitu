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
	Name   string
	Effect string
	Value  int
}

var player Character
var enemies = []Enemy{
	{"Goblin", 30, 5, 2, 10, 5},
	{"Orc", 50, 8, 4, 20, 10},
	{"Dragon", 100, 15, 10, 50, 30},
	{"Slime", 20, 3, 1, 5, 2},
	{"Skeleton", 40, 7, 3, 15, 8},
}

var items = []Item{
	{"Health Potion", "Restores 30 health", 30},
	{"Attack Boost", "Increases attack by 5", 5},
	{"Defense Boost", "Increases defense by 5", 5},
	{"Gold Pouch", "Adds 50 gold", 50},
}

func initGame() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Adventure Game Simulator!")
	fmt.Print("Enter your character name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	if name == "" {
		name = "Hero"
	}
	player = Character{
		Name:      name,
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
		Gold:      50,
		Inventory: []string{"Health Potion", "Health Potion"},
	}
	fmt.Printf("\nCharacter created: %s\n", player.Name)
	fmt.Println("Starting stats: Health 100, Attack 10, Defense 5, Level 1, Gold 50")
	fmt.Println("You have 2 Health Potions in your inventory.\n")
}

func showStatus() {
	fmt.Println("\n--- Status ---")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d, Defense: %d\n", player.Attack, player.Defense)
	fmt.Printf("Level: %d, Exp: %d/100\n", player.Level, player.Exp)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Inventory: %v\n", player.Inventory)
	fmt.Println("---------------")
}

func battle() {
	if len(enemies) == 0 {
		fmt.Println("No enemies left to fight!")
		return
	}
	rand.Seed(time.Now().UnixNano())
	enemy := enemies[rand.Intn(len(enemies))]
	fmt.Printf("\nA wild %s appears!\n", enemy.Name)
	fmt.Printf("Enemy stats: Health %d, Attack %d, Defense %d\n", enemy.Health, enemy.Attack, enemy.Defense)

	for player.Health > 0 && enemy.Health > 0 {
		fmt.Println("\n--- Battle Menu ---")
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Flee")
		fmt.Print("Choose an option: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			playerDamage := player.Attack - enemy.Defense
			if playerDamage < 1 {
				playerDamage = 1
			}
			enemy.Health -= playerDamage
			fmt.Printf("You attack the %s for %d damage.\n", enemy.Name, playerDamage)
			if enemy.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", enemy.Name)
				player.Exp += enemy.Exp
				player.Gold += enemy.Gold
				fmt.Printf("Gained %d exp and %d gold.\n", enemy.Exp, enemy.Gold)
				checkLevelUp()
				break
			}
			enemyDamage := enemy.Attack - player.Defense
			if enemyDamage < 1 {
				enemyDamage = 1
			}
			player.Health -= enemyDamage
			fmt.Printf("%s attacks you for %d damage.\n", enemy.Name, enemyDamage)
			if player.Health <= 0 {
				fmt.Println("You have been defeated! Game over.")
				os.Exit(0)
			}
		case "2":
			useItem()
		case "3":
			fmt.Println("You fled from the battle.")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func useItem() {
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty!")
		return
	}
	fmt.Println("\n--- Inventory ---")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Print("Choose an item to use (or 0 to cancel): ")
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
	for _, item := range items {
		if item.Name == itemName {
			switch item.Effect {
			case "Restores 30 health":
				player.Health += item.Value
				if player.Health > player.MaxHealth {
					player.Health = player.MaxHealth
				}
				fmt.Printf("Used %s. Health restored to %d.\n", item.Name, player.Health)
			case "Increases attack by 5":
				player.Attack += item.Value
				fmt.Printf("Used %s. Attack increased to %d.\n", item.Name, player.Attack)
			case "Increases defense by 5":
				player.Defense += item.Value
				fmt.Printf("Used %s. Defense increased to %d.\n", item.Name, player.Defense)
			case "Adds 50 gold":
				player.Gold += item.Value
				fmt.Printf("Used %s. Gold increased to %d.\n", item.Name, player.Gold)
			}
			player.Inventory = append(player.Inventory[:choice-1], player.Inventory[choice:]...)
			return
		}
	}
	fmt.Println("Item not found.")
}

func checkLevelUp() {
	for player.Exp >= 100 {
		player.Level++
		player.Exp -= 100
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("\nLevel up! You are now level %d.\n", player.Level)
		fmt.Println("Stats increased: Health +20, Attack +3, Defense +2")
	}
}

func shop() {
	fmt.Println("\n--- Shop ---")
	fmt.Println("Welcome to the shop! Here you can buy items.")
	fmt.Printf("Your gold: %d\n", player.Gold)
	for i, item := range items {
		fmt.Printf("%d. %s - Effect: %s - Price: %d gold\n", i+1, item.Name, item.Effect, item.Value*2)
	}
	fmt.Println("0. Exit shop")
	fmt.Print("Choose an item to buy: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 0 || choice > len(items) {
		fmt.Println("Invalid choice.")
		return
	}
	if choice == 0 {
		return
	}
	item := items[choice-1]
	price := item.Value * 2
	if player.Gold >= price {
		player.Gold -= price
		player.Inventory = append(player.Inventory, item.Name)
		fmt.Printf("You bought a %s for %d gold.\n", item.Name, price)
	} else {
		fmt.Println("Not enough gold!")
	}
}

func explore() {
	rand.Seed(time.Now().UnixNano())
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Println("\nYou explore the forest and find nothing.")
	case 1:
		goldFound := rand.Intn(20) + 10
		player.Gold += goldFound
		fmt.Printf("\nYou found %d gold while exploring!\n", goldFound)
	case 2:
		itemFound := items[rand.Intn(len(items))]
		player.Inventory = append(player.Inventory, itemFound.Name)
		fmt.Printf("\nYou found a %s!\n", itemFound.Name)
	}
}

func main() {
	initGame()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Show Status")
		fmt.Println("2. Battle")
		fmt.Println("3. Use Item")
		fmt.Println("4. Shop")
		fmt.Println("5. Explore")
		fmt.Println("6. Exit Game")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			showStatus()
		case "2":
			battle()
		case "3":
			useItem()
		case "4":
			shop()
		case "5":
			explore()
		case "6":
			fmt.Println("Thanks for playing! Goodbye.")
			os.Exit(0)
		default:
			fmt.Println("Invalid option. Please choose 1-6.")
		}
	}
}
