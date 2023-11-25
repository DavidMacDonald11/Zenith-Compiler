package semantic

import (
	"fmt"
	"strings"
	"zenith/parser"

	"github.com/antlr4-go/antlr/v4"
)

type Analyzer struct {
    parser.BaseZenithParserVisitor
    table Table

    ExprTypes map[parser.IExprContext]EType
}

type result struct {
    exprType EType
}

func MakeAnalyzer() Analyzer {
    return Analyzer {
        table: MakeTable(),
        ExprTypes: map[parser.IExprContext]EType{},
    }
}

func (a *Analyzer) Visit(tree antlr.ParseTree) any {
    return tree.Accept(a)
}

func (a *Analyzer) VisitFileStat(ctx *parser.FileStatContext) any {
    for _, stat := range ctx.AllEndedStat() {
        a.Visit(stat)
    }

    return nil
}

func (a *Analyzer) VisitEndedStat(ctx *parser.EndedStatContext) any {
    return a.Visit(ctx.Stat())
}

func (a *Analyzer) VisitDefineStat(ctx *parser.DefineStatContext) any {
    id := ctx.ID().GetText()

    expr := ctx.Expr()
    exprType := a.Visit(expr).(result).exprType

    if t := ctx.Type_(); t != nil {
        declType := a.Visit(t).(result).exprType
        common := deduceCommonType(declType, exprType)

        if common == nil || *common != declType {
            panic(fmt.Sprintf("Cannot assign %v to %v", exprType, declType))
        }

        exprType = declType
    }


    if !a.table.AddSymbol(id, exprType) {
        panic("Identifier is already defined")
    }

    a.ExprTypes[expr] = exprType
    return result{exprType: exprType}
}

func (a *Analyzer) VisitMultiStat(ctx *parser.MultiStatContext) any {
    a.table.StartScope()

    for _, stat := range ctx.AllEndedStat() {
        a.Visit(stat)
    }

    if stat := ctx.Stat(); stat != nil {
        a.Visit(stat)
    }

    a.table.EndScope()
    return nil
}

func (a *Analyzer) VisitType(ctx *parser.TypeContext) any {
    return result{exprType: EType(strings.ToLower((ctx.TYPE().GetText())))}
}

func (a *Analyzer) VisitExprStat(ctx *parser.ExprStatContext) any {
    return a.Visit(ctx.Expr())
}

func (a *Analyzer) VisitNumExpr(ctx *parser.NumExprContext) any {
    num := ctx.NUM().GetText()

    if strings.ContainsRune(num, '.') {
        return result{exprType: "anyfloat"}
    }

    return result{exprType: "anyint"}
}

func (a *Analyzer) VisitIdExpr(ctx *parser.IdExprContext) any {
    id := ctx.ID().GetText()
    exprType, ok := a.table.FindSymbol(id)

    if !ok {
        panic("No such identifier")
    }

    return result{exprType: exprType}
}

func (a *Analyzer) VisitKeyExpr(ctx *parser.KeyExprContext) any {
    return result{exprType: "bool"}
}

func (a *Analyzer) VisitParenExpr(ctx *parser.ParenExprContext) any {
    return a.Visit(ctx.Expr())
}

func (a *Analyzer) VisitCastExpr(ctx *parser.CastExprContext) any {
    a.Visit(ctx.Expr())
    a.ExprTypes[ctx] = EType(strings.ToLower((ctx.TYPE().GetText())))

    return result{exprType: a.ExprTypes[ctx]}
}

func (a *Analyzer) VisitPrefixExpr(ctx *parser.PrefixExprContext) any {
    return a.Visit(ctx.Right)
}

func (a *Analyzer) VisitMulExpr(ctx *parser.MulExprContext) any {
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
    t := deduceCommonType(left.exprType, right.exprType)

    if t == nil {
        panic("Cannot mul/div/rem different types")
    }

    if strings.Contains(string(*t), "float") && ctx.Op.GetText() == "%" {
        panic("Cannot take rem with float")
    }

    return result{exprType: *t}
}

func (a *Analyzer) VisitAddExpr(ctx *parser.AddExprContext) any {
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
    t := deduceCommonType(left.exprType, right.exprType)

    if t == nil {
        panic("Cannot add/sub different types")
    }

    return result{exprType: *t}
}

func (a *Analyzer) VisitIfExpr(ctx *parser.IfExprContext) any {
    left := a.Visit(ctx.Left).(result)
    condition := a.Visit(ctx.Condition).(result)
    right := a.Visit(ctx.Right).(result)

    if condition.exprType != "bool" {
        panic("Condition must be a boolean expression")
    }

    t := deduceCommonType(left.exprType, right.exprType)

    if t == nil {
        panic("Both sides of if expression must be the same type")
    }

    return result{exprType: *t}
}
