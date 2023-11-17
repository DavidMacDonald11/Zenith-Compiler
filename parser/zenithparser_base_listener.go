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

// EnterPrefixExpr is called when production prefixExpr is entered.
func (s *BaseZenithParserListener) EnterPrefixExpr(ctx *PrefixExprContext) {}

// ExitPrefixExpr is called when production prefixExpr is exited.
func (s *BaseZenithParserListener) ExitPrefixExpr(ctx *PrefixExprContext) {}

// EnterCastExpr is called when production castExpr is entered.
func (s *BaseZenithParserListener) EnterCastExpr(ctx *CastExprContext) {}

// ExitCastExpr is called when production castExpr is exited.
func (s *BaseZenithParserListener) ExitCastExpr(ctx *CastExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseZenithParserListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseZenithParserListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterMulExpr is called when production mulExpr is entered.
func (s *BaseZenithParserListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production mulExpr is exited.
func (s *BaseZenithParserListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterIfExpr is called when production ifExpr is entered.
func (s *BaseZenithParserListener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production ifExpr is exited.
func (s *BaseZenithParserListener) ExitIfExpr(ctx *IfExprContext) {}

// EnterKeyExpr is called when production keyExpr is entered.
func (s *BaseZenithParserListener) EnterKeyExpr(ctx *KeyExprContext) {}

// ExitKeyExpr is called when production keyExpr is exited.
func (s *BaseZenithParserListener) ExitKeyExpr(ctx *KeyExprContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseZenithParserListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseZenithParserListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterNumExpr is called when production numExpr is entered.
func (s *BaseZenithParserListener) EnterNumExpr(ctx *NumExprContext) {}

// ExitNumExpr is called when production numExpr is exited.
func (s *BaseZenithParserListener) ExitNumExpr(ctx *NumExprContext) {}
