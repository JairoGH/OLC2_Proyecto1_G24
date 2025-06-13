package instrucciones

import (
	"log"
	"main/parser"
	"main/tiposDeDato"

	"github.com/antlr4-go/antlr/v4"
)

type VisitanteDcl struct {
	parser.BaseVGrammarVisitor
	RegistroAmbito *RegistroAmbito
	TablaError     *TablaError
	StructNames    []string
}

func NewVisitanteDcl(TablaError *TablaError) *VisitanteDcl {
	return &VisitanteDcl{
		RegistroAmbito: NuevoRegistroAmbito(),
		TablaError:     TablaError,
		StructNames:    []string{},
	}
}

func (v *VisitanteDcl) Visit(tree antlr.ParseTree) interface{} {

	switch val := tree.(type) {
	case *antlr.ErrorNodeImpl:
		log.Fatal(val.GetText())
		return nil
	default:
		return tree.Accept(v)
	}

}

func (v *VisitanteDcl) VisitProgram(ctx *parser.ProgramContext) interface{} {

	for _, stmt := range ctx.AllStmt() {
		v.Visit(stmt)
	}

	return nil
}

func (v *VisitanteDcl) VisitStmt(ctx *parser.StmtContext) interface{} {

	if ctx.Func_dcl() != nil {
		v.Visit(ctx.Func_dcl())
	} else if ctx.Strct_dcl() != nil {
		v.Visit(ctx.Strct_dcl())
	}

	return nil
}

func (v *VisitanteDcl) VisitFuncionDeclerada(ctx *parser.FuncionDecleradaContext) interface{} {

	if v.RegistroAmbito.AmbitoActual != v.RegistroAmbito.AmbitoGlobal {
		v.TablaError.NewErrorSemantico(ctx.GetStart(), "Las funciones solo pueden ser declaradas en el scope global")
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
		returnType = ctx.Type_().GetText()
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
	}

	return nil
}

func (v *VisitanteDcl) VisitParamList(ctx *parser.ParamListContext) interface{} {

	params := make([]*Parametros, 0)

	for _, param := range ctx.AllFunc_param() {
		params = append(params, v.Visit(param).(*Parametros))
	}

	return params
}

func (v *VisitanteDcl) VisitFuncParam(ctx *parser.FuncParamContext) interface{} {

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

	paramType := ctx.Type_().GetText()

	return &Parametros{
		ExternName:      externName,
		InnerName:       innerName,
		PassByReference: passByReference,
		Type:            paramType,
		Token:           ctx.GetStart(),
	}

}

func (v *VisitanteDcl) VisitDeclararStruct(ctx *parser.DeclararStructContext) interface{} {
	v.StructNames = append(v.StructNames, ctx.ID().GetText())
	return nil
}
