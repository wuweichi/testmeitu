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
	"Your time is limited, so don't waste it living someone else's life. - Steve Jobs",
	"If life were predictable it would cease to be life, and be without flavor. - Eleanor Roosevelt",
	"If you look at what you have in life, you'll always have more. - Oprah Winfrey",
	"If you set your goals ridiculously high and it's a failure, you will fail above everyone else's success. - James Cameron",
	"You may say I'm a dreamer, but I'm not the only one. - John Lennon",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(quotes[rand.Intn(len(quotes))])
}
