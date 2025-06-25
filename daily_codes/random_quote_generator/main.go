package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quotes = []string{
	"The only limit to our realization of tomorrow is our doubts of today. - Franklin D. Roosevelt",
	"Life is what happens when you're busy making other plans. - John Lennon",
	"The way to get started is to quit talking and begin doing. - Walt Disney",
	"It's not whether you get knocked down, it's whether you get up. - Vince Lombardi",
	"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
	"Do not watch the clock. Do what it does. Keep going. - Sam Levenson",
	"You must be the change you wish to see in the world. - Mahatma Gandhi",
	"Whatever you are, be a good one. - Abraham Lincoln",
	"The only way to do great work is to love what you do. - Steve Jobs",
	"If you can dream it, you can do it. - Walt Disney",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomQuote := quotes[rand.Intn(len(quotes))]
	fmt.Println(randomQuote)
}
