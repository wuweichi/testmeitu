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
	g.Score = 0
	g.GameOver = false
	g.placeFood()
}

func (g *Game) placeFood() {
	for {
		g.Food = Point{
			X: rand.Intn(g.Width),
			Y: rand.Intn(g.Height),
		}
		if !g.isPointOnSnake(g.Food) {
			break
		}
	}
}

func (g *Game) isPointOnSnake(p Point) bool {
	for _, b := range g.Snake.Body {
		if b == p {
			return true
		}
	}
	return false
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

	for _, b := range g.Snake.Body[1:] {
		if b == newHead {
			g.GameOver = true
			return
		}
	}

	g.Snake.Body = append([]Point{newHead}, g.Snake.Body...)

	if newHead == g.Food {
		g.Score++
		g.placeFood()
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
			p := Point{X: x, Y: y}
			if p == g.Food {
				fmt.Print("F")
			} else if g.isPointOnSnake(p) {
				fmt.Print("S")
			} else {
				fmt.Print(".")
			}
			fmt.Print(" ")
		}
		fmt.Println()
	}
	fmt.Printf("Score: %d\n", g.Score)
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
				game.Snake.Dir = Point{X: 0, Y: -1}
			case "s":
				game.Snake.Dir = Point{X: 0, Y: 1}
			case "a":
				game.Snake.Dir = Point{X: -1, Y: 0}
			case "d":
				game.Snake.Dir = Point{X: 1, Y: 0}
			}
		}
	}()

	for {
		game.Draw()
		game.Update()
		time.Sleep(game.Snake.Speed)
		if game.GameOver {
			fmt.Println("Game Over!")
			break
		}
	}
}
