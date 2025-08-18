package main

import (
	"fmt"
	"time"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Channel to listen for interrupt signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Infinite loop to keep the clock running
	for {
		select {
		case <-time.After(1 * time.Second):
			// Get the current time
			now := time.Now()
			hour, min, sec := now.Hour(), now.Minute(), now.Second()

			// Generate a random color for the clock display
			color := fmt.Sprintf("\033[38;5;%dm", rand.Intn(256))
			reset := "\033[0m"

			// Display the time with the random color
			fmt.Printf("%s%02d:%02d:%02d%s\n", color, hour, min, sec, reset)

		case <-interrupt:
			// Handle interrupt signal
			fmt.Println("\nClock stopped.")
			return
		}
	}
}
