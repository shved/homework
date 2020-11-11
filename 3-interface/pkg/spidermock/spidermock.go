package spidermock

type Bot struct{}

func New() *Bot {
	return &Bot{}
}

func (b *Bot) Scan(url string) (map[string]string, error) {
	data := map[string]string{
		"http://example.com/":     "Super title",
		"http://example.com/asdf": "Duper title",
	}

	return data, nil
}
