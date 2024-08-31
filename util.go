package main

import (
	"fmt"
	"sort"
)

type reportPage struct {
	count int
	url   string
}

func (cfg *config) addPageVisit(normalizedURL string) (isFirstTime bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	if _, visited := cfg.pages[normalizedURL]; visited {
		cfg.pages[normalizedURL] += 1
		return false
	}

	cfg.pages[normalizedURL] = 1
	return true
}

func (cfg *config) pagesLen() int {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	return len(cfg.pages)
}

func printReport(pages map[string]int, baseURL string) {
	sortedResults := sortPages(pages)

	fmt.Printf(`
=============================
  REPORT for %s
=============================
`, baseURL)

	for _, report := range sortedResults {
		fmt.Printf("Found %d internal links to %s\n", report.count, report.url)
	}

}

func sortPages(pages map[string]int) []reportPage {
	results := make([]reportPage, 0, len(pages))

	for url, count := range pages {
		results = append(results, reportPage{count: count, url: url})
	}

	sort.Slice(results, func(i, j int) bool {
		if results[i].count == results[j].count {
			return results[i].url < results[j].url
		}
		return results[i].count > results[j].count
	})

	return results
}
