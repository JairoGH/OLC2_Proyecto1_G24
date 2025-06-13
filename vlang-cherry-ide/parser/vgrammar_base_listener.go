// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package parser // VGrammar
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

// EnterProgram is called when production program is entered.
func (s *BaseVGrammarListener) EnterProgram(ctx *ProgramContext) {}

// ExitProgram is called when production program is exited.
func (s *BaseVGrammarListener) ExitProgram(ctx *ProgramContext) {}

// EnterFuncionMain is called when production FuncionMain is entered.
func (s *BaseVGrammarListener) EnterFuncionMain(ctx *FuncionMainContext) {}

// ExitFuncionMain is called when production FuncionMain is exited.
func (s *BaseVGrammarListener) ExitFuncionMain(ctx *FuncionMainContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseVGrammarListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseVGrammarListener) ExitStmt(ctx *StmtContext) {}

// EnterDeclararInferenciaMut is called when production DeclararInferenciaMut is entered.
func (s *BaseVGrammarListener) EnterDeclararInferenciaMut(ctx *DeclararInferenciaMutContext) {}

// ExitDeclararInferenciaMut is called when production DeclararInferenciaMut is exited.
func (s *BaseVGrammarListener) ExitDeclararInferenciaMut(ctx *DeclararInferenciaMutContext) {}

// EnterDeclararInferencia is called when production DeclararInferencia is entered.
func (s *BaseVGrammarListener) EnterDeclararInferencia(ctx *DeclararInferenciaContext) {}

// ExitDeclararInferencia is called when production DeclararInferencia is exited.
func (s *BaseVGrammarListener) ExitDeclararInferencia(ctx *DeclararInferenciaContext) {}

// EnterDeclaraTipoValor is called when production DeclaraTipoValor is entered.
func (s *BaseVGrammarListener) EnterDeclaraTipoValor(ctx *DeclaraTipoValorContext) {}

// ExitDeclaraTipoValor is called when production DeclaraTipoValor is exited.
func (s *BaseVGrammarListener) ExitDeclaraTipoValor(ctx *DeclaraTipoValorContext) {}

// EnterDeclararTipo is called when production DeclararTipo is entered.
func (s *BaseVGrammarListener) EnterDeclararTipo(ctx *DeclararTipoContext) {}

// ExitDeclararTipo is called when production DeclararTipo is exited.
func (s *BaseVGrammarListener) ExitDeclararTipo(ctx *DeclararTipoContext) {}

// EnterDeclararSinMutValor is called when production DeclararSinMutValor is entered.
func (s *BaseVGrammarListener) EnterDeclararSinMutValor(ctx *DeclararSinMutValorContext) {}

// ExitDeclararSinMutValor is called when production DeclararSinMutValor is exited.
func (s *BaseVGrammarListener) ExitDeclararSinMutValor(ctx *DeclararSinMutValorContext) {}

// EnterDeclararSinMutTipo is called when production DeclararSinMutTipo is entered.
func (s *BaseVGrammarListener) EnterDeclararSinMutTipo(ctx *DeclararSinMutTipoContext) {}

// ExitDeclararSinMutTipo is called when production DeclararSinMutTipo is exited.
func (s *BaseVGrammarListener) ExitDeclararSinMutTipo(ctx *DeclararSinMutTipoContext) {}

// EnterDeclararVector is called when production DeclararVector is entered.
func (s *BaseVGrammarListener) EnterDeclararVector(ctx *DeclararVectorContext) {}

// ExitDeclararVector is called when production DeclararVector is exited.
func (s *BaseVGrammarListener) ExitDeclararVector(ctx *DeclararVectorContext) {}

// EnterListaItemsVector is called when production ListaItemsVector is entered.
func (s *BaseVGrammarListener) EnterListaItemsVector(ctx *ListaItemsVectorContext) {}

// ExitListaItemsVector is called when production ListaItemsVector is exited.
func (s *BaseVGrammarListener) ExitListaItemsVector(ctx *ListaItemsVectorContext) {}

// EnterVectorItem is called when production VectorItem is entered.
func (s *BaseVGrammarListener) EnterVectorItem(ctx *VectorItemContext) {}

// ExitVectorItem is called when production VectorItem is exited.
func (s *BaseVGrammarListener) ExitVectorItem(ctx *VectorItemContext) {}

// EnterPropVector is called when production PropVector is entered.
func (s *BaseVGrammarListener) EnterPropVector(ctx *PropVectorContext) {}

// ExitPropVector is called when production PropVector is exited.
func (s *BaseVGrammarListener) ExitPropVector(ctx *PropVectorContext) {}

// EnterFuncionVector is called when production FuncionVector is entered.
func (s *BaseVGrammarListener) EnterFuncionVector(ctx *FuncionVectorContext) {}

// ExitFuncionVector is called when production FuncionVector is exited.
func (s *BaseVGrammarListener) ExitFuncionVector(ctx *FuncionVectorContext) {}

// EnterRepeating is called when production repeating is entered.
func (s *BaseVGrammarListener) EnterRepeating(ctx *RepeatingContext) {}

// ExitRepeating is called when production repeating is exited.
func (s *BaseVGrammarListener) ExitRepeating(ctx *RepeatingContext) {}

// EnterVar_type is called when production var_type is entered.
func (s *BaseVGrammarListener) EnterVar_type(ctx *Var_typeContext) {}

// ExitVar_type is called when production var_type is exited.
func (s *BaseVGrammarListener) ExitVar_type(ctx *Var_typeContext) {}

// EnterType is called when production type is entered.
func (s *BaseVGrammarListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseVGrammarListener) ExitType(ctx *TypeContext) {}

// EnterVectorSimple is called when production VectorSimple is entered.
func (s *BaseVGrammarListener) EnterVectorSimple(ctx *VectorSimpleContext) {}

// ExitVectorSimple is called when production VectorSimple is exited.
func (s *BaseVGrammarListener) ExitVectorSimple(ctx *VectorSimpleContext) {}

// EnterMatrizDoble is called when production MatrizDoble is entered.
func (s *BaseVGrammarListener) EnterMatrizDoble(ctx *MatrizDobleContext) {}

// ExitMatrizDoble is called when production MatrizDoble is exited.
func (s *BaseVGrammarListener) ExitMatrizDoble(ctx *MatrizDobleContext) {}

// EnterMatrix_type is called when production matrix_type is entered.
func (s *BaseVGrammarListener) EnterMatrix_type(ctx *Matrix_typeContext) {}

// ExitMatrix_type is called when production matrix_type is exited.
func (s *BaseVGrammarListener) ExitMatrix_type(ctx *Matrix_typeContext) {}

// EnterAux_matrix_type is called when production aux_matrix_type is entered.
func (s *BaseVGrammarListener) EnterAux_matrix_type(ctx *Aux_matrix_typeContext) {}

// ExitAux_matrix_type is called when production aux_matrix_type is exited.
func (s *BaseVGrammarListener) ExitAux_matrix_type(ctx *Aux_matrix_typeContext) {}

// EnterAssignVectorItem is called when production AssignVectorItem is entered.
func (s *BaseVGrammarListener) EnterAssignVectorItem(ctx *AssignVectorItemContext) {}

// ExitAssignVectorItem is called when production AssignVectorItem is exited.
func (s *BaseVGrammarListener) ExitAssignVectorItem(ctx *AssignVectorItemContext) {}

// EnterAsignacionDirecta is called when production AsignacionDirecta is entered.
func (s *BaseVGrammarListener) EnterAsignacionDirecta(ctx *AsignacionDirectaContext) {}

// ExitAsignacionDirecta is called when production AsignacionDirecta is exited.
func (s *BaseVGrammarListener) ExitAsignacionDirecta(ctx *AsignacionDirectaContext) {}

// EnterAsignacionAritmetica is called when production AsignacionAritmetica is entered.
func (s *BaseVGrammarListener) EnterAsignacionAritmetica(ctx *AsignacionAritmeticaContext) {}

// ExitAsignacionAritmetica is called when production AsignacionAritmetica is exited.
func (s *BaseVGrammarListener) ExitAsignacionAritmetica(ctx *AsignacionAritmeticaContext) {}

// EnterAsignacionVector is called when production AsignacionVector is entered.
func (s *BaseVGrammarListener) EnterAsignacionVector(ctx *AsignacionVectorContext) {}

// ExitAsignacionVector is called when production AsignacionVector is exited.
func (s *BaseVGrammarListener) ExitAsignacionVector(ctx *AsignacionVectorContext) {}

// EnterID_Patron is called when production ID_Patron is entered.
func (s *BaseVGrammarListener) EnterID_Patron(ctx *ID_PatronContext) {}

// ExitID_Patron is called when production ID_Patron is exited.
func (s *BaseVGrammarListener) ExitID_Patron(ctx *ID_PatronContext) {}

// EnterIntLiteral is called when production IntLiteral is entered.
func (s *BaseVGrammarListener) EnterIntLiteral(ctx *IntLiteralContext) {}

// ExitIntLiteral is called when production IntLiteral is exited.
func (s *BaseVGrammarListener) ExitIntLiteral(ctx *IntLiteralContext) {}

// EnterFloatLiteral is called when production FloatLiteral is entered.
func (s *BaseVGrammarListener) EnterFloatLiteral(ctx *FloatLiteralContext) {}

// ExitFloatLiteral is called when production FloatLiteral is exited.
func (s *BaseVGrammarListener) ExitFloatLiteral(ctx *FloatLiteralContext) {}

// EnterStringLiteral is called when production StringLiteral is entered.
func (s *BaseVGrammarListener) EnterStringLiteral(ctx *StringLiteralContext) {}

// ExitStringLiteral is called when production StringLiteral is exited.
func (s *BaseVGrammarListener) ExitStringLiteral(ctx *StringLiteralContext) {}

// EnterBoolLiteral is called when production BoolLiteral is entered.
func (s *BaseVGrammarListener) EnterBoolLiteral(ctx *BoolLiteralContext) {}

// ExitBoolLiteral is called when production BoolLiteral is exited.
func (s *BaseVGrammarListener) ExitBoolLiteral(ctx *BoolLiteralContext) {}

// EnterNilLiteral is called when production NilLiteral is entered.
func (s *BaseVGrammarListener) EnterNilLiteral(ctx *NilLiteralContext) {}

// ExitNilLiteral is called when production NilLiteral is exited.
func (s *BaseVGrammarListener) ExitNilLiteral(ctx *NilLiteralContext) {}

// EnterLiteralExp is called when production LiteralExp is entered.
func (s *BaseVGrammarListener) EnterLiteralExp(ctx *LiteralExpContext) {}

// ExitLiteralExp is called when production LiteralExp is exited.
func (s *BaseVGrammarListener) ExitLiteralExp(ctx *LiteralExpContext) {}

// EnterStructMethodExp is called when production StructMethodExp is entered.
func (s *BaseVGrammarListener) EnterStructMethodExp(ctx *StructMethodExpContext) {}

// ExitStructMethodExp is called when production StructMethodExp is exited.
func (s *BaseVGrammarListener) ExitStructMethodExp(ctx *StructMethodExpContext) {}

// EnterStructInitExp is called when production StructInitExp is entered.
func (s *BaseVGrammarListener) EnterStructInitExp(ctx *StructInitExpContext) {}

// ExitStructInitExp is called when production StructInitExp is exited.
func (s *BaseVGrammarListener) ExitStructInitExp(ctx *StructInitExpContext) {}

// EnterRepeatingExp is called when production RepeatingExp is entered.
func (s *BaseVGrammarListener) EnterRepeatingExp(ctx *RepeatingExpContext) {}

// ExitRepeatingExp is called when production RepeatingExp is exited.
func (s *BaseVGrammarListener) ExitRepeatingExp(ctx *RepeatingExpContext) {}

// EnterStructExp is called when production StructExp is entered.
func (s *BaseVGrammarListener) EnterStructExp(ctx *StructExpContext) {}

// ExitStructExp is called when production StructExp is exited.
func (s *BaseVGrammarListener) ExitStructExp(ctx *StructExpContext) {}

// EnterExpBinario is called when production ExpBinario is entered.
func (s *BaseVGrammarListener) EnterExpBinario(ctx *ExpBinarioContext) {}

// ExitExpBinario is called when production ExpBinario is exited.
func (s *BaseVGrammarListener) ExitExpBinario(ctx *ExpBinarioContext) {}

// EnterVectorPropExp is called when production VectorPropExp is entered.
func (s *BaseVGrammarListener) EnterVectorPropExp(ctx *VectorPropExpContext) {}

// ExitVectorPropExp is called when production VectorPropExp is exited.
func (s *BaseVGrammarListener) ExitVectorPropExp(ctx *VectorPropExpContext) {}

// EnterVectorFuncExp is called when production VectorFuncExp is entered.
func (s *BaseVGrammarListener) EnterVectorFuncExp(ctx *VectorFuncExpContext) {}

// ExitVectorFuncExp is called when production VectorFuncExp is exited.
func (s *BaseVGrammarListener) ExitVectorFuncExp(ctx *VectorFuncExpContext) {}

// EnterParentecisExp is called when production ParentecisExp is entered.
func (s *BaseVGrammarListener) EnterParentecisExp(ctx *ParentecisExpContext) {}

// ExitParentecisExp is called when production ParentecisExp is exited.
func (s *BaseVGrammarListener) ExitParentecisExp(ctx *ParentecisExpContext) {}

// EnterExpUnary is called when production ExpUnary is entered.
func (s *BaseVGrammarListener) EnterExpUnary(ctx *ExpUnaryContext) {}

// ExitExpUnary is called when production ExpUnary is exited.
func (s *BaseVGrammarListener) ExitExpUnary(ctx *ExpUnaryContext) {}

// EnterIdExp is called when production IdExp is entered.
func (s *BaseVGrammarListener) EnterIdExp(ctx *IdExpContext) {}

// ExitIdExp is called when production IdExp is exited.
func (s *BaseVGrammarListener) ExitIdExp(ctx *IdExpContext) {}

// EnterFunctionCallExp is called when production FunctionCallExp is entered.
func (s *BaseVGrammarListener) EnterFunctionCallExp(ctx *FunctionCallExpContext) {}

// ExitFunctionCallExp is called when production FunctionCallExp is exited.
func (s *BaseVGrammarListener) ExitFunctionCallExp(ctx *FunctionCallExpContext) {}

// EnterVectorItemExp is called when production VectorItemExp is entered.
func (s *BaseVGrammarListener) EnterVectorItemExp(ctx *VectorItemExpContext) {}

// ExitVectorItemExp is called when production VectorItemExp is exited.
func (s *BaseVGrammarListener) ExitVectorItemExp(ctx *VectorItemExpContext) {}

// EnterVectorExp is called when production VectorExp is entered.
func (s *BaseVGrammarListener) EnterVectorExp(ctx *VectorExpContext) {}

// ExitVectorExp is called when production VectorExp is exited.
func (s *BaseVGrammarListener) ExitVectorExp(ctx *VectorExpContext) {}

// EnterIFstmt is called when production IFstmt is entered.
func (s *BaseVGrammarListener) EnterIFstmt(ctx *IFstmtContext) {}

// ExitIFstmt is called when production IFstmt is exited.
func (s *BaseVGrammarListener) ExitIFstmt(ctx *IFstmtContext) {}

// EnterIFcadena is called when production IFcadena is entered.
func (s *BaseVGrammarListener) EnterIFcadena(ctx *IFcadenaContext) {}

// ExitIFcadena is called when production IFcadena is exited.
func (s *BaseVGrammarListener) ExitIFcadena(ctx *IFcadenaContext) {}

// EnterElseStmt is called when production ElseStmt is entered.
func (s *BaseVGrammarListener) EnterElseStmt(ctx *ElseStmtContext) {}

// ExitElseStmt is called when production ElseStmt is exited.
func (s *BaseVGrammarListener) ExitElseStmt(ctx *ElseStmtContext) {}

// EnterSwitchStmt is called when production SwitchStmt is entered.
func (s *BaseVGrammarListener) EnterSwitchStmt(ctx *SwitchStmtContext) {}

// ExitSwitchStmt is called when production SwitchStmt is exited.
func (s *BaseVGrammarListener) ExitSwitchStmt(ctx *SwitchStmtContext) {}

// EnterSwitchCase is called when production SwitchCase is entered.
func (s *BaseVGrammarListener) EnterSwitchCase(ctx *SwitchCaseContext) {}

// ExitSwitchCase is called when production SwitchCase is exited.
func (s *BaseVGrammarListener) ExitSwitchCase(ctx *SwitchCaseContext) {}

// EnterDefaultCase is called when production DefaultCase is entered.
func (s *BaseVGrammarListener) EnterDefaultCase(ctx *DefaultCaseContext) {}

// ExitDefaultCase is called when production DefaultCase is exited.
func (s *BaseVGrammarListener) ExitDefaultCase(ctx *DefaultCaseContext) {}

// EnterWhileStmt is called when production WhileStmt is entered.
func (s *BaseVGrammarListener) EnterWhileStmt(ctx *WhileStmtContext) {}

// ExitWhileStmt is called when production WhileStmt is exited.
func (s *BaseVGrammarListener) ExitWhileStmt(ctx *WhileStmtContext) {}

// EnterForStmt is called when production ForStmt is entered.
func (s *BaseVGrammarListener) EnterForStmt(ctx *ForStmtContext) {}

// ExitForStmt is called when production ForStmt is exited.
func (s *BaseVGrammarListener) ExitForStmt(ctx *ForStmtContext) {}

// EnterRangoNum is called when production RangoNum is entered.
func (s *BaseVGrammarListener) EnterRangoNum(ctx *RangoNumContext) {}

// ExitRangoNum is called when production RangoNum is exited.
func (s *BaseVGrammarListener) ExitRangoNum(ctx *RangoNumContext) {}

// EnterGuardStmt is called when production GuardStmt is entered.
func (s *BaseVGrammarListener) EnterGuardStmt(ctx *GuardStmtContext) {}

// ExitGuardStmt is called when production GuardStmt is exited.
func (s *BaseVGrammarListener) ExitGuardStmt(ctx *GuardStmtContext) {}

// EnterReturnStmt is called when production ReturnStmt is entered.
func (s *BaseVGrammarListener) EnterReturnStmt(ctx *ReturnStmtContext) {}

// ExitReturnStmt is called when production ReturnStmt is exited.
func (s *BaseVGrammarListener) ExitReturnStmt(ctx *ReturnStmtContext) {}

// EnterBreakStmt is called when production BreakStmt is entered.
func (s *BaseVGrammarListener) EnterBreakStmt(ctx *BreakStmtContext) {}

// ExitBreakStmt is called when production BreakStmt is exited.
func (s *BaseVGrammarListener) ExitBreakStmt(ctx *BreakStmtContext) {}

// EnterContinueStmt is called when production ContinueStmt is entered.
func (s *BaseVGrammarListener) EnterContinueStmt(ctx *ContinueStmtContext) {}

// ExitContinueStmt is called when production ContinueStmt is exited.
func (s *BaseVGrammarListener) ExitContinueStmt(ctx *ContinueStmtContext) {}

// EnterLlamarFuncion is called when production LlamarFuncion is entered.
func (s *BaseVGrammarListener) EnterLlamarFuncion(ctx *LlamarFuncionContext) {}

// ExitLlamarFuncion is called when production LlamarFuncion is exited.
func (s *BaseVGrammarListener) ExitLlamarFuncion(ctx *LlamarFuncionContext) {}

// EnterArgList is called when production ArgList is entered.
func (s *BaseVGrammarListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production ArgList is exited.
func (s *BaseVGrammarListener) ExitArgList(ctx *ArgListContext) {}

// EnterFuncionArg is called when production FuncionArg is entered.
func (s *BaseVGrammarListener) EnterFuncionArg(ctx *FuncionArgContext) {}

// ExitFuncionArg is called when production FuncionArg is exited.
func (s *BaseVGrammarListener) ExitFuncionArg(ctx *FuncionArgContext) {}

// EnterFuncionDeclerada is called when production FuncionDeclerada is entered.
func (s *BaseVGrammarListener) EnterFuncionDeclerada(ctx *FuncionDecleradaContext) {}

// ExitFuncionDeclerada is called when production FuncionDeclerada is exited.
func (s *BaseVGrammarListener) ExitFuncionDeclerada(ctx *FuncionDecleradaContext) {}

// EnterMetodoStruct is called when production MetodoStruct is entered.
func (s *BaseVGrammarListener) EnterMetodoStruct(ctx *MetodoStructContext) {}

// ExitMetodoStruct is called when production MetodoStruct is exited.
func (s *BaseVGrammarListener) ExitMetodoStruct(ctx *MetodoStructContext) {}

// EnterMethodReceiver is called when production MethodReceiver is entered.
func (s *BaseVGrammarListener) EnterMethodReceiver(ctx *MethodReceiverContext) {}

// ExitMethodReceiver is called when production MethodReceiver is exited.
func (s *BaseVGrammarListener) ExitMethodReceiver(ctx *MethodReceiverContext) {}

// EnterParamList is called when production ParamList is entered.
func (s *BaseVGrammarListener) EnterParamList(ctx *ParamListContext) {}

// ExitParamList is called when production ParamList is exited.
func (s *BaseVGrammarListener) ExitParamList(ctx *ParamListContext) {}

// EnterFuncParam is called when production FuncParam is entered.
func (s *BaseVGrammarListener) EnterFuncParam(ctx *FuncParamContext) {}

// ExitFuncParam is called when production FuncParam is exited.
func (s *BaseVGrammarListener) ExitFuncParam(ctx *FuncParamContext) {}

// EnterDeclararStruct is called when production DeclararStruct is entered.
func (s *BaseVGrammarListener) EnterDeclararStruct(ctx *DeclararStructContext) {}

// ExitDeclararStruct is called when production DeclararStruct is exited.
func (s *BaseVGrammarListener) ExitDeclararStruct(ctx *DeclararStructContext) {}

// EnterStructAttr is called when production StructAttr is entered.
func (s *BaseVGrammarListener) EnterStructAttr(ctx *StructAttrContext) {}

// ExitStructAttr is called when production StructAttr is exited.
func (s *BaseVGrammarListener) ExitStructAttr(ctx *StructAttrContext) {}

// EnterStructFunc is called when production StructFunc is entered.
func (s *BaseVGrammarListener) EnterStructFunc(ctx *StructFuncContext) {}

// ExitStructFunc is called when production StructFunc is exited.
func (s *BaseVGrammarListener) ExitStructFunc(ctx *StructFuncContext) {}

// EnterStructVector is called when production StructVector is entered.
func (s *BaseVGrammarListener) EnterStructVector(ctx *StructVectorContext) {}

// ExitStructVector is called when production StructVector is exited.
func (s *BaseVGrammarListener) ExitStructVector(ctx *StructVectorContext) {}

// EnterStructInit is called when production StructInit is entered.
func (s *BaseVGrammarListener) EnterStructInit(ctx *StructInitContext) {}

// ExitStructInit is called when production StructInit is exited.
func (s *BaseVGrammarListener) ExitStructInit(ctx *StructInitContext) {}

// EnterStructInitList is called when production StructInitList is entered.
func (s *BaseVGrammarListener) EnterStructInitList(ctx *StructInitListContext) {}

// ExitStructInitList is called when production StructInitList is exited.
func (s *BaseVGrammarListener) ExitStructInitList(ctx *StructInitListContext) {}

// EnterStructMethodCall is called when production StructMethodCall is entered.
func (s *BaseVGrammarListener) EnterStructMethodCall(ctx *StructMethodCallContext) {}

// ExitStructMethodCall is called when production StructMethodCall is exited.
func (s *BaseVGrammarListener) ExitStructMethodCall(ctx *StructMethodCallContext) {}

// EnterStructInitField is called when production StructInitField is entered.
func (s *BaseVGrammarListener) EnterStructInitField(ctx *StructInitFieldContext) {}

// ExitStructInitField is called when production StructInitField is exited.
func (s *BaseVGrammarListener) ExitStructInitField(ctx *StructInitFieldContext) {}
