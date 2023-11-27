package semantic

import (
	"slices"
	"strings"
)

type Type interface {
    FullType() string
    InferType() Type
    MayBeAssignedTo(Type) bool
    IsBool() bool
    IsNumeric() bool
}

type BaseType struct {
    Name string
}

func (b BaseType) FullType() string {
    return b.Name
}

func (b BaseType) IsNumeric() bool {
    return !slices.Contains([]string {
        "Bool",
    }, b.Name)
}

func (b BaseType) InferType() Type {
    switch(b.Name) {
    case "AnyNum": return BaseType{Name: "Int"}
    case "AnyFloat": return BaseType{Name: "Float64"}
    default: return b
    }
}

func (b BaseType) MayBeAssignedTo(t Type) bool {
    if _, isPtr := t.(PtrType); isPtr {
        return false
    }

    name := t.(BaseType).Name
    isFloat := strings.Contains(name, "Float")
    isNum := strings.Contains(name, "Int") || isFloat

    switch(b.Name) {
    case name: return true
    case "AnyNum": return isNum
    case "AnyFloat": return isFloat
    default: return false
    }
}

func (b BaseType) IsBool() bool { return b.Name == "Bool" }

type PtrType struct {
    Base Type
    IsOwn bool
}

func (p PtrType) FullType() string {
    if p.Base == nil { return "NullType" }
    base := p.Base.FullType()

    if p.IsOwn { return "$" + base }
    return "&" + base
}

func (p PtrType) InferType() Type {
    if p.Base == nil { return nil }
    base := p.Base.InferType()

    if base == nil { return nil }
    return PtrType{Base: base, IsOwn: p.IsOwn}
}

func (p PtrType) MayBeAssignedTo(t Type) bool {
    ptr, isPtr := t.(PtrType)
    if !isPtr { return false }

    if p.IsOwn == ptr.IsOwn {
        return p.Base.MayBeAssignedTo(ptr.Base)
    }

    return false
}

func (p PtrType) IsBool() bool { return false }
func (p PtrType) IsNumeric() bool { return false }

func DeduceType(left, right Type) Type {
    if left.MayBeAssignedTo(right) { return right }
    if right.MayBeAssignedTo(left) { return left }
    return nil
}
