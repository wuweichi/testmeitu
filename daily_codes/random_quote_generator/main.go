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
	"Your time is limited, so don't waste it living someone else's life. - Steve Jobs",
	"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
	"Do not watch the clock. Do what it does. Keep going. - Sam Levenson",
	"You must be the change you wish to see in the world. - Mahatma Gandhi",
	"Whatever you are, be a good one. - Abraham Lincoln",
	"The best way to predict the future is to invent it. - Alan Kay",
	"It does not matter how slowly you go as long as you do not stop. - Confucius",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
