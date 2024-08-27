package main

import (
	"fmt"
	"os"
)

func main() {
	cmdArgs := os.Args[1:]

	if len(cmdArgs) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(cmdArgs) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl\n%v\n", cmdArgs[0])
	}
}
