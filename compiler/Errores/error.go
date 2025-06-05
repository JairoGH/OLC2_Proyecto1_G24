package Errores

import (
	error "compiler/Instrucciones"

	"github.com/antlr4-go/antlr/v4"
)

type SyntaxError struct {
	*antlr.DefaultErrorListener
	TablaErrores *error.TablaError
}

func NewSyntaxError(tablaErrores *error.TablaError) *SyntaxError {
	return &SyntaxError{
		TablaErrores: tablaErrores,
	}
}


func (e *SyntaxError) SyntaxError(_ antlr.Recognizer, _ interface{}, linea, columna int, msg string, _ antlr.RecognitionException) {

	e.TablaErrores.AddError(
		linea,
		columna,
		msg,
		error.ErrorSintactico,
	)

}

type ErrorLexico struct {
	*antlr.DefaultErrorListener
	TablaErrores *error.TablaError
}

func NewErrorLexico() *ErrorLexico {
	return &ErrorLexico{
		TablaErrores: error.NewErrorTable(),
	}
}

func (e *ErrorLexico) SyntaxError(_ antlr.Recognizer, _ interface{}, linea, columna int, msg string, _ antlr.RecognitionException) {

	e.TablaErrores.AddError(
		linea,
		columna,
		msg,
		error.ErrorLexico,
	)

}
