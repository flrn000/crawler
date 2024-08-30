package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

type config struct {
	pages              map[string]int
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

const maxConcurrent int = 5

func main() {
	cmdArgs := os.Args[1:]
	var baseURL string

	if len(cmdArgs) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(cmdArgs) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		baseURL = cmdArgs[0]
		fmt.Printf("starting crawl\n%v\n", baseURL)
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
	}

	cfg.crawlPage(cfg.baseURL.String())

	cfg.wg.Wait()

	fmt.Println(cfg.pages, len(cfg.pages))
}
