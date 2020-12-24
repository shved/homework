package main

import (
	"bufio"
	"fmt"
	"os"

	"websearch/pkg/crawler"
	"websearch/pkg/crawler/spider"
	"websearch/pkg/engine"
	"websearch/pkg/index"
	"websearch/pkg/resource"
)

func main() {
	var query string
	urls := []string{"https://golang.org", "https://go.dev"}
	depth := 2

	indx := index.New()
	scnr := spider.New()
	engn := engine.New()

	scanRes, err := scan(scnr, urls, depth)
	if err != nil {
		fmt.Println(err)
		return
	}

	indx.Build(scanRes)

	for {
		fmt.Printf("Search query: ")

		bscanner := bufio.NewScanner(os.Stdin)
		for bscanner.Scan() {
			query = bscanner.Text()
			break
		}

		if query == "" {
			continue
		}

		res := engn.Find(query, indx)

		printResults(res)
	}
}

func scan(scanner crawler.Interface, urls []string, depth int) ([]resource.Document, error) {
	res := []resource.Document{}

	for i := 0; i < len(urls); i++ {
		fmt.Printf("Scanning %s...", urls[i])

		data, err := scanner.Scan(urls[i], depth)
		if err != nil {
			return nil, err
		}

		for _, d := range data {
			res = append(res, d)
		}

		fmt.Printf(" done!\n")
	}

	return res, nil
}

func printResults(res map[string]string) {
	fmt.Printf("\n%v results\n", len(res))
	for k, v := range res {
		fmt.Println()
		fmt.Println(k)
		fmt.Println(v)
	}
}
