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
	hours, minutes, seconds int
}

// NewClock creates a new Clock instance.
func NewClock(h, m, s int) *Clock {
	return &Clock{hours: h, minutes: m, seconds: s}
}

// Display shows the current time of the clock.
func (c *Clock) Display() {
	fmt.Printf("%02d:%02d:%02d\n", c.hours, c.minutes, c.seconds)
}

// Tick advances the clock by one second.
func (c *Clock) Tick() {
	c.seconds++
	if c.seconds >= 60 {
		c.seconds = 0
		c.minutes++
		if c.minutes >= 60 {
			c.minutes = 0
			c.hours++
			if c.hours >= 24 {
				c.hours = 0
			}
		}
	}
}

// RandomClock generates a new Clock with random time.
func RandomClock() *Clock {
	rand.Seed(time.Now().UnixNano())
	return NewClock(rand.Intn(24), rand.Intn(60), rand.Intn(60)
}

func main() {
	// Create a random clock
	clock := RandomClock()
	// Setup signal handling for graceful shutdown
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	done := make(chan bool, 1)
	// Start a goroutine to handle the clock ticking
	go func() {
		for {
			select {
			case <-time.After(1 * time.Second):
				clock.Tick()
				clock.Display()
			case <-done:
				return
			}
		}
	}()
	// Wait for shutdown signal
	<-sigs
	done <- true
	fmt.Println("Clock stopped.")
}
