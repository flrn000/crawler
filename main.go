package main

import (
	"fmt"
	"net/url"
	"os"
	"strconv"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
	maxPages           int
}

func main() {
	cmdArgs := os.Args[1:]

	if len(cmdArgs) < 3 {
		fmt.Println("not enough arguments provided")
		os.Exit(1)
	} else if len(cmdArgs) > 3 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	baseURL := cmdArgs[0]
	maxConcurrent, err := strconv.Atoi(cmdArgs[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	maxPages, err := strconv.Atoi(cmdArgs[2])
	if err != nil {
		fmt.Println(err)
		return
	}

	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	cfg := config{
		pages:              make(map[string]int),
		baseURL:            parsedURL,
		mu:                 &sync.Mutex{},
		concurrencyControl: make(chan struct{}, maxConcurrent),
		wg:                 &sync.WaitGroup{},
		maxPages:           maxPages,
	}

	fmt.Printf("starting crawl\n%v\n", baseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(baseURL)
	cfg.wg.Wait()

	for normalizedURL, count := range cfg.pages {
		fmt.Printf("%s - %d\n", normalizedURL, count)
	}
}
