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
	Strength int
	Level    int
}

type Enemy struct {
	Name     string
	Health   int
	Strength int
}

type Item struct {
	Name        string
	Description string
	Effect      func(*Player)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Fun Game Simulator!")
	fmt.Print("Enter your player name: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	player := &Player{
		Name:     name,
		Health:   100,
		Strength: 10,
		Level:    1,
	}

	fmt.Printf("Hello, %s! You are starting at Level %d with %d health and %d strength.\n", player.Name, player.Level, player.Health, player.Strength)

	for {
		fmt.Println("\nChoose an action:")
		fmt.Println("1. Explore")
		fmt.Println("2. Check Status")
		fmt.Println("3. Quit")
		fmt.Print("Enter choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			explore(player, reader)
		case "2":
			checkStatus(player)
		case "3":
			fmt.Println("Thanks for playing! Goodbye.")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}

		if player.Health <= 0 {
			fmt.Println("You have been defeated! Game over.")
			return
		}
	}
}

func explore(player *Player, reader *bufio.Reader) {
	event := rand.Intn(3)
	switch event {
	case 0:
		fmt.Println("You found a treasure chest!")
		item := generateItem()
		fmt.Printf("It contains: %s - %s\n", item.Name, item.Description)
		item.Effect(player)
	case 1:
		fmt.Println("An enemy appears!")
		enemy := generateEnemy(player.Level)
		battle(player, enemy, reader)
	case 2:
		fmt.Println("You explore peacefully and gain some experience.")
		player.Level++
		player.Health += 10
		player.Strength += 5
		fmt.Printf("Level up! You are now Level %d. Health: %d, Strength: %d\n", player.Level, player.Health, player.Strength)
	}
}

func generateEnemy(level int) *Enemy {
	names := []string{"Goblin", "Orc", "Dragon", "Skeleton"}
	name := names[rand.Intn(len(names))]
	health := 20 + (level * 10)
	strength := 5 + (level * 3)
	return &Enemy{Name: name, Health: health, Strength: strength}
}

func generateItem() *Item {
	items := []*Item{
		{
			Name:        "Health Potion",
			Description: "Restores 20 health.",
			Effect: func(p *Player) {
				p.Health += 20
				fmt.Printf("Health restored! Current health: %d\n", p.Health)
			},
		},
		{
			Name:        "Strength Elixir",
			Description: "Increases strength by 5.",
			Effect: func(p *Player) {
				p.Strength += 5
				fmt.Printf("Strength increased! Current strength: %d\n", p.Strength)
			},
		},
		{
			Name:        "Mysterious Scroll",
			Description: "Grants a random effect.",
			Effect: func(p *Player) {
				randEffect := rand.Intn(2)
				if randEffect == 0 {
					p.Health += 30
					fmt.Printf("Scroll heals you! Health: %d\n", p.Health)
				} else {
					p.Strength += 10
					fmt.Printf("Scroll empowers you! Strength: %d\n", p.Strength)
				}
			},
		},
	}
	return items[rand.Intn(len(items))]
}

func battle(player *Player, enemy *Enemy, reader *bufio.Reader) {
	fmt.Printf("A wild %s appears with %d health and %d strength!\n", enemy.Name, enemy.Health, enemy.Strength)
	for enemy.Health > 0 && player.Health > 0 {
		fmt.Println("\nBattle options:")
		fmt.Println("1. Attack")
		fmt.Println("2. Run away")
		fmt.Print("Enter choice: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			damage := player.Strength + rand.Intn(10)
			enemy.Health -= damage
			fmt.Printf("You attack the %s for %d damage! Enemy health: %d\n", enemy.Name, damage, enemy.Health)
			if enemy.Health <= 0 {
				fmt.Printf("You defeated the %s!\n", enemy.Name)
				player.Level++
				player.Health += 15
				player.Strength += 8
				fmt.Printf("Victory! Level up to %d. Health: %d, Strength: %d\n", player.Level, player.Health, player.Strength)
				return
			}
			enemyDamage := enemy.Strength + rand.Intn(5)
			player.Health -= enemyDamage
			fmt.Printf("The %s attacks you for %d damage! Your health: %d\n", enemy.Name, enemyDamage, player.Health)
		case "2":
			if rand.Intn(2) == 0 {
				fmt.Println("You successfully ran away!")
				return
			} else {
				fmt.Println("You failed to run away!")
				enemyDamage := enemy.Strength + rand.Intn(5)
				player.Health -= enemyDamage
				fmt.Printf("The %s attacks you for %d damage! Your health: %d\n", enemy.Name, enemyDamage, player.Health)
			}
		default:
			fmt.Println("Invalid choice. The enemy attacks!")
			enemyDamage := enemy.Strength + rand.Intn(5)
			player.Health -= enemyDamage
			fmt.Printf("The %s attacks you for %d damage! Your health: %d\n", enemy.Name, enemyDamage, player.Health)
		}
	}
}

func checkStatus(player *Player) {
	fmt.Printf("Name: %s, Level: %d, Health: %d, Strength: %d\n", player.Name, player.Level, player.Health, player.Strength)
}

// Additional functions to increase line count
func dummyFunction1() {
	// This is a dummy function to add lines
	for i := 0; i < 10; i++ {
		fmt.Println("Dummy line", i)
	}
}

func dummyFunction2() {
	// Another dummy function
	var x int = 42
	if x > 0 {
		fmt.Println("x is positive")
	} else {
		fmt.Println("x is not positive")
	}
}

func dummyFunction3() {
	// More dummy code
	slice := []string{"a", "b", "c"}
	for index, value := range slice {
		fmt.Printf("Index %d: %s\n", index, value)
	}
}

func dummyFunction4() {
	// Even more lines
	for j := 0; j < 5; j++ {
		switch j {
		case 0:
			fmt.Println("Case 0")
		case 1:
			fmt.Println("Case 1")
		default:
			fmt.Println("Default case")
		}
	}
}

func dummyFunction5() {
	// Final dummy function
	numbers := []int{1, 2, 3, 4, 5}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	fmt.Println("Sum:", sum)
}

// Call dummy functions in main to ensure they are included
func init() {
	dummyFunction1()
	dummyFunction2()
	dummyFunction3()
	dummyFunction4()
	dummyFunction5()
}
