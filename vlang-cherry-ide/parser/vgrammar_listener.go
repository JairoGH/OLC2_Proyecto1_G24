// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VGrammar
import "github.com/antlr4-go/antlr/v4"

// VGrammarListener is a complete listener for a parse tree produced by VGrammar.
type VGrammarListener interface {
	antlr.ParseTreeListener

	// EnterProgram is called when entering the program production.
	EnterProgram(c *ProgramContext)

	// EnterFuncionMain is called when entering the FuncionMain production.
	EnterFuncionMain(c *FuncionMainContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterDeclararInferenciaMut is called when entering the DeclararInferenciaMut production.
	EnterDeclararInferenciaMut(c *DeclararInferenciaMutContext)

	// EnterDeclararInferencia is called when entering the DeclararInferencia production.
	EnterDeclararInferencia(c *DeclararInferenciaContext)

	// EnterDeclaraTipoValor is called when entering the DeclaraTipoValor production.
	EnterDeclaraTipoValor(c *DeclaraTipoValorContext)

	// EnterDeclararTipo is called when entering the DeclararTipo production.
	EnterDeclararTipo(c *DeclararTipoContext)

	// EnterDeclararSinMutValor is called when entering the DeclararSinMutValor production.
	EnterDeclararSinMutValor(c *DeclararSinMutValorContext)

	// EnterDeclararSinMutTipo is called when entering the DeclararSinMutTipo production.
	EnterDeclararSinMutTipo(c *DeclararSinMutTipoContext)

	// EnterDeclararVector is called when entering the DeclararVector production.
	EnterDeclararVector(c *DeclararVectorContext)

	// EnterListaItemsVector is called when entering the ListaItemsVector production.
	EnterListaItemsVector(c *ListaItemsVectorContext)

	// EnterVectorItem is called when entering the VectorItem production.
	EnterVectorItem(c *VectorItemContext)

	// EnterPropVector is called when entering the PropVector production.
	EnterPropVector(c *PropVectorContext)

	// EnterFuncionVector is called when entering the FuncionVector production.
	EnterFuncionVector(c *FuncionVectorContext)

	// EnterRepeating is called when entering the repeating production.
	EnterRepeating(c *RepeatingContext)

	// EnterVar_type is called when entering the var_type production.
	EnterVar_type(c *Var_typeContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterVectorSimple is called when entering the VectorSimple production.
	EnterVectorSimple(c *VectorSimpleContext)

	// EnterMatrizDoble is called when entering the MatrizDoble production.
	EnterMatrizDoble(c *MatrizDobleContext)

	// EnterMatrix_type is called when entering the matrix_type production.
	EnterMatrix_type(c *Matrix_typeContext)

	// EnterAux_matrix_type is called when entering the aux_matrix_type production.
	EnterAux_matrix_type(c *Aux_matrix_typeContext)

	// EnterAssignVectorItem is called when entering the AssignVectorItem production.
	EnterAssignVectorItem(c *AssignVectorItemContext)

	// EnterAsignacionDirecta is called when entering the AsignacionDirecta production.
	EnterAsignacionDirecta(c *AsignacionDirectaContext)

	// EnterAsignacionAritmetica is called when entering the AsignacionAritmetica production.
	EnterAsignacionAritmetica(c *AsignacionAritmeticaContext)

	// EnterAsignacionVector is called when entering the AsignacionVector production.
	EnterAsignacionVector(c *AsignacionVectorContext)

	// EnterID_Patron is called when entering the ID_Patron production.
	EnterID_Patron(c *ID_PatronContext)

	// EnterIntLiteral is called when entering the IntLiteral production.
	EnterIntLiteral(c *IntLiteralContext)

	// EnterFloatLiteral is called when entering the FloatLiteral production.
	EnterFloatLiteral(c *FloatLiteralContext)

	// EnterStringLiteral is called when entering the StringLiteral production.
	EnterStringLiteral(c *StringLiteralContext)

	// EnterBoolLiteral is called when entering the BoolLiteral production.
	EnterBoolLiteral(c *BoolLiteralContext)

	// EnterNilLiteral is called when entering the NilLiteral production.
	EnterNilLiteral(c *NilLiteralContext)

	// EnterLiteralExp is called when entering the LiteralExp production.
	EnterLiteralExp(c *LiteralExpContext)

	// EnterStructMethodExp is called when entering the StructMethodExp production.
	EnterStructMethodExp(c *StructMethodExpContext)

	// EnterStructInitExp is called when entering the StructInitExp production.
	EnterStructInitExp(c *StructInitExpContext)

	// EnterRepeatingExp is called when entering the RepeatingExp production.
	EnterRepeatingExp(c *RepeatingExpContext)

	// EnterStructExp is called when entering the StructExp production.
	EnterStructExp(c *StructExpContext)

	// EnterExpBinario is called when entering the ExpBinario production.
	EnterExpBinario(c *ExpBinarioContext)

	// EnterVectorPropExp is called when entering the VectorPropExp production.
	EnterVectorPropExp(c *VectorPropExpContext)

	// EnterVectorFuncExp is called when entering the VectorFuncExp production.
	EnterVectorFuncExp(c *VectorFuncExpContext)

	// EnterParentecisExp is called when entering the ParentecisExp production.
	EnterParentecisExp(c *ParentecisExpContext)

	// EnterExpUnary is called when entering the ExpUnary production.
	EnterExpUnary(c *ExpUnaryContext)

	// EnterIdExp is called when entering the IdExp production.
	EnterIdExp(c *IdExpContext)

	// EnterFunctionCallExp is called when entering the FunctionCallExp production.
	EnterFunctionCallExp(c *FunctionCallExpContext)

	// EnterVectorItemExp is called when entering the VectorItemExp production.
	EnterVectorItemExp(c *VectorItemExpContext)

	// EnterVectorExp is called when entering the VectorExp production.
	EnterVectorExp(c *VectorExpContext)

	// EnterIFstmt is called when entering the IFstmt production.
	EnterIFstmt(c *IFstmtContext)

	// EnterIFcadena is called when entering the IFcadena production.
	EnterIFcadena(c *IFcadenaContext)

	// EnterElseStmt is called when entering the ElseStmt production.
	EnterElseStmt(c *ElseStmtContext)

	// EnterSwitchStmt is called when entering the SwitchStmt production.
	EnterSwitchStmt(c *SwitchStmtContext)

	// EnterSwitchCase is called when entering the SwitchCase production.
	EnterSwitchCase(c *SwitchCaseContext)

	// EnterDefaultCase is called when entering the DefaultCase production.
	EnterDefaultCase(c *DefaultCaseContext)

	// EnterWhileStmt is called when entering the WhileStmt production.
	EnterWhileStmt(c *WhileStmtContext)

	// EnterForStmt is called when entering the ForStmt production.
	EnterForStmt(c *ForStmtContext)

	// EnterRangoNum is called when entering the RangoNum production.
	EnterRangoNum(c *RangoNumContext)

	// EnterGuardStmt is called when entering the GuardStmt production.
	EnterGuardStmt(c *GuardStmtContext)

	// EnterReturnStmt is called when entering the ReturnStmt production.
	EnterReturnStmt(c *ReturnStmtContext)

	// EnterBreakStmt is called when entering the BreakStmt production.
	EnterBreakStmt(c *BreakStmtContext)

	// EnterContinueStmt is called when entering the ContinueStmt production.
	EnterContinueStmt(c *ContinueStmtContext)

	// EnterLlamarFuncion is called when entering the LlamarFuncion production.
	EnterLlamarFuncion(c *LlamarFuncionContext)

	// EnterArgList is called when entering the ArgList production.
	EnterArgList(c *ArgListContext)

	// EnterFuncionArg is called when entering the FuncionArg production.
	EnterFuncionArg(c *FuncionArgContext)

	// EnterFuncionDeclerada is called when entering the FuncionDeclerada production.
	EnterFuncionDeclerada(c *FuncionDecleradaContext)

	// EnterMetodoStruct is called when entering the MetodoStruct production.
	EnterMetodoStruct(c *MetodoStructContext)

	// EnterMethodReceiver is called when entering the MethodReceiver production.
	EnterMethodReceiver(c *MethodReceiverContext)

	// EnterParamList is called when entering the ParamList production.
	EnterParamList(c *ParamListContext)

	// EnterFuncParam is called when entering the FuncParam production.
	EnterFuncParam(c *FuncParamContext)

	// EnterDeclararStruct is called when entering the DeclararStruct production.
	EnterDeclararStruct(c *DeclararStructContext)

	// EnterStructAttr is called when entering the StructAttr production.
	EnterStructAttr(c *StructAttrContext)

	// EnterStructFunc is called when entering the StructFunc production.
	EnterStructFunc(c *StructFuncContext)

	// EnterStructVector is called when entering the StructVector production.
	EnterStructVector(c *StructVectorContext)

	// EnterStructInit is called when entering the StructInit production.
	EnterStructInit(c *StructInitContext)

	// EnterStructInitList is called when entering the StructInitList production.
	EnterStructInitList(c *StructInitListContext)

	// EnterStructMethodCall is called when entering the StructMethodCall production.
	EnterStructMethodCall(c *StructMethodCallContext)

	// EnterStructInitField is called when entering the StructInitField production.
	EnterStructInitField(c *StructInitFieldContext)

	// ExitProgram is called when exiting the program production.
	ExitProgram(c *ProgramContext)

	// ExitFuncionMain is called when exiting the FuncionMain production.
	ExitFuncionMain(c *FuncionMainContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitDeclararInferenciaMut is called when exiting the DeclararInferenciaMut production.
	ExitDeclararInferenciaMut(c *DeclararInferenciaMutContext)

	// ExitDeclararInferencia is called when exiting the DeclararInferencia production.
	ExitDeclararInferencia(c *DeclararInferenciaContext)

	// ExitDeclaraTipoValor is called when exiting the DeclaraTipoValor production.
	ExitDeclaraTipoValor(c *DeclaraTipoValorContext)

	// ExitDeclararTipo is called when exiting the DeclararTipo production.
	ExitDeclararTipo(c *DeclararTipoContext)

	// ExitDeclararSinMutValor is called when exiting the DeclararSinMutValor production.
	ExitDeclararSinMutValor(c *DeclararSinMutValorContext)

	// ExitDeclararSinMutTipo is called when exiting the DeclararSinMutTipo production.
	ExitDeclararSinMutTipo(c *DeclararSinMutTipoContext)

	// ExitDeclararVector is called when exiting the DeclararVector production.
	ExitDeclararVector(c *DeclararVectorContext)

	// ExitListaItemsVector is called when exiting the ListaItemsVector production.
	ExitListaItemsVector(c *ListaItemsVectorContext)

	// ExitVectorItem is called when exiting the VectorItem production.
	ExitVectorItem(c *VectorItemContext)

	// ExitPropVector is called when exiting the PropVector production.
	ExitPropVector(c *PropVectorContext)

	// ExitFuncionVector is called when exiting the FuncionVector production.
	ExitFuncionVector(c *FuncionVectorContext)

	// ExitRepeating is called when exiting the repeating production.
	ExitRepeating(c *RepeatingContext)

	// ExitVar_type is called when exiting the var_type production.
	ExitVar_type(c *Var_typeContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitVectorSimple is called when exiting the VectorSimple production.
	ExitVectorSimple(c *VectorSimpleContext)

	// ExitMatrizDoble is called when exiting the MatrizDoble production.
	ExitMatrizDoble(c *MatrizDobleContext)

	// ExitMatrix_type is called when exiting the matrix_type production.
	ExitMatrix_type(c *Matrix_typeContext)

	// ExitAux_matrix_type is called when exiting the aux_matrix_type production.
	ExitAux_matrix_type(c *Aux_matrix_typeContext)

	// ExitAssignVectorItem is called when exiting the AssignVectorItem production.
	ExitAssignVectorItem(c *AssignVectorItemContext)

	// ExitAsignacionDirecta is called when exiting the AsignacionDirecta production.
	ExitAsignacionDirecta(c *AsignacionDirectaContext)

	// ExitAsignacionAritmetica is called when exiting the AsignacionAritmetica production.
	ExitAsignacionAritmetica(c *AsignacionAritmeticaContext)

	// ExitAsignacionVector is called when exiting the AsignacionVector production.
	ExitAsignacionVector(c *AsignacionVectorContext)

	// ExitID_Patron is called when exiting the ID_Patron production.
	ExitID_Patron(c *ID_PatronContext)

	// ExitIntLiteral is called when exiting the IntLiteral production.
	ExitIntLiteral(c *IntLiteralContext)

	// ExitFloatLiteral is called when exiting the FloatLiteral production.
	ExitFloatLiteral(c *FloatLiteralContext)

	// ExitStringLiteral is called when exiting the StringLiteral production.
	ExitStringLiteral(c *StringLiteralContext)

	// ExitBoolLiteral is called when exiting the BoolLiteral production.
	ExitBoolLiteral(c *BoolLiteralContext)

	// ExitNilLiteral is called when exiting the NilLiteral production.
	ExitNilLiteral(c *NilLiteralContext)

	// ExitLiteralExp is called when exiting the LiteralExp production.
	ExitLiteralExp(c *LiteralExpContext)

	// ExitStructMethodExp is called when exiting the StructMethodExp production.
	ExitStructMethodExp(c *StructMethodExpContext)

	// ExitStructInitExp is called when exiting the StructInitExp production.
	ExitStructInitExp(c *StructInitExpContext)

	// ExitRepeatingExp is called when exiting the RepeatingExp production.
	ExitRepeatingExp(c *RepeatingExpContext)

	// ExitStructExp is called when exiting the StructExp production.
	ExitStructExp(c *StructExpContext)

	// ExitExpBinario is called when exiting the ExpBinario production.
	ExitExpBinario(c *ExpBinarioContext)

	// ExitVectorPropExp is called when exiting the VectorPropExp production.
	ExitVectorPropExp(c *VectorPropExpContext)

	// ExitVectorFuncExp is called when exiting the VectorFuncExp production.
	ExitVectorFuncExp(c *VectorFuncExpContext)

	// ExitParentecisExp is called when exiting the ParentecisExp production.
	ExitParentecisExp(c *ParentecisExpContext)

	// ExitExpUnary is called when exiting the ExpUnary production.
	ExitExpUnary(c *ExpUnaryContext)

	// ExitIdExp is called when exiting the IdExp production.
	ExitIdExp(c *IdExpContext)

	// ExitFunctionCallExp is called when exiting the FunctionCallExp production.
	ExitFunctionCallExp(c *FunctionCallExpContext)

	// ExitVectorItemExp is called when exiting the VectorItemExp production.
	ExitVectorItemExp(c *VectorItemExpContext)

	// ExitVectorExp is called when exiting the VectorExp production.
	ExitVectorExp(c *VectorExpContext)

	// ExitIFstmt is called when exiting the IFstmt production.
	ExitIFstmt(c *IFstmtContext)

	// ExitIFcadena is called when exiting the IFcadena production.
	ExitIFcadena(c *IFcadenaContext)

	// ExitElseStmt is called when exiting the ElseStmt production.
	ExitElseStmt(c *ElseStmtContext)

	// ExitSwitchStmt is called when exiting the SwitchStmt production.
	ExitSwitchStmt(c *SwitchStmtContext)

	// ExitSwitchCase is called when exiting the SwitchCase production.
	ExitSwitchCase(c *SwitchCaseContext)

	// ExitDefaultCase is called when exiting the DefaultCase production.
	ExitDefaultCase(c *DefaultCaseContext)

	// ExitWhileStmt is called when exiting the WhileStmt production.
	ExitWhileStmt(c *WhileStmtContext)

	// ExitForStmt is called when exiting the ForStmt production.
	ExitForStmt(c *ForStmtContext)

	// ExitRangoNum is called when exiting the RangoNum production.
	ExitRangoNum(c *RangoNumContext)

	// ExitGuardStmt is called when exiting the GuardStmt production.
	ExitGuardStmt(c *GuardStmtContext)

	// ExitReturnStmt is called when exiting the ReturnStmt production.
	ExitReturnStmt(c *ReturnStmtContext)

	// ExitBreakStmt is called when exiting the BreakStmt production.
	ExitBreakStmt(c *BreakStmtContext)

	// ExitContinueStmt is called when exiting the ContinueStmt production.
	ExitContinueStmt(c *ContinueStmtContext)

	// ExitLlamarFuncion is called when exiting the LlamarFuncion production.
	ExitLlamarFuncion(c *LlamarFuncionContext)

	// ExitArgList is called when exiting the ArgList production.
	ExitArgList(c *ArgListContext)

	// ExitFuncionArg is called when exiting the FuncionArg production.
	ExitFuncionArg(c *FuncionArgContext)

	// ExitFuncionDeclerada is called when exiting the FuncionDeclerada production.
	ExitFuncionDeclerada(c *FuncionDecleradaContext)

	// ExitMetodoStruct is called when exiting the MetodoStruct production.
	ExitMetodoStruct(c *MetodoStructContext)

	// ExitMethodReceiver is called when exiting the MethodReceiver production.
	ExitMethodReceiver(c *MethodReceiverContext)

	// ExitParamList is called when exiting the ParamList production.
	ExitParamList(c *ParamListContext)

	// ExitFuncParam is called when exiting the FuncParam production.
	ExitFuncParam(c *FuncParamContext)

	// ExitDeclararStruct is called when exiting the DeclararStruct production.
	ExitDeclararStruct(c *DeclararStructContext)

	// ExitStructAttr is called when exiting the StructAttr production.
	ExitStructAttr(c *StructAttrContext)

	// ExitStructFunc is called when exiting the StructFunc production.
	ExitStructFunc(c *StructFuncContext)

	// ExitStructVector is called when exiting the StructVector production.
	ExitStructVector(c *StructVectorContext)

	// ExitStructInit is called when exiting the StructInit production.
	ExitStructInit(c *StructInitContext)

	// ExitStructInitList is called when exiting the StructInitList production.
	ExitStructInitList(c *StructInitListContext)

	// ExitStructMethodCall is called when exiting the StructMethodCall production.
	ExitStructMethodCall(c *StructMethodCallContext)

	// ExitStructInitField is called when exiting the StructInitField production.
	ExitStructInitField(c *StructInitFieldContext)
}
