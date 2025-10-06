package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type Player struct {
	Name      string
	Health    int
	MaxHealth int
	Attack    int
	Defense   int
	Level     int
	Exp       int
	Gold      int
	Inventory []string
}

type Enemy struct {
	Name    string
	Health  int
	Attack  int
	Defense int
	Exp     int
	Gold    int
}

type Item struct {
	Name        string
	Description string
	Effect      func(*Player)
}

func main() {
	fmt.Println("Welcome to the Adventure Game!")
	fmt.Println("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	player := &Player{
		Name:      name,
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
		Gold:      50,
		Inventory: []string{"Potion"},
	}

	gameLoop(player)
}

func gameLoop(player *Player) {
	for {
		if player.Health <= 0 {
			fmt.Println("You have been defeated. Game Over!")
			return
		}

		fmt.Println("\n=== Main Menu ===")
		fmt.Println("1. Explore")
		fmt.Println("2. Check Status")
		fmt.Println("3. Use Item")
		fmt.Println("4. Save Game")
		fmt.Println("5. Load Game")
		fmt.Println("6. Quit")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			explore(player)
		case "2":
			checkStatus(player)
		case "3":
			useItem(player)
		case "4":
			saveGame(player)
		case "5":
			loadGame(player)
		case "6":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func explore(player *Player) {
	fmt.Println("\nYou venture into the wilderness...")
	rand.Seed(time.Now().UnixNano())
	event := rand.Intn(10)

	switch event {
	case 0, 1, 2:
		fmt.Println("You found a treasure chest!")
		foundGold := rand.Intn(50) + 10
		player.Gold += foundGold
		fmt.Printf("You found %d gold!\n", foundGold)
	case 3, 4, 5:
		encounterEnemy(player)
	case 6, 7:
		fmt.Println("You discovered a hidden path.")
		player.Exp += 10
		fmt.Println("You gained 10 experience points.")
		checkLevelUp(player)
	case 8:
		fmt.Println("You stumbled upon a mysterious shrine.")
		player.Health = player.MaxHealth
		fmt.Println("Your health has been fully restored!")
	case 9:
		fmt.Println("You found a rare item!")
		items := []string{"Magic Sword", "Dragon Shield", "Ancient Amulet"}
		foundItem := items[rand.Intn(len(items))]
		player.Inventory = append(player.Inventory, foundItem)
		fmt.Printf("You found a %s!\n", foundItem)
	}
}

func encounterEnemy(player *Player) {
	enemies := []Enemy{
		{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Exp: 20, Gold: 15},
		{Name: "Orc", Health: 50, Attack: 8, Defense: 4, Exp: 35, Gold: 25},
		{Name: "Dragon", Health: 100, Attack: 15, Defense: 10, Exp: 100, Gold: 100},
	}

	enemy := enemies[rand.Intn(len(enemies))]
	fmt.Printf("\nA wild %s appears!\n", enemy.Name)

	for enemy.Health > 0 && player.Health > 0 {
		fmt.Println("\n=== Battle ===")
		fmt.Printf("Your Health: %d/%d\n", player.Health, player.MaxHealth)
		fmt.Printf("%s's Health: %d\n", enemy.Name, enemy.Health)
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Flee")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			playerAttack(player, &enemy)
		case "2":
			useItem(player)
		case "3":
			if rand.Intn(2) == 0 {
				fmt.Println("You successfully fled from the battle!")
				return
			} else {
				fmt.Println("You failed to flee!")
			}
		default:
			fmt.Println("Invalid choice. The enemy attacks!")
		}

		if enemy.Health > 0 {
			enemyAttack(player, &enemy)
		}
	}

	if enemy.Health <= 0 {
		fmt.Printf("\nYou defeated the %s!\n", enemy.Name)
		player.Exp += enemy.Exp
		player.Gold += enemy.Gold
		fmt.Printf("You gained %d experience and %d gold.\n", enemy.Exp, enemy.Gold)
		checkLevelUp(player)
	}
}

func playerAttack(player *Player, enemy *Enemy) {
	damage := player.Attack - enemy.Defense
	if damage < 1 {
		damage = 1
	}
	enemy.Health -= damage
	fmt.Printf("You attack the %s for %d damage!\n", enemy.Name, damage)
}

func enemyAttack(player *Player, enemy *Enemy) {
	damage := enemy.Attack - player.Defense
	if damage < 1 {
		damage = 1
	}
	player.Health -= damage
	fmt.Printf("The %s attacks you for %d damage!\n", enemy.Name, damage)
}

func checkStatus(player *Player) {
	fmt.Println("\n=== Player Status ===")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Experience: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Println("Inventory:")
	for _, item := range player.Inventory {
		fmt.Printf("  - %s\n", item)
	}
}

func useItem(player *Player) {
	if len(player.Inventory) == 0 {
		fmt.Println("You have no items to use.")
		return
	}

	fmt.Println("\n=== Use Item ===")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	fmt.Println("0. Back")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	choice, err := strconv.Atoi(input)
	if err != nil || choice < 0 || choice > len(player.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}

	if choice == 0 {
		return
	}

	item := player.Inventory[choice-1]
	fmt.Printf("You used %s.\n", item)

	switch item {
	case "Potion":
		healAmount := 30
		if player.Health+healAmount > player.MaxHealth {
			player.Health = player.MaxHealth
		} else {
			player.Health += healAmount
		}
		fmt.Printf("Restored %d health.\n", healAmount)
	case "Magic Sword":
		player.Attack += 5
		fmt.Println("Your attack increased by 5!")
	case "Dragon Shield":
		player.Defense += 5
		fmt.Println("Your defense increased by 5!")
	case "Ancient Amulet":
		player.MaxHealth += 20
		player.Health += 20
		fmt.Println("Your max health increased by 20!")
	}

	player.Inventory = append(player.Inventory[:choice-1], player.Inventory[choice:]...)
}

func checkLevelUp(player *Player) {
	for player.Exp >= player.Level*100 {
		player.Exp -= player.Level * 100
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("\nCongratulations! You reached level %d!\n", player.Level)
		fmt.Println("Your stats have improved!")
	}
}

func saveGame(player *Player) {
	file, err := os.Create("savegame.txt")
	if err != nil {
		fmt.Println("Error saving game:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	fmt.Fprintf(writer, "%s\n", player.Name)
	fmt.Fprintf(writer, "%d\n", player.Health)
	fmt.Fprintf(writer, "%d\n", player.MaxHealth)
	fmt.Fprintf(writer, "%d\n", player.Attack)
	fmt.Fprintf(writer, "%d\n", player.Defense)
	fmt.Fprintf(writer, "%d\n", player.Level)
	fmt.Fprintf(writer, "%d\n", player.Exp)
	fmt.Fprintf(writer, "%d\n", player.Gold)
	fmt.Fprintf(writer, "%d\n", len(player.Inventory))
	for _, item := range player.Inventory {
		fmt.Fprintf(writer, "%s\n", item)
	}

	writer.Flush()
	fmt.Println("Game saved successfully!")
}

func loadGame(player *Player) {
	file, err := os.Open("savegame.txt")
	if err != nil {
		fmt.Println("No saved game found.")
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	player.Name, _ = reader.ReadString('\n')
	player.Name = strings.TrimSpace(player.Name)

	healthStr, _ := reader.ReadString('\n')
	player.Health, _ = strconv.Atoi(strings.TrimSpace(healthStr))

	maxHealthStr, _ := reader.ReadString('\n')
	player.MaxHealth, _ = strconv.Atoi(strings.TrimSpace(maxHealthStr))

	attackStr, _ := reader.ReadString('\n')
	player.Attack, _ = strconv.Atoi(strings.TrimSpace(attackStr))

	defenseStr, _ := reader.ReadString('\n')
	player.Defense, _ = strconv.Atoi(strings.TrimSpace(defenseStr))

	levelStr, _ := reader.ReadString('\n')
	player.Level, _ = strconv.Atoi(strings.TrimSpace(levelStr))

	expStr, _ := reader.ReadString('\n')
	player.Exp, _ = strconv.Atoi(strings.TrimSpace(expStr))

	goldStr, _ := reader.ReadString('\n')
	player.Gold, _ = strconv.Atoi(strings.TrimSpace(goldStr))

	inventoryCountStr, _ := reader.ReadString('\n')
	inventoryCount, _ := strconv.Atoi(strings.TrimSpace(inventoryCountStr))

	player.Inventory = make([]string, inventoryCount)
	for i := 0; i < inventoryCount; i++ {
		player.Inventory[i], _ = reader.ReadString('\n')
		player.Inventory[i] = strings.TrimSpace(player.Inventory[i])
	}

	fmt.Println("Game loaded successfully!")
}
