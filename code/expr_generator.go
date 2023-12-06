package code

import (
	"fmt"
	"strings"
	"zenith/parser"
	"zenith/semantic"
)

func (g *Generator) VisitNumExpr(ctx *parser.NumExprContext) any {
    return semantic.GetNum(ctx.NUM().GetText())
}

func (g *Generator) VisitIdExpr(ctx *parser.IdExprContext) any {
    return ctx.ID().GetText()
}

func (g *Generator) VisitKeyExpr(ctx *parser.KeyExprContext) any {
    key := ctx.Key.GetText()

    if key == "null" { return "nullptr" }
    return key
}

func (g *Generator) VisitInitExpr(ctx *parser.InitExprContext) any {
    exprType := g.Analyzer.ExprTypes[ctx]
    expr := g.Visit(ctx.InitArgs())
    slice, isSlice := exprType.(semantic.SliceType)

    if !isSlice {
        return fmt.Sprintf("%v{%v}", getCType(exprType), expr)
    }

    size := slice.Size
    arrType := getCType(slice.Base)

    if size == 0 {
        return fmt.Sprintf("Zenith::Slice<%v>{%v}", arrType, expr)
    }

    return fmt.Sprintf("Zenith::Array<%v, %v>{%v}", arrType, size, expr)
}

func (g *Generator) VisitInitArgs(ctx *parser.InitArgsContext) any {
    exprs := ctx.AllExpr()
    b := strings.Builder{}

    for i, expr := range exprs {
        str := g.Visit(expr).(string)
        b.WriteString(str)

        if i + 1 < len(exprs) { b.WriteString(", ") }
    }

    return b.String()
}

func (g *Generator) VisitParenExpr(ctx *parser.ParenExprContext) any {
    expr := g.Visit(ctx.Expr())
    return fmt.Sprintf("(%v)", expr)
}

func (g *Generator) VisitPostfixExpr(ctx *parser.PostfixExprContext) any {
    expr := g.Visit(ctx.Expr())
    return fmt.Sprintf("(%v != nullptr)", expr)
}

func (g *Generator) VisitAllocExpr(ctx *parser.AllocExprContext) any {
    expr := g.Visit(ctx.Expr())
    exprType := getCType(g.Analyzer.ExprTypes[ctx.Expr()])

    return fmt.Sprintf("new %v(%v)", exprType, expr)
}

func (g *Generator) VisitDeallocExpr(ctx *parser.DeallocExprContext) any {
    expr := g.Visit(ctx.Expr())
    return fmt.Sprintf("delete %v", expr)
}

func (g *Generator) VisitCastExpr(ctx *parser.CastExprContext) any {
    exprType := getCType(g.Analyzer.ExprTypes[ctx])
    inner := g.Visit(ctx.Expr())

    return fmt.Sprintf("(%v)(%v)", exprType, inner)
}

func (g *Generator) VisitPtrExpr(ctx *parser.PtrExprContext) any {
    expr := g.Visit(ctx.Expr())
    op := ctx.Op.GetText()

    if op == "@" { op = "*" }
    return fmt.Sprintf("%v%v", op, expr)
}

func (g *Generator) VisitPrefixExpr(ctx *parser.PrefixExprContext) any {
    op := ctx.Op.GetText()
    right := g.Visit(ctx.Expr())

    return fmt.Sprintf("%v%v", op, right)
}

func (g *Generator) VisitPowExpr(ctx *parser.PowExprContext) any {
    left := g.Visit(ctx.Left)
    right := g.Visit(ctx.Right)

    exprType := g.Analyzer.ExprTypes[ctx]
    cType := getCType(exprType)

    if semantic.IsInt(exprType) {
        return fmt.Sprintf("Zenith::pow<%v>(%v, %v)", cType, left, right)
    }

    return fmt.Sprintf("std::pow(%v, %v)", left, right)
}

func (g *Generator) VisitMulExpr(ctx *parser.MulExprContext) any {
    left := g.Visit(ctx.Left)
    right := g.Visit(ctx.Right)
    op := ctx.Op.GetText()

    return fmt.Sprintf("%v %v %v", left, op, right)
}

func (g *Generator) VisitAddExpr(ctx *parser.AddExprContext) any {
    left := g.Visit(ctx.Left)
    right := g.Visit(ctx.Right)
    op := ctx.Op.GetText()

    return fmt.Sprintf("%v %v %v", left, op, right)
}

func (g *Generator) VisitShiftExpr(ctx *parser.ShiftExprContext) any {
    left := g.Visit(ctx.Left)
    right := g.Visit(ctx.Right)
    op := ctx.Op.GetText()

    return fmt.Sprintf("%v %v %v", left, op, right)
}

func (g *Generator) VisitBitAndExpr(ctx *parser.BitAndExprContext) any {
    left := g.Visit(ctx.Left)
    right := g.Visit(ctx.Right)

    return fmt.Sprintf("%v & %v", left, right)
}

func (g *Generator) VisitBitXorExpr(ctx *parser.BitXorExprContext) any {
    left := g.Visit(ctx.Left)
    right := g.Visit(ctx.Right)

    return fmt.Sprintf("%v ^ %v", left, right)
}

func (g *Generator) VisitBitOrExpr(ctx *parser.BitOrExprContext) any {
    left := g.Visit(ctx.Left)
    right := g.Visit(ctx.Right)

    return fmt.Sprintf("%v | %v", left, right)
}

func (g *Generator) VisitCompExpr(ctx *parser.CompExprContext) any {
    left := g.Visit(ctx.Left)
    right := g.Visit(ctx.Right)
    op := ctx.Op.GetText()

    // TODO implement correctly
    return fmt.Sprintf("%v %v %v", left, op, right)
}

func (g *Generator) VisitIfExpr(ctx *parser.IfExprContext) any {
    left := g.Visit(ctx.Left)
    condition := g.Visit(ctx.Condition)
    right := g.Visit(ctx.Right)

    return fmt.Sprintf("(%v)? %v : %v", condition, left, right)
}
