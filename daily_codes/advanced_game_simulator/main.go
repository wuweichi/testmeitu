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
    Inventory []string
}

type Enemy struct {
    Name    string
    Health  int
    Attack  int
    Defense int
    Reward  int
}

type GameState struct {
    Player      Character
    Enemies     []Enemy
    CurrentRoom int
    GameOver    bool
    Score       int
}

func (c *Character) TakeDamage(damage int) {
    actualDamage := damage - c.Defense
    if actualDamage < 1 {
        actualDamage = 1
    }
    c.Health -= actualDamage
    if c.Health < 0 {
        c.Health = 0
    }
}

func (c *Character) Heal(amount int) {
    c.Health += amount
    if c.Health > c.MaxHealth {
        c.Health = c.MaxHealth
    }
}

func (c *Character) GainExperience(exp int) {
    c.Experience += exp
    if c.Experience >= c.Level*100 {
        c.LevelUp()
    }
}

func (c *Character) LevelUp() {
    c.Level++
    c.MaxHealth += 20
    c.Health = c.MaxHealth
    c.Attack += 5
    c.Defense += 3
    c.Experience = 0
    fmt.Printf("\n%s leveled up to level %d!\n", c.Name, c.Level)
}

func (e *Enemy) TakeDamage(damage int) bool {
    actualDamage := damage - e.Defense
    if actualDamage < 1 {
        actualDamage = 1
    }
    e.Health -= actualDamage
    return e.Health <= 0
}

func createEnemies() []Enemy {
    return []Enemy{
        {Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Reward: 10},
        {Name: "Orc", Health: 50, Attack: 8, Defense: 5, Reward: 20},
        {Name: "Dragon", Health: 100, Attack: 15, Defense: 10, Reward: 50},
        {Name: "Slime", Health: 20, Attack: 3, Defense: 1, Reward: 5},
        {Name: "Skeleton", Health: 40, Attack: 6, Defense: 3, Reward: 15},
        {Name: "Wizard", Health: 60, Attack: 12, Defense: 4, Reward: 30},
        {Name: "Troll", Health: 80, Attack: 10, Defense: 8, Reward: 40},
        {Name: "Ghost", Health: 35, Attack: 7, Defense: 0, Reward: 12},
        {Name: "Minotaur", Health: 90, Attack: 14, Defense: 9, Reward: 45},
        {Name: "Demon", Health: 120, Attack: 18, Defense: 12, Reward: 60},
    }
}

func initializeGame() GameState {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter your character name: ")
    name, _ := reader.ReadString('\n')
    name = strings.TrimSpace(name)
    if name == "" {
        name = "Hero"
    }

    player := Character{
        Name:      name,
        Health:    100,
        MaxHealth: 100,
        Attack:    10,
        Defense:   5,
        Level:     1,
        Experience: 0,
        Inventory: []string{"Potion", "Sword"},
    }

    return GameState{
        Player:      player,
        Enemies:     createEnemies(),
        CurrentRoom: 1,
        GameOver:    false,
        Score:       0,
    }
}

func displayStatus(state *GameState) {
    fmt.Println("\n=== Game Status ===")
    fmt.Printf("Player: %s (Level %d)\n", state.Player.Name, state.Player.Level)
    fmt.Printf("Health: %d/%d\n", state.Player.Health, state.Player.MaxHealth)
    fmt.Printf("Attack: %d, Defense: %d\n", state.Player.Attack, state.Player.Defense)
    fmt.Printf("Experience: %d/%d\n", state.Player.Experience, state.Player.Level*100)
    fmt.Printf("Score: %d\n", state.Score)
    fmt.Printf("Current Room: %d\n", state.CurrentRoom)
    fmt.Printf("Inventory: %v\n", state.Player.Inventory)
    fmt.Println("===================\n")
}

func battle(state *GameState, enemyIndex int) {
    enemy := &state.Enemies[enemyIndex]
    fmt.Printf("\nA wild %s appears!\n", enemy.Name)

    for state.Player.Health > 0 && enemy.Health > 0 {
        fmt.Printf("\n%s's Health: %d\n", enemy.Name, enemy.Health)
        fmt.Println("1. Attack")
        fmt.Println("2. Use Potion")
        fmt.Println("3. Run Away")
        fmt.Print("Choose an action: ")

        reader := bufio.NewReader(os.Stdin)
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            damage := state.Player.Attack + rand.Intn(5)
            fmt.Printf("You attack the %s for %d damage!\n", enemy.Name, damage)
            if enemy.TakeDamage(damage) {
                fmt.Printf("You defeated the %s!\n", enemy.Name)
                state.Player.GainExperience(enemy.Reward)
                state.Score += enemy.Reward * 10
                state.Enemies = append(state.Enemies[:enemyIndex], state.Enemies[enemyIndex+1:]...)
                return
            }
        case "2":
            if len(state.Player.Inventory) > 0 {
                for i, item := range state.Player.Inventory {
                    if item == "Potion" {
                        state.Player.Heal(30)
                        state.Player.Inventory = append(state.Player.Inventory[:i], state.Player.Inventory[i+1:]...)
                        fmt.Println("You used a Potion and healed 30 health.")
                        break
                    }
                }
            } else {
                fmt.Println("No potions in inventory!")
                continue
            }
        case "3":
            fmt.Println("You ran away safely.")
            return
        default:
            fmt.Println("Invalid choice. Try again.")
            continue
        }

        if enemy.Health > 0 {
            enemyDamage := enemy.Attack + rand.Intn(3)
            fmt.Printf("The %s attacks you for %d damage!\n", enemy.Name, enemyDamage)
            state.Player.TakeDamage(enemyDamage)
            if state.Player.Health <= 0 {
                fmt.Println("You have been defeated!")
                state.GameOver = true
                return
            }
        }
    }
}

func exploreRoom(state *GameState) {
    fmt.Printf("\n=== Room %d ===\n", state.CurrentRoom)
    if len(state.Enemies) == 0 {
        fmt.Println("No enemies left. You have cleared all rooms!")
        state.GameOver = true
        return
    }

    enemyIndex := rand.Intn(len(state.Enemies))
    battle(state, enemyIndex)

    if !state.GameOver {
        lootChance := rand.Intn(100)
        if lootChance < 30 {
            loot := "Potion"
            if lootChance < 10 {
                loot = "Super Potion"
                state.Player.Attack += 2
            }
            state.Player.Inventory = append(state.Player.Inventory, loot)
            fmt.Printf("You found a %s!\n", loot)
        }
        state.CurrentRoom++
    }
}

func saveGame(state GameState) error {
    data, err := json.Marshal(state)
    if err != nil {
        return err
    }
    return os.WriteFile("save_game.json", data, 0644)
}

func loadGame() (GameState, error) {
    data, err := os.ReadFile("save_game.json")
    if err != nil {
        return GameState{}, err
    }
    var state GameState
    err = json.Unmarshal(data, &state)
    return state, err
}

func mainMenu(state *GameState) {
    for !state.GameOver {
        displayStatus(state)
        fmt.Println("Main Menu:")
        fmt.Println("1. Explore Next Room")
        fmt.Println("2. View Inventory")
        fmt.Println("3. Save Game")
        fmt.Println("4. Load Game")
        fmt.Println("5. Quit")
        fmt.Print("Choose an option: ")

        reader := bufio.NewReader(os.Stdin)
        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            exploreRoom(state)
        case "2":
            fmt.Println("\nInventory:")
            for i, item := range state.Player.Inventory {
                fmt.Printf("%d. %s\n", i+1, item)
            }
            fmt.Println()
        case "3":
            if err := saveGame(*state); err != nil {
                fmt.Println("Error saving game:", err)
            } else {
                fmt.Println("Game saved successfully!")
            }
        case "4":
            loadedState, err := loadGame()
            if err != nil {
                fmt.Println("Error loading game:", err)
            } else {
                *state = loadedState
                fmt.Println("Game loaded successfully!")
            }
        case "5":
            fmt.Println("Thanks for playing!")
            state.GameOver = true
        default:
            fmt.Println("Invalid choice. Please try again.")
        }
    }

    fmt.Printf("\nGame Over! Final Score: %d\n", state.Score)
}

func main() {
    rand.Seed(time.Now().UnixNano())
    fmt.Println("Welcome to the Advanced Game Simulator!")
    fmt.Println("This is a text-based RPG where you battle enemies, gain experience, and explore rooms.")
    fmt.Println("Your goal is to defeat all enemies and achieve the highest score possible.\n")

    var state GameState
    fmt.Println("Do you want to load a saved game? (yes/no)")
    reader := bufio.NewReader(os.Stdin)
    response, _ := reader.ReadString('\n')
    response = strings.TrimSpace(strings.ToLower(response))

    if response == "yes" {
        loadedState, err := loadGame()
        if err != nil {
            fmt.Println("No saved game found. Starting a new game.")
            state = initializeGame()
        } else {
            state = loadedState
            fmt.Println("Game loaded from save.")
        }
    } else {
        state = initializeGame()
    }

    mainMenu(&state)
}
