// Generated from /home/daaniieel/Escritorio/PROYECTO_1/Proyecto1/parser/V4LangGrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.tree.ParseTreeListener;

/**
 * This interface defines a complete listener for a parse tree produced by
 * {@link V4LangGrammar}.
 */
public interface V4LangGrammarListener extends ParseTreeListener {
	/**
	 * Enter a parse tree produced by {@link V4LangGrammar#prog}.
	 * @param ctx the parse tree
	 */
	void enterProg(V4LangGrammar.ProgContext ctx);
	/**
	 * Exit a parse tree produced by {@link V4LangGrammar#prog}.
	 * @param ctx the parse tree
	 */
	void exitProg(V4LangGrammar.ProgContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarDeclStmt}
	 * labeled alternative in {@link V4LangGrammar#statement}.
	 * @param ctx the parse tree
	 */
	void enterVarDeclStmt(V4LangGrammar.VarDeclStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarDeclStmt}
	 * labeled alternative in {@link V4LangGrammar#statement}.
	 * @param ctx the parse tree
	 */
	void exitVarDeclStmt(V4LangGrammar.VarDeclStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code ExprStmt}
	 * labeled alternative in {@link V4LangGrammar#statement}.
	 * @param ctx the parse tree
	 */
	void enterExprStmt(V4LangGrammar.ExprStmtContext ctx);
	/**
	 * Exit a parse tree produced by the {@code ExprStmt}
	 * labeled alternative in {@link V4LangGrammar#statement}.
	 * @param ctx the parse tree
	 */
	void exitExprStmt(V4LangGrammar.ExprStmtContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarDeclWithValue}
	 * labeled alternative in {@link V4LangGrammar#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterVarDeclWithValue(V4LangGrammar.VarDeclWithValueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarDeclWithValue}
	 * labeled alternative in {@link V4LangGrammar#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitVarDeclWithValue(V4LangGrammar.VarDeclWithValueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarDeclWithoutValue}
	 * labeled alternative in {@link V4LangGrammar#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterVarDeclWithoutValue(V4LangGrammar.VarDeclWithoutValueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarDeclWithoutValue}
	 * labeled alternative in {@link V4LangGrammar#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitVarDeclWithoutValue(V4LangGrammar.VarDeclWithoutValueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarDeclImmutableWithValue}
	 * labeled alternative in {@link V4LangGrammar#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterVarDeclImmutableWithValue(V4LangGrammar.VarDeclImmutableWithValueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarDeclImmutableWithValue}
	 * labeled alternative in {@link V4LangGrammar#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitVarDeclImmutableWithValue(V4LangGrammar.VarDeclImmutableWithValueContext ctx);
	/**
	 * Enter a parse tree produced by the {@code VarDeclImmutableWithoutValue}
	 * labeled alternative in {@link V4LangGrammar#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void enterVarDeclImmutableWithoutValue(V4LangGrammar.VarDeclImmutableWithoutValueContext ctx);
	/**
	 * Exit a parse tree produced by the {@code VarDeclImmutableWithoutValue}
	 * labeled alternative in {@link V4LangGrammar#varDeclaration}.
	 * @param ctx the parse tree
	 */
	void exitVarDeclImmutableWithoutValue(V4LangGrammar.VarDeclImmutableWithoutValueContext ctx);
	/**
	 * Enter a parse tree produced by {@link V4LangGrammar#type_spec}.
	 * @param ctx the parse tree
	 */
	void enterType_spec(V4LangGrammar.Type_specContext ctx);
	/**
	 * Exit a parse tree produced by {@link V4LangGrammar#type_spec}.
	 * @param ctx the parse tree
	 */
	void exitType_spec(V4LangGrammar.Type_specContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Float}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterFloat(V4LangGrammar.FloatContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Float}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitFloat(V4LangGrammar.FloatContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Variable}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterVariable(V4LangGrammar.VariableContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Variable}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitVariable(V4LangGrammar.VariableContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Bool}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterBool(V4LangGrammar.BoolContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Bool}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitBool(V4LangGrammar.BoolContext ctx);
	/**
	 * Enter a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterMulDiv(V4LangGrammar.MulDivContext ctx);
	/**
	 * Exit a parse tree produced by the {@code MulDiv}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitMulDiv(V4LangGrammar.MulDivContext ctx);
	/**
	 * Enter a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterAddSub(V4LangGrammar.AddSubContext ctx);
	/**
	 * Exit a parse tree produced by the {@code AddSub}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitAddSub(V4LangGrammar.AddSubContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterParens(V4LangGrammar.ParensContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Parens}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitParens(V4LangGrammar.ParensContext ctx);
	/**
	 * Enter a parse tree produced by the {@code String}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterString(V4LangGrammar.StringContext ctx);
	/**
	 * Exit a parse tree produced by the {@code String}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitString(V4LangGrammar.StringContext ctx);
	/**
	 * Enter a parse tree produced by the {@code Int}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void enterInt(V4LangGrammar.IntContext ctx);
	/**
	 * Exit a parse tree produced by the {@code Int}
	 * labeled alternative in {@link V4LangGrammar#expr}.
	 * @param ctx the parse tree
	 */
	void exitInt(V4LangGrammar.IntContext ctx);
}