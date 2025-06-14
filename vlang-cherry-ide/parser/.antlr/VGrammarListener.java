// Generated from /home/daaniieel/Escritorio/PROYECTO1/vlang-cherry-ide/parser/VGrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link VGrammar}.
 */
public interface VGrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link VGrammar#program}.
	 * @param ctx the parse tree
	 */
	void enterProgram(VGrammar.ProgramContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#program}.
	 * @param ctx the parse tree
	 */
	void exitProgram(VGrammar.ProgramContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncionMain}
	 * labeled alternative in {@link VGrammar#main_func}.
	 * @param ctx the parse tree
	 */
	void enterFuncionMain(VGrammar.FuncionMainContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncionMain}
	 * labeled alternative in {@link VGrammar#main_func}.
	 * @param ctx the parse tree
	 */
	void exitFuncionMain(VGrammar.FuncionMainContext ctx);
	/**
	 * Enter a parse tree produced by {@link VGrammar#stmt}.
	 * @param ctx the parse tree
	 */
	void enterStmt(VGrammar.StmtContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#stmt}.
	 * @param ctx the parse tree
	 */
	void exitStmt(VGrammar.StmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararInferenciaMut}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclararInferenciaMut(VGrammar.DeclararInferenciaMutContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararInferenciaMut}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclararInferenciaMut(VGrammar.DeclararInferenciaMutContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararInferencia}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclararInferencia(VGrammar.DeclararInferenciaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararInferencia}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclararInferencia(VGrammar.DeclararInferenciaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclaraTipoValor}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclaraTipoValor(VGrammar.DeclaraTipoValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclaraTipoValor}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclaraTipoValor(VGrammar.DeclaraTipoValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararTipo}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclararTipo(VGrammar.DeclararTipoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararTipo}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclararTipo(VGrammar.DeclararTipoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararSinMutValor}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclararSinMutValor(VGrammar.DeclararSinMutValorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararSinMutValor}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclararSinMutValor(VGrammar.DeclararSinMutValorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararSinMutTipo}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclararSinMutTipo(VGrammar.DeclararSinMutTipoContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararSinMutTipo}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclararSinMutTipo(VGrammar.DeclararSinMutTipoContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararVector}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDeclararVector(VGrammar.DeclararVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararVector}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDeclararVector(VGrammar.DeclararVectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ListaItemsVector}
	 * labeled alternative in {@link VGrammar#vector_expr}.
	 * @param ctx the parse tree
	 */
	void enterListaItemsVector(VGrammar.ListaItemsVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ListaItemsVector}
	 * labeled alternative in {@link VGrammar#vector_expr}.
	 * @param ctx the parse tree
	 */
	void exitListaItemsVector(VGrammar.ListaItemsVectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorItem}
	 * labeled alternative in {@link VGrammar#vector_item}.
	 * @param ctx the parse tree
	 */
	void enterVectorItem(VGrammar.VectorItemContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorItem}
	 * labeled alternative in {@link VGrammar#vector_item}.
	 * @param ctx the parse tree
	 */
	void exitVectorItem(VGrammar.VectorItemContext ctx);
	/**
	 * Enter a parse tree produced by the {@code PropVector}
	 * labeled alternative in {@link VGrammar#vector_prop}.
	 * @param ctx the parse tree
	 */
	void enterPropVector(VGrammar.PropVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code PropVector}
	 * labeled alternative in {@link VGrammar#vector_prop}.
	 * @param ctx the parse tree
	 */
	void exitPropVector(VGrammar.PropVectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncionVector}
	 * labeled alternative in {@link VGrammar#vector_func}.
	 * @param ctx the parse tree
	 */
	void enterFuncionVector(VGrammar.FuncionVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncionVector}
	 * labeled alternative in {@link VGrammar#vector_func}.
	 * @param ctx the parse tree
	 */
	void exitFuncionVector(VGrammar.FuncionVectorContext ctx);
	/**
	 * Enter a parse tree produced by {@link VGrammar#repeating}.
	 * @param ctx the parse tree
	 */
	void enterRepeating(VGrammar.RepeatingContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#repeating}.
	 * @param ctx the parse tree
	 */
	void exitRepeating(VGrammar.RepeatingContext ctx);
	/**
	 * Enter a parse tree produced by {@link VGrammar#var_type}.
	 * @param ctx the parse tree
	 */
	void enterVar_type(VGrammar.Var_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#var_type}.
	 * @param ctx the parse tree
	 */
	void exitVar_type(VGrammar.Var_typeContext ctx);
	/**
	 * Enter a parse tree produced by {@link VGrammar#type}.
	 * @param ctx the parse tree
	 */
	void enterType(VGrammar.TypeContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#type}.
	 * @param ctx the parse tree
	 */
	void exitType(VGrammar.TypeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorSimple}
	 * labeled alternative in {@link VGrammar#vector_type}.
	 * @param ctx the parse tree
	 */
	void enterVectorSimple(VGrammar.VectorSimpleContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorSimple}
	 * labeled alternative in {@link VGrammar#vector_type}.
	 * @param ctx the parse tree
	 */
	void exitVectorSimple(VGrammar.VectorSimpleContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MatrizDoble}
	 * labeled alternative in {@link VGrammar#vector_type}.
	 * @param ctx the parse tree
	 */
	void enterMatrizDoble(VGrammar.MatrizDobleContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MatrizDoble}
	 * labeled alternative in {@link VGrammar#vector_type}.
	 * @param ctx the parse tree
	 */
	void exitMatrizDoble(VGrammar.MatrizDobleContext ctx);
	/**
	 * Enter a parse tree produced by {@link VGrammar#matrix_type}.
	 * @param ctx the parse tree
	 */
	void enterMatrix_type(VGrammar.Matrix_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#matrix_type}.
	 * @param ctx the parse tree
	 */
	void exitMatrix_type(VGrammar.Matrix_typeContext ctx);
	/**
	 * Enter a parse tree produced by {@link VGrammar#aux_matrix_type}.
	 * @param ctx the parse tree
	 */
	void enterAux_matrix_type(VGrammar.Aux_matrix_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#aux_matrix_type}.
	 * @param ctx the parse tree
	 */
	void exitAux_matrix_type(VGrammar.Aux_matrix_typeContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AssignVectorItem}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void enterAssignVectorItem(VGrammar.AssignVectorItemContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AssignVectorItem}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void exitAssignVectorItem(VGrammar.AssignVectorItemContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignacionDirecta}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionDirecta(VGrammar.AsignacionDirectaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignacionDirecta}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionDirecta(VGrammar.AsignacionDirectaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignacionAritmetica}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionAritmetica(VGrammar.AsignacionAritmeticaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignacionAritmetica}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionAritmetica(VGrammar.AsignacionAritmeticaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AsignacionVector}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void enterAsignacionVector(VGrammar.AsignacionVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AsignacionVector}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void exitAsignacionVector(VGrammar.AsignacionVectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ID_Patron}
	 * labeled alternative in {@link VGrammar#id_pattern}.
	 * @param ctx the parse tree
	 */
	void enterID_Patron(VGrammar.ID_PatronContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ID_Patron}
	 * labeled alternative in {@link VGrammar#id_pattern}.
	 * @param ctx the parse tree
	 */
	void exitID_Patron(VGrammar.ID_PatronContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IntLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void enterIntLiteral(VGrammar.IntLiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IntLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void exitIntLiteral(VGrammar.IntLiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FloatLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void enterFloatLiteral(VGrammar.FloatLiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FloatLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void exitFloatLiteral(VGrammar.FloatLiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StringLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void enterStringLiteral(VGrammar.StringLiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StringLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void exitStringLiteral(VGrammar.StringLiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BoolLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void enterBoolLiteral(VGrammar.BoolLiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BoolLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void exitBoolLiteral(VGrammar.BoolLiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code NilLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void enterNilLiteral(VGrammar.NilLiteralContext ctx);
	/**
	 * Exit a parse tree produced by the {@code NilLiteral}
	 * labeled alternative in {@link VGrammar#literal}.
	 * @param ctx the parse tree
	 */
	void exitNilLiteral(VGrammar.NilLiteralContext ctx);
	/**
	 * Enter a parse tree produced by the {@code LiteralExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterLiteralExp(VGrammar.LiteralExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code LiteralExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitLiteralExp(VGrammar.LiteralExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructMethodExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterStructMethodExp(VGrammar.StructMethodExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructMethodExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitStructMethodExp(VGrammar.StructMethodExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructInitExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterStructInitExp(VGrammar.StructInitExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructInitExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitStructInitExp(VGrammar.StructInitExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code RepeatingExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterRepeatingExp(VGrammar.RepeatingExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code RepeatingExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitRepeatingExp(VGrammar.RepeatingExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterStructExp(VGrammar.StructExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitStructExp(VGrammar.StructExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ExpBinario}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpBinario(VGrammar.ExpBinarioContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ExpBinario}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpBinario(VGrammar.ExpBinarioContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorPropExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterVectorPropExp(VGrammar.VectorPropExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorPropExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitVectorPropExp(VGrammar.VectorPropExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorFuncExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterVectorFuncExp(VGrammar.VectorFuncExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorFuncExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitVectorFuncExp(VGrammar.VectorFuncExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ParentecisExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterParentecisExp(VGrammar.ParentecisExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ParentecisExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitParentecisExp(VGrammar.ParentecisExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ExpUnary}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterExpUnary(VGrammar.ExpUnaryContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ExpUnary}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitExpUnary(VGrammar.ExpUnaryContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IdExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterIdExp(VGrammar.IdExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IdExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitIdExp(VGrammar.IdExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FunctionCallExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterFunctionCallExp(VGrammar.FunctionCallExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FunctionCallExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitFunctionCallExp(VGrammar.FunctionCallExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorItemExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterVectorItemExp(VGrammar.VectorItemExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorItemExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitVectorItemExp(VGrammar.VectorItemExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterVectorExp(VGrammar.VectorExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitVectorExp(VGrammar.VectorExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IFstmt}
	 * labeled alternative in {@link VGrammar#if_stmt}.
	 * @param ctx the parse tree
	 */
	void enterIFstmt(VGrammar.IFstmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IFstmt}
	 * labeled alternative in {@link VGrammar#if_stmt}.
	 * @param ctx the parse tree
	 */
	void exitIFstmt(VGrammar.IFstmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IFcadena}
	 * labeled alternative in {@link VGrammar#if_chain}.
	 * @param ctx the parse tree
	 */
	void enterIFcadena(VGrammar.IFcadenaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IFcadena}
	 * labeled alternative in {@link VGrammar#if_chain}.
	 * @param ctx the parse tree
	 */
	void exitIFcadena(VGrammar.IFcadenaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ElseStmt}
	 * labeled alternative in {@link VGrammar#else_stmt}.
	 * @param ctx the parse tree
	 */
	void enterElseStmt(VGrammar.ElseStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ElseStmt}
	 * labeled alternative in {@link VGrammar#else_stmt}.
	 * @param ctx the parse tree
	 */
	void exitElseStmt(VGrammar.ElseStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SwitchStmt}
	 * labeled alternative in {@link VGrammar#switch_stmt}.
	 * @param ctx the parse tree
	 */
	void enterSwitchStmt(VGrammar.SwitchStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SwitchStmt}
	 * labeled alternative in {@link VGrammar#switch_stmt}.
	 * @param ctx the parse tree
	 */
	void exitSwitchStmt(VGrammar.SwitchStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code SwitchCase}
	 * labeled alternative in {@link VGrammar#switch_case}.
	 * @param ctx the parse tree
	 */
	void enterSwitchCase(VGrammar.SwitchCaseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code SwitchCase}
	 * labeled alternative in {@link VGrammar#switch_case}.
	 * @param ctx the parse tree
	 */
	void exitSwitchCase(VGrammar.SwitchCaseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DefaultCase}
	 * labeled alternative in {@link VGrammar#default_case}.
	 * @param ctx the parse tree
	 */
	void enterDefaultCase(VGrammar.DefaultCaseContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DefaultCase}
	 * labeled alternative in {@link VGrammar#default_case}.
	 * @param ctx the parse tree
	 */
	void exitDefaultCase(VGrammar.DefaultCaseContext ctx);
	/**
	 * Enter a parse tree produced by the {@code WhileStmt}
	 * labeled alternative in {@link VGrammar#while_stmt}.
	 * @param ctx the parse tree
	 */
	void enterWhileStmt(VGrammar.WhileStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code WhileStmt}
	 * labeled alternative in {@link VGrammar#while_stmt}.
	 * @param ctx the parse tree
	 */
	void exitWhileStmt(VGrammar.WhileStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ForStmt}
	 * labeled alternative in {@link VGrammar#for_stmt}.
	 * @param ctx the parse tree
	 */
	void enterForStmt(VGrammar.ForStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ForStmt}
	 * labeled alternative in {@link VGrammar#for_stmt}.
	 * @param ctx the parse tree
	 */
	void exitForStmt(VGrammar.ForStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code RangoNum}
	 * labeled alternative in {@link VGrammar#range}.
	 * @param ctx the parse tree
	 */
	void enterRangoNum(VGrammar.RangoNumContext ctx);
	/**
	 * Exit a parse tree produced by the {@code RangoNum}
	 * labeled alternative in {@link VGrammar#range}.
	 * @param ctx the parse tree
	 */
	void exitRangoNum(VGrammar.RangoNumContext ctx);
	/**
	 * Enter a parse tree produced by the {@code GuardStmt}
	 * labeled alternative in {@link VGrammar#guard_stmt}.
	 * @param ctx the parse tree
	 */
	void enterGuardStmt(VGrammar.GuardStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code GuardStmt}
	 * labeled alternative in {@link VGrammar#guard_stmt}.
	 * @param ctx the parse tree
	 */
	void exitGuardStmt(VGrammar.GuardStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ReturnStmt}
	 * labeled alternative in {@link VGrammar#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void enterReturnStmt(VGrammar.ReturnStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ReturnStmt}
	 * labeled alternative in {@link VGrammar#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void exitReturnStmt(VGrammar.ReturnStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BreakStmt}
	 * labeled alternative in {@link VGrammar#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void enterBreakStmt(VGrammar.BreakStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BreakStmt}
	 * labeled alternative in {@link VGrammar#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void exitBreakStmt(VGrammar.BreakStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ContinueStmt}
	 * labeled alternative in {@link VGrammar#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void enterContinueStmt(VGrammar.ContinueStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ContinueStmt}
	 * labeled alternative in {@link VGrammar#transfer_stmt}.
	 * @param ctx the parse tree
	 */
	void exitContinueStmt(VGrammar.ContinueStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code LlamarFuncion}
	 * labeled alternative in {@link VGrammar#func_call}.
	 * @param ctx the parse tree
	 */
	void enterLlamarFuncion(VGrammar.LlamarFuncionContext ctx);
	/**
	 * Exit a parse tree produced by the {@code LlamarFuncion}
	 * labeled alternative in {@link VGrammar#func_call}.
	 * @param ctx the parse tree
	 */
	void exitLlamarFuncion(VGrammar.LlamarFuncionContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ArgList}
	 * labeled alternative in {@link VGrammar#arg_list}.
	 * @param ctx the parse tree
	 */
	void enterArgList(VGrammar.ArgListContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ArgList}
	 * labeled alternative in {@link VGrammar#arg_list}.
	 * @param ctx the parse tree
	 */
	void exitArgList(VGrammar.ArgListContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncionArg}
	 * labeled alternative in {@link VGrammar#func_arg}.
	 * @param ctx the parse tree
	 */
	void enterFuncionArg(VGrammar.FuncionArgContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncionArg}
	 * labeled alternative in {@link VGrammar#func_arg}.
	 * @param ctx the parse tree
	 */
	void exitFuncionArg(VGrammar.FuncionArgContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncionDeclerada}
	 * labeled alternative in {@link VGrammar#func_dcl}.
	 * @param ctx the parse tree
	 */
	void enterFuncionDeclerada(VGrammar.FuncionDecleradaContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncionDeclerada}
	 * labeled alternative in {@link VGrammar#func_dcl}.
	 * @param ctx the parse tree
	 */
	void exitFuncionDeclerada(VGrammar.FuncionDecleradaContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MetodoStruct}
	 * labeled alternative in {@link VGrammar#func_dcl}.
	 * @param ctx the parse tree
	 */
	void enterMetodoStruct(VGrammar.MetodoStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MetodoStruct}
	 * labeled alternative in {@link VGrammar#func_dcl}.
	 * @param ctx the parse tree
	 */
	void exitMetodoStruct(VGrammar.MetodoStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MethodReceiver}
	 * labeled alternative in {@link VGrammar#method_receiver}.
	 * @param ctx the parse tree
	 */
	void enterMethodReceiver(VGrammar.MethodReceiverContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MethodReceiver}
	 * labeled alternative in {@link VGrammar#method_receiver}.
	 * @param ctx the parse tree
	 */
	void exitMethodReceiver(VGrammar.MethodReceiverContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ParamList}
	 * labeled alternative in {@link VGrammar#param_list}.
	 * @param ctx the parse tree
	 */
	void enterParamList(VGrammar.ParamListContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ParamList}
	 * labeled alternative in {@link VGrammar#param_list}.
	 * @param ctx the parse tree
	 */
	void exitParamList(VGrammar.ParamListContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncParam}
	 * labeled alternative in {@link VGrammar#func_param}.
	 * @param ctx the parse tree
	 */
	void enterFuncParam(VGrammar.FuncParamContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncParam}
	 * labeled alternative in {@link VGrammar#func_param}.
	 * @param ctx the parse tree
	 */
	void exitFuncParam(VGrammar.FuncParamContext ctx);
	/**
	 * Enter a parse tree produced by the {@code DeclararStruct}
	 * labeled alternative in {@link VGrammar#strct_dcl}.
	 * @param ctx the parse tree
	 */
	void enterDeclararStruct(VGrammar.DeclararStructContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DeclararStruct}
	 * labeled alternative in {@link VGrammar#strct_dcl}.
	 * @param ctx the parse tree
	 */
	void exitDeclararStruct(VGrammar.DeclararStructContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructAttr}
	 * labeled alternative in {@link VGrammar#struct_prop}.
	 * @param ctx the parse tree
	 */
	void enterStructAttr(VGrammar.StructAttrContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructAttr}
	 * labeled alternative in {@link VGrammar#struct_prop}.
	 * @param ctx the parse tree
	 */
	void exitStructAttr(VGrammar.StructAttrContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructFunc}
	 * labeled alternative in {@link VGrammar#struct_prop}.
	 * @param ctx the parse tree
	 */
	void enterStructFunc(VGrammar.StructFuncContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructFunc}
	 * labeled alternative in {@link VGrammar#struct_prop}.
	 * @param ctx the parse tree
	 */
	void exitStructFunc(VGrammar.StructFuncContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructVector}
	 * labeled alternative in {@link VGrammar#struct_vector}.
	 * @param ctx the parse tree
	 */
	void enterStructVector(VGrammar.StructVectorContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructVector}
	 * labeled alternative in {@link VGrammar#struct_vector}.
	 * @param ctx the parse tree
	 */
	void exitStructVector(VGrammar.StructVectorContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructInit}
	 * labeled alternative in {@link VGrammar#struct_init}.
	 * @param ctx the parse tree
	 */
	void enterStructInit(VGrammar.StructInitContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructInit}
	 * labeled alternative in {@link VGrammar#struct_init}.
	 * @param ctx the parse tree
	 */
	void exitStructInit(VGrammar.StructInitContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructInitList}
	 * labeled alternative in {@link VGrammar#struct_init_list}.
	 * @param ctx the parse tree
	 */
	void enterStructInitList(VGrammar.StructInitListContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructInitList}
	 * labeled alternative in {@link VGrammar#struct_init_list}.
	 * @param ctx the parse tree
	 */
	void exitStructInitList(VGrammar.StructInitListContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructMethodCall}
	 * labeled alternative in {@link VGrammar#struct_method_call}.
	 * @param ctx the parse tree
	 */
	void enterStructMethodCall(VGrammar.StructMethodCallContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructMethodCall}
	 * labeled alternative in {@link VGrammar#struct_method_call}.
	 * @param ctx the parse tree
	 */
	void exitStructMethodCall(VGrammar.StructMethodCallContext ctx);
	/**
	 * Enter a parse tree produced by the {@code StructInitField}
	 * labeled alternative in {@link VGrammar#struct_init_field}.
	 * @param ctx the parse tree
	 */
	void enterStructInitField(VGrammar.StructInitFieldContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructInitField}
	 * labeled alternative in {@link VGrammar#struct_init_field}.
	 * @param ctx the parse tree
	 */
	void exitStructInitField(VGrammar.StructInitFieldContext ctx);
}