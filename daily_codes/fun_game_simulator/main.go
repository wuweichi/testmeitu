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
	Name     string
	Health   int
	Attack   int
	Defense  int
	Level    int
	Experience int
}

type Monster struct {
	Name     string
	Health   int
	Attack   int
	Defense  int
	Experience int
}

type Item struct {
	Name        string
	Description string
	Effect      func(*Player)
}

func (p *Player) LevelUp() {
	if p.Experience >= p.Level*100 {
		p.Level++
		p.Health += 20
		p.Attack += 5
		p.Defense += 3
		p.Experience = 0
		fmt.Printf("Level up! You are now level %d.\n", p.Level)
	}
}

func (p *Player) AttackMonster(m *Monster) {
	damage := p.Attack - m.Defense
	if damage < 0 {
		damage = 0
	}
	m.Health -= damage
	fmt.Printf("You attack %s for %d damage.\n", m.Name, damage)
}

func (m *Monster) AttackPlayer(p *Player) {
	damage := m.Attack - p.Defense
	if damage < 0 {
		damage = 0
	}
	p.Health -= damage
	fmt.Printf("%s attacks you for %d damage.\n", m.Name, damage)
}

func generateMonster(level int) Monster {
	monsters := []Monster{
		{Name: "Goblin", Health: 30 + level*5, Attack: 5 + level, Defense: 2 + level, Experience: 20 + level*5},
		{Name: "Orc", Health: 50 + level*8, Attack: 8 + level*2, Defense: 5 + level, Experience: 40 + level*10},
		{Name: "Dragon", Health: 100 + level*15, Attack: 15 + level*3, Defense: 10 + level*2, Experience: 100 + level*20},
	}
	return monsters[rand.Intn(len(monsters))]
}

func generateItem() Item {
	items := []Item{
		{Name: "Health Potion", Description: "Restores 20 health.", Effect: func(p *Player) { p.Health += 20; fmt.Println("Health restored by 20.") }},
		{Name: "Attack Boost", Description: "Increases attack by 5.", Effect: func(p *Player) { p.Attack += 5; fmt.Println("Attack increased by 5.") }},
		{Name: "Defense Boost", Description: "Increases defense by 3.", Effect: func(p *Player) { p.Defense += 3; fmt.Println("Defense increased by 3.") }},
	}
	return items[rand.Intn(len(items))]
}

func battle(p *Player, m Monster) {
	fmt.Printf("A wild %s appears!\n", m.Name)
	for p.Health > 0 && m.Health > 0 {
		fmt.Printf("Your health: %d, %s's health: %d\n", p.Health, m.Name, m.Health)
		fmt.Println("Choose action: 1. Attack, 2. Run")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			p.AttackMonster(&m)
			if m.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", m.Name)
				p.Experience += m.Experience
				fmt.Printf("Gained %d experience.\n", m.Experience)
				p.LevelUp()
				if rand.Float64() < 0.3 {
					item := generateItem()
					fmt.Printf("You found a %s: %s\n", item.Name, item.Description)
					item.Effect(p)
				}
				return
			}
			m.AttackPlayer(p)
			if p.Health <= 0 {
				fmt.Println("You have been defeated! Game over.")
				os.Exit(0)
			}
		case "2":
			fmt.Println("You ran away safely.")
			return
		default:
			fmt.Println("Invalid choice. Try again.")
		}
	}
}

func explore(p *Player) {
	fmt.Println("You are exploring...")
	time.Sleep(1 * time.Second)
	event := rand.Intn(3)
	switch event {
	case 0:
		monster := generateMonster(p.Level)
		battle(p, monster)
	case 1:
		item := generateItem()
		fmt.Printf("You found a %s: %s\n", item.Name, item.Description)
		item.Effect(p)
	case 2:
		fmt.Println("You found nothing interesting.")
	}
}

func showStatus(p *Player) {
	fmt.Printf("Name: %s, Health: %d, Attack: %d, Defense: %d, Level: %d, Experience: %d/%d\n",
		p.Name, p.Health, p.Attack, p.Defense, p.Level, p.Experience, p.Level*100)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	player := Player{Name: name, Health: 100, Attack: 10, Defense: 5, Level: 1, Experience: 0}
	fmt.Printf("Welcome, %s! Your adventure begins.\n", player.Name)
	for {
		fmt.Println("\nChoose action: 1. Explore, 2. Show Status, 3. Quit")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			explore(&player)
		case "2":
			showStatus(&player)
		case "3":
			fmt.Println("Thanks for playing! Goodbye.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
