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
    Name      string `json:"name"`
    Health    int    `json:"health"`
    MaxHealth int    `json:"max_health"`
    Strength  int    `json:"strength"`
    Defense   int    `json:"defense"`
    Level     int    `json:"level"`
    XP        int    `json:"xp"`
    Gold      int    `json:"gold"`
    Inventory []Item `json:"inventory"`
}

type Item struct {
    Name        string `json:"name"`
    Description string `json:"description"`
    Value       int    `json:"value"`
    Type        string `json:"type"` // "weapon", "armor", "potion", "misc"
    Effect      int    `json:"effect"`
}

type Enemy struct {
    Name     string `json:"name"`
    Health   int    `json:"health"`
    Strength int    `json:"strength"`
    Defense  int    `json:"defense"`
    XP       int    `json:"xp"`
    Gold     int    `json:"gold"`
}

type Location struct {
    Name        string   `json:"name"`
    Description string   `json:"description"`
    Exits       []string `json:"exits"`
    Enemies     []Enemy  `json:"enemies"`
    Items       []Item   `json:"items"`
}

var player Character
var currentLocation Location
var locations map[string]Location
var rng *rand.Rand

func initGame() {
    rng = rand.New(rand.NewSource(time.Now().UnixNano()))
    player = Character{
        Name:      "Hero",
        Health:    100,
        MaxHealth: 100,
        Strength:  10,
        Defense:   5,
        Level:     1,
        XP:        0,
        Gold:      50,
        Inventory: []Item{
            {Name: "Rusty Sword", Description: "A basic sword for beginners.", Value: 10, Type: "weapon", Effect: 5},
            {Name: "Health Potion", Description: "Restores 30 health.", Value: 20, Type: "potion", Effect: 30},
        },
    }
    locations = map[string]Location{
        "forest": {
            Name:        "Mystic Forest",
            Description: "A dense forest with tall trees and mysterious sounds. Paths lead north and east.",
            Exits:       []string{"north", "east"},
            Enemies: []Enemy{
                {Name: "Goblin", Health: 30, Strength: 8, Defense: 3, XP: 20, Gold: 10},
                {Name: "Wolf", Health: 25, Strength: 10, Defense: 2, XP: 15, Gold: 5},
            },
            Items: []Item{
                {Name: "Herbs", Description: "Medicinal herbs that can be used for crafting.", Value: 5, Type: "misc", Effect: 0},
            },
        },
        "cave": {
            Name:        "Dark Cave",
            Description: "A dark and damp cave with glowing mushrooms. Paths lead south and west.",
            Exits:       []string{"south", "west"},
            Enemies: []Enemy{
                {Name: "Troll", Health: 50, Strength: 15, Defense: 8, XP: 40, Gold: 25},
            },
            Items: []Item{
                {Name: "Iron Shield", Description: "A sturdy shield that increases defense.", Value: 30, Type: "armor", Effect: 5},
            },
        },
        "village": {
            Name:        "Peaceful Village",
            Description: "A small village with friendly villagers. Paths lead south and east.",
            Exits:       []string{"south", "east"},
            Enemies:     []Enemy{},
            Items: []Item{
                {Name: "Gold Coin", Description: "A shiny gold coin.", Value: 1, Type: "misc", Effect: 0},
                {Name: "Bread", Description: "Freshly baked bread, restores 10 health.", Value: 5, Type: "potion", Effect: 10},
            },
        },
        "mountain": {
            Name:        "Snowy Mountain",
            Description: "A cold mountain peak with icy winds. Paths lead north and west.",
            Exits:       []string{"north", "west"},
            Enemies: []Enemy{
                {Name: "Yeti", Health: 80, Strength: 20, Defense: 12, XP: 60, Gold: 40},
            },
            Items: []Item{
                {Name: "Diamond", Description: "A rare and valuable gem.", Value: 100, Type: "misc", Effect: 0},
            },
        },
    }
    currentLocation = locations["forest"]
}

func printStatus() {
    fmt.Printf("\n=== Status ===\n")
    fmt.Printf("Name: %s\n", player.Name)
    fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
    fmt.Printf("Level: %d (XP: %d/100)\n", player.Level, player.XP)
    fmt.Printf("Strength: %d, Defense: %d\n", player.Strength, player.Defense)
    fmt.Printf("Gold: %d\n", player.Gold)
    fmt.Printf("Location: %s\n", currentLocation.Name)
    fmt.Printf("Description: %s\n", currentLocation.Description)
    if len(currentLocation.Exits) > 0 {
        fmt.Printf("Exits: %s\n", strings.Join(currentLocation.Exits, ", "))
    }
    if len(currentLocation.Enemies) > 0 {
        fmt.Printf("Enemies here: ")
        for i, enemy := range currentLocation.Enemies {
            if i > 0 {
                fmt.Printf(", ")
            }
            fmt.Printf("%s (Health: %d)", enemy.Name, enemy.Health)
        }
        fmt.Println()
    }
    if len(currentLocation.Items) > 0 {
        fmt.Printf("Items here: ")
        for i, item := range currentLocation.Items {
            if i > 0 {
                fmt.Printf(", ")
            }
            fmt.Printf("%s", item.Name)
        }
        fmt.Println()
    }
}

func handleCommand(cmd string) bool {
    parts := strings.Fields(cmd)
    if len(parts) == 0 {
        return false
    }
    switch parts[0] {
    case "go":
        if len(parts) < 2 {
            fmt.Println("Go where?")
            return false
        }
        direction := parts[1]
        found := false
        for _, exit := range currentLocation.Exits {
            if exit == direction {
                found = true
                break
            }
        }
        if !found {
            fmt.Printf("You cannot go %s from here.\n", direction)
            return false
        }
        var newLocName string
        switch direction {
        case "north":
            if currentLocation.Name == "Mystic Forest" {
                newLocName = "mountain"
            } else if currentLocation.Name == "Dark Cave" {
                newLocName = "forest"
            } else {
                fmt.Println("No location to the north.")
                return false
            }
        case "south":
            if currentLocation.Name == "Mystic Forest" {
                newLocName = "cave"
            } else if currentLocation.Name == "Peaceful Village" {
                newLocName = "forest"
            } else {
                fmt.Println("No location to the south.")
                return false
            }
        case "east":
            if currentLocation.Name == "Mystic Forest" {
                newLocName = "village"
            } else if currentLocation.Name == "Dark Cave" {
                newLocName = "mountain"
            } else {
                fmt.Println("No location to the east.")
                return false
            }
        case "west":
            if currentLocation.Name == "Peaceful Village" {
                newLocName = "forest"
            } else if currentLocation.Name == "Snowy Mountain" {
                newLocName = "cave"
            } else {
                fmt.Println("No location to the west.")
                return false
            }
        default:
            fmt.Printf("Invalid direction: %s\n", direction)
            return false
        }
        currentLocation = locations[newLocName]
        fmt.Printf("You move to %s.\n", currentLocation.Name)
        return false
    case "attack":
        if len(currentLocation.Enemies) == 0 {
            fmt.Println("No enemies to attack here.")
            return false
        }
        enemy := &currentLocation.Enemies[0]
        damage := player.Strength + rng.Intn(6) - enemy.Defense
        if damage < 1 {
            damage = 1
        }
        enemy.Health -= damage
        fmt.Printf("You attack %s for %d damage. %s's health: %d\n", enemy.Name, damage, enemy.Name, enemy.Health)
        if enemy.Health <= 0 {
            fmt.Printf("%s defeated! You gain %d XP and %d gold.\n", enemy.Name, enemy.XP, enemy.Gold)
            player.XP += enemy.XP
            player.Gold += enemy.Gold
            currentLocation.Enemies = currentLocation.Enemies[1:]
            if player.XP >= 100 {
                player.Level++
                player.XP -= 100
                player.MaxHealth += 20
                player.Health = player.MaxHealth
                player.Strength += 3
                player.Defense += 2
                fmt.Printf("Level up! You are now level %d.\n", player.Level)
            }
            return false
        }
        enemyDamage := enemy.Strength + rng.Intn(4) - player.Defense
        if enemyDamage < 1 {
            enemyDamage = 1
        }
        player.Health -= enemyDamage
        fmt.Printf("%s attacks you for %d damage. Your health: %d\n", enemy.Name, enemyDamage, player.Health)
        if player.Health <= 0 {
            fmt.Println("You have been defeated! Game over.")
            return true
        }
        return false
    case "use":
        if len(parts) < 2 {
            fmt.Println("Use what?")
            return false
        }
        itemName := strings.Join(parts[1:], " ")
        for i, item := range player.Inventory {
            if item.Name == itemName {
                if item.Type == "potion" {
                    player.Health += item.Effect
                    if player.Health > player.MaxHealth {
                        player.Health = player.MaxHealth
                    }
                    fmt.Printf("You use %s and restore %d health. Health: %d\n", item.Name, item.Effect, player.Health)
                    player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
                } else {
                    fmt.Printf("You cannot use %s.\n", item.Name)
                }
                return false
            }
        }
        fmt.Printf("You don't have %s in your inventory.\n", itemName)
        return false
    case "take":
        if len(parts) < 2 {
            fmt.Println("Take what?")
            return false
        }
        itemName := strings.Join(parts[1:], " ")
        for i, item := range currentLocation.Items {
            if item.Name == itemName {
                player.Inventory = append(player.Inventory, item)
                fmt.Printf("You take %s.\n", item.Name)
                currentLocation.Items = append(currentLocation.Items[:i], currentLocation.Items[i+1:]...)
                return false
            }
        }
        fmt.Printf("%s is not here.\n", itemName)
        return false
    case "inventory":
        if len(player.Inventory) == 0 {
            fmt.Println("Your inventory is empty.")
        } else {
            fmt.Println("Inventory:")
            for _, item := range player.Inventory {
                fmt.Printf("  - %s: %s (Value: %d)\n", item.Name, item.Description, item.Value)
            }
        }
        return false
    case "save":
        data, err := json.Marshal(player)
        if err != nil {
            fmt.Printf("Error saving game: %v\n", err)
            return false
        }
        err = os.WriteFile("save_game.json", data, 0644)
        if err != nil {
            fmt.Printf("Error writing save file: %v\n", err)
            return false
        }
        fmt.Println("Game saved to save_game.json")
        return false
    case "load":
        data, err := os.ReadFile("save_game.json")
        if err != nil {
            fmt.Printf("Error loading game: %v\n", err)
            return false
        }
        err = json.Unmarshal(data, &player)
        if err != nil {
            fmt.Printf("Error parsing save file: %v\n", err)
            return false
        }
        fmt.Println("Game loaded from save_game.json")
        return false
    case "quit":
        fmt.Println("Thanks for playing!")
        return true
    case "help":
        fmt.Println("Available commands:")
        fmt.Println("  go <direction> - Move north, south, east, or west")
        fmt.Println("  attack - Attack the first enemy in the location")
        fmt.Println("  use <item> - Use an item from your inventory (e.g., potion)")
        fmt.Println("  take <item> - Take an item from the current location")
        fmt.Println("  inventory - Show your inventory")
        fmt.Println("  save - Save your game progress")
        fmt.Println("  load - Load a saved game")
        fmt.Println("  status - Show your status and location info")
        fmt.Println("  help - Show this help message")
        fmt.Println("  quit - Quit the game")
        return false
    case "status":
        printStatus()
        return false
    default:
        fmt.Printf("Unknown command: %s. Type 'help' for a list of commands.\n", parts[0])
        return false
    }
}

func main() {
    initGame()
    fmt.Println("Welcome to the Adventure Game Simulator!")
    fmt.Println("Type 'help' for a list of commands.")
    scanner := bufio.NewScanner(os.Stdin)
    for {
        printStatus()
        fmt.Print("> ")
        if !scanner.Scan() {
            break
        }
        cmd := scanner.Text()
        if cmd == "" {
            continue
        }
        if handleCommand(cmd) {
            break
        }
    }
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading input: %v\n", err)
    }
}
