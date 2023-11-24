package code

import (
	"bytes"
	"fmt"
	"math/big"
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

func (g *Generator) VisitDefineStat(ctx *parser.DefineStatContext) any {
    id := getCId(ctx.ID().GetText())
    exprType := getCType(g.Analyzer.ExprTypes[ctx.Expr()])

    return fmt.Sprintf("%v %v = %v", exprType, id, g.Visit(ctx.Expr()))
}

func (g *Generator) VisitExprStat(ctx *parser.ExprStatContext) any {
    return g.Visit(ctx.Expr())
}

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
    return key
}

func (g *Generator) VisitParenExpr(ctx *parser.ParenExprContext) any {
    inner := g.Visit(ctx.Expr())
    return fmt.Sprintf("(%v)", inner)
}

func (g *Generator) VisitCastExpr(ctx *parser.CastExprContext) any {
    exprType := getCType(g.Analyzer.ExprTypes[ctx])
    inner := g.Visit(ctx.Expr())

    return fmt.Sprintf("(%v)(%v)", exprType, inner)
}

func (g *Generator) VisitPrefixExpr(ctx *parser.PrefixExprContext) any {
    op := ctx.Op.GetText()
    right := g.Visit(ctx.Right)

    return fmt.Sprintf("%v%v", op, right)
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

func (g *Generator) VisitIfExpr(ctx *parser.IfExprContext) any {
    left := g.Visit(ctx.Left)
    condition := g.Visit(ctx.Condition)
    right := g.Visit(ctx.Right)

    return fmt.Sprintf("(%v)? %v : %v", condition, left, right)
}

func getBase(num string) int64 {
    if strings.Contains(num, "0b") { return 2 }
    if strings.Contains(num, "0o") { return 8 }
    if strings.Contains(num, "0x") { return 16 }
    return 10
}

func getDigit(c byte) *big.Float {
    if c >= '0' && c <= '9' {
        return new(big.Float).SetInt64(int64(c - '0'))
    }

    c = bytes.ToLower([]byte{c})[0]
    return new(big.Float).SetInt64(int64(c - 'a' + 10))
}

func getCId(id string) string {
    return "__ZenithId" + id
}

func getCType(t semantic.EType) string {
    switch t {
    case "uint8": return "uint_least8_t"
    case "uint16": return "uint_least16_t"
    case "uint32": return "uint_least32_t"
    case "uint64": return "uint_least64_t"
    case "int8": return "int_least8_t"
    case "int16": return "int_least16_t"
    case "int32": return "int_least32_t"
    case "int64": return "int_least64_t"
    case "uint": return "unsigned int"
    case "anyint": fallthrough
    case "int": return "int"
    case "float32": return "float"
    case "anyfloat": fallthrough
    case "float64": return "double"
    case "bool": return "bool"
    }

    return "void"
}
