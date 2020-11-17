package main

import (
	"fmt"
	"os"
	"strings"

	"websearch/pkg/crawler"
	"websearch/pkg/crawler/spider"
	"websearch/pkg/index"
)

func main() {
	var searchTerm string
	url := "https://golang.org"
	depth := 2

	spdr := spider.New()
	indx := index.New()

	docs, err := scan(spdr, url, depth)

	if err != nil {
		fmt.Printf("Scanning %s: %v", url, err)
		os.Exit(1)
	}

	indx.Process(docs)

	for {
		fmt.Printf("Enter a search term: ")
		fmt.Scanf("%s", &searchTerm)
		fmt.Println("Results:")

		searchTerm = strings.ToLower(searchTerm)

		// TODO искать по фразам, не только по одному слову
		if entries, ok := indx.GIN[searchTerm]; ok {
			for _, id := range entries {
				docPtr, err := indx.FindDoc(id)
				if err != nil {
					continue
				}
				fmt.Printf("Entry found on %s page\n", docPtr.URL)
			}
		}
		fmt.Println()
	}
}

func scan(s crawler.Scanner, url string, depth int) ([]crawler.Document, error) {
	fmt.Printf("Scanning %s...", url)

	data, err := s.Scan(url, depth)
	if err != nil {
		return []crawler.Document{}, err
	}

	fmt.Println(" Done.")

	return data, nil
}
