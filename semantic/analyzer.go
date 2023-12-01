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
    hasAddress bool
    isNullable bool
    isDynamic bool
    right Type
}

func (r result) isNotNullable() result {
    if r.isNullable {
        panic("Expression cannot be nullable")
    }

    return r
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

func (a *Analyzer) VisitRefType(ctx *parser.RefTypeContext) any {
    base := a.Visit(ctx.Type_()).(result).exprType
    isDynamic := ctx.HASH() != nil
    isNullable := ctx.Ref.GetText() == "?"

    t := RefType{Base: base, IsDynamic: isDynamic, IsNullable: isNullable}
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

    a.ExprTypes[ctx] = exprType

    if ref, isRef := exprType.(RefType); isRef {
        return result{
            exprType: ref.Base,
            isDynamic: ref.IsDynamic,
            isNullable: ref.IsNullable,
            hasAddress: true,
        }
    }

    return result{exprType: exprType, hasAddress: true}
}

func (a *Analyzer) VisitKeyExpr(ctx *parser.KeyExprContext) any {
    var t Type

    if ctx.Key.GetText() == "null" {
        t = RefType{Base: nil, IsDynamic: true, IsNullable: true}
    } else {
        t = BaseType{Name: "Bool"}
    }

    return result{exprType: t}
}

func (a *Analyzer) VisitParenExpr(ctx *parser.ParenExprContext) any {
    return a.Visit(ctx.Expr())
}

func (a *Analyzer) VisitRefExpr(ctx *parser.RefExprContext) any {
    res := a.Visit(ctx.Left).(result)
    isDynamic := ctx.HASH() != nil
    isNullable := ctx.Op.GetText() == "?"

    if _, isRef := res.exprType.(RefType); isRef {
        panic("Cannot create reference to reference")
    }

    if !res.hasAddress {
        panic("Cannot make reference to non-variable")
    }

    if isDynamic && !res.isDynamic {
        panic("Cannot create dynamic reference to non-dynamic variable")
    }

    t := res.exprType
    t = RefType{Base: t, IsDynamic: isDynamic, IsNullable: isNullable}
    return result{exprType: t}
}

func (a *Analyzer) VisitNotNullExpr(ctx *parser.NotNullExprContext) any {
    res := a.Visit(ctx.Left).(result)

    if !res.isNullable {
        panic("May only assert nullable reference value to be not null")
    }

    if res.exprType.InferType() == nil {
        panic("Cannot assert null to be not null")
    }

    return result{exprType: res.exprType}
}

func (a *Analyzer) VisitAllocExpr(ctx *parser.AllocExprContext) any {
    res := a.Visit(ctx.Expr()).(result)
    isNullable := ctx.Ref.GetText() == "?"
    base, isBase := res.exprType.(BaseType)

    if !isBase {
        panic("Cannot alloc non-base type")
    }

    t := RefType{Base: base, IsDynamic: true, IsNullable: isNullable}
    return result{exprType: t}
}

func (a *Analyzer) VisitDeallocExpr(ctx *parser.DeallocExprContext) any {
    res := a.Visit(ctx.Expr()).(result)
    ref, isRef := res.exprType.(RefType)

    if !isRef || !ref.IsDynamic {
        panic("Can only dealloc dynamic reference type")
    }

    return result{exprType: NoneType{}}
}

func (a *Analyzer) VisitCastExpr(ctx *parser.CastExprContext) any {
    res := a.Visit(ctx.Expr()).(result).isNotNullable()

    if _, isBase := res.exprType.(BaseType); !isBase {
        panic("Cannot cast non-base type")
    }

    t := BaseType{Name: ctx.TYPE().GetText()}

    a.ExprTypes[ctx] = t
    return result{exprType: t}
}

func (a *Analyzer) VisitPrefixExpr(ctx *parser.PrefixExprContext) any {
    op := ctx.GetOp().GetText()
    exprType := a.Visit(ctx.Right).(result).isNotNullable().exprType

    if op == "!" && !IsBool(exprType) {
        panic("Cannot negate non-boolean value")
    }

    if (op == "+" || op == "-") && !IsNumeric(exprType) {
        panic("Cannot plus or minus non-numeric value")
    }

    return result{exprType: exprType}
}

func (a *Analyzer) VisitPowExpr(ctx *parser.PowExprContext) any {
    left := a.Visit(ctx.Left).(result).isNotNullable()
    right := a.Visit(ctx.Right).(result).isNotNullable()
    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Cannot pow different types")
    }

    if !IsNumeric(t) {
        panic("Cannot pow non-numeric types")
    }

    a.ExprTypes[ctx] = t
    return result{exprType: t}
}

func (a *Analyzer) VisitMulExpr(ctx *parser.MulExprContext) any {
    left := a.Visit(ctx.Left).(result).isNotNullable()
    right := a.Visit(ctx.Right).(result).isNotNullable()
    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Cannot mul/div/rem different types")
    }

    if !IsNumeric(t) {
        panic("Cannot mul/div/rem non-numeric types")
    }

    a.ExprTypes[ctx] = t
    return result{exprType: t}
}

func (a *Analyzer) VisitAddExpr(ctx *parser.AddExprContext) any {
    left := a.Visit(ctx.Left).(result).isNotNullable()
    right := a.Visit(ctx.Right).(result).isNotNullable()
    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Cannot add/sub different types")
    }

    if !IsNumeric(t) {
        panic("Cannot add/sub non-numeric types")
    }

    return result{exprType: t}
}

func (a *Analyzer) VisitShiftExpr(ctx *parser.ShiftExprContext) any {
    left := a.Visit(ctx.Left).(result).isNotNullable()
    right := a.Visit(ctx.Right).(result).isNotNullable()
    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Cannot shift different types")
    }

    if !IsInt(t) {
        panic("Cannot shift non-integer types")
    }

    return result{exprType: t}
}

func (a *Analyzer) VisitBitAndExpr(ctx *parser.BitAndExprContext) any {
    left := a.Visit(ctx.Left).(result).isNotNullable()
    right := a.Visit(ctx.Right).(result).isNotNullable()
    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Cannot bitwise-and different types")
    }

    if !IsInt(t) {
        panic("Cannot bitwise-and non-integer types")
    }

    return result{exprType: t}
}

func (a *Analyzer) VisitBitXorExpr(ctx *parser.BitXorExprContext) any {
    left := a.Visit(ctx.Left).(result).isNotNullable()
    right := a.Visit(ctx.Right).(result).isNotNullable()
    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Cannot bitwise-xor different types")
    }

    if !IsInt(t) {
        panic("Cannot bitwise-xor non-integer types")
    }

    return result{exprType: t}
}

func (a *Analyzer) VisitBitOrExpr(ctx *parser.BitOrExprContext) any {
    left := a.Visit(ctx.Left).(result).isNotNullable()
    right := a.Visit(ctx.Right).(result).isNotNullable()
    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Cannot bitwise-or different types")
    }

    if !IsInt(t) {
        panic("Cannot bitwise-or non-integer types")
    }

    return result{exprType: t}
}

func (a *Analyzer) VisitCompExpr(ctx *parser.CompExprContext) any {
    var leftType Type

    if comp, isComp := ctx.Left.(*parser.CompExprContext); isComp {
        leftType = a.Visit(comp).(result).right
    } else {
        leftType = a.Visit(ctx.Left).(result).isNotNullable().exprType
    }

    right := a.Visit(ctx.Right).(result).isNotNullable()
    t := DeduceType(leftType, right.exprType)

    if t == nil {
        panic("Cannot compare different types")
    }

    b := BaseType{Name: "Bool"}
    return result{exprType: b, right: right.exprType}
}

func (a *Analyzer) VisitCoalesceExpr(ctx *parser.CoalesceExprContext) any {
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Both sides of coalecse expression must be the same type")
    }

    return result{exprType: t, isNullable: right.isNullable}
}

func (a *Analyzer) VisitIfExpr(ctx *parser.IfExprContext) any {
    left := a.Visit(ctx.Left).(result).isNotNullable()
    condition := a.Visit(ctx.Condition).(result).isNotNullable()
    right := a.Visit(ctx.Right).(result).isNotNullable()

    if !IsBool(condition.exprType) {
        panic("Condition must be a boolean expression")
    }

    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Both sides of if expression must be the same type")
    }

    return result{exprType: t}
}
