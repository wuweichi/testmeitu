package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Character struct {
	Name      string
	Health    int
	MaxHealth int
	Attack    int
	Defense   int
	Level     int
	Exp       int
}

type Enemy struct {
	Name      string
	Health    int
	Attack    int
	Defense   int
	ExpReward int
}

type Item struct {
	Name        string
	Description string
	Effect      func(*Character)
}

type GameState struct {
	Player      Character
	Enemies     []Enemy
	Inventory   []Item
	CurrentRoom int
	GameOver    bool
}

func (c *Character) TakeDamage(damage int) {
	actualDamage := damage - c.Defense
	if actualDamage < 0 {
		actualDamage = 0
	}
	c.Health -= actualDamage
	if c.Health < 0 {
		c.Health = 0
	}
}

func (c *Character) Heal(amount int) {
	c.Health += amount
	if c.Health > c.MaxHealth {
		c.Health = c.MaxHealth
	}
}

func (c *Character) LevelUp() {
	c.Level++
	c.MaxHealth += 10
	c.Health = c.MaxHealth
	c.Attack += 2
	c.Defense += 1
	fmt.Printf("Level up! %s is now level %d\n", c.Name, c.Level)
}

func (c *Character) GainExp(exp int) {
	c.Exp += exp
	if c.Exp >= c.Level*100 {
		c.LevelUp()
		c.Exp = 0
	}
}

func (e *Enemy) TakeDamage(damage int) {
	actualDamage := damage - e.Defense
	if actualDamage < 0 {
		actualDamage = 0
	}
	e.Health -= actualDamage
	if e.Health < 0 {
		e.Health = 0
	}
}

func CreatePlayer(name string) Character {
	return Character{
		Name:      name,
		Health:    100,
		MaxHealth: 100,
		Attack:    10,
		Defense:   5,
		Level:     1,
		Exp:       0,
	}
}

func CreateEnemy(name string, health, attack, defense, expReward int) Enemy {
	return Enemy{
		Name:      name,
		Health:    health,
		Attack:    attack,
		Defense:   defense,
		ExpReward: expReward,
	}
}

func CreateHealthPotion() Item {
	return Item{
		Name:        "Health Potion",
		Description: "Restores 30 health points.",
		Effect: func(c *Character) {
			c.Heal(30)
			fmt.Printf("%s used a Health Potion and restored 30 health. Current health: %d\n", c.Name, c.Health)
		},
	}
}

func CreateAttackBoost() Item {
	return Item{
		Name:        "Attack Boost",
		Description: "Increases attack by 5 for the next battle.",
		Effect: func(c *Character) {
			c.Attack += 5
			fmt.Printf("%s used an Attack Boost. Attack increased by 5. Current attack: %d\n", c.Name, c.Attack)
		},
	}
}

func CreateDefenseBoost() Item {
	return Item{
		Name:        "Defense Boost",
		Description: "Increases defense by 3 for the next battle.",
		Effect: func(c *Character) {
			c.Defense += 3
			fmt.Printf("%s used a Defense Boost. Defense increased by 3. Current defense: %d\n", c.Name, c.Defense)
		},
	}
}

func InitializeGame() GameState {
	player := CreatePlayer("Hero")
	enemies := []Enemy{
		CreateEnemy("Goblin", 50, 8, 2, 20),
		CreateEnemy("Orc", 80, 12, 4, 40),
		CreateEnemy("Dragon", 150, 20, 8, 100),
	}
	inventory := []Item{
		CreateHealthPotion(),
		CreateAttackBoost(),
		CreateDefenseBoost(),
	}
	return GameState{
		Player:      player,
		Enemies:     enemies,
		Inventory:   inventory,
		CurrentRoom: 0,
		GameOver:    false,
	}
}

func (gs *GameState) DisplayStatus() {
	fmt.Printf("=== Player Status ===\n")
	fmt.Printf("Name: %s\n", gs.Player.Name)
	fmt.Printf("Health: %d/%d\n", gs.Player.Health, gs.Player.MaxHealth)
	fmt.Printf("Attack: %d\n", gs.Player.Attack)
	fmt.Printf("Defense: %d\n", gs.Player.Defense)
	fmt.Printf("Level: %d\n", gs.Player.Level)
	fmt.Printf("Exp: %d/%d\n", gs.Player.Exp, gs.Player.Level*100)
	fmt.Printf("Current Room: %d\n", gs.CurrentRoom)
	fmt.Printf("=====================\n")
}

func (gs *GameState) DisplayInventory() {
	fmt.Printf("=== Inventory ===\n")
	for i, item := range gs.Inventory {
		fmt.Printf("%d. %s - %s\n", i+1, item.Name, item.Description)
	}
	fmt.Printf("=================\n")
}

func (gs *GameState) UseItem(index int) {
	if index < 0 || index >= len(gs.Inventory) {
		fmt.Println("Invalid item index.")
		return
	}
	item := gs.Inventory[index]
	item.Effect(&gs.Player)
	gs.Inventory = append(gs.Inventory[:index], gs.Inventory[index+1:]...)
}

func (gs *GameState) Battle(enemyIndex int) {
	if enemyIndex < 0 || enemyIndex >= len(gs.Enemies) {
		fmt.Println("Invalid enemy index.")
		return
	}
	enemy := &gs.Enemies[enemyIndex]
	fmt.Printf("A wild %s appears!\n", enemy.Name)
	for gs.Player.Health > 0 && enemy.Health > 0 {
		fmt.Printf("\n--- Battle Round ---\n")
		fmt.Printf("%s: %d HP\n", gs.Player.Name, gs.Player.Health)
		fmt.Printf("%s: %d HP\n", enemy.Name, enemy.Health)
		fmt.Println("1. Attack")
		fmt.Println("2. Use Item")
		fmt.Println("3. Flee")
		var choice int
		fmt.Print("Choose an action: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			playerDamage := gs.Player.Attack + rand.Intn(5)
			enemy.TakeDamage(playerDamage)
			fmt.Printf("%s attacks %s for %d damage!\n", gs.Player.Name, enemy.Name, playerDamage)
			if enemy.Health <= 0 {
				fmt.Printf("%s defeated the %s!\n", gs.Player.Name, enemy.Name)
				gs.Player.GainExp(enemy.ExpReward)
				gs.Enemies = append(gs.Enemies[:enemyIndex], gs.Enemies[enemyIndex+1:]...)
				return
			}
			enemyDamage := enemy.Attack + rand.Intn(3)
			gs.Player.TakeDamage(enemyDamage)
			fmt.Printf("%s attacks %s for %d damage!\n", enemy.Name, gs.Player.Name, enemyDamage)
			if gs.Player.Health <= 0 {
				fmt.Printf("%s has been defeated! Game Over.\n", gs.Player.Name)
				gs.GameOver = true
				return
			}
		case 2:
			gs.DisplayInventory()
			var itemChoice int
			fmt.Print("Choose an item to use: ")
			fmt.Scan(&itemChoice)
			gs.UseItem(itemChoice - 1)
		case 3:
			fmt.Printf("%s fled from the battle!\n", gs.Player.Name)
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (gs *GameState) Explore() {
	fmt.Println("You are exploring the area...")
	randomEvent := rand.Intn(3)
	switch randomEvent {
	case 0:
		fmt.Println("You found a treasure chest!")
		newItem := CreateHealthPotion()
		gs.Inventory = append(gs.Inventory, newItem)
		fmt.Printf("You obtained a %s!\n", newItem.Name)
	case 1:
		if len(gs.Enemies) > 0 {
			enemyIndex := rand.Intn(len(gs.Enemies))
			gs.Battle(enemyIndex)
		} else {
			fmt.Println("You found a peaceful clearing. No enemies here.")
		}
	case 2:
		fmt.Println("You discovered a hidden shrine. Your stats are boosted!")
		gs.Player.Attack += 2
		gs.Player.Defense += 1
		fmt.Printf("Attack increased by 2, Defense increased by 1. Current stats: Attack=%d, Defense=%d\n", gs.Player.Attack, gs.Player.Defense)
	}
}

func (gs *GameState) MoveToNextRoom() {
	if gs.CurrentRoom >= 2 {
		fmt.Println("You have reached the final room. Defeat the boss to win!")
		return
	}
	gs.CurrentRoom++
	fmt.Printf("You moved to room %d.\n", gs.CurrentRoom)
	if gs.CurrentRoom == 2 {
		fmt.Println("This is the boss room! Prepare for a tough battle.")
	}
}

func (gs *GameState) SaveGame() {
	fmt.Println("Game saved successfully.")
}

func (gs *GameState) LoadGame() {
	fmt.Println("Game loaded successfully.")
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := InitializeGame()
	fmt.Println("Welcome to the Complex Game Simulator!")
	fmt.Println("You are a hero on a quest to defeat monsters and level up.")
	fmt.Println("Commands: status, inventory, use <item_index>, battle <enemy_index>, explore, move, save, load, quit")
	for !game.GameOver {
		fmt.Print("\nEnter command: ")
		var command string
		fmt.Scan(&command)
		switch command {
		case "status":
			game.DisplayStatus()
		case "inventory":
			game.DisplayInventory()
		case "use":
			var index int
			fmt.Scan(&index)
			game.UseItem(index - 1)
		case "battle":
			var index int
			fmt.Scan(&index)
			game.Battle(index - 1)
		case "explore":
			game.Explore()
		case "move":
			game.MoveToNextRoom()
		case "save":
			game.SaveGame()
		case "load":
			game.LoadGame()
		case "quit":
			fmt.Println("Thanks for playing!")
			return
		default:
			fmt.Println("Unknown command. Please try again.")
		}
		if len(game.Enemies) == 0 && game.CurrentRoom >= 2 {
			fmt.Println("Congratulations! You have defeated all enemies and won the game!")
			game.GameOver = true
		}
	}
	fmt.Println("Game Over.")
}
