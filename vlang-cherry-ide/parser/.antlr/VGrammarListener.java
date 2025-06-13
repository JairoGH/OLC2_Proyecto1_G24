// Generated from /home/jairogo/Descargas/OLC2PROYECTO/parser/VGrammar.g4 by ANTLR 4.13.1
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
	 * Enter a parse tree produced by {@link VGrammar#delimiter}.
	 * @param ctx the parse tree
	 */
	void enterDelimiter(VGrammar.DelimiterContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#delimiter}.
	 * @param ctx the parse tree
	 */
	void exitDelimiter(VGrammar.DelimiterContext ctx);
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
	 * Enter a parse tree produced by the {@code TypeValueDecl}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterTypeValueDecl(VGrammar.TypeValueDeclContext ctx);
	/**
	 * Exit a parse tree produced by the {@code TypeValueDecl}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitTypeValueDecl(VGrammar.TypeValueDeclContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ValueDecl}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterValueDecl(VGrammar.ValueDeclContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ValueDecl}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitValueDecl(VGrammar.ValueDeclContext ctx);
	/**
	 * Enter a parse tree produced by the {@code TypeDecl}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void enterTypeDecl(VGrammar.TypeDeclContext ctx);
	/**
	 * Exit a parse tree produced by the {@code TypeDecl}
	 * labeled alternative in {@link VGrammar#decl_stmt}.
	 * @param ctx the parse tree
	 */
	void exitTypeDecl(VGrammar.TypeDeclContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorItemList}
	 * labeled alternative in {@link VGrammar#vector_expr}.
	 * @param ctx the parse tree
	 */
	void enterVectorItemList(VGrammar.VectorItemListContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorItemList}
	 * labeled alternative in {@link VGrammar#vector_expr}.
	 * @param ctx the parse tree
	 */
	void exitVectorItemList(VGrammar.VectorItemListContext ctx);
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
	 * Enter a parse tree produced by the {@code VectorProp}
	 * labeled alternative in {@link VGrammar#vector_prop}.
	 * @param ctx the parse tree
	 */
	void enterVectorProp(VGrammar.VectorPropContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorProp}
	 * labeled alternative in {@link VGrammar#vector_prop}.
	 * @param ctx the parse tree
	 */
	void exitVectorProp(VGrammar.VectorPropContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorFunc}
	 * labeled alternative in {@link VGrammar#vector_func}.
	 * @param ctx the parse tree
	 */
	void enterVectorFunc(VGrammar.VectorFuncContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorFunc}
	 * labeled alternative in {@link VGrammar#vector_func}.
	 * @param ctx the parse tree
	 */
	void exitVectorFunc(VGrammar.VectorFuncContext ctx);
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
	 * Enter a parse tree produced by {@link VGrammar#vector_type}.
	 * @param ctx the parse tree
	 */
	void enterVector_type(VGrammar.Vector_typeContext ctx);
	/**
	 * Exit a parse tree produced by {@link VGrammar#vector_type}.
	 * @param ctx the parse tree
	 */
	void exitVector_type(VGrammar.Vector_typeContext ctx);
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
	 * Enter a parse tree produced by the {@code DirectAssign}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void enterDirectAssign(VGrammar.DirectAssignContext ctx);
	/**
	 * Exit a parse tree produced by the {@code DirectAssign}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void exitDirectAssign(VGrammar.DirectAssignContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ArithmeticAssign}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void enterArithmeticAssign(VGrammar.ArithmeticAssignContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ArithmeticAssign}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void exitArithmeticAssign(VGrammar.ArithmeticAssignContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VectorAssign}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void enterVectorAssign(VGrammar.VectorAssignContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VectorAssign}
	 * labeled alternative in {@link VGrammar#assign_stmt}.
	 * @param ctx the parse tree
	 */
	void exitVectorAssign(VGrammar.VectorAssignContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IdPattern}
	 * labeled alternative in {@link VGrammar#id_pattern}.
	 * @param ctx the parse tree
	 */
	void enterIdPattern(VGrammar.IdPatternContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IdPattern}
	 * labeled alternative in {@link VGrammar#id_pattern}.
	 * @param ctx the parse tree
	 */
	void exitIdPattern(VGrammar.IdPatternContext ctx);
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
	 * Enter a parse tree produced by the {@code StructVectorExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterStructVectorExp(VGrammar.StructVectorExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructVectorExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitStructVectorExp(VGrammar.StructVectorExpContext ctx);
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
	 * Enter a parse tree produced by the {@code ParenExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterParenExp(VGrammar.ParenExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ParenExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitParenExp(VGrammar.ParenExpContext ctx);
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
	 * Enter a parse tree produced by the {@code FuncCallExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterFuncCallExp(VGrammar.FuncCallExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncCallExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitFuncCallExp(VGrammar.FuncCallExpContext ctx);
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
	 * Enter a parse tree produced by the {@code UnaryExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterUnaryExp(VGrammar.UnaryExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code UnaryExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitUnaryExp(VGrammar.UnaryExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code BinaryExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterBinaryExp(VGrammar.BinaryExpContext ctx);
	/**
	 * Exit a parse tree produced by the {@code BinaryExp}
	 * labeled alternative in {@link VGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitBinaryExp(VGrammar.BinaryExpContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IfStmt}
	 * labeled alternative in {@link VGrammar#if_stmt}.
	 * @param ctx the parse tree
	 */
	void enterIfStmt(VGrammar.IfStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IfStmt}
	 * labeled alternative in {@link VGrammar#if_stmt}.
	 * @param ctx the parse tree
	 */
	void exitIfStmt(VGrammar.IfStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code IfChain}
	 * labeled alternative in {@link VGrammar#if_chain}.
	 * @param ctx the parse tree
	 */
	void enterIfChain(VGrammar.IfChainContext ctx);
	/**
	 * Exit a parse tree produced by the {@code IfChain}
	 * labeled alternative in {@link VGrammar#if_chain}.
	 * @param ctx the parse tree
	 */
	void exitIfChain(VGrammar.IfChainContext ctx);
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
	 * Enter a parse tree produced by the {@code NumericRange}
	 * labeled alternative in {@link VGrammar#range}.
	 * @param ctx the parse tree
	 */
	void enterNumericRange(VGrammar.NumericRangeContext ctx);
	/**
	 * Exit a parse tree produced by the {@code NumericRange}
	 * labeled alternative in {@link VGrammar#range}.
	 * @param ctx the parse tree
	 */
	void exitNumericRange(VGrammar.NumericRangeContext ctx);
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
	 * Enter a parse tree produced by the {@code FuncCall}
	 * labeled alternative in {@link VGrammar#func_call}.
	 * @param ctx the parse tree
	 */
	void enterFuncCall(VGrammar.FuncCallContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncCall}
	 * labeled alternative in {@link VGrammar#func_call}.
	 * @param ctx the parse tree
	 */
	void exitFuncCall(VGrammar.FuncCallContext ctx);
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
	 * Enter a parse tree produced by the {@code FuncArg}
	 * labeled alternative in {@link VGrammar#func_arg}.
	 * @param ctx the parse tree
	 */
	void enterFuncArg(VGrammar.FuncArgContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncArg}
	 * labeled alternative in {@link VGrammar#func_arg}.
	 * @param ctx the parse tree
	 */
	void exitFuncArg(VGrammar.FuncArgContext ctx);
	/**
	 * Enter a parse tree produced by the {@code FuncDecl}
	 * labeled alternative in {@link VGrammar#func_dcl}.
	 * @param ctx the parse tree
	 */
	void enterFuncDecl(VGrammar.FuncDeclContext ctx);
	/**
	 * Exit a parse tree produced by the {@code FuncDecl}
	 * labeled alternative in {@link VGrammar#func_dcl}.
	 * @param ctx the parse tree
	 */
	void exitFuncDecl(VGrammar.FuncDeclContext ctx);
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
	 * Enter a parse tree produced by the {@code StructDecl}
	 * labeled alternative in {@link VGrammar#strct_dcl}.
	 * @param ctx the parse tree
	 */
	void enterStructDecl(VGrammar.StructDeclContext ctx);
	/**
	 * Exit a parse tree produced by the {@code StructDecl}
	 * labeled alternative in {@link VGrammar#strct_dcl}.
	 * @param ctx the parse tree
	 */
	void exitStructDecl(VGrammar.StructDeclContext ctx);
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
}