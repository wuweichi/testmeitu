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
			fmt.Println("Why did the programmer quit his job? Because he didn't get arrays.")
		case 1:
			fmt.Println("Why do programmers always mix up Halloween and Christmas? Because Oct 31 == Dec 25!")
		case 2:
			fmt.Println("How many programmers does it take to change a light bulb? None, that's a hardware problem.")
		case 3:
			fmt.Println("Why do Java developers wear glasses? Because they can't C#.")
		case 4:
			fmt.Println("Why was the JavaScript developer sad? Because he didn't know how to 'null' his feelings.")
		}
	}
}
