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
	"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
	"It's not whether you get knocked down, it's whether you get up. - Vince Lombardi",
	"Do not watch the clock. Do what it does. Keep going. - Sam Levenson",
	"The only way to do great work is to love what you do. - Steve Jobs",
	"If you can dream it, you can do it. - Walt Disney",
	"Success is not the key to happiness. Happiness is the key to success. If you love what you are doing, you will be successful. - Albert Schweitzer",
	"Believe you can and you're halfway there. - Theodore Roosevelt",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
