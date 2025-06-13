package instrucciones

import (
	"main/tiposDeDato"
	"strconv"
	"strings"
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

// Funcion para Imprimir - MODIFICADA para aceptar vectores, matrices y structs
func Imprimir(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
	var output string

	for i, arg := range args {
		switch arg.Value.Type() {
		case tiposDeDato.TIPO_BOOLEAN:
			output += strconv.FormatBool(arg.Value.Value().(bool))
		case tiposDeDato.TIPO_ENTERO:
			output += strconv.Itoa(arg.Value.Value().(int))
		case tiposDeDato.TIPO_DECIMAL:
			output += strconv.FormatFloat(arg.Value.Value().(float64), 'f', 4, 64)
		case tiposDeDato.TIPO_CADENA:
			output += arg.Value.Value().(string)
		case tiposDeDato.TIPO_NULO:
			output += "nil"
		default:
			// Verificar si es un struct (TipoObjeto)
			if obj, ok := arg.Value.(*TipoObjeto); ok {
				output += formatStruct(obj)
			} else if EsTipoVector(arg.Value.Type()) || strings.HasPrefix(arg.Value.Type(), "[]") {
				// Manejar vectores y matrices
				if vector, ok := arg.Value.(*TipoVector); ok {
					output += "["
					for j, elemento := range vector.InternalValor {
						if j > 0 {
							output += ", "
						}
						// Recursivamente formatear cada elemento
						switch elemento.Type() {
						case tiposDeDato.TIPO_BOOLEAN:
							output += strconv.FormatBool(elemento.Value().(bool))
						case tiposDeDato.TIPO_ENTERO:
							output += strconv.Itoa(elemento.Value().(int))
						case tiposDeDato.TIPO_DECIMAL:
							output += strconv.FormatFloat(elemento.Value().(float64), 'f', 4, 64)
						case tiposDeDato.TIPO_CADENA:
							output += "\"" + elemento.Value().(string) + "\""
						case tiposDeDato.TIPO_NULO:
							output += "nil"
						default:
							if subVector, ok := elemento.(*TipoVector); ok {
								// Manejar matrices multidimensionales
								output += "["
								for k, subElemento := range subVector.InternalValor {
									if k > 0 {
										output += ", "
									}
									switch subElemento.Type() {
									case tiposDeDato.TIPO_BOOLEAN:
										output += strconv.FormatBool(subElemento.Value().(bool))
									case tiposDeDato.TIPO_ENTERO:
										output += strconv.Itoa(subElemento.Value().(int))
									case tiposDeDato.TIPO_DECIMAL:
										output += strconv.FormatFloat(subElemento.Value().(float64), 'f', 4, 64)
									case tiposDeDato.TIPO_CADENA:
										output += "\"" + subElemento.Value().(string) + "\""
									default:
										output += "nil"
									}
								}
								output += "]"
							} else if obj, ok := elemento.(*TipoObjeto); ok {
								output += formatStruct(obj)
							} else {
								output += "nil"
							}
						}
					}
					output += "]"
				} else {
					return tiposDeDato.NuloPorDefecto, false, "La función print no puede imprimir el tipo: " + arg.Value.Type()
				}
			} else {
				return tiposDeDato.NuloPorDefecto, false, "La función print no puede imprimir el tipo: " + arg.Value.Type()
			}
		}

		// Add a space between each Argumento
		if i < len(args)-1 {
			output += " "
		}
	}

	context.Console.Print(output)
	return tiposDeDato.NuloPorDefecto, true, ""
}

// Funcion para Imprimir con salto de línea - MODIFICADA para aceptar vectores, matrices y structs
func ImprimirLn(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
	var output string

	for i, arg := range args {
		switch arg.Value.Type() {
		case tiposDeDato.TIPO_BOOLEAN:
			output += strconv.FormatBool(arg.Value.Value().(bool))
		case tiposDeDato.TIPO_ENTERO:
			output += strconv.Itoa(arg.Value.Value().(int))
		case tiposDeDato.TIPO_DECIMAL:
			output += strconv.FormatFloat(arg.Value.Value().(float64), 'f', 4, 64)
		case tiposDeDato.TIPO_CADENA:
			output += arg.Value.Value().(string)
		case tiposDeDato.TIPO_NULO:
			output += "nil"
		default:
			// Verificar si es un struct (TipoObjeto)
			if obj, ok := arg.Value.(*TipoObjeto); ok {
				output += formatStruct(obj)
			} else if EsTipoVector(arg.Value.Type()) || strings.HasPrefix(arg.Value.Type(), "[]") {
				// Manejar vectores y matrices
				if vector, ok := arg.Value.(*TipoVector); ok {
					output += "["
					for j, elemento := range vector.InternalValor {
						if j > 0 {
							output += ", "
						}
						// Recursivamente formatear cada elemento
						switch elemento.Type() {
						case tiposDeDato.TIPO_BOOLEAN:
							output += strconv.FormatBool(elemento.Value().(bool))
						case tiposDeDato.TIPO_ENTERO:
							output += strconv.Itoa(elemento.Value().(int))
						case tiposDeDato.TIPO_DECIMAL:
							output += strconv.FormatFloat(elemento.Value().(float64), 'f', 4, 64)
						case tiposDeDato.TIPO_CADENA:
							output += "\"" + elemento.Value().(string) + "\""
						case tiposDeDato.TIPO_NULO:
							output += "nil"
						default:
							if subVector, ok := elemento.(*TipoVector); ok {
								// Manejar matrices multidimensionales
								output += "["
								for k, subElemento := range subVector.InternalValor {
									if k > 0 {
										output += ", "
									}
									switch subElemento.Type() {
									case tiposDeDato.TIPO_BOOLEAN:
										output += strconv.FormatBool(subElemento.Value().(bool))
									case tiposDeDato.TIPO_ENTERO:
										output += strconv.Itoa(subElemento.Value().(int))
									case tiposDeDato.TIPO_DECIMAL:
										output += strconv.FormatFloat(subElemento.Value().(float64), 'f', 4, 64)
									case tiposDeDato.TIPO_CADENA:
										output += "\"" + subElemento.Value().(string) + "\""
									default:
										output += "nil"
									}
								}
								output += "]"
							} else if obj, ok := elemento.(*TipoObjeto); ok {
								output += formatStruct(obj)
							} else {
								output += "nil"
							}
						}
					}
					output += "]"
				} else {
					return tiposDeDato.NuloPorDefecto, false, "La función println no puede imprimir el tipo: " + arg.Value.Type()
				}
			} else {
				return tiposDeDato.NuloPorDefecto, false, "La función println no puede imprimir el tipo: " + arg.Value.Type()
			}
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

// Función auxiliar para formatear structs
func formatStruct(obj *TipoObjeto) string {
	if obj == nil || obj.InternalScope == nil {
		return "nil"
	}

	result := obj.ConcretType + "{"
	first := true

	// Iterar sobre las variables del scope interno del struct
	for name, variable := range obj.InternalScope.variables {
		if !first {
			result += ", "
		}
		first = false

		result += name + ": "

		// Formatear el valor según su tipo
		switch variable.Value.Type() {
		case tiposDeDato.TIPO_CADENA:
			result += "\"" + variable.Value.Value().(string) + "\""
		case tiposDeDato.TIPO_ENTERO:
			result += strconv.Itoa(variable.Value.Value().(int))
		case tiposDeDato.TIPO_DECIMAL:
			result += strconv.FormatFloat(variable.Value.Value().(float64), 'f', 2, 64)
		case tiposDeDato.TIPO_BOOLEAN:
			result += strconv.FormatBool(variable.Value.Value().(bool))
		case tiposDeDato.TIPO_NULO:
			result += "nil"
		default:
			// Si es otro struct anidado
			if nestedObj, ok := variable.Value.(*TipoObjeto); ok {
				result += formatStruct(nestedObj)
			} else {
				result += "nil"
			}
		}
	}

	result += "}"
	return result
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

// Función Len - MODIFICADA para soportar matrices multidimensionales
func Len(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {

	if len(args) != 1 {
		return tiposDeDato.NuloPorDefecto, false, "La función len solo acepta un argumento"
	}

	argValue := args[0].Value

	switch v := argValue.(type) {
	case *TipoVector:
		return &tiposDeDato.ValorEntero{
			InternalValor: v.Size(),
		}, true, ""
	case *tiposDeDato.ValorCadena:
		return &tiposDeDato.ValorEntero{
			InternalValor: len(v.InternalValor),
		}, true, ""
	case *TipoMatriz:
		return &tiposDeDato.ValorEntero{
			InternalValor: v.Size(),
		}, true, ""
	case *TipoMatrizMulti: // NUEVO: Soporte para matrices multidimensionales
		return &tiposDeDato.ValorEntero{
			InternalValor: v.Filas, // Número de filas
		}, true, ""
	default:
		return tiposDeDato.NuloPorDefecto, false, "La función len solo acepta vectores, matrices o cadenas"
	}
}

// Función IndexOf - CORREGIDA para soportar matrices multidimensionales
func IndexOf(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {

	if len(args) != 2 {
		return tiposDeDato.NuloPorDefecto, false, "La función indexOf requiere exactamente 2 argumentos: indexOf(vector/matriz, valor)"
	}

	vectorArg := args[0].Value
	searchValue := args[1].Value

	// NUEVO: Verificar si es una matriz multidimensional
	if matrizMulti, ok := vectorArg.(*TipoMatrizMulti); ok {
		// Buscar en cada fila de la matriz
		for filaIdx, fila := range matrizMulti.InternalValor {
			for colIdx, item := range fila {
				if item.Value() == searchValue.Value() && item.Type() == searchValue.Type() {
					// CORREGIDO: Crear vector manualmente
					items := []tiposDeDato.ValorInterno{
						&tiposDeDato.ValorEntero{InternalValor: filaIdx},
						&tiposDeDato.ValorEntero{InternalValor: colIdx},
					}

					// Crear TipoVector directamente
					resultVector := &TipoVector{
						InternalValor: items,
						FullType:      "[]int",
						ItemType:      "int",
					}
					return resultVector, true, ""
				}
			}
		}
		// Si no se encuentra, retornar [-1, -1]
		items := []tiposDeDato.ValorInterno{
			&tiposDeDato.ValorEntero{InternalValor: -1},
			&tiposDeDato.ValorEntero{InternalValor: -1},
		}

		resultVector := &TipoVector{
			InternalValor: items,
			FullType:      "[]int",
			ItemType:      "int",
		}
		return resultVector, true, ""
	}

	// Verificar que el primer argumento sea un vector simple
	vector, ok := vectorArg.(*TipoVector)
	if !ok {
		return tiposDeDato.NuloPorDefecto, false, "El primer argumento de indexOf debe ser un vector o matriz"
	}

	// Buscar el valor en el vector
	for i, item := range vector.InternalValor {
		// Comparar valores usando el método Value()
		if item.Value() == searchValue.Value() && item.Type() == searchValue.Type() {
			return &tiposDeDato.ValorEntero{
				InternalValor: i,
			}, true, ""
		}
	}

	// Si no se encuentra, retornar -1
	return &tiposDeDato.ValorEntero{
		InternalValor: -1,
	}, true, ""
}

// Función Join - agregar después de la función IndexOf
func Join(ctx *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
	if len(args) != 2 {
		return tiposDeDato.NuloPorDefecto, false,
			"La función join requiere exactamente 2 argumentos: join(vector, separador)"
	}

	vectorArg := args[0].Value
	separatorArg := args[1].Value

	if vectorArg == nil {
		return tiposDeDato.NuloPorDefecto, false, "El primer argumento de join es nil"
	}
	if separatorArg == nil {
		return tiposDeDato.NuloPorDefecto, false, "El segundo argumento de join es nil"
	}

	// Debe ser un vector
	vector, ok := vectorArg.(*TipoVector)
	if !ok {
		return tiposDeDato.NuloPorDefecto, false,
			"El primer argumento de join debe ser un vector (tipo: " + vectorArg.Type() + ")"
	}

	// Sólo vectores de []string
	if vector.FullType != "[]string" {
		return tiposDeDato.NuloPorDefecto, false,
			"La función join solo acepta vectores de tipo []string (tipo recibido: " + vector.FullType + ")"
	}

	// El separador debe ser cadena
	sepCadena, ok := separatorArg.(*tiposDeDato.ValorCadena)
	if !ok {
		return tiposDeDato.NuloPorDefecto, false,
			"El segundo argumento de join debe ser una cadena (tipo recibido: " + separatorArg.Type() + ")"
	}
	separatorStr := sepCadena.InternalValor

	// Construir resultado
	var sb strings.Builder
	for i, item := range vector.InternalValor {
		// Cada elemento es ValorCadena
		strItem := item.(*tiposDeDato.ValorCadena).InternalValor
		sb.WriteString(strItem)
		if i < len(vector.InternalValor)-1 {
			sb.WriteString(separatorStr)
		}
	}

	return &tiposDeDato.ValorCadena{InternalValor: sb.String()}, true, ""
}

// Función Append - MODIFICADA para soportar matrices multidimensionales
func Append(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {

	if len(args) != 2 {
		return tiposDeDato.NuloPorDefecto, false, "La función append requiere exactamente 2 argumentos: append(vector/matriz, elemento)"
	}

	vectorArg := args[0].Value
	elementArg := args[1].Value

	// NUEVO: Verificar si es una matriz multidimensional
	if matrizMulti, ok := vectorArg.(*TipoMatrizMulti); ok {
		// Debe ser un vector (fila) compatible con la matriz
		if elementVector, ok := elementArg.(*TipoVector); ok {
			// Verificar que el tipo del vector sea compatible con el tipo de fila de la matriz
			if elementVector.FullType != matrizMulti.ItemType {
				return tiposDeDato.NuloPorDefecto, false, "No se puede agregar un vector de tipo " + elementVector.FullType + " a una matriz de tipo " + matrizMulti.FullType
			}

			// Crear una copia de la matriz original
			newMatrix := make([][]tiposDeDato.ValorInterno, len(matrizMulti.InternalValor))
			for i, row := range matrizMulti.InternalValor {
				newMatrix[i] = make([]tiposDeDato.ValorInterno, len(row))
				copy(newMatrix[i], row)
			}

			// Agregar la nueva fila
			newMatrix = append(newMatrix, elementVector.InternalValor)

			// Crear y retornar una nueva matriz con la fila agregada
			return NewTipoMatrizMulti(newMatrix, matrizMulti.FullType, matrizMulti.ItemType, matrizMulti.BaseType), true, ""
		} else {
			return tiposDeDato.NuloPorDefecto, false, "Solo se pueden agregar vectores (filas) a una matriz multidimensional"
		}
	}

	// Verificar que el primer argumento sea un vector simple
	vector, ok := vectorArg.(*TipoVector)
	if !ok {
		return tiposDeDato.NuloPorDefecto, false, "El primer argumento de append debe ser un vector o matriz"
	}

	// Verificar que el tipo del elemento coincida con el tipo del vector
	expectedElementType := vector.ItemType
	actualElementType := elementArg.Type()

	if expectedElementType != actualElementType {
		return tiposDeDato.NuloPorDefecto, false, "No se puede agregar un elemento de tipo " + actualElementType + " a un vector de tipo " + vector.FullType
	}

	// Crear una copia del vector original (no modificar el original)
	newItems := make([]tiposDeDato.ValorInterno, len(vector.InternalValor))
	copy(newItems, vector.InternalValor)

	// Agregar el nuevo elemento
	newItems = append(newItems, elementArg)

	// Crear y retornar un nuevo vector con los elementos actualizados
	newVector := NewTipoVector(newItems, vector.FullType, vector.ItemType)

	return newVector, true, ""
}

func Atoi(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
	if len(args) != 1 {
		return tiposDeDato.NuloPorDefecto, false, "La función 'atoi' solo acepta un argumento"
	}

	arg := args[0].Value

	if arg.Type() != tiposDeDato.TIPO_CADENA {
		return tiposDeDato.NuloPorDefecto, false, "La función 'atoi' solo acepta argumentos de tipo string"
	}

	str := arg.Value().(string)

	// Verificar si contiene punto (decimal)
	if _, err := strconv.ParseFloat(str, 64); err == nil && containsDot(str) {
		return tiposDeDato.NuloPorDefecto, false, "La función 'atoi' no permite valores decimales"
	}

	valor, err := strconv.Atoi(str)
	if err != nil {
		return tiposDeDato.NuloPorDefecto, false, "No se pudo convertir el string a entero"
	}

	return &tiposDeDato.ValorEntero{
		InternalValor: valor,
	}, true, ""
}

// Función auxiliar para detectar decimales
func containsDot(input string) bool {
	for _, ch := range input {
		if ch == '.' {
			return true
		}
	}
	return false
}

func ParseFloat(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
	if len(args) != 1 {
		return tiposDeDato.NuloPorDefecto, false, "La función 'parseFloat' solo acepta un argumento"
	}

	arg := args[0].Value

	if arg.Type() != tiposDeDato.TIPO_CADENA {
		return tiposDeDato.NuloPorDefecto, false, "La función 'parseFloat' solo acepta argumentos de tipo string"
	}

	str := arg.Value().(string)

	valor, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return tiposDeDato.NuloPorDefecto, false, "No se pudo convertir el string a decimal"
	}

	return &tiposDeDato.ValorDecimal{
		InternalValor: valor,
	}, true, ""
}

func TypeOf(context *InstruccionesContexto, args []*Argumento) (tiposDeDato.ValorInterno, bool, string) {
	// Validación más robusta
	if len(args) != 1 {
		return tiposDeDato.NuloPorDefecto, false, "La función typeOf espera exactamente un argumento"
	}

	if args[0] == nil {
		return tiposDeDato.NuloPorDefecto, false, "El argumento de typeOf no puede ser nil"
	}

	if args[0].Value == nil {
		return tiposDeDato.NuloPorDefecto, false, "El valor del argumento de typeOf no puede ser nil"
	}

	valor := args[0].Value
	var tipo string

	// DEBUG: Información adicional para diagnosticar
	//fmt.Printf("DEBUG TypeOf: valor tipo: %T, valor: %+v\n", valor, valor)

	// Verificación adicional de seguridad
	if valor == nil {
		return &tiposDeDato.ValorCadena{
			InternalValor: "nil",
		}, true, ""
	}

	// Manejo seguro de tipos
	switch v := valor.(type) {
	case *TipoVector:
		if v == nil {
			return tiposDeDato.NuloPorDefecto, false, "Vector es nil en typeOf"
		}

		// DEBUG: Información del vector
		//fmt.Printf("DEBUG Vector: ItemType='%s', FullType='%s', Size=%d\n", v.ItemType, v.FullType, len(v.InternalValor))

		// Validar que el vector tenga información de tipo
		if v.ItemType != "" {
			// Mapear tipos específicos para el test
			baseType := v.ItemType
			if baseType == "float64" || baseType == tiposDeDato.TIPO_DECIMAL {
				baseType = "f64"
			} else if baseType == tiposDeDato.TIPO_ENTERO || baseType == "int" {
				baseType = "int"
			} else if baseType == tiposDeDato.TIPO_CADENA || baseType == "string" {
				baseType = "string"
			} else if baseType == tiposDeDato.TIPO_BOOLEAN || baseType == "bool" {
				baseType = "bool"
			}
			tipo = "[]" + baseType
		} else if v.FullType != "" {
			// Si ItemType está vacío, usar FullType
			tipo = v.FullType
		} else if len(v.InternalValor) > 0 {
			// Si no hay información de tipo, inferir del primer elemento
			primerElemento := v.InternalValor[0]
			if primerElemento != nil {
				switch primerElemento.Type() {
				case tiposDeDato.TIPO_ENTERO:
					tipo = "[]int"
				case tiposDeDato.TIPO_DECIMAL:
					tipo = "[]f64"
				case tiposDeDato.TIPO_CADENA:
					tipo = "[]string"
				case tiposDeDato.TIPO_BOOLEAN:
					tipo = "[]bool"
				default:
					tipo = "[]any"
				}
			} else {
				tipo = "[]any"
			}
		} else {
			tipo = "[]any"
		}

	case *TipoMatriz:
		if v == nil {
			tipo = "nil"
		} else if v.ItemType != "" {
			baseType := v.ItemType
			if baseType == "float64" || baseType == tiposDeDato.TIPO_DECIMAL {
				baseType = "f64"
			} else if baseType == tiposDeDato.TIPO_ENTERO {
				baseType = "int"
			}
			tipo = "[][]" + baseType
		} else {
			tipo = "[][]any"
		}

	case *TipoMatrizMulti:
		if v == nil {
			tipo = "nil"
		} else if v.BaseType != "" {
			baseType := v.BaseType
			if baseType == "float64" || baseType == tiposDeDato.TIPO_DECIMAL {
				baseType = "f64"
			} else if baseType == tiposDeDato.TIPO_ENTERO {
				baseType = "int"
			}
			tipo = "[][]" + baseType
		} else {
			tipo = "[][]any"
		}

	default:
		// Para tipos primitivos, usar método Type() de forma segura
		tipoOriginal := ""
		if valor != nil {
			tipoOriginal = valor.Type()
		}

		switch tipoOriginal {
		case tiposDeDato.TIPO_ENTERO:
			tipo = "int"
		case tiposDeDato.TIPO_DECIMAL:
			tipo = "f64"
		case tiposDeDato.TIPO_CADENA:
			tipo = "string"
		case tiposDeDato.TIPO_BOOLEAN:
			tipo = "bool"
		case tiposDeDato.TIPO_NULO:
			tipo = "nil"
		default:
			if tipoOriginal != "" {
				tipo = tipoOriginal
			} else {
				tipo = "unknown"
			}
		}
	}

	// Verificar que el resultado no sea vacío
	if tipo == "" {
		tipo = "unknown"
	}

	return &tiposDeDato.ValorCadena{
		InternalValor: tipo,
	}, true, ""
}

var DefaultBuiltInFunctions = map[string]*FuncionNativa{
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
	"atoi": {
		Name: "atoi",
		Exec: Atoi,
	},
	"parseFloat": {
		Name: "parseFloat",
		Exec: ParseFloat,
	},
	"typeOf": {
		Name: "typeOf",
		Exec: TypeOf,
	},
	"len": {
		Name: "len",
		Exec: Len,
	},
	"indexOf": {
		Name: "indexOf",
		Exec: IndexOf,
	},
	"join": {
		Name: "join",
		Exec: Join,
	},
	"append": {
		Name: "append",
		Exec: Append,
	},
}
