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
	"Spread love everywhere you go. Let no one ever come to you without leaving happier. - Mother Teresa",
	"When you reach the end of your rope, tie a knot in it and hang on. - Franklin D. Roosevelt",
	"Always remember that you are absolutely unique. Just like everyone else. - Margaret Mead",
	"Don't judge each day by the harvest you reap but by the seeds that you plant. - Robert Louis Stevenson",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
