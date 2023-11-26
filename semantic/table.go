package semantic

type Table struct {
    table []map[string]Type
}

func MakeTable() Table {
    t := Table{table: []map[string]Type{}}
    t.StartScope()
    return t
}

func (t *Table) StartScope() {
    t.table = append(t.table, make(map[string]Type))
}

func (t *Table) EndScope() {
    t.table = t.table[:len(t.table) - 1]
}

func (t *Table) FindSymbol(id string) Type {
    for i := len(t.table) - 1; i >= 0; i-- {
        table := t.table[i]

        if exprType, ok := table[id]; ok {
            return exprType
        }
    }

    return nil
}

func (t *Table) AddSymbol(id string, exprType Type) bool {
    if t.FindSymbol(id) != nil {
        return false
    }

    t.table[len(t.table) - 1][id] = exprType
    return true
}
