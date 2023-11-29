package semantic

import (
	"slices"
	"strings"
)

type Type interface {
    FullType() string
    InferType() Type
    MayBeAssignedTo(Type) bool
}

func IsBool(t Type) bool {
    base, isBase := t.(BaseType)
    if !isBase { return false }

    return base.Name == "Bool"
}

func IsInt(t Type) bool {
    base, isBase := t.(BaseType)
    if !isBase { return false }

    return base.Name == "AnyNum" || strings.Contains(base.Name, "Int")
}

func IsNumeric(t Type) bool {
    base, isBase := t.(BaseType)
    if !isBase { return false }

    return !slices.Contains([]string {
        "Bool",
    }, base.Name)
}

func DeduceType(left, right Type) Type {
    if left.MayBeAssignedTo(right) { return right }
    if right.MayBeAssignedTo(left) { return left }
    return nil
}

type BaseType struct {
    Name string
}

func (b BaseType) FullType() string {
    return b.Name
}

func (b BaseType) InferType() Type {
    switch(b.Name) {
    case "AnyNum": return BaseType{Name: "Int"}
    case "AnyFloat": return BaseType{Name: "Float64"}
    default: return b
    }
}

func (b BaseType) MayBeAssignedTo(t Type) bool {
    base, isBase := t.(BaseType)
    if !isBase { return false }

    name := base.Name
    isFloat := strings.Contains(name, "Float")
    isNum := strings.Contains(name, "Int") || isFloat

    switch(b.Name) {
    case name: return true
    case "AnyNum": return isNum
    case "AnyFloat": return isFloat
    default: return false
    }
}

type RefType struct {
    Base Type
    IsNullable bool
}

func (r RefType) FullType() string {
    if r.Base == nil { return "NullType" }
    base := r.Base.FullType()

    if r.IsNullable { return base + "?" }
    return base + "!"
}

func (r RefType) InferType() Type {
    if r.Base == nil { return nil }
    base := r.Base.InferType()

    if base == nil { return nil }
    return RefType{Base: base, IsNullable: r.IsNullable}
}

func (r RefType) MayBeAssignedTo(t Type) bool {
    ref, isRef := t.(RefType)
    if !isRef { return false }

    if r.IsNullable == ref.IsNullable || ref.IsNullable {
        return r.Base.MayBeAssignedTo(ref.Base)
    }

    return false
}
