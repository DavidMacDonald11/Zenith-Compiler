package code

import (
	"fmt"
	"zenith/semantic"
)

func getCType(t semantic.Type) string {
    if ptr, isPtr := t.(semantic.PtrType); isPtr {
        return getCType(ptr.Base) + "*"
    }

    if slice, isSlice := t.(semantic.SliceType); isSlice {
        base := getCType(slice.Base)

        if slice.Size == 0 {
            return fmt.Sprintf("Zenith::Slice<%v>", base)
        }

        return fmt.Sprintf("Zenith::Array<%v, %v>", base, slice.Size)
    }

    name := t.(semantic.BaseType).Name

    if name == "AnyNum" { return "Zenith::Int" }
    if name == "AnyFloat" { return "Zenith::Float32" }
    return "Zenith::" + name
}
