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
}

func main() {
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]
	fmt.Println(quote)
}
