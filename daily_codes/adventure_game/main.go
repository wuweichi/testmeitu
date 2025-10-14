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
	{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Exp: 10, Gold: 5},
	{Name: "Orc", Health: 50, Attack: 8, Defense: 4, Exp: 20, Gold: 10},
	{Name: "Dragon", Health: 100, Attack: 15, Defense: 10, Exp: 50, Gold: 50},
}

var items = []Item{
	{Name: "Health Potion", Description: "Restores 20 health", Effect: "heal", Value: 20},
	{Name: "Strength Potion", Description: "Increases attack by 5", Effect: "buff_attack", Value: 5},
	{Name: "Defense Potion", Description: "Increases defense by 5", Effect: "buff_defense", Value: 5},
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
	player.Name = strings.TrimSpace(name)
	player.Health = 100
	player.MaxHealth = 100
	player.Attack = 10
	player.Defense = 5
	player.Level = 1
	player.Exp = 0
	player.Gold = 0
	player.Inventory = []string{}
	fmt.Printf("Welcome, %s! You are a brave adventurer.\n", player.Name)
}

func gameLoop() {
	for {
		showStatus()
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1. Explore")
		fmt.Println("2. Check Inventory")
		fmt.Println("3. Rest")
		fmt.Println("4. Quit")
		fmt.Print("Enter your choice: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
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
			fmt.Println("You have been defeated! Game over.")
			return
		}
	}
}

func showStatus() {
	fmt.Printf("\n=== %s's Status ===\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Exp: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
}

func explore() {
	fmt.Println("You venture into the wilderness...")
	event := rand.Intn(3)
	switch event {
	case 0:
		findItem()
	case 1:
		findGold()
	case 2:
		encounterEnemy()
	}
}

func findItem() {
	item := items[rand.Intn(len(items))]
	fmt.Printf("You found a %s! %s\n", item.Name, item.Description)
	player.Inventory = append(player.Inventory, item.Name)
}

func findGold() {
	gold := rand.Intn(20) + 5
	player.Gold += gold
	fmt.Printf("You found %d gold!\n", gold)
}

func encounterEnemy() {
	enemy := enemies[rand.Intn(len(enemies))]
	fmt.Printf("A wild %s appears!\n", enemy.Name)
	battle(enemy)
}

func battle(enemy Enemy) {
	for enemy.Health > 0 && player.Health > 0 {
		fmt.Printf("\n%s: %d HP\n", enemy.Name, enemy.Health)
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Flee")
		fmt.Print("Enter your choice: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			playerAttack(&enemy)
		case "2":
			useItemInBattle()
		case "3":
			if rand.Float32() < 0.5 {
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
	}

	if enemy.Health <= 0 {
		fmt.Printf("You defeated the %s!\n", enemy.Name)
		player.Exp += enemy.Exp
		player.Gold += enemy.Gold
		fmt.Printf("Gained %d exp and %d gold.\n", enemy.Exp, enemy.Gold)
		checkLevelUp()
	}
}

func playerAttack(enemy *Enemy) {
	damage := player.Attack - enemy.Defense
	if damage < 1 {
		damage = 1
	}
	enemy.Health -= damage
	fmt.Printf("You attack the %s for %d damage!\n", enemy.Name, damage)
	if enemy.Health > 0 {
		enemyAttack(enemy)
	}
}

func enemyAttack(enemy *Enemy) {
	damage := enemy.Attack - player.Defense
	if damage < 1 {
		damage = 1
	}
	player.Health -= damage
	fmt.Printf("The %s attacks you for %d damage!\n", enemy.Name, damage)
}

func useItemInBattle() {
	if len(player.Inventory) == 0 {
		fmt.Println("You have no items!")
		return
	}

	fmt.Println("Your inventory:")
	for i, itemName := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, itemName)
	}
	fmt.Print("Enter the number of the item to use: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(input)
	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(player.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}

	itemName := player.Inventory[index-1]
	var item Item
	for _, it := range items {
		if it.Name == itemName {
			item = it
			break
		}
	}

	useItem(item)
	player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
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
	choice := strings.TrimSpace(input)

	if choice == "y" || choice == "Y" {
		useItemFromInventory()
	}
}

func useItemFromInventory() {
	fmt.Print("Enter the number of the item to use: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	choice := strings.TrimSpace(input)
	index, err := strconv.Atoi(choice)
	if err != nil || index < 1 || index > len(player.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}

	itemName := player.Inventory[index-1]
	var item Item
	for _, it := range items {
		if it.Name == itemName {
			item = it
			break
		}
	}

	useItem(item)
	player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
}

func useItem(item Item) {
	switch item.Effect {
	case "heal":
		player.Health += item.Value
		if player.Health > player.MaxHealth {
			player.Health = player.MaxHealth
		}
		fmt.Printf("You used %s and restored %d health.\n", item.Name, item.Value)
	case "buff_attack":
		player.Attack += item.Value
		fmt.Printf("You used %s and increased your attack by %d.\n", item.Name, item.Value)
	case "buff_defense":
		player.Defense += item.Value
		fmt.Printf("You used %s and increased your defense by %d.\n", item.Name, item.Value)
	}
}

func rest() {
	if player.Health == player.MaxHealth {
		fmt.Println("You are already at full health.")
		return
	}

	player.Health = player.MaxHealth
	fmt.Println("You rest and recover all your health.")
}

func checkLevelUp() {
	requiredExp := player.Level * 100
	if player.Exp >= requiredExp {
		player.Level++
		player.Exp -= requiredExp
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("Congratulations! You reached level %d!\n", player.Level)
		fmt.Println("Your stats have improved!")
	}
}
