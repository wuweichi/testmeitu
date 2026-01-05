package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"strings"
	"os"
	"bufio"
	"encoding/json"
	"io/ioutil"
	"sync"
	"runtime"
	"sort"
	"errors"
)

type Player struct {
	ID          int
	Name        string
	Health      int
	MaxHealth   int
	Strength    int
	Agility     int
	Intelligence int
	Level       int
	Experience  int
	Gold        int
	Inventory   []Item
	Equipped    map[string]Item
	Skills      []Skill
}

type Item struct {
	ID          int
	Name        string
	Type        string // weapon, armor, potion, etc.
	Value       int
	Strength    int
	Agility     int
	Intelligence int
	HealthBonus int
	Description string
}

type Skill struct {
	ID          int
	Name        string
	Type        string // attack, heal, buff, etc.
	Power       int
	ManaCost    int
	Cooldown    int
	Description string
}

type Enemy struct {
	ID          int
	Name        string
	Health      int
	MaxHealth   int
	Strength    int
	Agility     int
	Intelligence int
	Experience  int
	GoldDrop    int
	LootTable   []Item
	Skills      []Skill
}

type GameState struct {
	Players     []Player
	Enemies     []Enemy
	CurrentTurn int
	Round       int
	GameOver    bool
	Winner      string
	Log         []string
}

var (
	gameState GameState
	mutex     sync.Mutex
	logger    []string
)

func init() {
	rand.Seed(time.Now().UnixNano())
	initializeGame()
}

func initializeGame() {
	// Initialize players
	player1 := Player{
		ID:          1,
		Name:        "Hero",
		Health:      100,
		MaxHealth:   100,
		Strength:    10,
		Agility:     8,
		Intelligence: 6,
		Level:       1,
		Experience:  0,
		Gold:        50,
		Inventory:   []Item{},
		Equipped:    make(map[string]Item),
		Skills:      []Skill{},
	}
	player2 := Player{
		ID:          2,
		Name:        "Mage",
		Health:      80,
		MaxHealth:   80,
		Strength:    4,
		Agility:     6,
		Intelligence: 12,
		Level:       1,
		Experience:  0,
		Gold:        30,
		Inventory:   []Item{},
		Equipped:    make(map[string]Item),
		Skills:      []Skill{},
	}
	gameState.Players = []Player{player1, player2}

	// Initialize enemies
	enemy1 := Enemy{
		ID:          1,
		Name:        "Goblin",
		Health:      50,
		MaxHealth:   50,
		Strength:    6,
		Agility:     5,
		Intelligence: 2,
		Experience:  20,
		GoldDrop:    10,
		LootTable:   []Item{},
		Skills:      []Skill{},
	}
	enemy2 := Enemy{
		ID:          2,
		Name:        "Orc",
		Health:      120,
		MaxHealth:   120,
		Strength:    12,
		Agility:     3,
		Intelligence: 4,
		Experience:  50,
		GoldDrop:    25,
		LootTable:   []Item{},
		Skills:      []Skill{},
	}
	gameState.Enemies = []Enemy{enemy1, enemy2}

	// Initialize items
	items := []Item{
		{ID: 1, Name: "Iron Sword", Type: "weapon", Value: 20, Strength: 5, Description: "A basic sword."},
		{ID: 2, Name: "Leather Armor", Type: "armor", Value: 15, Agility: 3, Description: "Light armor."},
		{ID: 3, Name: "Health Potion", Type: "potion", Value: 10, HealthBonus: 30, Description: "Restores health."},
		{ID: 4, Name: "Magic Staff", Type: "weapon", Value: 30, Intelligence: 8, Description: "Enhances magic."},
	}
	// Add items to players' inventories
	gameState.Players[0].Inventory = append(gameState.Players[0].Inventory, items[0], items[1])
	gameState.Players[1].Inventory = append(gameState.Players[1].Inventory, items[3])

	// Initialize skills
	skills := []Skill{
		{ID: 1, Name: "Slash", Type: "attack", Power: 15, ManaCost: 0, Cooldown: 0, Description: "Basic attack."},
		{ID: 2, Name: "Fireball", Type: "attack", Power: 25, ManaCost: 10, Cooldown: 2, Description: "Magical fire attack."},
		{ID: 3, Name: "Heal", Type: "heal", Power: 20, ManaCost: 15, Cooldown: 3, Description: "Restores health."},
	}
	gameState.Players[0].Skills = append(gameState.Players[0].Skills, skills[0])
	gameState.Players[1].Skills = append(gameState.Players[1].Skills, skills[1], skills[2])

	gameState.CurrentTurn = 0
	gameState.Round = 1
	gameState.GameOver = false
	gameState.Winner = ""
	gameState.Log = []string{"Game initialized with 2 players and 2 enemies."}
}

func (p *Player) Attack(target *Enemy, skill Skill) {
	damage := skill.Power + p.Strength/2
	if skill.Type == "attack" {
		damage += p.Intelligence / 4
	}
	target.Health -= damage
	logEntry := fmt.Sprintf("%s uses %s on %s for %d damage.", p.Name, skill.Name, target.Name, damage)
	gameState.Log = append(gameState.Log, logEntry)
	if target.Health <= 0 {
		target.Health = 0
		logEntry = fmt.Sprintf("%s has been defeated!", target.Name)
		gameState.Log = append(gameState.Log, logEntry)
		p.GainExperience(target.Experience)
		p.Gold += target.GoldDrop
	}
}

func (e *Enemy) Attack(target *Player, skill Skill) {
	damage := skill.Power + e.Strength/2
	target.Health -= damage
	logEntry := fmt.Sprintf("%s uses %s on %s for %d damage.", e.Name, skill.Name, target.Name, damage)
	gameState.Log = append(gameState.Log, logEntry)
	if target.Health <= 0 {
		target.Health = 0
		logEntry = fmt.Sprintf("%s has been defeated!", target.Name)
		gameState.Log = append(gameState.Log, logEntry)
	}
}

func (p *Player) GainExperience(exp int) {
	p.Experience += exp
	for p.Experience >= p.Level*100 {
		p.Level++
		p.MaxHealth += 20
		p.Health = p.MaxHealth
		p.Strength += 2
		p.Agility += 1
		p.Intelligence += 1
		logEntry := fmt.Sprintf("%s leveled up to level %d!", p.Name, p.Level)
		gameState.Log = append(gameState.Log, logEntry)
	}
}

func (p *Player) UseItem(item Item) {
	switch item.Type {
	case "potion":
		p.Health += item.HealthBonus
		if p.Health > p.MaxHealth {
			p.Health = p.MaxHealth
		}
		logEntry := fmt.Sprintf("%s uses %s and restores %d health.", p.Name, item.Name, item.HealthBonus)
		gameState.Log = append(gameState.Log, logEntry)
	case "weapon", "armor":
		p.Equipped[item.Type] = item
		logEntry := fmt.Sprintf("%s equips %s.", p.Name, item.Name)
		gameState.Log = append(gameState.Log, logEntry)
	}
	// Remove item from inventory
	for i, invItem := range p.Inventory {
		if invItem.ID == item.ID {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			break
		}
	}
}

func simulateCombat() {
	for !gameState.GameOver {
		mutex.Lock()
		logEntry := fmt.Sprintf("--- Round %d ---", gameState.Round)
		gameState.Log = append(gameState.Log, logEntry)

		// Player turns
		for i := range gameState.Players {
			if gameState.Players[i].Health > 0 {
				player := &gameState.Players[i]
				// Simple AI: attack random enemy with first skill
				if len(gameState.Enemies) > 0 {
					enemyIndex := rand.Intn(len(gameState.Enemies))
					enemy := &gameState.Enemies[enemyIndex]
					if len(player.Skills) > 0 {
						skill := player.Skills[0]
						player.Attack(enemy, skill)
					}
				}
			}
		}

		// Remove defeated enemies
		aliveEnemies := []Enemy{}
		for _, enemy := range gameState.Enemies {
			if enemy.Health > 0 {
				aliveEnemies = append(aliveEnemies, enemy)
			}
		}
		gameState.Enemies = aliveEnemies

		// Check if all enemies are defeated
		if len(gameState.Enemies) == 0 {
			gameState.GameOver = true
			gameState.Winner = "Players"
			logEntry = "All enemies defeated! Players win!"
			gameState.Log = append(gameState.Log, logEntry)
			mutex.Unlock()
			break
		}

		// Enemy turns
		for i := range gameState.Enemies {
			enemy := &gameState.Enemies[i]
			// Simple AI: attack random player with first skill
			if len(gameState.Players) > 0 {
				playerIndex := rand.Intn(len(gameState.Players))
				player := &gameState.Players[playerIndex]
				if len(enemy.Skills) > 0 {
					skill := enemy.Skills[0]
					enemy.Attack(player, skill)
				}
			}
		}

		// Remove defeated players
		alivePlayers := []Player{}
		for _, player := range gameState.Players {
			if player.Health > 0 {
				alivePlayers = append(alivePlayers, player)
			}
		}
		gameState.Players = alivePlayers

		// Check if all players are defeated
		if len(gameState.Players) == 0 {
			gameState.GameOver = true
			gameState.Winner = "Enemies"
			logEntry = "All players defeated! Enemies win!"
			gameState.Log = append(gameState.Log, logEntry)
			mutex.Unlock()
			break
		}

		gameState.Round++
		mutex.Unlock()
		time.Sleep(500 * time.Millisecond) // Simulate turn delay
	}
}

func saveGameState(filename string) error {
	mutex.Lock()
	defer mutex.Unlock()
	data, err := json.MarshalIndent(gameState, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

func loadGameState(filename string) error {
	mutex.Lock()
	defer mutex.Unlock()
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, &gameState)
}

func displayStatus() {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("=== Game Status ===")
	fmt.Println("Players:")
	for _, player := range gameState.Players {
		fmt.Printf("  %s: Health %d/%d, Level %d, Gold %d\n", player.Name, player.Health, player.MaxHealth, player.Level, player.Gold)
	}
	fmt.Println("Enemies:")
	for _, enemy := range gameState.Enemies {
		fmt.Printf("  %s: Health %d/%d\n", enemy.Name, enemy.Health, enemy.MaxHealth)
	}
	fmt.Printf("Round: %d\n", gameState.Round)
	if gameState.GameOver {
		fmt.Printf("Game Over! Winner: %s\n", gameState.Winner)
	}
}

func displayLog() {
	mutex.Lock()
	defer mutex.Unlock()
	fmt.Println("=== Game Log ===")
	for _, entry := range gameState.Log {
		fmt.Println(entry)
	}
}

func main() {
	fmt.Println("Welcome to Complex Game Simulator!")
	fmt.Println("Commands: start, status, log, save <filename>, load <filename>, exit")

	reader := bufio.NewReader(os.Stdin)
	go simulateCombat() // Start combat simulation in background

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Split(input, " ")
		command := parts[0]

		switch command {
		case "start":
			initializeGame()
			go simulateCombat()
			fmt.Println("Game restarted.")
		case "status":
			displayStatus()
		case "log":
			displayLog()
		case "save":
			if len(parts) < 2 {
				fmt.Println("Usage: save <filename>")
				continue
			}
			filename := parts[1]
			if err := saveGameState(filename); err != nil {
				fmt.Printf("Error saving game: %v\n", err)
			} else {
				fmt.Println("Game saved successfully.")
			}
		case "load":
			if len(parts) < 2 {
				fmt.Println("Usage: load <filename>")
				continue
			}
			filename := parts[1]
			if err := loadGameState(filename); err != nil {
				fmt.Printf("Error loading game: %v\n", err)
			} else {
				fmt.Println("Game loaded successfully.")
			}
		case "exit":
			fmt.Println("Exiting game.")
			return
		default:
			fmt.Println("Unknown command. Try: start, status, log, save, load, exit")
		}
	}
}
