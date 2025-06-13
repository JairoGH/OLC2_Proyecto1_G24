package instrucciones

import (
	"fmt"
	"log"
	"main/parser"
	"main/tiposDeDato"
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type PatronVIsitor struct {
	parser.BaseVGrammarVisitor
	RegistroAmbito *RegistroAmbito
	PilaLlamada    *PilaLlamada
	Console        *Console
	TablaError     *TablaError
	StructNames    []string
}

// defaultValueForType devuelve el ValorInterno por defecto para tipos primitivos
func defaultValueForType(typeName string) tiposDeDato.ValorInterno {
	switch typeName {
	case tiposDeDato.TIPO_ENTERO:
		return &tiposDeDato.ValorEntero{InternalValor: 0}
	case tiposDeDato.TIPO_DECIMAL:
		return &tiposDeDato.ValorDecimal{InternalValor: 0.0}
	case tiposDeDato.TIPO_CADENA:
		return &tiposDeDato.ValorCadena{InternalValor: ""}
	case tiposDeDato.TIPO_BOOLEAN:
		return &tiposDeDato.ValorBool{InternalValor: false}
	default:
		return tiposDeDato.NuloPorDefecto
	}
}

func NewVisitor(VisitanteDcl *VisitanteDcl) *PatronVIsitor {
	return &PatronVIsitor{
		RegistroAmbito: VisitanteDcl.RegistroAmbito,
		TablaError:     VisitanteDcl.TablaError,
		StructNames:    VisitanteDcl.StructNames,
		PilaLlamada:    NuevaLlamadaFuncion(),
		Console:        NewConsole(),
	}
}

func (v *PatronVIsitor) GetReplContext() *InstruccionesContexto {
	return &InstruccionesContexto{
		Console:        v.Console,
		RegistroAmbito: v.RegistroAmbito,
		PilaLlamada:    v.PilaLlamada,
		TablaError:     v.TablaError,
	}
}

func (v *PatronVIsitor) ValidType(_type string) bool {
	return v.RegistroAmbito.AmbitoGlobal.ValidType(_type)
}

func (v *PatronVIsitor) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}

func (v *PatronVIsitor) VisitProgram(ctx *parser.ProgramContext) interface{} {

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	// Agregar el manejo de main_func
	if ctx.Main_func() != nil {
		v.Visit(ctx.Main_func())
	}

	return nil
}
func (v *PatronVIsitor) VisitFuncionMain(ctx *parser.FuncionMainContext) interface{} {

	// Push scope para main
	v.RegistroAmbito.PushScope("main")

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	// Pop scope
	v.RegistroAmbito.PopScope()

	return nil
}

func (v *PatronVIsitor) VisitStmt(ctx *parser.StmtContext) interface{} {
	// 1) Declaraciones (decl_stmt) con sus 7 alternativas
	if ds := ctx.Decl_stmt(); ds != nil {
		switch decl := ds.(type) {
		case *parser.DeclararInferenciaMutContext:
			v.VisitDeclararInferenciaMut(decl) // mut x := expr
		case *parser.DeclararInferenciaContext:
			v.VisitDeclararInferencia(decl) // x := expr
		case *parser.DeclaraTipoValorContext:
			v.VisitDeclaraTipoValor(decl) // mut x type = expr
		case *parser.DeclararTipoContext:
			v.VisitDeclararTipo(decl) // mut x type
		case *parser.DeclararSinMutValorContext:
			v.VisitDeclararSinMutValor(decl) // x type = expr
		case *parser.DeclararSinMutTipoContext:
			v.VisitDeclararSinMutTipo(decl) // x type
		case *parser.DeclararVectorContext:
			v.VisitDeclararVector(decl) // mut x = [type] [values...]
		default:
			v.Visit(ds)
		}
		return nil
	}

	// 2) Asignaciones +, -, *, /, % y accesos a vectores
	if ctx.Assign_stmt() != nil {
		v.Visit(ctx.Assign_stmt())
		return nil
	}

	// 3) Resto de sentencias
	switch {
	case ctx.If_stmt() != nil:
		v.Visit(ctx.If_stmt())
	case ctx.Switch_stmt() != nil:
		v.Visit(ctx.Switch_stmt())
	case ctx.While_stmt() != nil:
		v.Visit(ctx.While_stmt())
	case ctx.For_stmt() != nil:
		v.Visit(ctx.For_stmt())
	case ctx.Guard_stmt() != nil:
		v.Visit(ctx.Guard_stmt())
	case ctx.Transfer_stmt() != nil:
		v.Visit(ctx.Transfer_stmt())
	case ctx.Func_call() != nil:
		v.Visit(ctx.Func_call())
	case ctx.Func_dcl() != nil:
		v.Visit(ctx.Func_dcl())
	case ctx.Strct_dcl() != nil:
		v.Visit(ctx.Strct_dcl())
	case ctx.Vector_func() != nil:
		v.Visit(ctx.Vector_func())
	default:
		log.Fatal("Declaracion No Fue Encontrada: " + ctx.GetText())
	}

	return nil
}

// 1) mut ID type = expr
func (v *PatronVIsitor) VisitDeclaraTipoValor(ctx *parser.DeclaraTipoValorContext) interface{} {
	varName := ctx.ID().GetText()
	varType := v.Visit(ctx.Type_()).(string)
	varValue := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	// 1) Verificar mismatch de tipo
	if varValue.Type() != varType {
		v.TablaError.NewErrorSemantico(ctx.GetStart(),
			fmt.Sprintf("No se puede asignar un valor de tipo %s a variable %s de tipo %s",
				varValue.Type(), varName, varType))
		return nil
	}

	// copy object/vector
	if obj, ok := varValue.(*TipoObjeto); ok {
		varValue = obj.Copy()
	}
	if EsTipoVector(varValue.Type()) {
		varValue = varValue.Copy()
	}

	// insertar mutable + inicializada
	variable, msg := v.RegistroAmbito.AgregarVariable(
		varName, varType, varValue,
		false, /* mutable */
		false, /* inicializada */
		ctx.GetStart(),
	)
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}
	return nil
}

// vector x type = expr
func (v *PatronVIsitor) VisitDeclararVector(ctx *parser.DeclararVectorContext) interface{} {
	varName := ctx.ID().GetText()

	// Verificar que ctx.Type_() no sea nil
	if ctx.Type_() == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Tipo no especificado en declaración de vector")
		return nil
	}

	typeResult := v.Visit(ctx.Type_())
	if typeResult == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Error al procesar el tipo del vector")
		return nil
	}
	varType := typeResult.(string)

	// Verificar que ctx.Expr() no sea nil
	if ctx.Expr() == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Valor no especificado en declaración de vector")
		return nil
	}

	valueResult := v.Visit(ctx.Expr())
	if valueResult == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Error al procesar el valor del vector")
		return nil
	}
	varValue := valueResult.(tiposDeDato.ValorInterno)

	// NUEVO: Manejar vectores vacíos especialmente
	if varValue.Type() == "[]" {
		// Es un vector vacío, asignarle el tipo correcto
		if vectorVal, ok := varValue.(*TipoVector); ok {
			vectorVal.FullType = varType
			vectorVal.ItemType = EliminarCorchetes(varType)
		}
	} else {
		// Verificar que el tipo del vector coincida
		if varValue.Type() != varType {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("No se puede asignar un vector de tipo %s a variable %s de tipo %s",
					varValue.Type(), varName, varType))
			return nil
		}
	}

	// copy object
	if obj, ok := varValue.(*TipoObjeto); ok {
		varValue = obj.Copy()
	}

	if EsTipoVector(varValue.Type()) {
		varValue = varValue.Copy()
	}

	// CAMBIO: true para mutable y TRUE para inicializada
	variable, msg := v.RegistroAmbito.AgregarVariable(varName, varType, varValue, false, true, ctx.GetStart())

	// Variable already exists
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}

	return nil
}

func isDeclConst(lexval string) bool {
	return lexval == "let"
}

// mut x := expr  (inferencia mutable)
func (v *PatronVIsitor) VisitDeclararInferenciaMut(ctx *parser.DeclararInferenciaMutContext) interface{} {
	name := ctx.ID().GetText()
	val := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	// copia objetos/vectores
	if obj, ok := val.(*TipoObjeto); ok {
		val = obj.Copy()
	}
	if EsTipoVector(val.Type()) {
		val = val.Copy()
	}

	typ := val.Type()
	// insertar mutable + inicializada
	variable, msg := v.RegistroAmbito.AgregarVariable(
		name, typ, val,
		false, /* mutable */
		false, /* inicializada */
		ctx.GetStart(),
	)
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}
	return nil
}

// x := expr  (inferencia sin mut)
func (v *PatronVIsitor) VisitDeclararInferencia(ctx *parser.DeclararInferenciaContext) interface{} {
	name := ctx.ID().GetText()
	val := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	if obj, ok := val.(*TipoObjeto); ok {
		val = obj.Copy()
	}
	if EsTipoVector(val.Type()) {
		val = val.Copy()
	}

	typ := val.Type()
	variable, msg := v.RegistroAmbito.AgregarVariable(
		name, typ, val,
		false, /* mutable = false */
		false, /* inicializada */
		ctx.GetStart(),
	)
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}
	return nil
}

func (v *PatronVIsitor) VisitDeclararTipo(ctx *parser.DeclararTipoContext) interface{} {
	varName := ctx.ID().GetText()
	varType := v.Visit(ctx.Type_()).(string)

	// valor por defecto según tipo
	defaultVal := defaultValueForType(varType)

	variable, msg := v.RegistroAmbito.AgregarVariable(
		varName, varType, defaultVal,
		false, /* mutable */
		false, /* inicializada */
		ctx.GetStart(),
	)
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}
	return nil
}

func (v *PatronVIsitor) VisitListaItemsVector(ctx *parser.ListaItemsVectorContext) interface{} {

	var vectorItems []tiposDeDato.ValorInterno

	if len(ctx.AllExpr()) == 0 {
		return NewTipoVector(vectorItems, "[]", "")
	}

	for _, item := range ctx.AllExpr() {
		vectorItems = append(vectorItems, v.Visit(item).(tiposDeDato.ValorInterno))
	}

	if len(vectorItems) == 0 {
		return NewTipoVector(vectorItems, "[]", "")
	}

	itemType := vectorItems[0].Type()

	// NUEVO: Detectar si es una matriz (vector de vectores)
	isMatrix := strings.HasPrefix(itemType, "[]")

	for _, item := range vectorItems {
		if item.Type() != itemType {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Todos los items de la colección deben ser del mismo tipo")
			return tiposDeDato.NuloPorDefecto
		}
	}

	if isMatrix {
		// Es una matriz [][]tipo
		baseType := strings.TrimPrefix(itemType, "[]") // extraer tipo base
		matrixType := "[]" + itemType                  // [][]int

		// Convertir vectores a matriz multidimensional
		matrix := make([][]tiposDeDato.ValorInterno, len(vectorItems))
		for i, item := range vectorItems {
			if vector, ok := item.(*TipoVector); ok {
				matrix[i] = vector.InternalValor
			} else {
				v.TablaError.NewErrorSemantico(ctx.GetStart(),
					"Error interno: elemento de matriz no es vector")
				return tiposDeDato.NuloPorDefecto
			}
		}

		return NewTipoMatrizMulti(matrix, matrixType, itemType, baseType)
	} else {
		// Es un vector simple []tipo
		_type := "[]" + itemType
		return NewTipoVector(vectorItems, _type, itemType)
	}
}

func (v *PatronVIsitor) VisitType(ctx *parser.TypeContext) interface{} {
	_type := ctx.GetText()

	// Verificar si es matriz multidimensional [][]tipo
	if strings.HasPrefix(_type, "[][]") && len(_type) > 4 {
		// Extraer el tipo base (después de [][])
		baseType := _type[4:] // eliminar [][]

		// Verificar si el tipo base es válido
		if v.ValidType(baseType) {
			return _type // retornar [][]int, [][]string, etc.
		}

		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El tipo "+baseType+" no es válido para una matriz")
		return tiposDeDato.TIPO_NULO
	}

	// Verificar si es vector simple []tipo
	if strings.HasPrefix(_type, "[]") && len(_type) > 2 {
		internType := _type[2:] // eliminar []

		if v.ValidType(internType) {
			return _type // retornar []int, []string, etc.
		}

		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El tipo "+internType+" no es válido para un vector")
		return tiposDeDato.TIPO_NULO
	}

	// Resto del código existente para tipos primitivos y formatos antiguos...
	if v.ValidType(_type) {
		return _type
	}

	// Código existente para EsTipoVector y EsTipoMatriz...

	v.TablaError.NewErrorSemantico(ctx.GetStart(), "Tipo "+ctx.GetText()+" no encontrado")
	return tiposDeDato.TIPO_NULO
}

// Función auxiliar para verificar si un tipo es vector (ambos formatos)
func esFormatoVector(tipo string) bool {
	// Formato []tipo
	if strings.HasPrefix(tipo, "[]") && len(tipo) > 2 {
		return true
	}
	// Formato [tipo]
	return EsTipoVector(tipo)
}

func (v *PatronVIsitor) VisitVector_type(ctx *parser.Vector_typeContext) interface{} {
	return ctx.GetText()
}

func (v *PatronVIsitor) VisitRepeating(ctx *parser.RepeatingContext) interface{} {

	if ctx.ID(0).GetText() != "repeating" {
		v.TablaError.NewErrorSemantico(ctx.ID(0).GetSymbol(), "La sintaxis de la función espera el parametro 'repeating'")
		return tiposDeDato.NuloPorDefecto
	}

	if ctx.ID(1).GetText() != "count" {
		v.TablaError.NewErrorSemantico(ctx.ID(1).GetSymbol(), "La sintaxis de la función espera el parametro 'count'")
		return tiposDeDato.NuloPorDefecto
	}

	reapeating_val := v.Visit(ctx.Expr(0)).(tiposDeDato.ValorInterno)
	count_val := v.Visit(ctx.Expr(1)).(tiposDeDato.ValorInterno)

	count, validCount := count_val.(*tiposDeDato.ValorEntero)
	if !validCount {
		v.TablaError.NewErrorSemantico(ctx.Expr(1).GetStart(), "El parametro count debe ser un entero")
		return tiposDeDato.NuloPorDefecto
	}

	if ctx.Vector_type() != nil {
		vector_type := ctx.Vector_type().GetText()
		primitive_type := EliminarCorchetes(vector_type)

		if primitive_type != reapeating_val.Type() {
			v.TablaError.NewErrorSemantico(ctx.Expr(0).GetStart(), "El tipo del valor repetido debe ser "+primitive_type)
			return tiposDeDato.NuloPorDefecto
		}

		var vectorItems []tiposDeDato.ValorInterno

		for i := 0; i < count.InternalValor; i++ {
			vectorItems = append(vectorItems, reapeating_val.Copy()) // ? indepedent values
		}

		return NewTipoVector(vectorItems, vector_type, primitive_type)

	} else if ctx.Matrix_type() != nil {

		matrix_type := ctx.Matrix_type().GetText()

		if !(EsTipoMatriz(reapeating_val.Type()) || EsTipoVector(reapeating_val.Type())) {
			v.TablaError.NewErrorSemantico(ctx.Expr(0).GetStart(), "Para crear una matriz con valores repetidos, el valor debe ser un vector o una matriz, se obtuvo '"+reapeating_val.Type()+"'")
			return tiposDeDato.NuloPorDefecto
		}

		if !tiposDeDato.PRIMITIVO(EliminarCorchetes(matrix_type)) {
			v.TablaError.NewErrorSemantico(ctx.Expr(0).GetStart(), "Las matrices solo pueden contener tipos primitivos")
			return tiposDeDato.NuloPorDefecto
		}

		// must be a lower order collection
		if matrix_type != "["+reapeating_val.Type()+"]" {
			v.TablaError.NewErrorSemantico(ctx.Expr(0).GetStart(), "Para conseguir un valor de tipo '"+matrix_type+"' no es posible contruirlo con un valores repetidos del tipo '"+reapeating_val.Type()+"'")
			return tiposDeDato.NuloPorDefecto
		}

		var matrixItems []tiposDeDato.ValorInterno

		for i := 0; i < count.InternalValor; i++ {
			matrixItems = append(matrixItems, reapeating_val.Copy()) // ? indepedent values
		}

		return NewTipoMatriz(matrixItems, matrix_type, EliminarCorchetes(reapeating_val.Type()))
	}
	return tiposDeDato.NuloPorDefecto
}

func (v *PatronVIsitor) VisitRepeatingExp(ctx *parser.RepeatingExpContext) interface{} {
	return v.Visit(ctx.Repeating())
}

func (v *PatronVIsitor) VisitVectorItem(ctx *parser.VectorItemContext) interface{} {

	varName := ctx.Id_pattern().GetText()
	variable := v.RegistroAmbito.GetVariable(varName)

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variable "+varName+" no encontrada")
		return nil
	}

	// Verificar que sea un vector o matriz (incluyendo formato [][]tipo)
	if !(EsTipoVector(variable.Type) || EsTipoMatriz(variable.Type) ||
		strings.HasPrefix(variable.Type, "[]")) {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La variable "+varName+" no es un vector o una matriz")
		return nil
	}

	indexes := []int{}
	for _, expr := range ctx.AllExpr() {
		val := v.Visit(expr).(tiposDeDato.ValorInterno)
		if val.Type() != tiposDeDato.TIPO_ENTERO {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Los indices de acceso deben ser enteros")
			return nil
		}
		indexes = append(indexes, val.Value().(int))
	}

	// NUEVO: Manejar TipoMatrizMulti primero
	if matrizMulti, ok := variable.Value.(*TipoMatrizMulti); ok {
		if len(indexes) == 2 {
			// Acceso completo: matriz[fila][columna]
			fila, columna := indexes[0], indexes[1]
			if !matrizMulti.ValidIndex(fila, columna) {
				v.TablaError.NewErrorSemantico(ctx.GetStart(),
					fmt.Sprintf("Índice fuera de rango: [%d][%d]. La matriz tiene %d filas",
						fila, columna, len(matrizMulti.InternalValor)))

				// Retornar un valor que indique error pero que no rompa el programa
				return &tiposDeDato.ValorEntero{InternalValor: -1}
			}
			return &MatrizMultiItemReference{
				Matriz:  matrizMulti,
				Fila:    fila,
				Columna: columna,
				Value:   matrizMulti.Get(fila, columna),
			}
		} else if len(indexes) == 1 {
			// Acceso a fila: matriz[fila] -> retorna vector
			fila := indexes[0]
			if fila < 0 || fila >= len(matrizMulti.InternalValor) {
				v.TablaError.NewErrorSemantico(ctx.GetStart(),
					fmt.Sprintf("Índice de fila fuera de rango: %d", fila))
				return &tiposDeDato.ValorEntero{InternalValor: -1}
			}
			// Retornar la fila como vector
			return NewTipoVector(matrizMulti.InternalValor[fila], matrizMulti.ItemType, matrizMulti.BaseType)
		} else {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Número incorrecto de índices para matriz multidimensional")
			return nil
		}
	}

	// Resto del código existente para vectores simples y matrices regulares
	structType := tiposDeDato.TIPO_VECTOR
	index := v.Visit(ctx.Expr(0)).(tiposDeDato.ValorInterno)

	if len(ctx.AllExpr()) != 1 {
		structType = tiposDeDato.TIPO_MATRIZ
	}

	if structType == tiposDeDato.TIPO_VECTOR {
		switch vectorValue := variable.Value.(type) {
		case *TipoVector:
			indexValue := index.(*tiposDeDato.ValorEntero).InternalValor
			if !vectorValue.ValidIndex(indexValue) {
				v.TablaError.NewErrorSemantico(ctx.GetStart(),
					fmt.Sprintf("Índice fuera de rango: %d. El vector tiene %d elementos (rango válido: 0-%d)",
						indexValue, len(vectorValue.InternalValor), len(vectorValue.InternalValor)-1))
				return &VectorItemReference{
					Vector: vectorValue,
					Index:  indexValue,
					Value:  &tiposDeDato.ValorEntero{InternalValor: -1},
				}
			}
			return &VectorItemReference{
				Vector: vectorValue,
				Index:  indexValue,
				Value:  vectorValue.Get(indexValue),
			}
		default:
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "La variable "+varName+" no es un vector")
			return nil
		}
	} else if structType == tiposDeDato.TIPO_MATRIZ {
		switch TipoMatriz := variable.Value.(type) {
		case *TipoMatriz:
			if !TipoMatriz.IndicesValidos(indexes) {
				v.TablaError.NewErrorSemantico(ctx.GetStart(),
					fmt.Sprintf("Índice fuera de rango: %v. La matriz tiene dimensiones válidas", indexes))
				return nil
			}
			return &ElementoMatriz{
				Matriz: TipoMatriz,
				Index:  indexes,
				Value:  TipoMatriz.Get(indexes),
			}
		default:
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "La variable "+varName+" no es una matriz")
			return nil
		}
	}

	return nil
}

func (v *PatronVIsitor) VisitAsignacionDirecta(ctx *parser.AsignacionDirectaContext) interface{} {

	varName := v.Visit(ctx.Id_pattern()).(string)
	varValue := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	variable := v.RegistroAmbito.GetVariable(varName)

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variable "+varName+" no encontrada")
	} else {

		// copy object
		if obj, ok := varValue.(*TipoObjeto); ok {
			varValue = obj.Copy()
		}

		if EsTipoVector(varValue.Type()) {
			varValue = varValue.Copy()
		}

		canMutate := true

		if v.RegistroAmbito.AmbitoActual.isStruct {
			canMutate = v.RegistroAmbito.EsEntornoMutable()
		}

		ok, msg := variable.AsignarVariable(varValue, canMutate)

		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		}
	}

	return nil

}

func (v *PatronVIsitor) VisitAsignacionAritmetica(ctx *parser.AsignacionAritmeticaContext) interface{} {
	varName := v.Visit(ctx.Id_pattern()).(string)

	variable := v.RegistroAmbito.GetVariable(varName)

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variable "+varName+" no encontrada")
	} else {

		leftValue := variable.Value
		rightValue := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

		op := string(ctx.GetOp().GetText()[0])

		strat, ok := ExpresionBinaria[op]

		if !ok {
			log.Fatal("Operador binario no encontrado: " + op)
		}

		ok, msg, varValue := strat.ValidarExp(leftValue, rightValue)

		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
			return nil
		}

		canMutate := true

		if v.RegistroAmbito.AmbitoActual.isStruct {
			canMutate = v.RegistroAmbito.EsEntornoMutable()
		}

		ok, msg = variable.AsignarVariable(varValue, canMutate)

		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		}
	}

	return nil
}

func (v *PatronVIsitor) VisitAsignacionVector(ctx *parser.AsignacionVectorContext) interface{} {

	rightValue := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	switch itemRef := v.Visit(ctx.Vector_item()).(type) {
	case *VectorItemReference:

		leftValue := itemRef.Value

		// check type, todo: improve cast -> ¿? idk what i was thinking
		if rightValue.Type() != itemRef.Vector.ItemType {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "No se puede asignar un valor de tipo "+rightValue.Type()+" a un vector de tipo "+itemRef.Vector.ItemType)
			return nil
		}
		op := string(ctx.GetOp().GetText()[0])

		if op == "=" {
			itemRef.Vector.InternalValor[itemRef.Index] = rightValue
			return nil
		}

		strat, ok := ExpresionBinaria[op]

		if !ok {
			log.Fatal("Binary operator not found")
		}

		ok, msg, varValue := strat.ValidarExp(leftValue, rightValue)

		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
			return nil
		}

		itemRef.Vector.InternalValor[itemRef.Index] = varValue

		return nil
	case *ElementoMatriz:
		leftValue := itemRef.Value

		// check type, todo: improve cast -> ¿? idk what i was thinking
		if rightValue.Type() != EliminarCorchetes(itemRef.Matriz.Type()) {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "No se puede asignar un valor de tipo "+rightValue.Type()+" a una matriz de tipo "+EliminarCorchetes(itemRef.Matriz.Type()))
			return nil
		}

		op := string(ctx.GetOp().GetText()[0])

		if op == "=" {
			itemRef.Matriz.Set(itemRef.Index, rightValue)
			return nil
		}

		strat, ok := ExpresionBinaria[op]

		if !ok {
			log.Fatal("Binary operator not found")
		}

		ok, msg, varValue := strat.ValidarExp(leftValue, rightValue)

		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
			return nil
		}

		itemRef.Matriz.Set(itemRef.Index, varValue)
		return nil
	}

	return nil
}

func (v *PatronVIsitor) VisitAssignVectorItem(ctx *parser.AssignVectorItemContext) interface{} {

	// Obtener la referencia al elemento del vector/matriz
	vectorItemRef := v.Visit(ctx.Vector_item())
	if vectorItemRef == nil {
		return nil
	}

	// NUEVO: Verificar si es un error (ValorEntero con -1)
	if errorVal, ok := vectorItemRef.(*tiposDeDato.ValorEntero); ok && errorVal.InternalValor == -1 {
		// Ya se reportó el error en VisitVectorItem, no hacer nada más
		return nil
	}

	// Evaluar el nuevo valor a asignar
	newValue := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	switch itemRef := vectorItemRef.(type) {
	case *VectorItemReference:
		// Caso: vector simple numeros[2] = valor

		// Verificar compatibilidad de tipos
		if newValue.Type() != itemRef.Vector.ItemType {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("No se puede asignar un valor de tipo %s a un elemento de tipo %s",
					newValue.Type(), itemRef.Vector.ItemType))
			return nil
		}

		// Asignar el nuevo valor
		itemRef.Vector.InternalValor[itemRef.Index] = newValue
		return nil

	case *ElementoMatriz:
		// Caso: matriz antigua TipoMatriz[indices] = valor

		// Verificar compatibilidad de tipos para matriz
		expectedType := EliminarCorchetes(itemRef.Matriz.Type())
		if newValue.Type() != expectedType {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("No se puede asignar un valor de tipo %s a un elemento de matriz de tipo %s",
					newValue.Type(), expectedType))
			return nil
		}

		// Asignar el nuevo valor en la matriz
		itemRef.Matriz.Set(itemRef.Index, newValue)
		return nil

	case *MatrizMultiItemReference:
		// Caso: matriz multidimensional mtx2[0][1] = valor

		// Verificar compatibilidad de tipos
		if newValue.Type() != itemRef.Matriz.BaseType {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("No se puede asignar un valor de tipo %s a un elemento de matriz de tipo %s",
					newValue.Type(), itemRef.Matriz.BaseType))
			return nil
		}

		// Verificar que los índices estén en rango ANTES de intentar asignar
		if !itemRef.Matriz.ValidIndex(itemRef.Fila, itemRef.Columna) {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("Índice fuera de rango: [%d][%d]. La matriz tiene %d filas",
					itemRef.Fila, itemRef.Columna, len(itemRef.Matriz.InternalValor)))
			return nil
		}

		// Asignar el nuevo valor
		itemRef.Matriz.Set(itemRef.Fila, itemRef.Columna, newValue)
		return nil

	default:
		// MEJORAR: No mostrar "Error interno" si ya se reportó el error de índice
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "No se puede realizar la asignación debido a un error previo")
		return nil
	}
}

func (v *PatronVIsitor) VisitID_Patron(ctx *parser.ID_PatronContext) interface{} {
	return ctx.GetText()
}

func (v *PatronVIsitor) VisitIntLiteral(ctx *parser.IntLiteralContext) interface{} {

	intVal, _ := strconv.Atoi(ctx.GetText())

	return &tiposDeDato.ValorEntero{
		InternalValor: intVal,
	}

}

func (v *PatronVIsitor) VisitFloatLiteral(ctx *parser.FloatLiteralContext) interface{} {

	floatVal, _ := strconv.ParseFloat(ctx.GetText(), 64)

	return &tiposDeDato.ValorDecimal{
		InternalValor: floatVal,
	}

}

func (v *PatronVIsitor) VisitStringLiteral(ctx *parser.StringLiteralContext) interface{} {
	// remove quotes
	stringVal := ctx.GetText()[1 : len(ctx.GetText())-1]

	// \" \\ \n \r \
	stringVal = strings.ReplaceAll(stringVal, "\\\"", "\"")
	stringVal = strings.ReplaceAll(stringVal, "\\\\", "\\")
	stringVal = strings.ReplaceAll(stringVal, "\\n", "\n")
	stringVal = strings.ReplaceAll(stringVal, "\\r", "\r")

	// CAMBIO: Las cadenas con comillas dobles siempre son string
	// Si quieres caracteres individuales, usa comillas simples (si las tienes implementadas)
	return &tiposDeDato.ValorCadena{
		InternalValor: stringVal,
	}
}

func (v *PatronVIsitor) VisitBoolLiteral(ctx *parser.BoolLiteralContext) interface{} {

	boolVal, _ := strconv.ParseBool(ctx.GetText())

	return &tiposDeDato.ValorBool{
		InternalValor: boolVal,
	}

}

func (v *PatronVIsitor) VisitNilLiteral(ctx *parser.NilLiteralContext) interface{} {
	return tiposDeDato.NuloPorDefecto
}

func (v *PatronVIsitor) VisitLiteralExp(ctx *parser.LiteralExpContext) interface{} {
	return v.Visit(ctx.Literal())
}

func (v *PatronVIsitor) VisitIdExp(ctx *parser.IdExpContext) interface{} {
	varName := ctx.Id_pattern().GetText()

	variable := v.RegistroAmbito.GetVariable(varName)

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variable "+varName+" no encontrada")
		return tiposDeDato.NuloPorDefecto
	}

	// ? pointer
	return variable.Value
}

func (v *PatronVIsitor) VisitParentecisExp(ctx *parser.ParentecisExpContext) interface{} {
	return v.Visit(ctx.Expr())
}

func (v *PatronVIsitor) VisitVectorItemExp(ctx *parser.VectorItemExpContext) interface{} {

	result := v.Visit(ctx.Vector_item())

	switch itemRef := result.(type) {
	case *VectorItemReference:
		return itemRef.Value
	case *ElementoMatriz:
		return itemRef.Value
	case *MatrizMultiItemReference:
		return itemRef.Value
	case *tiposDeDato.ValorEntero:
		// Caso de error (índice fuera de rango)
		if itemRef.InternalValor == -1 {
			return &tiposDeDato.ValorEntero{InternalValor: -1}
		}
		return itemRef
	}
	return tiposDeDato.NuloPorDefecto
}

func (v *PatronVIsitor) VisitFunctionCallExp(ctx *parser.FunctionCallExpContext) interface{} {
	return v.Visit(ctx.Func_call())
}

func (v *PatronVIsitor) VisitVectorExp(ctx *parser.VectorExpContext) interface{} {
	return v.Visit(ctx.Vector_expr())
}

func (v *PatronVIsitor) VisitExpUnary(ctx *parser.ExpUnaryContext) interface{} {

	exp := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	strat, ok := ExpresionesUnarias[ctx.GetOp().GetText()]

	if !ok {
		log.Fatal("Operador Unario no encontrado: " + ctx.GetOp().GetText())
	}

	ok, msg, result := strat.ValidarExp(exp)

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetOp(), msg)
		return tiposDeDato.NuloPorDefecto
	}

	return result

}

func (v *PatronVIsitor) VisitExpBinario(ctx *parser.ExpBinarioContext) interface{} {

	op := ctx.GetOp().GetText()
	left := v.Visit(ctx.GetLeft()).(tiposDeDato.ValorInterno)

	earlyCheck, ok := RetornoAnticipado[op]

	if ok {
		ok, _, result := earlyCheck.ValidarExp(left)

		if ok {
			return result
		}
	}

	right := v.Visit(ctx.GetRight()).(tiposDeDato.ValorInterno)

	strat, ok := ExpresionBinaria[op]

	if !ok {
		log.Fatal("Operador binario no encontrado: " + op)
	}

	ok, msg, result := strat.ValidarExp(left, right)

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetOp(), msg)
		return tiposDeDato.NuloPorDefecto
	}

	return result
}

func (v *PatronVIsitor) VisitIFstmt(ctx *parser.IFstmtContext) interface{} {

	runChain := true

	for _, ifStmt := range ctx.AllIf_chain() {

		runChain = !v.Visit(ifStmt).(bool)

		if !runChain {
			break
		}
	}

	if runChain && ctx.Else_stmt() != nil {
		v.Visit(ctx.Else_stmt())
	}

	return nil
}

func (v *PatronVIsitor) VisitIFcadena(ctx *parser.IFcadenaContext) interface{} {

	condition := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	if condition.Type() != tiposDeDato.TIPO_BOOLEAN {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La condicion del if debe ser un booleano")
		return false

	}

	if condition.(*tiposDeDato.ValorBool).InternalValor {

		// Push scope
		v.RegistroAmbito.PushScope("if")

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		// Pop scope
		v.RegistroAmbito.PopScope()

		return true
	}

	return false
}

func (v *PatronVIsitor) VisitElseStmt(ctx *parser.ElseStmtContext) interface{} {

	// Push scope
	v.RegistroAmbito.PushScope("else")

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	// Pop scope
	v.RegistroAmbito.PopScope()

	return nil
}

func (v *PatronVIsitor) VisitSwitchStmt(ctx *parser.SwitchStmtContext) interface{} {

	mainValue := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	v.RegistroAmbito.PushScope("switch")

	// Push break switchItem to call stack [breakable]
	switchItem := &LlamadaFunciones{
		RetornarValor: tiposDeDato.NuloPorDefecto,
		Type: []string{
			Detener,
		},
	}

	v.PilaLlamada.Push(switchItem)

	// handle break statements from call stack
	defer func() {

		v.RegistroAmbito.PopScope()       // pop switch scope
		v.PilaLlamada.Limpiar(switchItem) // Limpiar item if it's still in call stack

		if item, ok := recover().(*LlamadaFunciones); item != nil && ok {

			// Not a switch item, propagate panic
			if item != switchItem {
				panic(item)
			}

			return // break
		}
	}()

	visited := false

	// evaluate cases
	for _, switchCase := range ctx.AllSwitch_case() {

		caseValue := v.GetCaseValue(switchCase)

		// ? use binary strat
		if caseValue.Type() != mainValue.Type() {
			// warning
			continue
		}

		if caseValue.Value() == mainValue.Value() {
			v.Visit(switchCase)
			visited = true
			break // implicit break
		}

	}

	// evaluate default
	if ctx.Default_case() != nil && !visited {
		v.Visit(ctx.Default_case())
	}

	return nil
}

func (v *PatronVIsitor) GetCaseValue(tree antlr.ParseTree) tiposDeDato.ValorInterno {

	switch val := tree.(type) {
	case *parser.SwitchCaseContext:
		return v.Visit(val.Expr()).(tiposDeDato.ValorInterno)
	default:
		return nil
	}

}

func (v *PatronVIsitor) VisitSwitchCase(ctx *parser.SwitchCaseContext) interface{} {

	// * all cases inside switch case will share the same scope

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	return nil
}

func (v *PatronVIsitor) VisitDefaultCase(ctx *parser.DefaultCaseContext) interface{} {
	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}
	return nil
}

func (v *PatronVIsitor) VisitWhileStmt(ctx *parser.WhileStmtContext) interface{} {

	condition := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)
	// Push scope
	whileScope := v.RegistroAmbito.PushScope("while")

	// Push whileItem to call stack [breakable, continuable]
	whileItem := &LlamadaFunciones{
		RetornarValor: tiposDeDato.NuloPorDefecto,
		Type: []string{
			Detener,
			Continuar,
		},
	}

	v.PilaLlamada.Push(whileItem)

	v.VisitInnerWhile(ctx, condition, whileScope, whileItem)

	v.RegistroAmbito.PopScope()      // pop while scope
	v.PilaLlamada.Limpiar(whileItem) // clean item if it's still in call stack

	return nil
}

func (v *PatronVIsitor) VisitInnerWhile(ctx *parser.WhileStmtContext, condition tiposDeDato.ValorInterno, whileScope *AmbitoBase, whileItem *LlamadaFunciones) {

	// ? use binary strat
	if condition.Type() != tiposDeDato.TIPO_BOOLEAN {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La condicion del ciclo debe ser un booleano")
		return
	}

	// reset scope
	whileScope.Reset()

	// handle break and continue statements from call stack
	defer func() {

		if item, ok := recover().(*LlamadaFunciones); item != nil && ok {

			// Not a while item, propagate panic
			if item != whileItem {
				panic(item)
			}

			// Continue
			if item.EsAccion(Continuar) {
				item.ResetAccion()                                         // reset action, can be used again
				condition = v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno) // update condition
				v.VisitInnerWhile(ctx, condition, whileScope, whileItem)   // continue

			} else if item.EsAccion(Detener) {
				// Break
				return
			}
		}
	}()

	for condition.(*tiposDeDato.ValorBool).InternalValor {

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		condition = v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

		if condition.Type() != tiposDeDato.TIPO_BOOLEAN {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "La condicion del ciclo debe ser un booleano")
			return
		}

		// reset scope
		whileScope.Reset()
	}
}

func (v *PatronVIsitor) VisitForStmt(ctx *parser.ForStmtContext) interface{} {

	varName := ctx.ID().GetText()
	var iterableItem *TipoVector = TipoVectorVacio

	if ctx.Range_() != nil {
		rangeItem, ok := v.Visit(ctx.Range_()).(*TipoVector)

		if !ok {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "El valor del rango debe ser un vector")
			return nil
		}

		iterableItem = rangeItem
	}

	if ctx.Expr() != nil {
		iterableValue := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

		if EsTipoVector(iterableValue.Type()) {
			iterableItem = iterableValue.(*TipoVector)
		} else if iterableValue.Type() == tiposDeDato.TIPO_CADENA {
			iterableItem = CadenaAVector(iterableValue.(*tiposDeDato.ValorCadena))
		} else {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "El valor del rango debe ser un vector o una cadena")
			return nil
		}
	}

	if iterableItem.Size() == 0 {
		return nil
	}

	// Push scope outer scope
	outerForScope := v.RegistroAmbito.PushScope("outer_for")

	// create the associated variable to the iterable
	iterableVariable, msg := outerForScope.AgregarVariable(varName, iterableItem.ItemType, iterableItem.Current(), true, false, ctx.ID().GetSymbol())

	if iterableVariable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		log.Fatal("Esto no deberia pasar, variable no creada: " + varName)
		return nil
	}

	// Push forItem to call stack [breakable, continuable]

	forItem := &LlamadaFunciones{
		RetornarValor: tiposDeDato.NuloPorDefecto,
		Type: []string{
			Detener,
			Continuar,
		},
	}

	v.PilaLlamada.Push(forItem)

	// Push inner for scope
	innerForScope := v.RegistroAmbito.PushScope("inner_for")

	v.VisitInnerFor(ctx, outerForScope, innerForScope, forItem, iterableItem, iterableVariable)

	iterableItem.Reset()
	v.RegistroAmbito.PopScope()    // pop inner for scope
	v.RegistroAmbito.PopScope()    // pop outer for scope
	v.PilaLlamada.Limpiar(forItem) // ? Limpiar item if it's still in call stack

	return nil
}

func (v *PatronVIsitor) VisitInnerFor(ctx *parser.ForStmtContext, outerForScope *AmbitoBase, innerForScope *AmbitoBase, forItem *LlamadaFunciones, iterableItem *TipoVector, iterableVariable *Variable) {

	// handle break and continue statements from call stack
	defer func() {

		// reset scope
		innerForScope.Reset()
		if item, ok := recover().(*LlamadaFunciones); item != nil && ok {

			// Not a for item, propagate panic
			if item != forItem {
				panic(item)
			}

			// Continue
			if item.EsAccion(Continuar) {
				item.ResetAccion()                                                                          // reset action, can be used again
				iterableItem.Next()                                                                         // next item
				v.VisitInnerFor(ctx, outerForScope, innerForScope, forItem, iterableItem, iterableVariable) // continue
			}

			// Break
			if item.EsAccion(Detener) {
				return
			}

		}
	}()

	// iterableItem.Size()
	for iterableItem.CurrentIndex < iterableItem.Size() {

		// update variable value
		iterableVariable.Value = iterableItem.Current()

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		iterableItem.Next()
		innerForScope.Reset()
	}
}

func (v *PatronVIsitor) VisitRangoNum(ctx *parser.RangoNumContext) interface{} {

	leftExpr := v.Visit(ctx.Expr(0)).(tiposDeDato.ValorInterno)
	rightExpr := v.Visit(ctx.Expr(1)).(tiposDeDato.ValorInterno)

	if leftExpr.Type() != tiposDeDato.TIPO_ENTERO || rightExpr.Type() != tiposDeDato.TIPO_ENTERO {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Los valores de los rangos deben ser enteros")
		return tiposDeDato.NuloPorDefecto
	}

	left := leftExpr.(*tiposDeDato.ValorEntero).InternalValor
	right := rightExpr.(*tiposDeDato.ValorEntero).InternalValor

	if left > right {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El valor izquierdo del rango debe ser menor o igual al valor derecho")
	}

	var values []tiposDeDato.ValorInterno

	for i := left; i <= right; i++ {
		values = append(values, &tiposDeDato.ValorEntero{
			InternalValor: i,
		})
	}

	return &TipoVector{
		InternalValor: values,
		CurrentIndex:  0,
		ItemType:      tiposDeDato.TIPO_ENTERO,
	}
}

func (v *PatronVIsitor) VisitGuardStmt(ctx *parser.GuardStmtContext) interface{} {

	condition := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	if condition.Type() != tiposDeDato.TIPO_BOOLEAN {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La condicion del guard debe ser un booleano")
	}

	if !condition.(*tiposDeDato.ValorBool).InternalValor {

		// Push scope
		v.RegistroAmbito.PushScope("guard")

		for _, stmt := range ctx.AllStmt() {
			v.Visit(stmt)
		}

		// Pop scope
		v.RegistroAmbito.PopScope()
	}

	return nil
}

func (v *PatronVIsitor) VisitReturnStmt(ctx *parser.ReturnStmtContext) interface{} {

	exits, item := v.PilaLlamada.PuedeRetornar()

	if !exits {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La sentencia return debe estar dentro de una funcion")
		return nil
	}

	item.RetornarValor = tiposDeDato.NuloPorDefecto
	item.Accion = Retornar

	if ctx.Expr() != nil {
		item.RetornarValor = v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)
	}

	panic(item)
}

func (v *PatronVIsitor) VisitBreakStmt(ctx *parser.BreakStmtContext) interface{} {

	exits, item := v.PilaLlamada.PuedeDetener()

	if !exits {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La sentencia break debe estar dentro de un ciclo o un switch")
		return nil
	}

	item.Accion = Detener
	panic(item)
}

func (v *PatronVIsitor) VisitContinueStmt(ctx *parser.ContinueStmtContext) interface{} {

	exits, item := v.PilaLlamada.PuedeContinuar()

	if !exits {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "La sentencia continue debe estar dentro de un ciclo")
		return nil
	}

	item.Accion = Continuar
	panic(item)
}

func (v *PatronVIsitor) VisitLlamarFuncion(ctx *parser.LlamarFuncionContext) interface{} {

	canditateName := v.Visit(ctx.Id_pattern()).(string)

	// PRIMERO: Buscar en funciones nativas
	if nativeFunc, exists := DefaultBuiltInFunctions[canditateName]; exists {
		args := make([]*Argumento, 0)
		if ctx.Arg_list() != nil {
			args = v.Visit(ctx.Arg_list()).([]*Argumento)
		}

		RetornarValor, ok, msg := nativeFunc.Exec(v.GetReplContext(), args)
		if !ok {
			if msg != "" {
				v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
			}
			return tiposDeDato.NuloPorDefecto
		}
		return RetornarValor
	}

	// SEGUNDO: Buscar funciones definidas por el usuario
	funcObj, msg1 := v.RegistroAmbito.GetFuncion(canditateName)
	structObj, msg2 := v.RegistroAmbito.AmbitoGlobal.GetEstructura(canditateName)

	if funcObj == nil && structObj == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg1+msg2)
		return tiposDeDato.NuloPorDefecto
	}

	args := make([]*Argumento, 0)
	if ctx.Arg_list() != nil {
		args = v.Visit(ctx.Arg_list()).([]*Argumento)
	}

	// struct has priority over func
	if structObj != nil {
		if ArgumentoValidoEstructura(args) {
			return NewTipoObjeto(v, canditateName, ctx.Id_pattern().GetStart(), args, false)
		} else {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Si bien "+canditateName+" es un struct, no se puede llamar a su constructor con los argumentos especificados. Ni tampoco es una funcion.")
			return tiposDeDato.NuloPorDefecto
		}
	}

	switch funcObj := funcObj.(type) {
	case *FuncionNativa:
		RetornarValor, ok, msg := funcObj.Exec(v.GetReplContext(), args)

		if !ok {
			if msg != "" {
				v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
			}
			return tiposDeDato.NuloPorDefecto
		}

		return RetornarValor

	case *Funcion:
		funcObj.Exec(v, args, ctx.GetStart())
		return funcObj.RetornarValor

	case *FuncionNativaObjeto:
		funcObj.Exec(v, args, ctx.GetStart())
		return funcObj.RetornarValor

	default:
		log.Fatal("Tipo de funcion no soportado: ", funcObj)
	}

	return tiposDeDato.NuloPorDefecto
}

func (v *PatronVIsitor) VisitArgList(ctx *parser.ArgListContext) interface{} {

	args := make([]*Argumento, 0)

	for _, arg := range ctx.AllFunc_arg() {
		args = append(args, v.Visit(arg).(*Argumento))
	}

	return args

}

func (v *PatronVIsitor) VisitFuncionArg(ctx *parser.FuncionArgContext) interface{} {
	argName := ""
	passByReference := false

	var argValue tiposDeDato.ValorInterno = tiposDeDato.NuloPorDefecto
	var argVariableRef *Variable = nil

	if ctx.Id_pattern() != nil {
		// NO asignar el nombre de la variable como nombre del argumento
		variableName := ctx.Id_pattern().GetText()
		argVariableRef = v.RegistroAmbito.GetVariable(variableName)

		if argVariableRef != nil {
			argValue = argVariableRef.Value
		} else {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Variable "+variableName+" no encontrada")
		}
	} else {
		argValue = v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)
	}

	// Solo asignar argName si hay un nombre explícito con DOS_PUNTOS (argumentos con nombre)
	if ctx.ID() != nil {
		argName = ctx.ID().GetText()
	}

	if ctx.ANPERSAND() != nil {
		passByReference = true
	}

	return &Argumento{
		Name:            argName,
		Value:           argValue,
		PassByReference: passByReference,
		Token:           ctx.GetStart(),
		VariableRef:     argVariableRef,
	}
}

func (v *PatronVIsitor) VisitFuncionDeclerada(ctx *parser.FuncionDecleradaContext) interface{} {

	if v.RegistroAmbito.AmbitoActual == v.RegistroAmbito.AmbitoGlobal {
		// aready declared by dcl_visitor
		return nil
	}

	if v.RegistroAmbito.AmbitoActual != v.RegistroAmbito.AmbitoGlobal && !v.RegistroAmbito.AmbitoActual.isStruct {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Las funciones solo pueden ser declaradas en el scope global o en un struct")
	}

	funcName := ctx.ID().GetText()

	params := make([]*Parametros, 0)

	if ctx.Param_list() != nil {
		params = v.Visit(ctx.Param_list()).([]*Parametros)
	}

	if len(params) > 0 {

		baseParamType := params[0].ParamType()

		for _, param := range params {
			if param.ParamType() != baseParamType {
				v.TablaError.NewErrorSemantico(param.Token, "Todos los parametros de la funcion deben ser del mismo tipo")
				return nil
			}
		}
	}

	returnType := tiposDeDato.TIPO_NULO
	var returnTypeToken antlr.Token = nil

	if ctx.Type_() != nil {
		returnType = v.Visit(ctx.Type_()).(string)
		returnTypeToken = ctx.Type_().GetStart()
	}

	body := ctx.AllStmt()

	function := &Funcion{ // pointer ?
		Name:            funcName,
		Parametros:      params,
		ReturnType:      returnType,
		Body:            body,
		DeclScope:       v.RegistroAmbito.AmbitoActual,
		ReturnTypeToken: returnTypeToken,
		Token:           ctx.GetStart(),
	}

	ok, msg := v.RegistroAmbito.AgregarFuncion(funcName, function)

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		return nil
	}

	return function
}

func (v *PatronVIsitor) VisitParamList(ctx *parser.ParamListContext) interface{} {

	params := make([]*Parametros, 0)

	for _, param := range ctx.AllFunc_param() {
		params = append(params, v.Visit(param).(*Parametros))
	}

	return params
}

func (v *PatronVIsitor) VisitFuncParam(ctx *parser.FuncParamContext) interface{} {

	externName := ""
	innerName := ""

	// at least ID(0) is defined
	// only 1 ID defined
	if ctx.ID(1) == nil {
		// innerName : type
		// _ : type
		innerName = ctx.ID(0).GetText()
	} else {
		// externName innerName : type
		externName = ctx.ID(0).GetText()
		innerName = ctx.ID(1).GetText()
	}

	passByReference := false

	if ctx.RW_INOUT() != nil {
		passByReference = true
	}

	paramType := v.Visit(ctx.Type_()).(string)

	return &Parametros{
		ExternName:      externName,
		InnerName:       innerName,
		PassByReference: passByReference,
		Type:            paramType,
		Token:           ctx.GetStart(),
	}

}

// * Structs
func (v *PatronVIsitor) VisitDeclararStruct(ctx *parser.DeclararStructContext) interface{} {
	if v.RegistroAmbito.AmbitoActual != v.RegistroAmbito.AmbitoGlobal {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Los structs solo pueden ser declaradas en el scope global")
		return nil
	}

	structName := ctx.ID().GetText()

	// NUEVO: Crear scope temporal para procesar los campos
	structScope := NewAmbitoBase(structName+"_template", true)
	prevScope := v.RegistroAmbito.AmbitoActual
	v.RegistroAmbito.AmbitoActual = structScope

	// Procesar cada campo del struct
	for _, fieldCtx := range ctx.AllStruct_prop() {
		v.Visit(fieldCtx) // Esto llama a VisitStructAttr que ya tienes implementado
	}

	// Restaurar scope
	v.RegistroAmbito.AmbitoActual = prevScope

	// Crear el struct con los campos procesados (mantener compatibilidad)
	newStruct := &Struct{
		Name:   structName,
		Fields: ctx.AllStruct_prop(), // Mantener para compatibilidad con NewTipoObjeto
		Token:  ctx.GetStart(),
	}

	structAdded, msg := v.RegistroAmbito.AmbitoGlobal.AgregarEstructura(structName, newStruct)
	if !structAdded {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}

	return nil
}

func (v *PatronVIsitor) VisitStructAttr(ctx *parser.StructAttrContext) interface{} {

	// Verificar que ctx y sus elementos no sean nil
	if ctx == nil {
		return nil
	}

	if ctx.Type_() == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Tipo de campo no especificado")
		return nil
	}

	if ctx.ID() == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Nombre de campo no especificado")
		return nil
	}

	// Nueva sintaxis: type ID = expr?
	fieldTypeResult := v.Visit(ctx.Type_())
	if fieldTypeResult == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Error al procesar tipo de campo")
		return nil
	}

	fieldType := fieldTypeResult.(string) // string, int, bool
	fieldName := ctx.ID().GetText()       // Nombre, Edad, EsEstudiante

	var defaultValue tiposDeDato.ValorInterno = tiposDeDato.ValorNoIniPorDefecto

	// Si hay valor por defecto: string Nombre = "Alice"
	if ctx.Expr() != nil {
		exprResult := v.Visit(ctx.Expr())
		if exprResult == nil {
			v.TablaError.NewErrorSemantico(ctx.GetStart(), "Error al procesar valor por defecto")
			return nil
		}

		defaultValue = exprResult.(tiposDeDato.ValorInterno)

		// Verificar que el tipo coincida
		if defaultValue.Type() != fieldType {
			v.TablaError.NewErrorSemantico(ctx.GetStart(),
				fmt.Sprintf("No se puede asignar un valor de tipo %s al campo %s de tipo %s",
					defaultValue.Type(), fieldName, fieldType))
			return nil
		}
	}

	// Los campos de struct son siempre mutables por defecto
	variable, msg := v.RegistroAmbito.AgregarVariable(
		fieldName, fieldType, defaultValue,
		false, // mutable (siempre false para struct fields)
		false, // inicializada
		ctx.ID().GetSymbol())

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		return nil
	}

	return variable // Retornar la variable creada
}

func (v *PatronVIsitor) VisitStructFunc(ctx *parser.StructFuncContext) interface{} {

	funcDcl := v.Visit(ctx.Func_dcl())

	if ctx.RW_MUTATING() != nil {
		structFunc, ok := funcDcl.(*Funcion)

		if !ok {
			return nil
		}
		structFunc.IsMutating = true
	}

	return nil
}

func (v *PatronVIsitor) VisitStructInit(ctx *parser.StructInitContext) interface{} {

	if ctx == nil {
		return tiposDeDato.NuloPorDefecto
	}

	if ctx.ID() == nil {
		return tiposDeDato.NuloPorDefecto
	}

	structName := ctx.ID().GetText()

	// Verificar que el struct existe
	structDef, msg := v.RegistroAmbito.AmbitoGlobal.GetEstructura(structName)
	if structDef == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		return tiposDeDato.NuloPorDefecto
	}

	// Crear argumentos para el constructor existente
	args := make([]*Argumento, 0)

	if ctx.Struct_init_list() != nil {

		initFieldsResult := v.Visit(ctx.Struct_init_list())
		if initFieldsResult == nil {
			return tiposDeDato.NuloPorDefecto
		}

		initFields, ok := initFieldsResult.([]*StructInitValue)
		if !ok {
			return tiposDeDato.NuloPorDefecto
		}

		for _, initField := range initFields {
			if initField == nil {
				continue
			}

			args = append(args, &Argumento{
				Name:  initField.Name,
				Value: initField.Value,
				Token: initField.Token,
			})
		}
	}
	// Usar el constructor existente
	result := NewTipoObjeto(v, structName, ctx.GetStart(), args, false)

	return result
}

func (v *PatronVIsitor) VisitStructInitExp(ctx *parser.StructInitExpContext) interface{} {
	return v.Visit(ctx.Struct_init())
}

func (v *PatronVIsitor) VisitStructInitList(ctx *parser.StructInitListContext) interface{} {
	initFields := make([]*StructInitValue, 0)

	for _, fieldCtx := range ctx.AllStruct_init_field() {
		if field := v.Visit(fieldCtx); field != nil {
			if initField, ok := field.(*StructInitValue); ok {
				initFields = append(initFields, initField)
			}
		}
	}

	return initFields
}

func (v *PatronVIsitor) VisitStructInitField(ctx *parser.StructInitFieldContext) interface{} {
	fieldName := ctx.ID().GetText()
	fieldValue := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	return &StructInitValue{
		Name:  fieldName,
		Value: fieldValue,
		Token: ctx.GetStart(),
	}
}

func (v *PatronVIsitor) VisitMetodoStruct(ctx *parser.MetodoStructContext) interface{} {
	fmt.Println("=== DEBUG: ENTRANDO A VisitMetodoStruct ===")

	if v.RegistroAmbito.AmbitoActual != v.RegistroAmbito.AmbitoGlobal {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Los métodos solo pueden ser declarados en el scope global")
		return nil
	}

	// Verificar que el mapa de métodos esté inicializado
	if v.RegistroAmbito.AmbitoGlobal.methods == nil {
		fmt.Println("ERROR: El mapa de métodos NO está inicializado")
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Sistema de métodos no inicializado")
		return nil
	}

	fmt.Printf("DEBUG: Mapa de métodos inicializado correctamente: %v\n", v.RegistroAmbito.AmbitoGlobal.methods)

	// NUEVO: Con la nueva gramática, obtenemos receiver de method_receiver
	receiverCtx := ctx.Method_receiver()

	// Convertir a contexto específico y obtener los IDs
	methodReceiverCtx := receiverCtx.(*parser.MethodReceiverContext)
	receiverName := methodReceiverCtx.ID(0).GetText() // p
	receiverType := methodReceiverCtx.ID(1).GetText() // Persona
	methodName := ctx.ID().GetText()                  // Saludar

	fmt.Printf("DEBUG Método: receiver='%s', type='*%s', method='%s'\n",
		receiverName, receiverType, methodName)

	// Verificar que el struct existe
	_, msg := v.RegistroAmbito.AmbitoGlobal.GetEstructura(receiverType)
	if msg != "" {
		v.TablaError.NewErrorSemantico(ctx.GetStart(),
			fmt.Sprintf("El tipo %s no existe", receiverType))
		return nil
	}
	fmt.Printf("DEBUG: Struct %s existe correctamente\n", receiverType)

	// Procesar parámetros
	params := make([]*Parametros, 0)
	if ctx.Param_list() != nil {
		paramsResult := v.Visit(ctx.Param_list())
		if paramsResult != nil {
			params = paramsResult.([]*Parametros)
		}
	}
	fmt.Printf("DEBUG: Procesados %d parámetros\n", len(params))

	// Procesar tipo de retorno
	returnType := tiposDeDato.TIPO_NULO
	var returnTypeToken antlr.Token = nil
	if ctx.Type_() != nil {
		returnTypeResult := v.Visit(ctx.Type_())
		if returnTypeResult != nil {
			returnType = returnTypeResult.(string)
			returnTypeToken = ctx.Type_().GetStart()
		}
	}
	fmt.Printf("DEBUG: Tipo de retorno: %s\n", returnType)

	// Crear método
	method := &MetodoStruct{
		Name:            methodName,
		ReceiverName:    receiverName,
		ReceiverType:    receiverType,
		Parametros:      params,
		ReturnType:      returnType,
		ReturnTypeToken: returnTypeToken,
		Body:            ctx.AllStmt(),
		DeclScope:       v.RegistroAmbito.AmbitoActual,
		Token:           ctx.GetStart(),
	}
	fmt.Printf("DEBUG: Método %s creado correctamente\n", methodName)

	// Agregar método al ámbito global
	ok, msg := v.RegistroAmbito.AmbitoGlobal.AgregarMetodo(receiverType, methodName, method)
	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		return nil
	}

	fmt.Printf("DEBUG: Método %s agregado correctamente para tipo %s\n", methodName, receiverType)
	fmt.Printf("DEBUG: Total de métodos ahora: %d\n", len(v.RegistroAmbito.AmbitoGlobal.methods))

	// Listar todos los métodos registrados
	fmt.Println("DEBUG: Métodos registrados:")
	for key, _ := range v.RegistroAmbito.AmbitoGlobal.methods {
		fmt.Printf("  - %s\n", key)
	}

	fmt.Println("=== DEBUG: SALIENDO DE VisitMetodoStruct ===")
	return method
}

func (v *PatronVIsitor) VisitStructMethodCall(ctx *parser.StructMethodCallContext) interface{} {
	fmt.Println("=== DEBUG: LLAMANDO A MÉTODO ===")

	instanceName := ctx.Id_pattern().GetText()
	methodName := ctx.ID().GetText()

	fmt.Printf("DEBUG: Instancia='%s', Método='%s'\n", instanceName, methodName)

	// Obtener la instancia
	variable := v.RegistroAmbito.GetVariable(instanceName)
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(),
			fmt.Sprintf("Variable %s no encontrada", instanceName))
		return tiposDeDato.NuloPorDefecto
	}

	// Verificar que sea un struct
	receiver, ok := variable.Value.(*TipoObjeto)
	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(),
			fmt.Sprintf("La variable %s no es un struct", instanceName))
		return tiposDeDato.NuloPorDefecto
	}

	fmt.Printf("DEBUG: Buscando método %s para tipo %s\n", methodName, receiver.ConcretType)

	// Buscar el método
	method, msg := v.RegistroAmbito.AmbitoGlobal.GetMetodo(receiver.ConcretType, methodName)
	if method == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		return tiposDeDato.NuloPorDefecto
	}

	fmt.Printf("DEBUG: Método %s encontrado!\n", methodName)

	// Procesar argumentos
	args := make([]*Argumento, 0)
	if ctx.Arg_list() != nil {
		args = v.Visit(ctx.Arg_list()).([]*Argumento)
	}

	// Ejecutar método
	method.Exec(v, receiver, args, ctx.GetStart())

	fmt.Printf("DEBUG: Método ejecutado, retornando: %v\n", method.RetornarValor)
	return method.RetornarValor
}

func (v *PatronVIsitor) VisitStructMethodExp(ctx *parser.StructMethodExpContext) interface{} {
	return v.Visit(ctx.Struct_method_call())
}

func (v *PatronVIsitor) VisitStructVector(ctx *parser.StructVectorContext) interface{} {

	_type := ctx.ID().GetText()

	stc, msg := v.RegistroAmbito.AmbitoGlobal.GetEstructura(_type)

	if stc == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
		return tiposDeDato.NuloPorDefecto
	}

	return NewTipoVector(make([]tiposDeDato.ValorInterno, 0), "["+_type+"]", _type)
}

func (v *PatronVIsitor) VisitStructExp(ctx *parser.StructExpContext) interface{} {
	return v.Visit(ctx.Struct_vector())
}

func (v *PatronVIsitor) VisitVectorFuncExp(ctx *parser.VectorFuncExpContext) interface{} {
	return v.Visit(ctx.Vector_func())
}

func (v *PatronVIsitor) VisitVectorPropExp(ctx *parser.VectorPropExpContext) interface{} {
	return v.Visit(ctx.Vector_prop())
}

func (v *PatronVIsitor) VisitPropVector(ctx *parser.PropVectorContext) interface{} {

	var objectCandidate tiposDeDato.ValorInterno

	switch itemRef := v.Visit(ctx.Vector_item()).(type) {
	case *VectorItemReference:
		objectCandidate = itemRef.Value
	case *ElementoMatriz:
		objectCandidate = itemRef.Value
	}

	obj, ok := objectCandidate.(*TipoObjeto)

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El item del vector no es un struct")
		return tiposDeDato.NuloPorDefecto
	}

	lastScope := v.RegistroAmbito.AmbitoActual
	v.RegistroAmbito.AmbitoActual = obj.InternalScope

	defer func() {
		v.RegistroAmbito.AmbitoActual = lastScope
	}()

	variable := v.RegistroAmbito.GetVariable(ctx.Id_pattern().GetText())

	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Propiedad "+ctx.Id_pattern().GetText()+" no encontrada item del vector")
		return tiposDeDato.NuloPorDefecto
	}

	return variable.Value
}

func (v *PatronVIsitor) VisitFuncionVector(ctx *parser.FuncionVectorContext) interface{} {

	var objectCandidate tiposDeDato.ValorInterno

	switch itemRef := v.Visit(ctx.Vector_item()).(type) {
	case *VectorItemReference:
		objectCandidate = itemRef.Value
	case *ElementoMatriz:
		objectCandidate = itemRef.Value
	}

	obj, ok := objectCandidate.(*TipoObjeto)

	if !ok {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "El objeto no es un struct")
		return tiposDeDato.NuloPorDefecto
	}

	lastScope := v.RegistroAmbito.AmbitoActual
	v.RegistroAmbito.AmbitoActual = obj.InternalScope

	defer func() {
		v.RegistroAmbito.AmbitoActual = lastScope
	}()

	return v.Visit(ctx.Func_call())
}

func (s *RegistroAmbito) Print() {

	fmt.Println("Global Scope")
	fmt.Println("============")

	fmt.Println("Variables")
	for k, v := range s.AmbitoGlobal.variables {
		fmt.Println(k, v.Value.Value(), v.Type)
	}

	fmt.Println("Funciones")
	for k, v := range s.AmbitoGlobal.functions {
		fmt.Println(k, v)
	}

	fmt.Println("Child Scopes")
	fmt.Println("============")
	fmt.Println("")

	for _, child := range s.AmbitoGlobal.children {

		fmt.Println(child.name)
		fmt.Println("============")

		fmt.Println("Variables")
		for k, v := range child.variables {
			fmt.Println(k, v.Value.Value())
		}

		fmt.Println("Funciones")
		for k, v := range child.functions {
			fmt.Println(k, v)
		}

		fmt.Println("")
	}

}

// Declaración sin 'mut/let/var', con tipo y valor:  ID type = expr
func (v *PatronVIsitor) VisitDeclararSinMutValor(ctx *parser.DeclararSinMutValorContext) interface{} {
	varName := ctx.ID().GetText()
	varType := v.Visit(ctx.Type_()).(string)
	varValue := v.Visit(ctx.Expr()).(tiposDeDato.ValorInterno)

	// Copiar objetos y vectores para evitar aliasing
	if obj, ok := varValue.(*TipoObjeto); ok {
		varValue = obj.Copy()
	}
	if EsTipoVector(varValue.Type()) {
		varValue = varValue.Copy()
	}

	// isConst = false, inferir inicialización completa => inferir=false
	variable, msg := v.RegistroAmbito.AgregarVariable(varName, varType, varValue, false, false, ctx.GetStart())
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}
	return nil
}

// Declaración sin 'mut/let/var', solo tipo:  ID type
func (v *PatronVIsitor) VisitDeclararSinMutTipo(ctx *parser.DeclararSinMutTipoContext) interface{} {
	varName := ctx.ID().GetText()
	varType := v.Visit(ctx.Type_()).(string)

	// valor por defecto (NuloPorDefecto), isConst=false, inferir=true
	variable, msg := v.RegistroAmbito.AgregarVariable(varName, varType, tiposDeDato.NuloPorDefecto, false, true, ctx.GetStart())
	if variable == nil {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), msg)
	}
	return nil
}
