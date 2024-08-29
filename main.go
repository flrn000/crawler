package main

import (
	"fmt"
	"os"
)

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

	pages := crawlPage(baseURL, baseURL, make(map[string]int))

	fmt.Println(pages)
}
