package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quotes := []string{
		"The only limit to our realization of tomorrow is our doubts of today. - Franklin D. Roosevelt",
		"Life is what happens when you're busy making other plans. - John Lennon",
		"The way to get started is to quit talking and begin doing. - Walt Disney",
		"Whoever is happy will make others happy too. - Anne Frank",
		"Do not watch the clock. Do what it does. Keep going. - Sam Levenson",
		"You only live once, but if you do it right, once is enough. - Mae West",
		"Never let the fear of striking out keep you from playing the game. - Babe Ruth",
		"Life is either a daring adventure or nothing at all. - Helen Keller",
		"Many of life's failures are people who did not realize how close they were to success when they gave up. - Thomas A. Edison",
		"You miss 100% of the shots you don't take. - Wayne Gretzky",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	fmt.Println(quotes[randomIndex])
}
