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
	Agility  int
	Level    int
	Experience int
}

type Monster struct {
	Name     string
	Health   int
	Strength int
	Agility  int
}

type Item struct {
	Name        string
	Description string
	Value       int
}

type GameWorld struct {
	Players   []Player
	Monsters  []Monster
	Items     []Item
	Locations []string
}

func (p *Player) Attack(m *Monster) {
	damage := p.Strength + rand.Intn(10)
	m.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", p.Name, m.Name, damage)
}

func (m *Monster) Attack(p *Player) {
	damage := m.Strength + rand.Intn(5)
	p.Health -= damage
	fmt.Printf("%s attacks %s for %d damage!\n", m.Name, p.Name, damage)
}

func (p *Player) Heal() {
	healAmount := 20 + rand.Intn(10)
	p.Health += healAmount
	fmt.Printf("%s heals for %d health. Current health: %d\n", p.Name, healAmount, p.Health)
}

func (p *Player) LevelUp() {
	if p.Experience >= 100 {
		p.Level++
		p.Experience -= 100
		p.Strength += 5
		p.Agility += 3
		fmt.Printf("%s leveled up to level %d!\n", p.Name, p.Level)
	}
}

func (p *Player) GainExperience(amount int) {
	p.Experience += amount
	fmt.Printf("%s gained %d experience. Total: %d\n", p.Name, amount, p.Experience)
	p.LevelUp()
}

func CreatePlayer(name string) Player {
	return Player{
		Name:     name,
		Health:   100,
		Strength: 10,
		Agility:  10,
		Level:    1,
		Experience: 0,
	}
}

func CreateMonster(name string, health, strength, agility int) Monster {
	return Monster{
		Name:     name,
		Health:   health,
		Strength: strength,
		Agility:  agility,
	}
}

func CreateItem(name, description string, value int) Item {
	return Item{
		Name:        name,
		Description: description,
		Value:       value,
	}
}

func InitializeGameWorld() GameWorld {
	world := GameWorld{
		Players: []Player{
			CreatePlayer("Hero"),
			CreatePlayer("Mage"),
		},
		Monsters: []Monster{
			CreateMonster("Goblin", 50, 8, 12),
			CreateMonster("Orc", 80, 12, 6),
			CreateMonster("Dragon", 200, 20, 15),
		},
		Items: []Item{
			CreateItem("Health Potion", "Restores 50 health", 50),
			CreateItem("Strength Elixir", "Increases strength by 10", 100),
			CreateItem("Agility Boots", "Increases agility by 5", 75),
		},
		Locations: []string{
			"Forest",
			"Cave",
			"Mountain",
			"Castle",
		},
	}
	return world
}

func DisplayPlayerInfo(p Player) {
	fmt.Printf("Name: %s\n", p.Name)
	fmt.Printf("Health: %d\n", p.Health)
	fmt.Printf("Strength: %d\n", p.Strength)
	fmt.Printf("Agility: %d\n", p.Agility)
	fmt.Printf("Level: %d\n", p.Level)
	fmt.Printf("Experience: %d\n", p.Experience)
	fmt.Println()
}

func DisplayMonsterInfo(m Monster) {
	fmt.Printf("Name: %s\n", m.Name)
	fmt.Printf("Health: %d\n", m.Health)
	fmt.Printf("Strength: %d\n", m.Strength)
	fmt.Printf("Agility: %d\n", m.Agility)
	fmt.Println()
}

func DisplayItemInfo(i Item) {
	fmt.Printf("Name: %s\n", i.Name)
	fmt.Printf("Description: %s\n", i.Description)
	fmt.Printf("Value: %d\n", i.Value)
	fmt.Println()
}

func DisplayWorldInfo(world GameWorld) {
	fmt.Println("=== Game World Info ===")
	fmt.Println("Players:")
	for _, player := range world.Players {
		DisplayPlayerInfo(player)
	}
	fmt.Println("Monsters:")
	for _, monster := range world.Monsters {
		DisplayMonsterInfo(monster)
	}
	fmt.Println("Items:")
	for _, item := range world.Items {
		DisplayItemInfo(item)
	}
	fmt.Println("Locations:")
	for _, location := range world.Locations {
		fmt.Printf("- %s\n", location)
	}
	fmt.Println()
}

func SimulateBattle(player *Player, monster *Monster) {
	fmt.Printf("Battle between %s and %s begins!\n", player.Name, monster.Name)
	for player.Health > 0 && monster.Health > 0 {
		player.Attack(monster)
		if monster.Health <= 0 {
			fmt.Printf("%s has been defeated!\n", monster.Name)
			player.GainExperience(50)
			break
		}
		monster.Attack(player)
		if player.Health <= 0 {
			fmt.Printf("%s has been defeated!\n", player.Name)
			break
		}
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Battle ended.")
}

func ExploreLocation(player *Player, location string) {
	fmt.Printf("%s is exploring %s...\n", player.Name, location)
	// Simulate random events during exploration
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Printf("%s found a treasure chest!\n", player.Name)
		player.GainExperience(20)
	case 1:
		fmt.Printf("%s encountered a friendly NPC.\n", player.Name)
		player.Heal()
	case 2:
		fmt.Printf("%s discovered a hidden path.\n", player.Name)
		player.GainExperience(10)
	}
}

func main() {
	// Initialize random seed
	rand.Seed(time.Now().UnixNano())

	// Create game world
	world := InitializeGameWorld()

	// Display initial world info
	DisplayWorldInfo(world)

	// Simulate some gameplay
	fmt.Println("=== Starting Gameplay ===")

	// Player 1 explores and battles
	player1 := &world.Players[0]
	fmt.Printf("\n%s's Adventure:\n", player1.Name)
	for i := 0; i < 3; i++ {
		location := world.Locations[rand.Intn(len(world.Locations))]
		ExploreLocation(player1, location)
		if i < len(world.Monsters) {
			monster := &world.Monsters[i]
			SimulateBattle(player1, monster)
		}
		if player1.Health <= 0 {
			fmt.Printf("%s has fallen in battle! Game Over.\n", player1.Name)
			break
		}
	}

	// Player 2 explores and battles
	player2 := &world.Players[1]
	fmt.Printf("\n%s's Adventure:\n", player2.Name)
	for i := 0; i < 3; i++ {
		location := world.Locations[rand.Intn(len(world.Locations))]
		ExploreLocation(player2, location)
		if i < len(world.Monsters) {
			monster := &world.Monsters[i]
			SimulateBattle(player2, monster)
		}
		if player2.Health <= 0 {
			fmt.Printf("%s has fallen in battle! Game Over.\n", player2.Name)
			break
		}
	}

	// Display final world info
	fmt.Println("\n=== Final Game World Info ===")
	DisplayWorldInfo(world)

	fmt.Println("\nGame simulation completed!")
}
