// Code generated from parser/vGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // vGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by vGrammar.
type vGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by vGrammar#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by vGrammar#funcionPrincipal.
	VisitFuncionPrincipal(ctx *FuncionPrincipalContext) interface{}

	// Visit a parse tree produced by vGrammar#bloque.
	VisitBloque(ctx *BloqueContext) interface{}

	// Visit a parse tree produced by vGrammar#sentencia.
	VisitSentencia(ctx *SentenciaContext) interface{}

	// Visit a parse tree produced by vGrammar#declaracion.
	VisitDeclaracion(ctx *DeclaracionContext) interface{}

	// Visit a parse tree produced by vGrammar#asignacion.
	VisitAsignacion(ctx *AsignacionContext) interface{}

	// Visit a parse tree produced by vGrammar#ifStatement.
	VisitIfStatement(ctx *IfStatementContext) interface{}

	// Visit a parse tree produced by vGrammar#expresion.
	VisitExpresion(ctx *ExpresionContext) interface{}
}
