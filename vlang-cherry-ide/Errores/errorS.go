package errores

import (
	"github.com/antlr4-go/antlr/v4"
)

type ErrorPersonalizado struct {
	*antlr.DefaultErrorStrategy
}

func NewErrorPersonalizado() *ErrorPersonalizado {
	return &ErrorPersonalizado{
		DefaultErrorStrategy: antlr.NewDefaultErrorStrategy(),
	}
}

// spanish translation of the error message
func (es *ErrorPersonalizado) ReportarEntradaNoCoincidente(recognizer antlr.Parser, e *antlr.InputMisMatchException) {
	t1 := recognizer.GetTokenStream().LT(-1)
	msg := "Se recibió " + t1.GetText() + ", se esperaba " + es.GetExpectedTokens(recognizer).String()
	recognizer.NotifyErrorListeners(msg, e.GetOffendingToken(), e)
}

