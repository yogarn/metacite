package main

import (
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func italic(word string) string {
	return "\x1b[3m" + word + "\x1b[0m"
}

func capitalizeFirstChar(words string) string {
	title := cases.Title(language.English)
	return title.String(words)
}
