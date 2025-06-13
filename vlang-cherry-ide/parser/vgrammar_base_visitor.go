// Code generated from vlang-cherry-ide/parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VGrammar
import "github.com/antlr4-go/antlr/v4"

type BaseVGrammarVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseVGrammarVisitor) VisitProgram(ctx *ProgramContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFuncionMain(ctx *FuncionMainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararInferenciaMut(ctx *DeclararInferenciaMutContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararInferencia(ctx *DeclararInferenciaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclaraTipoValor(ctx *DeclaraTipoValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararTipo(ctx *DeclararTipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararSinMutValor(ctx *DeclararSinMutValorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararSinMutTipo(ctx *DeclararSinMutTipoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararVector(ctx *DeclararVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitListaItemsVector(ctx *ListaItemsVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitVectorItem(ctx *VectorItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitPropVector(ctx *PropVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFuncionVector(ctx *FuncionVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitRepeating(ctx *RepeatingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitVar_type(ctx *Var_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitVectorSimple(ctx *VectorSimpleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitMatrizDoble(ctx *MatrizDobleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitMatrix_type(ctx *Matrix_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAux_matrix_type(ctx *Aux_matrix_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAssignVectorItem(ctx *AssignVectorItemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAsignacionDirecta(ctx *AsignacionDirectaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAsignacionAritmetica(ctx *AsignacionAritmeticaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitAsignacionVector(ctx *AsignacionVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitID_Patron(ctx *ID_PatronContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitIntLiteral(ctx *IntLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFloatLiteral(ctx *FloatLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStringLiteral(ctx *StringLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitBoolLiteral(ctx *BoolLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitNilLiteral(ctx *NilLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLiteralExp(ctx *LiteralExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStructMethodExp(ctx *StructMethodExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStructInitExp(ctx *StructInitExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitRepeatingExp(ctx *RepeatingExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStructExp(ctx *StructExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExpBinario(ctx *ExpBinarioContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitVectorPropExp(ctx *VectorPropExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitVectorFuncExp(ctx *VectorFuncExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitParentecisExp(ctx *ParentecisExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitExpUnary(ctx *ExpUnaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitIdExp(ctx *IdExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFunctionCallExp(ctx *FunctionCallExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitVectorItemExp(ctx *VectorItemExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitVectorExp(ctx *VectorExpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitIFstmt(ctx *IFstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitIFcadena(ctx *IFcadenaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitElseStmt(ctx *ElseStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitSwitchStmt(ctx *SwitchStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitSwitchCase(ctx *SwitchCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDefaultCase(ctx *DefaultCaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitWhileStmt(ctx *WhileStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitForStmt(ctx *ForStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitRangoNum(ctx *RangoNumContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitGuardStmt(ctx *GuardStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitReturnStmt(ctx *ReturnStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitBreakStmt(ctx *BreakStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitContinueStmt(ctx *ContinueStmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitLlamarFuncion(ctx *LlamarFuncionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFuncionArg(ctx *FuncionArgContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFuncionDeclerada(ctx *FuncionDecleradaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitMetodoStruct(ctx *MetodoStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitMethodReceiver(ctx *MethodReceiverContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitParamList(ctx *ParamListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitFuncParam(ctx *FuncParamContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitDeclararStruct(ctx *DeclararStructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStructAttr(ctx *StructAttrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStructFunc(ctx *StructFuncContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStructVector(ctx *StructVectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStructInit(ctx *StructInitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStructInitList(ctx *StructInitListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStructMethodCall(ctx *StructMethodCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseVGrammarVisitor) VisitStructInitField(ctx *StructInitFieldContext) interface{} {
	return v.VisitChildren(ctx)
}
