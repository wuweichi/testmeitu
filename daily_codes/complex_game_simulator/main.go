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
	Exp       int
	Inventory []string
}

func (c *Character) DisplayStats() {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Level: %d\n", c.Level)
	fmt.Printf("Health: %d/%d\n", c.Health, c.MaxHealth)
	fmt.Printf("Attack: %d\n", c.Attack)
	fmt.Printf("Defense: %d\n", c.Defense)
	fmt.Printf("Experience: %d\n", c.Exp)
	fmt.Print("Inventory: ")
	if len(c.Inventory) == 0 {
		fmt.Println("Empty")
	} else {
		for i, item := range c.Inventory {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Print(item)
		}
		fmt.Println()
	}
}

func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.Attack += 2
	c.Defense += 1
	fmt.Printf("%s leveled up to level %d!\n", c.Name, c.Level)
}

func (c *Character) GainExp(amount int) {
	c.Exp += amount
	fmt.Printf("%s gained %d experience points.\n", c.Name, amount)
	for c.Exp >= c.Level*100 {
		c.Exp -= c.Level * 100
		c.LevelUp()
	}
}

func (c *Character) IsAlive() bool {
	return c.Health > 0
}

func (c *Character) AttackTarget(target *Character) {
	damage := c.Attack - target.Defense
	if damage < 1 {
		damage = 1
	}
	target.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", c.Name, target.Name, damage)
	if !target.IsAlive() {
		fmt.Printf("%s has been defeated!\n", target.Name)
		c.GainExp(target.Level * 10)
	}
}

func (c *Character) Heal(amount int) {
	c.Health += amount
	if c.Health > c.MaxHealth {
		c.Health = c.MaxHealth
	}
	fmt.Printf("%s healed for %d health. Current health: %d/%d\n", c.Name, amount, c.Health, c.MaxHealth)
}

func (c *Character) AddItem(item string) {
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("%s added to inventory.\n", item)
}

func (c *Character) UseItem(itemIndex int) {
	if itemIndex < 0 || itemIndex >= len(c.Inventory) {
		fmt.Println("Invalid item index.")
		return
	}
	item := c.Inventory[itemIndex]
	c.Inventory = append(c.Inventory[:itemIndex], c.Inventory[itemIndex+1:]...)
	switch item {
	case "Health Potion":
		c.Heal(30)
	case "Attack Boost":
		c.Attack += 5
		fmt.Printf("%s's attack increased by 5!\n", c.Name)
	case "Defense Boost":
		c.Defense += 5
		fmt.Printf("%s's defense increased by 5!\n", c.Name)
	default:
		fmt.Printf("Used %s.\n", item)
	}
}

type Enemy struct {
	Name    string
	Health  int
	Attack  int
	Defense int
	Level   int
}

func (e *Enemy) DisplayStats() {
	fmt.Printf("Enemy: %s\n", e.Name)
	fmt.Printf("Level: %d\n", e.Level)
	fmt.Printf("Health: %d\n", e.Health)
	fmt.Printf("Attack: %d\n", e.Attack)
	fmt.Printf("Defense: %d\n", e.Defense)
}

func (e *Enemy) IsAlive() bool {
	return e.Health > 0
}

func (e *Enemy) AttackTarget(target *Character) {
	damage := e.Attack - target.Defense
	if damage < 1 {
		damage = 1
	}
	target.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", e.Name, target.Name, damage)
	if !target.IsAlive() {
		fmt.Printf("%s has been defeated!\n", target.Name)
	}
}

func GenerateRandomEnemy(level int) *Enemy {
	names := []string{"Goblin", "Orc", "Skeleton", "Zombie", "Dragon", "Troll"}
	name := names[rand.Intn(len(names))]
	health := 20 + (level * 5)
	attack := 5 + (level * 2)
	defense := 2 + level
	return &Enemy{
		Name:    name,
		Health:  health,
		Attack:  attack,
		Defense: defense,
		Level:   level,
	}
}

type Game struct {
	Player *Character
	Round  int
}

func NewGame(playerName string) *Game {
	player := &Character{
		Name:      playerName,
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
		Inventory: []string{"Health Potion", "Health Potion"},
	}
	return &Game{
		Player: player,
		Round:  1,
	}
}

func (g *Game) StartBattle() {
	enemyLevel := g.Player.Level + rand.Intn(3) - 1
	if enemyLevel < 1 {
		enemyLevel = 1
	}
	enemy := GenerateRandomEnemy(enemyLevel)
	fmt.Printf("\n--- Round %d: Battle Start! ---\n", g.Round)
	fmt.Printf("A wild %s (Level %d) appears!\n", enemy.Name, enemy.Level)
	for g.Player.IsAlive() && enemy.IsAlive() {
		fmt.Println("\n--- Your Turn ---")
		g.PlayerTurn(enemy)
		if !enemy.IsAlive() {
			break
		}
		fmt.Println("\n--- Enemy Turn ---")
		enemy.AttackTarget(g.Player)
	}
	if g.Player.IsAlive() {
		fmt.Printf("\nYou defeated the %s!\n", enemy.Name)
		g.Round++
		// Chance to find item after battle
		if rand.Intn(100) < 30 {
			items := []string{"Health Potion", "Attack Boost", "Defense Boost"}
			foundItem := items[rand.Intn(len(items))]
			g.Player.AddItem(foundItem)
		}
	} else {
		fmt.Println("\nGame Over! You have been defeated.")
	}
}

func (g *Game) PlayerTurn(enemy *Enemy) {
	for {
		fmt.Println("Choose an action:")
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Check Stats")
		fmt.Println("4. Flee")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			g.Player.AttackTarget(enemy)
			return
		case 2:
			g.UseItemMenu()
			continue
		case 3:
			fmt.Println("\n--- Your Stats ---")
			g.Player.DisplayStats()
			fmt.Println("\n--- Enemy Stats ---")
			enemy.DisplayStats()
			continue
		case 4:
			if rand.Intn(100) < 50 {
				fmt.Println("You successfully fled from the battle!")
				g.Round++
				return
			} else {
				fmt.Println("You failed to flee!")
				return
			}
		default:
			fmt.Println("Invalid choice. Please try again.")
			continue
		}
	}
}

func (g *Game) UseItemMenu() {
	if len(g.Player.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("Your Inventory:")
	for i, item := range g.Player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Print("Choose an item to use (0 to cancel): ")
	var choice int
	fmt.Scan(&choice)
	if choice == 0 {
		return
	}
	if choice < 1 || choice > len(g.Player.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}
	g.Player.UseItem(choice - 1)
}

func (g *Game) DisplayGameStatus() {
	fmt.Printf("\n--- Game Status ---\n")
	fmt.Printf("Round: %d\n", g.Round)
	g.Player.DisplayStats()
}

func (g *Game) MainMenu() {
	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Start Battle")
		fmt.Println("2. Check Stats")
		fmt.Println("3. Use Item")
		fmt.Println("4. Exit Game")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			g.StartBattle()
			if !g.Player.IsAlive() {
				return
			}
		case 2:
			g.DisplayGameStatus()
		case 3:
			g.UseItemMenu()
		case 4:
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to the Complex Game Simulator!")
	fmt.Print("Enter your character's name: ")
	var playerName string
	fmt.Scan(&playerName)
	game := NewGame(playerName)
	game.MainMenu()
}
