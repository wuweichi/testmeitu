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
		"\033[35m", // Purple
		"\033[36m", // Cyan
		"\033[37m", // White
	}
	return colors[rand.Intn(len(colors))]
}

func main() {
	rand.Seed(time.Now().UnixNano())
	clock := Clock{}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
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
