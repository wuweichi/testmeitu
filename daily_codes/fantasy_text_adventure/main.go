package main

import (
    "bufio"
    "fmt"
    "math/rand"
    "os"
    "strconv"
    "strings"
    "time"
)

type Player struct {
    Name      string
    Health    int
    MaxHealth int
    Strength  int
    Agility   int
    Wisdom    int
    Gold      int
    Inventory []string
    Location  string
}

type Enemy struct {
    Name     string
    Health   int
    Strength int
    Agility  int
    Wisdom   int
    GoldDrop int
}

type Location struct {
    Name        string
    Description string
    Exits       map[string]string
    Enemies     []Enemy
    Items       []string
}

var player Player
var locations map[string]Location
var scanner *bufio.Scanner

func initGame() {
    rand.Seed(time.Now().UnixNano())
    scanner = bufio.NewScanner(os.Stdin)
    player = Player{
        Name:      "Hero",
        Health:    100,
        MaxHealth: 100,
        Strength:  10,
        Agility:   10,
        Wisdom:    10,
        Gold:      50,
        Inventory: []string{"rusty sword", "leather armor"},
        Location:  "town_square",
    }
    locations = make(map[string]Location)
    locations["town_square"] = Location{
        Name:        "Town Square",
        Description: "You are in the bustling town square. People are going about their daily business. To the north is the forest, to the east is the shop, and to the west is the tavern.",
        Exits: map[string]string{
            "north": "forest_entrance",
            "east":  "shop",
            "west":  "tavern",
        },
        Enemies: []Enemy{},
        Items:   []string{},
    }
    locations["forest_entrance"] = Location{
        Name:        "Forest Entrance",
        Description: "You stand at the edge of a dark forest. The trees loom overhead. You can hear strange noises. To the south is the town square, and deeper into the forest to the north.",
        Exits: map[string]string{
            "south": "town_square",
            "north": "deep_forest",
        },
        Enemies: []Enemy{
            {Name: "Goblin", Health: 30, Strength: 5, Agility: 8, Wisdom: 2, GoldDrop: 10},
        },
        Items: []string{"herbs"},
    }
    locations["deep_forest"] = Location{
        Name:        "Deep Forest",
        Description: "You are deep in the forest. It's dark and eerie. You see glowing eyes in the shadows. To the south is the forest entrance.",
        Exits: map[string]string{
            "south": "forest_entrance",
        },
        Enemies: []Enemy{
            {Name: "Wolf", Health: 50, Strength: 8, Agility: 12, Wisdom: 4, GoldDrop: 20},
            {Name: "Troll", Health: 100, Strength: 15, Agility: 3, Wisdom: 1, GoldDrop: 50},
        },
        Items: []string{"magic amulet"},
    }
    locations["shop"] = Location{
        Name:        "Shop",
        Description: "You are in a small shop. Shelves are lined with various goods. The shopkeeper eyes you curiously. To the west is the town square.",
        Exits: map[string]string{
            "west": "town_square",
        },
        Enemies: []Enemy{},
        Items:   []string{"health potion", "strength potion"},
    }
    locations["tavern"] = Location{
        Name:        "Tavern",
        Description: "The tavern is noisy and filled with adventurers. You smell ale and roasted meat. To the east is the town square.",
        Exits: map[string]string{
            "east": "town_square",
        },
        Enemies: []Enemy{},
        Items:   []string{"ale"},
    }
}

func printStatus() {
    fmt.Printf("\n--- Status ---\n")
    fmt.Printf("Name: %s\n", player.Name)
    fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
    fmt.Printf("Strength: %d, Agility: %d, Wisdom: %d\n", player.Strength, player.Agility, player.Wisdom)
    fmt.Printf("Gold: %d\n", player.Gold)
    fmt.Printf("Inventory: %v\n", player.Inventory)
    fmt.Printf("Location: %s\n", player.Location)
    fmt.Printf("----------------\n\n")
}

func move(direction string) {
    loc := locations[player.Location]
    if dest, ok := loc.Exits[direction]; ok {
        player.Location = dest
        fmt.Printf("You move %s to %s.\n", direction, locations[dest].Name)
        describeLocation()
    } else {
        fmt.Printf("You cannot go %s from here.\n", direction)
    }
}

func describeLocation() {
    loc := locations[player.Location]
    fmt.Printf("\n--- %s ---\n", loc.Name)
    fmt.Println(loc.Description)
    if len(loc.Enemies) > 0 {
        fmt.Println("Enemies here:")
        for _, enemy := range loc.Enemies {
            fmt.Printf("  - %s (Health: %d)\n", enemy.Name, enemy.Health)
        }
    }
    if len(loc.Items) > 0 {
        fmt.Println("Items here:")
        for _, item := range loc.Items {
            fmt.Printf("  - %s\n", item)
        }
    }
    fmt.Println("Exits:")
    for dir, dest := range loc.Exits {
        fmt.Printf("  - %s to %s\n", dir, locations[dest].Name)
    }
}

func fight() {
    loc := locations[player.Location]
    if len(loc.Enemies) == 0 {
        fmt.Println("There are no enemies here to fight.")
        return
    }
    fmt.Println("Choose an enemy to fight:")
    for i, enemy := range loc.Enemies {
        fmt.Printf("%d. %s (Health: %d)\n", i+1, enemy.Name, enemy.Health)
    }
    scanner.Scan()
    choice, err := strconv.Atoi(scanner.Text())
    if err != nil || choice < 1 || choice > len(loc.Enemies) {
        fmt.Println("Invalid choice.")
        return
    }
    enemy := loc.Enemies[choice-1]
    fmt.Printf("You engage in combat with %s!\n", enemy.Name)
    for player.Health > 0 && enemy.Health > 0 {
        playerAttack := rand.Intn(player.Strength) + 1
        enemyAttack := rand.Intn(enemy.Strength) + 1
        enemy.Health -= playerAttack
        fmt.Printf("You hit %s for %d damage. %s's health: %d\n", enemy.Name, playerAttack, enemy.Name, enemy.Health)
        if enemy.Health <= 0 {
            fmt.Printf("You defeated %s! You gain %d gold.\n", enemy.Name, enemy.GoldDrop)
            player.Gold += enemy.GoldDrop
            loc.Enemies = append(loc.Enemies[:choice-1], loc.Enemies[choice:]...)
            locations[player.Location] = loc
            break
        }
        player.Health -= enemyAttack
        fmt.Printf("%s hits you for %d damage. Your health: %d\n", enemy.Name, enemyAttack, player.Health)
        if player.Health <= 0 {
            fmt.Println("You have been defeated! Game over.")
            os.Exit(0)
        }
    }
}

func takeItem() {
    loc := locations[player.Location]
    if len(loc.Items) == 0 {
        fmt.Println("There are no items here to take.")
        return
    }
    fmt.Println("Choose an item to take:")
    for i, item := range loc.Items {
        fmt.Printf("%d. %s\n", i+1, item)
    }
    scanner.Scan()
    choice, err := strconv.Atoi(scanner.Text())
    if err != nil || choice < 1 || choice > len(loc.Items) {
        fmt.Println("Invalid choice.")
        return
    }
    item := loc.Items[choice-1]
    player.Inventory = append(player.Inventory, item)
    loc.Items = append(loc.Items[:choice-1], loc.Items[choice:]...)
    locations[player.Location] = loc
    fmt.Printf("You took the %s.\n", item)
}

func useItem() {
    if len(player.Inventory) == 0 {
        fmt.Println("Your inventory is empty.")
        return
    }
    fmt.Println("Choose an item to use:")
    for i, item := range player.Inventory {
        fmt.Printf("%d. %s\n", i+1, item)
    }
    scanner.Scan()
    choice, err := strconv.Atoi(scanner.Text())
    if err != nil || choice < 1 || choice > len(player.Inventory) {
        fmt.Println("Invalid choice.")
        return
    }
    item := player.Inventory[choice-1]
    switch item {
    case "health potion":
        player.Health = player.MaxHealth
        fmt.Println("You drank a health potion and restored your health to full.")
    case "strength potion":
        player.Strength += 5
        fmt.Println("You drank a strength potion and gained 5 strength.")
    case "ale":
        fmt.Println("You drink the ale. It tastes good but has no effect.")
    default:
        fmt.Printf("You cannot use the %s right now.\n", item)
        return
    }
    player.Inventory = append(player.Inventory[:choice-1], player.Inventory[choice:]...)
}

func shopAction() {
    if player.Location != "shop" {
        fmt.Println("You are not in the shop.")
        return
    }
    fmt.Println("Welcome to the shop! Items for sale:")
    fmt.Println("1. Health Potion - 20 gold")
    fmt.Println("2. Strength Potion - 30 gold")
    fmt.Println("3. Leave shop")
    scanner.Scan()
    choice, err := strconv.Atoi(scanner.Text())
    if err != nil {
        fmt.Println("Invalid input.")
        return
    }
    switch choice {
    case 1:
        if player.Gold >= 20 {
            player.Gold -= 20
            player.Inventory = append(player.Inventory, "health potion")
            fmt.Println("You bought a health potion.")
        } else {
            fmt.Println("Not enough gold.")
        }
    case 2:
        if player.Gold >= 30 {
            player.Gold -= 30
            player.Inventory = append(player.Inventory, "strength potion")
            fmt.Println("You bought a strength potion.")
        } else {
            fmt.Println("Not enough gold.")
        }
    case 3:
        fmt.Println("You leave the shop.")
    default:
        fmt.Println("Invalid choice.")
    }
}

func help() {
    fmt.Println("\n--- Commands ---")
    fmt.Println("status - Show player status")
    fmt.Println("look - Describe current location")
    fmt.Println("go [direction] - Move in a direction (e.g., go north)")
    fmt.Println("fight - Engage in combat with enemies")
    fmt.Println("take - Take an item from the location")
    fmt.Println("use - Use an item from inventory")
    fmt.Println("shop - Buy items in the shop")
    fmt.Println("help - Show this help message")
    fmt.Println("quit - Exit the game")
    fmt.Println("----------------\n")
}

func main() {
    initGame()
    fmt.Println("Welcome to the Fantasy Text Adventure!")
    fmt.Println("Type 'help' for a list of commands.\n")
    describeLocation()
    for {
        fmt.Print("> ")
        scanner.Scan()
        input := strings.ToLower(strings.TrimSpace(scanner.Text()))
        if input == "quit" {
            fmt.Println("Thanks for playing!")
            break
        }
        switch input {
        case "status":
            printStatus()
        case "look":
            describeLocation()
        case "fight":
            fight()
        case "take":
            takeItem()
        case "use":
            useItem()
        case "shop":
            shopAction()
        case "help":
            help()
        default:
            if strings.HasPrefix(input, "go ") {
                direction := strings.TrimPrefix(input, "go ")
                move(direction)
            } else {
                fmt.Println("Unknown command. Type 'help' for a list.")
            }
        }
    }
}
