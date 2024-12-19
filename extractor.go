package main

import (
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Metadata struct {
	Title        string
	Author       string
	Date         string
	SiteName     string
	CanonicalURL string
}

func fetchPage(targetURL string) (string, error) {
	resp, err := http.Get(targetURL)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func extractAuthor(doc *goquery.Document) string {
	authorFirstName := doc.Find("meta[property='article:author:first_name']").AttrOr("content", "")
	authorLastName := doc.Find("meta[property='article:author:last_name']").AttrOr("content", "")
	if authorFirstName != "" && authorLastName != "" {
		return authorFirstName + " " + authorLastName
	}

	author := doc.Find("meta[name='author']").AttrOr("content", "")
	if author != "" {
		return author
	}

	author = doc.Find("meta[property='article:author']").AttrOr("content", "")
	if author != "" {
		return author
	}

	return ""
}

func extractMetadata(doc *goquery.Document, targetURL string) Metadata {
	author := capitalizeFirstChar(extractAuthor(doc))
	ogTitle := doc.Find("meta[property='og:title']").AttrOr("content", "Untitled")
	ogSiteName := doc.Find("meta[property='og:site_name']").AttrOr("content", "")
	articleDate := doc.Find("meta[property='article:published_time']").AttrOr("content", "n.d.")

	if ogSiteName == "" {
		parsedURL, err := url.Parse(targetURL)
		if err == nil {
			ogSiteName = strings.TrimPrefix(parsedURL.Hostname(), "www.")
		} else {
			ogSiteName = "Unknown Site"
		}
	}

	canonicalURL := doc.Find("meta[property='og:url']").AttrOr("content", targetURL)

	return Metadata{
		Title:        ogTitle,
		Author:       author,
		Date:         articleDate,
		SiteName:     ogSiteName,
		CanonicalURL: canonicalURL,
	}
}
