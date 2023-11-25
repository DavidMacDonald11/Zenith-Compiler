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

	// Visit a parse tree produced by ZenithParser#blankStat.
	VisitBlankStat(ctx *BlankStatContext) interface{}

	// Visit a parse tree produced by ZenithParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by ZenithParser#prefixExpr.
	VisitPrefixExpr(ctx *PrefixExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#castExpr.
	VisitCastExpr(ctx *CastExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#mulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#ifExpr.
	VisitIfExpr(ctx *IfExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#keyExpr.
	VisitKeyExpr(ctx *KeyExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#numExpr.
	VisitNumExpr(ctx *NumExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#parenExpr.
	VisitParenExpr(ctx *ParenExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#idExpr.
	VisitIdExpr(ctx *IdExprContext) interface{}
}
