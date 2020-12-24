// Package index indexes, stores and retrieves webpages
package index

import (
	"strings"

	"websearch/pkg/binarytree"
	"websearch/pkg/resource"
)

// Service - search index service
type Service struct {
	Store *binarytree.Tree
	Data  Data
}

// New - creates and returns a new index service
func New() *Service {
	return &Service{
		Store: new(binarytree.Tree),
		Data:  make(Data),
	}
}

// Data contains a hashmap with words as keys and as values -
// an array of IDs of documents (webpages) containing given word.
type Data map[string][]uint

type wordIDPair struct {
	word string
	id   uint
}

var stopWords = []string{
	"i", "a", "about", "an", "are", "as", "at",
	"be", "by", "not", "for", "from", "how", "in",
	"is", "it", "of", "on", "or", "that", "the",
	"this", "to", "was", "what", "when", "where",
	"who", "will", "with", "the", "www",
}

// Build builds an index for an array of resource.Documents
func (s *Service) Build(docs []resource.Document) {
	allWordIDPairs := []wordIDPair{}

	// Extract all words from all docs into an array
	for _, d := range docs {
		// Store doc
		id := store(s, d)

		docWords := strings.Split(strip(d.Title), " ")

		for _, w := range docWords {
			widp := wordIDPair{word: w, id: id}
			allWordIDPairs = append(allWordIDPairs, widp)
		}
	}

	// Convert all words to lowercase
	for i, wdp := range allWordIDPairs {
		allWordIDPairs[i].word = strings.ToLower(wdp.word)
	}

	// Filter out stop words
	filtWordIDPairs := []wordIDPair{}

	for _, widp := range allWordIDPairs {
		if !contains(stopWords, widp.word) {
			filtWordIDPairs = append(filtWordIDPairs, widp)
		}
	}

	for _, widp := range filtWordIDPairs {
		s.Data[widp.word] = append(s.Data[widp.word], widp.id)
	}
}

func store(s *Service, d resource.Document) uint {
	return s.Store.Insert(&resource.Document{URL: d.URL, Title: d.Title})
}

func strip(s string) string {
	var result strings.Builder

	result.Grow(len(s))

	for i := 0; i < len(s); i++ {
		b := s[i]
		if ('a' <= b && b <= 'z') ||
			('A' <= b && b <= 'Z') ||
			('0' <= b && b <= '9') ||
			b == ' ' {
			result.WriteByte(b)
		}
	}

	return result.String()
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}

	return false
}
