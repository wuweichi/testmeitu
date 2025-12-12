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

type Character struct {
    Name      string
    Health    int
    MaxHealth int
    Strength  int
    Agility   int
    Level     int
    Experience int
    Inventory []Item
    Gold      int
}

type Item struct {
    Name        string
    Description string
    Value       int
    Type        string
}

type Monster struct {
    Name      string
    Health    int
    Strength  int
    Agility   int
    ExperienceReward int
    GoldReward int
}

type GameState struct {
    Player      Character
    Monsters    []Monster
    CurrentZone string
    Day         int
    GameActive  bool
    mu          sync.Mutex
}

func (c *Character) Attack(target *Monster) (int, bool) {
    damage := c.Strength + rand.Intn(10)
    target.Health -= damage
    if target.Health <= 0 {
        c.Experience += target.ExperienceReward
        c.Gold += target.GoldReward
        return damage, true
    }
    return damage, false
}

func (m *Monster) Attack(target *Character) int {
    damage := m.Strength + rand.Intn(5)
    target.Health -= damage
    return damage
}

func (c *Character) Heal(amount int) {
    c.Health += amount
    if c.Health > c.MaxHealth {
        c.Health = c.MaxHealth
    }
}

func (c *Character) LevelUp() {
    if c.Experience >= c.Level*100 {
        c.Level++
        c.MaxHealth += 20
        c.Health = c.MaxHealth
        c.Strength += 5
        c.Agility += 3
        c.Experience = 0
        fmt.Printf("\n%s leveled up to level %d!\n", c.Name, c.Level)
    }
}

func (c *Character) AddItem(item Item) {
    c.Inventory = append(c.Inventory, item)
}

func (c *Character) RemoveItem(index int) {
    if index >= 0 && index < len(c.Inventory) {
        c.Inventory = append(c.Inventory[:index], c.Inventory[index+1:]...)
    }
}

func (c *Character) ShowStatus() {
    fmt.Printf("\n=== %s Status ===\n", c.Name)
    fmt.Printf("Health: %d/%d\n", c.Health, c.MaxHealth)
    fmt.Printf("Strength: %d\n", c.Strength)
    fmt.Printf("Agility: %d\n", c.Agility)
    fmt.Printf("Level: %d\n", c.Level)
    fmt.Printf("Experience: %d/%d\n", c.Experience, c.Level*100)
    fmt.Printf("Gold: %d\n", c.Gold)
    fmt.Println("Inventory:")
    for i, item := range c.Inventory {
        fmt.Printf("  %d. %s - %s (Value: %d)\n", i+1, item.Name, item.Description, item.Value)
    }
}

func (gs *GameState) SaveGame(filename string) error {
    gs.mu.Lock()
    defer gs.mu.Unlock()
    data, err := json.MarshalIndent(gs, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}

func (gs *GameState) LoadGame(filename string) error {
    gs.mu.Lock()
    defer gs.mu.Unlock()
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, gs)
}

func (gs *GameState) GenerateMonsters() {
    monsterNames := []string{"Goblin", "Orc", "Troll", "Dragon", "Skeleton", "Zombie", "Vampire", "Werewolf"}
    for i := 0; i < 5; i++ {
        name := monsterNames[rand.Intn(len(monsterNames))]
        gs.Monsters = append(gs.Monsters, Monster{
            Name:            name,
            Health:          30 + rand.Intn(50),
            Strength:        5 + rand.Intn(10),
            Agility:         3 + rand.Intn(7),
            ExperienceReward: 20 + rand.Intn(30),
            GoldReward:      10 + rand.Intn(20),
        })
    }
}

func (gs *GameState) Combat() {
    if len(gs.Monsters) == 0 {
        fmt.Println("No monsters to fight!")
        return
    }
    monster := &gs.Monsters[0]
    fmt.Printf("\nA wild %s appears!\n", monster.Name)
    for gs.Player.Health > 0 && monster.Health > 0 {
        fmt.Printf("\nYour Health: %d/%d | %s Health: %d\n", gs.Player.Health, gs.Player.MaxHealth, monster.Name, monster.Health)
        fmt.Println("1. Attack")
        fmt.Println("2. Use Item")
        fmt.Println("3. Flee")
        fmt.Print("Choose an action: ")
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        switch input {
        case "1":
            damage, killed := gs.Player.Attack(monster)
            fmt.Printf("You dealt %d damage to %s!\n", damage, monster.Name)
            if killed {
                fmt.Printf("You defeated %s! Gained %d experience and %d gold.\n", monster.Name, monster.ExperienceReward, monster.GoldReward)
                gs.Monsters = gs.Monsters[1:]
                gs.Player.LevelUp()
                break
            }
            monsterDamage := monster.Attack(&gs.Player)
            fmt.Printf("%s dealt %d damage to you!\n", monster.Name, monsterDamage)
        case "2":
            if len(gs.Player.Inventory) == 0 {
                fmt.Println("No items in inventory!")
                continue
            }
            fmt.Println("Select an item to use:")
            for i, item := range gs.Player.Inventory {
                fmt.Printf("%d. %s\n", i+1, item.Name)
            }
            fmt.Print("Item number: ")
            itemInput, _ := reader.ReadString('\n')
            itemInput = strings.TrimSpace(itemInput)
            idx, err := strconv.Atoi(itemInput)
            if err != nil || idx < 1 || idx > len(gs.Player.Inventory) {
                fmt.Println("Invalid selection!")
                continue
            }
            item := gs.Player.Inventory[idx-1]
            if item.Type == "healing" {
                gs.Player.Heal(item.Value)
                fmt.Printf("Used %s, healed %d health.\n", item.Name, item.Value)
                gs.Player.RemoveItem(idx - 1)
            } else {
                fmt.Println("This item cannot be used in combat!")
            }
            monsterDamage := monster.Attack(&gs.Player)
            fmt.Printf("%s dealt %d damage to you!\n", monster.Name, monsterDamage)
        case "3":
            fmt.Println("You fled from the battle!")
            return
        default:
            fmt.Println("Invalid choice!")
        }
    }
    if gs.Player.Health <= 0 {
        fmt.Println("\nYou have been defeated! Game Over.")
        gs.GameActive = false
    }
}

func (gs *GameState) Shop() {
    items := []Item{
        {Name: "Health Potion", Description: "Restores 50 health", Value: 50, Type: "healing"},
        {Name: "Strength Elixir", Description: "Increases strength by 10", Value: 100, Type: "buff"},
        {Name: "Iron Sword", Description: "A basic sword", Value: 200, Type: "weapon"},
        {Name: "Leather Armor", Description: "Light armor", Value: 150, Type: "armor"},
    }
    fmt.Println("\n=== Shop ===")
    for i, item := range items {
        fmt.Printf("%d. %s - %s (Cost: %d gold)\n", i+1, item.Name, item.Description, item.Value)
    }
    fmt.Println("5. Exit Shop")
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Print("Choose an item to buy (or 5 to exit): ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if input == "5" {
            break
        }
        idx, err := strconv.Atoi(input)
        if err != nil || idx < 1 || idx > 4 {
            fmt.Println("Invalid choice!")
            continue
        }
        item := items[idx-1]
        if gs.Player.Gold >= item.Value {
            gs.Player.Gold -= item.Value
            gs.Player.AddItem(item)
            fmt.Printf("Purchased %s for %d gold.\n", item.Name, item.Value)
        } else {
            fmt.Println("Not enough gold!")
        }
    }
}

func (gs *GameState) Explore() {
    zones := []string{"Forest", "Cave", "Mountain", "Desert", "Swamp"}
    gs.CurrentZone = zones[rand.Intn(len(zones))]
    gs.Day++
    fmt.Printf("\nDay %d: You are exploring the %s.\n", gs.Day, gs.CurrentZone)
    eventChance := rand.Intn(100)
    if eventChance < 30 {
        fmt.Println("You found a treasure chest!")
        goldFound := 10 + rand.Intn(40)
        gs.Player.Gold += goldFound
        fmt.Printf("Found %d gold!\n", goldFound)
    } else if eventChance < 60 {
        fmt.Println("You encounter monsters!")
        gs.GenerateMonsters()
        gs.Combat()
    } else {
        fmt.Println("It's a peaceful day with no events.")
    }
}

func (gs *GameState) DisplayMenu() {
    fmt.Println("\n=== Main Menu ===")
    fmt.Println("1. Explore")
    fmt.Println("2. Shop")
    fmt.Println("3. Status")
    fmt.Println("4. Save Game")
    fmt.Println("5. Load Game")
    fmt.Println("6. Quit")
}

func main() {
    rand.Seed(time.Now().UnixNano())
    game := GameState{
        Player: Character{
            Name:      "Hero",
            Health:    100,
            MaxHealth: 100,
            Strength:  10,
            Agility:   5,
            Level:     1,
            Experience: 0,
            Inventory: []Item{
                {Name: "Small Health Potion", Description: "Restores 20 health", Value: 20, Type: "healing"},
            },
            Gold: 50,
        },
        CurrentZone: "Town",
        Day:         1,
        GameActive:  true,
    }
    game.GenerateMonsters()
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Welcome to Complex Game Simulator!")
    for game.GameActive {
        game.DisplayMenu()
        fmt.Print("Choose an option: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        switch input {
        case "1":
            game.Explore()
        case "2":
            game.Shop()
        case "3":
            game.Player.ShowStatus()
        case "4":
            fmt.Print("Enter filename to save (e.g., save.json): ")
            filename, _ := reader.ReadString('\n')
            filename = strings.TrimSpace(filename)
            if err := game.SaveGame(filename); err != nil {
                fmt.Printf("Error saving game: %v\n", err)
            } else {
                fmt.Println("Game saved successfully!")
            }
        case "5":
            fmt.Print("Enter filename to load (e.g., save.json): ")
            filename, _ := reader.ReadString('\n')
            filename = strings.TrimSpace(filename)
            if err := game.LoadGame(filename); err != nil {
                fmt.Printf("Error loading game: %v\n", err)
            } else {
                fmt.Println("Game loaded successfully!")
            }
        case "6":
            fmt.Println("Thanks for playing! Goodbye!")
            game.GameActive = false
        default:
            fmt.Println("Invalid option! Please choose 1-6.")
        }
    }
}
