// Code generated from parser/vGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // vGrammar
import "github.com/antlr4-go/antlr/v4"

// vGrammarListener is a complete listener for a parse tree produced by vGrammar.
type vGrammarListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterFuncionPrincipal is called when entering the funcionPrincipal production.
	EnterFuncionPrincipal(c *FuncionPrincipalContext)

	// EnterBloque is called when entering the bloque production.
	EnterBloque(c *BloqueContext)

	// EnterSentencia is called when entering the sentencia production.
	EnterSentencia(c *SentenciaContext)

	// EnterDeclaracion is called when entering the declaracion production.
	EnterDeclaracion(c *DeclaracionContext)

	// EnterAsignacion is called when entering the asignacion production.
	EnterAsignacion(c *AsignacionContext)

	// EnterIfStatement is called when entering the ifStatement production.
	EnterIfStatement(c *IfStatementContext)

	// EnterExpresion is called when entering the expresion production.
	EnterExpresion(c *ExpresionContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitFuncionPrincipal is called when exiting the funcionPrincipal production.
	ExitFuncionPrincipal(c *FuncionPrincipalContext)

	// ExitBloque is called when exiting the bloque production.
	ExitBloque(c *BloqueContext)

	// ExitSentencia is called when exiting the sentencia production.
	ExitSentencia(c *SentenciaContext)

	// ExitDeclaracion is called when exiting the declaracion production.
	ExitDeclaracion(c *DeclaracionContext)

	// ExitAsignacion is called when exiting the asignacion production.
	ExitAsignacion(c *AsignacionContext)

	// ExitIfStatement is called when exiting the ifStatement production.
	ExitIfStatement(c *IfStatementContext)

	// ExitExpresion is called when exiting the expresion production.
	ExitExpresion(c *ExpresionContext)
}
