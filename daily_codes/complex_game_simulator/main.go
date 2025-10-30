package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name     string
	Health   int
	Mana     int
	Strength int
	Agility  int
	Intellect int
}

type Enemy struct {
	Name     string
	Health   int
	Damage   int
	Defense  int
}

type Item struct {
	Name        string
	Description string
	Value       int
}

type GameState struct {
	Player      Player
	Enemies     []Enemy
	Inventory   []Item
	Gold        int
	Level       int
	Experience  int
}

func (p *Player) Attack(e *Enemy) int {
	damage := p.Strength + rand.Intn(10)
	e.Health -= damage
	return damage
}

func (p *Player) CastSpell(e *Enemy) int {
	if p.Mana < 10 {
		return 0
	}
	p.Mana -= 10
	damage := p.Intellect + rand.Intn(15)
	e.Health -= damage
	return damage
}

func (e *Enemy) Attack(p *Player) int {
	damage := e.Damage - p.Agility/2
	if damage < 0 {
		damage = 0
	}
	p.Health -= damage
	return damage
}

func (g *GameState) AddItem(item Item) {
	g.Inventory = append(g.Inventory, item)
}

func (g *GameState) RemoveItem(index int) {
	if index < 0 || index >= len(g.Inventory) {
		return
	}
	g.Inventory = append(g.Inventory[:index], g.Inventory[index+1:]...)
}

func (g *GameState) GainExperience(exp int) {
	g.Experience += exp
	if g.Experience >= g.Level*100 {
		g.LevelUp()
	}
}

func (g *GameState) LevelUp() {
	g.Level++
	g.Player.Health += 20
	g.Player.Mana += 10
	g.Player.Strength += 5
	g.Player.Agility += 3
	g.Player.Intellect += 4
	fmt.Printf("Level up! You are now level %d\n", g.Level)
}

func (g *GameState) DisplayStatus() {
	fmt.Printf("Player: %s\n", g.Player.Name)
	fmt.Printf("Health: %d, Mana: %d\n", g.Player.Health, g.Player.Mana)
	fmt.Printf("Strength: %d, Agility: %d, Intellect: %d\n", g.Player.Strength, g.Player.Agility, g.Player.Intellect)
	fmt.Printf("Level: %d, Experience: %d\n", g.Level, g.Experience)
	fmt.Printf("Gold: %d\n", g.Gold)
}

func (g *GameState) DisplayInventory() {
	if len(g.Inventory) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}
	fmt.Println("Inventory:")
	for i, item := range g.Inventory {
		fmt.Printf("%d. %s - %s (Value: %d)\n", i+1, item.Name, item.Description, item.Value)
	}
}

func (g *GameState) Fight(enemyIndex int) {
	if enemyIndex < 0 || enemyIndex >= len(g.Enemies) {
		fmt.Println("Invalid enemy index.")
		return
	}
	enemy := &g.Enemies[enemyIndex]
	fmt.Printf("A wild %s appears!\n", enemy.Name)
	for g.Player.Health > 0 && enemy.Health > 0 {
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Attack")
		fmt.Println("2. Cast Spell")
		fmt.Println("3. Run away")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			damage := g.Player.Attack(enemy)
			fmt.Printf("You attack the %s for %d damage.\n", enemy.Name, damage)
		case 2:
			damage := g.Player.CastSpell(enemy)
			if damage == 0 {
				fmt.Println("Not enough mana!")
			} else {
				fmt.Printf("You cast a spell on the %s for %d damage.\n", enemy.Name, damage)
			}
		case 3:
			fmt.Println("You run away!")
			return
		default:
			fmt.Println("Invalid choice.")
			continue
		}
		if enemy.Health <= 0 {
			fmt.Printf("You defeated the %s!\n", enemy.Name)
			g.GainExperience(50)
			g.Gold += 10
			g.Enemies = append(g.Enemies[:enemyIndex], g.Enemies[enemyIndex+1:]...)
			return
		}
		enemyDamage := enemy.Attack(&g.Player)
		fmt.Printf("The %s attacks you for %d damage.\n", enemy.Name, enemyDamage)
		if g.Player.Health <= 0 {
			fmt.Println("You have been defeated!")
			return
		}
	}
}

func (g *GameState) Explore() {
	fmt.Println("You explore the area...")
	rand.Seed(time.Now().UnixNano())
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Println("You found a treasure chest!")
		g.Gold += 50
		item := Item{Name: "Health Potion", Description: "Restores 50 health", Value: 25}
		g.AddItem(item)
		fmt.Println("You gained 50 gold and a Health Potion.")
	case 1:
		fmt.Println("You encountered an enemy!")
		enemy := Enemy{Name: "Goblin", Health: 30, Damage: 5, Defense: 2}
		g.Enemies = append(g.Enemies, enemy)
		g.Fight(len(g.Enemies) - 1)
	case 2:
		fmt.Println("You found nothing of interest.")
	}
}

func (g *GameState) Shop() {
	fmt.Println("Welcome to the shop!")
	fmt.Println("1. Health Potion - 50 gold")
	fmt.Println("2. Mana Potion - 50 gold")
	fmt.Println("3. Leave")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		if g.Gold >= 50 {
			g.Gold -= 50
			item := Item{Name: "Health Potion", Description: "Restores 50 health", Value: 25}
			g.AddItem(item)
			fmt.Println("You bought a Health Potion.")
		} else {
			fmt.Println("Not enough gold!")
		}
	case 2:
		if g.Gold >= 50 {
			g.Gold -= 50
			item := Item{Name: "Mana Potion", Description: "Restores 30 mana", Value: 25}
			g.AddItem(item)
			fmt.Println("You bought a Mana Potion.")
		} else {
			fmt.Println("Not enough gold!")
		}
	case 3:
		fmt.Println("You leave the shop.")
	default:
		fmt.Println("Invalid choice.")
	}
}

func (g *GameState) UseItem(index int) {
	if index < 0 || index >= len(g.Inventory) {
		fmt.Println("Invalid item index.")
		return
	}
	item := g.Inventory[index]
	switch item.Name {
	case "Health Potion":
		g.Player.Health += 50
		fmt.Println("You used a Health Potion and restored 50 health.")
	case "Mana Potion":
		g.Player.Mana += 30
		fmt.Println("You used a Mana Potion and restored 30 mana.")
	default:
		fmt.Println("This item cannot be used.")
		return
	}
	g.RemoveItem(index)
}

func main() {
	fmt.Println("Welcome to the Complex Game Simulator!")
	game := GameState{
		Player: Player{
			Name:      "Hero",
			Health:    100,
			Mana:      50,
			Strength:  10,
			Agility:   8,
			Intellect: 12,
		},
		Enemies: []Enemy{
			{Name: "Slime", Health: 20, Damage: 3, Defense: 1},
			{Name: "Skeleton", Health: 40, Damage: 7, Defense: 3},
		},
		Inventory: []Item{
			{Name: "Health Potion", Description: "Restores 50 health", Value: 25},
		},
		Gold:       100,
		Level:      1,
		Experience: 0,
	}
	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Display Status")
		fmt.Println("2. Display Inventory")
		fmt.Println("3. Explore")
		fmt.Println("4. Shop")
		fmt.Println("5. Use Item")
		fmt.Println("6. Fight Enemy")
		fmt.Println("7. Exit")
		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			game.DisplayStatus()
		case 2:
			game.DisplayInventory()
		case 3:
			game.Explore()
		case 4:
			game.Shop()
		case 5:
			fmt.Print("Enter item index to use: ")
			var index int
			fmt.Scan(&index)
			game.UseItem(index - 1)
		case 6:
			if len(game.Enemies) == 0 {
				fmt.Println("No enemies to fight.")
				continue
			}
			fmt.Println("Choose an enemy to fight:")
			for i, enemy := range game.Enemies {
				fmt.Printf("%d. %s (Health: %d)\n", i+1, enemy.Name, enemy.Health)
			}
			var enemyIndex int
			fmt.Scan(&enemyIndex)
			game.Fight(enemyIndex - 1)
		case 7:
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid choice.")
		}
	}
}
