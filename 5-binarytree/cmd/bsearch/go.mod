module websearch/cmd/search

go 1.15

replace websearch/pkg/binarytree v0.0.0 => ../../pkg/binarytree

replace websearch/pkg/crawler v0.0.0 => ../../pkg/crawler

replace websearch/pkg/crawler/spider v0.0.0 => ../../pkg/crawler/spider

replace websearch/pkg/engine v0.0.0 => ../../pkg/engine

replace websearch/pkg/index v0.0.0 => ../../pkg/index

replace websearch/pkg/resource v0.0.0 => ../../pkg/resource

require websearch/pkg/crawler v0.0.0

require websearch/pkg/crawler/spider v0.0.0

require websearch/pkg/engine v0.0.0

require websearch/pkg/index v0.0.0

require websearch/pkg/resource v0.0.0
