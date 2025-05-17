package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name string
	Health int
	Attack int
	Defense int
}

type Monster struct {
	Name string
	Health int
	Attack int
	Defense int
}

func (p *Player) AttackMonster(m *Monster) {
	damage := p.Attack - m.Defense
	if damage < 0 {
		damage = 0
	}
	m.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", p.Name, m.Name, damage)
}

func (m *Monster) AttackPlayer(p *Player) {
	damage := m.Attack - p.Defense
	if damage < 0 {
		damage = 0
	}
	p.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", m.Name, p.Name, damage)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	player := Player{Name: "Hero", Health: 100, Attack: 20, Defense: 10}
	monster := Monster{Name: "Goblin", Health: 50, Attack: 15, Defense: 5}

	fmt.Println("A wild", monster.Name, "appears!")
	for player.Health > 0 && monster.Health > 0 {
		player.AttackMonster(&monster)
		if monster.Health <= 0 {
			fmt.Println(monster.Name, "has been defeated!")
			break
		}
		monster.AttackPlayer(&player)
		if player.Health <= 0 {
			fmt.Println(player.Name, "has been defeated!")
			break
		}
		fmt.Printf("%s: %d HP, %s: %d HP\n", player.Name, player.Health, monster.Name, monster.Health)
		time.Sleep(1 * time.Second)
	}
}
