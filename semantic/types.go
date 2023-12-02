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


type NoneType struct{}

func (n NoneType) FullType() string { return "None" }
func (n NoneType) InferType() Type { return n }

func (n NoneType) MayBeAssignedTo(t Type) bool {
    return false
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


type PtrType struct {
    Base Type
    IsHeap bool
}

func (p PtrType) FullType() string {
    if p.Base == nil { return "NullType" }
    base := p.Base.FullType()

    if p.IsHeap { return "#" + base  }
    return "&" + base
}

func (p PtrType) InferType() Type {
    if p.Base == nil { return nil }
    base := p.Base.InferType()

    if base == nil { return nil }
    return PtrType{Base: base, IsHeap: p.IsHeap}
}

func (p PtrType) MayBeAssignedTo(t Type) bool {
    ptr, isPtr := t.(PtrType)
    if !isPtr { return false }

    if p.Base == nil { return true }
    if !p.IsHeap && ptr.IsHeap { return false }

    return p.Base.MayBeAssignedTo(ptr.Base)
}
