// Code generated from parser/vGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // vGrammar
import "github.com/antlr4-go/antlr/v4"

type BasevGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BasevGrammarVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevGrammarVisitor) VisitFuncionPrincipal(ctx *FuncionPrincipalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevGrammarVisitor) VisitBloque(ctx *BloqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevGrammarVisitor) VisitSentencia(ctx *SentenciaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevGrammarVisitor) VisitDeclaracion(ctx *DeclaracionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevGrammarVisitor) VisitAsignacion(ctx *AsignacionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevGrammarVisitor) VisitIfStatement(ctx *IfStatementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BasevGrammarVisitor) VisitExpresion(ctx *ExpresionContext) interface{} {
	return v.VisitChildren(ctx)
}
