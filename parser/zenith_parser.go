// Code generated from ZenithParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // ZenithParser

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

type ZenithParser struct {
	*antlr.BaseParser
}

var ZenithParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func zenithparserParserInit() {
	staticData := &ZenithParserParserStaticData
	staticData.LiteralNames = []string{
		"", "'$'", "'$?'", "'&'", "'&?'", "'@'", "'~'", "'!'", "'+'", "'-'",
		"'*'", "'/'", "'%'", "'if'", "'else'", "'true'", "'false'", "'null'",
		"'='", "':='", "", "", "", "'('", "')'", "'{'", "'}'", "';'",
	}
	staticData.SymbolicNames = []string{
		"", "OWN_PTR", "NULL_OWN_PTR", "PTR", "NULL_PTR", "AT", "BIT_NOT", "NOT",
		"PLUS", "MINUS", "TIMES", "DIVIDE", "REM", "IF", "ELSE", "TRUE", "FALSE",
		"NULL", "ASSIGN", "INIT_ASSIGN", "TYPE", "NUM", "ID", "LPAREN", "RPAREN",
		"LBRACE", "RBRACE", "SEMICOLON", "NL", "IGNORED",
	}
	staticData.RuleNames = []string{
		"fileStat", "endedStat", "lineEnd", "stat", "type", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 29, 141, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 1, 0, 4, 0, 14, 8, 0, 11, 0, 12, 0, 15, 1, 0, 3, 0, 19,
		8, 0, 1, 1, 3, 1, 22, 8, 1, 1, 1, 1, 1, 1, 1, 1, 2, 1, 2, 5, 2, 29, 8,
		2, 10, 2, 12, 2, 32, 9, 2, 1, 2, 1, 2, 5, 2, 36, 8, 2, 10, 2, 12, 2, 39,
		9, 2, 1, 2, 3, 2, 42, 8, 2, 1, 3, 1, 3, 3, 3, 46, 8, 3, 1, 3, 1, 3, 3,
		3, 50, 8, 3, 1, 3, 1, 3, 1, 3, 3, 3, 55, 8, 3, 1, 3, 5, 3, 58, 8, 3, 10,
		3, 12, 3, 61, 9, 3, 1, 3, 3, 3, 64, 8, 3, 1, 3, 3, 3, 67, 8, 3, 1, 3, 1,
		3, 3, 3, 71, 8, 3, 1, 4, 1, 4, 1, 4, 3, 4, 76, 8, 4, 1, 5, 1, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 3, 5, 84, 8, 5, 1, 5, 1, 5, 3, 5, 88, 8, 5, 1, 5, 1,
		5, 1, 5, 1, 5, 1, 5, 3, 5, 95, 8, 5, 1, 5, 1, 5, 3, 5, 99, 8, 5, 1, 5,
		1, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 107, 8, 5, 1, 5, 1, 5, 1, 5, 3, 5,
		112, 8, 5, 1, 5, 1, 5, 1, 5, 1, 5, 3, 5, 118, 8, 5, 1, 5, 1, 5, 1, 5, 1,
		5, 3, 5, 124, 8, 5, 1, 5, 1, 5, 3, 5, 128, 8, 5, 1, 5, 1, 5, 3, 5, 132,
		8, 5, 1, 5, 1, 5, 5, 5, 136, 8, 5, 10, 5, 12, 5, 139, 9, 5, 1, 5, 0, 1,
		10, 6, 0, 2, 4, 6, 8, 10, 0, 6, 1, 0, 27, 28, 1, 0, 1, 4, 1, 0, 15, 17,
		1, 0, 7, 9, 1, 0, 10, 12, 1, 0, 8, 9, 168, 0, 18, 1, 0, 0, 0, 2, 21, 1,
		0, 0, 0, 4, 41, 1, 0, 0, 0, 6, 70, 1, 0, 0, 0, 8, 75, 1, 0, 0, 0, 10, 106,
		1, 0, 0, 0, 12, 14, 3, 2, 1, 0, 13, 12, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0,
		15, 13, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 19, 1, 0, 0, 0, 17, 19, 5,
		0, 0, 1, 18, 13, 1, 0, 0, 0, 18, 17, 1, 0, 0, 0, 19, 1, 1, 0, 0, 0, 20,
		22, 5, 28, 0, 0, 21, 20, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 23, 1, 0,
		0, 0, 23, 24, 3, 6, 3, 0, 24, 25, 3, 4, 2, 0, 25, 3, 1, 0, 0, 0, 26, 30,
		5, 27, 0, 0, 27, 29, 7, 0, 0, 0, 28, 27, 1, 0, 0, 0, 29, 32, 1, 0, 0, 0,
		30, 28, 1, 0, 0, 0, 30, 31, 1, 0, 0, 0, 31, 42, 1, 0, 0, 0, 32, 30, 1,
		0, 0, 0, 33, 37, 5, 28, 0, 0, 34, 36, 7, 0, 0, 0, 35, 34, 1, 0, 0, 0, 36,
		39, 1, 0, 0, 0, 37, 35, 1, 0, 0, 0, 37, 38, 1, 0, 0, 0, 38, 42, 1, 0, 0,
		0, 39, 37, 1, 0, 0, 0, 40, 42, 5, 0, 0, 1, 41, 26, 1, 0, 0, 0, 41, 33,
		1, 0, 0, 0, 41, 40, 1, 0, 0, 0, 42, 5, 1, 0, 0, 0, 43, 45, 5, 22, 0, 0,
		44, 46, 3, 8, 4, 0, 45, 44, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 47, 1,
		0, 0, 0, 47, 49, 5, 19, 0, 0, 48, 50, 5, 28, 0, 0, 49, 48, 1, 0, 0, 0,
		49, 50, 1, 0, 0, 0, 50, 51, 1, 0, 0, 0, 51, 71, 3, 10, 5, 0, 52, 54, 5,
		25, 0, 0, 53, 55, 5, 28, 0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0,
		55, 59, 1, 0, 0, 0, 56, 58, 3, 2, 1, 0, 57, 56, 1, 0, 0, 0, 58, 61, 1,
		0, 0, 0, 59, 57, 1, 0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 63, 1, 0, 0, 0, 61,
		59, 1, 0, 0, 0, 62, 64, 3, 6, 3, 0, 63, 62, 1, 0, 0, 0, 63, 64, 1, 0, 0,
		0, 64, 66, 1, 0, 0, 0, 65, 67, 3, 4, 2, 0, 66, 65, 1, 0, 0, 0, 66, 67,
		1, 0, 0, 0, 67, 68, 1, 0, 0, 0, 68, 71, 5, 26, 0, 0, 69, 71, 3, 10, 5,
		0, 70, 43, 1, 0, 0, 0, 70, 52, 1, 0, 0, 0, 70, 69, 1, 0, 0, 0, 71, 7, 1,
		0, 0, 0, 72, 76, 5, 20, 0, 0, 73, 74, 7, 1, 0, 0, 74, 76, 3, 8, 4, 0, 75,
		72, 1, 0, 0, 0, 75, 73, 1, 0, 0, 0, 76, 9, 1, 0, 0, 0, 77, 78, 6, 5, -1,
		0, 78, 107, 5, 21, 0, 0, 79, 107, 5, 22, 0, 0, 80, 107, 7, 2, 0, 0, 81,
		83, 5, 23, 0, 0, 82, 84, 5, 28, 0, 0, 83, 82, 1, 0, 0, 0, 83, 84, 1, 0,
		0, 0, 84, 85, 1, 0, 0, 0, 85, 87, 3, 10, 5, 0, 86, 88, 5, 28, 0, 0, 87,
		86, 1, 0, 0, 0, 87, 88, 1, 0, 0, 0, 88, 89, 1, 0, 0, 0, 89, 90, 5, 24,
		0, 0, 90, 107, 1, 0, 0, 0, 91, 92, 5, 20, 0, 0, 92, 94, 5, 23, 0, 0, 93,
		95, 5, 28, 0, 0, 94, 93, 1, 0, 0, 0, 94, 95, 1, 0, 0, 0, 95, 96, 1, 0,
		0, 0, 96, 98, 3, 10, 5, 0, 97, 99, 5, 28, 0, 0, 98, 97, 1, 0, 0, 0, 98,
		99, 1, 0, 0, 0, 99, 100, 1, 0, 0, 0, 100, 101, 5, 24, 0, 0, 101, 107, 1,
		0, 0, 0, 102, 103, 7, 1, 0, 0, 103, 107, 3, 10, 5, 5, 104, 105, 7, 3, 0,
		0, 105, 107, 3, 10, 5, 4, 106, 77, 1, 0, 0, 0, 106, 79, 1, 0, 0, 0, 106,
		80, 1, 0, 0, 0, 106, 81, 1, 0, 0, 0, 106, 91, 1, 0, 0, 0, 106, 102, 1,
		0, 0, 0, 106, 104, 1, 0, 0, 0, 107, 137, 1, 0, 0, 0, 108, 109, 10, 3, 0,
		0, 109, 111, 7, 4, 0, 0, 110, 112, 5, 28, 0, 0, 111, 110, 1, 0, 0, 0, 111,
		112, 1, 0, 0, 0, 112, 113, 1, 0, 0, 0, 113, 136, 3, 10, 5, 4, 114, 115,
		10, 2, 0, 0, 115, 117, 7, 5, 0, 0, 116, 118, 5, 28, 0, 0, 117, 116, 1,
		0, 0, 0, 117, 118, 1, 0, 0, 0, 118, 119, 1, 0, 0, 0, 119, 136, 3, 10, 5,
		3, 120, 121, 10, 1, 0, 0, 121, 123, 5, 13, 0, 0, 122, 124, 5, 28, 0, 0,
		123, 122, 1, 0, 0, 0, 123, 124, 1, 0, 0, 0, 124, 125, 1, 0, 0, 0, 125,
		127, 3, 10, 5, 0, 126, 128, 5, 28, 0, 0, 127, 126, 1, 0, 0, 0, 127, 128,
		1, 0, 0, 0, 128, 129, 1, 0, 0, 0, 129, 131, 5, 14, 0, 0, 130, 132, 5, 28,
		0, 0, 131, 130, 1, 0, 0, 0, 131, 132, 1, 0, 0, 0, 132, 133, 1, 0, 0, 0,
		133, 134, 3, 10, 5, 2, 134, 136, 1, 0, 0, 0, 135, 108, 1, 0, 0, 0, 135,
		114, 1, 0, 0, 0, 135, 120, 1, 0, 0, 0, 136, 139, 1, 0, 0, 0, 137, 135,
		1, 0, 0, 0, 137, 138, 1, 0, 0, 0, 138, 11, 1, 0, 0, 0, 139, 137, 1, 0,
		0, 0, 26, 15, 18, 21, 30, 37, 41, 45, 49, 54, 59, 63, 66, 70, 75, 83, 87,
		94, 98, 106, 111, 117, 123, 127, 131, 135, 137,
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

// ZenithParserInit initializes any static state used to implement ZenithParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewZenithParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func ZenithParserInit() {
	staticData := &ZenithParserParserStaticData
	staticData.once.Do(zenithparserParserInit)
}

// NewZenithParser produces a new parser instance for the optional input antlr.TokenStream.
func NewZenithParser(input antlr.TokenStream) *ZenithParser {
	ZenithParserInit()
	this := new(ZenithParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &ZenithParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "ZenithParser.g4"

	return this
}

// ZenithParser tokens.
const (
	ZenithParserEOF          = antlr.TokenEOF
	ZenithParserOWN_PTR      = 1
	ZenithParserNULL_OWN_PTR = 2
	ZenithParserPTR          = 3
	ZenithParserNULL_PTR     = 4
	ZenithParserAT           = 5
	ZenithParserBIT_NOT      = 6
	ZenithParserNOT          = 7
	ZenithParserPLUS         = 8
	ZenithParserMINUS        = 9
	ZenithParserTIMES        = 10
	ZenithParserDIVIDE       = 11
	ZenithParserREM          = 12
	ZenithParserIF           = 13
	ZenithParserELSE         = 14
	ZenithParserTRUE         = 15
	ZenithParserFALSE        = 16
	ZenithParserNULL         = 17
	ZenithParserASSIGN       = 18
	ZenithParserINIT_ASSIGN  = 19
	ZenithParserTYPE         = 20
	ZenithParserNUM          = 21
	ZenithParserID           = 22
	ZenithParserLPAREN       = 23
	ZenithParserRPAREN       = 24
	ZenithParserLBRACE       = 25
	ZenithParserRBRACE       = 26
	ZenithParserSEMICOLON    = 27
	ZenithParserNL           = 28
	ZenithParserIGNORED      = 29
)

// ZenithParser rules.
const (
	ZenithParserRULE_fileStat  = 0
	ZenithParserRULE_endedStat = 1
	ZenithParserRULE_lineEnd   = 2
	ZenithParserRULE_stat      = 3
	ZenithParserRULE_type      = 4
	ZenithParserRULE_expr      = 5
)

// IFileStatContext is an interface to support dynamic dispatch.
type IFileStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllEndedStat() []IEndedStatContext
	EndedStat(i int) IEndedStatContext
	EOF() antlr.TerminalNode

	// IsFileStatContext differentiates from other interfaces.
	IsFileStatContext()
}

type FileStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFileStatContext() *FileStatContext {
	var p = new(FileStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_fileStat
	return p
}

func InitEmptyFileStatContext(p *FileStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_fileStat
}

func (*FileStatContext) IsFileStatContext() {}

func NewFileStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FileStatContext {
	var p = new(FileStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_fileStat

	return p
}

func (s *FileStatContext) GetParser() antlr.Parser { return s.parser }

func (s *FileStatContext) AllEndedStat() []IEndedStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEndedStatContext); ok {
			len++
		}
	}

	tst := make([]IEndedStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEndedStatContext); ok {
			tst[i] = t.(IEndedStatContext)
			i++
		}
	}

	return tst
}

func (s *FileStatContext) EndedStat(i int) IEndedStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndedStatContext); ok {
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

	return t.(IEndedStatContext)
}

func (s *FileStatContext) EOF() antlr.TerminalNode {
	return s.GetToken(ZenithParserEOF, 0)
}

func (s *FileStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FileStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FileStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterFileStat(s)
	}
}

func (s *FileStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitFileStat(s)
	}
}

func (s *FileStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitFileStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) FileStat() (localctx IFileStatContext) {
	localctx = NewFileStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ZenithParserRULE_fileStat)
	var _la int

	p.SetState(18)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ZenithParserOWN_PTR, ZenithParserNULL_OWN_PTR, ZenithParserPTR, ZenithParserNULL_PTR, ZenithParserNOT, ZenithParserPLUS, ZenithParserMINUS, ZenithParserTRUE, ZenithParserFALSE, ZenithParserNULL, ZenithParserTYPE, ZenithParserNUM, ZenithParserID, ZenithParserLPAREN, ZenithParserLBRACE, ZenithParserNL:
		p.EnterOuterAlt(localctx, 1)
		p.SetState(13)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&317948830) != 0) {
			{
				p.SetState(12)
				p.EndedStat()
			}

			p.SetState(15)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case ZenithParserEOF:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(17)
			p.Match(ZenithParserEOF)
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

// IEndedStatContext is an interface to support dynamic dispatch.
type IEndedStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Stat() IStatContext
	LineEnd() ILineEndContext
	NL() antlr.TerminalNode

	// IsEndedStatContext differentiates from other interfaces.
	IsEndedStatContext()
}

type EndedStatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyEndedStatContext() *EndedStatContext {
	var p = new(EndedStatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_endedStat
	return p
}

func InitEmptyEndedStatContext(p *EndedStatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_endedStat
}

func (*EndedStatContext) IsEndedStatContext() {}

func NewEndedStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *EndedStatContext {
	var p = new(EndedStatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_endedStat

	return p
}

func (s *EndedStatContext) GetParser() antlr.Parser { return s.parser }

func (s *EndedStatContext) Stat() IStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *EndedStatContext) LineEnd() ILineEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineEndContext)
}

func (s *EndedStatContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *EndedStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *EndedStatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *EndedStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterEndedStat(s)
	}
}

func (s *EndedStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitEndedStat(s)
	}
}

func (s *EndedStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitEndedStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) EndedStat() (localctx IEndedStatContext) {
	localctx = NewEndedStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ZenithParserRULE_endedStat)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(21)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == ZenithParserNL {
		{
			p.SetState(20)
			p.Match(ZenithParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(23)
		p.Stat()
	}
	{
		p.SetState(24)
		p.LineEnd()
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

// ILineEndContext is an interface to support dynamic dispatch.
type ILineEndContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode
	AllNL() []antlr.TerminalNode
	NL(i int) antlr.TerminalNode
	EOF() antlr.TerminalNode

	// IsLineEndContext differentiates from other interfaces.
	IsLineEndContext()
}

type LineEndContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLineEndContext() *LineEndContext {
	var p = new(LineEndContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_lineEnd
	return p
}

func InitEmptyLineEndContext(p *LineEndContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_lineEnd
}

func (*LineEndContext) IsLineEndContext() {}

func NewLineEndContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LineEndContext {
	var p = new(LineEndContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_lineEnd

	return p
}

func (s *LineEndContext) GetParser() antlr.Parser { return s.parser }

func (s *LineEndContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserSEMICOLON)
}

func (s *LineEndContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserSEMICOLON, i)
}

func (s *LineEndContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserNL)
}

func (s *LineEndContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, i)
}

func (s *LineEndContext) EOF() antlr.TerminalNode {
	return s.GetToken(ZenithParserEOF, 0)
}

func (s *LineEndContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LineEndContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LineEndContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterLineEnd(s)
	}
}

func (s *LineEndContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitLineEnd(s)
	}
}

func (s *LineEndContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitLineEnd(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) LineEnd() (localctx ILineEndContext) {
	localctx = NewLineEndContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ZenithParserRULE_lineEnd)
	var _la int

	var _alt int

	p.SetState(41)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ZenithParserSEMICOLON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(26)
			p.Match(ZenithParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(30)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(27)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ZenithParserSEMICOLON || _la == ZenithParserNL) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}
			p.SetState(32)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case ZenithParserNL:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(33)
			p.Match(ZenithParserNL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(37)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(34)
					_la = p.GetTokenStream().LA(1)

					if !(_la == ZenithParserSEMICOLON || _la == ZenithParserNL) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}

			}
			p.SetState(39)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 4, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}

	case ZenithParserEOF:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(40)
			p.Match(ZenithParserEOF)
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

// IStatContext is an interface to support dynamic dispatch.
type IStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsStatContext differentiates from other interfaces.
	IsStatContext()
}

type StatContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyStatContext() *StatContext {
	var p = new(StatContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_stat
	return p
}

func InitEmptyStatContext(p *StatContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_stat
}

func (*StatContext) IsStatContext() {}

func NewStatContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StatContext {
	var p = new(StatContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_stat

	return p
}

func (s *StatContext) GetParser() antlr.Parser { return s.parser }

func (s *StatContext) CopyAll(ctx *StatContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *StatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StatContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type MultiStatContext struct {
	StatContext
}

func NewMultiStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MultiStatContext {
	var p = new(MultiStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *MultiStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiStatContext) LBRACE() antlr.TerminalNode {
	return s.GetToken(ZenithParserLBRACE, 0)
}

func (s *MultiStatContext) RBRACE() antlr.TerminalNode {
	return s.GetToken(ZenithParserRBRACE, 0)
}

func (s *MultiStatContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *MultiStatContext) AllEndedStat() []IEndedStatContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IEndedStatContext); ok {
			len++
		}
	}

	tst := make([]IEndedStatContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IEndedStatContext); ok {
			tst[i] = t.(IEndedStatContext)
			i++
		}
	}

	return tst
}

func (s *MultiStatContext) EndedStat(i int) IEndedStatContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IEndedStatContext); ok {
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

	return t.(IEndedStatContext)
}

func (s *MultiStatContext) Stat() IStatContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IStatContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IStatContext)
}

func (s *MultiStatContext) LineEnd() ILineEndContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILineEndContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILineEndContext)
}

func (s *MultiStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterMultiStat(s)
	}
}

func (s *MultiStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitMultiStat(s)
	}
}

func (s *MultiStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitMultiStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type ExprStatContext struct {
	StatContext
}

func NewExprStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ExprStatContext {
	var p = new(ExprStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *ExprStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExprStatContext) Expr() IExprContext {
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

func (s *ExprStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterExprStat(s)
	}
}

func (s *ExprStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitExprStat(s)
	}
}

func (s *ExprStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitExprStat(s)

	default:
		return t.VisitChildren(s)
	}
}

type DefineStatContext struct {
	StatContext
}

func NewDefineStatContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *DefineStatContext {
	var p = new(DefineStatContext)

	InitEmptyStatContext(&p.StatContext)
	p.parser = parser
	p.CopyAll(ctx.(*StatContext))

	return p
}

func (s *DefineStatContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefineStatContext) ID() antlr.TerminalNode {
	return s.GetToken(ZenithParserID, 0)
}

func (s *DefineStatContext) INIT_ASSIGN() antlr.TerminalNode {
	return s.GetToken(ZenithParserINIT_ASSIGN, 0)
}

func (s *DefineStatContext) Expr() IExprContext {
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

func (s *DefineStatContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *DefineStatContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *DefineStatContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterDefineStat(s)
	}
}

func (s *DefineStatContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitDefineStat(s)
	}
}

func (s *DefineStatContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitDefineStat(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) Stat() (localctx IStatContext) {
	localctx = NewStatContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ZenithParserRULE_stat)
	var _la int

	var _alt int

	p.SetState(70)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		localctx = NewDefineStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(43)
			p.Match(ZenithParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(45)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1048606) != 0 {
			{
				p.SetState(44)
				p.Type_()
			}

		}
		{
			p.SetState(47)
			p.Match(ZenithParserINIT_ASSIGN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(49)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(48)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(51)
			p.expr(0)
		}

	case 2:
		localctx = NewMultiStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(52)
			p.Match(ZenithParserLBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(54)
		p.GetErrorHandler().Sync(p)

		if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 8, p.GetParserRuleContext()) == 1 {
			{
				p.SetState(53)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		} else if p.HasError() { // JIM
			goto errorExit
		}
		p.SetState(59)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
		for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
			if _alt == 1 {
				{
					p.SetState(56)
					p.EndedStat()
				}

			}
			p.SetState(61)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 9, p.GetParserRuleContext())
			if p.HasError() {
				goto errorExit
			}
		}
		p.SetState(63)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&49513374) != 0 {
			{
				p.SetState(62)
				p.Stat()
			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la - -1)) & ^0x3f) == 0 && ((int64(1)<<(_la - -1))&805306369) != 0 {
			{
				p.SetState(65)
				p.LineEnd()
			}

		}
		{
			p.SetState(68)
			p.Match(ZenithParserRBRACE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case 3:
		localctx = NewExprStatContext(p, localctx)
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(69)
			p.expr(0)
		}

	case antlr.ATNInvalidAltNumber:
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

// ITypeContext is an interface to support dynamic dispatch.
type ITypeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser
	// IsTypeContext differentiates from other interfaces.
	IsTypeContext()
}

type TypeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTypeContext() *TypeContext {
	var p = new(TypeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_type
	return p
}

func InitEmptyTypeContext(p *TypeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_type
}

func (*TypeContext) IsTypeContext() {}

func NewTypeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypeContext {
	var p = new(TypeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_type

	return p
}

func (s *TypeContext) GetParser() antlr.Parser { return s.parser }

func (s *TypeContext) CopyAll(ctx *TypeContext) {
	s.CopyFrom(&ctx.BaseParserRuleContext)
}

func (s *TypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

type PtrTypeContext struct {
	TypeContext
	Ptr antlr.Token
}

func NewPtrTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PtrTypeContext {
	var p = new(PtrTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *PtrTypeContext) GetPtr() antlr.Token { return s.Ptr }

func (s *PtrTypeContext) SetPtr(v antlr.Token) { s.Ptr = v }

func (s *PtrTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PtrTypeContext) Type_() ITypeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITypeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITypeContext)
}

func (s *PtrTypeContext) OWN_PTR() antlr.TerminalNode {
	return s.GetToken(ZenithParserOWN_PTR, 0)
}

func (s *PtrTypeContext) PTR() antlr.TerminalNode {
	return s.GetToken(ZenithParserPTR, 0)
}

func (s *PtrTypeContext) NULL_OWN_PTR() antlr.TerminalNode {
	return s.GetToken(ZenithParserNULL_OWN_PTR, 0)
}

func (s *PtrTypeContext) NULL_PTR() antlr.TerminalNode {
	return s.GetToken(ZenithParserNULL_PTR, 0)
}

func (s *PtrTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterPtrType(s)
	}
}

func (s *PtrTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitPtrType(s)
	}
}

func (s *PtrTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitPtrType(s)

	default:
		return t.VisitChildren(s)
	}
}

type BaseTypeContext struct {
	TypeContext
}

func NewBaseTypeContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *BaseTypeContext {
	var p = new(BaseTypeContext)

	InitEmptyTypeContext(&p.TypeContext)
	p.parser = parser
	p.CopyAll(ctx.(*TypeContext))

	return p
}

func (s *BaseTypeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseTypeContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ZenithParserTYPE, 0)
}

func (s *BaseTypeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterBaseType(s)
	}
}

func (s *BaseTypeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitBaseType(s)
	}
}

func (s *BaseTypeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitBaseType(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) Type_() (localctx ITypeContext) {
	localctx = NewTypeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ZenithParserRULE_type)
	var _la int

	p.SetState(75)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ZenithParserTYPE:
		localctx = NewBaseTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(72)
			p.Match(ZenithParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserOWN_PTR, ZenithParserNULL_OWN_PTR, ZenithParserPTR, ZenithParserNULL_PTR:
		localctx = NewPtrTypeContext(p, localctx)
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(73)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*PtrTypeContext).Ptr = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*PtrTypeContext).Ptr = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(74)
			p.Type_()
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
	p.RuleIndex = ZenithParserRULE_expr
	return p
}

func InitEmptyExprContext(p *ExprContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = ZenithParserRULE_expr
}

func (*ExprContext) IsExprContext() {}

func NewExprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExprContext {
	var p = new(ExprContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = ZenithParserRULE_expr

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

type PtrExprContext struct {
	ExprContext
	Op    antlr.Token
	Right IExprContext
}

func NewPtrExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PtrExprContext {
	var p = new(PtrExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PtrExprContext) GetOp() antlr.Token { return s.Op }

func (s *PtrExprContext) SetOp(v antlr.Token) { s.Op = v }

func (s *PtrExprContext) GetRight() IExprContext { return s.Right }

func (s *PtrExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *PtrExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PtrExprContext) Expr() IExprContext {
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

func (s *PtrExprContext) OWN_PTR() antlr.TerminalNode {
	return s.GetToken(ZenithParserOWN_PTR, 0)
}

func (s *PtrExprContext) PTR() antlr.TerminalNode {
	return s.GetToken(ZenithParserPTR, 0)
}

func (s *PtrExprContext) NULL_OWN_PTR() antlr.TerminalNode {
	return s.GetToken(ZenithParserNULL_OWN_PTR, 0)
}

func (s *PtrExprContext) NULL_PTR() antlr.TerminalNode {
	return s.GetToken(ZenithParserNULL_PTR, 0)
}

func (s *PtrExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterPtrExpr(s)
	}
}

func (s *PtrExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitPtrExpr(s)
	}
}

func (s *PtrExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitPtrExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type PrefixExprContext struct {
	ExprContext
	Op    antlr.Token
	Right IExprContext
}

func NewPrefixExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *PrefixExprContext {
	var p = new(PrefixExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *PrefixExprContext) GetOp() antlr.Token { return s.Op }

func (s *PrefixExprContext) SetOp(v antlr.Token) { s.Op = v }

func (s *PrefixExprContext) GetRight() IExprContext { return s.Right }

func (s *PrefixExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *PrefixExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixExprContext) Expr() IExprContext {
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

func (s *PrefixExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ZenithParserPLUS, 0)
}

func (s *PrefixExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ZenithParserMINUS, 0)
}

func (s *PrefixExprContext) NOT() antlr.TerminalNode {
	return s.GetToken(ZenithParserNOT, 0)
}

func (s *PrefixExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterPrefixExpr(s)
	}
}

func (s *PrefixExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitPrefixExpr(s)
	}
}

func (s *PrefixExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitPrefixExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type CastExprContext struct {
	ExprContext
}

func NewCastExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastExprContext {
	var p = new(CastExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CastExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastExprContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ZenithParserTYPE, 0)
}

func (s *CastExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserLPAREN, 0)
}

func (s *CastExprContext) Expr() IExprContext {
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

func (s *CastExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserRPAREN, 0)
}

func (s *CastExprContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserNL)
}

func (s *CastExprContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, i)
}

func (s *CastExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterCastExpr(s)
	}
}

func (s *CastExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitCastExpr(s)
	}
}

func (s *CastExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitCastExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type AddExprContext struct {
	ExprContext
	Left  IExprContext
	Op    antlr.Token
	Right IExprContext
}

func NewAddExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *AddExprContext {
	var p = new(AddExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *AddExprContext) GetOp() antlr.Token { return s.Op }

func (s *AddExprContext) SetOp(v antlr.Token) { s.Op = v }

func (s *AddExprContext) GetLeft() IExprContext { return s.Left }

func (s *AddExprContext) GetRight() IExprContext { return s.Right }

func (s *AddExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *AddExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *AddExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AddExprContext) AllExpr() []IExprContext {
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

func (s *AddExprContext) Expr(i int) IExprContext {
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

func (s *AddExprContext) PLUS() antlr.TerminalNode {
	return s.GetToken(ZenithParserPLUS, 0)
}

func (s *AddExprContext) MINUS() antlr.TerminalNode {
	return s.GetToken(ZenithParserMINUS, 0)
}

func (s *AddExprContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *AddExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterAddExpr(s)
	}
}

func (s *AddExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitAddExpr(s)
	}
}

func (s *AddExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitAddExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type MulExprContext struct {
	ExprContext
	Left  IExprContext
	Op    antlr.Token
	Right IExprContext
}

func NewMulExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *MulExprContext {
	var p = new(MulExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *MulExprContext) GetOp() antlr.Token { return s.Op }

func (s *MulExprContext) SetOp(v antlr.Token) { s.Op = v }

func (s *MulExprContext) GetLeft() IExprContext { return s.Left }

func (s *MulExprContext) GetRight() IExprContext { return s.Right }

func (s *MulExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *MulExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *MulExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MulExprContext) AllExpr() []IExprContext {
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

func (s *MulExprContext) Expr(i int) IExprContext {
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

func (s *MulExprContext) TIMES() antlr.TerminalNode {
	return s.GetToken(ZenithParserTIMES, 0)
}

func (s *MulExprContext) DIVIDE() antlr.TerminalNode {
	return s.GetToken(ZenithParserDIVIDE, 0)
}

func (s *MulExprContext) REM() antlr.TerminalNode {
	return s.GetToken(ZenithParserREM, 0)
}

func (s *MulExprContext) NL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, 0)
}

func (s *MulExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterMulExpr(s)
	}
}

func (s *MulExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitMulExpr(s)
	}
}

func (s *MulExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitMulExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IfExprContext struct {
	ExprContext
	Left      IExprContext
	Condition IExprContext
	Right     IExprContext
}

func NewIfExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IfExprContext {
	var p = new(IfExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IfExprContext) GetLeft() IExprContext { return s.Left }

func (s *IfExprContext) GetCondition() IExprContext { return s.Condition }

func (s *IfExprContext) GetRight() IExprContext { return s.Right }

func (s *IfExprContext) SetLeft(v IExprContext) { s.Left = v }

func (s *IfExprContext) SetCondition(v IExprContext) { s.Condition = v }

func (s *IfExprContext) SetRight(v IExprContext) { s.Right = v }

func (s *IfExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfExprContext) IF() antlr.TerminalNode {
	return s.GetToken(ZenithParserIF, 0)
}

func (s *IfExprContext) ELSE() antlr.TerminalNode {
	return s.GetToken(ZenithParserELSE, 0)
}

func (s *IfExprContext) AllExpr() []IExprContext {
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

func (s *IfExprContext) Expr(i int) IExprContext {
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

func (s *IfExprContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserNL)
}

func (s *IfExprContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, i)
}

func (s *IfExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterIfExpr(s)
	}
}

func (s *IfExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitIfExpr(s)
	}
}

func (s *IfExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitIfExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type KeyExprContext struct {
	ExprContext
	Key antlr.Token
}

func NewKeyExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *KeyExprContext {
	var p = new(KeyExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *KeyExprContext) GetKey() antlr.Token { return s.Key }

func (s *KeyExprContext) SetKey(v antlr.Token) { s.Key = v }

func (s *KeyExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *KeyExprContext) TRUE() antlr.TerminalNode {
	return s.GetToken(ZenithParserTRUE, 0)
}

func (s *KeyExprContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ZenithParserFALSE, 0)
}

func (s *KeyExprContext) NULL() antlr.TerminalNode {
	return s.GetToken(ZenithParserNULL, 0)
}

func (s *KeyExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterKeyExpr(s)
	}
}

func (s *KeyExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitKeyExpr(s)
	}
}

func (s *KeyExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitKeyExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type NumExprContext struct {
	ExprContext
}

func NewNumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumExprContext {
	var p = new(NumExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NumExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumExprContext) NUM() antlr.TerminalNode {
	return s.GetToken(ZenithParserNUM, 0)
}

func (s *NumExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterNumExpr(s)
	}
}

func (s *NumExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitNumExpr(s)
	}
}

func (s *NumExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitNumExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type ParenExprContext struct {
	ExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserLPAREN, 0)
}

func (s *ParenExprContext) Expr() IExprContext {
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

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserRPAREN, 0)
}

func (s *ParenExprContext) AllNL() []antlr.TerminalNode {
	return s.GetTokens(ZenithParserNL)
}

func (s *ParenExprContext) NL(i int) antlr.TerminalNode {
	return s.GetToken(ZenithParserNL, i)
}

func (s *ParenExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterParenExpr(s)
	}
}

func (s *ParenExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitParenExpr(s)
	}
}

func (s *ParenExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitParenExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

type IdExprContext struct {
	ExprContext
}

func NewIdExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *IdExprContext {
	var p = new(IdExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *IdExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IdExprContext) ID() antlr.TerminalNode {
	return s.GetToken(ZenithParserID, 0)
}

func (s *IdExprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.EnterIdExpr(s)
	}
}

func (s *IdExprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ZenithParserListener); ok {
		listenerT.ExitIdExpr(s)
	}
}

func (s *IdExprContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case ZenithParserVisitor:
		return t.VisitIdExpr(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *ZenithParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ZenithParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 10
	p.EnterRecursionRule(localctx, 10, ZenithParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ZenithParserNUM:
		localctx = NewNumExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(78)
			p.Match(ZenithParserNUM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserID:
		localctx = NewIdExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(79)
			p.Match(ZenithParserID)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserTRUE, ZenithParserFALSE, ZenithParserNULL:
		localctx = NewKeyExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(80)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*KeyExprContext).Key = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&229376) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*KeyExprContext).Key = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	case ZenithParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(81)
			p.Match(ZenithParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(83)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(82)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(85)
			p.expr(0)
		}
		p.SetState(87)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(86)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(89)
			p.Match(ZenithParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserTYPE:
		localctx = NewCastExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(91)
			p.Match(ZenithParserTYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(92)
			p.Match(ZenithParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(94)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(93)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(96)
			p.expr(0)
		}
		p.SetState(98)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(97)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(100)
			p.Match(ZenithParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserOWN_PTR, ZenithParserNULL_OWN_PTR, ZenithParserPTR, ZenithParserNULL_PTR:
		localctx = NewPtrExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(102)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*PtrExprContext).Op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&30) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*PtrExprContext).Op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(103)

			var _x = p.expr(5)

			localctx.(*PtrExprContext).Right = _x
		}

	case ZenithParserNOT, ZenithParserPLUS, ZenithParserMINUS:
		localctx = NewPrefixExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(104)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*PrefixExprContext).Op = _lt

			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&896) != 0) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*PrefixExprContext).Op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(105)

			var _x = p.expr(4)

			localctx.(*PrefixExprContext).Right = _x
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(137)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(135)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 24, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MulExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(108)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(109)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulExprContext).Op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&7168) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulExprContext).Op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(111)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(110)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(113)

					var _x = p.expr(4)

					localctx.(*MulExprContext).Right = _x
				}

			case 2:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(114)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(115)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*AddExprContext).Op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ZenithParserPLUS || _la == ZenithParserMINUS) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*AddExprContext).Op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(117)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(116)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(119)

					var _x = p.expr(3)

					localctx.(*AddExprContext).Right = _x
				}

			case 3:
				localctx = NewIfExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*IfExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(120)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
					goto errorExit
				}
				{
					p.SetState(121)
					p.Match(ZenithParserIF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(123)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(122)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(125)

					var _x = p.expr(0)

					localctx.(*IfExprContext).Condition = _x
				}
				p.SetState(127)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(126)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(129)
					p.Match(ZenithParserELSE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(131)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(130)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(133)

					var _x = p.expr(2)

					localctx.(*IfExprContext).Right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(139)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 25, p.GetParserRuleContext())
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

func (p *ZenithParser) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 5:
		var t *ExprContext = nil
		if localctx != nil {
			t = localctx.(*ExprContext)
		}
		return p.Expr_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *ZenithParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
