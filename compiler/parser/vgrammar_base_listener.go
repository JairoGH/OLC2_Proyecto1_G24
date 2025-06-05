// Code generated from parser/vGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // vGrammar
import "github.com/antlr4-go/antlr/v4"

// BasevGrammarListener is a complete listener for a parse tree produced by vGrammar.
type BasevGrammarListener struct{}

var _ vGrammarListener = &BasevGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BasevGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BasevGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BasevGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BasevGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BasevGrammarListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BasevGrammarListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterFuncionPrincipal is called when production funcionPrincipal is entered.
func (s *BasevGrammarListener) EnterFuncionPrincipal(ctx *FuncionPrincipalContext) {}

// ExitFuncionPrincipal is called when production funcionPrincipal is exited.
func (s *BasevGrammarListener) ExitFuncionPrincipal(ctx *FuncionPrincipalContext) {}

// EnterBloque is called when production bloque is entered.
func (s *BasevGrammarListener) EnterBloque(ctx *BloqueContext) {}

// ExitBloque is called when production bloque is exited.
func (s *BasevGrammarListener) ExitBloque(ctx *BloqueContext) {}

// EnterSentencia is called when production sentencia is entered.
func (s *BasevGrammarListener) EnterSentencia(ctx *SentenciaContext) {}

// ExitSentencia is called when production sentencia is exited.
func (s *BasevGrammarListener) ExitSentencia(ctx *SentenciaContext) {}

// EnterDeclaracion is called when production declaracion is entered.
func (s *BasevGrammarListener) EnterDeclaracion(ctx *DeclaracionContext) {}

// ExitDeclaracion is called when production declaracion is exited.
func (s *BasevGrammarListener) ExitDeclaracion(ctx *DeclaracionContext) {}

// EnterAsignacion is called when production asignacion is entered.
func (s *BasevGrammarListener) EnterAsignacion(ctx *AsignacionContext) {}

// ExitAsignacion is called when production asignacion is exited.
func (s *BasevGrammarListener) ExitAsignacion(ctx *AsignacionContext) {}

// EnterIfStatement is called when production ifStatement is entered.
func (s *BasevGrammarListener) EnterIfStatement(ctx *IfStatementContext) {}

// ExitIfStatement is called when production ifStatement is exited.
func (s *BasevGrammarListener) ExitIfStatement(ctx *IfStatementContext) {}

// EnterExpresion is called when production expresion is entered.
func (s *BasevGrammarListener) EnterExpresion(ctx *ExpresionContext) {}

// ExitExpresion is called when production expresion is exited.
func (s *BasevGrammarListener) ExitExpresion(ctx *ExpresionContext) {}
