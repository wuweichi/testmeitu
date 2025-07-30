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
		"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
		"It is during our darkest moments that we must focus to see the light. - Aristotle",
		"Do not go where the path may lead, go instead where there is no path and leave a trail. - Ralph Waldo Emerson",
		"You will face many defeats in life, but never let yourself be defeated. - Maya Angelou",
		"The greatest glory in living lies not in never falling, but in rising every time we fall. - Nelson Mandela",
		"In the end, it's not the years in your life that count. It's the life in your years. - Abraham Lincoln",
		"Life is either a daring adventure or nothing at all. - Helen Keller",
		"Many of life's failures are people who did not realize how close they were to success when they gave up. - Thomas A. Edison",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	fmt.Println(quotes[randomIndex])
}
