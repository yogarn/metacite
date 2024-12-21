package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

const citationFile = "citation.json"

func main() {
	urlFlag := flag.String("l", "", "The target URL")
	modeFlag := flag.String("m", "apa", "The mode to use (default, other modes)")
	actionFlag := flag.String("a", "add", "Action: add, show")
	flag.Parse()

	if *actionFlag == "show" {
		citations := loadCitations("citation.json")
		sortCitations(citations)
		showCitations(citations, *modeFlag)
		return
	}

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

	supportedActions := map[string]bool{
		"show":   true,
		"add":    true,
		"direct": true,
	}

	if _, ok := supportedActions[*actionFlag]; !ok {
		fmt.Println("Error: Unsupported action. Supported actions are:")
		for action := range supportedActions {
			fmt.Println("  -", action)
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

	if *actionFlag == "add" {
		saveCitation("citation.json", metadata)
		fmt.Println("Citation added to", citationFile)
	} else if *actionFlag == "direct" {
		var citation string
		if *modeFlag == "apa" {
			citation = generateAPACitation(metadata)
			fmt.Println(citation)
		}
	}
}
