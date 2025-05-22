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
	Speed int
}

type Game struct {
	Snake    Snake
	Food     Point
	Score    int
	GameOver bool
}

func (g *Game) Init() {
	g.Snake = Snake{
		Body:  []Point{{X: 10, Y: 10}},
		Dir:   Point{X: 1, Y: 0},
		Speed: 1,
	}
	g.SpawnFood()
	g.Score = 0
	g.GameOver = false
}

func (g *Game) SpawnFood() {
	rand.Seed(time.Now().UnixNano())
	g.Food = Point{
		X: rand.Intn(20),
		Y: rand.Intn(20),
	}
}

func (g *Game) Update() {
	if g.GameOver {
		return
	}

	head := g.Snake.Body[0]
	newHead := Point{X: head.X + g.Snake.Dir.X, Y: head.Y + g.Snake.Dir.Y}

	if newHead.X < 0 || newHead.X >= 20 || newHead.Y < 0 || newHead.Y >= 20 {
		g.GameOver = true
		return
	}

	for _, b := range g.Snake.Body {
		if b == newHead {
			g.GameOver = true
			return
		}
	}

	g.Snake.Body = append([]Point{newHead}, g.Snake.Body...)


	if newHead == g.Food {
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

	for y := 0; y < 20; y++ {
		for x := 0; x < 20; x++ {
			p := Point{X: x, Y: y}
			if p == g.Food {
				fmt.Print("F")
			} else {
				isBody := false
				for _, b := range g.Snake.Body {
					if b == p {
						fmt.Print("O")
						isBody = true
						break
					}
				}
				if !isBody {
					fmt.Print(".")
				}
			}
		}
		fmt.Println()
	}
	fmt.Printf("Score: %d\n", g.Score)
}

func main() {
	var game Game
	game.Init()

	for !game.GameOver {
		game.Draw()
		game.Update()
		time.Sleep(time.Second / time.Duration(game.Snake.Speed))
	}

	fmt.Println("Game Over!")
}
