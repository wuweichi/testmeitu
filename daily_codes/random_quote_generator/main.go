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
		"It's not whether you get knocked down, it's whether you get up. - Vince Lombardi",
		"Do not watch the clock. Do what it does. Keep going. - Sam Levenson",
		"The only way to do great work is to love what you do. - Steve Jobs",
		"If you can dream it, you can do it. - Walt Disney",
		"The best revenge is massive success. - Frank Sinatra",
		"You miss 100% of the shots you don't take. - Wayne Gretzky",
		"I have not failed. I've just found 10,000 ways that won't work. - Thomas Edison",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	fmt.Println(quotes[randomIndex])
}
