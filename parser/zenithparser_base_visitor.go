// Code generated from ZenithParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ZenithParser

import "github.com/antlr4-go/antlr/v4"

type BaseZenithParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseZenithParserVisitor) VisitFileStat(ctx *FileStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitEndedStat(ctx *EndedStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitLineEnd(ctx *LineEndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitDefineStat(ctx *DefineStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitMultiStat(ctx *MultiStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitExprStat(ctx *ExprStatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitType(ctx *TypeContext) interface{} {
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

func (v *BaseZenithParserVisitor) VisitNumExpr(ctx *NumExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitParenExpr(ctx *ParenExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseZenithParserVisitor) VisitIdExpr(ctx *IdExprContext) interface{} {
	return v.VisitChildren(ctx)
}
