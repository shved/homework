package crawler

type Document struct {
	ID    int
	URL   string
	Title string
	// Body  string
}

type Scanner interface {
	Scan(url string, depth int) ([]Document, error)
}
