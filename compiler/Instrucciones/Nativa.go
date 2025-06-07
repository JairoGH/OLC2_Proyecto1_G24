package instrucciones

import (
	"compiler/tiposDeDato"
	"strconv"
)

type FuncionNativa struct {
	Name string
	Exec func(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string)
}

func (b FuncionNativa) Type() string {
	return tiposDeDato.TIPO_FUNCIONNATIVA
}

func (b FuncionNativa) Value() interface{} {
	return b
}

func (b FuncionNativa) Copy() tiposDeDato.ValorInterno {
	return b
}

// Funcion para Imprimir
func Imprimir(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {

	var output string

	for i, arg := range args {

		if !tiposDeDato.PRIMITIVO(arg.Value.Type()) {
			return tiposDeDato.NuloPorDefecto, false, "La función print solo acepta tipos primitivos"
		}

		switch arg.Value.Type() {

		case tiposDeDato.TIPO_BOOLEAN:
			output += strconv.FormatBool(arg.Value.Value().(bool))
		case tiposDeDato.TIPO_ENTERO:
			output += strconv.Itoa(arg.Value.Value().(int))
		case tiposDeDato.TIPO_DECIMAL:
			output += strconv.FormatFloat(arg.Value.Value().(float64), 'f', 4, 64) // 4 Digitos de Precision
		case tiposDeDato.TIPO_CADENA:
			output += arg.Value.Value().(string)
		case tiposDeDato.TIPO_CARACTER:
			output += arg.Value.Value().(string)
		case tiposDeDato.TIPO_NULO:
			output += "nil"
		}

		// Add a space between each Argumento
		if i < len(args)-1 {
			output += " "
		}
	}

	context.Console.Print(output)

	return tiposDeDato.NuloPorDefecto, true, ""
}

// Funcion para Imprimir con salto de línea
func ImprimirLn(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {

    var output string

    for i, arg := range args {

        if !tiposDeDato.PRIMITIVO(arg.Value.Type()) {
            return tiposDeDato.NuloPorDefecto, false, "La función println solo acepta tipos primitivos"
        }

        switch arg.Value.Type() {

        case tiposDeDato.TIPO_BOOLEAN:
            output += strconv.FormatBool(arg.Value.Value().(bool))
        case tiposDeDato.TIPO_ENTERO:
            output += strconv.Itoa(arg.Value.Value().(int))
        case tiposDeDato.TIPO_DECIMAL:
            output += strconv.FormatFloat(arg.Value.Value().(float64), 'f', 4, 64) // 4 Digitos de Precision
        case tiposDeDato.TIPO_CADENA:
            output += arg.Value.Value().(string)
        case tiposDeDato.TIPO_CARACTER:
            output += arg.Value.Value().(string)
        case tiposDeDato.TIPO_NULO:
            output += "nil"
        }

        // Add a space between each Argumento
        if i < len(args)-1 {
            output += " "
        }
    }

    // Añadir salto de línea al final
    output += "\n"
    
    context.Console.Print(output)

    return tiposDeDato.NuloPorDefecto, true, ""
}

// Funcion Entero

func Entero(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {

	if len(args) != 1 {
		return tiposDeDato.NuloPorDefecto, false, "La función int solo acepta un Argumento"
	}

	argValue := args[0].Value

	if !(argValue.Type() == tiposDeDato.TIPO_CADENA || argValue.Type() == tiposDeDato.TIPO_DECIMAL) {
		return tiposDeDato.NuloPorDefecto, false, "La función Int solo acepta un Argumento de tipo string o float"
	}

	if argValue.Type() == tiposDeDato.TIPO_CADENA {
		ValorDecimal, err := strconv.ParseFloat(argValue.Value().(string), 64)

		if err != nil {
			return tiposDeDato.NuloPorDefecto, false, "No se pudo convertir el valor a int"
		}

		return &tiposDeDato.ValorEntero{
			InternalValor: int(ValorDecimal),
		}, true, ""
	}

	if argValue.Type() == tiposDeDato.TIPO_DECIMAL {
		// truncate the float

		ValorDecimal := argValue.Value().(float64)

		return &tiposDeDato.ValorEntero{
			InternalValor: int(ValorDecimal),
		}, true, ""
	}

	return tiposDeDato.NuloPorDefecto, false, "No se pudo convertir el valor a int"
}

// Funcion Decimal

func Decimal(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {

	if len(args) != 1 {
		return tiposDeDato.NuloPorDefecto, false, "La función float solo acepta un Argumento"
	}

	argValue := args[0].Value

	if !(argValue.Type() == tiposDeDato.TIPO_CADENA) {
		return tiposDeDato.NuloPorDefecto, false, "La función float solo acepta un Argumento de tipo string"
	}

	ValorDecimal, err := strconv.ParseFloat(argValue.Value().(string), 64)

	if err != nil {
		return tiposDeDato.NuloPorDefecto, false, "No se pudo convertir el valor a float"
	}

	return &tiposDeDato.ValorDecimal{
		InternalValor: ValorDecimal,
	}, true, ""
}

// Funcion Cadena

func Cadena(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {

	if len(args) != 1 {
		return tiposDeDato.NuloPorDefecto, false, "La función string solo acepta un Argumento"
	}

	argValue := args[0].Value

	if !(argValue.Type() == tiposDeDato.TIPO_ENTERO || argValue.Type() == tiposDeDato.TIPO_DECIMAL || argValue.Type() == tiposDeDato.TIPO_BOOLEAN) {
		return tiposDeDato.NuloPorDefecto, false, "La función string solo acepta un Argumento de tipo int, float o bool"
	}

	if argValue.Type() == tiposDeDato.TIPO_ENTERO {
		ValorCadena := strconv.Itoa(argValue.Value().(int))

		return &tiposDeDato.ValorCadena{
			InternalValor: ValorCadena,
		}, true, ""
	}

	if argValue.Type() == tiposDeDato.TIPO_DECIMAL {
		ValorCadena := strconv.FormatFloat(argValue.Value().(float64), 'f', 4, 64)

		return &tiposDeDato.ValorCadena{
			InternalValor: ValorCadena,
		}, true, ""
	}

	if argValue.Type() == tiposDeDato.TIPO_BOOLEAN {
		ValorCadena := strconv.FormatBool(argValue.Value().(bool))

		return &tiposDeDato.ValorCadena{
			InternalValor: ValorCadena,
		}, true, ""
	}

	return tiposDeDato.NuloPorDefecto, false, "No se pudo convertir el valor a string"
}

var FuncionesNativasPredeterminadas = map[string]*FuncionNativa{
	"print": {
		Name: "print",
		Exec: Imprimir,
	},
	"println": {
        Name: "println",
        Exec: ImprimirLn,
    },
	"Int": {
		Name: "Int",
		Exec: Entero,
	},
	"Float": {
		Name: "Float",
		Exec: Decimal,
	},
	"String": {
		Name: "String",
		Exec: Cadena,
	},
}
