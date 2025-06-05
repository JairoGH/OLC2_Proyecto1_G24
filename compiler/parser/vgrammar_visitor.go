// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VGrammar.
type VGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VGrammar#programa.
	VisitPrograma(ctx *ProgramaContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtDecl.
	VisitStmtDecl(ctx *StmtDeclContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtAssign.
	VisitStmtAssign(ctx *StmtAssignContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtFlow.
	VisitStmtFlow(ctx *StmtFlowContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtFuncCall.
	VisitStmtFuncCall(ctx *StmtFuncCallContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtIf.
	VisitStmtIf(ctx *StmtIfContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtSwitch.
	VisitStmtSwitch(ctx *StmtSwitchContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtFor.
	VisitStmtFor(ctx *StmtForContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtFuncDecl.
	VisitStmtFuncDecl(ctx *StmtFuncDeclContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtStructDecl.
	VisitStmtStructDecl(ctx *StmtStructDeclContext) interface{}

	// Visit a parse tree produced by VGrammar#InferDecl.
	VisitInferDecl(ctx *InferDeclContext) interface{}

	// Visit a parse tree produced by VGrammar#TypedDeclInt.
	VisitTypedDeclInt(ctx *TypedDeclIntContext) interface{}

	// Visit a parse tree produced by VGrammar#TypedDeclString.
	VisitTypedDeclString(ctx *TypedDeclStringContext) interface{}

	// Visit a parse tree produced by VGrammar#TypedDeclFloat.
	VisitTypedDeclFloat(ctx *TypedDeclFloatContext) interface{}

	// Visit a parse tree produced by VGrammar#TypedDeclBool.
	VisitTypedDeclBool(ctx *TypedDeclBoolContext) interface{}

	// Visit a parse tree produced by VGrammar#TypedDeclRune.
	VisitTypedDeclRune(ctx *TypedDeclRuneContext) interface{}

	// Visit a parse tree produced by VGrammar#AssignSimple.
	VisitAssignSimple(ctx *AssignSimpleContext) interface{}

	// Visit a parse tree produced by VGrammar#AssignAdd.
	VisitAssignAdd(ctx *AssignAddContext) interface{}

	// Visit a parse tree produced by VGrammar#AssignSub.
	VisitAssignSub(ctx *AssignSubContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprGreater.
	VisitExprGreater(ctx *ExprGreaterContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprParen.
	VisitExprParen(ctx *ExprParenContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprNotEqual.
	VisitExprNotEqual(ctx *ExprNotEqualContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprLessEq.
	VisitExprLessEq(ctx *ExprLessEqContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprLess.
	VisitExprLess(ctx *ExprLessContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprNot.
	VisitExprNot(ctx *ExprNotContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprDiv.
	VisitExprDiv(ctx *ExprDivContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprAnd.
	VisitExprAnd(ctx *ExprAndContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprGreaterEq.
	VisitExprGreaterEq(ctx *ExprGreaterEqContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprOr.
	VisitExprOr(ctx *ExprOrContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprSub.
	VisitExprSub(ctx *ExprSubContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprLiteral.
	VisitExprLiteral(ctx *ExprLiteralContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprMul.
	VisitExprMul(ctx *ExprMulContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprEqual.
	VisitExprEqual(ctx *ExprEqualContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprId.
	VisitExprId(ctx *ExprIdContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprAdd.
	VisitExprAdd(ctx *ExprAddContext) interface{}

	// Visit a parse tree produced by VGrammar#ExprMod.
	VisitExprMod(ctx *ExprModContext) interface{}

	// Visit a parse tree produced by VGrammar#LitInt.
	VisitLitInt(ctx *LitIntContext) interface{}

	// Visit a parse tree produced by VGrammar#LitFloat.
	VisitLitFloat(ctx *LitFloatContext) interface{}

	// Visit a parse tree produced by VGrammar#LitString.
	VisitLitString(ctx *LitStringContext) interface{}

	// Visit a parse tree produced by VGrammar#LitRune.
	VisitLitRune(ctx *LitRuneContext) interface{}

	// Visit a parse tree produced by VGrammar#LitTrue.
	VisitLitTrue(ctx *LitTrueContext) interface{}

	// Visit a parse tree produced by VGrammar#LitFalse.
	VisitLitFalse(ctx *LitFalseContext) interface{}

	// Visit a parse tree produced by VGrammar#LitNil.
	VisitLitNil(ctx *LitNilContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtBreak.
	VisitStmtBreak(ctx *StmtBreakContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtContinue.
	VisitStmtContinue(ctx *StmtContinueContext) interface{}

	// Visit a parse tree produced by VGrammar#StmtReturn.
	VisitStmtReturn(ctx *StmtReturnContext) interface{}

	// Visit a parse tree produced by VGrammar#FunctionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by VGrammar#PrintCall.
	VisitPrintCall(ctx *PrintCallContext) interface{}

	// Visit a parse tree produced by VGrammar#PrintLnCall.
	VisitPrintLnCall(ctx *PrintLnCallContext) interface{}

	// Visit a parse tree produced by VGrammar#if_stmt.
	VisitIf_stmt(ctx *If_stmtContext) interface{}

	// Visit a parse tree produced by VGrammar#switch_stmt.
	VisitSwitch_stmt(ctx *Switch_stmtContext) interface{}

	// Visit a parse tree produced by VGrammar#switch_case.
	VisitSwitch_case(ctx *Switch_caseContext) interface{}

	// Visit a parse tree produced by VGrammar#default_case.
	VisitDefault_case(ctx *Default_caseContext) interface{}

	// Visit a parse tree produced by VGrammar#for_stmt.
	VisitFor_stmt(ctx *For_stmtContext) interface{}

	// Visit a parse tree produced by VGrammar#func_decl.
	VisitFunc_decl(ctx *Func_declContext) interface{}

	// Visit a parse tree produced by VGrammar#param.
	VisitParam(ctx *ParamContext) interface{}

	// Visit a parse tree produced by VGrammar#struct_decl.
	VisitStruct_decl(ctx *Struct_declContext) interface{}

	// Visit a parse tree produced by VGrammar#struct_field.
	VisitStruct_field(ctx *Struct_fieldContext) interface{}
}
