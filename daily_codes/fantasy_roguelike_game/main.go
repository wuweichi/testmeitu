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
	Name          string
	Health        int
	MaxHealth     int
	Attack        int
	Defense       int
	Level         int
	Experience    int
	Gold          int
	Inventory     []Item
	EquippedWeapon *Item
	EquippedArmor  *Item
}

type Item struct {
	Name        string
	Description string
	Value       int
	AttackBonus int
	DefenseBonus int
	Type        string // "weapon", "armor", "potion", "key"
}

type Monster struct {
	Name      string
	Health    int
	Attack    int
	Defense   int
	Experience int
	Gold      int
	Loot      []Item
}

type Room struct {
	Description string
	Monsters    []Monster
	Items       []Item
	Exits       map[string]int // direction -> room index
	IsLocked    bool
}

var player Character
var rooms []Room
var currentRoom int
var gameOver bool

func initGame() {
	rand.Seed(time.Now().UnixNano())
	player = Character{
		Name:       "Hero",
		Health:     100,
		MaxHealth:  100,
		Attack:     10,
		Defense:    5,
		Level:      1,
		Experience: 0,
		Gold:       50,
		Inventory:  []Item{},
	}
	rooms = []Room{
		{
			Description: "You are in a dark, damp cave entrance. Torches flicker on the walls.",
			Monsters:    []Monster{},
			Items:       []Item{{Name: "Rusty Sword", Description: "A basic sword with some wear.", Value: 5, AttackBonus: 3, Type: "weapon"}},
			Exits:       map[string]int{"north": 1, "east": 2},
			IsLocked:    false,
		},
		{
			Description: "A large chamber with glowing mushrooms. Strange noises echo.",
			Monsters:    []Monster{{Name: "Goblin", Health: 30, Attack: 8, Defense: 2, Experience: 20, Gold: 10, Loot: []Item{{Name: "Health Potion", Description: "Restores 30 health.", Value: 10, Type: "potion"}}}},
			Items:       []Item{},
			Exits:       map[string]int{"south": 0, "west": 3},
			IsLocked:    false,
		},
		{
			Description: "A narrow tunnel with dripping water. It smells musty.",
			Monsters:    []Monster{},
			Items:       []Item{{Name: "Leather Armor", Description: "Light armor offering some protection.", Value: 15, DefenseBonus: 5, Type: "armor"}},
			Exits:       map[string]int{"west": 0, "north": 4},
			IsLocked:    false,
		},
		{
			Description: "A treasure room! Gold coins glitter, but a dragon guards it.",
			Monsters:    []Monster{{Name: "Dragon", Health: 200, Attack: 25, Defense: 15, Experience: 100, Gold: 100, Loot: []Item{{Name: "Dragon Key", Description: "A key to unlock special doors.", Value: 0, Type: "key"}}}},
			Items:       []Item{{Name: "Treasure Chest", Description: "Filled with 500 gold!", Value: 500, Type: "potion"}},
			Exits:       map[string]int{"east": 1},
			IsLocked:    true,
		},
		{
			Description: "A dead end with ancient writings on the wall. It feels eerie.",
			Monsters:    []Monster{},
			Items:       []Item{{Name: "Mysterious Scroll", Description: "Grants 50 experience when used.", Value: 0, Type: "potion"}},
			Exits:       map[string]int{"south": 2},
			IsLocked:    false,
		},
	}
	currentRoom = 0
	gameOver = false
}

func displayStatus() {
	fmt.Printf("\n=== %s (Level %d) ===\n", player.Name, player.Level)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Attack: %d, Defense: %d\n", player.Attack, player.Defense)
	fmt.Printf("Experience: %d/%d\n", player.Experience, player.Level*100)
	fmt.Printf("Gold: %d\n", player.Gold)
	weapon := "None"
	if player.EquippedWeapon != nil {
		weapon = player.EquippedWeapon.Name
	}
	armor := "None"
	if player.EquippedArmor != nil {
		armor = player.EquippedArmor.Name
	}
	fmt.Printf("Equipped: Weapon=%s, Armor=%s\n", weapon, armor)
}

func displayRoom() {
	room := rooms[currentRoom]
	fmt.Printf("\n%s\n", room.Description)
	if len(room.Monsters) > 0 {
		fmt.Println("Monsters here:")
		for _, m := range room.Monsters {
			fmt.Printf("  - %s (Health: %d)\n", m.Name, m.Health)
		}
	}
	if len(room.Items) > 0 {
		fmt.Println("Items here:")
		for _, item := range room.Items {
			fmt.Printf("  - %s: %s\n", item.Name, item.Description)
		}
	}
	fmt.Print("Exits: ")
	first := true
	for dir := range room.Exits {
		if !first {
			fmt.Print(", ")
		}
		fmt.Print(dir)
		first = false
	}
	fmt.Println()
}

func move(direction string) {
	room := rooms[currentRoom]
	if dest, ok := room.Exits[direction]; ok {
		if rooms[dest].IsLocked {
			fmt.Println("The door is locked! You need a key.")
			return
		}
		currentRoom = dest
		fmt.Printf("You move %s.\n", direction)
	} else {
		fmt.Println("You can't go that way.")
	}
}

func attackMonster() {
	room := rooms[currentRoom]
	if len(room.Monsters) == 0 {
		fmt.Println("There are no monsters to attack here.")
		return
	}
	monster := &room.Monsters[0]
	playerDamage := player.Attack - monster.Defense
	if playerDamage < 1 {
		playerDamage = 1
	}
	monster.Health -= playerDamage
	fmt.Printf("You attack the %s for %d damage.\n", monster.Name, playerDamage)
	if monster.Health <= 0 {
		fmt.Printf("You defeated the %s!\n", monster.Name)
		player.Experience += monster.Experience
		player.Gold += monster.Gold
		if len(monster.Loot) > 0 {
			fmt.Println("Loot dropped:")
			for _, item := range monster.Loot {
				player.Inventory = append(player.Inventory, item)
				fmt.Printf("  - %s\n", item.Name)
			}
		}
		rooms[currentRoom].Monsters = room.Monsters[1:]
		checkLevelUp()
	} else {
		monsterDamage := monster.Attack - player.Defense
		if monsterDamage < 1 {
			monsterDamage = 1
		}
		player.Health -= monsterDamage
		fmt.Printf("The %s attacks you for %d damage.\n", monster.Name, monsterDamage)
		if player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			gameOver = true
		}
	}
}

func checkLevelUp() {
	for player.Experience >= player.Level*100 {
		player.Experience -= player.Level * 100
		player.Level++
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Attack += 3
		player.Defense += 2
		fmt.Printf("Level up! You are now level %d. Health restored.\n", player.Level)
	}
}

func useItem(itemName string) {
	for i, item := range player.Inventory {
		if strings.ToLower(item.Name) == strings.ToLower(itemName) {
			switch item.Type {
			case "potion":
				if strings.Contains(item.Name, "Health") {
					healAmount := 30
					player.Health += healAmount
					if player.Health > player.MaxHealth {
						player.Health = player.MaxHealth
					}
					fmt.Printf("You used %s and healed %d health.\n", item.Name, healAmount)
				} else if strings.Contains(item.Name, "Scroll") {
					player.Experience += 50
					fmt.Printf("You used %s and gained 50 experience.\n", item.Name)
					checkLevelUp()
				}
			case "weapon":
				if player.EquippedWeapon != nil {
					player.Inventory = append(player.Inventory, *player.EquippedWeapon)
				}
				player.EquippedWeapon = &item
				player.Attack += item.AttackBonus
				fmt.Printf("You equipped %s. Attack increased by %d.\n", item.Name, item.AttackBonus)
			case "armor":
				if player.EquippedArmor != nil {
					player.Inventory = append(player.Inventory, *player.EquippedArmor)
				}
				player.EquippedArmor = &item
				player.Defense += item.DefenseBonus
				fmt.Printf("You equipped %s. Defense increased by %d.\n", item.Name, item.DefenseBonus)
			case "key":
				fmt.Println("Keys are used automatically when trying to open locked doors.")
			default:
				fmt.Printf("You used %s, but nothing happened.\n", item.Name)
			}
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			return
		}
	}
	fmt.Println("Item not found in inventory.")
}

func pickUpItem(itemName string) {
	room := rooms[currentRoom]
	for i, item := range room.Items {
		if strings.ToLower(item.Name) == strings.ToLower(itemName) {
			player.Inventory = append(player.Inventory, item)
			fmt.Printf("You picked up %s.\n", item.Name)
			if item.Type == "potion" && item.Name == "Treasure Chest" {
				player.Gold += item.Value
				fmt.Printf("You found %d gold in the chest!\n", item.Value)
			}
			rooms[currentRoom].Items = append(room.Items[:i], room.Items[i+1:]...)
			return
		}
	}
	fmt.Println("Item not found in the room.")
}

func showInventory() {
	if len(player.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
		return
	}
	fmt.Println("Inventory:")
	for _, item := range player.Inventory {
		fmt.Printf("  - %s: %s\n", item.Name, item.Description)
	}
}

func unlockDoor() {
	room := rooms[currentRoom]
	if !room.IsLocked {
		fmt.Println("This door is not locked.")
		return
	}
	for i, item := range player.Inventory {
		if item.Type == "key" && strings.Contains(item.Name, "Dragon") {
			room.IsLocked = false
			fmt.Println("You used the Dragon Key to unlock the door!")
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			rooms[currentRoom] = room
			return
		}
	}
	fmt.Println("You need a Dragon Key to unlock this door.")
}

func handleCommand(cmd string) {
	parts := strings.Fields(cmd)
	if len(parts) == 0 {
		return
	}
	action := strings.ToLower(parts[0])
	switch action {
	case "go", "move":
		if len(parts) > 1 {
			move(strings.ToLower(parts[1]))
		} else {
			fmt.Println("Specify a direction (e.g., go north).")
		}
	case "attack":
		attackMonster()
	case "use":
		if len(parts) > 1 {
			useItem(strings.Join(parts[1:], " "))
		} else {
			fmt.Println("Specify an item to use (e.g., use Health Potion).")
		}
	case "pickup", "take":
		if len(parts) > 1 {
			pickUpItem(strings.Join(parts[1:], " "))
		} else {
			fmt.Println("Specify an item to pick up (e.g., pickup Rusty Sword).")
		}
	case "inventory", "inv":
		showInventory()
	case "status":
		displayStatus()
	case "look":
		displayRoom()
	case "unlock":
		unlockDoor()
	case "quit", "exit":
		fmt.Println("Thanks for playing!")
		gameOver = true
	default:
		fmt.Println("Unknown command. Try: go [direction], attack, use [item], pickup [item], inventory, status, look, unlock, quit.")
	}
}

func main() {
	initGame()
	fmt.Println("Welcome to Fantasy Roguelike Game!")
	fmt.Println("Type 'help' for commands, or start exploring.")
	reader := bufio.NewReader(os.Stdin)
	for !gameOver {
		displayStatus()
		displayRoom()
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		if input == "help" {
			fmt.Println("Commands: go [north/south/east/west], attack, use [item], pickup [item], inventory, status, look, unlock, quit.")
			continue
		}
		handleCommand(input)
	}
	fmt.Println("Game ended.")
}
