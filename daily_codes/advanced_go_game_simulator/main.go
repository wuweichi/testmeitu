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
    Inventory   []string
    Gold        int
    IsAlive     bool
}

type Monster struct {
    ID        int
    Name      string
    Health    int
    Attack    int
    Defense   int
    RewardExp int
    RewardGold int
}

type GameState struct {
    Players   []Player
    Monsters  []Monster
    GameRound int
    GameOver  bool
    Mutex     sync.Mutex
}

func (p *Player) TakeDamage(damage int) {
    actualDamage := damage - p.Defense
    if actualDamage < 0 {
        actualDamage = 0
    }
    p.Health -= actualDamage
    if p.Health <= 0 {
        p.Health = 0
        p.IsAlive = false
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
        p.MaxHealth += 10
        p.Attack += 2
        p.Defense += 1
        p.Health = p.MaxHealth
        fmt.Printf("Player %s leveled up to level %d!\n", p.Name, p.Level)
    }
}

func (p *Player) AddItem(item string) {
    p.Inventory = append(p.Inventory, item)
}

func (p *Player) UseItem(itemIndex int) bool {
    if itemIndex < 0 || itemIndex >= len(p.Inventory) {
        return false
    }
    item := p.Inventory[itemIndex]
    switch item {
    case "Health Potion":
        p.Heal(50)
        fmt.Printf("Player %s used a Health Potion and healed 50 HP.\n", p.Name)
    case "Attack Boost":
        p.Attack += 5
        fmt.Printf("Player %s used an Attack Boost, increasing attack by 5.\n", p.Name)
    default:
        fmt.Printf("Item %s has no effect.\n", item)
    }
    p.Inventory = append(p.Inventory[:itemIndex], p.Inventory[itemIndex+1:]...)
    return true
}

func (m *Monster) TakeDamage(damage int) bool {
    actualDamage := damage - m.Defense
    if actualDamage < 0 {
        actualDamage = 0
    }
    m.Health -= actualDamage
    return m.Health <= 0
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
        Inventory:  []string{"Health Potion", "Health Potion"},
        Gold:       50,
        IsAlive:    true,
    }
}

func NewMonster(id int, name string, health, attack, defense, exp, gold int) Monster {
    return Monster{
        ID:         id,
        Name:       name,
        Health:     health,
        Attack:     attack,
        Defense:    defense,
        RewardExp:  exp,
        RewardGold: gold,
    }
}

func InitializeGame() *GameState {
    rand.Seed(time.Now().UnixNano())
    players := []Player{
        NewPlayer(1, "Hero"),
        NewPlayer(2, "Mage"),
        NewPlayer(3, "Warrior"),
    }
    monsters := []Monster{
        NewMonster(1, "Goblin", 30, 8, 2, 20, 10),
        NewMonster(2, "Orc", 50, 12, 5, 40, 20),
        NewMonster(3, "Dragon", 100, 20, 10, 100, 50),
        NewMonster(4, "Slime", 20, 5, 1, 10, 5),
        NewMonster(5, "Skeleton", 40, 10, 3, 30, 15),
    }
    return &GameState{
        Players:   players,
        Monsters:  monsters,
        GameRound: 1,
        GameOver:  false,
        Mutex:     sync.Mutex{},
    }
}

func (gs *GameState) DisplayStatus() {
    gs.Mutex.Lock()
    defer gs.Mutex.Unlock()
    fmt.Println("\n=== Game Status ===")
    fmt.Printf("Round: %d\n", gs.GameRound)
    fmt.Println("Players:")
    for _, p := range gs.Players {
        status := "Alive"
        if !p.IsAlive {
            status = "Dead"
        }
        fmt.Printf("  %s (ID: %d) - Level: %d, HP: %d/%d, Attack: %d, Defense: %d, Exp: %d, Gold: %d, Status: %s\n",
            p.Name, p.ID, p.Level, p.Health, p.MaxHealth, p.Attack, p.Defense, p.Experience, p.Gold, status)
        fmt.Printf("    Inventory: %v\n", p.Inventory)
    }
    fmt.Println("Monsters:")
    for _, m := range gs.Monsters {
        status := "Alive"
        if m.Health <= 0 {
            status = "Dead"
        }
        fmt.Printf("  %s (ID: %d) - HP: %d, Attack: %d, Defense: %d\n", m.Name, m.ID, m.Health, m.Attack, m.Defense)
    }
    fmt.Println("===================")
}

func (gs *GameState) PlayerAction(playerID int) {
    gs.Mutex.Lock()
    defer gs.Mutex.Unlock()
    var player *Player
    for i := range gs.Players {
        if gs.Players[i].ID == playerID && gs.Players[i].IsAlive {
            player = &gs.Players[i]
            break
        }
    }
    if player == nil {
        return
    }
    fmt.Printf("\nPlayer %s's turn (ID: %d):\n", player.Name, player.ID)
    fmt.Println("1. Attack a monster")
    fmt.Println("2. Use an item")
    fmt.Println("3. Do nothing")
    fmt.Print("Choose an action (1-3): ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    switch input {
    case "1":
        gs.AttackMonster(player)
    case "2":
        gs.UseItem(player)
    case "3":
        fmt.Printf("Player %s does nothing.\n", player.Name)
    default:
        fmt.Println("Invalid choice, doing nothing.")
    }
}

func (gs *GameState) AttackMonster(player *Player) {
    fmt.Println("Available monsters to attack:")
    for _, m := range gs.Monsters {
        if m.Health > 0 {
            fmt.Printf("  ID: %d - %s (HP: %d)\n", m.ID, m.Name, m.Health)
        }
    }
    fmt.Print("Enter monster ID to attack: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    monsterID, err := strconv.Atoi(input)
    if err != nil {
        fmt.Println("Invalid ID, attack failed.")
        return
    }
    var targetMonster *Monster
    for i := range gs.Monsters {
        if gs.Monsters[i].ID == monsterID && gs.Monsters[i].Health > 0 {
            targetMonster = &gs.Monsters[i]
            break
        }
    }
    if targetMonster == nil {
        fmt.Println("Monster not found or already dead, attack failed.")
        return
    }
    damage := player.Attack + rand.Intn(5)
    fmt.Printf("Player %s attacks %s for %d damage!\n", player.Name, targetMonster.Name, damage)
    if targetMonster.TakeDamage(damage) {
        fmt.Printf("%s has been defeated!\n", targetMonster.Name)
        player.GainExperience(targetMonster.RewardExp)
        player.Gold += targetMonster.RewardGold
        fmt.Printf("Player %s gains %d experience and %d gold.\n", player.Name, targetMonster.RewardExp, targetMonster.RewardGold)
    } else {
        fmt.Printf("%s now has %d HP.\n", targetMonster.Name, targetMonster.Health)
    }
}

func (gs *GameState) UseItem(player *Player) {
    if len(player.Inventory) == 0 {
        fmt.Println("Inventory is empty.")
        return
    }
    fmt.Println("Inventory items:")
    for i, item := range player.Inventory {
        fmt.Printf("  %d. %s\n", i, item)
    }
    fmt.Print("Enter item number to use: ")
    reader := bufio.NewReader(os.Stdin)
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)
    itemIndex, err := strconv.Atoi(input)
    if err != nil || !player.UseItem(itemIndex) {
        fmt.Println("Invalid item number, use failed.")
    }
}

func (gs *GameState) MonsterTurn() {
    gs.Mutex.Lock()
    defer gs.Mutex.Unlock()
    fmt.Println("\n--- Monster Turn ---")
    for i := range gs.Monsters {
        monster := &gs.Monsters[i]
        if monster.Health > 0 {
            alivePlayers := []*Player{}
            for j := range gs.Players {
                if gs.Players[j].IsAlive {
                    alivePlayers = append(alivePlayers, &gs.Players[j])
                }
            }
            if len(alivePlayers) == 0 {
                return
            }
            target := alivePlayers[rand.Intn(len(alivePlayers))]
            damage := monster.Attack + rand.Intn(3)
            fmt.Printf("%s attacks %s for %d damage!\n", monster.Name, target.Name, damage)
            target.TakeDamage(damage)
            if !target.IsAlive {
                fmt.Printf("Player %s has been defeated!\n", target.Name)
            }
        }
    }
}

func (gs *GameState) CheckGameOver() bool {
    gs.Mutex.Lock()
    defer gs.Mutex.Unlock()
    allPlayersDead := true
    for _, p := range gs.Players {
        if p.IsAlive {
            allPlayersDead = false
            break
        }
    }
    allMonstersDead := true
    for _, m := range gs.Monsters {
        if m.Health > 0 {
            allMonstersDead = false
            break
        }
    }
    if allPlayersDead {
        fmt.Println("\nAll players are dead! Game Over.")
        gs.GameOver = true
        return true
    }
    if allMonstersDead {
        fmt.Println("\nAll monsters are defeated! You Win!")
        gs.GameOver = true
        return true
    }
    return false
}

func (gs *GameState) SaveGame(filename string) error {
    gs.Mutex.Lock()
    defer gs.Mutex.Unlock()
    data, err := json.MarshalIndent(gs, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}

func (gs *GameState) LoadGame(filename string) error {
    gs.Mutex.Lock()
    defer gs.Mutex.Unlock()
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, gs)
}

func (gs *GameState) RunGame() {
    fmt.Println("Welcome to the Advanced Go Game Simulator!")
    fmt.Println("Commands: status, play, save <filename>, load <filename>, quit")
    reader := bufio.NewReader(os.Stdin)
    for !gs.GameOver {
        fmt.Print("\nEnter command: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        parts := strings.Fields(input)
        if len(parts) == 0 {
            continue
        }
        switch parts[0] {
        case "status":
            gs.DisplayStatus()
        case "play":
            gs.PlayRound()
        case "save":
            if len(parts) < 2 {
                fmt.Println("Usage: save <filename>")
            } else {
                if err := gs.SaveGame(parts[1]); err != nil {
                    fmt.Printf("Error saving game: %v\n", err)
                } else {
                    fmt.Println("Game saved successfully.")
                }
            }
        case "load":
            if len(parts) < 2 {
                fmt.Println("Usage: load <filename>")
            } else {
                if err := gs.LoadGame(parts[1]); err != nil {
                    fmt.Printf("Error loading game: %v\n", err)
                } else {
                    fmt.Println("Game loaded successfully.")
                }
            }
        case "quit":
            fmt.Println("Quitting game.")
            return
        default:
            fmt.Println("Unknown command. Available: status, play, save, load, quit")
        }
    }
}

func (gs *GameState) PlayRound() {
    fmt.Printf("\n--- Starting Round %d ---\n", gs.GameRound)
    for _, player := range gs.Players {
        if player.IsAlive {
            gs.PlayerAction(player.ID)
        }
    }
    gs.MonsterTurn()
    if gs.CheckGameOver() {
        return
    }
    gs.GameRound++
    fmt.Printf("--- End of Round %d ---\n", gs.GameRound-1)
}

func main() {
    game := InitializeGame()
    game.RunGame()
}
