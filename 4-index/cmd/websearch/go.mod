module websearch/cmd/engine

go 1.15

replace websearch/pkg/crawler/spider v0.0.0 => ../../pkg/crawler/spider

replace websearch/pkg/crawler v0.0.0 => ../../pkg/crawler

replace websearch/pkg/index v0.0.0 => ../../pkg/index

require websearch/pkg/crawler/spider v0.0.0

require websearch/pkg/crawler v0.0.0

require websearch/pkg/index v0.0.0
