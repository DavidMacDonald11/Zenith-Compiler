package semantic

import (
	"slices"
	"strings"
)

type EType string

func numerics() []EType {
    return []EType {
        "uint8", "uint16", "uint32", "uint64", "uint",
        "int8", "int16", "int32", "int64", "int",
        "float32", "float64",
        "anyint", "anyfloat",
    }
}

func nonNumerics() []EType {
    return []EType {
        "bool",
    }
}

func deduceCommonType(left, right EType) *EType {
    numericTypes := numerics()
    leftIsNumeric := slices.Contains(numericTypes, left)
    rightIsNumeric := slices.Contains(numericTypes, right)

    switch {
    case left == right: return &left
    case left == "anyint" && rightIsNumeric: return &right
    case right == "anyint" && leftIsNumeric: return &left
    case left == "anyfloat":
        if strings.Contains(string(right), "int") || !rightIsNumeric {
            return nil
        }

        return &right
    case right == "anyfloat":
        if strings.Contains(string(left), "int") || !leftIsNumeric {
            return nil
        }

        return &left
    default: return nil
    }
}
