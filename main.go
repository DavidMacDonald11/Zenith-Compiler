package main

import (
	"fmt"
	"zenith/code"
	"zenith/parser"
	"zenith/semantic"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
    in := antlr.NewInputStream("16 - 2 % uint(1.0 + 1)")
    lexer := parser.NewZenithLexer(in)
    tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    tree := parser.NewZenithParser(tokens).FileStat()

    analyzer := semantic.MakeAnalyzer()
    res := analyzer.Visit(tree)
    fmt.Println(res.(string))

    generator := code.Generator{Analyzer: &analyzer}
    res = generator.Visit(tree)
    fmt.Println(res.(string))
}
