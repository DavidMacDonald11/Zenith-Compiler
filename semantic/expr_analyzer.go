package semantic

import (
	"fmt"
	"math/big"
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

    constValue, _ := newConstValue(0).SetString(GetNum(num))
    return result{exprType: t, constValue: constValue}
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
    constValue := newConstValue(0)

    if ctx.Key.GetText() == "null" {
        t = PtrType{Base: nil, IsHeap: true}
    } else {
        t = BaseType{Name: "Bool"}

        if ctx.Key.GetText() == "true" {
            constValue.SetInt64(1)
        }
    }

    return result{exprType: t, constValue: constValue}
}

func (a *Analyzer) VisitInitExpr(ctx *parser.InitExprContext) any {
    t := a.Visit(ctx.Type).(result).exprType
    types := a.Visit(ctx.InitArgs()).([]Type)

    if !t.MayInitWith(types) {
        panic("Cannot initialize this type with this initializer")
    }

    a.ExprTypes[ctx] = t
    return result{exprType: t}
}

func (a *Analyzer) VisitInitArgs(ctx *parser.InitArgsContext) any {
    exprs := ctx.AllExpr()
    types := make([]Type, len(exprs))

    for i, expr := range exprs {
        types[i] = a.Visit(expr).(result).exprType
    }

    return types
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
    a.ExprTypes[ctx.Expr()] =  expr.exprType

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

    if expr.constValue != nil {
        expr.constValue = castToBaseType(t, expr.constValue)
    }

    a.ExprTypes[ctx] = t
    return result{exprType: t, constValue: expr.constValue}
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
    res := a.Visit(ctx.Expr()).(result)
    exprType := res.exprType
    constValue := res.constValue

    if op == "!" {
        if !IsBool(exprType) {
            panic("Cannot negate non-boolean value")
        }

        if constValue != nil {
            if constValue.Cmp(big.NewFloat(0)) == 0 {
                constValue.SetInt64(1)
            } else {
                constValue.SetInt64(0)
            }
        }
    }

    if (op == "+" || op == "-") {
        if !IsNumeric(exprType) {
            panic("Cannot plus or minus non-numeric value")
        }

        if constValue != nil && op == "-" {
            constValue.Neg(constValue)
        }
    }

    return result{exprType: exprType, constValue: constValue}
}

func (a *Analyzer) VisitPowExpr(ctx *parser.PowExprContext) any {
    _, _, t := a.binaryOpType(ctx.Left, ctx.Right, "pow")

    if !IsNumeric(t) {
        panic("Cannot pow non-numeric types")
    }

    a.ExprTypes[ctx] = t
    return result{exprType: t}
}

func (a *Analyzer) VisitMulExpr(ctx *parser.MulExprContext) any {
    left, right, t := a.binaryOpType(ctx.Left, ctx.Right, "mul/div/rem")

    if !IsNumeric(t) {
        panic("Cannot mul/div/rem non-numeric types")
    }

    var constValue *big.Float

    if left.constValue != nil && right.constValue != nil {
        constValue = newConstValue(0)

        switch ctx.GetOp().GetText() {
        case "*":
            constValue.Mul(left.constValue, right.constValue)
        case "/":
            if IsInt(t) && right.constValue.Sign() == 0 {
                panic("Cannot integer divide by 0")
            }

            constValue.Quo(left.constValue, right.constValue)
        case "%":
            if IsInt(t) {
                if right.constValue.Sign() == 0 {
                    panic("Cannot integer divide by 0")
                }

                l, r := big.NewInt(0), big.NewInt(0)

                left.constValue.Int(l)
                right.constValue.Int(r)

                l.QuoRem(l, r, r)
                constValue.SetInt(r)
            } else {
                constValue = nil
            }
        }
    }

    a.ExprTypes[ctx] = t
    return result{exprType: t, constValue: constValue}
}

func (a *Analyzer) VisitAddExpr(ctx *parser.AddExprContext) any {
    left, right, t := a.binaryOpType(ctx.Left, ctx.Right, "add/sub")

    if !IsNumeric(t) {
        panic("Cannot add/sub non-numeric types")
    }

    var constValue *big.Float

    if left.constValue != nil && right.constValue != nil {
        constValue = newConstValue(0)

        if ctx.Op.GetText() == "+" {
            constValue.Add(left.constValue, right.constValue)
        } else {
            constValue.Sub(left.constValue, right.constValue)
        }
    }

    return result{exprType: t, constValue: constValue}
}

func (a *Analyzer) VisitShiftExpr(ctx *parser.ShiftExprContext) any {
    left, right, t := a.binaryOpType(ctx.Left, ctx.Right, "shift")

    if !IsInt(t) {
        panic("Cannot shift non-integer types")
    }

    var constValue *big.Float

    if left.constValue != nil && right.constValue != nil {
        constValue = newConstValue(0)
        l, r := big.NewInt(0), big.NewInt(0)

        left.constValue.Int(l)
        right.constValue.Int(r)

        if r.Sign() < 0 {
            panic("Right side of shift cannot be negative")
        }

        if ctx.Op.GetText() == "<<" {
            constValue.SetInt(l.Lsh(l, uint(r.Uint64())))
        } else {
            constValue.SetInt(l.Rsh(l, uint(r.Uint64())))
        }
    }

    return result{exprType: t, constValue: constValue}
}

func (a *Analyzer) VisitBitAndExpr(ctx *parser.BitAndExprContext) any {
    left, right, t := a.binaryOpType(ctx.Left, ctx.Right, "bitwise-and")

    if !IsInt(t) {
        panic("Cannot bitwise-and non-integer types")
    }

    var constValue *big.Float

    if left.constValue != nil && right.constValue != nil {
        constValue = newConstValue(0)
        l, r := big.NewInt(0), big.NewInt(0)

        left.constValue.Int(l)
        right.constValue.Int(r)

        constValue.SetInt(l.And(l, r))
    }

    return result{exprType: t, constValue: constValue}
}

func (a *Analyzer) VisitBitXorExpr(ctx *parser.BitXorExprContext) any {
    left, right, t := a.binaryOpType(ctx.Left, ctx.Right, "bitwise-xor")

    if !IsInt(t) {
        panic("Cannot bitwise-xor non-integer types")
    }

    var constValue *big.Float

    if left.constValue != nil && right.constValue != nil {
        constValue = newConstValue(0)
        l, r := big.NewInt(0), big.NewInt(0)

        left.constValue.Int(l)
        right.constValue.Int(r)

        constValue.SetInt(l.Xor(l, r))
    }

    return result{exprType: t, constValue: constValue}
}

func (a *Analyzer) VisitBitOrExpr(ctx *parser.BitOrExprContext) any {
    left, right, t := a.binaryOpType(ctx.Left, ctx.Right, "bitwise-or")

    if !IsInt(t) {
        panic("Cannot bitwise-or non-integer types")
    }

    var constValue *big.Float

    if left.constValue != nil && right.constValue != nil {
        constValue = newConstValue(0)
        l, r := big.NewInt(0), big.NewInt(0)

        left.constValue.Int(l)
        right.constValue.Int(r)

        constValue.SetInt(l.Or(l, r))
    }

    return result{exprType: t, constValue: constValue}
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

    // TODO calculate constValue
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

    var constValue *big.Float

    if left.constValue != nil && condition.constValue != nil && right.constValue != nil {
        if condition.constValue.Sign() == 0 {
            constValue = right.constValue
        } else {
            constValue = left.constValue
        }
    }

    return result{exprType: t, constValue: constValue}
}

func (a *Analyzer) binaryOpType(left, right exprCtx, op string) (result, result, Type) {
    l := a.Visit(left).(result)
    r := a.Visit(right).(result)
    t := DeduceType(l.exprType, r.exprType)

    if t == nil {
        panic(fmt.Sprintf("Cannot %v different types", op))
    }

    return l, r, t
}
