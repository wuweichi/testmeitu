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
	"If you look at what you have in life, you'll always have more. - Oprah Winfrey",
	"The best and most beautiful things in the world cannot be seen or even touched - they must be felt with the heart. - Helen Keller",
	"Whoever is happy will make others happy too. - Anne Frank",
	"Do not go where the path may lead, go instead where there is no path and leave a trail. - Ralph Waldo Emerson",
	"You will face many defeats in life, but never let yourself be defeated. - Maya Angelou",
	"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
