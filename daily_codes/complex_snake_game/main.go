package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
	"github.com/nsf/termbox-go"
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
}

func (g *Game) Init() {
	g.Snake = Snake{Body: []Point{{X: 10, Y: 10}}, Dir: Point{X: 0, Y: 1}, Speed: 100 * time.Millisecond}
	g.Score = 0
	g.GameOver = false
	g.placeFood()
}

func (g *Game) placeFood() {
	width, height := termbox.Size()
	g.Food = Point{X: rand.Intn(width), Y: rand.Intn(height)}
}

func (g *Game) Update() {
	if g.GameOver {
		return
	}

	head := g.Snake.Body[0]
	newHead := Point{X: head.X + g.Snake.Dir.X, Y: head.Y + g.Snake.Dir.Y}

	if newHead.X < 0 || newHead.Y < 0 || newHead.X >= termbox.Size().X || newHead.Y >= termbox.Size().Y {
		g.GameOver = true
		return
	}

	for _, p := range g.Snake.Body {
		if p == newHead {
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
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	for _, p := range g.Snake.Body {
		termbox.SetCell(p.X, p.Y, 'O', termbox.ColorGreen, termbox.ColorDefault)
	}

	termbox.SetCell(g.Food.X, g.Food.Y, 'X', termbox.ColorRed, termbox.ColorDefault)

	termbox.Flush()
}

func main() {
	err := termbox.Init()
	if err != nil {
		fmt.Println("Failed to initialize termbox:", err)
		os.Exit(1)
	}
	defer termbox.Close()

	rand.Seed(time.Now().UnixNano())

	game := Game{}
	game.Init()

	eventQueue := make(chan termbox.Event)
	go func() {
		for {
			eventQueue <- termbox.PollEvent()
		}
	}()

	gameLoop:
	for {
		select {
		case ev := <-eventQueue:
			switch ev.Type {
			case termbox.EventKey:
				switch ev.Key {
				case termbox.KeyArrowUp:
					game.Snake.Dir = Point{X: 0, Y: -1}
				case termbox.KeyArrowDown:
					game.Snake.Dir = Point{X: 0, Y: 1}
				case termbox.KeyArrowLeft:
					game.Snake.Dir = Point{X: -1, Y: 0}
				case termbox.KeyArrowRight:
					game.Snake.Dir = Point{X: 1, Y: 0}
				case termbox.KeyEsc:
					break gameLoop
				}
			case termbox.EventError:
				panic(ev.Err)
			}
		default:
			game.Update()
			game.Draw()
			time.Sleep(game.Snake.Speed)
		}

		if game.GameOver {
			termbox.Close()
			cmd := exec.Command("clear")
			cmd.Stdout = os.Stdout
			cmd.Run()
			fmt.Printf("Game Over! Score: %d\n", game.Score)
			break
		}
	}
}
