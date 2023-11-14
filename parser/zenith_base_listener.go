// Code generated from Zenith.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // Zenith

import "github.com/antlr4-go/antlr/v4"

// BaseZenithListener is a complete listener for a parse tree produced by ZenithParser.
type BaseZenithListener struct{}

var _ ZenithListener = &BaseZenithListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseZenithListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseZenithListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseZenithListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseZenithListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterFileStat is called when production fileStat is entered.
func (s *BaseZenithListener) EnterFileStat(ctx *FileStatContext) {}

// ExitFileStat is called when production fileStat is exited.
func (s *BaseZenithListener) ExitFileStat(ctx *FileStatContext) {}

// EnterMulExpr is called when production MulExpr is entered.
func (s *BaseZenithListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production MulExpr is exited.
func (s *BaseZenithListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterNum is called when production Num is entered.
func (s *BaseZenithListener) EnterNum(ctx *NumContext) {}

// ExitNum is called when production Num is exited.
func (s *BaseZenithListener) ExitNum(ctx *NumContext) {}

// EnterAddExpr is called when production AddExpr is entered.
func (s *BaseZenithListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production AddExpr is exited.
func (s *BaseZenithListener) ExitAddExpr(ctx *AddExprContext) {}
