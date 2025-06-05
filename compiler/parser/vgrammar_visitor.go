// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VGrammar.
type VGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VGrammar#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by VGrammar#MulDiv.
	VisitMulDiv(ctx *MulDivContext) interface{}

	// Visit a parse tree produced by VGrammar#AddSub.
	VisitAddSub(ctx *AddSubContext) interface{}

	// Visit a parse tree produced by VGrammar#Parens.
	VisitParens(ctx *ParensContext) interface{}

	// Visit a parse tree produced by VGrammar#Int.
	VisitInt(ctx *IntContext) interface{}
}
