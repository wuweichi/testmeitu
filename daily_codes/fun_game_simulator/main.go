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

type Player struct {
    Name      string
    Health    int
    MaxHealth int
    Attack    int
    Defense   int
    Level     int
    Experience int
    Inventory []string
    Gold      int
}

type Monster struct {
    Name      string
    Health    int
    Attack    int
    Defense   int
    ExperienceReward int
    GoldReward int
}

type GameState struct {
    Player    Player
    Monsters  []Monster
    GameOver  bool
    TurnCount int
}

func (p *Player) LevelUp() {
    p.Level++
    p.MaxHealth += 10
    p.Health = p.MaxHealth
    p.Attack += 2
    p.Defense += 1
    fmt.Printf("\n%s leveled up to level %d!\n", p.Name, p.Level)
}

func (p *Player) GainExperience(exp int) {
    p.Experience += exp
    fmt.Printf("\n%s gained %d experience points.\n", p.Name, exp)
    for p.Experience >= p.Level*100 {
        p.Experience -= p.Level * 100
        p.LevelUp()
    }
}

func (p *Player) TakeDamage(damage int) {
    actualDamage := damage - p.Defense
    if actualDamage < 1 {
        actualDamage = 1
    }
    p.Health -= actualDamage
    if p.Health < 0 {
        p.Health = 0
    }
    fmt.Printf("%s took %d damage. Health: %d/%d\n", p.Name, actualDamage, p.Health, p.MaxHealth)
}

func (m *Monster) TakeDamage(damage int) {
    actualDamage := damage - m.Defense
    if actualDamage < 1 {
        actualDamage = 1
    }
    m.Health -= actualDamage
    if m.Health < 0 {
        m.Health = 0
    }
    fmt.Printf("%s took %d damage. Health: %d\n", m.Name, actualDamage, m.Health)
}

func (p *Player) Heal(amount int) {
    p.Health += amount
    if p.Health > p.MaxHealth {
        p.Health = p.MaxHealth
    }
    fmt.Printf("%s healed for %d. Health: %d/%d\n", p.Name, amount, p.Health, p.MaxHealth)
}

func (p *Player) AddItem(item string) {
    p.Inventory = append(p.Inventory, item)
    fmt.Printf("Added %s to inventory.\n", item)
}

func (p *Player) UseItem(item string) bool {
    for i, invItem := range p.Inventory {
        if invItem == item {
            p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
            switch item {
            case "Health Potion":
                p.Heal(30)
            case "Attack Boost":
                p.Attack += 5
                fmt.Printf("%s's attack increased by 5.\n", p.Name)
            case "Defense Boost":
                p.Defense += 3
                fmt.Printf("%s's defense increased by 3.\n", p.Name)
            }
            return true
        }
    }
    fmt.Printf("Item %s not found in inventory.\n", item)
    return false
}

func (p *Player) DisplayStatus() {
    fmt.Printf("\n=== Player Status ===\n")
    fmt.Printf("Name: %s\n", p.Name)
    fmt.Printf("Level: %d\n", p.Level)
    fmt.Printf("Health: %d/%d\n", p.Health, p.MaxHealth)
    fmt.Printf("Attack: %d\n", p.Attack)
    fmt.Printf("Defense: %d\n", p.Defense)
    fmt.Printf("Experience: %d/%d\n", p.Experience, p.Level*100)
    fmt.Printf("Gold: %d\n", p.Gold)
    fmt.Printf("Inventory: %v\n", p.Inventory)
    fmt.Printf("====================\n")
}

func (m Monster) DisplayStatus() {
    fmt.Printf("\n=== Monster Status ===\n")
    fmt.Printf("Name: %s\n", m.Name)
    fmt.Printf("Health: %d\n", m.Health)
    fmt.Printf("Attack: %d\n", m.Attack)
    fmt.Printf("Defense: %d\n", m.Defense)
    fmt.Printf("====================\n")
}

func InitializeGame() GameState {
    rand.Seed(time.Now().UnixNano())
    player := Player{
        Name:      "Hero",
        Health:    100,
        MaxHealth: 100,
        Attack:    10,
        Defense:   5,
        Level:     1,
        Experience: 0,
        Inventory: []string{"Health Potion", "Health Potion"},
        Gold:      50,
    }
    monsters := []Monster{
        {Name: "Goblin", Health: 30, Attack: 5, Defense: 2, ExperienceReward: 20, GoldReward: 10},
        {Name: "Orc", Health: 50, Attack: 8, Defense: 4, ExperienceReward: 40, GoldReward: 20},
        {Name: "Dragon", Health: 100, Attack: 15, Defense: 10, ExperienceReward: 100, GoldReward: 50},
    }
    return GameState{
        Player:   player,
        Monsters: monsters,
        GameOver: false,
        TurnCount: 0,
    }
}

func Battle(player *Player, monster *Monster) bool {
    fmt.Printf("\nA wild %s appears!\n", monster.Name)
    for player.Health > 0 && monster.Health > 0 {
        fmt.Printf("\n--- Turn Start ---\n")
        player.DisplayStatus()
        monster.DisplayStatus()
        fmt.Println("\nChoose action:")
        fmt.Println("1. Attack")
        fmt.Println("2. Use Item")
        fmt.Println("3. Flee")
        reader := bufio.NewReader(os.Stdin)
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        switch input {
        case "1":
            playerDamage := player.Attack + rand.Intn(5)
            fmt.Printf("%s attacks %s for %d damage.\n", player.Name, monster.Name, playerDamage)
            monster.TakeDamage(playerDamage)
            if monster.Health <= 0 {
                fmt.Printf("%s defeated!\n", monster.Name)
                player.GainExperience(monster.ExperienceReward)
                player.Gold += monster.GoldReward
                fmt.Printf("Gained %d gold. Total gold: %d\n", monster.GoldReward, player.Gold)
                return true
            }
            monsterDamage := monster.Attack + rand.Intn(3)
            fmt.Printf("%s attacks %s for %d damage.\n", monster.Name, player.Name, monsterDamage)
            player.TakeDamage(monsterDamage)
            if player.Health <= 0 {
                fmt.Printf("%s has been defeated!\n", player.Name)
                return false
            }
        case "2":
            fmt.Println("Available items:")
            for i, item := range player.Inventory {
                fmt.Printf("%d. %s\n", i+1, item)
            }
            fmt.Print("Enter item number to use: ")
            itemInput, _ := reader.ReadString('\n')
            itemInput = strings.TrimSpace(itemInput)
            idx, err := strconv.Atoi(itemInput)
            if err != nil || idx < 1 || idx > len(player.Inventory) {
                fmt.Println("Invalid choice.")
                continue
            }
            item := player.Inventory[idx-1]
            if player.UseItem(item) {
                monsterDamage := monster.Attack + rand.Intn(3)
                fmt.Printf("%s attacks %s for %d damage.\n", monster.Name, player.Name, monsterDamage)
                player.TakeDamage(monsterDamage)
                if player.Health <= 0 {
                    fmt.Printf("%s has been defeated!\n", player.Name)
                    return false
                }
            }
        case "3":
            fleeChance := rand.Intn(100)
            if fleeChance < 50 {
                fmt.Println("Successfully fled from battle!")
                return false
            } else {
                fmt.Println("Failed to flee!")
                monsterDamage := monster.Attack + rand.Intn(3)
                fmt.Printf("%s attacks %s for %d damage.\n", monster.Name, player.Name, monsterDamage)
                player.TakeDamage(monsterDamage)
                if player.Health <= 0 {
                    fmt.Printf("%s has been defeated!\n", player.Name)
                    return false
                }
            }
        default:
            fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
        }
    }
    return false
}

func Shop(player *Player) {
    fmt.Println("\nWelcome to the Shop!")
    fmt.Printf("Your gold: %d\n", player.Gold)
    shopItems := map[string]int{
        "Health Potion": 20,
        "Attack Boost":  50,
        "Defense Boost": 40,
    }
    for {
        fmt.Println("\nAvailable items:")
        i := 1
        for item, price := range shopItems {
            fmt.Printf("%d. %s - %d gold\n", i, item, price)
            i++
        }
        fmt.Println("0. Exit Shop")
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter item number to buy: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        if input == "0" {
            fmt.Println("Leaving shop.")
            return
        }
        idx, err := strconv.Atoi(input)
        if err != nil || idx < 1 || idx > len(shopItems) {
            fmt.Println("Invalid choice.")
            continue
        }
        itemList := make([]string, 0, len(shopItems))
        for item := range shopItems {
            itemList = append(itemList, item)
        }
        selectedItem := itemList[idx-1]
        price := shopItems[selectedItem]
        if player.Gold >= price {
            player.Gold -= price
            player.AddItem(selectedItem)
            fmt.Printf("Bought %s for %d gold. Remaining gold: %d\n", selectedItem, price, player.Gold)
        } else {
            fmt.Println("Not enough gold!")
        }
    }
}

func SaveGame(state GameState) error {
    data, err := json.Marshal(state)
    if err != nil {
        return err
    }
    return os.WriteFile("savegame.json", data, 0644)
}

func LoadGame() (GameState, error) {
    data, err := os.ReadFile("savegame.json")
    if err != nil {
        return GameState{}, err
    }
    var state GameState
    err = json.Unmarshal(data, &state)
    return state, err
}

func main() {
    fmt.Println("=== Fun Game Simulator ===")
    fmt.Println("A text-based RPG adventure!")
    var gameState GameState
    reader := bufio.NewReader(os.Stdin)
    for {
        fmt.Println("\nMain Menu:")
        fmt.Println("1. Start New Game")
        fmt.Println("2. Load Game")
        fmt.Println("3. Exit")
        fmt.Print("Enter choice: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        switch input {
        case "1":
            gameState = InitializeGame()
            fmt.Println("New game started!")
            break
        case "2":
            loadedState, err := LoadGame()
            if err != nil {
                fmt.Println("No saved game found or error loading.")
                continue
            }
            gameState = loadedState
            fmt.Println("Game loaded successfully!")
            break
        case "3":
            fmt.Println("Thanks for playing! Goodbye!")
            return
        default:
            fmt.Println("Invalid choice. Please enter 1, 2, or 3.")
            continue
        }
        break
    }
    for !gameState.GameOver {
        gameState.TurnCount++
        fmt.Printf("\n=== Turn %d ===\n", gameState.TurnCount)
        gameState.Player.DisplayStatus()
        fmt.Println("\nChoose action:")
        fmt.Println("1. Battle a Monster")
        fmt.Println("2. Visit Shop")
        fmt.Println("3. Save Game")
        fmt.Println("4. Quit Game")
        reader := bufio.NewReader(os.Stdin)
        fmt.Print("Enter choice: ")
        input, _ := reader.ReadString('\n')
        input = strings.TrimSpace(input)
        switch input {
        case "1":
            if len(gameState.Monsters) == 0 {
                fmt.Println("No monsters left to battle! You've defeated them all!")
                gameState.GameOver = true
                continue
            }
            monsterIndex := rand.Intn(len(gameState.Monsters))
            monster := &gameState.Monsters[monsterIndex]
            won := Battle(&gameState.Player, monster)
            if won {
                gameState.Monsters = append(gameState.Monsters[:monsterIndex], gameState.Monsters[monsterIndex+1:]...)
                if len(gameState.Monsters) == 0 {
                    fmt.Println("\nCongratulations! You have defeated all monsters and won the game!")
                    gameState.GameOver = true
                }
            } else {
                fmt.Println("\nGame Over! You were defeated in battle.")
                gameState.GameOver = true
            }
        case "2":
            Shop(&gameState.Player)
        case "3":
            err := SaveGame(gameState)
            if err != nil {
                fmt.Printf("Error saving game: %v\n", err)
            } else {
                fmt.Println("Game saved successfully!")
            }
        case "4":
            fmt.Println("Quitting game. Thanks for playing!")
            gameState.GameOver = true
        default:
            fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
        }
    }
    fmt.Println("\nFinal Stats:")
    gameState.Player.DisplayStatus()
    fmt.Printf("Total turns played: %d\n", gameState.TurnCount)
    fmt.Println("=== Game Ended ===")
}
