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
	Experience int
	Inventory []string
}

type Enemy struct {
	Name    string
	Health  int
	Attack  int
	Defense int
}

type GameWorld struct {
	Player    Character
	Enemies   []Enemy
	Locations []string
	CurrentLocation int
}

func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.Attack += 2
	c.Defense += 1
	fmt.Printf("Congratulations! %s leveled up to level %d!\n", c.Name, c.Level)
}

func (c *Character) GainExperience(exp int) {
	c.Experience += exp
	fmt.Printf("%s gained %d experience points. Total: %d\n", c.Name, exp, c.Experience)
	for c.Experience >= c.Level*100 {
		c.LevelUp()
	}
}

func (c *Character) AddItem(item string) {
	c.Inventory = append(c.Inventory, item)
	fmt.Printf("Added %s to inventory.\n", item)
}

func (c *Character) UseItem(itemName string) bool {
	for i, item := range c.Inventory {
		if item == itemName {
			c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
			fmt.Printf("Used %s.\n", itemName)
			return true
		}
	}
	fmt.Printf("Item %s not found in inventory.\n", itemName)
	return false
}

func (c *Character) DisplayStatus() {
	fmt.Printf("Name: %s\n", c.Name)
	fmt.Printf("Level: %d\n", c.Level)
	fmt.Printf("Health: %d/%d\n", c.Health, c.MaxHealth)
	fmt.Printf("Attack: %d\n", c.Attack)
	fmt.Printf("Defense: %d\n", c.Defense)
	fmt.Printf("Experience: %d\n", c.Experience)
	fmt.Printf("Inventory: %v\n", c.Inventory)
}

func (e *Enemy) DisplayStatus() {
	fmt.Printf("Enemy: %s\n", e.Name)
	fmt.Printf("Health: %d\n", e.Health)
	fmt.Printf("Attack: %d\n", e.Attack)
	fmt.Printf("Defense: %d\n", e.Defense)
}

func (gw *GameWorld) DisplayLocation() {
	fmt.Printf("Current Location: %s\n", gw.Locations[gw.CurrentLocation])
}

func (gw *GameWorld) MoveToLocation(locationIndex int) {
	if locationIndex >= 0 && locationIndex < len(gw.Locations) {
		gw.CurrentLocation = locationIndex
		fmt.Printf("Moved to %s.\n", gw.Locations[locationIndex])
	} else {
		fmt.Println("Invalid location.")
	}
}

func (gw *GameWorld) GenerateEnemies() {
	gw.Enemies = []Enemy{}
	rand.Seed(time.Now().UnixNano())
	numEnemies := rand.Intn(3) + 1
	enemyNames := []string{"Goblin", "Orc", "Skeleton", "Zombie", "Dragon"}
	for i := 0; i < numEnemies; i++ {
		name := enemyNames[rand.Intn(len(enemyNames))]
		health := rand.Intn(50) + 20
		attack := rand.Intn(10) + 5
		defense := rand.Intn(5) + 1
		gw.Enemies = append(gw.Enemies, Enemy{Name: name, Health: health, Attack: attack, Defense: defense})
	}
	fmt.Printf("Generated %d enemies in this location.\n", numEnemies)
}

func (gw *GameWorld) Battle() {
	if len(gw.Enemies) == 0 {
		fmt.Println("No enemies to battle.")
		return
	}
	fmt.Println("Battle started!")
	for len(gw.Enemies) > 0 && gw.Player.Health > 0 {
		fmt.Println("\n--- Your Turn ---")
		gw.Player.DisplayStatus()
		fmt.Println("Enemies:")
		for i, enemy := range gw.Enemies {
			fmt.Printf("%d: ", i)
			enemy.DisplayStatus()
		}
		fmt.Print("Choose an enemy to attack (by index) or 'run' to escape: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "run" {
			fmt.Println("You escaped from the battle!")
			return
		}
		index, err := strconv.Atoi(input)
		if err != nil || index < 0 || index >= len(gw.Enemies) {
			fmt.Println("Invalid enemy index.")
			continue
		}
		playerDamage := gw.Player.Attack - gw.Enemies[index].Defense
		if playerDamage < 0 {
			playerDamage = 0
		}
		gw.Enemies[index].Health -= playerDamage
		fmt.Printf("You attacked %s for %d damage.\n", gw.Enemies[index].Name, playerDamage)
		if gw.Enemies[index].Health <= 0 {
			fmt.Printf("%s defeated!\n", gw.Enemies[index].Name)
			gw.Player.GainExperience(20)
			gw.Enemies = append(gw.Enemies[:index], gw.Enemies[index+1:]...)
			if len(gw.Enemies) == 0 {
				fmt.Println("All enemies defeated!")
				gw.Player.AddItem("Gold")
				return
			}
		} else {
			fmt.Printf("%s has %d health remaining.\n", gw.Enemies[index].Name, gw.Enemies[index].Health)
		}
		if len(gw.Enemies) > 0 {
			fmt.Println("\n--- Enemy Turn ---")
			for i := 0; i < len(gw.Enemies); i++ {
				enemyDamage := gw.Enemies[i].Attack - gw.Player.Defense
				if enemyDamage < 0 {
					enemyDamage = 0
				}
				gw.Player.Health -= enemyDamage
				fmt.Printf("%s attacked you for %d damage.\n", gw.Enemies[i].Name, enemyDamage)
				if gw.Player.Health <= 0 {
					fmt.Println("You have been defeated! Game over.")
					return
				}
			}
		}
	}
}

func (gw *GameWorld) Explore() {
	fmt.Println("Exploring the area...")
	rand.Seed(time.Now().UnixNano())
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Println("You found a treasure chest!")
		gw.Player.AddItem("Potion")
	case 1:
		fmt.Println("You encountered enemies!")
		gw.GenerateEnemies()
		gw.Battle()
	case 2:
		fmt.Println("It's quiet here. Nothing happened.")
	}
}

func (gw *GameWorld) SaveGame(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "%s\n", gw.Player.Name)
	fmt.Fprintf(writer, "%d\n", gw.Player.Health)
	fmt.Fprintf(writer, "%d\n", gw.Player.MaxHealth)
	fmt.Fprintf(writer, "%d\n", gw.Player.Attack)
	fmt.Fprintf(writer, "%d\n", gw.Player.Defense)
	fmt.Fprintf(writer, "%d\n", gw.Player.Level)
	fmt.Fprintf(writer, "%d\n", gw.Player.Experience)
	fmt.Fprintf(writer, "%d\n", len(gw.Player.Inventory))
	for _, item := range gw.Player.Inventory {
		fmt.Fprintf(writer, "%s\n", item)
	}
	fmt.Fprintf(writer, "%d\n", gw.CurrentLocation)
	writer.Flush()
	fmt.Println("Game saved successfully.")
	return nil
}

func (gw *GameWorld) LoadGame(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	name, _ := reader.ReadString('\n')
	gw.Player.Name = strings.TrimSpace(name)
	healthStr, _ := reader.ReadString('\n')
	gw.Player.Health, _ = strconv.Atoi(strings.TrimSpace(healthStr))
	maxHealthStr, _ := reader.ReadString('\n')
	gw.Player.MaxHealth, _ = strconv.Atoi(strings.TrimSpace(maxHealthStr))
	attackStr, _ := reader.ReadString('\n')
	gw.Player.Attack, _ = strconv.Atoi(strings.TrimSpace(attackStr))
	defenseStr, _ := reader.ReadString('\n')
	gw.Player.Defense, _ = strconv.Atoi(strings.TrimSpace(defenseStr))
	levelStr, _ := reader.ReadString('\n')
	gw.Player.Level, _ = strconv.Atoi(strings.TrimSpace(levelStr))
	expStr, _ := reader.ReadString('\n')
	gw.Player.Experience, _ = strconv.Atoi(strings.TrimSpace(expStr))
	invCountStr, _ := reader.ReadString('\n')
	invCount, _ := strconv.Atoi(strings.TrimSpace(invCountStr))
	gw.Player.Inventory = []string{}
	for i := 0; i < invCount; i++ {
		item, _ := reader.ReadString('\n')
		gw.Player.Inventory = append(gw.Player.Inventory, strings.TrimSpace(item))
	}
	locationStr, _ := reader.ReadString('\n')
	gw.CurrentLocation, _ = strconv.Atoi(strings.TrimSpace(locationStr))
	fmt.Println("Game loaded successfully.")
	return nil
}

func main() {
	fmt.Println("Welcome to the Complex Game Simulator!")
	gw := GameWorld{
		Player: Character{
			Name:      "Hero",
			Health:    100,
			MaxHealth: 100,
			Attack:    10,
			Defense:   5,
			Level:     1,
			Experience: 0,
			Inventory: []string{"Sword", "Shield"},
		},
		Locations:        []string{"Forest", "Cave", "Mountain", "Castle"},
		CurrentLocation: 0,
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Display Player Status")
		fmt.Println("2. Display Current Location")
		fmt.Println("3. Move to Another Location")
		fmt.Println("4. Explore Current Location")
		fmt.Println("5. Battle Enemies")
		fmt.Println("6. Use Item")
		fmt.Println("7. Save Game")
		fmt.Println("8. Load Game")
		fmt.Println("9. Quit")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			gw.Player.DisplayStatus()
		case "2":
			gw.DisplayLocation()
		case "3":
			fmt.Println("Available locations:")
			for i, loc := range gw.Locations {
				fmt.Printf("%d: %s\n", i, loc)
			}
			fmt.Print("Enter location index to move to: ")
			locInput, _ := reader.ReadString('\n')
			locInput = strings.TrimSpace(locInput)
			index, err := strconv.Atoi(locInput)
			if err != nil {
				fmt.Println("Invalid input.")
			} else {
				gw.MoveToLocation(index)
			}
		case "4":
			gw.Explore()
		case "5":
			gw.Battle()
		case "6":
			fmt.Print("Enter item name to use: ")
			itemInput, _ := reader.ReadString('\n')
			itemInput = strings.TrimSpace(itemInput)
			gw.Player.UseItem(itemInput)
		case "7":
			fmt.Print("Enter filename to save: ")
			filename, _ := reader.ReadString('\n')
			filename = strings.TrimSpace(filename)
			gw.SaveGame(filename)
		case "8":
			fmt.Print("Enter filename to load: ")
			filename, _ := reader.ReadString('\n')
			filename = strings.TrimSpace(filename)
			gw.LoadGame(filename)
		case "9":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
