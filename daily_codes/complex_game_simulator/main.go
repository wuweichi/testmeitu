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
}

type Monster struct {
	Name      string
	Health    int
	Attack    int
	Defense   int
	ExpReward int
}

type Item struct {
	Name        string
	Description string
	Value       int
}

type GameState struct {
	Player     Character
	Monsters   []Monster
	Inventory  []Item
	Gold       int
	GameRound  int
}

func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.Attack += 2
	c.Defense += 1
	fmt.Printf("%s leveled up to level %d!\n", c.Name, c.Level)
}

func (c *Character) GainExp(exp int) {
	c.Exp += exp
	fmt.Printf("%s gained %d experience points.\n", c.Name, exp)
	for c.Exp >= c.Level*100 {
		c.Exp -= c.Level * 100
		c.LevelUp()
	}
}

func (c *Character) AttackMonster(m *Monster) int {
	damage := c.Attack - m.Defense
	if damage < 1 {
		damage = 1
	}
	m.Health -= damage
	return damage
}

func (m *Monster) AttackCharacter(c *Character) int {
	damage := m.Attack - c.Defense
	if damage < 1 {
		damage = 1
	}
	c.Health -= damage
	return damage
}

func (g *GameState) InitializeGame() {
	g.Player = Character{
		Name:      "Hero",
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
	}
	g.Monsters = []Monster{
		{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, ExpReward: 20},
		{Name: "Orc", Health: 50, Attack: 8, Defense: 4, ExpReward: 40},
		{Name: "Dragon", Health: 100, Attack: 15, Defense: 10, ExpReward: 100},
	}
	g.Inventory = []Item{
		{Name: "Health Potion", Description: "Restores 50 health", Value: 50},
		{Name: "Sword", Description: "Increases attack by 5", Value: 100},
	}
	g.Gold = 0
	g.GameRound = 1
}

func (g *GameState) DisplayStatus() {
	fmt.Printf("=== Game Round %d ===\n", g.GameRound)
	fmt.Printf("Player: %s (Level %d)\n", g.Player.Name, g.Player.Level)
	fmt.Printf("Health: %d/%d\n", g.Player.Health, g.Player.MaxHealth)
	fmt.Printf("Attack: %d, Defense: %d\n", g.Player.Attack, g.Player.Defense)
	fmt.Printf("Experience: %d/%d\n", g.Player.Exp, g.Player.Level*100)
	fmt.Printf("Gold: %d\n", g.Gold)
	fmt.Printf("Inventory: %d items\n", len(g.Inventory))
	for i, item := range g.Inventory {
		fmt.Printf("  %d. %s - %s (Value: %d)\n", i+1, item.Name, item.Description, item.Value)
	}
	fmt.Printf("Monsters remaining: %d\n", len(g.Monsters))
	for i, monster := range g.Monsters {
		fmt.Printf("  %d. %s (Health: %d, Attack: %d, Defense: %d)\n", i+1, monster.Name, monster.Health, monster.Attack, monster.Defense)
	}
}

func (g *GameState) Battle(monsterIndex int) {
	if monsterIndex < 0 || monsterIndex >= len(g.Monsters) {
		fmt.Println("Invalid monster index.")
		return
	}
	monster := &g.Monsters[monsterIndex]
	fmt.Printf("Battle started against %s!\n", monster.Name)
	for g.Player.Health > 0 && monster.Health > 0 {
		playerDamage := g.Player.AttackMonster(monster)
		fmt.Printf("%s attacks %s for %d damage. %s's health: %d\n", g.Player.Name, monster.Name, playerDamage, monster.Name, monster.Health)
		if monster.Health <= 0 {
			fmt.Printf("%s defeated!\n", monster.Name)
			g.Player.GainExp(monster.ExpReward)
			g.Gold += rand.Intn(20) + 10
			g.Monsters = append(g.Monsters[:monsterIndex], g.Monsters[monsterIndex+1:]...)
			break
		}
		monsterDamage := monster.AttackCharacter(&g.Player)
		fmt.Printf("%s attacks %s for %d damage. %s's health: %d\n", monster.Name, g.Player.Name, monsterDamage, g.Player.Name, g.Player.Health)
		if g.Player.Health <= 0 {
			fmt.Printf("%s has been defeated! Game over.\n", g.Player.Name)
			return
		}
		time.Sleep(1 * time.Second)
	}
}

func (g *GameState) UseItem(itemIndex int) {
	if itemIndex < 0 || itemIndex >= len(g.Inventory) {
		fmt.Println("Invalid item index.")
		return
	}
	item := g.Inventory[itemIndex]
	switch item.Name {
	case "Health Potion":
		g.Player.Health += 50
		if g.Player.Health > g.Player.MaxHealth {
			g.Player.Health = g.Player.MaxHealth
		}
		fmt.Printf("Used %s. Health restored to %d.\n", item.Name, g.Player.Health)
		g.Inventory = append(g.Inventory[:itemIndex], g.Inventory[itemIndex+1:]...)
	case "Sword":
		g.Player.Attack += 5
		fmt.Printf("Used %s. Attack increased to %d.\n", item.Name, g.Player.Attack)
		g.Inventory = append(g.Inventory[:itemIndex], g.Inventory[itemIndex+1:]...)
	default:
		fmt.Printf("Cannot use %s.\n", item.Name)
	}
}

func (g *GameState) Shop() {
	fmt.Println("Welcome to the shop!")
	shopItems := []Item{
		{Name: "Health Potion", Description: "Restores 50 health", Value: 50},
		{Name: "Sword", Description: "Increases attack by 5", Value: 100},
		{Name: "Shield", Description: "Increases defense by 5", Value: 150},
	}
	for i, item := range shopItems {
		fmt.Printf("%d. %s - %s (Cost: %d gold)\n", i+1, item.Name, item.Description, item.Value)
	}
	fmt.Printf("Your gold: %d\n", g.Gold)
	fmt.Print("Enter item number to buy (0 to exit): ")
	var choice int
	fmt.Scan(&choice)
	if choice == 0 {
		return
	}
	if choice < 1 || choice > len(shopItems) {
		fmt.Println("Invalid choice.")
		return
	}
	item := shopItems[choice-1]
	if g.Gold >= item.Value {
		g.Gold -= item.Value
		g.Inventory = append(g.Inventory, item)
		fmt.Printf("Bought %s. Remaining gold: %d\n", item.Name, g.Gold)
	} else {
		fmt.Println("Not enough gold.")
	}
}

func (g *GameState) NextRound() {
	g.GameRound++
	if g.GameRound%5 == 0 {
		newMonsters := []Monster{
			{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, ExpReward: 20},
			{Name: "Orc", Health: 50, Attack: 8, Defense: 4, ExpReward: 40},
			{Name: "Dragon", Health: 100, Attack: 15, Defense: 10, ExpReward: 100},
		}
		g.Monsters = append(g.Monsters, newMonsters...)
		fmt.Println("New monsters have appeared!")
	}
	g.Player.Health = g.Player.MaxHealth
	fmt.Printf("Advanced to round %d. Health restored.\n", g.GameRound)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := GameState{}
	game.InitializeGame()
	fmt.Println("Welcome to the Complex Game Simulator!")
	for {
		game.DisplayStatus()
		fmt.Println("Options:")
		fmt.Println("1. Battle a monster")
		fmt.Println("2. Use an item")
		fmt.Println("3. Visit shop")
		fmt.Println("4. Next round")
		fmt.Println("5. Quit")
		fmt.Print("Enter your choice: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			if len(game.Monsters) == 0 {
				fmt.Println("No monsters to battle.")
				continue
			}
			fmt.Print("Enter monster number: ")
			var monsterIndex int
			fmt.Scan(&monsterIndex)
			game.Battle(monsterIndex - 1)
		case 2:
			if len(game.Inventory) == 0 {
				fmt.Println("No items to use.")
				continue
			}
			fmt.Print("Enter item number: ")
			var itemIndex int
			fmt.Scan(&itemIndex)
			game.UseItem(itemIndex - 1)
		case 3:
			game.Shop()
		case 4:
			game.NextRound()
		case 5:
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
		if game.Player.Health <= 0 {
			fmt.Println("Game over!")
			break
		}
		if len(game.Monsters) == 0 && game.GameRound >= 10 {
			fmt.Println("Congratulations! You've cleared all rounds!")
			break
		}
	}
}
