module websearch/pkg/crawler/spider

go 1.15

replace websearch/pkg/crawler v0.0.0 => ../

replace websearch/pkg/resource v0.0.0 => ../../resource

require websearch/pkg/crawler v0.0.0

require websearch/pkg/resource v0.0.0

require golang.org/x/net v0.0.0-20201031054903-ff519b6c9102
