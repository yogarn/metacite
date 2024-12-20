package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	urlFlag := flag.String("l", "", "The target URL")
	modeFlag := flag.String("m", "apa", "The mode to use (default, other modes)")
	flag.Parse()

	supportedModes := map[string]bool{
		"apa": true,
	}

	if _, ok := supportedModes[*modeFlag]; !ok {
		fmt.Println("Error: Unsupported mode. Supported modes are:")
		for mode := range supportedModes {
			fmt.Println("  -", mode)
			os.Exit(1)
		}
	}

	if *urlFlag == "" {
		fmt.Println("Error: URL is required. Use -l <url>")
		os.Exit(1)
	}
	targetURL := *urlFlag

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

	var citation string
	if *modeFlag == "apa" {
		citation = generateAPACitation(metadata)
	}

	fmt.Println(citation)
}
