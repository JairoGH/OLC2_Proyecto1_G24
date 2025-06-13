// Generated from /home/jairogo/Documentos/GitHub/OLC2_Proyecto1/vlang-cherry-ide/parser/VGrammar.g4 by ANTLR 4.13.1
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
		WS=1, COMENTARIO=2, COMENTARIOMULT=3, RW_MUT=4, OP_ASSIGN=5, OP_INCREMENTO=6, 
		OP_DECREMENTO=7, OP_DECLARACION=8, RW_INT=9, RW_FLOAT64=10, RW_STRING=11, 
		RW_BOOL=12, INT_LITERAL=13, FLOAT_LITERAL=14, STRING_LITERAL=15, BOOL_LITERAL=16, 
		NIL_LITERAL=17, OP_SUMA=18, OP_RESTA=19, OP_MULT=20, OP_DIV=21, OP_MOD=22, 
		OP_IGUAL=23, OP_DIFERENTE=24, OP_MENORQ=25, OP_MAYORQ=26, OP_MENORIGUAL=27, 
		OP_MAYORIGUAL=28, OP_AND=29, OP_OR=30, OP_NOT=31, RW_MAIN=32, RW_FN=33, 
		RW_STRUCT=34, RW_IF=35, RW_ELSE=36, RW_SWITCH=37, RW_CASE=38, RW_DEFAULT=39, 
		RW_FOR=40, RW_WHILE=41, RW_BREAK=42, RW_CONTINUE=43, RW_RETURN=44, RW_IN=45, 
		RW_INOUT=46, RW_GUARD=47, RW_MUTATING=48, ID=49, PAR_IZQ=50, PAR_DER=51, 
		LLAVE_IZQ=52, LLAVE_DER=53, CORCHETE_IZQ=54, CORCHETE_DER=55, PUNTO=56, 
		COMA=57, PUNTO_Y_COMA=58, DOS_PUNTOS=59, ARROW=60, INTERROGATION=61, ANPERSAND=62;
	public static final int
		RULE_program = 0, RULE_main_func = 1, RULE_stmt = 2, RULE_decl_stmt = 3, 
		RULE_vector_expr = 4, RULE_vector_item = 5, RULE_vector_prop = 6, RULE_vector_func = 7, 
		RULE_repeating = 8, RULE_var_type = 9, RULE_type = 10, RULE_vector_type = 11, 
		RULE_matrix_type = 12, RULE_aux_matrix_type = 13, RULE_assign_stmt = 14, 
		RULE_id_pattern = 15, RULE_literal = 16, RULE_expr = 17, RULE_if_stmt = 18, 
		RULE_if_chain = 19, RULE_else_stmt = 20, RULE_switch_stmt = 21, RULE_switch_case = 22, 
		RULE_default_case = 23, RULE_while_stmt = 24, RULE_for_stmt = 25, RULE_range = 26, 
		RULE_guard_stmt = 27, RULE_transfer_stmt = 28, RULE_func_call = 29, RULE_arg_list = 30, 
		RULE_func_arg = 31, RULE_func_dcl = 32, RULE_method_receiver = 33, RULE_param_list = 34, 
		RULE_func_param = 35, RULE_strct_dcl = 36, RULE_struct_prop = 37, RULE_struct_vector = 38, 
		RULE_struct_init = 39, RULE_struct_init_list = 40, RULE_struct_method_call = 41, 
		RULE_struct_init_field = 42;
	private static String[] makeRuleNames() {
		return new String[] {
			"program", "main_func", "stmt", "decl_stmt", "vector_expr", "vector_item", 
			"vector_prop", "vector_func", "repeating", "var_type", "type", "vector_type", 
			"matrix_type", "aux_matrix_type", "assign_stmt", "id_pattern", "literal", 
			"expr", "if_stmt", "if_chain", "else_stmt", "switch_stmt", "switch_case", 
			"default_case", "while_stmt", "for_stmt", "range", "guard_stmt", "transfer_stmt", 
			"func_call", "arg_list", "func_arg", "func_dcl", "method_receiver", "param_list", 
			"func_param", "strct_dcl", "struct_prop", "struct_vector", "struct_init", 
			"struct_init_list", "struct_method_call", "struct_init_field"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, null, null, "'mut'", "':='", "'+='", "'-='", "'='", "'int'", 
			null, "'string'", "'bool'", null, null, null, null, "'nil'", "'+'", "'-'", 
			"'*'", "'/'", "'%'", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", "'&&'", 
			"'||'", "'!'", "'main'", "'fn'", "'struct'", "'if'", "'else'", "'switch'", 
			"'case'", "'default'", "'for'", "'while'", "'break'", "'continue'", "'return'", 
			"'in'", "'inout'", "'guard'", "'mutating'", null, "'('", "')'", "'{'", 
			"'}'", "'['", "']'", "'.'", "','", "';'", "':'", "'->'", "'?'", "'&'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "WS", "COMENTARIO", "COMENTARIOMULT", "RW_MUT", "OP_ASSIGN", "OP_INCREMENTO", 
			"OP_DECREMENTO", "OP_DECLARACION", "RW_INT", "RW_FLOAT64", "RW_STRING", 
			"RW_BOOL", "INT_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL", "BOOL_LITERAL", 
			"NIL_LITERAL", "OP_SUMA", "OP_RESTA", "OP_MULT", "OP_DIV", "OP_MOD", 
			"OP_IGUAL", "OP_DIFERENTE", "OP_MENORQ", "OP_MAYORQ", "OP_MENORIGUAL", 
			"OP_MAYORIGUAL", "OP_AND", "OP_OR", "OP_NOT", "RW_MAIN", "RW_FN", "RW_STRUCT", 
			"RW_IF", "RW_ELSE", "RW_SWITCH", "RW_CASE", "RW_DEFAULT", "RW_FOR", "RW_WHILE", 
			"RW_BREAK", "RW_CONTINUE", "RW_RETURN", "RW_IN", "RW_INOUT", "RW_GUARD", 
			"RW_MUTATING", "ID", "PAR_IZQ", "PAR_DER", "LLAVE_IZQ", "LLAVE_DER", 
			"CORCHETE_IZQ", "CORCHETE_DER", "PUNTO", "COMA", "PUNTO_Y_COMA", "DOS_PUNTOS", 
			"ARROW", "INTERROGATION", "ANPERSAND"
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
		public Main_funcContext main_func() {
			return getRuleContext(Main_funcContext.class,0);
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
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(89);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(86);
					stmt();
					}
					} 
				}
				setState(91);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,0,_ctx);
			}
			setState(93);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RW_FN) {
				{
				setState(92);
				main_func();
				}
			}

			setState(96);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				setState(95);
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
	public static class Main_funcContext extends ParserRuleContext {
		public Main_funcContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_main_func; }
	 
		public Main_funcContext() { }
		public void copyFrom(Main_funcContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionMainContext extends Main_funcContext {
		public TerminalNode RW_FN() { return getToken(VGrammar.RW_FN, 0); }
		public TerminalNode RW_MAIN() { return getToken(VGrammar.RW_MAIN, 0); }
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public FuncionMainContext(Main_funcContext ctx) { copyFrom(ctx); }
	}

	public final Main_funcContext main_func() throws RecognitionException {
		Main_funcContext _localctx = new Main_funcContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_main_func);
		int _la;
		try {
			_localctx = new FuncionMainContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(98);
			match(RW_FN);
			setState(99);
			match(RW_MAIN);
			setState(100);
			match(PAR_IZQ);
			setState(101);
			match(PAR_DER);
			setState(102);
			match(LLAVE_IZQ);
			setState(106);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 735770847477776L) != 0)) {
				{
				{
				setState(103);
				stmt();
				}
				}
				setState(108);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(109);
			match(LLAVE_DER);
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
		enterRule(_localctx, 4, RULE_stmt);
		try {
			setState(123);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(111);
				decl_stmt();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(112);
				assign_stmt();
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(113);
				transfer_stmt();
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(114);
				if_stmt();
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(115);
				switch_stmt();
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(116);
				while_stmt();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(117);
				for_stmt();
				}
				break;
			case 8:
				enterOuterAlt(_localctx, 8);
				{
				setState(118);
				guard_stmt();
				}
				break;
			case 9:
				enterOuterAlt(_localctx, 9);
				{
				setState(119);
				func_call();
				}
				break;
			case 10:
				enterOuterAlt(_localctx, 10);
				{
				setState(120);
				vector_func();
				}
				break;
			case 11:
				enterOuterAlt(_localctx, 11);
				{
				setState(121);
				func_dcl();
				}
				break;
			case 12:
				enterOuterAlt(_localctx, 12);
				{
				setState(122);
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
	public static class DeclararVectorContext extends Decl_stmtContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public DeclararVectorContext(Decl_stmtContext ctx) { copyFrom(ctx); }
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
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
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
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
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
		enterRule(_localctx, 6, RULE_decl_stmt);
		int _la;
		try {
			setState(157);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,7,_ctx) ) {
			case 1:
				_localctx = new DeclararInferenciaMutContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(125);
				match(RW_MUT);
				setState(126);
				match(ID);
				setState(127);
				match(OP_ASSIGN);
				setState(128);
				expr(0);
				}
				break;
			case 2:
				_localctx = new DeclararInferenciaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(129);
				match(ID);
				setState(130);
				match(OP_ASSIGN);
				setState(131);
				expr(0);
				}
				break;
			case 3:
				_localctx = new DeclaraTipoValorContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(132);
				match(RW_MUT);
				setState(133);
				match(ID);
				setState(134);
				type();
				setState(135);
				match(OP_DECLARACION);
				setState(136);
				expr(0);
				}
				break;
			case 4:
				_localctx = new DeclararTipoContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(138);
				match(RW_MUT);
				setState(139);
				match(ID);
				setState(140);
				type();
				}
				break;
			case 5:
				_localctx = new DeclararSinMutValorContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(141);
				match(ID);
				setState(142);
				type();
				setState(143);
				match(OP_DECLARACION);
				setState(144);
				expr(0);
				}
				break;
			case 6:
				_localctx = new DeclararSinMutTipoContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(146);
				match(ID);
				setState(147);
				type();
				}
				break;
			case 7:
				_localctx = new DeclararVectorContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(149);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==RW_MUT) {
					{
					setState(148);
					match(RW_MUT);
					}
				}

				setState(151);
				match(ID);
				setState(152);
				match(OP_DECLARACION);
				setState(153);
				type();
				setState(155);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
				case 1:
					{
					setState(154);
					expr(0);
					}
					break;
				}
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
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(VGrammar.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(VGrammar.COMA, i);
		}
		public ListaItemsVectorContext(Vector_exprContext ctx) { copyFrom(ctx); }
	}

	public final Vector_exprContext vector_expr() throws RecognitionException {
		Vector_exprContext _localctx = new Vector_exprContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_vector_expr);
		int _la;
		try {
			int _alt;
			_localctx = new ListaItemsVectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(159);
			match(LLAVE_IZQ);
			setState(168);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 24206850145378304L) != 0)) {
				{
				setState(160);
				expr(0);
				setState(165);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
				while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
					if ( _alt==1 ) {
						{
						{
						setState(161);
						match(COMA);
						setState(162);
						expr(0);
						}
						} 
					}
					setState(167);
					_errHandler.sync(this);
					_alt = getInterpreter().adaptivePredict(_input,8,_ctx);
				}
				}
			}

			setState(171);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMA) {
				{
				setState(170);
				match(COMA);
				}
			}

			setState(173);
			match(LLAVE_DER);
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
		public List<TerminalNode> CORCHETE_IZQ() { return getTokens(VGrammar.CORCHETE_IZQ); }
		public TerminalNode CORCHETE_IZQ(int i) {
			return getToken(VGrammar.CORCHETE_IZQ, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public List<TerminalNode> CORCHETE_DER() { return getTokens(VGrammar.CORCHETE_DER); }
		public TerminalNode CORCHETE_DER(int i) {
			return getToken(VGrammar.CORCHETE_DER, i);
		}
		public VectorItemContext(Vector_itemContext ctx) { copyFrom(ctx); }
	}

	public final Vector_itemContext vector_item() throws RecognitionException {
		Vector_itemContext _localctx = new Vector_itemContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_vector_item);
		try {
			int _alt;
			_localctx = new VectorItemContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			id_pattern();
			setState(180); 
			_errHandler.sync(this);
			_alt = 1;
			do {
				switch (_alt) {
				case 1:
					{
					{
					setState(176);
					match(CORCHETE_IZQ);
					setState(177);
					expr(0);
					setState(178);
					match(CORCHETE_DER);
					}
					}
					break;
				default:
					throw new NoViableAltException(this);
				}
				setState(182); 
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,11,_ctx);
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
		public TerminalNode PUNTO() { return getToken(VGrammar.PUNTO, 0); }
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public PropVectorContext(Vector_propContext ctx) { copyFrom(ctx); }
	}

	public final Vector_propContext vector_prop() throws RecognitionException {
		Vector_propContext _localctx = new Vector_propContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_vector_prop);
		try {
			_localctx = new PropVectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(184);
			vector_item();
			setState(185);
			match(PUNTO);
			setState(186);
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
		public TerminalNode PUNTO() { return getToken(VGrammar.PUNTO, 0); }
		public Func_callContext func_call() {
			return getRuleContext(Func_callContext.class,0);
		}
		public FuncionVectorContext(Vector_funcContext ctx) { copyFrom(ctx); }
	}

	public final Vector_funcContext vector_func() throws RecognitionException {
		Vector_funcContext _localctx = new Vector_funcContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_vector_func);
		try {
			_localctx = new FuncionVectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(188);
			vector_item();
			setState(189);
			match(PUNTO);
			setState(190);
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
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public List<TerminalNode> ID() { return getTokens(VGrammar.ID); }
		public TerminalNode ID(int i) {
			return getToken(VGrammar.ID, i);
		}
		public List<TerminalNode> DOS_PUNTOS() { return getTokens(VGrammar.DOS_PUNTOS); }
		public TerminalNode DOS_PUNTOS(int i) {
			return getToken(VGrammar.DOS_PUNTOS, i);
		}
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode COMA() { return getToken(VGrammar.COMA, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
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
		enterRule(_localctx, 16, RULE_repeating);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(194);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				{
				setState(192);
				vector_type();
				}
				break;
			case 2:
				{
				setState(193);
				matrix_type();
				}
				break;
			}
			setState(196);
			match(PAR_IZQ);
			setState(197);
			match(ID);
			setState(198);
			match(DOS_PUNTOS);
			setState(199);
			expr(0);
			setState(200);
			match(COMA);
			setState(201);
			match(ID);
			setState(202);
			match(DOS_PUNTOS);
			setState(203);
			expr(0);
			setState(204);
			match(PAR_DER);
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
		enterRule(_localctx, 18, RULE_var_type);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(206);
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
		enterRule(_localctx, 20, RULE_type);
		try {
			setState(215);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(208);
				match(ID);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(209);
				match(RW_INT);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(210);
				match(RW_FLOAT64);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(211);
				match(RW_STRING);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(212);
				match(RW_BOOL);
				}
				break;
			case 6:
				enterOuterAlt(_localctx, 6);
				{
				setState(213);
				vector_type();
				}
				break;
			case 7:
				enterOuterAlt(_localctx, 7);
				{
				setState(214);
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
		public Vector_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vector_type; }
	 
		public Vector_typeContext() { }
		public void copyFrom(Vector_typeContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorSimpleContext extends Vector_typeContext {
		public TerminalNode CORCHETE_IZQ() { return getToken(VGrammar.CORCHETE_IZQ, 0); }
		public TerminalNode CORCHETE_DER() { return getToken(VGrammar.CORCHETE_DER, 0); }
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public VectorSimpleContext(Vector_typeContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MatrizDobleContext extends Vector_typeContext {
		public List<TerminalNode> CORCHETE_IZQ() { return getTokens(VGrammar.CORCHETE_IZQ); }
		public TerminalNode CORCHETE_IZQ(int i) {
			return getToken(VGrammar.CORCHETE_IZQ, i);
		}
		public List<TerminalNode> CORCHETE_DER() { return getTokens(VGrammar.CORCHETE_DER); }
		public TerminalNode CORCHETE_DER(int i) {
			return getToken(VGrammar.CORCHETE_DER, i);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public MatrizDobleContext(Vector_typeContext ctx) { copyFrom(ctx); }
	}

	public final Vector_typeContext vector_type() throws RecognitionException {
		Vector_typeContext _localctx = new Vector_typeContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_vector_type);
		try {
			setState(225);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,14,_ctx) ) {
			case 1:
				_localctx = new VectorSimpleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(217);
				match(CORCHETE_IZQ);
				setState(218);
				match(CORCHETE_DER);
				setState(219);
				type();
				}
				break;
			case 2:
				_localctx = new MatrizDobleContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(220);
				match(CORCHETE_IZQ);
				setState(221);
				match(CORCHETE_DER);
				setState(222);
				match(CORCHETE_IZQ);
				setState(223);
				match(CORCHETE_DER);
				setState(224);
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
	public static class Matrix_typeContext extends ParserRuleContext {
		public Aux_matrix_typeContext aux_matrix_type() {
			return getRuleContext(Aux_matrix_typeContext.class,0);
		}
		public List<TerminalNode> CORCHETE_IZQ() { return getTokens(VGrammar.CORCHETE_IZQ); }
		public TerminalNode CORCHETE_IZQ(int i) {
			return getToken(VGrammar.CORCHETE_IZQ, i);
		}
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public List<TerminalNode> CORCHETE_DER() { return getTokens(VGrammar.CORCHETE_DER); }
		public TerminalNode CORCHETE_DER(int i) {
			return getToken(VGrammar.CORCHETE_DER, i);
		}
		public Matrix_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matrix_type; }
	}

	public final Matrix_typeContext matrix_type() throws RecognitionException {
		Matrix_typeContext _localctx = new Matrix_typeContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_matrix_type);
		try {
			setState(233);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(227);
				aux_matrix_type();
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(228);
				match(CORCHETE_IZQ);
				setState(229);
				match(CORCHETE_IZQ);
				setState(230);
				match(ID);
				setState(231);
				match(CORCHETE_DER);
				setState(232);
				match(CORCHETE_DER);
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
		public TerminalNode CORCHETE_IZQ() { return getToken(VGrammar.CORCHETE_IZQ, 0); }
		public Matrix_typeContext matrix_type() {
			return getRuleContext(Matrix_typeContext.class,0);
		}
		public TerminalNode CORCHETE_DER() { return getToken(VGrammar.CORCHETE_DER, 0); }
		public Aux_matrix_typeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_aux_matrix_type; }
	}

	public final Aux_matrix_typeContext aux_matrix_type() throws RecognitionException {
		Aux_matrix_typeContext _localctx = new Aux_matrix_typeContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_aux_matrix_type);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(235);
			match(CORCHETE_IZQ);
			setState(236);
			matrix_type();
			setState(237);
			match(CORCHETE_DER);
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
		public TerminalNode OP_INCREMENTO() { return getToken(VGrammar.OP_INCREMENTO, 0); }
		public TerminalNode OP_DECREMENTO() { return getToken(VGrammar.OP_DECREMENTO, 0); }
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
		public AsignacionVectorContext(Assign_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AsignacionDirectaContext extends Assign_stmtContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AsignacionDirectaContext(Assign_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignVectorItemContext extends Assign_stmtContext {
		public Vector_itemContext vector_item() {
			return getRuleContext(Vector_itemContext.class,0);
		}
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AssignVectorItemContext(Assign_stmtContext ctx) { copyFrom(ctx); }
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
		public TerminalNode OP_INCREMENTO() { return getToken(VGrammar.OP_INCREMENTO, 0); }
		public TerminalNode OP_DECREMENTO() { return getToken(VGrammar.OP_DECREMENTO, 0); }
		public AsignacionAritmeticaContext(Assign_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Assign_stmtContext assign_stmt() throws RecognitionException {
		Assign_stmtContext _localctx = new Assign_stmtContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_assign_stmt);
		int _la;
		try {
			setState(255);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,16,_ctx) ) {
			case 1:
				_localctx = new AssignVectorItemContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(239);
				vector_item();
				setState(240);
				match(OP_DECLARACION);
				setState(241);
				expr(0);
				}
				break;
			case 2:
				_localctx = new AsignacionDirectaContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(243);
				id_pattern();
				setState(244);
				match(OP_DECLARACION);
				setState(245);
				expr(0);
				}
				break;
			case 3:
				_localctx = new AsignacionAritmeticaContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(247);
				id_pattern();
				setState(248);
				((AsignacionAritmeticaContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==OP_INCREMENTO || _la==OP_DECREMENTO) ) {
					((AsignacionAritmeticaContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(249);
				expr(0);
				}
				break;
			case 4:
				_localctx = new AsignacionVectorContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(251);
				vector_item();
				setState(252);
				((AsignacionVectorContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 448L) != 0)) ) {
					((AsignacionVectorContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(253);
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
		public List<TerminalNode> PUNTO() { return getTokens(VGrammar.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(VGrammar.PUNTO, i);
		}
		public ID_PatronContext(Id_patternContext ctx) { copyFrom(ctx); }
	}

	public final Id_patternContext id_pattern() throws RecognitionException {
		Id_patternContext _localctx = new Id_patternContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_id_pattern);
		try {
			int _alt;
			_localctx = new ID_PatronContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(257);
			match(ID);
			setState(262);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(258);
					match(PUNTO);
					setState(259);
					match(ID);
					}
					} 
				}
				setState(264);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
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
		public TerminalNode INT_LITERAL() { return getToken(VGrammar.INT_LITERAL, 0); }
		public IntLiteralContext(LiteralContext ctx) { copyFrom(ctx); }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_literal);
		try {
			setState(270);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_LITERAL:
				_localctx = new IntLiteralContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(265);
				match(INT_LITERAL);
				}
				break;
			case FLOAT_LITERAL:
				_localctx = new FloatLiteralContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(266);
				match(FLOAT_LITERAL);
				}
				break;
			case STRING_LITERAL:
				_localctx = new StringLiteralContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(267);
				match(STRING_LITERAL);
				}
				break;
			case BOOL_LITERAL:
				_localctx = new BoolLiteralContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(268);
				match(BOOL_LITERAL);
				}
				break;
			case NIL_LITERAL:
				_localctx = new NilLiteralContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(269);
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
	public static class StructMethodExpContext extends ExprContext {
		public Struct_method_callContext struct_method_call() {
			return getRuleContext(Struct_method_callContext.class,0);
		}
		public StructMethodExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructInitExpContext extends ExprContext {
		public Struct_initContext struct_init() {
			return getRuleContext(Struct_initContext.class,0);
		}
		public StructInitExpContext(ExprContext ctx) { copyFrom(ctx); }
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
		public TerminalNode OP_MULT() { return getToken(VGrammar.OP_MULT, 0); }
		public TerminalNode OP_DIV() { return getToken(VGrammar.OP_DIV, 0); }
		public TerminalNode OP_MOD() { return getToken(VGrammar.OP_MOD, 0); }
		public TerminalNode OP_SUMA() { return getToken(VGrammar.OP_SUMA, 0); }
		public TerminalNode OP_RESTA() { return getToken(VGrammar.OP_RESTA, 0); }
		public TerminalNode OP_MENORQ() { return getToken(VGrammar.OP_MENORQ, 0); }
		public TerminalNode OP_MENORIGUAL() { return getToken(VGrammar.OP_MENORIGUAL, 0); }
		public TerminalNode OP_MAYORQ() { return getToken(VGrammar.OP_MAYORQ, 0); }
		public TerminalNode OP_MAYORIGUAL() { return getToken(VGrammar.OP_MAYORIGUAL, 0); }
		public TerminalNode OP_IGUAL() { return getToken(VGrammar.OP_IGUAL, 0); }
		public TerminalNode OP_DIFERENTE() { return getToken(VGrammar.OP_DIFERENTE, 0); }
		public TerminalNode OP_AND() { return getToken(VGrammar.OP_AND, 0); }
		public TerminalNode OP_OR() { return getToken(VGrammar.OP_OR, 0); }
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
	public static class VectorFuncExpContext extends ExprContext {
		public Vector_funcContext vector_func() {
			return getRuleContext(Vector_funcContext.class,0);
		}
		public VectorFuncExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParentecisExpContext extends ExprContext {
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public ParentecisExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExpUnaryContext extends ExprContext {
		public Token op;
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode OP_NOT() { return getToken(VGrammar.OP_NOT, 0); }
		public TerminalNode OP_RESTA() { return getToken(VGrammar.OP_RESTA, 0); }
		public ExpUnaryContext(ExprContext ctx) { copyFrom(ctx); }
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
	public static class VectorItemExpContext extends ExprContext {
		public Vector_itemContext vector_item() {
			return getRuleContext(Vector_itemContext.class,0);
		}
		public VectorItemExpContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class VectorExpContext extends ExprContext {
		public Vector_exprContext vector_expr() {
			return getRuleContext(Vector_exprContext.class,0);
		}
		public VectorExpContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 34;
		enterRecursionRule(_localctx, 34, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(290);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				{
				_localctx = new ParentecisExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(273);
				match(PAR_IZQ);
				setState(274);
				expr(0);
				setState(275);
				match(PAR_DER);
				}
				break;
			case 2:
				{
				_localctx = new FunctionCallExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(277);
				func_call();
				}
				break;
			case 3:
				{
				_localctx = new IdExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(278);
				id_pattern();
				}
				break;
			case 4:
				{
				_localctx = new VectorItemExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(279);
				vector_item();
				}
				break;
			case 5:
				{
				_localctx = new VectorPropExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(280);
				vector_prop();
				}
				break;
			case 6:
				{
				_localctx = new VectorFuncExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(281);
				vector_func();
				}
				break;
			case 7:
				{
				_localctx = new LiteralExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(282);
				literal();
				}
				break;
			case 8:
				{
				_localctx = new VectorExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(283);
				vector_expr();
				}
				break;
			case 9:
				{
				_localctx = new RepeatingExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(284);
				repeating();
				}
				break;
			case 10:
				{
				_localctx = new StructExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(285);
				struct_vector();
				}
				break;
			case 11:
				{
				_localctx = new StructInitExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(286);
				struct_init();
				}
				break;
			case 12:
				{
				_localctx = new StructMethodExpContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(287);
				struct_method_call();
				}
				break;
			case 13:
				{
				_localctx = new ExpUnaryContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(288);
				((ExpUnaryContext)_localctx).op = _input.LT(1);
				_la = _input.LA(1);
				if ( !(_la==OP_RESTA || _la==OP_NOT) ) {
					((ExpUnaryContext)_localctx).op = (Token)_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				setState(289);
				expr(7);
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(312);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(310);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
					case 1:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(292);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(293);
						((ExpBinarioContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 7340032L) != 0)) ) {
							((ExpBinarioContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(294);
						((ExpBinarioContext)_localctx).right = expr(7);
						}
						break;
					case 2:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(295);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(296);
						((ExpBinarioContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_SUMA || _la==OP_RESTA) ) {
							((ExpBinarioContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(297);
						((ExpBinarioContext)_localctx).right = expr(6);
						}
						break;
					case 3:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(298);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(299);
						((ExpBinarioContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 503316480L) != 0)) ) {
							((ExpBinarioContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(300);
						((ExpBinarioContext)_localctx).right = expr(5);
						}
						break;
					case 4:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(301);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(302);
						((ExpBinarioContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_IGUAL || _la==OP_DIFERENTE) ) {
							((ExpBinarioContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(303);
						((ExpBinarioContext)_localctx).right = expr(4);
						}
						break;
					case 5:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(304);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(305);
						((ExpBinarioContext)_localctx).op = match(OP_AND);
						setState(306);
						((ExpBinarioContext)_localctx).right = expr(3);
						}
						break;
					case 6:
						{
						_localctx = new ExpBinarioContext(new ExprContext(_parentctx, _parentState));
						((ExpBinarioContext)_localctx).left = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(307);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(308);
						((ExpBinarioContext)_localctx).op = match(OP_OR);
						setState(309);
						((ExpBinarioContext)_localctx).right = expr(2);
						}
						break;
					}
					} 
				}
				setState(314);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,21,_ctx);
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
		public List<TerminalNode> RW_ELSE() { return getTokens(VGrammar.RW_ELSE); }
		public TerminalNode RW_ELSE(int i) {
			return getToken(VGrammar.RW_ELSE, i);
		}
		public Else_stmtContext else_stmt() {
			return getRuleContext(Else_stmtContext.class,0);
		}
		public IFstmtContext(If_stmtContext ctx) { copyFrom(ctx); }
	}

	public final If_stmtContext if_stmt() throws RecognitionException {
		If_stmtContext _localctx = new If_stmtContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_if_stmt);
		int _la;
		try {
			int _alt;
			_localctx = new IFstmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(315);
			if_chain();
			setState(320);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(316);
					match(RW_ELSE);
					setState(317);
					if_chain();
					}
					} 
				}
				setState(322);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			}
			setState(324);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RW_ELSE) {
				{
				setState(323);
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
		public TerminalNode RW_IF() { return getToken(VGrammar.RW_IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
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
		enterRule(_localctx, 38, RULE_if_chain);
		int _la;
		try {
			_localctx = new IFcadenaContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(326);
			match(RW_IF);
			setState(327);
			expr(0);
			setState(328);
			match(LLAVE_IZQ);
			setState(332);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 735770847477776L) != 0)) {
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
			match(LLAVE_DER);
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
		public TerminalNode RW_ELSE() { return getToken(VGrammar.RW_ELSE, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
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
		enterRule(_localctx, 40, RULE_else_stmt);
		int _la;
		try {
			_localctx = new ElseStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(337);
			match(RW_ELSE);
			setState(338);
			match(LLAVE_IZQ);
			setState(342);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 735770847477776L) != 0)) {
				{
				{
				setState(339);
				stmt();
				}
				}
				setState(344);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(345);
			match(LLAVE_DER);
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
		public TerminalNode RW_SWITCH() { return getToken(VGrammar.RW_SWITCH, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
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
		enterRule(_localctx, 42, RULE_switch_stmt);
		int _la;
		try {
			_localctx = new SwitchStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(347);
			match(RW_SWITCH);
			setState(348);
			expr(0);
			setState(349);
			match(LLAVE_IZQ);
			setState(353);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==RW_CASE) {
				{
				{
				setState(350);
				switch_case();
				}
				}
				setState(355);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(357);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RW_DEFAULT) {
				{
				setState(356);
				default_case();
				}
			}

			setState(359);
			match(LLAVE_DER);
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
		public TerminalNode RW_CASE() { return getToken(VGrammar.RW_CASE, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode DOS_PUNTOS() { return getToken(VGrammar.DOS_PUNTOS, 0); }
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
		enterRule(_localctx, 44, RULE_switch_case);
		int _la;
		try {
			_localctx = new SwitchCaseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(361);
			match(RW_CASE);
			setState(362);
			expr(0);
			setState(363);
			match(DOS_PUNTOS);
			setState(367);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 735770847477776L) != 0)) {
				{
				{
				setState(364);
				stmt();
				}
				}
				setState(369);
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
		public TerminalNode RW_DEFAULT() { return getToken(VGrammar.RW_DEFAULT, 0); }
		public TerminalNode DOS_PUNTOS() { return getToken(VGrammar.DOS_PUNTOS, 0); }
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
		enterRule(_localctx, 46, RULE_default_case);
		int _la;
		try {
			_localctx = new DefaultCaseContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(370);
			match(RW_DEFAULT);
			setState(371);
			match(DOS_PUNTOS);
			setState(375);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 735770847477776L) != 0)) {
				{
				{
				setState(372);
				stmt();
				}
				}
				setState(377);
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
		public TerminalNode RW_FOR() { return getToken(VGrammar.RW_FOR, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
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
		enterRule(_localctx, 48, RULE_while_stmt);
		int _la;
		try {
			_localctx = new WhileStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(378);
			match(RW_FOR);
			setState(379);
			expr(0);
			setState(380);
			match(LLAVE_IZQ);
			setState(384);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 735770847477776L) != 0)) {
				{
				{
				setState(381);
				stmt();
				}
				}
				setState(386);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(387);
			match(LLAVE_DER);
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
		public TerminalNode RW_FOR() { return getToken(VGrammar.RW_FOR, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_IN() { return getToken(VGrammar.RW_IN, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
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
		enterRule(_localctx, 50, RULE_for_stmt);
		int _la;
		try {
			_localctx = new ForStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(389);
			match(RW_FOR);
			setState(390);
			match(ID);
			setState(391);
			match(RW_IN);
			setState(394);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,31,_ctx) ) {
			case 1:
				{
				setState(392);
				expr(0);
				}
				break;
			case 2:
				{
				setState(393);
				range();
				}
				break;
			}
			setState(396);
			match(LLAVE_IZQ);
			setState(400);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 735770847477776L) != 0)) {
				{
				{
				setState(397);
				stmt();
				}
				}
				setState(402);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(403);
			match(LLAVE_DER);
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
		public List<TerminalNode> PUNTO() { return getTokens(VGrammar.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(VGrammar.PUNTO, i);
		}
		public RangoNumContext(RangeContext ctx) { copyFrom(ctx); }
	}

	public final RangeContext range() throws RecognitionException {
		RangeContext _localctx = new RangeContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_range);
		try {
			_localctx = new RangoNumContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(405);
			expr(0);
			setState(406);
			match(PUNTO);
			setState(407);
			match(PUNTO);
			setState(408);
			match(PUNTO);
			setState(409);
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
		public TerminalNode RW_GUARD() { return getToken(VGrammar.RW_GUARD, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode RW_ELSE() { return getToken(VGrammar.RW_ELSE, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
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
		enterRule(_localctx, 54, RULE_guard_stmt);
		int _la;
		try {
			_localctx = new GuardStmtContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(411);
			match(RW_GUARD);
			setState(412);
			expr(0);
			setState(413);
			match(RW_ELSE);
			setState(414);
			match(LLAVE_IZQ);
			setState(418);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 735770847477776L) != 0)) {
				{
				{
				setState(415);
				stmt();
				}
				}
				setState(420);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(421);
			match(LLAVE_DER);
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
		public TerminalNode RW_CONTINUE() { return getToken(VGrammar.RW_CONTINUE, 0); }
		public ContinueStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class BreakStmtContext extends Transfer_stmtContext {
		public TerminalNode RW_BREAK() { return getToken(VGrammar.RW_BREAK, 0); }
		public BreakStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ReturnStmtContext extends Transfer_stmtContext {
		public TerminalNode RW_RETURN() { return getToken(VGrammar.RW_RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ReturnStmtContext(Transfer_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Transfer_stmtContext transfer_stmt() throws RecognitionException {
		Transfer_stmtContext _localctx = new Transfer_stmtContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_transfer_stmt);
		try {
			setState(429);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RW_RETURN:
				_localctx = new ReturnStmtContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(423);
				match(RW_RETURN);
				setState(425);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
				case 1:
					{
					setState(424);
					expr(0);
					}
					break;
				}
				}
				break;
			case RW_BREAK:
				_localctx = new BreakStmtContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(427);
				match(RW_BREAK);
				}
				break;
			case RW_CONTINUE:
				_localctx = new ContinueStmtContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(428);
				match(RW_CONTINUE);
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
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public Arg_listContext arg_list() {
			return getRuleContext(Arg_listContext.class,0);
		}
		public LlamarFuncionContext(Func_callContext ctx) { copyFrom(ctx); }
	}

	public final Func_callContext func_call() throws RecognitionException {
		Func_callContext _localctx = new Func_callContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_func_call);
		int _la;
		try {
			_localctx = new LlamarFuncionContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(431);
			id_pattern();
			setState(432);
			match(PAR_IZQ);
			setState(434);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4635892868572766208L) != 0)) {
				{
				setState(433);
				arg_list();
				}
			}

			setState(436);
			match(PAR_DER);
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
		public List<TerminalNode> COMA() { return getTokens(VGrammar.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(VGrammar.COMA, i);
		}
		public ArgListContext(Arg_listContext ctx) { copyFrom(ctx); }
	}

	public final Arg_listContext arg_list() throws RecognitionException {
		Arg_listContext _localctx = new Arg_listContext(_ctx, getState());
		enterRule(_localctx, 60, RULE_arg_list);
		int _la;
		try {
			int _alt;
			_localctx = new ArgListContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(438);
			func_arg();
			setState(443);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(439);
					match(COMA);
					setState(440);
					func_arg();
					}
					} 
				}
				setState(445);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,37,_ctx);
			}
			setState(447);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMA) {
				{
				setState(446);
				match(COMA);
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
		public TerminalNode DOS_PUNTOS() { return getToken(VGrammar.DOS_PUNTOS, 0); }
		public TerminalNode ANPERSAND() { return getToken(VGrammar.ANPERSAND, 0); }
		public FuncionArgContext(Func_argContext ctx) { copyFrom(ctx); }
	}

	public final Func_argContext func_arg() throws RecognitionException {
		Func_argContext _localctx = new Func_argContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_func_arg);
		int _la;
		try {
			_localctx = new FuncionArgContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(451);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,39,_ctx) ) {
			case 1:
				{
				setState(449);
				match(ID);
				setState(450);
				match(DOS_PUNTOS);
				}
				break;
			}
			setState(454);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ANPERSAND) {
				{
				setState(453);
				match(ANPERSAND);
				}
			}

			setState(458);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,41,_ctx) ) {
			case 1:
				{
				setState(456);
				id_pattern();
				}
				break;
			case 2:
				{
				setState(457);
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
	public static class MetodoStructContext extends Func_dclContext {
		public TerminalNode RW_FN() { return getToken(VGrammar.RW_FN, 0); }
		public List<TerminalNode> PAR_IZQ() { return getTokens(VGrammar.PAR_IZQ); }
		public TerminalNode PAR_IZQ(int i) {
			return getToken(VGrammar.PAR_IZQ, i);
		}
		public Method_receiverContext method_receiver() {
			return getRuleContext(Method_receiverContext.class,0);
		}
		public List<TerminalNode> PAR_DER() { return getTokens(VGrammar.PAR_DER); }
		public TerminalNode PAR_DER(int i) {
			return getToken(VGrammar.PAR_DER, i);
		}
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public Param_listContext param_list() {
			return getRuleContext(Param_listContext.class,0);
		}
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public MetodoStructContext(Func_dclContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FuncionDecleradaContext extends Func_dclContext {
		public TerminalNode RW_FN() { return getToken(VGrammar.RW_FN, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public Param_listContext param_list() {
			return getRuleContext(Param_listContext.class,0);
		}
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
		enterRule(_localctx, 64, RULE_func_dcl);
		int _la;
		try {
			setState(500);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,48,_ctx) ) {
			case 1:
				_localctx = new FuncionDecleradaContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(460);
				match(RW_FN);
				setState(461);
				match(ID);
				setState(462);
				match(PAR_IZQ);
				setState(464);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(463);
					param_list();
					}
				}

				setState(466);
				match(PAR_DER);
				setState(468);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 18577348462910976L) != 0)) {
					{
					setState(467);
					type();
					}
				}

				setState(470);
				match(LLAVE_IZQ);
				setState(474);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 735770847477776L) != 0)) {
					{
					{
					setState(471);
					stmt();
					}
					}
					setState(476);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(477);
				match(LLAVE_DER);
				}
				break;
			case 2:
				_localctx = new MetodoStructContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(478);
				match(RW_FN);
				setState(479);
				match(PAR_IZQ);
				setState(480);
				method_receiver();
				setState(481);
				match(PAR_DER);
				setState(482);
				match(ID);
				setState(483);
				match(PAR_IZQ);
				setState(485);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==ID) {
					{
					setState(484);
					param_list();
					}
				}

				setState(487);
				match(PAR_DER);
				setState(489);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 18577348462910976L) != 0)) {
					{
					setState(488);
					type();
					}
				}

				setState(491);
				match(LLAVE_IZQ);
				setState(495);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 735770847477776L) != 0)) {
					{
					{
					setState(492);
					stmt();
					}
					}
					setState(497);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(498);
				match(LLAVE_DER);
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
	public static class Method_receiverContext extends ParserRuleContext {
		public Method_receiverContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_method_receiver; }
	 
		public Method_receiverContext() { }
		public void copyFrom(Method_receiverContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class MethodReceiverContext extends Method_receiverContext {
		public List<TerminalNode> ID() { return getTokens(VGrammar.ID); }
		public TerminalNode ID(int i) {
			return getToken(VGrammar.ID, i);
		}
		public TerminalNode OP_MULT() { return getToken(VGrammar.OP_MULT, 0); }
		public MethodReceiverContext(Method_receiverContext ctx) { copyFrom(ctx); }
	}

	public final Method_receiverContext method_receiver() throws RecognitionException {
		Method_receiverContext _localctx = new Method_receiverContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_method_receiver);
		try {
			_localctx = new MethodReceiverContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(502);
			match(ID);
			setState(503);
			match(OP_MULT);
			setState(504);
			match(ID);
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
		public List<TerminalNode> COMA() { return getTokens(VGrammar.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(VGrammar.COMA, i);
		}
		public ParamListContext(Param_listContext ctx) { copyFrom(ctx); }
	}

	public final Param_listContext param_list() throws RecognitionException {
		Param_listContext _localctx = new Param_listContext(_ctx, getState());
		enterRule(_localctx, 68, RULE_param_list);
		int _la;
		try {
			int _alt;
			_localctx = new ParamListContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(506);
			func_param();
			setState(511);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,49,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(507);
					match(COMA);
					setState(508);
					func_param();
					}
					} 
				}
				setState(513);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,49,_ctx);
			}
			setState(515);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMA) {
				{
				setState(514);
				match(COMA);
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
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode RW_INOUT() { return getToken(VGrammar.RW_INOUT, 0); }
		public FuncParamContext(Func_paramContext ctx) { copyFrom(ctx); }
	}

	public final Func_paramContext func_param() throws RecognitionException {
		Func_paramContext _localctx = new Func_paramContext(_ctx, getState());
		enterRule(_localctx, 70, RULE_func_param);
		int _la;
		try {
			_localctx = new FuncParamContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(518);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,51,_ctx) ) {
			case 1:
				{
				setState(517);
				match(ID);
				}
				break;
			}
			setState(520);
			match(ID);
			setState(522);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RW_INOUT) {
				{
				setState(521);
				match(RW_INOUT);
				}
			}

			setState(524);
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
		public TerminalNode RW_STRUCT() { return getToken(VGrammar.RW_STRUCT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
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
		enterRule(_localctx, 72, RULE_strct_dcl);
		int _la;
		try {
			_localctx = new DeclararStructContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(526);
			match(RW_STRUCT);
			setState(527);
			match(ID);
			setState(528);
			match(LLAVE_IZQ);
			setState(532);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while ((((_la) & ~0x3f) == 0 && ((1L << _la) & 18858832029556224L) != 0)) {
				{
				{
				setState(529);
				struct_prop();
				}
				}
				setState(534);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(535);
			match(LLAVE_DER);
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
		public TypeContext type() {
			return getRuleContext(TypeContext.class,0);
		}
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_DECLARACION() { return getToken(VGrammar.OP_DECLARACION, 0); }
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
		public TerminalNode RW_MUTATING() { return getToken(VGrammar.RW_MUTATING, 0); }
		public StructFuncContext(Struct_propContext ctx) { copyFrom(ctx); }
	}

	public final Struct_propContext struct_prop() throws RecognitionException {
		Struct_propContext _localctx = new Struct_propContext(_ctx, getState());
		enterRule(_localctx, 74, RULE_struct_prop);
		int _la;
		try {
			setState(547);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RW_INT:
			case RW_FLOAT64:
			case RW_STRING:
			case RW_BOOL:
			case ID:
			case CORCHETE_IZQ:
				_localctx = new StructAttrContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(537);
				type();
				setState(538);
				match(ID);
				setState(541);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==OP_DECLARACION) {
					{
					setState(539);
					match(OP_DECLARACION);
					setState(540);
					expr(0);
					}
				}

				}
				break;
			case RW_FN:
			case RW_MUTATING:
				_localctx = new StructFuncContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(544);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==RW_MUTATING) {
					{
					setState(543);
					match(RW_MUTATING);
					}
				}

				setState(546);
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
		public TerminalNode CORCHETE_IZQ() { return getToken(VGrammar.CORCHETE_IZQ, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode CORCHETE_DER() { return getToken(VGrammar.CORCHETE_DER, 0); }
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public StructVectorContext(Struct_vectorContext ctx) { copyFrom(ctx); }
	}

	public final Struct_vectorContext struct_vector() throws RecognitionException {
		Struct_vectorContext _localctx = new Struct_vectorContext(_ctx, getState());
		enterRule(_localctx, 76, RULE_struct_vector);
		try {
			_localctx = new StructVectorContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(549);
			match(CORCHETE_IZQ);
			setState(550);
			match(ID);
			setState(551);
			match(CORCHETE_DER);
			setState(552);
			match(PAR_IZQ);
			setState(553);
			match(PAR_DER);
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
	public static class Struct_initContext extends ParserRuleContext {
		public Struct_initContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_init; }
	 
		public Struct_initContext() { }
		public void copyFrom(Struct_initContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructInitContext extends Struct_initContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public Struct_init_listContext struct_init_list() {
			return getRuleContext(Struct_init_listContext.class,0);
		}
		public StructInitContext(Struct_initContext ctx) { copyFrom(ctx); }
	}

	public final Struct_initContext struct_init() throws RecognitionException {
		Struct_initContext _localctx = new Struct_initContext(_ctx, getState());
		enterRule(_localctx, 78, RULE_struct_init);
		int _la;
		try {
			_localctx = new StructInitContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(555);
			match(ID);
			setState(556);
			match(LLAVE_IZQ);
			setState(558);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(557);
				struct_init_list();
				}
			}

			setState(560);
			match(LLAVE_DER);
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
	public static class Struct_init_listContext extends ParserRuleContext {
		public Struct_init_listContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_init_list; }
	 
		public Struct_init_listContext() { }
		public void copyFrom(Struct_init_listContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructInitListContext extends Struct_init_listContext {
		public List<Struct_init_fieldContext> struct_init_field() {
			return getRuleContexts(Struct_init_fieldContext.class);
		}
		public Struct_init_fieldContext struct_init_field(int i) {
			return getRuleContext(Struct_init_fieldContext.class,i);
		}
		public List<TerminalNode> COMA() { return getTokens(VGrammar.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(VGrammar.COMA, i);
		}
		public StructInitListContext(Struct_init_listContext ctx) { copyFrom(ctx); }
	}

	public final Struct_init_listContext struct_init_list() throws RecognitionException {
		Struct_init_listContext _localctx = new Struct_init_listContext(_ctx, getState());
		enterRule(_localctx, 80, RULE_struct_init_list);
		int _la;
		try {
			int _alt;
			_localctx = new StructInitListContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(562);
			struct_init_field();
			setState(567);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,58,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					{
					{
					setState(563);
					match(COMA);
					setState(564);
					struct_init_field();
					}
					} 
				}
				setState(569);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,58,_ctx);
			}
			setState(571);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==COMA) {
				{
				setState(570);
				match(COMA);
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
	public static class Struct_method_callContext extends ParserRuleContext {
		public Struct_method_callContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_method_call; }
	 
		public Struct_method_callContext() { }
		public void copyFrom(Struct_method_callContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructMethodCallContext extends Struct_method_callContext {
		public Id_patternContext id_pattern() {
			return getRuleContext(Id_patternContext.class,0);
		}
		public TerminalNode PUNTO() { return getToken(VGrammar.PUNTO, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public Arg_listContext arg_list() {
			return getRuleContext(Arg_listContext.class,0);
		}
		public StructMethodCallContext(Struct_method_callContext ctx) { copyFrom(ctx); }
	}

	public final Struct_method_callContext struct_method_call() throws RecognitionException {
		Struct_method_callContext _localctx = new Struct_method_callContext(_ctx, getState());
		enterRule(_localctx, 82, RULE_struct_method_call);
		int _la;
		try {
			_localctx = new StructMethodCallContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(573);
			id_pattern();
			setState(574);
			match(PUNTO);
			setState(575);
			match(ID);
			setState(576);
			match(PAR_IZQ);
			setState(578);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 4635892868572766208L) != 0)) {
				{
				setState(577);
				arg_list();
				}
			}

			setState(580);
			match(PAR_DER);
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
	public static class Struct_init_fieldContext extends ParserRuleContext {
		public Struct_init_fieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_init_field; }
	 
		public Struct_init_fieldContext() { }
		public void copyFrom(Struct_init_fieldContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StructInitFieldContext extends Struct_init_fieldContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode DOS_PUNTOS() { return getToken(VGrammar.DOS_PUNTOS, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StructInitFieldContext(Struct_init_fieldContext ctx) { copyFrom(ctx); }
	}

	public final Struct_init_fieldContext struct_init_field() throws RecognitionException {
		Struct_init_fieldContext _localctx = new Struct_init_fieldContext(_ctx, getState());
		enterRule(_localctx, 84, RULE_struct_init_field);
		try {
			_localctx = new StructInitFieldContext(_localctx);
			enterOuterAlt(_localctx, 1);
			{
			setState(582);
			match(ID);
			setState(583);
			match(DOS_PUNTOS);
			setState(584);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 17:
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
		"\u0004\u0001>\u024b\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
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
		"#\u0007#\u0002$\u0007$\u0002%\u0007%\u0002&\u0007&\u0002\'\u0007\'\u0002"+
		"(\u0007(\u0002)\u0007)\u0002*\u0007*\u0001\u0000\u0005\u0000X\b\u0000"+
		"\n\u0000\f\u0000[\t\u0000\u0001\u0000\u0003\u0000^\b\u0000\u0001\u0000"+
		"\u0003\u0000a\b\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0005\u0001i\b\u0001\n\u0001\f\u0001l\t\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002"+
		"\u0001\u0002\u0001\u0002\u0003\u0002|\b\u0002\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003\u0096\b\u0003"+
		"\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003\u009c\b\u0003"+
		"\u0003\u0003\u009e\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004"+
		"\u0005\u0004\u00a4\b\u0004\n\u0004\f\u0004\u00a7\t\u0004\u0003\u0004\u00a9"+
		"\b\u0004\u0001\u0004\u0003\u0004\u00ac\b\u0004\u0001\u0004\u0001\u0004"+
		"\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0004\u0005"+
		"\u00b5\b\u0005\u000b\u0005\f\u0005\u00b6\u0001\u0006\u0001\u0006\u0001"+
		"\u0006\u0001\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\b\u0001\b\u0003\b\u00c3\b\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0001\b\u0001\b\u0001\t\u0001\t\u0001\n\u0001\n\u0001"+
		"\n\u0001\n\u0001\n\u0001\n\u0001\n\u0003\n\u00d8\b\n\u0001\u000b\u0001"+
		"\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001\u000b\u0001"+
		"\u000b\u0003\u000b\u00e2\b\u000b\u0001\f\u0001\f\u0001\f\u0001\f\u0001"+
		"\f\u0001\f\u0003\f\u00ea\b\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e\u0100\b\u000e\u0001\u000f"+
		"\u0001\u000f\u0001\u000f\u0005\u000f\u0105\b\u000f\n\u000f\f\u000f\u0108"+
		"\t\u000f\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0003"+
		"\u0010\u010f\b\u0010\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0003\u0011\u0123\b\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0001"+
		"\u0011\u0001\u0011\u0001\u0011\u0001\u0011\u0005\u0011\u0137\b\u0011\n"+
		"\u0011\f\u0011\u013a\t\u0011\u0001\u0012\u0001\u0012\u0001\u0012\u0005"+
		"\u0012\u013f\b\u0012\n\u0012\f\u0012\u0142\t\u0012\u0001\u0012\u0003\u0012"+
		"\u0145\b\u0012\u0001\u0013\u0001\u0013\u0001\u0013\u0001\u0013\u0005\u0013"+
		"\u014b\b\u0013\n\u0013\f\u0013\u014e\t\u0013\u0001\u0013\u0001\u0013\u0001"+
		"\u0014\u0001\u0014\u0001\u0014\u0005\u0014\u0155\b\u0014\n\u0014\f\u0014"+
		"\u0158\t\u0014\u0001\u0014\u0001\u0014\u0001\u0015\u0001\u0015\u0001\u0015"+
		"\u0001\u0015\u0005\u0015\u0160\b\u0015\n\u0015\f\u0015\u0163\t\u0015\u0001"+
		"\u0015\u0003\u0015\u0166\b\u0015\u0001\u0015\u0001\u0015\u0001\u0016\u0001"+
		"\u0016\u0001\u0016\u0001\u0016\u0005\u0016\u016e\b\u0016\n\u0016\f\u0016"+
		"\u0171\t\u0016\u0001\u0017\u0001\u0017\u0001\u0017\u0005\u0017\u0176\b"+
		"\u0017\n\u0017\f\u0017\u0179\t\u0017\u0001\u0018\u0001\u0018\u0001\u0018"+
		"\u0001\u0018\u0005\u0018\u017f\b\u0018\n\u0018\f\u0018\u0182\t\u0018\u0001"+
		"\u0018\u0001\u0018\u0001\u0019\u0001\u0019\u0001\u0019\u0001\u0019\u0001"+
		"\u0019\u0003\u0019\u018b\b\u0019\u0001\u0019\u0001\u0019\u0005\u0019\u018f"+
		"\b\u0019\n\u0019\f\u0019\u0192\t\u0019\u0001\u0019\u0001\u0019\u0001\u001a"+
		"\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001a\u0001\u001b"+
		"\u0001\u001b\u0001\u001b\u0001\u001b\u0001\u001b\u0005\u001b\u01a1\b\u001b"+
		"\n\u001b\f\u001b\u01a4\t\u001b\u0001\u001b\u0001\u001b\u0001\u001c\u0001"+
		"\u001c\u0003\u001c\u01aa\b\u001c\u0001\u001c\u0001\u001c\u0003\u001c\u01ae"+
		"\b\u001c\u0001\u001d\u0001\u001d\u0001\u001d\u0003\u001d\u01b3\b\u001d"+
		"\u0001\u001d\u0001\u001d\u0001\u001e\u0001\u001e\u0001\u001e\u0005\u001e"+
		"\u01ba\b\u001e\n\u001e\f\u001e\u01bd\t\u001e\u0001\u001e\u0003\u001e\u01c0"+
		"\b\u001e\u0001\u001f\u0001\u001f\u0003\u001f\u01c4\b\u001f\u0001\u001f"+
		"\u0003\u001f\u01c7\b\u001f\u0001\u001f\u0001\u001f\u0003\u001f\u01cb\b"+
		"\u001f\u0001 \u0001 \u0001 \u0001 \u0003 \u01d1\b \u0001 \u0001 \u0003"+
		" \u01d5\b \u0001 \u0001 \u0005 \u01d9\b \n \f \u01dc\t \u0001 \u0001 "+
		"\u0001 \u0001 \u0001 \u0001 \u0001 \u0001 \u0003 \u01e6\b \u0001 \u0001"+
		" \u0003 \u01ea\b \u0001 \u0001 \u0005 \u01ee\b \n \f \u01f1\t \u0001 "+
		"\u0001 \u0003 \u01f5\b \u0001!\u0001!\u0001!\u0001!\u0001\"\u0001\"\u0001"+
		"\"\u0005\"\u01fe\b\"\n\"\f\"\u0201\t\"\u0001\"\u0003\"\u0204\b\"\u0001"+
		"#\u0003#\u0207\b#\u0001#\u0001#\u0003#\u020b\b#\u0001#\u0001#\u0001$\u0001"+
		"$\u0001$\u0001$\u0005$\u0213\b$\n$\f$\u0216\t$\u0001$\u0001$\u0001%\u0001"+
		"%\u0001%\u0001%\u0003%\u021e\b%\u0001%\u0003%\u0221\b%\u0001%\u0003%\u0224"+
		"\b%\u0001&\u0001&\u0001&\u0001&\u0001&\u0001&\u0001\'\u0001\'\u0001\'"+
		"\u0003\'\u022f\b\'\u0001\'\u0001\'\u0001(\u0001(\u0001(\u0005(\u0236\b"+
		"(\n(\f(\u0239\t(\u0001(\u0003(\u023c\b(\u0001)\u0001)\u0001)\u0001)\u0001"+
		")\u0003)\u0243\b)\u0001)\u0001)\u0001*\u0001*\u0001*\u0001*\u0001*\u0000"+
		"\u0001\"+\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016"+
		"\u0018\u001a\u001c\u001e \"$&(*,.02468:<>@BDFHJLNPRT\u0000\u0007\u0001"+
		"\u0000\u0006\u0007\u0001\u0000\u0006\b\u0002\u0000\u0013\u0013\u001f\u001f"+
		"\u0001\u0000\u0014\u0016\u0001\u0000\u0012\u0013\u0001\u0000\u0019\u001c"+
		"\u0001\u0000\u0017\u0018\u0285\u0000Y\u0001\u0000\u0000\u0000\u0002b\u0001"+
		"\u0000\u0000\u0000\u0004{\u0001\u0000\u0000\u0000\u0006\u009d\u0001\u0000"+
		"\u0000\u0000\b\u009f\u0001\u0000\u0000\u0000\n\u00af\u0001\u0000\u0000"+
		"\u0000\f\u00b8\u0001\u0000\u0000\u0000\u000e\u00bc\u0001\u0000\u0000\u0000"+
		"\u0010\u00c2\u0001\u0000\u0000\u0000\u0012\u00ce\u0001\u0000\u0000\u0000"+
		"\u0014\u00d7\u0001\u0000\u0000\u0000\u0016\u00e1\u0001\u0000\u0000\u0000"+
		"\u0018\u00e9\u0001\u0000\u0000\u0000\u001a\u00eb\u0001\u0000\u0000\u0000"+
		"\u001c\u00ff\u0001\u0000\u0000\u0000\u001e\u0101\u0001\u0000\u0000\u0000"+
		" \u010e\u0001\u0000\u0000\u0000\"\u0122\u0001\u0000\u0000\u0000$\u013b"+
		"\u0001\u0000\u0000\u0000&\u0146\u0001\u0000\u0000\u0000(\u0151\u0001\u0000"+
		"\u0000\u0000*\u015b\u0001\u0000\u0000\u0000,\u0169\u0001\u0000\u0000\u0000"+
		".\u0172\u0001\u0000\u0000\u00000\u017a\u0001\u0000\u0000\u00002\u0185"+
		"\u0001\u0000\u0000\u00004\u0195\u0001\u0000\u0000\u00006\u019b\u0001\u0000"+
		"\u0000\u00008\u01ad\u0001\u0000\u0000\u0000:\u01af\u0001\u0000\u0000\u0000"+
		"<\u01b6\u0001\u0000\u0000\u0000>\u01c3\u0001\u0000\u0000\u0000@\u01f4"+
		"\u0001\u0000\u0000\u0000B\u01f6\u0001\u0000\u0000\u0000D\u01fa\u0001\u0000"+
		"\u0000\u0000F\u0206\u0001\u0000\u0000\u0000H\u020e\u0001\u0000\u0000\u0000"+
		"J\u0223\u0001\u0000\u0000\u0000L\u0225\u0001\u0000\u0000\u0000N\u022b"+
		"\u0001\u0000\u0000\u0000P\u0232\u0001\u0000\u0000\u0000R\u023d\u0001\u0000"+
		"\u0000\u0000T\u0246\u0001\u0000\u0000\u0000VX\u0003\u0004\u0002\u0000"+
		"WV\u0001\u0000\u0000\u0000X[\u0001\u0000\u0000\u0000YW\u0001\u0000\u0000"+
		"\u0000YZ\u0001\u0000\u0000\u0000Z]\u0001\u0000\u0000\u0000[Y\u0001\u0000"+
		"\u0000\u0000\\^\u0003\u0002\u0001\u0000]\\\u0001\u0000\u0000\u0000]^\u0001"+
		"\u0000\u0000\u0000^`\u0001\u0000\u0000\u0000_a\u0005\u0000\u0000\u0001"+
		"`_\u0001\u0000\u0000\u0000`a\u0001\u0000\u0000\u0000a\u0001\u0001\u0000"+
		"\u0000\u0000bc\u0005!\u0000\u0000cd\u0005 \u0000\u0000de\u00052\u0000"+
		"\u0000ef\u00053\u0000\u0000fj\u00054\u0000\u0000gi\u0003\u0004\u0002\u0000"+
		"hg\u0001\u0000\u0000\u0000il\u0001\u0000\u0000\u0000jh\u0001\u0000\u0000"+
		"\u0000jk\u0001\u0000\u0000\u0000km\u0001\u0000\u0000\u0000lj\u0001\u0000"+
		"\u0000\u0000mn\u00055\u0000\u0000n\u0003\u0001\u0000\u0000\u0000o|\u0003"+
		"\u0006\u0003\u0000p|\u0003\u001c\u000e\u0000q|\u00038\u001c\u0000r|\u0003"+
		"$\u0012\u0000s|\u0003*\u0015\u0000t|\u00030\u0018\u0000u|\u00032\u0019"+
		"\u0000v|\u00036\u001b\u0000w|\u0003:\u001d\u0000x|\u0003\u000e\u0007\u0000"+
		"y|\u0003@ \u0000z|\u0003H$\u0000{o\u0001\u0000\u0000\u0000{p\u0001\u0000"+
		"\u0000\u0000{q\u0001\u0000\u0000\u0000{r\u0001\u0000\u0000\u0000{s\u0001"+
		"\u0000\u0000\u0000{t\u0001\u0000\u0000\u0000{u\u0001\u0000\u0000\u0000"+
		"{v\u0001\u0000\u0000\u0000{w\u0001\u0000\u0000\u0000{x\u0001\u0000\u0000"+
		"\u0000{y\u0001\u0000\u0000\u0000{z\u0001\u0000\u0000\u0000|\u0005\u0001"+
		"\u0000\u0000\u0000}~\u0005\u0004\u0000\u0000~\u007f\u00051\u0000\u0000"+
		"\u007f\u0080\u0005\u0005\u0000\u0000\u0080\u009e\u0003\"\u0011\u0000\u0081"+
		"\u0082\u00051\u0000\u0000\u0082\u0083\u0005\u0005\u0000\u0000\u0083\u009e"+
		"\u0003\"\u0011\u0000\u0084\u0085\u0005\u0004\u0000\u0000\u0085\u0086\u0005"+
		"1\u0000\u0000\u0086\u0087\u0003\u0014\n\u0000\u0087\u0088\u0005\b\u0000"+
		"\u0000\u0088\u0089\u0003\"\u0011\u0000\u0089\u009e\u0001\u0000\u0000\u0000"+
		"\u008a\u008b\u0005\u0004\u0000\u0000\u008b\u008c\u00051\u0000\u0000\u008c"+
		"\u009e\u0003\u0014\n\u0000\u008d\u008e\u00051\u0000\u0000\u008e\u008f"+
		"\u0003\u0014\n\u0000\u008f\u0090\u0005\b\u0000\u0000\u0090\u0091\u0003"+
		"\"\u0011\u0000\u0091\u009e\u0001\u0000\u0000\u0000\u0092\u0093\u00051"+
		"\u0000\u0000\u0093\u009e\u0003\u0014\n\u0000\u0094\u0096\u0005\u0004\u0000"+
		"\u0000\u0095\u0094\u0001\u0000\u0000\u0000\u0095\u0096\u0001\u0000\u0000"+
		"\u0000\u0096\u0097\u0001\u0000\u0000\u0000\u0097\u0098\u00051\u0000\u0000"+
		"\u0098\u0099\u0005\b\u0000\u0000\u0099\u009b\u0003\u0014\n\u0000\u009a"+
		"\u009c\u0003\"\u0011\u0000\u009b\u009a\u0001\u0000\u0000\u0000\u009b\u009c"+
		"\u0001\u0000\u0000\u0000\u009c\u009e\u0001\u0000\u0000\u0000\u009d}\u0001"+
		"\u0000\u0000\u0000\u009d\u0081\u0001\u0000\u0000\u0000\u009d\u0084\u0001"+
		"\u0000\u0000\u0000\u009d\u008a\u0001\u0000\u0000\u0000\u009d\u008d\u0001"+
		"\u0000\u0000\u0000\u009d\u0092\u0001\u0000\u0000\u0000\u009d\u0095\u0001"+
		"\u0000\u0000\u0000\u009e\u0007\u0001\u0000\u0000\u0000\u009f\u00a8\u0005"+
		"4\u0000\u0000\u00a0\u00a5\u0003\"\u0011\u0000\u00a1\u00a2\u00059\u0000"+
		"\u0000\u00a2\u00a4\u0003\"\u0011\u0000\u00a3\u00a1\u0001\u0000\u0000\u0000"+
		"\u00a4\u00a7\u0001\u0000\u0000\u0000\u00a5\u00a3\u0001\u0000\u0000\u0000"+
		"\u00a5\u00a6\u0001\u0000\u0000\u0000\u00a6\u00a9\u0001\u0000\u0000\u0000"+
		"\u00a7\u00a5\u0001\u0000\u0000\u0000\u00a8\u00a0\u0001\u0000\u0000\u0000"+
		"\u00a8\u00a9\u0001\u0000\u0000\u0000\u00a9\u00ab\u0001\u0000\u0000\u0000"+
		"\u00aa\u00ac\u00059\u0000\u0000\u00ab\u00aa\u0001\u0000\u0000\u0000\u00ab"+
		"\u00ac\u0001\u0000\u0000\u0000\u00ac\u00ad\u0001\u0000\u0000\u0000\u00ad"+
		"\u00ae\u00055\u0000\u0000\u00ae\t\u0001\u0000\u0000\u0000\u00af\u00b4"+
		"\u0003\u001e\u000f\u0000\u00b0\u00b1\u00056\u0000\u0000\u00b1\u00b2\u0003"+
		"\"\u0011\u0000\u00b2\u00b3\u00057\u0000\u0000\u00b3\u00b5\u0001\u0000"+
		"\u0000\u0000\u00b4\u00b0\u0001\u0000\u0000\u0000\u00b5\u00b6\u0001\u0000"+
		"\u0000\u0000\u00b6\u00b4\u0001\u0000\u0000\u0000\u00b6\u00b7\u0001\u0000"+
		"\u0000\u0000\u00b7\u000b\u0001\u0000\u0000\u0000\u00b8\u00b9\u0003\n\u0005"+
		"\u0000\u00b9\u00ba\u00058\u0000\u0000\u00ba\u00bb\u0003\u001e\u000f\u0000"+
		"\u00bb\r\u0001\u0000\u0000\u0000\u00bc\u00bd\u0003\n\u0005\u0000\u00bd"+
		"\u00be\u00058\u0000\u0000\u00be\u00bf\u0003:\u001d\u0000\u00bf\u000f\u0001"+
		"\u0000\u0000\u0000\u00c0\u00c3\u0003\u0016\u000b\u0000\u00c1\u00c3\u0003"+
		"\u0018\f\u0000\u00c2\u00c0\u0001\u0000\u0000\u0000\u00c2\u00c1\u0001\u0000"+
		"\u0000\u0000\u00c3\u00c4\u0001\u0000\u0000\u0000\u00c4\u00c5\u00052\u0000"+
		"\u0000\u00c5\u00c6\u00051\u0000\u0000\u00c6\u00c7\u0005;\u0000\u0000\u00c7"+
		"\u00c8\u0003\"\u0011\u0000\u00c8\u00c9\u00059\u0000\u0000\u00c9\u00ca"+
		"\u00051\u0000\u0000\u00ca\u00cb\u0005;\u0000\u0000\u00cb\u00cc\u0003\""+
		"\u0011\u0000\u00cc\u00cd\u00053\u0000\u0000\u00cd\u0011\u0001\u0000\u0000"+
		"\u0000\u00ce\u00cf\u0005\u0004\u0000\u0000\u00cf\u0013\u0001\u0000\u0000"+
		"\u0000\u00d0\u00d8\u00051\u0000\u0000\u00d1\u00d8\u0005\t\u0000\u0000"+
		"\u00d2\u00d8\u0005\n\u0000\u0000\u00d3\u00d8\u0005\u000b\u0000\u0000\u00d4"+
		"\u00d8\u0005\f\u0000\u0000\u00d5\u00d8\u0003\u0016\u000b\u0000\u00d6\u00d8"+
		"\u0003\u0018\f\u0000\u00d7\u00d0\u0001\u0000\u0000\u0000\u00d7\u00d1\u0001"+
		"\u0000\u0000\u0000\u00d7\u00d2\u0001\u0000\u0000\u0000\u00d7\u00d3\u0001"+
		"\u0000\u0000\u0000\u00d7\u00d4\u0001\u0000\u0000\u0000\u00d7\u00d5\u0001"+
		"\u0000\u0000\u0000\u00d7\u00d6\u0001\u0000\u0000\u0000\u00d8\u0015\u0001"+
		"\u0000\u0000\u0000\u00d9\u00da\u00056\u0000\u0000\u00da\u00db\u00057\u0000"+
		"\u0000\u00db\u00e2\u0003\u0014\n\u0000\u00dc\u00dd\u00056\u0000\u0000"+
		"\u00dd\u00de\u00057\u0000\u0000\u00de\u00df\u00056\u0000\u0000\u00df\u00e0"+
		"\u00057\u0000\u0000\u00e0\u00e2\u0003\u0014\n\u0000\u00e1\u00d9\u0001"+
		"\u0000\u0000\u0000\u00e1\u00dc\u0001\u0000\u0000\u0000\u00e2\u0017\u0001"+
		"\u0000\u0000\u0000\u00e3\u00ea\u0003\u001a\r\u0000\u00e4\u00e5\u00056"+
		"\u0000\u0000\u00e5\u00e6\u00056\u0000\u0000\u00e6\u00e7\u00051\u0000\u0000"+
		"\u00e7\u00e8\u00057\u0000\u0000\u00e8\u00ea\u00057\u0000\u0000\u00e9\u00e3"+
		"\u0001\u0000\u0000\u0000\u00e9\u00e4\u0001\u0000\u0000\u0000\u00ea\u0019"+
		"\u0001\u0000\u0000\u0000\u00eb\u00ec\u00056\u0000\u0000\u00ec\u00ed\u0003"+
		"\u0018\f\u0000\u00ed\u00ee\u00057\u0000\u0000\u00ee\u001b\u0001\u0000"+
		"\u0000\u0000\u00ef\u00f0\u0003\n\u0005\u0000\u00f0\u00f1\u0005\b\u0000"+
		"\u0000\u00f1\u00f2\u0003\"\u0011\u0000\u00f2\u0100\u0001\u0000\u0000\u0000"+
		"\u00f3\u00f4\u0003\u001e\u000f\u0000\u00f4\u00f5\u0005\b\u0000\u0000\u00f5"+
		"\u00f6\u0003\"\u0011\u0000\u00f6\u0100\u0001\u0000\u0000\u0000\u00f7\u00f8"+
		"\u0003\u001e\u000f\u0000\u00f8\u00f9\u0007\u0000\u0000\u0000\u00f9\u00fa"+
		"\u0003\"\u0011\u0000\u00fa\u0100\u0001\u0000\u0000\u0000\u00fb\u00fc\u0003"+
		"\n\u0005\u0000\u00fc\u00fd\u0007\u0001\u0000\u0000\u00fd\u00fe\u0003\""+
		"\u0011\u0000\u00fe\u0100\u0001\u0000\u0000\u0000\u00ff\u00ef\u0001\u0000"+
		"\u0000\u0000\u00ff\u00f3\u0001\u0000\u0000\u0000\u00ff\u00f7\u0001\u0000"+
		"\u0000\u0000\u00ff\u00fb\u0001\u0000\u0000\u0000\u0100\u001d\u0001\u0000"+
		"\u0000\u0000\u0101\u0106\u00051\u0000\u0000\u0102\u0103\u00058\u0000\u0000"+
		"\u0103\u0105\u00051\u0000\u0000\u0104\u0102\u0001\u0000\u0000\u0000\u0105"+
		"\u0108\u0001\u0000\u0000\u0000\u0106\u0104\u0001\u0000\u0000\u0000\u0106"+
		"\u0107\u0001\u0000\u0000\u0000\u0107\u001f\u0001\u0000\u0000\u0000\u0108"+
		"\u0106\u0001\u0000\u0000\u0000\u0109\u010f\u0005\r\u0000\u0000\u010a\u010f"+
		"\u0005\u000e\u0000\u0000\u010b\u010f\u0005\u000f\u0000\u0000\u010c\u010f"+
		"\u0005\u0010\u0000\u0000\u010d\u010f\u0005\u0011\u0000\u0000\u010e\u0109"+
		"\u0001\u0000\u0000\u0000\u010e\u010a\u0001\u0000\u0000\u0000\u010e\u010b"+
		"\u0001\u0000\u0000\u0000\u010e\u010c\u0001\u0000\u0000\u0000\u010e\u010d"+
		"\u0001\u0000\u0000\u0000\u010f!\u0001\u0000\u0000\u0000\u0110\u0111\u0006"+
		"\u0011\uffff\uffff\u0000\u0111\u0112\u00052\u0000\u0000\u0112\u0113\u0003"+
		"\"\u0011\u0000\u0113\u0114\u00053\u0000\u0000\u0114\u0123\u0001\u0000"+
		"\u0000\u0000\u0115\u0123\u0003:\u001d\u0000\u0116\u0123\u0003\u001e\u000f"+
		"\u0000\u0117\u0123\u0003\n\u0005\u0000\u0118\u0123\u0003\f\u0006\u0000"+
		"\u0119\u0123\u0003\u000e\u0007\u0000\u011a\u0123\u0003 \u0010\u0000\u011b"+
		"\u0123\u0003\b\u0004\u0000\u011c\u0123\u0003\u0010\b\u0000\u011d\u0123"+
		"\u0003L&\u0000\u011e\u0123\u0003N\'\u0000\u011f\u0123\u0003R)\u0000\u0120"+
		"\u0121\u0007\u0002\u0000\u0000\u0121\u0123\u0003\"\u0011\u0007\u0122\u0110"+
		"\u0001\u0000\u0000\u0000\u0122\u0115\u0001\u0000\u0000\u0000\u0122\u0116"+
		"\u0001\u0000\u0000\u0000\u0122\u0117\u0001\u0000\u0000\u0000\u0122\u0118"+
		"\u0001\u0000\u0000\u0000\u0122\u0119\u0001\u0000\u0000\u0000\u0122\u011a"+
		"\u0001\u0000\u0000\u0000\u0122\u011b\u0001\u0000\u0000\u0000\u0122\u011c"+
		"\u0001\u0000\u0000\u0000\u0122\u011d\u0001\u0000\u0000\u0000\u0122\u011e"+
		"\u0001\u0000\u0000\u0000\u0122\u011f\u0001\u0000\u0000\u0000\u0122\u0120"+
		"\u0001\u0000\u0000\u0000\u0123\u0138\u0001\u0000\u0000\u0000\u0124\u0125"+
		"\n\u0006\u0000\u0000\u0125\u0126\u0007\u0003\u0000\u0000\u0126\u0137\u0003"+
		"\"\u0011\u0007\u0127\u0128\n\u0005\u0000\u0000\u0128\u0129\u0007\u0004"+
		"\u0000\u0000\u0129\u0137\u0003\"\u0011\u0006\u012a\u012b\n\u0004\u0000"+
		"\u0000\u012b\u012c\u0007\u0005\u0000\u0000\u012c\u0137\u0003\"\u0011\u0005"+
		"\u012d\u012e\n\u0003\u0000\u0000\u012e\u012f\u0007\u0006\u0000\u0000\u012f"+
		"\u0137\u0003\"\u0011\u0004\u0130\u0131\n\u0002\u0000\u0000\u0131\u0132"+
		"\u0005\u001d\u0000\u0000\u0132\u0137\u0003\"\u0011\u0003\u0133\u0134\n"+
		"\u0001\u0000\u0000\u0134\u0135\u0005\u001e\u0000\u0000\u0135\u0137\u0003"+
		"\"\u0011\u0002\u0136\u0124\u0001\u0000\u0000\u0000\u0136\u0127\u0001\u0000"+
		"\u0000\u0000\u0136\u012a\u0001\u0000\u0000\u0000\u0136\u012d\u0001\u0000"+
		"\u0000\u0000\u0136\u0130\u0001\u0000\u0000\u0000\u0136\u0133\u0001\u0000"+
		"\u0000\u0000\u0137\u013a\u0001\u0000\u0000\u0000\u0138\u0136\u0001\u0000"+
		"\u0000\u0000\u0138\u0139\u0001\u0000\u0000\u0000\u0139#\u0001\u0000\u0000"+
		"\u0000\u013a\u0138\u0001\u0000\u0000\u0000\u013b\u0140\u0003&\u0013\u0000"+
		"\u013c\u013d\u0005$\u0000\u0000\u013d\u013f\u0003&\u0013\u0000\u013e\u013c"+
		"\u0001\u0000\u0000\u0000\u013f\u0142\u0001\u0000\u0000\u0000\u0140\u013e"+
		"\u0001\u0000\u0000\u0000\u0140\u0141\u0001\u0000\u0000\u0000\u0141\u0144"+
		"\u0001\u0000\u0000\u0000\u0142\u0140\u0001\u0000\u0000\u0000\u0143\u0145"+
		"\u0003(\u0014\u0000\u0144\u0143\u0001\u0000\u0000\u0000\u0144\u0145\u0001"+
		"\u0000\u0000\u0000\u0145%\u0001\u0000\u0000\u0000\u0146\u0147\u0005#\u0000"+
		"\u0000\u0147\u0148\u0003\"\u0011\u0000\u0148\u014c\u00054\u0000\u0000"+
		"\u0149\u014b\u0003\u0004\u0002\u0000\u014a\u0149\u0001\u0000\u0000\u0000"+
		"\u014b\u014e\u0001\u0000\u0000\u0000\u014c\u014a\u0001\u0000\u0000\u0000"+
		"\u014c\u014d\u0001\u0000\u0000\u0000\u014d\u014f\u0001\u0000\u0000\u0000"+
		"\u014e\u014c\u0001\u0000\u0000\u0000\u014f\u0150\u00055\u0000\u0000\u0150"+
		"\'\u0001\u0000\u0000\u0000\u0151\u0152\u0005$\u0000\u0000\u0152\u0156"+
		"\u00054\u0000\u0000\u0153\u0155\u0003\u0004\u0002\u0000\u0154\u0153\u0001"+
		"\u0000\u0000\u0000\u0155\u0158\u0001\u0000\u0000\u0000\u0156\u0154\u0001"+
		"\u0000\u0000\u0000\u0156\u0157\u0001\u0000\u0000\u0000\u0157\u0159\u0001"+
		"\u0000\u0000\u0000\u0158\u0156\u0001\u0000\u0000\u0000\u0159\u015a\u0005"+
		"5\u0000\u0000\u015a)\u0001\u0000\u0000\u0000\u015b\u015c\u0005%\u0000"+
		"\u0000\u015c\u015d\u0003\"\u0011\u0000\u015d\u0161\u00054\u0000\u0000"+
		"\u015e\u0160\u0003,\u0016\u0000\u015f\u015e\u0001\u0000\u0000\u0000\u0160"+
		"\u0163\u0001\u0000\u0000\u0000\u0161\u015f\u0001\u0000\u0000\u0000\u0161"+
		"\u0162\u0001\u0000\u0000\u0000\u0162\u0165\u0001\u0000\u0000\u0000\u0163"+
		"\u0161\u0001\u0000\u0000\u0000\u0164\u0166\u0003.\u0017\u0000\u0165\u0164"+
		"\u0001\u0000\u0000\u0000\u0165\u0166\u0001\u0000\u0000\u0000\u0166\u0167"+
		"\u0001\u0000\u0000\u0000\u0167\u0168\u00055\u0000\u0000\u0168+\u0001\u0000"+
		"\u0000\u0000\u0169\u016a\u0005&\u0000\u0000\u016a\u016b\u0003\"\u0011"+
		"\u0000\u016b\u016f\u0005;\u0000\u0000\u016c\u016e\u0003\u0004\u0002\u0000"+
		"\u016d\u016c\u0001\u0000\u0000\u0000\u016e\u0171\u0001\u0000\u0000\u0000"+
		"\u016f\u016d\u0001\u0000\u0000\u0000\u016f\u0170\u0001\u0000\u0000\u0000"+
		"\u0170-\u0001\u0000\u0000\u0000\u0171\u016f\u0001\u0000\u0000\u0000\u0172"+
		"\u0173\u0005\'\u0000\u0000\u0173\u0177\u0005;\u0000\u0000\u0174\u0176"+
		"\u0003\u0004\u0002\u0000\u0175\u0174\u0001\u0000\u0000\u0000\u0176\u0179"+
		"\u0001\u0000\u0000\u0000\u0177\u0175\u0001\u0000\u0000\u0000\u0177\u0178"+
		"\u0001\u0000\u0000\u0000\u0178/\u0001\u0000\u0000\u0000\u0179\u0177\u0001"+
		"\u0000\u0000\u0000\u017a\u017b\u0005(\u0000\u0000\u017b\u017c\u0003\""+
		"\u0011\u0000\u017c\u0180\u00054\u0000\u0000\u017d\u017f\u0003\u0004\u0002"+
		"\u0000\u017e\u017d\u0001\u0000\u0000\u0000\u017f\u0182\u0001\u0000\u0000"+
		"\u0000\u0180\u017e\u0001\u0000\u0000\u0000\u0180\u0181\u0001\u0000\u0000"+
		"\u0000\u0181\u0183\u0001\u0000\u0000\u0000\u0182\u0180\u0001\u0000\u0000"+
		"\u0000\u0183\u0184\u00055\u0000\u0000\u01841\u0001\u0000\u0000\u0000\u0185"+
		"\u0186\u0005(\u0000\u0000\u0186\u0187\u00051\u0000\u0000\u0187\u018a\u0005"+
		"-\u0000\u0000\u0188\u018b\u0003\"\u0011\u0000\u0189\u018b\u00034\u001a"+
		"\u0000\u018a\u0188\u0001\u0000\u0000\u0000\u018a\u0189\u0001\u0000\u0000"+
		"\u0000\u018b\u018c\u0001\u0000\u0000\u0000\u018c\u0190\u00054\u0000\u0000"+
		"\u018d\u018f\u0003\u0004\u0002\u0000\u018e\u018d\u0001\u0000\u0000\u0000"+
		"\u018f\u0192\u0001\u0000\u0000\u0000\u0190\u018e\u0001\u0000\u0000\u0000"+
		"\u0190\u0191\u0001\u0000\u0000\u0000\u0191\u0193\u0001\u0000\u0000\u0000"+
		"\u0192\u0190\u0001\u0000\u0000\u0000\u0193\u0194\u00055\u0000\u0000\u0194"+
		"3\u0001\u0000\u0000\u0000\u0195\u0196\u0003\"\u0011\u0000\u0196\u0197"+
		"\u00058\u0000\u0000\u0197\u0198\u00058\u0000\u0000\u0198\u0199\u00058"+
		"\u0000\u0000\u0199\u019a\u0003\"\u0011\u0000\u019a5\u0001\u0000\u0000"+
		"\u0000\u019b\u019c\u0005/\u0000\u0000\u019c\u019d\u0003\"\u0011\u0000"+
		"\u019d\u019e\u0005$\u0000\u0000\u019e\u01a2\u00054\u0000\u0000\u019f\u01a1"+
		"\u0003\u0004\u0002\u0000\u01a0\u019f\u0001\u0000\u0000\u0000\u01a1\u01a4"+
		"\u0001\u0000\u0000\u0000\u01a2\u01a0\u0001\u0000\u0000\u0000\u01a2\u01a3"+
		"\u0001\u0000\u0000\u0000\u01a3\u01a5\u0001\u0000\u0000\u0000\u01a4\u01a2"+
		"\u0001\u0000\u0000\u0000\u01a5\u01a6\u00055\u0000\u0000\u01a67\u0001\u0000"+
		"\u0000\u0000\u01a7\u01a9\u0005,\u0000\u0000\u01a8\u01aa\u0003\"\u0011"+
		"\u0000\u01a9\u01a8\u0001\u0000\u0000\u0000\u01a9\u01aa\u0001\u0000\u0000"+
		"\u0000\u01aa\u01ae\u0001\u0000\u0000\u0000\u01ab\u01ae\u0005*\u0000\u0000"+
		"\u01ac\u01ae\u0005+\u0000\u0000\u01ad\u01a7\u0001\u0000\u0000\u0000\u01ad"+
		"\u01ab\u0001\u0000\u0000\u0000\u01ad\u01ac\u0001\u0000\u0000\u0000\u01ae"+
		"9\u0001\u0000\u0000\u0000\u01af\u01b0\u0003\u001e\u000f\u0000\u01b0\u01b2"+
		"\u00052\u0000\u0000\u01b1\u01b3\u0003<\u001e\u0000\u01b2\u01b1\u0001\u0000"+
		"\u0000\u0000\u01b2\u01b3\u0001\u0000\u0000\u0000\u01b3\u01b4\u0001\u0000"+
		"\u0000\u0000\u01b4\u01b5\u00053\u0000\u0000\u01b5;\u0001\u0000\u0000\u0000"+
		"\u01b6\u01bb\u0003>\u001f\u0000\u01b7\u01b8\u00059\u0000\u0000\u01b8\u01ba"+
		"\u0003>\u001f\u0000\u01b9\u01b7\u0001\u0000\u0000\u0000\u01ba\u01bd\u0001"+
		"\u0000\u0000\u0000\u01bb\u01b9\u0001\u0000\u0000\u0000\u01bb\u01bc\u0001"+
		"\u0000\u0000\u0000\u01bc\u01bf\u0001\u0000\u0000\u0000\u01bd\u01bb\u0001"+
		"\u0000\u0000\u0000\u01be\u01c0\u00059\u0000\u0000\u01bf\u01be\u0001\u0000"+
		"\u0000\u0000\u01bf\u01c0\u0001\u0000\u0000\u0000\u01c0=\u0001\u0000\u0000"+
		"\u0000\u01c1\u01c2\u00051\u0000\u0000\u01c2\u01c4\u0005;\u0000\u0000\u01c3"+
		"\u01c1\u0001\u0000\u0000\u0000\u01c3\u01c4\u0001\u0000\u0000\u0000\u01c4"+
		"\u01c6\u0001\u0000\u0000\u0000\u01c5\u01c7\u0005>\u0000\u0000\u01c6\u01c5"+
		"\u0001\u0000\u0000\u0000\u01c6\u01c7\u0001\u0000\u0000\u0000\u01c7\u01ca"+
		"\u0001\u0000\u0000\u0000\u01c8\u01cb\u0003\u001e\u000f\u0000\u01c9\u01cb"+
		"\u0003\"\u0011\u0000\u01ca\u01c8\u0001\u0000\u0000\u0000\u01ca\u01c9\u0001"+
		"\u0000\u0000\u0000\u01cb?\u0001\u0000\u0000\u0000\u01cc\u01cd\u0005!\u0000"+
		"\u0000\u01cd\u01ce\u00051\u0000\u0000\u01ce\u01d0\u00052\u0000\u0000\u01cf"+
		"\u01d1\u0003D\"\u0000\u01d0\u01cf\u0001\u0000\u0000\u0000\u01d0\u01d1"+
		"\u0001\u0000\u0000\u0000\u01d1\u01d2\u0001\u0000\u0000\u0000\u01d2\u01d4"+
		"\u00053\u0000\u0000\u01d3\u01d5\u0003\u0014\n\u0000\u01d4\u01d3\u0001"+
		"\u0000\u0000\u0000\u01d4\u01d5\u0001\u0000\u0000\u0000\u01d5\u01d6\u0001"+
		"\u0000\u0000\u0000\u01d6\u01da\u00054\u0000\u0000\u01d7\u01d9\u0003\u0004"+
		"\u0002\u0000\u01d8\u01d7\u0001\u0000\u0000\u0000\u01d9\u01dc\u0001\u0000"+
		"\u0000\u0000\u01da\u01d8\u0001\u0000\u0000\u0000\u01da\u01db\u0001\u0000"+
		"\u0000\u0000\u01db\u01dd\u0001\u0000\u0000\u0000\u01dc\u01da\u0001\u0000"+
		"\u0000\u0000\u01dd\u01f5\u00055\u0000\u0000\u01de\u01df\u0005!\u0000\u0000"+
		"\u01df\u01e0\u00052\u0000\u0000\u01e0\u01e1\u0003B!\u0000\u01e1\u01e2"+
		"\u00053\u0000\u0000\u01e2\u01e3\u00051\u0000\u0000\u01e3\u01e5\u00052"+
		"\u0000\u0000\u01e4\u01e6\u0003D\"\u0000\u01e5\u01e4\u0001\u0000\u0000"+
		"\u0000\u01e5\u01e6\u0001\u0000\u0000\u0000\u01e6\u01e7\u0001\u0000\u0000"+
		"\u0000\u01e7\u01e9\u00053\u0000\u0000\u01e8\u01ea\u0003\u0014\n\u0000"+
		"\u01e9\u01e8\u0001\u0000\u0000\u0000\u01e9\u01ea\u0001\u0000\u0000\u0000"+
		"\u01ea\u01eb\u0001\u0000\u0000\u0000\u01eb\u01ef\u00054\u0000\u0000\u01ec"+
		"\u01ee\u0003\u0004\u0002\u0000\u01ed\u01ec\u0001\u0000\u0000\u0000\u01ee"+
		"\u01f1\u0001\u0000\u0000\u0000\u01ef\u01ed\u0001\u0000\u0000\u0000\u01ef"+
		"\u01f0\u0001\u0000\u0000\u0000\u01f0\u01f2\u0001\u0000\u0000\u0000\u01f1"+
		"\u01ef\u0001\u0000\u0000\u0000\u01f2\u01f3\u00055\u0000\u0000\u01f3\u01f5"+
		"\u0001\u0000\u0000\u0000\u01f4\u01cc\u0001\u0000\u0000\u0000\u01f4\u01de"+
		"\u0001\u0000\u0000\u0000\u01f5A\u0001\u0000\u0000\u0000\u01f6\u01f7\u0005"+
		"1\u0000\u0000\u01f7\u01f8\u0005\u0014\u0000\u0000\u01f8\u01f9\u00051\u0000"+
		"\u0000\u01f9C\u0001\u0000\u0000\u0000\u01fa\u01ff\u0003F#\u0000\u01fb"+
		"\u01fc\u00059\u0000\u0000\u01fc\u01fe\u0003F#\u0000\u01fd\u01fb\u0001"+
		"\u0000\u0000\u0000\u01fe\u0201\u0001\u0000\u0000\u0000\u01ff\u01fd\u0001"+
		"\u0000\u0000\u0000\u01ff\u0200\u0001\u0000\u0000\u0000\u0200\u0203\u0001"+
		"\u0000\u0000\u0000\u0201\u01ff\u0001\u0000\u0000\u0000\u0202\u0204\u0005"+
		"9\u0000\u0000\u0203\u0202\u0001\u0000\u0000\u0000\u0203\u0204\u0001\u0000"+
		"\u0000\u0000\u0204E\u0001\u0000\u0000\u0000\u0205\u0207\u00051\u0000\u0000"+
		"\u0206\u0205\u0001\u0000\u0000\u0000\u0206\u0207\u0001\u0000\u0000\u0000"+
		"\u0207\u0208\u0001\u0000\u0000\u0000\u0208\u020a\u00051\u0000\u0000\u0209"+
		"\u020b\u0005.\u0000\u0000\u020a\u0209\u0001\u0000\u0000\u0000\u020a\u020b"+
		"\u0001\u0000\u0000\u0000\u020b\u020c\u0001\u0000\u0000\u0000\u020c\u020d"+
		"\u0003\u0014\n\u0000\u020dG\u0001\u0000\u0000\u0000\u020e\u020f\u0005"+
		"\"\u0000\u0000\u020f\u0210\u00051\u0000\u0000\u0210\u0214\u00054\u0000"+
		"\u0000\u0211\u0213\u0003J%\u0000\u0212\u0211\u0001\u0000\u0000\u0000\u0213"+
		"\u0216\u0001\u0000\u0000\u0000\u0214\u0212\u0001\u0000\u0000\u0000\u0214"+
		"\u0215\u0001\u0000\u0000\u0000\u0215\u0217\u0001\u0000\u0000\u0000\u0216"+
		"\u0214\u0001\u0000\u0000\u0000\u0217\u0218\u00055\u0000\u0000\u0218I\u0001"+
		"\u0000\u0000\u0000\u0219\u021a\u0003\u0014\n\u0000\u021a\u021d\u00051"+
		"\u0000\u0000\u021b\u021c\u0005\b\u0000\u0000\u021c\u021e\u0003\"\u0011"+
		"\u0000\u021d\u021b\u0001\u0000\u0000\u0000\u021d\u021e\u0001\u0000\u0000"+
		"\u0000\u021e\u0224\u0001\u0000\u0000\u0000\u021f\u0221\u00050\u0000\u0000"+
		"\u0220\u021f\u0001\u0000\u0000\u0000\u0220\u0221\u0001\u0000\u0000\u0000"+
		"\u0221\u0222\u0001\u0000\u0000\u0000\u0222\u0224\u0003@ \u0000\u0223\u0219"+
		"\u0001\u0000\u0000\u0000\u0223\u0220\u0001\u0000\u0000\u0000\u0224K\u0001"+
		"\u0000\u0000\u0000\u0225\u0226\u00056\u0000\u0000\u0226\u0227\u00051\u0000"+
		"\u0000\u0227\u0228\u00057\u0000\u0000\u0228\u0229\u00052\u0000\u0000\u0229"+
		"\u022a\u00053\u0000\u0000\u022aM\u0001\u0000\u0000\u0000\u022b\u022c\u0005"+
		"1\u0000\u0000\u022c\u022e\u00054\u0000\u0000\u022d\u022f\u0003P(\u0000"+
		"\u022e\u022d\u0001\u0000\u0000\u0000\u022e\u022f\u0001\u0000\u0000\u0000"+
		"\u022f\u0230\u0001\u0000\u0000\u0000\u0230\u0231\u00055\u0000\u0000\u0231"+
		"O\u0001\u0000\u0000\u0000\u0232\u0237\u0003T*\u0000\u0233\u0234\u0005"+
		"9\u0000\u0000\u0234\u0236\u0003T*\u0000\u0235\u0233\u0001\u0000\u0000"+
		"\u0000\u0236\u0239\u0001\u0000\u0000\u0000\u0237\u0235\u0001\u0000\u0000"+
		"\u0000\u0237\u0238\u0001\u0000\u0000\u0000\u0238\u023b\u0001\u0000\u0000"+
		"\u0000\u0239\u0237\u0001\u0000\u0000\u0000\u023a\u023c\u00059\u0000\u0000"+
		"\u023b\u023a\u0001\u0000\u0000\u0000\u023b\u023c\u0001\u0000\u0000\u0000"+
		"\u023cQ\u0001\u0000\u0000\u0000\u023d\u023e\u0003\u001e\u000f\u0000\u023e"+
		"\u023f\u00058\u0000\u0000\u023f\u0240\u00051\u0000\u0000\u0240\u0242\u0005"+
		"2\u0000\u0000\u0241\u0243\u0003<\u001e\u0000\u0242\u0241\u0001\u0000\u0000"+
		"\u0000\u0242\u0243\u0001\u0000\u0000\u0000\u0243\u0244\u0001\u0000\u0000"+
		"\u0000\u0244\u0245\u00053\u0000\u0000\u0245S\u0001\u0000\u0000\u0000\u0246"+
		"\u0247\u00051\u0000\u0000\u0247\u0248\u0005;\u0000\u0000\u0248\u0249\u0003"+
		"\"\u0011\u0000\u0249U\u0001\u0000\u0000\u0000=Y]`j{\u0095\u009b\u009d"+
		"\u00a5\u00a8\u00ab\u00b6\u00c2\u00d7\u00e1\u00e9\u00ff\u0106\u010e\u0122"+
		"\u0136\u0138\u0140\u0144\u014c\u0156\u0161\u0165\u016f\u0177\u0180\u018a"+
		"\u0190\u01a2\u01a9\u01ad\u01b2\u01bb\u01bf\u01c3\u01c6\u01ca\u01d0\u01d4"+
		"\u01da\u01e5\u01e9\u01ef\u01f4\u01ff\u0203\u0206\u020a\u0214\u021d\u0220"+
		"\u0223\u022e\u0237\u023b\u0242";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}