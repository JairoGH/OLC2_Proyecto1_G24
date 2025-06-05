package tipos

type ValorEntero struct {
	InternalValor int
}

func (i ValorEntero) Value() interface{} {
	return i.InternalValor
}

func (i ValorEntero) Type() string {
	return TIPO_ENTERO 
}

func (i ValorEntero) Copy() ValorInterno {
	return &ValorEntero{i.InternalValor}
}
