package semantic

type Table struct {
    table []map[string]EType
}

func MakeTable() Table {
    t := Table{table: []map[string]EType{}}
    t.StartScope()
    return t
}

func (t *Table) StartScope() {
    t.table = append(t.table, make(map[string]EType))
}

func (t *Table) EndScope() {
    t.table = t.table[:len(t.table) - 1]
}

func (t *Table) FindSymbol(id string) (EType, bool) {
    for i := len(t.table) - 1; i >= 0; i-- {
        table := t.table[i]

        if exprType, ok := table[id]; ok {
            return exprType, true
        }
    }

    return "", false
}

func (t *Table) AddSymbol(id string, exprType EType) bool {
    if _, ok := t.FindSymbol(id); ok {
        return false
    }

    t.table[len(t.table) - 1][id] = exprType
    return true
}
