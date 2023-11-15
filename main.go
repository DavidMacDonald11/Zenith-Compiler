package main

import (
	"fmt"
	"zenith/parser"
	"zenith/semantic"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
    in := antlr.NewInputStream("100.3")
    lexer := parser.NewZenithLexer(in)
    tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    tree := parser.NewZenithParser(tokens).FileStat()

    analyzer := semantic.Analyzer {}

    res := analyzer.Visit(tree)
    fmt.Println(res.(string))
}
