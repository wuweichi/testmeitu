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
		"Life is either a daring adventure or nothing at all. - Helen Keller",
		"You will face many defeats in life, but never let yourself be defeated. - Maya Angelou",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	fmt.Println(quotes[randomIndex])
}
