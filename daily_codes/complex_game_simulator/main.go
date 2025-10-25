package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"strings"
)

type Player struct {
	Name string
	Health int
	Mana int
	Level int
	Experience int
	Inventory []string
	Gold int
}

type Monster struct {
	Name string
	Health int
	Damage int
	Experience int
	Gold int
}

type Item struct {
	Name string
	Type string
	Value int
}

func (p *Player) Attack(m *Monster) {
	damage := rand.Intn(10) + p.Level
	m.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", p.Name, m.Name, damage)
}

func (m *Monster) Attack(p *Player) {
	damage := rand.Intn(m.Damage) + 1
	p.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", m.Name, p.Name, damage)
}

func (p *Player) Heal() {
	if p.Mana >= 5 {
		healAmount := rand.Intn(15) + 10
		p.Health += healAmount
		p.Mana -= 5
		fmt.Printf("%s heals for %d health!\n", p.Name, healAmount)
	} else {
		fmt.Println("Not enough mana to heal!")
	}
}

func (p *Player) LevelUp() {
	if p.Experience >= p.Level*100 {
		p.Level++
		p.Health += 20
		p.Mana += 10
		p.Experience = 0
		fmt.Printf("%s leveled up to level %d!\n", p.Name, p.Level)
	}
}

func (p *Player) AddItem(item string) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s added to inventory.\n", item)
}

func (p *Player) ShowStats() {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Health: %d\n", p.Health)
	fmt.Printf("Mana: %d\n", p.Mana)
	fmt.Printf("Level: %d\n", p.Level)
	fmt.Printf("Experience: %d\n", p.Experience)
	fmt.Printf("Gold: %d\n", p.Gold)
	fmt.Printf("Inventory: %v\n", p.Inventory)
}

func GenerateMonster() Monster {
	monsters := []Monster{
		{"Goblin", 30, 5, 10, 5},
		{"Orc", 50, 8, 20, 10},
		{"Dragon", 100, 15, 50, 25},
		{"Slime", 20, 3, 5, 2},
		{"Skeleton", 40, 6, 15, 8},
		{"Zombie", 35, 7, 12, 6},
		{"Witch", 45, 10, 25, 12},
		{"Giant", 80, 12, 40, 20},
		{"Vampire", 60, 9, 30, 15},
		{"Werewolf", 55, 11, 28, 14},
	}
	return monsters[rand.Intn(len(monsters))]
}

func GenerateItem() Item {
	items := []Item{
		{"Health Potion", "potion", 20},
		{"Mana Potion", "potion", 15},
		{"Sword", "weapon", 10},
		{"Shield", "armor", 8},
		{"Magic Wand", "weapon", 12},
		{"Helmet", "armor", 6},
		{"Boots", "armor", 4},
		{"Ring", "accessory", 5},
		{"Amulet", "accessory", 7},
		{"Scroll", "consumable", 3},
	}
	return items[rand.Intn(len(items))]
}

func Battle(player *Player, monster Monster) {
	fmt.Printf("A wild %s appears!\n", monster.Name)
	for player.Health > 0 && monster.Health > 0 {
		fmt.Println("\n--- Battle Menu ---")
		fmt.Println("1. Attack")
		fmt.Println("2. Heal")
		fmt.Println("3. Run")
		var choice string
		fmt.Print("Choose an action: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			player.Attack(&monster)
			if monster.Health > 0 {
				monster.Attack(player)
			}
		case "2":
			player.Heal()
			monster.Attack(player)
		case "3":
			fmt.Println("You ran away!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
		fmt.Printf("%s Health: %d\n", player.Name, player.Health)
		fmt.Printf("%s Health: %d\n", monster.Name, monster.Health)
	}
	if player.Health <= 0 {
		fmt.Println("You have been defeated!")
	} else {
		fmt.Printf("You defeated the %s!\n", monster.Name)
		player.Experience += monster.Experience
		player.Gold += monster.Gold
		fmt.Printf("Gained %d experience and %d gold.\n", monster.Experience, monster.Gold)
		player.LevelUp()
		if rand.Intn(100) < 30 {
			item := GenerateItem()
			player.AddItem(item.Name)
			fmt.Printf("Found a %s!\n", item.Name)
		}
	}
}

func Explore(player *Player) {
	events := []string{
		"You find a hidden path.",
		"You discover an ancient ruin.",
		"You stumble upon a treasure chest.",
		"You encounter a friendly traveler.",
		"You find a magical spring.",
		"You discover a secret cave.",
		"You find a abandoned camp.",
		"You encounter a mysterious old man.",
		"You find a glowing crystal.",
		"You discover a hidden waterfall.",
	}
	event := events[rand.Intn(len(events))]
	fmt.Println(event)
	if rand.Intn(100) < 50 {
		monster := GenerateMonster()
		Battle(player, monster)
	} else {
		if rand.Intn(100) < 30 {
			gold := rand.Intn(20) + 5
			player.Gold += gold
			fmt.Printf("You found %d gold!\n", gold)
		} else if rand.Intn(100) < 20 {
			item := GenerateItem()
			player.AddItem(item.Name)
			fmt.Printf("You found a %s!\n", item.Name)
		} else {
			fmt.Println("Nothing interesting happens.")
		}
	}
}

func Shop(player *Player) {
	items := []Item{
		{"Health Potion", "potion", 10},
		{"Mana Potion", "potion", 8},
		{"Sword", "weapon", 50},
		{"Shield", "armor", 40},
		{"Magic Wand", "weapon", 60},
		{"Helmet", "armor", 30},
		{"Boots", "armor", 20},
		{"Ring", "accessory", 25},
		{"Amulet", "accessory", 35},
		{"Scroll", "consumable", 15},
	}
	fmt.Println("\n--- Shop ---")
	for i, item := range items {
		fmt.Printf("%d. %s - %d gold\n", i+1, item.Name, item.Value)
	}
	fmt.Println("0. Exit")
	var choice string
	fmt.Print("Choose an item to buy: ")
	fmt.Scanln(&choice)
	index, err := strconv.Atoi(choice)
	if err != nil || index < 0 || index > len(items) {
		fmt.Println("Invalid choice!")
		return
	}
	if index == 0 {
		return
	}
	item := items[index-1]
	if player.Gold >= item.Value {
		player.Gold -= item.Value
		player.AddItem(item.Name)
		fmt.Printf("You bought a %s!\n", item.Name)
	} else {
		fmt.Println("Not enough gold!")
	}
}

func Rest(player *Player) {
	player.Health = 100 + player.Level*10
	player.Mana = 50 + player.Level*5
	fmt.Println("You rest and recover your health and mana.")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("Welcome to the Complex Game Simulator!")
	var playerName string
	fmt.Print("Enter your character name: ")
	fmt.Scanln(&playerName)
	player := &Player{
		Name: playerName,
		Health: 100,
		Mana: 50,
		Level: 1,
		Experience: 0,
		Inventory: []string{},
		Gold: 10,
	}
	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Explore")
		fmt.Println("2. Shop")
		fmt.Println("3. Rest")
		fmt.Println("4. Show Stats")
		fmt.Println("5. Quit")
		var choice string
		fmt.Print("Choose an option: ")
		fmt.Scanln(&choice)
		switch choice {
		case "1":
			Explore(player)
		case "2":
			Shop(player)
		case "3":
			Rest(player)
		case "4":
			player.ShowStats()
		case "5":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid choice!")
		}
		if player.Health <= 0 {
			fmt.Println("Game Over!")
			break
		}
	}
}
