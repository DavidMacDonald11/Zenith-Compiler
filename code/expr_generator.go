package code

import (
	"fmt"
	"math/big"
	"strings"
	"zenith/parser"
)

func (g *Generator) VisitNumExpr(ctx *parser.NumExprContext) any {
    str := strings.ReplaceAll(ctx.NUM().GetText(), "_", "")
    b := getBase(str)

    base := new(big.Float).SetInt64(b)
    num := new(big.Float).SetPrec(64)
    i := 0

    if b != 10 {
        i = 2
    }

    for ; i < len(str) && str[i] != '.'; i++ {
        num.Mul(num, base)
        num.Add(num, getDigit(str[i]))
    }

    inv := new(big.Float).SetPrec(64).SetInt64(1)

    for i++; i < len(str); i++ {
        inv.Quo(inv, base)

        digit := getDigit(str[i])
        digit.Mul(digit, inv)
        num.Add(num, digit)
    }

    isFloat := strings.ContainsRune(str, '.')
    str = num.Text('g', 64)

    if isFloat && num.IsInt() {
        return str + ".0"
    }

    return str
}

func (g *Generator) VisitIdExpr(ctx *parser.IdExprContext) any {
    return getCId(ctx.ID().GetText())
}

func (g *Generator) VisitKeyExpr(ctx *parser.KeyExprContext) any {
    key := ctx.Key.GetText()

    if key == "null" { return "0" }
    return key
}

func (g *Generator) VisitParenExpr(ctx *parser.ParenExprContext) any {
    inner := g.Visit(ctx.Expr())
    return fmt.Sprintf("(%v)", inner)
}

func (g *Generator) VisitPostfixExpr(ctx *parser.PostfixExprContext) any {
    expr := g.Visit(ctx.Expr())
    return fmt.Sprintf("(!%v)", expr)
}

func (g *Generator) VisitAllocExpr(ctx *parser.AllocExprContext) any {
    expr := g.Visit(ctx.Expr())
    // TODO change
    return fmt.Sprintf("__Zenith_alloc(%v)", expr)
}

func (g *Generator) VisitDeallocExpr(ctx *parser.DeallocExprContext) any {
    expr := g.Visit(ctx.Expr())
    // TODO change
    return fmt.Sprintf("__Zenith_dealloc(%v)", expr)
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
    exprType := getCType(g.Analyzer.ExprTypes[ctx])

    // TODO implement better
    return fmt.Sprintf("(%v)powl(%v, %v)", exprType, left, right)
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
