package instrucciones

import "main/tiposDeDato"

type TipoPuntero struct {
	VariableAsociada *Variable
}

func (p TipoPuntero) Value() interface{} {
	return p
}

func (p TipoPuntero) Type() string {
	return tiposDeDato.TIPO_PUNTERO
}

func (p TipoPuntero) Copy() tiposDeDato.ValorInterno {
	return p
}
