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
	hour, minute, second int
}

func (c *Clock) Update() {
	now := time.Now()
	c.hour = now.Hour()
	c.minute = now.Minute()
	c.second = now.Second()
}

func (c Clock) Display() {
	fmt.Printf("\r%02d:%02d:%02d", c.hour, c.minute, c.second)
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
	rand.Seed(time.Now().UnixNano())
	clock := Clock{}
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("Press Ctrl+C to stop the clock")
	fmt.Println("The clock will change color every second:")

	for {
		select {
		case <-stop:
			fmt.Println("\nClock stopped.")
			return
		default:
			clock.Update()
			color := generateRandomColor()
			fmt.Print(color)
			clock.Display()
			time.Sleep(1 * time.Second)
		}
	}
}
