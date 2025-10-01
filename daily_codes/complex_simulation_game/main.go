package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
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
	Gold      int
}

type Monster struct {
	Name      string
	Health    int
	MaxHealth int
	Attack    int
	Defense   int
	ExpReward int
	GoldReward int
}

type GameState struct {
	Player     Character
	Monsters   []Monster
	Locations  []string
	CurrentLocation int
	GameOver   bool
}

func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.Attack += 2
	c.Defense += 1
	fmt.Printf("Level up! You are now level %d\n", c.Level)
}

func (c *Character) GainExp(exp int) {
	c.Exp += exp
	fmt.Printf("Gained %d experience points.\n", exp)
	if c.Exp >= c.Level*100 {
		c.LevelUp()
		c.Exp = 0
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

func (c *Character) IsAlive() bool {
	return c.Health > 0
}

func (m *Monster) IsAlive() bool {
	return m.Health > 0
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
			return true
		}
	}
	return false
}

func (c *Character) ShowStatus() {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Level: %d\n", c.Level)
	fmt.Printf("Health: %d/%d\n", c.Health, c.MaxHealth)
	fmt.Printf("Attack: %d\n", c.Attack)
	fmt.Printf("Defense: %d\n", c.Defense)
	fmt.Printf("Experience: %d/%d\n", c.Exp, c.Level*100)
	fmt.Printf("Gold: %d\n", c.Gold)
	fmt.Printf("Inventory: %v\n", c.Inventory)
}

func (gs *GameState) SaveGame(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	return encoder.Encode(gs)
}

func (gs *GameState) LoadGame(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	return decoder.Decode(gs)
}

func (gs *GameState) GenerateMonster() Monster {
	monsterTypes := []string{"Goblin", "Orc", "Dragon", "Slime", "Skeleton"}
	name := monsterTypes[rand.Intn(len(monsterTypes))]
	level := gs.Player.Level + rand.Intn(3) - 1
	if level < 1 {
		level = 1
	}
	return Monster{
		Name:      name,
		Health:    20 + level*10,
		MaxHealth: 20 + level*10,
		Attack:    5 + level*2,
		Defense:   2 + level,
		ExpReward: 10 + level*5,
		GoldReward: 5 + level*3,
	}
}

func (gs *GameState) Battle() {
	monster := gs.GenerateMonster()
	fmt.Printf("A wild %s appears!\n", monster.Name)
	fmt.Printf("Monster stats - Health: %d/%d, Attack: %d, Defense: %d\n", 
		monster.Health, monster.MaxHealth, monster.Attack, monster.Defense)

	for gs.Player.IsAlive() && monster.IsAlive() {
		fmt.Println("\n--- Battle Menu ---")
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Flee")
		fmt.Print("Choose an option: ")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			damage := gs.Player.AttackMonster(&monster)
			fmt.Printf("You attack the %s for %d damage!\n", monster.Name, damage)
			if monster.IsAlive() {
				damage = monster.AttackCharacter(&gs.Player)
				fmt.Printf("The %s attacks you for %d damage!\n", monster.Name, damage)
			}
		case "2":
			if len(gs.Player.Inventory) == 0 {
				fmt.Println("Your inventory is empty!")
				continue
			}
			fmt.Println("Your inventory:")
			for i, item := range gs.Player.Inventory {
				fmt.Printf("%d. %s\n", i+1, item)
			}
			fmt.Print("Choose an item to use: ")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			index, err := strconv.Atoi(input)
			if err != nil || index < 1 || index > len(gs.Player.Inventory) {
				fmt.Println("Invalid choice!")
				continue
			}
			item := gs.Player.Inventory[index-1]
			if item == "Health Potion" {
				if gs.Player.UseItem(item) {
					gs.Player.Heal(30)
				} else {
					fmt.Println("Failed to use item!")
				}
			} else {
				fmt.Printf("Cannot use %s in battle!\n", item)
			}
		case "3":
			if rand.Float32() < 0.7 {
				fmt.Println("You successfully fled from the battle!")
				return
			} else {
				fmt.Println("You failed to flee!")
				damage := monster.AttackCharacter(&gs.Player)
				fmt.Printf("The %s attacks you for %d damage!\n", monster.Name, damage)
			}
		default:
			fmt.Println("Invalid option!")
		}

		fmt.Printf("Your health: %d/%d, Monster health: %d/%d\n", 
			gs.Player.Health, gs.Player.MaxHealth, monster.Health, monster.MaxHealth)
	}

	if gs.Player.IsAlive() {
		fmt.Printf("\nYou defeated the %s!\n", monster.Name)
		gs.Player.GainExp(monster.ExpReward)
		gs.Player.Gold += monster.GoldReward
		fmt.Printf("You gained %d gold!\n", monster.GoldReward)
		
		// Chance to find item
		if rand.Float32() < 0.3 {
			items := []string{"Health Potion", "Magic Scroll", "Ancient Coin"}
			item := items[rand.Intn(len(items))]
			gs.Player.AddItem(item)
		}
	} else {
		fmt.Println("\nYou have been defeated!")
		gs.GameOver = true
	}
}

func (gs *GameState) Explore() {
	fmt.Printf("\nYou are at: %s\n", gs.Locations[gs.CurrentLocation])
	fmt.Println("What would you like to do?")
	fmt.Println("1. Look for monsters")
	fmt.Println("2. Move to another location")
	fmt.Println("3. Rest")
	fmt.Println("4. Check status")
	fmt.Println("5. Save game")
	fmt.Println("6. Load game")
	fmt.Println("7. Quit")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	switch input {
	case "1":
		if rand.Float32() < 0.6 {
			gs.Battle()
		} else {
			fmt.Println("You search the area but find nothing interesting.")
		}
	case "2":
		fmt.Println("Available locations:")
		for i, location := range gs.Locations {
			fmt.Printf("%d. %s\n", i+1, location)
		}
		fmt.Print("Choose a location: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		index, err := strconv.Atoi(input)
		if err != nil || index < 1 || index > len(gs.Locations) {
			fmt.Println("Invalid location!")
			return
		}
		gs.CurrentLocation = index - 1
		fmt.Printf("You travel to %s\n", gs.Locations[gs.CurrentLocation])
	case "3":
		gs.Player.Heal(gs.Player.MaxHealth / 4)
		fmt.Println("You rest and recover some health.")
	case "4":
		gs.Player.ShowStatus()
	case "5":
		fmt.Print("Enter save file name: ")
		filename, _ := reader.ReadString('\n')
		filename = strings.TrimSpace(filename)
		if err := gs.SaveGame(filename); err != nil {
			fmt.Printf("Error saving game: %v\n", err)
		} else {
			fmt.Println("Game saved successfully!")
		}
	case "6":
		fmt.Print("Enter save file name: ")
		filename, _ := reader.ReadString('\n')
		filename = strings.TrimSpace(filename)
		if err := gs.LoadGame(filename); err != nil {
			fmt.Printf("Error loading game: %v\n", err)
		} else {
			fmt.Println("Game loaded successfully!")
		}
	case "7":
		gs.GameOver = true
	default:
		fmt.Println("Invalid option!")
	}
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create initial game state
	gameState := GameState{
		Player: Character{
			Name:      "Hero",
			Health:    100,
			MaxHealth: 100,
			Attack:    10,
			Defense:   5,
			Level:     1,
			Exp:       0,
			Inventory: []string{"Health Potion", "Health Potion"},
			Gold:      50,
		},
		Locations: []string{
			"Forest",
			"Mountains", 
			"Desert",
			"Swamp",
			"Ancient Ruins",
			"Crystal Cave",
			"Volcano",
			"Frozen Tundra",
			"Enchanted Garden",
			"Abandoned Castle",
		},
		CurrentLocation: 0,
		GameOver:        false,
	}

	fmt.Println("Welcome to the Complex Simulation Game!")
	fmt.Println("You are a brave adventurer exploring a dangerous world.")
	fmt.Println("Defeat monsters, gain experience, and collect treasures!")
	fmt.Println()

	// Game loop
	for !gameState.GameOver {
		gameState.Explore()
		
		// Check for win condition
		if gameState.Player.Level >= 10 {
			fmt.Println("\nCongratulations! You have reached level 10 and become a legendary hero!")
			fmt.Println("You win the game!")
			gameState.GameOver = true
		}
	}

	fmt.Println("\nThanks for playing!")
	fmt.Printf("Final stats - Level: %d, Gold: %d, Inventory: %v\n", 
		gameState.Player.Level, gameState.Player.Gold, gameState.Player.Inventory)
}