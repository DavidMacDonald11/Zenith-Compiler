package main

import (
	"fmt"
	"zenith/code"
	"zenith/parser"
	"zenith/semantic"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
    in := antlr.NewInputStream("0b101.0 + 0xF.0 * 0.0")
    lexer := parser.NewZenithLexer(in)
    tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    tree := parser.NewZenithParser(tokens).FileStat()

    analyzer := semantic.Analyzer{}
    res := analyzer.Visit(tree)
    fmt.Println(res.(string))

    generator := code.Generator{}
    res = generator.Visit(tree)
    fmt.Println(res.(string))
}
