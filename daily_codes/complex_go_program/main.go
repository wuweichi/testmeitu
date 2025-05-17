package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ComplexStruct represents a complex data structure
type ComplexStruct struct {
	ID          int
	Name        string
	Description string
	Values      []int
	Timestamp   time.Time
}

// generateRandomComplexStruct generates a ComplexStruct with random values
func generateRandomComplexStruct() ComplexStruct {
	rand.Seed(time.Now().UnixNano())
	values := make([]int, rand.Intn(10)+5)
	for i := range values {
		values[i] = rand.Intn(100)
	}
	return ComplexStruct{
		ID:          rand.Intn(1000),
		Name:        fmt.Sprintf("Name%d", rand.Intn(100)),
		Description: fmt.Sprintf("Description%d", rand.Intn(100)),
		Values:      values,
		Timestamp:   time.Now(),
	}
}

// printComplexStruct prints the details of a ComplexStruct
func printComplexStruct(cs ComplexStruct) {
	fmt.Printf("ID: %d\n", cs.ID)
	fmt.Printf("Name: %s\n", cs.Name)
	fmt.Printf("Description: %s\n", cs.Description)
	fmt.Printf("Values: %v\n", cs.Values)
	fmt.Printf("Timestamp: %s\n", cs.Timestamp.Format(time.RFC3339))
}

func main() {
	// Generate and print 100 random ComplexStructs
	for i := 0; i < 100; i++ {
		cs := generateRandomComplexStruct()
		printComplexStruct(cs)
		fmt.Println("----------------------------------------")
	}
}
