package instrucciones

import "github.com/antlr4-go/antlr/v4"

const (
	ErrorLexico     = "Error léxico"
	ErrorSintactico = "Error sintáctico"
	ErrorSemantico  = "Error semántico"
	ErrorEjecucion  = "Error en tiempo de ejecución"
)

type Error struct {
	Line   int
	Columna int
	Descripcion    string
	Tipo   string
}



type TablaError struct {
	Errores []Error
}

func (et *TablaError) AddError(Linea int, Columna int, descripcion string, TipoError string) {
	et.Errores = append(et.Errores, Error{Linea, Columna, descripcion, TipoError})
}

func (et *TablaError) NewLErrorLexico(Linea int, Columna int, descripcion string) {
	et.AddError(Linea, Columna, descripcion, ErrorLexico)
}

func (et *TablaError) NewErrorSintactico(Linea int, Columna int, descripcion string) {
	et.AddError(Linea, Columna, descripcion, ErrorSintactico)
}

func (et *TablaError) NewErrorSemantico(token antlr.Token, descripcion string) {
	et.AddError(token.GetLine(), token.GetColumn(), descripcion, ErrorSemantico)
}

func (et *TablaError) NewErrorEjecucion(Linea int, Columna int, descripcion string) {
	et.AddError(Linea, Columna, descripcion, ErrorEjecucion)
}

func (et *TablaError) GetErrors() []Error {
	return et.Errores
}
func NewTablaError() *TablaError {
	return &TablaError{}
}
