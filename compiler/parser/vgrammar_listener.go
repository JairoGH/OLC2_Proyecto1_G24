// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VGrammar
import "github.com/antlr4-go/antlr/v4"

// VGrammarListener is a complete listener for a parse tree produced by VGrammar.
type VGrammarListener interface {
	antlr.ParseTreeListener

	// EnterPrograma is called when entering the programa production.
	EnterPrograma(c *ProgramaContext)

	// EnterStmtDecl is called when entering the StmtDecl production.
	EnterStmtDecl(c *StmtDeclContext)

	// EnterStmtAssign is called when entering the StmtAssign production.
	EnterStmtAssign(c *StmtAssignContext)

	// EnterStmtFlow is called when entering the StmtFlow production.
	EnterStmtFlow(c *StmtFlowContext)

	// EnterStmtFuncCall is called when entering the StmtFuncCall production.
	EnterStmtFuncCall(c *StmtFuncCallContext)

	// EnterStmtIf is called when entering the StmtIf production.
	EnterStmtIf(c *StmtIfContext)

	// EnterStmtSwitch is called when entering the StmtSwitch production.
	EnterStmtSwitch(c *StmtSwitchContext)

	// EnterStmtFor is called when entering the StmtFor production.
	EnterStmtFor(c *StmtForContext)

	// EnterStmtFuncDecl is called when entering the StmtFuncDecl production.
	EnterStmtFuncDecl(c *StmtFuncDeclContext)

	// EnterStmtStructDecl is called when entering the StmtStructDecl production.
	EnterStmtStructDecl(c *StmtStructDeclContext)

	// EnterInferDecl is called when entering the InferDecl production.
	EnterInferDecl(c *InferDeclContext)

	// EnterTypedDeclInt is called when entering the TypedDeclInt production.
	EnterTypedDeclInt(c *TypedDeclIntContext)

	// EnterTypedDeclString is called when entering the TypedDeclString production.
	EnterTypedDeclString(c *TypedDeclStringContext)

	// EnterTypedDeclFloat is called when entering the TypedDeclFloat production.
	EnterTypedDeclFloat(c *TypedDeclFloatContext)

	// EnterTypedDeclBool is called when entering the TypedDeclBool production.
	EnterTypedDeclBool(c *TypedDeclBoolContext)

	// EnterTypedDeclRune is called when entering the TypedDeclRune production.
	EnterTypedDeclRune(c *TypedDeclRuneContext)

	// EnterAssignSimple is called when entering the AssignSimple production.
	EnterAssignSimple(c *AssignSimpleContext)

	// EnterAssignAdd is called when entering the AssignAdd production.
	EnterAssignAdd(c *AssignAddContext)

	// EnterAssignSub is called when entering the AssignSub production.
	EnterAssignSub(c *AssignSubContext)

	// EnterExprGreater is called when entering the ExprGreater production.
	EnterExprGreater(c *ExprGreaterContext)

	// EnterExprParen is called when entering the ExprParen production.
	EnterExprParen(c *ExprParenContext)

	// EnterExprNotEqual is called when entering the ExprNotEqual production.
	EnterExprNotEqual(c *ExprNotEqualContext)

	// EnterExprLessEq is called when entering the ExprLessEq production.
	EnterExprLessEq(c *ExprLessEqContext)

	// EnterExprLess is called when entering the ExprLess production.
	EnterExprLess(c *ExprLessContext)

	// EnterExprNot is called when entering the ExprNot production.
	EnterExprNot(c *ExprNotContext)

	// EnterExprDiv is called when entering the ExprDiv production.
	EnterExprDiv(c *ExprDivContext)

	// EnterExprAnd is called when entering the ExprAnd production.
	EnterExprAnd(c *ExprAndContext)

	// EnterExprGreaterEq is called when entering the ExprGreaterEq production.
	EnterExprGreaterEq(c *ExprGreaterEqContext)

	// EnterExprOr is called when entering the ExprOr production.
	EnterExprOr(c *ExprOrContext)

	// EnterExprSub is called when entering the ExprSub production.
	EnterExprSub(c *ExprSubContext)

	// EnterExprLiteral is called when entering the ExprLiteral production.
	EnterExprLiteral(c *ExprLiteralContext)

	// EnterExprMul is called when entering the ExprMul production.
	EnterExprMul(c *ExprMulContext)

	// EnterExprEqual is called when entering the ExprEqual production.
	EnterExprEqual(c *ExprEqualContext)

	// EnterExprId is called when entering the ExprId production.
	EnterExprId(c *ExprIdContext)

	// EnterExprAdd is called when entering the ExprAdd production.
	EnterExprAdd(c *ExprAddContext)

	// EnterExprMod is called when entering the ExprMod production.
	EnterExprMod(c *ExprModContext)

	// EnterLitInt is called when entering the LitInt production.
	EnterLitInt(c *LitIntContext)

	// EnterLitFloat is called when entering the LitFloat production.
	EnterLitFloat(c *LitFloatContext)

	// EnterLitString is called when entering the LitString production.
	EnterLitString(c *LitStringContext)

	// EnterLitRune is called when entering the LitRune production.
	EnterLitRune(c *LitRuneContext)

	// EnterLitTrue is called when entering the LitTrue production.
	EnterLitTrue(c *LitTrueContext)

	// EnterLitFalse is called when entering the LitFalse production.
	EnterLitFalse(c *LitFalseContext)

	// EnterLitNil is called when entering the LitNil production.
	EnterLitNil(c *LitNilContext)

	// EnterStmtBreak is called when entering the StmtBreak production.
	EnterStmtBreak(c *StmtBreakContext)

	// EnterStmtContinue is called when entering the StmtContinue production.
	EnterStmtContinue(c *StmtContinueContext)

	// EnterStmtReturn is called when entering the StmtReturn production.
	EnterStmtReturn(c *StmtReturnContext)

	// EnterFunctionCall is called when entering the FunctionCall production.
	EnterFunctionCall(c *FunctionCallContext)

	// EnterPrintCall is called when entering the PrintCall production.
	EnterPrintCall(c *PrintCallContext)

	// EnterPrintLnCall is called when entering the PrintLnCall production.
	EnterPrintLnCall(c *PrintLnCallContext)

	// EnterIf_stmt is called when entering the if_stmt production.
	EnterIf_stmt(c *If_stmtContext)

	// EnterSwitch_stmt is called when entering the switch_stmt production.
	EnterSwitch_stmt(c *Switch_stmtContext)

	// EnterSwitch_case is called when entering the switch_case production.
	EnterSwitch_case(c *Switch_caseContext)

	// EnterDefault_case is called when entering the default_case production.
	EnterDefault_case(c *Default_caseContext)

	// EnterFor_stmt is called when entering the for_stmt production.
	EnterFor_stmt(c *For_stmtContext)

	// EnterFunc_decl is called when entering the func_decl production.
	EnterFunc_decl(c *Func_declContext)

	// EnterParam is called when entering the param production.
	EnterParam(c *ParamContext)

	// EnterStruct_decl is called when entering the struct_decl production.
	EnterStruct_decl(c *Struct_declContext)

	// EnterStruct_field is called when entering the struct_field production.
	EnterStruct_field(c *Struct_fieldContext)

	// ExitPrograma is called when exiting the programa production.
	ExitPrograma(c *ProgramaContext)

	// ExitStmtDecl is called when exiting the StmtDecl production.
	ExitStmtDecl(c *StmtDeclContext)

	// ExitStmtAssign is called when exiting the StmtAssign production.
	ExitStmtAssign(c *StmtAssignContext)

	// ExitStmtFlow is called when exiting the StmtFlow production.
	ExitStmtFlow(c *StmtFlowContext)

	// ExitStmtFuncCall is called when exiting the StmtFuncCall production.
	ExitStmtFuncCall(c *StmtFuncCallContext)

	// ExitStmtIf is called when exiting the StmtIf production.
	ExitStmtIf(c *StmtIfContext)

	// ExitStmtSwitch is called when exiting the StmtSwitch production.
	ExitStmtSwitch(c *StmtSwitchContext)

	// ExitStmtFor is called when exiting the StmtFor production.
	ExitStmtFor(c *StmtForContext)

	// ExitStmtFuncDecl is called when exiting the StmtFuncDecl production.
	ExitStmtFuncDecl(c *StmtFuncDeclContext)

	// ExitStmtStructDecl is called when exiting the StmtStructDecl production.
	ExitStmtStructDecl(c *StmtStructDeclContext)

	// ExitInferDecl is called when exiting the InferDecl production.
	ExitInferDecl(c *InferDeclContext)

	// ExitTypedDeclInt is called when exiting the TypedDeclInt production.
	ExitTypedDeclInt(c *TypedDeclIntContext)

	// ExitTypedDeclString is called when exiting the TypedDeclString production.
	ExitTypedDeclString(c *TypedDeclStringContext)

	// ExitTypedDeclFloat is called when exiting the TypedDeclFloat production.
	ExitTypedDeclFloat(c *TypedDeclFloatContext)

	// ExitTypedDeclBool is called when exiting the TypedDeclBool production.
	ExitTypedDeclBool(c *TypedDeclBoolContext)

	// ExitTypedDeclRune is called when exiting the TypedDeclRune production.
	ExitTypedDeclRune(c *TypedDeclRuneContext)

	// ExitAssignSimple is called when exiting the AssignSimple production.
	ExitAssignSimple(c *AssignSimpleContext)

	// ExitAssignAdd is called when exiting the AssignAdd production.
	ExitAssignAdd(c *AssignAddContext)

	// ExitAssignSub is called when exiting the AssignSub production.
	ExitAssignSub(c *AssignSubContext)

	// ExitExprGreater is called when exiting the ExprGreater production.
	ExitExprGreater(c *ExprGreaterContext)

	// ExitExprParen is called when exiting the ExprParen production.
	ExitExprParen(c *ExprParenContext)

	// ExitExprNotEqual is called when exiting the ExprNotEqual production.
	ExitExprNotEqual(c *ExprNotEqualContext)

	// ExitExprLessEq is called when exiting the ExprLessEq production.
	ExitExprLessEq(c *ExprLessEqContext)

	// ExitExprLess is called when exiting the ExprLess production.
	ExitExprLess(c *ExprLessContext)

	// ExitExprNot is called when exiting the ExprNot production.
	ExitExprNot(c *ExprNotContext)

	// ExitExprDiv is called when exiting the ExprDiv production.
	ExitExprDiv(c *ExprDivContext)

	// ExitExprAnd is called when exiting the ExprAnd production.
	ExitExprAnd(c *ExprAndContext)

	// ExitExprGreaterEq is called when exiting the ExprGreaterEq production.
	ExitExprGreaterEq(c *ExprGreaterEqContext)

	// ExitExprOr is called when exiting the ExprOr production.
	ExitExprOr(c *ExprOrContext)

	// ExitExprSub is called when exiting the ExprSub production.
	ExitExprSub(c *ExprSubContext)

	// ExitExprLiteral is called when exiting the ExprLiteral production.
	ExitExprLiteral(c *ExprLiteralContext)

	// ExitExprMul is called when exiting the ExprMul production.
	ExitExprMul(c *ExprMulContext)

	// ExitExprEqual is called when exiting the ExprEqual production.
	ExitExprEqual(c *ExprEqualContext)

	// ExitExprId is called when exiting the ExprId production.
	ExitExprId(c *ExprIdContext)

	// ExitExprAdd is called when exiting the ExprAdd production.
	ExitExprAdd(c *ExprAddContext)

	// ExitExprMod is called when exiting the ExprMod production.
	ExitExprMod(c *ExprModContext)

	// ExitLitInt is called when exiting the LitInt production.
	ExitLitInt(c *LitIntContext)

	// ExitLitFloat is called when exiting the LitFloat production.
	ExitLitFloat(c *LitFloatContext)

	// ExitLitString is called when exiting the LitString production.
	ExitLitString(c *LitStringContext)

	// ExitLitRune is called when exiting the LitRune production.
	ExitLitRune(c *LitRuneContext)

	// ExitLitTrue is called when exiting the LitTrue production.
	ExitLitTrue(c *LitTrueContext)

	// ExitLitFalse is called when exiting the LitFalse production.
	ExitLitFalse(c *LitFalseContext)

	// ExitLitNil is called when exiting the LitNil production.
	ExitLitNil(c *LitNilContext)

	// ExitStmtBreak is called when exiting the StmtBreak production.
	ExitStmtBreak(c *StmtBreakContext)

	// ExitStmtContinue is called when exiting the StmtContinue production.
	ExitStmtContinue(c *StmtContinueContext)

	// ExitStmtReturn is called when exiting the StmtReturn production.
	ExitStmtReturn(c *StmtReturnContext)

	// ExitFunctionCall is called when exiting the FunctionCall production.
	ExitFunctionCall(c *FunctionCallContext)

	// ExitPrintCall is called when exiting the PrintCall production.
	ExitPrintCall(c *PrintCallContext)

	// ExitPrintLnCall is called when exiting the PrintLnCall production.
	ExitPrintLnCall(c *PrintLnCallContext)

	// ExitIf_stmt is called when exiting the if_stmt production.
	ExitIf_stmt(c *If_stmtContext)

	// ExitSwitch_stmt is called when exiting the switch_stmt production.
	ExitSwitch_stmt(c *Switch_stmtContext)

	// ExitSwitch_case is called when exiting the switch_case production.
	ExitSwitch_case(c *Switch_caseContext)

	// ExitDefault_case is called when exiting the default_case production.
	ExitDefault_case(c *Default_caseContext)

	// ExitFor_stmt is called when exiting the for_stmt production.
	ExitFor_stmt(c *For_stmtContext)

	// ExitFunc_decl is called when exiting the func_decl production.
	ExitFunc_decl(c *Func_declContext)

	// ExitParam is called when exiting the param production.
	ExitParam(c *ParamContext)

	// ExitStruct_decl is called when exiting the struct_decl production.
	ExitStruct_decl(c *Struct_declContext)

	// ExitStruct_field is called when exiting the struct_field production.
	ExitStruct_field(c *Struct_fieldContext)
}
