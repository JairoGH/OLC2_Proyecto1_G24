package tiposDeDato

type ValorCadena struct {
	InternalValor string
}

func (s ValorCadena) Value() interface{} {
	return s.InternalValor
}

func (s ValorCadena) Type() string {
	return TIPO_CADENA  
}

func (s ValorCadena) Copy() ValorInterno {
	return &ValorCadena{InternalValor: s.InternalValor}
}
