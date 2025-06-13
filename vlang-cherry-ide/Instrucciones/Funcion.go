package instrucciones

import (
	"main/parser"
	"main/tiposDeDato"

	"fmt"

	"github.com/antlr4-go/antlr/v4"
)

type Funcion struct {
	Name            string
	Parametros      []*Parametros
	ReturnType      string
	ReturnTypeToken antlr.Token
	Body            []parser.IStmtContext
	DeclScope       *AmbitoBase
	RetornarValor   tiposDeDato.ValorInterno
	IsMutating      bool
	DefaultScope    *AmbitoBase
	Token           antlr.Token
}

type MetodoStruct struct {
	Name            string
	ReceiverName    string
	ReceiverType    string
	Parametros      []*Parametros
	ReturnType      string
	ReturnTypeToken antlr.Token
	Body            []parser.IStmtContext
	DeclScope       *AmbitoBase
	RetornarValor   tiposDeDato.ValorInterno
	IsMutating      bool
	Token           antlr.Token
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

func (f *Funcion) Exec(visitor *PatronVIsitor, args []*Argumento, token antlr.Token) {

	context := visitor.GetReplContext()

	// validate args
	argsOk, argsMap := f.ValidarArgumentos(context, args, token)

	if !argsOk {
		f.RetornarValor = tiposDeDato.NuloPorDefecto
		return
	}

	// create new scope
	initialScope := context.RegistroAmbito.AmbitoActual // save current scope, scope at call time

	if f.DefaultScope != nil {
		context.RegistroAmbito.AmbitoActual = f.DefaultScope // set Funcion default scope as current scope
	} else {
		context.RegistroAmbito.AmbitoActual = f.DeclScope            // set Funcion declaration scope as current scope
		context.RegistroAmbito.PushScope("func: " + token.GetText()) // push a new Funcion scope
	}

	wasMutating := context.RegistroAmbito.AmbitoActual.IsMutating
	context.RegistroAmbito.AmbitoActual.IsMutating = f.IsMutating

	// push return item to PilaLlamada
	funcItem := &LlamadaFunciones{
		RetornarValor: tiposDeDato.NuloPorDefecto,
		Type: []string{
			Retornar,
		},
	}
	context.PilaLlamada.Push(funcItem)

	// handle return from PilaLlamada

	defer func() {

		context.PilaLlamada.Limpiar(funcItem)                        // clean callstack
		context.RegistroAmbito.PopScope()                            // pop Funcion scope
		context.RegistroAmbito.AmbitoActual.IsMutating = wasMutating // restore mutating flag
		context.RegistroAmbito.AmbitoActual = initialScope           // restore the call time scope

		if item, ok := recover().(*LlamadaFunciones); item != nil && ok {

			if item != funcItem {
				context.TablaError.NewErrorSemantico(token, "Return Invalido")
			}

			// validate return type
			f.ValidateReturn(context, item.RetornarValor, token) // return value from return statement
			return
		}

		f.ValidateReturn(context, tiposDeDato.NuloPorDefecto, token)
	}()

	// push args to scope
	for varName, arg := range argsMap {

		// special treatment for pass by reference
		if arg.PassByReference {

			if arg.VariableRef == nil {
				context.TablaError.NewErrorSemantico(arg.Token, "No es posible pasar por referencia un valor que no este asociado a una variable")
				f.ValidateReturn(context, tiposDeDato.NuloPorDefecto, token)
				return
			}

			// create the pointer
			pointer := &TipoPuntero{
				VariableAsociada: arg.VariableRef,
			}

			// add pointer to scope
			context.RegistroAmbito.AmbitoActual.AgregarVariable(varName, tiposDeDato.TIPO_PUNTERO, pointer, false, false, arg.Token)
			continue
		}

		context.RegistroAmbito.AmbitoActual.AgregarVariable(varName, arg.Value.Type(), arg.Value.Copy(), false, false, arg.Token)
	}

	// evaluate body
	for _, stmt := range f.Body {
		visitor.Visit(stmt)
	}

	// f.ValidateReturn(context, value.NuloPorDefecto, token)
	// return
}

func (f *Funcion) ValidarArgumentos(context *InstruccionesContexto, args []*Argumento, token antlr.Token) (bool, map[string]*Argumento) {

	// validate arg count
	if len(args) != len(f.Parametros) {
		context.TablaError.NewErrorSemantico(token, "Numero de Argumentos invalido")
		return false, nil
	}

	argsMap := make(map[string]*Argumento)
	finalArgsMap := make(map[string]*Argumento)

	// Verificar si hay argumentos con nombres
	hasNamedArgs := false
	for _, arg := range args {
		if arg.Name != "" {
			argsMap[arg.Name] = arg
			hasNamedArgs = true
		}
	}

	errorFound := false

	for i, param := range f.Parametros {

		// determine param type
		var argToValidate *Argumento = nil

		if !hasNamedArgs {
			// Si no hay argumentos con nombres, usar argumentos posicionales
			argToValidate = args[i]
		} else if param.ExternName == "" {
			// inner = arg name
			argToValidate = argsMap[param.InnerName]
		} else if param.ExternName == "_" {
			// positional arg
			argToValidate = args[i]
		} else {
			// extern = arg name
			argToValidate = argsMap[param.ExternName]
		}

		// validate arg exists
		if argToValidate == nil {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Argumento %s no especificado", param.InnerName))
			errorFound = true
			continue
		}

		// validate type
		if argToValidate.Value.Type() != param.Type && param.Type != tiposDeDato.TIPO_ANY {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Tipo de Argumento %s invalido", param.InnerName))
			errorFound = true
			continue
		}

		// validate pass by reference
		if argToValidate.PassByReference != param.PassByReference {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Argumento %s no es pasado por referencia", param.InnerName))
			errorFound = true
			continue
		}

		// add to final args map
		finalArgsMap[param.InnerName] = argToValidate
	}

	if errorFound {
		return false, nil
	}

	return true, finalArgsMap
}

func (f *Funcion) ValidateReturn(context *InstruccionesContexto, val tiposDeDato.ValorInterno, token antlr.Token) {

	if val.Type() != f.ReturnType {
		if f.ReturnTypeToken != nil {
			context.TablaError.NewErrorSemantico(f.ReturnTypeToken, fmt.Sprintf("Tipo de retorno invalido, se esperaba %s, se obtuvo %s", f.ReturnType, val.Type()))
		} else {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Tipo de retorno invalido, se esperaba %s, se obtuvo %s", f.ReturnType, val.Type()))
		}

		f.RetornarValor = tiposDeDato.NuloPorDefecto
		return
	}

	f.RetornarValor = val
}

func (m *MetodoStruct) Value() interface{} {
	return m
}

func (m *MetodoStruct) Type() string {
	return tiposDeDato.TIPO_METODO
}

func (m *MetodoStruct) Copy() tiposDeDato.ValorInterno {
	return m
}

func (m *MetodoStruct) Exec(visitor *PatronVIsitor, receiver *TipoObjeto, args []*Argumento, token antlr.Token) {
	context := visitor.GetReplContext()

	// Validar argumentos
	argsOk, argsMap := m.ValidarArgumentos(context, args, token)
	if !argsOk {
		m.RetornarValor = tiposDeDato.NuloPorDefecto
		return
	}

	// Crear nuevo scope
	initialScope := context.RegistroAmbito.AmbitoActual
	context.RegistroAmbito.AmbitoActual = m.DeclScope
	context.RegistroAmbito.PushScope("method: " + m.Name)

	// Push return item
	funcItem := &LlamadaFunciones{
		RetornarValor: tiposDeDato.NuloPorDefecto,
		Type:          []string{Retornar},
	}
	context.PilaLlamada.Push(funcItem)

	defer func() {
		context.PilaLlamada.Limpiar(funcItem)
		context.RegistroAmbito.PopScope()
		context.RegistroAmbito.AmbitoActual = initialScope

		if item, ok := recover().(*LlamadaFunciones); item != nil && ok {
			if item != funcItem {
				context.TablaError.NewErrorSemantico(token, "Return Invalido")
			}
			m.ValidateReturn(context, item.RetornarValor, token)
			return
		}
		m.ValidateReturn(context, tiposDeDato.NuloPorDefecto, token)
	}()

	// Agregar receiver al scope como la primera variable
	context.RegistroAmbito.AmbitoActual.AgregarVariable(
		m.ReceiverName, receiver.ConcretType, receiver, false, false, token)

	// Agregar argumentos al scope
	for varName, arg := range argsMap {
		context.RegistroAmbito.AmbitoActual.AgregarVariable(
			varName, arg.Value.Type(), arg.Value.Copy(), false, false, arg.Token)
	}

	// Ejecutar cuerpo del método
	for _, stmt := range m.Body {
		visitor.Visit(stmt)
	}
}

func (m *MetodoStruct) ValidarArgumentos(context *InstruccionesContexto, args []*Argumento, token antlr.Token) (bool, map[string]*Argumento) {
	if len(args) != len(m.Parametros) {
		context.TablaError.NewErrorSemantico(token, "Número de Argumento Invalido")
		return false, nil
	}

	argsMap := make(map[string]*Argumento)
	finalArgsMap := make(map[string]*Argumento)

	// Verificar argumentos con nombres
	hasNamedArgs := false
	for _, arg := range args {
		if arg.Name != "" {
			argsMap[arg.Name] = arg
			hasNamedArgs = true
		}
	}

	errorFound := false
	for i, param := range m.Parametros {
		var argToValidate *Argumento = nil

		if !hasNamedArgs {
			argToValidate = args[i]
		} else if param.ExternName == "" {
			argToValidate = argsMap[param.InnerName]
		} else if param.ExternName == "_" {
			argToValidate = args[i]
		} else {
			argToValidate = argsMap[param.ExternName]
		}

		if argToValidate == nil {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Argumento %s no especificado", param.InnerName))
			errorFound = true
			continue
		}

		if argToValidate.Value.Type() != param.Type && param.Type != tiposDeDato.TIPO_ANY {
			context.TablaError.NewErrorSemantico(token, fmt.Sprintf("Tipo de argumento %s inválido", param.InnerName))
			errorFound = true
			continue
		}

		finalArgsMap[param.InnerName] = argToValidate
	}

	if errorFound {
		return false, nil
	}
	return true, finalArgsMap
}

func (m *MetodoStruct) ValidateReturn(context *InstruccionesContexto, val tiposDeDato.ValorInterno, token antlr.Token) {
	if val.Type() != m.ReturnType {
		if m.ReturnTypeToken != nil {
			context.TablaError.NewErrorSemantico(m.ReturnTypeToken,
				fmt.Sprintf("Tipo de retorno inválido, se esperaba %s, se obtuvo %s", m.ReturnType, val.Type()))
		} else {
			context.TablaError.NewErrorSemantico(token,
				fmt.Sprintf("Tipo de retorno inválido, se esperaba %s, se obtuvo %s", m.ReturnType, val.Type()))
		}
		m.RetornarValor = tiposDeDato.NuloPorDefecto
		return
	}
	m.RetornarValor = val
}
