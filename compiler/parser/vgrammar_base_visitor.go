// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseVGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVGrammarVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitMulDiv(ctx *MulDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAddSub(ctx *AddSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitParens(ctx *ParensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitInt(ctx *IntContext) interface{} {
	return v.VisitChildren(ctx)
}
