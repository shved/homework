package index

import (
	"strings"

	"websearch/pkg/crawler"
)

type Service struct {
	plain map[string]crawler.Document
	GIN   map[string][]string
}

func New() *Service {
	return &Service{plain: make(map[string]crawler.Document), GIN: make(map[string][]string)}
}

var stopWords = []string{
	"i", "a", "about", "an", "are", "as", "at",
	"be", "by", "com", "for", "from", "how", "in",
	"is", "it", "of", "on", "or", "that", "the",
	"this", "to", "was", "what", "when", "where",
	"who", "will", "with", "the", "www",
}

func (s *Service) Process(data []crawler.Document) {
	for _, doc := range data {
		// register a doc if it never met
		if _, ok := s.plain[doc.URL]; !ok {
			s.plain[doc.URL] = doc
		}

		// index doc words
		words := titleToIndexKeys(doc.Title)
		if len(words) > 0 { // title may contain only stop words
			s.registerWords(doc, words)
		}
	}
}

func titleToIndexKeys(title string) []string {
	// TODO почистить слова от несловесного мусора типа двоеточий или запятых
	var words []string

	raw := strings.Fields(title)

	for _, word := range raw {
		if len(word) < 3 {
			continue
		}

		word = strings.ToLower(word)

		if contains(stopWords, word) {
			continue
		}

		words = append(words, word)
	}

	return words
}

func (s *Service) registerWords(doc crawler.Document, words []string) {
	var found bool

	for _, word := range words {
		if _, ok := s.GIN[word]; !ok {
			s.GIN[word] = []string{doc.URL}
			continue
		}

		for _, url := range s.GIN[word] {
			if doc.URL == url {
				found = true
				break // doc already added
			}

		}

		if found {
			continue
		}

		s.GIN[word] = append(s.GIN[word], doc.URL)
	}
}

func contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}
