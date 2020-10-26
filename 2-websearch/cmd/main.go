package main

import (
	"fmt"
	"strings"

	"websearch/pkg/spider"
)

func main() {
	var searchTerm string
	scanUrl := "https://golang.org"

	pages, err := spider.Scan(scanUrl, 2)
	if err != nil {
		fmt.Printf("scanning site %s: %v\n", scanUrl, err)
		return
	}

	fmt.Printf("Scanning site %s completed\n", scanUrl)

	for {
		fmt.Printf("Enter a search term: ")
		fmt.Scanf("%s", &searchTerm)
		fmt.Println("I found:")

		for url, data := range pages {
			if strings.Contains(data, searchTerm) {
				fmt.Printf("Entry found on %s page\n", url)
			}
		}
		fmt.Println()
	}
}
