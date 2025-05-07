package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name string
	Health int
	AttackPower int
}

type Monster struct {
	Name string
	Health int
	AttackPower int
}

func (p *Player) Attack(m *Monster) {
	damage := rand.Intn(p.AttackPower)
	m.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", p.Name, m.Name, damage)
}

func (m *Monster) Attack(p *Player) {
	damage := rand.Intn(m.AttackPower)
	p.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", m.Name, p.Name, damage)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	player := Player{Name: "Hero", Health: 100, AttackPower: 20}
	monster := Monster{Name: "Dragon", Health: 150, AttackPower: 25}

	fmt.Println("A wild", monster.Name, "appears!")

	for player.Health > 0 && monster.Health > 0 {
		player.Attack(&monster)
		if monster.Health <= 0 {
			fmt.Println(monster.Name, "has been defeated!")
			break
		}
		monster.Attack(&player)
		if player.Health <= 0 {
			fmt.Println(player.Name, "has been defeated!")
			break
		}
		fmt.Printf("%s's Health: %d, %s's Health: %d\n", player.Name, player.Health, monster.Name, monster.Health)
		time.Sleep(1 * time.Second)
	}

	if player.Health > 0 {
		fmt.Println("Congratulations! You've won!")
	} else {
		fmt.Println("Game Over!")
	}
}
