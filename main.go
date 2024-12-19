package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: metacite <url>")
		os.Exit(1)
	}

	targetURL := os.Args[1]
	htmlContent, err := fetchPage(targetURL)
	if err != nil {
		fmt.Println("Error fetching page:", err)
		os.Exit(1)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlContent))
	if err != nil {
		fmt.Println("Error parsing HTML:", err)
		os.Exit(1)
	}

	metadata := extractMetadata(doc, targetURL)
	citation := generateAPACitation(metadata)
	fmt.Println(citation)
}
