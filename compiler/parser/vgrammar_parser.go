// Code generated from parser/vGrammar.g4 by ANTLR 4.13.2. DO NOT EDIT.

package compiler // vGrammar
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

type vGrammar struct {
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
		"programa", "funcionPrincipal", "bloque", "sentencia", "declaracion",
		"asignacion", "ifStatement", "expresion",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 67, 99, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 28, 8, 2, 10, 2, 12, 2, 31, 9, 2,
		1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 1, 3, 3, 3, 43, 8,
		3, 1, 4, 1, 4, 1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 1, 5, 1, 5, 1, 6, 1, 6, 1,
		6, 1, 6, 1, 6, 3, 6, 59, 8, 6, 1, 6, 1, 6, 3, 6, 63, 8, 6, 1, 7, 1, 7,
		1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 77, 8,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 5, 7, 94, 8, 7, 10, 7, 12, 7, 97, 9, 7, 1, 7, 0, 1,
		14, 8, 0, 2, 4, 6, 8, 10, 12, 14, 0, 0, 108, 0, 16, 1, 0, 0, 0, 2, 19,
		1, 0, 0, 0, 4, 25, 1, 0, 0, 0, 6, 42, 1, 0, 0, 0, 8, 44, 1, 0, 0, 0, 10,
		49, 1, 0, 0, 0, 12, 53, 1, 0, 0, 0, 14, 76, 1, 0, 0, 0, 16, 17, 3, 2, 1,
		0, 17, 18, 5, 0, 0, 1, 18, 1, 1, 0, 0, 0, 19, 20, 5, 2, 0, 0, 20, 21, 5,
		1, 0, 0, 21, 22, 5, 54, 0, 0, 22, 23, 5, 55, 0, 0, 23, 24, 3, 4, 2, 0,
		24, 3, 1, 0, 0, 0, 25, 29, 5, 56, 0, 0, 26, 28, 3, 6, 3, 0, 27, 26, 1,
		0, 0, 0, 28, 31, 1, 0, 0, 0, 29, 27, 1, 0, 0, 0, 29, 30, 1, 0, 0, 0, 30,
		32, 1, 0, 0, 0, 31, 29, 1, 0, 0, 0, 32, 33, 5, 57, 0, 0, 33, 5, 1, 0, 0,
		0, 34, 43, 3, 8, 4, 0, 35, 43, 3, 10, 5, 0, 36, 43, 3, 12, 6, 0, 37, 38,
		5, 23, 0, 0, 38, 39, 5, 54, 0, 0, 39, 40, 3, 14, 7, 0, 40, 41, 5, 55, 0,
		0, 41, 43, 1, 0, 0, 0, 42, 34, 1, 0, 0, 0, 42, 35, 1, 0, 0, 0, 42, 36,
		1, 0, 0, 0, 42, 37, 1, 0, 0, 0, 43, 7, 1, 0, 0, 0, 44, 45, 5, 3, 0, 0,
		45, 46, 5, 64, 0, 0, 46, 47, 5, 49, 0, 0, 47, 48, 3, 14, 7, 0, 48, 9, 1,
		0, 0, 0, 49, 50, 5, 64, 0, 0, 50, 51, 5, 37, 0, 0, 51, 52, 3, 14, 7, 0,
		52, 11, 1, 0, 0, 0, 53, 54, 5, 5, 0, 0, 54, 55, 3, 14, 7, 0, 55, 58, 3,
		4, 2, 0, 56, 57, 5, 6, 0, 0, 57, 59, 3, 12, 6, 0, 58, 56, 1, 0, 0, 0, 58,
		59, 1, 0, 0, 0, 59, 62, 1, 0, 0, 0, 60, 61, 5, 6, 0, 0, 61, 63, 3, 4, 2,
		0, 62, 60, 1, 0, 0, 0, 62, 63, 1, 0, 0, 0, 63, 13, 1, 0, 0, 0, 64, 65,
		6, 7, -1, 0, 65, 66, 5, 54, 0, 0, 66, 67, 3, 14, 7, 0, 67, 68, 5, 55, 0,
		0, 68, 77, 1, 0, 0, 0, 69, 77, 5, 60, 0, 0, 70, 77, 5, 61, 0, 0, 71, 77,
		5, 62, 0, 0, 72, 77, 5, 63, 0, 0, 73, 77, 5, 15, 0, 0, 74, 77, 5, 16, 0,
		0, 75, 77, 5, 64, 0, 0, 76, 64, 1, 0, 0, 0, 76, 69, 1, 0, 0, 0, 76, 70,
		1, 0, 0, 0, 76, 71, 1, 0, 0, 0, 76, 72, 1, 0, 0, 0, 76, 73, 1, 0, 0, 0,
		76, 74, 1, 0, 0, 0, 76, 75, 1, 0, 0, 0, 77, 95, 1, 0, 0, 0, 78, 79, 10,
		13, 0, 0, 79, 80, 5, 32, 0, 0, 80, 94, 3, 14, 7, 14, 81, 82, 10, 12, 0,
		0, 82, 83, 5, 33, 0, 0, 83, 94, 3, 14, 7, 13, 84, 85, 10, 11, 0, 0, 85,
		86, 5, 34, 0, 0, 86, 94, 3, 14, 7, 12, 87, 88, 10, 10, 0, 0, 88, 89, 5,
		35, 0, 0, 89, 94, 3, 14, 7, 11, 90, 91, 10, 9, 0, 0, 91, 92, 5, 36, 0,
		0, 92, 94, 3, 14, 7, 10, 93, 78, 1, 0, 0, 0, 93, 81, 1, 0, 0, 0, 93, 84,
		1, 0, 0, 0, 93, 87, 1, 0, 0, 0, 93, 90, 1, 0, 0, 0, 94, 97, 1, 0, 0, 0,
		95, 93, 1, 0, 0, 0, 95, 96, 1, 0, 0, 0, 96, 15, 1, 0, 0, 0, 97, 95, 1,
		0, 0, 0, 7, 29, 42, 58, 62, 76, 93, 95,
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

// vGrammarInit initializes any static state used to implement vGrammar. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewvGrammar(). You can call this function if you wish to initialize the static state ahead
// of time.
func VGrammarInit() {
	staticData := &VGrammarParserStaticData
	staticData.once.Do(vgrammarParserInit)
}

// NewvGrammar produces a new parser instance for the optional input antlr.TokenStream.
func NewvGrammar(input antlr.TokenStream) *vGrammar {
	VGrammarInit()
	this := new(vGrammar)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &VGrammarParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "vGrammar.g4"

	return this
}

// vGrammar tokens.
const (
	vGrammarEOF               = antlr.TokenEOF
	vGrammarRW_MAIN           = 1
	vGrammarRW_FN             = 2
	vGrammarRW_MUT            = 3
	vGrammarRW_STRUCT         = 4
	vGrammarRW_IF             = 5
	vGrammarRW_ELSE           = 6
	vGrammarRW_SWITCH         = 7
	vGrammarRW_CASE           = 8
	vGrammarRW_DEFAULT        = 9
	vGrammarRW_FOR            = 10
	vGrammarRW_IN             = 11
	vGrammarRW_BREAK          = 12
	vGrammarRW_CONTINUE       = 13
	vGrammarRW_RETURN         = 14
	vGrammarRW_TRUE           = 15
	vGrammarRW_FALSE          = 16
	vGrammarRW_NIL            = 17
	vGrammarRW_INT            = 18
	vGrammarRW_FLOAT64        = 19
	vGrammarRW_STRING         = 20
	vGrammarRW_BOOL           = 21
	vGrammarRW_RUNE           = 22
	vGrammarRW_PRINT          = 23
	vGrammarRW_PRINTLN        = 24
	vGrammarRW_ATOI           = 25
	vGrammarRW_PARSEFLOAT     = 26
	vGrammarRW_TYPEOF         = 27
	vGrammarRW_APPEND         = 28
	vGrammarRW_LEN            = 29
	vGrammarRW_JOIN           = 30
	vGrammarRW_INDEXOF        = 31
	vGrammarOP_SUMA           = 32
	vGrammarOP_RESTA          = 33
	vGrammarOP_MULT           = 34
	vGrammarOP_DIV            = 35
	vGrammarOP_MOD            = 36
	vGrammarOP_ASSIGN         = 37
	vGrammarOP_ADD_ASSIGN     = 38
	vGrammarOP_SUB_ASSIGN     = 39
	vGrammarOP_IGUAL          = 40
	vGrammarOP_DIFERENTE      = 41
	vGrammarOP_MENORQ         = 42
	vGrammarOP_MAYORQ         = 43
	vGrammarOP_MENORIGUAL     = 44
	vGrammarOP_MAYORIGUAL     = 45
	vGrammarOP_AND            = 46
	vGrammarOP_OR             = 47
	vGrammarOP_NOT            = 48
	vGrammarOP_DECLARATION    = 49
	vGrammarPUNTO             = 50
	vGrammarCOMA              = 51
	vGrammarPUNTO_Y_COMA      = 52
	vGrammarDOS_PUNTOS        = 53
	vGrammarPAR_IZQ           = 54
	vGrammarPAR_DER           = 55
	vGrammarLLAVE_IZQ         = 56
	vGrammarLLAVE_DER         = 57
	vGrammarCORCHETE_IZQ      = 58
	vGrammarCORCHETE_DER      = 59
	vGrammarINT_LITERAL       = 60
	vGrammarFLOAT_LITERAL     = 61
	vGrammarSTRING_LITERAL    = 62
	vGrammarRUNE_LITERAL      = 63
	vGrammarID                = 64
	vGrammarWS                = 65
	vGrammarCOMMENT           = 66
	vGrammarMULTILINE_COMMENT = 67
)

// vGrammar rules.
const (
	vGrammarRULE_programa         = 0
	vGrammarRULE_funcionPrincipal = 1
	vGrammarRULE_bloque           = 2
	vGrammarRULE_sentencia        = 3
	vGrammarRULE_declaracion      = 4
	vGrammarRULE_asignacion       = 5
	vGrammarRULE_ifStatement      = 6
	vGrammarRULE_expresion        = 7
)

// IProgramaContext is an interface to support dynamic dispatch.
type IProgramaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FuncionPrincipal() IFuncionPrincipalContext
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
	p.RuleIndex = vGrammarRULE_programa
	return p
}

func InitEmptyProgramaContext(p *ProgramaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_programa
}

func (*ProgramaContext) IsProgramaContext() {}

func NewProgramaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ProgramaContext {
	var p = new(ProgramaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = vGrammarRULE_programa

	return p
}

func (s *ProgramaContext) GetParser() antlr.Parser { return s.parser }

func (s *ProgramaContext) FuncionPrincipal() IFuncionPrincipalContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFuncionPrincipalContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFuncionPrincipalContext)
}

func (s *ProgramaContext) EOF() antlr.TerminalNode {
	return s.GetToken(vGrammarEOF, 0)
}

func (s *ProgramaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ProgramaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ProgramaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.EnterPrograma(s)
	}
}

func (s *ProgramaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.ExitPrograma(s)
	}
}

func (s *ProgramaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case vGrammarVisitor:
		return t.VisitPrograma(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *vGrammar) Programa() (localctx IProgramaContext) {
	localctx = NewProgramaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, vGrammarRULE_programa)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(16)
		p.FuncionPrincipal()
	}
	{
		p.SetState(17)
		p.Match(vGrammarEOF)
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

// IFuncionPrincipalContext is an interface to support dynamic dispatch.
type IFuncionPrincipalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_FN() antlr.TerminalNode
	RW_MAIN() antlr.TerminalNode
	PAR_IZQ() antlr.TerminalNode
	PAR_DER() antlr.TerminalNode
	Bloque() IBloqueContext

	// IsFuncionPrincipalContext differentiates from other interfaces.
	IsFuncionPrincipalContext()
}

type FuncionPrincipalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFuncionPrincipalContext() *FuncionPrincipalContext {
	var p = new(FuncionPrincipalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_funcionPrincipal
	return p
}

func InitEmptyFuncionPrincipalContext(p *FuncionPrincipalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_funcionPrincipal
}

func (*FuncionPrincipalContext) IsFuncionPrincipalContext() {}

func NewFuncionPrincipalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FuncionPrincipalContext {
	var p = new(FuncionPrincipalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = vGrammarRULE_funcionPrincipal

	return p
}

func (s *FuncionPrincipalContext) GetParser() antlr.Parser { return s.parser }

func (s *FuncionPrincipalContext) RW_FN() antlr.TerminalNode {
	return s.GetToken(vGrammarRW_FN, 0)
}

func (s *FuncionPrincipalContext) RW_MAIN() antlr.TerminalNode {
	return s.GetToken(vGrammarRW_MAIN, 0)
}

func (s *FuncionPrincipalContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(vGrammarPAR_IZQ, 0)
}

func (s *FuncionPrincipalContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(vGrammarPAR_DER, 0)
}

func (s *FuncionPrincipalContext) Bloque() IBloqueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBloqueContext)
}

func (s *FuncionPrincipalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FuncionPrincipalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FuncionPrincipalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.EnterFuncionPrincipal(s)
	}
}

func (s *FuncionPrincipalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.ExitFuncionPrincipal(s)
	}
}

func (s *FuncionPrincipalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case vGrammarVisitor:
		return t.VisitFuncionPrincipal(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *vGrammar) FuncionPrincipal() (localctx IFuncionPrincipalContext) {
	localctx = NewFuncionPrincipalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, vGrammarRULE_funcionPrincipal)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(19)
		p.Match(vGrammarRW_FN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(20)
		p.Match(vGrammarRW_MAIN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(21)
		p.Match(vGrammarPAR_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(22)
		p.Match(vGrammarPAR_DER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(23)
		p.Bloque()
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

// IBloqueContext is an interface to support dynamic dispatch.
type IBloqueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LLAVE_IZQ() antlr.TerminalNode
	LLAVE_DER() antlr.TerminalNode
	AllSentencia() []ISentenciaContext
	Sentencia(i int) ISentenciaContext

	// IsBloqueContext differentiates from other interfaces.
	IsBloqueContext()
}

type BloqueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBloqueContext() *BloqueContext {
	var p = new(BloqueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_bloque
	return p
}

func InitEmptyBloqueContext(p *BloqueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_bloque
}

func (*BloqueContext) IsBloqueContext() {}

func NewBloqueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BloqueContext {
	var p = new(BloqueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = vGrammarRULE_bloque

	return p
}

func (s *BloqueContext) GetParser() antlr.Parser { return s.parser }

func (s *BloqueContext) LLAVE_IZQ() antlr.TerminalNode {
	return s.GetToken(vGrammarLLAVE_IZQ, 0)
}

func (s *BloqueContext) LLAVE_DER() antlr.TerminalNode {
	return s.GetToken(vGrammarLLAVE_DER, 0)
}

func (s *BloqueContext) AllSentencia() []ISentenciaContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISentenciaContext); ok {
			len++
		}
	}

	tst := make([]ISentenciaContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISentenciaContext); ok {
			tst[i] = t.(ISentenciaContext)
			i++
		}
	}

	return tst
}

func (s *BloqueContext) Sentencia(i int) ISentenciaContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISentenciaContext); ok {
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

	return t.(ISentenciaContext)
}

func (s *BloqueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BloqueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BloqueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.EnterBloque(s)
	}
}

func (s *BloqueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.ExitBloque(s)
	}
}

func (s *BloqueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case vGrammarVisitor:
		return t.VisitBloque(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *vGrammar) Bloque() (localctx IBloqueContext) {
	localctx = NewBloqueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, vGrammarRULE_bloque)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(25)
		p.Match(vGrammarLLAVE_IZQ)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(29)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64((_la-3)) & ^0x3f) == 0 && ((int64(1)<<(_la-3))&2305843009214742533) != 0 {
		{
			p.SetState(26)
			p.Sentencia()
		}

		p.SetState(31)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(32)
		p.Match(vGrammarLLAVE_DER)
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

// ISentenciaContext is an interface to support dynamic dispatch.
type ISentenciaContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Declaracion() IDeclaracionContext
	Asignacion() IAsignacionContext
	IfStatement() IIfStatementContext
	RW_PRINT() antlr.TerminalNode
	PAR_IZQ() antlr.TerminalNode
	Expresion() IExpresionContext
	PAR_DER() antlr.TerminalNode

	// IsSentenciaContext differentiates from other interfaces.
	IsSentenciaContext()
}

type SentenciaContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySentenciaContext() *SentenciaContext {
	var p = new(SentenciaContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_sentencia
	return p
}

func InitEmptySentenciaContext(p *SentenciaContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_sentencia
}

func (*SentenciaContext) IsSentenciaContext() {}

func NewSentenciaContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SentenciaContext {
	var p = new(SentenciaContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = vGrammarRULE_sentencia

	return p
}

func (s *SentenciaContext) GetParser() antlr.Parser { return s.parser }

func (s *SentenciaContext) Declaracion() IDeclaracionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDeclaracionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDeclaracionContext)
}

func (s *SentenciaContext) Asignacion() IAsignacionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAsignacionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAsignacionContext)
}

func (s *SentenciaContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *SentenciaContext) RW_PRINT() antlr.TerminalNode {
	return s.GetToken(vGrammarRW_PRINT, 0)
}

func (s *SentenciaContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(vGrammarPAR_IZQ, 0)
}

func (s *SentenciaContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *SentenciaContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(vGrammarPAR_DER, 0)
}

func (s *SentenciaContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SentenciaContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SentenciaContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.EnterSentencia(s)
	}
}

func (s *SentenciaContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.ExitSentencia(s)
	}
}

func (s *SentenciaContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case vGrammarVisitor:
		return t.VisitSentencia(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *vGrammar) Sentencia() (localctx ISentenciaContext) {
	localctx = NewSentenciaContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, vGrammarRULE_sentencia)
	p.SetState(42)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case vGrammarRW_MUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(34)
			p.Declaracion()
		}

	case vGrammarID:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(35)
			p.Asignacion()
		}

	case vGrammarRW_IF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(36)
			p.IfStatement()
		}

	case vGrammarRW_PRINT:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(37)
			p.Match(vGrammarRW_PRINT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(38)
			p.Match(vGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(39)
			p.expresion(0)
		}
		{
			p.SetState(40)
			p.Match(vGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
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

// IDeclaracionContext is an interface to support dynamic dispatch.
type IDeclaracionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_MUT() antlr.TerminalNode
	ID() antlr.TerminalNode
	OP_DECLARATION() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsDeclaracionContext differentiates from other interfaces.
	IsDeclaracionContext()
}

type DeclaracionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDeclaracionContext() *DeclaracionContext {
	var p = new(DeclaracionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_declaracion
	return p
}

func InitEmptyDeclaracionContext(p *DeclaracionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_declaracion
}

func (*DeclaracionContext) IsDeclaracionContext() {}

func NewDeclaracionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DeclaracionContext {
	var p = new(DeclaracionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = vGrammarRULE_declaracion

	return p
}

func (s *DeclaracionContext) GetParser() antlr.Parser { return s.parser }

func (s *DeclaracionContext) RW_MUT() antlr.TerminalNode {
	return s.GetToken(vGrammarRW_MUT, 0)
}

func (s *DeclaracionContext) ID() antlr.TerminalNode {
	return s.GetToken(vGrammarID, 0)
}

func (s *DeclaracionContext) OP_DECLARATION() antlr.TerminalNode {
	return s.GetToken(vGrammarOP_DECLARATION, 0)
}

func (s *DeclaracionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *DeclaracionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DeclaracionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DeclaracionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.EnterDeclaracion(s)
	}
}

func (s *DeclaracionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.ExitDeclaracion(s)
	}
}

func (s *DeclaracionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case vGrammarVisitor:
		return t.VisitDeclaracion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *vGrammar) Declaracion() (localctx IDeclaracionContext) {
	localctx = NewDeclaracionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, vGrammarRULE_declaracion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(44)
		p.Match(vGrammarRW_MUT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(45)
		p.Match(vGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(46)
		p.Match(vGrammarOP_DECLARATION)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(47)
		p.expresion(0)
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

// IAsignacionContext is an interface to support dynamic dispatch.
type IAsignacionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ID() antlr.TerminalNode
	OP_ASSIGN() antlr.TerminalNode
	Expresion() IExpresionContext

	// IsAsignacionContext differentiates from other interfaces.
	IsAsignacionContext()
}

type AsignacionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAsignacionContext() *AsignacionContext {
	var p = new(AsignacionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_asignacion
	return p
}

func InitEmptyAsignacionContext(p *AsignacionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_asignacion
}

func (*AsignacionContext) IsAsignacionContext() {}

func NewAsignacionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AsignacionContext {
	var p = new(AsignacionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = vGrammarRULE_asignacion

	return p
}

func (s *AsignacionContext) GetParser() antlr.Parser { return s.parser }

func (s *AsignacionContext) ID() antlr.TerminalNode {
	return s.GetToken(vGrammarID, 0)
}

func (s *AsignacionContext) OP_ASSIGN() antlr.TerminalNode {
	return s.GetToken(vGrammarOP_ASSIGN, 0)
}

func (s *AsignacionContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *AsignacionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AsignacionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AsignacionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.EnterAsignacion(s)
	}
}

func (s *AsignacionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.ExitAsignacion(s)
	}
}

func (s *AsignacionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case vGrammarVisitor:
		return t.VisitAsignacion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *vGrammar) Asignacion() (localctx IAsignacionContext) {
	localctx = NewAsignacionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, vGrammarRULE_asignacion)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(49)
		p.Match(vGrammarID)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(50)
		p.Match(vGrammarOP_ASSIGN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(51)
		p.expresion(0)
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

// IIfStatementContext is an interface to support dynamic dispatch.
type IIfStatementContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RW_IF() antlr.TerminalNode
	Expresion() IExpresionContext
	AllBloque() []IBloqueContext
	Bloque(i int) IBloqueContext
	AllRW_ELSE() []antlr.TerminalNode
	RW_ELSE(i int) antlr.TerminalNode
	IfStatement() IIfStatementContext

	// IsIfStatementContext differentiates from other interfaces.
	IsIfStatementContext()
}

type IfStatementContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIfStatementContext() *IfStatementContext {
	var p = new(IfStatementContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_ifStatement
	return p
}

func InitEmptyIfStatementContext(p *IfStatementContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_ifStatement
}

func (*IfStatementContext) IsIfStatementContext() {}

func NewIfStatementContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfStatementContext {
	var p = new(IfStatementContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = vGrammarRULE_ifStatement

	return p
}

func (s *IfStatementContext) GetParser() antlr.Parser { return s.parser }

func (s *IfStatementContext) RW_IF() antlr.TerminalNode {
	return s.GetToken(vGrammarRW_IF, 0)
}

func (s *IfStatementContext) Expresion() IExpresionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpresionContext)
}

func (s *IfStatementContext) AllBloque() []IBloqueContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IBloqueContext); ok {
			len++
		}
	}

	tst := make([]IBloqueContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IBloqueContext); ok {
			tst[i] = t.(IBloqueContext)
			i++
		}
	}

	return tst
}

func (s *IfStatementContext) Bloque(i int) IBloqueContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBloqueContext); ok {
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

	return t.(IBloqueContext)
}

func (s *IfStatementContext) AllRW_ELSE() []antlr.TerminalNode {
	return s.GetTokens(vGrammarRW_ELSE)
}

func (s *IfStatementContext) RW_ELSE(i int) antlr.TerminalNode {
	return s.GetToken(vGrammarRW_ELSE, i)
}

func (s *IfStatementContext) IfStatement() IIfStatementContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIfStatementContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIfStatementContext)
}

func (s *IfStatementContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfStatementContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfStatementContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.EnterIfStatement(s)
	}
}

func (s *IfStatementContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.ExitIfStatement(s)
	}
}

func (s *IfStatementContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case vGrammarVisitor:
		return t.VisitIfStatement(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *vGrammar) IfStatement() (localctx IIfStatementContext) {
	localctx = NewIfStatementContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, vGrammarRULE_ifStatement)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(53)
		p.Match(vGrammarRW_IF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(54)
		p.expresion(0)
	}
	{
		p.SetState(55)
		p.Bloque()
	}
	p.SetState(58)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(56)
			p.Match(vGrammarRW_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(57)
			p.IfStatement()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}
	p.SetState(62)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(60)
			p.Match(vGrammarRW_ELSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(61)
			p.Bloque()
		}

	} else if p.HasError() { // JIM
		goto errorExit
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

// IExpresionContext is an interface to support dynamic dispatch.
type IExpresionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PAR_IZQ() antlr.TerminalNode
	AllExpresion() []IExpresionContext
	Expresion(i int) IExpresionContext
	PAR_DER() antlr.TerminalNode
	INT_LITERAL() antlr.TerminalNode
	FLOAT_LITERAL() antlr.TerminalNode
	STRING_LITERAL() antlr.TerminalNode
	RUNE_LITERAL() antlr.TerminalNode
	RW_TRUE() antlr.TerminalNode
	RW_FALSE() antlr.TerminalNode
	ID() antlr.TerminalNode
	OP_SUMA() antlr.TerminalNode
	OP_RESTA() antlr.TerminalNode
	OP_MULT() antlr.TerminalNode
	OP_DIV() antlr.TerminalNode
	OP_MOD() antlr.TerminalNode

	// IsExpresionContext differentiates from other interfaces.
	IsExpresionContext()
}

type ExpresionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpresionContext() *ExpresionContext {
	var p = new(ExpresionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_expresion
	return p
}

func InitEmptyExpresionContext(p *ExpresionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = vGrammarRULE_expresion
}

func (*ExpresionContext) IsExpresionContext() {}

func NewExpresionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpresionContext {
	var p = new(ExpresionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = vGrammarRULE_expresion

	return p
}

func (s *ExpresionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpresionContext) PAR_IZQ() antlr.TerminalNode {
	return s.GetToken(vGrammarPAR_IZQ, 0)
}

func (s *ExpresionContext) AllExpresion() []IExpresionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpresionContext); ok {
			len++
		}
	}

	tst := make([]IExpresionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpresionContext); ok {
			tst[i] = t.(IExpresionContext)
			i++
		}
	}

	return tst
}

func (s *ExpresionContext) Expresion(i int) IExpresionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpresionContext); ok {
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

	return t.(IExpresionContext)
}

func (s *ExpresionContext) PAR_DER() antlr.TerminalNode {
	return s.GetToken(vGrammarPAR_DER, 0)
}

func (s *ExpresionContext) INT_LITERAL() antlr.TerminalNode {
	return s.GetToken(vGrammarINT_LITERAL, 0)
}

func (s *ExpresionContext) FLOAT_LITERAL() antlr.TerminalNode {
	return s.GetToken(vGrammarFLOAT_LITERAL, 0)
}

func (s *ExpresionContext) STRING_LITERAL() antlr.TerminalNode {
	return s.GetToken(vGrammarSTRING_LITERAL, 0)
}

func (s *ExpresionContext) RUNE_LITERAL() antlr.TerminalNode {
	return s.GetToken(vGrammarRUNE_LITERAL, 0)
}

func (s *ExpresionContext) RW_TRUE() antlr.TerminalNode {
	return s.GetToken(vGrammarRW_TRUE, 0)
}

func (s *ExpresionContext) RW_FALSE() antlr.TerminalNode {
	return s.GetToken(vGrammarRW_FALSE, 0)
}

func (s *ExpresionContext) ID() antlr.TerminalNode {
	return s.GetToken(vGrammarID, 0)
}

func (s *ExpresionContext) OP_SUMA() antlr.TerminalNode {
	return s.GetToken(vGrammarOP_SUMA, 0)
}

func (s *ExpresionContext) OP_RESTA() antlr.TerminalNode {
	return s.GetToken(vGrammarOP_RESTA, 0)
}

func (s *ExpresionContext) OP_MULT() antlr.TerminalNode {
	return s.GetToken(vGrammarOP_MULT, 0)
}

func (s *ExpresionContext) OP_DIV() antlr.TerminalNode {
	return s.GetToken(vGrammarOP_DIV, 0)
}

func (s *ExpresionContext) OP_MOD() antlr.TerminalNode {
	return s.GetToken(vGrammarOP_MOD, 0)
}

func (s *ExpresionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpresionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpresionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.EnterExpresion(s)
	}
}

func (s *ExpresionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(vGrammarListener); ok {
		listenerT.ExitExpresion(s)
	}
}

func (s *ExpresionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case vGrammarVisitor:
		return t.VisitExpresion(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *vGrammar) Expresion() (localctx IExpresionContext) {
	return p.expresion(0)
}

func (p *vGrammar) expresion(_p int) (localctx IExpresionContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExpresionContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpresionContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 14
	p.EnterRecursionRule(localctx, 14, vGrammarRULE_expresion, _p)
	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(76)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case vGrammarPAR_IZQ:
		{
			p.SetState(65)
			p.Match(vGrammarPAR_IZQ)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(66)
			p.expresion(0)
		}
		{
			p.SetState(67)
			p.Match(vGrammarPAR_DER)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case vGrammarINT_LITERAL:
		{
			p.SetState(69)
			p.Match(vGrammarINT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case vGrammarFLOAT_LITERAL:
		{
			p.SetState(70)
			p.Match(vGrammarFLOAT_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case vGrammarSTRING_LITERAL:
		{
			p.SetState(71)
			p.Match(vGrammarSTRING_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case vGrammarRUNE_LITERAL:
		{
			p.SetState(72)
			p.Match(vGrammarRUNE_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case vGrammarRW_TRUE:
		{
			p.SetState(73)
			p.Match(vGrammarRW_TRUE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case vGrammarRW_FALSE:
		{
			p.SetState(74)
			p.Match(vGrammarRW_FALSE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case vGrammarID:
		{
			p.SetState(75)
			p.Match(vGrammarID)
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
	p.SetState(95)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(93)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 5, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, vGrammarRULE_expresion)
				p.SetState(78)

				if !(p.Precpred(p.GetParserRuleContext(), 13)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 13)", ""))
					goto errorExit
				}
				{
					p.SetState(79)
					p.Match(vGrammarOP_SUMA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(80)
					p.expresion(14)
				}

			case 2:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, vGrammarRULE_expresion)
				p.SetState(81)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
					goto errorExit
				}
				{
					p.SetState(82)
					p.Match(vGrammarOP_RESTA)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(83)
					p.expresion(13)
				}

			case 3:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, vGrammarRULE_expresion)
				p.SetState(84)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
					goto errorExit
				}
				{
					p.SetState(85)
					p.Match(vGrammarOP_MULT)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(86)
					p.expresion(12)
				}

			case 4:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, vGrammarRULE_expresion)
				p.SetState(87)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
					goto errorExit
				}
				{
					p.SetState(88)
					p.Match(vGrammarOP_DIV)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(89)
					p.expresion(11)
				}

			case 5:
				localctx = NewExpresionContext(p, _parentctx, _parentState)
				p.PushNewRecursionContext(localctx, _startState, vGrammarRULE_expresion)
				p.SetState(90)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
					goto errorExit
				}
				{
					p.SetState(91)
					p.Match(vGrammarOP_MOD)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				{
					p.SetState(92)
					p.expresion(10)
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(97)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 6, p.GetParserRuleContext())
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

func (p *vGrammar) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 7:
		var t *ExpresionContext = nil
		if localctx != nil {
			t = localctx.(*ExpresionContext)
		}
		return p.Expresion_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *vGrammar) Expresion_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 13)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
