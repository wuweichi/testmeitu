package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"strings"
)

type Player struct {
	Name     string
	Health   int
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
}

type Item struct {
	Name        string
	Description string
	Value       int
}

type GameState struct {
	Player      Player
	Enemies     []Enemy
	Inventory   []Item
	GameActive  bool
	CurrentRoom int
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

func (p *Player) LevelUp() {
	if p.Exp >= p.Level*100 {
		p.Level++
		p.Strength += 5
		p.Agility += 3
		p.Health += 20
		p.Exp = 0
		fmt.Printf("Level up! You are now level %d\n", p.Level)
	}
}

func (g *GameState) GenerateEnemies() {
	enemyNames := []string{"Goblin", "Orc", "Troll", "Dragon", "Skeleton"}
	g.Enemies = nil
	numEnemies := rand.Intn(3) + 1
	for i := 0; i < numEnemies; i++ {
		name := enemyNames[rand.Intn(len(enemyNames))]
		enemy := Enemy{
			Name:     name,
			Health:   rand.Intn(50) + 20,
			Strength: rand.Intn(10) + 5,
			Agility:  rand.Intn(5) + 1,
		}
		g.Enemies = append(g.Enemies, enemy)
	}
}

func (g *GameState) GenerateItems() {
	itemNames := []string{"Health Potion", "Strength Elixir", "Agility Boost", "Gold Coin"}
	g.Inventory = nil
	numItems := rand.Intn(3) + 1
	for i := 0; i < numItems; i++ {
		name := itemNames[rand.Intn(len(itemNames))]
		item := Item{
			Name:        name,
			Description: fmt.Sprintf("A useful %s", name),
			Value:       rand.Intn(50) + 10,
		}
		g.Inventory = append(g.Inventory, item)
	}
}

func (g *GameState) DisplayStatus() {
	fmt.Printf("Player: %s (Level %d)\n", g.Player.Name, g.Player.Level)
	fmt.Printf("Health: %d, Strength: %d, Agility: %d\n", g.Player.Health, g.Player.Strength, g.Player.Agility)
	fmt.Printf("Experience: %d/%d\n", g.Player.Exp, g.Player.Level*100)
	fmt.Printf("Current Room: %d\n", g.CurrentRoom)
}

func (g *GameState) DisplayEnemies() {
	if len(g.Enemies) == 0 {
		fmt.Println("No enemies in this room.")
		return
	}
	fmt.Println("Enemies in the room:")
	for i, enemy := range g.Enemies {
		fmt.Printf("%d. %s (Health: %d, Strength: %d, Agility: %d)\n", i+1, enemy.Name, enemy.Health, enemy.Strength, enemy.Agility)
	}
}

func (g *GameState) DisplayInventory() {
	if len(g.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("Your inventory:")
	for i, item := range g.Inventory {
		fmt.Printf("%d. %s - %s (Value: %d)\n", i+1, item.Name, item.Description, item.Value)
	}
}

func (g *GameState) UseItem(index int) {
	if index < 0 || index >= len(g.Inventory) {
		fmt.Println("Invalid item index.")
		return
	}
	item := g.Inventory[index]
	switch item.Name {
	case "Health Potion":
		g.Player.Health += 20
		fmt.Println("You used a Health Potion and restored 20 health.")
	case "Strength Elixir":
		g.Player.Strength += 5
		fmt.Println("You used a Strength Elixir and increased strength by 5.")
	case "Agility Boost":
		g.Player.Agility += 3
		fmt.Println("You used an Agility Boost and increased agility by 3.")
	default:
		fmt.Printf("You used %s but nothing happened.\n", item.Name)
	}
	g.Inventory = append(g.Inventory[:index], g.Inventory[index+1:]...)
}

func (g *GameState) FightEnemy(index int) {
	if index < 0 || index >= len(g.Enemies) {
		fmt.Println("Invalid enemy index.")
		return
	}
	enemy := &g.Enemies[index]
	fmt.Printf("You are fighting %s!\n", enemy.Name)
	for g.Player.Health > 0 && enemy.Health > 0 {
		playerDamage := g.Player.Attack(enemy)
		fmt.Printf("You attack %s for %d damage. %s's health: %d\n", enemy.Name, playerDamage, enemy.Name, enemy.Health)
		if enemy.Health <= 0 {
			fmt.Printf("You defeated %s!\n", enemy.Name)
			g.Player.Exp += 30
			g.Player.LevelUp()
			g.Enemies = append(g.Enemies[:index], g.Enemies[index+1:]...)
			return
		}
		enemyDamage := enemy.Attack(&g.Player)
		fmt.Printf("%s attacks you for %d damage. Your health: %d\n", enemy.Name, enemyDamage, g.Player.Health)
		if g.Player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			g.GameActive = false
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func (g *GameState) MoveToNextRoom() {
	g.CurrentRoom++
	fmt.Printf("You move to room %d.\n", g.CurrentRoom)
	g.GenerateEnemies()
	g.GenerateItems()
}

func (g *GameState) HandleCommand(command string) {
	parts := strings.Fields(command)
	if len(parts) == 0 {
		fmt.Println("Please enter a command.")
		return
	}
	switch parts[0] {
	case "status":
		g.DisplayStatus()
	case "enemies":
		g.DisplayEnemies()
	case "inventory":
		g.DisplayInventory()
	case "use":
		if len(parts) < 2 {
			fmt.Println("Usage: use <item_index>")
			return
		}
		index, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Invalid item index.")
			return
		}
		g.UseItem(index - 1)
	case "fight":
		if len(parts) < 2 {
			fmt.Println("Usage: fight <enemy_index>")
			return
		}
		index, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Println("Invalid enemy index.")
			return
		}
		g.FightEnemy(index - 1)
	case "move":
		g.MoveToNextRoom()
	case "quit":
		fmt.Println("Thanks for playing!")
		g.GameActive = false
	default:
		fmt.Println("Unknown command. Available commands: status, enemies, inventory, use, fight, move, quit")
	}
}

func main() {
	fmt.Println("Welcome to the Complex Game Simulator!")
	fmt.Println("You are an adventurer exploring a dungeon.")
	fmt.Println("Commands: status, enemies, inventory, use <item_index>, fight <enemy_index>, move, quit")
	
	game := GameState{
		Player: Player{
			Name:     "Hero",
			Health:   100,
			Strength: 10,
			Agility:  5,
			Level:    1,
			Exp:      0,
		},
		GameActive:  true,
		CurrentRoom: 1,
	}
	
	game.GenerateEnemies()
	game.GenerateItems()
	
	for game.GameActive {
		fmt.Print("> ")
		var command string
		fmt.Scanln(&command)
		game.HandleCommand(command)
	}
}
