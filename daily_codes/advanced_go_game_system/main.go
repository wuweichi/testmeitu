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
    Inventory   []Item
    Skills      []Skill
    Gold        int
    Position    Coordinate
}

type Item struct {
    ID          int
    Name        string
    Description string
    Type        string // weapon, armor, potion, etc.
    Value       int
    AttackBonus int
    DefenseBonus int
    HealthBonus int
}

type Skill struct {
    ID          int
    Name        string
    Description string
    Damage      int
    Cost        int // mana or stamina cost
    Cooldown    int
}

type Coordinate struct {
    X int
    Y int
}

type GameWorld struct {
    Width   int
    Height  int
    Map     [][]string
    Monsters []Monster
    Items   []Item
}

type Monster struct {
    ID          int
    Name        string
    Health      int
    Attack      int
    Defense     int
    Experience  int
    Drops       []Item
}

type GameState struct {
    Player      Player
    World       GameWorld
    Turn        int
    GameOver    bool
    MessageLog  []string
    mu          sync.Mutex
}

func NewPlayer(id int, name string) Player {
    return Player{
        ID:         id,
        Name:       name,
        Health:     100,
        MaxHealth:  100,
        Attack:     10,
        Defense:    5,
        Level:      1,
        Experience: 0,
        Inventory:  []Item{},
        Skills:     []Skill{},
        Gold:       50,
        Position:   Coordinate{X: 0, Y: 0},
    }
}

func NewGameWorld(width, height int) GameWorld {
    world := GameWorld{
        Width:   width,
        Height:  height,
        Map:     make([][]string, height),
        Monsters: []Monster{},
        Items:   []Item{},
    }
    for i := range world.Map {
        world.Map[i] = make([]string, width)
        for j := range world.Map[i] {
            world.Map[i][j] = "."
        }
    }
    return world
}

func (gs *GameState) AddMessage(msg string) {
    gs.mu.Lock()
    defer gs.mu.Unlock()
    gs.MessageLog = append(gs.MessageLog, msg)
    if len(gs.MessageLog) > 10 {
        gs.MessageLog = gs.MessageLog[1:]
    }
}

func (gs *GameState) DisplayMessages() {
    fmt.Println("\n--- Message Log ---")
    for _, msg := range gs.MessageLog {
        fmt.Println(msg)
    }
    fmt.Println("-------------------")
}

func (gs *GameState) MovePlayer(dx, dy int) {
    newX := gs.Player.Position.X + dx
    newY := gs.Player.Position.Y + dy
    if newX >= 0 && newX < gs.World.Width && newY >= 0 && newY < gs.World.Height {
        gs.Player.Position.X = newX
        gs.Player.Position.Y = newY
        gs.AddMessage(fmt.Sprintf("Moved to (%d, %d)", newX, newY))
        gs.CheckForEvents()
    } else {
        gs.AddMessage("Cannot move outside the map!")
    }
}

func (gs *GameState) CheckForEvents() {
    // Check for monsters
    for i, monster := range gs.World.Monsters {
        if monster.Health > 0 && gs.Player.Position.X == i%gs.World.Width && gs.Player.Position.Y == i/gs.World.Width {
            gs.AddMessage(fmt.Sprintf("Encountered a %s!", monster.Name))
            gs.Combat(&gs.World.Monsters[i])
            return
        }
    }
    // Check for items
    for i, item := range gs.World.Items {
        if gs.Player.Position.X == i%gs.World.Width && gs.Player.Position.Y == i/gs.World.Width {
            gs.AddMessage(fmt.Sprintf("Found a %s!", item.Name))
            gs.Player.Inventory = append(gs.Player.Inventory, item)
            gs.World.Items = append(gs.World.Items[:i], gs.World.Items[i+1:]...)
            return
        }
    }
}

func (gs *GameState) Combat(monster *Monster) {
    gs.AddMessage(fmt.Sprintf("Combat started with %s", monster.Name))
    for gs.Player.Health > 0 && monster.Health > 0 {
        // Player attack
        damage := gs.Player.Attack - monster.Defense
        if damage < 1 {
            damage = 1
        }
        monster.Health -= damage
        gs.AddMessage(fmt.Sprintf("You hit %s for %d damage", monster.Name, damage))
        if monster.Health <= 0 {
            gs.AddMessage(fmt.Sprintf("You defeated %s!", monster.Name))
            gs.Player.Experience += monster.Experience
            gs.AddMessage(fmt.Sprintf("Gained %d experience", monster.Experience))
            // Drop items
            for _, drop := range monster.Drops {
                gs.Player.Inventory = append(gs.Player.Inventory, drop)
                gs.AddMessage(fmt.Sprintf("Obtained %s", drop.Name))
            }
            gs.CheckLevelUp()
            return
        }
        // Monster attack
        damage = monster.Attack - gs.Player.Defense
        if damage < 1 {
            damage = 1
        }
        gs.Player.Health -= damage
        gs.AddMessage(fmt.Sprintf("%s hit you for %d damage", monster.Name, damage))
        if gs.Player.Health <= 0 {
            gs.AddMessage("You have been defeated!")
            gs.GameOver = true
            return
        }
    }
}

func (gs *GameState) CheckLevelUp() {
    expNeeded := gs.Player.Level * 100
    if gs.Player.Experience >= expNeeded {
        gs.Player.Level++
        gs.Player.Experience -= expNeeded
        gs.Player.MaxHealth += 20
        gs.Player.Health = gs.Player.MaxHealth
        gs.Player.Attack += 5
        gs.Player.Defense += 3
        gs.AddMessage(fmt.Sprintf("Level up! Now level %d", gs.Player.Level))
    }
}

func (gs *GameState) UseItem(itemIndex int) {
    if itemIndex < 0 || itemIndex >= len(gs.Player.Inventory) {
        gs.AddMessage("Invalid item index")
        return
    }
    item := gs.Player.Inventory[itemIndex]
    switch item.Type {
    case "potion":
        gs.Player.Health += item.HealthBonus
        if gs.Player.Health > gs.Player.MaxHealth {
            gs.Player.Health = gs.Player.MaxHealth
        }
        gs.AddMessage(fmt.Sprintf("Used %s, healed %d HP", item.Name, item.HealthBonus))
    case "weapon":
        gs.Player.Attack += item.AttackBonus
        gs.AddMessage(fmt.Sprintf("Equipped %s, attack increased by %d", item.Name, item.AttackBonus))
    case "armor":
        gs.Player.Defense += item.DefenseBonus
        gs.AddMessage(fmt.Sprintf("Equipped %s, defense increased by %d", item.Name, item.DefenseBonus))
    default:
        gs.AddMessage("Cannot use this item")
        return
    }
    gs.Player.Inventory = append(gs.Player.Inventory[:itemIndex], gs.Player.Inventory[itemIndex+1:]...)
}

func (gs *GameState) DisplayStatus() {
    fmt.Printf("\n--- Status ---\n")
    fmt.Printf("Name: %s\n", gs.Player.Name)
    fmt.Printf("Health: %d/%d\n", gs.Player.Health, gs.Player.MaxHealth)
    fmt.Printf("Level: %d\n", gs.Player.Level)
    fmt.Printf("Experience: %d/%d\n", gs.Player.Experience, gs.Player.Level*100)
    fmt.Printf("Attack: %d\n", gs.Player.Attack)
    fmt.Printf("Defense: %d\n", gs.Player.Defense)
    fmt.Printf("Gold: %d\n", gs.Player.Gold)
    fmt.Printf("Position: (%d, %d)\n", gs.Player.Position.X, gs.Player.Position.Y)
    fmt.Printf("---------------\n")
}

func (gs *GameState) DisplayInventory() {
    fmt.Println("\n--- Inventory ---")
    if len(gs.Player.Inventory) == 0 {
        fmt.Println("Empty")
    } else {
        for i, item := range gs.Player.Inventory {
            fmt.Printf("%d: %s - %s\n", i, item.Name, item.Description)
        }
    }
    fmt.Println("-----------------")
}

func (gs *GameState) DisplayMap() {
    fmt.Println("\n--- Map ---")
    for y := 0; y < gs.World.Height; y++ {
        for x := 0; x < gs.World.Width; x++ {
            if x == gs.Player.Position.X && y == gs.Player.Position.Y {
                fmt.Print("P ")
            } else {
                fmt.Print(gs.World.Map[y][x], " ")
            }
        }
        fmt.Println()
    }
    fmt.Println("------------")
}

func (gs *GameState) SaveGame(filename string) error {
    data, err := json.Marshal(gs)
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}

func (gs *GameState) LoadGame(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, gs)
}

func InitializeGame() *GameState {
    rand.Seed(time.Now().UnixNano())
    player := NewPlayer(1, "Hero")
    world := NewGameWorld(10, 10)
    // Add some monsters
    monsters := []Monster{
        {ID: 1, Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Experience: 20, Drops: []Item{{ID: 101, Name: "Goblin Ear", Description: "A trophy from a goblin", Type: "misc", Value: 5}}},
        {ID: 2, Name: "Orc", Health: 50, Attack: 8, Defense: 4, Experience: 40, Drops: []Item{{ID: 102, Name: "Orc Tooth", Description: "A sharp orc tooth", Type: "misc", Value: 10}}},
        {ID: 3, Name: "Dragon", Health: 100, Attack: 15, Defense: 10, Experience: 100, Drops: []Item{{ID: 103, Name: "Dragon Scale", Description: "A shiny dragon scale", Type: "armor", DefenseBonus: 5, Value: 50}}},
    }
    for i := range monsters {
        monsters[i].Health = monsters[i].Health // Ensure health is set
    }
    world.Monsters = monsters
    // Add some items
    items := []Item{
        {ID: 201, Name: "Health Potion", Description: "Restores 20 HP", Type: "potion", HealthBonus: 20, Value: 10},
        {ID: 202, Name: "Iron Sword", Description: "A basic sword", Type: "weapon", AttackBonus: 5, Value: 30},
        {ID: 203, Name: "Leather Armor", Description: "Light armor", Type: "armor", DefenseBonus: 3, Value: 25},
    }
    world.Items = items
    // Randomly place monsters and items on the map
    for i := range world.Monsters {
        world.Map[i/world.Width][i%world.Width] = "M"
    }
    for i := range world.Items {
        pos := len(world.Monsters) + i
        world.Map[pos/world.Width][pos%world.Width] = "I"
    }
    return &GameState{
        Player:     player,
        World:      world,
        Turn:       0,
        GameOver:   false,
        MessageLog: []string{"Welcome to the Advanced Go Game System!"},
    }
}

func main() {
    game := InitializeGame()
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("Advanced Go Game System")
    fmt.Println("Commands: move [n/s/e/w], status, inventory, map, use [index], save [filename], load [filename], quit")
    for !game.GameOver {
        game.DisplayStatus()
        game.DisplayMap()
        game.DisplayMessages()
        fmt.Print("> ")
        if !scanner.Scan() {
            break
        }
        input := scanner.Text()
        parts := strings.Fields(input)
        if len(parts) == 0 {
            continue
        }
        command := parts[0]
        switch command {
        case "move":
            if len(parts) < 2 {
                game.AddMessage("Usage: move [n/s/e/w]")
                continue
            }
            direction := parts[1]
            switch direction {
            case "n":
                game.MovePlayer(0, -1)
            case "s":
                game.MovePlayer(0, 1)
            case "e":
                game.MovePlayer(1, 0)
            case "w":
                game.MovePlayer(-1, 0)
            default:
                game.AddMessage("Invalid direction. Use n, s, e, w")
            }
        case "status":
            game.DisplayStatus()
        case "inventory":
            game.DisplayInventory()
        case "map":
            game.DisplayMap()
        case "use":
            if len(parts) < 2 {
                game.AddMessage("Usage: use [item index]")
                continue
            }
            index, err := strconv.Atoi(parts[1])
            if err != nil {
                game.AddMessage("Invalid index")
                continue
            }
            game.UseItem(index)
        case "save":
            if len(parts) < 2 {
                game.AddMessage("Usage: save [filename]")
                continue
            }
            err := game.SaveGame(parts[1])
            if err != nil {
                game.AddMessage(fmt.Sprintf("Save failed: %v", err))
            } else {
                game.AddMessage("Game saved successfully")
            }
        case "load":
            if len(parts) < 2 {
                game.AddMessage("Usage: load [filename]")
                continue
            }
            err := game.LoadGame(parts[1])
            if err != nil {
                game.AddMessage(fmt.Sprintf("Load failed: %v", err))
            } else {
                game.AddMessage("Game loaded successfully")
            }
        case "quit":
            fmt.Println("Thanks for playing!")
            return
        default:
            game.AddMessage("Unknown command")
        }
        game.Turn++
    }
    fmt.Println("Game Over!")
}
