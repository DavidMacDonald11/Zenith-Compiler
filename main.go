package main

import (
	"fmt"
	"os"
	"zenith/code"
	"zenith/parser"
	"zenith/semantic"

	"github.com/antlr4-go/antlr/v4"
)

func main() {
    for _, path := range os.Args[1:] {
        file, err := antlr.NewFileStream(path)

        if err != nil {
            fmt.Printf("No such file: %s\n", path)
            continue
        }

        fmt.Printf("Compiling %s:\n\n", path)

        lexer := parser.NewZenithLexer(file)
        tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
        tree := parser.NewZenithParser(tokens).FileStat()

        analyzer := semantic.MakeAnalyzer()
        analyzer.Visit(tree)

        generator := code.Generator{Analyzer: &analyzer}
        res := generator.Visit(tree)
        fmt.Printf("%s\n\n", res.(string))
    }
}
