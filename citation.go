package main

import (
	"fmt"
	"strings"
	"time"
)

func formatAuthor(author string) string {
	parts := strings.Fields(author)
	if len(parts) == 0 {
		return "No Author"
	}
	if len(parts) == 1 {
		return parts[0]
	}
	lastName := parts[len(parts)-1]
	initials := ""
	for _, part := range parts[:len(parts)-1] {
		initials += string(part[0]) + ". "
	}
	return fmt.Sprintf("%s, %s", lastName, initials)
}

func parseDate(date string) string {
	parsedTime, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return "n.d."
	}
	return parsedTime.Format("2 January 2006")
}

func generateAPACitation(metadata Metadata) string {
	formattedAuthor := formatAuthor(metadata.Author)
	formattedDate := parseDate(metadata.Date)
	if formattedDate == "" {
		formattedDate = "n.d."
	}
	return fmt.Sprintf("%s (%s). %s. %s. %s", formattedAuthor, formattedDate, metadata.Title, italic(metadata.SiteName), metadata.CanonicalURL)
}
