package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quotes = []string{
	"The only way to do great work is to love what you do. - Steve Jobs",
	"Life is what happens when you're busy making other plans. - John Lennon",
	"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
	"It is during our darkest moments that we must focus to see the light. - Aristotle",
	"Whoever is happy will make others happy too. - Anne Frank",
	"Do not go where the path may lead, go instead where there is no path and leave a trail. - Ralph Waldo Emerson",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	quote := quotes[rand.Intn(len(quotes))]
	fmt.Println(quote)
}
