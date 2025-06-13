package errores

import (
	"main/Instrucciones"

	"github.com/antlr4-go/antlr/v4"
)

type SyntaxError struct {
	*antlr.DefaultErrorListener
	TablaError *instrucciones.TablaError
}

func NewSyntaxError(TablaError *instrucciones.TablaError) *SyntaxError {
	return &SyntaxError{
		TablaError: TablaError,
	}
}

func (e *SyntaxError) SyntaxError(_ antlr.Recognizer, _ interface{}, linea, columna int, descripcion string, _ antlr.RecognitionException) {

	e.TablaError.AddError(
		linea,
		columna,
		descripcion,
		instrucciones.ErrorSintactico,
	)

}

type LexicalErrorListener struct {
	*antlr.DefaultErrorListener
	TablaError *instrucciones.TablaError
}

func NewErrorLexico() *LexicalErrorListener {
	return &LexicalErrorListener{
		TablaError: instrucciones.NewTablaError(),
	}
}

func (e *LexicalErrorListener) SyntaxError(_ antlr.Recognizer, _ interface{}, linea, columna int, descripcion string, _ antlr.RecognitionException) {

	e.TablaError.AddError(
		linea,
		columna,
		descripcion,
		instrucciones.ErrorLexico,
	)

}
