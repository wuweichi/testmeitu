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

	// Create a channel to listen for interrupt signals
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	// Start a goroutine to handle the clock
	go func() {
		for {
			// Get the current time
			now := time.Now()
			hour, min, sec := now.Hour(), now.Minute(), now.Second()

			// Generate a random color for the clock
			color := fmt.Sprintf("\033[38;5;%dm", rand.Intn(256))
			reset := "\033[0m"

			// Print the time with the random color
			fmt.Printf("%s%02d:%02d:%02d%s\r", color, hour, min, sec, reset)

			// Wait for one second
			time.Sleep(time.Second)
		}
	}()

	// Wait for an interrupt signal
	<-interrupt
	fmt.Println("\nClock stopped.")
}
