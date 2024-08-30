package main

import (
	"fmt"
	"net/url"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	htmlBody, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error getting html:%v", err)
		return
	}

	urlsToCrawl, err := getURLsFromHTML(htmlBody, cfg.baseURL)
	if err != nil {
		fmt.Printf("error getting urls:%v", err)
		return
	}
	fmt.Println("crawling: ", rawCurrentURL)

	for _, currURL := range urlsToCrawl {
		currentURL, err := url.Parse(currURL)
		if err != nil {
			fmt.Println(err)
			return
		}

		if currentURL.Hostname() != cfg.baseURL.Hostname() {
			return
		}

		normalizedURL, err := normalizeURL(currURL)
		if err != nil {
			fmt.Println(err)
			return
		}

		cfg.mu.Lock()
		_, visited := cfg.pages[normalizedURL]
		cfg.mu.Unlock()

		if visited {
			cfg.mu.Lock()
			cfg.pages[normalizedURL] += 1
			cfg.mu.Unlock()
			continue
		} else {
			cfg.mu.Lock()
			cfg.pages[normalizedURL] = 1
			cfg.mu.Unlock()
		}

		cfg.wg.Add(1)
		go func() {
			cfg.concurrencyControl <- struct{}{}
			defer func() {
				<-cfg.concurrencyControl
				cfg.wg.Done()
			}()
			cfg.crawlPage(currURL)
		}()
	}
}
