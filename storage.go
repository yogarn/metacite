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
	encoder.SetIndent("", "  ")
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

func removeCitation(fileName string, canonicalUrl string) {
	citations := loadCitations(fileName)

	updatedCitations := []Metadata{}
	for _, citation := range citations {
		if citation.CanonicalURL != canonicalUrl {
			updatedCitations = append(updatedCitations, citation)
		}
	}

	if len(citations) == len(updatedCitations) {
		fmt.Println("No matching citation found to remove.")
		return
	}

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(updatedCitations)
	if err != nil {
		fmt.Println("Error saving to JSON:", err)
		os.Exit(1)
	}

	fmt.Println("Citation removed successfully!")
}
