package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	
	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	}

	 if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	}

	url := args[0]
	
	fmt.Printf("starting crawl of %v\n", url)

	pages := make(map[string]int)

	crawlPage(url, url, pages)

	for normalizedURL, count := range pages {
		fmt.Printf("%d - %s\n", count, normalizedURL)
	}
}
