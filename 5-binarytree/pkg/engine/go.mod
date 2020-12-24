module websearch/pkg/engine

go 1.15

replace websearch/pkg/binarytree v0.0.0 => ../../pkg/binarytree

replace websearch/pkg/index v0.0.0 => ../index

replace websearch/pkg/resource v0.0.0 => ../resource

require websearch/pkg/index v0.0.0

require websearch/pkg/resource v0.0.0
