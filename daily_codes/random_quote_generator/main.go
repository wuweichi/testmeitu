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
		"Your time is limited, don't waste it living someone else's life. - Steve Jobs",
		"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
		"Do not watch the clock. Do what it does. Keep going. - Sam Levenson",
		"You must be the change you wish to see in the world. - Mahatma Gandhi",
		"Whatever you are, be a good one. - Abraham Lincoln",
		"The best way to predict the future is to invent it. - Alan Kay",
		"Strive not to be a success, but rather to be of value. - Albert Einstein",
	}

	rand.Seed(time.Now().UnixNano())
	randomQuote := quotes[rand.Intn(len(quotes))]
	fmt.Println(randomQuote)
}
