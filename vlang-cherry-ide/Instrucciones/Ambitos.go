package instrucciones

import (
	"fmt"
	"log"
	"main/tiposDeDato"
	"strings"

	"github.com/antlr4-go/antlr/v4"
)

type AmbitoBase struct {
    name      string
    variables map[string]*Variable
    functions map[string]tiposDeDato.ValorInterno
    methods   map[string]*MetodoStruct 
    children  []*AmbitoBase
    structs   map[string]*Struct
    parent    *AmbitoBase
    isStruct  bool
    IsMutating bool
}

func (s *AmbitoBase) Name() string {
	return s.name
}

func (s *AmbitoBase) Parent() *AmbitoBase {
	return s.parent
}

func (s *AmbitoBase) Children() []*AmbitoBase {
	return s.children
}

func (s *AmbitoBase) ValidType(_type string) bool {

	_, isStructType := s.structs[_type]

	return tiposDeDato.PRIMITIVO(_type) || isStructType
}

func (s *AmbitoBase) AddChild(child *AmbitoBase) {
	s.children = append(s.children, child)
	child.parent = s
}

func (s *AmbitoBase) existeVariable(variable *Variable) bool {

	if _, ok := s.variables[variable.Name]; ok {
		return true
	}

	return false

}

// Crear función NewAmbitoBase:
func NewAmbitoBase(name string, isStruct bool) *AmbitoBase {
    return &AmbitoBase{
        name:       name,
        variables:  make(map[string]*Variable),
        functions:  make(map[string]tiposDeDato.ValorInterno),
        methods:    make(map[string]*MetodoStruct),  // NUEVO
        children:   make([]*AmbitoBase, 0),
        structs:    make(map[string]*Struct),
        parent:     nil,
        isStruct:   isStruct,
        IsMutating: false,
    }
}

// Agregar método para obtener métodos:
func (a *AmbitoBase) GetMetodo(receiverType, methodName string) (*MetodoStruct, string) {
    methodKey := receiverType + "." + methodName
    if method, exists := a.methods[methodKey]; exists {
        return method, ""
    }
    return nil, fmt.Sprintf("Método %s no encontrado para tipo %s", methodName, receiverType)
}

// Agregar método para agregar métodos:
func (a *AmbitoBase) AgregarMetodo(receiverType, methodName string, method *MetodoStruct) (bool, string) {
    methodKey := receiverType + "." + methodName
    if _, exists := a.methods[methodKey]; exists {
        return false, fmt.Sprintf("Método %s ya existe para tipo %s", methodName, receiverType)
    }
    a.methods[methodKey] = method
    return true, ""
}

func (s *AmbitoBase) AgregarVariable(name string, varType string, value tiposDeDato.ValorInterno, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {

	variable := &Variable{
		Name:     name,
		Value:    value,
		Type:     varType,
		IsConst:  isConst,
		AllowNil: allowNil,
		Token:    token,
	}

	if s.existeVariable(variable) {
		return nil, "La variable " + name + " ya existe"
	}

	typesOk, msg := variable.TipoValidacion()

	// even if the variable is not valid, we add it to the scope, (internally it will be nil)
	s.variables[name] = variable

	if !typesOk {
		// report error
		return nil, msg
	}

	return variable, ""
}

func (s *AmbitoBase) GetVariable(name string) *Variable {
	// verify if is refering to and object/struct function
	if strings.Contains(name, ".") {
		return s.searchObjectVariable(name, nil)
	}

	initialScope := s

	for {
		if variable, ok := initialScope.variables[name]; ok {

			// verify if is refering to a pointer
			if variable.Type == tiposDeDato.TIPO_PUNTERO {
				return variable.Value.(*TipoPuntero).VariableAsociada // pointer of a pointer ?
			}

			return variable
		}

		if initialScope.parent == nil {
			break
		}

		initialScope = initialScope.parent
	}

	return nil
}

// obj1.obj2.prop1

func (s *AmbitoBase) searchObjectVariable(name string, lastObj tiposDeDato.ValorInterno) *Variable {

	// split name by dot
	parts := strings.Split(name, ".")

	if len(parts) == 0 {
		log.Fatal("idk what u did, cant split by dot")
		return nil
	}

	if len(parts) == 1 {
		obj, ok := lastObj.(*TipoObjeto)

		if ok {
			return obj.InternalScope.GetVariable(name)
		}

		log.Fatal("idk what u did, cant convert to object")
		return nil
	}

	// then parts should be 2 or more

	if lastObj == nil {
		variable := s.GetVariable(parts[0])

		if variable == nil {
			return nil
		}

		obj := variable.Value

		// obj must be an object/struct or vector
		switch obj := obj.(type) {
		case *TipoObjeto:
			lastObj = obj
		case *TipoVector:
			lastObj = obj.TipoObjeto
		default:
			return nil
		}

		return s.searchObjectVariable(strings.Join(parts[1:], "."), lastObj)
	}

	obj, ok := lastObj.(*TipoObjeto)

	if ok {
		lastObj = obj.InternalScope.GetVariable(parts[0]).Value

		return s.searchObjectVariable(strings.Join(parts[1:], "."), lastObj)
	} else {
		log.Fatal("idk what u did, cant convert to object")
		return nil
	}
}

func (s *AmbitoBase) AgregarFuncion(name string, function tiposDeDato.ValorInterno) (bool, string) {
	// check if function already exists

	if _, ok := s.functions[name]; ok {
		return false, "La funcion " + name + " ya existe"
	}

	s.functions[name] = function

	return true, ""
}

func (s *AmbitoBase) GetFuncion(name string) (tiposDeDato.ValorInterno, string) {

	// verify if is refering to and object/struct function
	if strings.Contains(name, ".") {
		return s.buscarFuncion(name, nil)
	}

	initialScope := s

	for {
		if function, ok := initialScope.functions[name]; ok {
			return function, ""
		}

		if initialScope.parent == nil {
			break
		}

		initialScope = initialScope.parent
	}

	return nil, "La funcion " + name + " no existe"
}

// obj1.obj2.func1()

func (s *AmbitoBase) buscarFuncion(name string, lastObj tiposDeDato.ValorInterno) (tiposDeDato.ValorInterno, string) {

	// split name by dot
	parts := strings.Split(name, ".")

	if len(parts) == 0 {
		log.Fatal("idk what u did, cant split by dot")
		return nil, ""
	}

	if len(parts) == 1 {
		obj, ok := lastObj.(*TipoObjeto)

		if ok {
			return obj.InternalScope.GetFuncion(name)
		}

		log.Fatal("idk what u did, cant convert to object")
		return nil, ""
	}

	// then parts should be 2 or more

	if lastObj == nil {
		variable := s.GetVariable(parts[0])

		if variable == nil {
			return nil, "No se puede acceder a la propiedad " + parts[0]
		}

		obj := variable.Value

		// obj must be an object/struct or vector

		switch obj := obj.(type) {
		case *TipoObjeto:
			lastObj = obj
		case *TipoVector:
			lastObj = obj.TipoObjeto
		default:
			return nil, "La propiedad '" + variable.Name + "' de tipo " + obj.Type() + " no tiene propiedades"
		}

		return s.buscarFuncion(strings.Join(parts[1:], "."), lastObj)
	}

	obj, ok := lastObj.(*TipoObjeto)

	if ok {
		lastObj = obj.InternalScope.GetVariable(parts[0]).Value

		return s.buscarFuncion(strings.Join(parts[1:], "."), lastObj)
	} else {
		log.Fatal("idk what u did, cant convert to object")
		return nil, ""
	}
}

func (s *AmbitoBase) AgregarEstructura(name string, structValue *Struct) (bool, string) {

	if _, ok := s.structs[name]; ok {
		return false, "La estructura " + name + " ya existe"
	}

	s.structs[name] = structValue
	return true, ""
}

func (s *AmbitoBase) GetEstructura(name string) (*Struct, string) {

	initialScope := s

	for {
		if structValue, ok := initialScope.structs[name]; ok {
			return structValue, ""
		}

		if initialScope.parent == nil {
			break
		}

		initialScope = initialScope.parent
	}

	return nil, "La estructura " + name + " no existe"
}

func (s *AmbitoBase) Reset() {
	s.variables = make(map[string]*Variable)
	s.children = make([]*AmbitoBase, 0)
	s.functions = make(map[string]tiposDeDato.ValorInterno)
}

func (s *AmbitoBase) EsAmbitoMutante() bool {
	aux := s

	for {
		if aux.IsMutating {
			return true
		}

		if aux.parent == nil {
			break
		}

		aux = aux.parent
	}

	return false
}

func NuevoAmbitoGlobal() *AmbitoBase {
    funcs := make(map[string]tiposDeDato.ValorInterno)

    for k, v := range DefaultBuiltInFunctions {
        funcs[k] = v
    }

    return &AmbitoBase{
        name:      "global",
        variables: make(map[string]*Variable),
        children:  make([]*AmbitoBase, 0),
        structs:   make(map[string]*Struct),
        parent:    nil,
        functions: funcs,
        methods:   make(map[string]*MetodoStruct), 
    }
}

func NewLocalScope(name string) *AmbitoBase {
    return &AmbitoBase{
        name:      name,
        variables: make(map[string]*Variable),
        functions: make(map[string]tiposDeDato.ValorInterno),
        methods:   make(map[string]*MetodoStruct), 
        children:  make([]*AmbitoBase, 0),
        parent:    nil,
    }
}

type RegistroAmbito struct {
	AmbitoGlobal *AmbitoBase
	AmbitoActual *AmbitoBase
}

func (s *RegistroAmbito) PushScope(name string) *AmbitoBase {

	newScope := NewLocalScope(name)
	s.AmbitoActual.AddChild(newScope)
	s.AmbitoActual = newScope

	return s.AmbitoActual
}

func (s *RegistroAmbito) PopScope() {
	s.AmbitoActual = s.AmbitoActual.Parent()
}

func (s *RegistroAmbito) Reset() {
	s.AmbitoActual = s.AmbitoGlobal
}

func (s *RegistroAmbito) AgregarVariable(name string, varType string, value tiposDeDato.ValorInterno, isConst bool, allowNil bool, token antlr.Token) (*Variable, string) {
	return s.AmbitoActual.AgregarVariable(name, varType, value, isConst, allowNil, token)
}

func (s *RegistroAmbito) GetVariable(name string) *Variable {
	return s.AmbitoActual.GetVariable(name)
}

func (s *RegistroAmbito) AgregarFuncion(name string, function tiposDeDato.ValorInterno) (bool, string) {
	return s.AmbitoActual.AgregarFuncion(name, function)
}

func (s *RegistroAmbito) GetFuncion(name string) (tiposDeDato.ValorInterno, string) {
	return s.AmbitoActual.GetFuncion(name)
}

func (s *RegistroAmbito) EsEntornoMutable() bool {
	return s.AmbitoActual.EsAmbitoMutante()
}

func NuevoRegistroAmbito() *RegistroAmbito {
	AmbitoGlobal := NuevoAmbitoGlobal()
	return &RegistroAmbito{
		AmbitoGlobal: AmbitoGlobal,
		AmbitoActual: AmbitoGlobal,
	}
}

func NuevoVectorGlobal() *AmbitoBase {
	var scope = &AmbitoBase{
		name:      "vector",
		variables: make(map[string]*Variable),
		children:  make([]*AmbitoBase, 0),
		functions: make(map[string]tiposDeDato.ValorInterno),
		parent:    nil,
	}

	// register object built-in functions

	return scope
}

// * Report

type ReporteTabla struct {
	AmbitoGlobal ReporteAmbito
}

type ReporteAmbito struct {
	Name        string
	Vars        []ReporteSimbolos
	Funcs       []ReporteSimbolos
	Structs     []ReporteSimbolos
	ChildScopes []ReporteAmbito
}

type ReporteSimbolos struct {
	Name   string
	Type   string
	Line   int
	Column int
}

func (s *RegistroAmbito) Report() ReporteTabla {
	return ReporteTabla{
		AmbitoGlobal: s.AmbitoActual.Report(),
	}
}

func (s *AmbitoBase) Report() ReporteAmbito {

	ReporteAmbito := ReporteAmbito{
		Name:        s.name,
		Vars:        make([]ReporteSimbolos, 0),
		Funcs:       make([]ReporteSimbolos, 0),
		Structs:     make([]ReporteSimbolos, 0),
		ChildScopes: make([]ReporteAmbito, 0),
	}

	for _, v := range s.variables {

		token := v.Token
		line := 0
		column := 0

		if token != nil {
			line = token.GetLine()
			column = token.GetColumn()
		}

		ReporteAmbito.Vars = append(ReporteAmbito.Vars, ReporteSimbolos{
			Name:   v.Name,
			Type:   v.Type,
			Line:   line,
			Column: column,
		})
	}

	for _, f := range s.functions {
		switch function := f.(type) {
		case *FuncionNativa:
			ReporteAmbito.Funcs = append(ReporteAmbito.Funcs, ReporteSimbolos{
				Name:   function.Name,
				Type:   "Embebida: " + function.Name,
				Line:   0,
				Column: 0,
			})
		case *Funcion:

			line := 0
			column := 0

			if function.Token != nil {
				line = function.Token.GetLine()
				column = function.Token.GetColumn()
			}

			ReporteAmbito.Funcs = append(ReporteAmbito.Funcs, ReporteSimbolos{
				Name:   function.Name,
				Type:   function.ReturnType,
				Line:   line,
				Column: column,
			})
		case *FuncionNativaObjeto:
			break
		default:
			log.Fatal("Function type not found")
		}
	}

	for _, v := range s.structs {
		ReporteAmbito.Structs = append(ReporteAmbito.Structs, ReporteSimbolos{
			Name:   v.Name,
			Type:   v.Name,
			Line:   v.Token.GetLine(),
			Column: v.Token.GetColumn(),
		})
	}

	for _, v := range s.children {
		ReporteAmbito.ChildScopes = append(ReporteAmbito.ChildScopes, v.Report())
	}

	return ReporteAmbito
}
