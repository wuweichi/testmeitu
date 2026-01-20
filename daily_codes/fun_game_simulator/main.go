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
    XP       int
    Gold     int
    Inventory []string
}

type Monster struct {
    Name     string
    Health   int
    Damage   int
    XP       int
    Gold     int
}

type GameState struct {
    Player      Player
    Day         int
    MonstersDefeated int
    GameOver    bool
}

func (p *Player) Attack(m *Monster) int {
    damage := p.Strength + rand.Intn(10)
    m.Health -= damage
    return damage
}

func (m *Monster) Attack(p *Player) int {
    damage := m.Damage + rand.Intn(5)
    p.Health -= damage
    return damage
}

func (p *Player) LevelUp() {
    p.Level++
    p.Strength += 5
    p.Agility += 3
    p.Health += 20
    fmt.Printf("\nCongratulations! You leveled up to level %d!\n", p.Level)
}

func (p *Player) GainXP(xp int) {
    p.XP += xp
    fmt.Printf("You gained %d XP. Total XP: %d\n", xp, p.XP)
    if p.XP >= p.Level*100 {
        p.LevelUp()
        p.XP = 0
    }
}

func (p *Player) AddItem(item string) {
    p.Inventory = append(p.Inventory, item)
    fmt.Printf("You found: %s\n", item)
}

func (p *Player) UseItem(item string) bool {
    for i, it := range p.Inventory {
        if it == item {
            p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
            switch item {
            case "Health Potion":
                p.Health += 50
                fmt.Println("You used a Health Potion and restored 50 health.")
            case "Strength Elixir":
                p.Strength += 10
                fmt.Println("You used a Strength Elixir and gained 10 strength.")
            }
            return true
        }
    }
    return false
}

func (p *Player) DisplayStatus() {
    fmt.Printf("\n=== Player Status ===\n")
    fmt.Printf("Name: %s\n", p.Name)
    fmt.Printf("Health: %d\n", p.Health)
    fmt.Printf("Strength: %d\n", p.Strength)
    fmt.Printf("Agility: %d\n", p.Agility)
    fmt.Printf("Level: %d\n", p.Level)
    fmt.Printf("XP: %d\n", p.XP)
    fmt.Printf("Gold: %d\n", p.Gold)
    fmt.Printf("Inventory: %v\n", p.Inventory)
    fmt.Printf("====================\n")
}

func generateMonster(day int) Monster {
    names := []string{"Goblin", "Orc", "Dragon", "Troll", "Skeleton"}
    name := names[rand.Intn(len(names))]
    health := 30 + day*10
    damage := 5 + day*2
    xp := 20 + day*5
    gold := 10 + day*3
    return Monster{Name: name, Health: health, Damage: damage, XP: xp, Gold: gold}
}

func battle(p *Player, m *Monster) bool {
    fmt.Printf("\nA wild %s appears!\n", m.Name)
    for p.Health > 0 && m.Health > 0 {
        fmt.Printf("\nYour Health: %d, %s's Health: %d\n", p.Health, m.Name, m.Health)
        fmt.Println("Choose action: (1) Attack, (2) Use Item, (3) Flee")
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        switch input {
        case "1":
            damage := p.Attack(m)
            fmt.Printf("You attack the %s for %d damage.\n", m.Name, damage)
            if m.Health <= 0 {
                fmt.Printf("You defeated the %s!\n", m.Name)
                p.GainXP(m.XP)
                p.Gold += m.Gold
                fmt.Printf("You found %d gold.\n", m.Gold)
                if rand.Intn(100) < 30 {
                    items := []string{"Health Potion", "Strength Elixir"}
                    item := items[rand.Intn(len(items))]
                    p.AddItem(item)
                }
                return true
            }
            damage = m.Attack(p)
            fmt.Printf("The %s attacks you for %d damage.\n", m.Name, damage)
            if p.Health <= 0 {
                fmt.Println("You have been defeated!")
                return false
            }
        case "2":
            fmt.Println("Your inventory:")
            for i, item := range p.Inventory {
                fmt.Printf("%d: %s\n", i+1, item)
            }
            fmt.Println("Enter item name to use (or 'back' to cancel):")
            itemInput, _ := reader.ReadString('\n')
            itemInput = strings.TrimSpace(itemInput)
            if itemInput == "back" {
                continue
            }
            if !p.UseItem(itemInput) {
                fmt.Println("Item not found in inventory.")
            }
        case "3":
            if rand.Intn(100) < p.Agility {
                fmt.Println("You successfully fled from the battle!")
                return false
            } else {
                fmt.Println("You failed to flee!")
                damage := m.Attack(p)
                fmt.Printf("The %s attacks you for %d damage.\n", m.Name, damage)
                if p.Health <= 0 {
                    fmt.Println("You have been defeated!")
                    return false
                }
            }
        default:
            fmt.Println("Invalid choice. Try again.")
        }
    }
    return false
}

func shop(p *Player) {
    fmt.Println("\nWelcome to the Shop!")
    fmt.Println("Items available:")
    fmt.Println("1. Health Potion - 50 gold (restores 50 health)")
    fmt.Println("2. Strength Elixir - 100 gold (increases strength by 10)")
    fmt.Println("3. Exit shop")
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Printf("\nYour gold: %d\n", p.Gold)
        fmt.Print("Enter choice (1-3): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        switch input {
        case "1":
            if p.Gold >= 50 {
                p.Gold -= 50
                p.AddItem("Health Potion")
                fmt.Println("You bought a Health Potion.")
            } else {
                fmt.Println("Not enough gold!")
            }
        case "2":
            if p.Gold >= 100 {
                p.Gold -= 100
                p.AddItem("Strength Elixir")
                fmt.Println("You bought a Strength Elixir.")
            } else {
                fmt.Println("Not enough gold!")
            }
        case "3":
            fmt.Println("Exiting shop.")
            return
        default:
            fmt.Println("Invalid choice. Try again.")
        }
    }
}

func saveGame(state GameState) error {
    data, err := json.Marshal(state)
    if err != nil {
        return err
    }
    return os.WriteFile("savegame.json", data, 0644)
}

func loadGame() (GameState, error) {
    var state GameState
    data, err := os.ReadFile("savegame.json")
    if err != nil {
        return state, err
    }
    err = json.Unmarshal(data, &state)
    return state, err
}

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("Welcome to Fun Game Simulator!")
    fmt.Println("This is a text-based adventure game where you battle monsters, level up, and collect items.")
    fmt.Println("Game commands: explore, status, shop, save, load, quit")
    
    var player Player
    var day int = 1
    var monstersDefeated int = 0
    var gameOver bool = false
    
    fmt.Print("Enter your character name: ")
    reader := bufio.NewReader(os.Stdin)
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)
    if name == "" {
        name = "Hero"
    }
    
    player = Player{
        Name:     name,
        Health:   100,
        Strength: 10,
        Agility:  10,
        Level:    1,
        XP:       0,
        Gold:     50,
        Inventory: []string{"Health Potion"},
    }
    
    state := GameState{Player: player, Day: day, MonstersDefeated: monstersDefeated, GameOver: gameOver}
    
    for !state.GameOver {
        fmt.Printf("\n=== Day %d ===\n", state.Day)
        fmt.Print("Enter command: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        
        switch input {
        case "explore":
            fmt.Println("You venture into the wilderness...")
            time.Sleep(1 * time.Second)
            monster := generateMonster(state.Day)
            if battle(&state.Player, &monster) {
                state.MonstersDefeated++
                state.Day++
            } else {
                if state.Player.Health <= 0 {
                    fmt.Println("Game Over!")
                    state.GameOver = true
                }
            }
        case "status":
            state.Player.DisplayStatus()
            fmt.Printf("Day: %d, Monsters Defeated: %d\n", state.Day, state.MonstersDefeated)
        case "shop":
            shop(&state.Player)
        case "save":
            err := saveGame(state)
            if err != nil {
                fmt.Println("Error saving game:", err)
            } else {
                fmt.Println("Game saved successfully.")
            }
        case "load":
            loadedState, err := loadGame()
            if err != nil {
                fmt.Println("Error loading game:", err)
            } else {
                state = loadedState
                fmt.Println("Game loaded successfully.")
            }
        case "quit":
            fmt.Println("Thanks for playing!")
            state.GameOver = true
        default:
            fmt.Println("Unknown command. Available commands: explore, status, shop, save, load, quit")
        }
        
        if state.Player.Health <= 0 {
            state.GameOver = true
            fmt.Println("Game Over! You have been defeated.")
        }
        
        if state.Day > 10 && state.MonstersDefeated >= 5 {
            fmt.Println("\nCongratulations! You have completed the game by surviving 10 days and defeating 5 monsters!")
            state.GameOver = true
        }
    }
    
    fmt.Printf("\nFinal Stats:\n")
    state.Player.DisplayStatus()
    fmt.Printf("Total Days: %d, Monsters Defeated: %d\n", state.Day, state.MonstersDefeated)
    fmt.Println("Game ended.")
}
