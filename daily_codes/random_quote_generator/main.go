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
	"Whoever is happy will make others happy too. - Anne Frank",
	"Do not go where the path may lead, go instead where there is no path and leave a trail. - Ralph Waldo Emerson",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
