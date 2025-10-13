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

type Player struct {
	Name     string
	Health   int
	Mana     int
	Level    int
	Exp      int
	Gold     int
	Inventory []string
}

type Enemy struct {
	Name   string
	Health int
	Damage int
	Reward int
}

type GameState struct {
	Player      Player
	Enemies     []Enemy
	CurrentRoom int
	GameOver    bool
}

func (p *Player) Attack(e *Enemy) int {
	damage := rand.Intn(10) + p.Level
	e.Health -= damage
	return damage
}

func (p *Player) Heal() {
	if p.Mana >= 5 {
		healAmount := rand.Intn(15) + 5
		p.Health += healAmount
		p.Mana -= 5
		fmt.Printf("Healed for %d health. Current health: %d\n", healAmount, p.Health)
	} else {
		fmt.Println("Not enough mana to heal!")
	}
}

func (p *Player) LevelUp() {
	if p.Exp >= p.Level*100 {
		p.Level++
		p.Health += 20
		p.Mana += 10
		p.Exp = 0
		fmt.Printf("Level up! You are now level %d. Health: %d, Mana: %d\n", p.Level, p.Health, p.Mana)
	}
}

func (e *Enemy) Attack(p *Player) int {
	damage := rand.Intn(e.Damage) + 1
	p.Health -= damage
	return damage
}

func NewPlayer(name string) Player {
	return Player{
		Name:     name,
		Health:   100,
		Mana:     50,
		Level:    1,
		Exp:      0,
		Gold:     0,
		Inventory: []string{"Potion", "Sword"},
	}
}

func GenerateEnemies() []Enemy {
	enemies := []Enemy{
		{Name: "Goblin", Health: 30, Damage: 5, Reward: 10},
		{Name: "Orc", Health: 50, Damage: 10, Reward: 20},
		{Name: "Dragon", Health: 100, Damage: 20, Reward: 50},
	}
	return enemies
}

func SaveGame(state GameState) error {
	file, err := os.Create("savegame.json")
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(state)
}

func LoadGame() (GameState, error) {
	var state GameState
	file, err := os.Open("savegame.json")
	if err != nil {
		return state, err
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&state)
	return state, err
}

func DisplayStatus(p Player) {
	fmt.Printf("Name: %s, Health: %d, Mana: %d, Level: %d, Exp: %d, Gold: %d\n", p.Name, p.Health, p.Mana, p.Level, p.Exp, p.Gold)
	fmt.Printf("Inventory: %v\n", p.Inventory)
}

func Combat(p *Player, e *Enemy) bool {
	fmt.Printf("A wild %s appears!\n", e.Name)
	for p.Health > 0 && e.Health > 0 {
		fmt.Printf("Your health: %d, %s's health: %d\n", p.Health, e.Name, e.Health)
		fmt.Println("Choose action: (1) Attack, (2) Heal, (3) Run")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		switch input {
		case "1":
			damage := p.Attack(e)
			fmt.Printf("You dealt %d damage to %s!\n", damage, e.Name)
			if e.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", e.Name)
				p.Exp += e.Reward
				p.Gold += e.Reward
				p.LevelUp()
				return true
			}
		case "2":
			p.Heal()
		case "3":
			fmt.Println("You ran away!")
			return false
		default:
			fmt.Println("Invalid choice, try again.")
			continue
		}
		if e.Health > 0 {
			damage := e.Attack(p)
			fmt.Printf("%s attacked you for %d damage!\n", e.Name, damage)
		}
	}
	if p.Health <= 0 {
		fmt.Println("You have been defeated!")
		return false
	}
	return true
}

func ExploreRoom(p *Player, enemies []Enemy) bool {
	fmt.Println("You are in a dark room. What do you want to do?")
	fmt.Println("(1) Look for enemies, (2) Rest, (3) Check status, (4) Save game, (5) Quit")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "1":
		enemyIndex := rand.Intn(len(enemies))
		enemy := enemies[enemyIndex]
		if Combat(p, &enemy) {
			fmt.Println("You survived the combat!")
		} else {
			return false
		}
	case "2":
		p.Health += 10
		p.Mana += 5
		fmt.Printf("You rest and recover. Health: %d, Mana: %d\n", p.Health, p.Mana)
	case "3":
		DisplayStatus(*p)
	case "4":
		state := GameState{Player: *p, Enemies: enemies, CurrentRoom: 1, GameOver: false}
		err := SaveGame(state)
		if err != nil {
			fmt.Println("Error saving game:", err)
		} else {
			fmt.Println("Game saved successfully!")
		}
	case "5":
		fmt.Println("Thanks for playing!")
		return false
	default:
		fmt.Println("Invalid choice, try again.")
	}
	return true
}

func main() {
	fmt.Println("Welcome to the Fun Game Simulator!")
	fmt.Print("Enter your character's name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	var player Player
	var enemies []Enemy
	var currentRoom int
	var gameOver bool
	fmt.Print("Do you want to load a saved game? (y/n): ")
	loadInput, _ := reader.ReadString('\n')
	loadInput = strings.TrimSpace(loadInput)
	if loadInput == "y" || loadInput == "Y" {
		state, err := LoadGame()
		if err != nil {
			fmt.Println("No saved game found or error loading. Starting new game.")
			player = NewPlayer(name)
			enemies = GenerateEnemies()
			currentRoom = 1
			gameOver = false
		} else {
			player = state.Player
			enemies = state.Enemies
			currentRoom = state.CurrentRoom
			gameOver = state.GameOver
			fmt.Println("Game loaded successfully!")
		}
	} else {
		player = NewPlayer(name)
			enemies = GenerateEnemies()
			currentRoom = 1
			gameOver = false
	}
	for !gameOver {
		fmt.Printf("\n--- Room %d ---\n", currentRoom)
		if !ExploreRoom(&player, enemies) {
			gameOver = true
		} else {
			currentRoom++
			if currentRoom > 10 {
				fmt.Println("Congratulations! You have completed all rooms!")
				gameOver = true
			}
		}
	}
	fmt.Println("Game over. Thanks for playing!")
}
