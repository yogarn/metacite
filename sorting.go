package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func sortCitations(citations []Metadata) {
	sort.Slice(citations, func(i, j int) bool {
		keyI := getSortingKey(citations[i])
		keyJ := getSortingKey(citations[j])

		return keyI < keyJ
	})
}

func getSortingKey(metadata Metadata) string {
	if metadata.Author != "" {
		return extractLastName(metadata.Author)
	}
	return metadata.Title
}

func extractLastName(author string) string {
	parts := strings.Fields(author)
	if len(parts) == 0 {
		return ""
	}
	return parts[len(parts)-1]
}

func showCitations(citations []Metadata, mode string) {
	for _, metadata := range citations {
		var citation string
		if mode == "apa" {
			citation = generateAPACitation(metadata)
		} else {
			fmt.Println("Error: Unsupported citation mode.")
			os.Exit(1)
		}
		fmt.Println(citation)
	}
}
