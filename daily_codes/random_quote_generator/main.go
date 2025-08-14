package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	quotes := []string{
		"The only limit to our realization of tomorrow is our doubts of today. - Franklin D. Roosevelt",
		"Life is what happens when you're busy making other plans. - John Lennon",
		"The way to get started is to quit talking and begin doing. - Walt Disney",
		"Your time is limited, don't waste it living someone else's life. - Steve Jobs",
		"It is during our darkest moments that we must focus to see the light. - Aristotle",
		"Whoever is happy will make others happy too. - Anne Frank",
		"Do not go where the path may lead, go instead where there is no path and leave a trail. - Ralph Waldo Emerson",
	}

	rand.Seed(time.Now().UnixNano())
	randomQuote := quotes[rand.Intn(len(quotes))]
	fmt.Println(randomQuote)
}
