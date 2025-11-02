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
	Strength int
	Agility  int
	Level    int
	Exp      int
}

type Monster struct {
	Name     string
	Health   int
	Strength int
	Agility  int
}

type Item struct {
	Name        string
	Description string
	Value       int
}

type GameState struct {
	Player      Player
	Monsters    []Monster
	Items       []Item
	CurrentRoom string
	GameOver    bool
}

func (p *Player) Attack(m *Monster) int {
	damage := p.Strength + rand.Intn(10)
	m.Health -= damage
	return damage
}

func (m *Monster) Attack(p *Player) int {
	damage := m.Strength + rand.Intn(5)
	p.Health -= damage
	return damage
}

func (p *Player) LevelUp() {
	if p.Exp >= p.Level*100 {
		p.Level++
		p.Strength += 5
		p.Agility += 3
		p.Health += 20
		p.Exp = 0
		fmt.Printf("Level up! You are now level %d\n", p.Level)
	}
}

func (p *Player) Heal(amount int) {
	p.Health += amount
	if p.Health > 100+p.Level*10 {
		p.Health = 100 + p.Level*10
	}
}

func (p *Player) IsAlive() bool {
	return p.Health > 0
}

func (m *Monster) IsAlive() bool {
	return m.Health > 0
}

func NewPlayer(name string) Player {
	return Player{
		Name:     name,
		Health:   100,
		Strength: 10,
		Agility:  10,
		Level:    1,
		Exp:      0,
	}
}

func NewMonster(name string, health, strength, agility int) Monster {
	return Monster{
		Name:     name,
		Health:   health,
		Strength: strength,
		Agility:  agility,
	}
}

func NewItem(name, description string, value int) Item {
	return Item{
		Name:        name,
		Description: description,
		Value:       value,
	}
}

func InitGame() GameState {
	player := NewPlayer("Hero")
	monsters := []Monster{
		NewMonster("Goblin", 30, 5, 8),
		NewMonster("Orc", 50, 10, 5),
		NewMonster("Dragon", 100, 20, 3),
	}
	items := []Item{
		NewItem("Health Potion", "Restores 50 health", 50),
		NewItem("Strength Elixir", "Increases strength by 5", 5),
		NewItem("Agility Boots", "Increases agility by 3", 3),
	}
	return GameState{
		Player:      player,
		Monsters:    monsters,
		Items:       items,
		CurrentRoom: "start",
		GameOver:    false,
	}
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

func (gs *GameState) DisplayStatus() {
	fmt.Printf("Player: %s\n", gs.Player.Name)
	fmt.Printf("Health: %d\n", gs.Player.Health)
	fmt.Printf("Strength: %d\n", gs.Player.Strength)
	fmt.Printf("Agility: %d\n", gs.Player.Agility)
	fmt.Printf("Level: %d\n", gs.Player.Level)
	fmt.Printf("Exp: %d\n", gs.Player.Exp)
	fmt.Printf("Current Room: %s\n", gs.CurrentRoom)
}

func (gs *GameState) DisplayMonsters() {
	fmt.Println("Monsters in the area:")
	for i, monster := range gs.Monsters {
		if monster.IsAlive() {
			fmt.Printf("%d. %s (Health: %d, Strength: %d, Agility: %d)\n", i+1, monster.Name, monster.Health, monster.Strength, monster.Agility)
		}
	}
}

func (gs *GameState) DisplayItems() {
	fmt.Println("Items available:")
	for i, item := range gs.Items {
		fmt.Printf("%d. %s - %s (Value: %d)\n", i+1, item.Name, item.Description, item.Value)
	}
}

func (gs *GameState) Fight(monsterIndex int) {
	if monsterIndex < 0 || monsterIndex >= len(gs.Monsters) {
		fmt.Println("Invalid monster index")
		return
	}
	monster := &gs.Monsters[monsterIndex]
	if !monster.IsAlive() {
		fmt.Println("This monster is already defeated")
		return
	}

	fmt.Printf("You are fighting %s!\n", monster.Name)
	for gs.Player.IsAlive() && monster.IsAlive() {
		playerDamage := gs.Player.Attack(monster)
		fmt.Printf("You attack %s for %d damage. %s's health: %d\n", monster.Name, playerDamage, monster.Name, monster.Health)
		if !monster.IsAlive() {
			fmt.Printf("You defeated %s!\n", monster.Name)
			gs.Player.Exp += 50
			gs.Player.LevelUp()
			break
		}

		monsterDamage := monster.Attack(&gs.Player)
		fmt.Printf("%s attacks you for %d damage. Your health: %d\n", monster.Name, monsterDamage, gs.Player.Health)
		if !gs.Player.IsAlive() {
			fmt.Println("You have been defeated! Game over.")
			gs.GameOver = true
			break
		}
	}
}

func (gs *GameState) UseItem(itemIndex int) {
	if itemIndex < 0 || itemIndex >= len(gs.Items) {
		fmt.Println("Invalid item index")
		return
	}
	item := gs.Items[itemIndex]
	switch item.Name {
	case "Health Potion":
		gs.Player.Heal(item.Value)
		fmt.Printf("You used %s and restored %d health. Your health: %d\n", item.Name, item.Value, gs.Player.Health)
	case "Strength Elixir":
		gs.Player.Strength += item.Value
		fmt.Printf("You used %s and increased strength by %d. Your strength: %d\n", item.Name, item.Value, gs.Player.Strength)
	case "Agility Boots":
		gs.Player.Agility += item.Value
		fmt.Printf("You used %s and increased agility by %d. Your agility: %d\n", item.Name, item.Value, gs.Player.Agility)
	default:
		fmt.Println("Unknown item")
	}
	gs.Items = append(gs.Items[:itemIndex], gs.Items[itemIndex+1:]...)
}

func (gs *GameState) MoveToRoom(room string) {
	gs.CurrentRoom = room
	fmt.Printf("You moved to %s\n", room)
}

func (gs *GameState) Explore() {
	fmt.Println("You explore the area...")
	// Simulate finding something
	if rand.Intn(100) < 30 {
		newItem := NewItem("Mysterious Potion", "A potion with unknown effects", 25)
		gs.Items = append(gs.Items, newItem)
		fmt.Printf("You found a %s!\n", newItem.Name)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := InitGame()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Fun Game Simulator!")
	fmt.Println("Type 'help' for a list of commands.")

	for !game.GameOver {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		switch args[0] {
		case "help":
			fmt.Println("Available commands:")
			fmt.Println("  status - Display player status")
			fmt.Println("  monsters - List monsters")
			fmt.Println("  items - List items")
			fmt.Println("  fight <monster_index> - Fight a monster")
			fmt.Println("  use <item_index> - Use an item")
			fmt.Println("  move <room> - Move to a different room")
			fmt.Println("  explore - Explore the current room")
			fmt.Println("  save <filename> - Save game")
			fmt.Println("  load <filename> - Load game")
			fmt.Println("  quit - Quit the game")
		case "status":
			game.DisplayStatus()
		case "monsters":
			game.DisplayMonsters()
		case "items":
			game.DisplayItems()
		case "fight":
			if len(args) < 2 {
				fmt.Println("Usage: fight <monster_index>")
			} else {
				index, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Println("Invalid monster index")
				} else {
					game.Fight(index - 1)
				}
			}
		case "use":
			if len(args) < 2 {
				fmt.Println("Usage: use <item_index>")
			} else {
				index, err := strconv.Atoi(args[1])
				if err != nil {
					fmt.Println("Invalid item index")
				} else {
					game.UseItem(index - 1)
				}
			}
		case "move":
			if len(args) < 2 {
				fmt.Println("Usage: move <room>")
			} else {
				game.MoveToRoom(args[1])
			}
		case "explore":
			game.Explore()
		case "save":
			if len(args) < 2 {
				fmt.Println("Usage: save <filename>")
			} else {
				err := game.SaveGame(args[1])
				if err != nil {
					fmt.Printf("Error saving game: %v\n", err)
				} else {
					fmt.Println("Game saved successfully")
				}
			}
		case "load":
			if len(args) < 2 {
				fmt.Println("Usage: load <filename>")
			} else {
				err := game.LoadGame(args[1])
				if err != nil {
					fmt.Printf("Error loading game: %v\n", err)
				} else {
					fmt.Println("Game loaded successfully")
				}
			}
		case "quit":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Unknown command. Type 'help' for a list of commands.")
		}
	}

	if game.GameOver {
		fmt.Println("Game over! Better luck next time.")
	}
}
