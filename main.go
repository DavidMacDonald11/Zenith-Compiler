package main

import (
	"fmt"
	"zenith/parser"

	"github.com/antlr4-go/antlr/v4"
)

type treeShapeListener struct {
    *parser.BaseZenithParserListener
}

func (s *treeShapeListener) VisitTerminal(node antlr.TerminalNode) {
    fmt.Printf("Terminal: %v\n", node.GetText())
}

func (s *treeShapeListener) VisitErrorNode(node antlr.ErrorNode) {
    fmt.Printf("Error: %v\n", node.GetText())
}

func (s *treeShapeListener) EnterEveryRule(ctx antlr.ParserRuleContext) {
    fmt.Printf("Enter: %v\n", ctx.GetRuleIndex())
}

func (s *treeShapeListener) ExitEveryRule(ctx antlr.ParserRuleContext) {
    fmt.Printf("Exit: %v\n", ctx.GetRuleIndex())
}

func main() {
    in := antlr.NewInputStream("1 + 2 / 3")
    lexer := parser.NewZenithLexer(in)
    tokens := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
    parser := parser.NewZenithParser(tokens)
    tree := parser.FileStat()

    antlr.ParseTreeWalkerDefault.Walk(&treeShapeListener{}, tree)
}
