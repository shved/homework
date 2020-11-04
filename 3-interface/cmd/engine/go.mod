module websearch/cmd/engine

go 1.15

replace websearch/pkg/spider v0.0.0 => ../../pkg/spider

require (
	golang.org/x/net v0.0.0-20201031054903-ff519b6c9102 // indirect
	websearch/pkg/spider v0.0.0
)
