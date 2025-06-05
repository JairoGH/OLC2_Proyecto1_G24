// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VGrammar
import "github.com/antlr4-go/antlr/v4"

// BaseVGrammarListener is a complete listener for a parse tree produced by VGrammar.
type BaseVGrammarListener struct{}

var _ VGrammarListener = &BaseVGrammarListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseVGrammarListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseVGrammarListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseVGrammarListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseVGrammarListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterPrograma is called when production programa is entered.
func (s *BaseVGrammarListener) EnterPrograma(ctx *ProgramaContext) {}

// ExitPrograma is called when production programa is exited.
func (s *BaseVGrammarListener) ExitPrograma(ctx *ProgramaContext) {}

// EnterMulDiv is called when production MulDiv is entered.
func (s *BaseVGrammarListener) EnterMulDiv(ctx *MulDivContext) {}

// ExitMulDiv is called when production MulDiv is exited.
func (s *BaseVGrammarListener) ExitMulDiv(ctx *MulDivContext) {}

// EnterAddSub is called when production AddSub is entered.
func (s *BaseVGrammarListener) EnterAddSub(ctx *AddSubContext) {}

// ExitAddSub is called when production AddSub is exited.
func (s *BaseVGrammarListener) ExitAddSub(ctx *AddSubContext) {}

// EnterParens is called when production Parens is entered.
func (s *BaseVGrammarListener) EnterParens(ctx *ParensContext) {}

// ExitParens is called when production Parens is exited.
func (s *BaseVGrammarListener) ExitParens(ctx *ParensContext) {}

// EnterInt is called when production Int is entered.
func (s *BaseVGrammarListener) EnterInt(ctx *IntContext) {}

// ExitInt is called when production Int is exited.
func (s *BaseVGrammarListener) ExitInt(ctx *IntContext) {}
