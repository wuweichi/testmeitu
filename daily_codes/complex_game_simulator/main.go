package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"strings"
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
	Gold      int
}

type Monster struct {
	Name      string
	Health    int
	Attack    int
	Defense   int
	ExpReward int
	GoldReward int
}

type Item struct {
	Name        string
	Description string
	Value       int
	Type        string
}

func (c *Character) DisplayStatus() {
	fmt.Printf("=== %s's Status ===\n", c.Name)
	fmt.Printf("Level: %d\n", c.Level)
	fmt.Printf("Health: %d/%d\n", c.Health, c.MaxHealth)
	fmt.Printf("Attack: %d\n", c.Attack)
	fmt.Printf("Defense: %d\n", c.Defense)
	fmt.Printf("Experience: %d/%d\n", c.Exp, c.Level*100)
	fmt.Printf("Gold: %d\n", c.Gold)
	fmt.Printf("Inventory: %v\n", c.Inventory)
	fmt.Println("===================")
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

func (c *Character) GainExp(exp int) {
	c.Exp += exp
	for c.Exp >= c.Level*100 {
		c.LevelUp()
	}
}

func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.Attack += 2
	c.Defense += 1
	c.Exp = 0
	fmt.Printf("🎉 %s leveled up to level %d!\n", c.Name, c.Level)
}

func (c *Character) AddItem(item string) {
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("📦 Added %s to inventory.\n", item)
}

func (c *Character) UseItem(itemName string) bool {
	for i, item := range c.Inventory {
		if item == itemName {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Printf("✅ Used %s.\n", itemName)
			return true
		}
	}
	fmt.Printf("❌ Item %s not found in inventory.\n", itemName)
	return false
}

func GenerateMonster(level int) Monster {
	monsters := []Monster{
		{Name: "Goblin", Health: 20 + level*5, Attack: 5 + level, Defense: 2 + level/2, ExpReward: 10 + level*2, GoldReward: 5 + level},
		{Name: "Orc", Health: 30 + level*8, Attack: 8 + level*2, Defense: 5 + level, ExpReward: 20 + level*3, GoldReward: 10 + level*2},
		{Name: "Dragon", Health: 50 + level*15, Attack: 15 + level*3, Defense: 10 + level*2, ExpReward: 50 + level*5, GoldReward: 25 + level*3},
	}
	return monsters[rand.Intn(len(monsters))]
}

func Battle(c *Character, m *Monster) bool {
	fmt.Printf("⚔️  A wild %s appears!\n", m.Name)
	for c.Health > 0 && m.Health > 0 {
		fmt.Printf("\n%s: %d HP | %s: %d HP\n", c.Name, c.Health, m.Name, m.Health)
		fmt.Print("Choose action: (1) Attack (2) Use Item (3) Flee: ")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			damage := c.AttackMonster(m)
			fmt.Printf("💥 %s attacks %s for %d damage!\n", c.Name, m.Name, damage)
			if m.Health <= 0 {
				fmt.Printf("🎊 %s defeated the %s!\n", c.Name, m.Name)
				c.GainExp(m.ExpReward)
				c.Gold += m.GoldReward
				fmt.Printf("💰 Gained %d gold.\n", m.GoldReward)
				return true
			}
			mDamage := m.AttackCharacter(c)
			fmt.Printf("💢 %s attacks %s for %d damage!\n", m.Name, c.Name, mDamage)
			if c.Health <= 0 {
				fmt.Printf("💀 %s has been defeated...\n", c.Name)
				return false
			}
		case 2:
			if len(c.Inventory) == 0 {
				fmt.Println("❌ No items in inventory.")
				continue
			}
			fmt.Printf("Inventory: %v\n", c.Inventory)
			fmt.Print("Enter item name to use: ")
			var item string
			fmt.Scan(&item)
			if c.UseItem(item) {
				if item == "Health Potion" {
					healAmount := 20
					c.Health += healAmount
					if c.Health > c.MaxHealth {
						c.Health = c.MaxHealth
					}
					fmt.Printf("❤️  Restored %d health. Current health: %d\n", healAmount, c.Health)
				}
			}
		case 3:
			fmt.Println("🏃 Fleeing from battle...")
			return false
		default:
			fmt.Println("❌ Invalid choice.")
		}
	}
	return false
}

func Shop(c *Character) {
	items := map[string]int{
		"Health Potion": 10,
		"Attack Boost": 20,
		"Defense Boost": 15,
	}
	fmt.Println("🏪 Welcome to the Shop!")
	for {
		fmt.Println("Available items:")
		for item, price := range items {
			fmt.Printf("- %s: %d gold\n", item, price)
		}
		fmt.Printf("Your gold: %d\n", c.Gold)
		fmt.Print("Enter item name to buy (or 'exit' to leave): ")
		var choice string
		fmt.Scan(&choice)
		if strings.ToLower(choice) == "exit" {
			break
		}
		if price, exists := items[choice]; exists {
			if c.Gold >= price {
				c.Gold -= price
				c.AddItem(choice)
				fmt.Printf("✅ Purchased %s for %d gold.\n", choice, price)
			} else {
				fmt.Println("❌ Not enough gold.")
			}
		} else {
			fmt.Println("❌ Item not available.")
		}
	}
}

func Explore(c *Character) {
	events := []string{
		"You find a treasure chest!",
		"You discover a hidden path.",
		"You encounter a mysterious stranger.",
		"You stumble upon an ancient ruin.",
	}
	event := events[rand.Intn(len(events))]
	fmt.Println(event)
	switch rand.Intn(4) {
	case 0:
		goldFound := rand.Intn(20) + 5
		c.Gold += goldFound
		fmt.Printf("💰 Found %d gold!\n", goldFound)
	case 1:
		item := "Health Potion"
		c.AddItem(item)
		fmt.Printf("📦 Found a %s!\n", item)
	case 2:
		monster := GenerateMonster(c.Level)
		Battle(c, &monster)
	case 3:
		fmt.Println("Nothing interesting happens.")
	}
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Welcome message
	fmt.Println("🎮 Welcome to the Complex Game Simulator!")
	fmt.Println("Create your character:")

	// Character creation
	var name string
	fmt.Print("Enter character name: ")
	fmt.Scan(&name)

	player := Character{
		Name:      name,
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
		Inventory: []string{"Health Potion"},
		Gold:      50,
	}

	fmt.Printf("🌟 Character %s created!\n", player.Name)
	player.DisplayStatus()

	// Main game loop
	for {
		fmt.Println("\n=== Main Menu ===")
		fmt.Println("1. Explore")
		fmt.Println("2. Battle")
		fmt.Println("3. Shop")
		fmt.Println("4. Status")
		fmt.Println("5. Quit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			Explore(&player)
		case 2:
			monster := GenerateMonster(player.Level)
			Battle(&player, &monster)
		case 3:
			Shop(&player)
		case 4:
			player.DisplayStatus()
		case 5:
			fmt.Println("👋 Thanks for playing!")
			return
		default:
			fmt.Println("❌ Invalid choice.")
		}

		// Check if player is dead
		if player.Health <= 0 {
			fmt.Println("\n💀 Game Over!")
			fmt.Print("Play again? (y/n): ")
			var restart string
			fmt.Scan(&restart)
			if strings.ToLower(restart) == "y" {
				main()
			}
			return
		}
	}
}
