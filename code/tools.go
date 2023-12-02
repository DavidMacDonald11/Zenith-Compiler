package code

import (
	"bytes"
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

func getCId(id string) string {
    return "__ZenithId_" + id
}

func getCType(t semantic.Type) string {
    if ptr, isPtr := t.(semantic.PtrType); isPtr {
        return getCType(ptr.Base) + "*"
    }

    name := t.(semantic.BaseType).Name

    switch name {
    case "UInt8": return "uint8_t"
    case "UInt16": return "uint16_t"
    case "UInt32": return "uint32_t"
    case "UInt64": return "uint64_t"
    case "Int8": return "int8_t"
    case "Int16": return "int16_t"
    case "Int32": return "int32_t"
    case "Int64": return "int64_t"
    case "UInt": return "uint_fast32_t"
    case "AnyNum": fallthrough
    case "Int": return "int_fast32_t"
    case "Float32": return "float"
    case "AnyFloat": fallthrough
    case "Float64": return "double"
    case "Bool": return "bool"
    }

    return "void"
}
