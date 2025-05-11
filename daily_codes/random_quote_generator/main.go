package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quotes = []string{
	"The only limit to our realization of tomorrow is our doubts of today. - Franklin D. Roosevelt",
	"Life is what happens when you're busy making other plans. - John Lennon",
	"You must be the change you wish to see in the world. - Mahatma Gandhi",
	"In the end, it's not the years in your life that count. It's the life in your years. - Abraham Lincoln",
	"Do not go where the path may lead, go instead where there is no path and leave a trail. - Ralph Waldo Emerson",
	"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
	"To live is the rarest thing in the world. Most people exist, that is all. - Oscar Wilde",
	"That which does not kill us makes us stronger. - Friedrich Nietzsche",
	"It is never too late to be what you might have been. - George Eliot",
	"Strive not to be a success, but rather to be of value. - Albert Einstein",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
