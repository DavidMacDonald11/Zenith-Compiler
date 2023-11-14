// Code generated from ZenithParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ZenithParser

import "github.com/antlr4-go/antlr/v4"

// BaseZenithParserListener is a complete listener for a parse tree produced by ZenithParser.
type BaseZenithParserListener struct{}

var _ ZenithParserListener = &BaseZenithParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseZenithParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseZenithParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseZenithParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseZenithParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFileStat is called when production fileStat is entered.
func (s *BaseZenithParserListener) EnterFileStat(ctx *FileStatContext) {}

// ExitFileStat is called when production fileStat is exited.
func (s *BaseZenithParserListener) ExitFileStat(ctx *FileStatContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseZenithParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseZenithParserListener) ExitExpr(ctx *ExprContext) {}
