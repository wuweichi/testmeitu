package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Character struct {
	Name      string
	Health    int
	MaxHealth int
	Attack    int
	Defense   int
	Level     int
	Experience int
}

type Item struct {
	Name        string
	Description string
	Value       int
}

type GameState struct {
	Player     Character
	Inventory  []Item
	GameTime   time.Duration
	IsRunning  bool
}

func (c *Character) TakeDamage(damage int) {
	actualDamage := damage - c.Defense
	if actualDamage < 0 {
		actualDamage = 0
	}
	c.Health -= actualDamage
	if c.Health < 0 {
		c.Health = 0
	}
}

func (c *Character) Heal(amount int) {
	c.Health += amount
	if c.Health > c.MaxHealth {
		c.Health = c.MaxHealth
	}
}

func (c *Character) GainExperience(exp int) {
	c.Experience += exp
	if c.Experience >= c.Level*100 {
		c.LevelUp()
	}
}

func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.Attack += 2
	c.Defense += 1
	c.Experience = 0
}

func NewCharacter(name string) Character {
	return Character{
		Name:      name,
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Experience: 0,
	}
}

func NewGameState() GameState {
	return GameState{
		Player:    NewCharacter("Hero"),
		Inventory: []Item{},
		GameTime:  0,
		IsRunning: true,
	}
}

func (gs *GameState) AddItem(item Item) {
	gs.Inventory = append(gs.Inventory, item)
}

func (gs *GameState) RemoveItem(index int) {
	if index >= 0 && index < len(gs.Inventory) {
		gs.Inventory = append(gs.Inventory[:index], gs.Inventory[index+1:]...)
	}
}

func (gs *GameState) PrintStatus() {
	fmt.Printf("Player: %s\n", gs.Player.Name)
	fmt.Printf("Health: %d/%d\n", gs.Player.Health, gs.Player.MaxHealth)
	fmt.Printf("Level: %d\n", gs.Player.Level)
	fmt.Printf("Experience: %d\n", gs.Player.Experience)
	fmt.Printf("Attack: %d\n", gs.Player.Attack)
	fmt.Printf("Defense: %d\n", gs.Player.Defense)
	fmt.Printf("Game Time: %v\n", gs.GameTime)
	fmt.Printf("Inventory: %d items\n", len(gs.Inventory))
	for i, item := range gs.Inventory {
		fmt.Printf("  %d. %s - %s (Value: %d)\n", i+1, item.Name, item.Description, item.Value)
	}
}

func (gs *GameState) Battle(enemy Character) {
	fmt.Printf("A wild %s appears!\n", enemy.Name)
	for gs.Player.Health > 0 && enemy.Health > 0 {
		playerDamage := gs.Player.Attack + rand.Intn(5)
		enemy.TakeDamage(playerDamage)
		fmt.Printf("You attack %s for %d damage!\n", enemy.Name, playerDamage)
		if enemy.Health <= 0 {
			fmt.Printf("You defeated %s!\n", enemy.Name)
			expGained := enemy.Level * 10
			gs.Player.GainExperience(expGained)
			fmt.Printf("Gained %d experience!\n", expGained)
			break
		}
		enemyDamage := enemy.Attack + rand.Intn(3)
		gs.Player.TakeDamage(enemyDamage)
		fmt.Printf("%s attacks you for %d damage!\n", enemy.Name, enemyDamage)
		if gs.Player.Health <= 0 {
			fmt.Printf("You have been defeated by %s!\n", enemy.Name)
			gs.IsRunning = false
			break
		}
	}
}

func (gs *GameState) Explore() {
	event := rand.Intn(10)
	switch event {
	case 0, 1, 2:
		fmt.Println("You find a treasure chest!")
		items := []Item{
			{"Health Potion", "Restores 50 health", 50},
			{"Sword", "Increases attack by 5", 100},
			{"Shield", "Increases defense by 3", 80},
		}
		foundItem := items[rand.Intn(len(items))]
		gs.AddItem(foundItem)
		fmt.Printf("You found: %s - %s\n", foundItem.Name, foundItem.Description)
	case 3, 4, 5:
		fmt.Println("You encounter an enemy!")
		enemies := []Character{
			{"Goblin", 30, 30, 8, 2, 1, 0},
			{"Orc", 50, 50, 12, 4, 2, 0},
			{"Dragon", 100, 100, 20, 10, 5, 0},
		}
		enemy := enemies[rand.Intn(len(enemies))]
		gs.Battle(enemy)
	case 6, 7:
		fmt.Println("You find a healing spring.")
		healAmount := 20 + rand.Intn(20)
		gs.Player.Heal(healAmount)
		fmt.Printf("Restored %d health!\n", healAmount)
	case 8, 9:
		fmt.Println("Nothing interesting happens.")
	}
}

func (gs *GameState) UseItem(index int) {
	if index < 0 || index >= len(gs.Inventory) {
		fmt.Println("Invalid item index.")
		return
	}
	item := gs.Inventory[index]
	switch item.Name {
	case "Health Potion":
		gs.Player.Heal(50)
		fmt.Println("Used Health Potion! Restored 50 health.")
		gs.RemoveItem(index)
	case "Sword":
		gs.Player.Attack += 5
		fmt.Println("Equipped Sword! Attack increased by 5.")
		gs.RemoveItem(index)
	case "Shield":
		gs.Player.Defense += 3
		fmt.Println("Equipped Shield! Defense increased by 3.")
		gs.RemoveItem(index)
	default:
		fmt.Printf("Cannot use %s.\n", item.Name)
	}
}

func (gs *GameState) HandleCommand(command string) {
	switch command {
	case "status":
		gs.PrintStatus()
	case "explore":
		gs.Explore()
		gs.GameTime += time.Minute * 10
	case "use":
		if len(gs.Inventory) == 0 {
			fmt.Println("No items in inventory.")
			return
		}
		fmt.Println("Select an item to use:")
		for i, item := range gs.Inventory {
			fmt.Printf("%d. %s\n", i+1, item.Name)
		}
		var index int
		fmt.Print("Enter item number: ")
		_, err := fmt.Scan(&index)
		if err != nil {
			fmt.Println("Invalid input.")
			return
		}
		gs.UseItem(index - 1)
	case "quit":
		gs.IsRunning = false
		fmt.Println("Thanks for playing!")
	default:
		fmt.Println("Unknown command. Available commands: status, explore, use, quit")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := NewGameState()
	fmt.Println("Welcome to the Complex Game Simulator!")
	fmt.Println("Available commands: status, explore, use, quit")
	for game.IsRunning {
		var command string
		fmt.Print("> ")
		_, err := fmt.Scan(&command)
		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}
		game.HandleCommand(command)
	}
}
