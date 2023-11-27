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

func (g *Generator) VisitCastExpr(ctx *parser.CastExprContext) any {
    exprType := getCType(g.Analyzer.ExprTypes[ctx])
    inner := g.Visit(ctx.Expr())

    return fmt.Sprintf("(%v)(%v)", exprType, inner)
}

func (g *Generator) VisitPtrExpr(ctx *parser.PtrExprContext) any {
    var op string

    if ctx.Op.GetText() == "@" {
        op = "*"
    } else {
        op = "&"
    }

    right := g.Visit(ctx.Right)
    return fmt.Sprintf("%v%v", op, right)
}

func (g *Generator) VisitPrefixExpr(ctx *parser.PrefixExprContext) any {
    op := ctx.Op.GetText()
    right := g.Visit(ctx.Right)

    return fmt.Sprintf("%v%v", op, right)
}

func (g *Generator) VisitPowExpr(ctx *parser.PowExprContext) any {
    left := g.Visit(ctx.Left)
    right := g.Visit(ctx.Right)

    // TODO implement correctly
    return fmt.Sprintf("/*%v^%v*/", left, right)
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

func getCType(t semantic.Type) string {
    if ptr, isPtr := t.(semantic.PtrType); isPtr {
        return getCType(ptr.Base) + "*"
    }

    name := t.(semantic.BaseType).Name

    switch name {
    case "UInt8": return "uint8_t"
    case "UInt16": return "uint16_t"
    case "UInt32": return "uint32_t"
    case "UInt64": return "uint64_t"
    case "Int8": return "int8_t"
    case "Int16": return "int16_t"
    case "Int32": return "int32_t"
    case "Int64": return "int64_t"
    case "UInt": return "uint_fast32_t"
    case "AnyNum": fallthrough
    case "Int": return "int_fast32_t"
    case "Float32": return "float"
    case "AnyFloat": fallthrough
    case "Float64": return "double"
    case "Bool": return "bool"
    }

    return "void"
}
