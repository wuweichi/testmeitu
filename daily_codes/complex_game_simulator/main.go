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
	Experience int
	Inventory []string
}

type Enemy struct {
	Name    string
	Health  int
	Attack  int
	Defense int
}

func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.Attack += 2
	c.Defense += 1
	fmt.Printf("%s leveled up to level %d!\n", c.Name, c.Level)
}

func (c *Character) GainExperience(exp int) {
	c.Experience += exp
	fmt.Printf("%s gained %d experience points.\n", c.Name, exp)
	for c.Experience >= c.Level*100 {
		c.Experience -= c.Level * 100
		c.LevelUp()
	}
}

func (c *Character) AttackEnemy(e *Enemy) {
	damage := c.Attack - e.Defense
	if damage < 1 {
		damage = 1
	}
	e.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", c.Name, e.Name, damage)
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
	fmt.Printf("%s added %s to inventory.\n", c.Name, item)
}

func (c *Character) UseItem(item string) bool {
	for i, invItem := range c.Inventory {
		if invItem == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Printf("%s used %s.\n", c.Name, item)
			return true
		}
	}
	fmt.Printf("%s does not have %s in inventory.\n", c.Name, item)
	return false
}

func (e *Enemy) AttackCharacter(c *Character) {
	damage := e.Attack - c.Defense
	if damage < 1 {
		damage = 1
	}
	c.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", e.Name, c.Name, damage)
}

func (e *Enemy) IsAlive() bool {
	return e.Health > 0
}

func (c *Character) IsAlive() bool {
	return c.Health > 0
}

func GenerateEnemy(level int) *Enemy {
	names := []string{"Goblin", "Orc", "Troll", "Dragon", "Skeleton", "Zombie"}
	name := names[rand.Intn(len(names))]
	health := 20 + (level * 5)
	attack := 5 + (level * 2)
	defense := 2 + level
	return &Enemy{Name: name, Health: health, Attack: attack, Defense: defense}
}

func Battle(c *Character, e *Enemy) {
	fmt.Printf("A wild %s appears!\n", e.Name)
	for c.IsAlive() && e.IsAlive() {
		fmt.Println("\n--- Battle Menu ---")
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Flee")
		var choice string
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			c.AttackEnemy(e)
			if e.IsAlive() {
				e.AttackCharacter(c)
			}
		case "2":
			if len(c.Inventory) > 0 {
				fmt.Println("Inventory:")
				for i, item := range c.Inventory {
					fmt.Printf("%d. %s\n", i+1, item)
				}
				fmt.Print("Choose item to use: ")
				var itemChoice string
				fmt.Scanln(&itemChoice)
				index, err := strconv.Atoi(itemChoice)
				if err == nil && index > 0 && index <= len(c.Inventory) {
					item := c.Inventory[index-1]
					if item == "Health Potion" {
						c.Heal(20)
						c.UseItem(item)
					} else {
						fmt.Println("Cannot use that item in battle.")
					}
				} else {
					fmt.Println("Invalid choice.")
				}
			} else {
				fmt.Println("Inventory is empty.")
			}
			if e.IsAlive() {
				e.AttackCharacter(c)
			}
		case "3":
			if rand.Float32() < 0.5 {
				fmt.Println("You successfully fled!")
				return
			} else {
				fmt.Println("You failed to flee!")
				e.AttackCharacter(c)
			}
		default:
			fmt.Println("Invalid choice.")
		}
		fmt.Printf("%s Health: %d/%d\n", c.Name, c.Health, c.MaxHealth)
		fmt.Printf("%s Health: %d\n", e.Name, e.Health)
	}
	if !e.IsAlive() {
		exp := 10 + (c.Level * 5)
		c.GainExperience(exp)
		fmt.Printf("You defeated the %s!\n", e.Name)
		if rand.Float32() < 0.3 {
			item := "Health Potion"
			c.AddItem(item)
			fmt.Printf("The %s dropped a %s!\n", e.Name, item)
		}
	}
	if !c.IsAlive() {
		fmt.Printf("%s has been defeated!\n", c.Name)
	}
}

func Explore(c *Character) {
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Println("You found a treasure chest!")
		item := "Health Potion"
		c.AddItem(item)
		fmt.Printf("You found a %s!\n", item)
	case 1:
		fmt.Println("You encountered an enemy!")
		enemy := GenerateEnemy(c.Level)
		Battle(c, enemy)
	case 2:
		fmt.Println("You found a peaceful area and rest.")
		c.Heal(10)
	}
}

func DisplayStatus(c *Character) {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Level: %d\n", c.Level)
	fmt.Printf("Health: %d/%d\n", c.Health, c.MaxHealth)
	fmt.Printf("Attack: %d\n", c.Attack)
	fmt.Printf("Defense: %d\n", c.Defense)
	fmt.Printf("Experience: %d/%d\n", c.Experience, c.Level*100)
	fmt.Printf("Inventory: %s\n", strings.Join(c.Inventory, ", "))
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create a character
	var name string
	fmt.Print("Enter your character's name: ")
	fmt.Scanln(&name)
	player := &Character{
		Name:      name,
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Experience: 0,
		Inventory: []string{"Health Potion"},
	}

	fmt.Printf("Welcome, %s!\n", player.Name)
	fmt.Println("You are about to embark on an adventure!")

	// Main game loop
	for player.IsAlive() {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Explore")
		fmt.Println("2. Display Status")
		fmt.Println("3. Use Item")
		fmt.Println("4. Quit")
		var choice string
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			Explore(player)
		case "2":
			DisplayStatus(player)
		case "3":
			if len(player.Inventory) > 0 {
				fmt.Println("Inventory:")
				for i, item := range player.Inventory {
					fmt.Printf("%d. %s\n", i+1, item)
				}
				fmt.Print("Choose item to use: ")
				var itemChoice string
				fmt.Scanln(&itemChoice)
				index, err := strconv.Atoi(itemChoice)
				if err == nil && index > 0 && index <= len(player.Inventory) {
					item := player.Inventory[index-1]
					if item == "Health Potion" {
						player.Heal(20)
						player.UseItem(item)
					} else {
						fmt.Println("Cannot use that item now.")
					}
				} else {
					fmt.Println("Invalid choice.")
				}
			} else {
				fmt.Println("Inventory is empty.")
			}
		case "4":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
	fmt.Printf("Game Over! %s has been defeated.\n", player.Name)
}
