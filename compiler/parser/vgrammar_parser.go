// Code generated from parser/VGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // VGrammar
import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type VGrammar struct {
	*antlr.BaseParser
}

var VGrammarParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func vgrammarParserInit() {
	staticData := &VGrammarParserStaticData
	staticData.LiteralNames = []string{
		"", "'main'", "'fn'", "'mut'", "'struct'", "'if'", "'else'", "'switch'",
		"'case'", "'default'", "'for'", "'in'", "'break'", "'continue'", "'return'",
		"'true'", "'false'", "'nil'", "'int'", "'float64'", "'string'", "'bool'",
		"'rune'", "'print'", "'println'", "'Atoi'", "'parseFloat'", "'typeof'",
		"'append'", "'len'", "'join'", "'indexOf'", "'+'", "'-'", "'*'", "'/'",
		"'%'", "'='", "'+='", "'-='", "'=='", "'!='", "'<'", "'>'", "'<='",
		"'>='", "'&&'", "'||'", "'!'", "':='", "'.'", "','", "';'", "':'", "'('",
		"')'", "'{'", "'}'", "'['", "']'",
	}
	staticData.SymbolicNames = []string{
		"", "RW_MAIN", "RW_FN", "RW_MUT", "RW_STRUCT", "RW_IF", "RW_ELSE", "RW_SWITCH",
		"RW_CASE", "RW_DEFAULT", "RW_FOR", "RW_IN", "RW_BREAK", "RW_CONTINUE",
		"RW_RETURN", "RW_TRUE", "RW_FALSE", "RW_NIL", "RW_INT", "RW_FLOAT64",
		"RW_STRING", "RW_BOOL", "RW_RUNE", "RW_PRINT", "RW_PRINTLN", "RW_ATOI",
		"RW_PARSEFLOAT", "RW_TYPEOF", "RW_APPEND", "RW_LEN", "RW_JOIN", "RW_INDEXOF",
		"OP_SUMA", "OP_RESTA", "OP_MULT", "OP_DIV", "OP_MOD", "OP_ASSIGN", "OP_ADD_ASSIGN",
		"OP_SUB_ASSIGN", "OP_IGUAL", "OP_DIFERENTE", "OP_MENORQ", "OP_MAYORQ",
		"OP_MENORIGUAL", "OP_MAYORIGUAL", "OP_AND", "OP_OR", "OP_NOT", "OP_DECLARATION",
		"PUNTO", "COMA", "PUNTO_Y_COMA", "DOS_PUNTOS", "PAR_IZQ", "PAR_DER",
		"LLAVE_IZQ", "LLAVE_DER", "CORCHETE_IZQ", "CORCHETE_DER", "INT_LITERAL",
		"FLOAT_LITERAL", "STRING_LITERAL", "RUNE_LITERAL", "ID", "WS", "COMMENT",
		"MULTILINE_COMMENT",
	}
	staticData.RuleNames = []string{
		"programa", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 67, 27, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 3, 1, 14, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		5, 1, 22, 8, 1, 10, 1, 12, 1, 25, 9, 1, 1, 1, 0, 1, 2, 2, 0, 2, 0, 2, 1,
		0, 34, 35, 1, 0, 32, 33, 27, 0, 4, 1, 0, 0, 0, 2, 13, 1, 0, 0, 0, 4, 5,
		3, 2, 1, 0, 5, 6, 5, 0, 0, 1, 6, 1, 1, 0, 0, 0, 7, 8, 6, 1, -1, 0, 8, 14,
		5, 60, 0, 0, 9, 10, 5, 54, 0, 0, 10, 11, 3, 2, 1, 0, 11, 12, 5, 55, 0,
		0, 12, 14, 1, 0, 0, 0, 13, 7, 1, 0, 0, 0, 13, 9, 1, 0, 0, 0, 14, 23, 1,
		0, 0, 0, 15, 16, 10, 4, 0, 0, 16, 17, 7, 0, 0, 0, 17, 22, 3, 2, 1, 5, 18,
		19, 10, 3, 0, 0, 19, 20, 7, 1, 0, 0, 20, 22, 3, 2, 1, 4, 21, 15, 1, 0,
		0, 0, 21, 18, 1, 0, 0, 0, 22, 25, 1, 0, 0, 0, 23, 21, 1, 0, 0, 0, 23, 24,
		1, 0, 0, 0, 24, 3, 1, 0, 0, 0, 25, 23, 1, 0, 0, 0, 3, 13, 21, 23,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// VGrammarInit initializes any static state used to implement VGrammar. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewVGrammar(). You can call this function if you wish to initialize the static state ahead
// of time.
func VGrammarInit() {
	staticData := &VGrammarParserStaticData
	staticData.once.Do(vgrammarParserInit)
}

// NewVGrammar produces a new parser instance for the optional input antlr.TokenStream.
func NewVGrammar(input antlr.TokenStream) *VGrammar {
	VGrammarInit()
	this := new(VGrammar)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "VGrammar.g4"

	return this
}

// VGrammar tokens.
const (
	VGrammarEOF               = antlr.TokenEOF
	VGrammarRW_MAIN           = 1
	VGrammarRW_FN             = 2
	VGrammarRW_MUT            = 3
	VGrammarRW_STRUCT         = 4
	VGrammarRW_IF             = 5
	VGrammarRW_ELSE           = 6
	VGrammarRW_SWITCH         = 7
	VGrammarRW_CASE           = 8
	VGrammarRW_DEFAULT        = 9
	VGrammarRW_FOR            = 10
	VGrammarRW_IN             = 11
	VGrammarRW_BREAK          = 12
	VGrammarRW_CONTINUE       = 13
	VGrammarRW_RETURN         = 14
	VGrammarRW_TRUE           = 15
	VGrammarRW_FALSE          = 16
	VGrammarRW_NIL            = 17
	VGrammarRW_INT            = 18
	VGrammarRW_FLOAT64        = 19
	VGrammarRW_STRING         = 20
	VGrammarRW_BOOL           = 21
	VGrammarRW_RUNE           = 22
	VGrammarRW_PRINT          = 23
	VGrammarRW_PRINTLN        = 24
	VGrammarRW_ATOI           = 25
	VGrammarRW_PARSEFLOAT     = 26
	VGrammarRW_TYPEOF         = 27
	VGrammarRW_APPEND         = 28
	VGrammarRW_LEN            = 29
	VGrammarRW_JOIN           = 30
	VGrammarRW_INDEXOF        = 31
	VGrammarOP_SUMA           = 32
	VGrammarOP_RESTA          = 33
	VGrammarOP_MULT           = 34
	VGrammarOP_DIV            = 35
	VGrammarOP_MOD            = 36
	VGrammarOP_ASSIGN         = 37
	VGrammarOP_ADD_ASSIGN     = 38
	VGrammarOP_SUB_ASSIGN     = 39
	VGrammarOP_IGUAL          = 40
	VGrammarOP_DIFERENTE      = 41
	VGrammarOP_MENORQ         = 42
	VGrammarOP_MAYORQ         = 43
	VGrammarOP_MENORIGUAL     = 44
	VGrammarOP_MAYORIGUAL     = 45
	VGrammarOP_AND            = 46
	VGrammarOP_OR             = 47
	VGrammarOP_NOT            = 48
	VGrammarOP_DECLARATION    = 49
	VGrammarPUNTO             = 50
	VGrammarCOMA              = 51
	VGrammarPUNTO_Y_COMA      = 52
	VGrammarDOS_PUNTOS        = 53
	VGrammarPAR_IZQ           = 54
	VGrammarPAR_DER           = 55
	VGrammarLLAVE_IZQ         = 56
	VGrammarLLAVE_DER         = 57
	VGrammarCORCHETE_IZQ      = 58
	VGrammarCORCHETE_DER      = 59
	VGrammarINT_LITERAL       = 60
	VGrammarFLOAT_LITERAL     = 61
	VGrammarSTRING_LITERAL    = 62
	VGrammarRUNE_LITERAL      = 63
	VGrammarID                = 64
	VGrammarWS                = 65
	VGrammarCOMMENT           = 66
	VGrammarMULTILINE_COMMENT = 67
)

// VGrammar rules.
const (
	VGrammarRULE_programa = 0
	VGrammarRULE_expr     = 1
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
	EOF() antlr.TerminalNode

	// IsProgramaContext differentiates from other interfaces.
	IsProgramaContext()
}

type ProgramaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyProgramaContext() *ProgramaContext {
	var p = new(ProgramaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ProgramaContext) EOF() antlr.TerminalNode {
	return s.GetToken(VGrammarEOF, 0)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, VGrammarRULE_programa)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.expr(0)
	}
	{
		p.SetState(5)
		p.Match(VGrammarEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExprContext is an interface to support dynamic dispatch.
type IExprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsExprContext differentiates from other interfaces.
	IsExprContext()
}

type ExprContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExprContext() *ExprContext {
	var p = new(ExprContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = VGrammarRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = VGrammarRULE_expr

	return p
}

func (s *ExprContext) GetParser() antlr.Parser { return s.parser }

func (s *ExprContext) CopyAll(ctx *ExprContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *ExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MulDivContext struct {
	ExprContext
	op antlr.Token
}

func NewMulDivContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulDivContext {
	var p = new(MulDivContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulDivContext) GetOp() antlr.Token { return s.op }

func (s *MulDivContext) SetOp(v antlr.Token) { s.op = v }

func (s *MulDivContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulDivContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *MulDivContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *MulDivContext) OP_MULT() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_MULT, 0)
}

func (s *MulDivContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_DIV, 0)
}

func (s *MulDivContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterMulDiv(s)
	}
}

func (s *MulDivContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitMulDiv(s)
	}
}

func (s *MulDivContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitMulDiv(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddSubContext struct {
	ExprContext
	op antlr.Token
}

func NewAddSubContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddSubContext {
	var p = new(AddSubContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddSubContext) GetOp() antlr.Token { return s.op }

func (s *AddSubContext) SetOp(v antlr.Token) { s.op = v }

func (s *AddSubContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddSubContext) AllExpr() []IExprContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExprContext); ok {
			len++
		}
	}

	tst := make([]IExprContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExprContext); ok {
			tst[i] = t.(IExprContext)
			i++
		}
	}

	return tst
}

func (s *AddSubContext) Expr(i int) IExprContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *AddSubContext) OP_SUMA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_SUMA, 0)
}

func (s *AddSubContext) OP_RESTA() antlr.TerminalNode {
	return s.GetToken(VGrammarOP_RESTA, 0)
}

func (s *AddSubContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterAddSub(s)
	}
}

func (s *AddSubContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitAddSub(s)
	}
}

func (s *AddSubContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitAddSub(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParensContext struct {
	ExprContext
}

func NewParensContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParensContext {
	var p = new(ParensContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParensContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParensContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_IZQ, 0)
}

func (s *ParensContext) Expr() IExprContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExprContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExprContext)
}

func (s *ParensContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(VGrammarPAR_DER, 0)
}

func (s *ParensContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterParens(s)
	}
}

func (s *ParensContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitParens(s)
	}
}

func (s *ParensContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitParens(s)

	default:
		return t.VisitChildren(s)
	}
}

type IntContext struct {
	ExprContext
}

func NewIntContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IntContext {
	var p = new(IntContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IntContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IntContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(VGrammarINT_LITERAL, 0)
}

func (s *IntContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.EnterInt(s)
	}
}

func (s *IntContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(VGrammarListener); ok {
		listenerT.ExitInt(s)
	}
}

func (s *IntContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case VGrammarVisitor:
		return t.VisitInt(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *VGrammar) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *VGrammar) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, VGrammarRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(13)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case VGrammarINT_LITERAL:
		localctx = NewIntContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(8)
			p.Match(VGrammarINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case VGrammarPAR_IZQ:
		localctx = NewParensContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(9)
			p.Match(VGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(10)
			p.expr(0)
		}
		{
			p.SetState(11)
			p.Match(VGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(23)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(21)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 1, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulDivContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(15)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(16)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulDivContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VGrammarOP_MULT || _la == VGrammarOP_DIV) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulDivContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(17)
					p.expr(5)
				}

			case 2:
				localctx = NewAddSubContext(p, NewExprContext(p, _parentctx, _parentState))
				p.PushNewRecursionContext(localctx, _startState, VGrammarRULE_expr)
				p.SetState(18)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(19)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddSubContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == VGrammarOP_SUMA || _la == VGrammarOP_RESTA) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddSubContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(20)
					p.expr(4)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.UnrollRecursionContexts(_parentctx)
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

func (p *VGrammar) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 1:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *VGrammar) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
