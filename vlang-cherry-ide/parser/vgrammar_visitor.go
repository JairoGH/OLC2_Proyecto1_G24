// Code generated from vlang-cherry-ide/parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VGrammar
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by VGrammar.
type VGrammarVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by VGrammar#program.
	VisitProgram(ctx *ProgramContext) interface{}

	// Visit a parse tree produced by VGrammar#FuncionMain.
	VisitFuncionMain(ctx *FuncionMainContext) interface{}

	// Visit a parse tree produced by VGrammar#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararInferenciaMut.
	VisitDeclararInferenciaMut(ctx *DeclararInferenciaMutContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararInferencia.
	VisitDeclararInferencia(ctx *DeclararInferenciaContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclaraTipoValor.
	VisitDeclaraTipoValor(ctx *DeclaraTipoValorContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararTipo.
	VisitDeclararTipo(ctx *DeclararTipoContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararSinMutValor.
	VisitDeclararSinMutValor(ctx *DeclararSinMutValorContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararSinMutTipo.
	VisitDeclararSinMutTipo(ctx *DeclararSinMutTipoContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararVector.
	VisitDeclararVector(ctx *DeclararVectorContext) interface{}

	// Visit a parse tree produced by VGrammar#ListaItemsVector.
	VisitListaItemsVector(ctx *ListaItemsVectorContext) interface{}

	// Visit a parse tree produced by VGrammar#VectorItem.
	VisitVectorItem(ctx *VectorItemContext) interface{}

	// Visit a parse tree produced by VGrammar#PropVector.
	VisitPropVector(ctx *PropVectorContext) interface{}

	// Visit a parse tree produced by VGrammar#FuncionVector.
	VisitFuncionVector(ctx *FuncionVectorContext) interface{}

	// Visit a parse tree produced by VGrammar#repeating.
	VisitRepeating(ctx *RepeatingContext) interface{}

	// Visit a parse tree produced by VGrammar#var_type.
	VisitVar_type(ctx *Var_typeContext) interface{}

	// Visit a parse tree produced by VGrammar#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by VGrammar#VectorSimple.
	VisitVectorSimple(ctx *VectorSimpleContext) interface{}

	// Visit a parse tree produced by VGrammar#MatrizDoble.
	VisitMatrizDoble(ctx *MatrizDobleContext) interface{}

	// Visit a parse tree produced by VGrammar#matrix_type.
	VisitMatrix_type(ctx *Matrix_typeContext) interface{}

	// Visit a parse tree produced by VGrammar#aux_matrix_type.
	VisitAux_matrix_type(ctx *Aux_matrix_typeContext) interface{}

	// Visit a parse tree produced by VGrammar#AssignVectorItem.
	VisitAssignVectorItem(ctx *AssignVectorItemContext) interface{}

	// Visit a parse tree produced by VGrammar#AsignacionDirecta.
	VisitAsignacionDirecta(ctx *AsignacionDirectaContext) interface{}

	// Visit a parse tree produced by VGrammar#AsignacionAritmetica.
	VisitAsignacionAritmetica(ctx *AsignacionAritmeticaContext) interface{}

	// Visit a parse tree produced by VGrammar#AsignacionVector.
	VisitAsignacionVector(ctx *AsignacionVectorContext) interface{}

	// Visit a parse tree produced by VGrammar#ID_Patron.
	VisitID_Patron(ctx *ID_PatronContext) interface{}

	// Visit a parse tree produced by VGrammar#IntLiteral.
	VisitIntLiteral(ctx *IntLiteralContext) interface{}

	// Visit a parse tree produced by VGrammar#FloatLiteral.
	VisitFloatLiteral(ctx *FloatLiteralContext) interface{}

	// Visit a parse tree produced by VGrammar#StringLiteral.
	VisitStringLiteral(ctx *StringLiteralContext) interface{}

	// Visit a parse tree produced by VGrammar#BoolLiteral.
	VisitBoolLiteral(ctx *BoolLiteralContext) interface{}

	// Visit a parse tree produced by VGrammar#NilLiteral.
	VisitNilLiteral(ctx *NilLiteralContext) interface{}

	// Visit a parse tree produced by VGrammar#LiteralExp.
	VisitLiteralExp(ctx *LiteralExpContext) interface{}

	// Visit a parse tree produced by VGrammar#StructMethodExp.
	VisitStructMethodExp(ctx *StructMethodExpContext) interface{}

	// Visit a parse tree produced by VGrammar#StructInitExp.
	VisitStructInitExp(ctx *StructInitExpContext) interface{}

	// Visit a parse tree produced by VGrammar#RepeatingExp.
	VisitRepeatingExp(ctx *RepeatingExpContext) interface{}

	// Visit a parse tree produced by VGrammar#StructExp.
	VisitStructExp(ctx *StructExpContext) interface{}

	// Visit a parse tree produced by VGrammar#ExpBinario.
	VisitExpBinario(ctx *ExpBinarioContext) interface{}

	// Visit a parse tree produced by VGrammar#VectorPropExp.
	VisitVectorPropExp(ctx *VectorPropExpContext) interface{}

	// Visit a parse tree produced by VGrammar#VectorFuncExp.
	VisitVectorFuncExp(ctx *VectorFuncExpContext) interface{}

	// Visit a parse tree produced by VGrammar#ParentecisExp.
	VisitParentecisExp(ctx *ParentecisExpContext) interface{}

	// Visit a parse tree produced by VGrammar#ExpUnary.
	VisitExpUnary(ctx *ExpUnaryContext) interface{}

	// Visit a parse tree produced by VGrammar#IdExp.
	VisitIdExp(ctx *IdExpContext) interface{}

	// Visit a parse tree produced by VGrammar#FunctionCallExp.
	VisitFunctionCallExp(ctx *FunctionCallExpContext) interface{}

	// Visit a parse tree produced by VGrammar#VectorItemExp.
	VisitVectorItemExp(ctx *VectorItemExpContext) interface{}

	// Visit a parse tree produced by VGrammar#VectorExp.
	VisitVectorExp(ctx *VectorExpContext) interface{}

	// Visit a parse tree produced by VGrammar#IFstmt.
	VisitIFstmt(ctx *IFstmtContext) interface{}

	// Visit a parse tree produced by VGrammar#IFcadena.
	VisitIFcadena(ctx *IFcadenaContext) interface{}

	// Visit a parse tree produced by VGrammar#ElseStmt.
	VisitElseStmt(ctx *ElseStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#SwitchStmt.
	VisitSwitchStmt(ctx *SwitchStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#SwitchCase.
	VisitSwitchCase(ctx *SwitchCaseContext) interface{}

	// Visit a parse tree produced by VGrammar#DefaultCase.
	VisitDefaultCase(ctx *DefaultCaseContext) interface{}

	// Visit a parse tree produced by VGrammar#WhileStmt.
	VisitWhileStmt(ctx *WhileStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#ForStmt.
	VisitForStmt(ctx *ForStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#RangoNum.
	VisitRangoNum(ctx *RangoNumContext) interface{}

	// Visit a parse tree produced by VGrammar#GuardStmt.
	VisitGuardStmt(ctx *GuardStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#ReturnStmt.
	VisitReturnStmt(ctx *ReturnStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#BreakStmt.
	VisitBreakStmt(ctx *BreakStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#ContinueStmt.
	VisitContinueStmt(ctx *ContinueStmtContext) interface{}

	// Visit a parse tree produced by VGrammar#LlamarFuncion.
	VisitLlamarFuncion(ctx *LlamarFuncionContext) interface{}

	// Visit a parse tree produced by VGrammar#ArgList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by VGrammar#FuncionArg.
	VisitFuncionArg(ctx *FuncionArgContext) interface{}

	// Visit a parse tree produced by VGrammar#FuncionDeclerada.
	VisitFuncionDeclerada(ctx *FuncionDecleradaContext) interface{}

	// Visit a parse tree produced by VGrammar#MetodoStruct.
	VisitMetodoStruct(ctx *MetodoStructContext) interface{}

	// Visit a parse tree produced by VGrammar#MethodReceiver.
	VisitMethodReceiver(ctx *MethodReceiverContext) interface{}

	// Visit a parse tree produced by VGrammar#ParamList.
	VisitParamList(ctx *ParamListContext) interface{}

	// Visit a parse tree produced by VGrammar#FuncParam.
	VisitFuncParam(ctx *FuncParamContext) interface{}

	// Visit a parse tree produced by VGrammar#DeclararStruct.
	VisitDeclararStruct(ctx *DeclararStructContext) interface{}

	// Visit a parse tree produced by VGrammar#StructAttr.
	VisitStructAttr(ctx *StructAttrContext) interface{}

	// Visit a parse tree produced by VGrammar#StructFunc.
	VisitStructFunc(ctx *StructFuncContext) interface{}

	// Visit a parse tree produced by VGrammar#StructVector.
	VisitStructVector(ctx *StructVectorContext) interface{}

	// Visit a parse tree produced by VGrammar#StructInit.
	VisitStructInit(ctx *StructInitContext) interface{}

	// Visit a parse tree produced by VGrammar#StructInitList.
	VisitStructInitList(ctx *StructInitListContext) interface{}

	// Visit a parse tree produced by VGrammar#StructMethodCall.
	VisitStructMethodCall(ctx *StructMethodCallContext) interface{}

	// Visit a parse tree produced by VGrammar#StructInitField.
	VisitStructInitField(ctx *StructInitFieldContext) interface{}
}
