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
	"It is during our darkest moments that we must focus to see the light. - Aristotle",
	"Do not go where the path may lead, go instead where there is no path and leave a trail. - Ralph Waldo Emerson",
	"You will face many defeats in life, but never let yourself be defeated. - Maya Angelou",
	"The greatest glory in living lies not in never falling, but in rising every time we fall. - Nelson Mandela",
	"In the end, it's not the years in your life that count. It's the life in your years. - Abraham Lincoln",
	"Never let the fear of striking out keep you from playing the game. - Babe Ruth",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]
	fmt.Println(quote)
}
