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

// EnterStmtDecl is called when production StmtDecl is entered.
func (s *BaseVGrammarListener) EnterStmtDecl(ctx *StmtDeclContext) {}

// ExitStmtDecl is called when production StmtDecl is exited.
func (s *BaseVGrammarListener) ExitStmtDecl(ctx *StmtDeclContext) {}

// EnterStmtAssign is called when production StmtAssign is entered.
func (s *BaseVGrammarListener) EnterStmtAssign(ctx *StmtAssignContext) {}

// ExitStmtAssign is called when production StmtAssign is exited.
func (s *BaseVGrammarListener) ExitStmtAssign(ctx *StmtAssignContext) {}

// EnterStmtFlow is called when production StmtFlow is entered.
func (s *BaseVGrammarListener) EnterStmtFlow(ctx *StmtFlowContext) {}

// ExitStmtFlow is called when production StmtFlow is exited.
func (s *BaseVGrammarListener) ExitStmtFlow(ctx *StmtFlowContext) {}

// EnterStmtFuncCall is called when production StmtFuncCall is entered.
func (s *BaseVGrammarListener) EnterStmtFuncCall(ctx *StmtFuncCallContext) {}

// ExitStmtFuncCall is called when production StmtFuncCall is exited.
func (s *BaseVGrammarListener) ExitStmtFuncCall(ctx *StmtFuncCallContext) {}

// EnterStmtIf is called when production StmtIf is entered.
func (s *BaseVGrammarListener) EnterStmtIf(ctx *StmtIfContext) {}

// ExitStmtIf is called when production StmtIf is exited.
func (s *BaseVGrammarListener) ExitStmtIf(ctx *StmtIfContext) {}

// EnterStmtSwitch is called when production StmtSwitch is entered.
func (s *BaseVGrammarListener) EnterStmtSwitch(ctx *StmtSwitchContext) {}

// ExitStmtSwitch is called when production StmtSwitch is exited.
func (s *BaseVGrammarListener) ExitStmtSwitch(ctx *StmtSwitchContext) {}

// EnterStmtFor is called when production StmtFor is entered.
func (s *BaseVGrammarListener) EnterStmtFor(ctx *StmtForContext) {}

// ExitStmtFor is called when production StmtFor is exited.
func (s *BaseVGrammarListener) ExitStmtFor(ctx *StmtForContext) {}

// EnterStmtFuncDecl is called when production StmtFuncDecl is entered.
func (s *BaseVGrammarListener) EnterStmtFuncDecl(ctx *StmtFuncDeclContext) {}

// ExitStmtFuncDecl is called when production StmtFuncDecl is exited.
func (s *BaseVGrammarListener) ExitStmtFuncDecl(ctx *StmtFuncDeclContext) {}

// EnterStmtStructDecl is called when production StmtStructDecl is entered.
func (s *BaseVGrammarListener) EnterStmtStructDecl(ctx *StmtStructDeclContext) {}

// ExitStmtStructDecl is called when production StmtStructDecl is exited.
func (s *BaseVGrammarListener) ExitStmtStructDecl(ctx *StmtStructDeclContext) {}

// EnterInferDecl is called when production InferDecl is entered.
func (s *BaseVGrammarListener) EnterInferDecl(ctx *InferDeclContext) {}

// ExitInferDecl is called when production InferDecl is exited.
func (s *BaseVGrammarListener) ExitInferDecl(ctx *InferDeclContext) {}

// EnterTypedDeclInt is called when production TypedDeclInt is entered.
func (s *BaseVGrammarListener) EnterTypedDeclInt(ctx *TypedDeclIntContext) {}

// ExitTypedDeclInt is called when production TypedDeclInt is exited.
func (s *BaseVGrammarListener) ExitTypedDeclInt(ctx *TypedDeclIntContext) {}

// EnterTypedDeclString is called when production TypedDeclString is entered.
func (s *BaseVGrammarListener) EnterTypedDeclString(ctx *TypedDeclStringContext) {}

// ExitTypedDeclString is called when production TypedDeclString is exited.
func (s *BaseVGrammarListener) ExitTypedDeclString(ctx *TypedDeclStringContext) {}

// EnterTypedDeclFloat is called when production TypedDeclFloat is entered.
func (s *BaseVGrammarListener) EnterTypedDeclFloat(ctx *TypedDeclFloatContext) {}

// ExitTypedDeclFloat is called when production TypedDeclFloat is exited.
func (s *BaseVGrammarListener) ExitTypedDeclFloat(ctx *TypedDeclFloatContext) {}

// EnterTypedDeclBool is called when production TypedDeclBool is entered.
func (s *BaseVGrammarListener) EnterTypedDeclBool(ctx *TypedDeclBoolContext) {}

// ExitTypedDeclBool is called when production TypedDeclBool is exited.
func (s *BaseVGrammarListener) ExitTypedDeclBool(ctx *TypedDeclBoolContext) {}

// EnterTypedDeclRune is called when production TypedDeclRune is entered.
func (s *BaseVGrammarListener) EnterTypedDeclRune(ctx *TypedDeclRuneContext) {}

// ExitTypedDeclRune is called when production TypedDeclRune is exited.
func (s *BaseVGrammarListener) ExitTypedDeclRune(ctx *TypedDeclRuneContext) {}

// EnterAssignSimple is called when production AssignSimple is entered.
func (s *BaseVGrammarListener) EnterAssignSimple(ctx *AssignSimpleContext) {}

// ExitAssignSimple is called when production AssignSimple is exited.
func (s *BaseVGrammarListener) ExitAssignSimple(ctx *AssignSimpleContext) {}

// EnterAssignAdd is called when production AssignAdd is entered.
func (s *BaseVGrammarListener) EnterAssignAdd(ctx *AssignAddContext) {}

// ExitAssignAdd is called when production AssignAdd is exited.
func (s *BaseVGrammarListener) ExitAssignAdd(ctx *AssignAddContext) {}

// EnterAssignSub is called when production AssignSub is entered.
func (s *BaseVGrammarListener) EnterAssignSub(ctx *AssignSubContext) {}

// ExitAssignSub is called when production AssignSub is exited.
func (s *BaseVGrammarListener) ExitAssignSub(ctx *AssignSubContext) {}

// EnterExprGreater is called when production ExprGreater is entered.
func (s *BaseVGrammarListener) EnterExprGreater(ctx *ExprGreaterContext) {}

// ExitExprGreater is called when production ExprGreater is exited.
func (s *BaseVGrammarListener) ExitExprGreater(ctx *ExprGreaterContext) {}

// EnterExprParen is called when production ExprParen is entered.
func (s *BaseVGrammarListener) EnterExprParen(ctx *ExprParenContext) {}

// ExitExprParen is called when production ExprParen is exited.
func (s *BaseVGrammarListener) ExitExprParen(ctx *ExprParenContext) {}

// EnterExprNotEqual is called when production ExprNotEqual is entered.
func (s *BaseVGrammarListener) EnterExprNotEqual(ctx *ExprNotEqualContext) {}

// ExitExprNotEqual is called when production ExprNotEqual is exited.
func (s *BaseVGrammarListener) ExitExprNotEqual(ctx *ExprNotEqualContext) {}

// EnterExprLessEq is called when production ExprLessEq is entered.
func (s *BaseVGrammarListener) EnterExprLessEq(ctx *ExprLessEqContext) {}

// ExitExprLessEq is called when production ExprLessEq is exited.
func (s *BaseVGrammarListener) ExitExprLessEq(ctx *ExprLessEqContext) {}

// EnterExprLess is called when production ExprLess is entered.
func (s *BaseVGrammarListener) EnterExprLess(ctx *ExprLessContext) {}

// ExitExprLess is called when production ExprLess is exited.
func (s *BaseVGrammarListener) ExitExprLess(ctx *ExprLessContext) {}

// EnterExprNot is called when production ExprNot is entered.
func (s *BaseVGrammarListener) EnterExprNot(ctx *ExprNotContext) {}

// ExitExprNot is called when production ExprNot is exited.
func (s *BaseVGrammarListener) ExitExprNot(ctx *ExprNotContext) {}

// EnterExprDiv is called when production ExprDiv is entered.
func (s *BaseVGrammarListener) EnterExprDiv(ctx *ExprDivContext) {}

// ExitExprDiv is called when production ExprDiv is exited.
func (s *BaseVGrammarListener) ExitExprDiv(ctx *ExprDivContext) {}

// EnterExprAnd is called when production ExprAnd is entered.
func (s *BaseVGrammarListener) EnterExprAnd(ctx *ExprAndContext) {}

// ExitExprAnd is called when production ExprAnd is exited.
func (s *BaseVGrammarListener) ExitExprAnd(ctx *ExprAndContext) {}

// EnterExprGreaterEq is called when production ExprGreaterEq is entered.
func (s *BaseVGrammarListener) EnterExprGreaterEq(ctx *ExprGreaterEqContext) {}

// ExitExprGreaterEq is called when production ExprGreaterEq is exited.
func (s *BaseVGrammarListener) ExitExprGreaterEq(ctx *ExprGreaterEqContext) {}

// EnterExprOr is called when production ExprOr is entered.
func (s *BaseVGrammarListener) EnterExprOr(ctx *ExprOrContext) {}

// ExitExprOr is called when production ExprOr is exited.
func (s *BaseVGrammarListener) ExitExprOr(ctx *ExprOrContext) {}

// EnterExprSub is called when production ExprSub is entered.
func (s *BaseVGrammarListener) EnterExprSub(ctx *ExprSubContext) {}

// ExitExprSub is called when production ExprSub is exited.
func (s *BaseVGrammarListener) ExitExprSub(ctx *ExprSubContext) {}

// EnterExprLiteral is called when production ExprLiteral is entered.
func (s *BaseVGrammarListener) EnterExprLiteral(ctx *ExprLiteralContext) {}

// ExitExprLiteral is called when production ExprLiteral is exited.
func (s *BaseVGrammarListener) ExitExprLiteral(ctx *ExprLiteralContext) {}

// EnterExprMul is called when production ExprMul is entered.
func (s *BaseVGrammarListener) EnterExprMul(ctx *ExprMulContext) {}

// ExitExprMul is called when production ExprMul is exited.
func (s *BaseVGrammarListener) ExitExprMul(ctx *ExprMulContext) {}

// EnterExprEqual is called when production ExprEqual is entered.
func (s *BaseVGrammarListener) EnterExprEqual(ctx *ExprEqualContext) {}

// ExitExprEqual is called when production ExprEqual is exited.
func (s *BaseVGrammarListener) ExitExprEqual(ctx *ExprEqualContext) {}

// EnterExprId is called when production ExprId is entered.
func (s *BaseVGrammarListener) EnterExprId(ctx *ExprIdContext) {}

// ExitExprId is called when production ExprId is exited.
func (s *BaseVGrammarListener) ExitExprId(ctx *ExprIdContext) {}

// EnterExprAdd is called when production ExprAdd is entered.
func (s *BaseVGrammarListener) EnterExprAdd(ctx *ExprAddContext) {}

// ExitExprAdd is called when production ExprAdd is exited.
func (s *BaseVGrammarListener) ExitExprAdd(ctx *ExprAddContext) {}

// EnterExprMod is called when production ExprMod is entered.
func (s *BaseVGrammarListener) EnterExprMod(ctx *ExprModContext) {}

// ExitExprMod is called when production ExprMod is exited.
func (s *BaseVGrammarListener) ExitExprMod(ctx *ExprModContext) {}

// EnterLitInt is called when production LitInt is entered.
func (s *BaseVGrammarListener) EnterLitInt(ctx *LitIntContext) {}

// ExitLitInt is called when production LitInt is exited.
func (s *BaseVGrammarListener) ExitLitInt(ctx *LitIntContext) {}

// EnterLitFloat is called when production LitFloat is entered.
func (s *BaseVGrammarListener) EnterLitFloat(ctx *LitFloatContext) {}

// ExitLitFloat is called when production LitFloat is exited.
func (s *BaseVGrammarListener) ExitLitFloat(ctx *LitFloatContext) {}

// EnterLitString is called when production LitString is entered.
func (s *BaseVGrammarListener) EnterLitString(ctx *LitStringContext) {}

// ExitLitString is called when production LitString is exited.
func (s *BaseVGrammarListener) ExitLitString(ctx *LitStringContext) {}

// EnterLitRune is called when production LitRune is entered.
func (s *BaseVGrammarListener) EnterLitRune(ctx *LitRuneContext) {}

// ExitLitRune is called when production LitRune is exited.
func (s *BaseVGrammarListener) ExitLitRune(ctx *LitRuneContext) {}

// EnterLitTrue is called when production LitTrue is entered.
func (s *BaseVGrammarListener) EnterLitTrue(ctx *LitTrueContext) {}

// ExitLitTrue is called when production LitTrue is exited.
func (s *BaseVGrammarListener) ExitLitTrue(ctx *LitTrueContext) {}

// EnterLitFalse is called when production LitFalse is entered.
func (s *BaseVGrammarListener) EnterLitFalse(ctx *LitFalseContext) {}

// ExitLitFalse is called when production LitFalse is exited.
func (s *BaseVGrammarListener) ExitLitFalse(ctx *LitFalseContext) {}

// EnterLitNil is called when production LitNil is entered.
func (s *BaseVGrammarListener) EnterLitNil(ctx *LitNilContext) {}

// ExitLitNil is called when production LitNil is exited.
func (s *BaseVGrammarListener) ExitLitNil(ctx *LitNilContext) {}

// EnterStmtBreak is called when production StmtBreak is entered.
func (s *BaseVGrammarListener) EnterStmtBreak(ctx *StmtBreakContext) {}

// ExitStmtBreak is called when production StmtBreak is exited.
func (s *BaseVGrammarListener) ExitStmtBreak(ctx *StmtBreakContext) {}

// EnterStmtContinue is called when production StmtContinue is entered.
func (s *BaseVGrammarListener) EnterStmtContinue(ctx *StmtContinueContext) {}

// ExitStmtContinue is called when production StmtContinue is exited.
func (s *BaseVGrammarListener) ExitStmtContinue(ctx *StmtContinueContext) {}

// EnterStmtReturn is called when production StmtReturn is entered.
func (s *BaseVGrammarListener) EnterStmtReturn(ctx *StmtReturnContext) {}

// ExitStmtReturn is called when production StmtReturn is exited.
func (s *BaseVGrammarListener) ExitStmtReturn(ctx *StmtReturnContext) {}

// EnterFunctionCall is called when production FunctionCall is entered.
func (s *BaseVGrammarListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production FunctionCall is exited.
func (s *BaseVGrammarListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterPrintCall is called when production PrintCall is entered.
func (s *BaseVGrammarListener) EnterPrintCall(ctx *PrintCallContext) {}

// ExitPrintCall is called when production PrintCall is exited.
func (s *BaseVGrammarListener) ExitPrintCall(ctx *PrintCallContext) {}

// EnterPrintLnCall is called when production PrintLnCall is entered.
func (s *BaseVGrammarListener) EnterPrintLnCall(ctx *PrintLnCallContext) {}

// ExitPrintLnCall is called when production PrintLnCall is exited.
func (s *BaseVGrammarListener) ExitPrintLnCall(ctx *PrintLnCallContext) {}

// EnterIf_stmt is called when production if_stmt is entered.
func (s *BaseVGrammarListener) EnterIf_stmt(ctx *If_stmtContext) {}

// ExitIf_stmt is called when production if_stmt is exited.
func (s *BaseVGrammarListener) ExitIf_stmt(ctx *If_stmtContext) {}

// EnterSwitch_stmt is called when production switch_stmt is entered.
func (s *BaseVGrammarListener) EnterSwitch_stmt(ctx *Switch_stmtContext) {}

// ExitSwitch_stmt is called when production switch_stmt is exited.
func (s *BaseVGrammarListener) ExitSwitch_stmt(ctx *Switch_stmtContext) {}

// EnterSwitch_case is called when production switch_case is entered.
func (s *BaseVGrammarListener) EnterSwitch_case(ctx *Switch_caseContext) {}

// ExitSwitch_case is called when production switch_case is exited.
func (s *BaseVGrammarListener) ExitSwitch_case(ctx *Switch_caseContext) {}

// EnterDefault_case is called when production default_case is entered.
func (s *BaseVGrammarListener) EnterDefault_case(ctx *Default_caseContext) {}

// ExitDefault_case is called when production default_case is exited.
func (s *BaseVGrammarListener) ExitDefault_case(ctx *Default_caseContext) {}

// EnterFor_stmt is called when production for_stmt is entered.
func (s *BaseVGrammarListener) EnterFor_stmt(ctx *For_stmtContext) {}

// ExitFor_stmt is called when production for_stmt is exited.
func (s *BaseVGrammarListener) ExitFor_stmt(ctx *For_stmtContext) {}

// EnterFunc_decl is called when production func_decl is entered.
func (s *BaseVGrammarListener) EnterFunc_decl(ctx *Func_declContext) {}

// ExitFunc_decl is called when production func_decl is exited.
func (s *BaseVGrammarListener) ExitFunc_decl(ctx *Func_declContext) {}

// EnterParam is called when production param is entered.
func (s *BaseVGrammarListener) EnterParam(ctx *ParamContext) {}

// ExitParam is called when production param is exited.
func (s *BaseVGrammarListener) ExitParam(ctx *ParamContext) {}

// EnterStruct_decl is called when production struct_decl is entered.
func (s *BaseVGrammarListener) EnterStruct_decl(ctx *Struct_declContext) {}

// ExitStruct_decl is called when production struct_decl is exited.
func (s *BaseVGrammarListener) ExitStruct_decl(ctx *Struct_declContext) {}

// EnterStruct_field is called when production struct_field is entered.
func (s *BaseVGrammarListener) EnterStruct_field(ctx *Struct_fieldContext) {}

// ExitStruct_field is called when production struct_field is exited.
func (s *BaseVGrammarListener) ExitStruct_field(ctx *Struct_fieldContext) {}
