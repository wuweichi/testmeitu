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
	fmt.Println("Enter your name:")
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
		fmt.Printf("\n=== %s's Adventure ===\n", player.Name)
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
			showStatus(player)
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

		if player.Health <= 0 {
			fmt.Println("\nYou have been defeated. Game Over!")
			return
		}
	}
}

func explore(player *Player) {
	fmt.Println("\nYou venture into the unknown...")
	rand.Seed(time.Now().UnixNano())
	event := rand.Intn(10)

	switch event {
	case 0, 1, 2:
		encounterEnemy(player)
	case 3, 4:
		findTreasure(player)
	case 5, 6:
		findItem(player)
	case 7, 8:
		restArea(player)
	case 9:
		mysteriousEvent(player)
	}
}

func encounterEnemy(player *Player) {
	enemies := []Enemy{
		{"Goblin", 30, 8, 2, 20, 10},
		{"Orc", 50, 12, 5, 35, 20},
		{"Dragon", 100, 20, 10, 100, 50},
		{"Slime", 20, 5, 1, 10, 5},
		{"Bandit", 40, 10, 3, 25, 15},
	}

	rand.Seed(time.Now().UnixNano())
	enemy := enemies[rand.Intn(len(enemies))]

	fmt.Printf("\nA wild %s appears!\n", enemy.Name)

	for enemy.Health > 0 && player.Health > 0 {
		fmt.Println("\n1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Flee")

		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			playerAttack(player, &enemy)
			if enemy.Health > 0 {
				enemyAttack(player, &enemy)
			}
		case "2":
			useItem(player)
		case "3":
			if rand.Intn(2) == 0 {
				fmt.Println("You successfully fled!")
				return
			} else {
				fmt.Println("You failed to flee!")
				enemyAttack(player, &enemy)
			}
		default:
			fmt.Println("Invalid choice. The enemy attacks!")
			enemyAttack(player, &enemy)
		}
	}

	if enemy.Health <= 0 {
		fmt.Printf("\nYou defeated the %s!\n", enemy.Name)
		player.Exp += enemy.Exp
		player.Gold += enemy.Gold
		fmt.Printf("Gained %d EXP and %d Gold!\n", enemy.Exp, enemy.Gold)
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
	fmt.Printf("%s's Health: %d\n", enemy.Name, enemy.Health)
}

func enemyAttack(player *Player, enemy *Enemy) {
	damage := enemy.Attack - player.Defense
	if damage < 1 {
		damage = 1
	}
	player.Health -= damage
	fmt.Printf("The %s attacks you for %d damage!\n", enemy.Name, damage)
	fmt.Printf("Your Health: %d\n", player.Health)
}

func findTreasure(player *Player) {
	goldFound := rand.Intn(50) + 10
	player.Gold += goldFound
	fmt.Printf("\nYou found a treasure chest!\n")
	fmt.Printf("You gained %d Gold!\n", goldFound)
}

func findItem(player *Player) {
	items := []Item{
		{"Potion", "Restores 30 HP", func(p *Player) {
			p.Health += 30
			if p.Health > p.MaxHealth {
				p.Health = p.MaxHealth
			}
			fmt.Println("Restored 30 HP!")
		}},
		{"Super Potion", "Restores 60 HP", func(p *Player) {
			p.Health += 60
			if p.Health > p.MaxHealth {
				p.Health = p.MaxHealth
			}
			fmt.Println("Restored 60 HP!")
		}},
		{"Attack Boost", "Increases Attack by 5", func(p *Player) {
			p.Attack += 5
			fmt.Println("Attack increased by 5!")
		}},
		{"Defense Boost", "Increases Defense by 3", func(p *Player) {
			p.Defense += 3
			fmt.Println("Defense increased by 3!")
		}},
	}

	rand.Seed(time.Now().UnixNano())
	item := items[rand.Intn(len(items))]
	player.Inventory = append(player.Inventory, item.Name)
	fmt.Printf("\nYou found a %s!\n", item.Name)
	fmt.Println(item.Description)
}

func restArea(player *Player) {
	fmt.Println("\nYou find a peaceful resting area.")
	player.Health = player.MaxHealth
	fmt.Println("Your health has been fully restored!")
}

func mysteriousEvent(player *Player) {
	fmt.Println("\nYou encounter a mysterious event!")
	rand.Seed(time.Now().UnixNano())
	event := rand.Intn(3)

	switch event {
	case 0:
		fmt.Println("A wise old wizard grants you a blessing!")
		player.MaxHealth += 10
		player.Health = player.MaxHealth
		fmt.Println("Your max health increased by 10!")
	case 1:
		fmt.Println("You stumble upon an ancient shrine.")
		player.Exp += 50
		fmt.Println("Gained 50 EXP!")
		checkLevelUp(player)
	case 2:
		fmt.Println("A mischievous fairy plays a trick on you!")
		player.Gold -= 20
		if player.Gold < 0 {
			player.Gold = 0
		}
		fmt.Println("You lost 20 Gold!")
	}
}

func showStatus(player *Player) {
	fmt.Printf("\n=== %s's Status ===\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("EXP: %d/%d\n", player.Exp, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Inventory: %v\n", player.Inventory)
}

func useItem(player *Player) {
	if len(player.Inventory) == 0 {
		fmt.Println("\nYour inventory is empty!")
		return
	}

	fmt.Println("\nSelect an item to use:")
	for i, item := range player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(player.Inventory) {
		fmt.Println("Invalid choice.")
		return
	}

	itemName := player.Inventory[index-1]
	items := map[string]func(*Player){
		"Potion": func(p *Player) {
			p.Health += 30
			if p.Health > p.MaxHealth {
				p.Health = p.MaxHealth
			}
			fmt.Println("Used Potion. Restored 30 HP!")
		},
		"Super Potion": func(p *Player) {
			p.Health += 60
			if p.Health > p.MaxHealth {
				p.Health = p.MaxHealth
			}
			fmt.Println("Used Super Potion. Restored 60 HP!")
		},
		"Attack Boost": func(p *Player) {
			p.Attack += 5
			fmt.Println("Used Attack Boost. Attack increased by 5!")
		},
		"Defense Boost": func(p *Player) {
			p.Defense += 3
			fmt.Println("Used Defense Boost. Defense increased by 3!")
		},
	}

	if effect, exists := items[itemName]; exists {
		effect(player)
		player.Inventory = append(player.Inventory[:index-1], player.Inventory[index:]...)
	} else {
		fmt.Println("Item not found or cannot be used.")
	}
}

func checkLevelUp(player *Player) {
	for player.Exp >= player.Level*100 {
		player.Exp -= player.Level * 100
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("\nLevel Up! You are now level %d!\n", player.Level)
		fmt.Println("Max Health +20, Attack +3, Defense +2")
	}
}

func saveGame(player *Player) {
	file, err := os.Create("savegame.txt")
	if err != nil {
		fmt.Println("Error saving game:", err)
		return
	}
	defer file.Close()

	fmt.Fprintf(file, "%s\n", player.Name)
	fmt.Fprintf(file, "%d\n", player.Health)
	fmt.Fprintf(file, "%d\n", player.MaxHealth)
	fmt.Fprintf(file, "%d\n", player.Attack)
	fmt.Fprintf(file, "%d\n", player.Defense)
	fmt.Fprintf(file, "%d\n", player.Level)
	fmt.Fprintf(file, "%d\n", player.Exp)
	fmt.Fprintf(file, "%d\n", player.Gold)
	fmt.Fprintf(file, "%s\n", strings.Join(player.Inventory, ","))

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
	name, _ := reader.ReadString('\n')
	healthStr, _ := reader.ReadString('\n')
	maxHealthStr, _ := reader.ReadString('\n')
	attackStr, _ := reader.ReadString('\n')
	defenseStr, _ := reader.ReadString('\n')
	levelStr, _ := reader.ReadString('\n')
	expStr, _ := reader.ReadString('\n')
	goldStr, _ := reader.ReadString('\n')
	inventoryStr, _ := reader.ReadString('\n')

	player.Name = strings.TrimSpace(name)
	player.Health, _ = strconv.Atoi(strings.TrimSpace(healthStr))
	player.MaxHealth, _ = strconv.Atoi(strings.TrimSpace(maxHealthStr))
	player.Attack, _ = strconv.Atoi(strings.TrimSpace(attackStr))
	player.Defense, _ = strconv.Atoi(strings.TrimSpace(defenseStr))
	player.Level, _ = strconv.Atoi(strings.TrimSpace(levelStr))
	player.Exp, _ = strconv.Atoi(strings.TrimSpace(expStr))
	player.Gold, _ = strconv.Atoi(strings.TrimSpace(goldStr))
	player.Inventory = strings.Split(strings.TrimSpace(inventoryStr), ",")

	fmt.Println("Game loaded successfully!")
}
