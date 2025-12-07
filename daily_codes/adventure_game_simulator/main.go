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

type Character struct {
    Name      string
    Health    int
    MaxHealth int
    Strength  int
    Agility   int
    Level     int
    XP        int
    Gold      int
    Inventory []string
}

type Enemy struct {
    Name     string
    Health   int
    Strength int
    Agility  int
    XP       int
    Gold     int
}

type Location struct {
    Name        string
    Description string
    Enemies     []Enemy
    Items       []string
    Connections []string
}

type GameState struct {
    Player    Character
    Location  string
    Locations map[string]Location
    GameLog   []string
}

func (c *Character) Attack(e *Enemy) int {
    damage := c.Strength + rand.Intn(5)
    e.Health -= damage
    return damage
}

func (e *Enemy) Attack(c *Character) int {
    damage := e.Strength + rand.Intn(3)
    c.Health -= damage
    return damage
}

func (c *Character) Heal(amount int) {
    c.Health += amount
    if c.Health > c.MaxHealth {
        c.Health = c.MaxHealth
    }
}

func (c *Character) GainXP(xp int) {
    c.XP += xp
    for c.XP >= c.Level*100 {
        c.XP -= c.Level * 100
        c.Level++
        c.MaxHealth += 10
        c.Health = c.MaxHealth
        c.Strength += 2
        c.Agility += 1
    }
}

func (g *GameState) AddLog(message string) {
    g.GameLog = append(g.GameLog, message)
    if len(g.GameLog) > 20 {
        g.GameLog = g.GameLog[1:]
    }
}

func (g *GameState) DisplayLog() {
    fmt.Println("\n=== Game Log ===")
    for _, entry := range g.GameLog {
        fmt.Println(entry)
    }
    fmt.Println("================")
}

func (g *GameState) SaveGame(filename string) error {
    data, err := json.MarshalIndent(g, "", "  ")
    if err != nil {
        return err
    }
    return os.WriteFile(filename, data, 0644)
}

func (g *GameState) LoadGame(filename string) error {
    data, err := os.ReadFile(filename)
    if err != nil {
        return err
    }
    return json.Unmarshal(data, g)
}

func (g *GameState) MoveTo(locationName string) bool {
    if loc, exists := g.Locations[locationName]; exists {
        for _, conn := range loc.Connections {
            if conn == g.Location {
                g.Location = locationName
                g.AddLog(fmt.Sprintf("Moved to %s", locationName))
                return true
            }
        }
        g.AddLog("Cannot move to that location from here.")
        return false
    }
    g.AddLog("Location not found.")
    return false
}

func (g *GameState) Explore() {
    loc := g.Locations[g.Location]
    fmt.Printf("\nYou are in %s. %s\n", loc.Name, loc.Description)
    if len(loc.Enemies) > 0 {
        fmt.Println("Enemies here:")
        for i, enemy := range loc.Enemies {
            fmt.Printf("  %d. %s (Health: %d)\n", i+1, enemy.Name, enemy.Health)
        }
    }
    if len(loc.Items) > 0 {
        fmt.Println("Items here:")
        for _, item := range loc.Items {
            fmt.Printf("  - %s\n", item)
        }
    }
    fmt.Println("Connections:")
    for _, conn := range loc.Connections {
        fmt.Printf("  - %s\n", conn)
    }
}

func (g *GameState) Fight(enemyIndex int) {
    loc := g.Locations[g.Location]
    if enemyIndex < 0 || enemyIndex >= len(loc.Enemies) {
        g.AddLog("Invalid enemy selection.")
        return
    }
    enemy := &loc.Enemies[enemyIndex]
    g.AddLog(fmt.Sprintf("You engage in combat with %s!", enemy.Name))
    for g.Player.Health > 0 && enemy.Health > 0 {
        playerDamage := g.Player.Attack(enemy)
        g.AddLog(fmt.Sprintf("You hit %s for %d damage.", enemy.Name, playerDamage))
        if enemy.Health <= 0 {
            g.AddLog(fmt.Sprintf("You defeated %s!", enemy.Name))
            g.Player.GainXP(enemy.XP)
            g.Player.Gold += enemy.Gold
            g.AddLog(fmt.Sprintf("Gained %d XP and %d gold.", enemy.XP, enemy.Gold))
            loc.Enemies = append(loc.Enemies[:enemyIndex], loc.Enemies[enemyIndex+1:]...)
            g.Locations[g.Location] = loc
            return
        }
        enemyDamage := enemy.Attack(&g.Player)
        g.AddLog(fmt.Sprintf("%s hits you for %d damage.", enemy.Name, enemyDamage))
        if g.Player.Health <= 0 {
            g.AddLog("You have been defeated! Game over.")
            return
        }
    }
}

func (g *GameState) PickupItem(itemName string) {
    loc := g.Locations[g.Location]
    for i, item := range loc.Items {
        if strings.EqualFold(item, itemName) {
            g.Player.Inventory = append(g.Player.Inventory, item)
            g.AddLog(fmt.Sprintf("Picked up %s.", item))
            loc.Items = append(loc.Items[:i], loc.Items[i+1:]...)
            g.Locations[g.Location] = loc
            return
        }
    }
    g.AddLog("Item not found here.")
}

func (g *GameState) UseItem(itemName string) {
    for i, item := range g.Player.Inventory {
        if strings.EqualFold(item, itemName) {
            switch item {
            case "Health Potion":
                g.Player.Heal(20)
                g.AddLog("Used Health Potion. Healed 20 health.")
            case "Strength Elixir":
                g.Player.Strength += 5
                g.AddLog("Used Strength Elixir. Strength increased by 5.")
            default:
                g.AddLog(fmt.Sprintf("Used %s. (No effect)", item))
            }
            g.Player.Inventory = append(g.Player.Inventory[:i], g.Player.Inventory[i+1:]...)
            return
        }
    }
    g.AddLog("Item not in inventory.")
}

func (g *GameState) ShowStatus() {
    fmt.Printf("\n=== Player Status ===\n")
    fmt.Printf("Name: %s\n", g.Player.Name)
    fmt.Printf("Health: %d/%d\n", g.Player.Health, g.Player.MaxHealth)
    fmt.Printf("Level: %d (XP: %d/%d)\n", g.Player.Level, g.Player.XP, g.Player.Level*100)
    fmt.Printf("Strength: %d, Agility: %d\n", g.Player.Strength, g.Player.Agility)
    fmt.Printf("Gold: %d\n", g.Player.Gold)
    fmt.Printf("Inventory: %v\n", g.Player.Inventory)
    fmt.Printf("Current Location: %s\n", g.Location)
}

func (g *GameState) ShowHelp() {
    fmt.Println("\n=== Commands ===")
    fmt.Println("explore - Explore the current location")
    fmt.Println("move <location> - Move to a connected location")
    fmt.Println("fight <enemy number> - Fight an enemy")
    fmt.Println("pickup <item> - Pick up an item")
    fmt.Println("use <item> - Use an item from inventory")
    fmt.Println("status - Show player status")
    fmt.Println("log - Show game log")
    fmt.Println("save <filename> - Save the game")
    fmt.Println("load <filename> - Load a saved game")
    fmt.Println("quit - Quit the game")
    fmt.Println("help - Show this help message")
}

func initializeGame() GameState {
    rand.Seed(time.Now().UnixNano())
    player := Character{
        Name:      "Hero",
        Health:    100,
        MaxHealth: 100,
        Strength:  10,
        Agility:   5,
        Level:     1,
        XP:        0,
        Gold:      50,
        Inventory: []string{"Health Potion", "Health Potion"},
    }
    locations := map[string]Location{
        "forest": {
            Name:        "Dark Forest",
            Description: "A dense forest with tall trees and eerie sounds.",
            Enemies: []Enemy{
                {Name: "Goblin", Health: 30, Strength: 5, Agility: 3, XP: 20, Gold: 10},
                {Name: "Wolf", Health: 25, Strength: 7, Agility: 5, XP: 15, Gold: 5},
            },
            Items:       []string{"Health Potion", "Rusty Sword"},
            Connections: []string{"village", "cave"},
        },
        "village": {
            Name:        "Quiet Village",
            Description: "A peaceful village with friendly inhabitants.",
            Enemies:     []Enemy{},
            Items:       []string{"Gold Coin", "Bread"},
            Connections: []string{"forest"},
        },
        "cave": {
            Name:        "Dark Cave",
            Description: "A dark and damp cave with glowing mushrooms.",
            Enemies: []Enemy{
                {Name: "Troll", Health: 50, Strength: 15, Agility: 2, XP: 50, Gold: 30},
            },
            Items:       []string{"Strength Elixir", "Treasure Chest"},
            Connections: []string{"forest"},
        },
    }
    return GameState{
        Player:    player,
        Location:  "forest",
        Locations: locations,
        GameLog:   []string{"Welcome to the Adventure Game Simulator!"},
    }
}

func main() {
    game := initializeGame()
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Println("=== Adventure Game Simulator ===")
    fmt.Println("Type 'help' for a list of commands.")
    game.ShowStatus()
    for {
        fmt.Print("\n> ")
        if !scanner.Scan() {
            break
        }
        input := scanner.Text()
        parts := strings.Fields(input)
        if len(parts) == 0 {
            continue
        }
        command := strings.ToLower(parts[0])
        switch command {
        case "explore":
            game.Explore()
        case "move":
            if len(parts) < 2 {
                game.AddLog("Please specify a location to move to.")
            } else {
                game.MoveTo(parts[1])
            }
        case "fight":
            if len(parts) < 2 {
                game.AddLog("Please specify an enemy number to fight.")
            } else {
                if idx, err := strconv.Atoi(parts[1]); err == nil {
                    game.Fight(idx - 1)
                } else {
                    game.AddLog("Invalid enemy number.")
                }
            }
        case "pickup":
            if len(parts) < 2 {
                game.AddLog("Please specify an item to pick up.")
            } else {
                game.PickupItem(parts[1])
            }
        case "use":
            if len(parts) < 2 {
                game.AddLog("Please specify an item to use.")
            } else {
                game.UseItem(parts[1])
            }
        case "status":
            game.ShowStatus()
        case "log":
            game.DisplayLog()
        case "save":
            if len(parts) < 2 {
                game.AddLog("Please specify a filename to save.")
            } else {
                if err := game.SaveGame(parts[1]); err != nil {
                    game.AddLog(fmt.Sprintf("Error saving game: %v", err))
                } else {
                    game.AddLog("Game saved successfully.")
                }
            }
        case "load":
            if len(parts) < 2 {
                game.AddLog("Please specify a filename to load.")
            } else {
                if err := game.LoadGame(parts[1]); err != nil {
                    game.AddLog(fmt.Sprintf("Error loading game: %v", err))
                } else {
                    game.AddLog("Game loaded successfully.")
                }
            }
        case "quit":
            fmt.Println("Thanks for playing!")
            return
        case "help":
            game.ShowHelp()
        default:
            game.AddLog("Unknown command. Type 'help' for a list.")
        }
        if game.Player.Health <= 0 {
            fmt.Println("\nGame Over! You have been defeated.")
            break
        }
    }
}
