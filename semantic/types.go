package semantic

import (
	"slices"
	"strings"
)

func deduceCommonType(left, right string) *string {
    nonnumerics := []string{"bool"}
    leftIsNumeric := !slices.Contains(nonnumerics, left)
    rightIsNumeric := !slices.Contains(nonnumerics, right)

    switch {
    case left == right: return &left
    case left == "anyint" && rightIsNumeric: return &right
    case right == "anyint" && leftIsNumeric: return &left
    case left == "anyfloat":
        if strings.Contains(right, "int") || !rightIsNumeric {
            return nil
        }

        return &right
    case right == "anyfloat":
        if strings.Contains(left, "int") || !leftIsNumeric {
            return nil
        }

        return &left
    default: return nil
    }
}
