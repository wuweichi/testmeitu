package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"strings"
	"bufio"
	"os"
)

type Player struct {
	Name     string
	Health   int
	Strength int
	Level    int
	XP       int
}

type Enemy struct {
	Name     string
	Health   int
	Strength int
}

type Item struct {
	Name        string
	Description string
	Value       int
}

func (p *Player) Attack(e *Enemy) {
	damage := p.Strength + rand.Intn(10)
	e.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", p.Name, e.Name, damage)
}

func (e *Enemy) Attack(p *Player) {
	damage := e.Strength + rand.Intn(5)
	p.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", e.Name, p.Name, damage)
}

func (p *Player) Heal() {
	healAmount := 20 + rand.Intn(10)
	p.Health += healAmount
	fmt.Printf("%s heals for %d health. Current health: %d\n", p.Name, healAmount, p.Health)
}

func (p *Player) LevelUp() {
	if p.XP >= p.Level*100 {
		p.Level++
		p.Strength += 5
		p.Health += 20
		p.XP = 0
		fmt.Printf("%s leveled up to level %d! Strength: %d, Health: %d\n", p.Name, p.Level, p.Strength, p.Health)
	}
}

func generateEnemy(level int) Enemy {
	enemyNames := []string{"Goblin", "Orc", "Dragon", "Skeleton", "Zombie"}
	name := enemyNames[rand.Intn(len(enemyNames))]
	health := 30 + (level * 10)
	strength := 5 + (level * 2)
	return Enemy{Name: name, Health: health, Strength: strength}
}

func generateItem() Item {
	itemNames := []string{"Health Potion", "Strength Elixir", "Magic Scroll", "Golden Coin"}
	descriptions := []string{"Restores health", "Increases strength", "Casts a spell", "Valuable currency"}
	index := rand.Intn(len(itemNames))
	value := 10 + rand.Intn(50)
	return Item{Name: itemNames[index], Description: descriptions[index], Value: value}
}

func battle(p *Player, e *Enemy) bool {
	fmt.Printf("A wild %s appears!\n", e.Name)
	for p.Health > 0 && e.Health > 0 {
		fmt.Println("\n1. Attack")
		fmt.Println("2. Heal")
		fmt.Print("Choose an action: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			p.Attack(e)
		case "2":
			p.Heal()
		default:
			fmt.Println("Invalid choice. Try again.")
			continue
		}
		if e.Health > 0 {
			e.Attack(p)
		}
	}
	if p.Health <= 0 {
		fmt.Printf("%s has been defeated! Game Over.\n", p.Name)
		return false
	} else {
		xpGain := 50 + (e.Strength * 5)
		p.XP += xpGain
		fmt.Printf("%s defeated the %s and gained %d XP!\n", p.Name, e.Name, xpGain)
		p.LevelUp()
		return true
	}
}

func explore(p *Player) {
	fmt.Println("\nYou are exploring...")
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Println("You found nothing.")
	case 1:
		item := generateItem()
		fmt.Printf("You found a %s: %s (Value: %d)\n", item.Name, item.Description, item.Value)
	case 2:
		enemy := generateEnemy(p.Level)
		if !battle(p, &enemy) {
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your character's name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	player := Player{Name: name, Health: 100, Strength: 10, Level: 1, XP: 0}
	fmt.Printf("Welcome, %s! Your journey begins.\n", player.Name)
	for {
		fmt.Println("\n1. Explore")
		fmt.Println("2. Check Status")
		fmt.Println("3. Quit")
		fmt.Print("Choose an option: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			explore(&player)
			if player.Health <= 0 {
				return
			}
		case "2":
			fmt.Printf("Name: %s, Health: %d, Strength: %d, Level: %d, XP: %d\n", player.Name, player.Health, player.Strength, player.Level, player.XP)
		case "3":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}
