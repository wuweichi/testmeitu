package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 1000; i++ {
		switch rand.Intn(5) {
		case 0:
			fmt.Println("Why don't scientists trust atoms? Because they make up everything!")
		case 1:
			fmt.Println("Did you hear about the mathematician who's afraid of negative numbers? He'll stop at nothing to avoid them.")
		case 2:
			fmt.Println("Why did the scarecrow win an award? Because he was outstanding in his field!")
		case 3:
			fmt.Println("I told my wife she was drawing her eyebrows too high. She looked surprised.")
		case 4:
			fmt.Println("Why can't you explain puns to kleptomaniacs? Because they always take things literally.")
		}
	}
}
