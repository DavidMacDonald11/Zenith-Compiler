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

// EnterEndedStat is called when production endedStat is entered.
func (s *BaseZenithParserListener) EnterEndedStat(ctx *EndedStatContext) {}

// ExitEndedStat is called when production endedStat is exited.
func (s *BaseZenithParserListener) ExitEndedStat(ctx *EndedStatContext) {}

// EnterLineEnd is called when production lineEnd is entered.
func (s *BaseZenithParserListener) EnterLineEnd(ctx *LineEndContext) {}

// ExitLineEnd is called when production lineEnd is exited.
func (s *BaseZenithParserListener) ExitLineEnd(ctx *LineEndContext) {}

// EnterDefineStat is called when production defineStat is entered.
func (s *BaseZenithParserListener) EnterDefineStat(ctx *DefineStatContext) {}

// ExitDefineStat is called when production defineStat is exited.
func (s *BaseZenithParserListener) ExitDefineStat(ctx *DefineStatContext) {}

// EnterMultiStat is called when production multiStat is entered.
func (s *BaseZenithParserListener) EnterMultiStat(ctx *MultiStatContext) {}

// ExitMultiStat is called when production multiStat is exited.
func (s *BaseZenithParserListener) ExitMultiStat(ctx *MultiStatContext) {}

// EnterExprStat is called when production exprStat is entered.
func (s *BaseZenithParserListener) EnterExprStat(ctx *ExprStatContext) {}

// ExitExprStat is called when production exprStat is exited.
func (s *BaseZenithParserListener) ExitExprStat(ctx *ExprStatContext) {}

// EnterBaseType is called when production baseType is entered.
func (s *BaseZenithParserListener) EnterBaseType(ctx *BaseTypeContext) {}

// ExitBaseType is called when production baseType is exited.
func (s *BaseZenithParserListener) ExitBaseType(ctx *BaseTypeContext) {}

// EnterPtrType is called when production ptrType is entered.
func (s *BaseZenithParserListener) EnterPtrType(ctx *PtrTypeContext) {}

// ExitPtrType is called when production ptrType is exited.
func (s *BaseZenithParserListener) ExitPtrType(ctx *PtrTypeContext) {}

// EnterCastExpr is called when production castExpr is entered.
func (s *BaseZenithParserListener) EnterCastExpr(ctx *CastExprContext) {}

// ExitCastExpr is called when production castExpr is exited.
func (s *BaseZenithParserListener) ExitCastExpr(ctx *CastExprContext) {}

// EnterKeyExpr is called when production keyExpr is entered.
func (s *BaseZenithParserListener) EnterKeyExpr(ctx *KeyExprContext) {}

// ExitKeyExpr is called when production keyExpr is exited.
func (s *BaseZenithParserListener) ExitKeyExpr(ctx *KeyExprContext) {}

// EnterNumExpr is called when production numExpr is entered.
func (s *BaseZenithParserListener) EnterNumExpr(ctx *NumExprContext) {}

// ExitNumExpr is called when production numExpr is exited.
func (s *BaseZenithParserListener) ExitNumExpr(ctx *NumExprContext) {}

// EnterParenExpr is called when production parenExpr is entered.
func (s *BaseZenithParserListener) EnterParenExpr(ctx *ParenExprContext) {}

// ExitParenExpr is called when production parenExpr is exited.
func (s *BaseZenithParserListener) ExitParenExpr(ctx *ParenExprContext) {}

// EnterBitXorExpr is called when production bitXorExpr is entered.
func (s *BaseZenithParserListener) EnterBitXorExpr(ctx *BitXorExprContext) {}

// ExitBitXorExpr is called when production bitXorExpr is exited.
func (s *BaseZenithParserListener) ExitBitXorExpr(ctx *BitXorExprContext) {}

// EnterShiftExpr is called when production shiftExpr is entered.
func (s *BaseZenithParserListener) EnterShiftExpr(ctx *ShiftExprContext) {}

// ExitShiftExpr is called when production shiftExpr is exited.
func (s *BaseZenithParserListener) ExitShiftExpr(ctx *ShiftExprContext) {}

// EnterPtrExpr is called when production ptrExpr is entered.
func (s *BaseZenithParserListener) EnterPtrExpr(ctx *PtrExprContext) {}

// ExitPtrExpr is called when production ptrExpr is exited.
func (s *BaseZenithParserListener) ExitPtrExpr(ctx *PtrExprContext) {}

// EnterPrefixExpr is called when production prefixExpr is entered.
func (s *BaseZenithParserListener) EnterPrefixExpr(ctx *PrefixExprContext) {}

// ExitPrefixExpr is called when production prefixExpr is exited.
func (s *BaseZenithParserListener) ExitPrefixExpr(ctx *PrefixExprContext) {}

// EnterBitOrExpr is called when production bitOrExpr is entered.
func (s *BaseZenithParserListener) EnterBitOrExpr(ctx *BitOrExprContext) {}

// ExitBitOrExpr is called when production bitOrExpr is exited.
func (s *BaseZenithParserListener) ExitBitOrExpr(ctx *BitOrExprContext) {}

// EnterAddExpr is called when production addExpr is entered.
func (s *BaseZenithParserListener) EnterAddExpr(ctx *AddExprContext) {}

// ExitAddExpr is called when production addExpr is exited.
func (s *BaseZenithParserListener) ExitAddExpr(ctx *AddExprContext) {}

// EnterCompExpr is called when production compExpr is entered.
func (s *BaseZenithParserListener) EnterCompExpr(ctx *CompExprContext) {}

// ExitCompExpr is called when production compExpr is exited.
func (s *BaseZenithParserListener) ExitCompExpr(ctx *CompExprContext) {}

// EnterMulExpr is called when production mulExpr is entered.
func (s *BaseZenithParserListener) EnterMulExpr(ctx *MulExprContext) {}

// ExitMulExpr is called when production mulExpr is exited.
func (s *BaseZenithParserListener) ExitMulExpr(ctx *MulExprContext) {}

// EnterIfExpr is called when production ifExpr is entered.
func (s *BaseZenithParserListener) EnterIfExpr(ctx *IfExprContext) {}

// ExitIfExpr is called when production ifExpr is exited.
func (s *BaseZenithParserListener) ExitIfExpr(ctx *IfExprContext) {}

// EnterBitAndExpr is called when production bitAndExpr is entered.
func (s *BaseZenithParserListener) EnterBitAndExpr(ctx *BitAndExprContext) {}

// ExitBitAndExpr is called when production bitAndExpr is exited.
func (s *BaseZenithParserListener) ExitBitAndExpr(ctx *BitAndExprContext) {}

// EnterPowExpr is called when production powExpr is entered.
func (s *BaseZenithParserListener) EnterPowExpr(ctx *PowExprContext) {}

// ExitPowExpr is called when production powExpr is exited.
func (s *BaseZenithParserListener) ExitPowExpr(ctx *PowExprContext) {}

// EnterIdExpr is called when production idExpr is entered.
func (s *BaseZenithParserListener) EnterIdExpr(ctx *IdExprContext) {}

// ExitIdExpr is called when production idExpr is exited.
func (s *BaseZenithParserListener) ExitIdExpr(ctx *IdExprContext) {}
