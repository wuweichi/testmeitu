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

type Character struct {
    Name      string
    Health    int
    MaxHealth int
    Strength  int
    Agility   int
    Magic     int
    Gold      int
    Inventory []string
    Level     int
    Exp       int
}

type Enemy struct {
    Name      string
    Health    int
    Strength  int
    Agility   int
    Magic     int
    GoldDrop  int
    ExpDrop   int
}

type Location struct {
    Name        string
    Description string
    Exits       map[string]string
    Enemies     []Enemy
    Items       []string
    IsSafe      bool
}

var player Character
var currentLocation Location
var locations map[string]Location
var gameRunning bool
var scanner *bufio.Scanner

func initGame() {
    rand.Seed(time.Now().UnixNano())
    scanner = bufio.NewScanner(os.Stdin)
    player = Character{
        Name:      "Hero",
        Health:    100,
        MaxHealth: 100,
        Strength:  10,
        Agility:   10,
        Magic:     10,
        Gold:      50,
        Inventory: []string{"Health Potion", "Rusty Sword"},
        Level:     1,
        Exp:       0,
    }
    locations = make(map[string]Location)
    initializeLocations()
    currentLocation = locations["town_square"]
    gameRunning = true
}

func initializeLocations() {
    locations["town_square"] = Location{
        Name:        "Town Square",
        Description: "You are in the bustling town square. People are going about their daily business. You see a fountain in the center.",
        Exits:       map[string]string{"north": "market", "east": "tavern", "south": "forest_entrance"},
        Enemies:     []Enemy{},
        Items:       []string{"Gold Coin"},
        IsSafe:      true,
    }
    locations["market"] = Location{
        Name:        "Market",
        Description: "A lively market with stalls selling various goods. You smell fresh bread and spices.",
        Exits:       map[string]string{"south": "town_square", "west": "blacksmith"},
        Enemies:     []Enemy{},
        Items:       []string{"Apple", "Bread"},
        IsSafe:      true,
    }
    locations["tavern"] = Location{
        Name:        "Tavern",
        Description: "A cozy tavern with a fireplace. Adventurers are sharing stories over ale.",
        Exits:       map[string]string{"west": "town_square", "upstairs": "inn_room"},
        Enemies:     []Enemy{},
        Items:       []string{"Ale"},
        IsSafe:      true,
    }
    locations["forest_entrance"] = Location{
        Name:        "Forest Entrance",
        Description: "The edge of a dark forest. You hear strange noises from within.",
        Exits:       map[string]string{"north": "town_square", "east": "forest_path"},
        Enemies:     []Enemy{{Name: "Goblin", Health: 30, Strength: 5, Agility: 8, Magic: 0, GoldDrop: 10, ExpDrop: 20}},
        Items:       []string{"Herbs"},
        IsSafe:      false,
    }
    locations["forest_path"] = Location{
        Name:        "Forest Path",
        Description: "A narrow path through dense trees. It's getting darker.",
        Exits:       map[string]string{"west": "forest_entrance", "north": "cave_entrance"},
        Enemies:     []Enemy{{Name: "Wolf", Health: 40, Strength: 8, Agility: 12, Magic: 0, GoldDrop: 15, ExpDrop: 30}},
        Items:       []string{"Mushrooms"},
        IsSafe:      false,
    }
    locations["cave_entrance"] = Location{
        Name:        "Cave Entrance",
        Description: "A dark cave opening. You feel a cold breeze from inside.",
        Exits:       map[string]string{"south": "forest_path", "inside": "cave_chamber"},
        Enemies:     []Enemy{{Name: "Troll", Health: 80, Strength: 15, Agility: 5, Magic: 0, GoldDrop: 50, ExpDrop: 100}},
        Items:       []string{"Torch"},
        IsSafe:      false,
    }
    locations["cave_chamber"] = Location{
        Name:        "Cave Chamber",
        Description: "A large chamber with glowing crystals. A treasure chest sits in the corner.",
        Exits:       map[string]string{"outside": "cave_entrance"},
        Enemies:     []Enemy{{Name: "Dragon", Health: 150, Strength: 25, Agility: 10, Magic: 20, GoldDrop: 200, ExpDrop: 500}},
        Items:       []string{"Treasure Chest"},
        IsSafe:      false,
    }
    locations["blacksmith"] = Location{
        Name:        "Blacksmith",
        Description: "A hot forge where weapons are being crafted. The blacksmith is hammering away.",
        Exits:       map[string]string{"east": "market"},
        Enemies:     []Enemy{},
        Items:       []string{"Iron Sword", "Shield"},
        IsSafe:      true,
    }
    locations["inn_room"] = Location{
        Name:        "Inn Room",
        Description: "A comfortable room with a bed. You can rest here to recover health.",
        Exits:       map[string]string{"downstairs": "tavern"},
        Enemies:     []Enemy{},
        Items:       []string{},
        IsSafe:      true,
    }
}

func showStatus() {
    fmt.Println("\n=== Status ===")
    fmt.Printf("Name: %s\n", player.Name)
    fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
    fmt.Printf("Level: %d (Exp: %d/%d)\n", player.Level, player.Exp, player.Level*100)
    fmt.Printf("Strength: %d, Agility: %d, Magic: %d\n", player.Strength, player.Agility, player.Magic)
    fmt.Printf("Gold: %d\n", player.Gold)
    fmt.Printf("Inventory: %v\n", player.Inventory)
    fmt.Println("==============")
}

func showLocation() {
    fmt.Printf("\nYou are at: %s\n", currentLocation.Name)
    fmt.Println(currentLocation.Description)
    if len(currentLocation.Exits) > 0 {
        fmt.Print("Exits: ")
        for dir, loc := range currentLocation.Exits {
            fmt.Printf("%s to %s, ", dir, loc)
        }
        fmt.Println()
    }
    if len(currentLocation.Items) > 0 {
        fmt.Printf("Items here: %v\n", currentLocation.Items)
    }
    if len(currentLocation.Enemies) > 0 && !currentLocation.IsSafe {
        fmt.Printf("Enemies here: ")
        for _, enemy := range currentLocation.Enemies {
            fmt.Printf("%s (Health: %d), ", enemy.Name, enemy.Health)
        }
        fmt.Println()
    }
}

func move(direction string) {
    dest, exists := currentLocation.Exits[direction]
    if !exists {
        fmt.Println("You cannot go that way.")
        return
    }
    newLoc, exists := locations[dest]
    if !exists {
        fmt.Println("That location does not exist.")
        return
    }
    currentLocation = newLoc
    fmt.Printf("You move %s to %s.\n", direction, currentLocation.Name)
    if !currentLocation.IsSafe && len(currentLocation.Enemies) > 0 {
        fmt.Println("Be careful! Enemies are nearby.")
    }
}

func fight() {
    if len(currentLocation.Enemies) == 0 {
        fmt.Println("There are no enemies to fight here.")
        return
    }
    enemy := &currentLocation.Enemies[0]
    fmt.Printf("You engage in combat with %s!\n", enemy.Name)
    for player.Health > 0 && enemy.Health > 0 {
        playerAttack := player.Strength + rand.Intn(10)
        enemyAttack := enemy.Strength + rand.Intn(10)
        if rand.Intn(100) < player.Agility {
            fmt.Printf("You dodge %s's attack!\n", enemy.Name)
            enemyAttack = 0
        }
        if rand.Intn(100) < enemy.Agility {
            fmt.Printf("%s dodges your attack!\n", enemy.Name)
            playerAttack = 0
        }
        enemy.Health -= playerAttack
        player.Health -= enemyAttack
        fmt.Printf("You hit %s for %d damage. %s's Health: %d\n", enemy.Name, playerAttack, enemy.Name, enemy.Health)
        fmt.Printf("%s hits you for %d damage. Your Health: %d\n", enemy.Name, enemyAttack, player.Health)
        time.Sleep(500 * time.Millisecond)
    }
    if player.Health <= 0 {
        fmt.Println("You have been defeated! Game Over.")
        gameRunning = false
        return
    }
    fmt.Printf("You defeated %s!\n", enemy.Name)
    player.Gold += enemy.GoldDrop
    player.Exp += enemy.ExpDrop
    fmt.Printf("You gained %d gold and %d experience.\n", enemy.GoldDrop, enemy.ExpDrop)
    currentLocation.Enemies = currentLocation.Enemies[1:]
    checkLevelUp()
}

func checkLevelUp() {
    for player.Exp >= player.Level*100 {
        player.Exp -= player.Level * 100
        player.Level++
        player.MaxHealth += 20
        player.Health = player.MaxHealth
        player.Strength += 2
        player.Agility += 2
        player.Magic += 2
        fmt.Printf("Congratulations! You leveled up to Level %d!\n", player.Level)
    }
}

func rest() {
    if currentLocation.Name != "Inn Room" {
        fmt.Println("You can only rest in the Inn Room.")
        return
    }
    player.Health = player.MaxHealth
    fmt.Println("You rest and recover all health.")
}

func useItem(item string) {
    for i, invItem := range player.Inventory {
        if invItem == item {
            switch item {
            case "Health Potion":
                heal := 30
                player.Health += heal
                if player.Health > player.MaxHealth {
                    player.Health = player.MaxHealth
                }
                fmt.Printf("You used a Health Potion and healed %d health. Health: %d/%d\n", heal, player.Health, player.MaxHealth)
            case "Rusty Sword", "Iron Sword":
                fmt.Println("You already have a sword equipped.")
            default:
                fmt.Printf("You used %s, but nothing happens.\n", item)
            }
            player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
            return
        }
    }
    fmt.Println("You don't have that item in your inventory.")
}

func pickUpItem(item string) {
    for i, locItem := range currentLocation.Items {
        if locItem == item {
            player.Inventory = append(player.Inventory, item)
            currentLocation.Items = append(currentLocation.Items[:i], currentLocation.Items[i+1:]...)
            fmt.Printf("You picked up %s.\n", item)
            return
        }
    }
    fmt.Println("That item is not here.")
}

func buyItem(item string) {
    if currentLocation.Name != "Market" && currentLocation.Name != "Blacksmith" {
        fmt.Println("You can only buy items in the Market or Blacksmith.")
        return
    }
    prices := map[string]int{
        "Apple":      5,
        "Bread":      10,
        "Ale":        15,
        "Health Potion": 50,
        "Iron Sword":  100,
        "Shield":      80,
    }
    price, exists := prices[item]
    if !exists {
        fmt.Println("That item is not for sale here.")
        return
    }
    if player.Gold >= price {
        player.Gold -= price
        player.Inventory = append(player.Inventory, item)
        fmt.Printf("You bought %s for %d gold. Remaining gold: %d\n", item, price, player.Gold)
    } else {
        fmt.Println("You don't have enough gold.")
    }
}

func handleCommand(command string) {
    parts := strings.Fields(command)
    if len(parts) == 0 {
        return
    }
    switch parts[0] {
    case "go", "move":
        if len(parts) < 2 {
            fmt.Println("Go where?")
        } else {
            move(parts[1])
        }
    case "fight":
        fight()
    case "rest":
        rest()
    case "use":
        if len(parts) < 2 {
            fmt.Println("Use what?")
        } else {
            useItem(parts[1])
        }
    case "pickup":
        if len(parts) < 2 {
            fmt.Println("Pick up what?")
        } else {
            pickUpItem(parts[1])
        }
    case "buy":
        if len(parts) < 2 {
            fmt.Println("Buy what?")
        } else {
            buyItem(parts[1])
        }
    case "status":
        showStatus()
    case "look":
        showLocation()
    case "help":
        printHelp()
    case "quit", "exit":
        fmt.Println("Thanks for playing! Goodbye.")
        gameRunning = false
    default:
        fmt.Println("Unknown command. Type 'help' for a list of commands.")
    }
}

func printHelp() {
    fmt.Println("\nAvailable commands:")
    fmt.Println("  go <direction> - Move in a direction (e.g., go north)")
    fmt.Println("  fight - Engage in combat with enemies")
    fmt.Println("  rest - Rest to recover health (only in Inn Room)")
    fmt.Println("  use <item> - Use an item from inventory")
    fmt.Println("  pickup <item> - Pick up an item from the location")
    fmt.Println("  buy <item> - Buy an item (only in Market or Blacksmith)")
    fmt.Println("  status - Show player status")
    fmt.Println("  look - Describe current location")
    fmt.Println("  help - Show this help message")
    fmt.Println("  quit - Exit the game")
}

func main() {
    initGame()
    fmt.Println("Welcome to the Fantasy Text Adventure!")
    fmt.Println("Type 'help' for a list of commands.")
    showLocation()
    for gameRunning {
        fmt.Print("\n> ")
        if !scanner.Scan() {
            break
        }
        command := scanner.Text()
        handleCommand(command)
    }
}
