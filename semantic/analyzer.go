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

    ExprTypes map[parser.IExprContext]Type
}

type result struct {
    exprType Type
    isVar bool
}

func MakeAnalyzer() Analyzer {
    return Analyzer {
        table: MakeTable(),
        ExprTypes: map[parser.IExprContext]Type{},
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

        if !exprType.MayBeAssignedTo(declType) {
            e, d := exprType.FullType(), declType.FullType()
            panic(fmt.Sprintf("Cannot assign %v to %v", e, d))
        }

        exprType = declType
    } else {
        exprType = exprType.InferType()

        if exprType == nil {
            panic("Cannot infer type; requires a type annotation")
        }
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

func (a *Analyzer) VisitBaseType(ctx *parser.BaseTypeContext) any {
    t := BaseType{Name: ctx.TYPE().GetText()}
    return result{exprType: t}
}

func (a *Analyzer) VisitPtrType(ctx *parser.PtrTypeContext) any {
    base := a.Visit(ctx.Type_()).(result).exprType
    ptr := ctx.Ptr.GetText()

    t := PtrType{Base: base, Kind: PtrKind(ptr)}
    return result{exprType: t}
}

func (a *Analyzer) VisitExprStat(ctx *parser.ExprStatContext) any {
    return a.Visit(ctx.Expr())
}

func (a *Analyzer) VisitNumExpr(ctx *parser.NumExprContext) any {
    num := ctx.NUM().GetText()
    var t Type

    if strings.ContainsRune(num, '.') {
        t = BaseType{Name: "AnyFloat"}
    } else {
        t = BaseType{Name: "AnyNum"}
    }

    return result{exprType: t}
}

func (a *Analyzer) VisitIdExpr(ctx *parser.IdExprContext) any {
    id := ctx.ID().GetText()
    exprType := a.table.FindSymbol(id)

    if exprType == nil {
        panic("No such identifier")
    }

    return result{exprType: exprType, isVar: true}
}

func (a *Analyzer) VisitKeyExpr(ctx *parser.KeyExprContext) any {
    var t Type

    if ctx.Key.GetText() == "null" {
        t = PtrType{Base: nil}
    } else {
        t = BaseType{Name: "Bool"}
    }

    return result{exprType: t}
}

func (a *Analyzer) VisitParenExpr(ctx *parser.ParenExprContext) any {
    return a.Visit(ctx.Expr())
}

func (a *Analyzer) VisitCastExpr(ctx *parser.CastExprContext) any {
    a.Visit(ctx.Expr())
    t := BaseType{Name: ctx.TYPE().GetText()}

    a.ExprTypes[ctx] = t
    return result{exprType: t}
}

func (a *Analyzer) VisitPtrExpr(ctx *parser.PtrExprContext) any {
    res := a.Visit(ctx.Right).(result)

    if !res.isVar {
        panic("Cannot make pointer to non-variable")
    }

    kind := ctx.Op.GetText()

    if kind[0] == '$' {
        panic("Cannot take ownership of variable")
    }

    t := PtrType{Base: res.exprType, Kind: PtrKind(kind)}
    return result{exprType: t}
}

func (a *Analyzer) VisitPrefixExpr(ctx *parser.PrefixExprContext) any {
    op := ctx.GetOp().GetText()
    exprType := a.Visit(ctx.Right).(result).exprType

    if op == "!" && !exprType.IsBool() {
        panic("Cannot negate non-boolean value")
    }

    if (op == "+" || op == "-") && !exprType.IsNumeric() {
        panic("Cannot plus or minus non-numeric value")
    }

    return result{exprType: exprType}
}

func (a *Analyzer) VisitMulExpr(ctx *parser.MulExprContext) any {
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Cannot mul/div/rem different types")
    }

    if !t.IsNumeric() {
        panic("Cannot mul/div/rem non-numeric types")
    }

    a.ExprTypes[ctx] = t
    return result{exprType: t}
}

func (a *Analyzer) VisitAddExpr(ctx *parser.AddExprContext) any {
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Cannot add/sub different types")
    }

    if !t.IsNumeric() {
        panic("Cannot add/sub non-numeric types")
    }

    return result{exprType: t}
}

func (a *Analyzer) VisitIfExpr(ctx *parser.IfExprContext) any {
    left := a.Visit(ctx.Left).(result)
    condition := a.Visit(ctx.Condition).(result)
    right := a.Visit(ctx.Right).(result)

    if !condition.exprType.IsBool() {
        panic("Condition must be a boolean expression")
    }

    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Both sides of if expression must be the same type")
    }

    return result{exprType: t}
}
