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
	Name          string
	Health        int
	MaxHealth     int
	Attack        int
	Defense       int
	Level         int
	Experience    int
	Gold          int
	Inventory     []string
	CurrentRoom   *Room
}

type Room struct {
	ID          int
	Description string
	Exits       map[string]*Room
	Enemies     []*Enemy
	Items       []string
	IsLocked    bool
}

type Enemy struct {
	Name    string
	Health  int
	Attack  int
	Defense int
	Loot    []string
	Gold    int
}

var player Player
var rooms map[int]*Room
var currentRoomID int
var scanner *bufio.Scanner

func main() {
	initializeGame()
	fmt.Println("Welcome to the Adventure Game!")
	fmt.Println("Type 'help' for a list of commands.")
	gameLoop()
}

func initializeGame() {
	rand.Seed(time.Now().UnixNano())
	scanner = bufio.NewScanner(os.Stdin)

	player = Player{
		Name:        "Hero",
		Health:      100,
		MaxHealth:   100,
		Attack:      10,
		Defense:     5,
		Level:       1,
		Experience:  0,
		Gold:        50,
		Inventory:   []string{"Potion"},
		CurrentRoom: nil,
	}

	rooms = make(map[int]*Room)
	createRooms()
	currentRoomID = 1
	player.CurrentRoom = rooms[currentRoomID]
}

func createRooms() {
	rooms[1] = &Room{
		ID:          1,
		Description: "You are in a dark forest. There are paths to the north and east.",
		Exits:       make(map[string]*Room),
		Enemies:     []*Enemy{},
		Items:       []string{"Berry"},
		IsLocked:    false,
	}

	rooms[2] = &Room{
		ID:          2,
		Description: "You are in a clearing. A small cabin is to the west. Paths lead north and south.",
		Exits:       make(map[string]*Room),
		Enemies:     []*Enemy{{Name: "Goblin", Health: 30, Attack: 5, Defense: 2, Loot: []string{"Rusty Sword"}, Gold: 10}},
		Items:       []string{},
		IsLocked:    false,
	}

	rooms[3] = &Room{
		ID:          3,
		Description: "You are at a river. The water flows swiftly. There is a bridge to the east.",
		Exits:       make(map[string]*Room),
		Enemies:     []*Enemy{},
		Items:       []string{"Fishing Rod"},
		IsLocked:    false,
	}

	rooms[4] = &Room{
		ID:          4,
		Description: "You are in a cave. It's damp and dark. An exit is to the south.",
		Exits:       make(map[string]*Room),
		Enemies:     []*Enemy{{Name: "Troll", Health: 50, Attack: 8, Defense: 5, Loot: []string{"Gold Nugget"}, Gold: 25}},
		Items:       []string{"Torch"},
		IsLocked:    false,
	}

	rooms[5] = &Room{
		ID:          5,
		Description: "You are in a treasure room! Gold and jewels are everywhere, but the door is locked.",
		Exits:       make(map[string]*Room),
		Enemies:     []*Enemy{},
		Items:       []string{"Diamond", "Ruby"},
		IsLocked:    true,
	}

	rooms[1].Exits["north"] = rooms[2]
	rooms[1].Exits["east"] = rooms[3]
	rooms[2].Exits["south"] = rooms[1]
	rooms[2].Exits["north"] = rooms[4]
	rooms[2].Exits["west"] = rooms[5]
	rooms[3].Exits["west"] = rooms[1]
	rooms[3].Exits["east"] = rooms[4]
	rooms[4].Exits["south"] = rooms[2]
	rooms[4].Exits["west"] = rooms[3]
	rooms[5].Exits["east"] = rooms[2]
}

func gameLoop() {
	for {
		fmt.Printf("\n%s\n", player.CurrentRoom.Description)
		if len(player.CurrentRoom.Enemies) > 0 {
			fmt.Println("Enemies here:")
			for _, enemy := range player.CurrentRoom.Enemies {
				fmt.Printf(" - %s (Health: %d)\n", enemy.Name, enemy.Health)
			}
		}
		if len(player.CurrentRoom.Items) > 0 {
			fmt.Println("Items here:")
			for _, item := range player.CurrentRoom.Items {
				fmt.Printf(" - %s\n", item)
			}
		}
		fmt.Print("\nWhat do you want to do? ")
		scanner.Scan()
		input := strings.ToLower(scanner.Text())
		handleInput(input)
	}
}

func handleInput(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		fmt.Println("Please enter a command.")
		return
	}

	command := parts[0]
	switch command {
	case "help":
		printHelp()
	case "look":
		lookAround()
	case "go":
		if len(parts) < 2 {
			fmt.Println("Go where?")
			return
		}
		movePlayer(parts[1])
	case "attack":
		if len(parts) < 2 {
			fmt.Println("Attack what?")
			return
		}
		attackEnemy(parts[1])
	case "take":
		if len(parts) < 2 {
			fmt.Println("Take what?")
			return
		}
		takeItem(parts[1])
	case "inventory":
		showInventory()
	case "use":
		if len(parts) < 2 {
			fmt.Println("Use what?")
			return
		}
		useItem(parts[1])
	case "status":
		showStatus()
	case "quit":
		fmt.Println("Thanks for playing!")
		os.Exit(0)
	default:
		fmt.Println("Unknown command. Type 'help' for a list of commands.")
	}
}

func printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  help - Show this help message")
	fmt.Println("  look - Look around the current room")
	fmt.Println("  go <direction> - Move in a direction (north, south, east, west)")
	fmt.Println("  attack <enemy> - Attack an enemy")
	fmt.Println("  take <item> - Take an item from the room")
	fmt.Println("  inventory - Show your inventory")
	fmt.Println("  use <item> - Use an item from your inventory")
	fmt.Println("  status - Show your current status")
	fmt.Println("  quit - Quit the game")
}

func lookAround() {
	fmt.Printf("\n%s\n", player.CurrentRoom.Description)
	if len(player.CurrentRoom.Enemies) > 0 {
		fmt.Println("Enemies here:")
		for _, enemy := range player.CurrentRoom.Enemies {
			fmt.Printf(" - %s (Health: %d)\n", enemy.Name, enemy.Health)
		}
	}
	if len(player.CurrentRoom.Items) > 0 {
		fmt.Println("Items here:")
		for _, item := range player.CurrentRoom.Items {
			fmt.Printf(" - %s\n", item)
		}
	}
	fmt.Println("Exits:")
	for direction := range player.CurrentRoom.Exits {
		fmt.Printf(" - %s\n", direction)
	}
}

func movePlayer(direction string) {
	room, exists := player.CurrentRoom.Exits[direction]
	if !exists {
		fmt.Println("You can't go that way.")
		return
	}
	if room.IsLocked {
		fmt.Println("The door is locked. You need a key to enter.")
		return
	}
	player.CurrentRoom = room
	currentRoomID = room.ID
	fmt.Printf("You move %s.\n", direction)
}

func attackEnemy(enemyName string) {
	var target *Enemy
	for _, enemy := range player.CurrentRoom.Enemies {
		if strings.ToLower(enemy.Name) == enemyName {
			target = enemy
			break
		}
	}
	if target == nil {
		fmt.Println("No such enemy here.")
		return
	}

	playerDamage := player.Attack - target.Defense
	if playerDamage < 1 {
		playerDamage = 1
	}
	target.Health -= playerDamage
	fmt.Printf("You attack the %s for %d damage!\n", target.Name, playerDamage)

	if target.Health <= 0 {
		fmt.Printf("You defeated the %s!\n", target.Name)
		player.Gold += target.Gold
		player.Experience += 10
		fmt.Printf("You gained %d gold and 10 experience.\n", target.Gold)
		if len(target.Loot) > 0 {
			player.Inventory = append(player.Inventory, target.Loot...)
			fmt.Printf("You found: %s\n", strings.Join(target.Loot, ", "))
		}
		removeEnemy(target)
		checkLevelUp()
	} else {
		enemyDamage := target.Attack - player.Defense
		if enemyDamage < 1 {
			enemyDamage = 1
		}
		player.Health -= enemyDamage
		fmt.Printf("The %s attacks you for %d damage!\n", target.Name, enemyDamage)
		if player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			os.Exit(0)
		}
	}
}

func removeEnemy(enemy *Enemy) {
	for i, e := range player.CurrentRoom.Enemies {
		if e == enemy {
			player.CurrentRoom.Enemies = append(player.CurrentRoom.Enemies[:i], player.CurrentRoom.Enemies[i+1:]...)
			break
		}
	}
}

func checkLevelUp() {
	if player.Experience >= player.Level*50 {
		player.Level++
		player.Experience = 0
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 2
		player.Defense += 1
		fmt.Printf("Congratulations! You reached level %d!\n", player.Level)
		fmt.Println("Your stats have improved.")
	}
}

func takeItem(itemName string) {
	for i, item := range player.CurrentRoom.Items {
		if strings.ToLower(item) == itemName {
			player.Inventory = append(player.Inventory, item)
			player.CurrentRoom.Items = append(player.CurrentRoom.Items[:i], player.CurrentRoom.Items[i+1:]...)
			fmt.Printf("You took the %s.\n", item)
			return
		}
	}
	fmt.Println("No such item here.")
}

func showInventory() {
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("Your inventory:")
	for _, item := range player.Inventory {
		fmt.Printf(" - %s\n", item)
	}
}

func useItem(itemName string) {
	for i, item := range player.Inventory {
		if strings.ToLower(item) == itemName {
			switch item {
			case "potion":
				player.Health += 30
				if player.Health > player.MaxHealth {
					player.Health = player.MaxHealth
				}
				fmt.Println("You used a potion and restored 30 health.")
				player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			case "key":
				if player.CurrentRoom.ID == 5 {
					player.CurrentRoom.IsLocked = false
					fmt.Println("You used the key to unlock the treasure room!")
					player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
				} else {
					fmt.Println("There's nothing to unlock here.")
				}
			default:
				fmt.Printf("You can't use the %s right now.\n", item)
			}
			return
		}
	}
	fmt.Println("You don't have that item.")
}

func showStatus() {
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Level: %d\n", player.Level)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d\n", player.Attack)
	fmt.Printf("Defense: %d\n", player.Defense)
	fmt.Printf("Experience: %d/%d\n", player.Experience, player.Level*50)
	fmt.Printf("Gold: %d\n", player.Gold)
}
