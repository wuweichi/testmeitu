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
}

type Enemy struct {
	Name     string
	Health   int
	Strength int
}

func (p *Player) Attack(e *Enemy) {
	damage := rand.Intn(p.Strength) + 1
	e.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", p.Name, e.Name, damage)
}

func (e *Enemy) Attack(p *Player) {
	damage := rand.Intn(e.Strength) + 1
	p.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", e.Name, p.Name, damage)
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create player and enemy
	player := Player{Name: "Hero", Health: 100, Strength: 20}
	enemy := Enemy{Name: "Dragon", Health: 150, Strength: 15}

	fmt.Println("Game Start!")
	fmt.Printf("Player: %s, Health: %d, Strength: %d\n", player.Name, player.Health, player.Strength)
	fmt.Printf("Enemy: %s, Health: %d, Strength: %d\n", enemy.Name, enemy.Health, enemy.Strength)

	// Game loop
	for player.Health > 0 && enemy.Health > 0 {
		fmt.Println("\n--- Turn Start ---")
		player.Attack(&enemy)
		if enemy.Health <= 0 {
			fmt.Printf("%s has been defeated!\n", enemy.Name)
			break
		}
		enemy.Attack(&player)
		if player.Health <= 0 {
			fmt.Printf("%s has been defeated!\n", player.Name)
			break
		}
		fmt.Printf("Current Health - Player: %d, Enemy: %d\n", player.Health, enemy.Health)
	}

	// Game over message
	if player.Health > 0 {
		fmt.Println("\nYou win! Congratulations!")
	} else {
		fmt.Println("\nGame over. You lost!")
	}
}
