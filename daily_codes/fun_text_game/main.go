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

// GameState holds the state of the game
type GameState struct {
	PlayerName  string
	Score       int
	Level       int
	Inventory   []string
	IsGameOver  bool
}

// Item represents an item in the game
type Item struct {
	Name        string
	Description string
	Value       int
}

// Enemy represents an enemy in the game
type Enemy struct {
	Name        string
	Health      int
	AttackPower int
}

// Function to initialize the game
func initializeGame() GameState {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	return GameState{
		PlayerName: name,
		Score:      0,
		Level:      1,
		Inventory:  []string{"Sword", "Potion"},
		IsGameOver: false,
	}
}

// Function to display the main menu
func displayMenu() {
	fmt.Println("\n=== Main Menu ===")
	fmt.Println("1. Start Game")
	fmt.Println("2. View Inventory")
	fmt.Println("3. View Score")
	fmt.Println("4. Exit")
	fmt.Print("Choose an option: ")
}

// Function to handle the game loop
func gameLoop(state *GameState) {
	reader := bufio.NewReader(os.Stdin)
	for !state.IsGameOver {
		fmt.Printf("\n=== Level %d ===\n", state.Level)
		fmt.Println("You are in a dark forest. What do you want to do?")
		fmt.Println("1. Explore")
		fmt.Println("2. Rest")
		fmt.Println("3. Check Inventory")
		fmt.Println("4. Quit Game")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			explore(state)
		case "2":
			rest(state)
		case "3":
			viewInventory(state)
		case "4":
			state.IsGameOver = true
			fmt.Println("Game over! Thanks for playing.")
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

// Function to handle exploration
func explore(state *GameState) {
	fmt.Println("You venture deeper into the forest...")
	rand.Seed(time.Now().UnixNano())
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Println("You found a treasure chest! +10 points.")
		state.Score += 10
		state.Inventory = append(state.Inventory, "Gold Coin")
	case 1:
		fmt.Println("You encountered a goblin!")
		fightEnemy(state, Enemy{Name: "Goblin", Health: 20, AttackPower: 5})
	case 2:
		fmt.Println("You discovered a hidden path. Level up!")
		state.Level++
		state.Score += 5
	}
}

// Function to handle resting
func rest(state *GameState) {
	fmt.Println("You take a rest and recover your strength.")
	state.Score += 2
}

// Function to view inventory
func viewInventory(state *GameState) {
	fmt.Println("\n=== Inventory ===")
	for _, item := range state.Inventory {
		fmt.Printf("- %s\n", item)
	}
}

// Function to handle enemy fights
func fightEnemy(state *GameState, enemy Enemy) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("A %s appears with %d health.\n", enemy.Name, enemy.Health)
	for enemy.Health > 0 && !state.IsGameOver {
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Run Away")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			damage := rand.Intn(10) + 5 // Player attack between 5-14
			fmt.Printf("You attack the %s for %d damage.\n", enemy.Name, damage)
			enemy.Health -= damage
			if enemy.Health <= 0 {
				fmt.Printf("You defeated the %s! +15 points.\n", enemy.Name)
				state.Score += 15
				break
			}
			// Enemy counterattack
			enemyDamage := rand.Intn(enemy.AttackPower) + 1
			fmt.Printf("The %s attacks you for %d damage.\n", enemy.Name, enemyDamage)
			// Simple health system: if score drops below 0, game over
			state.Score -= enemyDamage
			if state.Score < 0 {
				state.IsGameOver = true
				fmt.Println("You have been defeated! Game over.")
			}
		case "2":
			useItem(state)
		case "3":
			fmt.Println("You run away safely.")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

// Function to use an item from inventory
func useItem(state *GameState) {
	if len(state.Inventory) == 0 {
		fmt.Println("Your inventory is empty!")
		return
	}
	fmt.Println("Select an item to use:")
	for i, item := range state.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter item number: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input)
	if err != nil || index < 1 || index > len(state.Inventory) {
		fmt.Println("Invalid selection.")
		return
	}
	item := state.Inventory[index-1]
	if item == "Potion" {
		fmt.Println("You used a Potion and gained 10 health.")
		state.Score += 10
		// Remove the used item
		state.Inventory = append(state.Inventory[:index-1], state.Inventory[index:]...)
	} else {
		fmt.Printf("You can't use %s right now.\n", item)
	}
}

// Function to save game state (placeholder)
func saveGame(state GameState) {
	fmt.Println("Game saved (not implemented).")
}

// Function to load game state (placeholder)
func loadGame() GameState {
	fmt.Println("Game loaded (not implemented).")
	return GameState{}
}

// Main function
func main() {
	fmt.Println("Welcome to the Fun Text Adventure Game!")
	gameState := initializeGame()
	for {
		displayMenu()
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			gameLoop(&gameState)
		case "2":
			viewInventory(&gameState)
		case "3":
			fmt.Printf("Your score: %d\n", gameState.Score)
		case "4":
			fmt.Println("Exiting game. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}
