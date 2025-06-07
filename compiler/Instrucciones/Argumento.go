package instrucciones

import (
	"compiler/tiposDeDato"

	"github.com/antlr4-go/antlr/v4"
)

type Argumento struct {
	Name            string
	Value           tiposDeDato.ValorInterno
	PassByReference bool
	Token           antlr.Token
	//VariableRef     *Variable
}
