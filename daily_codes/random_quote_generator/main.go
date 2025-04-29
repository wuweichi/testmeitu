package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Quote struct {
	Text   string
	Author string
}

func main() {
	quotes := []Quote{
		{"The only limit to our realization of tomorrow is our doubts of today.", "Franklin D. Roosevelt"},
		{"Do not watch the clock. Do what it does. Keep going.", "Sam Levenson"},
		{"The way to get started is to quit talking and begin doing.", "Walt Disney"},
		{"Life is what happens when you're busy making other plans.", "John Lennon"},
		{"The future belongs to those who believe in the beauty of their dreams.", "Eleanor Roosevelt"},
		{"You must be the change you wish to see in the world.", "Mahatma Gandhi"},
		{"Spread love everywhere you go. Let no one ever come to you without leaving happier.", "Mother Teresa"},
		{"The only thing we have to fear is fear itself.", "Franklin D. Roosevelt"},
		{"It is during our darkest moments that we must focus to see the light.", "Aristotle"},
		{"Whoever is happy will make others happy too.", "Anne Frank"},
	}

	rand.Seed(time.Now().UnixNano())
	randomQuote := quotes[rand.Intn(len(quotes))]

	fmt.Printf("\"%s\" - %s\n", randomQuote.Text, randomQuote.Author)
}
