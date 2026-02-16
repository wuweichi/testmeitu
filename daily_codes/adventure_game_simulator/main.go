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
    Attack    int
    Defense   int
    Level     int
    Experience int
    Gold      int
    Inventory []string
    Skills    map[string]int
}

type Enemy struct {
    Name    string
    Health  int
    Attack  int
    Defense int
    GoldDrop int
    ExpDrop int
}

type Item struct {
    Name   string
    Effect string
    Value  int
}

type Location struct {
    Name        string
    Description string
    Enemies     []Enemy
    Items       []Item
    Connections []string
}

var player Character
var locations map[string]Location
var currentLocation string
var gameRunning bool

func initGame() {
    rand.Seed(time.Now().UnixNano())
    player = Character{
        Name:      "Hero",
        Health:    100,
        MaxHealth: 100,
        Attack:    10,
        Defense:   5,
        Level:     1,
        Experience: 0,
        Gold:      50,
        Inventory: []string{"Potion", "Sword"},
        Skills:    map[string]int{"Slash": 15, "Heal": 20},
    }
    locations = make(map[string]Location)
    locations["forest"] = Location{
        Name:        "Dark Forest",
        Description: "A dense forest with tall trees and mysterious sounds.",
        Enemies: []Enemy{
            {Name: "Goblin", Health: 30, Attack: 8, Defense: 2, GoldDrop: 10, ExpDrop: 20},
            {Name: "Wolf", Health: 25, Attack: 10, Defense: 1, GoldDrop: 5, ExpDrop: 15},
        },
        Items: []Item{
            {Name: "Herb", Effect: "Restores 20 health", Value: 5},
            {Name: "Iron Shield", Effect: "Increases defense by 3", Value: 30},
        },
        Connections: []string{"town", "cave"},
    }
    locations["town"] = Location{
        Name:        "Peaceful Town",
        Description: "A small town with shops and friendly villagers.",
        Enemies:     []Enemy{},
        Items: []Item{
            {Name: "Health Potion", Effect: "Restores 50 health", Value: 20},
            {Name: "Steel Sword", Effect: "Increases attack by 5", Value: 50},
        },
        Connections: []string{"forest", "mountain"},
    }
    locations["cave"] = Location{
        Name:        "Dark Cave",
        Description: "A dark and damp cave with glowing mushrooms.",
        Enemies: []Enemy{
            {Name: "Troll", Health: 80, Attack: 15, Defense: 10, GoldDrop: 30, ExpDrop: 50},
            {Name: "Bat Swarm", Health: 20, Attack: 5, Defense: 0, GoldDrop: 2, ExpDrop: 10},
        },
        Items: []Item{
            {Name: "Mystic Gem", Effect: "Unknown power", Value: 100},
            {Name: "Cave Mushroom", Effect: "Restores 10 health", Value: 3},
        },
        Connections: []string{"forest"},
    }
    locations["mountain"] = Location{
        Name:        "Snowy Mountain",
        Description: "A cold mountain peak with icy winds.",
        Enemies: []Enemy{
            {Name: "Yeti", Health: 120, Attack: 20, Defense: 15, GoldDrop: 50, ExpDrop: 80},
            {Name: "Ice Golem", Health: 100, Attack: 18, Defense: 12, GoldDrop: 40, ExpDrop: 60},
        },
        Items: []Item{
            {Name: "Ice Crystal", Effect: "Freezes enemies", Value: 80},
            {Name: "Warm Cloak", Effect: "Protects from cold", Value: 40},
        },
        Connections: []string{"town"},
    }
    currentLocation = "town"
    gameRunning = true
}

func showStatus() {
    fmt.Printf("\n=== Status ===\n")
    fmt.Printf("Name: %s\n", player.Name)
    fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
    fmt.Printf("Attack: %d, Defense: %d\n", player.Attack, player.Defense)
    fmt.Printf("Level: %d, Experience: %d/%d\n", player.Level, player.Experience, player.Level*100)
    fmt.Printf("Gold: %d\n", player.Gold)
    fmt.Printf("Inventory: %v\n", player.Inventory)
    fmt.Printf("Skills: %v\n", player.Skills)
}

func showLocation() {
    loc := locations[currentLocation]
    fmt.Printf("\n=== %s ===\n", loc.Name)
    fmt.Printf("Description: %s\n", loc.Description)
    if len(loc.Enemies) > 0 {
        fmt.Printf("Enemies here: ")
        for i, e := range loc.Enemies {
            fmt.Printf("%s (Health: %d)", e.Name, e.Health)
            if i < len(loc.Enemies)-1 {
                fmt.Printf(", ")
            }
        }
        fmt.Println()
    } else {
        fmt.Println("No enemies here.")
    }
    if len(loc.Items) > 0 {
        fmt.Printf("Items here: ")
        for i, item := range loc.Items {
            fmt.Printf("%s (%s)", item.Name, item.Effect)
            if i < len(loc.Items)-1 {
                fmt.Printf(", ")
            }
        }
        fmt.Println()
    } else {
        fmt.Println("No items here.")
    }
    fmt.Printf("Connections: %v\n", loc.Connections)
}

func move(destination string) {
    loc := locations[currentLocation]
    for _, conn := range loc.Connections {
        if conn == destination {
            currentLocation = destination
            fmt.Printf("Moved to %s.\n", destination)
            return
        }
    }
    fmt.Printf("Cannot move to %s from here.\n", destination)
}

func fight() {
    loc := locations[currentLocation]
    if len(loc.Enemies) == 0 {
        fmt.Println("No enemies to fight here.")
        return
    }
    enemy := loc.Enemies[rand.Intn(len(loc.Enemies))]
    fmt.Printf("You encounter a %s!\n", enemy.Name)
    for player.Health > 0 && enemy.Health > 0 {
        fmt.Printf("\nYour Health: %d, %s's Health: %d\n", player.Health, enemy.Name, enemy.Health)
        fmt.Println("Choose action: 1. Attack, 2. Use Skill, 3. Use Item, 4. Flee")
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        switch input {
        case "1":
            damage := player.Attack - enemy.Defense
            if damage < 1 {
                damage = 1
            }
            enemy.Health -= damage
            fmt.Printf("You attack the %s for %d damage.\n", enemy.Name, damage)
        case "2":
            fmt.Printf("Skills: %v\n", player.Skills)
            fmt.Print("Enter skill name: ")
            skill, _ := reader.ReadString('\n')
            skill = strings.TrimSpace(skill)
            if dmg, ok := player.Skills[skill]; ok {
                damage := dmg - enemy.Defense
                if damage < 1 {
                    damage = 1
                }
                enemy.Health -= damage
                fmt.Printf("You use %s for %d damage.\n", skill, damage)
            } else {
                fmt.Println("Invalid skill.")
                continue
            }
        case "3":
            fmt.Printf("Inventory: %v\n", player.Inventory)
            fmt.Print("Enter item name: ")
            itemName, _ := reader.ReadString('\n')
            itemName = strings.TrimSpace(itemName)
            found := false
            for i, item := range player.Inventory {
                if item == itemName {
                    found = true
                    player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
                    if itemName == "Potion" || itemName == "Health Potion" || itemName == "Herb" || itemName == "Cave Mushroom" {
                        heal := 0
                        switch itemName {
                        case "Potion":
                            heal = 30
                        case "Health Potion":
                            heal = 50
                        case "Herb":
                            heal = 20
                        case "Cave Mushroom":
                            heal = 10
                        }
                        player.Health += heal
                        if player.Health > player.MaxHealth {
                            player.Health = player.MaxHealth
                        }
                        fmt.Printf("You use %s and heal %d health.\n", itemName, heal)
                    }
                    break
                }
            }
            if !found {
                fmt.Println("Item not found in inventory.")
                continue
            }
        case "4":
            if rand.Intn(100) < 50 {
                fmt.Println("You successfully flee from the battle!")
                return
            } else {
                fmt.Println("You failed to flee!")
            }
        default:
            fmt.Println("Invalid choice.")
            continue
        }
        if enemy.Health > 0 {
            damage := enemy.Attack - player.Defense
            if damage < 1 {
                damage = 1
            }
            player.Health -= damage
            fmt.Printf("The %s attacks you for %d damage.\n", enemy.Name, damage)
        }
    }
    if player.Health <= 0 {
        fmt.Println("You have been defeated! Game Over.")
        gameRunning = false
        return
    }
    if enemy.Health <= 0 {
        fmt.Printf("You defeated the %s!\n", enemy.Name)
        player.Gold += enemy.GoldDrop
        player.Experience += enemy.ExpDrop
        fmt.Printf("Gained %d gold and %d experience.\n", enemy.GoldDrop, enemy.ExpDrop)
        if player.Experience >= player.Level*100 {
            player.Level++
            player.MaxHealth += 20
            player.Health = player.MaxHealth
            player.Attack += 3
            player.Defense += 2
            fmt.Printf("Level up! Now level %d. Stats increased.\n", player.Level)
        }
    }
}

func explore() {
    loc := locations[currentLocation]
    if len(loc.Items) == 0 {
        fmt.Println("Nothing to find here.")
        return
    }
    item := loc.Items[rand.Intn(len(loc.Items))]
    fmt.Printf("You found a %s! (%s)\n", item.Name, item.Effect)
    player.Inventory = append(player.Inventory, item.Name)
    loc.Items = removeItem(loc.Items, item)
    locations[currentLocation] = loc
}

func removeItem(items []Item, target Item) []Item {
    for i, item := range items {
        if item.Name == target.Name && item.Effect == target.Effect && item.Value == target.Value {
            return append(items[:i], items[i+1:]...)
        }
    }
    return items
}

func shop() {
    if currentLocation != "town" {
        fmt.Println("You can only shop in the town.")
        return
    }
    fmt.Println("Welcome to the shop! Items for sale:")
    loc := locations[currentLocation]
    for i, item := range loc.Items {
        fmt.Printf("%d. %s - %s (Cost: %d gold)\n", i+1, item.Name, item.Effect, item.Value)
    }
    fmt.Println("Enter item number to buy, or 'exit' to leave.")
    reader := bufio.NewReader(os.Stdin)
    for {
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if input == "exit" {
            fmt.Println("Leaving shop.")
            return
        }
        idx, err := strconv.Atoi(input)
        if err != nil || idx < 1 || idx > len(loc.Items) {
            fmt.Println("Invalid choice.")
            continue
        }
        item := loc.Items[idx-1]
        if player.Gold >= item.Value {
            player.Gold -= item.Value
            player.Inventory = append(player.Inventory, item.Name)
            fmt.Printf("You bought a %s for %d gold.\n", item.Name, item.Value)
            loc.Items = append(loc.Items[:idx-1], loc.Items[idx:]...)
            locations[currentLocation] = loc
        } else {
            fmt.Println("Not enough gold.")
        }
    }
}

func saveGame() {
    data, err := json.MarshalIndent(struct {
        Player          Character
        Locations       map[string]Location
        CurrentLocation string
    }{player, locations, currentLocation}, "", "  ")
    if err != nil {
        fmt.Printf("Error saving game: %v\n", err)
        return
    }
    err = os.WriteFile("savegame.json", data, 0644)
    if err != nil {
        fmt.Printf("Error writing file: %v\n", err)
    } else {
        fmt.Println("Game saved to savegame.json.")
    }
}

func loadGame() {
    data, err := os.ReadFile("savegame.json")
    if err != nil {
        fmt.Printf("No save file found: %v\n", err)
        return
    }
    var save struct {
        Player          Character
        Locations       map[string]Location
        CurrentLocation string
    }
    err = json.Unmarshal(data, &save)
    if err != nil {
        fmt.Printf("Error loading game: %v\n", err)
        return
    }
    player = save.Player
    locations = save.Locations
    currentLocation = save.CurrentLocation
    fmt.Println("Game loaded from savegame.json.")
}

func main() {
    initGame()
    fmt.Println("Welcome to Adventure Game Simulator!")
    fmt.Println("Type 'help' for commands.")
    reader := bufio.NewReader(os.Stdin)
    for gameRunning {
        fmt.Print("\n> ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        switch input {
        case "help":
            fmt.Println("Commands: status, look, move <location>, fight, explore, shop, save, load, quit")
        case "status":
            showStatus()
        case "look":
            showLocation()
        case "fight":
            fight()
        case "explore":
            explore()
        case "shop":
            shop()
        case "save":
            saveGame()
        case "load":
            loadGame()
        case "quit":
            fmt.Println("Thanks for playing!")
            gameRunning = false
        default:
            if strings.HasPrefix(input, "move ") {
                parts := strings.Split(input, " ")
                if len(parts) == 2 {
                    move(parts[1])
                } else {
                    fmt.Println("Usage: move <location>")
                }
            } else {
                fmt.Println("Unknown command. Type 'help' for commands.")
            }
        }
    }
}
