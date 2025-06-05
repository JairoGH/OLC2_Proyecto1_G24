// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseVGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVGrammarVisitor) VisitPrograma(ctx *ProgramaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtDecl(ctx *StmtDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtAssign(ctx *StmtAssignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtFlow(ctx *StmtFlowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtFuncCall(ctx *StmtFuncCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtIf(ctx *StmtIfContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtSwitch(ctx *StmtSwitchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtFor(ctx *StmtForContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtFuncDecl(ctx *StmtFuncDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtStructDecl(ctx *StmtStructDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitInferDecl(ctx *InferDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitTypedDeclInt(ctx *TypedDeclIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitTypedDeclString(ctx *TypedDeclStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitTypedDeclFloat(ctx *TypedDeclFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitTypedDeclBool(ctx *TypedDeclBoolContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitTypedDeclRune(ctx *TypedDeclRuneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAssignSimple(ctx *AssignSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAssignAdd(ctx *AssignAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAssignSub(ctx *AssignSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprGreater(ctx *ExprGreaterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprParen(ctx *ExprParenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprNotEqual(ctx *ExprNotEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprLessEq(ctx *ExprLessEqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprLess(ctx *ExprLessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprNot(ctx *ExprNotContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprDiv(ctx *ExprDivContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprAnd(ctx *ExprAndContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprGreaterEq(ctx *ExprGreaterEqContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprOr(ctx *ExprOrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprSub(ctx *ExprSubContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprLiteral(ctx *ExprLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprMul(ctx *ExprMulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprEqual(ctx *ExprEqualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprId(ctx *ExprIdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprAdd(ctx *ExprAddContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExprMod(ctx *ExprModContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLitInt(ctx *LitIntContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLitFloat(ctx *LitFloatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLitString(ctx *LitStringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLitRune(ctx *LitRuneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLitTrue(ctx *LitTrueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLitFalse(ctx *LitFalseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLitNil(ctx *LitNilContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtBreak(ctx *StmtBreakContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtContinue(ctx *StmtContinueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmtReturn(ctx *StmtReturnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitPrintCall(ctx *PrintCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitPrintLnCall(ctx *PrintLnCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitIf_stmt(ctx *If_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitSwitch_stmt(ctx *Switch_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitSwitch_case(ctx *Switch_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDefault_case(ctx *Default_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFor_stmt(ctx *For_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFunc_decl(ctx *Func_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitParam(ctx *ParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStruct_decl(ctx *Struct_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStruct_field(ctx *Struct_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}
