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
	Effect      func(*Player)
}

var player Player
var enemies = []Enemy{
	{"Goblin", 30, 5, 2, 10, 5},
	{"Orc", 50, 8, 4, 20, 10},
	{"Dragon", 100, 15, 10, 50, 25},
}

var items = []Item{
	{"Health Potion", "Restores 20 health", func(p *Player) {
		p.Health = min(p.Health+20, p.MaxHealth)
		fmt.Println("You restored 20 health!")
	}},
	{"Strength Potion", "Increases attack by 5", func(p *Player) {
		p.Attack += 5
		fmt.Println("Your attack increased by 5!")
	}},
	{"Defense Potion", "Increases defense by 5", func(p *Player) {
		p.Defense += 5
		fmt.Println("Your defense increased by 5!")
	}},
}

func main() {
	fmt.Println("Welcome to the Adventure Game!")
	initializePlayer()
	gameLoop()
}

func initializePlayer() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	player = Player{
		Name:      strings.TrimSpace(name),
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
		Gold:      10,
		Inventory: []string{},
	}
}

func gameLoop() {
	for {
		if player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			return
		}
		fmt.Println("\n=== Main Menu ===")
		fmt.Println("1. Explore")
		fmt.Println("2. Check Status")
		fmt.Println("3. Use Item")
		fmt.Println("4. Shop")
		fmt.Println("5. Quit")
		fmt.Print("Choose an option: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			explore()
		case "2":
			checkStatus()
		case "3":
			useItem()
		case "4":
			shop()
		case "5":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func explore() {
	fmt.Println("\nYou venture into the wilderness...")
	rand.Seed(time.Now().UnixNano())
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Println("You found a treasure chest!")
		foundGold := rand.Intn(20) + 5
		player.Gold += foundGold
		fmt.Printf("You found %d gold!\n", foundGold)
	case 1:
		fmt.Println("You encountered an enemy!")
		enemyIndex := rand.Intn(len(enemies))
		enemy := enemies[enemyIndex]
		battle(&enemy)
	case 2:
		fmt.Println("You found a mysterious item!")
		itemIndex := rand.Intn(len(items))
		item := items[itemIndex]
		player.Inventory = append(player.Inventory, item.Name)
		fmt.Printf("You found a %s!\n", item.Name)
	}
}

func battle(enemy *Enemy) {
	fmt.Printf("A wild %s appears!\n", enemy.Name)
	for player.Health > 0 && enemy.Health > 0 {
		fmt.Println("\n=== Battle ===")
		fmt.Printf("Your Health: %d/%d\n", player.Health, player.MaxHealth)
		fmt.Printf("%s Health: %d\n", enemy.Name, enemy.Health)
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Flee")
		fmt.Print("Choose an action: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			playerAttack := max(player.Attack-enemy.Defense, 1)
			enemy.Health -= playerAttack
			fmt.Printf("You dealt %d damage to the %s!\n", playerAttack, enemy.Name)
			if enemy.Health > 0 {
				enemyAttack := max(enemy.Attack-player.Defense, 1)
				player.Health -= enemyAttack
				fmt.Printf("The %s dealt %d damage to you!\n", enemy.Name, enemyAttack)
			}
		case "2":
			useItem()
		case "3":
			fmt.Println("You fled from the battle!")
			return
		default:
			fmt.Println("Invalid action. Please try again.")
		}
	}
	if enemy.Health <= 0 {
		fmt.Printf("You defeated the %s!\n", enemy.Name)
		player.Exp += enemy.Exp
		player.Gold += enemy.Gold
		fmt.Printf("You gained %d experience and %d gold!\n", enemy.Exp, enemy.Gold)
		checkLevelUp()
	}
}

func checkLevelUp() {
	expNeeded := player.Level * 50
	if player.Exp >= expNeeded {
		player.Level++
		player.Exp -= expNeeded
		player.MaxHealth += 10
		player.Health = player.MaxHealth
		player.Attack += 2
		player.Defense += 1
		fmt.Printf("You leveled up to level %d!\n", player.Level)
		fmt.Println("Your stats have improved!")
	}
}

func checkStatus() {
	fmt.Println("\n=== Player Status ===")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Experience: %d/%d\n", player.Exp, player.Level*50)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Inventory: %v\n", player.Inventory)
}

func useItem() {
	if len(player.Inventory) == 0 {
		fmt.Println("You have no items to use.")
		return
	}
	fmt.Println("\n=== Inventory ===")
	for i, itemName := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, itemName)
	}
	fmt.Print("Choose an item to use (or 0 to cancel): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(player.Inventory) {
		if index == 0 {
			return
		}
		fmt.Println("Invalid selection.")
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

func shop() {
	fmt.Println("\n=== Shop ===")
	fmt.Println("Welcome to the shop!")
	fmt.Printf("You have %d gold.\n", player.Gold)
	for i, item := range items {
		fmt.Printf("%d. %s - 10 gold\n", i+1, item.Name)
	}
	fmt.Println("0. Leave")
	fmt.Print("Choose an item to buy: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(items) {
		if index == 0 {
			return
		}
		fmt.Println("Invalid selection.")
		return
	}
	if player.Gold < 10 {
		fmt.Println("Not enough gold!")
		return
	}
	item := items[index-1]
	player.Inventory = append(player.Inventory, item.Name)
	player.Gold -= 10
	fmt.Printf("You bought a %s!\n", item.Name)
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