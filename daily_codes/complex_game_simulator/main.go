package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"strings"
	"bufio"
	"os"
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
	Name      string
	Health    int
	Attack    int
	Defense   int
	Loot      []string
}

type GameState struct {
	Player    Character
	Enemies   []Enemy
	Locations []Location
	CurrentLocation int
}

type Location struct {
	Name        string
	Description string
	Enemies     []int
	Items       []string
}

func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.Attack += 2
	c.Defense += 1
	fmt.Printf("Level up! %s is now level %d. Stats improved.\n", c.Name, c.Level)
}

func (c *Character) GainExperience(exp int) {
	c.Experience += exp
	fmt.Printf("Gained %d experience. Total: %d/%d\n", exp, c.Experience, c.Level*100)
	if c.Experience >= c.Level*100 {
		c.Experience -= c.Level * 100
		c.LevelUp()
	}
}

func (c *Character) AttackEnemy(e *Enemy) int {
	damage := c.Attack - e.Defense
	if damage < 1 {
		damage = 1
	}
	e.Health -= damage
	return damage
}

func (e *Enemy) AttackPlayer(c *Character) int {
	damage := e.Attack - c.Defense
	if damage < 1 {
		damage = 1
	}
	c.Health -= damage
	return damage
}

func (c *Character) Heal(amount int) {
	c.Health += amount
	if c.Health > c.MaxHealth {
		c.Health = c.MaxHealth
	}
	fmt.Printf("Healed for %d. Current health: %d/%d\n", amount, c.Health, c.MaxHealth)
}

func (c *Character) AddItem(item string) {
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("Added %s to inventory.\n", item)
}

func (c *Character) UseItem(item string) bool {
	for i, invItem := range c.Inventory {
		if invItem == item {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			if item == "Health Potion" {
				c.Heal(30)
			} else if item == "Attack Boost" {
				c.Attack += 5
				fmt.Printf("Attack increased by 5! Current attack: %d\n", c.Attack)
			}
			return true
		}
	}
	return false
}

func (g *GameState) MoveToLocation(locationIndex int) {
	if locationIndex < 0 || locationIndex >= len(g.Locations) {
		fmt.Println("Invalid location.")
		return
	}
	g.CurrentLocation = locationIndex
	loc := g.Locations[locationIndex]
	fmt.Printf("Moved to %s. %s\n", loc.Name, loc.Description)
}

func (g *GameState) ExploreLocation() {
	loc := g.Locations[g.CurrentLocation]
	if len(loc.Enemies) > 0 {
		enemyIndex := loc.Enemies[0]
		if enemyIndex < len(g.Enemies) {
			enemy := g.Enemies[enemyIndex]
			fmt.Printf("You encounter a %s!\n", enemy.Name)
			g.Combat(enemyIndex)
		}
	}
	if len(loc.Items) > 0 {
		for _, item := range loc.Items {
			g.Player.AddItem(item)
		}
		loc.Items = []string{}
	}
}

func (g *GameState) Combat(enemyIndex int) {
	player := &g.Player
	enemy := &g.Enemies[enemyIndex]
	fmt.Printf("Combat started! %s vs %s\n", player.Name, enemy.Name)
	for player.Health > 0 && enemy.Health > 0 {
		fmt.Printf("\n%s: %d/%d HP | %s: %d/%d HP\n", player.Name, player.Health, player.MaxHealth, enemy.Name, enemy.Health, enemy.Health)
		fmt.Println("Choose action: (1) Attack (2) Use Item (3) Flee")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			damage := player.AttackEnemy(enemy)
			fmt.Printf("%s attacks %s for %d damage.\n", player.Name, enemy.Name, damage)
		case "2":
			fmt.Println("Inventory:")
			for i, item := range player.Inventory {
				fmt.Printf("%d: %s\n", i+1, item)
			}
			fmt.Print("Enter item number to use: ")
			itemInput, _ := reader.ReadString('\n')
			itemInput = strings.TrimSpace(itemInput)
			if index, err := strconv.Atoi(itemInput); err == nil && index > 0 && index <= len(player.Inventory) {
				item := player.Inventory[index-1]
				if player.UseItem(item) {
					fmt.Printf("Used %s.\n", item)
				} else {
					fmt.Println("Failed to use item.")
				}
			} else {
				fmt.Println("Invalid item selection.")
			}
			continue
		case "3":
			fmt.Println("You fled from combat.")
			return
		default:
			fmt.Println("Invalid action.")
			continue
		}
		if enemy.Health <= 0 {
			fmt.Printf("%s defeated!\n", enemy.Name)
			player.GainExperience(50)
			for _, loot := range enemy.Loot {
				player.AddItem(loot)
			}
			g.RemoveEnemyFromLocation(enemyIndex)
			return
		}
		enemyDamage := enemy.AttackPlayer(player)
		fmt.Printf("%s attacks %s for %d damage.\n", enemy.Name, player.Name, enemyDamage)
		if player.Health <= 0 {
			fmt.Printf("%s has been defeated. Game over.\n", player.Name)
			os.Exit(0)
		}
	}
}

func (g *GameState) RemoveEnemyFromLocation(enemyIndex int) {
	loc := &g.Locations[g.CurrentLocation]
	for i, idx := range loc.Enemies {
		if idx == enemyIndex {
			loc.Enemies = append(loc.Enemies[:i], loc.Enemies[i+1:]...)
			return
		}
	}
}

func (g *GameState) DisplayStatus() {
	player := g.Player
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Experience: %d/%d\n", player.Experience, player.Level*100)
	fmt.Println("Inventory:")
	for _, item := range player.Inventory {
		fmt.Printf(" - %s\n", item)
	}
}

func (g *GameState) DisplayLocations() {
	fmt.Println("Available locations:")
	for i, loc := range g.Locations {
		fmt.Printf("%d: %s\n", i, loc.Name)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := GameState{
		Player: Character{
			Name:      "Hero",
			Health:    100,
			MaxHealth: 100,
			Attack:    10,
			Defense:   5,
			Level:     1,
			Experience: 0,
			Inventory: []string{"Health Potion"},
		},
		Enemies: []Enemy{
			{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Loot: []string{"Gold"}},
			{Name: "Orc", Health: 50, Attack: 8, Defense: 4, Loot: []string{"Health Potion"}},
			{Name: "Dragon", Health: 100, Attack: 15, Defense: 10, Loot: []string{"Dragon Scale", "Attack Boost"}},
		},
		Locations: []Location{
			{Name: "Forest", Description: "A dense forest with tall trees.", Enemies: []int{0}, Items: []string{"Health Potion"}},
			{Name: "Cave", Description: "A dark cave with echoing sounds.", Enemies: []int{1}, Items: []string{"Gold"}},
			{Name: "Mountain", Description: "A high mountain with a treacherous path.", Enemies: []int{2}, Items: []string{"Attack Boost"}},
		},
		CurrentLocation: 0,
	}
	fmt.Println("Welcome to the Complex Game Simulator!")
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Display Status")
		fmt.Println("2. Explore Current Location")
		fmt.Println("3. Move to Another Location")
		fmt.Println("4. Use Item")
		fmt.Println("5. Quit")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			game.DisplayStatus()
		case "2":
			game.ExploreLocation()
		case "3":
			game.DisplayLocations()
			fmt.Print("Enter location number: ")
			locInput, _ := reader.ReadString('\n')
			locInput = strings.TrimSpace(locInput)
			if index, err := strconv.Atoi(locInput); err == nil {
				game.MoveToLocation(index)
			} else {
				fmt.Println("Invalid location number.")
			}
		case "4":
			fmt.Println("Inventory:")
			for i, item := range game.Player.Inventory {
				fmt.Printf("%d: %s\n", i+1, item)
			}
			fmt.Print("Enter item number to use: ")
			itemInput, _ := reader.ReadString('\n')
			itemInput = strings.TrimSpace(itemInput)
			if index, err := strconv.Atoi(itemInput); err == nil && index > 0 && index <= len(game.Player.Inventory) {
				item := game.Player.Inventory[index-1]
				if !game.Player.UseItem(item) {
					fmt.Println("Failed to use item.")
				}
			} else {
				fmt.Println("Invalid item selection.")
			}
		case "5":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid action.")
		}
	}
}
