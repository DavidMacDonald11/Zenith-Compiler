// Code generated from ZenithParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ZenithParser

import "github.com/antlr4-go/antlr/v4"

type BaseZenithParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseZenithParserVisitor) VisitFileStat(ctx *FileStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitNumExpr(ctx *NumExprContext) interface{} {
	return v.VisitChildren(ctx)
}