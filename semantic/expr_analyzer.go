package semantic

import (
	"strings"
	"zenith/parser"
)

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
        t = PtrType{Base: nil, IsHeap: true}
    } else {
        t = BaseType{Name: "Bool"}
    }

    return result{exprType: t}
}

func (a *Analyzer) VisitParenExpr(ctx *parser.ParenExprContext) any {
    return a.Visit(ctx.Expr())
}

func (a *Analyzer) VisitPostfixExpr(ctx *parser.PostfixExprContext) any {
    exprType := a.Visit(ctx.Expr()).(result).exprType

    if _, isPtr := exprType.(PtrType); !isPtr {
        panic("May only check if pointer is null")
    }

    t := BaseType{Name: "Bool"}
    return result{exprType: t}
}

func (a *Analyzer) VisitAllocExpr(ctx *parser.AllocExprContext) any {
    expr := a.Visit(ctx.Expr()).(result)
    t := PtrType{Base: expr.exprType, IsHeap: true}
    return result{exprType: t}
}

func (a *Analyzer) VisitDeallocExpr(ctx *parser.DeallocExprContext) any {
    expr := a.Visit(ctx.Expr()).(result)
    ptr, isPtr := expr.exprType.(PtrType)

    if !isPtr || !ptr.IsHeap {
        panic("May only dealloc dynamic pointer")
    }

    return result{exprType: NoneType{}}
}

func (a *Analyzer) VisitCastExpr(ctx *parser.CastExprContext) any {
    expr := a.Visit(ctx.Expr()).(result)

    if _, isBase := expr.exprType.(BaseType); !isBase {
        panic("Cannot cast non-base type")
    }

    t := BaseType{Name: ctx.TYPE().GetText()}

    a.ExprTypes[ctx] = t
    return result{exprType: t}
}

func (a *Analyzer) VisitPtrExpr(ctx *parser.PtrExprContext) any {
    expr := a.Visit(ctx.Expr()).(result)
    op := ctx.Op.GetText()

    if op != "@" {
        if !expr.isVar {
            panic("Cannot create pointer to non-variable")
        }

        t := PtrType{Base: expr.exprType, IsHeap: false}
        return result{exprType: t}
    }

    ptr, isPtr := expr.exprType.(PtrType)

    if !isPtr {
        panic("Cannot dereference non-pointer")
    }

    if ptr.Base == nil {
        panic("Cannot dereference null")
    }

    t := ptr.Base
    return result{exprType: t}
}

func (a *Analyzer) VisitPrefixExpr(ctx *parser.PrefixExprContext) any {
    op := ctx.GetOp().GetText()
    exprType := a.Visit(ctx.Expr()).(result).exprType

    if op == "!" && !IsBool(exprType) {
        panic("Cannot negate non-boolean value")
    }

    if (op == "+" || op == "-") && !IsNumeric(exprType) {
        panic("Cannot plus or minus non-numeric value")
    }

    return result{exprType: exprType}
}

func (a *Analyzer) VisitPowExpr(ctx *parser.PowExprContext) any {
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
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
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
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
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
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
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
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
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
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
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
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
    left := a.Visit(ctx.Left).(result)
    right := a.Visit(ctx.Right).(result)
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
        leftType = a.Visit(comp).(result).compRightType
    } else {
        leftType = a.Visit(ctx.Left).(result).exprType
    }

    right := a.Visit(ctx.Right).(result)
    t := DeduceType(leftType, right.exprType)

    if t == nil {
        panic("Cannot compare different types")
    }

    b := BaseType{Name: "Bool"}
    return result{exprType: b, compRightType: right.exprType}
}

func (a *Analyzer) VisitIfExpr(ctx *parser.IfExprContext) any {
    left := a.Visit(ctx.Left).(result)
    condition := a.Visit(ctx.Condition).(result)
    right := a.Visit(ctx.Right).(result)

    if !IsBool(condition.exprType) {
        panic("Condition must be a boolean expression")
    }

    t := DeduceType(left.exprType, right.exprType)

    if t == nil {
        panic("Both sides of if expression must be the same type")
    }

    return result{exprType: t}
}
