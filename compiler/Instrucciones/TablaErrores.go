package instrucciones

import "github.com/antlr4-go/antlr/v4"

const (
	ErrorLexico     = "Error léxico"
	ErrorSintactico = "Error sintáctico"
	ErrorSemantico  = "Error semántico"
	ErrorEjecucion  = "Error en tiempo de ejecución"
)

type Error struct {
	Linea       int
	Columna     int
	Descripcion string
	Tipo        string
}

type TablaError struct {
	Errors []Error
}

func (et *TablaError) AddError(linea int, colunma int, descr string, tipo string) {
	et.Errors = append(et.Errors, Error{linea, colunma, descr, tipo})
}

func (et *TablaError) NewLErrorLexico(linea int, colunma int, descr string) {
	et.AddError(linea, colunma, descr, ErrorLexico)
}

func (et *TablaError) NewErrorSintactico(linea int, colunma int, descr string) {
	et.AddError(linea, colunma, descr, ErrorSintactico)
}

func (et *TablaError) NewErrorSemantico(token antlr.Token, descr string) {
	et.AddError(token.GetLine(), token.GetColumn(), descr, ErrorSemantico)
}

func (et *TablaError) NewErrorEjecucion(linea int, colunma int, descr string) {
	et.AddError(linea, colunma, descr, ErrorEjecucion)
}

func (et *TablaError) GetErrors() []Error {
	return et.Errors
}

func NewErrorTable() *TablaError {
	return &TablaError{}
}
