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
	Name      string
	Health    int
	Attack    int
	Defense   int
	ExpReward int
	GoldReward int
}

type Item struct {
	Name        string
	Description string
	Effect      func(*Character)
}

var player Character
var enemies = []Enemy{
	{"Goblin", 20, 5, 2, 10, 5},
	{"Orc", 30, 8, 4, 20, 10},
	{"Dragon", 50, 15, 10, 50, 25},
}

var items = []Item{
	{"Health Potion", "Restores 20 health.", func(c *Character) { c.Health = min(c.Health+20, c.MaxHealth) }},
	{"Attack Boost", "Increases attack by 5 for this battle.", func(c *Character) { c.Attack += 5 }},
	{"Defense Boost", "Increases defense by 5 for this battle.", func(c *Character) { c.Defense += 5 }},
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func initializeGame() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your character's name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	player = Character{
		Name:      name,
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
		Gold:      0,
		Inventory: []string{},
	}
	fmt.Printf("Welcome, %s! Your adventure begins.\n", player.Name)
}

func displayStatus() {
	fmt.Printf("Name: %s, Level: %d, Health: %d/%d, Attack: %d, Defense: %d, Exp: %d, Gold: %d\n",
		player.Name, player.Level, player.Health, player.MaxHealth, player.Attack, player.Defense, player.Exp, player.Gold)
	fmt.Println("Inventory:", player.Inventory)
}

func battle() {
	rand.Seed(time.Now().UnixNano())
	enemy := enemies[rand.Intn(len(enemies))]
	fmt.Printf("A wild %s appears!\n", enemy.Name)
	for player.Health > 0 && enemy.Health > 0 {
		fmt.Printf("Your health: %d, %s's health: %d\n", player.Health, enemy.Name, enemy.Health)
		fmt.Print("Choose action: (1) Attack, (2) Use Item, (3) Flee: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			damage := max(player.Attack-enemy.Defense, 1)
			enemy.Health -= damage
			fmt.Printf("You deal %d damage to %s.\n", damage, enemy.Name)
		case "2":
			if len(player.Inventory) == 0 {
				fmt.Println("No items in inventory!")
				continue
			}
			fmt.Println("Available items:")
			for i, item := range player.Inventory {
				fmt.Printf("%d: %s\n", i+1, item)
			}
			fmt.Print("Select item to use: ")
			choiceStr, _ := reader.ReadString('\n')
			choiceStr = strings.TrimSpace(choiceStr)
			choice, err := strconv.Atoi(choiceStr)
			if err != nil || choice < 1 || choice > len(player.Inventory) {
				fmt.Println("Invalid choice!")
				continue
			}
			itemName := player.Inventory[choice-1]
			for _, item := range items {
				if item.Name == itemName {
					item.Effect(&player)
					fmt.Printf("Used %s.\n", item.Name)
					player.Inventory = append(player.Inventory[:choice-1], player.Inventory[choice:]...)
					break
				}
			}
		case "3":
			if rand.Intn(2) == 0 {
				fmt.Println("You fled successfully!")
				return
			} else {
				fmt.Println("Failed to flee!")
			}
		default:
			fmt.Println("Invalid action!")
			continue
		}
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			player.Exp += enemy.ExpReward
			player.Gold += enemy.GoldReward
			fmt.Printf("Gained %d EXP and %d gold.\n", enemy.ExpReward, enemy.GoldReward)
			checkLevelUp()
			return
		}
		enemyDamage := max(enemy.Attack-player.Defense, 1)
		player.Health -= enemyDamage
		fmt.Printf("%s deals %d damage to you.\n", enemy.Name, enemyDamage)
	}
	if player.Health <= 0 {
		fmt.Println("You have been defeated! Game over.")
		os.Exit(0)
	}
}

func checkLevelUp() {
	expNeeded := player.Level * 100
	if player.Exp >= expNeeded {
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 5
		player.Defense += 3
		player.Exp = 0
		fmt.Printf("Level up! You are now level %d. Stats increased.\n", player.Level)
	}
}

func shop() {
	fmt.Println("Welcome to the shop! Available items:")
	for i, item := range items {
		fmt.Printf("%d: %s - 10 gold\n", i+1, item.Name)
	}
	fmt.Print("Enter item number to buy (or 0 to exit): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 0 || choice > len(items) {
		fmt.Println("Invalid choice!")
		return
	}
	if choice == 0 {
		return
	}
	if player.Gold < 10 {
		fmt.Println("Not enough gold!")
		return
	}
	player.Gold -= 10
	item := items[choice-1]
	player.Inventory = append(player.Inventory, item.Name)
	fmt.Printf("Bought %s. Added to inventory.\n", item.Name)
}

func saveGame() {
	file, err := os.Create("savegame.txt")
	if err != nil {
		fmt.Println("Error saving game:", err)
		return
	}
	defer file.Close()
	fmt.Fprintf(file, "%s\n%d\n%d\n%d\n%d\n%d\n%d\n%d\n",
		player.Name, player.Health, player.MaxHealth, player.Attack, player.Defense, player.Level, player.Exp, player.Gold)
	for _, item := range player.Inventory {
		fmt.Fprintln(file, item)
	}
	fmt.Println("Game saved successfully!")
}

func loadGame() {
	file, err := os.Open("savegame.txt")
	if err != nil {
		fmt.Println("No save game found.")
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if scanner.Scan() { player.Name = scanner.Text() }
	if scanner.Scan() { player.Health, _ = strconv.Atoi(scanner.Text()) }
	if scanner.Scan() { player.MaxHealth, _ = strconv.Atoi(scanner.Text()) }
	if scanner.Scan() { player.Attack, _ = strconv.Atoi(scanner.Text()) }
	if scanner.Scan() { player.Defense, _ = strconv.Atoi(scanner.Text()) }
	if scanner.Scan() { player.Level, _ = strconv.Atoi(scanner.Text()) }
	if scanner.Scan() { player.Exp, _ = strconv.Atoi(scanner.Text()) }
	if scanner.Scan() { player.Gold, _ = strconv.Atoi(scanner.Text()) }
	player.Inventory = nil
	for scanner.Scan() {
		player.Inventory = append(player.Inventory, scanner.Text())
	}
	fmt.Println("Game loaded successfully!")
}

func main() {
	fmt.Println("Starting Complex Simulation Game...")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Start New Game")
		fmt.Println("2. Load Game")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			initializeGame()
			gameLoop()
		case "2":
			loadGame()
			gameLoop()
		case "3":
			fmt.Println("Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid option!")
		}
	}
}

func gameLoop() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nGame Menu:")
		fmt.Println("1. Display Status")
		fmt.Println("2. Battle")
		fmt.Println("3. Visit Shop")
		fmt.Println("4. Save Game")
		fmt.Println("5. Return to Main Menu")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			displayStatus()
		case "2":
			battle()
		case "3":
			shop()
		case "4":
			saveGame()
		case "5":
			return
		default:
			fmt.Println("Invalid option!")
		}
	}
}
