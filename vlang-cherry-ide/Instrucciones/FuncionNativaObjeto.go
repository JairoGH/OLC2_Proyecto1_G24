package instrucciones

import (
	"main/tiposDeDato"

	"github.com/antlr4-go/antlr/v4"
)

type FuncionNativaObjeto struct {
	*Funcion
	Object     *TipoObjeto
	CustomExec func(builtinRef *FuncionNativaObjeto, visitor *PatronVIsitor, args map[string]*Argumento, token antlr.Token)
}

// implementing ValorInterno
func (b FuncionNativaObjeto) Type() string {
	return tiposDeDato.TIPO_FUNCION
}

func (b FuncionNativaObjeto) Value() interface{} {
	return b
}

func (b FuncionNativaObjeto) Copy() tiposDeDato.ValorInterno {
	return b
}

func (f *FuncionNativaObjeto) Exec(visitor *PatronVIsitor, args []*Argumento, token antlr.Token) {

	context := visitor.GetReplContext()

	// validate args
	argsOk, argsMap := f.ValidarArgumentos(context, args, token)

	if !argsOk {
		f.RetornarValor = tiposDeDato.NuloPorDefecto
		return
	}

	f.CustomExec(f, visitor, argsMap, token)

}

// * Vector Built In Functions

// 1. Append
// vector.append(value)

// Parameters:

var appendParams = []*Parametros{
	// Just one positional Argumento
	{
		ExternName:      "_",
		InnerName:       "_",
		Type:            tiposDeDato.TIPO_ANY,
		PassByReference: false,
		Token:           nil,
	},
}

func EjecucionPersonalizada(builtinRef *FuncionNativaObjeto, visitor *PatronVIsitor, args map[string]*Argumento, token antlr.Token) {

	builtinRef.RetornarValor = tiposDeDato.NuloPorDefecto

	// get the vector
	vector := builtinRef.Object.AuxObject.(*TipoVector)

	// get the value
	arg := args["_"]

	if vector.ItemType != arg.Value.Type() {
		visitor.TablaError.NewErrorSemantico(arg.Token, "No se puede agregar un valor de tipo "+arg.Value.Type()+" a un vector de tipo "+vector.ItemType)
		return
	}
	vector.InternalValor = append(vector.InternalValor, arg.Value)
	vector.updateProps()
}

// 2. vector.remove(at: Int) -> nil

// parameters:

var removeParams = []*Parametros{
	// at: Int
	{
		ExternName:      "at",
		InnerName:       "at",
		Type:            tiposDeDato.TIPO_ENTERO,
		PassByReference: false,
		Token:           nil,
	},
}

func RemoverEjecucion(builtinRef *FuncionNativaObjeto, visitor *PatronVIsitor, args map[string]*Argumento, token antlr.Token) {

	builtinRef.RetornarValor = tiposDeDato.NuloPorDefecto

	// get the vector
	vector := builtinRef.Object.AuxObject.(*TipoVector)

	// get the value
	arg := args["at"]

	if arg.Value.Type() != tiposDeDato.TIPO_ENTERO {
		visitor.TablaError.NewErrorSemantico(arg.Token, "El Argumento 'at' debe ser de tipo Int")
		return
	}

	// out of bounds
	if arg.Value.Value().(int) >= vector.Size() || arg.Value.Value().(int) < 0 {
		visitor.TablaError.NewErrorSemantico(arg.Token, "El indice esta fuera de rango")
		return
	}

	// remove the element
	vector.InternalValor = append(vector.InternalValor[:arg.Value.Value().(int)], vector.InternalValor[arg.Value.Value().(int)+1:]...)
	vector.updateProps()
}

// 3. removeLast
// vector.removeLast() -> nil

// parameters:

var removeLastParams = []*Parametros{}

func EliminarUltimaEjecucion(builtinRef *FuncionNativaObjeto, visitor *PatronVIsitor, args map[string]*Argumento, token antlr.Token) {

	builtinRef.RetornarValor = tiposDeDato.NuloPorDefecto

	// get the vector
	vector := builtinRef.Object.AuxObject.(*TipoVector)

	if vector.Size() == 0 {
		visitor.TablaError.NewErrorSemantico(token, "El vector esta vacio y no se puede remover el ultimo elemento")
		return
	}

	// remove the last element
	vector.InternalValor = vector.InternalValor[:vector.Size()-1]
	vector.updateProps()
}

func AgregarVectorNativo(vectorRef *TipoVector) {

	vectorScope := NuevoVectorGlobal()

	vectorInternalObject := &TipoObjeto{
		InternalScope: vectorScope,
		AuxObject:     vectorRef,
	}

	// Register built in functions
	vectorScope.AgregarFuncion("append", &FuncionNativaObjeto{
		Funcion: &Funcion{
			Parametros: appendParams,
		},
		Object:     vectorInternalObject,
		CustomExec: EjecucionPersonalizada,
	})

	vectorScope.AgregarFuncion("remove", &FuncionNativaObjeto{
		Funcion: &Funcion{
			Parametros: removeParams,
		},
		Object:     vectorInternalObject,
		CustomExec: RemoverEjecucion,
	})

	vectorScope.AgregarFuncion("removeLast", &FuncionNativaObjeto{
		Funcion: &Funcion{
			Parametros: removeLastParams,
		},
		Object:     vectorInternalObject,
		CustomExec: EliminarUltimaEjecucion,
	})

	// make isEmpty a property
	vectorScope.AgregarVariable("isEmpty", tiposDeDato.TIPO_BOOLEAN, vectorRef.IsEmpty, true, false, nil)

	// make count a property
	vectorScope.AgregarVariable("count", tiposDeDato.TIPO_ENTERO, vectorRef.SizeValue, true, false, nil)

	vectorRef.TipoObjeto = vectorInternalObject
}
