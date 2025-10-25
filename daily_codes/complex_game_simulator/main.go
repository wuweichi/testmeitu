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

type Player struct {
	Name     string
	Health   int
	Mana     int
	Level    int
	Experience int
	Inventory []string
	Skills   map[string]int
}

type Enemy struct {
	Name   string
	Health int
	Damage int
	Loot   []string
}

type GameState struct {
	Player      Player
	Enemies     []Enemy
	CurrentRoom string
	GameLog     []string
}

func (p *Player) Attack(e *Enemy) string {
	damage := p.Level * 10
	e.Health -= damage
	return fmt.Sprintf("%s attacks %s for %d damage!", p.Name, e.Name, damage)
}

func (p *Player) CastSpell(spell string, e *Enemy) string {
	if cost, exists := p.Skills[spell]; exists && p.Mana >= cost {
		p.Mana -= cost
		damage := cost * 15
		e.Health -= damage
		return fmt.Sprintf("%s casts %s on %s for %d damage!", p.Name, spell, e.Name, damage)
	}
	return "Not enough mana or unknown spell!"
}

func (p *Player) LevelUp() {
	if p.Experience >= p.Level*100 {
		p.Level++
		p.Health += 20
		p.Mana += 10
		p.Experience = 0
		fmt.Printf("%s leveled up to level %d!\n", p.Name, p.Level)
	}
}

func (e *Enemy) IsAlive() bool {
	return e.Health > 0
}

func (gs *GameState) AddLog(message string) {
	gs.GameLog = append(gs.GameLog, message)
}

func (gs *GameState) DisplayStatus() {
	fmt.Printf("Player: %s (Level %d) - Health: %d, Mana: %d, Exp: %d\n",
		gs.Player.Name, gs.Player.Level, gs.Player.Health, gs.Player.Mana, gs.Player.Experience)
	fmt.Printf("Current Room: %s\n", gs.CurrentRoom)
	if len(gs.Enemies) > 0 {
		fmt.Println("Enemies in room:")
		for i, enemy := range gs.Enemies {
			fmt.Printf("  %d. %s - Health: %d\n", i+1, enemy.Name, enemy.Health)
		}
	}
}

func (gs *GameState) HandleCombat() {
	reader := bufio.NewReader(os.Stdin)
	for len(gs.Enemies) > 0 {
		gs.DisplayStatus()
		fmt.Print("Choose action (attack, spell, inventory, flee): ")
		action, _ := reader.ReadString('\n')
		action = strings.TrimSpace(action)
		switch action {
		case "attack":
			if len(gs.Enemies) > 0 {
				result := gs.Player.Attack(&gs.Enemies[0])
				gs.AddLog(result)
				fmt.Println(result)
				if !gs.Enemies[0].IsAlive() {
					gs.AddLog(fmt.Sprintf("%s defeated!", gs.Enemies[0].Name))
					gs.Player.Inventory = append(gs.Player.Inventory, gs.Enemies[0].Loot...)
					gs.Player.Experience += 50
					gs.Enemies = gs.Enemies[1:]
					gs.Player.LevelUp()
				}
			}
		case "spell":
			fmt.Print("Enter spell name: ")
			spell, _ := reader.ReadString('\n')
			spell = strings.TrimSpace(spell)
			if len(gs.Enemies) > 0 {
				result := gs.Player.CastSpell(spell, &gs.Enemies[0])
				gs.AddLog(result)
				fmt.Println(result)
				if !gs.Enemies[0].IsAlive() {
					gs.AddLog(fmt.Sprintf("%s defeated!", gs.Enemies[0].Name))
					gs.Player.Inventory = append(gs.Player.Inventory, gs.Enemies[0].Loot...)
					gs.Player.Experience += 50
					gs.Enemies = gs.Enemies[1:]
					gs.Player.LevelUp()
				}
			}
		case "inventory":
			fmt.Println("Inventory:", gs.Player.Inventory)
		case "flee":
			gs.AddLog("Player fled from combat!")
			return
		default:
			fmt.Println("Invalid action!")
		}
		// Enemy counterattack
		for i := range gs.Enemies {
			if gs.Enemies[i].IsAlive() {
				gs.Player.Health -= gs.Enemies[i].Damage
				gs.AddLog(fmt.Sprintf("%s attacks player for %d damage!", gs.Enemies[i].Name, gs.Enemies[i].Damage))
				fmt.Printf("%s attacks you for %d damage!\n", gs.Enemies[i].Name, gs.Enemies[i].Damage)
			}
		}
		if gs.Player.Health <= 0 {
			gs.AddLog("Player has been defeated!")
			fmt.Println("Game Over!")
			return
		}
	}
}

func (gs *GameState) ExploreRoom() {
	rooms := []string{"Forest", "Cave", "Castle", "Dungeon", "Mountain"}
	gs.CurrentRoom = rooms[rand.Intn(len(rooms))]
	gs.AddLog(fmt.Sprintf("Entered %s", gs.CurrentRoom))
	
	enemyTypes := []Enemy{
		{Name: "Goblin", Health: 30, Damage: 5, Loot: []string{"Gold", "Potion"}},
		{Name: "Orc", Health: 50, Damage: 10, Loot: []string{"Sword", "Shield"}},
		{Name: "Dragon", Health: 100, Damage: 20, Loot: []string{"Dragon Scale", "Treasure"}},
	}
	
	numEnemies := rand.Intn(3) + 1
	gs.Enemies = nil
	for i := 0; i < numEnemies; i++ {
		enemy := enemyTypes[rand.Intn(len(enemyTypes))]
		gs.Enemies = append(gs.Enemies, enemy)
	}
	
	gs.AddLog(fmt.Sprintf("Encountered %d enemies!", numEnemies))
}

func (gs *GameState) SaveGame() {
	data, err := json.Marshal(gs)
	if err != nil {
		fmt.Println("Error saving game:", err)
		return
	}
	err = os.WriteFile("savegame.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing save file:", err)
		return
	}
	fmt.Println("Game saved successfully!")
}

func (gs *GameState) LoadGame() {
	data, err := os.ReadFile("savegame.json")
	if err != nil {
		fmt.Println("No save game found.")
		return
	}
	err = json.Unmarshal(data, gs)
	if err != nil {
		fmt.Println("Error loading game:", err)
		return
	}
	fmt.Println("Game loaded successfully!")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	game := GameState{
		Player: Player{
			Name:     "Hero",
			Health:   100,
			Mana:     50,
			Level:    1,
			Experience: 0,
			Inventory: []string{"Potion", "Bread"},
			Skills:   map[string]int{"Fireball": 10, "Heal": 5},
		},
		CurrentRoom: "Starting Area",
		GameLog:     []string{"Game started!"},
	}
	
	reader := bufio.NewReader(os.Stdin)
	
	for {
		fmt.Println("\n=== Game Menu ===")
		fmt.Println("1. Explore new room")
		fmt.Println("2. Check status")
		fmt.Println("3. View inventory")
		fmt.Println("4. View game log")
		fmt.Println("5. Save game")
		fmt.Println("6. Load game")
		fmt.Println("7. Quit")
		fmt.Print("Choose an option: ")
		
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		
		switch input {
		case "1":
			game.ExploreRoom()
			if len(game.Enemies) > 0 {
				game.HandleCombat()
			}
		case "2":
			game.DisplayStatus()
		case "3":
			fmt.Println("Inventory:", game.Player.Inventory)
		case "4":
			fmt.Println("Game Log:")
			for _, log := range game.GameLog {
				fmt.Println(" -", log)
			}
		case "5":
			game.SaveGame()
		case "6":
			game.LoadGame()
		case "7":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid option!")
		}
		
		if game.Player.Health <= 0 {
			fmt.Println("Game Over!")
			return
		}
	}
}
