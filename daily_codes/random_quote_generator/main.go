package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quotes = []string{
	"The only limit to our realization of tomorrow is our doubts of today. - Franklin D. Roosevelt",
	"Life is what happens when you're busy making other plans. - John Lennon",
	"In the end, it's not the years in your life that count. It's the life in your years. - Abraham Lincoln",
	"You must be the change you wish to see in the world. - Mahatma Gandhi",
	"Spread love everywhere you go. Let no one ever come to you without leaving happier. - Mother Teresa",
	"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
	"Tell me and I forget. Teach me and I remember. Involve me and I learn. - Benjamin Franklin",
	"Do not go where the path may lead, go instead where there is no path and leave a trail. - Ralph Waldo Emerson",
	"It is during our darkest moments that we must focus to see the light. - Aristotle",
	"Whoever is happy will make others happy too. - Anne Frank",
	"Do not dwell in the past, do not dream of the future, concentrate the mind on the present moment. - Buddha",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
