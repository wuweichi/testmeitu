package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
	"strings"
)

type Character struct {
	Name      string
	Health    int
	MaxHealth int
	Attack    int
	Defense   int
	Level     int
	Exp       int
	Inventory []string
}

type Enemy struct {
	Name      string
	Health    int
	Attack    int
	Defense   int
	ExpReward int
}

type GameState struct {
	Player     Character
	Enemies    []Enemy
	GameOver   bool
	Messages   []string
}

func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.Attack += 2
	c.Defense += 1
	c.Exp = 0
}

func (c *Character) GainExp(exp int) {
	c.Exp += exp
	if c.Exp >= c.Level*100 {
		c.LevelUp()
	}
}

func (c *Character) AttackEnemy(e *Enemy) int {
	damage := c.Attack - e.Defense
	if damage < 1 {
		damage = 1
	}
	e.Health -= damage
	if e.Health < 0 {
		e.Health = 0
	}
	return damage
}

func (e *Enemy) AttackPlayer(c *Character) int {
	damage := e.Attack - c.Defense
	if damage < 1 {
		damage = 1
	}
	c.Health -= damage
	if c.Health < 0 {
		c.Health = 0
	}
	return damage
}

func (gs *GameState) AddMessage(msg string) {
	gs.Messages = append(gs.Messages, msg)
}

func (gs *GameState) PrintMessages() {
	for _, msg := range gs.Messages {
		fmt.Println(msg)
	}
	gs.Messages = []string{}
}

func (gs *GameState) GenerateEnemies() {
	gs.Enemies = []Enemy{}
	enemyTypes := []string{"Goblin", "Orc", "Dragon", "Skeleton", "Zombie"}
	numEnemies := rand.Intn(3) + 1
	for i := 0; i < numEnemies; i++ {
		enemyType := enemyTypes[rand.Intn(len(enemyTypes))]
		health := rand.Intn(20) + 10
		attack := rand.Intn(5) + 5
		defense := rand.Intn(3) + 2
		expReward := rand.Intn(10) + 5
		gs.Enemies = append(gs.Enemies, Enemy{
			Name:      enemyType,
			Health:    health,
			Attack:    attack,
			Defense:   defense,
			ExpReward: expReward,
		})
	}
}

func (gs *GameState) PlayerTurn() {
	fmt.Println("\nYour turn!")
	fmt.Println("Choose an action:")
	fmt.Println("1. Attack")
	fmt.Println("2. Use Item")
	fmt.Println("3. Check Status")
	fmt.Println("4. Flee")

	var choice int
	fmt.Scan(&choice)

	switch choice {
	case 1:
		gs.AttackAction()
	case 2:
		gs.UseItemAction()
	case 3:
		gs.CheckStatusAction()
	case 4:
		gs.FleeAction()
	default:
		gs.AddMessage("Invalid choice!")
	}
}

func (gs *GameState) AttackAction() {
	if len(gs.Enemies) == 0 {
		gs.AddMessage("No enemies to attack!")
		return
	}

	fmt.Println("Choose an enemy to attack:")
	for i, enemy := range gs.Enemies {
		fmt.Printf("%d. %s (Health: %d)\n", i+1, enemy.Name, enemy.Health)
	}

	var choice int
	fmt.Scan(&choice)

	if choice < 1 || choice > len(gs.Enemies) {
		gs.AddMessage("Invalid enemy choice!")
		return
	}

	enemyIndex := choice - 1
	damage := gs.Player.AttackEnemy(&gs.Enemies[enemyIndex])
	gs.AddMessage(fmt.Sprintf("You attack %s for %d damage!", gs.Enemies[enemyIndex].Name, damage))

	if gs.Enemies[enemyIndex].Health == 0 {
		gs.AddMessage(fmt.Sprintf("%s defeated! Gained %d exp.", gs.Enemies[enemyIndex].Name, gs.Enemies[enemyIndex].ExpReward))
		gs.Player.GainExp(gs.Enemies[enemyIndex].ExpReward)
		gs.Enemies = append(gs.Enemies[:enemyIndex], gs.Enemies[enemyIndex+1:]...)
	}
}

func (gs *GameState) UseItemAction() {
	if len(gs.Player.Inventory) == 0 {
		gs.AddMessage("No items in inventory!")
		return
	}

	fmt.Println("Choose an item to use:")
	for i, item := range gs.Player.Inventory {
		fmt.Printf("%d. %s\n", i+1, item)
	}

	var choice int
	fmt.Scan(&choice)

	if choice < 1 || choice > len(gs.Player.Inventory) {
		gs.AddMessage("Invalid item choice!")
		return
	}

	itemIndex := choice - 1
	item := gs.Player.Inventory[itemIndex]

	switch item {
	case "Health Potion":
		gs.Player.Health += 20
		if gs.Player.Health > gs.Player.MaxHealth {
			gs.Player.Health = gs.Player.MaxHealth
		}
		gs.AddMessage("Used Health Potion! Restored 20 health.")
	case "Attack Boost":
		gs.Player.Attack += 5
		gs.AddMessage("Used Attack Boost! Increased attack by 5.")
	case "Defense Boost":
		gs.Player.Defense += 3
		gs.AddMessage("Used Defense Boost! Increased defense by 3.")
	default:
		gs.AddMessage("Item effect not implemented!")
	}

	gs.Player.Inventory = append(gs.Player.Inventory[:itemIndex], gs.Player.Inventory[itemIndex+1:]...)
}

func (gs *GameState) CheckStatusAction() {
	gs.AddMessage(fmt.Sprintf("Name: %s", gs.Player.Name))
	gs.AddMessage(fmt.Sprintf("Level: %d", gs.Player.Level))
	gs.AddMessage(fmt.Sprintf("Health: %d/%d", gs.Player.Health, gs.Player.MaxHealth))
	gs.AddMessage(fmt.Sprintf("Attack: %d", gs.Player.Attack))
	gs.AddMessage(fmt.Sprintf("Defense: %d", gs.Player.Defense))
	gs.AddMessage(fmt.Sprintf("Exp: %d/%d", gs.Player.Exp, gs.Player.Level*100))
	gs.AddMessage(fmt.Sprintf("Inventory: %v", gs.Player.Inventory))
}

func (gs *GameState) FleeAction() {
	if rand.Float32() < 0.5 {
		gs.AddMessage("You successfully fled!")
		gs.GameOver = true
	} else {
		gs.AddMessage("Failed to flee!")
	}
}

func (gs *GameState) EnemyTurn() {
	if len(gs.Enemies) == 0 {
		return
	}

	for i := range gs.Enemies {
		if gs.Enemies[i].Health > 0 {
			damage := gs.Enemies[i].AttackPlayer(&gs.Player)
			gs.AddMessage(fmt.Sprintf("%s attacks you for %d damage!", gs.Enemies[i].Name, damage))
		}
	}
}

func (gs *GameState) CheckGameOver() {
	if gs.Player.Health <= 0 {
		gs.AddMessage("You have been defeated! Game Over.")
		gs.GameOver = true
	} else if len(gs.Enemies) == 0 {
		gs.AddMessage("All enemies defeated! You win!")
		gs.GameOver = true
	}
}

func (gs *GameState) Battle() {
	gs.GenerateEnemies()
	gs.AddMessage("A battle begins!")
	gs.AddMessage(fmt.Sprintf("Enemies: %v", func() []string {
		var names []string
		for _, enemy := range gs.Enemies {
			names = append(names, enemy.Name)
		}
		return names
	}()))

	for !gs.GameOver {
		gs.PlayerTurn()
		gs.CheckGameOver()
		if gs.GameOver {
			break
		}
		gs.EnemyTurn()
		gs.CheckGameOver()
		gs.PrintMessages()
	}
}

func (gs *GameState) Explore() {
	events := []string{
		"You find a treasure chest!",
		"You discover a hidden path.",
		"You encounter a friendly traveler.",
		"You stumble upon an ancient ruin.",
		"You find a mysterious artifact.",
	}

	event := events[rand.Intn(len(events))]
	gs.AddMessage(event)

	switch event {
	case "You find a treasure chest!":
		items := []string{"Health Potion", "Attack Boost", "Defense Boost"}
		item := items[rand.Intn(len(items))]
		gs.Player.Inventory = append(gs.Player.Inventory, item)
		gs.AddMessage(fmt.Sprintf("You found a %s!", item))
	case "You discover a hidden path.":
		gs.Player.GainExp(10)
		gs.AddMessage("Gained 10 exp for exploration!")
	case "You encounter a friendly traveler.":
		gs.Player.Health = gs.Player.MaxHealth
		gs.AddMessage("The traveler heals you to full health!")
	case "You stumble upon an ancient ruin.":
		gs.Player.Attack += 1
		gs.AddMessage("You gain 1 attack from the ruin's knowledge!")
	case "You find a mysterious artifact.":
		gs.Player.Defense += 1
		gs.AddMessage("You gain 1 defense from the artifact's power!")
	}
}

func (gs *GameState) MainMenu() {
	for {
		fmt.Println("\nMain Menu:")
		fmt.Println("1. Start Battle")
		fmt.Println("2. Explore")
		fmt.Println("3. Check Status")
		fmt.Println("4. Quit")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			gs.Battle()
		case 2:
			gs.Explore()
			gs.PrintMessages()
		case 3:
			gs.CheckStatusAction()
			gs.PrintMessages()
		case 4:
			gs.AddMessage("Thanks for playing!")
			gs.PrintMessages()
			return
		default:
			gs.AddMessage("Invalid choice!")
			gs.PrintMessages()
		}

		if gs.GameOver {
			gs.AddMessage("Game Over!")
			gs.PrintMessages()
			return
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	player := Character{
		Name:      "Hero",
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
		Inventory: []string{"Health Potion", "Attack Boost"},
	}

	gameState := GameState{
		Player:   player,
		GameOver: false,
		Messages: []string{},
	}

	fmt.Println("Welcome to the Complex Game Simulator!")
	gameState.MainMenu()
}
