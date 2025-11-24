package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name     string
	Health   int
	Strength int
	Agility  int
	Magic    int
}

type Monster struct {
	Name     string
	Health   int
	Strength int
	Agility  int
	Magic    int
}

type Item struct {
	Name        string
	Description string
	Value       int
}

type GameState struct {
	Player    Player
	Monsters  []Monster
	Inventory []Item
	Level     int
	Score     int
}

func (p *Player) Attack(m *Monster) int {
	damage := p.Strength + rand.Intn(10)
	m.Health -= damage
	return damage
}

func (m *Monster) Attack(p *Player) int {
	damage := m.Strength + rand.Intn(8)
	p.Health -= damage
	return damage
}

func (p *Player) Heal() int {
	healAmount := p.Magic + rand.Intn(15)
	p.Health += healAmount
	return healAmount
}

func (p *Player) Dodge() bool {
	return rand.Intn(100) < p.Agility
}

func (m *Monster) Dodge() bool {
	return rand.Intn(100) < m.Agility
}

func generateMonster(level int) Monster {
	baseHealth := 50 + level*10
	baseStrength := 10 + level*2
	baseAgility := 5 + level
	baseMagic := 3 + level
	
	names := []string{"Goblin", "Orc", "Troll", "Dragon", "Vampire", "Werewolf"}
	
	return Monster{
		Name:     names[rand.Intn(len(names))],
		Health:   baseHealth + rand.Intn(20),
		Strength: baseStrength + rand.Intn(5),
		Agility:  baseAgility + rand.Intn(3),
		Magic:    baseMagic + rand.Intn(2),
	}
}

func generateItem(level int) Item {
	items := []Item{
		{Name: "Health Potion", Description: "Restores health", Value: 20},
		{Name: "Strength Potion", Description: "Increases strength", Value: 15},
		{Name: "Agility Potion", Description: "Increases agility", Value: 15},
		{Name: "Magic Potion", Description: "Increases magic", Value: 15},
		{Name: "Sword", Description: "A sharp blade", Value: 30},
		{Name: "Shield", Description: "Protects from attacks", Value: 25},
		{Name: "Armor", Description: "Increases defense", Value: 40},
	}
	
	item := items[rand.Intn(len(items))]
	item.Value += level * 5
	return item
}

func battle(player *Player, monster *Monster) bool {
	fmt.Printf("A wild %s appears!\n", monster.Name)
	
	for player.Health > 0 && monster.Health > 0 {
		fmt.Printf("\nYour Health: %d | %s's Health: %d\n", player.Health, monster.Name, monster.Health)
		fmt.Println("Choose your action:")
		fmt.Println("1. Attack")
		fmt.Println("2. Heal")
		fmt.Println("3. Dodge")
		
		var choice int
		fmt.Scan(&choice)
		
		switch choice {
		case 1:
			if !monster.Dodge() {
				damage := player.Attack(monster)
				fmt.Printf("You attack the %s for %d damage!\n", monster.Name, damage)
			} else {
				fmt.Printf("The %s dodged your attack!\n", monster.Name)
			}
		case 2:
			healAmount := player.Heal()
			fmt.Printf("You heal yourself for %d health!\n", healAmount)
		case 3:
			if player.Dodge() {
				fmt.Println("You successfully dodge!")
				continue
			} else {
				fmt.Println("Dodge failed!")
			}
		default:
			fmt.Println("Invalid choice!")
			continue
		}
		
		if monster.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", monster.Name)
			return true
		}
		
		if !player.Dodge() {
			damage := monster.Attack(player)
			fmt.Printf("The %s attacks you for %d damage!\n", monster.Name, damage)
		} else {
			fmt.Printf("You dodged the %s's attack!\n", monster.Name)
		}
		
		if player.Health <= 0 {
			fmt.Println("You have been defeated!")
			return false
		}
	}
	
	return player.Health > 0
}

func explore(player *Player, gameState *GameState) {
	fmt.Println("You are exploring...")
	
	event := rand.Intn(3)
	switch event {
	case 0:
		monster := generateMonster(gameState.Level)
		if battle(player, &monster) {
			gameState.Score += 100
			item := generateItem(gameState.Level)
			gameState.Inventory = append(gameState.Inventory, item)
			fmt.Printf("You found a %s!\n", item.Name)
		}
	case 1:
		item := generateItem(gameState.Level)
		gameState.Inventory = append(gameState.Inventory, item)
		fmt.Printf("You found a %s!\n", item.Name)
	case 2:
		fmt.Println("You found nothing interesting.")
	}
}

func showInventory(gameState *GameState) {
	fmt.Println("\n--- Inventory ---")
	if len(gameState.Inventory) == 0 {
		fmt.Println("Your inventory is empty.")
	} else {
		for i, item := range gameState.Inventory {
			fmt.Printf("%d. %s - %s (Value: %d)\n", i+1, item.Name, item.Description, item.Value)
		}
	}
}

func useItem(player *Player, gameState *GameState) {
	if len(gameState.Inventory) == 0 {
		fmt.Println("Your inventory is empty!")
		return
	}
	
	showInventory(gameState)
	fmt.Print("Enter the number of the item to use: ")
	var choice int
	fmt.Scan(&choice)
	
	if choice < 1 || choice > len(gameState.Inventory) {
		fmt.Println("Invalid choice!")
		return
	}
	
	item := gameState.Inventory[choice-1]
	switch item.Name {
	case "Health Potion":
		player.Health += item.Value
		fmt.Printf("You used %s and restored %d health!\n", item.Name, item.Value)
	case "Strength Potion":
		player.Strength += 5
		fmt.Printf("You used %s and increased your strength by 5!\n", item.Name)
	case "Agility Potion":
		player.Agility += 5
		fmt.Printf("You used %s and increased your agility by 5!\n", item.Name)
	case "Magic Potion":
		player.Magic += 5
		fmt.Printf("You used %s and increased your magic by 5!\n", item.Name)
	default:
		fmt.Printf("You can't use %s right now.\n", item.Name)
		return
	}
	
	gameState.Inventory = append(gameState.Inventory[:choice-1], gameState.Inventory[choice:]...)
}

func showStatus(player *Player, gameState *GameState) {
	fmt.Println("\n--- Player Status ---")
	fmt.Printf("Name: %s\n", player.Name)
	fmt.Printf("Health: %d\n", player.Health)
	fmt.Printf("Strength: %d\n", player.Strength)
	fmt.Printf("Agility: %d\n", player.Agility)
	fmt.Printf("Magic: %d\n", player.Magic)
	fmt.Printf("Level: %d\n", gameState.Level)
	fmt.Printf("Score: %d\n", gameState.Score)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	
	fmt.Println("Welcome to the Complex Game Simulator!")
	fmt.Print("Enter your character name: ")
	
	var playerName string
	fmt.Scan(&playerName)
	
	player := Player{
		Name:     playerName,
		Health:   100,
		Strength: 15,
		Agility:  10,
		Magic:    8,
	}
	
	gameState := GameState{
		Player:   player,
		Monsters: []Monster{},
		Level:    1,
		Score:    0,
	}
	
	for {
		fmt.Println("\n--- Main Menu ---")
		fmt.Println("1. Explore")
		fmt.Println("2. Show Inventory")
		fmt.Println("3. Use Item")
		fmt.Println("4. Show Status")
		fmt.Println("5. Quit")
		
		var choice int
		fmt.Scan(&choice)
		
		switch choice {
		case 1:
			explore(&player, &gameState)
			if player.Health <= 0 {
				fmt.Println("Game Over!")
				fmt.Printf("Final Score: %d\n", gameState.Score)
				return
			}
			
			if gameState.Score >= gameState.Level*200 {
				gameState.Level++
				player.Health += 20
				player.Strength += 2
				player.Agility += 1
				player.Magic += 1
				fmt.Printf("\nCongratulations! You reached level %d!\n", gameState.Level)
			}
		case 2:
			showInventory(&gameState)
		case 3:
			useItem(&player, &gameState)
		case 4:
			showStatus(&player, &gameState)
		case 5:
			fmt.Println("Thanks for playing!")
			fmt.Printf("Final Score: %d\n", gameState.Score)
			return
		default:
			fmt.Println("Invalid choice!")
		}
	}
}
