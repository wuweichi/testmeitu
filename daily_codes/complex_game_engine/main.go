package main

import (
	"fmt"
	"math/rand"
	"time"
	"os"
	"bufio"
	"strconv"
	"strings"
)

type Player struct {
	Name     string
	Health   int
	Mana     int
	Strength int
	Agility  int
	Level    int
	Exp      int
}

type Enemy struct {
	Name     string
	Health   int
	Strength int
	Agility  int
	ExpGiven int
}

type Item struct {
	Name        string
	Description string
	Type        string
	Value       int
}

type GameState struct {
	Player      Player
	Enemies     []Enemy
	Inventory   []Item
	CurrentRoom int
	GameOver    bool
}

func (p *Player) Attack(e *Enemy) int {
	damage := p.Strength + rand.Intn(10)
	e.Health -= damage
	return damage
}

func (e *Enemy) Attack(p *Player) int {
	damage := e.Strength + rand.Intn(5)
	p.Health -= damage
	return damage
}

func (p *Player) Heal() bool {
	if p.Mana >= 10 {
		healAmount := 20 + rand.Intn(10)
		p.Health += healAmount
		p.Mana -= 10
		return true
	}
	return false
}

func (p *Player) LevelUp() {
	if p.Exp >= p.Level*100 {
		p.Exp -= p.Level * 100
		p.Level++
		p.Health += 20
		p.Mana += 10
		p.Strength += 2
		p.Agility += 1
		fmt.Printf("Level up! You are now level %d\n", p.Level)
	}
}

func (g *GameState) GenerateEnemies() {
	enemyTypes := []Enemy{
		{"Goblin", 30, 5, 3, 10},
		{"Orc", 50, 8, 2, 20},
		{"Dragon", 100, 15, 1, 50},
		{"Slime", 20, 3, 4, 5},
		{"Skeleton", 40, 6, 3, 15},
		{"Witch", 35, 7, 5, 25},
		{"Giant", 80, 12, 1, 40},
		{"Wolf", 25, 4, 6, 8},
		{"Zombie", 45, 5, 2, 12},
		{"Vampire", 60, 9, 7, 30},
	}
	g.Enemies = nil
	numEnemies := rand.Intn(3) + 1
	for i := 0; i < numEnemies; i++ {
		enemy := enemyTypes[rand.Intn(len(enemyTypes))]
		g.Enemies = append(g.Enemies, enemy)
	}
}

func (g *GameState) Battle() {
	fmt.Println("A battle begins!")
	for len(g.Enemies) > 0 && g.Player.Health > 0 {
		fmt.Printf("Your health: %d, Mana: %d\n", g.Player.Health, g.Player.Mana)
		fmt.Println("Enemies:")
		for i, enemy := range g.Enemies {
			fmt.Printf("%d. %s (Health: %d)\n", i+1, enemy.Name, enemy.Health)
		}
		fmt.Print("Choose action: (1) Attack, (2) Heal, (3) Flee: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			fmt.Print("Choose enemy to attack: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			index, err := strconv.Atoi(input)
			if err != nil || index < 1 || index > len(g.Enemies) {
				fmt.Println("Invalid choice")
				continue
			}
			enemy := &g.Enemies[index-1]
			damage := g.Player.Attack(enemy)
			fmt.Printf("You attack %s for %d damage!\n", enemy.Name, damage)
			if enemy.Health <= 0 {
				fmt.Printf("%s defeated! You gain %d exp.\n", enemy.Name, enemy.ExpGiven)
				g.Player.Exp += enemy.ExpGiven
				g.Enemies = append(g.Enemies[:index-1], g.Enemies[index:]...)
			}
		case "2":
			if g.Player.Heal() {
				fmt.Println("You heal yourself!")
			} else {
				fmt.Println("Not enough mana!")
			}
		case "3":
			fmt.Println("You flee from battle!")
			return
		default:
			fmt.Println("Invalid action")
			continue
		}
		for i := range g.Enemies {
			if g.Enemies[i].Health > 0 {
				damage := g.Enemies[i].Attack(&g.Player)
				fmt.Printf("%s attacks you for %d damage!\n", g.Enemies[i].Name, damage)
			}
		}
		g.Player.LevelUp()
		if g.Player.Health <= 0 {
			fmt.Println("You have been defeated!")
			g.GameOver = true
			return
		}
	}
	if g.Player.Health > 0 {
		fmt.Println("You win the battle!")
	}
}

func (g *GameState) Explore() {
	g.CurrentRoom++
	fmt.Printf("You enter room %d\n", g.CurrentRoom)
	g.GenerateEnemies()
	if len(g.Enemies) > 0 {
		g.Battle()
	} else {
		fmt.Println("The room is empty.")
	}
}

func (g *GameState) Rest() {
	g.Player.Health = 100
	g.Player.Mana = 50
	fmt.Println("You rest and recover your health and mana.")
}

func (g *GameState) ShowStatus() {
	fmt.Printf("Name: %s\n", g.Player.Name)
	fmt.Printf("Level: %d\n", g.Player.Level)
	fmt.Printf("Health: %d\n", g.Player.Health)
	fmt.Printf("Mana: %d\n", g.Player.Mana)
	fmt.Printf("Strength: %d\n", g.Player.Strength)
	fmt.Printf("Agility: %d\n", g.Player.Agility)
	fmt.Printf("Exp: %d/%d\n", g.Player.Exp, g.Player.Level*100)
	fmt.Printf("Current Room: %d\n", g.CurrentRoom)
}

func (g *GameState) ShowInventory() {
	if len(g.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("Inventory:")
	for i, item := range g.Inventory {
		fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.Description)
	}
}

func (g *GameState) AddItem(item Item) {
	g.Inventory = append(g.Inventory, item)
	fmt.Printf("You obtained: %s\n", item.Name)
}

func (g *GameState) UseItem() {
	if len(g.Inventory) == 0 {
		fmt.Println("No items to use.")
		return
	}
	g.ShowInventory()
	fmt.Print("Choose item to use: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(g.Inventory) {
		fmt.Println("Invalid choice")
		return
	}
	item := g.Inventory[index-1]
	switch item.Type {
	case "potion":
		g.Player.Health += item.Value
		fmt.Printf("You use %s and recover %d health.\n", item.Name, item.Value)
	case "elixir":
		g.Player.Mana += item.Value
		fmt.Printf("You use %s and recover %d mana.\n", item.Name, item.Value)
	default:
		fmt.Println("This item cannot be used.")
		return
	}
	g.Inventory = append(g.Inventory[:index-1], g.Inventory[index:]...)
}

func (g *GameState) RandomEvent() {
	eventChance := rand.Intn(100)
	if eventChance < 20 {
		items := []Item{
			{"Health Potion", "Restores 20 health", "potion", 20},
			{"Mana Elixir", "Restores 15 mana", "elixir", 15},
			{"Gold Coin", "Shiny coin", "currency", 1},
		}
		item := items[rand.Intn(len(items))]
		g.AddItem(item)
	} else if eventChance < 40 {
		fmt.Println("You find a hidden treasure!")
		g.Player.Exp += 10
	} else if eventChance < 60 {
		fmt.Println("A trap is triggered!")
		damage := 10 + rand.Intn(10)
		g.Player.Health -= damage
		fmt.Printf("You take %d damage.\n", damage)
	} else {
		fmt.Println("Nothing happens.")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := GameState{
		Player: Player{
			Name:     "Hero",
			Health:   100,
			Mana:     50,
			Strength: 10,
			Agility:  5,
			Level:    1,
			Exp:      0,
		},
		CurrentRoom: 0,
		GameOver:    false,
	}
	fmt.Println("Welcome to the Complex Game Engine!")
	fmt.Print("Enter your character name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	game.Player.Name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s! Let's begin your adventure.\n", game.Player.Name)
	for !game.GameOver {
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Explore next room")
		fmt.Println("2. Rest and recover")
		fmt.Println("3. Check status")
		fmt.Println("4. Check inventory")
		fmt.Println("5. Use item")
		fmt.Println("6. Quit game")
		fmt.Print("Enter your choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			game.Explore()
			game.RandomEvent()
		case "2":
			game.Rest()
		case "3":
			game.ShowStatus()
		case "4":
			game.ShowInventory()
		case "5":
			game.UseItem()
		case "6":
			fmt.Println("Thanks for playing!")
			game.GameOver = true
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
		if game.Player.Health <= 0 {
			fmt.Println("Game Over!")
			game.GameOver = true
		}
	}
}
