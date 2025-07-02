package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"math/rand"
	"time"
)

type Point struct {
	X int
	Y int
}

type Snake struct {
	Body  []Point
	Dir   Point
	Speed int
}

type Game struct {
	Snake Snake
	Food  Point
	Score int
	Over  bool
}

func (g *Game) Init() {
	g.Snake = Snake{Body: []Point{{X: 10, Y: 10}}, Dir: Point{X: 0, Y: 1}, Speed: 1}
	g.SpawnFood()
	g.Score = 0
	g.Over = false
}

func (g *Game) SpawnFood() {
	rand.Seed(time.Now().UnixNano())
	g.Food = Point{X: rand.Intn(20), Y: rand.Intn(20)}
}

func (g *Game) Update() {
	if g.Over {
		return
	}

	head := g.Snake.Body[0]
	newHead := Point{X: head.X + g.Snake.Dir.X, Y: head.Y + g.Snake.Dir.Y}

	if newHead.X < 0 || newHead.X >= 20 || newHead.Y < 0 || newHead.Y >= 20 {
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

func (g *Game) Draw() {
	fmt.Printf("Score: %d\n", g.Score)
	for y := 0; y < 20; y++ {
		for x := 0; x < 20; x++ {
			cell := ' '
			for _, p := range g.Snake.Body {
				if p.X == x && p.Y == y {
					cell = '#'
					break
				}
			}
			if g.Food.X == x && g.Food.Y == y {
				cell = '*'
			}
			fmt.Printf("%c", cell)
		}
		fmt.Println()
	}
}

func main() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	game := Game{}
	game.Init()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	for !game.Over {
		select {
		case ev := <-eventQueue:
			if ev.Type == termbox.EventKey {
				switch ev.Key {
				case termbox.KeyArrowUp:
					game.Snake.Dir = Point{X: 0, Y: -1}
				case termbox.KeyArrowDown:
					game.Snake.Dir = Point{X: 0, Y: 1}
				case termbox.KeyArrowLeft:
					game.Snake.Dir = Point{X: -1, Y: 0}
				case termbox.KeyArrowRight:
					game.Snake.Dir = Point{X: 1, Y: 0}
				}
			}
		default:
			game.Update()
			termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)
			game.Draw()
			termbox.Flush()
			time.Sleep(time.Second / time.Duration(game.Snake.Speed))
		}
	}

	fmt.Println("Game Over!")
}