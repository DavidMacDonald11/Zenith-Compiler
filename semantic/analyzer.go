package semantic

import (
	"fmt"
	"math/big"
	"zenith/parser"

	"github.com/antlr4-go/antlr/v4"
)

type exprCtx = parser.IExprContext

type Analyzer struct {
    parser.BaseZenithParserVisitor
    table Table

    ExprTypes map[exprCtx]Type
}

type result struct {
    exprType Type
    isVar bool
    constValue *big.Float
    compRightType Type
}

func MakeAnalyzer() Analyzer {
    return Analyzer {
        table: MakeTable(),
        ExprTypes: map[exprCtx]Type{},
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

    if t := ctx.Type; t != nil {
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

func (a *Analyzer) VisitExprStat(ctx *parser.ExprStatContext) any {
    return a.Visit(ctx.Expr())
}

func (a *Analyzer) VisitFullType(ctx *parser.FullTypeContext) any {
    if ctx.PartType() != nil { return a.Visit(ctx.PartType()) }
    return a.Visit(ctx.PtrType())
}

func (a *Analyzer) VisitPtrType(ctx *parser.PtrTypeContext) any {
    base := a.Visit(ctx.Type).(result).exprType
    isHeap := ctx.Ptr.GetText() == "#"

    t := PtrType{Base: base, IsHeap: isHeap}
    return result{exprType: t}
}

func (a *Analyzer) VisitBaseType(ctx *parser.BaseTypeContext) any {
    t := BaseType{Name: ctx.TYPE().GetText()}
    return result{exprType: t}
}

func (a *Analyzer) VisitSliceType(ctx *parser.SliceTypeContext) any {
    var size uint

    if ctx.Expr() != nil {
        res := a.Visit(ctx.Expr()).(result)

        if !IsInt(res.exprType) {
            panic("Array size must be an integer")
        }

        if res.constValue == nil {
            panic("Array size must be constant expression")
        }

        if res.constValue.Sign() < 1 {
            panic("Array size must be at least 1")
        }

        n, _ := res.constValue.Uint64()
        size = uint(n)
    }

    base := a.Visit(ctx.Type).(result).exprType
    t := SliceType{Base: base, Size: size}
    return result{exprType: t}
}
