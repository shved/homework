package main

import (
	"fmt"
	"os"
	"strings"

	"websearch/pkg/spider"
)

// Scanner - interface for scanners/crawlers
type Scanner interface {
	Scan(url string) (data map[string]string, err error)
}

func main() {
	var searchTerm string
	url := "https://golang.org"
	store := make(map[string]string)

	s := spider.New(2)

	store, err := scan(s, store, url)

	if err != nil {
		fmt.Printf("Scanning %s: %v", url, err)
		os.Exit(1)
	}

	for {
		fmt.Printf("Enter a search term: ")
		fmt.Scanf("%s", &searchTerm)
		fmt.Println("Results:")

		for url, data := range store {
			if strings.Contains(data, searchTerm) {
				fmt.Printf("Entry found on %s page\n", url)
			}
		}
		fmt.Println()
	}
}

func scan(s Scanner, store map[string]string, url string) (map[string]string, error) {
	fmt.Printf("Scanning %s...", url)

	data, err := s.Scan(url)
	if err != nil {
		return store, err
	}

	for k, v := range data {
		store[k] = v
	}

	fmt.Println(" Done.")

	return store, nil
}
