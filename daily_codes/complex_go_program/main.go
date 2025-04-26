package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ComplexStruct represents a complex data structure
type ComplexStruct struct {
	ID       int
	Name     string
	Details  []string
	Metadata map[string]interface{}
}

// generateRandomString generates a random string of given length
func generateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// populateComplexStruct creates and populates a ComplexStruct with random data
func populateComplexStruct() ComplexStruct {
	cs := ComplexStruct{
		ID:   rand.Intn(1000),
		Name: generateRandomString(10),
	}

	detailsCount := rand.Intn(5) + 1
	for i := 0; i < detailsCount; i++ {
		cs.Details = append(cs.Details, generateRandomString(20))
	}

	cs.Metadata = make(map[string]interface{})
	metadataCount := rand.Intn(3) + 1
	for i := 0; i < metadataCount; i++ {
		cs.Metadata[generateRandomString(5)] = rand.Intn(100)
	}

	return cs
}

// main is the entry point of the program
func main() {
	rand.Seed(time.Now().UnixNano())

	// Generate a slice of ComplexStruct
	var complexStructs []ComplexStruct
	for i := 0; i < 100; i++ {
		complexStructs = append(complexStructs, populateComplexStruct())
	}

	// Print each ComplexStruct
	for _, cs := range complexStructs {
		fmt.Printf("ID: %d, Name: %s\n", cs.ID, cs.Name)
		fmt.Println("Details:")
		for _, detail := range cs.Details {
			fmt.Printf("\t%s\n", detail)
		}
		fmt.Println("Metadata:")
		for key, value := range cs.Metadata {
			fmt.Printf("\t%s: %v\n", key, value)
		}
		fmt.Println()
	}
}
