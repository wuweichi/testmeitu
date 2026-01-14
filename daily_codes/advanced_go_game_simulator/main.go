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
    ID        int
    Name      string
    Health    int
    MaxHealth int
    Attack    int
    Defense   int
    Level     int
    Experience int
    Inventory []string
    Gold      int
    Position  Point
}

type Point struct {
    X int
    Y int
}

type Monster struct {
    ID      int
    Name    string
    Health  int
    Attack  int
    Defense int
    Reward  Reward
}

type Reward struct {
    Experience int
    Gold       int
    Items      []string
}

type GameWorld struct {
    Width    int
    Height   int
    Players  []Player
    Monsters []Monster
    Map      [][]string
    mu       sync.Mutex
}

func NewGameWorld(width, height int) *GameWorld {
    gw := &GameWorld{
        Width:   width,
        Height:  height,
        Players: []Player{},
        Map:     make([][]string, height),
    }
    for i := range gw.Map {
        gw.Map[i] = make([]string, width)
        for j := range gw.Map[i] {
            gw.Map[i][j] = "."
        }
    }
    return gw
}

func (gw *GameWorld) AddPlayer(p Player) {
    gw.mu.Lock()
    defer gw.mu.Unlock()
    gw.Players = append(gw.Players, p)
    gw.Map[p.Position.Y][p.Position.X] = "P" + strconv.Itoa(p.ID)
}

func (gw *GameWorld) RemovePlayer(id int) {
    gw.mu.Lock()
    defer gw.mu.Unlock()
    for i, p := range gw.Players {
        if p.ID == id {
            gw.Map[p.Position.Y][p.Position.X] = "."
            gw.Players = append(gw.Players[:i], gw.Players[i+1:]...)
            break
        }
    }
}

func (gw *GameWorld) MovePlayer(id int, dx, dy int) bool {
    gw.mu.Lock()
    defer gw.mu.Unlock()
    for i, p := range gw.Players {
        if p.ID == id {
            newX := p.Position.X + dx
            newY := p.Position.Y + dy
            if newX >= 0 && newX < gw.Width && newY >= 0 && newY < gw.Height {
                gw.Map[p.Position.Y][p.Position.X] = "."
                p.Position.X = newX
                p.Position.Y = newY
                gw.Map[newY][newX] = "P" + strconv.Itoa(p.ID)
                gw.Players[i] = p
                return true
            }
            return false
        }
    }
    return false
}

func (gw *GameWorld) Display() {
    gw.mu.Lock()
    defer gw.mu.Unlock()
    fmt.Println("Game World Map:")
    for _, row := range gw.Map {
        fmt.Println(strings.Join(row, " "))
    }
    fmt.Println("Players:")
    for _, p := range gw.Players {
        fmt.Printf("ID: %d, Name: %s, Health: %d/%d, Position: (%d, %d)\n", p.ID, p.Name, p.Health, p.MaxHealth, p.Position.X, p.Position.Y)
    }
}

func NewPlayer(id int, name string) Player {
    return Player{
        ID:        id,
        Name:      name,
        Health:    100,
        MaxHealth: 100,
        Attack:    10,
        Defense:   5,
        Level:     1,
        Experience: 0,
        Inventory: []string{"Potion", "Sword"},
        Gold:      50,
        Position:  Point{X: 0, Y: 0},
    }
}

func (p *Player) TakeDamage(damage int) {
    actualDamage := damage - p.Defense
    if actualDamage < 0 {
        actualDamage = 0
    }
    p.Health -= actualDamage
    if p.Health < 0 {
        p.Health = 0
    }
}

func (p *Player) Heal(amount int) {
    p.Health += amount
    if p.Health > p.MaxHealth {
        p.Health = p.MaxHealth
    }
}

func (p *Player) GainExperience(exp int) {
    p.Experience += exp
    for p.Experience >= p.Level*100 {
        p.Experience -= p.Level * 100
        p.Level++
        p.MaxHealth += 20
        p.Health = p.MaxHealth
        p.Attack += 5
        p.Defense += 2
        fmt.Printf("%s leveled up to level %d!\n", p.Name, p.Level)
    }
}

func NewMonster(id int, name string, health, attack, defense int, reward Reward) Monster {
    return Monster{
        ID:      id,
        Name:    name,
        Health:  health,
        Attack:  attack,
        Defense: defense,
        Reward:  reward,
    }
}

func (m *Monster) TakeDamage(damage int) {
    actualDamage := damage - m.Defense
    if actualDamage < 0 {
        actualDamage = 0
    }
    m.Health -= actualDamage
    if m.Health < 0 {
        m.Health = 0
    }
}

func Battle(player *Player, monster *Monster) bool {
    fmt.Printf("Battle between %s and %s!\n", player.Name, monster.Name)
    for player.Health > 0 && monster.Health > 0 {
        playerDamage := player.Attack + rand.Intn(10)
        monster.TakeDamage(playerDamage)
        fmt.Printf("%s attacks %s for %d damage. Monster health: %d\n", player.Name, monster.Name, playerDamage, monster.Health)
        if monster.Health <= 0 {
            fmt.Printf("%s defeated %s!\n", player.Name, monster.Name)
            player.GainExperience(monster.Reward.Experience)
            player.Gold += monster.Reward.Gold
            player.Inventory = append(player.Inventory, monster.Reward.Items...)
            return true
        }
        monsterDamage := monster.Attack + rand.Intn(5)
        player.TakeDamage(monsterDamage)
        fmt.Printf("%s attacks %s for %d damage. Player health: %d\n", monster.Name, player.Name, monsterDamage, player.Health)
        if player.Health <= 0 {
            fmt.Printf("%s was defeated by %s!\n", player.Name, monster.Name)
            return false
        }
        time.Sleep(500 * time.Millisecond)
    }
    return false
}

func SaveGame(gw *GameWorld, filename string) error {
    data, err := json.Marshal(gw)
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}

func LoadGame(filename string) (*GameWorld, error) {
    data, err := os.ReadFile(filename)
    if err != nil {
        return nil, err
    }
    var gw GameWorld
    err = json.Unmarshal(data, &gw)
    if err != nil {
        return nil, err
    }
    return &gw, nil
}

func main() {
    rand.Seed(time.Now().UnixNano())
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Welcome to Advanced Go Game Simulator!")
    fmt.Println("Commands: new, load, move, battle, save, quit")
    var gameWorld *GameWorld
    playerIDCounter := 1
    monsterIDCounter := 1
    for {
        fmt.Print("> ")
        scanner.Scan()
        command := strings.ToLower(strings.TrimSpace(scanner.Text()))
        switch command {
        case "new":
            fmt.Print("Enter world width: ")
            scanner.Scan()
            width, _ := strconv.Atoi(scanner.Text())
            fmt.Print("Enter world height: ")
            scanner.Scan()
            height, _ := strconv.Atoi(scanner.Text())
            gameWorld = NewGameWorld(width, height)
            fmt.Print("Enter player name: ")
            scanner.Scan()
            name := scanner.Text()
            player := NewPlayer(playerIDCounter, name)
            gameWorld.AddPlayer(player)
            playerIDCounter++
            fmt.Println("New game created with player", name)
        case "load":
            fmt.Print("Enter filename to load: ")
            scanner.Scan()
            filename := scanner.Text()
            loadedWorld, err := LoadGame(filename)
            if err != nil {
                fmt.Println("Error loading game:", err)
            } else {
                gameWorld = loadedWorld
                fmt.Println("Game loaded from", filename)
            }
        case "move":
            if gameWorld == nil {
                fmt.Println("No game world. Use 'new' or 'load' first.")
                break
            }
            fmt.Print("Enter player ID to move: ")
            scanner.Scan()
            id, _ := strconv.Atoi(scanner.Text())
            fmt.Print("Enter direction (up, down, left, right): ")
            scanner.Scan()
            dir := scanner.Text()
            var dx, dy int
            switch dir {
            case "up":
                dy = -1
            case "down":
                dy = 1
            case "left":
                dx = -1
            case "right":
                dx = 1
            default:
                fmt.Println("Invalid direction")
                break
            }
            if gameWorld.MovePlayer(id, dx, dy) {
                fmt.Println("Player moved successfully")
            } else {
                fmt.Println("Move failed")
            }
            gameWorld.Display()
        case "battle":
            if gameWorld == nil {
                fmt.Println("No game world. Use 'new' or 'load' first.")
                break
            }
            fmt.Print("Enter player ID for battle: ")
            scanner.Scan()
            playerID, _ := strconv.Atoi(scanner.Text())
            var player *Player
            for i, p := range gameWorld.Players {
                if p.ID == playerID {
                    player = &gameWorld.Players[i]
                    break
                }
            }
            if player == nil {
                fmt.Println("Player not found")
                break
            }
            monster := NewMonster(monsterIDCounter, "Goblin", 50, 8, 3, Reward{Experience: 30, Gold: 20, Items: []string{"Goblin Ear"}})
            monsterIDCounter++
            if Battle(player, &monster) {
                fmt.Println("Battle won!")
            } else {
                fmt.Println("Battle lost!")
                gameWorld.RemovePlayer(player.ID)
            }
        case "save":
            if gameWorld == nil {
                fmt.Println("No game world to save.")
                break
            }
            fmt.Print("Enter filename to save: ")
            scanner.Scan()
            filename := scanner.Text()
            err := SaveGame(gameWorld, filename)
            if err != nil {
                fmt.Println("Error saving game:", err)
            } else {
                fmt.Println("Game saved to", filename)
            }
        case "quit":
            fmt.Println("Goodbye!")
            return
        default:
            fmt.Println("Unknown command. Try: new, load, move, battle, save, quit")
        }
    }
}
