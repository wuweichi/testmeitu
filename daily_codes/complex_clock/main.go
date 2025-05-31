package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

// Clock represents a clock with hours, minutes, and seconds.
type Clock struct {
	Hours   int
	Minutes int
	Seconds int
}

// NewClock creates a new Clock instance.
func NewClock(h, m, s int) *Clock {
	return &Clock{Hours: h, Minutes: m, Seconds: s}
}

// Display prints the current time of the clock.
func (c *Clock) Display() {
	fmt.Printf("%02d:%02d:%02d\n", c.Hours, c.Minutes, c.Seconds)
}

// Tick advances the clock by one second.
func (c *Clock) Tick() {
	c.Seconds++
	if c.Seconds >= 60 {
		c.Seconds = 0
		c.Minutes++
		if c.Minutes >= 60 {
			c.Minutes = 0
			c.Hours++
			if c.Hours >= 24 {
				c.Hours = 0
			}
		}
	}
}

// RandomClock generates a new Clock with random time.
func RandomClock() *Clock {
	rand.Seed(time.Now().UnixNano())
	return NewClock(rand.Intn(24), rand.Intn(60), rand.Intn(60))
}

func main() {
	// Create a random clock
	clock := RandomClock()

	// Set up channel to listen for interrupt signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Create a ticker that ticks every second
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	// Run the clock
	for {
		select {
		case <-ticker.C:
			clock.Tick()
			clock.Display()
		case <-interrupt:
			fmt.Println("\nClock stopped.")
			return
		}
	}
}
