package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

type Player struct {
	ID          int
	Name        string
	Health      int
	MaxHealth   int
	Attack      int
	Defense     int
	Level       int
	Experience  int
	Gold        int
	Inventory   []string
	Skills      map[string]int
	Position    Coordinates
	IsAlive     bool
}

type Coordinates struct {
	X int
	Y int
}

type Monster struct {
	ID        int
	Name      string
	Health    int
	Attack    int
	Defense   int
	RewardExp int
	RewardGold int
	Position  Coordinates
}

type GameWorld struct {
	Width     int
	Height    int
	Grid      [][]string
	Monsters  []Monster
	Players   []Player
	Items     map[Coordinates]string
	Mutex     sync.Mutex
}

func NewGameWorld(width, height int) *GameWorld {
	grid := make([][]string, height)
	for i := range grid {
		grid[i] = make([]string, width)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}
	return &GameWorld{
		Width:   width,
		Height:  height,
		Grid:    grid,
		Monsters: []Monster{},
		Players:  []Player{},
		Items:    make(map[Coordinates]string),
	}
}

func (gw *GameWorld) AddPlayer(p Player) {
	gw.Mutex.Lock()
	defer gw.Mutex.Unlock()
	gw.Players = append(gw.Players, p)
	gw.Grid[p.Position.Y][p.Position.X] = "P"
}

func (gw *GameWorld) AddMonster(m Monster) {
	gw.Mutex.Lock()
	defer gw.Mutex.Unlock()
	gw.Monsters = append(gw.Monsters, m)
	gw.Grid[m.Position.Y][m.Position.X] = "M"
}

func (gw *GameWorld) MovePlayer(playerID int, direction string) bool {
	gw.Mutex.Lock()
	defer gw.Mutex.Unlock()
	for i, p := range gw.Players {
		if p.ID == playerID && p.IsAlive {
			newPos := p.Position
			switch direction {
			case "up":
				newPos.Y--
			case "down":
				newPos.Y++
			case "left":
				newPos.X--
			case "right":
				newPos.X++
			default:
				return false
			}
			if newPos.X < 0 || newPos.X >= gw.Width || newPos.Y < 0 || newPos.Y >= gw.Height {
				return false
			}
			if gw.Grid[newPos.Y][newPos.X] == "M" {
				return false
			}
			gw.Grid[p.Position.Y][p.Position.X] = "."
			gw.Players[i].Position = newPos
			gw.Grid[newPos.Y][newPos.X] = "P"
			return true
		}
	}
	return false
}

func (gw *GameWorld) AttackMonster(playerID, monsterID int) (bool, string) {
	gw.Mutex.Lock()
	defer gw.Mutex.Unlock()
	var player *Player
	var monster *Monster
	for i := range gw.Players {
		if gw.Players[i].ID == playerID && gw.Players[i].IsAlive {
			player = &gw.Players[i]
			break
		}
	}
	if player == nil {
		return false, "Player not found or dead"
	}
	for i := range gw.Monsters {
		if gw.Monsters[i].ID == monsterID {
			monster = &gw.Monsters[i]
			break
		}
	}
	if monster == nil {
		return false, "Monster not found"
	}
	distance := abs(player.Position.X-monster.Position.X) + abs(player.Position.Y-monster.Position.Y)
	if distance > 1 {
		return false, "Monster too far to attack"
	}
	damage := player.Attack - monster.Defense
	if damage < 0 {
		damage = 0
	}
	monster.Health -= damage
	if monster.Health <= 0 {
		player.Experience += monster.RewardExp
		player.Gold += monster.RewardGold
		gw.Grid[monster.Position.Y][monster.Position.X] = "."
		gw.Monsters = append(gw.Monsters[:monsterID], gw.Monsters[monsterID+1:]...)
		return true, fmt.Sprintf("Monster defeated! Gained %d exp and %d gold", monster.RewardExp, monster.RewardGold)
	}
	return true, fmt.Sprintf("Dealt %d damage to monster", damage)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (gw *GameWorld) Display() {
	gw.Mutex.Lock()
	defer gw.Mutex.Unlock()
	fmt.Println("Game World:")
	for y := 0; y < gw.Height; y++ {
		for x := 0; x < gw.Width; x++ {
			fmt.Print(gw.Grid[y][x], " ")
		}
		fmt.Println()
	}
	fmt.Println("Players:")
	for _, p := range gw.Players {
		if p.IsAlive {
			fmt.Printf("ID: %d, Name: %s, Health: %d/%d, Position: (%d,%d)\n", p.ID, p.Name, p.Health, p.MaxHealth, p.Position.X, p.Position.Y)
		}
	}
	fmt.Println("Monsters:")
	for _, m := range gw.Monsters {
		fmt.Printf("ID: %d, Name: %s, Health: %d, Position: (%d,%d)\n", m.ID, m.Name, m.Health, m.Position.X, m.Position.Y)
	}
}

func (p *Player) LevelUp() {
	if p.Experience >= p.Level*100 {
		p.Experience -= p.Level * 100
		p.Level++
		p.MaxHealth += 10
		p.Health = p.MaxHealth
		p.Attack += 5
		p.Defense += 2
		fmt.Printf("%s leveled up to level %d!\n", p.Name, p.Level)
	}
}

func (gw *GameWorld) SaveGame(filename string) error {
	gw.Mutex.Lock()
	defer gw.Mutex.Unlock()
	data, err := json.Marshal(gw)
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func (gw *GameWorld) LoadGame(filename string) error {
	gw.Mutex.Lock()
	defer gw.Mutex.Unlock()
	data, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	return json.Unmarshal(data, gw)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	world := NewGameWorld(10, 10)
	player1 := Player{
		ID:         1,
		Name:       "Hero",
		Health:     100,
		MaxHealth:  100,
		Attack:     20,
		Defense:    5,
		Level:      1,
		Experience: 0,
		Gold:       50,
		Inventory:  []string{"Sword", "Potion"},
		Skills:     map[string]int{"Slash": 10},
		Position:   Coordinates{X: 0, Y: 0},
		IsAlive:    true,
	}
	world.AddPlayer(player1)
	for i := 0; i < 5; i++ {
		monster := Monster{
			ID:         i,
			Name:       fmt.Sprintf("Goblin%d", i),
			Health:     30 + rand.Intn(20),
			Attack:     10 + rand.Intn(10),
			Defense:    3 + rand.Intn(5),
			RewardExp:  50,
			RewardGold: 10 + rand.Intn(20),
			Position:   Coordinates{X: rand.Intn(10), Y: rand.Intn(10)},
		}
		world.AddMonster(monster)
	}
	reader := bufio.NewReader(os.Stdin)
	for {
		world.Display()
		fmt.Print("Enter command (move/attack/save/load/quit): ")
		cmd, _ := reader.ReadString('\n')
		cmd = strings.TrimSpace(cmd)
		switch cmd {
		case "move":
			fmt.Print("Enter direction (up/down/left/right): ")
			dir, _ := reader.ReadString('\n')
			dir = strings.TrimSpace(dir)
			if world.MovePlayer(1, dir) {
				fmt.Println("Moved successfully")
			} else {
				fmt.Println("Move failed")
			}
		case "attack":
			fmt.Print("Enter monster ID to attack: ")
			idStr, _ := reader.ReadString('\n')
			idStr = strings.TrimSpace(idStr)
			monsterID, err := strconv.Atoi(idStr)
			if err != nil {
				fmt.Println("Invalid monster ID")
				continue
			}
			success, msg := world.AttackMonster(1, monsterID)
			fmt.Println(msg)
			if success {
				for i := range world.Players {
					if world.Players[i].ID == 1 {
						world.Players[i].LevelUp()
						break
					}
				}
			}
		case "save":
			err := world.SaveGame("savegame.json")
			if err != nil {
				fmt.Println("Save failed:", err)
			} else {
				fmt.Println("Game saved")
			}
		case "load":
			err := world.LoadGame("savegame.json")
			if err != nil {
				fmt.Println("Load failed:", err)
			} else {
				fmt.Println("Game loaded")
			}
		case "quit":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Unknown command")
		}
		time.Sleep(500 * time.Millisecond)
	}
}
