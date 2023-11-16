package code

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
	"zenith/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Generator struct {
    parser.BaseZenithParserVisitor
}

func (g *Generator) Visit(tree antlr.ParseTree) any {
    return tree.Accept(g)
}

func (g *Generator) VisitFileStat(ctx *parser.FileStatContext) any {
    return g.Visit(ctx.Expr())
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

func (g *Generator) VisitNumExpr(ctx *parser.NumExprContext) any {
    str := strings.ReplaceAll(ctx.Num.GetText(), "_", "")
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
