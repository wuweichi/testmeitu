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
	Strength  int
	Agility   int
	Intellect int
	Gold      int
	Inventory []string
	Level     int
	Exp       int
}

type Monster struct {
	Name      string
	Health    int
	Strength  int
	Agility   int
	RewardExp int
	RewardGold int
}

type Item struct {
	Name        string
	Description string
	Effect      string
	Value       int
}

type Location struct {
	Name        string
	Description string
	Monsters    []Monster
	Items       []Item
	Connections []string
}

var player Character
var locations map[string]Location
var currentLocation string
var rng *rand.Rand

func initGame() {
	rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	player = Character{
		Name:      "Hero",
		Health:    100,
		MaxHealth: 100,
		Strength:  10,
		Agility:   10,
		Intellect: 10,
		Gold:      50,
		Inventory: []string{"Health Potion", "Rusty Sword"},
		Level:     1,
		Exp:       0,
	}
	locations = make(map[string]Location)
	locations["forest"] = Location{
		Name:        "Enchanted Forest",
		Description: "A mystical forest with glowing mushrooms and ancient trees.",
		Monsters: []Monster{
			{Name: "Goblin", Health: 30, Strength: 5, Agility: 8, RewardExp: 20, RewardGold: 10},
			{Name: "Wolf", Health: 25, Strength: 7, Agility: 10, RewardExp: 15, RewardGold: 5},
		},
		Items: []Item{
			{Name: "Herbs", Description: "Medicinal herbs for healing.", Effect: "heal", Value: 20},
			{Name: "Silver Coin", Description: "A shiny silver coin.", Effect: "gold", Value: 5},
		},
		Connections: []string{"village", "cave"},
	}
	locations["village"] = Location{
		Name:        "Quaint Village",
		Description: "A peaceful village with friendly villagers and a market.",
		Monsters:    []Monster{},
		Items: []Item{
			{Name: "Health Potion", Description: "Restores health.", Effect: "heal", Value: 50},
			{Name: "Strength Elixir", Description: "Boosts strength temporarily.", Effect: "buff", Value: 5},
		},
		Connections: []string{"forest", "mountains"},
	}
	locations["cave"] = Location{
		Name:        "Dark Cave",
		Description: "A spooky cave with echoing drips and hidden dangers.",
		Monsters: []Monster{
			{Name: "Troll", Health: 80, Strength: 15, Agility: 3, RewardExp: 50, RewardGold: 30},
			{Name: "Bat Swarm", Health: 20, Strength: 3, Agility: 15, RewardExp: 10, RewardGold: 2},
		},
		Items: []Item{
			{Name: "Treasure Chest", Description: "A chest filled with gold.", Effect: "gold", Value: 100},
			{Name: "Mystic Orb", Description: "A glowing orb with unknown powers.", Effect: "special", Value: 1},
		},
		Connections: []string{"forest"},
	}
	locations["mountains"] = Location{
		Name:        "Snowy Mountains",
		Description: "Icy peaks with treacherous paths and rare creatures.",
		Monsters: []Monster{
			{Name: "Yeti", Health: 120, Strength: 20, Agility: 5, RewardExp: 80, RewardGold: 50},
			{Name: "Ice Elemental", Health: 60, Strength: 12, Agility: 10, RewardExp: 40, RewardGold: 25},
		},
		Items: []Item{
			{Name: "Ice Crystal", Description: "A crystal that radiates cold.", Effect: "special", Value: 1},
			{Name: "Warm Cloak", Description: "Keeps you warm in cold climates.", Effect: "armor", Value: 10},
		},
		Connections: []string{"village"},
	}
	currentLocation = "forest"
}

func displayStatus() {
	fmt.Println("\n=== Player Status ===")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Health: %d/%d\n", player.Health, player.MaxHealth)
	fmt.Printf("Strength: %d, Agility: %d, Intellect: %d\n", player.Strength, player.Agility, player.Intellect)
	fmt.Printf("Gold: %d\n", player.Gold)
	fmt.Printf("Level: %d, Exp: %d/%d\n", player.Level, player.Exp, player.Level*100)
	fmt.Printf("Inventory: %v\n", player.Inventory)
	fmt.Println("====================")
}

func displayLocation() {
	loc := locations[currentLocation]
	fmt.Printf("\nYou are at: %s\n", loc.Name)
	fmt.Printf("Description: %s\n", loc.Description)
	if len(loc.Monsters) > 0 {
		fmt.Print("Monsters here: ")
		for i, m := range loc.Monsters {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s (HP: %d)", m.Name, m.Health)
		}
		fmt.Println()
	} else {
		fmt.Println("No monsters here.")
	}
	if len(loc.Items) > 0 {
		fmt.Print("Items here: ")
		for i, item := range loc.Items {
			if i > 0 {
				fmt.Print(", ")
			}
			fmt.Printf("%s", item.Name)
		}
		fmt.Println()
	} else {
		fmt.Println("No items here.")
	}
	fmt.Printf("Connections: %v\n", loc.Connections)
}

func move(destination string) {
	loc := locations[currentLocation]
	for _, conn := range loc.Connections {
		if conn == destination {
			currentLocation = destination
			fmt.Printf("You move to %s.\n", locations[destination].Name)
			return
		}
	}
	fmt.Println("You cannot move there from here.")
}

func fight(monsterName string) {
	loc := locations[currentLocation]
	var monster *Monster
	for i := range loc.Monsters {
		if loc.Monsters[i].Name == monsterName {
			monster = &loc.Monsters[i]
			break
		}
	}
	if monster == nil {
		fmt.Println("Monster not found here.")
		return
	}
	fmt.Printf("You engage in battle with %s!\n", monster.Name)
	for player.Health > 0 && monster.Health > 0 {
		playerDamage := player.Strength + rng.Intn(5)
		monster.Health -= playerDamage
		fmt.Printf("You hit %s for %d damage. %s's HP: %d\n", monster.Name, playerDamage, monster.Name, monster.Health)
		if monster.Health <= 0 {
			fmt.Printf("You defeated %s!\n", monster.Name)
			player.Exp += monster.RewardExp
			player.Gold += monster.RewardGold
			fmt.Printf("Gained %d Exp and %d Gold.\n", monster.RewardExp, monster.RewardGold)
			removeMonster(currentLocation, monsterName)
			checkLevelUp()
			return
		}
		monsterDamage := monster.Strength + rng.Intn(3)
		player.Health -= monsterDamage
		fmt.Printf("%s hits you for %d damage. Your HP: %d\n", monster.Name, monsterDamage, player.Health)
		if player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			os.Exit(0)
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func removeMonster(locName, monsterName string) {
	loc := locations[locName]
	newMonsters := []Monster{}
	for _, m := range loc.Monsters {
		if m.Name != monsterName {
			newMonsters = append(newMonsters, m)
		}
	}
	loc.Monsters = newMonsters
	locations[locName] = loc
}

func checkLevelUp() {
	requiredExp := player.Level * 100
	if player.Exp >= requiredExp {
		player.Level++
		player.Exp -= requiredExp
		player.MaxHealth += 20
		player.Health = player.MaxHealth
		player.Strength += 2
		player.Agility += 2
		player.Intellect += 2
		fmt.Printf("Congratulations! You leveled up to Level %d!\n", player.Level)
	}
}

func useItem(itemName string) {
	for i, invItem := range player.Inventory {
		if invItem == itemName {
			switch itemName {
			case "Health Potion":
				healAmount := 50
				player.Health += healAmount
				if player.Health > player.MaxHealth {
					player.Health = player.MaxHealth
				}
				fmt.Printf("Used Health Potion. Healed %d HP. Current HP: %d\n", healAmount, player.Health)
			case "Strength Elixir":
				player.Strength += 5
				fmt.Printf("Used Strength Elixir. Strength increased by 5. Current Strength: %d\n", player.Strength)
			default:
				fmt.Printf("Used %s.\n", itemName)
			}
			player.Inventory = append(player.Inventory[:i], player.Inventory[i+1:]...)
			return
		}
	}
	fmt.Println("Item not found in inventory.")
}

func pickUpItem(itemName string) {
	loc := locations[currentLocation]
	for i, item := range loc.Items {
		if item.Name == itemName {
			player.Inventory = append(player.Inventory, item.Name)
			fmt.Printf("You picked up %s.\n", item.Name)
			loc.Items = append(loc.Items[:i], loc.Items[i+1:]...)
			locations[currentLocation] = loc
			return
		}
	}
	fmt.Println("Item not found here.")
}

func saveGame() {
	data, err := json.MarshalIndent(struct {
		Player          Character
		Locations       map[string]Location
		CurrentLocation string
	}{player, locations, currentLocation}, "", "  ")
	if err != nil {
		fmt.Println("Error saving game:", err)
		return
	}
	err = os.WriteFile("savegame.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing save file:", err)
		return
	}
	fmt.Println("Game saved successfully.")
}

func loadGame() {
	data, err := os.ReadFile("savegame.json")
	if err != nil {
		fmt.Println("No save file found.")
		return
	}
	var save struct {
		Player          Character
		Locations       map[string]Location
		CurrentLocation string
	}
	err = json.Unmarshal(data, &save)
	if err != nil {
		fmt.Println("Error loading save:", err)
		return
	}
	player = save.Player
	locations = save.Locations
	currentLocation = save.CurrentLocation
	fmt.Println("Game loaded successfully.")
}

func main() {
	initGame()
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Adventure Game Simulator!")
	fmt.Println("Type 'help' for commands.")
	for {
		displayStatus()
		displayLocation()
		fmt.Print("\nEnter command: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		parts := strings.Fields(input)
		if len(parts) == 0 {
			continue
		}
		command := strings.ToLower(parts[0])
		switch command {
		case "help":
			fmt.Println("Commands:")
			fmt.Println("  status - Show player status")
			fmt.Println("  look - Show current location")
			fmt.Println("  move <destination> - Move to connected location")
			fmt.Println("  fight <monster> - Fight a monster")
			fmt.Println("  use <item> - Use an item from inventory")
			fmt.Println("  pickup <item> - Pick up an item from location")
			fmt.Println("  save - Save game")
			fmt.Println("  load - Load game")
			fmt.Println("  quit - Exit game")
		case "status":
			displayStatus()
		case "look":
			displayLocation()
		case "move":
			if len(parts) < 2 {
				fmt.Println("Usage: move <destination>")
			} else {
				move(parts[1])
			}
		case "fight":
			if len(parts) < 2 {
				fmt.Println("Usage: fight <monster>")
			} else {
				fight(parts[1])
			}
		case "use":
			if len(parts) < 2 {
				fmt.Println("Usage: use <item>")
			} else {
				useItem(parts[1])
			}
		case "pickup":
			if len(parts) < 2 {
				fmt.Println("Usage: pickup <item>")
			} else {
				pickUpItem(parts[1])
			}
		case "save":
			saveGame()
		case "load":
			loadGame()
		case "quit":
			fmt.Println("Thanks for playing!")
			os.Exit(0)
		default:
			fmt.Println("Unknown command. Type 'help' for list.")
		}
	}
}
