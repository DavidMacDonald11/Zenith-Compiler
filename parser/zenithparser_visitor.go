// Code generated from ZenithParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ZenithParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by ZenithParser.
type ZenithParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by ZenithParser#fileStat.
	VisitFileStat(ctx *FileStatContext) interface{}

	// Visit a parse tree produced by ZenithParser#addExpr.
	VisitAddExpr(ctx *AddExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#mulExpr.
	VisitMulExpr(ctx *MulExprContext) interface{}

	// Visit a parse tree produced by ZenithParser#numExpr.
	VisitNumExpr(ctx *NumExprContext) interface{}
}
