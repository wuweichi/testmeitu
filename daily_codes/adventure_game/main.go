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
	Name        string
	Description string
	Effect      func(*Character)
}

var player Character
var enemies = []Enemy{
	{"Goblin", 30, 5, 2, 10, 5},
	{"Orc", 50, 8, 4, 20, 10},
	{"Dragon", 100, 15, 10, 50, 25},
}

var items = []Item{
	{"Health Potion", "Restores 20 health", func(c *Character) {
		c.Health = min(c.Health+20, c.MaxHealth)
		fmt.Println("You restored 20 health!")
	}},
	{"Strength Potion", "Increases attack by 5", func(c *Character) {
		c.Attack += 5
		fmt.Println("Your attack increased by 5!")
	}},
}

func main() {
	rand.Seed(time.Now().UnixNano())
	initializeGame()
	gameLoop()
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
		Gold:      50,
		Inventory: []string{"Health Potion"},
	}
	
	fmt.Printf("Welcome, %s! Your adventure begins now.\n", player.Name)
}

func gameLoop() {
	for {
		displayStatus()
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Explore")
		fmt.Println("2. Check Inventory")
		fmt.Println("3. Rest")
		fmt.Println("4. Quit")
		
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		switch input {
		case "1":
			explore()
		case "2":
			checkInventory()
		case "3":
			rest()
		case "4":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		
		if player.Health <= 0 {
			fmt.Println("You have been defeated. Game over!")
			return
		}
	}
}

func displayStatus() {
	fmt.Printf("\n--- %s's Status ---\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Exp: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Attack: %d, Defense: %d\n", player.Attack, player.Defense)
}

func explore() {
	fmt.Println("You venture into the wilderness...")
	
	// Random encounter
	encounter := rand.Intn(3)
	if encounter == 0 {
		fmt.Println("You found a treasure chest!")
		foundGold := rand.Intn(20) + 10
		player.Gold += foundGold
		fmt.Printf("You found %d gold!\n", foundGold)
	} else if encounter == 1 {
		fmt.Println("You discovered a hidden item!")
		itemIndex := rand.Intn(len(items))
		player.Inventory = append(player.Inventory, items[itemIndex].Name)
		fmt.Printf("You found a %s!\n", items[itemIndex].Name)
	} else {
		battle()
	}
}

func battle() {
	enemy := enemies[rand.Intn(len(enemies))]
	fmt.Printf("A wild %s appears!\n", enemy.Name)
	
	for player.Health > 0 && enemy.Health > 0 {
		fmt.Printf("\n%s's Health: %d\n", enemy.Name, enemy.Health)
		fmt.Println("Choose your action:")
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Flee")
		
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		switch input {
		case "1":
			playerAttack(&enemy)
		case "2":
			useItemInBattle()
		case "3":
			if rand.Intn(2) == 0 {
				fmt.Println("You successfully fled!")
				return
			} else {
				fmt.Println("You failed to flee!")
				enemyAttack(&enemy)
			}
		default:
			fmt.Println("Invalid choice. The enemy attacks!")
			enemyAttack(&enemy)
		}
		
		if enemy.Health > 0 {
			enemyAttack(&enemy)
		}
	}
	
	if enemy.Health <= 0 {
		fmt.Printf("You defeated the %s!\n", enemy.Name)
		player.Exp += enemy.Exp
		player.Gold += enemy.Gold
		fmt.Printf("Gained %d exp and %d gold!\n", enemy.Exp, enemy.Gold)
		checkLevelUp()
	}
}

func playerAttack(enemy *Enemy) {
	damage := max(player.Attack-enemy.Defense, 1)
	enemy.Health -= damage
	fmt.Printf("You dealt %d damage to the %s!\n", damage, enemy.Name)
}

func enemyAttack(enemy *Enemy) {
	damage := max(enemy.Attack-player.Defense, 1)
	player.Health -= damage
	fmt.Printf("The %s dealt %d damage to you!\n", enemy.Name, damage)
}

func useItemInBattle() {
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty!")
		return
	}
	
	fmt.Println("Your inventory:")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(player.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}
	
	itemName := player.Inventory[index-1]
	for _, item := range items {
		if item.Name == itemName {
			item.Effect(&player)
			player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
			return
		}
	}
	
	fmt.Println("Item not found.")
}

func checkInventory() {
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	
	fmt.Println("Your inventory:")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	
	fmt.Println("\nDo you want to use an item? (y/n)")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	if strings.ToLower(input) == "y" {
		useItem()
	}
}

func useItem() {
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty!")
		return
	}
	
	fmt.Println("Choose an item to use:")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(player.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}
	
	itemName := player.Inventory[index-1]
	for _, item := range items {
		if item.Name == itemName {
			item.Effect(&player)
			player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
			return
		}
	}
	
	fmt.Println("Item not found.")
}

func rest() {
	fmt.Println("You take a rest and recover some health.")
	player.Health = min(player.Health+30, player.MaxHealth)
	fmt.Printf("Health restored to %d/%d\n", player.Health, player.MaxHealth)
}

func checkLevelUp() {
	for player.Exp >= player.Level*100 {
		player.Exp -= player.Level * 100
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("Level up! You are now level %d!\n", player.Level)
		fmt.Println("Your stats have improved!")
	}
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
