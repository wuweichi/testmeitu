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
    Inventory []Item
    Gold      int
    Position  Point
}

type Item struct {
    ID          int
    Name        string
    Description string
    Type        string // "weapon", "armor", "potion", "key"
    Value       int
    Durability  int
}

type Monster struct {
    ID        int
    Name      string
    Health    int
    Attack    int
    Defense   int
    Experience int
    Drops     []Item
}

type Point struct {
    X int
    Y int
}

type GameMap struct {
    Width   int
    Height  int
    Tiles   [][]Tile
    Monsters []Monster
    Chests  []Chest
}

type Tile struct {
    Type      string // "grass", "water", "mountain", "forest"
    Passable  bool
    Symbol    string
}

type Chest struct {
    Position Point
    Locked   bool
    Items    []Item
}

type GameState struct {
    Player    Player
    Map       GameMap
    Turn      int
    GameOver  bool
    Messages  []string
    mu        sync.Mutex
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
        Inventory: []Item{
            {ID: 1, Name: "Wooden Sword", Description: "A basic sword", Type: "weapon", Value: 5, Durability: 20},
            {ID: 2, Name: "Leather Armor", Description: "Light armor", Type: "armor", Value: 3, Durability: 30},
            {ID: 3, Name: "Health Potion", Description: "Restores 50 health", Type: "potion", Value: 20, Durability: 1},
        },
        Gold:     50,
        Position: Point{X: 0, Y: 0},
    }
}

func NewMonster(id int, name string, health, attack, defense, exp int) Monster {
    drops := []Item{}
    if rand.Intn(100) < 30 {
        drops = append(drops, Item{ID: 4, Name: "Monster Claw", Description: "A sharp claw", Type: "key", Value: 10, Durability: 1})
    }
    if rand.Intn(100) < 20 {
        drops = append(drops, Item{ID: 5, Name: "Gold Coin", Description: "Worth 10 gold", Type: "potion", Value: 10, Durability: 1})
    }
    return Monster{
        ID:        id,
        Name:      name,
        Health:    health,
        Attack:    attack,
        Defense:   defense,
        Experience: exp,
        Drops:     drops,
    }
}

func NewGameMap(width, height int) GameMap {
    tiles := make([][]Tile, height)
    for y := 0; y < height; y++ {
        tiles[y] = make([]Tile, width)
        for x := 0; x < width; x++ {
            tileType := "grass"
            passable := true
            symbol := "."
            r := rand.Intn(100)
            if r < 10 {
                tileType = "water"
                passable = false
                symbol = "~"
            } else if r < 20 {
                tileType = "mountain"
                passable = false
                symbol = "^"
            } else if r < 40 {
                tileType = "forest"
                symbol = "*"
            }
            tiles[y][x] = Tile{Type: tileType, Passable: passable, Symbol: symbol}
        }
    }
    monsters := []Monster{}
    for i := 0; i < 10; i++ {
        x := rand.Intn(width)
        y := rand.Intn(height)
        if tiles[y][x].Passable {
            monster := NewMonster(i+1, fmt.Sprintf("Goblin %d", i+1), 30, 8, 2, 20)
            monsters = append(monsters, monster)
        }
    }
    chests := []Chest{}
    for i := 0; i < 5; i++ {
        x := rand.Intn(width)
        y := rand.Intn(height)
        if tiles[y][x].Passable {
            items := []Item{
                {ID: 6, Name: "Iron Sword", Description: "A stronger sword", Type: "weapon", Value: 15, Durability: 40},
                {ID: 7, Name: "Steel Armor", Description: "Heavy armor", Type: "armor", Value: 10, Durability: 50},
            }
            chests = append(chests, Chest{Position: Point{X: x, Y: y}, Locked: rand.Intn(100) < 50, Items: items})
        }
    }
    return GameMap{Width: width, Height: height, Tiles: tiles, Monsters: monsters, Chests: chests}
}

func (p *Player) Move(dx, dy int, m *GameMap) bool {
    newX := p.Position.X + dx
    newY := p.Position.Y + dy
    if newX < 0 || newX >= m.Width || newY < 0 || newY >= m.Height {
        return false
    }
    if !m.Tiles[newY][newX].Passable {
        return false
    }
    p.Position.X = newX
    p.Position.Y = newY
    return true
}

func (p *Player) AttackMonster(m *Monster) (bool, int) {
    damage := p.Attack - m.Defense
    if damage < 1 {
        damage = 1
    }
    m.Health -= damage
    if m.Health <= 0 {
        p.Experience += m.Experience
        p.Gold += rand.Intn(20) + 5
        return true, damage
    }
    return false, damage
}

func (m *Monster) AttackPlayer(p *Player) (bool, int) {
    damage := m.Attack - p.Defense
    if damage < 1 {
        damage = 1
    }
    p.Health -= damage
    if p.Health <= 0 {
        return true, damage
    }
    return false, damage
}

func (p *Player) LevelUp() {
    if p.Experience >= p.Level*100 {
        p.Level++
        p.MaxHealth += 20
        p.Health = p.MaxHealth
        p.Attack += 5
        p.Defense += 3
        p.Experience = 0
    }
}

func (p *Player) UseItem(itemIndex int) bool {
    if itemIndex < 0 || itemIndex >= len(p.Inventory) {
        return false
    }
    item := &p.Inventory[itemIndex]
    if item.Type == "potion" {
        p.Health += 50
        if p.Health > p.MaxHealth {
            p.Health = p.MaxHealth
        }
        item.Durability--
        if item.Durability <= 0 {
            p.Inventory = append(p.Inventory[:itemIndex], p.Inventory[itemIndex+1:]...)
        }
        return true
    }
    return false
}

func (gs *GameState) AddMessage(msg string) {
    gs.mu.Lock()
    defer gs.mu.Unlock()
    gs.Messages = append(gs.Messages, msg)
    if len(gs.Messages) > 10 {
        gs.Messages = gs.Messages[1:]
    }
}

func (gs *GameState) Display() {
    fmt.Println("=== Complex Game Simulator ===")
    fmt.Printf("Turn: %d | Player: %s (Level %d) | Health: %d/%d | Gold: %d | Exp: %d\n",
        gs.Turn, gs.Player.Name, gs.Player.Level, gs.Player.Health, gs.Player.MaxHealth, gs.Player.Gold, gs.Player.Experience)
    fmt.Println("Position:", gs.Player.Position)
    fmt.Println("Inventory:")
    for i, item := range gs.Player.Inventory {
        fmt.Printf("  %d: %s (%s) - %s\n", i, item.Name, item.Type, item.Description)
    }
    fmt.Println("Map (P=Player, M=Monster, C=Chest):")
    for y := 0; y < gs.Map.Height; y++ {
        for x := 0; x < gs.Map.Width; x++ {
            symbol := gs.Map.Tiles[y][x].Symbol
            if gs.Player.Position.X == x && gs.Player.Position.Y == y {
                symbol = "P"
            } else {
                for _, monster := range gs.Map.Monsters {
                    // Simplified: monsters not displayed on map for brevity
                }
                for _, chest := range gs.Map.Chests {
                    if chest.Position.X == x && chest.Position.Y == y {
                        symbol = "C"
                    }
                }
            }
            fmt.Print(symbol)
        }
        fmt.Println()
    }
    fmt.Println("Recent Messages:")
    for _, msg := range gs.Messages {
        fmt.Println("  ", msg)
    }
    fmt.Println("Commands: move [n/s/e/w], attack, use [item_index], save, load, quit")
}

func (gs *GameState) Save(filename string) error {
    data, err := json.Marshal(gs)
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}

func (gs *GameState) Load(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, gs)
}

func (gs *GameState) ProcessCommand(cmd string) {
    parts := strings.Fields(cmd)
    if len(parts) == 0 {
        gs.AddMessage("Invalid command")
        return
    }
    switch parts[0] {
    case "move":
        if len(parts) < 2 {
            gs.AddMessage("Usage: move [n/s/e/w]")
            return
        }
        dx, dy := 0, 0
        switch parts[1] {
        case "n":
            dy = -1
        case "s":
            dy = 1
        case "e":
            dx = 1
        case "w":
            dx = -1
        default:
            gs.AddMessage("Invalid direction")
            return
        }
        if gs.Player.Move(dx, dy, &gs.Map) {
            gs.AddMessage(fmt.Sprintf("Moved to (%d, %d)", gs.Player.Position.X, gs.Player.Position.Y))
            // Check for encounters
            for i, monster := range gs.Map.Monsters {
                // Simplified: combat not triggered by movement in this version
            }
            for i, chest := range gs.Map.Chests {
                if chest.Position.X == gs.Player.Position.X && chest.Position.Y == gs.Player.Position.Y {
                    gs.AddMessage("You found a chest!")
                    if chest.Locked {
                        gs.AddMessage("Chest is locked. Need a key.")
                    } else {
                        gs.AddMessage("Chest opened. Items added to inventory.")
                        gs.Player.Inventory = append(gs.Player.Inventory, chest.Items...)
                        gs.Map.Chests = append(gs.Map.Chests[:i], gs.Map.Chests[i+1:]...)
                    }
                }
            }
        } else {
            gs.AddMessage("Cannot move there")
        }
    case "attack":
        // Simplified: attack a monster at current position
        gs.AddMessage("No monster here to attack")
    case "use":
        if len(parts) < 2 {
            gs.AddMessage("Usage: use [item_index]")
            return
        }
        index, err := strconv.Atoi(parts[1])
        if err != nil {
            gs.AddMessage("Invalid item index")
            return
        }
        if gs.Player.UseItem(index) {
            gs.AddMessage("Item used")
        } else {
            gs.AddMessage("Cannot use that item")
        }
    case "save":
        err := gs.Save("savegame.json")
        if err != nil {
            gs.AddMessage(fmt.Sprintf("Save failed: %v", err))
        } else {
            gs.AddMessage("Game saved")
        }
    case "load":
        err := gs.Load("savegame.json")
        if err != nil {
            gs.AddMessage(fmt.Sprintf("Load failed: %v", err))
        } else {
            gs.AddMessage("Game loaded")
        }
    case "quit":
        gs.GameOver = true
        gs.AddMessage("Thanks for playing!")
    default:
        gs.AddMessage("Unknown command")
    }
    gs.Turn++
    gs.Player.LevelUp()
}

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("Welcome to Complex Game Simulator!")
    fmt.Print("Enter your name: ")
    reader := bufio.NewReader(os.Stdin)
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)
    if name == "" {
        name = "Hero"
    }
    player := NewPlayer(1, name)
    gameMap := NewGameMap(20, 10)
    gameState := GameState{
        Player:   player,
        Map:      gameMap,
        Turn:     1,
        GameOver: false,
        Messages: []string{"Game started!"},
    }
    for !gameState.GameOver {
        gameState.Display()
        fmt.Print("> ")
        cmd, _ := reader.ReadString('\n')
        cmd = strings.TrimSpace(cmd)
        gameState.ProcessCommand(cmd)
    }
}
