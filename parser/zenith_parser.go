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
		"", "'+'", "'-'", "'*'", "'/'", "'%'", "'if'", "'else'", "'true'", "'false'",
		"", "", "'('", "')'",
	}
	staticData.SymbolicNames = []string{
		"", "PLUS", "MINUS", "TIMES", "DIVIDE", "REM", "IF", "ELSE", "TRUE",
		"FALSE", "TYPE", "NUM", "LPAREN", "RPAREN", "NL", "SPACE",
	}
	staticData.RuleNames = []string{
		"fileStat", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 15, 68, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1,
		1, 3, 1, 11, 8, 1, 1, 1, 1, 1, 3, 1, 15, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 3, 1, 22, 8, 1, 1, 1, 1, 1, 3, 1, 26, 8, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 3, 1, 34, 8, 1, 1, 1, 1, 1, 1, 1, 3, 1, 39, 8, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 45, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 51, 8,
		1, 1, 1, 1, 1, 3, 1, 55, 8, 1, 1, 1, 1, 1, 3, 1, 59, 8, 1, 1, 1, 1, 1,
		5, 1, 63, 8, 1, 10, 1, 12, 1, 66, 9, 1, 1, 1, 0, 1, 2, 2, 0, 2, 0, 3, 1,
		0, 1, 2, 1, 0, 8, 9, 1, 0, 3, 5, 81, 0, 4, 1, 0, 0, 0, 2, 33, 1, 0, 0,
		0, 4, 5, 3, 2, 1, 0, 5, 6, 5, 0, 0, 1, 6, 1, 1, 0, 0, 0, 7, 8, 6, 1, -1,
		0, 8, 10, 5, 12, 0, 0, 9, 11, 5, 14, 0, 0, 10, 9, 1, 0, 0, 0, 10, 11, 1,
		0, 0, 0, 11, 12, 1, 0, 0, 0, 12, 14, 3, 2, 1, 0, 13, 15, 5, 14, 0, 0, 14,
		13, 1, 0, 0, 0, 14, 15, 1, 0, 0, 0, 15, 16, 1, 0, 0, 0, 16, 17, 5, 13,
		0, 0, 17, 34, 1, 0, 0, 0, 18, 19, 5, 10, 0, 0, 19, 21, 5, 12, 0, 0, 20,
		22, 5, 14, 0, 0, 21, 20, 1, 0, 0, 0, 21, 22, 1, 0, 0, 0, 22, 23, 1, 0,
		0, 0, 23, 25, 3, 2, 1, 0, 24, 26, 5, 14, 0, 0, 25, 24, 1, 0, 0, 0, 25,
		26, 1, 0, 0, 0, 26, 27, 1, 0, 0, 0, 27, 28, 5, 13, 0, 0, 28, 34, 1, 0,
		0, 0, 29, 30, 7, 0, 0, 0, 30, 34, 3, 2, 1, 6, 31, 34, 5, 11, 0, 0, 32,
		34, 7, 1, 0, 0, 33, 7, 1, 0, 0, 0, 33, 18, 1, 0, 0, 0, 33, 29, 1, 0, 0,
		0, 33, 31, 1, 0, 0, 0, 33, 32, 1, 0, 0, 0, 34, 64, 1, 0, 0, 0, 35, 36,
		10, 5, 0, 0, 36, 38, 7, 2, 0, 0, 37, 39, 5, 14, 0, 0, 38, 37, 1, 0, 0,
		0, 38, 39, 1, 0, 0, 0, 39, 40, 1, 0, 0, 0, 40, 63, 3, 2, 1, 6, 41, 42,
		10, 4, 0, 0, 42, 44, 7, 0, 0, 0, 43, 45, 5, 14, 0, 0, 44, 43, 1, 0, 0,
		0, 44, 45, 1, 0, 0, 0, 45, 46, 1, 0, 0, 0, 46, 63, 3, 2, 1, 5, 47, 48,
		10, 3, 0, 0, 48, 50, 5, 6, 0, 0, 49, 51, 5, 14, 0, 0, 50, 49, 1, 0, 0,
		0, 50, 51, 1, 0, 0, 0, 51, 52, 1, 0, 0, 0, 52, 54, 3, 2, 1, 0, 53, 55,
		5, 14, 0, 0, 54, 53, 1, 0, 0, 0, 54, 55, 1, 0, 0, 0, 55, 56, 1, 0, 0, 0,
		56, 58, 5, 7, 0, 0, 57, 59, 5, 14, 0, 0, 58, 57, 1, 0, 0, 0, 58, 59, 1,
		0, 0, 0, 59, 60, 1, 0, 0, 0, 60, 61, 3, 2, 1, 4, 61, 63, 1, 0, 0, 0, 62,
		35, 1, 0, 0, 0, 62, 41, 1, 0, 0, 0, 62, 47, 1, 0, 0, 0, 63, 66, 1, 0, 0,
		0, 64, 62, 1, 0, 0, 0, 64, 65, 1, 0, 0, 0, 65, 3, 1, 0, 0, 0, 66, 64, 1,
		0, 0, 0, 12, 10, 14, 21, 25, 33, 38, 44, 50, 54, 58, 62, 64,
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
	ZenithParserEOF    = antlr.TokenEOF
	ZenithParserPLUS   = 1
	ZenithParserMINUS  = 2
	ZenithParserTIMES  = 3
	ZenithParserDIVIDE = 4
	ZenithParserREM    = 5
	ZenithParserIF     = 6
	ZenithParserELSE   = 7
	ZenithParserTRUE   = 8
	ZenithParserFALSE  = 9
	ZenithParserTYPE   = 10
	ZenithParserNUM    = 11
	ZenithParserLPAREN = 12
	ZenithParserRPAREN = 13
	ZenithParserNL     = 14
	ZenithParserSPACE  = 15
)

// ZenithParser rules.
const (
	ZenithParserRULE_fileStat = 0
	ZenithParserRULE_expr     = 1
)

// IFileStatContext is an interface to support dynamic dispatch.
type IFileStatContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expr() IExprContext
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

func (s *FileStatContext) Expr() IExprContext {
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
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(4)
		p.expr(0)
	}
	{
		p.SetState(5)
		p.Match(ZenithParserEOF)
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
	Type  antlr.Token
	Inner IExprContext
}

func NewCastExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *CastExprContext {
	var p = new(CastExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *CastExprContext) GetType() antlr.Token { return s.Type }

func (s *CastExprContext) SetType(v antlr.Token) { s.Type = v }

func (s *CastExprContext) GetInner() IExprContext { return s.Inner }

func (s *CastExprContext) SetInner(v IExprContext) { s.Inner = v }

func (s *CastExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CastExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserLPAREN, 0)
}

func (s *CastExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserRPAREN, 0)
}

func (s *CastExprContext) TYPE() antlr.TerminalNode {
	return s.GetToken(ZenithParserTYPE, 0)
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

type ParenExprContext struct {
	ExprContext
	Inner IExprContext
}

func NewParenExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *ParenExprContext {
	var p = new(ParenExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *ParenExprContext) GetInner() IExprContext { return s.Inner }

func (s *ParenExprContext) SetInner(v IExprContext) { s.Inner = v }

func (s *ParenExprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ParenExprContext) LPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserLPAREN, 0)
}

func (s *ParenExprContext) RPAREN() antlr.TerminalNode {
	return s.GetToken(ZenithParserRPAREN, 0)
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

type NumExprContext struct {
	ExprContext
	Num antlr.Token
}

func NewNumExprContext(parser antlr.Parser, ctx antlr.ParserRuleContext) *NumExprContext {
	var p = new(NumExprContext)

	InitEmptyExprContext(&p.ExprContext)
	p.parser = parser
	p.CopyAll(ctx.(*ExprContext))

	return p
}

func (s *NumExprContext) GetNum() antlr.Token { return s.Num }

func (s *NumExprContext) SetNum(v antlr.Token) { s.Num = v }

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

func (p *ZenithParser) Expr() (localctx IExprContext) {
	return p.expr(0)
}

func (p *ZenithParser) expr(_p int) (localctx IExprContext) {
	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()

	_parentState := p.GetState()
	localctx = NewExprContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExprContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 2
	p.EnterRecursionRule(localctx, 2, ZenithParserRULE_expr, _p)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(33)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case ZenithParserLPAREN:
		localctx = NewParenExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx

		{
			p.SetState(8)
			p.Match(ZenithParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(10)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(9)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(12)

			var _x = p.expr(0)

			localctx.(*ParenExprContext).Inner = _x
		}
		p.SetState(14)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(13)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(16)
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
			p.SetState(18)

			var _m = p.Match(ZenithParserTYPE)

			localctx.(*CastExprContext).Type = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(19)
			p.Match(ZenithParserLPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
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

			var _x = p.expr(0)

			localctx.(*CastExprContext).Inner = _x
		}
		p.SetState(25)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == ZenithParserNL {
			{
				p.SetState(24)
				p.Match(ZenithParserNL)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(27)
			p.Match(ZenithParserRPAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserPLUS, ZenithParserMINUS:
		localctx = NewPrefixExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(29)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*PrefixExprContext).Op = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ZenithParserPLUS || _la == ZenithParserMINUS) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*PrefixExprContext).Op = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(30)

			var _x = p.expr(6)

			localctx.(*PrefixExprContext).Right = _x
		}

	case ZenithParserNUM:
		localctx = NewNumExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(31)

			var _m = p.Match(ZenithParserNUM)

			localctx.(*NumExprContext).Num = _m
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case ZenithParserTRUE, ZenithParserFALSE:
		localctx = NewKeyExprContext(p, localctx)
		p.SetParserRuleContext(localctx)
		_prevctx = localctx
		{
			p.SetState(32)

			var _lt = p.GetTokenStream().LT(1)

			localctx.(*KeyExprContext).Key = _lt

			_la = p.GetTokenStream().LA(1)

			if !(_la == ZenithParserTRUE || _la == ZenithParserFALSE) {
				var _ri = p.GetErrorHandler().RecoverInline(p)

				localctx.(*KeyExprContext).Key = _ri
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(64)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(62)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 10, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MulExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(35)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
					goto errorExit
				}
				{
					p.SetState(36)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulExprContext).Op = _lt

					_la = p.GetTokenStream().LA(1)

					if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&56) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulExprContext).Op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(38)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(37)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(40)

					var _x = p.expr(6)

					localctx.(*MulExprContext).Right = _x
				}

			case 2:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(41)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
					goto errorExit
				}
				{
					p.SetState(42)

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
				p.SetState(44)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(43)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(46)

					var _x = p.expr(5)

					localctx.(*AddExprContext).Right = _x
				}

			case 3:
				localctx = NewIfExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*IfExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(47)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(48)
					p.Match(ZenithParserIF)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(50)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(49)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(52)

					var _x = p.expr(0)

					localctx.(*IfExprContext).Condition = _x
				}
				p.SetState(54)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(53)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(56)
					p.Match(ZenithParserELSE)
					if p.HasError() {
						// Recognition error - abort rule
						goto errorExit
					}
				}
				p.SetState(58)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(57)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(60)

					var _x = p.expr(4)

					localctx.(*IfExprContext).Right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(66)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 11, p.GetParserRuleContext())
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

func (p *ZenithParser) Expr_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
