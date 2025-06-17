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
	Body      []Point
	Direction Point
	Length    int
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
		Body:      []Point{{X: g.Width / 2, Y: g.Height / 2}},
		Direction: Point{X: 0, Y: 1},
		Length:    1,
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
	newHead := Point{X: head.X + g.Snake.Direction.X, Y: head.Y + g.Snake.Direction.Y}

	if newHead.X < 0 || newHead.X >= g.Width || newHead.Y < 0 || newHead.Y >= g.Height || g.isPointOnSnake(newHead) {
		g.GameOver = true
		return
	}

	g.Snake.Body = append([]Point{newHead}, g.Snake.Body...)
	if len(g.Snake.Body) > g.Snake.Length {
		g.Snake.Body = g.Snake.Body[:len(g.Snake.Body)-1]
	}

	if newHead == g.Food {
		g.Score++
		g.Snake.Length++
		g.placeFood()
	}
}

func (g *Game) Draw() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			p := Point{X: x, Y: y}
			if g.isPointOnSnake(p) {
				fmt.Print("■")
			} else if p == g.Food {
				fmt.Print("★")
			} else {
				fmt.Print("□")
			}
			fmt.Print(" ")
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

	for !game.GameOver {
		game.Draw()
		time.Sleep(200 * time.Millisecond)
		game.Update()
	}
	game.Draw()
}
