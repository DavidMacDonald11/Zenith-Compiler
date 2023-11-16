package semantic

import (
	"strings"
	"zenith/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Analyzer struct {
    parser.BaseZenithParserVisitor
}

func (a *Analyzer) Visit(tree antlr.ParseTree) any {
    return tree.Accept(a)
}

func (a *Analyzer) VisitFileStat(ctx *parser.FileStatContext) any {
    return a.Visit(ctx.Expr())
}

func (a *Analyzer) VisitMulExpr(ctx *parser.MulExprContext) any {
    left := a.Visit(ctx.Left)
    right := a.Visit(ctx.Right)

    if left.(string) != right.(string) {
        panic("Cannot mul/div/rem different types")
    }

    return left
}

func (a *Analyzer) VisitAddExpr(ctx *parser.AddExprContext) any {
    left := a.Visit(ctx.Left)
    right := a.Visit(ctx.Right)

    if left.(string) != right.(string) {
        panic("Cannot add/sub different types")
    }

    return left
}

func (a *Analyzer) VisitNumExpr(ctx *parser.NumExprContext) any {
    num := ctx.Num.GetText()

    if strings.ContainsRune(num, '.') {
        return "float64"
    }

    return "int"
}
