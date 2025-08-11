package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
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
	Snake Snake
	Food  Point
	Score int
	Over  bool
}

func (g *Game) Init() {
	termbox.Init()
	width, height := termbox.Size()
	g.Snake = Snake{
		Body:  []Point{{width / 2, height / 2}},
		Dir:   Point{1, 0},
		Speed: 100 * time.Millisecond,
	}
	g.SpawnFood()
	g.Score = 0
	g.Over = false
}

func (g *Game) SpawnFood() {
	width, height := termbox.Size()
	g.Food = Point{rand.Intn(width-2) + 1, rand.Intn(height-2) + 1}
}

func (g *Game) Draw() {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
	for _, p := range g.Snake.Body {
		termbox.SetCell(p.X, p.Y, 'O', termbox.ColorGreen, termbox.ColorDefault)
	}
	termbox.SetCell(g.Food.X, g.Food.Y, '@', termbox.ColorRed, termbox.ColorDefault)
	termbox.Flush()
}

func (g *Game) Update() {
	head := g.Snake.Body[0]
	newHead := Point{head.X + g.Snake.Dir.X, head.Y + g.Snake.Dir.Y}
	if newHead.X <= 0 || newHead.X >= termbox.Size().X-1 || newHead.Y <= 0 || newHead.Y >= termbox.Size().Y-1 {
		g.Over = true
		return
	}
	for _, p := range g.Snake.Body {
		if p == newHead {
			g.Over = true
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

func (g *Game) HandleInput() {
	switch ev := termbox.PollEvent(); ev.Type {
	case termbox.EventKey:
		switch ev.Key {
		case termbox.KeyArrowUp:
			if g.Snake.Dir.Y == 0 {
				g.Snake.Dir = Point{0, -1}
			}
		case termbox.KeyArrowDown:
			if g.Snake.Dir.Y == 0 {
				g.Snake.Dir = Point{0, 1}
			}
		case termbox.KeyArrowLeft:
			if g.Snake.Dir.X == 0 {
				g.Snake.Dir = Point{-1, 0}
			}
		case termbox.KeyArrowRight:
			if g.Snake.Dir.X == 0 {
				g.Snake.Dir = Point{1, 0}
			}
		case termbox.KeyEsc:
			g.Over = true
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := Game{}
	game.Init()
	defer termbox.Close()

	go func() {
		for !game.Over {
			game.HandleInput()
		}
	}()

	for !game.Over {
		game.Update()
		game.Draw()
		time.Sleep(game.Snake.Speed)
	}

	fmt.Printf("Game Over! Your score: %d\n", game.Score)
}
