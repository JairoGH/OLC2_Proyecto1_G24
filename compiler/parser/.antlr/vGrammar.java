// Generated from /home/jairogo/Documentos/GitHub/OLC2_Proyecto1/compiler/parser/vGrammar.g4 by ANTLR 4.13.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast", "CheckReturnValue"})
public class vGrammar extends Parser {
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
		RULE_programa = 0, RULE_funcionPrincipal = 1, RULE_bloque = 2, RULE_sentencia = 3, 
		RULE_declaracion = 4, RULE_asignacion = 5, RULE_ifStatement = 6, RULE_expresion = 7;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "funcionPrincipal", "bloque", "sentencia", "declaracion", 
			"asignacion", "ifStatement", "expresion"
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
	public String getGrammarFileName() { return "vGrammar.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public vGrammar(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	@SuppressWarnings("CheckReturnValue")
	public static class ProgramaContext extends ParserRuleContext {
		public FuncionPrincipalContext funcionPrincipal() {
			return getRuleContext(FuncionPrincipalContext.class,0);
		}
		public TerminalNode EOF() { return getToken(vGrammar.EOF, 0); }
		public ProgramaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_programa; }
	}

	public final ProgramaContext programa() throws RecognitionException {
		ProgramaContext _localctx = new ProgramaContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_programa);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(16);
			funcionPrincipal();
			setState(17);
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
	public static class FuncionPrincipalContext extends ParserRuleContext {
		public TerminalNode RW_FN() { return getToken(vGrammar.RW_FN, 0); }
		public TerminalNode RW_MAIN() { return getToken(vGrammar.RW_MAIN, 0); }
		public TerminalNode PAR_IZQ() { return getToken(vGrammar.PAR_IZQ, 0); }
		public TerminalNode PAR_DER() { return getToken(vGrammar.PAR_DER, 0); }
		public BloqueContext bloque() {
			return getRuleContext(BloqueContext.class,0);
		}
		public FuncionPrincipalContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_funcionPrincipal; }
	}

	public final FuncionPrincipalContext funcionPrincipal() throws RecognitionException {
		FuncionPrincipalContext _localctx = new FuncionPrincipalContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_funcionPrincipal);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(19);
			match(RW_FN);
			setState(20);
			match(RW_MAIN);
			setState(21);
			match(PAR_IZQ);
			setState(22);
			match(PAR_DER);
			setState(23);
			bloque();
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
	public static class BloqueContext extends ParserRuleContext {
		public TerminalNode LLAVE_IZQ() { return getToken(vGrammar.LLAVE_IZQ, 0); }
		public TerminalNode LLAVE_DER() { return getToken(vGrammar.LLAVE_DER, 0); }
		public List<SentenciaContext> sentencia() {
			return getRuleContexts(SentenciaContext.class);
		}
		public SentenciaContext sentencia(int i) {
			return getRuleContext(SentenciaContext.class,i);
		}
		public BloqueContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_bloque; }
	}

	public final BloqueContext bloque() throws RecognitionException {
		BloqueContext _localctx = new BloqueContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_bloque);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(25);
			match(LLAVE_IZQ);
			setState(29);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 3)) & ~0x3f) == 0 && ((1L << (_la - 3)) & 2305843009214742533L) != 0)) {
				{
				{
				setState(26);
				sentencia();
				}
				}
				setState(31);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(32);
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
	public static class SentenciaContext extends ParserRuleContext {
		public DeclaracionContext declaracion() {
			return getRuleContext(DeclaracionContext.class,0);
		}
		public AsignacionContext asignacion() {
			return getRuleContext(AsignacionContext.class,0);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public TerminalNode RW_PRINT() { return getToken(vGrammar.RW_PRINT, 0); }
		public TerminalNode PAR_IZQ() { return getToken(vGrammar.PAR_IZQ, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public TerminalNode PAR_DER() { return getToken(vGrammar.PAR_DER, 0); }
		public SentenciaContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_sentencia; }
	}

	public final SentenciaContext sentencia() throws RecognitionException {
		SentenciaContext _localctx = new SentenciaContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_sentencia);
		try {
			setState(42);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case RW_MUT:
				enterOuterAlt(_localctx, 1);
				{
				setState(34);
				declaracion();
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				setState(35);
				asignacion();
				}
				break;
			case RW_IF:
				enterOuterAlt(_localctx, 3);
				{
				setState(36);
				ifStatement();
				}
				break;
			case RW_PRINT:
				enterOuterAlt(_localctx, 4);
				{
				setState(37);
				match(RW_PRINT);
				setState(38);
				match(PAR_IZQ);
				setState(39);
				expresion(0);
				setState(40);
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
	public static class DeclaracionContext extends ParserRuleContext {
		public TerminalNode RW_MUT() { return getToken(vGrammar.RW_MUT, 0); }
		public TerminalNode ID() { return getToken(vGrammar.ID, 0); }
		public TerminalNode OP_DECLARATION() { return getToken(vGrammar.OP_DECLARATION, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public DeclaracionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion; }
	}

	public final DeclaracionContext declaracion() throws RecognitionException {
		DeclaracionContext _localctx = new DeclaracionContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_declaracion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(44);
			match(RW_MUT);
			setState(45);
			match(ID);
			setState(46);
			match(OP_DECLARATION);
			setState(47);
			expresion(0);
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
	public static class AsignacionContext extends ParserRuleContext {
		public TerminalNode ID() { return getToken(vGrammar.ID, 0); }
		public TerminalNode OP_ASSIGN() { return getToken(vGrammar.OP_ASSIGN, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public AsignacionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion; }
	}

	public final AsignacionContext asignacion() throws RecognitionException {
		AsignacionContext _localctx = new AsignacionContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_asignacion);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(49);
			match(ID);
			setState(50);
			match(OP_ASSIGN);
			setState(51);
			expresion(0);
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
	public static class IfStatementContext extends ParserRuleContext {
		public TerminalNode RW_IF() { return getToken(vGrammar.RW_IF, 0); }
		public ExpresionContext expresion() {
			return getRuleContext(ExpresionContext.class,0);
		}
		public List<BloqueContext> bloque() {
			return getRuleContexts(BloqueContext.class);
		}
		public BloqueContext bloque(int i) {
			return getRuleContext(BloqueContext.class,i);
		}
		public List<TerminalNode> RW_ELSE() { return getTokens(vGrammar.RW_ELSE); }
		public TerminalNode RW_ELSE(int i) {
			return getToken(vGrammar.RW_ELSE, i);
		}
		public IfStatementContext ifStatement() {
			return getRuleContext(IfStatementContext.class,0);
		}
		public IfStatementContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifStatement; }
	}

	public final IfStatementContext ifStatement() throws RecognitionException {
		IfStatementContext _localctx = new IfStatementContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_ifStatement);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(53);
			match(RW_IF);
			setState(54);
			expresion(0);
			setState(55);
			bloque();
			setState(58);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,2,_ctx) ) {
			case 1:
				{
				setState(56);
				match(RW_ELSE);
				setState(57);
				ifStatement();
				}
				break;
			}
			setState(62);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,3,_ctx) ) {
			case 1:
				{
				setState(60);
				match(RW_ELSE);
				setState(61);
				bloque();
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
	public static class ExpresionContext extends ParserRuleContext {
		public TerminalNode PAR_IZQ() { return getToken(vGrammar.PAR_IZQ, 0); }
		public List<ExpresionContext> expresion() {
			return getRuleContexts(ExpresionContext.class);
		}
		public ExpresionContext expresion(int i) {
			return getRuleContext(ExpresionContext.class,i);
		}
		public TerminalNode PAR_DER() { return getToken(vGrammar.PAR_DER, 0); }
		public TerminalNode INT_LITERAL() { return getToken(vGrammar.INT_LITERAL, 0); }
		public TerminalNode FLOAT_LITERAL() { return getToken(vGrammar.FLOAT_LITERAL, 0); }
		public TerminalNode STRING_LITERAL() { return getToken(vGrammar.STRING_LITERAL, 0); }
		public TerminalNode RUNE_LITERAL() { return getToken(vGrammar.RUNE_LITERAL, 0); }
		public TerminalNode RW_TRUE() { return getToken(vGrammar.RW_TRUE, 0); }
		public TerminalNode RW_FALSE() { return getToken(vGrammar.RW_FALSE, 0); }
		public TerminalNode ID() { return getToken(vGrammar.ID, 0); }
		public TerminalNode OP_SUMA() { return getToken(vGrammar.OP_SUMA, 0); }
		public TerminalNode OP_RESTA() { return getToken(vGrammar.OP_RESTA, 0); }
		public TerminalNode OP_MULT() { return getToken(vGrammar.OP_MULT, 0); }
		public TerminalNode OP_DIV() { return getToken(vGrammar.OP_DIV, 0); }
		public TerminalNode OP_MOD() { return getToken(vGrammar.OP_MOD, 0); }
		public ExpresionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expresion; }
	}

	public final ExpresionContext expresion() throws RecognitionException {
		return expresion(0);
	}

	private ExpresionContext expresion(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpresionContext _localctx = new ExpresionContext(_ctx, _parentState);
		ExpresionContext _prevctx = _localctx;
		int _startState = 14;
		enterRecursionRule(_localctx, 14, RULE_expresion, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case PAR_IZQ:
				{
				setState(65);
				match(PAR_IZQ);
				setState(66);
				expresion(0);
				setState(67);
				match(PAR_DER);
				}
				break;
			case INT_LITERAL:
				{
				setState(69);
				match(INT_LITERAL);
				}
				break;
			case FLOAT_LITERAL:
				{
				setState(70);
				match(FLOAT_LITERAL);
				}
				break;
			case STRING_LITERAL:
				{
				setState(71);
				match(STRING_LITERAL);
				}
				break;
			case RUNE_LITERAL:
				{
				setState(72);
				match(RUNE_LITERAL);
				}
				break;
			case RW_TRUE:
				{
				setState(73);
				match(RW_TRUE);
				}
				break;
			case RW_FALSE:
				{
				setState(74);
				match(RW_FALSE);
				}
				break;
			case ID:
				{
				setState(75);
				match(ID);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(95);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(93);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
					case 1:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(78);
						if (!(precpred(_ctx, 13))) throw new FailedPredicateException(this, "precpred(_ctx, 13)");
						setState(79);
						match(OP_SUMA);
						setState(80);
						expresion(14);
						}
						break;
					case 2:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(81);
						if (!(precpred(_ctx, 12))) throw new FailedPredicateException(this, "precpred(_ctx, 12)");
						setState(82);
						match(OP_RESTA);
						setState(83);
						expresion(13);
						}
						break;
					case 3:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(84);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(85);
						match(OP_MULT);
						setState(86);
						expresion(12);
						}
						break;
					case 4:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(87);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(88);
						match(OP_DIV);
						setState(89);
						expresion(11);
						}
						break;
					case 5:
						{
						_localctx = new ExpresionContext(_parentctx, _parentState);
						pushNewRecursionContext(_localctx, _startState, RULE_expresion);
						setState(90);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(91);
						match(OP_MOD);
						setState(92);
						expresion(10);
						}
						break;
					}
					} 
				}
				setState(97);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,6,_ctx);
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

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 7:
			return expresion_sempred((ExpresionContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expresion_sempred(ExpresionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 13);
		case 1:
			return precpred(_ctx, 12);
		case 2:
			return precpred(_ctx, 11);
		case 3:
			return precpred(_ctx, 10);
		case 4:
			return precpred(_ctx, 9);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001Cc\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0002\u0002"+
		"\u0007\u0002\u0002\u0003\u0007\u0003\u0002\u0004\u0007\u0004\u0002\u0005"+
		"\u0007\u0005\u0002\u0006\u0007\u0006\u0002\u0007\u0007\u0007\u0001\u0000"+
		"\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0002\u0001\u0002\u0005\u0002\u001c\b\u0002"+
		"\n\u0002\f\u0002\u001f\t\u0002\u0001\u0002\u0001\u0002\u0001\u0003\u0001"+
		"\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001\u0003\u0001"+
		"\u0003\u0003\u0003+\b\u0003\u0001\u0004\u0001\u0004\u0001\u0004\u0001"+
		"\u0004\u0001\u0004\u0001\u0005\u0001\u0005\u0001\u0005\u0001\u0005\u0001"+
		"\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0001\u0006\u0003\u0006;\b"+
		"\u0006\u0001\u0006\u0001\u0006\u0003\u0006?\b\u0006\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0003\u0007M\b"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0001"+
		"\u0007\u0001\u0007\u0001\u0007\u0001\u0007\u0005\u0007^\b\u0007\n\u0007"+
		"\f\u0007a\t\u0007\u0001\u0007\u0000\u0001\u000e\b\u0000\u0002\u0004\u0006"+
		"\b\n\f\u000e\u0000\u0000l\u0000\u0010\u0001\u0000\u0000\u0000\u0002\u0013"+
		"\u0001\u0000\u0000\u0000\u0004\u0019\u0001\u0000\u0000\u0000\u0006*\u0001"+
		"\u0000\u0000\u0000\b,\u0001\u0000\u0000\u0000\n1\u0001\u0000\u0000\u0000"+
		"\f5\u0001\u0000\u0000\u0000\u000eL\u0001\u0000\u0000\u0000\u0010\u0011"+
		"\u0003\u0002\u0001\u0000\u0011\u0012\u0005\u0000\u0000\u0001\u0012\u0001"+
		"\u0001\u0000\u0000\u0000\u0013\u0014\u0005\u0002\u0000\u0000\u0014\u0015"+
		"\u0005\u0001\u0000\u0000\u0015\u0016\u00056\u0000\u0000\u0016\u0017\u0005"+
		"7\u0000\u0000\u0017\u0018\u0003\u0004\u0002\u0000\u0018\u0003\u0001\u0000"+
		"\u0000\u0000\u0019\u001d\u00058\u0000\u0000\u001a\u001c\u0003\u0006\u0003"+
		"\u0000\u001b\u001a\u0001\u0000\u0000\u0000\u001c\u001f\u0001\u0000\u0000"+
		"\u0000\u001d\u001b\u0001\u0000\u0000\u0000\u001d\u001e\u0001\u0000\u0000"+
		"\u0000\u001e \u0001\u0000\u0000\u0000\u001f\u001d\u0001\u0000\u0000\u0000"+
		" !\u00059\u0000\u0000!\u0005\u0001\u0000\u0000\u0000\"+\u0003\b\u0004"+
		"\u0000#+\u0003\n\u0005\u0000$+\u0003\f\u0006\u0000%&\u0005\u0017\u0000"+
		"\u0000&\'\u00056\u0000\u0000\'(\u0003\u000e\u0007\u0000()\u00057\u0000"+
		"\u0000)+\u0001\u0000\u0000\u0000*\"\u0001\u0000\u0000\u0000*#\u0001\u0000"+
		"\u0000\u0000*$\u0001\u0000\u0000\u0000*%\u0001\u0000\u0000\u0000+\u0007"+
		"\u0001\u0000\u0000\u0000,-\u0005\u0003\u0000\u0000-.\u0005@\u0000\u0000"+
		"./\u00051\u0000\u0000/0\u0003\u000e\u0007\u00000\t\u0001\u0000\u0000\u0000"+
		"12\u0005@\u0000\u000023\u0005%\u0000\u000034\u0003\u000e\u0007\u00004"+
		"\u000b\u0001\u0000\u0000\u000056\u0005\u0005\u0000\u000067\u0003\u000e"+
		"\u0007\u00007:\u0003\u0004\u0002\u000089\u0005\u0006\u0000\u00009;\u0003"+
		"\f\u0006\u0000:8\u0001\u0000\u0000\u0000:;\u0001\u0000\u0000\u0000;>\u0001"+
		"\u0000\u0000\u0000<=\u0005\u0006\u0000\u0000=?\u0003\u0004\u0002\u0000"+
		"><\u0001\u0000\u0000\u0000>?\u0001\u0000\u0000\u0000?\r\u0001\u0000\u0000"+
		"\u0000@A\u0006\u0007\uffff\uffff\u0000AB\u00056\u0000\u0000BC\u0003\u000e"+
		"\u0007\u0000CD\u00057\u0000\u0000DM\u0001\u0000\u0000\u0000EM\u0005<\u0000"+
		"\u0000FM\u0005=\u0000\u0000GM\u0005>\u0000\u0000HM\u0005?\u0000\u0000"+
		"IM\u0005\u000f\u0000\u0000JM\u0005\u0010\u0000\u0000KM\u0005@\u0000\u0000"+
		"L@\u0001\u0000\u0000\u0000LE\u0001\u0000\u0000\u0000LF\u0001\u0000\u0000"+
		"\u0000LG\u0001\u0000\u0000\u0000LH\u0001\u0000\u0000\u0000LI\u0001\u0000"+
		"\u0000\u0000LJ\u0001\u0000\u0000\u0000LK\u0001\u0000\u0000\u0000M_\u0001"+
		"\u0000\u0000\u0000NO\n\r\u0000\u0000OP\u0005 \u0000\u0000P^\u0003\u000e"+
		"\u0007\u000eQR\n\f\u0000\u0000RS\u0005!\u0000\u0000S^\u0003\u000e\u0007"+
		"\rTU\n\u000b\u0000\u0000UV\u0005\"\u0000\u0000V^\u0003\u000e\u0007\fW"+
		"X\n\n\u0000\u0000XY\u0005#\u0000\u0000Y^\u0003\u000e\u0007\u000bZ[\n\t"+
		"\u0000\u0000[\\\u0005$\u0000\u0000\\^\u0003\u000e\u0007\n]N\u0001\u0000"+
		"\u0000\u0000]Q\u0001\u0000\u0000\u0000]T\u0001\u0000\u0000\u0000]W\u0001"+
		"\u0000\u0000\u0000]Z\u0001\u0000\u0000\u0000^a\u0001\u0000\u0000\u0000"+
		"_]\u0001\u0000\u0000\u0000_`\u0001\u0000\u0000\u0000`\u000f\u0001\u0000"+
		"\u0000\u0000a_\u0001\u0000\u0000\u0000\u0007\u001d*:>L]_";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}