// Generated from /home/daaniieel/Escritorio/PROYECTO_1/Proyecto1/parser/VGrammar.g4 by ANTLR 4.13.1
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
		INT=1, RW_MAIN=2, RW_FN=3, RW_MUT=4, RW_STRUCT=5, RW_IF=6, RW_ELSE=7, 
		RW_SWITCH=8, RW_CASE=9, RW_DEFAULT=10, RW_FOR=11, RW_IN=12, RW_BREAK=13, 
		RW_CONTINUE=14, RW_RETURN=15, RW_TRUE=16, RW_FALSE=17, RW_NIL=18, RW_INT=19, 
		RW_FLOAT64=20, RW_STRING=21, RW_BOOL=22, RW_RUNE=23, RW_PRINT=24, RW_PRINTLN=25, 
		RW_ATOI=26, RW_PARSEFLOAT=27, RW_TYPEOF=28, RW_APPEND=29, RW_LEN=30, RW_JOIN=31, 
		RW_INDEXOF=32, OP_SUMA=33, OP_RESTA=34, OP_MULT=35, OP_DIV=36, OP_MOD=37, 
		OP_ASSIGN=38, OP_ADD_ASSIGN=39, OP_SUB_ASSIGN=40, OP_IGUAL=41, OP_DIFERENTE=42, 
		OP_MENORQ=43, OP_MAYORQ=44, OP_MENORIGUAL=45, OP_MAYORIGUAL=46, OP_AND=47, 
		OP_OR=48, OP_NOT=49, OP_DECLARATION=50, PUNTO=51, COMA=52, PUNTO_Y_COMA=53, 
		DOS_PUNTOS=54, PAR_IZQ=55, PAR_DER=56, LLAVE_IZQ=57, LLAVE_DER=58, CORCHETE_IZQ=59, 
		CORCHETE_DER=60, INT_LITERAL=61, FLOAT_LITERAL=62, STRING_LITERAL=63, 
		RUNE_LITERAL=64, ID=65, WS=66, COMMENT=67, MULTILINE_COMMENT=68;
	public static final int
		RULE_programa = 0, RULE_expr = 1;
	private static String[] makeRuleNames() {
		return new String[] {
			"programa", "expr"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, null, "'main'", "'fn'", "'mut'", "'struct'", "'if'", "'else'", 
			"'switch'", "'case'", "'default'", "'for'", "'in'", "'break'", "'continue'", 
			"'return'", "'true'", "'false'", "'nil'", "'int'", "'float64'", "'string'", 
			"'bool'", "'rune'", "'print'", "'println'", "'Atoi'", "'parseFloat'", 
			"'typeof'", "'append'", "'len'", "'join'", "'indexOf'", "'+'", "'-'", 
			"'*'", "'/'", "'%'", "'='", "'+='", "'-='", "'=='", "'!='", "'<'", "'>'", 
			"'<='", "'>='", "'&&'", "'||'", "'!'", "':='", "'.'", "','", "';'", "':'", 
			"'('", "')'", "'{'", "'}'", "'['", "']'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "INT", "RW_MAIN", "RW_FN", "RW_MUT", "RW_STRUCT", "RW_IF", "RW_ELSE", 
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
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode EOF() { return getToken(VGrammar.EOF, 0); }
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
			setState(4);
			expr(0);
			setState(5);
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
	public static class MulDivContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_MULT() { return getToken(VGrammar.OP_MULT, 0); }
		public TerminalNode OP_DIV() { return getToken(VGrammar.OP_DIV, 0); }
		public MulDivContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class AddSubContext extends ExprContext {
		public Token op;
		public List<ExprContext> expr() {
			return getRuleContexts(ExprContext.class);
		}
		public ExprContext expr(int i) {
			return getRuleContext(ExprContext.class,i);
		}
		public TerminalNode OP_SUMA() { return getToken(VGrammar.OP_SUMA, 0); }
		public TerminalNode OP_RESTA() { return getToken(VGrammar.OP_RESTA, 0); }
		public AddSubContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class ParensContext extends ExprContext {
		public TerminalNode PAR_IZQ() { return getToken(VGrammar.PAR_IZQ, 0); }
		public ExprContext expr() {
			return getRuleContext(ExprContext.class,0);
		}
		public TerminalNode PAR_DER() { return getToken(VGrammar.PAR_DER, 0); }
		public ParensContext(ExprContext ctx) { copyFrom(ctx); }
	}
	@SuppressWarnings("CheckReturnValue")
	public static class IntContext extends ExprContext {
		public TerminalNode INT_LITERAL() { return getToken(VGrammar.INT_LITERAL, 0); }
		public IntContext(ExprContext ctx) { copyFrom(ctx); }
	}

	public final ExprContext expr() throws RecognitionException {
		return expr(0);
	}

	private ExprContext expr(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExprContext _localctx = new ExprContext(_ctx, _parentState);
		ExprContext _prevctx = _localctx;
		int _startState = 2;
		enterRecursionRule(_localctx, 2, RULE_expr, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(13);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT_LITERAL:
				{
				_localctx = new IntContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;

				setState(8);
				match(INT_LITERAL);
				}
				break;
			case PAR_IZQ:
				{
				_localctx = new ParensContext(_localctx);
				_ctx = _localctx;
				_prevctx = _localctx;
				setState(9);
				match(PAR_IZQ);
				setState(10);
				expr(0);
				setState(11);
				match(PAR_DER);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(23);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(21);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,1,_ctx) ) {
					case 1:
						{
						_localctx = new MulDivContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(15);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(16);
						((MulDivContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_MULT || _la==OP_DIV) ) {
							((MulDivContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(17);
						expr(5);
						}
						break;
					case 2:
						{
						_localctx = new AddSubContext(new ExprContext(_parentctx, _parentState));
						pushNewRecursionContext(_localctx, _startState, RULE_expr);
						setState(18);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(19);
						((AddSubContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OP_SUMA || _la==OP_RESTA) ) {
							((AddSubContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(20);
						expr(4);
						}
						break;
					}
					} 
				}
				setState(25);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,2,_ctx);
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
		case 1:
			return expr_sempred((ExprContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expr_sempred(ExprContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 4);
		case 1:
			return precpred(_ctx, 3);
		}
		return true;
	}

	public static final String _serializedATN =
		"\u0004\u0001D\u001b\u0002\u0000\u0007\u0000\u0002\u0001\u0007\u0001\u0001"+
		"\u0000\u0001\u0000\u0001\u0000\u0001\u0001\u0001\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0003\u0001\u000e\b\u0001\u0001\u0001\u0001"+
		"\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0001\u0005\u0001\u0016"+
		"\b\u0001\n\u0001\f\u0001\u0019\t\u0001\u0001\u0001\u0000\u0001\u0002\u0002"+
		"\u0000\u0002\u0000\u0002\u0001\u0000#$\u0001\u0000!\"\u001b\u0000\u0004"+
		"\u0001\u0000\u0000\u0000\u0002\r\u0001\u0000\u0000\u0000\u0004\u0005\u0003"+
		"\u0002\u0001\u0000\u0005\u0006\u0005\u0000\u0000\u0001\u0006\u0001\u0001"+
		"\u0000\u0000\u0000\u0007\b\u0006\u0001\uffff\uffff\u0000\b\u000e\u0005"+
		"=\u0000\u0000\t\n\u00057\u0000\u0000\n\u000b\u0003\u0002\u0001\u0000\u000b"+
		"\f\u00058\u0000\u0000\f\u000e\u0001\u0000\u0000\u0000\r\u0007\u0001\u0000"+
		"\u0000\u0000\r\t\u0001\u0000\u0000\u0000\u000e\u0017\u0001\u0000\u0000"+
		"\u0000\u000f\u0010\n\u0004\u0000\u0000\u0010\u0011\u0007\u0000\u0000\u0000"+
		"\u0011\u0016\u0003\u0002\u0001\u0005\u0012\u0013\n\u0003\u0000\u0000\u0013"+
		"\u0014\u0007\u0001\u0000\u0000\u0014\u0016\u0003\u0002\u0001\u0004\u0015"+
		"\u000f\u0001\u0000\u0000\u0000\u0015\u0012\u0001\u0000\u0000\u0000\u0016"+
		"\u0019\u0001\u0000\u0000\u0000\u0017\u0015\u0001\u0000\u0000\u0000\u0017"+
		"\u0018\u0001\u0000\u0000\u0000\u0018\u0003\u0001\u0000\u0000\u0000\u0019"+
		"\u0017\u0001\u0000\u0000\u0000\u0003\r\u0015\u0017";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}