package crawler

import "websearch/pkg/resource"

// Interface определяет контракт поискового робота.
type Interface interface {
	Scan(url string, depth int) ([]resource.Document, error)
}
