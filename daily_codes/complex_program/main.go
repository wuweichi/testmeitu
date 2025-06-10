package main

import (
	"fmt"
	"math/rand"
	"time"
	"sync"
)

func generateRandomNumbers(count int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	rand.Seed(time.Now().UnixNano()
	for i := 0; i < count; i++ {
		ch <- rand.Intn(100)
	}
}

func processNumbers(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Printf("Processed number: %d\n", num*2)
	}
}

func main() {
	const numCount = 1000
	var wg sync.WaitGroup
	ch := make(chan int, numCount)

	wg.Add(1)
	go generateRandomNumbers(numCount, &wg, ch)

	wg.Add(1)
	go processNumbers(ch, &wg)

	wg.Wait()
	close(ch)
	fmt.Println("All numbers processed.")
}
