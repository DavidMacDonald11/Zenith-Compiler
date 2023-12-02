package code

import (
	"fmt"
	"strings"
	"zenith/parser"
	"zenith/semantic"

	"github.com/antlr4-go/antlr/v4"
)

type Generator struct {
    parser.BaseZenithParserVisitor
    depth int

    Analyzer *semantic.Analyzer
}

func (g *Generator) Visit(tree antlr.ParseTree) any {
    return tree.Accept(g)
}

func (g *Generator) VisitFileStat(ctx *parser.FileStatContext) any {
    result := strings.Builder{}

    for _, stat := range ctx.AllEndedStat() {
        result.WriteString(fmt.Sprintf("%s\n", g.Visit(stat).(string)))
    }

    return result.String()
}

func (g *Generator) VisitEndedStat(ctx *parser.EndedStatContext) any {
    indent := strings.Repeat("    ", g.depth)
    return indent + g.Visit(ctx.Stat()).(string) + ";"
}

func (g *Generator) VisitDefineStat(ctx *parser.DefineStatContext) any {
    id := getCId(ctx.ID().GetText())
    exprType := getCType(g.Analyzer.ExprTypes[ctx.Expr()])

    return fmt.Sprintf("%v %v = %v", exprType, id, g.Visit(ctx.Expr()))
}

func (g *Generator) VisitMultiStat(ctx *parser.MultiStatContext) any {
    result := strings.Builder{}
    result.WriteString("{\n")
    g.depth++

    for _, stat := range ctx.AllEndedStat() {
        result.WriteString(fmt.Sprintf("%s\n", g.Visit(stat)))
    }

    if stat := ctx.Stat(); stat != nil {
        result.WriteString(fmt.Sprintf("%s;\n", g.Visit(stat)))
    }

    g.depth--
    indent := strings.Repeat("    ", g.depth)
    result.WriteString(indent + "}")

    return result.String()
}

func (g *Generator) VisitExprStat(ctx *parser.ExprStatContext) any {
    return g.Visit(ctx.Expr())
}
