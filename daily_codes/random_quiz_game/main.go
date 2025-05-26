package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Question struct {
	Text   string
	Answer string
}

func main() {
	questions := []Question{
		{"What is the capital of France?", "Paris"},
		{"Which planet is known as the Red Planet?", "Mars"},
		{"What is the largest ocean on Earth?", "Pacific"},
		{"Who wrote 'Romeo and Juliet'?", "Shakespeare"},
		{"What is the chemical symbol for gold?", "Au"},
		{"Which country is home to the kangaroo?", "Australia"},
		{"What is the square root of 64?", "8"},
		{"In which year did the Titanic sink?", "1912"},
		{"What is the main ingredient in guacamole?", "Avocado"},
		{"Who painted the Mona Lisa?", "Da Vinci"},
	}

	rand.Seed(time.Now().UnixNano())
	score := 0
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to the Random Quiz Game!")
	fmt.Println("You will be asked 5 random questions. Let's see how many you can answer correctly!")

	for i := 0; i < 5; i++ {
		index := rand.Intn(len(questions))
		question := questions[index]
		questions = append(questions[:index], questions[index+1:]...)

		fmt.Printf("Question %d: %s\n", i+1, question.Text)
		fmt.Print("Your answer: ")
		answer, _ := reader.ReadString('\n')
		answer = strings.TrimSpace(answer)

		if strings.EqualFold(answer, question.Answer) {
			fmt.Println("Correct!")
			score++
		} else {
			fmt.Printf("Wrong! The correct answer is %s.\n", question.Answer)
		}
	}

	fmt.Printf("Your final score is %d out of 5.\n", score)
}
