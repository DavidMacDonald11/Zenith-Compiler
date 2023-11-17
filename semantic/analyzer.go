package semantic

import (
	"strings"
	"zenith/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Analyzer struct {
    parser.BaseZenithParserVisitor
    ExprTypes map[parser.IExprContext]string
}

func MakeAnalyzer() Analyzer {
    return Analyzer{ExprTypes: make(map[parser.IExprContext]string)}
}

func (a *Analyzer) Visit(tree antlr.ParseTree) any {
    return tree.Accept(a)
}

func (a *Analyzer) VisitFileStat(ctx *parser.FileStatContext) any {
    return a.Visit(ctx.Expr())
}

func (a *Analyzer) VisitMulExpr(ctx *parser.MulExprContext) any {
    left := a.Visit(ctx.Left).(string)
    right := a.Visit(ctx.Right).(string)
    t := deduceCommonType(left, right)

    if t == nil {
        panic("Cannot mul/div/rem different types")
    }

    if strings.Contains(*t, "float") && ctx.Op.GetText() == "%" {
        panic("Cannot take rem with float")
    }

    return *t
}

func (a *Analyzer) VisitAddExpr(ctx *parser.AddExprContext) any {
    left := a.Visit(ctx.Left).(string)
    right := a.Visit(ctx.Right).(string)
    t := deduceCommonType(left, right)

    if t == nil {
        panic("Cannot add/sub different types")
    }

    return *t
}

func (a *Analyzer) VisitCastExpr(ctx *parser.CastExprContext) any {
    a.Visit(ctx.Inner)
    a.ExprTypes[ctx] = ctx.Type.GetText()

    return a.ExprTypes[ctx]
}

func (a *Analyzer) VisitParenExpr(ctx *parser.ParenExprContext) any {
    return a.Visit(ctx.Inner)
}

func (a *Analyzer) VisitNumExpr(ctx *parser.NumExprContext) any {
    num := ctx.Num.GetText()

    if strings.ContainsRune(num, '.') {
        return "anyfloat"
    }

    return "anyint"
}

func deduceCommonType(left, right string) *string {
    if left == "anyint" {
        return &right
    }

    if left == "anyfloat" {
        if right == "anyint" {
            return &left
        }

        if strings.Contains(right, "int") {
            return nil
        }

        return &right
    }

    if right == "anyint" {
        return &left
    }

    if right == "anyfloat" {
        if strings.Contains(left, "int") {
            return nil
        }

        return &left
    }

    if left == right {
        return &left
    }

    return nil
}
