package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	ID     int
	Health int
	Score  int
}

type Game struct {
	Players []*Player
	Mutex   sync.Mutex
}

func (g *Game) AddPlayer(id int) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	player := &Player{ID: id, Health: 100, Score: 0}
	g.Players = append(g.Players, player)
	fmt.Printf("Player %d joined the game.\n", id)
}

func (g *Game) RemovePlayer(id int) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	for i, p := range g.Players {
		if p.ID == id {
			g.Players = append(g.Players[:i], g.Players[i+1:]...)
			fmt.Printf("Player %d left the game.\n", id)
			return
		}
	}
}

func (g *Game) UpdatePlayer(id int, healthChange int, scoreChange int) {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	for _, p := range g.Players {
		if p.ID == id {
			p.Health += healthChange
			if p.Health < 0 {
				p.Health = 0
			}
			p.Score += scoreChange
			fmt.Printf("Player %d: Health=%d, Score=%d\n", id, p.Health, p.Score)
			return
		}
	}
}

func (g *Game) SimulateAction(playerID int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		action := rand.Intn(3)
		switch action {
		case 0:
			g.UpdatePlayer(playerID, -10, 5)
		case 1:
			g.UpdatePlayer(playerID, 5, 10)
		case 2:
			g.UpdatePlayer(playerID, 0, 15)
		}
	}
}

func (g *Game) PrintStatus() {
	g.Mutex.Lock()
	defer g.Mutex.Unlock()
	fmt.Println("Current game status:")
	for _, p := range g.Players {
		fmt.Printf("Player %d: Health=%d, Score=%d\n", p.ID, p.Health, p.Score)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	game := &Game{}
	var wg sync.WaitGroup

	// Add players
	for i := 1; i <= 10; i++ {
		game.AddPlayer(i)
	}

	// Start simulation for each player
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go game.SimulateAction(i, &wg)
	}

	// Print status periodically
	go func() {
		for {
			time.Sleep(2 * time.Second)
			game.PrintStatus()
		}
	}()

	// Wait for all simulations to finish
	wg.Wait()

	// Final status
	fmt.Println("Simulation ended. Final status:")
	game.PrintStatus()
}
