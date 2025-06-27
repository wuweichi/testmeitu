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
		"You must be the change you wish to see in the world. - Mahatma Gandhi",
		"The future belongs to those who believe in the beauty of their dreams. - Eleanor Roosevelt",
		"Do not dwell in the past, do not dream of the future, concentrate the mind on the present moment. - Buddha",
		"Success is not final, failure is not fatal: It is the courage to continue that counts. - Winston Churchill",
		"In the middle of difficulty lies opportunity. - Albert Einstein",
		"The only way to do great work is to love what you do. - Steve Jobs",
		"It is never too late to be what you might have been. - George Eliot",
		"Happiness is not something ready made. It comes from your own actions. - Dalai Lama",
	}

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(quotes))
	fmt.Println(quotes[randomIndex])
}
