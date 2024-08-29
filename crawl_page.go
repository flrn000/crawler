package main

import (
	"fmt"
	"net/url"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) map[string]int {
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html:%v", err)
		return pages
	}

	urlsToCrawl, err := getURLsFromHTML(htmlBody, rawBaseURL)
	if err != nil {
		fmt.Printf("error getting urls:%v", err)
		return pages
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Println(err)
		return pages
	}

	for _, currURL := range urlsToCrawl {
		currentURL, err := url.Parse(currURL)
		if err != nil {
			fmt.Println(err)
			return pages
		}

		if currentURL.Hostname() != baseURL.Hostname() {
			return pages
		}

		normalizedURL, err := normalizeURL(currURL)
		if err != nil {
			fmt.Println(err)
			return pages
		}

		if _, ok := pages[normalizedURL]; ok {
			pages[normalizedURL] += 1
			continue
		} else {
			pages[normalizedURL] = 1
		}

		crawlPage(rawBaseURL, currURL, pages)
	}

	return pages
}
