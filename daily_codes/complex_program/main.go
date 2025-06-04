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

// Detail represents details of the ComplexStruct
type Detail struct {
	Field1 string
	Field2 int
	Field3 bool
}

func generateRandomDetail() Detail {
	rand.Seed(time.Now().UnixNano())
	return Detail{
		Field1: fmt.Sprintf("Field1-%d", rand.Intn(100)),
		Field2: rand.Intn(1000),
		Field3: rand.Intn(2) == 1,
	}
}

func generateComplexStruct(id int, name string, detailCount int) ComplexStruct {
	details := make([]Detail, detailCount)
	for i := range details {
		details[i] = generateRandomDetail()
	}
	return ComplexStruct{
		ID:      id,
		Name:    name,
		Details: details,
	}
}

func printComplexStruct(cs ComplexStruct) {
	fmt.Printf("ID: %d, Name: %s\n", cs.ID, cs.Name)
	for i, detail := range cs.Details {
		fmt.Printf("\tDetail %d: Field1=%s, Field2=%d, Field3=%t\n", i+1, detail.Field1, detail.Field2, detail.Field3)
	}
}

func main() {
	// Generate a slice of ComplexStruct
	var complexStructs []ComplexStruct
	for i := 1; i <= 100; i++ {
		cs := generateComplexStruct(i, fmt.Sprintf("Name-%d", i), 10)
		complexStructs = append(complexStructs, cs)
	}

	// Print all ComplexStruct
	for _, cs := range complexStructs {
		printComplexStruct(cs)
	}

	fmt.Println("Program completed successfully.")
}
