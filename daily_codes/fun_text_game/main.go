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

// Player struct to hold player information
type Player struct {
	Name     string
	Score    int
	Health   int
	Inventory []string
}

// GameState struct to manage the game state
type GameState struct {
	Player      Player
	CurrentRoom string
	Rooms       map[string]Room
	Quests      []Quest
	Enemies     []Enemy
}

// Room struct to define a room in the game
type Room struct {
	Name        string
	Description string
	Exits       map[string]string
	Items       []string
	Enemy       *Enemy
}

// Quest struct to define a quest
type Quest struct {
	ID          int
	Name        string
	Description string
	Completed   bool
}

// Enemy struct to define an enemy
type Enemy struct {
	Name   string
	Health int
	Damage int
}

// Function to initialize the game
func initGame() GameState {
	player := Player{
		Name:     "",
		Score:    0,
		Health:   100,
		Inventory: []string{},
	}

	rooms := map[string]Room{
		"start": {
			Name:        "Start Room",
			Description: "You are in a dimly lit room with a single door to the north.",
			Exits:       map[string]string{"north": "hallway"},
			Items:       []string{"key"},
			Enemy:       nil,
		},
		"hallway": {
			Name:        "Hallway",
			Description: "A long hallway stretches east and west. There is a door to the south.",
			Exits:       map[string]string{"south": "start", "east": "treasure", "west": "enemy"},
			Items:       []string{},
			Enemy:       nil,
		},
		"treasure": {
			Name:        "Treasure Room",
			Description: "You found a room filled with gold and jewels! But be careful, there might be traps.",
			Exits:       map[string]string{"west": "hallway"},
			Items:       []string{"gold", "potion"},
			Enemy:       nil,
		},
		"enemy": {
			Name:        "Enemy Room",
			Description: "A fierce goblin blocks your path! Prepare for battle.",
			Exits:       map[string]string{"east": "hallway"},
			Items:       []string{},
			Enemy:       &Enemy{Name: "Goblin", Health: 30, Damage: 10},
		},
	}

	quests := []Quest{
		{ID: 1, Name: "Find the Key", Description: "Locate the key in the start room.", Completed: false},
		{ID: 2, Name: "Defeat the Goblin", Description: "Defeat the goblin in the enemy room.", Completed: false},
		{ID: 3, Name: "Collect Treasure", Description: "Gather gold from the treasure room.", Completed: false},
	}

	enemies := []Enemy{
		{Name: "Goblin", Health: 30, Damage: 10},
	}

	return GameState{
		Player:      player,
		CurrentRoom: "start",
		Rooms:       rooms,
		Quests:      quests,
		Enemies:     enemies,
	}
}

// Function to print the current room description
func printRoomDescription(gs *GameState) {
	room := gs.Rooms[gs.CurrentRoom]
	fmt.Printf("You are in the %s. %s\n", room.Name, room.Description)
	if len(room.Items) > 0 {
		fmt.Printf("Items here: %s\n", strings.Join(room.Items, ", "))
	}
	if room.Enemy != nil {
		fmt.Printf("An enemy is here: %s (Health: %d)\n", room.Enemy.Name, room.Enemy.Health)
	}
	fmt.Printf("Exits: %s\n", getExitsString(room.Exits))
}

// Helper function to get exits as a string
func getExitsString(exits map[string]string) string {
	keys := make([]string, 0, len(exits))
	for k := range exits {
		keys = append(keys, k)
	}
	return strings.Join(keys, ", ")
}

// Function to handle player movement
func movePlayer(gs *GameState, direction string) {
	room := gs.Rooms[gs.CurrentRoom]
	if exit, ok := room.Exits[direction]; ok {
		gs.CurrentRoom = exit
		fmt.Printf("You move %s to the %s.\n", direction, gs.Rooms[exit].Name)
	} else {
		fmt.Println("You can't go that way.")
	}
}

// Function to handle taking items
func takeItem(gs *GameState, item string) {
	room := gs.Rooms[gs.CurrentRoom]
	for i, it := range room.Items {
		if it == item {
			gs.Player.Inventory = append(gs.Player.Inventory, item)
			room.Items = append(room.Items[:i], room.Items[i+1:]...)
			gs.Rooms[gs.CurrentRoom] = room // Update the room in the map
			fmt.Printf("You took the %s.\n", item)
			updateQuests(gs, item)
			return
		}
	}
	fmt.Println("Item not found here.")
}

// Function to update quests based on actions
func updateQuests(gs *GameState, action string) {
	for i := range gs.Quests {
		if !gs.Quests[i].Completed {
			switch gs.Quests[i].Name {
			case "Find the Key":
				if action == "key" {
					gs.Quests[i].Completed = true
					gs.Player.Score += 10
					fmt.Println("Quest completed: Find the Key! +10 points.")
				}
			case "Collect Treasure":
				if action == "gold" {
					gs.Quests[i].Completed = true
					gs.Player.Score += 20
					fmt.Println("Quest completed: Collect Treasure! +20 points.")
				}
			}
		}
	}
}

// Function to handle combat with an enemy
func combat(gs *GameState) {
	room := gs.Rooms[gs.CurrentRoom]
	if room.Enemy == nil {
		fmt.Println("No enemy here to fight.")
		return
	}
	enemy := room.Enemy
	fmt.Printf("Combat with %s!\n", enemy.Name)
	for enemy.Health > 0 && gs.Player.Health > 0 {
		// Player attacks
		damage := rand.Intn(15) + 5 // Random damage between 5 and 20
		enemy.Health -= damage
		fmt.Printf("You attack the %s for %d damage. Enemy health: %d\n", enemy.Name, damage, enemy.Health)
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			gs.Player.Score += 15
			room.Enemy = nil
			gs.Rooms[gs.CurrentRoom] = room // Update the room
			updateQuestCombat(gs, enemy.Name)
			break
		}
		// Enemy attacks
		enemyDamage := rand.Intn(enemy.Damage) + 1
		gs.Player.Health -= enemyDamage
		fmt.Printf("%s attacks you for %d damage. Your health: %d\n", enemy.Name, enemyDamage, gs.Player.Health)
		if gs.Player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			os.Exit(0)
		}
		time.Sleep(1 * time.Second) // Pause for effect
	}
}

// Function to update quest for combat
func updateQuestCombat(gs *GameState, enemyName string) {
	for i := range gs.Quests {
		if gs.Quests[i].Name == "Defeat the Goblin" && enemyName == "Goblin" && !gs.Quests[i].Completed {
			gs.Quests[i].Completed = true
			fmt.Println("Quest completed: Defeat the Goblin! +15 points.")
		}
	}
}

// Function to display player status
func showStatus(gs *GameState) {
	fmt.Printf("Name: %s, Health: %d, Score: %d\n", gs.Player.Name, gs.Player.Health, gs.Player.Score)
	fmt.Printf("Inventory: %s\n", strings.Join(gs.Player.Inventory, ", "))
	fmt.Println("Active Quests:")
	for _, quest := range gs.Quests {
		if !quest.Completed {
			fmt.Printf("- %s: %s\n", quest.Name, quest.Description)
		}
	}
}

// Function to handle using items
func useItem(gs *GameState, item string) {
	for i, it := range gs.Player.Inventory {
		if it == item {
			switch item {
			case "potion":
				gs.Player.Health += 20
				fmt.Printf("You used a potion. Health restored by 20. Current health: %d\n", gs.Player.Health)
				gs.Player.Inventory = append(gs.Player.Inventory[:i], gs.Player.Inventory[i+1:]...)
				return
			default:
				fmt.Println("You can't use that item.")
			}
		}
	}
	fmt.Println("Item not in inventory.")
}

// Main function to run the game
func main() {
	rand.Seed(time.Now().UnixNano())
	gs := initGame()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Text Adventure Game!")
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	gs.Player.Name = strings.TrimSpace(name)
	fmt.Printf("Hello, %s! Let's begin.\n", gs.Player.Name)
	fmt.Println("Type 'help' for a list of commands.")

	for {
		printRoomDescription(&gs)
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(strings.ToLower(input))

		switch input {
		case "help":
			fmt.Println("Commands: go [direction], take [item], use [item], fight, status, quit")
		case "quit":
			fmt.Println("Thanks for playing!")
			os.Exit(0)
		case "status":
			showStatus(&gs)
		case "fight":
			combat(&gs)
		default:
			if strings.HasPrefix(input, "go ") {
				direction := strings.TrimPrefix(input, "go ")
				movePlayer(&gs, direction)
			} else if strings.HasPrefix(input, "take ") {
				item := strings.TrimPrefix(input, "take ")
				takeItem(&gs, item)
			} else if strings.HasPrefix(input, "use ") {
				item := strings.TrimPrefix(input, "use ")
				useItem(&gs, item)
			} else {
				fmt.Println("Unknown command. Type 'help' for options.")
			}
		}

		// Check if all quests are completed
		allCompleted := true
		for _, quest := range gs.Quests {
			if !quest.Completed {
				allCompleted = false
				break
			}
		}
		if allCompleted {
			fmt.Println("Congratulations! You've completed all quests. Final score:", gs.Player.Score)
			os.Exit(0)
		}
	}
}
