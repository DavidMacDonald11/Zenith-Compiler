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
    return nil
}

func (a *Analyzer) VisitAddExpr(ctx *parser.AddExprContext) any {
    return nil
}

func (a *Analyzer) VisitNumExpr(ctx *parser.NumExprContext) any {
    if strings.ContainsRune(ctx.Num.GetText(), '.') {
        return "float64"
    }

    return "int"
}
