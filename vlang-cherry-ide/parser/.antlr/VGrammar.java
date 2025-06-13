// Generated from /home/jairogo/Descargas/OLC2PROYECTO/parser/VGrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class VGrammar extends Parser {
	static { RuntimeMetaData.checkVersion("4.13.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		WS=1, COMMENT=2, MULTILINE_COMMENT=3, SEMICOLON=4, RW_MUT=5, OP_ASSIGN=6, 
		FUNC_KW=7, STRUCT_KW=8, IF_KW=9, ELSE_KW=10, SWITCH_KW=11, CASE_KW=12, 
		DEFAULT_KW=13, FOR_KW=14, WHILE_KW=15, BREAK_KW=16, CONTINUE_KW=17, RETURN_KW=18, 
		GUARD_KW=19, INOUT_KW=20, MUTATING_KW=21, IN_KW=22, INTEGER_LITERAL=23, 
		FLOAT_LITERAL=24, STRING_LITERAL=25, BOOL_LITERAL=26, NIL_LITERAL=27, 
		RW_INT=28, RW_FLOAT64=29, RW_STRING=30, RW_BOOL=31, RW_RUNE=32, ID=33, 
		PLUS=34, MINUS=35, MULT=36, DIV=37, MOD=38, EQUALS=39, PLUS_EQUALS=40, 
		MINUS_EQUALS=41, EQUALS_EQUALS=42, NOT_EQUALS=43, LESS_THAN=44, LESS_THAN_OR_EQUAL=45, 
		GREATER_THAN=46, GREATER_THAN_OR_EQUAL=47, AND=48, OR=49, NOT=50, LPAREN=51, 
		RPAREN=52, LBRACE=53, RBRACE=54, LBRACK=55, RBRACK=56, COMMA=57, DOT=58, 
		COLON=59, ARROW=60, INTERROGATION=61, ANPERSAND=62;
	public static final int
		RULE_program = 0, RULE_stmt = 1, RULE_decl_stmt = 2, RULE_vector_expr = 3, 
		RULE_vector_item = 4, RULE_vector_prop = 5, RULE_vector_func = 6, RULE_repeating = 7, 
		RULE_var_type = 8, RULE_type = 9, RULE_vector_type = 10, RULE_matrix_type = 11, 
		RULE_aux_matrix_type = 12, RULE_assign_stmt = 13, RULE_id_pattern = 14, 
		RULE_literal = 15, RULE_expr = 16, RULE_if_stmt = 17, RULE_if_chain = 18, 
		RULE_else_stmt = 19, RULE_switch_stmt = 20, RULE_switch_case = 21, RULE_default_case = 22, 
		RULE_while_stmt = 23, RULE_for_stmt = 24, RULE_range = 25, RULE_guard_stmt = 26, 
		RULE_transfer_stmt = 27, RULE_func_call = 28, RULE_arg_list = 29, RULE_func_arg = 30, 
		RULE_func_dcl = 31, RULE_param_list = 32, RULE_func_param = 33, RULE_strct_dcl = 34, 
		RULE_struct_prop = 35, RULE_struct_vector = 36;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "stmt", "decl_stmt", "vector_expr", "vector_item", "vector_prop", 
			"vector_func", "repeating", "var_type", "type", "vector_type", "matrix_type", 
			"aux_matrix_type", "assign_stmt", "id_pattern", "literal", "expr", "if_stmt", 
			"if_chain", "else_stmt", "switch_stmt", "switch_case", "default_case", 
			"while_stmt", "for_stmt", "range", "guard_stmt", "transfer_stmt", "func_call", 
			"arg_list", "func_arg", "func_dcl", "param_list", "func_param", "strct_dcl", 
			"struct_prop", "struct_vector"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, "';'", "'mut'", "':='", "'func'", "'struct'", 
			"'if'", "'else'", "'switch'", "'case'", "'default'", "'for'", "'while'", 
			"'break'", "'continue'", "'return'", "'guard'", "'inout'", "'mutating'", 
			"'in'", null, null, null, null, "'nil'", "'int'", null, "'string'", "'bool'", 
			"'rune'", null, "'+'", "'-'", "'*'", "'/'", "'%'", "'='", "'+='", "'-='", 
			"'=='", "'!='", "'<'", "'<='", "'>'", "'>='", "'&&'", "'||'", "'!'", 
			"'('", "')'", "'{'", "'}'", "'['", "']'", "','", "'.'", "':'", "'->'", 
			"'?'", "'&'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WS", "COMMENT", "MULTILINE_COMMENT", "SEMICOLON", "RW_MUT", "OP_ASSIGN", 
			"FUNC_KW", "STRUCT_KW", "IF_KW", "ELSE_KW", "SWITCH_KW", "CASE_KW", "DEFAULT_KW", 
			"FOR_KW", "WHILE_KW", "BREAK_KW", "CONTINUE_KW", "RETURN_KW", "GUARD_KW", 
			"INOUT_KW", "MUTATING_KW", "IN_KW", "INTEGER_LITERAL", "FLOAT_LITERAL", 
			"STRING_LITERAL", "BOOL_LITERAL", "NIL_LITERAL", "RW_INT", "RW_FLOAT64", 
			"RW_STRING", "RW_BOOL", "RW_RUNE", "ID", "PLUS", "MINUS", "MULT", "DIV", 
			"MOD", "EQUALS", "PLUS_EQUALS", "MINUS_EQUALS", "EQUALS_EQUALS", "NOT_EQUALS", 
			"LESS_THAN", "LESS_THAN_OR_EQUAL", "GREATER_THAN", "GREATER_THAN_OR_EQUAL", 
			"AND", "OR", "NOT", "LPAREN", "RPAREN", "LBRACE", "RBRACE", "LBRACK", 
			"RBRACK", "COMMA", "DOT", "COLON", "ARROW", "INTERROGATION", "ANPERSAND"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "VGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public VGrammar(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramContext extends ParserRuleContext {
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public TerminalNode EOF() { return getToken(VGrammar.EOF, 0); }
		public ProgramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_program; }
	}

	public final ProgramContext program() throws RecognitionException {
		ProgramContext _localctx = new ProgramContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_program);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(77);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8590969760L) != 0)) {
				{
				{
				setState(74);
				stmt();
				}
				}
				setState(79);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(81);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
			case 1:
				{
				setState(80);
				match(EOF);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class StmtContext extends ParserRuleContext {
		public Decl_stmtContext decl_stmt() {
			return getRuleContext(Decl_stmtContext.class,0);
		}
		public Assign_stmtContext assign_stmt() {
			return getRuleContext(Assign_stmtContext.class,0);
		}
		public Transfer_stmtContext transfer_stmt() {
			return getRuleContext(Transfer_stmtContext.class,0);
		}
		public If_stmtContext if_stmt() {
			return getRuleContext(If_stmtContext.class,0);
		}
		public Switch_stmtContext switch_stmt() {
			return getRuleContext(Switch_stmtContext.class,0);
		}
		public While_stmtContext while_stmt() {
			return getRuleContext(While_stmtContext.class,0);
		}
		public For_stmtContext for_stmt() {
			return getRuleContext(For_stmtContext.class,0);
		}
		public Guard_stmtContext guard_stmt() {
			return getRuleContext(Guard_stmtContext.class,0);
		}
		public Func_callContext func_call() {
			return getRuleContext(Func_callContext.class,0);
		}
		public Vector_funcContext vector_func() {
			return getRuleContext(Vector_funcContext.class,0);
		}
		public Func_dclContext func_dcl() {
			return getRuleContext(Func_dclContext.class,0);
		}
		public Strct_dclContext strct_dcl() {
			return getRuleContext(Strct_dclContext.class,0);
		}
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_stmt);
		try {
			setState(95);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(83);
				decl_stmt();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(84);
				assign_stmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(85);
				transfer_stmt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(86);
				if_stmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(87);
				switch_stmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(88);
				while_stmt();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(89);
				for_stmt();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(90);
				guard_stmt();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(91);
				func_call();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(92);
				vector_func();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(93);
				func_dcl();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(94);
				strct_dcl();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Decl_stmtContext extends ParserRuleContext {
		public Decl_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decl_stmt; }
	 
		public Decl_stmtContext() { }
		public void copyFrom(Decl_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararInferenciaContext extends Decl_stmtContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_ASSIGN() { return getToken(VGrammar.OP_ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclararInferenciaContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclaraTipoValorContext extends Decl_stmtContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode EQUALS() { return getToken(VGrammar.EQUALS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclaraTipoValorContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararSinMutValorContext extends Decl_stmtContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode EQUALS() { return getToken(VGrammar.EQUALS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclararSinMutValorContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararInferenciaMutContext extends Decl_stmtContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_ASSIGN() { return getToken(VGrammar.OP_ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclararInferenciaMutContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararTipoContext extends Decl_stmtContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public DeclararTipoContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararSinMutTipoContext extends Decl_stmtContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public DeclararSinMutTipoContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Decl_stmtContext decl_stmt() throws RecognitionException {
		Decl_stmtContext _localctx = new Decl_stmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_decl_stmt);
		try {
			setState(120);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				_localctx = new DeclararInferenciaMutContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(97);
				match(RW_MUT);
				setState(98);
				match(ID);
				setState(99);
				match(OP_ASSIGN);
				setState(100);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DeclararInferenciaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(101);
				match(ID);
				setState(102);
				match(OP_ASSIGN);
				setState(103);
				expr(0);
				}
				break;
			case 3:
				_localctx = new DeclaraTipoValorContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(104);
				match(RW_MUT);
				setState(105);
				match(ID);
				setState(106);
				type();
				setState(107);
				match(EQUALS);
				setState(108);
				expr(0);
				}
				break;
			case 4:
				_localctx = new DeclararTipoContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(110);
				match(RW_MUT);
				setState(111);
				match(ID);
				setState(112);
				type();
				}
				break;
			case 5:
				_localctx = new DeclararSinMutValorContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(113);
				match(ID);
				setState(114);
				type();
				setState(115);
				match(EQUALS);
				setState(116);
				expr(0);
				}
				break;
			case 6:
				_localctx = new DeclararSinMutTipoContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(118);
				match(ID);
				setState(119);
				type();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Vector_exprContext extends ParserRuleContext {
		public Vector_exprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vector_expr; }
	 
		public Vector_exprContext() { }
		public void copyFrom(Vector_exprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ListaItemsVectorContext extends Vector_exprContext {
		public TerminalNode LBRACK() { return getToken(VGrammar.LBRACK, 0); }
		public TerminalNode RBRACK() { return getToken(VGrammar.RBRACK, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VGrammar.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VGrammar.COMMA, i);
		}
		public ListaItemsVectorContext(Vector_exprContext ctx) { copyFrom(ctx); }
	}

	public final Vector_exprContext vector_expr() throws RecognitionException {
		Vector_exprContext _localctx = new Vector_exprContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_vector_expr);
		int _la;
		try {
			_localctx = new ListaItemsVectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(122);
			match(LBRACK);
			setState(131);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 39406539949211648L) != 0)) {
				{
				setState(123);
				expr(0);
				setState(128);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMMA) {
					{
					{
					setState(124);
					match(COMMA);
					setState(125);
					expr(0);
					}
					}
					setState(130);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(133);
			match(RBRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Vector_itemContext extends ParserRuleContext {
		public Vector_itemContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vector_item; }
	 
		public Vector_itemContext() { }
		public void copyFrom(Vector_itemContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorItemContext extends Vector_itemContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public List<TerminalNode> LBRACK() { return getTokens(VGrammar.LBRACK); }
		public TerminalNode LBRACK(int i) {
			return getToken(VGrammar.LBRACK, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> RBRACK() { return getTokens(VGrammar.RBRACK); }
		public TerminalNode RBRACK(int i) {
			return getToken(VGrammar.RBRACK, i);
		}
		public VectorItemContext(Vector_itemContext ctx) { copyFrom(ctx); }
	}

	public final Vector_itemContext vector_item() throws RecognitionException {
		Vector_itemContext _localctx = new Vector_itemContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_vector_item);
		try {
			int _alt;
			_localctx = new VectorItemContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(135);
			id_pattern();
			setState(140); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(136);
					match(LBRACK);
					setState(137);
					expr(0);
					setState(138);
					match(RBRACK);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(142); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			} while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Vector_propContext extends ParserRuleContext {
		public Vector_propContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vector_prop; }
	 
		public Vector_propContext() { }
		public void copyFrom(Vector_propContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PropVectorContext extends Vector_propContext {
		public Vector_itemContext vector_item() {
			return getRuleContext(Vector_itemContext.class,0);
		}
		public TerminalNode DOT() { return getToken(VGrammar.DOT, 0); }
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public PropVectorContext(Vector_propContext ctx) { copyFrom(ctx); }
	}

	public final Vector_propContext vector_prop() throws RecognitionException {
		Vector_propContext _localctx = new Vector_propContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_vector_prop);
		try {
			_localctx = new PropVectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(144);
			vector_item();
			setState(145);
			match(DOT);
			setState(146);
			id_pattern();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Vector_funcContext extends ParserRuleContext {
		public Vector_funcContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vector_func; }
	 
		public Vector_funcContext() { }
		public void copyFrom(Vector_funcContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionVectorContext extends Vector_funcContext {
		public Vector_itemContext vector_item() {
			return getRuleContext(Vector_itemContext.class,0);
		}
		public TerminalNode DOT() { return getToken(VGrammar.DOT, 0); }
		public Func_callContext func_call() {
			return getRuleContext(Func_callContext.class,0);
		}
		public FuncionVectorContext(Vector_funcContext ctx) { copyFrom(ctx); }
	}

	public final Vector_funcContext vector_func() throws RecognitionException {
		Vector_funcContext _localctx = new Vector_funcContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_vector_func);
		try {
			_localctx = new FuncionVectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(148);
			vector_item();
			setState(149);
			match(DOT);
			setState(150);
			func_call();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RepeatingContext extends ParserRuleContext {
		public TerminalNode LPAREN() { return getToken(VGrammar.LPAREN, 0); }
		public List<TerminalNode> ID() { return getTokens(VGrammar.ID); }
		public TerminalNode ID(int i) {
			return getToken(VGrammar.ID, i);
		}
		public List<TerminalNode> COLON() { return getTokens(VGrammar.COLON); }
		public TerminalNode COLON(int i) {
			return getToken(VGrammar.COLON, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode COMMA() { return getToken(VGrammar.COMMA, 0); }
		public TerminalNode RPAREN() { return getToken(VGrammar.RPAREN, 0); }
		public Vector_typeContext vector_type() {
			return getRuleContext(Vector_typeContext.class,0);
		}
		public Matrix_typeContext matrix_type() {
			return getRuleContext(Matrix_typeContext.class,0);
		}
		public RepeatingContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_repeating; }
	}

	public final RepeatingContext repeating() throws RecognitionException {
		RepeatingContext _localctx = new RepeatingContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_repeating);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(154);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				{
				setState(152);
				vector_type();
				}
				break;
			case 2:
				{
				setState(153);
				matrix_type();
				}
				break;
			}
			setState(156);
			match(LPAREN);
			setState(157);
			match(ID);
			setState(158);
			match(COLON);
			setState(159);
			expr(0);
			setState(160);
			match(COMMA);
			setState(161);
			match(ID);
			setState(162);
			match(COLON);
			setState(163);
			expr(0);
			setState(164);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Var_typeContext extends ParserRuleContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public Var_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_var_type; }
	}

	public final Var_typeContext var_type() throws RecognitionException {
		Var_typeContext _localctx = new Var_typeContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_var_type);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(166);
			match(RW_MUT);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class TypeContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_INT() { return getToken(VGrammar.RW_INT, 0); }
		public TerminalNode RW_FLOAT64() { return getToken(VGrammar.RW_FLOAT64, 0); }
		public TerminalNode RW_STRING() { return getToken(VGrammar.RW_STRING, 0); }
		public TerminalNode RW_BOOL() { return getToken(VGrammar.RW_BOOL, 0); }
		public Vector_typeContext vector_type() {
			return getRuleContext(Vector_typeContext.class,0);
		}
		public Matrix_typeContext matrix_type() {
			return getRuleContext(Matrix_typeContext.class,0);
		}
		public TypeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_type; }
	}

	public final TypeContext type() throws RecognitionException {
		TypeContext _localctx = new TypeContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_type);
		try {
			setState(175);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(168);
				match(ID);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(169);
				match(RW_INT);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(170);
				match(RW_FLOAT64);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(171);
				match(RW_STRING);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(172);
				match(RW_BOOL);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(173);
				vector_type();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(174);
				matrix_type();
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Vector_typeContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(VGrammar.LBRACK, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RBRACK() { return getToken(VGrammar.RBRACK, 0); }
		public Vector_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vector_type; }
	}

	public final Vector_typeContext vector_type() throws RecognitionException {
		Vector_typeContext _localctx = new Vector_typeContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_vector_type);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(177);
			match(LBRACK);
			setState(178);
			match(ID);
			setState(179);
			match(RBRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Matrix_typeContext extends ParserRuleContext {
		public Aux_matrix_typeContext aux_matrix_type() {
			return getRuleContext(Aux_matrix_typeContext.class,0);
		}
		public List<TerminalNode> LBRACK() { return getTokens(VGrammar.LBRACK); }
		public TerminalNode LBRACK(int i) {
			return getToken(VGrammar.LBRACK, i);
		}
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public List<TerminalNode> RBRACK() { return getTokens(VGrammar.RBRACK); }
		public TerminalNode RBRACK(int i) {
			return getToken(VGrammar.RBRACK, i);
		}
		public Matrix_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matrix_type; }
	}

	public final Matrix_typeContext matrix_type() throws RecognitionException {
		Matrix_typeContext _localctx = new Matrix_typeContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_matrix_type);
		try {
			setState(187);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(181);
				aux_matrix_type();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(182);
				match(LBRACK);
				setState(183);
				match(LBRACK);
				setState(184);
				match(ID);
				setState(185);
				match(RBRACK);
				setState(186);
				match(RBRACK);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Aux_matrix_typeContext extends ParserRuleContext {
		public TerminalNode LBRACK() { return getToken(VGrammar.LBRACK, 0); }
		public Matrix_typeContext matrix_type() {
			return getRuleContext(Matrix_typeContext.class,0);
		}
		public TerminalNode RBRACK() { return getToken(VGrammar.RBRACK, 0); }
		public Aux_matrix_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_aux_matrix_type; }
	}

	public final Aux_matrix_typeContext aux_matrix_type() throws RecognitionException {
		Aux_matrix_typeContext _localctx = new Aux_matrix_typeContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_aux_matrix_type);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(189);
			match(LBRACK);
			setState(190);
			matrix_type();
			setState(191);
			match(RBRACK);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Assign_stmtContext extends ParserRuleContext {
		public Assign_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assign_stmt; }
	 
		public Assign_stmtContext() { }
		public void copyFrom(Assign_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionVectorContext extends Assign_stmtContext {
		public Token op;
		public Vector_itemContext vector_item() {
			return getRuleContext(Vector_itemContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PLUS_EQUALS() { return getToken(VGrammar.PLUS_EQUALS, 0); }
		public TerminalNode MINUS_EQUALS() { return getToken(VGrammar.MINUS_EQUALS, 0); }
		public TerminalNode EQUALS() { return getToken(VGrammar.EQUALS, 0); }
		public AsignacionVectorContext(Assign_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionDirectaContext extends Assign_stmtContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public TerminalNode EQUALS() { return getToken(VGrammar.EQUALS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AsignacionDirectaContext(Assign_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionAritmeticaContext extends Assign_stmtContext {
		public Token op;
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PLUS_EQUALS() { return getToken(VGrammar.PLUS_EQUALS, 0); }
		public TerminalNode MINUS_EQUALS() { return getToken(VGrammar.MINUS_EQUALS, 0); }
		public AsignacionAritmeticaContext(Assign_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Assign_stmtContext assign_stmt() throws RecognitionException {
		Assign_stmtContext _localctx = new Assign_stmtContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_assign_stmt);
		int _la;
		try {
			setState(205);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
			case 1:
				_localctx = new AsignacionDirectaContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(193);
				id_pattern();
				setState(194);
				match(EQUALS);
				setState(195);
				expr(0);
				}
				break;
			case 2:
				_localctx = new AsignacionAritmeticaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(197);
				id_pattern();
				setState(198);
				((AsignacionAritmeticaContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==PLUS_EQUALS || _la==MINUS_EQUALS) ) {
					((AsignacionAritmeticaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(199);
				expr(0);
				}
				break;
			case 3:
				_localctx = new AsignacionVectorContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(201);
				vector_item();
				setState(202);
				((AsignacionVectorContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 3848290697216L) != 0)) ) {
					((AsignacionVectorContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(203);
				expr(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Id_patternContext extends ParserRuleContext {
		public Id_patternContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_id_pattern; }
	 
		public Id_patternContext() { }
		public void copyFrom(Id_patternContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ID_PatronContext extends Id_patternContext {
		public List<TerminalNode> ID() { return getTokens(VGrammar.ID); }
		public TerminalNode ID(int i) {
			return getToken(VGrammar.ID, i);
		}
		public List<TerminalNode> DOT() { return getTokens(VGrammar.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(VGrammar.DOT, i);
		}
		public ID_PatronContext(Id_patternContext ctx) { copyFrom(ctx); }
	}

	public final Id_patternContext id_pattern() throws RecognitionException {
		Id_patternContext _localctx = new Id_patternContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_id_pattern);
		try {
			int _alt;
			_localctx = new ID_PatronContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(207);
			match(ID);
			setState(212);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(208);
					match(DOT);
					setState(209);
					match(ID);
					}
					} 
				}
				setState(214);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class LiteralContext extends ParserRuleContext {
		public LiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_literal; }
	 
		public LiteralContext() { }
		public void copyFrom(LiteralContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StringLiteralContext extends LiteralContext {
		public TerminalNode STRING_LITERAL() { return getToken(VGrammar.STRING_LITERAL, 0); }
		public StringLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BoolLiteralContext extends LiteralContext {
		public TerminalNode BOOL_LITERAL() { return getToken(VGrammar.BOOL_LITERAL, 0); }
		public BoolLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FloatLiteralContext extends LiteralContext {
		public TerminalNode FLOAT_LITERAL() { return getToken(VGrammar.FLOAT_LITERAL, 0); }
		public FloatLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class NilLiteralContext extends LiteralContext {
		public TerminalNode NIL_LITERAL() { return getToken(VGrammar.NIL_LITERAL, 0); }
		public NilLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntLiteralContext extends LiteralContext {
		public TerminalNode INTEGER_LITERAL() { return getToken(VGrammar.INTEGER_LITERAL, 0); }
		public IntLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_literal);
		try {
			setState(220);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INTEGER_LITERAL:
				_localctx = new IntLiteralContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(215);
				match(INTEGER_LITERAL);
				}
				break;
			case FLOAT_LITERAL:
				_localctx = new FloatLiteralContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(216);
				match(FLOAT_LITERAL);
				}
				break;
			case STRING_LITERAL:
				_localctx = new StringLiteralContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(217);
				match(STRING_LITERAL);
				}
				break;
			case BOOL_LITERAL:
				_localctx = new BoolLiteralContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(218);
				match(BOOL_LITERAL);
				}
				break;
			case NIL_LITERAL:
				_localctx = new NilLiteralContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(219);
				match(NIL_LITERAL);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ExprContext extends ParserRuleContext {
		public ExprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr; }
	 
		public ExprContext() { }
		public void copyFrom(ExprContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LiteralExpContext extends ExprContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public LiteralExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IdExpContext extends ExprContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public IdExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionCallExpContext extends ExprContext {
		public Func_callContext func_call() {
			return getRuleContext(Func_callContext.class,0);
		}
		public FunctionCallExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RepeatingExpContext extends ExprContext {
		public RepeatingContext repeating() {
			return getRuleContext(RepeatingContext.class,0);
		}
		public RepeatingExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructExpContext extends ExprContext {
		public Struct_vectorContext struct_vector() {
			return getRuleContext(Struct_vectorContext.class,0);
		}
		public StructExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpBinarioContext extends ExprContext {
		public ExprContext left;
		public Token op;
		public ExprContext right;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode MULT() { return getToken(VGrammar.MULT, 0); }
		public TerminalNode DIV() { return getToken(VGrammar.DIV, 0); }
		public TerminalNode MOD() { return getToken(VGrammar.MOD, 0); }
		public TerminalNode PLUS() { return getToken(VGrammar.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(VGrammar.MINUS, 0); }
		public TerminalNode LESS_THAN() { return getToken(VGrammar.LESS_THAN, 0); }
		public TerminalNode LESS_THAN_OR_EQUAL() { return getToken(VGrammar.LESS_THAN_OR_EQUAL, 0); }
		public TerminalNode GREATER_THAN() { return getToken(VGrammar.GREATER_THAN, 0); }
		public TerminalNode GREATER_THAN_OR_EQUAL() { return getToken(VGrammar.GREATER_THAN_OR_EQUAL, 0); }
		public TerminalNode EQUALS_EQUALS() { return getToken(VGrammar.EQUALS_EQUALS, 0); }
		public TerminalNode NOT_EQUALS() { return getToken(VGrammar.NOT_EQUALS, 0); }
		public TerminalNode AND() { return getToken(VGrammar.AND, 0); }
		public TerminalNode OR() { return getToken(VGrammar.OR, 0); }
		public ExpBinarioContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorPropExpContext extends ExprContext {
		public Vector_propContext vector_prop() {
			return getRuleContext(Vector_propContext.class,0);
		}
		public VectorPropExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorItemExpContext extends ExprContext {
		public Vector_itemContext vector_item() {
			return getRuleContext(Vector_itemContext.class,0);
		}
		public VectorItemExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorFuncExpContext extends ExprContext {
		public Vector_funcContext vector_func() {
			return getRuleContext(Vector_funcContext.class,0);
		}
		public VectorFuncExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorExpContext extends ExprContext {
		public Vector_exprContext vector_expr() {
			return getRuleContext(Vector_exprContext.class,0);
		}
		public VectorExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParentecisExpContext extends ExprContext {
		public TerminalNode LPAREN() { return getToken(VGrammar.LPAREN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode RPAREN() { return getToken(VGrammar.RPAREN, 0); }
		public ParentecisExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpUnaryContext extends ExprContext {
		public Token op;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode NOT() { return getToken(VGrammar.NOT, 0); }
		public TerminalNode MINUS() { return getToken(VGrammar.MINUS, 0); }
		public ExpUnaryContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 32;
		enterRecursionRule(_localctx, 32, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(238);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				{
				_localctx = new ParentecisExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(223);
				match(LPAREN);
				setState(224);
				expr(0);
				setState(225);
				match(RPAREN);
				}
				break;
			case 2:
				{
				_localctx = new FunctionCallExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(227);
				func_call();
				}
				break;
			case 3:
				{
				_localctx = new IdExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(228);
				id_pattern();
				}
				break;
			case 4:
				{
				_localctx = new VectorItemExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(229);
				vector_item();
				}
				break;
			case 5:
				{
				_localctx = new VectorPropExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(230);
				vector_prop();
				}
				break;
			case 6:
				{
				_localctx = new VectorFuncExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(231);
				vector_func();
				}
				break;
			case 7:
				{
				_localctx = new LiteralExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(232);
				literal();
				}
				break;
			case 8:
				{
				_localctx = new VectorExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(233);
				vector_expr();
				}
				break;
			case 9:
				{
				_localctx = new RepeatingExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(234);
				repeating();
				}
				break;
			case 10:
				{
				_localctx = new StructExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(235);
				struct_vector();
				}
				break;
			case 11:
				{
				_localctx = new ExpUnaryContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(236);
				((ExpUnaryContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==MINUS || _la==NOT) ) {
					((ExpUnaryContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(237);
				expr(7);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(260);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(258);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
					case 1:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(240);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(241);
						((ExpBinarioContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 481036337152L) != 0)) ) {
							((ExpBinarioContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(242);
						((ExpBinarioContext)_localctx).right = expr(7);
						}
						break;
					case 2:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(243);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(244);
						((ExpBinarioContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==PLUS || _la==MINUS) ) {
							((ExpBinarioContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(245);
						((ExpBinarioContext)_localctx).right = expr(6);
						}
						break;
					case 3:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(246);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(247);
						((ExpBinarioContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 263882790666240L) != 0)) ) {
							((ExpBinarioContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(248);
						((ExpBinarioContext)_localctx).right = expr(5);
						}
						break;
					case 4:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(249);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(250);
						((ExpBinarioContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==EQUALS_EQUALS || _la==NOT_EQUALS) ) {
							((ExpBinarioContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(251);
						((ExpBinarioContext)_localctx).right = expr(4);
						}
						break;
					case 5:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(252);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(253);
						((ExpBinarioContext)_localctx).op = match(AND);
						setState(254);
						((ExpBinarioContext)_localctx).right = expr(3);
						}
						break;
					case 6:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(255);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(256);
						((ExpBinarioContext)_localctx).op = match(OR);
						setState(257);
						((ExpBinarioContext)_localctx).right = expr(2);
						}
						break;
					}
					} 
				}
				setState(262);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,15,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_stmtContext extends ParserRuleContext {
		public If_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_stmt; }
	 
		public If_stmtContext() { }
		public void copyFrom(If_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IFstmtContext extends If_stmtContext {
		public List<If_chainContext> if_chain() {
			return getRuleContexts(If_chainContext.class);
		}
		public If_chainContext if_chain(int i) {
			return getRuleContext(If_chainContext.class,i);
		}
		public List<TerminalNode> ELSE_KW() { return getTokens(VGrammar.ELSE_KW); }
		public TerminalNode ELSE_KW(int i) {
			return getToken(VGrammar.ELSE_KW, i);
		}
		public Else_stmtContext else_stmt() {
			return getRuleContext(Else_stmtContext.class,0);
		}
		public IFstmtContext(If_stmtContext ctx) { copyFrom(ctx); }
	}

	public final If_stmtContext if_stmt() throws RecognitionException {
		If_stmtContext _localctx = new If_stmtContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_if_stmt);
		int _la;
		try {
			int _alt;
			_localctx = new IFstmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(263);
			if_chain();
			setState(268);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(264);
					match(ELSE_KW);
					setState(265);
					if_chain();
					}
					} 
				}
				setState(270);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			}
			setState(272);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ELSE_KW) {
				{
				setState(271);
				else_stmt();
				}
			}

			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class If_chainContext extends ParserRuleContext {
		public If_chainContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_chain; }
	 
		public If_chainContext() { }
		public void copyFrom(If_chainContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IFcadenaContext extends If_chainContext {
		public TerminalNode IF_KW() { return getToken(VGrammar.IF_KW, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(VGrammar.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VGrammar.RBRACE, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public IFcadenaContext(If_chainContext ctx) { copyFrom(ctx); }
	}

	public final If_chainContext if_chain() throws RecognitionException {
		If_chainContext _localctx = new If_chainContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_if_chain);
		int _la;
		try {
			_localctx = new IFcadenaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(274);
			match(IF_KW);
			setState(275);
			expr(0);
			setState(276);
			match(LBRACE);
			setState(280);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8590969760L) != 0)) {
				{
				{
				setState(277);
				stmt();
				}
				}
				setState(282);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(283);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Else_stmtContext extends ParserRuleContext {
		public Else_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_else_stmt; }
	 
		public Else_stmtContext() { }
		public void copyFrom(Else_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ElseStmtContext extends Else_stmtContext {
		public TerminalNode ELSE_KW() { return getToken(VGrammar.ELSE_KW, 0); }
		public TerminalNode LBRACE() { return getToken(VGrammar.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VGrammar.RBRACE, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ElseStmtContext(Else_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Else_stmtContext else_stmt() throws RecognitionException {
		Else_stmtContext _localctx = new Else_stmtContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_else_stmt);
		int _la;
		try {
			_localctx = new ElseStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(285);
			match(ELSE_KW);
			setState(286);
			match(LBRACE);
			setState(290);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8590969760L) != 0)) {
				{
				{
				setState(287);
				stmt();
				}
				}
				setState(292);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(293);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Switch_stmtContext extends ParserRuleContext {
		public Switch_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_stmt; }
	 
		public Switch_stmtContext() { }
		public void copyFrom(Switch_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchStmtContext extends Switch_stmtContext {
		public TerminalNode SWITCH_KW() { return getToken(VGrammar.SWITCH_KW, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(VGrammar.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VGrammar.RBRACE, 0); }
		public List<Switch_caseContext> switch_case() {
			return getRuleContexts(Switch_caseContext.class);
		}
		public Switch_caseContext switch_case(int i) {
			return getRuleContext(Switch_caseContext.class,i);
		}
		public Default_caseContext default_case() {
			return getRuleContext(Default_caseContext.class,0);
		}
		public SwitchStmtContext(Switch_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Switch_stmtContext switch_stmt() throws RecognitionException {
		Switch_stmtContext _localctx = new Switch_stmtContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_switch_stmt);
		int _la;
		try {
			_localctx = new SwitchStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(295);
			match(SWITCH_KW);
			setState(296);
			expr(0);
			setState(297);
			match(LBRACE);
			setState(301);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==CASE_KW) {
				{
				{
				setState(298);
				switch_case();
				}
				}
				setState(303);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(305);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DEFAULT_KW) {
				{
				setState(304);
				default_case();
				}
			}

			setState(307);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Switch_caseContext extends ParserRuleContext {
		public Switch_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_case; }
	 
		public Switch_caseContext() { }
		public void copyFrom(Switch_caseContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class SwitchCaseContext extends Switch_caseContext {
		public TerminalNode CASE_KW() { return getToken(VGrammar.CASE_KW, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode COLON() { return getToken(VGrammar.COLON, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public SwitchCaseContext(Switch_caseContext ctx) { copyFrom(ctx); }
	}

	public final Switch_caseContext switch_case() throws RecognitionException {
		Switch_caseContext _localctx = new Switch_caseContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_switch_case);
		int _la;
		try {
			_localctx = new SwitchCaseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(309);
			match(CASE_KW);
			setState(310);
			expr(0);
			setState(311);
			match(COLON);
			setState(315);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8590969760L) != 0)) {
				{
				{
				setState(312);
				stmt();
				}
				}
				setState(317);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Default_caseContext extends ParserRuleContext {
		public Default_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default_case; }
	 
		public Default_caseContext() { }
		public void copyFrom(Default_caseContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DefaultCaseContext extends Default_caseContext {
		public TerminalNode DEFAULT_KW() { return getToken(VGrammar.DEFAULT_KW, 0); }
		public TerminalNode COLON() { return getToken(VGrammar.COLON, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public DefaultCaseContext(Default_caseContext ctx) { copyFrom(ctx); }
	}

	public final Default_caseContext default_case() throws RecognitionException {
		Default_caseContext _localctx = new Default_caseContext(_ctx, getState());
		enterRule(_localctx, 44, RULE_default_case);
		int _la;
		try {
			_localctx = new DefaultCaseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(318);
			match(DEFAULT_KW);
			setState(319);
			match(COLON);
			setState(323);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8590969760L) != 0)) {
				{
				{
				setState(320);
				stmt();
				}
				}
				setState(325);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class While_stmtContext extends ParserRuleContext {
		public While_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_while_stmt; }
	 
		public While_stmtContext() { }
		public void copyFrom(While_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class WhileStmtContext extends While_stmtContext {
		public TerminalNode WHILE_KW() { return getToken(VGrammar.WHILE_KW, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LBRACE() { return getToken(VGrammar.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VGrammar.RBRACE, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public WhileStmtContext(While_stmtContext ctx) { copyFrom(ctx); }
	}

	public final While_stmtContext while_stmt() throws RecognitionException {
		While_stmtContext _localctx = new While_stmtContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_while_stmt);
		int _la;
		try {
			_localctx = new WhileStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(326);
			match(WHILE_KW);
			setState(327);
			expr(0);
			setState(328);
			match(LBRACE);
			setState(332);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8590969760L) != 0)) {
				{
				{
				setState(329);
				stmt();
				}
				}
				setState(334);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(335);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class For_stmtContext extends ParserRuleContext {
		public For_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_stmt; }
	 
		public For_stmtContext() { }
		public void copyFrom(For_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ForStmtContext extends For_stmtContext {
		public TerminalNode FOR_KW() { return getToken(VGrammar.FOR_KW, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode IN_KW() { return getToken(VGrammar.IN_KW, 0); }
		public TerminalNode LBRACE() { return getToken(VGrammar.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VGrammar.RBRACE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public RangeContext range() {
			return getRuleContext(RangeContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ForStmtContext(For_stmtContext ctx) { copyFrom(ctx); }
	}

	public final For_stmtContext for_stmt() throws RecognitionException {
		For_stmtContext _localctx = new For_stmtContext(_ctx, getState());
		enterRule(_localctx, 48, RULE_for_stmt);
		int _la;
		try {
			_localctx = new ForStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(337);
			match(FOR_KW);
			setState(338);
			match(ID);
			setState(339);
			match(IN_KW);
			setState(342);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,25,_ctx) ) {
			case 1:
				{
				setState(340);
				expr(0);
				}
				break;
			case 2:
				{
				setState(341);
				range();
				}
				break;
			}
			setState(344);
			match(LBRACE);
			setState(348);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8590969760L) != 0)) {
				{
				{
				setState(345);
				stmt();
				}
				}
				setState(350);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(351);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class RangeContext extends ParserRuleContext {
		public RangeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_range; }
	 
		public RangeContext() { }
		public void copyFrom(RangeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class RangoNumContext extends RangeContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> DOT() { return getTokens(VGrammar.DOT); }
		public TerminalNode DOT(int i) {
			return getToken(VGrammar.DOT, i);
		}
		public RangoNumContext(RangeContext ctx) { copyFrom(ctx); }
	}

	public final RangeContext range() throws RecognitionException {
		RangeContext _localctx = new RangeContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_range);
		try {
			_localctx = new RangoNumContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(353);
			expr(0);
			setState(354);
			match(DOT);
			setState(355);
			match(DOT);
			setState(356);
			match(DOT);
			setState(357);
			expr(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Guard_stmtContext extends ParserRuleContext {
		public Guard_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_guard_stmt; }
	 
		public Guard_stmtContext() { }
		public void copyFrom(Guard_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class GuardStmtContext extends Guard_stmtContext {
		public TerminalNode GUARD_KW() { return getToken(VGrammar.GUARD_KW, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ELSE_KW() { return getToken(VGrammar.ELSE_KW, 0); }
		public TerminalNode LBRACE() { return getToken(VGrammar.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VGrammar.RBRACE, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public GuardStmtContext(Guard_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Guard_stmtContext guard_stmt() throws RecognitionException {
		Guard_stmtContext _localctx = new Guard_stmtContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_guard_stmt);
		int _la;
		try {
			_localctx = new GuardStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(359);
			match(GUARD_KW);
			setState(360);
			expr(0);
			setState(361);
			match(ELSE_KW);
			setState(362);
			match(LBRACE);
			setState(366);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8590969760L) != 0)) {
				{
				{
				setState(363);
				stmt();
				}
				}
				setState(368);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(369);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Transfer_stmtContext extends ParserRuleContext {
		public Transfer_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_transfer_stmt; }
	 
		public Transfer_stmtContext() { }
		public void copyFrom(Transfer_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ContinueStmtContext extends Transfer_stmtContext {
		public TerminalNode CONTINUE_KW() { return getToken(VGrammar.CONTINUE_KW, 0); }
		public ContinueStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStmtContext extends Transfer_stmtContext {
		public TerminalNode BREAK_KW() { return getToken(VGrammar.BREAK_KW, 0); }
		public BreakStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStmtContext extends Transfer_stmtContext {
		public TerminalNode RETURN_KW() { return getToken(VGrammar.RETURN_KW, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ReturnStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Transfer_stmtContext transfer_stmt() throws RecognitionException {
		Transfer_stmtContext _localctx = new Transfer_stmtContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_transfer_stmt);
		try {
			setState(377);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RETURN_KW:
				_localctx = new ReturnStmtContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(371);
				match(RETURN_KW);
				setState(373);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,28,_ctx) ) {
				case 1:
					{
					setState(372);
					expr(0);
					}
					break;
				}
				}
				break;
			case BREAK_KW:
				_localctx = new BreakStmtContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(375);
				match(BREAK_KW);
				}
				break;
			case CONTINUE_KW:
				_localctx = new ContinueStmtContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(376);
				match(CONTINUE_KW);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Func_callContext extends ParserRuleContext {
		public Func_callContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_call; }
	 
		public Func_callContext() { }
		public void copyFrom(Func_callContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LlamarFuncionContext extends Func_callContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public TerminalNode LPAREN() { return getToken(VGrammar.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VGrammar.RPAREN, 0); }
		public Arg_listContext arg_list() {
			return getRuleContext(Arg_listContext.class,0);
		}
		public LlamarFuncionContext(Func_callContext ctx) { copyFrom(ctx); }
	}

	public final Func_callContext func_call() throws RecognitionException {
		Func_callContext _localctx = new Func_callContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_func_call);
		int _la;
		try {
			_localctx = new LlamarFuncionContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(379);
			id_pattern();
			setState(380);
			match(LPAREN);
			setState(382);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4651092558376599552L) != 0)) {
				{
				setState(381);
				arg_list();
				}
			}

			setState(384);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Arg_listContext extends ParserRuleContext {
		public Arg_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arg_list; }
	 
		public Arg_listContext() { }
		public void copyFrom(Arg_listContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ArgListContext extends Arg_listContext {
		public List<Func_argContext> func_arg() {
			return getRuleContexts(Func_argContext.class);
		}
		public Func_argContext func_arg(int i) {
			return getRuleContext(Func_argContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VGrammar.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VGrammar.COMMA, i);
		}
		public ArgListContext(Arg_listContext ctx) { copyFrom(ctx); }
	}

	public final Arg_listContext arg_list() throws RecognitionException {
		Arg_listContext _localctx = new Arg_listContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_arg_list);
		int _la;
		try {
			_localctx = new ArgListContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(386);
			func_arg();
			setState(391);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(387);
				match(COMMA);
				setState(388);
				func_arg();
				}
				}
				setState(393);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Func_argContext extends ParserRuleContext {
		public Func_argContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_arg; }
	 
		public Func_argContext() { }
		public void copyFrom(Func_argContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionArgContext extends Func_argContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode COLON() { return getToken(VGrammar.COLON, 0); }
		public TerminalNode ANPERSAND() { return getToken(VGrammar.ANPERSAND, 0); }
		public FuncionArgContext(Func_argContext ctx) { copyFrom(ctx); }
	}

	public final Func_argContext func_arg() throws RecognitionException {
		Func_argContext _localctx = new Func_argContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_func_arg);
		int _la;
		try {
			_localctx = new FuncionArgContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(396);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				{
				setState(394);
				match(ID);
				setState(395);
				match(COLON);
				}
				break;
			}
			setState(399);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ANPERSAND) {
				{
				setState(398);
				match(ANPERSAND);
				}
			}

			setState(403);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				{
				setState(401);
				id_pattern();
				}
				break;
			case 2:
				{
				setState(402);
				expr(0);
				}
				break;
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Func_dclContext extends ParserRuleContext {
		public Func_dclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_dcl; }
	 
		public Func_dclContext() { }
		public void copyFrom(Func_dclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionDecleradaContext extends Func_dclContext {
		public TerminalNode FUNC_KW() { return getToken(VGrammar.FUNC_KW, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode LPAREN() { return getToken(VGrammar.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VGrammar.RPAREN, 0); }
		public TerminalNode LBRACE() { return getToken(VGrammar.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VGrammar.RBRACE, 0); }
		public Param_listContext param_list() {
			return getRuleContext(Param_listContext.class,0);
		}
		public TerminalNode ARROW() { return getToken(VGrammar.ARROW, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public FuncionDecleradaContext(Func_dclContext ctx) { copyFrom(ctx); }
	}

	public final Func_dclContext func_dcl() throws RecognitionException {
		Func_dclContext _localctx = new Func_dclContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_func_dcl);
		int _la;
		try {
			_localctx = new FuncionDecleradaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(405);
			match(FUNC_KW);
			setState(406);
			match(ID);
			setState(407);
			match(LPAREN);
			setState(409);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(408);
				param_list();
				}
			}

			setState(411);
			match(RPAREN);
			setState(414);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ARROW) {
				{
				setState(412);
				match(ARROW);
				setState(413);
				type();
				}
			}

			setState(416);
			match(LBRACE);
			setState(420);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8590969760L) != 0)) {
				{
				{
				setState(417);
				stmt();
				}
				}
				setState(422);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(423);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Param_listContext extends ParserRuleContext {
		public Param_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param_list; }
	 
		public Param_listContext() { }
		public void copyFrom(Param_listContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParamListContext extends Param_listContext {
		public List<Func_paramContext> func_param() {
			return getRuleContexts(Func_paramContext.class);
		}
		public Func_paramContext func_param(int i) {
			return getRuleContext(Func_paramContext.class,i);
		}
		public List<TerminalNode> COMMA() { return getTokens(VGrammar.COMMA); }
		public TerminalNode COMMA(int i) {
			return getToken(VGrammar.COMMA, i);
		}
		public ParamListContext(Param_listContext ctx) { copyFrom(ctx); }
	}

	public final Param_listContext param_list() throws RecognitionException {
		Param_listContext _localctx = new Param_listContext(_ctx, getState());
		enterRule(_localctx, 64, RULE_param_list);
		int _la;
		try {
			_localctx = new ParamListContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(425);
			func_param();
			setState(430);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==COMMA) {
				{
				{
				setState(426);
				match(COMMA);
				setState(427);
				func_param();
				}
				}
				setState(432);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Func_paramContext extends ParserRuleContext {
		public Func_paramContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_param; }
	 
		public Func_paramContext() { }
		public void copyFrom(Func_paramContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncParamContext extends Func_paramContext {
		public List<TerminalNode> ID() { return getTokens(VGrammar.ID); }
		public TerminalNode ID(int i) {
			return getToken(VGrammar.ID, i);
		}
		public TerminalNode COLON() { return getToken(VGrammar.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode INOUT_KW() { return getToken(VGrammar.INOUT_KW, 0); }
		public FuncParamContext(Func_paramContext ctx) { copyFrom(ctx); }
	}

	public final Func_paramContext func_param() throws RecognitionException {
		Func_paramContext _localctx = new Func_paramContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_func_param);
		int _la;
		try {
			_localctx = new FuncParamContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(434);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				{
				setState(433);
				match(ID);
				}
				break;
			}
			setState(436);
			match(ID);
			setState(437);
			match(COLON);
			setState(439);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==INOUT_KW) {
				{
				setState(438);
				match(INOUT_KW);
				}
			}

			setState(441);
			type();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Strct_dclContext extends ParserRuleContext {
		public Strct_dclContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_strct_dcl; }
	 
		public Strct_dclContext() { }
		public void copyFrom(Strct_dclContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class DeclararStructContext extends Strct_dclContext {
		public TerminalNode STRUCT_KW() { return getToken(VGrammar.STRUCT_KW, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode LBRACE() { return getToken(VGrammar.LBRACE, 0); }
		public TerminalNode RBRACE() { return getToken(VGrammar.RBRACE, 0); }
		public List<Struct_propContext> struct_prop() {
			return getRuleContexts(Struct_propContext.class);
		}
		public Struct_propContext struct_prop(int i) {
			return getRuleContext(Struct_propContext.class,i);
		}
		public DeclararStructContext(Strct_dclContext ctx) { copyFrom(ctx); }
	}

	public final Strct_dclContext strct_dcl() throws RecognitionException {
		Strct_dclContext _localctx = new Strct_dclContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_strct_dcl);
		int _la;
		try {
			_localctx = new DeclararStructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(443);
			match(STRUCT_KW);
			setState(444);
			match(ID);
			setState(445);
			match(LBRACE);
			setState(449);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 2097312L) != 0)) {
				{
				{
				setState(446);
				struct_prop();
				}
				}
				setState(451);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(452);
			match(RBRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Struct_propContext extends ParserRuleContext {
		public Struct_propContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_prop; }
	 
		public Struct_propContext() { }
		public void copyFrom(Struct_propContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructAttrContext extends Struct_propContext {
		public Var_typeContext var_type() {
			return getRuleContext(Var_typeContext.class,0);
		}
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode COLON() { return getToken(VGrammar.COLON, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode EQUALS() { return getToken(VGrammar.EQUALS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StructAttrContext(Struct_propContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructFuncContext extends Struct_propContext {
		public Func_dclContext func_dcl() {
			return getRuleContext(Func_dclContext.class,0);
		}
		public TerminalNode MUTATING_KW() { return getToken(VGrammar.MUTATING_KW, 0); }
		public StructFuncContext(Struct_propContext ctx) { copyFrom(ctx); }
	}

	public final Struct_propContext struct_prop() throws RecognitionException {
		Struct_propContext _localctx = new Struct_propContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_struct_prop);
		int _la;
		try {
			setState(468);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RW_MUT:
				_localctx = new StructAttrContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(454);
				var_type();
				setState(455);
				match(ID);
				setState(458);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==COLON) {
					{
					setState(456);
					match(COLON);
					setState(457);
					type();
					}
				}

				setState(462);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==EQUALS) {
					{
					setState(460);
					match(EQUALS);
					setState(461);
					expr(0);
					}
				}

				}
				break;
			case FUNC_KW:
			case MUTATING_KW:
				_localctx = new StructFuncContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(465);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==MUTATING_KW) {
					{
					setState(464);
					match(MUTATING_KW);
					}
				}

				setState(467);
				func_dcl();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	@SuppressWarnings("CheckReturnValue")
	public static class Struct_vectorContext extends ParserRuleContext {
		public Struct_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_vector; }
	 
		public Struct_vectorContext() { }
		public void copyFrom(Struct_vectorContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructVectorContext extends Struct_vectorContext {
		public TerminalNode LBRACK() { return getToken(VGrammar.LBRACK, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RBRACK() { return getToken(VGrammar.RBRACK, 0); }
		public TerminalNode LPAREN() { return getToken(VGrammar.LPAREN, 0); }
		public TerminalNode RPAREN() { return getToken(VGrammar.RPAREN, 0); }
		public StructVectorContext(Struct_vectorContext ctx) { copyFrom(ctx); }
	}

	public final Struct_vectorContext struct_vector() throws RecognitionException {
		Struct_vectorContext _localctx = new Struct_vectorContext(_ctx, getState());
		enterRule(_localctx, 72, RULE_struct_vector);
		try {
			_localctx = new StructVectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(470);
			match(LBRACK);
			setState(471);
			match(ID);
			setState(472);
			match(RBRACK);
			setState(473);
			match(LPAREN);
			setState(474);
			match(RPAREN);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 16:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 6);
		case 1:
			return precpred(_ctx, 5);
		case 2:
			return precpred(_ctx, 4);
		case 3:
			return precpred(_ctx, 3);
		case 4:
			return precpred(_ctx, 2);
		case 5:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001>\u01dd\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0002\u0011\u0007\u0011\u0002\u0012\u0007\u0012"+
		"\u0002\u0013\u0007\u0013\u0002\u0014\u0007\u0014\u0002\u0015\u0007\u0015"+
		"\u0002\u0016\u0007\u0016\u0002\u0017\u0007\u0017\u0002\u0018\u0007\u0018"+
		"\u0002\u0019\u0007\u0019\u0002\u001a\u0007\u001a\u0002\u001b\u0007\u001b"+
		"\u0002\u001c\u0007\u001c\u0002\u001d\u0007\u001d\u0002\u001e\u0007\u001e"+
		"\u0002\u001f\u0007\u001f\u0002 \u0007 \u0002!\u0007!\u0002\"\u0007\"\u0002"+
		"#\u0007#\u0002$\u0007$\u0001\u0000\u0005\u0000L\b\u0000\n\u0000\f\u0000"+
		"O\t\u0000\u0001\u0000\u0003\u0000R\b\u0000\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0003\u0001`\b\u0001\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003\u0002y\b"+
		"\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0005\u0003\u007f"+
		"\b\u0003\n\u0003\f\u0003\u0082\t\u0003\u0003\u0003\u0084\b\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0004\u0004\u008d\b\u0004\u000b\u0004\f\u0004\u008e\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0006\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0007\u0001\u0007\u0003\u0007\u009b\b\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0001\u0007\u0001\u0007\u0001\b\u0001\b\u0001\t\u0001\t\u0001"+
		"\t\u0001\t\u0001\t\u0001\t\u0001\t\u0003\t\u00b0\b\t\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0001\u000b\u0003\u000b\u00bc\b\u000b\u0001\f\u0001\f\u0001\f\u0001\f"+
		"\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001"+
		"\r\u0001\r\u0001\r\u0001\r\u0003\r\u00ce\b\r\u0001\u000e\u0001\u000e\u0001"+
		"\u000e\u0005\u000e\u00d3\b\u000e\n\u000e\f\u000e\u00d6\t\u000e\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0003\u000f\u00dd\b\u000f"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u00ef\b\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010"+
		"\u0005\u0010\u0103\b\u0010\n\u0010\f\u0010\u0106\t\u0010\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0005\u0011\u010b\b\u0011\n\u0011\f\u0011\u010e\t\u0011"+
		"\u0001\u0011\u0003\u0011\u0111\b\u0011\u0001\u0012\u0001\u0012\u0001\u0012"+
		"\u0001\u0012\u0005\u0012\u0117\b\u0012\n\u0012\f\u0012\u011a\t\u0012\u0001"+
		"\u0012\u0001\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0005\u0013\u0121"+
		"\b\u0013\n\u0013\f\u0013\u0124\t\u0013\u0001\u0013\u0001\u0013\u0001\u0014"+
		"\u0001\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u012c\b\u0014\n\u0014"+
		"\f\u0014\u012f\t\u0014\u0001\u0014\u0003\u0014\u0132\b\u0014\u0001\u0014"+
		"\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015\u0001\u0015\u0005\u0015"+
		"\u013a\b\u0015\n\u0015\f\u0015\u013d\t\u0015\u0001\u0016\u0001\u0016\u0001"+
		"\u0016\u0005\u0016\u0142\b\u0016\n\u0016\f\u0016\u0145\t\u0016\u0001\u0017"+
		"\u0001\u0017\u0001\u0017\u0001\u0017\u0005\u0017\u014b\b\u0017\n\u0017"+
		"\f\u0017\u014e\t\u0017\u0001\u0017\u0001\u0017\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0001\u0018\u0001\u0018\u0003\u0018\u0157\b\u0018\u0001\u0018"+
		"\u0001\u0018\u0005\u0018\u015b\b\u0018\n\u0018\f\u0018\u015e\t\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0001\u0019\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001"+
		"\u001a\u0005\u001a\u016d\b\u001a\n\u001a\f\u001a\u0170\t\u001a\u0001\u001a"+
		"\u0001\u001a\u0001\u001b\u0001\u001b\u0003\u001b\u0176\b\u001b\u0001\u001b"+
		"\u0001\u001b\u0003\u001b\u017a\b\u001b\u0001\u001c\u0001\u001c\u0001\u001c"+
		"\u0003\u001c\u017f\b\u001c\u0001\u001c\u0001\u001c\u0001\u001d\u0001\u001d"+
		"\u0001\u001d\u0005\u001d\u0186\b\u001d\n\u001d\f\u001d\u0189\t\u001d\u0001"+
		"\u001e\u0001\u001e\u0003\u001e\u018d\b\u001e\u0001\u001e\u0003\u001e\u0190"+
		"\b\u001e\u0001\u001e\u0001\u001e\u0003\u001e\u0194\b\u001e\u0001\u001f"+
		"\u0001\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u019a\b\u001f\u0001\u001f"+
		"\u0001\u001f\u0001\u001f\u0003\u001f\u019f\b\u001f\u0001\u001f\u0001\u001f"+
		"\u0005\u001f\u01a3\b\u001f\n\u001f\f\u001f\u01a6\t\u001f\u0001\u001f\u0001"+
		"\u001f\u0001 \u0001 \u0001 \u0005 \u01ad\b \n \f \u01b0\t \u0001!\u0003"+
		"!\u01b3\b!\u0001!\u0001!\u0001!\u0003!\u01b8\b!\u0001!\u0001!\u0001\""+
		"\u0001\"\u0001\"\u0001\"\u0005\"\u01c0\b\"\n\"\f\"\u01c3\t\"\u0001\"\u0001"+
		"\"\u0001#\u0001#\u0001#\u0001#\u0003#\u01cb\b#\u0001#\u0001#\u0003#\u01cf"+
		"\b#\u0001#\u0003#\u01d2\b#\u0001#\u0003#\u01d5\b#\u0001$\u0001$\u0001"+
		"$\u0001$\u0001$\u0001$\u0001$\u0000\u0001 %\u0000\u0002\u0004\u0006\b"+
		"\n\f\u000e\u0010\u0012\u0014\u0016\u0018\u001a\u001c\u001e \"$&(*,.02"+
		"468:<>@BDFH\u0000\u0007\u0001\u0000()\u0001\u0000\')\u0002\u0000##22\u0001"+
		"\u0000$&\u0001\u0000\"#\u0001\u0000,/\u0001\u0000*+\u020a\u0000M\u0001"+
		"\u0000\u0000\u0000\u0002_\u0001\u0000\u0000\u0000\u0004x\u0001\u0000\u0000"+
		"\u0000\u0006z\u0001\u0000\u0000\u0000\b\u0087\u0001\u0000\u0000\u0000"+
		"\n\u0090\u0001\u0000\u0000\u0000\f\u0094\u0001\u0000\u0000\u0000\u000e"+
		"\u009a\u0001\u0000\u0000\u0000\u0010\u00a6\u0001\u0000\u0000\u0000\u0012"+
		"\u00af\u0001\u0000\u0000\u0000\u0014\u00b1\u0001\u0000\u0000\u0000\u0016"+
		"\u00bb\u0001\u0000\u0000\u0000\u0018\u00bd\u0001\u0000\u0000\u0000\u001a"+
		"\u00cd\u0001\u0000\u0000\u0000\u001c\u00cf\u0001\u0000\u0000\u0000\u001e"+
		"\u00dc\u0001\u0000\u0000\u0000 \u00ee\u0001\u0000\u0000\u0000\"\u0107"+
		"\u0001\u0000\u0000\u0000$\u0112\u0001\u0000\u0000\u0000&\u011d\u0001\u0000"+
		"\u0000\u0000(\u0127\u0001\u0000\u0000\u0000*\u0135\u0001\u0000\u0000\u0000"+
		",\u013e\u0001\u0000\u0000\u0000.\u0146\u0001\u0000\u0000\u00000\u0151"+
		"\u0001\u0000\u0000\u00002\u0161\u0001\u0000\u0000\u00004\u0167\u0001\u0000"+
		"\u0000\u00006\u0179\u0001\u0000\u0000\u00008\u017b\u0001\u0000\u0000\u0000"+
		":\u0182\u0001\u0000\u0000\u0000<\u018c\u0001\u0000\u0000\u0000>\u0195"+
		"\u0001\u0000\u0000\u0000@\u01a9\u0001\u0000\u0000\u0000B\u01b2\u0001\u0000"+
		"\u0000\u0000D\u01bb\u0001\u0000\u0000\u0000F\u01d4\u0001\u0000\u0000\u0000"+
		"H\u01d6\u0001\u0000\u0000\u0000JL\u0003\u0002\u0001\u0000KJ\u0001\u0000"+
		"\u0000\u0000LO\u0001\u0000\u0000\u0000MK\u0001\u0000\u0000\u0000MN\u0001"+
		"\u0000\u0000\u0000NQ\u0001\u0000\u0000\u0000OM\u0001\u0000\u0000\u0000"+
		"PR\u0005\u0000\u0000\u0001QP\u0001\u0000\u0000\u0000QR\u0001\u0000\u0000"+
		"\u0000R\u0001\u0001\u0000\u0000\u0000S`\u0003\u0004\u0002\u0000T`\u0003"+
		"\u001a\r\u0000U`\u00036\u001b\u0000V`\u0003\"\u0011\u0000W`\u0003(\u0014"+
		"\u0000X`\u0003.\u0017\u0000Y`\u00030\u0018\u0000Z`\u00034\u001a\u0000"+
		"[`\u00038\u001c\u0000\\`\u0003\f\u0006\u0000]`\u0003>\u001f\u0000^`\u0003"+
		"D\"\u0000_S\u0001\u0000\u0000\u0000_T\u0001\u0000\u0000\u0000_U\u0001"+
		"\u0000\u0000\u0000_V\u0001\u0000\u0000\u0000_W\u0001\u0000\u0000\u0000"+
		"_X\u0001\u0000\u0000\u0000_Y\u0001\u0000\u0000\u0000_Z\u0001\u0000\u0000"+
		"\u0000_[\u0001\u0000\u0000\u0000_\\\u0001\u0000\u0000\u0000_]\u0001\u0000"+
		"\u0000\u0000_^\u0001\u0000\u0000\u0000`\u0003\u0001\u0000\u0000\u0000"+
		"ab\u0005\u0005\u0000\u0000bc\u0005!\u0000\u0000cd\u0005\u0006\u0000\u0000"+
		"dy\u0003 \u0010\u0000ef\u0005!\u0000\u0000fg\u0005\u0006\u0000\u0000g"+
		"y\u0003 \u0010\u0000hi\u0005\u0005\u0000\u0000ij\u0005!\u0000\u0000jk"+
		"\u0003\u0012\t\u0000kl\u0005\'\u0000\u0000lm\u0003 \u0010\u0000my\u0001"+
		"\u0000\u0000\u0000no\u0005\u0005\u0000\u0000op\u0005!\u0000\u0000py\u0003"+
		"\u0012\t\u0000qr\u0005!\u0000\u0000rs\u0003\u0012\t\u0000st\u0005\'\u0000"+
		"\u0000tu\u0003 \u0010\u0000uy\u0001\u0000\u0000\u0000vw\u0005!\u0000\u0000"+
		"wy\u0003\u0012\t\u0000xa\u0001\u0000\u0000\u0000xe\u0001\u0000\u0000\u0000"+
		"xh\u0001\u0000\u0000\u0000xn\u0001\u0000\u0000\u0000xq\u0001\u0000\u0000"+
		"\u0000xv\u0001\u0000\u0000\u0000y\u0005\u0001\u0000\u0000\u0000z\u0083"+
		"\u00057\u0000\u0000{\u0080\u0003 \u0010\u0000|}\u00059\u0000\u0000}\u007f"+
		"\u0003 \u0010\u0000~|\u0001\u0000\u0000\u0000\u007f\u0082\u0001\u0000"+
		"\u0000\u0000\u0080~\u0001\u0000\u0000\u0000\u0080\u0081\u0001\u0000\u0000"+
		"\u0000\u0081\u0084\u0001\u0000\u0000\u0000\u0082\u0080\u0001\u0000\u0000"+
		"\u0000\u0083{\u0001\u0000\u0000\u0000\u0083\u0084\u0001\u0000\u0000\u0000"+
		"\u0084\u0085\u0001\u0000\u0000\u0000\u0085\u0086\u00058\u0000\u0000\u0086"+
		"\u0007\u0001\u0000\u0000\u0000\u0087\u008c\u0003\u001c\u000e\u0000\u0088"+
		"\u0089\u00057\u0000\u0000\u0089\u008a\u0003 \u0010\u0000\u008a\u008b\u0005"+
		"8\u0000\u0000\u008b\u008d\u0001\u0000\u0000\u0000\u008c\u0088\u0001\u0000"+
		"\u0000\u0000\u008d\u008e\u0001\u0000\u0000\u0000\u008e\u008c\u0001\u0000"+
		"\u0000\u0000\u008e\u008f\u0001\u0000\u0000\u0000\u008f\t\u0001\u0000\u0000"+
		"\u0000\u0090\u0091\u0003\b\u0004\u0000\u0091\u0092\u0005:\u0000\u0000"+
		"\u0092\u0093\u0003\u001c\u000e\u0000\u0093\u000b\u0001\u0000\u0000\u0000"+
		"\u0094\u0095\u0003\b\u0004\u0000\u0095\u0096\u0005:\u0000\u0000\u0096"+
		"\u0097\u00038\u001c\u0000\u0097\r\u0001\u0000\u0000\u0000\u0098\u009b"+
		"\u0003\u0014\n\u0000\u0099\u009b\u0003\u0016\u000b\u0000\u009a\u0098\u0001"+
		"\u0000\u0000\u0000\u009a\u0099\u0001\u0000\u0000\u0000\u009b\u009c\u0001"+
		"\u0000\u0000\u0000\u009c\u009d\u00053\u0000\u0000\u009d\u009e\u0005!\u0000"+
		"\u0000\u009e\u009f\u0005;\u0000\u0000\u009f\u00a0\u0003 \u0010\u0000\u00a0"+
		"\u00a1\u00059\u0000\u0000\u00a1\u00a2\u0005!\u0000\u0000\u00a2\u00a3\u0005"+
		";\u0000\u0000\u00a3\u00a4\u0003 \u0010\u0000\u00a4\u00a5\u00054\u0000"+
		"\u0000\u00a5\u000f\u0001\u0000\u0000\u0000\u00a6\u00a7\u0005\u0005\u0000"+
		"\u0000\u00a7\u0011\u0001\u0000\u0000\u0000\u00a8\u00b0\u0005!\u0000\u0000"+
		"\u00a9\u00b0\u0005\u001c\u0000\u0000\u00aa\u00b0\u0005\u001d\u0000\u0000"+
		"\u00ab\u00b0\u0005\u001e\u0000\u0000\u00ac\u00b0\u0005\u001f\u0000\u0000"+
		"\u00ad\u00b0\u0003\u0014\n\u0000\u00ae\u00b0\u0003\u0016\u000b\u0000\u00af"+
		"\u00a8\u0001\u0000\u0000\u0000\u00af\u00a9\u0001\u0000\u0000\u0000\u00af"+
		"\u00aa\u0001\u0000\u0000\u0000\u00af\u00ab\u0001\u0000\u0000\u0000\u00af"+
		"\u00ac\u0001\u0000\u0000\u0000\u00af\u00ad\u0001\u0000\u0000\u0000\u00af"+
		"\u00ae\u0001\u0000\u0000\u0000\u00b0\u0013\u0001\u0000\u0000\u0000\u00b1"+
		"\u00b2\u00057\u0000\u0000\u00b2\u00b3\u0005!\u0000\u0000\u00b3\u00b4\u0005"+
		"8\u0000\u0000\u00b4\u0015\u0001\u0000\u0000\u0000\u00b5\u00bc\u0003\u0018"+
		"\f\u0000\u00b6\u00b7\u00057\u0000\u0000\u00b7\u00b8\u00057\u0000\u0000"+
		"\u00b8\u00b9\u0005!\u0000\u0000\u00b9\u00ba\u00058\u0000\u0000\u00ba\u00bc"+
		"\u00058\u0000\u0000\u00bb\u00b5\u0001\u0000\u0000\u0000\u00bb\u00b6\u0001"+
		"\u0000\u0000\u0000\u00bc\u0017\u0001\u0000\u0000\u0000\u00bd\u00be\u0005"+
		"7\u0000\u0000\u00be\u00bf\u0003\u0016\u000b\u0000\u00bf\u00c0\u00058\u0000"+
		"\u0000\u00c0\u0019\u0001\u0000\u0000\u0000\u00c1\u00c2\u0003\u001c\u000e"+
		"\u0000\u00c2\u00c3\u0005\'\u0000\u0000\u00c3\u00c4\u0003 \u0010\u0000"+
		"\u00c4\u00ce\u0001\u0000\u0000\u0000\u00c5\u00c6\u0003\u001c\u000e\u0000"+
		"\u00c6\u00c7\u0007\u0000\u0000\u0000\u00c7\u00c8\u0003 \u0010\u0000\u00c8"+
		"\u00ce\u0001\u0000\u0000\u0000\u00c9\u00ca\u0003\b\u0004\u0000\u00ca\u00cb"+
		"\u0007\u0001\u0000\u0000\u00cb\u00cc\u0003 \u0010\u0000\u00cc\u00ce\u0001"+
		"\u0000\u0000\u0000\u00cd\u00c1\u0001\u0000\u0000\u0000\u00cd\u00c5\u0001"+
		"\u0000\u0000\u0000\u00cd\u00c9\u0001\u0000\u0000\u0000\u00ce\u001b\u0001"+
		"\u0000\u0000\u0000\u00cf\u00d4\u0005!\u0000\u0000\u00d0\u00d1\u0005:\u0000"+
		"\u0000\u00d1\u00d3\u0005!\u0000\u0000\u00d2\u00d0\u0001\u0000\u0000\u0000"+
		"\u00d3\u00d6\u0001\u0000\u0000\u0000\u00d4\u00d2\u0001\u0000\u0000\u0000"+
		"\u00d4\u00d5\u0001\u0000\u0000\u0000\u00d5\u001d\u0001\u0000\u0000\u0000"+
		"\u00d6\u00d4\u0001\u0000\u0000\u0000\u00d7\u00dd\u0005\u0017\u0000\u0000"+
		"\u00d8\u00dd\u0005\u0018\u0000\u0000\u00d9\u00dd\u0005\u0019\u0000\u0000"+
		"\u00da\u00dd\u0005\u001a\u0000\u0000\u00db\u00dd\u0005\u001b\u0000\u0000"+
		"\u00dc\u00d7\u0001\u0000\u0000\u0000\u00dc\u00d8\u0001\u0000\u0000\u0000"+
		"\u00dc\u00d9\u0001\u0000\u0000\u0000\u00dc\u00da\u0001\u0000\u0000\u0000"+
		"\u00dc\u00db\u0001\u0000\u0000\u0000\u00dd\u001f\u0001\u0000\u0000\u0000"+
		"\u00de\u00df\u0006\u0010\uffff\uffff\u0000\u00df\u00e0\u00053\u0000\u0000"+
		"\u00e0\u00e1\u0003 \u0010\u0000\u00e1\u00e2\u00054\u0000\u0000\u00e2\u00ef"+
		"\u0001\u0000\u0000\u0000\u00e3\u00ef\u00038\u001c\u0000\u00e4\u00ef\u0003"+
		"\u001c\u000e\u0000\u00e5\u00ef\u0003\b\u0004\u0000\u00e6\u00ef\u0003\n"+
		"\u0005\u0000\u00e7\u00ef\u0003\f\u0006\u0000\u00e8\u00ef\u0003\u001e\u000f"+
		"\u0000\u00e9\u00ef\u0003\u0006\u0003\u0000\u00ea\u00ef\u0003\u000e\u0007"+
		"\u0000\u00eb\u00ef\u0003H$\u0000\u00ec\u00ed\u0007\u0002\u0000\u0000\u00ed"+
		"\u00ef\u0003 \u0010\u0007\u00ee\u00de\u0001\u0000\u0000\u0000\u00ee\u00e3"+
		"\u0001\u0000\u0000\u0000\u00ee\u00e4\u0001\u0000\u0000\u0000\u00ee\u00e5"+
		"\u0001\u0000\u0000\u0000\u00ee\u00e6\u0001\u0000\u0000\u0000\u00ee\u00e7"+
		"\u0001\u0000\u0000\u0000\u00ee\u00e8\u0001\u0000\u0000\u0000\u00ee\u00e9"+
		"\u0001\u0000\u0000\u0000\u00ee\u00ea\u0001\u0000\u0000\u0000\u00ee\u00eb"+
		"\u0001\u0000\u0000\u0000\u00ee\u00ec\u0001\u0000\u0000\u0000\u00ef\u0104"+
		"\u0001\u0000\u0000\u0000\u00f0\u00f1\n\u0006\u0000\u0000\u00f1\u00f2\u0007"+
		"\u0003\u0000\u0000\u00f2\u0103\u0003 \u0010\u0007\u00f3\u00f4\n\u0005"+
		"\u0000\u0000\u00f4\u00f5\u0007\u0004\u0000\u0000\u00f5\u0103\u0003 \u0010"+
		"\u0006\u00f6\u00f7\n\u0004\u0000\u0000\u00f7\u00f8\u0007\u0005\u0000\u0000"+
		"\u00f8\u0103\u0003 \u0010\u0005\u00f9\u00fa\n\u0003\u0000\u0000\u00fa"+
		"\u00fb\u0007\u0006\u0000\u0000\u00fb\u0103\u0003 \u0010\u0004\u00fc\u00fd"+
		"\n\u0002\u0000\u0000\u00fd\u00fe\u00050\u0000\u0000\u00fe\u0103\u0003"+
		" \u0010\u0003\u00ff\u0100\n\u0001\u0000\u0000\u0100\u0101\u00051\u0000"+
		"\u0000\u0101\u0103\u0003 \u0010\u0002\u0102\u00f0\u0001\u0000\u0000\u0000"+
		"\u0102\u00f3\u0001\u0000\u0000\u0000\u0102\u00f6\u0001\u0000\u0000\u0000"+
		"\u0102\u00f9\u0001\u0000\u0000\u0000\u0102\u00fc\u0001\u0000\u0000\u0000"+
		"\u0102\u00ff\u0001\u0000\u0000\u0000\u0103\u0106\u0001\u0000\u0000\u0000"+
		"\u0104\u0102\u0001\u0000\u0000\u0000\u0104\u0105\u0001\u0000\u0000\u0000"+
		"\u0105!\u0001\u0000\u0000\u0000\u0106\u0104\u0001\u0000\u0000\u0000\u0107"+
		"\u010c\u0003$\u0012\u0000\u0108\u0109\u0005\n\u0000\u0000\u0109\u010b"+
		"\u0003$\u0012\u0000\u010a\u0108\u0001\u0000\u0000\u0000\u010b\u010e\u0001"+
		"\u0000\u0000\u0000\u010c\u010a\u0001\u0000\u0000\u0000\u010c\u010d\u0001"+
		"\u0000\u0000\u0000\u010d\u0110\u0001\u0000\u0000\u0000\u010e\u010c\u0001"+
		"\u0000\u0000\u0000\u010f\u0111\u0003&\u0013\u0000\u0110\u010f\u0001\u0000"+
		"\u0000\u0000\u0110\u0111\u0001\u0000\u0000\u0000\u0111#\u0001\u0000\u0000"+
		"\u0000\u0112\u0113\u0005\t\u0000\u0000\u0113\u0114\u0003 \u0010\u0000"+
		"\u0114\u0118\u00055\u0000\u0000\u0115\u0117\u0003\u0002\u0001\u0000\u0116"+
		"\u0115\u0001\u0000\u0000\u0000\u0117\u011a\u0001\u0000\u0000\u0000\u0118"+
		"\u0116\u0001\u0000\u0000\u0000\u0118\u0119\u0001\u0000\u0000\u0000\u0119"+
		"\u011b\u0001\u0000\u0000\u0000\u011a\u0118\u0001\u0000\u0000\u0000\u011b"+
		"\u011c\u00056\u0000\u0000\u011c%\u0001\u0000\u0000\u0000\u011d\u011e\u0005"+
		"\n\u0000\u0000\u011e\u0122\u00055\u0000\u0000\u011f\u0121\u0003\u0002"+
		"\u0001\u0000\u0120\u011f\u0001\u0000\u0000\u0000\u0121\u0124\u0001\u0000"+
		"\u0000\u0000\u0122\u0120\u0001\u0000\u0000\u0000\u0122\u0123\u0001\u0000"+
		"\u0000\u0000\u0123\u0125\u0001\u0000\u0000\u0000\u0124\u0122\u0001\u0000"+
		"\u0000\u0000\u0125\u0126\u00056\u0000\u0000\u0126\'\u0001\u0000\u0000"+
		"\u0000\u0127\u0128\u0005\u000b\u0000\u0000\u0128\u0129\u0003 \u0010\u0000"+
		"\u0129\u012d\u00055\u0000\u0000\u012a\u012c\u0003*\u0015\u0000\u012b\u012a"+
		"\u0001\u0000\u0000\u0000\u012c\u012f\u0001\u0000\u0000\u0000\u012d\u012b"+
		"\u0001\u0000\u0000\u0000\u012d\u012e\u0001\u0000\u0000\u0000\u012e\u0131"+
		"\u0001\u0000\u0000\u0000\u012f\u012d\u0001\u0000\u0000\u0000\u0130\u0132"+
		"\u0003,\u0016\u0000\u0131\u0130\u0001\u0000\u0000\u0000\u0131\u0132\u0001"+
		"\u0000\u0000\u0000\u0132\u0133\u0001\u0000\u0000\u0000\u0133\u0134\u0005"+
		"6\u0000\u0000\u0134)\u0001\u0000\u0000\u0000\u0135\u0136\u0005\f\u0000"+
		"\u0000\u0136\u0137\u0003 \u0010\u0000\u0137\u013b\u0005;\u0000\u0000\u0138"+
		"\u013a\u0003\u0002\u0001\u0000\u0139\u0138\u0001\u0000\u0000\u0000\u013a"+
		"\u013d\u0001\u0000\u0000\u0000\u013b\u0139\u0001\u0000\u0000\u0000\u013b"+
		"\u013c\u0001\u0000\u0000\u0000\u013c+\u0001\u0000\u0000\u0000\u013d\u013b"+
		"\u0001\u0000\u0000\u0000\u013e\u013f\u0005\r\u0000\u0000\u013f\u0143\u0005"+
		";\u0000\u0000\u0140\u0142\u0003\u0002\u0001\u0000\u0141\u0140\u0001\u0000"+
		"\u0000\u0000\u0142\u0145\u0001\u0000\u0000\u0000\u0143\u0141\u0001\u0000"+
		"\u0000\u0000\u0143\u0144\u0001\u0000\u0000\u0000\u0144-\u0001\u0000\u0000"+
		"\u0000\u0145\u0143\u0001\u0000\u0000\u0000\u0146\u0147\u0005\u000f\u0000"+
		"\u0000\u0147\u0148\u0003 \u0010\u0000\u0148\u014c\u00055\u0000\u0000\u0149"+
		"\u014b\u0003\u0002\u0001\u0000\u014a\u0149\u0001\u0000\u0000\u0000\u014b"+
		"\u014e\u0001\u0000\u0000\u0000\u014c\u014a\u0001\u0000\u0000\u0000\u014c"+
		"\u014d\u0001\u0000\u0000\u0000\u014d\u014f\u0001\u0000\u0000\u0000\u014e"+
		"\u014c\u0001\u0000\u0000\u0000\u014f\u0150\u00056\u0000\u0000\u0150/\u0001"+
		"\u0000\u0000\u0000\u0151\u0152\u0005\u000e\u0000\u0000\u0152\u0153\u0005"+
		"!\u0000\u0000\u0153\u0156\u0005\u0016\u0000\u0000\u0154\u0157\u0003 \u0010"+
		"\u0000\u0155\u0157\u00032\u0019\u0000\u0156\u0154\u0001\u0000\u0000\u0000"+
		"\u0156\u0155\u0001\u0000\u0000\u0000\u0157\u0158\u0001\u0000\u0000\u0000"+
		"\u0158\u015c\u00055\u0000\u0000\u0159\u015b\u0003\u0002\u0001\u0000\u015a"+
		"\u0159\u0001\u0000\u0000\u0000\u015b\u015e\u0001\u0000\u0000\u0000\u015c"+
		"\u015a\u0001\u0000\u0000\u0000\u015c\u015d\u0001\u0000\u0000\u0000\u015d"+
		"\u015f\u0001\u0000\u0000\u0000\u015e\u015c\u0001\u0000\u0000\u0000\u015f"+
		"\u0160\u00056\u0000\u0000\u01601\u0001\u0000\u0000\u0000\u0161\u0162\u0003"+
		" \u0010\u0000\u0162\u0163\u0005:\u0000\u0000\u0163\u0164\u0005:\u0000"+
		"\u0000\u0164\u0165\u0005:\u0000\u0000\u0165\u0166\u0003 \u0010\u0000\u0166"+
		"3\u0001\u0000\u0000\u0000\u0167\u0168\u0005\u0013\u0000\u0000\u0168\u0169"+
		"\u0003 \u0010\u0000\u0169\u016a\u0005\n\u0000\u0000\u016a\u016e\u0005"+
		"5\u0000\u0000\u016b\u016d\u0003\u0002\u0001\u0000\u016c\u016b\u0001\u0000"+
		"\u0000\u0000\u016d\u0170\u0001\u0000\u0000\u0000\u016e\u016c\u0001\u0000"+
		"\u0000\u0000\u016e\u016f\u0001\u0000\u0000\u0000\u016f\u0171\u0001\u0000"+
		"\u0000\u0000\u0170\u016e\u0001\u0000\u0000\u0000\u0171\u0172\u00056\u0000"+
		"\u0000\u01725\u0001\u0000\u0000\u0000\u0173\u0175\u0005\u0012\u0000\u0000"+
		"\u0174\u0176\u0003 \u0010\u0000\u0175\u0174\u0001\u0000\u0000\u0000\u0175"+
		"\u0176\u0001\u0000\u0000\u0000\u0176\u017a\u0001\u0000\u0000\u0000\u0177"+
		"\u017a\u0005\u0010\u0000\u0000\u0178\u017a\u0005\u0011\u0000\u0000\u0179"+
		"\u0173\u0001\u0000\u0000\u0000\u0179\u0177\u0001\u0000\u0000\u0000\u0179"+
		"\u0178\u0001\u0000\u0000\u0000\u017a7\u0001\u0000\u0000\u0000\u017b\u017c"+
		"\u0003\u001c\u000e\u0000\u017c\u017e\u00053\u0000\u0000\u017d\u017f\u0003"+
		":\u001d\u0000\u017e\u017d\u0001\u0000\u0000\u0000\u017e\u017f\u0001\u0000"+
		"\u0000\u0000\u017f\u0180\u0001\u0000\u0000\u0000\u0180\u0181\u00054\u0000"+
		"\u0000\u01819\u0001\u0000\u0000\u0000\u0182\u0187\u0003<\u001e\u0000\u0183"+
		"\u0184\u00059\u0000\u0000\u0184\u0186\u0003<\u001e\u0000\u0185\u0183\u0001"+
		"\u0000\u0000\u0000\u0186\u0189\u0001\u0000\u0000\u0000\u0187\u0185\u0001"+
		"\u0000\u0000\u0000\u0187\u0188\u0001\u0000\u0000\u0000\u0188;\u0001\u0000"+
		"\u0000\u0000\u0189\u0187\u0001\u0000\u0000\u0000\u018a\u018b\u0005!\u0000"+
		"\u0000\u018b\u018d\u0005;\u0000\u0000\u018c\u018a\u0001\u0000\u0000\u0000"+
		"\u018c\u018d\u0001\u0000\u0000\u0000\u018d\u018f\u0001\u0000\u0000\u0000"+
		"\u018e\u0190\u0005>\u0000\u0000\u018f\u018e\u0001\u0000\u0000\u0000\u018f"+
		"\u0190\u0001\u0000\u0000\u0000\u0190\u0193\u0001\u0000\u0000\u0000\u0191"+
		"\u0194\u0003\u001c\u000e\u0000\u0192\u0194\u0003 \u0010\u0000\u0193\u0191"+
		"\u0001\u0000\u0000\u0000\u0193\u0192\u0001\u0000\u0000\u0000\u0194=\u0001"+
		"\u0000\u0000\u0000\u0195\u0196\u0005\u0007\u0000\u0000\u0196\u0197\u0005"+
		"!\u0000\u0000\u0197\u0199\u00053\u0000\u0000\u0198\u019a\u0003@ \u0000"+
		"\u0199\u0198\u0001\u0000\u0000\u0000\u0199\u019a\u0001\u0000\u0000\u0000"+
		"\u019a\u019b\u0001\u0000\u0000\u0000\u019b\u019e\u00054\u0000\u0000\u019c"+
		"\u019d\u0005<\u0000\u0000\u019d\u019f\u0003\u0012\t\u0000\u019e\u019c"+
		"\u0001\u0000\u0000\u0000\u019e\u019f\u0001\u0000\u0000\u0000\u019f\u01a0"+
		"\u0001\u0000\u0000\u0000\u01a0\u01a4\u00055\u0000\u0000\u01a1\u01a3\u0003"+
		"\u0002\u0001\u0000\u01a2\u01a1\u0001\u0000\u0000\u0000\u01a3\u01a6\u0001"+
		"\u0000\u0000\u0000\u01a4\u01a2\u0001\u0000\u0000\u0000\u01a4\u01a5\u0001"+
		"\u0000\u0000\u0000\u01a5\u01a7\u0001\u0000\u0000\u0000\u01a6\u01a4\u0001"+
		"\u0000\u0000\u0000\u01a7\u01a8\u00056\u0000\u0000\u01a8?\u0001\u0000\u0000"+
		"\u0000\u01a9\u01ae\u0003B!\u0000\u01aa\u01ab\u00059\u0000\u0000\u01ab"+
		"\u01ad\u0003B!\u0000\u01ac\u01aa\u0001\u0000\u0000\u0000\u01ad\u01b0\u0001"+
		"\u0000\u0000\u0000\u01ae\u01ac\u0001\u0000\u0000\u0000\u01ae\u01af\u0001"+
		"\u0000\u0000\u0000\u01afA\u0001\u0000\u0000\u0000\u01b0\u01ae\u0001\u0000"+
		"\u0000\u0000\u01b1\u01b3\u0005!\u0000\u0000\u01b2\u01b1\u0001\u0000\u0000"+
		"\u0000\u01b2\u01b3\u0001\u0000\u0000\u0000\u01b3\u01b4\u0001\u0000\u0000"+
		"\u0000\u01b4\u01b5\u0005!\u0000\u0000\u01b5\u01b7\u0005;\u0000\u0000\u01b6"+
		"\u01b8\u0005\u0014\u0000\u0000\u01b7\u01b6\u0001\u0000\u0000\u0000\u01b7"+
		"\u01b8\u0001\u0000\u0000\u0000\u01b8\u01b9\u0001\u0000\u0000\u0000\u01b9"+
		"\u01ba\u0003\u0012\t\u0000\u01baC\u0001\u0000\u0000\u0000\u01bb\u01bc"+
		"\u0005\b\u0000\u0000\u01bc\u01bd\u0005!\u0000\u0000\u01bd\u01c1\u0005"+
		"5\u0000\u0000\u01be\u01c0\u0003F#\u0000\u01bf\u01be\u0001\u0000\u0000"+
		"\u0000\u01c0\u01c3\u0001\u0000\u0000\u0000\u01c1\u01bf\u0001\u0000\u0000"+
		"\u0000\u01c1\u01c2\u0001\u0000\u0000\u0000\u01c2\u01c4\u0001\u0000\u0000"+
		"\u0000\u01c3\u01c1\u0001\u0000\u0000\u0000\u01c4\u01c5\u00056\u0000\u0000"+
		"\u01c5E\u0001\u0000\u0000\u0000\u01c6\u01c7\u0003\u0010\b\u0000\u01c7"+
		"\u01ca\u0005!\u0000\u0000\u01c8\u01c9\u0005;\u0000\u0000\u01c9\u01cb\u0003"+
		"\u0012\t\u0000\u01ca\u01c8\u0001\u0000\u0000\u0000\u01ca\u01cb\u0001\u0000"+
		"\u0000\u0000\u01cb\u01ce\u0001\u0000\u0000\u0000\u01cc\u01cd\u0005\'\u0000"+
		"\u0000\u01cd\u01cf\u0003 \u0010\u0000\u01ce\u01cc\u0001\u0000\u0000\u0000"+
		"\u01ce\u01cf\u0001\u0000\u0000\u0000\u01cf\u01d5\u0001\u0000\u0000\u0000"+
		"\u01d0\u01d2\u0005\u0015\u0000\u0000\u01d1\u01d0\u0001\u0000\u0000\u0000"+
		"\u01d1\u01d2\u0001\u0000\u0000\u0000\u01d2\u01d3\u0001\u0000\u0000\u0000"+
		"\u01d3\u01d5\u0003>\u001f\u0000\u01d4\u01c6\u0001\u0000\u0000\u0000\u01d4"+
		"\u01d1\u0001\u0000\u0000\u0000\u01d5G\u0001\u0000\u0000\u0000\u01d6\u01d7"+
		"\u00057\u0000\u0000\u01d7\u01d8\u0005!\u0000\u0000\u01d8\u01d9\u00058"+
		"\u0000\u0000\u01d9\u01da\u00053\u0000\u0000\u01da\u01db\u00054\u0000\u0000"+
		"\u01dbI\u0001\u0000\u0000\u0000.MQ_x\u0080\u0083\u008e\u009a\u00af\u00bb"+
		"\u00cd\u00d4\u00dc\u00ee\u0102\u0104\u010c\u0110\u0118\u0122\u012d\u0131"+
		"\u013b\u0143\u014c\u0156\u015c\u016e\u0175\u0179\u017e\u0187\u018c\u018f"+
		"\u0193\u0199\u019e\u01a4\u01ae\u01b2\u01b7\u01c1\u01ca\u01ce\u01d1\u01d4";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}