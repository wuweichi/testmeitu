package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"strings"
	"os"
	"bufio"
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

type Enemy struct {
	Name    string
	Health  int
	Attack  int
	Defense int
	Exp     int
}

type GameWorld struct {
	Player    Character
	Enemies   []Enemy
	Locations []string
	Quests    []Quest
}

type Quest struct {
	Name        string
	Description string
	Completed   bool
	Reward      int
}

func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.Attack += 2
	c.Defense += 1
	fmt.Printf("Level up! %s is now level %d\n", c.Name, c.Level)
}

func (c *Character) GainExp(amount int) {
	c.Exp += amount
	fmt.Printf("Gained %d experience points. Total: %d\n", amount, c.Exp)
	for c.Exp >= c.Level*100 {
		c.Exp -= c.Level * 100
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
	fmt.Printf("%s heals for %d health. Current health: %d\n", c.Name, amount, c.Health)
}

func (e *Enemy) AttackPlayer(c *Character) {
	damage := e.Attack - c.Defense
	if damage < 1 {
		damage = 1
	}
	c.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", e.Name, c.Name, damage)
}

func (g *GameWorld) GenerateEnemies() {
	g.Enemies = []Enemy{
		{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Exp: 20},
		{Name: "Orc", Health: 50, Attack: 8, Defense: 4, Exp: 35},
		{Name: "Dragon", Health: 100, Attack: 15, Defense: 10, Exp: 100},
		{Name: "Skeleton", Health: 25, Attack: 4, Defense: 1, Exp: 15},
		{Name: "Wizard", Health: 40, Attack: 12, Defense: 3, Exp: 50},
		{Name: "Troll", Health: 80, Attack: 10, Defense: 6, Exp: 70},
		{Name: "Ghost", Health: 20, Attack: 3, Defense: 0, Exp: 10},
		{Name: "Knight", Health: 60, Attack: 9, Defense: 7, Exp: 45},
		{Name: "Spider", Health: 15, Attack: 2, Defense: 1, Exp: 8},
		{Name: "Giant", Health: 120, Attack: 18, Defense: 12, Exp: 120},
	}
}

func (g *GameWorld) GenerateLocations() {
	g.Locations = []string{
		"Forest",
		"Cave",
		"Mountain",
		"Desert",
		"Castle",
		"Swamp",
		"Village",
		"Dungeon",
		"Beach",
		"Volcano",
	}
}

func (g *GameWorld) GenerateQuests() {
	g.Quests = []Quest{
		{Name: "Slay the Dragon", Description: "Defeat the mighty dragon in the volcano.", Completed: false, Reward: 200},
		{Name: "Rescue the Villager", Description: "Find and rescue the lost villager in the forest.", Completed: false, Reward: 100},
		{Name: "Collect Artifacts", Description: "Gather 5 ancient artifacts from the dungeon.", Completed: false, Reward: 150},
		{Name: "Defeat the Goblin King", Description: "Conquer the goblin king in the cave.", Completed: false, Reward: 80},
		{Name: "Explore the Castle", Description: "Fully explore the abandoned castle.", Completed: false, Reward: 120},
	}
}

func (g *GameWorld) StartGame() {
	fmt.Println("Welcome to the Complex Game Simulator!")
	fmt.Println("You are about to embark on an epic adventure.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your character's name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	g.Player = Character{
		Name:      name,
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
		Inventory: []string{"Health Potion", "Sword"},
	}
	g.GenerateEnemies()
	g.GenerateLocations()
	g.GenerateQuests()
	fmt.Printf("Hello, %s! Your adventure begins now.\n", g.Player.Name)
}

func (g *GameWorld) DisplayStatus() {
	fmt.Printf("Name: %s\n", g.Player.Name)
	fmt.Printf("Level: %d\n", g.Player.Level)
	fmt.Printf("Health: %d/%d\n", g.Player.Health, g.Player.MaxHealth)
	fmt.Printf("Attack: %d\n", g.Player.Attack)
	fmt.Printf("Defense: %d\n", g.Player.Defense)
	fmt.Printf("Experience: %d\n", g.Player.Exp)
	fmt.Println("Inventory:", g.Player.Inventory)
}

func (g *GameWorld) DisplayQuests() {
	fmt.Println("Active Quests:")
	for i, quest := range g.Quests {
		status := "Incomplete"
		if quest.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s - %s [%s]\n", i+1, quest.Name, quest.Description, status)
	}
}

func (g *GameWorld) DisplayLocations() {
	fmt.Println("Available Locations:")
	for i, location := range g.Locations {
		fmt.Printf("%d. %s\n", i+1, location)
	}
}

func (g *GameWorld) ExploreLocation(locationIndex int) {
	if locationIndex < 0 || locationIndex >= len(g.Locations) {
		fmt.Println("Invalid location index.")
		return
	}
	location := g.Locations[locationIndex]
	fmt.Printf("You are exploring the %s.\n", location)
	rand.Seed(time.Now().UnixNano())
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Println("You found a treasure chest!")
		g.Player.Inventory = append(g.Player.Inventory, "Gold Coin")
		fmt.Println("Added Gold Coin to inventory.")
	case 1:
		enemyIndex := rand.Intn(len(g.Enemies))
		enemy := g.Enemies[enemyIndex]
		fmt.Printf("A wild %s appears!\n", enemy.Name)
		g.Combat(&enemy)
	case 2:
		fmt.Println("You discover a hidden path but find nothing of interest.")
	}
}

func (g *GameWorld) Combat(enemy *Enemy) {
	fmt.Printf("Combat started with %s!\n", enemy.Name)
	for g.Player.Health > 0 && enemy.Health > 0 {
		fmt.Println("\n--- Your Turn ---")
		fmt.Println("1. Attack")
		fmt.Println("2. Use Health Potion")
		fmt.Println("3. Flee")
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Choose an action: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			g.Player.AttackEnemy(enemy)
		case "2":
			if contains(g.Player.Inventory, "Health Potion") {
				g.Player.Heal(30)
				g.Player.Inventory = removeItem(g.Player.Inventory, "Health Potion")
			} else {
				fmt.Println("No Health Potion in inventory!")
			}
		case "3":
			fmt.Println("You fled from the battle!")
			return
		default:
			fmt.Println("Invalid action. Try again.")
			continue
		}
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			g.Player.GainExp(enemy.Exp)
			return
		}
		fmt.Println("\n--- Enemy Turn ---")
		enemy.AttackPlayer(&g.Player)
		if g.Player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			os.Exit(0)
		}
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

func removeItem(slice []string, item string) []string {
	for i, s := range slice {
		if s == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

func (g *GameWorld) CompleteQuest(questIndex int) {
	if questIndex < 0 || questIndex >= len(g.Quests) {
		fmt.Println("Invalid quest index.")
		return
	}
	quest := &g.Quests[questIndex]
	if quest.Completed {
		fmt.Println("This quest is already completed.")
		return
	}
	quest.Completed = true
	g.Player.GainExp(quest.Reward)
	fmt.Printf("Quest '%s' completed! Gained %d experience.\n", quest.Name, quest.Reward)
}

func (g *GameWorld) MainMenu() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Display Status")
		fmt.Println("2. Display Quests")
		fmt.Println("3. Display Locations")
		fmt.Println("4. Explore Location")
		fmt.Println("5. Complete Quest")
		fmt.Println("6. Exit Game")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			g.DisplayStatus()
		case "2":
			g.DisplayQuests()
		case "3":
			g.DisplayLocations()
		case "4":
			g.DisplayLocations()
			fmt.Print("Enter location number to explore: ")
			locInput, _ := reader.ReadString('\n')
			locInput = strings.TrimSpace(locInput)
			locIndex, err := strconv.Atoi(locInput)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			g.ExploreLocation(locIndex - 1)
		case "5":
			g.DisplayQuests()
			fmt.Print("Enter quest number to complete: ")
			questInput, _ := reader.ReadString('\n')
			questInput = strings.TrimSpace(questInput)
			questIndex, err := strconv.Atoi(questInput)
			if err != nil {
				fmt.Println("Invalid input. Please enter a number.")
				continue
			}
			g.CompleteQuest(questIndex - 1)
		case "6":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func main() {
	game := GameWorld{}
	game.StartGame()
	game.MainMenu()
}
