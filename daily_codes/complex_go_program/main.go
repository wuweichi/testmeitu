package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ComplexStruct represents a complex data structure
type ComplexStruct struct {
	ID      int
	Name    string
	Details []Detail
}

// Detail represents additional details
type Detail struct {
	Key   string
	Value string
}

func generateRandomComplexStruct() ComplexStruct {
	rand.Seed(time.Now().UnixNano())
	details := make([]Detail, rand.Intn(5)+1)
	for i := range details {
		details[i] = Detail{
			Key:   fmt.Sprintf("Key%d", i),
			Value: fmt.Sprintf("Value%d", rand.Intn(100)),
		}
	}
	return ComplexStruct{
		ID:      rand.Intn(1000),
		Name:    fmt.Sprintf("Name%d", rand.Intn(100)),
		Details: details,
	}
}

func printComplexStruct(cs ComplexStruct) {
	fmt.Printf("ID: %d\n", cs.ID)
	fmt.Printf("Name: %s\n", cs.Name)
	fmt.Println("Details:")
	for _, detail := range cs.Details {
		fmt.Printf("\t%s: %s\n", detail.Key, detail.Value)
	}
}

func main() {
	for i := 0; i < 100; i++ {
		cs := generateRandomComplexStruct()
		printComplexStruct(cs)
		fmt.Println("-----------------")
	}
}
