package instrucciones

import (
	"log"
	"main/tiposDeDato"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

// AmbitoBase: representa un ámbito/contexto que contiene variables, funciones y estructuras
type AmbitoBase struct {
	Nombre       string
	Variables    map[string]*Variable
	Funciones    map[string]tiposDeDato.ValorInterno
	Hijos        []*AmbitoBase
	Estructuras  map[string]*Struct
	Padre        *AmbitoBase
	EsEstructura bool
	EsMutante    bool
}

func (s *AmbitoBase) Name() string {
	return s.Nombre
}

func (s *AmbitoBase) Parent() *AmbitoBase {
	return s.Padre
}

func (s *AmbitoBase) Children() []*AmbitoBase {
	return s.Hijos
}

// ValidType: verifica si un tipo es válido (primitivo o estructura definida)
func (s *AmbitoBase) ValidType(_type string) bool {
	_, esTipoEstructura := s.Estructuras[_type]

	return tiposDeDato.PRIMITIVO(_type) || esTipoEstructura
}

// AddChild: agrega un ámbito hijo y establece la relación padre-hijo
func (s *AmbitoBase) AddChild(hijo *AmbitoBase) {
	s.Hijos = append(s.Hijos, hijo)
	hijo.Padre = s
}

// existeVariable: verifica si una variable ya existe en el ámbito actual
func (s *AmbitoBase) existeVariable(variable *Variable) bool {

	if _, ok := s.Variables[variable.Nombre]; ok {
		return true
	}

	return false

}

// Crear función NewAmbitoBase:
func NewAmbitoBase(name string, esEstructura bool) *AmbitoBase {
	return &AmbitoBase{
		Nombre:       name,
		Variables:    make(map[string]*Variable),
		Funciones:    make(map[string]tiposDeDato.ValorInterno),
		Hijos:        make([]*AmbitoBase, 0),
		Estructuras:  make(map[string]*Struct),
		Padre:        nil,
		EsEstructura: esEstructura,
		EsMutante:    false,
	}
}

func (s *AmbitoBase) AgregarVariable(nombre string, tipoVariable string, valor tiposDeDato.ValorInterno, esConstante bool, pudeSerNulo bool, token antlr.Token) (*Variable, string) {

	variable := &Variable{
		Nombre:       nombre,
		Valor:        valor,
		Tipo:         tipoVariable,
		esConstante:  esConstante,
		PuedeSerNulo: pudeSerNulo,
		Token:        token,
	}

	if s.existeVariable(variable) {
		return nil, "La variable " + nombre + " ya existe"
	}

	typesOk, msg := variable.TipoValidacion()

	// incluso si la variable no es válida, la agregamos al ámbito (internamente será nil)
	s.Variables[nombre] = variable

	if !typesOk {
		// reportar error
		return nil, msg
	}

	return variable, ""
}

func (s *AmbitoBase) GetVariable(name string) *Variable {
	// verificar si se refiere a una función/variable de objeto/struct
	if strings.Contains(name, ".") {
		return s.buscarVariableObjeto(name, nil)
	}

	ambitoInicial := s

	for {
		if variable, ok := ambitoInicial.Variables[name]; ok {

			// verificar si se refiere a un puntero
			if variable.Tipo == tiposDeDato.TIPO_PUNTERO {
				return variable.Valor.(*TipoPuntero).VariableAsociada
			}

			return variable
		}

		if ambitoInicial.Padre == nil {
			break
		}

		ambitoInicial = ambitoInicial.Padre
	}

	return nil
}

// obj1.obj2.prop1
func (s *AmbitoBase) buscarVariableObjeto(name string, lastObj tiposDeDato.ValorInterno) *Variable {

	// split name by dot
	parts := strings.Split(name, ".")

	if len(parts) == 0 {
		log.Fatal("no se puede dividir por punto")
		return nil
	}

	if len(parts) == 1 {
		obj, ok := lastObj.(*TipoObjeto)

		if ok {
			return obj.AmbitoInterno.GetVariable(name)
		}

		log.Fatal("no se puede convertir a objeto")
		return nil
	}

	// then parts should be 2 or more

	if lastObj == nil {
		variable := s.GetVariable(parts[0])

		if variable == nil {
			return nil
		}

		obj := variable.Valor

		// obj debe ser un objeto/struct o vector
		switch obj := obj.(type) {
		case *TipoObjeto:
			lastObj = obj
		case *TipoVector:
			lastObj = obj.TipoObjeto
		default:
			return nil
		}

		return s.buscarVariableObjeto(strings.Join(parts[1:], "."), lastObj)
	}

	obj, ok := lastObj.(*TipoObjeto)

	if ok {
		lastObj = obj.AmbitoInterno.GetVariable(parts[0]).Valor

		return s.buscarVariableObjeto(strings.Join(parts[1:], "."), lastObj)
	} else {
		log.Fatal("no se puede convertir a objeto")
		return nil
	}
}

func (s *AmbitoBase) AgregarFuncion(name string, function tiposDeDato.ValorInterno) (bool, string) {

	if _, ok := s.Funciones[name]; ok {
		return false, "La funcion " + name + " ya existe"
	}

	s.Funciones[name] = function

	return true, ""
}

func (s *AmbitoBase) GetFuncion(name string) (tiposDeDato.ValorInterno, string) {

	// verificar si se refiere a una función de objeto/struct
	if strings.Contains(name, ".") {
		return s.buscarFuncion(name, nil)
	}

	ambitoInicial := s

	for {
		if function, ok := ambitoInicial.Funciones[name]; ok {
			return function, ""
		}

		if ambitoInicial.Padre == nil {
			break
		}

		ambitoInicial = ambitoInicial.Padre
	}

	return nil, "La funcion " + name + " no existe"
}

// obj1.obj2.func1()
func (s *AmbitoBase) buscarFuncion(name string, lastObj tiposDeDato.ValorInterno) (tiposDeDato.ValorInterno, string) {

	// split name by dot
	parts := strings.Split(name, ".")

	if len(parts) == 0 {
		log.Fatal("no se puede dividir por punto")
		return nil, ""
	}

	if len(parts) == 1 {
		obj, ok := lastObj.(*TipoObjeto)

		if ok {
			return obj.AmbitoInterno.GetFuncion(name)
		}

		log.Fatal("no se puede convertir a objeto")
		return nil, ""
	}

	// then parts should be 2 or more

	if lastObj == nil {
		variable := s.GetVariable(parts[0])

		if variable == nil {
			return nil, "No se puede acceder a la propiedad " + parts[0]
		}

		obj := variable.Valor

		// obj debe ser un objeto/struct o vector
		switch obj := obj.(type) {
		case *TipoObjeto:
			lastObj = obj
		case *TipoVector:
			lastObj = obj.TipoObjeto
		default:
			return nil, "La propiedad '" + variable.Nombre + "' de tipo " + obj.Type() + " no tiene propiedades"
		}

		return s.buscarFuncion(strings.Join(parts[1:], "."), lastObj)
	}

	obj, ok := lastObj.(*TipoObjeto)

	if ok {
		lastObj = obj.AmbitoInterno.GetVariable(parts[0]).Valor

		return s.buscarFuncion(strings.Join(parts[1:], "."), lastObj)
	} else {
		log.Fatal("no se puede convertir a objeto")
		return nil, ""
	}
}

func (s *AmbitoBase) AgregarEstructura(name string, structValue *Struct) (bool, string) {

	if _, ok := s.Estructuras[name]; ok {
		return false, "La estructura " + name + " ya existe"
	}

	s.Estructuras[name] = structValue
	return true, ""
}

func (s *AmbitoBase) GetEstructura(name string) (*Struct, string) {

	ambitoInicial := s

	for {
		if structValue, ok := ambitoInicial.Estructuras[name]; ok {
			return structValue, ""
		}

		if ambitoInicial.Padre == nil {
			break
		}

		ambitoInicial = ambitoInicial.Padre
	}

	return nil, "La estructura " + name + " no existe"
}

// Reset: limpia todas las variables, hijos y funciones del ámbito
func (s *AmbitoBase) Reset() {
	s.Variables = make(map[string]*Variable)
	s.Hijos = make([]*AmbitoBase, 0)
	s.Funciones = make(map[string]tiposDeDato.ValorInterno)
}

// EsAmbitoMutante: verifica si el ámbito actual o algún padre es mutante
func (s *AmbitoBase) EsAmbitoMutante() bool {
	aux := s

	for {
		if aux.EsMutante {
			return true
		}

		if aux.Padre == nil {
			break
		}

		aux = aux.Padre
	}

	return false
}

// NuevoAmbitoGlobal: constructor que crea el ámbito global con funciones embebidas
func NuevoAmbitoGlobal() *AmbitoBase {
	funcs := make(map[string]tiposDeDato.ValorInterno)

	for k, v := range FuncionesEmbebidas {
		funcs[k] = v
	}

	return &AmbitoBase{
		Nombre:      "global",
		Variables:   make(map[string]*Variable),
		Hijos:       make([]*AmbitoBase, 0),
		Estructuras: make(map[string]*Struct),
		Padre:       nil,
		Funciones:   funcs,
	}
}

// NuevoAmbitoLocal: constructor que crea un nuevo ámbito local
func NuevoAmbitoLocal(name string) *AmbitoBase {
	return &AmbitoBase{
		Nombre:    name,
		Variables: make(map[string]*Variable),
		Funciones: make(map[string]tiposDeDato.ValorInterno),
		Hijos:     make([]*AmbitoBase, 0),
		Padre:     nil,
	}
}

// RegistroAmbito: maneja la pila de ámbitos y el ámbito actual
type RegistroAmbito struct {
	AmbitoGlobal *AmbitoBase
	AmbitoActual *AmbitoBase
}

// PushAmbito: crea y empuja un nuevo ámbito a la pila
func (s *RegistroAmbito) PushAmbito(name string) *AmbitoBase {

	nuevoAmbito := NuevoAmbitoLocal(name)
	s.AmbitoActual.AddChild(nuevoAmbito)
	s.AmbitoActual = nuevoAmbito

	return s.AmbitoActual
}

// PopAmbito: retorna al ámbito padre en la pila
func (s *RegistroAmbito) PopAmbito() {
	s.AmbitoActual = s.AmbitoActual.Parent()
}

// Reset: reinicia el registro al ámbito global
func (s *RegistroAmbito) Reset() {
	s.AmbitoActual = s.AmbitoGlobal
}

// AgregarVariable: delegada para agregar variable en el ámbito actual
func (s *RegistroAmbito) AgregarVariable(name string, varType string, value tiposDeDato.ValorInterno, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {
	return s.AmbitoActual.AgregarVariable(name, varType, value, isConst, allowNil, token)
}

// GetVariable: delegada para obtener variable del ámbito actual
func (s *RegistroAmbito) GetVariable(name string) *Variable {
	return s.AmbitoActual.GetVariable(name)
}

// AgregarFuncion: delegada para agregar función en el ámbito actual
func (s *RegistroAmbito) AgregarFuncion(name string, function tiposDeDato.ValorInterno) (bool, string) {
	return s.AmbitoActual.AgregarFuncion(name, function)
}

// GetFuncion: delegada para obtener función del ámbito actual
func (s *RegistroAmbito) GetFuncion(name string) (tiposDeDato.ValorInterno, string) {
	return s.AmbitoActual.GetFuncion(name)
}

// EsEntornoMutable: verifica si el entorno actual permite mutaciones
func (s *RegistroAmbito) EsEntornoMutable() bool {
	return s.AmbitoActual.EsAmbitoMutante()
}

// NuevoRegistroAmbito: constructor que crea un nuevo registro con ámbito global
func NuevoRegistroAmbito() *RegistroAmbito {
	AmbitoGlobal := NuevoAmbitoGlobal()
	return &RegistroAmbito{
		AmbitoGlobal: AmbitoGlobal,
		AmbitoActual: AmbitoGlobal,
	}
}

// NuevoVectorGlobal: constructor que crea un ámbito específico para vectores
func NuevoVectorGlobal() *AmbitoBase {
	var scope = &AmbitoBase{
		Nombre:    "vector",
		Variables: make(map[string]*Variable),
		Hijos:     make([]*AmbitoBase, 0),
		Funciones: make(map[string]tiposDeDato.ValorInterno),
		Padre:     nil,
	}

	return scope
}

// Reporte: estructuras para generar reportes de la tabla de símbolos
type ReporteTabla struct {
	AmbitoGlobal ReporteAmbito
}

type ReporteAmbito struct {
	Nombre       string
	Variables    []ReporteSimbolos
	Funciones    []ReporteSimbolos
	Estructuras      []ReporteSimbolos
	AmbitosHijos []ReporteAmbito
}

type ReporteSimbolos struct {
	Nombre  string
	Tipo    string
	Linea   int
	Columna int
}

// Report: genera un reporte completo de la tabla de símbolos
func (s *RegistroAmbito) Report() ReporteTabla {
	return ReporteTabla{
		AmbitoGlobal: s.AmbitoActual.Report(),
	}
}

// Report: genera un reporte del ámbito actual incluyendo todos sus símbolos
func (s *AmbitoBase) Report() ReporteAmbito {

	ReporteAmbito := ReporteAmbito{
		Nombre:       s.Nombre,
		Variables:    make([]ReporteSimbolos, 0),
		Funciones:    make([]ReporteSimbolos, 0),
		Estructuras:      make([]ReporteSimbolos, 0),
		AmbitosHijos: make([]ReporteAmbito, 0),
	}

	for _, v := range s.Variables {

		token := v.Token
		line := 0
		column := 0

		if token != nil {
			line = token.GetLine()
			column = token.GetColumn()
		}

		ReporteAmbito.Variables = append(ReporteAmbito.Variables, ReporteSimbolos{
			Nombre:  v.Nombre,
			Tipo:    v.Tipo,
			Linea:   line,
			Columna: column,
		})
	}

	for _, f := range s.Funciones {
		switch funcion := f.(type) {
		case *FuncionNativa:
			// No incluir funciones embebidas en el reporte
			if _, esEmbebida := FuncionesEmbebidas[funcion.Nombre]; !esEmbebida {
				ReporteAmbito.Funciones = append(ReporteAmbito.Funciones, ReporteSimbolos{
					Nombre:  funcion.Nombre,
					Tipo:    funcion.Nombre,
					Linea:   0,
					Columna: 0,
				})
			}
		case *Funcion:

			linea := 0
			columna := 0

			if funcion.Token != nil {
				linea = funcion.Token.GetLine()
				columna = funcion.Token.GetColumn()
			}

			ReporteAmbito.Funciones = append(ReporteAmbito.Funciones, ReporteSimbolos{
				Nombre:  funcion.Nombre,
				Tipo:    funcion.TipoRetorno,
				Linea:   linea,
				Columna: columna,
			})
		case *FuncionNativaObjeto:
			// No incluir FuncionNativaObjeto en el reporte
			break
		default:
			log.Fatal("Tipo de función no encontrado")
		}
	}

	for _, v := range s.Estructuras {
		ReporteAmbito.Estructuras = append(ReporteAmbito.Estructuras, ReporteSimbolos{
			Nombre:  v.Nombre,
			Tipo:    v.Nombre,
			Linea:   v.Token.GetLine(),
			Columna: v.Token.GetColumn(),
		})
	}

	for _, v := range s.Hijos {
		ReporteAmbito.AmbitosHijos = append(ReporteAmbito.AmbitosHijos, v.Report())
	}

	return ReporteAmbito
}
