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

type Enemy struct {
	Name string
	Health int
	Attack int
	Defense int
}

func (p *Player) AttackEnemy(e *Enemy) {
	damage := p.Attack - e.Defense
	if damage < 0 {
		damage = 0
	}
	e.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", p.Name, e.Name, damage)
}

func (e *Enemy) AttackPlayer(p *Player) {
	damage := e.Attack - p.Defense
	if damage < 0 {
		damage = 0
	}
	p.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", e.Name, p.Name, damage)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	player := Player{
		Name: "Hero",
		Health: 100,
		Attack: 20,
		Defense: 10,
	}

	enemies := []Enemy{
		{Name: "Goblin", Health: 30, Attack: 15, Defense: 5},
		{Name: "Orc", Health: 50, Attack: 25, Defense: 15},
		{Name: "Dragon", Health: 100, Attack: 40, Defense: 30},
	}

	fmt.Println("Welcome to the Complex Game!")
	fmt.Printf("You are %s, with %d health, %d attack, and %d defense.\n", player.Name, player.Health, player.Attack, player.Defense)
	fmt.Println("You will face 3 enemies: Goblin, Orc, and Dragon.")
	fmt.Println("Defeat them all to win the game!")

	for _, enemy := range enemies {
		fmt.Printf("\nA wild %s appears!\n", enemy.Name)
		for player.Health > 0 && enemy.Health > 0 {
			player.AttackEnemy(&enemy)
			if enemy.Health <= 0 {
				fmt.Printf("%s has been defeated!\n", enemy.Name)
				break
			}
			enemy.AttackPlayer(&player)
			if player.Health <= 0 {
				fmt.Println("You have been defeated! Game over.")
				return
			}
		}
	}

	fmt.Println("\nCongratulations! You have defeated all enemies and won the game!")
}
