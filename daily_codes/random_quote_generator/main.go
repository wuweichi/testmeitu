package main

import (
	"fmt"
	"math/rand"
	"time"
)

var quotes = []string{
	"The only limit to our realization of tomorrow is our doubts of today. - Franklin D. Roosevelt",
	"Do not watch the clock. Do what it does. Keep going. - Sam Levenson",
	"The secret of getting ahead is getting started. - Mark Twain",
	"You miss 100% of the shots you don’t take. - Wayne Gretzky",
	"Whether you think you can or you think you can’t, you’re right. - Henry Ford",
	"The best time to plant a tree was 20 years ago. The second best time is now. - Chinese Proverb",
	"It’s not whether you get knocked down, it’s whether you get up. - Vince Lombardi",
	"If you’re going through hell, keep going. - Winston Churchill",
	"People who are crazy enough to think they can change the world, are the ones who do. - Rob Siltanen",
	"We may encounter many defeats but we must not be defeated. - Maya Angelou",
}

func main() {
	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	fmt.Println(quotes[randomIndex])
}
