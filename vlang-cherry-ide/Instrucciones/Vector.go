package instrucciones

import "main/tiposDeDato"

// Implementar el  ValorInterno

type TipoVector struct {
	*TipoObjeto

	InternalValor []tiposDeDato.ValorInterno
	CurrentIndex  int
	ItemType      string
	FullType      string
	SizeValue     *tiposDeDato.ValorEntero
	IsEmpty       *tiposDeDato.ValorBool
}

func (v TipoVector) Value() interface{} {
	return v
}

func (v TipoVector) Type() string {
	return v.FullType
}

func (v TipoVector) Size() int {
	return len(v.InternalValor)
}

func (v TipoVector) ValidIndex(index int) bool {

	if index < 0 || index >= len(v.InternalValor) {
		return false
	}

	return true

}

func (v TipoVector) Get(index int) tiposDeDato.ValorInterno {
	return v.InternalValor[index]
}

func (v *TipoVector) Next() bool {
	if v.CurrentIndex < len(v.InternalValor) {
		v.CurrentIndex++
		return true
	}
	return false
}

func (v *TipoVector) Current() tiposDeDato.ValorInterno {
	return v.InternalValor[v.CurrentIndex]
}

func (v *TipoVector) Reset() {
	v.CurrentIndex = 0
}

func (v *TipoVector) Copy() tiposDeDato.ValorInterno {

	internalCopy := make([]tiposDeDato.ValorInterno, len(v.InternalValor))

	for i, item := range v.InternalValor {
		internalCopy[i] = item.Copy()
	}

	return NewTipoVector(internalCopy, v.FullType, v.ItemType)

}

func (v *TipoVector) updateProps() {

	v.SizeValue.InternalValor = len(v.InternalValor)
	v.IsEmpty.InternalValor = len(v.InternalValor) == 0

}

func NewTipoVector(vectorItems []tiposDeDato.ValorInterno, fullType, itemType string) *TipoVector {
	vector := &TipoVector{
		TipoObjeto:    &TipoObjeto{},
		InternalValor: vectorItems,
		CurrentIndex:  0,
		ItemType:      itemType,
		FullType:      fullType,
		SizeValue:     &tiposDeDato.ValorEntero{InternalValor: len(vectorItems)},
		IsEmpty:       &tiposDeDato.ValorBool{InternalValor: len(vectorItems) == 0},
	}

	AgregarVectorNativo(vector)

	return vector
}

var TipoVectorVacio = &TipoVector{
	TipoObjeto:    &TipoObjeto{},
	InternalValor: []tiposDeDato.ValorInterno{},
	CurrentIndex:  0,
	ItemType:      tiposDeDato.TIPO_ANY,
	FullType:      "[" + tiposDeDato.TIPO_ANY + "]",
	SizeValue:     &tiposDeDato.ValorEntero{InternalValor: 0},
	IsEmpty:       &tiposDeDato.ValorBool{InternalValor: true},
}

type VectorItemReference struct {
	Vector *TipoVector
	Index  int
	Value  tiposDeDato.ValorInterno
}
