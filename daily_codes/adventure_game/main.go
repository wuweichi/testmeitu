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
	Effect      string
	Value       int
}

var player Player
var enemies = []Enemy{
	{"Goblin", 20, 5, 2, 10, 5},
	{"Orc", 40, 8, 4, 20, 10},
	{"Dragon", 100, 15, 10, 50, 50},
}

var items = []Item{
	{"Health Potion", "Restores 20 health", "heal", 20},
	{"Strength Potion", "Increases attack by 5", "buff_attack", 5},
	{"Defense Potion", "Increases defense by 5", "buff_defense", 5},
}

func init() {
	rand.Seed(time.Now().UnixNano())
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
		Gold:      50,
		Inventory: []string{},
	}
}

func gameLoop() {
	for {
		displayStatus()
		action := getPlayerAction()
		switch action {
		case "1":
			explore()
		case "2":
			rest()
		case "3":
			shop()
		case "4":
			viewInventory()
		case "5":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid action. Please try again.")
		}
		checkLevelUp()
		if player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			return
		}
	}
}

func displayStatus() {
	fmt.Printf("\n=== %s's Status ===\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Exp: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Inventory: %d items\n", len(player.Inventory))
}

func getPlayerAction() string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("\nChoose an action:")
	fmt.Println("1. Explore")
	fmt.Println("2. Rest")
	fmt.Println("3. Shop")
	fmt.Println("4. View Inventory")
	fmt.Println("5. Quit")
	fmt.Print("Enter your choice: ")
	choice, _ := reader.ReadString('\n')
	return strings.TrimSpace(choice)
}

func explore() {
	fmt.Println("You venture into the wilderness...")
	if rand.Intn(100) < 70 {
		encounterEnemy()
	} else {
		findItem()
	}
}

func encounterEnemy() {
	enemy := enemies[rand.Intn(len(enemies))]
	fmt.Printf("A wild %s appears!\n", enemy.Name)
	battle(&enemy)
}

func battle(enemy *Enemy) {
	for player.Health > 0 && enemy.Health > 0 {
		fmt.Printf("\n%s: %d HP\n", enemy.Name, enemy.Health)
		fmt.Printf("%s: %d HP\n", player.Name, player.Health)
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Flee")
		fmt.Print("Choose action: ")
		choice, _ := reader.ReadString('\n')
		switch strings.TrimSpace(choice) {
		case "1":
			playerAttack(enemy)
		case "2":
			useItemInBattle()
		case "3":
			if rand.Intn(100) < 50 {
				fmt.Println("You successfully fled!")
				return
			} else {
				fmt.Println("You failed to flee!")
				enemyAttack(enemy)
			}
		default:
			fmt.Println("Invalid choice!")
		}
		if enemy.Health > 0 {
			enemyAttack(enemy)
		}
	}
	if player.Health > 0 {
		fmt.Printf("You defeated the %s!\n", enemy.Name)
		player.Exp += enemy.Exp
		player.Gold += enemy.Gold
		fmt.Printf("Gained %d exp and %d gold!\n", enemy.Exp, enemy.Gold)
	}
}

func playerAttack(enemy *Enemy) {
	damage := player.Attack - enemy.Defense
	if damage < 1 {
		damage = 1
	}
	enemy.Health -= damage
	fmt.Printf("You dealt %d damage to the %s!\n", damage, enemy.Name)
}

func enemyAttack(enemy *Enemy) {
	damage := enemy.Attack - player.Defense
	if damage < 1 {
		damage = 1
	}
	player.Health -= damage
	fmt.Printf("The %s dealt %d damage to you!\n", enemy.Name, damage)
}

func useItemInBattle() {
	if len(player.Inventory) == 0 {
		fmt.Println("You have no items!")
		return
	}
	fmt.Println("Your inventory:")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose item to use (0 to cancel): ")
	choice, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(choice))
	if err != nil || index < 1 || index > len(player.Inventory) {
		if index != 0 {
			fmt.Println("Invalid choice!")
		}
		return
	}
	itemName := player.Inventory[index-1]
	useItem(itemName)
	player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
}

func findItem() {
	item := items[rand.Intn(len(items))]
	fmt.Printf("You found a %s!\n", item.Name)
	player.Inventory = append(player.Inventory, item.Name)
}

func rest() {
	if player.Gold >= 10 {
		player.Health = player.MaxHealth
		player.Gold -= 10
		fmt.Println("You rest at the inn and recover all health. Cost: 10 gold")
	} else {
		fmt.Println("You don't have enough gold to rest! (Need 10 gold)")
	}
}

func shop() {
	fmt.Println("Welcome to the shop!")
	for {
		fmt.Printf("\nYour gold: %d\n", player.Gold)
		fmt.Println("Items for sale:")
		for i, item := range items {
			fmt.Printf("%d. %s - %s (%d gold)\n", i+1, item.Name, item.Description, item.Value)
		}
		fmt.Println("0. Leave shop")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Choose item to buy: ")
		choice, _ := reader.ReadString('\n')
		index, err := strconv.Atoi(strings.TrimSpace(choice))
		if err != nil || index < 0 || index > len(items) {
			fmt.Println("Invalid choice!")
			continue
		}
		if index == 0 {
			break
		}
		item := items[index-1]
		if player.Gold >= item.Value {
			player.Gold -= item.Value
			player.Inventory = append(player.Inventory, item.Name)
			fmt.Printf("You bought a %s!\n", item.Name)
		} else {
			fmt.Println("You don't have enough gold!")
		}
	}
}

func viewInventory() {
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty!")
		return
	}
	fmt.Println("Your inventory:")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Choose item to use (0 to cancel): ")
	choice, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(choice))
	if err != nil || index < 1 || index > len(player.Inventory) {
		if index != 0 {
			fmt.Println("Invalid choice!")
		}
		return
	}
	itemName := player.Inventory[index-1]
	useItem(itemName)
	player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
}

func useItem(itemName string) {
	for _, item := range items {
		if item.Name == itemName {
			switch item.Effect {
			case "heal":
				player.Health += item.Value
				if player.Health > player.MaxHealth {
					player.Health = player.MaxHealth
				}
				fmt.Printf("Used %s. Healed %d health.\n", item.Name, item.Value)
			case "buff_attack":
				player.Attack += item.Value
				fmt.Printf("Used %s. Attack increased by %d.\n", item.Name, item.Value)
			case "buff_defense":
				player.Defense += item.Value
				fmt.Printf("Used %s. Defense increased by %d.\n", item.Name, item.Value)
			}
			return
		}
	}
	fmt.Println("Item not found!")
}

func checkLevelUp() {
	for player.Exp >= player.Level*100 {
		player.Exp -= player.Level * 100
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("\n!!! Level Up !!! You are now level %d!\n", player.Level)
		fmt.Println("Max health, attack, and defense increased!")
	}
}
