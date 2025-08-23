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
	AttackPower   int
	Defense       int
	Gold          int
	Inventory     []string
	CurrentRoom   string
}

type Room struct {
	Name        string
	Description string
	Exits       map[string]string
	Items       []string
	Enemies     []Enemy
}

type Enemy struct {
	Name        string
	Health      int
	AttackPower int
	Defense     int
	LootGold    int
	LootItems   []string
}

var rooms map[string]Room
var player Player

func initGame() {
	rand.Seed(time.Now().UnixNano())
	rooms = make(map[string]Room)
	rooms["start"] = Room{
		Name:        "Starting Room",
		Description: "You are in a dimly lit room with stone walls. There is a door to the north and a chest in the corner.",
		Exits:       map[string]string{"north": "hallway"},
		Items:       []string{"health_potion"},
		Enemies:     []Enemy{},
	}
	rooms["hallway"] = Room{
		Name:        "Hallway",
		Description: "A long hallway stretches east and west. Torches flicker on the walls.",
		Exits:       map[string]string{"east": "treasure_room", "west": "armory", "south": "start"},
		Items:       []string{},
		Enemies:     []Enemy{{Name: "Goblin", Health: 20, AttackPower: 5, Defense: 2, LootGold: 10, LootItems: []string{"small_key"}}},
	}
	rooms["treasure_room"] = Room{
		Name:        "Treasure Room",
		Description: "A room filled with gold coins and jewels. A large door is to the north, but it's locked.",
		Exits:       map[string]string{"west": "hallway"},
		Items:       []string{"gold", "gold", "gem"},
		Enemies:     []Enemy{},
	}
	rooms["armory"] = Room{
		Name:        "Armory",
		Description: "Weapons and armor line the walls. There's an exit to the east.",
		Exits:       map[string]string{"east": "hallway"},
		Items:       []string{"sword", "shield"},
		Enemies:     []Enemy{{Name: "Skeleton", Health: 25, AttackPower: 7, Defense: 3, LootGold: 15, LootItems: []string{"health_potion"}}},
	}
	rooms["boss_room"] = Room{
		Name:        "Boss Room",
		Description: "A massive chamber with a dragon sleeping on a pile of gold. The only exit is south.",
		Exits:       map[string]string{"south": "treasure_room"},
		Items:       []string{},
		Enemies:     []Enemy{{Name: "Dragon", Health: 50, AttackPower: 15, Defense: 10, LootGold: 100, LootItems: []string{"dragon_scale", "victory_key"}}},
	}
	player = Player{
		Name:        "Hero",
		Health:      30,
		MaxHealth:   30,
		AttackPower: 8,
		Defense:     4,
		Gold:        0,
		Inventory:   []string{},
		CurrentRoom: "start",
	}
}

func main() {
	initGame()
	fmt.Println("Welcome to the Text Adventure Game! Type 'help' for commands.")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		room := rooms[player.CurrentRoom]
		fmt.Printf("\nYou are in the %s. %s\n", room.Name, room.Description)
		fmt.Printf("Exits: ")
		for dir := range room.Exits {
			fmt.Printf("%s ", dir)
		}
		fmt.Println()
		if len(room.Items) > 0 {
			fmt.Printf("Items here: %v\n", room.Items)
		}
		if len(room.Enemies) > 0 {
			fmt.Printf("Enemies here: ")
			for _, enemy := range room.Enemies {
				fmt.Printf("%s (Health: %d) ", enemy.Name, enemy.Health)
			}
			fmt.Println()
		}
		fmt.Printf("Your health: %d/%d, Gold: %d, Inventory: %v\n", player.Health, player.MaxHealth, player.Gold, player.Inventory)
		fmt.Print("> ")
		scanner.Scan()
		input := strings.ToLower(strings.TrimSpace(scanner.Text()))
		if input == "quit" {
			fmt.Println("Thanks for playing!")
			break
		}
		handleCommand(input)
		if player.Health <= 0 {
			fmt.Println("You have died! Game over.")
			break
		}
		if hasItem("victory_key") && player.CurrentRoom == "start" {
			fmt.Println("You have the victory key and returned to the start! You win!")
			break
		}
	}
}

func handleCommand(input string) {
	parts := strings.Fields(input)
	if len(parts) == 0 {
		return
	}
	command := parts[0]
	switch command {
	case "go":
		if len(parts) < 2 {
			fmt.Println("Go where?")
			return
		}
		direction := parts[1]
		room := rooms[player.CurrentRoom]
		if exit, ok := room.Exits[direction]; ok {
			if direction == "north" && player.CurrentRoom == "treasure_room" && !hasItem("small_key") {
				fmt.Println("The door to the north is locked. You need a key.")
				return
			}
			player.CurrentRoom = exit
			fmt.Printf("You go %s to the %s.\n", direction, rooms[exit].Name)
		} else {
			fmt.Println("You can't go that way.")
		}
	case "get":
		if len(parts) < 2 {
			fmt.Println("Get what?")
			return
		}
		item := parts[1]
		room := rooms[player.CurrentRoom]
		for i, it := range room.Items {
			if it == item {
				player.Inventory = append(player.Inventory, item)
				room.Items = append(room.Items[:i], room.Items[i+1:]...)
				rooms[player.CurrentRoom] = room
				fmt.Printf("You picked up the %s.\n", item)
				return
			}
		}
		fmt.Println("That item isn't here.")
	case "use":
		if len(parts) < 2 {
			fmt.Println("Use what?")
			return
		}
		item := parts[1]
		if !hasItem(item) {
			fmt.Println("You don't have that item.")
			return
		}
		if item == "health_potion" {
			player.Health = player.MaxHealth
			removeItem(item)
			fmt.Println("You used a health potion and restored your health to full.")
		} else {
			fmt.Println("You can't use that item.")
		}
	case "attack":
		room := rooms[player.CurrentRoom]
		if len(room.Enemies) == 0 {
			fmt.Println("There's nothing to attack here.")
			return
		}
		enemy := &room.Enemies[0]
		damage := player.AttackPower - enemy.Defense
		if damage < 0 {
			damage = 0
		}
		enemy.Health -= damage
		fmt.Printf("You attack the %s for %d damage.\n", enemy.Name, damage)
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			player.Gold += enemy.LootGold
			for _, loot := range enemy.LootItems {
				player.Inventory = append(player.Inventory, loot)
			}
			room.Enemies = room.Enemies[1:]
			rooms[player.CurrentRoom] = room
			return
		}
		enemyDamage := enemy.AttackPower - player.Defense
		if enemyDamage < 0 {
			enemyDamage = 0
		}
		player.Health -= enemyDamage
		fmt.Printf("The %s attacks you for %d damage.\n", enemy.Name, enemyDamage)
	case "look":
		room := rooms[player.CurrentRoom]
		fmt.Printf("You look around. %s Exits: %v. Items: %v. Enemies: %v\n", room.Description, room.Exits, room.Items, room.Enemies)
	case "inventory":
		fmt.Printf("Inventory: %v\n", player.Inventory)
	case "stats":
		fmt.Printf("Health: %d/%d, Attack: %d, Defense: %d, Gold: %d\n", player.Health, player.MaxHealth, player.AttackPower, player.Defense, player.Gold)
	case "help":
		fmt.Println("Commands: go [direction], get [item], use [item], attack, look, inventory, stats, quit, help")
	default:
		fmt.Println("Unknown command. Type 'help' for a list of commands.")
	}
}

func hasItem(item string) bool {
	for _, i := range player.Inventory {
		if i == item {
			return true
		}
	}
	return false
}

func removeItem(item string) {
	for i, it := range player.Inventory {
		if it == item {
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			return
		}
	}
}
