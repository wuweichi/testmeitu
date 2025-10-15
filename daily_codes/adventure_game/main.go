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
	Gold      int
	Inventory []string
	Level     int
	Exp       int
}

type Enemy struct {
	Name    string
	Health  int
	Attack  int
	Defense int
	Gold    int
	Exp     int
}

type Item struct {
	Name        string
	Description string
	Effect      string
	Value       int
}

var player Player
var enemies = []Enemy{
	{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Gold: 10, Exp: 20},
	{Name: "Orc", Health: 50, Attack: 8, Defense: 4, Gold: 20, Exp: 40},
	{Name: "Dragon", Health: 100, Attack: 15, Defense: 8, Gold: 100, Exp: 100},
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
	player.Gold = 50
	player.Inventory = []string{}
	player.Level = 1
	player.Exp = 0
	fmt.Printf("Welcome, %s! You start with %d health, %d attack, %d defense, and %d gold.\n", player.Name, player.Health, player.Attack, player.Defense, player.Gold)
}

func gameLoop() {
	for {
		if player.Health <= 0 {
			fmt.Println("You have died. Game over!")
			return
		}
		showMenu()
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			explore()
		case "2":
			fight()
		case "3":
			useItem()
		case "4":
			checkStatus()
		case "5":
			shop()
		case "6":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid option. Please choose a number between 1 and 6.")
		}
	}
}

func showMenu() {
	fmt.Println("\n=== Main Menu ===")
	fmt.Println("1. Explore")
	fmt.Println("2. Fight")
	fmt.Println("3. Use Item")
	fmt.Println("4. Check Status")
	fmt.Println("5. Shop")
	fmt.Println("6. Quit")
	fmt.Print("Choose an option: ")
}

func explore() {
	fmt.Println("You venture into the wilderness...")
	events := []string{
		"You find a hidden treasure chest!",
		"You discover a peaceful village.",
		"You stumble upon an ancient ruin.",
		"You get lost in a dense forest.",
		"You find a healing spring.",
	}
	event := events[rand.Intn(len(events))]
	fmt.Println(event)
	switch event {
	case "You find a hidden treasure chest!":
		goldFound := rand.Intn(50) + 10
		player.Gold += goldFound
		fmt.Printf("You found %d gold! You now have %d gold.\n", goldFound, player.Gold)
	case "You discover a peaceful village.":
		fmt.Println("The villagers welcome you and offer rest.")
		player.Health = player.MaxHealth
		fmt.Println("Your health has been fully restored!")
	case "You stumble upon an ancient ruin.":
		itemFound := items[rand.Intn(len(items))]
		player.Inventory = append(player.Inventory, itemFound.Name)
		fmt.Printf("You found a %s! Added to inventory.\n", itemFound.Name)
	case "You get lost in a dense forest.":
		damage := rand.Intn(10) + 5
		player.Health -= damage
		fmt.Printf("You take %d damage from thorns and wild animals. Health: %d/%d\n", damage, player.Health, player.MaxHealth)
	case "You find a healing spring.":
		healAmount := rand.Intn(20) + 10
		player.Health += healAmount
		if player.Health > player.MaxHealth {
			player.Health = player.MaxHealth
		}
		fmt.Printf("You drink from the spring and heal %d health. Health: %d/%d\n", healAmount, player.Health, player.MaxHealth)
	}
}

func fight() {
	if len(enemies) == 0 {
		fmt.Println("No enemies available to fight!")
		return
	}
	enemy := enemies[rand.Intn(len(enemies))]
	fmt.Printf("A wild %s appears!\n", enemy.Name)
	battle(enemy)
}

func battle(enemy Enemy) {
	for {
		if player.Health <= 0 {
			fmt.Println("You have been defeated!")
			return
		}
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			player.Gold += enemy.Gold
			player.Exp += enemy.Exp
			fmt.Printf("You gained %d gold and %d experience. Total gold: %d, Exp: %d\n", enemy.Gold, enemy.Exp, player.Gold, player.Exp)
			checkLevelUp()
			return
		}
		fmt.Printf("\nYour Health: %d/%d, %s's Health: %d\n", player.Health, player.MaxHealth, enemy.Name, enemy.Health)
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Flee")
		fmt.Print("Choose an action: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			playerAttack := player.Attack + rand.Intn(5)
			enemyDamage := playerAttack - enemy.Defense
			if enemyDamage < 0 {
				enemyDamage = 0
			}
			enemy.Health -= enemyDamage
			fmt.Printf("You attack the %s for %d damage!\n", enemy.Name, enemyDamage)
			if enemy.Health > 0 {
				enemyAttack := enemy.Attack + rand.Intn(3)
				playerDamage := enemyAttack - player.Defense
				if playerDamage < 0 {
					playerDamage = 0
				}
				player.Health -= playerDamage
				fmt.Printf("The %s attacks you for %d damage!\n", enemy.Name, playerDamage)
			}
		case "2":
			useItemInBattle()
		case "3":
			if rand.Intn(2) == 0 {
				fmt.Println("You successfully fled from the battle!")
				return
			} else {
				fmt.Println("You failed to flee!")
				enemyAttack := enemy.Attack + rand.Intn(3)
				playerDamage := enemyAttack - player.Defense
				if playerDamage < 0 {
					playerDamage = 0
				}
				player.Health -= playerDamage
				fmt.Printf("The %s attacks you for %d damage as you try to flee!\n", enemy.Name, playerDamage)
			}
		default:
			fmt.Println("Invalid action. Please choose 1, 2, or 3.")
		}
	}
}

func useItemInBattle() {
	if len(player.Inventory) == 0 {
		fmt.Println("You have no items to use!")
		return
	}
	fmt.Println("Your inventory:")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Print("Choose an item to use (or 0 to cancel): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 0 || index > len(player.Inventory) {
		fmt.Println("Invalid selection.")
		return
	}
	if index == 0 {
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
	applyItemEffect(item)
	player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
	fmt.Printf("Used %s.\n", item.Name)
}

func useItem() {
	if len(player.Inventory) == 0 {
		fmt.Println("You have no items to use!")
		return
	}
	fmt.Println("Your inventory:")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Print("Choose an item to use (or 0 to cancel): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 0 || index > len(player.Inventory) {
		fmt.Println("Invalid selection.")
		return
	}
	if index == 0 {
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
	applyItemEffect(item)
	player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
	fmt.Printf("Used %s.\n", item.Name)
}

func applyItemEffect(item Item) {
	switch item.Effect {
	case "heal":
		player.Health += item.Value
		if player.Health > player.MaxHealth {
			player.Health = player.MaxHealth
		}
		fmt.Printf("Healed for %d health. Health: %d/%d\n", item.Value, player.Health, player.MaxHealth)
	case "buff_attack":
		player.Attack += item.Value
		fmt.Printf("Attack increased by %d. Attack: %d\n", item.Value, player.Attack)
	case "buff_defense":
		player.Defense += item.Value
		fmt.Printf("Defense increased by %d. Defense: %d\n", item.Value, player.Defense)
	}
}

func checkStatus() {
	fmt.Printf("\n=== Status ===\n")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Experience: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Inventory: %v\n", player.Inventory)
}

func shop() {
	fmt.Println("Welcome to the shop!")
	for {
		fmt.Println("\nAvailable items:")
		for i, item := range items {
			fmt.Printf("%d. %s - %s (Cost: %d gold)\n", i+1, item.Name, item.Description, item.Value*2)
		}
		fmt.Printf("Your gold: %d\n", player.Gold)
		fmt.Println("0. Leave shop")
		fmt.Print("Choose an item to buy: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		index, err := strconv.Atoi(input)
		if err != nil || index < 0 || index > len(items) {
			fmt.Println("Invalid selection.")
			continue
		}
		if index == 0 {
			fmt.Println("Thanks for visiting the shop!")
			return
		}
		item := items[index-1]
		cost := item.Value * 2
		if player.Gold >= cost {
			player.Gold -= cost
			player.Inventory = append(player.Inventory, item.Name)
			fmt.Printf("You bought a %s for %d gold. Remaining gold: %d\n", item.Name, cost, player.Gold)
		} else {
			fmt.Println("Not enough gold!")
		}
	}
}

func checkLevelUp() {
	expNeeded := player.Level * 100
	if player.Exp >= expNeeded {
		player.Level++
		player.Exp -= expNeeded
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("Level up! You are now level %d! Health, attack, and defense increased.\n", player.Level)
	}
}
