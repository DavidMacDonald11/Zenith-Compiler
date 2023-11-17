package main

import (
	"fmt"
	"zenith/code"
	"zenith/parser"
	"zenith/semantic"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
    str := "5 if true else uint(6)"
    fmt.Printf(`Compiling "%s"` + "\n", str)
    in := antlr.NewInputStream(str)

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
