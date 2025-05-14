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

	// Start the clock
	fmt.Println("Starting the complex clock...")
	clock := time.NewTicker(1 * time.Second)
	defer clock.Stop()

	// Run the clock
	for {
		select {
		case <-clock.C:
			hour, min, sec := time.Now().Clock()
			fmt.Printf("Current time: %02d:%02d:%02d\n", hour, min, sec)
			// Generate a random number and print it
			randomNum := rand.Intn(100)
			fmt.Printf("Random number: %d\n", randomNum)
		case <-interrupt:
			fmt.Println("\nReceived interrupt signal. Stopping the clock...")
			return
		}
	}
}
