package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Player represents a game player
type Player struct {
	Name     string
	Health   int
	Strength int
	Level    int
}

// Enemy represents an enemy in the game
type Enemy struct {
	Name     string
	Health   int
	Strength int
}

// Item represents a collectible item
type Item struct {
	Name        string
	Description string
	Value       int
}

// GameState holds the current state of the game
type GameState struct {
	Player    Player
	Enemies   []Enemy
	Items     []Item
	GameOver  bool
	Score     int
}

// NewPlayer creates a new player with default values
func NewPlayer(name string) Player {
	return Player{
		Name:     name,
		Health:   100,
		Strength: 10,
		Level:    1,
	}
}

// NewEnemy creates a new enemy with random stats
func NewEnemy() Enemy {
	names := []string{"Goblin", "Orc", "Dragon", "Skeleton"}
	rand.Seed(time.Now().UnixNano())
	name := names[rand.Intn(len(names))]
	health := rand.Intn(50) + 20
	strength := rand.Intn(15) + 5
	return Enemy{
		Name:     name,
		Health:   health,
		Strength: strength,
	}
}

// NewItem creates a new item with random properties
func NewItem() Item {
	names := []string{"Potion", "Sword", "Shield", "Gem"}
	descriptions := []string{"Restores health", "Increases strength", "Provides defense", "Valuable treasure"}
	rand.Seed(time.Now().UnixNano())
	name := names[rand.Intn(len(names))]
	description := descriptions[rand.Intn(len(descriptions))]
	value := rand.Intn(100) + 10
	return Item{
		Name:        name,
		Description: description,
		Value:       value,
	}
}

// InitializeGame sets up the initial game state
func InitializeGame(playerName string) GameState {
	player := NewPlayer(playerName)
	enemies := make([]Enemy, 5)
	for i := range enemies {
		enemies[i] = NewEnemy()
	}
	items := make([]Item, 10)
	for i := range items {
		items[i] = NewItem()
	}
	return GameState{
		Player:   player,
		Enemies:  enemies,
		Items:    items,
		GameOver: false,
		Score:    0,
	}
}

// Attack simulates an attack between attacker and defender
func Attack(attackerStrength int, defenderHealth int) (int, bool) {
	damage := rand.Intn(attackerStrength) + 1
	newHealth := defenderHealth - damage
	if newHealth <= 0 {
		return 0, true // defender is defeated
	}
	return newHealth, false
}

// UseItem applies the effect of an item on the player
func UseItem(item Item, player *Player) {
	switch item.Name {
	case "Potion":
		player.Health += 20
		fmt.Printf("Used %s. Health increased by 20.\n", item.Name)
	case "Sword":
		player.Strength += 5
		fmt.Printf("Used %s. Strength increased by 5.\n", item.Name)
	case "Shield":
		player.Health += 10
		fmt.Printf("Used %s. Health increased by 10.\n", item.Name)
	case "Gem":
		player.Level += 1
		fmt.Printf("Used %s. Level increased by 1.\n", item.Name)
	}
}

// DisplayStatus shows the current status of the player
func DisplayStatus(player Player) {
	fmt.Printf("Player: %s, Health: %d, Strength: %d, Level: %d\n", player.Name, player.Health, player.Strength, player.Level)
}

// DisplayEnemies lists all enemies
func DisplayEnemies(enemies []Enemy) {
	fmt.Println("Enemies:")
	for i, enemy := range enemies {
		fmt.Printf("%d: %s (Health: %d, Strength: %d)\n", i+1, enemy.Name, enemy.Health, enemy.Strength)
	}
}

// DisplayItems lists all items
func DisplayItems(items []Item) {
	fmt.Println("Items:")
	for i, item := range items {
		fmt.Printf("%d: %s - %s (Value: %d)\n", i+1, item.Name, item.Description, item.Value)
	}
}

// Main game loop
func GameLoop(state *GameState) {
	for !state.GameOver {
		fmt.Println("\n--- Game Menu ---")
		fmt.Println("1. View Status")
		fmt.Println("2. View Enemies")
		fmt.Println("3. View Items")
		fmt.Println("4. Attack an Enemy")
		fmt.Println("5. Use an Item")
		fmt.Println("6. Quit Game")
		fmt.Print("Choose an option: ")
		var choice int
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}
		switch choice {
		case 1:
			DisplayStatus(state.Player)
		case 2:
			DisplayEnemies(state.Enemies)
		case 3:
			DisplayItems(state.Items)
		case 4:
			if len(state.Enemies) == 0 {
				fmt.Println("No enemies left!")
				continue
			}
			DisplayEnemies(state.Enemies)
			fmt.Print("Select enemy to attack (number): ")
			var enemyIndex int
			_, err := fmt.Scan(&enemyIndex)
			if err != nil || enemyIndex < 1 || enemyIndex > len(state.Enemies) {
				fmt.Println("Invalid enemy selection.")
				continue
			}
			enemyIndex-- // adjust for 0-based index
			newHealth, defeated := Attack(state.Player.Strength, state.Enemies[enemyIndex].Health)
			if defeated {
				fmt.Printf("You defeated the %s!\n", state.Enemies[enemyIndex].Name)
				state.Score += 10
				// Remove defeated enemy
				state.Enemies = append(state.Enemies[:enemyIndex], state.Enemies[enemyIndex+1:]...)
			} else {
				state.Enemies[enemyIndex].Health = newHealth
				fmt.Printf("You attacked the %s. Its health is now %d.\n", state.Enemies[enemyIndex].Name, newHealth)
			}
			// Enemy counter-attack
			if len(state.Enemies) > 0 && !defeated {
				enemy := state.Enemies[enemyIndex]
				newPlayerHealth, playerDefeated := Attack(enemy.Strength, state.Player.Health)
				if playerDefeated {
					state.Player.Health = 0
					fmt.Printf("You were defeated by the %s! Game Over.\n", enemy.Name)
					state.GameOver = true
				} else {
					state.Player.Health = newPlayerHealth
					fmt.Printf("The %s attacked you. Your health is now %d.\n", enemy.Name, newPlayerHealth)
				}
			}
		case 5:
			if len(state.Items) == 0 {
				fmt.Println("No items left!")
				continue
			}
			DisplayItems(state.Items)
			fmt.Print("Select item to use (number): ")
			var itemIndex int
			_, err := fmt.Scan(&itemIndex)
			if err != nil || itemIndex < 1 || itemIndex > len(state.Items) {
				fmt.Println("Invalid item selection.")
				continue
			}
			itemIndex-- // adjust for 0-based index
			UseItem(state.Items[itemIndex], &state.Player)
			// Remove used item
			state.Items = append(state.Items[:itemIndex], state.Items[itemIndex+1:]...)
		case 6:
			fmt.Println("Quitting game. Thanks for playing!")
			state.GameOver = true
		default:
			fmt.Println("Invalid option, please try again.")
		}
		// Check win condition: all enemies defeated
		if len(state.Enemies) == 0 {
			fmt.Println("Congratulations! You defeated all enemies and won the game!")
			fmt.Printf("Final Score: %d\n", state.Score)
			state.GameOver = true
		}
		// Check lose condition: player health zero
		if state.Player.Health <= 0 {
			fmt.Println("Game Over! Your health reached zero.")
			state.GameOver = true
		}
	}
}

func main() {
	fmt.Println("Welcome to the Fun Game Simulator!")
	fmt.Print("Enter your player name: ")
	var playerName string
	_, err := fmt.Scan(&playerName)
	if err != nil {
		fmt.Println("Error reading input, using default name.")
		playerName = "Player"
	}
	gameState := InitializeGame(playerName)
	GameLoop(&gameState)
	fmt.Println("Game ended. Score:", gameState.Score)
}
