// Code generated from ZenithParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ZenithParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ZenithParser.
type ZenithParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ZenithParser#fileStat.
	VisitFileStat(ctx *FileStatContext) interface{}

	// Visit a parse tree produced by ZenithParser#endedStat.
	VisitEndedStat(ctx *EndedStatContext) interface{}

	// Visit a parse tree produced by ZenithParser#lineEnd.
	VisitLineEnd(ctx *LineEndContext) interface{}

	// Visit a parse tree produced by ZenithParser#defineStat.
	VisitDefineStat(ctx *DefineStatContext) interface{}

	// Visit a parse tree produced by ZenithParser#multiStat.
	VisitMultiStat(ctx *MultiStatContext) interface{}

	// Visit a parse tree produced by ZenithParser#exprStat.
	VisitExprStat(ctx *ExprStatContext) interface{}

	// Visit a parse tree produced by ZenithParser#baseType.
	VisitBaseType(ctx *BaseTypeContext) interface{}

	// Visit a parse tree produced by ZenithParser#ptrType.
	VisitPtrType(ctx *PtrTypeContext) interface{}

	// Visit a parse tree produced by ZenithParser#castExpr.
	VisitCastExpr(ctx *CastExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#allocExpr.
	VisitAllocExpr(ctx *AllocExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#keyExpr.
	VisitKeyExpr(ctx *KeyExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#numExpr.
	VisitNumExpr(ctx *NumExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#bitXorExpr.
	VisitBitXorExpr(ctx *BitXorExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#shiftExpr.
	VisitShiftExpr(ctx *ShiftExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#ptrExpr.
	VisitPtrExpr(ctx *PtrExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#prefixExpr.
	VisitPrefixExpr(ctx *PrefixExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#deallocExpr.
	VisitDeallocExpr(ctx *DeallocExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#bitOrExpr.
	VisitBitOrExpr(ctx *BitOrExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#compExpr.
	VisitCompExpr(ctx *CompExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#mulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#ifExpr.
	VisitIfExpr(ctx *IfExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#postfixExpr.
	VisitPostfixExpr(ctx *PostfixExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#bitAndExpr.
	VisitBitAndExpr(ctx *BitAndExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#powExpr.
	VisitPowExpr(ctx *PowExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#idExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}
}
