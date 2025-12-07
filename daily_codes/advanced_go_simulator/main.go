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

type GameCharacter struct {
    Name      string
    Health    int
    MaxHealth int
    Strength  int
    Agility   int
    Level     int
    Experience int
    Inventory []string
}

type GameWorld struct {
    Name        string
    Description string
    Locations   []Location
    NPCs        []NPC
    Monsters    []Monster
}

type Location struct {
    ID          int
    Name        string
    Description string
    Exits       map[string]int
    Items       []string
}

type NPC struct {
    ID          int
    Name        string
    Dialogue    []string
    Quest       *Quest
}

type Monster struct {
    ID          int
    Name        string
    Health      int
    Strength    int
    Experience  int
    Drops       []string
}

type Quest struct {
    ID          int
    Title       string
    Description string
    Objective   string
    Reward      int
    Completed   bool
}

type GameState struct {
    Character   GameCharacter
    World       GameWorld
    CurrentLocationID int
    Quests      []Quest
    GameTime    time.Time
    mu          sync.Mutex
}

func NewGameCharacter(name string) GameCharacter {
    return GameCharacter{
        Name:      name,
        Health:    100,
        MaxHealth: 100,
        Strength:  10,
        Agility:   10,
        Level:     1,
        Experience: 0,
        Inventory: []string{"Health Potion", "Sword"},
    }
}

func NewGameWorld() GameWorld {
    return GameWorld{
        Name: "Fantasy Land",
        Description: "A magical world full of adventure and danger.",
        Locations: []Location{
            {ID: 1, Name: "Starting Village", Description: "A peaceful village where your journey begins.", Exits: map[string]int{"north": 2, "east": 3}, Items: []string{"Map", "Torch"}},
            {ID: 2, Name: "Dark Forest", Description: "A dense forest with mysterious creatures.", Exits: map[string]int{"south": 1, "east": 4}, Items: []string{"Herbs", "Mushrooms"}},
            {ID: 3, Name: "Riverbank", Description: "A serene river with clear water.", Exits: map[string]int{"west": 1, "north": 4}, Items: []string{"Fishing Rod", "Water Flask"}},
            {ID: 4, Name: "Ancient Ruins", Description: "Old ruins with hidden treasures.", Exits: map[string]int{"south": 2, "west": 3}, Items: []string{"Gold Coin", "Ancient Artifact"}},
        },
        NPCs: []NPC{
            {ID: 1, Name: "Old Sage", Dialogue: []string{"Welcome, traveler!", "Beware of the monsters in the forest.", "I have a quest for you if you're interested."}, Quest: &Quest{ID: 1, Title: "Defeat the Goblin", Description: "A goblin is terrorizing the village.", Objective: "Kill the goblin in the Dark Forest.", Reward: 50, Completed: false}},
            {ID: 2, Name: "Merchant", Dialogue: []string{"I sell potions and weapons.", "Check out my wares!"}, Quest: nil},
        },
        Monsters: []Monster{
            {ID: 1, Name: "Goblin", Health: 30, Strength: 5, Experience: 20, Drops: []string{"Goblin Ear", "Copper Coin"}},
            {ID: 2, Name: "Wolf", Health: 25, Strength: 8, Experience: 15, Drops: []string{"Wolf Pelt", "Meat"}},
            {ID: 3, Name: "Troll", Health: 50, Strength: 12, Experience: 40, Drops: []string{"Troll Blood", "Gold Coin"}},
        },
    }
}

func (gs *GameState) SaveGame(filename string) error {
    gs.mu.Lock()
    defer gs.mu.Unlock()
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    encoder := json.NewEncoder(file)
    encoder.SetIndent("", "  ")
    return encoder.Encode(gs)
}

func (gs *GameState) LoadGame(filename string) error {
    gs.mu.Lock()
    defer gs.mu.Unlock()
    file, err := os.Open(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    decoder := json.NewDecoder(file)
    return decoder.Decode(gs)
}

func (c *GameCharacter) Attack(target *Monster) (int, bool) {
    damage := c.Strength + rand.Intn(5)
    target.Health -= damage
    if target.Health <= 0 {
        c.Experience += target.Experience
        return damage, true
    }
    return damage, false
}

func (c *GameCharacter) TakeDamage(damage int) {
    c.Health -= damage
    if c.Health < 0 {
        c.Health = 0
    }
}

func (c *GameCharacter) LevelUp() {
    for c.Experience >= c.Level*100 {
        c.Experience -= c.Level * 100
        c.Level++
        c.MaxHealth += 20
        c.Health = c.MaxHealth
        c.Strength += 5
        c.Agility += 3
        fmt.Printf("Level up! You are now level %d.\n", c.Level)
    }
}

func (c *GameCharacter) UseItem(item string) bool {
    for i, invItem := range c.Inventory {
        if invItem == item {
            if item == "Health Potion" {
                c.Health += 30
                if c.Health > c.MaxHealth {
                    c.Health = c.MaxHealth
                }
                fmt.Println("Used Health Potion. Health restored.")
            }
            c.Inventory = append(c.Inventory[:i], c.Inventory[i+1:]...)
            return true
        }
    }
    return false
}

func (gs *GameState) Move(direction string) bool {
    currentLoc := gs.GetCurrentLocation()
    if nextID, ok := currentLoc.Exits[direction]; ok {
        gs.CurrentLocationID = nextID
        fmt.Printf("Moved %s to %s.\n", direction, gs.GetCurrentLocation().Name)
        return true
    }
    fmt.Println("You cannot go that way.")
    return false
}

func (gs *GameState) GetCurrentLocation() *Location {
    for i := range gs.World.Locations {
        if gs.World.Locations[i].ID == gs.CurrentLocationID {
            return &gs.World.Locations[i]
        }
    }
    return nil
}

func (gs *GameState) Look() {
    loc := gs.GetCurrentLocation()
    if loc == nil {
        fmt.Println("Location not found.")
        return
    }
    fmt.Printf("You are at: %s\n", loc.Name)
    fmt.Printf("Description: %s\n", loc.Description)
    fmt.Print("Exits: ")
    for dir := range loc.Exits {
        fmt.Printf("%s ", dir)
    }
    fmt.Println()
    if len(loc.Items) > 0 {
        fmt.Printf("Items here: %v\n", loc.Items)
    }
    for _, npc := range gs.World.NPCs {
        fmt.Printf("NPC: %s\n", npc.Name)
    }
}

func (gs *GameState) TalkToNPC(name string) {
    for i := range gs.World.NPCs {
        if gs.World.NPCs[i].Name == name {
            npc := &gs.World.NPCs[i]
            fmt.Printf("%s says:\n", npc.Name)
            for _, line := range npc.Dialogue {
                fmt.Printf("  - %s\n", line)
            }
            if npc.Quest != nil && !npc.Quest.Completed {
                fmt.Printf("Quest available: %s - %s\n", npc.Quest.Title, npc.Quest.Description)
                fmt.Print("Accept quest? (yes/no): ")
                reader := bufio.NewReader(os.Stdin)
                input, _ := reader.ReadString('\n')
                input = strings.TrimSpace(input)
                if input == "yes" {
                    gs.Quests = append(gs.Quests, *npc.Quest)
                    fmt.Println("Quest accepted!")
                }
            }
            return
        }
    }
    fmt.Println("NPC not found.")
}

func (gs *GameState) Fight(monsterName string) {
    var target *Monster
    for i := range gs.World.Monsters {
        if gs.World.Monsters[i].Name == monsterName {
            target = &gs.World.Monsters[i]
            break
        }
    }
    if target == nil {
        fmt.Println("Monster not found.")
        return
    }
    fmt.Printf("You engage in combat with %s!\n", target.Name)
    for gs.Character.Health > 0 && target.Health > 0 {
        damage, killed := gs.Character.Attack(target)
        fmt.Printf("You hit %s for %d damage.\n", target.Name, damage)
        if killed {
            fmt.Printf("%s defeated! You gain %d experience.\n", target.Name, target.Experience)
            gs.Character.LevelUp()
            for _, drop := range target.Drops {
                gs.Character.Inventory = append(gs.Character.Inventory, drop)
                fmt.Printf("You found: %s\n", drop)
            }
            for i := range gs.Quests {
                if gs.Quests[i].Objective == fmt.Sprintf("Kill the %s", target.Name) {
                    gs.Quests[i].Completed = true
                    gs.Character.Experience += gs.Quests[i].Reward
                    fmt.Printf("Quest completed! Reward: %d experience.\n", gs.Quests[i].Reward)
                }
            }
            target.Health = 0
            break
        }
        monsterDamage := target.Strength + rand.Intn(3)
        gs.Character.TakeDamage(monsterDamage)
        fmt.Printf("%s hits you for %d damage. Your health: %d/%d\n", target.Name, monsterDamage, gs.Character.Health, gs.Character.MaxHealth)
        time.Sleep(500 * time.Millisecond)
    }
    if gs.Character.Health <= 0 {
        fmt.Println("You have been defeated! Game over.")
        os.Exit(0)
    }
}

func (gs *GameState) ShowInventory() {
    fmt.Printf("Inventory: %v\n", gs.Character.Inventory)
    fmt.Printf("Health: %d/%d, Level: %d, Experience: %d\n", gs.Character.Health, gs.Character.MaxHealth, gs.Character.Level, gs.Character.Experience)
}

func (gs *GameState) ShowQuests() {
    if len(gs.Quests) == 0 {
        fmt.Println("No active quests.")
        return
    }
    fmt.Println("Active Quests:")
    for _, quest := range gs.Quests {
        status := "In Progress"
        if quest.Completed {
            status = "Completed"
        }
        fmt.Printf("  - %s: %s [%s]\n", quest.Title, quest.Objective, status)
    }
}

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("Welcome to Advanced Go Simulator!")
    fmt.Print("Enter your character name: ")
    reader := bufio.NewReader(os.Stdin)
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)
    if name == "" {
        name = "Hero"
    }
    character := NewGameCharacter(name)
    world := NewGameWorld()
    gameState := GameState{
        Character: character,
        World:     world,
        CurrentLocationID: 1,
        Quests:    []Quest{},
        GameTime:  time.Now(),
    }
    fmt.Printf("Hello, %s! Your adventure begins in %s.\n", character.Name, world.Name)
    gameState.Look()
    for {
        fmt.Print("> ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        parts := strings.Fields(input)
        if len(parts) == 0 {
            continue
        }
        command := parts[0]
        switch command {
        case "look":
            gameState.Look()
        case "go":
            if len(parts) < 2 {
                fmt.Println("Usage: go <direction>")
            } else {
                gameState.Move(parts[1])
            }
        case "talk":
            if len(parts) < 2 {
                fmt.Println("Usage: talk <npc_name>")
            } else {
                gameState.TalkToNPC(parts[1])
            }
        case "fight":
            if len(parts) < 2 {
                fmt.Println("Usage: fight <monster_name>")
            } else {
                gameState.Fight(parts[1])
            }
        case "inventory", "inv":
            gameState.ShowInventory()
        case "use":
            if len(parts) < 2 {
                fmt.Println("Usage: use <item>")
            } else {
                if !gameState.Character.UseItem(parts[1]) {
                    fmt.Println("Item not found in inventory.")
                }
            }
        case "quests":
            gameState.ShowQuests()
        case "save":
            if len(parts) < 2 {
                fmt.Println("Usage: save <filename>")
            } else {
                err := gameState.SaveGame(parts[1])
                if err != nil {
                    fmt.Printf("Error saving game: %v\n", err)
                } else {
                    fmt.Println("Game saved successfully.")
                }
            }
        case "load":
            if len(parts) < 2 {
                fmt.Println("Usage: load <filename>")
            } else {
                err := gameState.LoadGame(parts[1])
                if err != nil {
                    fmt.Printf("Error loading game: %v\n", err)
                } else {
                    fmt.Println("Game loaded successfully.")
                    gameState.Look()
                }
            }
        case "help":
            fmt.Println("Available commands:")
            fmt.Println("  look - Look around the current location")
            fmt.Println("  go <direction> - Move in a direction (e.g., north, south)")
            fmt.Println("  talk <npc_name> - Talk to an NPC")
            fmt.Println("  fight <monster_name> - Fight a monster")
            fmt.Println("  inventory/inv - Show inventory and stats")
            fmt.Println("  use <item> - Use an item from inventory")
            fmt.Println("  quests - Show active quests")
            fmt.Println("  save <filename> - Save game to file")
            fmt.Println("  load <filename> - Load game from file")
            fmt.Println("  help - Show this help message")
            fmt.Println("  quit - Exit the game")
        case "quit":
            fmt.Println("Thanks for playing! Goodbye.")
            return
        default:
            fmt.Println("Unknown command. Type 'help' for a list of commands.")
        }
    }
}
