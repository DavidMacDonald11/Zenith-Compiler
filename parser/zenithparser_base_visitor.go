// Code generated from ZenithParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ZenithParser

import "github.com/antlr4-go/antlr/v4"

type BaseZenithParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseZenithParserVisitor) VisitFileStat(ctx *FileStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitPrefixExpr(ctx *PrefixExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitCastExpr(ctx *CastExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitAddExpr(ctx *AddExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitMulExpr(ctx *MulExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitIfExpr(ctx *IfExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitKeyExpr(ctx *KeyExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitNumExpr(ctx *NumExprContext) interface{} {
	return v.VisitChildren(ctx)
}
