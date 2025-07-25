package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quotes = []string{
	"The only limit to our realization of tomorrow is our doubts of today. - Franklin D. Roosevelt",
	"Life is what happens when you're busy making other plans. - John Lennon",
	"It is during our darkest moments that we must focus to see the light. - Aristotle",
	"Do not go where the path may lead, go instead where there is no path and leave a trail. - Ralph Waldo Emerson",
	"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
	"You must be the change you wish to see in the world. - Mahatma Gandhi",
	"In the end, it's not the years in your life that count. It's the life in your years. - Abraham Lincoln",
	"You only live once, but if you do it right, once is enough. - Mae West",
	"The greatest glory in living lies not in never falling, but in rising every time we fall. - Nelson Mandela",
	"Life is either a daring adventure or nothing at all. - Helen Keller",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
