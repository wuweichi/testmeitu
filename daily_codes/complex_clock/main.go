package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

type Clock struct {
	hours, minutes, seconds int
}

func (c *Clock) Update() {
	now := time.Now()
	c.hours = now.Hour()
	c.minutes = now.Minute()
	c.seconds = now.Second()
}

func (c *Clock) Display() {
	fmt.Printf("\r%02d:%02d:%02d", c.hours, c.minutes, c.seconds)
}

func generateRandomColor() string {
	colors := []string{
		"\033[31m", // Red
		"\033[32m", // Green
		"\033[33m", // Yellow
		"\033[34m", // Blue
		"\033[35m", // Magenta
		"\033[36m", // Cyan
	}
	return colors[rand.Intn(len(colors))]
}

func main() {
	clock := Clock{}
	rand.Seed(time.Now().UnixNano())

	// Handle interrupt signal for graceful shutdown
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interrupt
		fmt.Println("\nExiting...")
		os.Exit(0)
	}()

	for {
		clock.Update()
		color := generateRandomColor()
		fmt.Print(color)
		clock.Display()
		time.Sleep(1 * time.Second)
	}
}
