package tipos

type ValorCaracter struct {
	InternalValor string
}

func (s ValorCaracter) Value() interface{} {
	return s.InternalValor
}

func (s ValorCaracter) Type() string {
	return TIPO_CARACTER
}

func (s ValorCaracter) Copy() ValorInterno {
	return &ValorCaracter{s.InternalValor}
}
