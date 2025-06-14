package instrucciones

import (
	"main/parser"
	"main/tiposDeDato"

	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

// Funcion: representa una función definida por el usuario con sus parámetros,
// tipo de retorno y contexto de ejecución
type Funcion struct {
	Nombre           string
	Parametros       []*Parametros
	TipoRetorno      string
	TokenTipoRetorno antlr.Token
	Cuerpo           []parser.IStmtContext
	AmbitoDeclaro    *AmbitoBase
	RetornarValor    tiposDeDato.ValorInterno
	IsMutating       bool
	AmbitoDefault    *AmbitoBase
	Token            antlr.Token
}

func (f *Funcion) Value() interface{} {
	return f
}

func (f *Funcion) Type() string {
	return tiposDeDato.TIPO_FUNCION
}

func (f *Funcion) Copy() tiposDeDato.ValorInterno {
	return f
}

// Exec: ejecuta la función validando argumentos, manejando el ámbito
// y controlando el flujo de retorno
func (f *Funcion) Exec(visitor *PatronVIsitor, args []*Argumento, token antlr.Token) {

	context := visitor.GetInstruccionesContexto()

	// Validar argumentos
	argsOk, argsMap := f.ValidarArgumentos(context, args, token)

	if !argsOk {
		f.RetornarValor = tiposDeDato.NuloPorDefecto
		return
	}

	// Crear nuevo ámbito
	initialScope := context.RegistroAmbito.AmbitoActual

	if f.AmbitoDefault != nil {
		context.RegistroAmbito.AmbitoActual = f.AmbitoDefault
	} else {
		context.RegistroAmbito.AmbitoActual = f.AmbitoDeclaro
		context.RegistroAmbito.PushAmbito("func: " + token.GetText())
	}

	wasMutating := context.RegistroAmbito.AmbitoActual.EsMutante
	context.RegistroAmbito.AmbitoActual.EsMutante = f.IsMutating

	// Agregar elemento de retorno a la pila de llamadas
	funcItem := &LlamadaFunciones{
		RetornarValor: tiposDeDato.NuloPorDefecto,
		Tipo: []string{
			Retornar,
		},
	}
	context.PilaLlamada.Push(funcItem)

	// Manejar retorno y limpieza del ámbito

	defer func() {

		context.PilaLlamada.Limpiar(funcItem)
		context.RegistroAmbito.PopAmbito()
		context.RegistroAmbito.AmbitoActual.EsMutante = wasMutating
		context.RegistroAmbito.AmbitoActual = initialScope

		if item, ok := recover().(*LlamadaFunciones); item != nil && ok {

			if item != funcItem {
				context.TablaError.NewErrorSemantico(token, "Return Invalido")
			}

			f.ValidarRetorno(context, item.RetornarValor, token)
			return
		}

		f.ValidarRetorno(context, tiposDeDato.NuloPorDefecto, token)
	}()

	// Agregar argumentos al ámbito
	for varName, arg := range argsMap {

		// Tratamiento especial para paso por referencia
		if arg.esReferencia {

			if arg.VariableRef == nil {
				context.TablaError.NewErrorSemantico(arg.Token, "No es posible pasar por referencia un valor que no este asociado a una variable")
				f.ValidarRetorno(context, tiposDeDato.NuloPorDefecto, token)
				return
			}

			// Crear el puntero
			pointer := &TipoPuntero{
				VariableAsociada: arg.VariableRef,
			}

			context.RegistroAmbito.AmbitoActual.AgregarVariable(varName, tiposDeDato.TIPO_PUNTERO, pointer, false, false, arg.Token)
			continue
		}

		context.RegistroAmbito.AmbitoActual.AgregarVariable(varName, arg.Valor.Type(), arg.Valor.Copy(), false, false, arg.Token)
	}

	// Ejecutar cuerpo de la función
	for _, stmt := range f.Cuerpo {
		visitor.Visit(stmt)
	}

}

// ValidarArgumentos: valida que los argumentos pasados coincidan con los parámetros
// esperados en número, tipo y forma de paso (valor/referencia)
func (f *Funcion) ValidarArgumentos(context *InstruccionesContexto, args []*Argumento, token antlr.Token) (bool, map[string]*Argumento) {

	if len(args) != len(f.Parametros) {
		context.TablaError.NewErrorSemantico(token, "Numero de Argumentos invalido")
		return false, nil
	}

	argsMap := make(map[string]*Argumento)
	finalArgsMap := make(map[string]*Argumento)

	// Verificar si hay argumentos con nombres
	hasNamedArgs := false
	for _, arg := range args {
		if arg.Nombre != "" {
			argsMap[arg.Nombre] = arg
			hasNamedArgs = true
		}
	}

	errorFound := false

	for i, param := range f.Parametros {

		var argToValidate *Argumento = nil

		if !hasNamedArgs {
			argToValidate = args[i]
		} else if param.NombreExterno == "" {
			argToValidate = argsMap[param.NombreInterno]
		} else if param.NombreExterno == "_" {
			argToValidate = args[i]
		} else {
			argToValidate = argsMap[param.NombreExterno]
		}

		if argToValidate == nil {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Argumento %s no especificado", param.NombreInterno))
			errorFound = true
			continue
		}

		if argToValidate.Valor.Type() != param.Tipo && param.Tipo != tiposDeDato.TIPO_ANY {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Tipo de Argumento %s invalido", param.NombreInterno))
			errorFound = true
			continue
		}

		if argToValidate.esReferencia != param.PasarPorReferencia {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Argumento %s no es pasado por referencia", param.NombreInterno))
			errorFound = true
			continue
		}

		finalArgsMap[param.NombreInterno] = argToValidate
	}

	if errorFound {
		return false, nil
	}

	return true, finalArgsMap
}

// ValidarRetorno: verifica que el tipo del valor de retorno coincida
// con el tipo declarado en la función
func (f *Funcion) ValidarRetorno(context *InstruccionesContexto, val tiposDeDato.ValorInterno, token antlr.Token) {

	if val.Type() != f.TipoRetorno {
		if f.TokenTipoRetorno != nil {
			context.TablaError.NewErrorSemantico(f.TokenTipoRetorno, fmt.Sprintf("Tipo de retorno invalido, se esperaba %s, se obtuvo %s", f.TipoRetorno, val.Type()))
		} else {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Tipo de retorno invalido, se esperaba %s, se obtuvo %s", f.TipoRetorno, val.Type()))
		}

		f.RetornarValor = tiposDeDato.NuloPorDefecto
		return
	}

	f.RetornarValor = val
}
