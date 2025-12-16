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

type GameState struct {
	PlayerName     string
	PlayerHealth   int
	PlayerMaxHealth int
	PlayerAttack   int
	PlayerDefense  int
	PlayerGold     int
	PlayerLevel    int
	PlayerXP       int
	PlayerXPToNext int
	Inventory      map[string]int
	CurrentArea    string
	GameOver       bool
}

type Enemy struct {
	Name     string
	Health   int
	Attack   int
	Defense  int
	GoldDrop int
	XPReward int
}

type Item struct {
	Name        string
	Description string
	Effect      string
	Value       int
}

var gameState GameState
var enemies = map[string]Enemy{
	"goblin": {"Goblin", 30, 5, 2, 10, 20},
	"orc": {"Orc", 50, 8, 4, 25, 40},
	"dragon": {"Dragon", 100, 15, 10, 100, 100},
}
var items = map[string]Item{
	"potion": {"Health Potion", "Restores 20 health", "heal", 20},
	"sword": {"Iron Sword", "Increases attack by 5", "attack", 5},
	"shield": {"Wooden Shield", "Increases defense by 3", "defense", 3},
}

func initGame() {
	gameState = GameState{
		PlayerName:     "",
		PlayerHealth:   100,
		PlayerMaxHealth: 100,
		PlayerAttack:   10,
		PlayerDefense:  5,
		PlayerGold:     50,
		PlayerLevel:    1,
		PlayerXP:       0,
		PlayerXPToNext: 100,
		Inventory:      make(map[string]int),
		CurrentArea:    "town",
		GameOver:       false,
	}
	gameState.Inventory["potion"] = 3
	gameState.Inventory["gold"] = gameState.PlayerGold
}

func printStatus() {
	fmt.Println("\n=== Status ===")
	fmt.Printf("Name: %s\n", gameState.PlayerName)
	fmt.Printf("Health: %d/%d\n", gameState.PlayerHealth, gameState.PlayerMaxHealth)
	fmt.Printf("Attack: %d\n", gameState.PlayerAttack)
	fmt.Printf("Defense: %d\n", gameState.PlayerDefense)
	fmt.Printf("Gold: %d\n", gameState.PlayerGold)
	fmt.Printf("Level: %d\n", gameState.PlayerLevel)
	fmt.Printf("XP: %d/%d\n", gameState.PlayerXP, gameState.PlayerXPToNext)
	fmt.Println("Inventory:")
	for item, count := range gameState.Inventory {
		if item != "gold" {
			fmt.Printf("  %s: %d\n", items[item].Name, count)
		}
	}
	fmt.Printf("Current Area: %s\n", gameState.CurrentArea)
}

func handleCombat(enemyKey string) {
	enemy := enemies[enemyKey]
	fmt.Printf("\nA wild %s appears!\n", enemy.Name)
	for gameState.PlayerHealth > 0 && enemy.Health > 0 {
		fmt.Println("\n--- Combat Round ---")
		fmt.Printf("Your Health: %d, %s's Health: %d\n", gameState.PlayerHealth, enemy.Name, enemy.Health)
		fmt.Println("Choose action: (1) Attack, (2) Use Item, (3) Flee")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			playerDamage := gameState.PlayerAttack - enemy.Defense
			if playerDamage < 1 {
				playerDamage = 1
			}
			enemy.Health -= playerDamage
			fmt.Printf("You deal %d damage to %s.\n", playerDamage, enemy.Name)
		case "2":
			useItem()
			continue
		case "3":
			fmt.Println("You flee from combat!")
			return
		default:
			fmt.Println("Invalid choice, you hesitate!")
			continue
		}
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			gameState.PlayerGold += enemy.GoldDrop
			gameState.Inventory["gold"] = gameState.PlayerGold
			gameState.PlayerXP += enemy.XPReward
			fmt.Printf("Gained %d gold and %d XP.\n", enemy.GoldDrop, enemy.XPReward)
			checkLevelUp()
			return
		}
		enemyDamage := enemy.Attack - gameState.PlayerDefense
		if enemyDamage < 1 {
			enemyDamage = 1
		}
		gameState.PlayerHealth -= enemyDamage
		fmt.Printf("%s deals %d damage to you.\n", enemy.Name, enemyDamage)
	}
	if gameState.PlayerHealth <= 0 {
		fmt.Println("\nYou have been defeated! Game Over.")
		gameState.GameOver = true
	}
}

func useItem() {
	fmt.Println("\nAvailable items:")
	for item, count := range gameState.Inventory {
		if item != "gold" && count > 0 {
			fmt.Printf("%s: %d - %s\n", items[item].Name, count, items[item].Description)
		}
	}
	fmt.Print("Enter item name to use (or 'back' to cancel): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "back" {
		return
	}
	itemKey := ""
	for key, item := range items {
		if strings.EqualFold(item.Name, input) {
			itemKey = key
			break
		}
	}
	if itemKey == "" {
		fmt.Println("Item not found.")
		return
	}
	if gameState.Inventory[itemKey] <= 0 {
		fmt.Println("You don't have any of that item.")
		return
	}
	item := items[itemKey]
	switch item.Effect {
	case "heal":
		gameState.PlayerHealth += item.Value
		if gameState.PlayerHealth > gameState.PlayerMaxHealth {
			gameState.PlayerHealth = gameState.PlayerMaxHealth
		}
		fmt.Printf("Used %s, healed %d health. Current health: %d\n", item.Name, item.Value, gameState.PlayerHealth)
	case "attack":
		gameState.PlayerAttack += item.Value
		fmt.Printf("Used %s, attack increased by %d. Current attack: %d\n", item.Name, item.Value, gameState.PlayerAttack)
	case "defense":
		gameState.PlayerDefense += item.Value
		fmt.Printf("Used %s, defense increased by %d. Current defense: %d\n", item.Name, item.Value, gameState.PlayerDefense)
	}
	gameState.Inventory[itemKey]--
}

func checkLevelUp() {
	for gameState.PlayerXP >= gameState.PlayerXPToNext {
		gameState.PlayerLevel++
		gameState.PlayerXP -= gameState.PlayerXPToNext
		gameState.PlayerXPToNext = gameState.PlayerLevel * 100
		gameState.PlayerMaxHealth += 20
		gameState.PlayerHealth = gameState.PlayerMaxHealth
		gameState.PlayerAttack += 2
		gameState.PlayerDefense += 1
		fmt.Printf("\nLevel up! You are now level %d. Stats increased!\n", gameState.PlayerLevel)
	}
}

func visitShop() {
	fmt.Println("\n=== Shop ===")
	fmt.Println("Welcome to the shop! Available items:")
	for key, item := range items {
		fmt.Printf("%s - %s (Cost: %d gold)\n", item.Name, item.Description, item.Value*2)
	}
	fmt.Printf("Your gold: %d\n", gameState.PlayerGold)
	fmt.Println("Enter item name to buy (or 'back' to leave):")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	if input == "back" {
		return
	}
	itemKey := ""
	for key, item := range items {
		if strings.EqualFold(item.Name, input) {
			itemKey = key
			break
		}
	}
	if itemKey == "" {
		fmt.Println("Item not found.")
		return
	}
	cost := items[itemKey].Value * 2
	if gameState.PlayerGold < cost {
		fmt.Println("Not enough gold!")
		return
	}
	gameState.PlayerGold -= cost
	gameState.Inventory["gold"] = gameState.PlayerGold
	gameState.Inventory[itemKey]++
	fmt.Printf("Bought %s for %d gold. You now have %d.\n", items[itemKey].Name, cost, gameState.Inventory[itemKey])
}

func exploreArea() {
	fmt.Println("\nYou venture into the wilderness...")
	rand.Seed(time.Now().UnixNano())
	encounterChance := rand.Intn(100)
	if encounterChance < 60 {
		enemyKeys := []string{"goblin", "orc", "dragon"}
		enemyKey := enemyKeys[rand.Intn(len(enemyKeys))]
		handleCombat(enemyKey)
	} else {
		fmt.Println("You find nothing of interest.")
	}
}

func main() {
	fmt.Println("Welcome to the Fun Terminal Game!")
	fmt.Print("Enter your character name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	gameState.PlayerName = strings.TrimSpace(name)
	initGame()
	for !gameState.GameOver {
		printStatus()
		fmt.Println("\n=== Main Menu ===")
		fmt.Println("Choose an option:")
		fmt.Println("1. Explore")
		fmt.Println("2. Visit Shop")
		fmt.Println("3. Use Item")
		fmt.Println("4. Quit Game")
		fmt.Print("Enter choice (1-4): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			exploreArea()
		case "2":
			visitShop()
		case "3":
			useItem()
		case "4":
			fmt.Println("Thanks for playing! Goodbye.")
			return
		default:
			fmt.Println("Invalid choice, please try again.")
		}
	}
	fmt.Println("Game Over. Thanks for playing!")
}
