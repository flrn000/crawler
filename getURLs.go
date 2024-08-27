package main

import (
	"fmt"
	"net/url"
	"strings"

	"golang.org/x/net/html"
)

func findAnchorElements(n *html.Node, results *[]string, baseURL *url.URL) error {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				hrefURL, err := url.Parse(attr.Val)
				if err != nil {
					return fmt.Errorf("error parsing href: %w", err)
				}

				if hrefURL.Path != "" {
					*results = append(*results, baseURL.ResolveReference(hrefURL).String())
				} else {
					break
				}
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		findAnchorElements(c, results, baseURL)
	}

	return nil
}

func getURLsFromHTML(htmlBody, rawBaseURL string) ([]string, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return []string{}, fmt.Errorf("error parsing url: %w", err)
	}
	r := strings.NewReader(htmlBody)
	doc, err := html.Parse(r)
	if err != nil {
		return []string{}, err
	}

	var results []string
	if err := findAnchorElements(doc, &results, baseURL); err != nil {
		return results, err
	}

	return results, nil
}
