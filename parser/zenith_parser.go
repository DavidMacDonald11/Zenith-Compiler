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
		"", "'+'", "'-'", "'*'", "'/'",
	}
	staticData.SymbolicNames = []string{
		"", "PLUS", "MINUS", "TIMES", "DIVIDE", "NUM", "NL", "SPACE",
	}
	staticData.RuleNames = []string{
		"fileStat", "expr",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 7, 28, 2, 0, 7, 0, 2, 1, 7, 1, 1, 0, 1, 0, 1, 0, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 3, 1, 14, 8, 1, 1, 1, 1, 1, 1, 1, 1, 1, 3, 1, 20, 8,
		1, 1, 1, 5, 1, 23, 8, 1, 10, 1, 12, 1, 26, 9, 1, 1, 1, 0, 1, 2, 2, 0, 2,
		0, 2, 1, 0, 3, 4, 1, 0, 1, 2, 29, 0, 4, 1, 0, 0, 0, 2, 7, 1, 0, 0, 0, 4,
		5, 3, 2, 1, 0, 5, 6, 5, 0, 0, 1, 6, 1, 1, 0, 0, 0, 7, 8, 6, 1, -1, 0, 8,
		9, 5, 5, 0, 0, 9, 24, 1, 0, 0, 0, 10, 11, 10, 3, 0, 0, 11, 13, 7, 0, 0,
		0, 12, 14, 5, 6, 0, 0, 13, 12, 1, 0, 0, 0, 13, 14, 1, 0, 0, 0, 14, 15,
		1, 0, 0, 0, 15, 23, 3, 2, 1, 4, 16, 17, 10, 2, 0, 0, 17, 19, 7, 1, 0, 0,
		18, 20, 5, 6, 0, 0, 19, 18, 1, 0, 0, 0, 19, 20, 1, 0, 0, 0, 20, 21, 1,
		0, 0, 0, 21, 23, 3, 2, 1, 3, 22, 10, 1, 0, 0, 0, 22, 16, 1, 0, 0, 0, 23,
		26, 1, 0, 0, 0, 24, 22, 1, 0, 0, 0, 24, 25, 1, 0, 0, 0, 25, 3, 1, 0, 0,
		0, 26, 24, 1, 0, 0, 0, 4, 13, 19, 22, 24,
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
	ZenithParserNUM    = 5
	ZenithParserNL     = 6
	ZenithParserSPACE  = 7
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
	localctx = NewNumExprContext(p, localctx)
	p.SetParserRuleContext(localctx)
	_prevctx = localctx

	{
		p.SetState(8)

		var _m = p.Match(ZenithParserNUM)

		localctx.(*NumExprContext).Num = _m
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(24)
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
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(22)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 2, p.GetParserRuleContext()) {
			case 1:
				localctx = NewMulExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*MulExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(10)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
					goto errorExit
				}
				{
					p.SetState(11)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*MulExprContext).Op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ZenithParserTIMES || _la == ZenithParserDIVIDE) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*MulExprContext).Op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				p.SetState(13)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(12)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(15)

					var _x = p.expr(4)

					localctx.(*MulExprContext).Right = _x
				}

			case 2:
				localctx = NewAddExprContext(p, NewExprContext(p, _parentctx, _parentState))
				localctx.(*AddExprContext).Left = _prevctx

				p.PushNewRecursionContext(localctx, _startState, ZenithParserRULE_expr)
				p.SetState(16)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					p.SetError(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
					goto errorExit
				}
				{
					p.SetState(17)

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
				p.SetState(19)
				p.GetErrorHandler().Sync(p)
				if p.HasError() {
					goto errorExit
				}
				_la = p.GetTokenStream().LA(1)

				if _la == ZenithParserNL {
					{
						p.SetState(18)
						p.Match(ZenithParserNL)
						if p.HasError() {
							// Recognition error - abort rule
							goto errorExit
						}
					}

				}
				{
					p.SetState(21)

					var _x = p.expr(3)

					localctx.(*AddExprContext).Right = _x
				}

			case antlr.ATNInvalidAltNumber:
				goto errorExit
			}

		}
		p.SetState(26)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 3, p.GetParserRuleContext())
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
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
