package code

import (
	"bytes"
	"fmt"
	"math/big"
	"strings"
	"zenith/semantic"
)

func getBase(num string) int64 {
    if strings.Contains(num, "0b") { return 2 }
    if strings.Contains(num, "0o") { return 8 }
    if strings.Contains(num, "0x") { return 16 }
    return 10
}

func getDigit(c byte) *big.Float {
    if c >= '0' && c <= '9' {
        return new(big.Float).SetInt64(int64(c - '0'))
    }

    c = bytes.ToLower([]byte{c})[0]
    return new(big.Float).SetInt64(int64(c - 'a' + 10))
}

func getCType(t semantic.Type) string {
    if ptr, isPtr := t.(semantic.PtrType); isPtr {
        return getCType(ptr.Base) + "*"
    }

    if slice, isSlice := t.(semantic.SliceType); isSlice {
        base := getCType(slice.Base)

        if slice.Size == 0 {
            return fmt.Sprintf("Zenith::Slice<%v>", base)
        }

        return fmt.Sprintf("Zenith::Array<%v>", base)
    }

    name := t.(semantic.BaseType).Name

    if name == "AnyNum" { return "Zenith::Int" }
    if name == "AnyFloat" { return "Zenith::Float32" }
    return "Zenith::" + name
}
