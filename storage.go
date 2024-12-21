package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func saveCitation(fileName string, metadata Metadata) {
	citations := loadCitations(fileName)
	citations = append(citations, metadata)

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(citations)
	if err != nil {
		fmt.Println("Error saving to JSON:", err)
		os.Exit(1)
	}
}

func loadCitations(fileName string) []Metadata {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return []Metadata{}
		}
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var citations []Metadata
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&citations)
	if err != nil && err.Error() != "EOF" {
		fmt.Println("Error decoding JSON:", err)
		os.Exit(1)
	}

	return citations
}
