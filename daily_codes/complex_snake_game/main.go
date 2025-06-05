package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

type Point struct {
	X, Y int
}

type Snake struct {
	Body  []Point
	Dir   Point
	Speed time.Duration
}

type Game struct {
	Snake    Snake
	Food     Point
	Score    int
	GameOver bool
	Width    int
	Height   int
}

func (g *Game) Init() {
	g.Width = 20
	g.Height = 20
	g.Snake = Snake{
		Body: []Point{{X: g.Width / 2, Y: g.Height / 2}},
		Dir:  Point{X: 0, Y: 1},
		Speed: 200 * time.Millisecond,
	}
	g.SpawnFood()
	g.Score = 0
	g.GameOver = false
}

func (g *Game) SpawnFood() {
	for {
		x := rand.Intn(g.Width)
		y := rand.Intn(g.Height)
		g.Food = Point{X: x, Y: y}
		collision := false
		for _, b := range g.Snake.Body {
			if b.X == x && b.Y == y {
				collision = true
				break
			}
		}
		if !collision {
			break
		}
	}
}

func (g *Game) Update() {
	if g.GameOver {
		return
	}

	head := g.Snake.Body[0]
	newHead := Point{X: head.X + g.Snake.Dir.X, Y: head.Y + g.Snake.Dir.Y}

	if newHead.X < 0 || newHead.X >= g.Width || newHead.Y < 0 || newHead.Y >= g.Height {
		g.GameOver = true
		return
	}

	for _, b := range g.Snake.Body {
		if b.X == newHead.X && b.Y == newHead.Y {
			g.GameOver = true
			return
		}
	}

	g.Snake.Body = append([]Point{newHead}, g.Snake.Body...)

	if newHead.X == g.Food.X && newHead.Y == g.Food.Y {
		g.Score++
		g.SpawnFood()
	} else {
		g.Snake.Body = g.Snake.Body[:len(g.Snake.Body)-1]
	}
}

func (g *Game) Draw() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			cell := " "
			if x == g.Food.X && y == g.Food.Y {
				cell = "🍎"
			} else {
				for _, b := range g.Snake.Body {
					if b.X == x && b.Y == y {
						cell = "🟩"
						break
					}
				}
			}
			fmt.Print(cell)
		}
		fmt.Println()
	}
	fmt.Printf("Score: %d\n", g.Score)
	if g.GameOver {
		fmt.Println("Game Over!")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := Game{}
	game.Init()

	go func() {
		for {
			var input string
			fmt.Scanln(&input)
			switch input {
			case "w":
				if game.Snake.Dir.Y == 0 {
					game.Snake.Dir = Point{X: 0, Y: -1}
				}
			case "s":
				if game.Snake.Dir.Y == 0 {
					game.Snake.Dir = Point{X: 0, Y: 1}
				}
			case "a":
				if game.Snake.Dir.X == 0 {
					game.Snake.Dir = Point{X: -1, Y: 0}
				}
			case "d":
				if game.Snake.Dir.X == 0 {
					game.Snake.Dir = Point{X: 1, Y: 0}
				}
			}
		}
	}()

	for !game.GameOver {
		game.Update()
		game.Draw()
		time.Sleep(game.Snake.Speed)
	}
}
