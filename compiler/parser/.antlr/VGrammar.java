// Generated from /home/jairogo/Documentos/GitHub/OLC2_Proyecto1/compiler/parser/VGrammar.g4 by ANTLR 4.13.1
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
		RW_MAIN=1, RW_FN=2, RW_MUT=3, RW_STRUCT=4, RW_IF=5, RW_ELSE=6, RW_SWITCH=7, 
		RW_CASE=8, RW_DEFAULT=9, RW_FOR=10, RW_IN=11, RW_BREAK=12, RW_CONTINUE=13, 
		RW_RETURN=14, RW_TRUE=15, RW_FALSE=16, RW_NIL=17, RW_INT=18, RW_FLOAT64=19, 
		RW_STRING=20, RW_BOOL=21, RW_RUNE=22, RW_PRINT=23, RW_PRINTLN=24, RW_ATOI=25, 
		RW_PARSEFLOAT=26, RW_TYPEOF=27, RW_APPEND=28, RW_LEN=29, RW_JOIN=30, RW_INDEXOF=31, 
		OP_SUMA=32, OP_RESTA=33, OP_MULT=34, OP_DIV=35, OP_MOD=36, OP_ASSIGN=37, 
		OP_ADD_ASSIGN=38, OP_SUB_ASSIGN=39, OP_IGUAL=40, OP_DIFERENTE=41, OP_MENORQ=42, 
		OP_MAYORQ=43, OP_MENORIGUAL=44, OP_MAYORIGUAL=45, OP_AND=46, OP_OR=47, 
		OP_NOT=48, OP_DECLARATION=49, PUNTO=50, COMA=51, PUNTO_Y_COMA=52, DOS_PUNTOS=53, 
		PAR_IZQ=54, PAR_DER=55, LLAVE_IZQ=56, LLAVE_DER=57, CORCHETE_IZQ=58, CORCHETE_DER=59, 
		INT_LITERAL=60, FLOAT_LITERAL=61, STRING_LITERAL=62, RUNE_LITERAL=63, 
		ID=64, WS=65, COMMENT=66, MULTILINE_COMMENT=67;
	public static final int
		RULE_programa = 0, RULE_stmt = 1, RULE_decl_stmt = 2, RULE_assign_stmt = 3, 
		RULE_expr = 4, RULE_literal = 5, RULE_flow_stmt = 6, RULE_func_call = 7, 
		RULE_if_stmt = 8, RULE_switch_stmt = 9, RULE_switch_case = 10, RULE_default_case = 11, 
		RULE_for_stmt = 12, RULE_func_decl = 13, RULE_param = 14, RULE_struct_decl = 15, 
		RULE_struct_field = 16;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "stmt", "decl_stmt", "assign_stmt", "expr", "literal", "flow_stmt", 
			"func_call", "if_stmt", "switch_stmt", "switch_case", "default_case", 
			"for_stmt", "func_decl", "param", "struct_decl", "struct_field"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'main'", "'fn'", "'mut'", "'struct'", "'if'", "'else'", "'switch'", 
			"'case'", "'default'", "'for'", "'in'", "'break'", "'continue'", "'return'", 
			"'true'", "'false'", "'nil'", "'int'", "'float64'", "'string'", "'bool'", 
			"'rune'", "'print'", "'println'", "'Atoi'", "'parseFloat'", "'typeof'", 
			"'append'", "'len'", "'join'", "'indexOf'", "'+'", "'-'", "'*'", "'/'", 
			"'%'", "'='", "'+='", "'-='", "'=='", "'!='", "'<'", "'>'", "'<='", "'>='", 
			"'&&'", "'||'", "'!'", "':='", "'.'", "','", "';'", "':'", "'('", "')'", 
			"'{'", "'}'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "RW_MAIN", "RW_FN", "RW_MUT", "RW_STRUCT", "RW_IF", "RW_ELSE", 
			"RW_SWITCH", "RW_CASE", "RW_DEFAULT", "RW_FOR", "RW_IN", "RW_BREAK", 
			"RW_CONTINUE", "RW_RETURN", "RW_TRUE", "RW_FALSE", "RW_NIL", "RW_INT", 
			"RW_FLOAT64", "RW_STRING", "RW_BOOL", "RW_RUNE", "RW_PRINT", "RW_PRINTLN", 
			"RW_ATOI", "RW_PARSEFLOAT", "RW_TYPEOF", "RW_APPEND", "RW_LEN", "RW_JOIN", 
			"RW_INDEXOF", "OP_SUMA", "OP_RESTA", "OP_MULT", "OP_DIV", "OP_MOD", "OP_ASSIGN", 
			"OP_ADD_ASSIGN", "OP_SUB_ASSIGN", "OP_IGUAL", "OP_DIFERENTE", "OP_MENORQ", 
			"OP_MAYORQ", "OP_MENORIGUAL", "OP_MAYORIGUAL", "OP_AND", "OP_OR", "OP_NOT", 
			"OP_DECLARATION", "PUNTO", "COMA", "PUNTO_Y_COMA", "DOS_PUNTOS", "PAR_IZQ", 
			"PAR_DER", "LLAVE_IZQ", "LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER", 
			"INT_LITERAL", "FLOAT_LITERAL", "STRING_LITERAL", "RUNE_LITERAL", "ID", 
			"WS", "COMMENT", "MULTILINE_COMMENT"
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
	public static class ProgramaContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(VGrammar.EOF, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public ProgramaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_programa; }
	}

	public final ProgramaContext programa() throws RecognitionException {
		ProgramaContext _localctx = new ProgramaContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_programa);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(37);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 2)) & ~0x3f) == 0 && ((1L << (_la - 2)) & 4611686018433686831L) != 0)) {
				{
				{
				setState(34);
				stmt();
				}
				}
				setState(39);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(40);
			match(EOF);
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
		public StmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stmt; }
	 
		public StmtContext() { }
		public void copyFrom(StmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtFuncDeclContext extends StmtContext {
		public Func_declContext func_decl() {
			return getRuleContext(Func_declContext.class,0);
		}
		public StmtFuncDeclContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtDeclContext extends StmtContext {
		public Decl_stmtContext decl_stmt() {
			return getRuleContext(Decl_stmtContext.class,0);
		}
		public TerminalNode PUNTO_Y_COMA() { return getToken(VGrammar.PUNTO_Y_COMA, 0); }
		public StmtDeclContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtStructDeclContext extends StmtContext {
		public Struct_declContext struct_decl() {
			return getRuleContext(Struct_declContext.class,0);
		}
		public StmtStructDeclContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtSwitchContext extends StmtContext {
		public Switch_stmtContext switch_stmt() {
			return getRuleContext(Switch_stmtContext.class,0);
		}
		public StmtSwitchContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtForContext extends StmtContext {
		public For_stmtContext for_stmt() {
			return getRuleContext(For_stmtContext.class,0);
		}
		public StmtForContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtAssignContext extends StmtContext {
		public Assign_stmtContext assign_stmt() {
			return getRuleContext(Assign_stmtContext.class,0);
		}
		public TerminalNode PUNTO_Y_COMA() { return getToken(VGrammar.PUNTO_Y_COMA, 0); }
		public StmtAssignContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtIfContext extends StmtContext {
		public If_stmtContext if_stmt() {
			return getRuleContext(If_stmtContext.class,0);
		}
		public StmtIfContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtFlowContext extends StmtContext {
		public Flow_stmtContext flow_stmt() {
			return getRuleContext(Flow_stmtContext.class,0);
		}
		public StmtFlowContext(StmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtFuncCallContext extends StmtContext {
		public Func_callContext func_call() {
			return getRuleContext(Func_callContext.class,0);
		}
		public TerminalNode PUNTO_Y_COMA() { return getToken(VGrammar.PUNTO_Y_COMA, 0); }
		public StmtFuncCallContext(StmtContext ctx) { copyFrom(ctx); }
	}

	public final StmtContext stmt() throws RecognitionException {
		StmtContext _localctx = new StmtContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_stmt);
		int _la;
		try {
			setState(60);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				_localctx = new StmtDeclContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(42);
				decl_stmt();
				setState(44);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTO_Y_COMA) {
					{
					setState(43);
					match(PUNTO_Y_COMA);
					}
				}

				}
				break;
			case 2:
				_localctx = new StmtAssignContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(46);
				assign_stmt();
				setState(48);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTO_Y_COMA) {
					{
					setState(47);
					match(PUNTO_Y_COMA);
					}
				}

				}
				break;
			case 3:
				_localctx = new StmtFlowContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(50);
				flow_stmt();
				}
				break;
			case 4:
				_localctx = new StmtFuncCallContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(51);
				func_call();
				setState(53);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==PUNTO_Y_COMA) {
					{
					setState(52);
					match(PUNTO_Y_COMA);
					}
				}

				}
				break;
			case 5:
				_localctx = new StmtIfContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(55);
				if_stmt();
				}
				break;
			case 6:
				_localctx = new StmtSwitchContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(56);
				switch_stmt();
				}
				break;
			case 7:
				_localctx = new StmtForContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(57);
				for_stmt();
				}
				break;
			case 8:
				_localctx = new StmtFuncDeclContext(_localctx);
				enterOuterAlt(_localctx, 8);
				{
				setState(58);
				func_decl();
				}
				break;
			case 9:
				_localctx = new StmtStructDeclContext(_localctx);
				enterOuterAlt(_localctx, 9);
				{
				setState(59);
				struct_decl();
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
	public static class TypedDeclBoolContext extends Decl_stmtContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_BOOL() { return getToken(VGrammar.RW_BOOL, 0); }
		public TypedDeclBoolContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedDeclFloatContext extends Decl_stmtContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_FLOAT64() { return getToken(VGrammar.RW_FLOAT64, 0); }
		public TypedDeclFloatContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedDeclStringContext extends Decl_stmtContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_STRING() { return getToken(VGrammar.RW_STRING, 0); }
		public TypedDeclStringContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class InferDeclContext extends Decl_stmtContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_DECLARATION() { return getToken(VGrammar.OP_DECLARATION, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public InferDeclContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedDeclRuneContext extends Decl_stmtContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_RUNE() { return getToken(VGrammar.RW_RUNE, 0); }
		public TypedDeclRuneContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class TypedDeclIntContext extends Decl_stmtContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_INT() { return getToken(VGrammar.RW_INT, 0); }
		public TypedDeclIntContext(Decl_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Decl_stmtContext decl_stmt() throws RecognitionException {
		Decl_stmtContext _localctx = new Decl_stmtContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_decl_stmt);
		try {
			setState(81);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				_localctx = new InferDeclContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(62);
				match(RW_MUT);
				setState(63);
				match(ID);
				setState(64);
				match(OP_DECLARATION);
				setState(65);
				expr(0);
				}
				break;
			case 2:
				_localctx = new TypedDeclIntContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(66);
				match(RW_MUT);
				setState(67);
				match(ID);
				setState(68);
				match(RW_INT);
				}
				break;
			case 3:
				_localctx = new TypedDeclStringContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(69);
				match(RW_MUT);
				setState(70);
				match(ID);
				setState(71);
				match(RW_STRING);
				}
				break;
			case 4:
				_localctx = new TypedDeclFloatContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(72);
				match(RW_MUT);
				setState(73);
				match(ID);
				setState(74);
				match(RW_FLOAT64);
				}
				break;
			case 5:
				_localctx = new TypedDeclBoolContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(75);
				match(RW_MUT);
				setState(76);
				match(ID);
				setState(77);
				match(RW_BOOL);
				}
				break;
			case 6:
				_localctx = new TypedDeclRuneContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(78);
				match(RW_MUT);
				setState(79);
				match(ID);
				setState(80);
				match(RW_RUNE);
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
	public static class AssignSimpleContext extends Assign_stmtContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_ASSIGN() { return getToken(VGrammar.OP_ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AssignSimpleContext(Assign_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignAddContext extends Assign_stmtContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_ADD_ASSIGN() { return getToken(VGrammar.OP_ADD_ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AssignAddContext(Assign_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AssignSubContext extends Assign_stmtContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode OP_SUB_ASSIGN() { return getToken(VGrammar.OP_SUB_ASSIGN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public AssignSubContext(Assign_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Assign_stmtContext assign_stmt() throws RecognitionException {
		Assign_stmtContext _localctx = new Assign_stmtContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_assign_stmt);
		try {
			setState(92);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				_localctx = new AssignSimpleContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(83);
				match(ID);
				setState(84);
				match(OP_ASSIGN);
				setState(85);
				expr(0);
				}
				break;
			case 2:
				_localctx = new AssignAddContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(86);
				match(ID);
				setState(87);
				match(OP_ADD_ASSIGN);
				setState(88);
				expr(0);
				}
				break;
			case 3:
				_localctx = new AssignSubContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(89);
				match(ID);
				setState(90);
				match(OP_SUB_ASSIGN);
				setState(91);
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
	public static class ExprGreaterContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_MAYORQ() { return getToken(VGrammar.OP_MAYORQ, 0); }
		public ExprGreaterContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprParenContext extends ExprContext {
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public ExprParenContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprNotEqualContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_DIFERENTE() { return getToken(VGrammar.OP_DIFERENTE, 0); }
		public ExprNotEqualContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprLessEqContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_MENORIGUAL() { return getToken(VGrammar.OP_MENORIGUAL, 0); }
		public ExprLessEqContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprLessContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_MENORQ() { return getToken(VGrammar.OP_MENORQ, 0); }
		public ExprLessContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprNotContext extends ExprContext {
		public TerminalNode OP_NOT() { return getToken(VGrammar.OP_NOT, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public ExprNotContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprDivContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_DIV() { return getToken(VGrammar.OP_DIV, 0); }
		public ExprDivContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprAndContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_AND() { return getToken(VGrammar.OP_AND, 0); }
		public ExprAndContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprGreaterEqContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_MAYORIGUAL() { return getToken(VGrammar.OP_MAYORIGUAL, 0); }
		public ExprGreaterEqContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprOrContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_OR() { return getToken(VGrammar.OP_OR, 0); }
		public ExprOrContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprSubContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_RESTA() { return getToken(VGrammar.OP_RESTA, 0); }
		public ExprSubContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprLiteralContext extends ExprContext {
		public LiteralContext literal() {
			return getRuleContext(LiteralContext.class,0);
		}
		public ExprLiteralContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprMulContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_MULT() { return getToken(VGrammar.OP_MULT, 0); }
		public ExprMulContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprEqualContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_IGUAL() { return getToken(VGrammar.OP_IGUAL, 0); }
		public ExprEqualContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprIdContext extends ExprContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public ExprIdContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprAddContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_SUMA() { return getToken(VGrammar.OP_SUMA, 0); }
		public ExprAddContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ExprModContext extends ExprContext {
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_MOD() { return getToken(VGrammar.OP_MOD, 0); }
		public ExprModContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 8;
		enterRecursionRule(_localctx, 8, RULE_expr, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(103);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RW_TRUE:
			case RW_FALSE:
			case RW_NIL:
			case INT_LITERAL:
			case FLOAT_LITERAL:
			case STRING_LITERAL:
			case RUNE_LITERAL:
				{
				_localctx = new ExprLiteralContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(95);
				literal();
				}
				break;
			case ID:
				{
				_localctx = new ExprIdContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(96);
				match(ID);
				}
				break;
			case OP_NOT:
				{
				_localctx = new ExprNotContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(97);
				match(OP_NOT);
				setState(98);
				expr(2);
				}
				break;
			case PAR_IZQ:
				{
				_localctx = new ExprParenContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(99);
				match(PAR_IZQ);
				setState(100);
				expr(0);
				setState(101);
				match(PAR_DER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(146);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(144);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
					case 1:
						{
						_localctx = new ExprAddContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(105);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(106);
						match(OP_SUMA);
						setState(107);
						expr(16);
						}
						break;
					case 2:
						{
						_localctx = new ExprSubContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(108);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(109);
						match(OP_RESTA);
						setState(110);
						expr(15);
						}
						break;
					case 3:
						{
						_localctx = new ExprMulContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(111);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(112);
						match(OP_MULT);
						setState(113);
						expr(14);
						}
						break;
					case 4:
						{
						_localctx = new ExprDivContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(114);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(115);
						match(OP_DIV);
						setState(116);
						expr(13);
						}
						break;
					case 5:
						{
						_localctx = new ExprModContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(117);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(118);
						match(OP_MOD);
						setState(119);
						expr(12);
						}
						break;
					case 6:
						{
						_localctx = new ExprEqualContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(120);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(121);
						match(OP_IGUAL);
						setState(122);
						expr(11);
						}
						break;
					case 7:
						{
						_localctx = new ExprNotEqualContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(123);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(124);
						match(OP_DIFERENTE);
						setState(125);
						expr(10);
						}
						break;
					case 8:
						{
						_localctx = new ExprLessContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(126);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(127);
						match(OP_MENORQ);
						setState(128);
						expr(9);
						}
						break;
					case 9:
						{
						_localctx = new ExprGreaterContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(129);
						if (!(precpred(_ctx, 7))) throw new FailedPredicateException(this, "precpred(_ctx, 7)");
						setState(130);
						match(OP_MAYORQ);
						setState(131);
						expr(8);
						}
						break;
					case 10:
						{
						_localctx = new ExprLessEqContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(132);
						if (!(precpred(_ctx, 6))) throw new FailedPredicateException(this, "precpred(_ctx, 6)");
						setState(133);
						match(OP_MENORIGUAL);
						setState(134);
						expr(7);
						}
						break;
					case 11:
						{
						_localctx = new ExprGreaterEqContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(135);
						if (!(precpred(_ctx, 5))) throw new FailedPredicateException(this, "precpred(_ctx, 5)");
						setState(136);
						match(OP_MAYORIGUAL);
						setState(137);
						expr(6);
						}
						break;
					case 12:
						{
						_localctx = new ExprAndContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(138);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(139);
						match(OP_AND);
						setState(140);
						expr(5);
						}
						break;
					case 13:
						{
						_localctx = new ExprOrContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(141);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(142);
						match(OP_OR);
						setState(143);
						expr(4);
						}
						break;
					}
					} 
				}
				setState(148);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,9,_ctx);
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
	public static class LitStringContext extends LiteralContext {
		public TerminalNode STRING_LITERAL() { return getToken(VGrammar.STRING_LITERAL, 0); }
		public LitStringContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LitRuneContext extends LiteralContext {
		public TerminalNode RUNE_LITERAL() { return getToken(VGrammar.RUNE_LITERAL, 0); }
		public LitRuneContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LitFloatContext extends LiteralContext {
		public TerminalNode FLOAT_LITERAL() { return getToken(VGrammar.FLOAT_LITERAL, 0); }
		public LitFloatContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LitTrueContext extends LiteralContext {
		public TerminalNode RW_TRUE() { return getToken(VGrammar.RW_TRUE, 0); }
		public LitTrueContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LitIntContext extends LiteralContext {
		public TerminalNode INT_LITERAL() { return getToken(VGrammar.INT_LITERAL, 0); }
		public LitIntContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LitNilContext extends LiteralContext {
		public TerminalNode RW_NIL() { return getToken(VGrammar.RW_NIL, 0); }
		public LitNilContext(LiteralContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class LitFalseContext extends LiteralContext {
		public TerminalNode RW_FALSE() { return getToken(VGrammar.RW_FALSE, 0); }
		public LitFalseContext(LiteralContext ctx) { copyFrom(ctx); }
	}

	public final LiteralContext literal() throws RecognitionException {
		LiteralContext _localctx = new LiteralContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_literal);
		try {
			setState(156);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_LITERAL:
				_localctx = new LitIntContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(149);
				match(INT_LITERAL);
				}
				break;
			case FLOAT_LITERAL:
				_localctx = new LitFloatContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(150);
				match(FLOAT_LITERAL);
				}
				break;
			case STRING_LITERAL:
				_localctx = new LitStringContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(151);
				match(STRING_LITERAL);
				}
				break;
			case RUNE_LITERAL:
				_localctx = new LitRuneContext(_localctx);
				enterOuterAlt(_localctx, 4);
				{
				setState(152);
				match(RUNE_LITERAL);
				}
				break;
			case RW_TRUE:
				_localctx = new LitTrueContext(_localctx);
				enterOuterAlt(_localctx, 5);
				{
				setState(153);
				match(RW_TRUE);
				}
				break;
			case RW_FALSE:
				_localctx = new LitFalseContext(_localctx);
				enterOuterAlt(_localctx, 6);
				{
				setState(154);
				match(RW_FALSE);
				}
				break;
			case RW_NIL:
				_localctx = new LitNilContext(_localctx);
				enterOuterAlt(_localctx, 7);
				{
				setState(155);
				match(RW_NIL);
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
	public static class Flow_stmtContext extends ParserRuleContext {
		public Flow_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_flow_stmt; }
	 
		public Flow_stmtContext() { }
		public void copyFrom(Flow_stmtContext ctx) {
			super.copyFrom(ctx);
		}
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtReturnContext extends Flow_stmtContext {
		public TerminalNode RW_RETURN() { return getToken(VGrammar.RW_RETURN, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public StmtReturnContext(Flow_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtContinueContext extends Flow_stmtContext {
		public TerminalNode RW_CONTINUE() { return getToken(VGrammar.RW_CONTINUE, 0); }
		public StmtContinueContext(Flow_stmtContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class StmtBreakContext extends Flow_stmtContext {
		public TerminalNode RW_BREAK() { return getToken(VGrammar.RW_BREAK, 0); }
		public StmtBreakContext(Flow_stmtContext ctx) { copyFrom(ctx); }
	}

	public final Flow_stmtContext flow_stmt() throws RecognitionException {
		Flow_stmtContext _localctx = new Flow_stmtContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_flow_stmt);
		try {
			setState(164);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RW_BREAK:
				_localctx = new StmtBreakContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(158);
				match(RW_BREAK);
				}
				break;
			case RW_CONTINUE:
				_localctx = new StmtContinueContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(159);
				match(RW_CONTINUE);
				}
				break;
			case RW_RETURN:
				_localctx = new StmtReturnContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(160);
				match(RW_RETURN);
				setState(162);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
				case 1:
					{
					setState(161);
					expr(0);
					}
					break;
				}
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
	public static class PrintCallContext extends Func_callContext {
		public TerminalNode RW_PRINT() { return getToken(VGrammar.RW_PRINT, 0); }
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
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
		public PrintCallContext(Func_callContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class FunctionCallContext extends Func_callContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
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
		public FunctionCallContext(Func_callContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class PrintLnCallContext extends Func_callContext {
		public TerminalNode RW_PRINTLN() { return getToken(VGrammar.RW_PRINTLN, 0); }
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
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
		public PrintLnCallContext(Func_callContext ctx) { copyFrom(ctx); }
	}

	public final Func_callContext func_call() throws RecognitionException {
		Func_callContext _localctx = new Func_callContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_func_call);
		int _la;
		try {
			setState(205);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case ID:
				_localctx = new FunctionCallContext(_localctx);
				enterOuterAlt(_localctx, 1);
				{
				setState(166);
				match(ID);
				setState(167);
				match(PAR_IZQ);
				setState(176);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 1091273880502279L) != 0)) {
					{
					setState(168);
					expr(0);
					setState(173);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMA) {
						{
						{
						setState(169);
						match(COMA);
						setState(170);
						expr(0);
						}
						}
						setState(175);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(178);
				match(PAR_DER);
				}
				break;
			case RW_PRINT:
				_localctx = new PrintCallContext(_localctx);
				enterOuterAlt(_localctx, 2);
				{
				setState(179);
				match(RW_PRINT);
				setState(180);
				match(PAR_IZQ);
				setState(189);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 1091273880502279L) != 0)) {
					{
					setState(181);
					expr(0);
					setState(186);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMA) {
						{
						{
						setState(182);
						match(COMA);
						setState(183);
						expr(0);
						}
						}
						setState(188);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(191);
				match(PAR_DER);
				}
				break;
			case RW_PRINTLN:
				_localctx = new PrintLnCallContext(_localctx);
				enterOuterAlt(_localctx, 3);
				{
				setState(192);
				match(RW_PRINTLN);
				setState(193);
				match(PAR_IZQ);
				setState(202);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (((((_la - 15)) & ~0x3f) == 0 && ((1L << (_la - 15)) & 1091273880502279L) != 0)) {
					{
					setState(194);
					expr(0);
					setState(199);
					_errHandler.sync(this);
					_la = _input.LA(1);
					while (_la==COMA) {
						{
						{
						setState(195);
						match(COMA);
						setState(196);
						expr(0);
						}
						}
						setState(201);
						_errHandler.sync(this);
						_la = _input.LA(1);
					}
					}
				}

				setState(204);
				match(PAR_DER);
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
	public static class If_stmtContext extends ParserRuleContext {
		public TerminalNode RW_IF() { return getToken(VGrammar.RW_IF, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public List<TerminalNode> LLAVE_IZQ() { return getTokens(VGrammar.LLAVE_IZQ); }
		public TerminalNode LLAVE_IZQ(int i) {
			return getToken(VGrammar.LLAVE_IZQ, i);
		}
		public List<TerminalNode> LLAVE_DER() { return getTokens(VGrammar.LLAVE_DER); }
		public TerminalNode LLAVE_DER(int i) {
			return getToken(VGrammar.LLAVE_DER, i);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public TerminalNode RW_ELSE() { return getToken(VGrammar.RW_ELSE, 0); }
		public If_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_if_stmt; }
	}

	public final If_stmtContext if_stmt() throws RecognitionException {
		If_stmtContext _localctx = new If_stmtContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_if_stmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(207);
			match(RW_IF);
			setState(208);
			expr(0);
			setState(209);
			match(LLAVE_IZQ);
			setState(213);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 2)) & ~0x3f) == 0 && ((1L << (_la - 2)) & 4611686018433686831L) != 0)) {
				{
				{
				setState(210);
				stmt();
				}
				}
				setState(215);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(216);
			match(LLAVE_DER);
			setState(226);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RW_ELSE) {
				{
				setState(217);
				match(RW_ELSE);
				setState(218);
				match(LLAVE_IZQ);
				setState(222);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (((((_la - 2)) & ~0x3f) == 0 && ((1L << (_la - 2)) & 4611686018433686831L) != 0)) {
					{
					{
					setState(219);
					stmt();
					}
					}
					setState(224);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				setState(225);
				match(LLAVE_DER);
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
	public static class Switch_stmtContext extends ParserRuleContext {
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
		public Switch_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_stmt; }
	}

	public final Switch_stmtContext switch_stmt() throws RecognitionException {
		Switch_stmtContext _localctx = new Switch_stmtContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_switch_stmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(228);
			match(RW_SWITCH);
			setState(229);
			expr(0);
			setState(230);
			match(LLAVE_IZQ);
			setState(234);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==RW_CASE) {
				{
				{
				setState(231);
				switch_case();
				}
				}
				setState(236);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(238);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==RW_DEFAULT) {
				{
				setState(237);
				default_case();
				}
			}

			setState(240);
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
		public Switch_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_switch_case; }
	}

	public final Switch_caseContext switch_case() throws RecognitionException {
		Switch_caseContext _localctx = new Switch_caseContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_switch_case);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(242);
			match(RW_CASE);
			setState(243);
			expr(0);
			setState(244);
			match(DOS_PUNTOS);
			setState(248);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 2)) & ~0x3f) == 0 && ((1L << (_la - 2)) & 4611686018433686831L) != 0)) {
				{
				{
				setState(245);
				stmt();
				}
				}
				setState(250);
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
		public TerminalNode RW_DEFAULT() { return getToken(VGrammar.RW_DEFAULT, 0); }
		public TerminalNode DOS_PUNTOS() { return getToken(VGrammar.DOS_PUNTOS, 0); }
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public Default_caseContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_default_case; }
	}

	public final Default_caseContext default_case() throws RecognitionException {
		Default_caseContext _localctx = new Default_caseContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_default_case);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(251);
			match(RW_DEFAULT);
			setState(252);
			match(DOS_PUNTOS);
			setState(256);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 2)) & ~0x3f) == 0 && ((1L << (_la - 2)) & 4611686018433686831L) != 0)) {
				{
				{
				setState(253);
				stmt();
				}
				}
				setState(258);
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
	public static class For_stmtContext extends ParserRuleContext {
		public TerminalNode RW_FOR() { return getToken(VGrammar.RW_FOR, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_IN() { return getToken(VGrammar.RW_IN, 0); }
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
		public For_stmtContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_for_stmt; }
	}

	public final For_stmtContext for_stmt() throws RecognitionException {
		For_stmtContext _localctx = new For_stmtContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_for_stmt);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(259);
			match(RW_FOR);
			setState(260);
			match(ID);
			setState(261);
			match(RW_IN);
			setState(262);
			expr(0);
			setState(263);
			match(LLAVE_IZQ);
			setState(267);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 2)) & ~0x3f) == 0 && ((1L << (_la - 2)) & 4611686018433686831L) != 0)) {
				{
				{
				setState(264);
				stmt();
				}
				}
				setState(269);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(270);
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
	public static class Func_declContext extends ParserRuleContext {
		public TerminalNode RW_FN() { return getToken(VGrammar.RW_FN, 0); }
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_MAIN() { return getToken(VGrammar.RW_MAIN, 0); }
		public List<ParamContext> param() {
			return getRuleContexts(ParamContext.class);
		}
		public ParamContext param(int i) {
			return getRuleContext(ParamContext.class,i);
		}
		public List<StmtContext> stmt() {
			return getRuleContexts(StmtContext.class);
		}
		public StmtContext stmt(int i) {
			return getRuleContext(StmtContext.class,i);
		}
		public TerminalNode RW_INT() { return getToken(VGrammar.RW_INT, 0); }
		public TerminalNode RW_STRING() { return getToken(VGrammar.RW_STRING, 0); }
		public TerminalNode RW_FLOAT64() { return getToken(VGrammar.RW_FLOAT64, 0); }
		public TerminalNode RW_BOOL() { return getToken(VGrammar.RW_BOOL, 0); }
		public TerminalNode RW_RUNE() { return getToken(VGrammar.RW_RUNE, 0); }
		public List<TerminalNode> COMA() { return getTokens(VGrammar.COMA); }
		public TerminalNode COMA(int i) {
			return getToken(VGrammar.COMA, i);
		}
		public Func_declContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_func_decl; }
	}

	public final Func_declContext func_decl() throws RecognitionException {
		Func_declContext _localctx = new Func_declContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_func_decl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(272);
			match(RW_FN);
			setState(273);
			_la = _input.LA(1);
			if ( !(_la==RW_MAIN || _la==ID) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			setState(274);
			match(PAR_IZQ);
			setState(283);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==ID) {
				{
				setState(275);
				param();
				setState(280);
				_errHandler.sync(this);
				_la = _input.LA(1);
				while (_la==COMA) {
					{
					{
					setState(276);
					match(COMA);
					setState(277);
					param();
					}
					}
					setState(282);
					_errHandler.sync(this);
					_la = _input.LA(1);
				}
				}
			}

			setState(285);
			match(PAR_DER);
			setState(287);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & 8126464L) != 0)) {
				{
				setState(286);
				_la = _input.LA(1);
				if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & 8126464L) != 0)) ) {
				_errHandler.recoverInline(this);
				}
				else {
					if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
					_errHandler.reportMatch(this);
					consume();
				}
				}
			}

			setState(289);
			match(LLAVE_IZQ);
			setState(293);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 2)) & ~0x3f) == 0 && ((1L << (_la - 2)) & 4611686018433686831L) != 0)) {
				{
				{
				setState(290);
				stmt();
				}
				}
				setState(295);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(296);
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
	public static class ParamContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_INT() { return getToken(VGrammar.RW_INT, 0); }
		public TerminalNode RW_STRING() { return getToken(VGrammar.RW_STRING, 0); }
		public TerminalNode RW_FLOAT64() { return getToken(VGrammar.RW_FLOAT64, 0); }
		public TerminalNode RW_BOOL() { return getToken(VGrammar.RW_BOOL, 0); }
		public TerminalNode RW_RUNE() { return getToken(VGrammar.RW_RUNE, 0); }
		public ParamContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_param; }
	}

	public final ParamContext param() throws RecognitionException {
		ParamContext _localctx = new ParamContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_param);
		try {
			setState(308);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,32,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(298);
				match(ID);
				setState(299);
				match(RW_INT);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(300);
				match(ID);
				setState(301);
				match(RW_STRING);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(302);
				match(ID);
				setState(303);
				match(RW_FLOAT64);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(304);
				match(ID);
				setState(305);
				match(RW_BOOL);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(306);
				match(ID);
				setState(307);
				match(RW_RUNE);
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
	public static class Struct_declContext extends ParserRuleContext {
		public TerminalNode RW_STRUCT() { return getToken(VGrammar.RW_STRUCT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode LLAVE_IZQ() { return getToken(VGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(VGrammar.LLAVE_DER, 0); }
		public List<Struct_fieldContext> struct_field() {
			return getRuleContexts(Struct_fieldContext.class);
		}
		public Struct_fieldContext struct_field(int i) {
			return getRuleContext(Struct_fieldContext.class,i);
		}
		public Struct_declContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_decl; }
	}

	public final Struct_declContext struct_decl() throws RecognitionException {
		Struct_declContext _localctx = new Struct_declContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_struct_decl);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(310);
			match(RW_STRUCT);
			setState(311);
			match(ID);
			setState(312);
			match(LLAVE_IZQ);
			setState(316);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==RW_MUT) {
				{
				{
				setState(313);
				struct_field();
				}
				}
				setState(318);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(319);
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
	public static class Struct_fieldContext extends ParserRuleContext {
		public TerminalNode RW_MUT() { return getToken(VGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(VGrammar.ID, 0); }
		public TerminalNode RW_INT() { return getToken(VGrammar.RW_INT, 0); }
		public TerminalNode RW_STRING() { return getToken(VGrammar.RW_STRING, 0); }
		public TerminalNode RW_FLOAT64() { return getToken(VGrammar.RW_FLOAT64, 0); }
		public TerminalNode RW_BOOL() { return getToken(VGrammar.RW_BOOL, 0); }
		public TerminalNode RW_RUNE() { return getToken(VGrammar.RW_RUNE, 0); }
		public Struct_fieldContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_struct_field; }
	}

	public final Struct_fieldContext struct_field() throws RecognitionException {
		Struct_fieldContext _localctx = new Struct_fieldContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_struct_field);
		try {
			setState(336);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,34,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(321);
				match(RW_MUT);
				setState(322);
				match(ID);
				setState(323);
				match(RW_INT);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(324);
				match(RW_MUT);
				setState(325);
				match(ID);
				setState(326);
				match(RW_STRING);
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(327);
				match(RW_MUT);
				setState(328);
				match(ID);
				setState(329);
				match(RW_FLOAT64);
				}
				break;
			case 4:
				enterOuterAlt(_localctx, 4);
				{
				setState(330);
				match(RW_MUT);
				setState(331);
				match(ID);
				setState(332);
				match(RW_BOOL);
				}
				break;
			case 5:
				enterOuterAlt(_localctx, 5);
				{
				setState(333);
				match(RW_MUT);
				setState(334);
				match(ID);
				setState(335);
				match(RW_RUNE);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 4:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 15);
		case 1:
			return precpred(_ctx, 14);
		case 2:
			return precpred(_ctx, 13);
		case 3:
			return precpred(_ctx, 12);
		case 4:
			return precpred(_ctx, 11);
		case 5:
			return precpred(_ctx, 10);
		case 6:
			return precpred(_ctx, 9);
		case 7:
			return precpred(_ctx, 8);
		case 8:
			return precpred(_ctx, 7);
		case 9:
			return precpred(_ctx, 6);
		case 10:
			return precpred(_ctx, 5);
		case 11:
			return precpred(_ctx, 4);
		case 12:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001C\u0153\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002"+
		"\u0002\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002"+
		"\u0005\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0002"+
		"\b\u0007\b\u0002\t\u0007\t\u0002\n\u0007\n\u0002\u000b\u0007\u000b\u0002"+
		"\f\u0007\f\u0002\r\u0007\r\u0002\u000e\u0007\u000e\u0002\u000f\u0007\u000f"+
		"\u0002\u0010\u0007\u0010\u0001\u0000\u0005\u0000$\b\u0000\n\u0000\f\u0000"+
		"\'\t\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0003\u0001"+
		"-\b\u0001\u0001\u0001\u0001\u0001\u0003\u00011\b\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0003\u00016\b\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0003\u0001=\b\u0001\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001"+
		"\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0001\u0002\u0003"+
		"\u0002R\b\u0002\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0003\u0003]\b"+
		"\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0003\u0004h\b\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0004\u0005\u0004\u0091\b\u0004\n\u0004\f\u0004"+
		"\u0094\t\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005"+
		"\u0001\u0005\u0001\u0005\u0003\u0005\u009d\b\u0005\u0001\u0006\u0001\u0006"+
		"\u0001\u0006\u0001\u0006\u0003\u0006\u00a3\b\u0006\u0003\u0006\u00a5\b"+
		"\u0006\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0005"+
		"\u0007\u00ac\b\u0007\n\u0007\f\u0007\u00af\t\u0007\u0003\u0007\u00b1\b"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0005\u0007\u00b9\b\u0007\n\u0007\f\u0007\u00bc\t\u0007\u0003\u0007"+
		"\u00be\b\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007"+
		"\u0001\u0007\u0005\u0007\u00c6\b\u0007\n\u0007\f\u0007\u00c9\t\u0007\u0003"+
		"\u0007\u00cb\b\u0007\u0001\u0007\u0003\u0007\u00ce\b\u0007\u0001\b\u0001"+
		"\b\u0001\b\u0001\b\u0005\b\u00d4\b\b\n\b\f\b\u00d7\t\b\u0001\b\u0001\b"+
		"\u0001\b\u0001\b\u0005\b\u00dd\b\b\n\b\f\b\u00e0\t\b\u0001\b\u0003\b\u00e3"+
		"\b\b\u0001\t\u0001\t\u0001\t\u0001\t\u0005\t\u00e9\b\t\n\t\f\t\u00ec\t"+
		"\t\u0001\t\u0003\t\u00ef\b\t\u0001\t\u0001\t\u0001\n\u0001\n\u0001\n\u0001"+
		"\n\u0005\n\u00f7\b\n\n\n\f\n\u00fa\t\n\u0001\u000b\u0001\u000b\u0001\u000b"+
		"\u0005\u000b\u00ff\b\u000b\n\u000b\f\u000b\u0102\t\u000b\u0001\f\u0001"+
		"\f\u0001\f\u0001\f\u0001\f\u0001\f\u0005\f\u010a\b\f\n\f\f\f\u010d\t\f"+
		"\u0001\f\u0001\f\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0001\r\u0005"+
		"\r\u0117\b\r\n\r\f\r\u011a\t\r\u0003\r\u011c\b\r\u0001\r\u0001\r\u0003"+
		"\r\u0120\b\r\u0001\r\u0001\r\u0005\r\u0124\b\r\n\r\f\r\u0127\t\r\u0001"+
		"\r\u0001\r\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e"+
		"\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0001\u000e\u0003\u000e"+
		"\u0135\b\u000e\u0001\u000f\u0001\u000f\u0001\u000f\u0001\u000f\u0005\u000f"+
		"\u013b\b\u000f\n\u000f\f\u000f\u013e\t\u000f\u0001\u000f\u0001\u000f\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001\u0010\u0001"+
		"\u0010\u0001\u0010\u0001\u0010\u0003\u0010\u0151\b\u0010\u0001\u0010\u0000"+
		"\u0001\b\u0011\u0000\u0002\u0004\u0006\b\n\f\u000e\u0010\u0012\u0014\u0016"+
		"\u0018\u001a\u001c\u001e \u0000\u0002\u0002\u0000\u0001\u0001@@\u0001"+
		"\u0000\u0012\u0016\u018a\u0000%\u0001\u0000\u0000\u0000\u0002<\u0001\u0000"+
		"\u0000\u0000\u0004Q\u0001\u0000\u0000\u0000\u0006\\\u0001\u0000\u0000"+
		"\u0000\bg\u0001\u0000\u0000\u0000\n\u009c\u0001\u0000\u0000\u0000\f\u00a4"+
		"\u0001\u0000\u0000\u0000\u000e\u00cd\u0001\u0000\u0000\u0000\u0010\u00cf"+
		"\u0001\u0000\u0000\u0000\u0012\u00e4\u0001\u0000\u0000\u0000\u0014\u00f2"+
		"\u0001\u0000\u0000\u0000\u0016\u00fb\u0001\u0000\u0000\u0000\u0018\u0103"+
		"\u0001\u0000\u0000\u0000\u001a\u0110\u0001\u0000\u0000\u0000\u001c\u0134"+
		"\u0001\u0000\u0000\u0000\u001e\u0136\u0001\u0000\u0000\u0000 \u0150\u0001"+
		"\u0000\u0000\u0000\"$\u0003\u0002\u0001\u0000#\"\u0001\u0000\u0000\u0000"+
		"$\'\u0001\u0000\u0000\u0000%#\u0001\u0000\u0000\u0000%&\u0001\u0000\u0000"+
		"\u0000&(\u0001\u0000\u0000\u0000\'%\u0001\u0000\u0000\u0000()\u0005\u0000"+
		"\u0000\u0001)\u0001\u0001\u0000\u0000\u0000*,\u0003\u0004\u0002\u0000"+
		"+-\u00054\u0000\u0000,+\u0001\u0000\u0000\u0000,-\u0001\u0000\u0000\u0000"+
		"-=\u0001\u0000\u0000\u0000.0\u0003\u0006\u0003\u0000/1\u00054\u0000\u0000"+
		"0/\u0001\u0000\u0000\u000001\u0001\u0000\u0000\u00001=\u0001\u0000\u0000"+
		"\u00002=\u0003\f\u0006\u000035\u0003\u000e\u0007\u000046\u00054\u0000"+
		"\u000054\u0001\u0000\u0000\u000056\u0001\u0000\u0000\u00006=\u0001\u0000"+
		"\u0000\u00007=\u0003\u0010\b\u00008=\u0003\u0012\t\u00009=\u0003\u0018"+
		"\f\u0000:=\u0003\u001a\r\u0000;=\u0003\u001e\u000f\u0000<*\u0001\u0000"+
		"\u0000\u0000<.\u0001\u0000\u0000\u0000<2\u0001\u0000\u0000\u0000<3\u0001"+
		"\u0000\u0000\u0000<7\u0001\u0000\u0000\u0000<8\u0001\u0000\u0000\u0000"+
		"<9\u0001\u0000\u0000\u0000<:\u0001\u0000\u0000\u0000<;\u0001\u0000\u0000"+
		"\u0000=\u0003\u0001\u0000\u0000\u0000>?\u0005\u0003\u0000\u0000?@\u0005"+
		"@\u0000\u0000@A\u00051\u0000\u0000AR\u0003\b\u0004\u0000BC\u0005\u0003"+
		"\u0000\u0000CD\u0005@\u0000\u0000DR\u0005\u0012\u0000\u0000EF\u0005\u0003"+
		"\u0000\u0000FG\u0005@\u0000\u0000GR\u0005\u0014\u0000\u0000HI\u0005\u0003"+
		"\u0000\u0000IJ\u0005@\u0000\u0000JR\u0005\u0013\u0000\u0000KL\u0005\u0003"+
		"\u0000\u0000LM\u0005@\u0000\u0000MR\u0005\u0015\u0000\u0000NO\u0005\u0003"+
		"\u0000\u0000OP\u0005@\u0000\u0000PR\u0005\u0016\u0000\u0000Q>\u0001\u0000"+
		"\u0000\u0000QB\u0001\u0000\u0000\u0000QE\u0001\u0000\u0000\u0000QH\u0001"+
		"\u0000\u0000\u0000QK\u0001\u0000\u0000\u0000QN\u0001\u0000\u0000\u0000"+
		"R\u0005\u0001\u0000\u0000\u0000ST\u0005@\u0000\u0000TU\u0005%\u0000\u0000"+
		"U]\u0003\b\u0004\u0000VW\u0005@\u0000\u0000WX\u0005&\u0000\u0000X]\u0003"+
		"\b\u0004\u0000YZ\u0005@\u0000\u0000Z[\u0005\'\u0000\u0000[]\u0003\b\u0004"+
		"\u0000\\S\u0001\u0000\u0000\u0000\\V\u0001\u0000\u0000\u0000\\Y\u0001"+
		"\u0000\u0000\u0000]\u0007\u0001\u0000\u0000\u0000^_\u0006\u0004\uffff"+
		"\uffff\u0000_h\u0003\n\u0005\u0000`h\u0005@\u0000\u0000ab\u00050\u0000"+
		"\u0000bh\u0003\b\u0004\u0002cd\u00056\u0000\u0000de\u0003\b\u0004\u0000"+
		"ef\u00057\u0000\u0000fh\u0001\u0000\u0000\u0000g^\u0001\u0000\u0000\u0000"+
		"g`\u0001\u0000\u0000\u0000ga\u0001\u0000\u0000\u0000gc\u0001\u0000\u0000"+
		"\u0000h\u0092\u0001\u0000\u0000\u0000ij\n\u000f\u0000\u0000jk\u0005 \u0000"+
		"\u0000k\u0091\u0003\b\u0004\u0010lm\n\u000e\u0000\u0000mn\u0005!\u0000"+
		"\u0000n\u0091\u0003\b\u0004\u000fop\n\r\u0000\u0000pq\u0005\"\u0000\u0000"+
		"q\u0091\u0003\b\u0004\u000ers\n\f\u0000\u0000st\u0005#\u0000\u0000t\u0091"+
		"\u0003\b\u0004\ruv\n\u000b\u0000\u0000vw\u0005$\u0000\u0000w\u0091\u0003"+
		"\b\u0004\fxy\n\n\u0000\u0000yz\u0005(\u0000\u0000z\u0091\u0003\b\u0004"+
		"\u000b{|\n\t\u0000\u0000|}\u0005)\u0000\u0000}\u0091\u0003\b\u0004\n~"+
		"\u007f\n\b\u0000\u0000\u007f\u0080\u0005*\u0000\u0000\u0080\u0091\u0003"+
		"\b\u0004\t\u0081\u0082\n\u0007\u0000\u0000\u0082\u0083\u0005+\u0000\u0000"+
		"\u0083\u0091\u0003\b\u0004\b\u0084\u0085\n\u0006\u0000\u0000\u0085\u0086"+
		"\u0005,\u0000\u0000\u0086\u0091\u0003\b\u0004\u0007\u0087\u0088\n\u0005"+
		"\u0000\u0000\u0088\u0089\u0005-\u0000\u0000\u0089\u0091\u0003\b\u0004"+
		"\u0006\u008a\u008b\n\u0004\u0000\u0000\u008b\u008c\u0005.\u0000\u0000"+
		"\u008c\u0091\u0003\b\u0004\u0005\u008d\u008e\n\u0003\u0000\u0000\u008e"+
		"\u008f\u0005/\u0000\u0000\u008f\u0091\u0003\b\u0004\u0004\u0090i\u0001"+
		"\u0000\u0000\u0000\u0090l\u0001\u0000\u0000\u0000\u0090o\u0001\u0000\u0000"+
		"\u0000\u0090r\u0001\u0000\u0000\u0000\u0090u\u0001\u0000\u0000\u0000\u0090"+
		"x\u0001\u0000\u0000\u0000\u0090{\u0001\u0000\u0000\u0000\u0090~\u0001"+
		"\u0000\u0000\u0000\u0090\u0081\u0001\u0000\u0000\u0000\u0090\u0084\u0001"+
		"\u0000\u0000\u0000\u0090\u0087\u0001\u0000\u0000\u0000\u0090\u008a\u0001"+
		"\u0000\u0000\u0000\u0090\u008d\u0001\u0000\u0000\u0000\u0091\u0094\u0001"+
		"\u0000\u0000\u0000\u0092\u0090\u0001\u0000\u0000\u0000\u0092\u0093\u0001"+
		"\u0000\u0000\u0000\u0093\t\u0001\u0000\u0000\u0000\u0094\u0092\u0001\u0000"+
		"\u0000\u0000\u0095\u009d\u0005<\u0000\u0000\u0096\u009d\u0005=\u0000\u0000"+
		"\u0097\u009d\u0005>\u0000\u0000\u0098\u009d\u0005?\u0000\u0000\u0099\u009d"+
		"\u0005\u000f\u0000\u0000\u009a\u009d\u0005\u0010\u0000\u0000\u009b\u009d"+
		"\u0005\u0011\u0000\u0000\u009c\u0095\u0001\u0000\u0000\u0000\u009c\u0096"+
		"\u0001\u0000\u0000\u0000\u009c\u0097\u0001\u0000\u0000\u0000\u009c\u0098"+
		"\u0001\u0000\u0000\u0000\u009c\u0099\u0001\u0000\u0000\u0000\u009c\u009a"+
		"\u0001\u0000\u0000\u0000\u009c\u009b\u0001\u0000\u0000\u0000\u009d\u000b"+
		"\u0001\u0000\u0000\u0000\u009e\u00a5\u0005\f\u0000\u0000\u009f\u00a5\u0005"+
		"\r\u0000\u0000\u00a0\u00a2\u0005\u000e\u0000\u0000\u00a1\u00a3\u0003\b"+
		"\u0004\u0000\u00a2\u00a1\u0001\u0000\u0000\u0000\u00a2\u00a3\u0001\u0000"+
		"\u0000\u0000\u00a3\u00a5\u0001\u0000\u0000\u0000\u00a4\u009e\u0001\u0000"+
		"\u0000\u0000\u00a4\u009f\u0001\u0000\u0000\u0000\u00a4\u00a0\u0001\u0000"+
		"\u0000\u0000\u00a5\r\u0001\u0000\u0000\u0000\u00a6\u00a7\u0005@\u0000"+
		"\u0000\u00a7\u00b0\u00056\u0000\u0000\u00a8\u00ad\u0003\b\u0004\u0000"+
		"\u00a9\u00aa\u00053\u0000\u0000\u00aa\u00ac\u0003\b\u0004\u0000\u00ab"+
		"\u00a9\u0001\u0000\u0000\u0000\u00ac\u00af\u0001\u0000\u0000\u0000\u00ad"+
		"\u00ab\u0001\u0000\u0000\u0000\u00ad\u00ae\u0001\u0000\u0000\u0000\u00ae"+
		"\u00b1\u0001\u0000\u0000\u0000\u00af\u00ad\u0001\u0000\u0000\u0000\u00b0"+
		"\u00a8\u0001\u0000\u0000\u0000\u00b0\u00b1\u0001\u0000\u0000\u0000\u00b1"+
		"\u00b2\u0001\u0000\u0000\u0000\u00b2\u00ce\u00057\u0000\u0000\u00b3\u00b4"+
		"\u0005\u0017\u0000\u0000\u00b4\u00bd\u00056\u0000\u0000\u00b5\u00ba\u0003"+
		"\b\u0004\u0000\u00b6\u00b7\u00053\u0000\u0000\u00b7\u00b9\u0003\b\u0004"+
		"\u0000\u00b8\u00b6\u0001\u0000\u0000\u0000\u00b9\u00bc\u0001\u0000\u0000"+
		"\u0000\u00ba\u00b8\u0001\u0000\u0000\u0000\u00ba\u00bb\u0001\u0000\u0000"+
		"\u0000\u00bb\u00be\u0001\u0000\u0000\u0000\u00bc\u00ba\u0001\u0000\u0000"+
		"\u0000\u00bd\u00b5\u0001\u0000\u0000\u0000\u00bd\u00be\u0001\u0000\u0000"+
		"\u0000\u00be\u00bf\u0001\u0000\u0000\u0000\u00bf\u00ce\u00057\u0000\u0000"+
		"\u00c0\u00c1\u0005\u0018\u0000\u0000\u00c1\u00ca\u00056\u0000\u0000\u00c2"+
		"\u00c7\u0003\b\u0004\u0000\u00c3\u00c4\u00053\u0000\u0000\u00c4\u00c6"+
		"\u0003\b\u0004\u0000\u00c5\u00c3\u0001\u0000\u0000\u0000\u00c6\u00c9\u0001"+
		"\u0000\u0000\u0000\u00c7\u00c5\u0001\u0000\u0000\u0000\u00c7\u00c8\u0001"+
		"\u0000\u0000\u0000\u00c8\u00cb\u0001\u0000\u0000\u0000\u00c9\u00c7\u0001"+
		"\u0000\u0000\u0000\u00ca\u00c2\u0001\u0000\u0000\u0000\u00ca\u00cb\u0001"+
		"\u0000\u0000\u0000\u00cb\u00cc\u0001\u0000\u0000\u0000\u00cc\u00ce\u0005"+
		"7\u0000\u0000\u00cd\u00a6\u0001\u0000\u0000\u0000\u00cd\u00b3\u0001\u0000"+
		"\u0000\u0000\u00cd\u00c0\u0001\u0000\u0000\u0000\u00ce\u000f\u0001\u0000"+
		"\u0000\u0000\u00cf\u00d0\u0005\u0005\u0000\u0000\u00d0\u00d1\u0003\b\u0004"+
		"\u0000\u00d1\u00d5\u00058\u0000\u0000\u00d2\u00d4\u0003\u0002\u0001\u0000"+
		"\u00d3\u00d2\u0001\u0000\u0000\u0000\u00d4\u00d7\u0001\u0000\u0000\u0000"+
		"\u00d5\u00d3\u0001\u0000\u0000\u0000\u00d5\u00d6\u0001\u0000\u0000\u0000"+
		"\u00d6\u00d8\u0001\u0000\u0000\u0000\u00d7\u00d5\u0001\u0000\u0000\u0000"+
		"\u00d8\u00e2\u00059\u0000\u0000\u00d9\u00da\u0005\u0006\u0000\u0000\u00da"+
		"\u00de\u00058\u0000\u0000\u00db\u00dd\u0003\u0002\u0001\u0000\u00dc\u00db"+
		"\u0001\u0000\u0000\u0000\u00dd\u00e0\u0001\u0000\u0000\u0000\u00de\u00dc"+
		"\u0001\u0000\u0000\u0000\u00de\u00df\u0001\u0000\u0000\u0000\u00df\u00e1"+
		"\u0001\u0000\u0000\u0000\u00e0\u00de\u0001\u0000\u0000\u0000\u00e1\u00e3"+
		"\u00059\u0000\u0000\u00e2\u00d9\u0001\u0000\u0000\u0000\u00e2\u00e3\u0001"+
		"\u0000\u0000\u0000\u00e3\u0011\u0001\u0000\u0000\u0000\u00e4\u00e5\u0005"+
		"\u0007\u0000\u0000\u00e5\u00e6\u0003\b\u0004\u0000\u00e6\u00ea\u00058"+
		"\u0000\u0000\u00e7\u00e9\u0003\u0014\n\u0000\u00e8\u00e7\u0001\u0000\u0000"+
		"\u0000\u00e9\u00ec\u0001\u0000\u0000\u0000\u00ea\u00e8\u0001\u0000\u0000"+
		"\u0000\u00ea\u00eb\u0001\u0000\u0000\u0000\u00eb\u00ee\u0001\u0000\u0000"+
		"\u0000\u00ec\u00ea\u0001\u0000\u0000\u0000\u00ed\u00ef\u0003\u0016\u000b"+
		"\u0000\u00ee\u00ed\u0001\u0000\u0000\u0000\u00ee\u00ef\u0001\u0000\u0000"+
		"\u0000\u00ef\u00f0\u0001\u0000\u0000\u0000\u00f0\u00f1\u00059\u0000\u0000"+
		"\u00f1\u0013\u0001\u0000\u0000\u0000\u00f2\u00f3\u0005\b\u0000\u0000\u00f3"+
		"\u00f4\u0003\b\u0004\u0000\u00f4\u00f8\u00055\u0000\u0000\u00f5\u00f7"+
		"\u0003\u0002\u0001\u0000\u00f6\u00f5\u0001\u0000\u0000\u0000\u00f7\u00fa"+
		"\u0001\u0000\u0000\u0000\u00f8\u00f6\u0001\u0000\u0000\u0000\u00f8\u00f9"+
		"\u0001\u0000\u0000\u0000\u00f9\u0015\u0001\u0000\u0000\u0000\u00fa\u00f8"+
		"\u0001\u0000\u0000\u0000\u00fb\u00fc\u0005\t\u0000\u0000\u00fc\u0100\u0005"+
		"5\u0000\u0000\u00fd\u00ff\u0003\u0002\u0001\u0000\u00fe\u00fd\u0001\u0000"+
		"\u0000\u0000\u00ff\u0102\u0001\u0000\u0000\u0000\u0100\u00fe\u0001\u0000"+
		"\u0000\u0000\u0100\u0101\u0001\u0000\u0000\u0000\u0101\u0017\u0001\u0000"+
		"\u0000\u0000\u0102\u0100\u0001\u0000\u0000\u0000\u0103\u0104\u0005\n\u0000"+
		"\u0000\u0104\u0105\u0005@\u0000\u0000\u0105\u0106\u0005\u000b\u0000\u0000"+
		"\u0106\u0107\u0003\b\u0004\u0000\u0107\u010b\u00058\u0000\u0000\u0108"+
		"\u010a\u0003\u0002\u0001\u0000\u0109\u0108\u0001\u0000\u0000\u0000\u010a"+
		"\u010d\u0001\u0000\u0000\u0000\u010b\u0109\u0001\u0000\u0000\u0000\u010b"+
		"\u010c\u0001\u0000\u0000\u0000\u010c\u010e\u0001\u0000\u0000\u0000\u010d"+
		"\u010b\u0001\u0000\u0000\u0000\u010e\u010f\u00059\u0000\u0000\u010f\u0019"+
		"\u0001\u0000\u0000\u0000\u0110\u0111\u0005\u0002\u0000\u0000\u0111\u0112"+
		"\u0007\u0000\u0000\u0000\u0112\u011b\u00056\u0000\u0000\u0113\u0118\u0003"+
		"\u001c\u000e\u0000\u0114\u0115\u00053\u0000\u0000\u0115\u0117\u0003\u001c"+
		"\u000e\u0000\u0116\u0114\u0001\u0000\u0000\u0000\u0117\u011a\u0001\u0000"+
		"\u0000\u0000\u0118\u0116\u0001\u0000\u0000\u0000\u0118\u0119\u0001\u0000"+
		"\u0000\u0000\u0119\u011c\u0001\u0000\u0000\u0000\u011a\u0118\u0001\u0000"+
		"\u0000\u0000\u011b\u0113\u0001\u0000\u0000\u0000\u011b\u011c\u0001\u0000"+
		"\u0000\u0000\u011c\u011d\u0001\u0000\u0000\u0000\u011d\u011f\u00057\u0000"+
		"\u0000\u011e\u0120\u0007\u0001\u0000\u0000\u011f\u011e\u0001\u0000\u0000"+
		"\u0000\u011f\u0120\u0001\u0000\u0000\u0000\u0120\u0121\u0001\u0000\u0000"+
		"\u0000\u0121\u0125\u00058\u0000\u0000\u0122\u0124\u0003\u0002\u0001\u0000"+
		"\u0123\u0122\u0001\u0000\u0000\u0000\u0124\u0127\u0001\u0000\u0000\u0000"+
		"\u0125\u0123\u0001\u0000\u0000\u0000\u0125\u0126\u0001\u0000\u0000\u0000"+
		"\u0126\u0128\u0001\u0000\u0000\u0000\u0127\u0125\u0001\u0000\u0000\u0000"+
		"\u0128\u0129\u00059\u0000\u0000\u0129\u001b\u0001\u0000\u0000\u0000\u012a"+
		"\u012b\u0005@\u0000\u0000\u012b\u0135\u0005\u0012\u0000\u0000\u012c\u012d"+
		"\u0005@\u0000\u0000\u012d\u0135\u0005\u0014\u0000\u0000\u012e\u012f\u0005"+
		"@\u0000\u0000\u012f\u0135\u0005\u0013\u0000\u0000\u0130\u0131\u0005@\u0000"+
		"\u0000\u0131\u0135\u0005\u0015\u0000\u0000\u0132\u0133\u0005@\u0000\u0000"+
		"\u0133\u0135\u0005\u0016\u0000\u0000\u0134\u012a\u0001\u0000\u0000\u0000"+
		"\u0134\u012c\u0001\u0000\u0000\u0000\u0134\u012e\u0001\u0000\u0000\u0000"+
		"\u0134\u0130\u0001\u0000\u0000\u0000\u0134\u0132\u0001\u0000\u0000\u0000"+
		"\u0135\u001d\u0001\u0000\u0000\u0000\u0136\u0137\u0005\u0004\u0000\u0000"+
		"\u0137\u0138\u0005@\u0000\u0000\u0138\u013c\u00058\u0000\u0000\u0139\u013b"+
		"\u0003 \u0010\u0000\u013a\u0139\u0001\u0000\u0000\u0000\u013b\u013e\u0001"+
		"\u0000\u0000\u0000\u013c\u013a\u0001\u0000\u0000\u0000\u013c\u013d\u0001"+
		"\u0000\u0000\u0000\u013d\u013f\u0001\u0000\u0000\u0000\u013e\u013c\u0001"+
		"\u0000\u0000\u0000\u013f\u0140\u00059\u0000\u0000\u0140\u001f\u0001\u0000"+
		"\u0000\u0000\u0141\u0142\u0005\u0003\u0000\u0000\u0142\u0143\u0005@\u0000"+
		"\u0000\u0143\u0151\u0005\u0012\u0000\u0000\u0144\u0145\u0005\u0003\u0000"+
		"\u0000\u0145\u0146\u0005@\u0000\u0000\u0146\u0151\u0005\u0014\u0000\u0000"+
		"\u0147\u0148\u0005\u0003\u0000\u0000\u0148\u0149\u0005@\u0000\u0000\u0149"+
		"\u0151\u0005\u0013\u0000\u0000\u014a\u014b\u0005\u0003\u0000\u0000\u014b"+
		"\u014c\u0005@\u0000\u0000\u014c\u0151\u0005\u0015\u0000\u0000\u014d\u014e"+
		"\u0005\u0003\u0000\u0000\u014e\u014f\u0005@\u0000\u0000\u014f\u0151\u0005"+
		"\u0016\u0000\u0000\u0150\u0141\u0001\u0000\u0000\u0000\u0150\u0144\u0001"+
		"\u0000\u0000\u0000\u0150\u0147\u0001\u0000\u0000\u0000\u0150\u014a\u0001"+
		"\u0000\u0000\u0000\u0150\u014d\u0001\u0000\u0000\u0000\u0151!\u0001\u0000"+
		"\u0000\u0000#%,05<Q\\g\u0090\u0092\u009c\u00a2\u00a4\u00ad\u00b0\u00ba"+
		"\u00bd\u00c7\u00ca\u00cd\u00d5\u00de\u00e2\u00ea\u00ee\u00f8\u0100\u010b"+
		"\u0118\u011b\u011f\u0125\u0134\u013c\u0150";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}