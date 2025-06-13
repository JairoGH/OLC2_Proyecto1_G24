package instrucciones

import (
	"main/tiposDeDato"
)

type TipoMatriz struct {
	*TipoVector
}

func (v *TipoMatriz) Copy() tiposDeDato.ValorInterno {

	internalCopy := make([]tiposDeDato.ValorInterno, len(v.InternalValor))

	for i, item := range v.InternalValor {
		internalCopy[i] = item.Copy()
	}

	return NewTipoMatriz(internalCopy, v.FullType, v.ItemType)

}

func (v *TipoMatriz) IndicesValidos(indexes []int) bool {

	// check if indexes are valid
	pivot := v.TipoVector

	for i, index := range indexes {
		if index < 0 || index >= pivot.Size() {
			return false
		}

		item := pivot.Get(index)

		// vector, Matriz or value
		switch s := item.(type) {
		case *TipoVector:
			pivot = s

			if i == len(indexes)-1 {
				return true
			}

		case *TipoMatriz:
			pivot = s.TipoVector

			if i == len(indexes)-1 {
				return true
			}

		default:
			if i != len(indexes)-1 {
				return false
			} else {
				return true
			}
		}
	}

	return false
}

func (v *TipoMatriz) Get(index []int) tiposDeDato.ValorInterno {

	// check if indexes are valid
	if !v.IndicesValidos(index) {
		return nil
	}

	pivot := v.TipoVector

	for i := 0; i < len(index); i++ {
		item := pivot.Get(index[i])

		// vector, Matriz or value
		switch s := item.(type) {
		case *TipoVector:
			pivot = s

			if i == len(index)-1 {
				return pivot
			}

		case *TipoMatriz:
			pivot = s.TipoVector

			if i == len(index)-1 {
				return pivot
			}
		default:
			return item
		}
	}

	return nil
}

func (v *TipoMatriz) Set(index []int, value tiposDeDato.ValorInterno) bool {

	// check if indexes are valid
	if !v.IndicesValidos(index) {
		return false
	}

	pivot := v.TipoVector

	for i := 0; i < len(index); i++ {
		item := pivot.Get(index[i])

		// vector, Matriz or value
		switch s := item.(type) {
		case *TipoVector:
			pivot = s
		case *TipoMatriz:
			pivot = s.TipoVector
		default:
			if i == len(index)-1 {
				pivot.InternalValor[index[i]] = value
				return true
			}
		}
	}

	return false
}

func NewTipoMatriz(vectorItems []tiposDeDato.ValorInterno, fullType, itemType string) *TipoMatriz {
	vector := &TipoVector{
		InternalValor: vectorItems,
		CurrentIndex:  0,
		ItemType:      itemType,
		FullType:      fullType,
		SizeValue:     &tiposDeDato.ValorEntero{InternalValor: len(vectorItems)},
		IsEmpty:       &tiposDeDato.ValorBool{InternalValor: len(vectorItems) == 0},
	}

	// remove builtins from vector value
	EliminarIntegradosDeVector(vectorItems)

	return &TipoMatriz{
		TipoVector: vector,
	}
}

func EliminarIntegradosDeVector(vectorItems []tiposDeDato.ValorInterno) {
	for i := 0; i < len(vectorItems); i++ {
		if item, ok := vectorItems[i].(*TipoVector); ok {
			item.TipoObjeto.InternalScope.Reset()
		} else {
			break
		}
	}
}

type ElementoMatriz struct {
	Matriz *TipoMatriz
	Index  []int
	Value  tiposDeDato.ValorInterno
}

// Nuevo tipo para matrices multidimensionales con sintaxis [][]tipo
type TipoMatrizMulti struct {
    InternalValor [][]tiposDeDato.ValorInterno
    FullType      string // "[][]int", "[][]string", etc.
    ItemType      string // "[]int", "[]string", etc. (tipo de cada fila)
    BaseType      string // "int", "string", etc. (tipo de elementos)
    Filas         int
    Columnas      []int // Cada fila puede tener diferente tamaño
}

func (t *TipoMatrizMulti) Type() string {
    return t.FullType
}

func (t *TipoMatrizMulti) Value() interface{} {
    return t.InternalValor
}

func (t *TipoMatrizMulti) Copy() tiposDeDato.ValorInterno {
    newMatrix := make([][]tiposDeDato.ValorInterno, len(t.InternalValor))
    for i, row := range t.InternalValor {
        newMatrix[i] = make([]tiposDeDato.ValorInterno, len(row))
        for j, item := range row {
            newMatrix[i][j] = item.Copy()
        }
    }
    
    return &TipoMatrizMulti{
        InternalValor: newMatrix,
        FullType:      t.FullType,
        ItemType:      t.ItemType,
        BaseType:      t.BaseType,
        Filas:         t.Filas,
        Columnas:      append([]int{}, t.Columnas...),
    }
}

func (t *TipoMatrizMulti) ValidIndex(fila, columna int) bool {
    // Verificar que la fila esté en rango
    if fila < 0 || fila >= len(t.InternalValor) {
        return false
    }
    
    // Verificar que la columna esté en rango para esa fila específica
    if columna < 0 || columna >= len(t.InternalValor[fila]) {
        return false
    }
    
    return true
}

func (t *TipoMatrizMulti) Get(fila, columna int) tiposDeDato.ValorInterno {
    if !t.ValidIndex(fila, columna) {
        return &tiposDeDato.ValorEntero{InternalValor: -1}
    }
    return t.InternalValor[fila][columna]
}

func (t *TipoMatrizMulti) Set(fila, columna int, valor tiposDeDato.ValorInterno) {
    if t.ValidIndex(fila, columna) {
        t.InternalValor[fila][columna] = valor
    }
}

func NewTipoMatrizMulti(filas [][]tiposDeDato.ValorInterno, fullType, itemType, baseType string) *TipoMatrizMulti {
    columnas := make([]int, len(filas))
    for i, row := range filas {
        columnas[i] = len(row)
    }
    
    return &TipoMatrizMulti{
        InternalValor: filas,
        FullType:      fullType,
        ItemType:      itemType,
        BaseType:      baseType,
        Filas:         len(filas),
        Columnas:      columnas,
    }
}

// Tipo para referencias a elementos de matriz multidimensional
type MatrizMultiItemReference struct {
    Matriz *TipoMatrizMulti
    Fila   int
    Columna int
    Value  tiposDeDato.ValorInterno
}