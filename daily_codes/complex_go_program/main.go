package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Define a struct for a person
type Person struct {
	Name string
	Age  int
}

// Function to generate a random person
func generateRandomPerson() Person {
	names := []string{"Alice", "Bob", "Charlie", "Diana", "Eve", "Frank", "Grace", "Hank"}
	return Person{
		Name: names[rand.Intn(len(names))],
		Age:  rand.Intn(100) + 1,
	}
}

// Function to print a person's details
func printPerson(p Person) {
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
}

// Main function
func main() {
	rand.Seed(time.Now().UnixNano())
	
	// Generate and print 1000 random people
	for i := 0; i < 1000; i++ {
		person := generateRandomPerson()
		printPerson(person)
	}
	
	fmt.Println("Generated 1000 random people!")
}
