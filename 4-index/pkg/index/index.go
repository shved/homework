package index

import (
	"fmt"
	"sort"
	"strings"

	"websearch/pkg/crawler"
)

type Service struct {
	plain []crawler.Document
	GIN   map[string][]int
	curID int
}

func New() *Service {
	return &Service{GIN: make(map[string][]int)}
}

var stopWords = []string{
	"i", "a", "about", "an", "are", "as", "at",
	"be", "by", "not", "for", "from", "how", "in",
	"is", "it", "of", "on", "or", "that", "the",
	"this", "to", "was", "what", "when", "where",
	"who", "will", "with", "the", "www",
}

func (s *Service) Process(data []crawler.Document) {
	for _, doc := range data {
		// register docs
		doc.ID = s.curID + 1
		s.curID++
		s.plain = append(s.plain, doc)

		// index words
		words := titleToIndexKeys(doc.Title)
		if len(words) > 0 { // title may contain only stop words
			s.registerWords(doc, words)
		}
	}
}

func (s *Service) FindDoc(id int) (*crawler.Document, error) {
	i := sort.Search(len(s.plain), func(i int) bool { return s.plain[i].ID >= id })
	if i < len(s.plain) && s.plain[i].ID == id {
		return &s.plain[i], nil
	}

	return nil, fmt.Errorf("document with id %v not found", id)
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
			s.GIN[word] = []int{doc.ID}
			continue
		}

		for _, id := range s.GIN[word] {
			if doc.ID == id {
				found = true
				break // doc already added
			}
		}

		if found {
			continue
		}

		s.GIN[word] = append(s.GIN[word], doc.ID)
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
