package semantic

import (
	"bytes"
	"math/big"
	"strings"
)

func GetNum(s string) string {
    str := strings.ReplaceAll(s, "_", "")
    b := getNumBase(str)

    base := new(big.Float).SetInt64(b)
    num := new(big.Float).SetPrec(64)
    i := 0

    if b != 10 {
        i = 2
    }

    for ; i < len(str) && str[i] != '.'; i++ {
        num.Mul(num, base)
        num.Add(num, getDigit(str[i]))
    }

    inv := new(big.Float).SetPrec(64).SetInt64(1)

    for i++; i < len(str); i++ {
        inv.Quo(inv, base)

        digit := getDigit(str[i])
        digit.Mul(digit, inv)
        num.Add(num, digit)
    }

    isFloat := strings.ContainsRune(str, '.')
    str = num.Text('g', 64)

    if isFloat && num.IsInt() {
        return str + ".0"
    }

    return str
}

func getNumBase(num string) int64 {
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

func newConstValue(n float64) *big.Float {
    return big.NewFloat(n).SetPrec(128)
}

func castToBaseType(t BaseType, n *big.Float) *big.Float {
    i := big.NewInt(0)
    n.Int(i)

    if t.Name == "Bool" {
        if i.Sign() != 0 { return newConstValue(1) }
        return newConstValue(0)
    }

    if strings.Contains("Int", t.Name) {
        return n.SetInt(i)
    }

    return n
}
