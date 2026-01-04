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
	Skills    []Skill
}

type Skill struct {
	Name        string
	Description string
	Damage      int
	ManaCost    int
}

type Enemy struct {
	Name      string
	Health    int
	Attack    int
	Defense   int
	ExpReward int
	GoldReward int
	Drops     []string
}

type Item struct {
	Name        string
	Description string
	Effect      string
	Value       int
}

var player Character
var enemies = []Enemy{
	{"Goblin", 30, 5, 2, 10, 5, []string{"Rusty Dagger"}},
	{"Orc", 50, 8, 4, 20, 10, []string{"Orcish Axe"}},
	{"Dragon", 100, 15, 10, 50, 30, []string{"Dragon Scale"}},
}

var items = []Item{
	{"Health Potion", "Restores 20 health", "heal", 10},
	{"Mana Potion", "Restores 10 mana", "mana", 15},
	{"Iron Sword", "A basic sword", "attack", 25},
	{"Leather Armor", "Light armor", "defense", 20},
}

func initGame() {
	player = Character{
		Name:      "Hero",
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
		Gold:      50,
		Inventory: []string{"Health Potion", "Iron Sword"},
		Skills: []Skill{
			{"Slash", "A basic attack", 15, 0},
			{"Fireball", "A magical attack", 25, 10},
		},
	}
}

func displayStatus() {
	fmt.Printf("\n=== Status ===\n")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Exp: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Inventory: %v\n", player.Inventory)
	fmt.Printf("Skills: ")
	for _, skill := range player.Skills {
		fmt.Printf("%s ", skill.Name)
	}
	fmt.Println()
}

func explore() {
	fmt.Println("\nYou venture into the wilderness...")
	rand.Seed(time.Now().UnixNano())
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Println("You find a hidden treasure chest!")
		goldFound := rand.Intn(20) + 10
		player.Gold += goldFound
		fmt.Printf("You found %d gold!\n", goldFound)
	case 1:
		fmt.Println("You encounter a wild enemy!")
		enemy := enemies[rand.Intn(len(enemies))]
		battle(enemy)
	case 2:
		fmt.Println("You discover a peaceful village.")
		visitVillage()
	}
}

func battle(enemy Enemy) {
	fmt.Printf("\nA wild %s appears!\n", enemy.Name)
	for enemy.Health > 0 && player.Health > 0 {
		fmt.Printf("\nYour Health: %d/%d | %s's Health: %d\n", player.Health, player.MaxHealth, enemy.Name, enemy.Health)
		fmt.Println("Choose an action:")
		fmt.Println("1. Attack")
		fmt.Println("2. Use Skill")
		fmt.Println("3. Use Item")
		fmt.Println("4. Flee")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			damage := player.Attack - enemy.Defense
			if damage < 0 {
				damage = 0
			}
			enemy.Health -= damage
			fmt.Printf("You attack the %s for %d damage!\n", enemy.Name, damage)
		case "2":
			fmt.Println("Choose a skill:")
			for i, skill := range player.Skills {
				fmt.Printf("%d. %s (Damage: %d, Mana Cost: %d)\n", i+1, skill.Name, skill.Damage, skill.ManaCost)
			}
			skillInput, _ := reader.ReadString('\n')
			skillInput = strings.TrimSpace(skillInput)
			index, err := strconv.Atoi(skillInput)
			if err == nil && index > 0 && index <= len(player.Skills) {
				skill := player.Skills[index-1]
				damage := skill.Damage - enemy.Defense
				if damage < 0 {
					damage = 0
				}
				enemy.Health -= damage
				fmt.Printf("You use %s for %d damage!\n", skill.Name, damage)
			} else {
				fmt.Println("Invalid skill choice.")
			}
		case "3":
			fmt.Println("Choose an item:")
			for i, item := range player.Inventory {
				fmt.Printf("%d. %s\n", i+1, item)
			}
			itemInput, _ := reader.ReadString('\n')
			itemInput = strings.TrimSpace(itemInput)
			index, err := strconv.Atoi(itemInput)
			if err == nil && index > 0 && index <= len(player.Inventory) {
				itemName := player.Inventory[index-1]
				useItem(itemName)
			} else {
				fmt.Println("Invalid item choice.")
			}
		case "4":
			fmt.Println("You attempt to flee...")
			if rand.Intn(2) == 0 {
				fmt.Println("You successfully flee!")
				return
			} else {
				fmt.Println("You failed to flee!")
			}
		default:
			fmt.Println("Invalid action.")
		}
		if enemy.Health > 0 {
			enemyDamage := enemy.Attack - player.Defense
			if enemyDamage < 0 {
				enemyDamage = 0
			}
			player.Health -= enemyDamage
			fmt.Printf("The %s attacks you for %d damage!\n", enemy.Name, enemyDamage)
		}
	}
	if player.Health <= 0 {
		fmt.Println("\nYou have been defeated! Game Over.")
		os.Exit(0)
	} else {
		fmt.Printf("\nYou defeated the %s!\n", enemy.Name)
		player.Exp += enemy.ExpReward
		player.Gold += enemy.GoldReward
		fmt.Printf("Gained %d exp and %d gold.\n", enemy.ExpReward, enemy.GoldReward)
		if rand.Intn(2) == 0 && len(enemy.Drops) > 0 {
			drop := enemy.Drops[rand.Intn(len(enemy.Drops))]
			player.Inventory = append(player.Inventory, drop)
			fmt.Printf("The %s dropped: %s\n", enemy.Name, drop)
		}
		checkLevelUp()
	}
}

func useItem(itemName string) {
	for _, item := range items {
		if item.Name == itemName {
			switch item.Effect {
			case "heal":
				player.Health += 20
				if player.Health > player.MaxHealth {
					player.Health = player.MaxHealth
				}
				fmt.Println("You used a Health Potion and restored 20 health.")
			case "mana":
				fmt.Println("Mana potion effect not implemented in this version.")
			default:
				fmt.Printf("Used %s.\n", item.Name)
			}
			for i, invItem := range player.Inventory {
				if invItem == itemName {
					player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
					break
				}
			}
			return
		}
	}
	fmt.Println("Item not found.")
}

func checkLevelUp() {
	for player.Exp >= player.Level*100 {
		player.Exp -= player.Level * 100
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("\nCongratulations! You leveled up to Level %d!\n", player.Level)
		fmt.Println("Health, Attack, and Defense increased.")
	}
}

func visitVillage() {
	fmt.Println("\nWelcome to the village!")
	for {
		fmt.Println("What would you like to do?")
		fmt.Println("1. Visit the Shop")
		fmt.Println("2. Rest at the Inn (Restores Health)")
		fmt.Println("3. Leave Village")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			visitShop()
		case "2":
			player.Health = player.MaxHealth
			fmt.Println("You rest at the inn and fully restore your health.")
		case "3":
			fmt.Println("You leave the village.")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}

func visitShop() {
	fmt.Println("\nWelcome to the shop!")
	for {
		fmt.Printf("Your Gold: %d\n", player.Gold)
		fmt.Println("Items for sale:")
		for i, item := range items {
			fmt.Printf("%d. %s - %s (Cost: %d gold)\n", i+1, item.Name, item.Description, item.Value)
		}
		fmt.Println("0. Exit Shop")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "0" {
			fmt.Println("You leave the shop.")
			return
		}
		index, err := strconv.Atoi(input)
		if err == nil && index > 0 && index <= len(items) {
			item := items[index-1]
			if player.Gold >= item.Value {
				player.Gold -= item.Value
				player.Inventory = append(player.Inventory, item.Name)
				fmt.Printf("You bought %s for %d gold.\n", item.Name, item.Value)
			} else {
				fmt.Println("Not enough gold!")
			}
		} else {
			fmt.Println("Invalid choice.")
		}
	}
}

func main() {
	initGame()
	fmt.Println("Welcome to Fantasy Adventure Game!")
	fmt.Println("You are a hero on a quest. Explore, battle enemies, and level up!")
	reader := bufio.NewReader(os.Stdin)
	for {
		displayStatus()
		fmt.Println("\nWhat would you like to do?")
		fmt.Println("1. Explore")
		fmt.Println("2. Visit Village")
		fmt.Println("3. Quit Game")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			explore()
		case "2":
			visitVillage()
		case "3":
			fmt.Println("Thanks for playing! Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
		}
	}
}
