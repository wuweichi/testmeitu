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
	"If you look at what you have in life, you'll always have more. If you look at what you don't have in life, you'll never have enough. - Oprah Winfrey",
	"The best and most beautiful things in the world cannot be seen or even touched - they must be felt with the heart. - Helen Keller",
	"Keep your face always toward the sunshine - and shadows will fall behind you. - Walt Whitman",
	"You will face many defeats in life, but never let yourself be defeated. - Maya Angelou",
	"The only impossible journey is the one you never begin. - Tony Robbins",
	"In the end, it's not the years in your life that count. It's the life in your years. - Abraham Lincoln",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
